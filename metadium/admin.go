/* admin.go */

package metadium

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"path"
	"reflect"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/embed"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	metaapi "github.com/ethereum/go-ethereum/metadium/api"
	"github.com/ethereum/go-ethereum/metadium/metclient"
	metaminer "github.com/ethereum/go-ethereum/metadium/miner"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discv5"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

type metaNode struct {
	Partner   bool   `json:"partner"`
	Name      string `json:"name"`
	Id        string `json:"id"`
	V4Id      string `json:"v4id"`
	Ip        string `json:"ip"`
	Port      int    `json:"port"`
	NotBefore int    `json:"notBefore"`
	NotAfter  int    `json:"notAfter"`
	Prev      string `json:"prev"`
	Next      string `json:"next"`

	Status string `json:"status"`
	Miner  bool   `json:"miner"`
}

type metaMember struct {
	Addr  common.Address `json:"address"`
	Stake int            `json:"stake"`
	Prev  common.Address `json:"prev"`
	Next  common.Address `json:"next"`
}

type metaAdmin struct {
	stack *node.Node

	bootNodeId  string // allowed to generate block without admin contract
	bootAccount common.Address
	nodeInfo    *p2p.NodeInfo
	anchor      common.Address
	admin       common.Address
	Updates     chan bool
	rpcCli      *rpc.Client
	cli         *ethclient.Client

	etcd        *embed.Etcd
	etcdCli     *clientv3.Client
	etcdDir     string
	etcdPort    int
	etcdTimeout time.Duration

	adminAbi  abi.ABI
	anchorAbi abi.ABI

	lastBlock     int
	modifiedBlock int
	blocksPer     int
	self          *metaNode

	lock  *sync.Mutex
	nodes map[string]*metaNode

	// # of blocks consecutively mined by this node
	blocksMined int
}

var (
	magic           = 0x4d6574616469756d // Metadium
	etcdClusterName = "Metadium"
	big0            = big.NewInt(0)
	nilAddress      = common.Address{}
	admin           *metaAdmin

	ErrNotRunning     = errors.New("Not Running")
	ErrAlreadyRunning = errors.New("Already Running")
)

func (n *metaNode) eq(m *metaNode) bool {
	if n.Partner == m.Partner && n.Name == m.Name && n.Id == m.Id &&
		n.Ip == m.Ip && n.Port == m.Port && n.NotBefore == m.NotBefore &&
		n.NotAfter == m.NotAfter {
		return true
	} else {
		return false
	}
}

// convert v5 id to v4 id
func toV4Id(id string) (string, error) {
	if len(id) == 64 {
		return id, nil
	} else if len(id) == 128 {
		v4id, err := enode.ParseV4(fmt.Sprintf("enode://%v@127.0.0.1:8589", id))
		if err != nil {
			return "", err
		} else {
			return v4id.ID().String(), nil
		}
	} else {
		return "", fmt.Errorf("Invalid V5 Identifier")
	}
}

// returns
// 1) extradata of genesis block, which is the id of the node that is allowed
//   to generated blocks before admin contract is established.
// 2) returns the coinbase of genesis block, which should be the admin
//   contract creator
func (ma *metaAdmin) getGenesisInfo() (string, common.Address, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	block, err := ma.cli.HeaderByNumber(ctx, big0)
	if err != nil {
		return "", common.Address{}, err
	}

	nodeId := hex.EncodeToString(block.Extra)
	if len(nodeId) < 128 {
		panic("Invalid bootnode id in the genesis block.")
	}
	nodeId, err = toV4Id(nodeId[len(nodeId)-128:])
	if err != nil {
		panic("Invalid bootnode id in the genesis block.")
	}
	return nodeId, block.Coinbase, nil
}

// it should be the first transaction of the coinbase of the genesis block
func (ma *metaAdmin) getAdminAnchorAddress() common.Address {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := uint64(0); i < 10; i++ {
		anchorAddress := crypto.CreateAddress(ma.bootAccount, i)
		anchorContract := &metclient.RemoteContract{
			Cli: ma.cli,
			To:  &anchorAddress,
			Abi: ma.anchorAbi,
		}

		var v *big.Int
		err := metclient.CallContract(ctx, anchorContract, "magic", nil, &v, nil)
		if err == nil && int(v.Int64()) == magic {
			return anchorAddress
		}
	}

	return nilAddress
}

func (ma *metaAdmin) getAdminAddress() common.Address {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	anchorContract := &metclient.RemoteContract{
		Cli: ma.cli,
		To:  &ma.anchor,
		Abi: ma.anchorAbi,
	}

	var adminAddress common.Address
	err := metclient.CallContract(ctx, anchorContract, "admin", nil, &adminAddress, nil)
	if err != nil {
		log.Debug("Metadium", "Anchored admin address is invalid", err)
		ma.anchor = nilAddress
		return nilAddress
	}

	return adminAddress
}

func (ma *metaAdmin) getInt(ctx context.Context, to *common.Address, block *big.Int, name string) (int, error) {
	var v *big.Int
	err := ma.callContract(ctx, to, name, nil, &v, block)
	if err != nil {
		return 0, err
	} else {
		return int(v.Int64()), nil
	}
}

// returns []*metaNode from map[string]*metaNode
func (ma *metaAdmin) getNodes() []*metaNode {
	ma.lock.Lock()
	defer ma.lock.Unlock()

	var nodes []*metaNode
	for _, i := range ma.nodes {
		nodes = append(nodes, i)
	}
	return nodes
}

// returns
// 1. currentMiner *metaNode: the current leader
// 2. nextMiner *metaNode: the most eligible miner for the given height,
//   which is up and running
// 3. nodes []*metaNode: copies of map[string]*metaNode, not references,
//   sorted by id, i.e. mining order
// 'locked' indicates whether ma.lock is held by the caller or not
func (ma *metaAdmin) getMinerNodes(height int, locked bool) (*metaNode, *metaNode, []*metaNode) {
	var nodes []*metaNode
	if !locked {
		ma.lock.Lock()
	}
	for _, i := range ma.nodes {
		n := new(metaNode)
		*n = *i
		nodes = append(nodes, n)
	}
	if !locked {
		ma.lock.Unlock()
	}
	if len(nodes) == 0 {
		return nil, nil, nodes
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Name < nodes[j].Name
	})

	for _, n := range nodes {
		if (ma.self != nil && n.Id == ma.self.Id) || ma.isPeerUp(n.Id) {
			n.Status = "up"
		} else {
			n.Status = "down"
		}
	}

	_, leaderNode := ma.etcdLeader(locked)
	var miner, next *metaNode
	ix := height / admin.blocksPer % len(nodes)
	i := ix
	for j := 0; j < len(nodes); j++ {
		if miner != nil && next != nil {
			break
		}
		n := nodes[i]
		if miner == nil && leaderNode != nil && n.Name == leaderNode.Name {
			miner = n
			miner.Miner = true
		}
		if next == nil && n.Status == "up" {
			next = n
		}
		i = (i + 1) % len(nodes)
	}

	return miner, next, nodes
}

func (ma *metaAdmin) getNode(ctx context.Context, id string, block *big.Int) (*metaNode, error) {
	var (
		present       bool
		jsonOut       string
		input, output []interface{}
		err           error
	)

	output = []interface{}{&present, &jsonOut}
	if len(id) == 0 {
		err = ma.callContract(ctx, &ma.admin, "firstNode", input, &output, block)
	} else {
		input = []interface{}{[]byte(id)}
		err = ma.callContract(ctx, &ma.admin, "getNode", input, &output, block)
	}
	if err != nil {
		return nil, err
	}

	// trim the json output
	if ix := strings.Index(jsonOut, "\000"); ix >= 0 {
		jsonOut = jsonOut[:ix]
	}

	n := new(metaNode)
	err = json.Unmarshal([]byte(jsonOut), n)
	if err == nil {
		n.V4Id, err = toV4Id(n.Id)
	}
	return n, err
}

func isArray(x interface{}) bool {
	if x == nil {
		return false
	}
	y := reflect.TypeOf(x)
	switch y.Kind() {
	case reflect.Slice:
		return true
	case reflect.Array:
		return true
	default:
		return false
	}
}

/* input:
 *   no arguments -> nil
 *   one argument -> arg
 *   array        -> []interface{ arg1, arg2 }
 * output:
 *   one argument -> &arg
 *   array        -> &[]interface{ &arg1, &arg2 }
 */
func (ma *metaAdmin) callContract(ctx context.Context, to *common.Address, name string, input, output interface{}, block *big.Int) error {
	contract := &metclient.RemoteContract{
		Cli: ma.cli,
		To:  to,
		Abi: ma.adminAbi,
	}

	return metclient.CallContract(ctx, contract, name, input, output, block)
}

func (ma *metaAdmin) getMembers(ctx context.Context, block *big.Int) ([]*metaMember, error) {
	var (
		members          []*metaMember
		present          bool
		addr, prev, next common.Address
		stake            *big.Int
		input, output    []interface{}
		err              error
	)

	output = []interface{}{&present, &addr, &stake, &prev, &next}
	err = ma.callContract(ctx, &ma.admin, "firstMember", input, &output, block)
	if err != nil {
		return nil, err
	}
	for present {
		members = append(members, &metaMember{
			Addr:  addr,
			Stake: int(stake.Int64()),
			Prev:  prev,
			Next:  next,
		})

		if next == nilAddress {
			break
		} else {
			input = []interface{}{next}
			output = []interface{}{&present, &addr, &stake, &prev, &next}
			err = ma.callContract(ctx, &ma.admin, "getMember", input, &output, block)
			if err != nil {
				return nil, err
			}
		}
	}

	return members, nil
}

func (ma *metaAdmin) getData(refresh bool) (blockNum, modifiedBlock, blocksPer int, nodes, addedNodes, updatedNodes, deletedNodes []*metaNode, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	block, err := ma.cli.HeaderByNumber(ctx, nil)
	if err != nil {
		return
	}
	blockNum = int(block.Number.Int64())
	if !refresh && blockNum <= ma.lastBlock {
		return
	}

	modifiedBlock, err = ma.getInt(ctx, &ma.admin, block.Number, "modifiedBlock")
	if err != nil {
		return
	}
	if !refresh && ma.modifiedBlock == modifiedBlock {
		return
	}

	blocksPer, err = ma.getInt(ctx, &ma.admin, block.Number, "blocksPer")
	if err != nil {
		return
	}

	var n *metaNode
	for {
		if n == nil {
			n, err = ma.getNode(ctx, "", block.Number)
		} else if len(n.Next) == 0 {
			break
		} else {
			n, err = ma.getNode(ctx, n.Next, block.Number)
		}

		if err != nil {
			return
		} else {
			nodes = append(nodes, n)
		}
	}

	oldNodes := ma.getNodes()
	sort.Slice(oldNodes, func(i, j int) bool {
		return oldNodes[i].Name < oldNodes[j].Name
	})
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Name < nodes[j].Name
	})

	i, j := 0, 0
	for {
		if i >= len(oldNodes) || j >= len(nodes) {
			break
		}
		v := strings.Compare(oldNodes[i].Name, nodes[j].Name)
		if v == 0 {
			if !oldNodes[i].eq(nodes[j]) {
				updatedNodes = append(updatedNodes, nodes[j])
			}
			i++
			j++
		} else if v < 0 {
			deletedNodes = append(deletedNodes, oldNodes[i])
			i++
		} else if v > 0 {
			addedNodes = append(addedNodes, nodes[j])
			j++
		}
	}

	if i < len(oldNodes) {
		for ; i < len(oldNodes); i++ {
			deletedNodes = append(deletedNodes, oldNodes[i])
		}
	}
	if j < len(nodes) {
		for ; j < len(nodes); j++ {
			addedNodes = append(addedNodes, nodes[j])
		}
	}

	return
}

func StartAdmin(stack *node.Node, datadir string) {
	if !(params.ConsensusMethod == params.ConsensusPoA ||
		params.ConsensusMethod == params.ConsensusETCD ||
		params.ConsensusMethod == params.ConsensusPBFT) {
		utils.Fatalf("Invalid Consensus Method: %d\n", params.ConsensusMethod)
	}

	metaminer.IsMinerFunc = IsMiner
	metaminer.AmPartnerFunc = AmPartner
	metaminer.IsPartnerFunc = IsPartner
	metaminer.LogBlockFunc = LogBlock
	metaminer.CalculateRewardsFunc = calculateRewards
	metaminer.VerifyRewardsFunc = verifyRewards
	metaminer.SignBlockFunc = signBlock
	metaminer.VerifyBlockSigFunc = verifyBlockSig
	metaminer.RequirePendingTxsFunc = requirePendingTxs
	metaapi.Info = Info
	metaapi.GetMiners = getMiners
	metaapi.GetMinerStatus = getMinerStatus
	metaapi.EtcdInit = EtcdInit
	metaapi.EtcdAddMember = EtcdAddMember
	metaapi.EtcdRemoveMember = EtcdRemoveMember
	metaapi.EtcdJoin = EtcdJoin
	metaapi.EtcdMoveLeader = EtcdMoveLeader

	rpcCli, err := stack.Attach()
	if err != nil {
		utils.Fatalf("Failed to attach to self: %v", err)
	}

	anchorContract, err := metclient.LoadJsonContract(strings.NewReader(AdminAnchorAbi))
	if err != nil {
		utils.Fatalf("Loading ABI failed: %v", err)
	}

	adminContract, err := metclient.LoadJsonContract(strings.NewReader(AdminAbi))
	if err != nil {
		utils.Fatalf("Loading ABI failed: %v", err)
	}

	admin = &metaAdmin{
		stack:       stack,
		lock:        &sync.Mutex{},
		anchor:      nilAddress,
		admin:       nilAddress,
		Updates:     make(chan bool, 10),
		rpcCli:      rpcCli,
		cli:         ethclient.NewClient(rpcCli),
		etcdDir:     path.Join(datadir, "etcd"),
		etcdTimeout: 30 * time.Second,
		anchorAbi:   anchorContract.Abi,
		adminAbi:    adminContract.Abi,
	}

	admin.bootNodeId, admin.bootAccount, err = admin.getGenesisInfo()
	if err != nil {
		utils.Fatalf("Cannot get contract address from genesis block: %v", err)
	}

	go admin.run()
	go func() {
		for {
			admin.updateMiner(false)
			time.Sleep(1 * time.Second)
		}
	}()
}

func (ma *metaAdmin) addPeer(node *metaNode) error {
	if node.V4Id == ma.nodeInfo.ID {
		return nil
	}

	var v *bool
	ctx, cancel := context.WithCancel(context.Background())
	id := fmt.Sprintf("enode://%s@%s:%d", node.Id, node.Ip, node.Port)
	// TODO: trusted peers need more work
	//e := ma.rpcCli.CallContext(ctx, &v, "admin_addTrustedPeer", id)
	e := ma.rpcCli.CallContext(ctx, &v, "admin_addPeer", id)
	cancel()
	if e != nil || !*v {
		log.Error(fmt.Sprintf("Cannot add peer %s: %v", id, e))
	} else {
		log.Info(fmt.Sprintf("Added %s.", id))
	}

	return nil
}

func (ma *metaAdmin) update() {
	refresh := false

	newAdmin := ma.getAdminAddress()
	if newAdmin == nilAddress {
		return
	} else if newAdmin != ma.admin {
		ma.admin = newAdmin
		refresh = true
	}

	blockNum, modifiedBlock, blocksPer, nodes, addedNodes, updatedNodes, deletedNodes, err := ma.getData(refresh)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to get nodes: %v", err))
	} else if refresh ||
		(modifiedBlock != 0 && ma.modifiedBlock != modifiedBlock) {
		log.Debug(fmt.Sprintf("Modified Block: %d", modifiedBlock))

		ma.modifiedBlock = modifiedBlock
		ma.blocksPer = blocksPer

		_nodes := map[string]*metaNode{}
		for _, i := range nodes {
			_nodes[i.Id] = i
			if i.V4Id == ma.nodeInfo.ID {
				ma.self = i
			}
		}
		ma.nodes = _nodes

		log.Debug(fmt.Sprintf("Added:\n"))
		for _, i := range addedNodes {
			log.Debug(fmt.Sprintf("%v\n", i))
			ma.addPeer(i)
		}
		log.Debug(fmt.Sprintf("Updated:\n"))
		for _, i := range updatedNodes {
			log.Debug(fmt.Sprintf("%v\n", i))
		}
		log.Debug(fmt.Sprintf("Deleted:\n"))
		for _, i := range deletedNodes {
			log.Debug(fmt.Sprintf("%v\n", i))
		}
	}

	if blockNum != 0 {
		ma.lastBlock = blockNum
	}
}

func (ma *metaAdmin) checkMining() {
	on := false
	if ma.nodeInfo != nil && ma.nodeInfo.ID == admin.bootNodeId {
		on = true
	} else if ma.self != nil && ma.self.Partner {
		on = true
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mining *bool
	err := ma.rpcCli.CallContext(ctx, &mining, "eth_mining")
	if err != nil {
		log.Error("Checking mining status", "failure", err)
		return
	}

	if on == *mining {
		return
	} else if on {
		err := ma.rpcCli.CallContext(ctx, &mining, "miner_start", 1)
		if err != nil {
			log.Error("Starting miner", "failed", err)
			return
		} else {
			log.Info("Started miner")
		}
	} else {
		err := ma.rpcCli.CallContext(ctx, &mining, "miner_stop", 1)
		if err != nil {
			log.Error("Stopping miner", "failed", err)
			return
		} else {
			log.Info("Stopped miner")
		}
	}
}

func (ma *metaAdmin) run() {
	for {
		if ma.nodeInfo == nil {
			nodeInfo, err := ma.getNodeInfo()
			if err != nil {
				log.Error("Failed to get node info", "error", err)
			} else {
				ma.nodeInfo = nodeInfo
			}
		}

		if ma.anchor == nilAddress {
			ma.anchor = ma.getAdminAnchorAddress()
		}
		if ma.anchor != nilAddress && ma.admin == nilAddress {
			ma.admin = ma.getAdminAddress()
		}
		if ma.admin != nilAddress && ma.nodeInfo != nil {
			ma.update()
			if !ma.etcdIsRunning() {
				EtcdStart()
			}
		}

		ma.checkMining()

		to := make(chan bool, 1)
		go func() {
			time.Sleep(5 * time.Second)
			to <- true
		}()
		select {
		case <-ma.Updates:
		case <-to:
		}
	}
}

type reward struct {
	Addr   common.Address `json:"addr"`
	Reward uint64         `json:"reward"`
}

// to get around 64 bit boundary. big.Float didn't help here.
func distributeRewards(six int, members []*metaMember, rewards []reward, amount int64) {
	n := len(members)
	var u int64
	for i := 0; i < n; i++ {
		rewards[i].Addr = members[i].Addr
		u += int64(members[i].Stake)
	}

	var h, l uint64 = uint64(amount) >> 32, uint64(amount) & uint64(0x0FFFFFFFF)
	var hd, ld float64 = float64(h) / float64(u), float64(l) / float64(u) // slopes
	var hv, lv, vi float64 = 0, 0, 0
	var s, vj uint64

	for i := 0; i < n; i++ {
		s = uint64(members[six].Stake)
		vi = hv + hd*float64(s)
		vj = uint64(math.Floor(vi+.5)-math.Floor(hv+.5)) << 32
		hv = vi
		vi = lv + ld*float64(s)
		vj += uint64(math.Floor(vi+.5) - math.Floor(lv+.5))
		lv = vi
		rewards[six].Reward = vj

		six = (six + 1) % n
	}
}

func (ma *metaAdmin) calculateRewards(num, blockReward, fees *big.Int, addBalance func(common.Address, *big.Int)) (rewards []byte, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	members, err := ma.getMembers(ctx, big.NewInt(num.Int64()-1))
	if err != nil {
		return
	}
	if len(members) == 0 {
		err = fmt.Errorf("Not initialized")
		return
	}

	n := len(members)
	rr := make([]reward, n)
	distributeRewards(int(new(big.Int).Set(num).Mod(num, big.NewInt(int64(len(members)))).Int64()),
		members, rr, new(big.Int).Add(blockReward, fees).Int64())

	if addBalance != nil {
		bi := new(big.Int)
		for _, i := range rr {
			addBalance(i.Addr, bi.SetUint64(i.Reward))
		}
	}

	rewards, err = json.Marshal(rr)
	return
}

func (ma *metaAdmin) verifyRewards(r1, r2 []byte) error {
	var err error
	var a, b []reward

	if err = json.Unmarshal(r1, &a); err != nil {
		return err
	}
	if err = json.Unmarshal(r2, &b); err != nil {
		return err
	}

	err = fmt.Errorf("Incorrect Rewards")
	if len(a) != len(b) {
		return err
	}
	for i := 0; i < len(a); i++ {
		if !bytes.Equal(a[i].Addr.Bytes(), b[i].Addr.Bytes()) ||
			a[i].Reward != b[i].Reward {
			return err
		}
	}

	return nil
}

func calculateRewards(num, blockReward, fees *big.Int, addBalance func(common.Address, *big.Int)) ([]byte, error) {
	return admin.calculateRewards(num, blockReward, fees, addBalance)
}

func verifyRewards(num *big.Int, rewards string) error {
	return nil
	//return admin.verifyRewards(num, rewards)
}

func signBlock(hash common.Hash) (nodeId, sig []byte, err error) {
	if admin == nil {
		err = errors.New("Not initialized")
		return
	}

	prvKey := admin.stack.Server().PrivateKey
	sig, err = crypto.Sign(hash.Bytes(), prvKey)
	v5id := discv5.PubkeyID(&prvKey.PublicKey) 
	nodeId = v5id[:]
	return
}

func verifyBlockSig(height *big.Int, nodeId []byte, hash common.Hash, sig []byte) bool {
	// TODO: need to check if nodeId is a valid partner in the 'height' block.
	pubKey, err := crypto.Ecrecover(hash.Bytes(), sig)
	return err == nil && nodeId != nil && len(pubKey) > 1 && bytes.Equal(nodeId, pubKey[1:])
}

func (ma *metaAdmin) getNodeInfo() (*p2p.NodeInfo, error) {
	var nodeInfo *p2p.NodeInfo
	ctx, cancel := context.WithCancel(context.Background())
	err := ma.rpcCli.CallContext(ctx, &nodeInfo, "admin_nodeInfo")
	cancel()
	if err != nil {
		log.Error("Cannot get node info", "error", err)
	}
	return nodeInfo, err
}

func (ma *metaAdmin) getPeerInfo(id string) (*p2p.NodeInfo, error) {
	var nodeInfo *p2p.NodeInfo
	ctx, cancel := context.WithCancel(context.Background())
	err := ma.rpcCli.CallContext(ctx, &nodeInfo, "admin_peerInfo", id)
	cancel()
	if err != nil {
		log.Error("Cannot get peer info", "id", id, "error", err)
	}
	return nodeInfo, err
}

func (ma *metaAdmin) isPeerUp(id string) bool {
	nodeInfo, err := ma.getPeerInfo(id)
	return err == nil && nodeInfo != nil
}

func AmPartner() bool {
	if admin == nil {
		return false
	}

	admin.lock.Lock()
	defer admin.lock.Unlock()

	return (admin.nodeInfo != nil && admin.nodeInfo.ID == admin.bootNodeId) ||
		(admin.self != nil && admin.self.Partner)
}

func IsPartner(id string) bool {
	if admin == nil {
		return false
	}

	admin.lock.Lock()
	defer admin.lock.Unlock()

	n, ok := admin.nodes[id]
	if !ok {
		if id == admin.bootNodeId {
			return true
		} else {
			return false
		}
	}

	// TODO: need to check NotBefore and NotAfter

	return n.Partner
}

func (ma *metaAdmin) pendingEmpty() bool {
	type txpool_status struct {
		Pending hexutil.Uint `json:"pending"`
		Queued  hexutil.Uint `json:"queued"`
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var status txpool_status
	if err := admin.rpcCli.CallContext(ctx, &status, "txpool_status"); err != nil {
		log.Error("Canot get txpool.status", "error", err)
		return false
	}

	return status.Pending == 0
}

func LogBlock(height int64) {
	if admin == nil || admin.self == nil {
		return
	}

	admin.lock.Lock()
	defer admin.lock.Unlock()

	admin.blocksMined++
	height++
	if admin.blocksMined >= admin.blocksPer &&
		int(height)%admin.blocksPer == 0 {
		// time to yield leader role

		if !admin.pendingEmpty() {
			log.Info("Metadium - not yielding due to pending txs...")
			return
		}
		_, next, _ := admin.getMinerNodes(int(height), true)
		if next.Id == admin.self.Id {
			log.Info("Metadium - yield to self", "mined", admin.blocksMined,
				"new miner", "self")
		} else {
			if err := admin.etcdMoveLeader(next.Name); err == nil {
				log.Info("Metadium - yielded", "mined", admin.blocksMined,
					"new miner", next.Name)
				admin.blocksMined = 0
			} else {
				log.Error("Metadium - yield failed", "mined", admin.blocksMined,
					"new miner", next.Name, "error", err)
			}
		}
	}
}

func (ma *metaAdmin) toMiningPeers(nodes []*metaNode) string {
	var bb bytes.Buffer
	for _, n := range nodes {
		if bb.Len() != 0 {
			bb.Write([]byte(" "))
		}
		bb.Write([]byte(fmt.Sprintf("%s/%s", n.Name, n.Status)))
		if n.Miner {
			bb.Write([]byte("/*"))
		}
	}
	return string(bb.Bytes())
}

func (ma *metaAdmin) miners() string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	block, err := ma.cli.HeaderByNumber(ctx, nil)
	if err != nil {
		return ""
	}
	height := int(block.Number.Int64())

	_, _, nodes := ma.getMinerNodes(height+1, false)
	return ma.toMiningPeers(nodes)
}

func Info() interface{} {
	if admin == nil {
		return ""
	} else {
		self := admin.self
		var nodes []*metaNode
		for _, i := range admin.nodes {
			nodes = append(nodes, i)
		}
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].Name < nodes[j].Name
		})

		info := &map[string]interface{}{
			"consensus":     params.ConsensusMethod,
			"anchor":        admin.anchor,
			"admin":         admin.admin,
			"modifiedblock": admin.modifiedBlock,
			"blocksPer":     admin.blocksPer,
			"self":          self,
			"nodes":         nodes,
			"miners":        admin.miners(),
			"etcd":          admin.etcdInfo(),
		}
		return info
	}
}

func getMinerStatus() *metaapi.MetadiumMinerStatus {
	if admin == nil || admin.self == nil {
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	header, err := admin.cli.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil
	}
	height := int(header.Number.Int64())

	_, _, nodes := admin.getMinerNodes(height+1, false)
	miningPeers := admin.toMiningPeers(nodes)

	admin.lock.Lock()
	defer admin.lock.Unlock()

	var (
		mining, syncing bool
		syncProgress    *ethereum.SyncProgress
	)
	if err = admin.rpcCli.CallContext(ctx, &mining, "eth_mining"); err != nil {
		mining = false
	}
	if syncProgress, err = admin.cli.SyncProgress(ctx); err != nil {
		log.Error("SyncProgress failed", "error", err)
	}
	if syncProgress != nil {
		syncing = true
	} else {
		syncing = false
		syncProgress = &ethereum.SyncProgress{}
	}

	return &metaapi.MetadiumMinerStatus{
		Name:              admin.self.Name,
		Id:                admin.self.Id,
		Addr:              fmt.Sprintf("%s:%d", admin.self.Ip, admin.self.Port),
		Status:            "up",
		Miner:             admin.self.Miner,
		MiningPeers:       miningPeers,
		Mining:            mining,
		LatestBlockHeight: header.Number,
		LatestBlockHash:   header.Hash(),
		Syncing:           syncing,
		SyncProgress:      *syncProgress,
	}
}

// Returns the array of peer status
// 'id' could be null, a name, node id (public key) or ip address of a miner
func getMiners(id string) []*metaapi.MetadiumMinerStatus {
	if admin == nil {
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodes := admin.getNodes()

	var node *metaNode
	for _, n := range nodes {
		if strings.EqualFold(n.Name, id) || strings.EqualFold(n.Id, id) || strings.EqualFold(n.Ip, id) {
			node = n
			break
		}
	}

	getDownStatus := func(node *metaNode) *metaapi.MetadiumMinerStatus {
		return &metaapi.MetadiumMinerStatus{
			Name:   node.Name,
			Id:     node.Id,
			Addr:   fmt.Sprintf("%s:%d", node.Ip, node.Port),
			Status: "down",
		}
	}

	var miners []*metaapi.MetadiumMinerStatus
	var err error
	msgch := make(chan interface{}, 1)
	metaapi.SetMsgChannel(msgch)
	defer func() {
		metaapi.SetMsgChannel(nil)
		close(msgch)
	}()

	timer := time.NewTimer(5 * time.Second)
	peers := map[string]bool{}
	count := 0

	if node != nil {
		if admin.self != nil && admin.self.Id == node.Id {
			miners = append(miners, getMinerStatus())
			return miners
		} else if !admin.isPeerUp(node.Id) {
			miners = append(miners, getDownStatus(node))
			return miners
		}

		err = admin.rpcCli.CallContext(ctx, nil, "admin_requestMinerStatus", &node.Id)
		if err != nil {
			log.Error("Metadium RequestMinerStatus Failed", "id", node.Id, "error", err)
		} else {
			peers[node.Name] = false
			count++
		}
	} else {
		for _, n := range nodes {
			if admin.self != nil && admin.self.Id == n.Id {
				miners = append(miners, getMinerStatus())
				continue
			} else if !admin.isPeerUp(n.Id) {
				miners = append(miners, getDownStatus(n))
				continue
			}

			err = admin.rpcCli.CallContext(ctx, nil, "admin_requestMinerStatus", n.Id)
			if err != nil {
				log.Error("Metadium RequestMinerStatus Failed", "id", n.Id, "error", err)
			} else {
				peers[n.Name] = false
				count++
			}
		}
	}

	done := false
	if count == 0 {
		done = true
	}
	for {
		if done {
			break
		}
		select {
		case msg := <-msgch:
			s, ok := msg.(*metaapi.MetadiumMinerStatus)
			if !ok {
				continue
			}
			if b, exists := peers[s.Name]; exists {
				miners = append(miners, s)
				if b == false {
					peers[s.Name] = true
					count--
					if count <= 0 {
						done = true
					}
				}
			}
		case <-timer.C:
			done = true
		}
	}

	if len(miners) > 1 {
		sort.Slice(miners, func(i, j int) bool {
			return miners[i].Name < miners[j].Name
		})
	}
	return miners
}

func (ma *metaAdmin) getTxPoolStatus() (pending, queued uint, err error) {
	var data map[string]hexutil.Uint

	ctx, cancel := context.WithCancel(context.Background())
	err = ma.rpcCli.CallContext(ctx, &data, "txpool_status")
	cancel()

	if err != nil {
		return
	}
	p, b1 := data["pending"]
	q, b2 := data["queued"]
	if !b1 || !b2 {
		err = fmt.Errorf("Invalid Data")
	} else {
		pending = uint(p)
		queued = uint(q)
	}

	return
}

func requirePendingTxs() bool {
	if admin == nil {
		return false
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	block, err := admin.cli.HeaderByNumber(ctx, nil)
	if err != nil {
		return false
	}
	height := int(block.Number.Int64())

	if !IsMiner(height + 1) {
		return false
	}

	p, _, e := admin.getTxPoolStatus()
	if e != nil {
		return false
	} else if p > 0 {
		return false
	}

	return true
}

/* EOF */
