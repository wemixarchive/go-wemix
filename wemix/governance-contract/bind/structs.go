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
	"github.com/ethereum/go-ethereum/wemix/metclient"
)

//go:generate go run ../contracts/abigen.go

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
	// deploy registry
	if address, tx, contract, err := DeployRegistry(opts, backend); err != nil {
		return nil, err
	} else {
		txs = append(txs, tx)
		contractAddresses[tx.Hash()] = address
		gov.Registry = &Contract[Registry]{contract, address}
	}
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
	// deploy proxys
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
	// check deployed contracts
	for _, tx := range txs {
		actural, err := bind.WaitDeployed(opts.Context, backend, tx)
		if err != nil {
			return nil, err
		}
		if contractAddresses[tx.Hash()] != actural {

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
	for _, tx := range txs {
		receipt, err := bind.WaitMined(opts.Context, backend, tx)
		if err != nil {
			return nil, err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return nil, errors.New("")
		}
	}
	return gov, nil
}

func GetGovContractsByOwner(opts *bind.CallOpts, backend bind.ContractBackend, owner common.Address) (*GovContracts, error) {
	gov := new(GovContracts)
	registry, err := GetRegistryByOwner(opts, backend, owner)
	if err != nil {
		return nil, err
	}
	return gov, gov.Init(opts, backend, registry)
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

// TODO 이게 필요하지 않을까...?
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
