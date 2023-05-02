/* admin.go */

package wemix

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"path"
	"sort"
	"strings"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	wemixapi "github.com/ethereum/go-ethereum/wemix/api"
	"github.com/ethereum/go-ethereum/wemix/metclient"
	wemixminer "github.com/ethereum/go-ethereum/wemix/miner"
)

type wemixNode struct {
	Name  string         `json:"name"`
	Enode string         `json:"enode"`
	Id    string         `json:"id"`
	Ip    string         `json:"ip"`
	Port  int            `json:"port"`
	Addr  common.Address `json:"addr"`

	Status string `json:"status"`
	Miner  bool   `json:"miner"`
}

type wemixMember struct {
	Staker common.Address `json:"address"`
	Reward common.Address `json:"reward"`
	Stake  *big.Int       `json:"stake"`
}

type wemixAdmin struct {
	stack *node.Node

	bootNodeId  string // allowed to generate block without admin contract
	bootAccount common.Address
	nodeInfo    *p2p.NodeInfo
	registry    *metclient.RemoteContract
	gov         *metclient.RemoteContract
	staking     *metclient.RemoteContract
	envStorage  *metclient.RemoteContract
	Updates     chan bool
	rpcCli      *rpc.Client
	cli         *ethclient.Client

	etcd        *embed.Etcd
	etcdCli     *clientv3.Client
	etcdDir     string
	etcdPort    int
	etcdTimeout time.Duration

	lastBlock     int64
	modifiedBlock int64

	blockInterval        int64
	blocksPer            int64
	blockReward          *big.Int
	maxPriorityFeePerGas *big.Int
	maxBaseFee           *big.Int
	gasLimit             *big.Int
	baseFeeMaxChangeRate int64
	gasTargetPercentage  int64

	self *wemixNode

	lock  *sync.Mutex
	nodes map[string]*wemixNode

	// # of blocks consecutively mined by this node
	blocksMined int64
}

// latest block generated
type wemixWork struct {
	Height int64       `json:"height"`
	Hash   common.Hash `json:"hash"`
}

// block build parameters for caching
type blockBuildParameters struct {
	height               uint64
	blockInterval        int64
	maxBaseFee           *big.Int
	gasLimit             *big.Int
	baseFeeMaxChangeRate int64
	gasTargetPercentage  int64
}

// reward related parameters
type rewardParameters struct {
	rewardAmount                                 *big.Int
	staker, ecoSystem, maintenance, feeCollector *common.Address
	members                                      []*wemixMember
	distributionMethod                           []*big.Int
	blocksPer                                    int64
}

var (
	// "Wemix Registry"
	magic, _        = big.NewInt(0).SetString("0x57656d6978205265676973747279", 0)
	etcdClusterName = "Wemix"
	big0            = big.NewInt(0)
	nilAddress      = common.Address{}
	admin           *wemixAdmin

	ErrAlreadyRunning = errors.New("already running")
	ErrExists         = errors.New("already exists")
	ErrIneligible     = errors.New("not eligible")
	ErrInvalidEnode   = errors.New("invalid enode")
	ErrInvalidToken   = errors.New("invalid token")
	ErrInvalidWork    = errors.New("invalid work")
	ErrNotFound       = errors.New("not found")
	ErrNotRunning     = errors.New("not running")

	etcdCompactFrequency = int64(100)
	etcdCompactWindow    = int64(100)

	// cached block build parameters
	blockBuildParamsLock = &sync.Mutex{}
	blockBuildParams     *blockBuildParameters

	// testnet block 94 rewards
	testnetBlock94Rewards       []reward
	testnetBlock94RewardsString = `[
		{ "addr": "0x6f488615e6b462ce8909e9cd34c3f103994ab2fb", "reward": 100000000000000000 },
		{ "addr": "0x6bd26c4a45e7d7cac2a389142f99f12e5713d719", "reward": 250000000000000000 },
		{ "addr": "0x816e30b6c314ba5d1a67b1b54be944ce4554ed87", "reward": 306213253695614752 }]`
)

func (n *wemixNode) eq(m *wemixNode) bool {
	if n.Name == m.Name && n.Id == m.Id && n.Ip == m.Ip && n.Port == m.Port {
		return true
	} else {
		return false
	}
}

// convert v5 id to v4 id
func toIdv4(id string) (string, error) {
	if len(id) == 64 {
		return id, nil
	} else if len(id) == 128 {
		idv4, err := enode.ParseV4(fmt.Sprintf("enode://%v@127.0.0.1:8589", id))
		if err != nil {
			return "", err
		} else {
			return idv4.ID().String(), nil
		}
	} else {
		return "", fmt.Errorf("Invalid V5 Identifier")
	}
}

// returns
//  1. extradata of genesis block, which is the id of the node that is allowed
//     to generated blocks before admin contract is established.
//  2. returns the coinbase of genesis block, which should be the admin
//     contract creator
func (ma *wemixAdmin) getGenesisInfo() (string, common.Address, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	block, err := ma.cli.HeaderByNumber(ctx, big0)
	if err != nil {
		return "", common.Address{}, err
	}

	var nodeId string
	if len(block.Extra) < 64 {
		return "", common.Address{}, fmt.Errorf("Invalid bootnode id in the genesis block")
	} else if len(block.Extra) == 64 {
		nodeId = hex.EncodeToString(block.Extra)
	} else if len(block.Extra) <= 128 {
		return "", common.Address{}, fmt.Errorf("Invalid bootnode id in the genesis block")
	} else {
		nodeId = string(block.Extra[len(block.Extra)-128:])
	}
	nodeId, _ = toIdv4(nodeId)
	return nodeId, block.Coinbase, nil
}

func (ma *wemixAdmin) getRegistryAddress(ctx context.Context, cli *ethclient.Client, registryAbi abi.ABI, height *big.Int) (*common.Address, error) {
	contract := &metclient.RemoteContract{
		Cli: cli,
		Abi: registryAbi,
	}
	for i := uint64(0); i < 10; i++ {
		addr := crypto.CreateAddress(ma.bootAccount, i)
		contract.To = &addr

		var v *big.Int
		err := metclient.CallContract(ctx, contract, "magic", nil, &v, height)
		if err == nil && v.Cmp(magic) == 0 {
			return &addr, nil
		}
	}
	return nil, wemixminer.ErrNotInitialized
}

// it should be the first transaction of the coinbase of the genesis block
func (ma *wemixAdmin) getAdminAddresses() (registry, gov, staking, envStorage *common.Address, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	registry, gov, staking, envStorage = nil, nil, nil, nil
	contract := &metclient.RemoteContract{
		Cli: ma.cli,
		Abi: ma.registry.Abi,
	}
	if ma.registry != nil && ma.registry.To != nil {
		registry = ma.registry.To
	} else {
		registry, err = ma.getRegistryAddress(ctx, ma.cli, ma.registry.Abi, nil)
		if err != nil {
			err = ethereum.NotFound
			return
		}
	}
	contract.To = registry

	n1 := metclient.ToBytes32("GovernanceContract")
	n2 := metclient.ToBytes32("Staking")
	n3 := metclient.ToBytes32("EnvStorage")
	var a1, a2, a3 common.Address
	input := []interface{}{n1}
	if err = metclient.CallContract(ctx, contract, "getContractAddress", input, &a1, nil); err != nil {
		return
	}
	input = []interface{}{n2}
	if err = metclient.CallContract(ctx, contract, "getContractAddress", input, &a2, nil); err != nil {
		return
	}
	input = []interface{}{n3}
	if err = metclient.CallContract(ctx, contract, "getContractAddress", input, &a3, nil); err != nil {
		return
	}

	log.Debug("Wemix Contract Address",
		hex.EncodeToString(n1[:]), a1.Hex(),
		hex.EncodeToString(n2[:]), a2.Hex(),
		hex.EncodeToString(n3[:]), a3.Hex())

	gov, staking, envStorage = &a1, &a2, &a3
	return
}

func (ma *wemixAdmin) getInt(ctx context.Context, contract *metclient.RemoteContract, block *big.Int, name string) (int64, error) {
	var v *big.Int
	err := metclient.CallContract(ctx, contract, name, nil, &v, block)
	if err != nil {
		return 0, err
	} else {
		return v.Int64(), nil
	}
}

// TODO: error handling
func (ma *wemixAdmin) getRegGovEnvContracts(ctx context.Context, height *big.Int) (reg, gov, env, staking *metclient.RemoteContract, err error) {
	if ma.registry == nil {
		err = wemixminer.ErrNotInitialized
		return
	}
	reg = &metclient.RemoteContract{
		Cli: ma.cli,
		Abi: ma.registry.Abi,
	}
	env = &metclient.RemoteContract{
		Cli: ma.cli,
		Abi: ma.envStorage.Abi,
	}
	gov = &metclient.RemoteContract{
		Cli: ma.cli,
		Abi: ma.gov.Abi,
	}
	staking = &metclient.RemoteContract{
		Cli: ma.cli,
		Abi: ma.staking.Abi,
	}
	if ma.registry.To != nil {
		reg.To = ma.registry.To
	} else {
		var addr *common.Address
		if addr, err = ma.getRegistryAddress(ctx, ma.cli, reg.Abi, height); err != nil {
			err = wemixminer.ErrNotInitialized
			return
		}
		reg.To = addr
	}

	var addr common.Address
	input := []interface{}{metclient.ToBytes32("GovernanceContract")}
	if err = metclient.CallContract(ctx, reg, "getContractAddress", input, &addr, height); err != nil {
		err = wemixminer.ErrNotInitialized
		return
	}
	gov.To = &common.Address{}
	gov.To.SetBytes(addr.Bytes())

	input = []interface{}{metclient.ToBytes32("EnvStorage")}
	if err = metclient.CallContract(ctx, reg, "getContractAddress", input, &addr, height); err != nil {
		err = wemixminer.ErrNotInitialized
		return
	}
	env.To = &common.Address{}
	env.To.SetBytes(addr.Bytes())

	input = []interface{}{metclient.ToBytes32("Staking")}
	if err = metclient.CallContract(ctx, reg, "getContractAddress", input, &addr, height); err != nil {
		err = wemixminer.ErrNotInitialized
		return
	}
	staking.To = &common.Address{}
	staking.To.SetBytes(addr.Bytes())

	return
}

// returns []*wemixNode from map[string]*wemixNode
func (ma *wemixAdmin) getNodes() []*wemixNode {
	ma.lock.Lock()
	defer ma.lock.Unlock()

	var nodes []*wemixNode
	for _, i := range ma.nodes {
		nodes = append(nodes, i)
	}
	return nodes
}

// returns
//  1. currentMiner *wemixNode: the current leader
//  2. nextMiner *wemixNode: the most eligible miner for the given height,
//     which is up and running
//  3. nodes []*wemixNode: copies of map[string]*wemixNode, not references,
//     sorted by id, i.e. mining order
//
// 'locked' indicates whether ma.lock is held by the caller or not
func (ma *wemixAdmin) getMinerNodes(height int64, locked bool) (*wemixNode, *wemixNode, []*wemixNode) {
	var nodes []*wemixNode
	if !locked {
		ma.lock.Lock()
	}
	for _, i := range ma.nodes {
		n := new(wemixNode)
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
	var miner, next *wemixNode
	ix := int(height/admin.blocksPer) % len(nodes)
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

// get nodes from the Governance contract
func (ma *wemixAdmin) getWemixNodes(ctx context.Context, block *big.Int) ([]*wemixNode, error) {
	var (
		nodes           []*wemixNode
		addr            common.Address
		name, enode, ip []byte
		port            *big.Int
		count           int64
		input, output   []interface{}
		err             error
	)

	count, err = ma.getInt(ctx, ma.gov, block, "getNodeLength")
	for i := int64(1); i <= count; i++ {
		input = []interface{}{big.NewInt(i)}
		output = []interface{}{&name, &enode, &ip, &port}
		if err = metclient.CallContract(ctx, ma.gov, "getNode", input, &output, block); err != nil {
			return nil, err
		}

		if err = metclient.CallContract(ctx, ma.gov, "getMember", input, &addr, block); err != nil {
			return nil, err
		}

		sid := hex.EncodeToString(enode)
		if len(sid) != 128 {
			return nil, ErrInvalidEnode
		}
		idv4, _ := toIdv4(sid)
		nodes = append(nodes, &wemixNode{
			Name:  string(name),
			Enode: sid,
			Ip:    string(ip),
			Id:    idv4,
			Port:  int(port.Int64()),
			Addr:  addr,
		})
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Name < nodes[j].Name
	})
	return nodes, err
}

func (ma *wemixAdmin) getRewardParams(ctx context.Context, height *big.Int) (*rewardParameters, error) {
	rp := &rewardParameters{}
	reg, gov, env, staking, err := ma.getRegGovEnvContracts(ctx, height)
	if err != nil {
		return nil, err
	}

	if err = metclient.CallContract(ctx, env, "getBlockRewardAmount", nil, &rp.rewardAmount, height); err != nil {
		return nil, err
	}

	rp.distributionMethod = make([]*big.Int, 4)
	if err = metclient.CallContract(ctx, env, "getBlockRewardDistributionMethod", nil, &rp.distributionMethod, height); err != nil {
		return nil, err
	}

	var addr common.Address
	input := []interface{}{metclient.ToBytes32("StakingReward")}
	if err = metclient.CallContract(ctx, reg, "getContractAddress", input, &addr, height); err != nil {
		return nil, err
	}
	rp.staker = &common.Address{}
	rp.staker.SetBytes(addr.Bytes())

	input = []interface{}{metclient.ToBytes32("Ecosystem")}
	if err = metclient.CallContract(ctx, reg, "getContractAddress", input, &addr, height); err != nil {
		return nil, err
	}
	rp.ecoSystem = &common.Address{}
	rp.ecoSystem.SetBytes(addr.Bytes())

	input = []interface{}{metclient.ToBytes32("Maintenance")}
	if err = metclient.CallContract(ctx, reg, "getContractAddress", input, &addr, height); err != nil {
		return nil, err
	}
	rp.maintenance = &common.Address{}
	rp.maintenance.SetBytes(addr.Bytes())

	input = []interface{}{metclient.ToBytes32("FeeCollector")}
	if err = metclient.CallContract(ctx, reg, "getContractAddress", input, &addr, height); err != nil {
		// ignore error
		rp.feeCollector = nil
	} else {
		rp.feeCollector = &common.Address{}
		rp.feeCollector.SetBytes(addr.Bytes())
	}

	rp.blocksPer, err = ma.getInt(ctx, env, height, "getBlocksPer")
	if err != nil {
		return nil, err
	}

	if count, err := ma.getInt(ctx, gov, height, "getMemberLength"); err != nil {
		return nil, err
	} else {
		for i := int64(1); i <= count; i++ {
			var rewardAddress common.Address
			var stake *big.Int

			input = []interface{}{big.NewInt(i)}
			if err = metclient.CallContract(ctx, gov, "getMember", input, &addr, height); err != nil {
				return nil, err
			}
			input = []interface{}{big.NewInt(i)}
			if err = metclient.CallContract(ctx, gov, "getReward", input, &rewardAddress, height); err != nil {
				return nil, err
			}
			input = []interface{}{addr}
			if err = metclient.CallContract(ctx, staking, "lockedBalanceOf", input, &stake, height); err != nil {
				return nil, err
			}
			rp.members = append(rp.members, &wemixMember{
				Staker: addr,
				Reward: rewardAddress,
				Stake:  stake,
			})
		}
	}

	return rp, nil
}

func (ma *wemixAdmin) getRewardAccounts(ctx context.Context, block *big.Int) (rewardPoolAccount, maintenanceAccount *common.Address, members []*wemixMember, err error) {
	var (
		addr  common.Address
		count int64
		stake *big.Int
		input []interface{}
	)

	if ma.registry == nil || ma.registry.To == nil {
		err = wemixminer.ErrNotInitialized
		return
	}

	input = []interface{}{metclient.ToBytes32("RewardPool")}
	err = metclient.CallContract(ctx, ma.registry, "getContractAddress", input, &addr, block)
	if err == nil {
		rewardPoolAccount = &common.Address{}
		rewardPoolAccount.SetBytes(addr.Bytes())
	}

	input = []interface{}{metclient.ToBytes32("Maintenance")}
	err = metclient.CallContract(ctx, ma.registry, "getContractAddress", input, &addr, block)
	if err == nil {
		maintenanceAccount = &common.Address{}
		maintenanceAccount.SetBytes(addr.Bytes())
	}

	count, err = ma.getInt(ctx, ma.gov, block, "getMemberLength")
	if err != nil {
		return
	}

	for i := int64(1); i <= count; i++ {
		var rewardAddress common.Address

		input = []interface{}{big.NewInt(i)}
		err = metclient.CallContract(ctx, ma.gov, "getMember", input,
			&addr, block)
		if err != nil {
			return
		}
		err = metclient.CallContract(ctx, ma.gov, "getReward", input,
			&rewardAddress, block)
		if err != nil {
			return
		}
		input = []interface{}{addr}
		err = metclient.CallContract(ctx, ma.staking, "lockedBalanceOf", input,
			&stake, block)
		if err != nil {
			return
		}

		members = append(members, &wemixMember{
			Staker: addr,
			Reward: rewardAddress,
			Stake:  stake,
		})
	}

	return
}

// temporary internal structure to collect data from governance contracts
type govdata struct {
	blockNum, modifiedBlock                        int64
	blockInterval, blocksPer, maxIdleBlockInterval int64
	blockReward, maxPriorityFeePerGas              *big.Int
	maxBaseFee, gasLimit                           *big.Int
	baseFeeMaxChangeRate, gasTargetPercentage      int64
	nodes, addedNodes, updatedNodes, deletedNodes  []*wemixNode
}

func (ma *wemixAdmin) getGovData(refresh bool) (data *govdata, err error) {
	data = &govdata{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	block, err := ma.cli.HeaderByNumber(ctx, nil)
	if err != nil {
		return
	}
	data.blockNum = block.Number.Int64()
	if !refresh && data.blockNum <= ma.lastBlock {
		return
	}

	data.modifiedBlock, err = ma.getInt(ctx, ma.gov, block.Number,
		"modifiedBlock")
	if err != nil {
		return
	}
	if !refresh && ma.modifiedBlock == data.modifiedBlock {
		return
	}

	data.blockInterval, err = ma.getInt(ctx, ma.envStorage, block.Number, "getBlockCreationTime")
	if err != nil {
		// TODO: ignore this error for now
		data.blockInterval = ma.blockInterval
		//return
	}
	data.blocksPer, err = ma.getInt(ctx, ma.envStorage, block.Number, "getBlocksPer")
	if err != nil {
		// TODO: ignore this error for now
		data.blocksPer = ma.blocksPer
		//return
	}
	data.maxIdleBlockInterval, err = ma.getInt(ctx, ma.envStorage, block.Number, "getMaxIdleBlockInterval")
	if err != nil {
		// TODO: ignore this error for now
		data.maxIdleBlockInterval = int64(params.MaxIdleBlockInterval)
		//return
	}
	err = metclient.CallContract(ctx, ma.envStorage, "getBlockRewardAmount", nil, &data.blockReward, block.Number)
	if err != nil {
		return
	}
	err = metclient.CallContract(ctx, ma.envStorage, "getMaxPriorityFeePerGas", nil, &data.maxPriorityFeePerGas, block.Number)
	if err != nil {
		return
	}
	gasLimitAndBaseFee := make([]*big.Int, 3)
	err = metclient.CallContract(ctx, ma.envStorage, "getGasLimitAndBaseFee", nil, &gasLimitAndBaseFee, block.Number)
	if err != nil {
		return
	}
	data.gasLimit = gasLimitAndBaseFee[0]
	data.baseFeeMaxChangeRate = gasLimitAndBaseFee[1].Int64()
	data.gasTargetPercentage = gasLimitAndBaseFee[2].Int64()

	err = metclient.CallContract(ctx, ma.envStorage, "getMaxBaseFee", nil, &data.maxBaseFee, block.Number)
	if err != nil {
		return
	}

	data.nodes, err = ma.getWemixNodes(ctx, block.Number)
	if err != nil {
		return
	}

	oldNodes := ma.getNodes()
	sort.Slice(oldNodes, func(i, j int) bool {
		return oldNodes[i].Name < oldNodes[j].Name
	})
	sort.Slice(data.nodes, func(i, j int) bool {
		return data.nodes[i].Name < data.nodes[j].Name
	})

	i, j := 0, 0
	for {
		if i >= len(oldNodes) || j >= len(data.nodes) {
			break
		}
		v := strings.Compare(oldNodes[i].Name, data.nodes[j].Name)
		if v == 0 {
			if !oldNodes[i].eq(data.nodes[j]) {
				data.updatedNodes = append(data.updatedNodes, data.nodes[j])
			}
			i++
			j++
		} else if v < 0 {
			data.deletedNodes = append(data.deletedNodes, oldNodes[i])
			i++
		} else if v > 0 {
			data.addedNodes = append(data.addedNodes, data.nodes[j])
			j++
		}
	}

	if i < len(oldNodes) {
		for ; i < len(oldNodes); i++ {
			data.deletedNodes = append(data.deletedNodes, oldNodes[i])
		}
	}
	if j < len(data.nodes) {
		for ; j < len(data.nodes); j++ {
			data.addedNodes = append(data.addedNodes, data.nodes[j])
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

	rpcCli, err := stack.Attach()
	if err != nil {
		utils.Fatalf("Failed to attach to self: %v", err)
	}

	registryContract, err := metclient.LoadJsonContract(strings.NewReader(RegistryAbi))
	if err != nil {
		utils.Fatalf("Loading ABI failed: %v", err)
	}
	govContract, err := metclient.LoadJsonContract(strings.NewReader(GovAbi))
	if err != nil {
		utils.Fatalf("Loading ABI failed: %v", err)
	}
	stakingContract, err := metclient.LoadJsonContract(strings.NewReader(StakingAbi))
	if err != nil {
		utils.Fatalf("Loading ABI failed: %v", err)
	}
	envStorageImpContract, err := metclient.LoadJsonContract(strings.NewReader(EnvStorageImpAbi))
	if err != nil {
		utils.Fatalf("Loading ABI failed: %v", err)
	}

	cli := ethclient.NewClient(rpcCli)
	admin = &wemixAdmin{
		stack: stack,
		lock:  &sync.Mutex{},
		registry: &metclient.RemoteContract{
			Cli: cli, Abi: registryContract.Abi},
		gov: &metclient.RemoteContract{
			Cli: cli, Abi: govContract.Abi},
		staking: &metclient.RemoteContract{
			Cli: cli, Abi: stakingContract.Abi},
		envStorage: &metclient.RemoteContract{
			Cli: cli, Abi: envStorageImpContract.Abi},
		Updates:     make(chan bool, 10),
		rpcCli:      rpcCli,
		cli:         cli,
		blocksPer:   100,
		etcdDir:     path.Join(datadir, "etcd"),
		etcdTimeout: 30 * time.Second,
	}

	admin.bootNodeId, admin.bootAccount, err = admin.getGenesisInfo()
	if err != nil {
		return
	}

	go admin.run()
	go admin.handleNewBlocks()
	go func() {
		for {
			time.Sleep(time.Duration(SyncIdleThreshold/2) * time.Second)
			syncCheck()
		}
	}()
}

func (ma *wemixAdmin) addPeer(node *wemixNode) error {
	if node.Id == ma.nodeInfo.ID || ma.self == nil {
		return nil
	}

	var v *bool
	ctx, cancel := context.WithCancel(context.Background())
	id := fmt.Sprintf("enode://%s@%s:%d", node.Enode, node.Ip, node.Port)
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

func (ma *wemixAdmin) update() {
	if ma.registry == nil || ma.registry.To == nil {
		return
	}

	refresh := false
	registry, gov, staking, envStorage, err := ma.getAdminAddresses()
	if err != nil {
		return
	} else if !bytes.Equal(registry[:], ma.registry.To[:]) ||
		!bytes.Equal(gov[:], ma.gov.To[:]) ||
		!bytes.Equal(staking[:], ma.staking.To[:]) ||
		!bytes.Equal(envStorage[:], ma.envStorage.To[:]) {
		ma.registry.To = registry
		ma.gov.To = gov
		ma.staking.To = staking
		ma.envStorage.To = envStorage
		refresh = true
	}

	data, err := ma.getGovData(refresh)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to get nodes: %v", err))
	} else if refresh ||
		(data.modifiedBlock != 0 && ma.modifiedBlock != data.modifiedBlock) {
		log.Debug(fmt.Sprintf("Modified Block: %d", data.modifiedBlock))

		ma.modifiedBlock = data.modifiedBlock
		ma.blockInterval = data.blockInterval
		ma.blocksPer = data.blocksPer
		ma.blockReward = data.blockReward
		ma.maxPriorityFeePerGas = data.maxPriorityFeePerGas
		ma.maxBaseFee = data.maxBaseFee
		ma.gasLimit = data.gasLimit
		ma.baseFeeMaxChangeRate = data.baseFeeMaxChangeRate
		ma.gasTargetPercentage = data.gasTargetPercentage

		_nodes := map[string]*wemixNode{}
		for _, i := range data.nodes {
			_nodes[i.Id] = i
			if i.Id == ma.nodeInfo.ID {
				ma.self = i
			}
		}
		ma.nodes = _nodes

		if len(data.addedNodes) > 0 {
			log.Debug("Added:\n")
			for _, i := range data.addedNodes {
				log.Debug(fmt.Sprintf("%v\n", i))
				ma.addPeer(i)
			}
		}
		if len(data.addedNodes) > 0 {
			log.Debug("Updated:\n")
			for _, i := range data.updatedNodes {
				log.Debug(fmt.Sprintf("%v\n", i))
			}
		}
		if len(data.addedNodes) > 0 {
			log.Debug("Deleted:\n")
			for _, i := range data.deletedNodes {
				log.Debug(fmt.Sprintf("%v\n", i))
			}
		}

		if params.MaxIdleBlockInterval != uint64(data.maxIdleBlockInterval) {
			params.MaxIdleBlockInterval = uint64(data.maxIdleBlockInterval)
		}

		// set coinbase and minimum gas price
		setGasCoinbase := func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var v *bool
			err := ma.rpcCli.CallContext(ctx, &v, "miner_setGasPrice",
				"0x"+data.maxPriorityFeePerGas.Text(16))
			if err != nil || !*v {
				log.Info("set minimum gas price failed",
					"maxPriorityFeePerGas", data.maxPriorityFeePerGas, "error", err)
			} else {
				log.Info("successfully set",
					"maxPriorityFeePerGas", data.maxPriorityFeePerGas)
			}

			if ma.self != nil && !bytes.Equal(ma.self.Addr[:], nilAddress[:]) {
				err = ma.rpcCli.CallContext(ctx, &v, "miner_setEtherbase", &ma.self.Addr)
				if err != nil || !*v {
					log.Info("set the coinbase", "error", err)
				} else {
					log.Info("successfully set the coinbase")
				}
			}
		}
		setGasCoinbase()
	}

	if data.blockNum != 0 {
		ma.lastBlock = data.blockNum
	}
}

func (ma *wemixAdmin) checkMining() {
	on := false
	if ma.nodeInfo != nil && ma.nodeInfo.ID == admin.bootNodeId {
		on = true
	} else if ma.self != nil {
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
	if mining != nil && !*mining {
		// in case we're leader, transfer leadership
		ma.etcdTransferLeadership()
	}
}

func (ma *wemixAdmin) run() {
	lt := time.Now()
	for {
		if ma.nodeInfo == nil {
			nodeInfo, err := ma.getNodeInfo()
			if err != nil {
				log.Error("Failed to get node info", "error", err)
			} else {
				ma.nodeInfo = nodeInfo
			}
		}
		if ma.registry.To == nil {
			registry, gov, staking, envStorage, err := ma.getAdminAddresses()
			if err == nil {
				ma.registry.To = registry
				ma.gov.To = gov
				ma.staking.To = staking
				ma.envStorage.To = envStorage
			}
		}
		if ma.registry.To != nil && ma.nodeInfo != nil {
			ma.update()
			if ma.amPartner() && ma.self != nil && !ma.etcdIsRunning() {
				EtcdStart()
			}
		}

		if ma.amPartner() {
			ma.checkMining()

			t := time.Now()
			if t.Sub(lt).Seconds() >= 30 {
				lt = t
				nodes := ma.getNodes()
				for _, n := range nodes {
					if !ma.isPeerUp(n.Id) {
						ma.addPeer(n)
					}
				}
			}
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

type reward struct {
	Addr   common.Address `json:"addr"`
	Reward *big.Int       `json:"reward"`
}

func (ma *wemixAdmin) verifyRewards(r1, r2 []byte) error {
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

// handles rewards in testnet block 94
func handleBlock94Rewards(height *big.Int, rp *rewardParameters, fees *big.Int) []reward {
	if height.Int64() != 94 || len(rp.members) != 0 ||
		!bytes.Equal(rp.staker[:], testnetBlock94Rewards[0].Addr[:]) ||
		!bytes.Equal(rp.ecoSystem[:], testnetBlock94Rewards[1].Addr[:]) ||
		!bytes.Equal(rp.maintenance[:], testnetBlock94Rewards[2].Addr[:]) {
		return nil
	}
	return testnetBlock94Rewards
}

// distributeRewards divides the rewardAmount among members according to their
// stakes, and allocates rewards to staker, ecoSystem, and maintenance accounts.
func distributeRewards(height *big.Int, rp *rewardParameters, fees *big.Int) ([]reward, error) {
	dm := new(big.Int)
	for i := 0; i < len(rp.distributionMethod); i++ {
		dm.Add(dm, rp.distributionMethod[i])
	}
	if dm.Int64() != 10000 {
		return nil, wemixminer.ErrNotInitialized
	}

	v10000 := big.NewInt(10000)
	minerAmount := new(big.Int).Set(rp.rewardAmount)
	minerAmount.Div(minerAmount.Mul(minerAmount, rp.distributionMethod[0]), v10000)
	stakerAmount := new(big.Int).Set(rp.rewardAmount)
	stakerAmount.Div(stakerAmount.Mul(stakerAmount, rp.distributionMethod[1]), v10000)
	ecoSystemAmount := new(big.Int).Set(rp.rewardAmount)
	ecoSystemAmount.Div(ecoSystemAmount.Mul(ecoSystemAmount, rp.distributionMethod[2]), v10000)
	// the rest goes to maintenance
	maintenanceAmount := new(big.Int).Set(rp.rewardAmount)
	maintenanceAmount.Sub(maintenanceAmount, minerAmount)
	maintenanceAmount.Sub(maintenanceAmount, stakerAmount)
	maintenanceAmount.Sub(maintenanceAmount, ecoSystemAmount)

	// if feeCollector is not specified, i.e. nil, fees go to maintenance
	if rp.feeCollector == nil {
		maintenanceAmount.Add(maintenanceAmount, fees)
	}

	var rewards []reward
	if n := len(rp.members); n > 0 {
		stakeTotal, equalStakes := big.NewInt(0), true
		for i := 0; i < n; i++ {
			if equalStakes && i < n-1 && rp.members[i].Stake.Cmp(rp.members[i+1].Stake) != 0 {
				equalStakes = false
			}
			stakeTotal.Add(stakeTotal, rp.members[i].Stake)
		}

		if equalStakes {
			v0, v1 := big.NewInt(0), big.NewInt(1)
			vn := big.NewInt(int64(n))
			b := new(big.Int).Set(minerAmount)
			d := new(big.Int)
			d.Div(b, vn)
			for i := 0; i < n; i++ {
				rewards = append(rewards, reward{
					Addr:   rp.members[i].Reward,
					Reward: new(big.Int).Set(d),
				})
			}
			d.Mul(d, vn)
			b.Sub(b, d)
			for i, ix := 0, height.Int64()%int64(n); b.Cmp(v0) > 0; i, ix = i+1, (ix+1)%int64(n) {
				rewards[ix].Reward.Add(rewards[ix].Reward, v1)
				b.Sub(b, v1)
			}
		} else {
			// rewards distributed according to stakes
			v0, v1 := big.NewInt(0), big.NewInt(1)
			remainder := new(big.Int).Set(minerAmount)
			for i := 0; i < n; i++ {
				memberReward := new(big.Int).Mul(minerAmount, rp.members[i].Stake)
				memberReward.Div(memberReward, stakeTotal)
				remainder.Sub(remainder, memberReward)
				rewards = append(rewards, reward{
					Addr:   rp.members[i].Reward,
					Reward: memberReward,
				})
			}
			for ix := height.Int64() % int64(n); remainder.Cmp(v0) > 0; ix = (ix + 1) % int64(n) {
				rewards[ix].Reward.Add(rewards[ix].Reward, v1)
				remainder.Sub(remainder, v1)
			}
		}
	}
	rewards = append(rewards, reward{
		Addr:   *rp.staker,
		Reward: stakerAmount,
	})
	rewards = append(rewards, reward{
		Addr:   *rp.ecoSystem,
		Reward: ecoSystemAmount,
	})
	rewards = append(rewards, reward{
		Addr:   *rp.maintenance,
		Reward: maintenanceAmount,
	})
	if rp.feeCollector != nil {
		rewards = append(rewards, reward{
			Addr:   *rp.feeCollector,
			Reward: fees,
		})
	}
	return rewards, nil
}

func (ma *wemixAdmin) calculateRewards(num, blockReward, fees *big.Int, addBalance func(common.Address, *big.Int)) (coinbase *common.Address, rewards []byte, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rp, err := ma.getRewardParams(ctx, big.NewInt(num.Int64()-1))
	if err != nil {
		// all goes to the coinbase
		err = wemixminer.ErrNotInitialized
		return
	}

	if (rp.staker == nil && rp.ecoSystem == nil && rp.maintenance == nil) || len(rp.members) == 0 {
		// handle testnet block 94 rewards
		if rewards94 := handleBlock94Rewards(num, rp, fees); rewards94 != nil {
			if addBalance != nil {
				for _, i := range rewards94 {
					addBalance(i.Addr, i.Reward)
				}
			}
			rewards, err = json.Marshal(rewards94)
			return
		}
		err = wemixminer.ErrNotInitialized
		return
	}

	// determine coinbase
	if len(rp.members) > 0 {
		mix := int(num.Int64()/ma.blocksPer) % len(rp.members)
		coinbase = &common.Address{}
		coinbase.SetBytes(rp.members[mix].Reward.Bytes())
	}

	rr, errr := distributeRewards(num, rp, fees)
	if errr != nil {
		err = errr
		return
	}

	if addBalance != nil {
		for _, i := range rr {
			addBalance(i.Addr, i.Reward)
		}
	}

	rewards, err = json.Marshal(rr)
	return
}

func calculateRewards(num, blockReward, fees *big.Int, addBalance func(common.Address, *big.Int)) (*common.Address, []byte, error) {
	return admin.calculateRewards(num, blockReward, fees, addBalance)
}

func verifyRewards(num *big.Int, rewards string) error {
	return nil
	//return admin.verifyRewards(num, rewards)
}

func signBlock(height *big.Int, hash common.Hash) (coinbase common.Address, sig []byte, err error) {
	if admin == nil {
		err = wemixminer.ErrNotInitialized
		return
	}
	data := append(height.Bytes(), hash.Bytes()...)
	data = crypto.Keccak256(data)
	prvKey := admin.stack.Server().PrivateKey
	sig, err = crypto.Sign(data, prvKey)
	if admin.self != nil {
		coinbase = admin.self.Addr
	} else if admin.nodeInfo != nil && admin.nodeInfo.ID == admin.bootNodeId {
		coinbase = admin.bootAccount
	}
	return
}

func verifyBlockSig(height *big.Int, coinbase common.Address, nodeId []byte, hash common.Hash, sig []byte, checkMinerLimit bool) bool {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get nodeid from the coinbase
	num := new(big.Int).Sub(height, common.Big1)
	_, gov, _, _, err := admin.getRegGovEnvContracts(ctx, num)
	if err != nil {
		return err == wemixminer.ErrNotInitialized
	} else if count, err := admin.getInt(ctx, gov, num, "getMemberLength"); err != nil || count == 0 {
		return err == wemixminer.ErrNotInitialized || count == 0
	}
	// if minerNodeId is given, i.e. present in block header, use it,
	// otherwise, derive it from the codebase
	var data []byte
	if len(nodeId) == 0 {
		nodeId, err = coinbaseExists(ctx, height, gov, &coinbase)
		if err != nil || len(nodeId) == 0 {
			return false
		}
		data = append(height.Bytes(), hash.Bytes()...)
		data = crypto.Keccak256(data)
	} else {
		if ok, err := enodeExists(ctx, height, gov, nodeId); err != nil || !ok {
			return false
		}
		data = hash.Bytes()
	}
	pubKey, err := crypto.Ecrecover(data, sig)
	if err != nil || len(pubKey) < 1 || !bytes.Equal(nodeId, pubKey[1:]) {
		return false
	}
	// check miner limit
	if !checkMinerLimit {
		return true
	}
	ok, err := admin.verifyMinerLimit(ctx, height, gov, &coinbase, nodeId)
	return err == nil && ok
}

func (ma *wemixAdmin) getNodeInfo() (*p2p.NodeInfo, error) {
	var nodeInfo *p2p.NodeInfo
	ctx, cancel := context.WithCancel(context.Background())
	err := ma.rpcCli.CallContext(ctx, &nodeInfo, "admin_nodeInfo")
	cancel()
	if err != nil {
		log.Error("Cannot get node info", "error", err)
	}
	return nodeInfo, err
}

func (ma *wemixAdmin) getPeerInfo(id string) (*p2p.NodeInfo, error) {
	var nodeInfo *p2p.NodeInfo
	ctx, cancel := context.WithCancel(context.Background())
	err := ma.rpcCli.CallContext(ctx, &nodeInfo, "admin_peerInfo", id)
	cancel()
	if err != nil {
		log.Error("Cannot get peer info", "id", id, "error", err)
	}
	return nodeInfo, err
}

func (ma *wemixAdmin) isPeerUp(id string) bool {
	nodeInfo, err := ma.getPeerInfo(id)
	return err == nil && nodeInfo != nil
}

func (ma *wemixAdmin) amPartner() bool {
	if admin == nil {
		return false
	}
	return admin.self != nil || (admin.nodeInfo != nil && admin.nodeInfo.ID == admin.bootNodeId)
}

func AmPartner() bool {
	if admin == nil {
		return false
	}

	admin.lock.Lock()
	defer admin.lock.Unlock()

	return admin.amPartner()
}

// id is v4 id
func IsPartner(id string) bool {
	if admin == nil {
		return false
	}

	admin.lock.Lock()
	defer admin.lock.Unlock()

	_, ok := admin.nodes[id]
	if !ok {
		if id == admin.bootNodeId {
			return true
		} else {
			return false
		}
	}

	return true
}

// id is v4 id
func AmHub(id string) int {
	if admin == nil || admin.self == nil {
		return -1
	}

	admin.lock.Lock()
	defer admin.lock.Unlock()
	if strings.HasPrefix(strings.ToUpper(admin.self.Id), strings.ToUpper(id)) {
		return 1
	} else {
		return 0
	}
}

func (ma *wemixAdmin) pendingEmpty() bool {
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

func getMaxPriorityFeePerGas() *big.Int {
	defaultFee := big.NewInt(100 * params.GWei)
	if admin == nil || admin.envStorage == nil || admin.envStorage.To == nil {
		return defaultFee
	}
	var fee *big.Int
	if err := metclient.CallContract(context.Background(), admin.envStorage, "getMaxPriorityFeePerGas", nil, &fee, nil); err != nil {
		return defaultFee
	}
	return fee
}

func suggestGasPrice() *big.Int {
	defaultFee := big.NewInt(100 * params.GWei)
	if admin == nil || admin.envStorage == nil || admin.envStorage.To == nil {
		return defaultFee
	}
	var fee *big.Int
	if err := metclient.CallContract(context.Background(), admin.envStorage, "getMaxPriorityFeePerGas", nil, &fee, nil); err != nil {
		return defaultFee
	}
	return fee
}

func getBlockBuildParameters(height *big.Int) (blockInterval int64, maxBaseFee, gasLimit *big.Int, baseFeeMaxChangeRate, gasTargetPercentage int64, err error) {
	err = wemixminer.ErrNotInitialized

	blockBuildParamsLock.Lock()
	if blockBuildParams != nil && blockBuildParams.height == height.Uint64() {
		// use chached
		blockInterval = blockBuildParams.blockInterval
		maxBaseFee = blockBuildParams.maxBaseFee
		gasLimit = blockBuildParams.gasLimit
		baseFeeMaxChangeRate = blockBuildParams.baseFeeMaxChangeRate
		gasTargetPercentage = blockBuildParams.gasTargetPercentage
		blockBuildParamsLock.Unlock()
		err = nil
		return
	}
	blockBuildParamsLock.Unlock()

	// default values
	blockInterval = 15
	maxBaseFee = big.NewInt(0)
	gasLimit = big.NewInt(0)
	baseFeeMaxChangeRate = 0
	gasTargetPercentage = 100

	if admin == nil {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var env, gov *metclient.RemoteContract
	if _, gov, env, _, err = admin.getRegGovEnvContracts(ctx, height); err != nil {
		err = wemixminer.ErrNotInitialized
		return
	} else if count, err2 := admin.getInt(ctx, gov, height, "getMemberLength"); err2 != nil || count == 0 {
		err = wemixminer.ErrNotInitialized
		return
	}
	var v *big.Int
	if err = metclient.CallContract(ctx, env, "getBlockCreationTime", nil, &v, height); err != nil {
		err = wemixminer.ErrNotInitialized
		return
	}
	blockInterval = v.Int64()

	gasLimitAndBaseFee := make([]*big.Int, 3)
	if err = metclient.CallContract(ctx, env, "getGasLimitAndBaseFee", nil, &gasLimitAndBaseFee, height); err != nil {
		err = wemixminer.ErrNotInitialized
		return
	}
	gasLimit = gasLimitAndBaseFee[0]
	baseFeeMaxChangeRate = gasLimitAndBaseFee[1].Int64()
	gasTargetPercentage = gasLimitAndBaseFee[2].Int64()

	if err = metclient.CallContract(ctx, env, "getMaxBaseFee", nil, &maxBaseFee, height); err != nil {
		err = wemixminer.ErrNotInitialized
		return
	}

	// cache it
	blockBuildParamsLock.Lock()
	blockBuildParams = &blockBuildParameters{
		height:               height.Uint64(),
		blockInterval:        blockInterval,
		maxBaseFee:           maxBaseFee,
		gasLimit:             gasLimit,
		baseFeeMaxChangeRate: baseFeeMaxChangeRate,
		gasTargetPercentage:  gasTargetPercentage,
	}
	blockBuildParamsLock.Unlock()
	err = nil
	return
}

func (ma *wemixAdmin) toMiningPeers(nodes []*wemixNode) string {
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
	return bb.String()
}

func (ma *wemixAdmin) miners() string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	block, err := ma.cli.HeaderByNumber(ctx, nil)
	if err != nil {
		return ""
	}
	height := block.Number.Int64()

	_, _, nodes := ma.getMinerNodes(height+1, false)
	return ma.toMiningPeers(nodes)
}

func Info() interface{} {
	if admin == nil {
		return ""
	} else {
		self := admin.self
		var nodes []*wemixNode
		for _, i := range admin.nodes {
			nodes = append(nodes, i)
		}
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].Name < nodes[j].Name
		})

		info := &map[string]interface{}{
			"consensus":            params.ConsensusMethod,
			"registry":             admin.registry.To,
			"governance":           admin.gov.To,
			"staking":              admin.staking.To,
			"modifiedblock":        admin.modifiedBlock,
			"blocksPer":            admin.blocksPer,
			"blockInterval":        admin.blockInterval,
			"blockReward":          admin.blockReward,
			"maxPriorityFeePerGas": admin.maxPriorityFeePerGas,
			"blockGasLimit":        admin.gasLimit,
			"maxBaseFee":           admin.maxBaseFee,
			"baseFeeMaxChangeRate": admin.baseFeeMaxChangeRate,
			"gasTargetPercentage":  admin.gasTargetPercentage,
			"self":                 self,
			"nodes":                nodes,
			"miners":               admin.miners(),
			"etcd":                 admin.etcdInfo(),
			"maxIdle":              params.MaxIdleBlockInterval,
		}
		return info
	}
}

func getMinerStatus() *wemixapi.WemixMinerStatus {
	if admin == nil || admin.self == nil {
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	header, err := admin.cli.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil
	}
	height := header.Number.Int64()

	_, _, nodes := admin.getMinerNodes(height+1, false)
	miningPeers := admin.toMiningPeers(nodes)

	admin.lock.Lock()
	defer admin.lock.Unlock()

	return &wemixapi.WemixMinerStatus{
		NodeName:          admin.self.Name,
		Enode:             admin.self.Enode,
		Id:                admin.self.Id,
		Addr:              fmt.Sprintf("%s:%d", admin.self.Ip, admin.self.Port),
		Status:            "up",
		Miner:             admin.self.Miner,
		MiningPeers:       miningPeers,
		LatestBlockHeight: header.Number,
		LatestBlockHash:   header.Hash(),
		RttMs:             big0,
	}
}

// Returns the array of peer status
// 'id' could be null, a name, node id (public key) or ip address of a miner
func getMiners(id string, timeout int) []*wemixapi.WemixMinerStatus {
	if admin == nil {
		return nil
	}

	if timeout <= 0 {
		timeout = 5
	} else if timeout > 60 {
		timeout = 60
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodes := admin.getNodes()

	var node *wemixNode
	for _, n := range nodes {
		if strings.EqualFold(n.Name, id) || strings.EqualFold(n.Id, id) || strings.EqualFold(n.Ip, id) {
			node = n
			break
		}
	}

	getDownStatus := func(node *wemixNode) *wemixapi.WemixMinerStatus {
		return &wemixapi.WemixMinerStatus{
			NodeName: node.Name,
			Enode:    node.Enode,
			Id:       node.Id,
			Addr:     fmt.Sprintf("%s:%d", node.Ip, node.Port),
			Status:   "down",
			RttMs:    big0,
		}
	}

	var miners []*wemixapi.WemixMinerStatus
	var err error
	ch := make(chan *wemixapi.WemixMinerStatus, len(nodes)*2+1)
	sub := wemixapi.SubscribeToMinerStatus(ch)
	defer func() {
		sub.Unsubscribe()
		close(ch)
	}()

	startTime := time.Now().UnixNano()
	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	peers := map[string]*wemixNode{}
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
			log.Error("RequestMinerStatus Failed", "id", node.Id, "error", err)
			status := getDownStatus(node)
			status.RttMs = big.NewInt((time.Now().UnixNano() - startTime) / 1000000)
			miners = append(miners, status)
		} else {
			peers[node.Name] = node
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
				status := getDownStatus(n)
				status.RttMs = big.NewInt((time.Now().UnixNano() - startTime) / 1000000)
				miners = append(miners, status)
				log.Error("RequestMinerStatus Failed", "id", n.Id, "error", err)
			} else {
				peers[n.Name] = n
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
		case status := <-ch:
			if done {
				continue
			}
			if n, exists := peers[status.NodeName]; exists {
				status.RttMs = big.NewInt((time.Now().UnixNano() - startTime) / 1000000)
				miners = append(miners, status)
				if n != nil {
					peers[status.NodeName] = nil
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

	for _, n := range peers {
		if n != nil {
			status := getDownStatus(n)
			status.RttMs = big.NewInt((time.Now().UnixNano() - startTime) / 1000000)
			miners = append(miners, status)
		}
	}

	if len(miners) > 1 {
		sort.Slice(miners, func(i, j int) bool {
			return miners[i].NodeName < miners[j].NodeName
		})
	}
	return miners
}

func (ma *wemixAdmin) getTxPoolStatus() (pending, queued uint, err error) {
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
	p, _, e := admin.getTxPoolStatus()
	if e != nil {
		return false
	} else if p > 0 {
		return false
	}

	return true
}

// checks
//  1. fees total and per governance accounts are accurate
//  2. sum(rewards) == fees + block reward
//  3. rewards distribution is correct
//  4. reward members, reward pool and maintenance account are correct
//  5. balances of governance accounts are accurate.
//     Note that it doesn't take account of internal transactions,
//     so balance checks won't be accurate if there are contract transactions.
func verifyBlockRewards(height *big.Int) interface{} {
	type result struct {
		Status bool `json:"status"`
		// txs counts: total, contract calls and simple ether transfers
		Txs         int `json:"txs"` // # of txs
		ContractTxs int `json:"contractTxs"`
		SimpleTxs   int `json:"simpleTxs"`
		// this will be 0 for now
		BlockReward *big.Int `json:"blockReward"`
		// fees: total and per accounts in governance contract
		Fees map[string]*big.Int `json:"fees"`
		// error and messsages if any
		Error   string `json:"error"`
		Message string `json:"message"`
	}

	r := &result{
		Status: false,
		Error:  "Not initialized",
	}

	if admin == nil {
		return r
	}

	return r
}

func init() {
	wemixminer.AmPartnerFunc = AmPartner
	wemixminer.IsPartnerFunc = IsPartner
	wemixminer.AmHubFunc = AmHub
	wemixminer.SuggestGasPriceFunc = suggestGasPrice
	wemixminer.CalculateRewardsFunc = calculateRewards
	wemixminer.VerifyRewardsFunc = verifyRewards
	wemixminer.SignBlockFunc = signBlock
	wemixminer.VerifyBlockSigFunc = verifyBlockSig
	wemixminer.RequirePendingTxsFunc = requirePendingTxs
	wemixminer.VerifyBlockRewardsFunc = verifyBlockRewards
	wemixminer.GetBlockBuildParametersFunc = getBlockBuildParameters
	wemixminer.AcquireMiningTokenFunc = acquireMiningToken
	wemixminer.ReleaseMiningTokenFunc = releaseMiningToken
	wemixminer.HasMiningTokenFunc = hasMiningToken
	wemixapi.Info = Info
	wemixapi.GetMiners = getMiners
	wemixapi.GetMinerStatus = getMinerStatus
	wemixapi.EtcdInit = EtcdInit
	wemixapi.EtcdAddMember = EtcdAddMember
	wemixapi.EtcdRemoveMember = EtcdRemoveMember
	wemixapi.EtcdJoin = EtcdJoin
	wemixapi.EtcdMoveLeader = EtcdMoveLeader
	wemixapi.EtcdGetWork = EtcdGetWork
	wemixapi.EtcdDeleteWork = EtcdDeleteWork
	wemixapi.EtcdGet = EtcdGet
	wemixapi.EtcdPut = EtcdPut
	wemixapi.EtcdDelete = EtcdDelete

	// handle testnet block 94 rewards
	if err := json.Unmarshal([]byte(testnetBlock94RewardsString), &testnetBlock94Rewards); err != nil {
		panic("failed to unmarshal testnet block 94 rewards")
	}

	// handle mining peers' status update
	go handleMinerStatusUpdate()
}

/* EOF */
