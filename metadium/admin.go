/* admin.go */

package metadium

import (
	"bufio"
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	metaapi "github.com/ethereum/go-ethereum/metadium/api"
	metaminer "github.com/ethereum/go-ethereum/metadium/miner"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

type metaNode struct {
	Partner   bool   `json:"partner"`
	Name      string `json:"name"`
	Id        string `json:"id"`
	Ip        string `json:"ip"`
	Port      int    `json:"port"`
	NotBefore int    `json:"notBefore"`
	NotAfter  int    `json:"notAfter"`
	Prev      string `json:"prev"`
	Next      string `json:"next"`
}

type metaStake struct {
	account common.Address `json:"account"`
	stake   int            `json:"stake"`
}

type metaAdmin struct {
	bootNodeId  string // allowed to generate block without admin contract
	bootAccount common.Address
	nodeInfo    *p2p.NodeInfo
	from        common.Address
	to          common.Address
	Updates     chan bool
	rpcCli      *rpc.Client
	cli         *ethclient.Client
	abi         abi.ABI

	lastBlock     int
	modifiedBlock int
	blocksPer     int
	self          *metaNode

	lock  *sync.Mutex
	nodes map[string]*metaNode
}

var (
	magic      = 0x4d6574616469756d // Metadium
	big0       = big.NewInt(0)
	nilAddress = common.Address{}
	admin      *metaAdmin
)

// var <name>_contract = web3.eth.contract([{<abi>}]);
func loadAbis(r io.Reader) (map[string]abi.ABI, error) {
	abis := map[string]abi.ABI{}

	re := regexp.MustCompile(`^var ([^_]+)_contract[^(]+\(([^)]+)\);$`)
	b := bufio.NewReaderSize(r, 1024*1024)
	for {
		line, _, err := b.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		submatches := re.FindSubmatch(line)
		if submatches == nil || len(submatches) != 3 {
			continue
		}

		def, err := abi.JSON(bytes.NewReader(submatches[2]))
		if err != nil {
			// ignore error
			continue
		}

		abis[string(submatches[1])] = def
	}

	return abis, nil
}

func (n *metaNode) eq(m *metaNode) bool {
	if n.Partner == m.Partner && n.Name == m.Name && n.Id == m.Id &&
		n.Ip == m.Ip && n.Port == m.Port && n.NotBefore == m.NotBefore &&
		n.NotAfter == m.NotAfter {
		return true
	} else {
		return false
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
	return nodeId, block.Coinbase, nil
}

// it should be the first transaction of the coinbase of the genesis block
func (ma *metaAdmin) getAdminContractAddress() common.Address {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := uint64(0); i < 10; i++ {
		addr := crypto.CreateAddress(ma.bootAccount, i)
		m, _ := ma.getInt(ctx, &addr, nil, "magic")
		if m == magic {
			return addr
		}
	}

	return nilAddress
}

func (ma *metaAdmin) getInt(ctx context.Context, to *common.Address, block *big.Int, name string) (int, error) {
	input, err := ma.abi.Pack(name)
	if err != nil {
		return 0, err
	}

	msg := ethereum.CallMsg{From: ma.from, To: to, Data: input}
	output, err := ma.cli.CallContract(ctx, msg, block)
	if err != nil {
		return 0, err
	}

	var v *big.Int
	err = ma.abi.Unpack(&v, name, output)
	if err != nil {
		return 0, err
	} else {
		return int(v.Int64()), nil
	}
}

func (ma *metaAdmin) igetNodes() []*metaNode {
	ma.lock.Lock()
	defer ma.lock.Unlock()

	var nodes []*metaNode
	for _, i := range ma.nodes {
		nodes = append(nodes, i)
	}
	return nodes
}

func (ma *metaAdmin) getNode(ctx context.Context, id string, block *big.Int) (*metaNode, error) {
	var method string
	var input []byte
	var err error

	if len(id) == 0 {
		method = "firstNode"
		input, err = ma.abi.Pack(method)
	} else {
		method = "getNode"
		input, err = ma.abi.Pack(method, []byte(id))
	}

	msg := ethereum.CallMsg{From: ma.from, To: &ma.to, Data: input}
	output, err := ma.cli.CallContract(ctx, msg, block)
	if err != nil {
		return nil, err
	}

	var present bool
	var jsonOut string
	o := []interface{}{&present, &jsonOut}
	err = ma.abi.Unpack(&o, method, output)
	if err != nil {
		return nil, err
	}

	// trim the json output
	if ix := strings.Index(jsonOut, "\000"); ix >= 0 {
		jsonOut = jsonOut[:ix]
	}

	n := new(metaNode)
	err = json.Unmarshal([]byte(jsonOut), n)
	return n, err
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

	modifiedBlock, err = ma.getInt(ctx, &ma.to, block.Number, "modifiedBlock")
	if err != nil {
		return
	}
	if !refresh && ma.modifiedBlock == modifiedBlock {
		return
	}

	blocksPer, err = ma.getInt(ctx, &ma.to, block.Number, "blocksPer")
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

	oldNodes := ma.igetNodes()
	sort.Slice(oldNodes, func(i, j int) bool {
		return oldNodes[i].Id < oldNodes[j].Id
	})
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Id < nodes[j].Id
	})

	i, j := 0, 0
	for {
		if i >= len(oldNodes) || j >= len(nodes) {
			break
		}
		v := strings.Compare(oldNodes[i].Id, nodes[j].Id)
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

func StartAdmin(stack *node.Node, abiFile string) {
	if !(params.ConsensusMethod == params.ConsensusPoA ||
		params.ConsensusMethod == params.ConsensusETCD ||
		params.ConsensusMethod == params.ConsensusPBFT) {
		utils.Fatalf("Invalid Consensus Method: %d\n", params.ConsensusMethod)
	}

	metaminer.IsMinerFunc = IsMiner
	metaapi.Info = Info

	rpcCli, err := stack.Attach()
	if err != nil {
		utils.Fatalf("Failed to attach to self: %v", err)
	}

	f, err := os.Open(abiFile)
	if err != nil {
		utils.Fatalf("Failed to open %s: %v\n", abiFile, err)
	}

	abis, err := loadAbis(f)
	if err != nil {
		utils.Fatalf("Loading ABI failed: %v", err)
	}
	if _, ok := abis["Admin"]; !ok {
		utils.Fatalf("Cannot find governance abi")
	}

	admin = &metaAdmin{
		lock:    &sync.Mutex{},
		from:    nilAddress,
		to:      nilAddress,
		Updates: make(chan bool, 10),
		rpcCli:  rpcCli,
		cli:     ethclient.NewClient(rpcCli),
		abi:     abis["Admin"],
	}

	admin.bootNodeId, admin.bootAccount, err = admin.getGenesisInfo()
	if err != nil {
		utils.Fatalf("Cannot get contract address from genesis block: %v", err)
	}

	go admin.run()
}

func (ma *metaAdmin) addPeer(node *metaNode) error {
	if node.Id == ma.nodeInfo.ID {
		return nil
	}

	var v *bool
	ctx, cancel := context.WithCancel(context.Background())
	id := fmt.Sprintf("enode://%s@%s:%d", node.Id, node.Ip, node.Port)
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
	//refresh := params.ConsensusMethod == "etcd" && !etcdhelper.IsRunning()
	refresh := false
	blockNum, modifiedBlock, blocksPer, nodes, addedNodes, updatedNodes, deletedNodes, err := ma.getData(refresh)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to get nodes: %v", err))
	} else if refresh ||
		(modifiedBlock != 0 && ma.modifiedBlock != modifiedBlock) {
		log.Debug(fmt.Sprintf("Modified Block: %d", modifiedBlock))

		self := ma.self
		ma.modifiedBlock = modifiedBlock
		ma.blocksPer = blocksPer

		_nodes := map[string]*metaNode{}
		for _, i := range nodes {
			_nodes[i.Id] = i
			if i.Id == ma.nodeInfo.ID {
				ma.self = i
			}
		}
		ma.nodes = _nodes

		if params.ConsensusMethod == params.ConsensusETCD {
			needToStartEtcd := self == nil && ma.self != nil
			if !needToStartEtcd && ma.self != nil {
				//needToStartEtcd = !etcdhelper.IsRunning()
			}
			if needToStartEtcd {
				//ma.startEtcd()
			}
		}

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

func (ma *metaAdmin) run() {
	for {
		if ma.nodeInfo == nil {
			nodeInfo, err := ma.getNodeInfo()
			if err != nil {
				log.Error(fmt.Sprintf("Failed to get node info: %v", err))
			} else {
				ma.nodeInfo = nodeInfo
			}
		}

		if ma.to == nilAddress {
			ma.to = ma.getAdminContractAddress()
		}

		if ma.to != nilAddress && ma.nodeInfo != nil {
			ma.update()
		}

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

func (ma *metaAdmin) getNodeInfo() (*p2p.NodeInfo, error) {
	var nodeInfo *p2p.NodeInfo
	ctx, cancel := context.WithCancel(context.Background())
	err := ma.rpcCli.CallContext(ctx, &nodeInfo, "admin_nodeInfo")
	cancel()
	if err != nil {
		log.Error(fmt.Sprintf("Cannot get node info: %v", err))
	}
	return nodeInfo, err
}

func (ma *metaAdmin) getPeerInfo(id string) (*p2p.NodeInfo, error) {
	nodeId, err := discover.HexID(id)
	if err != nil {
		return nil, err
	}

	var nodeInfo *p2p.NodeInfo
	ctx, cancel := context.WithCancel(context.Background())
	err = ma.rpcCli.CallContext(ctx, &nodeInfo, "admin_peerInfo", nodeId)
	cancel()
	if err != nil {
		log.Error(fmt.Sprintf("Cannot get peer info(%s): %v", id, err))
	}
	return nodeInfo, err
}

func (ma *metaAdmin) isPeerUp(id string) bool {
	nodeInfo, err := ma.getPeerInfo(id)
	return err == nil && nodeInfo != nil
}

func IsMiner(height int) bool {
	if params.ConsensusMethod == params.ConsensusPoW {
		return true
	} else if params.ConsensusMethod == params.ConsensusETCD {
		return false
	} else if params.ConsensusMethod == params.ConsensusPoA {
		if admin == nil {
			return false
		} else if admin.self == nil || len(admin.nodes) <= 0 {
			if admin.nodeInfo != nil && admin.nodeInfo.ID == admin.bootNodeId {
				return true
			} else {
				return false
			}
		}

		if height != admin.lastBlock {
			admin.update()
		}

		nodes := admin.igetNodes()
		var vnodes []*metaNode
		for _, i := range nodes {
			if (i.NotBefore != 0 && height < i.NotBefore) ||
				(i.NotAfter != 0 && height > i.NotAfter) {
			} else {
				vnodes = append(vnodes, i)
			}
		}

		if len(vnodes) == 0 {
			return true
		}

		sort.Slice(vnodes, func(i, j int) bool {
			return vnodes[i].Id < vnodes[j].Id
		})

		ix := height / admin.blocksPer % len(vnodes)
		i := ix
		for n := 0; n < len(vnodes); n++ {
			if vnodes[i].Id == admin.self.Id {
				if i == ix {
					log.Debug("Am the miner")
				} else {
					log.Debug(fmt.Sprintf("Am the miner as %s (or others) are not up.", vnodes[ix].Name))
				}
				return true
			} else if admin.isPeerUp(vnodes[i].Id) {
				log.Debug(fmt.Sprintf("%s is the miner", vnodes[i].Name))
				return false
			}
			i = (i + 1) % len(vnodes)
		}

		log.Error("No miners found")
		return false
	} else {
		return false
	}
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
			"to":            admin.to,
			"modifiedblock": admin.modifiedBlock,
			"blocksPer":     admin.blocksPer,
			"self":          self,
			"nodes":         nodes,
		}
		return info
	}
}

/* EOF */
