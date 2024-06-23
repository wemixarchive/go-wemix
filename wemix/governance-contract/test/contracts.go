package test

import (
	"context"
	"encoding/json"
	"math/big"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	compile "github.com/ethereum/go-ethereum/wemix/governance-contract"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

type Governance struct {
	backend   *backends.SimulatedBackend
	owner     *bind.TransactOpts
	nodeInfos []nodeInfo

	registry common.Address
	Registry,
	Gov, GovImp,
	NCPExit, NCPExitImp,
	Staking, StakingImp,
	BallotStorage, BallotStorageImp,
	EnvStorage, EnvStorageImp *bind.BoundContract
}

type nodeInfo struct {
	name  []byte
	enode []byte
	ip    []byte
	port  *big.Int
}

func NewGovernance(t *testing.T) *Governance {
	owner := getTxOpt(t, "owner")
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		owner.From: {Balance: new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 128), common.Big1)},
	}, params.MaxGasLimit)

	return &Governance{
		backend: backend,
		owner:   owner,
		nodeInfos: []nodeInfo{{
			[]byte("name"),
			hexutil.MustDecode("0x6f8a80d14311c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
			[]byte("127.0.0.1"),
			big.NewInt(8542),
		}},
	}
}

func (g *Governance) DeployContracts(t *testing.T) *Governance {
	// deploy registry
	registry, Registry, err := g.Deploy(compiled.Registry.Deploy(g.backend, g.owner))
	require.NoError(t, err)
	// deploy impls
	govImp, _, err := g.Deploy(compiled.GovImp.Deploy(g.backend, g.owner))
	require.NoError(t, err)
	ncpExitImp, _, err := g.Deploy(compiled.NCPExitImp.Deploy(g.backend, g.owner))
	require.NoError(t, err)
	stakingImp, _, err := g.Deploy(compiled.StakingImp.Deploy(g.backend, g.owner))
	require.NoError(t, err)
	ballotStorageImp, _, err := g.Deploy(compiled.BallotStorageImp.Deploy(g.backend, g.owner))
	require.NoError(t, err)
	envStorageImp, _, err := g.Deploy(compiled.EnvStorageImp.Deploy(g.backend, g.owner))
	require.NoError(t, err)

	// deploy proxies
	gov, Gov, err := g.Deploy(compiled.Gov.Deploy(g.backend, g.owner, govImp))
	require.NoError(t, err)
	ncpExit, NCPExit, err := g.Deploy(compiled.NCPExit.Deploy(g.backend, g.owner, ncpExitImp))
	require.NoError(t, err)
	staking, Staking, err := g.Deploy(compiled.Staking.Deploy(g.backend, g.owner, stakingImp))
	require.NoError(t, err)
	ballotStorage, BallotStorage, err := g.Deploy(compiled.BallotStorage.Deploy(g.backend, g.owner, ballotStorageImp))
	require.NoError(t, err)
	envStorage, EnvStorage, err := g.Deploy(compiled.EnvStorage.Deploy(g.backend, g.owner, envStorageImp))
	require.NoError(t, err)

	// set up g
	g.registry = registry
	g.Registry = Registry
	g.Gov = Gov
	g.NCPExit = NCPExit
	g.Staking = Staking
	g.BallotStorage = BallotStorage
	g.EnvStorage = EnvStorage

	g.GovImp = compiled.GovImp.New(g.backend, gov)
	g.NCPExitImp = compiled.NCPExitImp.New(g.backend, ncpExit)
	g.StakingImp = compiled.StakingImp.New(g.backend, staking)
	g.BallotStorageImp = compiled.BallotStorageImp.New(g.backend, ballotStorage)
	g.EnvStorageImp = compiled.EnvStorageImp.New(g.backend, envStorage)

	// set Domains
	require.NoError(t, g.ExpectedOk(g.Registry.Transact(g.owner, "setContractDomain", ToBytes32("GovernanceContract"), gov)))
	require.NoError(t, g.ExpectedOk(g.Registry.Transact(g.owner, "setContractDomain", ToBytes32("Staking"), staking)))
	require.NoError(t, g.ExpectedOk(g.Registry.Transact(g.owner, "setContractDomain", ToBytes32("EnvStorage"), envStorage)))
	require.NoError(t, g.ExpectedOk(g.Registry.Transact(g.owner, "setContractDomain", ToBytes32("BallotStorage"), ballotStorage)))

	// initialize
	require.NoError(t, g.ExpectedOk(g.NCPExitImp.Transact(g.owner, "initialize", registry)))
	require.NoError(t, g.ExpectedOk(g.StakingImp.Transact(g.owner, "init", registry, []byte{})))
	require.NoError(t, g.ExpectedOk(g.BallotStorageImp.Transact(g.owner, "initialize", registry)))
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
	require.NoError(t, g.ExpectedOk(g.EnvStorageImp.Transact(g.owner, "initialize", registry, envNames, envValues)))

	// staking
	{
		g.owner.Value = LOCK_AMOUNT
		require.NoError(t, g.ExpectedOk(g.StakingImp.Transact(g.owner, "deposit")))
		g.owner.Value = nil
	}
	node := g.nodeInfos[0]
	require.NoError(t, g.ExpectedOk(g.GovImp.Transact(g.owner, "init", registry, LOCK_AMOUNT, node.name, node.enode, node.ip, node.port)))

	return g
}

func (r *Governance) Deploy(address common.Address, tx *types.Transaction, contract *bind.BoundContract, err error) (common.Address, *bind.BoundContract, error) {
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, contract, r.ExpectedOk(tx, err)
}

func (r *Governance) ExpectedOk(tx *types.Transaction, err error) error {
	r.backend.Commit()
	if err != nil {
		return err
	} else if receipt, err := bind.WaitMined(context.TODO(), r.backend, tx); err != nil {
		return err
	} else if receipt.Status != types.ReceiptStatusSuccessful {
		panic(vm.ErrExecutionReverted)
	} else {
		return nil
	}
}

func (r *Governance) ExpectedFail(tx *types.Transaction, err error) error {
	r.backend.Commit()
	if err != nil {
		return err
	} else if receipt, err := bind.WaitMined(context.TODO(), r.backend, tx); err != nil {
		return err
	} else if receipt.Status == types.ReceiptStatusSuccessful {
		panic("execution not reverted")
	} else {
		return vm.ErrExecutionReverted
	}
}

type MemberInfo struct {
	Staker     common.Address `json:"staker"`
	Voter      common.Address `json:"voter"`
	Reward     common.Address `json:"reward"`
	Name       []byte         `json:"name"`
	Enode      []byte         `json:"enode"`
	Ip         []byte         `json:"ip"`
	Port       *big.Int       `json:"port"`
	LockAmount *big.Int       `json:"lockAmount"`
	Memo       []byte         `json:"memo"`
	Duration   *big.Int       `json:"duration"`
}

type bindBackend interface {
	bind.ContractBackend
	bind.DeployBackend
}

type bindContract struct {
	Bin []byte
	Abi abi.ABI
}

func newBindContract(contract *compiler.Contract) (*bindContract, error) {
	if contract == nil {
		return nil, errors.New("nil contracts")
	}

	if parsedAbi, err := parseABI(contract.Info.AbiDefinition); err != nil {
		return nil, err
	} else {
		code := contract.Code
		if !strings.HasPrefix(code, "0x") {
			code = "0x" + code
		}
		return &bindContract{Bin: hexutil.MustDecode(code), Abi: *parsedAbi}, err
	}
}

func (bc *bindContract) New(backend bindBackend, address common.Address) *bind.BoundContract {
	return bind.NewBoundContract(address, bc.Abi, backend, backend, backend)
}

func (bc *bindContract) Deploy(backend bindBackend, opts *bind.TransactOpts, args ...interface{}) (common.Address, *types.Transaction, *bind.BoundContract, error) {
	return bind.DeployContract(opts, bc.Abi, bc.Bin, backend, args...)
}

var (
	compiled compiledContract
)

func init() {
	compiled.Compile("../contracts")
}

type compiledContract struct {
	Registry,
	Gov, GovImp,
	NCPExit, NCPExitImp,
	Staking, StakingImp,
	BallotStorage, BallotStorageImp,
	EnvStorage, EnvStorageImp *bindContract
}

func (c *compiledContract) Compile(root string) {
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
	} else {
		if c.Registry, err = newBindContract(contracts["Registry"]); err != nil {
			panic(err)
		} else if c.Gov, err = newBindContract(contracts["Gov"]); err != nil {
			panic(err)
		} else if c.GovImp, err = newBindContract(contracts["GovImp"]); err != nil {
			panic(err)
		} else if c.NCPExit, err = newBindContract(contracts["NCPExit"]); err != nil {
			panic(err)
		} else if c.NCPExitImp, err = newBindContract(contracts["NCPExitImp"]); err != nil {
			panic(err)
		} else if c.Staking, err = newBindContract(contracts["Staking"]); err != nil {
			panic(err)
		} else if c.StakingImp, err = newBindContract(contracts["StakingImp"]); err != nil {
			panic(err)
		} else if c.BallotStorage, err = newBindContract(contracts["BallotStorage"]); err != nil {
			panic(err)
		} else if c.BallotStorageImp, err = newBindContract(contracts["BallotStorageImp"]); err != nil {
			panic(err)
		} else if c.EnvStorage, err = newBindContract(contracts["EnvStorage"]); err != nil {
			panic(err)
		} else if c.EnvStorageImp, err = newBindContract(contracts["EnvStorageImp"]); err != nil {
			panic(err)
		}
	}
}

func parseABI(abiDefinition interface{}) (*abi.ABI, error) {
	s, ok := abiDefinition.(string)
	if !ok {
		if bytes, err := json.Marshal(abiDefinition); err != nil {
			return nil, err
		} else {
			s = string(bytes)
		}
	}
	if abi, err := abi.JSON(strings.NewReader(s)); err != nil {
		return nil, err
	} else {
		return &abi, nil
	}
}
