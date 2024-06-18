package test

import (
	"math/big"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/wemix/governance-contract/common/compile"
	sim "github.com/ethereum/go-ethereum/wemix/governance-contract/common/simulated-backend"
)

type Governance struct {
	client *sim.Client
	Compiled
	nodeInfos []nodeInfo
}

type nodeInfo struct {
	name  []byte
	enode []byte
	ip    []byte
	port  *big.Int
}

func DeployGovernance(t *testing.T) *Governance {
	gov := compiled.Copy(&Governance{client: sim.NewClient(t), nodeInfos: make([]nodeInfo, 0)})
	// deploy Registry
	gov.Registry.Deploy(t, gov.client)
	// deploy impl
	gov.GovImp.Deploy(t, gov.client)
	gov.NCPExitImp.Deploy(t, gov.client)
	gov.StakingImp.Deploy(t, gov.client)
	gov.BallotStorageImp.Deploy(t, gov.client)
	gov.EnvStorageImp.Deploy(t, gov.client)
	// deploy proxy
	gov.Gov.Deploy(t, gov.client, gov.GovImp.Address)
	gov.NCPExit.Deploy(t, gov.client, gov.NCPExitImp.Address)
	gov.Staking.Deploy(t, gov.client, gov.StakingImp.Address)
	gov.BallotStorage.Deploy(t, gov.client, gov.BallotStorageImp.Address)
	gov.EnvStorage.Deploy(t, gov.client, gov.EnvStorageImp.Address)
	// set Domains
	gov.Registry.ExecuteOk(t, gov.client.Owner, "setContractDomain", sim.StringToBytes32("GovernanceContract"), gov.Gov.Address)
	gov.Registry.ExecuteOk(t, gov.client.Owner, "setContractDomain", sim.StringToBytes32("Staking"), gov.Staking.Address)
	gov.Registry.ExecuteOk(t, gov.client.Owner, "setContractDomain", sim.StringToBytes32("EnvStorage"), gov.EnvStorage.Address)
	gov.Registry.ExecuteOk(t, gov.client.Owner, "setContractDomain", sim.StringToBytes32("BallotStorage"), gov.BallotStorage.Address)
	// initialize
	gov.NCPExit.ExecuteOk(t, gov.client.Owner, "initialize", gov.Registry.Address)
	gov.Staking.ExecuteOk(t, gov.client.Owner, "init", gov.Registry.Address, []byte{})
	gov.BallotStorage.ExecuteOk(t, gov.client.Owner, "initialize", gov.Registry.Address)
	envNames, envValues := makeEnvParams(
		EnvConstants.BLOCKS_PER,
		EnvConstants.BALLOT_DURATION_MIN,
		EnvConstants.BALLOT_DURATION_MAX,
		EnvConstants.STAKING_MIN,
		EnvConstants.STAKING_MAX,
		EnvConstants.MAX_IDLE_BLOCK_INTERVAL,
		EnvConstants.BLOCK_CREATION_TIME,
		EnvConstants.BLOCK_REWARD_AMOUNT,
		EnvConstants.MAX_PRIORITY_FEE_PER_GAS,
		EnvConstants.BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER,
		EnvConstants.BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD,
		EnvConstants.BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM,
		EnvConstants.BLOCK_REWARD_DISTRIBUTION_MAINTENANCE,
		EnvConstants.MAX_BASE_FEE,
		EnvConstants.BLOCK_GASLIMIT,
		EnvConstants.BASE_FEE_MAX_CHANGE_RATE,
		EnvConstants.GAS_TARGET_PERCENTAGE,
	)
	gov.EnvStorage.ExecuteOk(t, gov.client.Owner, "initialize", gov.Registry.Address, envNames, envValues)
	gov.Staking.ExecuteWithETHOk(t, gov.client.Owner, LOCK_AMOUNT, "deposit")
	node := nodeInfo{
		[]byte("name"),
		hexutil.MustDecode("0x6f8a80d14311c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
		[]byte("127.0.0.1"),
		big.NewInt(8542),
	}
	gov.Gov.ExecuteOk(t, gov.client.Owner, "init", gov.Registry.Address, LOCK_AMOUNT, node.name, node.enode, node.ip, node.port)
	gov.nodeInfos = append(gov.nodeInfos, node)

	return gov
}

var (
	compiled Compiled
)

func init() {
	compiled.Compile("../contracts")
}

type Compiled struct {
	Registry,
	Gov, GovImp,
	NCPExit, NCPExitImp,
	Staking, StakingImp,
	BallotStorage, BallotStorageImp,
	EnvStorage, EnvStorageImp *sim.Contract
}

func (c *Compiled) Compile(root string) {
	if contracts, err := compile.Compile("",
		filepath.Join(root, "Registry.sol"),
		filepath.Join(root, "Gov.sol"),
		filepath.Join(root, "GovImp.sol"),
		filepath.Join(root, "NCPExit.sol"),
		filepath.Join(root, "NCPExitImp.sol"),
		filepath.Join(root, "Staking.sol"),
		filepath.Join(root, "StakingImp.sol"),
		filepath.Join(root, "storage", "BallotStorage.sol"),
		filepath.Join(root, "storage", "BallotStorageImp.sol"),
		filepath.Join(root, "storage", "EnvStorage.sol"),
		filepath.Join(root, "storage", "EnvStorageImp.sol"),
	); err != nil {
		panic(err)
	} else if c.Registry, err = sim.NewContract("Registry", "Registry", contracts); err != nil {
		panic(err)
	} else if c.Gov, err = sim.NewContract("Gov", "Gov", contracts); err != nil {
		panic(err)
	} else if c.GovImp, err = sim.NewContract("GovImp", "GovImp", contracts); err != nil {
		panic(err)
	} else if c.NCPExit, err = sim.NewContract("NCPExit", "NCPExit", contracts); err != nil {
		panic(err)
	} else if c.NCPExitImp, err = sim.NewContract("NCPExitImp", "NCPExitImp", contracts); err != nil {
		panic(err)
	} else if c.Staking, err = sim.NewContract("Staking", "Staking", contracts); err != nil {
		panic(err)
	} else if c.StakingImp, err = sim.NewContract("StakingImp", "StakingImp", contracts); err != nil {
		panic(err)
	} else if c.BallotStorage, err = sim.NewContract("BallotStorage", "BallotStorage", contracts); err != nil {
		panic(err)
	} else if c.BallotStorageImp, err = sim.NewContract("BallotStorageImp", "BallotStorageImp", contracts); err != nil {
		panic(err)
	} else if c.EnvStorage, err = sim.NewContract("EnvStorage", "EnvStorage", contracts); err != nil {
		panic(err)
	} else if c.EnvStorageImp, err = sim.NewContract("EnvStorageImp", "EnvStorageImp", contracts); err != nil {
		panic(err)
	}
}

func (c *Compiled) Copy(g *Governance) *Governance {
	typ := reflect.TypeOf(c).Elem()
	for i := 0; i < typ.NumField(); i++ {
		if v := reflect.ValueOf(g).Elem().FieldByName(typ.Field(i).Name); v.IsValid() {
			v.Set(reflect.ValueOf(c).Elem().Field(i))
		}
	}
	return g
}

var EnvTypes = struct {
	Invalid *big.Int
	Int     *big.Int
	Uint    *big.Int
	Address *big.Int
	Bytes32 *big.Int
	Bytes   *big.Int
	String  *big.Int
}{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5), big.NewInt(6)}

var BallotStates = struct {
	Invalid    *big.Int
	Ready      *big.Int
	InProgress *big.Int
	Accepted   *big.Int
	Rejected   *big.Int
}{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4)}
