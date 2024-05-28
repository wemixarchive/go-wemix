package gov

import (
	"bytes"
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/wemix/metclient"
)

//go:generate go run ../governance-contract/contracts/abigen.go --root="../governance-contract/contracts"

var (
	magic = new(big.Int).SetBytes(hexutil.MustDecode("0x57656d6978205265676973747279"))
)

type Contract[T any] struct {
	Funcs   *T
	address common.Address
}

func (c *Contract[T]) Address() common.Address {
	return c.address
}

type GovContracts struct {
	Registry   *Contract[Registry]
	Gov        *Contract[Gov]
	GovImp     *Contract[GovImp]
	Staking    *Contract[Staking]
	StakingImp *Contract[StakingImp]
	// NCPExit          *Contract[NCPExit]
	// NCPExitImp       *Contract[NCPExitImp]
	BallotStorage    *Contract[BallotStorage]
	BallotStorageImp *Contract[BallotStorageImp]
	EnvStorage       *Contract[EnvStorage]
	EnvStorageImp    *Contract[EnvStorageImp]
}

func DeployGovContracts(opts *bind.TransactOpts, backend interface {
	bind.ContractBackend
	bind.DeployBackend
}) (*GovContracts, error) {
	if opts.Context == nil {
		ctx, cancel := context.WithCancel(context.Background())
		opts.Context = ctx
		defer func() { cancel(); opts.Context = nil }()
	}
	gov := new(GovContracts)
	contractAddresses := make(map[common.Hash]common.Address)
	txs := make([]*types.Transaction, 0)
	// deploy imps
	if address, tx, contract, err := DeployGovImp(opts, backend); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
		contractAddresses[tx.Hash()] = address
		gov.GovImp = &Contract[GovImp]{contract, address}
	}
	// if address, tx, contract, err := DeployNCPExitImp(opts, backend); err != nil {
	// 	return nil, err
	// } else {
	// 	contractAddresses[tx.Hash()] = address
	// 	gov.NCPExitImp = &Contract[NCPExitImp]{contract, address}
	// }
	if address, tx, contract, err := DeployStakingImp(opts, backend); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
		contractAddresses[tx.Hash()] = address
		gov.StakingImp = &Contract[StakingImp]{contract, address}
	}
	if address, tx, contract, err := DeployBallotStorageImp(opts, backend); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
		contractAddresses[tx.Hash()] = address
		gov.BallotStorageImp = &Contract[BallotStorageImp]{contract, address}
	}
	if address, tx, contract, err := DeployEnvStorageImp(opts, backend); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
		contractAddresses[tx.Hash()] = address
		gov.EnvStorageImp = &Contract[EnvStorageImp]{contract, address}
	}
	for _, tx := range txs {
		address, err := bind.WaitDeployed(context.TODO(), backend, tx)
		if err != nil {
			return nil, err
		}
		if contractAddresses[tx.Hash()] != address {
			return nil, errors.New("deployed error")
		}
	}
	// deploy proxys
	txs = make([]*types.Transaction, 0)
	if address, tx, contract, err := DeployGov(opts, backend, gov.GovImp.Address()); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
		contractAddresses[tx.Hash()] = address
		gov.Gov = &Contract[Gov]{contract, address}
	}
	// if address, tx, contract, err := DeployNCPExit(opts, backend, gov.NCPExitImp.Address()); err != nil {
	// 	return nil, err
	// } else {
	// 	contractAddresses[tx.Hash()] = address
	// 	gov.NCPExit = &Contract[NCPExit]{contract, address}
	// }
	if address, tx, contract, err := DeployStaking(opts, backend, gov.StakingImp.Address()); err != nil {
		return nil, err
	} else {
		contractAddresses[tx.Hash()] = address
		gov.Staking = &Contract[Staking]{contract, address}
	}
	if address, tx, contract, err := DeployBallotStorage(opts, backend, gov.BallotStorageImp.Address()); err != nil {
		return nil, err
	} else {
		contractAddresses[tx.Hash()] = address
		gov.BallotStorage = &Contract[BallotStorage]{contract, address}
	}
	if address, tx, contract, err := DeployEnvStorage(opts, backend, gov.EnvStorageImp.Address()); err != nil {
		return nil, err
	} else {
		contractAddresses[tx.Hash()] = address
		gov.EnvStorage = &Contract[EnvStorage]{contract, address}
	}
	// deploy registry
	if address, tx, contract, err := DeployRegistry(opts, backend); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
		contractAddresses[tx.Hash()] = address
		gov.Registry = &Contract[Registry]{contract, address}
	}

	// check deployed contracts
	for _, tx := range txs {
		address, err := bind.WaitDeployed(opts.Context, backend, tx)
		if err != nil {
			return nil, err
		}
		if !bytes.Equal(contractAddresses[tx.Hash()].Bytes(), address.Bytes()) {
			return nil, errors.New("deployed error")
		}
	}

	// setup registry
	txs = make([]*types.Transaction, 0)
	if tx, err := gov.Registry.Funcs.SetContractDomain(opts, metclient.ToBytes32("GovernanceContract"), gov.Gov.Address()); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
	}
	if tx, err := gov.Registry.Funcs.SetContractDomain(opts, metclient.ToBytes32("Staking"), gov.Staking.Address()); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
	}
	if tx, err := gov.Registry.Funcs.SetContractDomain(opts, metclient.ToBytes32("BallotStorage"), gov.BallotStorage.Address()); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
	}
	if tx, err := gov.Registry.Funcs.SetContractDomain(opts, metclient.ToBytes32("EnvStorage"), gov.EnvStorage.Address()); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
	}
	if tx, err := gov.Registry.Funcs.SetContractDomain(opts, metclient.ToBytes32("StakingReward"), opts.From); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
	}
	if tx, err := gov.Registry.Funcs.SetContractDomain(opts, metclient.ToBytes32("Ecosystem"), opts.From); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
	}
	for _, tx := range txs {
		receipt, err := bind.WaitMined(opts.Context, backend, tx)
		if err != nil {
			return nil, err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return nil, errors.New("reverted SetContractDomain")
		}
	}
	return gov, nil
}

func ExecuteInitialize(gov *GovContracts, opts *bind.TransactOpts,
	backend interface {
		bind.ContractBackend
		bind.DeployBackend
	},
	envConfig *EnvInitializeConfig,
	govInit func(Gov IGovInitFuncs) (*types.Transaction, error),
) (*GovContracts, error) {
	contracts := new(GovContracts)
	if err := gov.Copy(contracts, backend); err != nil {
		return nil, err
	}

	if !bytes.Equal(contracts.Gov.address[:], contracts.GovImp.address[:]) ||
		!bytes.Equal(contracts.Staking.address[:], contracts.StakingImp.address[:]) ||
		!bytes.Equal(contracts.BallotStorage.address[:], contracts.BallotStorageImp.address[:]) ||
		!bytes.Equal(contracts.EnvStorage.address[:], contracts.EnvStorageImp.address[:]) {
		if err := contracts.Init(new(bind.CallOpts), backend, contracts.Registry); err != nil {
			return nil, err
		}
	}

	waitMined := func(txs ...*types.Transaction) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5e9)
		defer cancel()
		for _, tx := range txs {
			receipt, err := bind.WaitMined(ctx, backend, tx)
			if err != nil {
				return err
			}
			if receipt.Status != types.ReceiptStatusSuccessful {
				return errors.New("execute reverted")
			}
		}
		return nil
	}

	txs := make([]*types.Transaction, 0)
	if tx, err := contracts.StakingImp.Funcs.Init(opts, contracts.Registry.Address(), []byte{}); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
	}
	if tx, err := contracts.BallotStorageImp.Funcs.Initialize(opts, contracts.Registry.Address()); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
	}
	envNames, envValues := envConfig.Args()
	if tx, err := contracts.EnvStorageImp.Funcs.Initialize(opts, contracts.Registry.Address(), envNames, envValues); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
	}
	if err := waitMined(txs...); err != nil {
		return nil, err
	}

	opts.Value = envConfig.STAKING_MIN
	if tx, err := contracts.StakingImp.Funcs.Deposit(opts); err != nil {
		return nil, err
	} else if err := waitMined(tx); err != nil {
		return nil, err
	}
	opts.Value = nil

	if tx, err := govInit(contracts.GovImp.Funcs); err != nil {
		return nil, err
	} else if tx != nil {
		if err := waitMined(tx); err != nil {
			return nil, err
		}
	}
	return contracts, nil
}

type IGovInitFuncs interface {
	Init(opts *bind.TransactOpts, registry common.Address, lockAmount *big.Int, name []byte, enode []byte, ip []byte, port *big.Int) (*types.Transaction, error)
	InitOnce(opts *bind.TransactOpts, registry common.Address, lockAmount *big.Int, data []byte) (*types.Transaction, error)
	InitMigration(opts *bind.TransactOpts, registry common.Address, oldModifiedBlock *big.Int, oldOwner common.Address) (*types.Transaction, error)
}

func (gov *GovContracts) Init(opts *bind.CallOpts, backend bind.ContractBackend, registry *Contract[Registry]) error {
	if gov.Registry == nil || gov.Gov.Address() != gov.GovImp.Address() {
		gov.Registry = registry
	} else {
		return errors.New("already inited")
	}

	if address, err := registry.Funcs.GetContractAddress(opts, metclient.ToBytes32("GovernanceContract")); err != nil {
		return err
	} else if proxy, err := NewGov(address, backend); err != nil {
		return err
	} else if imp, err := NewGovImp(address, backend); err != nil {
		return err
	} else {
		gov.Gov = &Contract[Gov]{proxy, address}
		gov.GovImp = &Contract[GovImp]{imp, address}
	}

	if address, err := registry.Funcs.GetContractAddress(opts, metclient.ToBytes32("Staking")); err != nil {
		return err
	} else if proxy, err := NewStaking(address, backend); err != nil {
		return err
	} else if imp, err := NewStakingImp(address, backend); err != nil {
		return err
	} else {
		gov.Staking = &Contract[Staking]{proxy, address}
		gov.StakingImp = &Contract[StakingImp]{imp, address}
	}

	if address, err := registry.Funcs.GetContractAddress(opts, metclient.ToBytes32("EnvStorage")); err != nil {
		return err
	} else if proxy, err := NewEnvStorage(address, backend); err != nil {
		return err
	} else if imp, err := NewEnvStorageImp(address, backend); err != nil {
		return err
	} else {
		gov.EnvStorage = &Contract[EnvStorage]{proxy, address}
		gov.EnvStorageImp = &Contract[EnvStorageImp]{imp, address}
	}

	if address, err := registry.Funcs.GetContractAddress(opts, metclient.ToBytes32("BallotStorage")); err != nil {
		return err
	} else if proxy, err := NewBallotStorage(address, backend); err != nil {
		return err
	} else if imp, err := NewBallotStorageImp(address, backend); err != nil {
		return err
	} else {
		gov.BallotStorage = &Contract[BallotStorage]{proxy, address}
		gov.BallotStorageImp = &Contract[BallotStorageImp]{imp, address}
	}
	return nil
}

func (src *GovContracts) Copy(dst *GovContracts, backend bind.ContractBackend) error {
	if dst == nil {
		return errors.New("nil destination")
	}

	if contract, err := NewRegistry(src.Registry.Address(), backend); err != nil {
		return err
	} else {
		dst.Registry = &Contract[Registry]{contract, src.Registry.Address()}
	}
	if contract, err := NewGov(src.Gov.Address(), backend); err != nil {
		return err
	} else {
		dst.Gov = &Contract[Gov]{contract, src.Gov.Address()}
	}
	if contract, err := NewGovImp(src.GovImp.Address(), backend); err != nil {
		return err
	} else {
		dst.GovImp = &Contract[GovImp]{contract, src.GovImp.Address()}
	}
	if contract, err := NewStaking(src.Staking.Address(), backend); err != nil {
		return err
	} else {
		dst.Staking = &Contract[Staking]{contract, src.Staking.Address()}
	}
	if contract, err := NewStakingImp(src.StakingImp.Address(), backend); err != nil {
		return err
	} else {
		dst.StakingImp = &Contract[StakingImp]{contract, src.StakingImp.Address()}
	}
	if contract, err := NewBallotStorage(src.BallotStorage.Address(), backend); err != nil {
		return err
	} else {
		dst.BallotStorage = &Contract[BallotStorage]{contract, src.BallotStorage.Address()}
	}
	if contract, err := NewBallotStorageImp(src.BallotStorageImp.Address(), backend); err != nil {
		return err
	} else {
		dst.BallotStorageImp = &Contract[BallotStorageImp]{contract, src.BallotStorageImp.Address()}
	}
	if contract, err := NewEnvStorage(src.EnvStorage.Address(), backend); err != nil {
		return err
	} else {
		dst.EnvStorage = &Contract[EnvStorage]{contract, src.EnvStorage.Address()}
	}
	if contract, err := NewEnvStorageImp(src.EnvStorageImp.Address(), backend); err != nil {
		return err
	} else {
		dst.EnvStorageImp = &Contract[EnvStorageImp]{contract, src.EnvStorageImp.Address()}
	}

	return nil
}

func (src *GovContracts) Equal(dst *GovContracts) bool {
	return bytes.Equal(src.Registry.address.Bytes(), dst.Registry.address.Bytes()) &&
		bytes.Equal(src.Gov.address.Bytes(), dst.Gov.address.Bytes()) &&
		bytes.Equal(src.GovImp.address.Bytes(), dst.GovImp.address.Bytes()) &&
		bytes.Equal(src.Staking.address.Bytes(), dst.Staking.address.Bytes()) &&
		bytes.Equal(src.StakingImp.address.Bytes(), dst.StakingImp.address.Bytes()) &&
		bytes.Equal(src.BallotStorage.address.Bytes(), dst.BallotStorage.address.Bytes()) &&
		bytes.Equal(src.BallotStorageImp.address.Bytes(), dst.BallotStorageImp.address.Bytes()) &&
		bytes.Equal(src.EnvStorage.address.Bytes(), dst.EnvStorage.address.Bytes()) &&
		bytes.Equal(src.EnvStorageImp.address.Bytes(), dst.EnvStorageImp.address.Bytes())
}

func GetRegistryByOwner(opts *bind.CallOpts, backend bind.ContractBackend, owner common.Address) (*Contract[Registry], error) {
	for i := uint64(0); i < 10; i++ {
		registry, err := GetRegistryByAddress(opts, backend, crypto.CreateAddress(owner, i))
		if err == nil {
			return registry, nil
		}
	}
	return nil, ethereum.NotFound
}

func GetGovContractsByOwner(opts *bind.CallOpts, backend bind.ContractBackend, owner common.Address) (*GovContracts, error) {
	gov := new(GovContracts)
	registry, err := GetRegistryByOwner(opts, backend, owner)
	if err != nil {
		return nil, err
	}
	return gov, gov.Init(opts, backend, registry)
}

// XXX 이게 필요하지 않을까...?
func GetRegistryByAddress(opts *bind.CallOpts, backend bind.ContractBackend, address common.Address) (*Contract[Registry], error) {
	if registry, err := NewRegistry(address, backend); err != nil {
		return nil, err
	} else if callMagic, err := registry.Magic(opts); err != nil {
		return nil, err
	} else if callMagic != nil && magic.Cmp(callMagic) == 0 {
		return &Contract[Registry]{registry, address}, nil
	} else {
		return nil, ethereum.NotFound
	}
}

// //////////////////////
// EnvInitializeConfig //
// //////////////////////

type EnvInitializeConfig struct {
	BLOCKS_PER                               *big.Int
	BALLOT_DURATION_MIN                      *big.Int
	BALLOT_DURATION_MAX                      *big.Int
	STAKING_MIN                              *big.Int
	STAKING_MAX                              *big.Int
	MAX_IDLE_BLOCK_INTERVAL                  *big.Int
	BLOCK_CREATION_TIME                      *big.Int
	BLOCK_REWARD_AMOUNT                      *big.Int
	MAX_PRIORITY_FEE_PER_GAS                 *big.Int
	BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER *big.Int
	BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD *big.Int
	BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM      *big.Int
	BLOCK_REWARD_DISTRIBUTION_MAINTANANCE    *big.Int
	MAX_BASE_FEE                             *big.Int
	BLOCK_GASLIMIT                           *big.Int
	BASE_FEE_MAX_CHANGE_RATE                 *big.Int
	GAS_TARGET_PERCENTAGE                    *big.Int
	Options                                  map[string]*big.Int
}

func (cfg *EnvInitializeConfig) Args() (names [][32]byte, values []*big.Int) {
	names = make([][32]byte, 0)
	values = make([]*big.Int, 0)
	if value := cfg.BLOCKS_PER; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("blocksPer")))
		values = append(values, value)
	}
	if value := cfg.BALLOT_DURATION_MIN; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("ballotDurationMin")))
		values = append(values, value)
	}
	if value := cfg.BALLOT_DURATION_MAX; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("ballotDurationMax")))
		values = append(values, value)
	}
	if value := cfg.STAKING_MIN; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("stakingMin")))
		values = append(values, value)
	}
	if value := cfg.STAKING_MAX; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("stakingMax")))
		values = append(values, value)
	}
	if value := cfg.MAX_IDLE_BLOCK_INTERVAL; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("MaxIdleBlockInterval")))
		values = append(values, value)
	}
	if value := cfg.BLOCK_CREATION_TIME; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("blockCreationTime")))
		values = append(values, value)
	}
	if value := cfg.BLOCK_REWARD_AMOUNT; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("blockRewardAmount")))
		values = append(values, value)
	}
	if value := cfg.MAX_PRIORITY_FEE_PER_GAS; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("maxPriorityFeePerGas")))
		values = append(values, value)
	}
	if value := cfg.BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("blockRewardDistributionBlockProducer")))
		values = append(values, value)
	}
	if value := cfg.BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("blockRewardDistributionStakingReward")))
		values = append(values, value)
	}
	if value := cfg.BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("blockRewardDistributionEcosystem")))
		values = append(values, value)
	}
	if value := cfg.BLOCK_REWARD_DISTRIBUTION_MAINTANANCE; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("blockRewardDistributionMaintenance")))
		values = append(values, value)
	}
	if value := cfg.MAX_BASE_FEE; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("maxBaseFee")))
		values = append(values, value)
	}
	if value := cfg.BLOCK_GASLIMIT; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("blockGasLimit")))
		values = append(values, value)
	}
	if value := cfg.BASE_FEE_MAX_CHANGE_RATE; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("baseFeeMaxChangeRate")))
		values = append(values, value)
	}
	if value := cfg.GAS_TARGET_PERCENTAGE; value != nil && value.Sign() > 0 {
		names = append(names, crypto.Keccak256Hash([]byte("gasTargetPercentage")))
		values = append(values, value)
	}
	for key, value := range cfg.Options {
		if value != nil && value.Sign() > 0 {
			names = append(names, crypto.Keccak256Hash([]byte(key)))
			values = append(values, value)
		}
	}

	return
}

var DefaultEnvInitializeConfig EnvInitializeConfig = EnvInitializeConfig{
	BLOCKS_PER:                               big.NewInt(1),
	BALLOT_DURATION_MIN:                      big.NewInt(86400),
	BALLOT_DURATION_MAX:                      big.NewInt(604800),
	STAKING_MIN:                              new(big.Int).Mul(big.NewInt(1500000), big.NewInt(params.Ether)),
	STAKING_MAX:                              new(big.Int).Mul(big.NewInt(1500000), big.NewInt(params.Ether)),
	MAX_IDLE_BLOCK_INTERVAL:                  big.NewInt(5),
	BLOCK_CREATION_TIME:                      big.NewInt(1000),
	BLOCK_REWARD_AMOUNT:                      new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether)),
	MAX_PRIORITY_FEE_PER_GAS:                 new(big.Int).Mul(big.NewInt(100), big.NewInt(params.GWei)),
	BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER: big.NewInt(4000),
	BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD: big.NewInt(1000),
	BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM:      big.NewInt(2500),
	BLOCK_REWARD_DISTRIBUTION_MAINTANANCE:    big.NewInt(2500),
	MAX_BASE_FEE:                             new(big.Int).Mul(big.NewInt(5000), big.NewInt(params.GWei)),
	BLOCK_GASLIMIT:                           big.NewInt(1050000000),
	BASE_FEE_MAX_CHANGE_RATE:                 big.NewInt(46),
	GAS_TARGET_PERCENTAGE:                    big.NewInt(30),
	Options:                                  make(map[string]*big.Int),
}
