// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gov

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetContractDomain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_contract\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"granted\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"SetPermission\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"contracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"getContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_contract\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_granted\",\"type\":\"address\"}],\"name\":\"getPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"magic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modifiedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"permissions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setContractDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_contract\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_granted\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ec56a373": "contracts(bytes32)",
		"0d2020dd": "getContractAddress(bytes32)",
		"60d6c7cf": "getPermission(bytes32,address)",
		"0d854646": "magic()",
		"7d10dd1b": "modifiedBlock()",
		"8da5cb5b": "owner()",
		"3ec50c6c": "permissions(bytes32,address)",
		"715018a6": "renounceOwnership()",
		"04af66ad": "setContractDomain(bytes32,address)",
		"599e4c70": "setPermission(bytes32,address,bool)",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x60806040526d57656d697820526567697374727960015534801561002257600080fd5b5061002c33610031565b610081565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610659806100906000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806360d6c7cf1161007157806360d6c7cf14610159578063715018a6146101925780637d10dd1b1461019c5780638da5cb5b146101a5578063ec56a373146101b6578063f2fde38b146101df57600080fd5b806304af66ad146100ae5780630d2020dd146100d65780630d854646146101015780633ec50c6c14610118578063599e4c7014610146575b600080fd5b6100c16100bc36600461050b565b6101f2565b60405190151581526020015b60405180910390f35b6100e96100e4366004610537565b6102b9565b6040516001600160a01b0390911681526020016100cd565b61010a60015481565b6040519081526020016100cd565b6100c161012636600461050b565b600460209081526000928352604080842090915290825290205460ff1681565b6100c1610154366004610550565b610309565b6100c161016736600461050b565b60009182526004602090815260408084206001600160a01b0393909316845291905290205460ff1690565b61019a6103ce565b005b61010a60025481565b6000546001600160a01b03166100e9565b6100e96101c4366004610537565b6003602052600090815260409020546001600160a01b031681565b61019a6101ed366004610595565b610404565b600080546001600160a01b031633146102265760405162461bcd60e51b815260040161021d906105b7565b60405180910390fd5b6001600160a01b03821661024c5760405162461bcd60e51b815260040161021d906105ec565b60008381526003602090815260409182902080546001600160a01b0319166001600160a01b03861690811790915543600255915133815285917f37724a4a9968ac9654e6ee52f3d0c93e5ef8863e057254ee2e36e8ad3e8429db910160405180910390a350600192915050565b6000818152600360205260408120546001600160a01b03166102ed5760405162461bcd60e51b815260040161021d906105ec565b506000908152600360205260409020546001600160a01b031690565b600080546001600160a01b031633146103345760405162461bcd60e51b815260040161021d906105b7565b6001600160a01b03831661035a5760405162461bcd60e51b815260040161021d906105ec565b60008481526004602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915543600255905190815286917fe9f5231bbfb4b32867755b94562215cff6c8998489de8ba20926f8d0980e7818910160405180910390a35060019392505050565b6000546001600160a01b031633146103f85760405162461bcd60e51b815260040161021d906105b7565b610402600061049f565b565b6000546001600160a01b0316331461042e5760405162461bcd60e51b815260040161021d906105b7565b6001600160a01b0381166104935760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161021d565b61049c8161049f565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80356001600160a01b038116811461050657600080fd5b919050565b6000806040838503121561051e57600080fd5b8235915061052e602084016104ef565b90509250929050565b60006020828403121561054957600080fd5b5035919050565b60008060006060848603121561056557600080fd5b83359250610575602085016104ef565b91506040840135801515811461058a57600080fd5b809150509250925092565b6000602082840312156105a757600080fd5b6105b0826104ef565b9392505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601a908201527f616464726573732073686f756c64206265206e6f6e2d7a65726f00000000000060408201526060019056fea2646970667358221220b7b950e503cb9742f3e07f2288e28f096b282d71a7b32fe533576ba4d480652064736f6c634300080e0033",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Deprecated: Use RegistryMetaData.Sigs instead.
// RegistryFuncSigs maps the 4-byte function signature to its string representation.
var RegistryFuncSigs = RegistryMetaData.Sigs

// RegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegistryMetaData.Bin instead.
var RegistryBin = RegistryMetaData.Bin

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// Contracts is a free data retrieval call binding the contract method 0xec56a373.
//
// Solidity: function contracts(bytes32 ) view returns(address)
func (_Registry *RegistryCaller) Contracts(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "contracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Contracts is a free data retrieval call binding the contract method 0xec56a373.
//
// Solidity: function contracts(bytes32 ) view returns(address)
func (_Registry *RegistrySession) Contracts(arg0 [32]byte) (common.Address, error) {
	return _Registry.Contract.Contracts(&_Registry.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0xec56a373.
//
// Solidity: function contracts(bytes32 ) view returns(address)
func (_Registry *RegistryCallerSession) Contracts(arg0 [32]byte) (common.Address, error) {
	return _Registry.Contract.Contracts(&_Registry.CallOpts, arg0)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(bytes32 _name) view returns(address addr)
func (_Registry *RegistryCaller) GetContractAddress(opts *bind.CallOpts, _name [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getContractAddress", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(bytes32 _name) view returns(address addr)
func (_Registry *RegistrySession) GetContractAddress(_name [32]byte) (common.Address, error) {
	return _Registry.Contract.GetContractAddress(&_Registry.CallOpts, _name)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(bytes32 _name) view returns(address addr)
func (_Registry *RegistryCallerSession) GetContractAddress(_name [32]byte) (common.Address, error) {
	return _Registry.Contract.GetContractAddress(&_Registry.CallOpts, _name)
}

// GetPermission is a free data retrieval call binding the contract method 0x60d6c7cf.
//
// Solidity: function getPermission(bytes32 _contract, address _granted) view returns(bool found)
func (_Registry *RegistryCaller) GetPermission(opts *bind.CallOpts, _contract [32]byte, _granted common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getPermission", _contract, _granted)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPermission is a free data retrieval call binding the contract method 0x60d6c7cf.
//
// Solidity: function getPermission(bytes32 _contract, address _granted) view returns(bool found)
func (_Registry *RegistrySession) GetPermission(_contract [32]byte, _granted common.Address) (bool, error) {
	return _Registry.Contract.GetPermission(&_Registry.CallOpts, _contract, _granted)
}

// GetPermission is a free data retrieval call binding the contract method 0x60d6c7cf.
//
// Solidity: function getPermission(bytes32 _contract, address _granted) view returns(bool found)
func (_Registry *RegistryCallerSession) GetPermission(_contract [32]byte, _granted common.Address) (bool, error) {
	return _Registry.Contract.GetPermission(&_Registry.CallOpts, _contract, _granted)
}

// Magic is a free data retrieval call binding the contract method 0x0d854646.
//
// Solidity: function magic() view returns(uint256)
func (_Registry *RegistryCaller) Magic(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "magic")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Magic is a free data retrieval call binding the contract method 0x0d854646.
//
// Solidity: function magic() view returns(uint256)
func (_Registry *RegistrySession) Magic() (*big.Int, error) {
	return _Registry.Contract.Magic(&_Registry.CallOpts)
}

// Magic is a free data retrieval call binding the contract method 0x0d854646.
//
// Solidity: function magic() view returns(uint256)
func (_Registry *RegistryCallerSession) Magic() (*big.Int, error) {
	return _Registry.Contract.Magic(&_Registry.CallOpts)
}

// ModifiedBlock is a free data retrieval call binding the contract method 0x7d10dd1b.
//
// Solidity: function modifiedBlock() view returns(uint256)
func (_Registry *RegistryCaller) ModifiedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "modifiedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModifiedBlock is a free data retrieval call binding the contract method 0x7d10dd1b.
//
// Solidity: function modifiedBlock() view returns(uint256)
func (_Registry *RegistrySession) ModifiedBlock() (*big.Int, error) {
	return _Registry.Contract.ModifiedBlock(&_Registry.CallOpts)
}

// ModifiedBlock is a free data retrieval call binding the contract method 0x7d10dd1b.
//
// Solidity: function modifiedBlock() view returns(uint256)
func (_Registry *RegistryCallerSession) ModifiedBlock() (*big.Int, error) {
	return _Registry.Contract.ModifiedBlock(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Permissions is a free data retrieval call binding the contract method 0x3ec50c6c.
//
// Solidity: function permissions(bytes32 , address ) view returns(bool)
func (_Registry *RegistryCaller) Permissions(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "permissions", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Permissions is a free data retrieval call binding the contract method 0x3ec50c6c.
//
// Solidity: function permissions(bytes32 , address ) view returns(bool)
func (_Registry *RegistrySession) Permissions(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Registry.Contract.Permissions(&_Registry.CallOpts, arg0, arg1)
}

// Permissions is a free data retrieval call binding the contract method 0x3ec50c6c.
//
// Solidity: function permissions(bytes32 , address ) view returns(bool)
func (_Registry *RegistryCallerSession) Permissions(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Registry.Contract.Permissions(&_Registry.CallOpts, arg0, arg1)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// SetContractDomain is a paid mutator transaction binding the contract method 0x04af66ad.
//
// Solidity: function setContractDomain(bytes32 _name, address _addr) returns(bool success)
func (_Registry *RegistryTransactor) SetContractDomain(opts *bind.TransactOpts, _name [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setContractDomain", _name, _addr)
}

// SetContractDomain is a paid mutator transaction binding the contract method 0x04af66ad.
//
// Solidity: function setContractDomain(bytes32 _name, address _addr) returns(bool success)
func (_Registry *RegistrySession) SetContractDomain(_name [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetContractDomain(&_Registry.TransactOpts, _name, _addr)
}

// SetContractDomain is a paid mutator transaction binding the contract method 0x04af66ad.
//
// Solidity: function setContractDomain(bytes32 _name, address _addr) returns(bool success)
func (_Registry *RegistryTransactorSession) SetContractDomain(_name [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetContractDomain(&_Registry.TransactOpts, _name, _addr)
}

// SetPermission is a paid mutator transaction binding the contract method 0x599e4c70.
//
// Solidity: function setPermission(bytes32 _contract, address _granted, bool _status) returns(bool success)
func (_Registry *RegistryTransactor) SetPermission(opts *bind.TransactOpts, _contract [32]byte, _granted common.Address, _status bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setPermission", _contract, _granted, _status)
}

// SetPermission is a paid mutator transaction binding the contract method 0x599e4c70.
//
// Solidity: function setPermission(bytes32 _contract, address _granted, bool _status) returns(bool success)
func (_Registry *RegistrySession) SetPermission(_contract [32]byte, _granted common.Address, _status bool) (*types.Transaction, error) {
	return _Registry.Contract.SetPermission(&_Registry.TransactOpts, _contract, _granted, _status)
}

// SetPermission is a paid mutator transaction binding the contract method 0x599e4c70.
//
// Solidity: function setPermission(bytes32 _contract, address _granted, bool _status) returns(bool success)
func (_Registry *RegistryTransactorSession) SetPermission(_contract [32]byte, _granted common.Address, _status bool) (*types.Transaction, error) {
	return _Registry.Contract.SetPermission(&_Registry.TransactOpts, _contract, _granted, _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryOwnershipTransferred, error) {
	event := new(RegistryOwnershipTransferred)
	if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrySetContractDomainIterator is returned from FilterSetContractDomain and is used to iterate over the raw logs and unpacked data for SetContractDomain events raised by the Registry contract.
type RegistrySetContractDomainIterator struct {
	Event *RegistrySetContractDomain // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistrySetContractDomainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrySetContractDomain)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistrySetContractDomain)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistrySetContractDomainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrySetContractDomainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrySetContractDomain represents a SetContractDomain event raised by the Registry contract.
type RegistrySetContractDomain struct {
	Setter common.Address
	Name   [32]byte
	Addr   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetContractDomain is a free log retrieval operation binding the contract event 0x37724a4a9968ac9654e6ee52f3d0c93e5ef8863e057254ee2e36e8ad3e8429db.
//
// Solidity: event SetContractDomain(address setter, bytes32 indexed name, address indexed addr)
func (_Registry *RegistryFilterer) FilterSetContractDomain(opts *bind.FilterOpts, name [][32]byte, addr []common.Address) (*RegistrySetContractDomainIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "SetContractDomain", nameRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &RegistrySetContractDomainIterator{contract: _Registry.contract, event: "SetContractDomain", logs: logs, sub: sub}, nil
}

// WatchSetContractDomain is a free log subscription operation binding the contract event 0x37724a4a9968ac9654e6ee52f3d0c93e5ef8863e057254ee2e36e8ad3e8429db.
//
// Solidity: event SetContractDomain(address setter, bytes32 indexed name, address indexed addr)
func (_Registry *RegistryFilterer) WatchSetContractDomain(opts *bind.WatchOpts, sink chan<- *RegistrySetContractDomain, name [][32]byte, addr []common.Address) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "SetContractDomain", nameRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrySetContractDomain)
				if err := _Registry.contract.UnpackLog(event, "SetContractDomain", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetContractDomain is a log parse operation binding the contract event 0x37724a4a9968ac9654e6ee52f3d0c93e5ef8863e057254ee2e36e8ad3e8429db.
//
// Solidity: event SetContractDomain(address setter, bytes32 indexed name, address indexed addr)
func (_Registry *RegistryFilterer) ParseSetContractDomain(log types.Log) (*RegistrySetContractDomain, error) {
	event := new(RegistrySetContractDomain)
	if err := _Registry.contract.UnpackLog(event, "SetContractDomain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrySetPermissionIterator is returned from FilterSetPermission and is used to iterate over the raw logs and unpacked data for SetPermission events raised by the Registry contract.
type RegistrySetPermissionIterator struct {
	Event *RegistrySetPermission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistrySetPermissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrySetPermission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistrySetPermission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistrySetPermissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrySetPermissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrySetPermission represents a SetPermission event raised by the Registry contract.
type RegistrySetPermission struct {
	Contract [32]byte
	Granted  common.Address
	Status   bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPermission is a free log retrieval operation binding the contract event 0xe9f5231bbfb4b32867755b94562215cff6c8998489de8ba20926f8d0980e7818.
//
// Solidity: event SetPermission(bytes32 indexed _contract, address indexed granted, bool status)
func (_Registry *RegistryFilterer) FilterSetPermission(opts *bind.FilterOpts, _contract [][32]byte, granted []common.Address) (*RegistrySetPermissionIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}
	var grantedRule []interface{}
	for _, grantedItem := range granted {
		grantedRule = append(grantedRule, grantedItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "SetPermission", _contractRule, grantedRule)
	if err != nil {
		return nil, err
	}
	return &RegistrySetPermissionIterator{contract: _Registry.contract, event: "SetPermission", logs: logs, sub: sub}, nil
}

// WatchSetPermission is a free log subscription operation binding the contract event 0xe9f5231bbfb4b32867755b94562215cff6c8998489de8ba20926f8d0980e7818.
//
// Solidity: event SetPermission(bytes32 indexed _contract, address indexed granted, bool status)
func (_Registry *RegistryFilterer) WatchSetPermission(opts *bind.WatchOpts, sink chan<- *RegistrySetPermission, _contract [][32]byte, granted []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}
	var grantedRule []interface{}
	for _, grantedItem := range granted {
		grantedRule = append(grantedRule, grantedItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "SetPermission", _contractRule, grantedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrySetPermission)
				if err := _Registry.contract.UnpackLog(event, "SetPermission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPermission is a log parse operation binding the contract event 0xe9f5231bbfb4b32867755b94562215cff6c8998489de8ba20926f8d0980e7818.
//
// Solidity: event SetPermission(bytes32 indexed _contract, address indexed granted, bool status)
func (_Registry *RegistryFilterer) ParseSetPermission(log types.Log) (*RegistrySetPermission, error) {
	event := new(RegistrySetPermission)
	if err := _Registry.contract.UnpackLog(event, "SetPermission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
