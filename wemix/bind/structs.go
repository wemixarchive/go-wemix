package gov

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/wemix/metclient"
	"github.com/pkg/errors"
)

//go:generate go run ../governance-contract/contracts/abigen.go --root=../governance-contract/contracts

var (
	magic = new(big.Int).SetBytes(hexutil.MustDecode("0x57656d6978205265676973747279"))
)

type GovContracts struct {
	Registry         *Registry
	Gov              *Gov
	GovImp           *GovImp
	Staking          *Staking
	StakingImp       *StakingImp
	BallotStorage    *BallotStorage
	BallotStorageImp *BallotStorageImp
	EnvStorage       *EnvStorage
	EnvStorageImp    *EnvStorageImp
	NCPExit          *NCPExit
	NCPExitImp       *NCPExitImp
	address          struct {
		Registry      common.Address
		Gov           common.Address
		Staking       common.Address
		BallotStorage common.Address
		EnvStorage    common.Address
		NCPExit       common.Address
	}
}

func NewGovContracts(backend bind.ContractBackend, registry common.Address) (*GovContracts, error) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	gov := new(GovContracts)

	if contract, err := NewRegistry(registry, backend); err != nil {
		return nil, err
	} else {
		gov.Registry, gov.address.Registry = contract, registry
	}
	var (
		err error
		cfg = &uupsConfig{
			callOpts: &bind.CallOpts{Context: ctx},
			backend:  backend,
			registry: gov.Registry,
		}
	)

	if gov.address.Gov, gov.Gov, gov.GovImp, err = newUUPSContracts(cfg, DOMAIN_Gov, NewGov, NewGovImp); err != nil {
		return nil, errors.Wrap(err, DOMAIN_Gov)
	} else if gov.address.Staking, gov.Staking, gov.StakingImp, err = newUUPSContracts(cfg, DOMAIN_Staking, NewStaking, NewStakingImp); err != nil {
		return nil, errors.Wrap(err, DOMAIN_Staking)
	} else if gov.address.BallotStorage, gov.BallotStorage, gov.BallotStorageImp, err = newUUPSContracts(cfg, DOMAIN_BallotStorage, NewBallotStorage, NewBallotStorageImp); err != nil {
		return nil, errors.Wrap(err, DOMAIN_BallotStorage)
	} else if gov.address.EnvStorage, gov.EnvStorage, gov.EnvStorageImp, err = newUUPSContracts(cfg, DOMAIN_EnvStorage, NewEnvStorage, NewEnvStorageImp); err != nil {
		return nil, errors.Wrap(err, DOMAIN_EnvStorage)
	} else if gov.address.NCPExit, gov.NCPExit, gov.NCPExitImp, err = newUUPSContracts(cfg, DOMAIN_NCPExit, NewNCPExit, NewNCPExitImp); err != nil {
		return nil, errors.Wrap(err, DOMAIN_NCPExit)
	} else {
		return gov, nil
	}
}

func DeployGovContracts(opts *bind.TransactOpts, backend iBackend, optionDomains map[string]common.Address) (*GovContracts, error) {
	if opts.Context == nil {
		ctx, cancel := context.WithCancel(context.Background())
		opts.Context = ctx
		defer func() { cancel(); opts.Context = nil }()
	}

	if optionDomains == nil {
		optionDomains = make(map[string]common.Address)
	}

	var (
		err error

		gov        = new(GovContracts)
		logger     = log.New("func", "DeployGovContracts")
		txPool     = &txPool{backend: backend, logger: logger, opts: opts, txs: make([]*types.Transaction, 0)}
		impAddress = struct {
			Gov           common.Address
			Staking       common.Address
			BallotStorage common.Address
			EnvStorage    common.Address
			NCPExit       common.Address
		}{}
	)

	// deploy registry
	if address, tx, contract, err := DeployRegistry(opts, backend); err != nil {
		return nil, errors.Wrap(err, REGISTRY)
	} else {
		logger.Info(fmt.Sprintf("Deploying Registry at %s...", address))
		txPool.AppendTx(tx, nil)
		gov.address.Registry, gov.Registry = address, contract
	}

	// deploy imps
	logger.Info("Deploy Logic Contracts...")
	if impAddress.Gov, err = deployLogic(txPool, DeployGovImp); err != nil {
		return nil, errors.Wrap(err, GOV_IMP)
	} else if impAddress.Staking, err = deployLogic(txPool, DeployStakingImp); err != nil {
		return nil, errors.Wrap(err, STAKING_IMP)
	} else if impAddress.BallotStorage, err = deployLogic(txPool, DeployBallotStorageImp); err != nil {
		return nil, errors.Wrap(err, BALLOTSTORAGE_IMP)
	} else if impAddress.EnvStorage, err = deployLogic(txPool, DeployEnvStorageImp); err != nil {
		return nil, errors.Wrap(err, ENVSTORAGE_IMP)
	} else if impAddress.NCPExit, err = deployLogic(txPool, DeployNCPExitImp); err != nil {
		return nil, errors.Wrap(err, NCPEXIT_IMP)
	} else if err = txPool.WaitMined(); err != nil {
		return nil, err
	}

	// deploy proxies
	logger.Info("Deploy Governance Contracts...")
	if gov.address.Gov, gov.Gov, err = deployProxy(txPool, impAddress.Gov, DeployGov); err != nil {
		return nil, errors.Wrap(err, GOV)
	} else if gov.address.Staking, gov.Staking, err = deployProxy(txPool, impAddress.Staking, DeployStaking); err != nil {
		return nil, errors.Wrap(err, STAKING)
	} else if gov.address.BallotStorage, gov.BallotStorage, err = deployProxy(txPool, impAddress.BallotStorage, DeployBallotStorage); err != nil {
		return nil, errors.Wrap(err, BALLOTSTORAGE)
	} else if gov.address.EnvStorage, gov.EnvStorage, err = deployProxy(txPool, impAddress.EnvStorage, DeployEnvStorage); err != nil {
		return nil, errors.Wrap(err, ENVSTORAGE)
	} else if gov.address.NCPExit, gov.NCPExit, err = deployProxy(txPool, impAddress.NCPExit, DeployNCPExit); err != nil {
		return nil, errors.Wrap(err, NCPEXIT)
	} else if err = txPool.WaitMined(); err != nil {
		return nil, err
	} else {
		optionDomains[DOMAIN_Gov] = gov.address.Gov
		optionDomains[DOMAIN_Staking] = gov.address.Staking
		optionDomains[DOMAIN_BallotStorage] = gov.address.BallotStorage
		optionDomains[DOMAIN_EnvStorage] = gov.address.EnvStorage
		optionDomains[DOMAIN_NCPExit] = gov.address.NCPExit
	}

	// setup registry
	logger.Info("Setting Domains...")
	for name, address := range optionDomains {
		if err := txPool.AppendTx(gov.Registry.SetContractDomain(opts, metclient.ToBytes32(name), address)); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("SetContractDomain(%s)", name))
		}
	}
	if err := txPool.WaitMined(); err != nil {
		return nil, err
	}

	// init impContracts
	if gov.GovImp, err = NewGovImp(gov.address.Gov, backend); err != nil {
		return nil, errors.Wrap(err, GOV_IMP)
	} else if gov.StakingImp, err = NewStakingImp(gov.address.Staking, backend); err != nil {
		return nil, errors.Wrap(err, STAKING_IMP)
	} else if gov.BallotStorageImp, err = NewBallotStorageImp(gov.address.BallotStorage, backend); err != nil {
		return nil, errors.Wrap(err, BALLOTSTORAGE_IMP)
	} else if gov.EnvStorageImp, err = NewEnvStorageImp(gov.address.EnvStorage, backend); err != nil {
		return nil, errors.Wrap(err, ENVSTORAGE_IMP)
	} else if gov.NCPExitImp, err = NewNCPExitImp(gov.address.NCPExit, backend); err != nil {
		return nil, errors.Wrap(err, NCPEXIT_IMP)
	} else {
		return gov, nil
	}
}

func ExecuteInitialize(gov *GovContracts, opts *bind.TransactOpts, backend iBackend, lockAmount *big.Int, envConfig InitEnvStorage, members InitMembers) error {
	if lockAmount.Cmp(envConfig.STAKING_MIN) < 0 || lockAmount.Cmp(envConfig.STAKING_MAX) > 0 {
		return fmt.Errorf("invalid lock amount, input:%v, min:%v, max:%v", lockAmount, envConfig.STAKING_MIN, envConfig.STAKING_MAX)
	}

	var (
		registryAddress = gov.address.Registry
		logger          = log.New("func", "ExecuteInitialize")
		txPool          = &txPool{backend: backend, logger: logger, opts: opts, txs: make([]*types.Transaction, 0)}
	)

	logger.Info("Initializing Staking...")
	if err := txPool.AppendTx(gov.StakingImp.Init(opts, registryAddress, members.StakingInit())); err != nil {
		return errors.Wrap(err, "StakingImp.Init")
	}

	logger.Info("Initializing BallotStorage...")
	if err := txPool.AppendTx(gov.BallotStorageImp.Initialize(opts, registryAddress)); err != nil {
		return errors.Wrap(err, "BallotStorageImp.Initialize")
	}

	logger.Info("Initializing NCPExit...")
	if err := txPool.AppendTx(gov.NCPExitImp.Initialize(opts, registryAddress)); err != nil {
		return errors.Wrap(err, "NCPExitImp.Initialize")
	}

	logger.Info("Initializing EnvStorage...")
	envNames, envValues := envConfig.Args()
	if err := txPool.AppendTx(gov.EnvStorageImp.Initialize(opts, registryAddress, envNames, envValues)); err != nil {
		return errors.Wrap(err, "EnvStorageImp.Initialize")
	}
	if err := txPool.WaitMined(); err != nil {
		return err
	}

	logger.Info("Initializing Gov...")
	if datas, err := members.GovInitOnce(); err != nil {
		return errors.Wrap(err, "members.GovInitOnce")
	} else if err := txPool.AppendTx(gov.GovImp.InitOnce(opts, registryAddress, lockAmount, datas)); err != nil {
		return errors.Wrap(err, "GovImp.InitOnce")
	} else {
		return txPool.WaitMined()
	}
}

func (gov *GovContracts) Address() struct {
	Registry      common.Address
	Gov           common.Address
	Staking       common.Address
	BallotStorage common.Address
	EnvStorage    common.Address
	NCPExit       common.Address
} {
	return gov.address
}

func (src *GovContracts) Equal(dst *GovContracts) bool {
	return reflect.DeepEqual(src.address, dst.address)
}

func GetGovContractsByOwner(opts *bind.CallOpts, backend bind.ContractBackend, owner common.Address) (*GovContracts, error) {
	address, _, err := GetRegistryByOwner(opts, backend, owner)
	if err != nil {
		return nil, err
	}
	return NewGovContracts(backend, address)
}

func GetRegistryByOwner(opts *bind.CallOpts, backend bind.ContractBackend, owner common.Address) (common.Address, *Registry, error) {
	for i := uint64(0); i < 10; i++ {
		address := crypto.CreateAddress(owner, i)
		registry, err := GetRegistryByAddress(opts, backend, address)
		if err == nil {
			return address, registry, nil
		}
	}
	return common.Address{}, nil, errors.Wrap(ethereum.NotFound, REGISTRY)
}

func GetRegistryByAddress(opts *bind.CallOpts, backend bind.ContractBackend, address common.Address) (*Registry, error) {
	if registry, err := NewRegistry(address, backend); err != nil {
		return nil, err
	} else if callMagic, err := registry.Magic(opts); err != nil {
		return nil, err
	} else if callMagic != nil && magic.Cmp(callMagic) == 0 {
		return registry, nil
	} else {
		return nil, ethereum.NotFound
	}
}

// //////////////////////
//    InitEnvStorage   //
// //////////////////////

type InitEnvStorage struct {
	BLOCKS_PER                               *big.Int `name:"blocksPer"`
	BALLOT_DURATION_MIN                      *big.Int `name:"ballotDurationMin"`
	BALLOT_DURATION_MAX                      *big.Int `name:"ballotDurationMax"`
	STAKING_MIN                              *big.Int `name:"stakingMin"`
	STAKING_MAX                              *big.Int `name:"stakingMax"`
	MAX_IDLE_BLOCK_INTERVAL                  *big.Int `name:"MaxIdleBlockInterval"`
	BLOCK_CREATION_TIME                      *big.Int `name:"blockCreationTime"`
	BLOCK_REWARD_AMOUNT                      *big.Int `name:"blockRewardAmount"`
	MAX_PRIORITY_FEE_PER_GAS                 *big.Int `name:"maxPriorityFeePerGas"`
	BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER *big.Int `name:"blockRewardDistributionBlockProducer"`
	BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD *big.Int `name:"blockRewardDistributionStakingReward"`
	BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM      *big.Int `name:"blockRewardDistributionEcosystem"`
	BLOCK_REWARD_DISTRIBUTION_MAINTENANCE    *big.Int `name:"blockRewardDistributionMaintenance"`
	MAX_BASE_FEE                             *big.Int `name:"maxBaseFee"`
	BLOCK_GASLIMIT                           *big.Int `name:"blockGasLimit"`
	BASE_FEE_MAX_CHANGE_RATE                 *big.Int `name:"baseFeeMaxChangeRate"`
	GAS_TARGET_PERCENTAGE                    *big.Int `name:"gasTargetPercentage"`
}

func (cfg InitEnvStorage) Args() (names [][32]byte, values []*big.Int) {
	v := reflect.ValueOf(cfg)
	t := reflect.TypeOf(cfg)
	numField := v.NumField()
	names = [][32]byte{}
	values = []*big.Int{}
	for i := 0; i < numField; i++ {
		value, ok := v.Field(i).Interface().(*big.Int)
		if ok && value != nil && value.Sign() > 0 {
			tag := t.Field(i).Tag.Get("name")
			names = append(names, crypto.Keccak256Hash([]byte(tag)))
			values = append(values, value)
		}
	}
	return
}

var DefaultInitEnvStorage = InitEnvStorage{
	BLOCKS_PER:                               big.NewInt(1),
	BALLOT_DURATION_MIN:                      big.NewInt(86400),
	BALLOT_DURATION_MAX:                      big.NewInt(604800),
	STAKING_MIN:                              new(big.Int).Mul(big.NewInt(1500000), big.NewInt(params.Ether)),
	STAKING_MAX:                              new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1), // type(uint256).max
	MAX_IDLE_BLOCK_INTERVAL:                  big.NewInt(5),
	BLOCK_CREATION_TIME:                      big.NewInt(1000),
	BLOCK_REWARD_AMOUNT:                      new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether)),
	MAX_PRIORITY_FEE_PER_GAS:                 new(big.Int).Mul(big.NewInt(100), big.NewInt(params.GWei)),
	BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER: big.NewInt(5000),
	BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD: big.NewInt(0),
	BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM:      big.NewInt(2500),
	BLOCK_REWARD_DISTRIBUTION_MAINTENANCE:    big.NewInt(2500),
	MAX_BASE_FEE:                             new(big.Int).Mul(big.NewInt(5000), big.NewInt(params.GWei)),
	BLOCK_GASLIMIT:                           big.NewInt(1050000000),
	BASE_FEE_MAX_CHANGE_RATE:                 big.NewInt(46),
	GAS_TARGET_PERCENTAGE:                    big.NewInt(30),
}

// /////////////////////////
// Init Governance Member //
// /////////////////////////
type InitMember struct {
	Staker  common.Address `json:"staker"`
	Voter   common.Address `json:"voter"`
	Reward  common.Address `json:"reward"`
	Name    string         `json:"name"`
	Enode   string         `json:"enode"`
	Ip      string         `json:"ip"`
	Port    int            `json:"port"`
	Deposit *big.Int       `json:"deposit"`
}

type InitMembers []InitMember

func (members InitMembers) GovInitOnce() ([]byte, error) {
	var datas bytes.Buffer
	for _, m := range members {
		var (
			staker, voter, reward *big.Int
			id                    []byte
			err                   error
		)
		staker = new(big.Int).SetBytes(m.Staker[:])
		voter = new(big.Int).SetBytes(m.Voter[:])
		reward = new(big.Int).SetBytes(m.Reward[:])
		switch len(m.Enode) {
		case 128:
			id, err = hex.DecodeString(m.Enode)
		case 130:
			id, err = hex.DecodeString(m.Enode[2:])
		default:
			err = fmt.Errorf("invalid enode id %s", m.Enode)
		}
		if err != nil {
			return nil, err
		}

		datas.Write(metclient.PackNum(reflect.ValueOf(staker)))
		datas.Write(metclient.PackNum(reflect.ValueOf(voter)))
		datas.Write(metclient.PackNum(reflect.ValueOf(reward)))
		datas.Write(metclient.PackNum(reflect.ValueOf(len(m.Name))))
		datas.Write([]byte(m.Name))
		datas.Write(metclient.PackNum(reflect.ValueOf(len(id))))
		datas.Write(id)
		datas.Write(metclient.PackNum(reflect.ValueOf(len(m.Ip))))
		datas.Write([]byte(m.Ip))
		datas.Write(metclient.PackNum(reflect.ValueOf(m.Port)))
	}

	return datas.Bytes(), nil
}

func (members InitMembers) StakingInit() []byte {
	var datas bytes.Buffer
	for _, m := range members {
		staker := new(big.Int).SetBytes(m.Staker[:])
		datas.Write(metclient.PackNum(reflect.ValueOf(staker)))
		datas.Write(metclient.PackNum(reflect.ValueOf(m.Deposit)))
	}
	return datas.Bytes()
}

// ////////
// utils //
// ////////

type uupsConfig struct {
	backend  bind.ContractBackend
	callOpts *bind.CallOpts
	registry *Registry
}

func newUUPSContracts[P, L any](
	cfg *uupsConfig,
	name string,
	newProxy func(common.Address, bind.ContractBackend) (*P, error),
	newLogic func(common.Address, bind.ContractBackend) (*L, error),
) (common.Address, *P, *L, error) {
	if address, err := cfg.registry.GetContractAddress(cfg.callOpts, metclient.ToBytes32(name)); err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "GetContractAddress")
	} else if proxy, err := newProxy(address, cfg.backend); err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "newProxy")
	} else if imp, err := newLogic(address, cfg.backend); err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "newLogic")
	} else {
		return address, proxy, imp, nil
	}
}

type iBackend interface {
	bind.ContractBackend
	bind.DeployBackend
}

type txPool struct {
	backend iBackend
	logger  log.Logger
	opts    *bind.TransactOpts
	txs     []*types.Transaction
}

func (pool *txPool) WaitMined() error {
	pool.logger.Info("Waiting for receipts...", "count", len(pool.txs))

	ctx, cancel := context.WithTimeout(context.Background(), 5e9)
	defer cancel()
	for _, tx := range pool.txs {
		receipt, err := bind.WaitMined(ctx, pool.backend, tx)
		if err != nil {
			return err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return errors.New("execute reverted")
		}
	}
	pool.txs = pool.txs[:0]
	return nil
}

func (pool *txPool) AppendTx(tx *types.Transaction, err error) error {
	if tx != nil {
		pool.txs = append(pool.txs, tx)
	}
	return err
}

func deployProxy[T any](pool *txPool, logic common.Address, deploy func(*bind.TransactOpts, bind.ContractBackend, common.Address) (common.Address, *types.Transaction, *T, error)) (common.Address, *T, error) {
	address, tx, proxy, err := deploy(pool.opts, pool.backend, logic)
	return address, proxy, pool.AppendTx(tx, err)
}

func deployLogic[T any](pool *txPool, deploy func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *T, error)) (common.Address, error) {
	address, tx, _, err := deploy(pool.opts, pool.backend)
	return address, pool.AppendTx(tx, err)
}
