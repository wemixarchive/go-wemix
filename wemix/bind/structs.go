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
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/wemix/metclient"
	"github.com/pkg/errors"
)

//go:generate go run ../governance-contract/contracts/abigen.go --root='../governance-contract/contracts'

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
	address          struct {
		Registry      common.Address
		Gov           common.Address
		Staking       common.Address
		BallotStorage common.Address
		EnvStorage    common.Address
	}
}

func NewGovContracts(backend bind.ContractBackend, registry common.Address) (*GovContracts, error) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	callOpts := &bind.CallOpts{Context: ctx}
	gov := new(GovContracts)

	if contract, err := NewRegistry(registry, backend); err != nil {
		return nil, err
	} else {
		gov.Registry, gov.address.Registry = contract, registry
	}

	if address, err := gov.Registry.GetContractAddress(callOpts, metclient.ToBytes32("GovernanceContract")); err != nil {
		return nil, err
	} else if proxy, err := NewGov(address, backend); err != nil {
		return nil, err
	} else if imp, err := NewGovImp(address, backend); err != nil {
		return nil, err
	} else {
		gov.Gov, gov.GovImp, gov.address.Gov = proxy, imp, address
	}

	if address, err := gov.Registry.GetContractAddress(callOpts, metclient.ToBytes32("Staking")); err != nil {
		return nil, err
	} else if proxy, err := NewStaking(address, backend); err != nil {
		return nil, err
	} else if imp, err := NewStakingImp(address, backend); err != nil {
		return nil, err
	} else {
		gov.Staking, gov.StakingImp, gov.address.Staking = proxy, imp, address
	}

	if address, err := gov.Registry.GetContractAddress(callOpts, metclient.ToBytes32("BallotStorage")); err != nil {
		return nil, err
	} else if proxy, err := NewBallotStorage(address, backend); err != nil {
		return nil, err
	} else if imp, err := NewBallotStorageImp(address, backend); err != nil {
		return nil, err
	} else {
		gov.BallotStorage, gov.BallotStorageImp, gov.address.BallotStorage = proxy, imp, address
	}

	if address, err := gov.Registry.GetContractAddress(callOpts, metclient.ToBytes32("EnvStorage")); err != nil {
		return nil, err
	} else if proxy, err := NewEnvStorage(address, backend); err != nil {
		return nil, err
	} else if imp, err := NewEnvStorageImp(address, backend); err != nil {
		return nil, err
	} else {
		gov.EnvStorage, gov.EnvStorageImp, gov.address.EnvStorage = proxy, imp, address
	}

	return gov, nil
}

type IBackend interface {
	bind.ContractBackend
	bind.DeployBackend
}

func DeployGovContracts(opts *bind.TransactOpts, backend IBackend, optionDomains map[string]common.Address) (*GovContracts, map[string]common.Address, error) {
	if opts.Context == nil {
		ctx, cancel := context.WithCancel(context.Background())
		opts.Context = ctx
		defer func() { cancel(); opts.Context = nil }()
	}
	gov := new(GovContracts)
	contractAddresses := make(map[common.Hash]common.Address)
	impAddress := make(map[string]common.Address)

	// deploy registry
	txs := make([]*types.Transaction, 0)
	if address, tx, contract, err := DeployRegistry(opts, backend); err != nil {
		return nil, nil, errors.Wrap(err, "DeployRegistry")
	} else {
		txs = append(txs, tx)
		contractAddresses[tx.Hash()] = address
		gov.Registry = contract
		gov.address.Registry = address
	}

	// deploy imps
	if address, tx, _, err := DeployGovImp(opts, backend); err != nil {
		return nil, nil, errors.Wrap(err, "DeployGovImp")
	} else {
		txs = append(txs, tx)
		contractAddresses[tx.Hash()] = address
		impAddress["GovImp"] = address
	}
	if address, tx, _, err := DeployStakingImp(opts, backend); err != nil {
		return nil, nil, errors.Wrap(err, "DeployStakingImp")
	} else {
		txs = append(txs, tx)
		contractAddresses[tx.Hash()] = address
		impAddress["StakingImp"] = address
	}
	if address, tx, _, err := DeployBallotStorageImp(opts, backend); err != nil {
		return nil, nil, errors.Wrap(err, "DeployBallotStorageImp")
	} else {
		txs = append(txs, tx)
		contractAddresses[tx.Hash()] = address
		impAddress["BallotStorageImp"] = address
	}
	if address, tx, _, err := DeployEnvStorageImp(opts, backend); err != nil {
		return nil, nil, errors.Wrap(err, "DeployEnvStorageImp")
	} else {
		txs = append(txs, tx)
		contractAddresses[tx.Hash()] = address
		impAddress["EnvStorageImp"] = address
	}
	// check deployed contracts (registry + imps)
	for _, tx := range txs {
		address, err := bind.WaitDeployed(context.TODO(), backend, tx)
		if err != nil {
			return nil, nil, err
		}
		if contractAddresses[tx.Hash()] != address {
			return nil, nil, errors.New("deployed error")
		}
	}

	// deploy proxys
	txs = make([]*types.Transaction, 0)
	if address, tx, contract, err := DeployGov(opts, backend, impAddress["GovImp"]); err != nil {
		return nil, nil, errors.Wrap(err, "DeployGov")
	} else {
		txs = append(txs, tx)
		contractAddresses[tx.Hash()] = address
		gov.Gov = contract
		gov.address.Gov = address
	}
	if address, tx, contract, err := DeployStaking(opts, backend, impAddress["StakingImp"]); err != nil {
		return nil, nil, errors.Wrap(err, "DeployStaking")
	} else {
		contractAddresses[tx.Hash()] = address
		gov.Staking = contract
		gov.address.Staking = address
	}
	if address, tx, contract, err := DeployBallotStorage(opts, backend, impAddress["BallotStorageImp"]); err != nil {
		return nil, nil, errors.Wrap(err, "DeployBallotStorage")
	} else {
		contractAddresses[tx.Hash()] = address
		gov.BallotStorage = contract
		gov.address.BallotStorage = address
	}
	if address, tx, contract, err := DeployEnvStorage(opts, backend, impAddress["EnvStorageImp"]); err != nil {
		return nil, nil, errors.Wrap(err, "DeployEnvStorage")
	} else {
		contractAddresses[tx.Hash()] = address
		gov.EnvStorage = contract
		gov.address.EnvStorage = address
	}
	// check deployed contracts (proxies)
	for _, tx := range txs {
		address, err := bind.WaitDeployed(opts.Context, backend, tx)
		if err != nil {
			return nil, nil, err
		}
		if contractAddresses[tx.Hash()] != address {
			return nil, nil, errors.New("deployed error")
		}
	}

	// setup registry
	txs = make([]*types.Transaction, 0)
	if tx, err := gov.Registry.SetContractDomain(opts, metclient.ToBytes32("GovernanceContract"), gov.address.Gov); err != nil {
		return nil, nil, errors.Wrap(err, "SetContractDomain(GovernanceContract)")
	} else {
		txs = append(txs, tx)
	}
	if tx, err := gov.Registry.SetContractDomain(opts, metclient.ToBytes32("Staking"), gov.address.Staking); err != nil {
		return nil, nil, errors.Wrap(err, "SetContractDomain(Staking)")
	} else {
		txs = append(txs, tx)
	}
	if tx, err := gov.Registry.SetContractDomain(opts, metclient.ToBytes32("BallotStorage"), gov.address.BallotStorage); err != nil {
		return nil, nil, errors.Wrap(err, "SetContractDomain(BallotStorage)")
	} else {
		txs = append(txs, tx)
	}
	if tx, err := gov.Registry.SetContractDomain(opts, metclient.ToBytes32("EnvStorage"), gov.address.EnvStorage); err != nil {
		return nil, nil, errors.Wrap(err, "SetContractDomain(EnvStorage)")
	} else {
		txs = append(txs, tx)
	}
	for name, address := range optionDomains {
		if tx, err := gov.Registry.SetContractDomain(opts, metclient.ToBytes32(name), address); err != nil {
			return nil, nil, errors.Wrap(err, fmt.Sprintf("SetContractDomain(%s)", name))
		} else {
			txs = append(txs, tx)
		}
	}
	// check setup registry
	for _, tx := range txs {
		receipt, err := bind.WaitMined(opts.Context, backend, tx)
		if err != nil {
			return nil, nil, err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return nil, nil, fmt.Errorf("reverted SetContractDomain")
		}
	}

	var err error
	if gov.GovImp, err = NewGovImp(gov.address.Gov, backend); err != nil {
		return nil, nil, errors.Wrap(err, "NewGovImp")
	} else if gov.StakingImp, err = NewStakingImp(gov.address.Staking, backend); err != nil {
		return nil, nil, errors.Wrap(err, "NewStakingImp")
	} else if gov.BallotStorageImp, err = NewBallotStorageImp(gov.address.BallotStorage, backend); err != nil {
		return nil, nil, errors.Wrap(err, "NewBallotStorageImp")
	} else if gov.EnvStorageImp, err = NewEnvStorageImp(gov.address.EnvStorage, backend); err != nil {
		return nil, nil, errors.Wrap(err, "NewEnvStorageImp")
	} else {
		return gov, impAddress, nil
	}
}

func ExecuteInitialize(gov *GovContracts, opts *bind.TransactOpts, backend IBackend, envConfig InitEnvStorage, members InitMembers) error {
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

	registryAddress := gov.address.Registry

	txs := make([]*types.Transaction, 0)
	if tx, err := gov.StakingImp.Init(opts, registryAddress, members.StakingInit()); err != nil {
		return errors.Wrap(err, "StakingImp.Init")
	} else {
		txs = append(txs, tx)
	}
	if tx, err := gov.BallotStorageImp.Initialize(opts, registryAddress); err != nil {
		return errors.Wrap(err, "BallotStorageImp.Initialize")
	} else {
		txs = append(txs, tx)
	}
	envNames, envValues := envConfig.Args()
	if tx, err := gov.EnvStorageImp.Initialize(opts, registryAddress, envNames, envValues); err != nil {
		return errors.Wrap(err, "EnvStorageImp.Initialize")
	} else {
		txs = append(txs, tx)
	}

	if err := waitMined(txs...); err != nil {
		return err
	}

	if datas, err := members.GovInitOnce(); err != nil {
		return errors.Wrap(err, "members.GovInitOnce")
	} else if tx, err := gov.GovImp.InitOnce(opts, registryAddress, envConfig.STAKING_MIN, datas); err != nil {
		return errors.Wrap(err, "GovImp.InitOnce")
	} else if err := waitMined(tx); err != nil {
		return err
	} else {
		return nil
	}
}

func (gov *GovContracts) Address() struct {
	Registry      common.Address
	Gov           common.Address
	Staking       common.Address
	BallotStorage common.Address
	EnvStorage    common.Address
} {
	return gov.address
}

func (src *GovContracts) Copy(dst *GovContracts, backend bind.ContractBackend) error {
	if dst == nil {
		return errors.New("nil destination")
	}
	zeroAddress := common.Address{}
	errZeroAddress := errors.New("invalid src address")

	srcAddress := src.Address()
	if address := srcAddress.Registry; address == zeroAddress {
		return errors.Wrap(errZeroAddress, "Registry")
	} else if contract, err := NewRegistry(address, backend); err != nil {
		return errors.Wrap(err, "Registry")
	} else {
		dst.Registry, dst.address.Registry = contract, address
	}

	if address := srcAddress.Gov; address == zeroAddress {
		return errors.Wrap(errZeroAddress, "Gov")
	} else if proxy, err := NewGov(address, backend); err != nil {
		return errors.Wrap(err, "Gov")
	} else if imp, err := NewGovImp(address, backend); err != nil {
		return errors.Wrap(err, "GovImp")
	} else {
		dst.Gov, dst.GovImp, dst.address.Gov = proxy, imp, address
	}

	if address := srcAddress.Staking; address == zeroAddress {
		return errors.Wrap(errZeroAddress, "Staking")
	} else if proxy, err := NewStaking(address, backend); err != nil {
		return errors.Wrap(err, "Staking")
	} else if imp, err := NewStakingImp(address, backend); err != nil {
		return errors.Wrap(err, "StakingImp")
	} else {
		dst.Staking, dst.StakingImp, dst.address.Staking = proxy, imp, address
	}

	if address := srcAddress.BallotStorage; address == zeroAddress {
		return errors.Wrap(errZeroAddress, "BallotStorage")
	} else if proxy, err := NewBallotStorage(address, backend); err != nil {
		return errors.Wrap(err, "BallotStorage")
	} else if imp, err := NewBallotStorageImp(address, backend); err != nil {
		return errors.Wrap(err, "BallotStorageImp")
	} else {
		dst.BallotStorage, dst.BallotStorageImp, dst.address.BallotStorage = proxy, imp, address
	}

	if address := srcAddress.EnvStorage; address == zeroAddress {
		return errors.Wrap(errZeroAddress, "EnvStorage")
	} else if proxy, err := NewEnvStorage(address, backend); err != nil {
		return errors.Wrap(err, "EnvStorage")
	} else if imp, err := NewEnvStorageImp(address, backend); err != nil {
		return errors.Wrap(err, "EnvStorageImp")
	} else {
		dst.EnvStorage, dst.EnvStorageImp, dst.address.EnvStorage = proxy, imp, address
	}

	return nil
}

func (src *GovContracts) Equal(dst *GovContracts) bool {
	return reflect.DeepEqual(src.address, dst.address)
}

func GetGovContractsByOwner(opts *bind.CallOpts, backend bind.ContractBackend, owner common.Address) (*GovContracts, error) {
	_, address, err := GetRegistryByOwner(opts, backend, owner)
	if err != nil {
		return nil, err
	}
	return NewGovContracts(backend, address)
}

func GetRegistryByOwner(opts *bind.CallOpts, backend bind.ContractBackend, owner common.Address) (*Registry, common.Address, error) {
	for i := uint64(0); i < 10; i++ {
		address := crypto.CreateAddress(owner, i)
		registry, err := GetRegistryByAddress(opts, backend, address)
		if err == nil {
			return registry, address, nil
		}
	}
	return nil, common.Address{}, errors.Wrap(ethereum.NotFound, "Registry")
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
// InitEnvStorage //
// //////////////////////

type InitEnvStorage struct {
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
	BLOCK_REWARD_DISTRIBUTION_MAINTENANCE    *big.Int
	MAX_BASE_FEE                             *big.Int
	BLOCK_GASLIMIT                           *big.Int
	BASE_FEE_MAX_CHANGE_RATE                 *big.Int
	GAS_TARGET_PERCENTAGE                    *big.Int
	Options                                  map[string]*big.Int
}

func (cfg InitEnvStorage) Args() (names [][32]byte, values []*big.Int) {
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
	if value := cfg.BLOCK_REWARD_DISTRIBUTION_MAINTENANCE; value != nil && value.Sign() > 0 {
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

var DefaultInitEnvStorage InitEnvStorage = InitEnvStorage{
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
	BLOCK_REWARD_DISTRIBUTION_MAINTENANCE:    big.NewInt(2500),
	MAX_BASE_FEE:                             new(big.Int).Mul(big.NewInt(5000), big.NewInt(params.GWei)),
	BLOCK_GASLIMIT:                           big.NewInt(1050000000),
	BASE_FEE_MAX_CHANGE_RATE:                 big.NewInt(46),
	GAS_TARGET_PERCENTAGE:                    big.NewInt(30),
	Options:                                  make(map[string]*big.Int),
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
