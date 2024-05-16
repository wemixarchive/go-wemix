package metclient

// import (
// 	"math/big"

// 	"github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	gov "github.com/ethereum/go-ethereum/wemix/governance-contract/bind"
// 	"github.com/pkg/errors"
// )

// var (
// 	magic, _ = new(big.Int).SetString("0x57656d6978205265676973747279", 16)
// )

// func GetRegistryByOwnerAddress(opts *bind.CallOpts, client *ethclient.Client, owner common.Address) (common.Address, *gov.Registry) {
// 	for i := uint64(0); i < 10; i++ {
// 		addr := crypto.CreateAddress(owner, i)
// 		if registry, err := gov.NewRegistry(addr, client); err != nil {
// 			continue
// 		} else if callMagic, err := registry.Magic(opts); err != nil {
// 			continue
// 		} else if magic.Cmp(callMagic) == 0 {
// 			return addr, registry
// 		}
// 	}
// 	return common.Address{}, nil
// }

// func GetGovernanceContract(opts *bind.CallOpts, client *ethclient.Client, registry *gov.Registry) (common.Address, *gov.Gov, *gov.GovImp, error) {
// 	if addr, err := registry.GetContractAddress(opts, ToBytes32("GovernanceContract")); err != nil {
// 		return common.Address{}, nil, nil, err
// 	} else if addr == (common.Address{}) {
// 		return common.Address{}, nil, nil, errors.Wrap(ethereum.NotFound, "GovernanceContract")
// 	} else if Gov, err := gov.NewGov(addr, client); err != nil {
// 		return common.Address{}, nil, nil, err
// 	} else if GovImp, err := gov.NewGovImp(addr, client); err != nil {
// 		return common.Address{}, nil, nil, err
// 	} else {
// 		return addr, Gov, GovImp, nil
// 	}
// }

// func GetStaking(opts *bind.CallOpts, client *ethclient.Client, registry *gov.Registry) (common.Address, *gov.Staking, *gov.StakingImp, error) {
// 	if addr, err := registry.GetContractAddress(opts, ToBytes32("Staking")); err != nil {
// 		return common.Address{}, nil, nil, err
// 	} else if addr == (common.Address{}) {
// 		return common.Address{}, nil, nil, errors.Wrap(ethereum.NotFound, "Staking")
// 	} else if Staking, err := gov.NewStaking(addr, client); err != nil {
// 		return common.Address{}, nil, nil, err
// 	} else if StakingImp, err := gov.NewStakingImp(addr, client); err != nil {
// 		return common.Address{}, nil, nil, err
// 	} else {
// 		return addr, Staking, StakingImp, nil
// 	}
// }

// func GetEnvStorage(opts *bind.CallOpts, client *ethclient.Client, registry *gov.Registry) (common.Address, *gov.EnvStorage, *gov.EnvStorageImp, error) {
// 	if addr, err := registry.GetContractAddress(opts, ToBytes32("EnvStorage")); err != nil {
// 		return common.Address{}, nil, nil, err
// 	} else if addr == (common.Address{}) {
// 		return common.Address{}, nil, nil, errors.Wrap(ethereum.NotFound, "EnvStorage")
// 	} else if EnvStorage, err := gov.NewEnvStorage(addr, client); err != nil {
// 		return common.Address{}, nil, nil, err
// 	} else if EnvStorageImp, err := gov.NewEnvStorageImp(addr, client); err != nil {
// 		return common.Address{}, nil, nil, err
// 	} else {
// 		return addr, EnvStorage, EnvStorageImp, nil
// 	}
// }

// func GetBallotStorage(opts *bind.CallOpts, client *ethclient.Client, registry *gov.Registry) (common.Address, *gov.BallotStorage, *gov.BallotStorageImp, error) {
// 	if addr, err := registry.GetContractAddress(opts, ToBytes32("BallotStorage")); err != nil {
// 		return common.Address{}, nil, nil, err
// 	} else if addr == (common.Address{}) {
// 		return common.Address{}, nil, nil, errors.Wrap(ethereum.NotFound, "BallotStorage")
// 	} else if BallotStorage, err := gov.NewBallotStorage(addr, client); err != nil {
// 		return common.Address{}, nil, nil, err
// 	} else if BallotStorageImp, err := gov.NewBallotStorageImp(addr, client); err != nil {
// 		return common.Address{}, nil, nil, err
// 	} else {
// 		return addr, BallotStorage, BallotStorageImp, nil
// 	}
// }
