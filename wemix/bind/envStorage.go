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

// EnvStorageMetaData contains all meta data concerning the EnvStorage contract.
var EnvStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"5c60da1b": "implementation()",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161077238038061077283398101604081905261002f91610326565b604080516020810190915260008152819061006b60017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd61034f565b60008051602061072b8339815191521461008757610087610374565b6100938282600061009b565b505050610405565b6100a4836100d1565b6000825111806100b15750805b156100cc576100ca838361011160201b61008b1760201c565b505b505050565b6100da8161013d565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060610136838360405180606001604052806027815260200161074b602791396101fd565b9392505050565b610150816102db60201b6100b71760201c565b6101b75760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b806101dc60008051602061072b83398151915260001b6102ea60201b6100c61760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606001600160a01b0384163b6102655760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016101ae565b600080856001600160a01b03168560405161028091906103b6565b600060405180830381855af49150503d80600081146102bb576040519150601f19603f3d011682016040523d82523d6000602084013e6102c0565b606091505b5090925090506102d18282866102ed565b9695505050505050565b6001600160a01b03163b151590565b90565b606083156102fc575081610136565b82511561030c5782518084602001fd5b8160405162461bcd60e51b81526004016101ae91906103d2565b60006020828403121561033857600080fd5b81516001600160a01b038116811461013657600080fd5b60008282101561036f57634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052600160045260246000fd5b60005b838110156103a557818101518382015260200161038d565b838111156100ca5750506000910152565b600082516103c881846020870161038a565b9190910192915050565b60208152600082518060208401526103f181604085016020870161038a565b601f01601f19169190910160400192915050565b610317806104146000396000f3fe6080604052600436106100225760003560e01c80635c60da1b1461003957610031565b366100315761002f61006a565b005b61002f61006a565b34801561004557600080fd5b5061004e61007c565b6040516001600160a01b03909116815260200160405180910390f35b61007a6100756100c9565b6100fc565b565b60006100866100c9565b905090565b60606100b083836040518060600160405280602781526020016102bb60279139610120565b9392505050565b6001600160a01b03163b151590565b90565b60006100867f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b3660008037600080366000845af43d6000803e80801561011b573d6000f35b3d6000fd5b60606001600160a01b0384163b61018d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084015b60405180910390fd5b600080856001600160a01b0316856040516101a8919061026b565b600060405180830381855af49150503d80600081146101e3576040519150601f19603f3d011682016040523d82523d6000602084013e6101e8565b606091505b50915091506101f8828286610202565b9695505050505050565b606083156102115750816100b0565b8251156102215782518084602001fd5b8160405162461bcd60e51b81526004016101849190610287565b60005b8381101561025657818101518382015260200161023e565b83811115610265576000848401525b50505050565b6000825161027d81846020870161023b565b9190910192915050565b60208152600082518060208401526102a681604085016020870161023b565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122073a93b04f755a8d9dc8a21cd7ced80e99dae623e3f5317cf75e12209357071a664736f6c634300080e0033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
}

// EnvStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use EnvStorageMetaData.ABI instead.
var EnvStorageABI = EnvStorageMetaData.ABI

// Deprecated: Use EnvStorageMetaData.Sigs instead.
// EnvStorageFuncSigs maps the 4-byte function signature to its string representation.
var EnvStorageFuncSigs = EnvStorageMetaData.Sigs

// EnvStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnvStorageMetaData.Bin instead.
var EnvStorageBin = EnvStorageMetaData.Bin

// DeployEnvStorage deploys a new Ethereum contract, binding an instance of EnvStorage to it.
func DeployEnvStorage(auth *bind.TransactOpts, backend bind.ContractBackend, _implementation common.Address) (common.Address, *types.Transaction, *EnvStorage, error) {
	parsed, err := EnvStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnvStorageBin), backend, _implementation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnvStorage{EnvStorageCaller: EnvStorageCaller{contract: contract}, EnvStorageTransactor: EnvStorageTransactor{contract: contract}, EnvStorageFilterer: EnvStorageFilterer{contract: contract}}, nil
}

// EnvStorage is an auto generated Go binding around an Ethereum contract.
type EnvStorage struct {
	EnvStorageCaller     // Read-only binding to the contract
	EnvStorageTransactor // Write-only binding to the contract
	EnvStorageFilterer   // Log filterer for contract events
}

// EnvStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnvStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnvStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnvStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnvStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnvStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnvStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnvStorageSession struct {
	Contract     *EnvStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnvStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnvStorageCallerSession struct {
	Contract *EnvStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EnvStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnvStorageTransactorSession struct {
	Contract     *EnvStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EnvStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnvStorageRaw struct {
	Contract *EnvStorage // Generic contract binding to access the raw methods on
}

// EnvStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnvStorageCallerRaw struct {
	Contract *EnvStorageCaller // Generic read-only contract binding to access the raw methods on
}

// EnvStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnvStorageTransactorRaw struct {
	Contract *EnvStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnvStorage creates a new instance of EnvStorage, bound to a specific deployed contract.
func NewEnvStorage(address common.Address, backend bind.ContractBackend) (*EnvStorage, error) {
	contract, err := bindEnvStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnvStorage{EnvStorageCaller: EnvStorageCaller{contract: contract}, EnvStorageTransactor: EnvStorageTransactor{contract: contract}, EnvStorageFilterer: EnvStorageFilterer{contract: contract}}, nil
}

// NewEnvStorageCaller creates a new read-only instance of EnvStorage, bound to a specific deployed contract.
func NewEnvStorageCaller(address common.Address, caller bind.ContractCaller) (*EnvStorageCaller, error) {
	contract, err := bindEnvStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnvStorageCaller{contract: contract}, nil
}

// NewEnvStorageTransactor creates a new write-only instance of EnvStorage, bound to a specific deployed contract.
func NewEnvStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*EnvStorageTransactor, error) {
	contract, err := bindEnvStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnvStorageTransactor{contract: contract}, nil
}

// NewEnvStorageFilterer creates a new log filterer instance of EnvStorage, bound to a specific deployed contract.
func NewEnvStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*EnvStorageFilterer, error) {
	contract, err := bindEnvStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnvStorageFilterer{contract: contract}, nil
}

// bindEnvStorage binds a generic wrapper to an already deployed contract.
func bindEnvStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnvStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnvStorage *EnvStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnvStorage.Contract.EnvStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnvStorage *EnvStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnvStorage.Contract.EnvStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnvStorage *EnvStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnvStorage.Contract.EnvStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnvStorage *EnvStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnvStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnvStorage *EnvStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnvStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnvStorage *EnvStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnvStorage.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_EnvStorage *EnvStorageCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnvStorage.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_EnvStorage *EnvStorageSession) Implementation() (common.Address, error) {
	return _EnvStorage.Contract.Implementation(&_EnvStorage.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_EnvStorage *EnvStorageCallerSession) Implementation() (common.Address, error) {
	return _EnvStorage.Contract.Implementation(&_EnvStorage.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_EnvStorage *EnvStorageTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _EnvStorage.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_EnvStorage *EnvStorageSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _EnvStorage.Contract.Fallback(&_EnvStorage.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_EnvStorage *EnvStorageTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _EnvStorage.Contract.Fallback(&_EnvStorage.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EnvStorage *EnvStorageTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnvStorage.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EnvStorage *EnvStorageSession) Receive() (*types.Transaction, error) {
	return _EnvStorage.Contract.Receive(&_EnvStorage.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EnvStorage *EnvStorageTransactorSession) Receive() (*types.Transaction, error) {
	return _EnvStorage.Contract.Receive(&_EnvStorage.TransactOpts)
}

// EnvStorageAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the EnvStorage contract.
type EnvStorageAdminChangedIterator struct {
	Event *EnvStorageAdminChanged // Event containing the contract specifics and raw log

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
func (it *EnvStorageAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageAdminChanged)
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
		it.Event = new(EnvStorageAdminChanged)
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
func (it *EnvStorageAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageAdminChanged represents a AdminChanged event raised by the EnvStorage contract.
type EnvStorageAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EnvStorage *EnvStorageFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*EnvStorageAdminChangedIterator, error) {

	logs, sub, err := _EnvStorage.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &EnvStorageAdminChangedIterator{contract: _EnvStorage.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EnvStorage *EnvStorageFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *EnvStorageAdminChanged) (event.Subscription, error) {

	logs, sub, err := _EnvStorage.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageAdminChanged)
				if err := _EnvStorage.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EnvStorage *EnvStorageFilterer) ParseAdminChanged(log types.Log) (*EnvStorageAdminChanged, error) {
	event := new(EnvStorageAdminChanged)
	if err := _EnvStorage.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the EnvStorage contract.
type EnvStorageBeaconUpgradedIterator struct {
	Event *EnvStorageBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *EnvStorageBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageBeaconUpgraded)
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
		it.Event = new(EnvStorageBeaconUpgraded)
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
func (it *EnvStorageBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageBeaconUpgraded represents a BeaconUpgraded event raised by the EnvStorage contract.
type EnvStorageBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EnvStorage *EnvStorageFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*EnvStorageBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EnvStorage.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageBeaconUpgradedIterator{contract: _EnvStorage.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EnvStorage *EnvStorageFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *EnvStorageBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EnvStorage.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageBeaconUpgraded)
				if err := _EnvStorage.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EnvStorage *EnvStorageFilterer) ParseBeaconUpgraded(log types.Log) (*EnvStorageBeaconUpgraded, error) {
	event := new(EnvStorageBeaconUpgraded)
	if err := _EnvStorage.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EnvStorage contract.
type EnvStorageUpgradedIterator struct {
	Event *EnvStorageUpgraded // Event containing the contract specifics and raw log

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
func (it *EnvStorageUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageUpgraded)
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
		it.Event = new(EnvStorageUpgraded)
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
func (it *EnvStorageUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageUpgraded represents a Upgraded event raised by the EnvStorage contract.
type EnvStorageUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EnvStorage *EnvStorageFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EnvStorageUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnvStorage.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageUpgradedIterator{contract: _EnvStorage.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EnvStorage *EnvStorageFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EnvStorageUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnvStorage.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageUpgraded)
				if err := _EnvStorage.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EnvStorage *EnvStorageFilterer) ParseUpgraded(log types.Log) (*EnvStorageUpgraded, error) {
	event := new(EnvStorageUpgraded)
	if err := _EnvStorage.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpMetaData contains all meta data concerning the EnvStorageImp contract.
var EnvStorageImpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"AddressVarableChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"Bytes32VarableChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"BytesVarableChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_value\",\"type\":\"int256\"}],\"name\":\"IntVarableChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"StringVarableChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"UintVarableChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"UpgradeImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"VarableChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_DURATION_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BALLOT_DURATION_MIN_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BALLOT_DURATION_MIN_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BALLOT_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_FEE_MAX_CHANGE_RATE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCKS_PER_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_CREATION_TIME_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_GASLIMIT_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_AMOUNT_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_METHOD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ECOSYSTEM_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ENV_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GASLIMIT_AND_BASE_FEE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAS_TARGET_PERCENTAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAINTENANCE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BASE_FEE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_IDLE_BLOCK_INTERVAL_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PRIORITY_FEE_PER_GAS_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_POOL_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MIN_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MIN_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_REWARD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"envKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"envVal\",\"type\":\"bytes\"}],\"name\":\"checkVariableCondition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBallotDurationMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBallotDurationMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBallotDurationMinMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockCreationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockRewardDistributionMethod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlocksPer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"getBoolean\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"getBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"getBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasLimitAndBaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"getInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxBaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxIdleBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxPriorityFeePerGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingMinMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"getString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"getUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"names\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"infos\",\"type\":\"uint256[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reg\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setBallotDurationMax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setBallotDurationMaxByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setBallotDurationMin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setBallotDurationMinByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"setBallotDurationMinMax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setBallotDurationMinMaxByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setBlockCreationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setBlockCreationTimeByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setBlockRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setBlockRewardAmountByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block_producer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_staking_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ecofund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maintenance\",\"type\":\"uint256\"}],\"name\":\"setBlockRewardDistributionMethod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setBlockRewardDistributionMethodByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setBlocksPer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setBlocksPerByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block_GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseFeeMaxChangeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasTargetPercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBaseFee\",\"type\":\"uint256\"}],\"name\":\"setGasLimitAndBaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setGasLimitAndBaseFeeByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setMaxBaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setMaxBaseFeeByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setMaxIdleBlockInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setMaxIdleBlockIntervalByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setMaxPriorityFeePerGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setMaxPriorityFeePerGasByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setStakingMax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setStakingMaxByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setStakingMin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setStakingMinByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"setStakingMinMax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setStakingMinMaxByBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"envKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"envVal\",\"type\":\"bytes\"}],\"name\":\"setVariable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0015a73b": "BALLOT_DURATION_MAX_NAME()",
		"656e3052": "BALLOT_DURATION_MIN_MAX_NAME()",
		"6d583ca7": "BALLOT_DURATION_MIN_NAME()",
		"9986e4b9": "BALLOT_STORAGE_NAME()",
		"c42a0abc": "BASE_FEE_MAX_CHANGE_RATE_NAME()",
		"3f35c8fe": "BLOCKS_PER_NAME()",
		"c0b4fe15": "BLOCK_CREATION_TIME_NAME()",
		"238737b6": "BLOCK_GASLIMIT_NAME()",
		"a9b629b2": "BLOCK_REWARD_AMOUNT_NAME()",
		"c6713baf": "BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME()",
		"7b2bfb01": "BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME()",
		"46946416": "BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME()",
		"278bb12a": "BLOCK_REWARD_DISTRIBUTION_METHOD_NAME()",
		"6167eb45": "BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME()",
		"918f8674": "DENOMINATOR()",
		"34125c84": "ECOSYSTEM_NAME()",
		"7bf46530": "ENV_STORAGE_NAME()",
		"c7d3da34": "GASLIMIT_AND_BASE_FEE_NAME()",
		"4d273e28": "GAS_TARGET_PERCENTAGE_NAME()",
		"6c78d2cf": "GOV_NAME()",
		"4bd1ed76": "MAINTENANCE_NAME()",
		"f38ecf47": "MAX_BASE_FEE_NAME()",
		"b128f880": "MAX_IDLE_BLOCK_INTERVAL_NAME()",
		"38294419": "MAX_PRIORITY_FEE_PER_GAS_NAME()",
		"2f40992e": "REWARD_POOL_NAME()",
		"c00ace6c": "STAKING_MAX_NAME()",
		"a6868b7d": "STAKING_MIN_MAX_NAME()",
		"6fde207a": "STAKING_MIN_NAME()",
		"1e0cba0d": "STAKING_NAME()",
		"5a731cca": "STAKING_REWARD_NAME()",
		"9801bff9": "checkVariableCondition(bytes32,bytes)",
		"21f8a721": "getAddress(bytes32)",
		"1b27e01b": "getBallotDurationMax()",
		"33be496e": "getBallotDurationMin()",
		"cdb9f643": "getBallotDurationMinMax()",
		"33e31184": "getBlockCreationTime()",
		"5dba8c4a": "getBlockRewardAmount()",
		"8f6fe725": "getBlockRewardDistributionMethod()",
		"3690057a": "getBlocksPer()",
		"3848207a": "getBoolean(bytes32)",
		"c031a180": "getBytes(bytes32)",
		"a6ed563e": "getBytes32(bytes32)",
		"67d1a2e0": "getGasLimitAndBaseFee()",
		"dc97d962": "getInt(bytes32)",
		"c95437be": "getMaxBaseFee()",
		"185582f1": "getMaxIdleBlockInterval()",
		"2b2eaa92": "getMaxPriorityFeePerGas()",
		"737c59b8": "getStakingMax()",
		"076cd77f": "getStakingMin()",
		"aba01fee": "getStakingMinMax()",
		"986e791a": "getString(bytes32)",
		"bd02d0f5": "getUint(bytes32)",
		"539927be": "initialize(address,bytes32[],uint256[])",
		"8da5cb5b": "owner()",
		"52d1902d": "proxiableUUID()",
		"738fdd1a": "reg()",
		"715018a6": "renounceOwnership()",
		"852cf38f": "setBallotDurationMax(uint256)",
		"a5d6a581": "setBallotDurationMaxByBytes(bytes)",
		"a36f259b": "setBallotDurationMin(uint256)",
		"124cea37": "setBallotDurationMinByBytes(bytes)",
		"b4d6bd3b": "setBallotDurationMinMax(uint256,uint256)",
		"45b5ec29": "setBallotDurationMinMaxByBytes(bytes)",
		"3305508e": "setBlockCreationTime(uint256)",
		"44b89914": "setBlockCreationTimeByBytes(bytes)",
		"2ed19cd5": "setBlockRewardAmount(uint256)",
		"3d4c65f3": "setBlockRewardAmountByBytes(bytes)",
		"0add66dd": "setBlockRewardDistributionMethod(uint256,uint256,uint256,uint256)",
		"8c8b887e": "setBlockRewardDistributionMethodByBytes(bytes)",
		"79e85859": "setBlocksPer(uint256)",
		"3e8daafe": "setBlocksPerByBytes(bytes)",
		"0fca11d2": "setGasLimitAndBaseFee(uint256,uint256,uint256,uint256)",
		"be33732a": "setGasLimitAndBaseFeeByBytes(bytes)",
		"6fe13177": "setMaxBaseFee(uint256)",
		"408d79cf": "setMaxBaseFeeByBytes(bytes)",
		"f6fd7129": "setMaxIdleBlockInterval(uint256)",
		"0b90a39a": "setMaxIdleBlockIntervalByBytes(bytes)",
		"ec3df879": "setMaxPriorityFeePerGas(uint256)",
		"0fc238bf": "setMaxPriorityFeePerGasByBytes(bytes)",
		"a91ee0dc": "setRegistry(address)",
		"f8214591": "setStakingMax(uint256)",
		"2eccd832": "setStakingMaxByBytes(bytes)",
		"2eb57c65": "setStakingMin(uint256)",
		"71c6960d": "setStakingMinByBytes(bytes)",
		"8d5cdd7e": "setStakingMinMax(uint256,uint256)",
		"e078869f": "setStakingMinMaxByBytes(bytes)",
		"88c28019": "setVariable(bytes32,bytes)",
		"f2fde38b": "transferOwnership(address)",
		"3659cfe6": "upgradeTo(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000e8565b603854610100900460ff1615620000935760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60385460ff9081161015620000e6576038805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60805161366f6200011260003960008181611a1e01528181611aa30152611b7b015261366f6000f3fe6080604052600436106104e05760003560e01c80636fde207a11610281578063a6ed563e1161015a578063c42a0abc116100cc578063e078869f11610085578063e078869f146113d0578063ec3df879146113f0578063f2fde38b14611410578063f38ecf4714611430578063f6fd712914611452578063f82145911461147257600080fd5b8063c42a0abc146112a7578063c6713baf146112db578063c7d3da341461130f578063c95437be14611343578063cdb9f6431461138e578063dc97d962146113a357600080fd5b8063b4d6bd3b1161011e578063b4d6bd3b146111d6578063bd02d0f5146111f6578063be33732a14611223578063c00ace6c14611243578063c031a18014611265578063c0b4fe151461128557600080fd5b8063a6ed563e1461111b578063a91ee0dc14611148578063a9b629b214611168578063aba01fee1461118a578063b128f880146111b457600080fd5b80638c8b887e116101f35780639801bff9116101b75780639801bff914611036578063986e791a146110565780639986e4b914611083578063a36f259b146110a7578063a5d6a581146110c7578063a6868b7d146110e757600080fd5b80638c8b887e14610ee45780638d5cdd7e14610f045780638da5cb5b14610f245780638f6fe72514610f42578063918f86741461102057600080fd5b8063738fdd1a11610245578063738fdd1a14610e0f57806379e8585914610e2f5780637b2bfb0114610e4f5780637bf4653014610e83578063852cf38f14610ea457806388c2801914610ec457600080fd5b80636fde207a14610d4d5780636fe1317714610d6f578063715018a614610d8f57806371c6960d14610da4578063737c59b814610dc457600080fd5b80633659cfe6116103be5780634bd1ed76116103305780635dba8c4a116102e95780635dba8c4a14610b985780636167eb4514610be3578063656e305214610c1757806367d1a2e014610c4b5780636c78d2cf14610d025780636d583ca714610d2b57600080fd5b80634bd1ed7614610ad65780634d273e2814610af85780634f1ef28614610b2c57806352d1902d14610b3f578063539927be14610b545780635a731cca14610b7457600080fd5b80633e8daafe116103825780633e8daafe14610a005780633f35c8fe14610a20578063408d79cf14610a4257806344b8991414610a6257806345b5ec2914610a825780634694641614610aa257600080fd5b80633659cfe6146109135780633690057a14610933578063382944191461097e5780633848207a146109a05780633d4c65f3146109e057600080fd5b8063238737b6116104575780632ed19cd51161041b5780632ed19cd5146107fc5780632f40992e1461081c5780633305508e1461083d57806333be496e1461085d57806333e31184146108a857806334125c84146108f357600080fd5b8063238737b614610709578063278bb12a1461073d5780632b2eaa92146107715780632eb57c65146107bc5780632eccd832146107dc57600080fd5b80630fca11d2116104a95780630fca11d2146105c7578063124cea37146105e7578063185582f1146106075780631b27e01b146106525780631e0cba0d1461069d57806321f8a721146106bb57600080fd5b806215a73b146104e5578063076cd77f1461051a5780630add66dd146105655780630b90a39a146105875780630fc238bf146105a7575b600080fd5b3480156104f157600080fd5b5061050760008051602061351383398151915281565b6040519081526020015b60405180910390f35b34801561052657600080fd5b5060008051602061357383398151915260005260026020527fd7a437163435b79030132855d9146a10e26570f1767311219983e0c430cfa69054610507565b34801561057157600080fd5b50610585610580366004612f22565b611497565b005b34801561059357600080fd5b506105856105a236600461300b565b6115ed565b3480156105b357600080fd5b506105856105c236600461300b565b61162d565b3480156105d357600080fd5b506105856105e2366004612f22565b61166a565b3480156105f357600080fd5b5061058561060236600461300b565b61172f565b34801561061357600080fd5b506000805160206134d383398151915260005260026020527fbc779bbce8d4fef90926dc39b56d9135e8e3795964b3adae769767961d4a94d854610507565b34801561065e57600080fd5b5060008051602061351383398151915260005260026020527f5a1993193b8890c75d83157b2c2c12b144545357ab60021d7429856434de051b54610507565b3480156106a957600080fd5b50610507665374616b696e6760c81b81565b3480156106c757600080fd5b506106f16106d6366004613048565b6000908152600460205260409020546001600160a01b031690565b6040516001600160a01b039091168152602001610511565b34801561071557600080fd5b506105077f1d36f8ce53f59e624857e1d8dc7932d19981a2ea1b8faa4eb8ff843fc3e5a27881565b34801561074957600080fd5b506105077f9b2e0c7fdae148f225bae7deb92d7e7bd24bb77edb12956e8fa7204900dd8a2281565b34801561077d57600080fd5b5060008051602061355383398151915260005260026020527f270fb31b28afd06ad35288bc3cb901686a4c5a15239e966fd9b64269abd226c354610507565b3480156107c857600080fd5b506105856107d7366004613048565b61176c565b3480156107e857600080fd5b506105856107f736600461300b565b6117b3565b34801561080857600080fd5b50610585610817366004613048565b6117f0565b34801561082857600080fd5b506105076914995dd85c99141bdbdb60b21b81565b34801561084957600080fd5b50610585610858366004613048565b611837565b34801561086957600080fd5b506000805160206135fa83398151915260005260026020527f8e8c0bef3b4557d74ba01e0433728c574990153e35cacae69cb26f23a8a4017554610507565b3480156108b457600080fd5b506000805160206134f383398151915260005260026020527ff3497fe1496bfc19eaef6d253f75ec2e9cfe1b4f6080c29192641567c6fc5d8754610507565b3480156108ff57600080fd5b506105076845636f73797374656d60b81b81565b34801561091f57600080fd5b5061058561092e366004613076565b61187e565b34801561093f57600080fd5b506000805160206135da83398151915260005260026020527f14e650c949c7cc3fca2efe79b3711fc734b4df488cc27c4137f35bd0adaae7fa54610507565b34801561098a57600080fd5b5061050760008051602061355383398151915281565b3480156109ac57600080fd5b506109d06109bb366004613048565b60009081526020819052604090205460ff1690565b6040519015158152602001610511565b3480156109ec57600080fd5b506105856109fb36600461300b565b6118d8565b348015610a0c57600080fd5b50610585610a1b36600461300b565b611915565b348015610a2c57600080fd5b506105076000805160206135da83398151915281565b348015610a4e57600080fd5b50610585610a5d36600461300b565b611952565b348015610a6e57600080fd5b50610585610a7d36600461300b565b61198f565b348015610a8e57600080fd5b50610585610a9d36600461300b565b6119cc565b348015610aae57600080fd5b506105077fdd5a41a7fc01f5c6d30816b17f638d6531625f1e1eaa599673ab2f6079f2dd9d81565b348015610ae257600080fd5b506105076a4d61696e74656e616e636560a81b81565b348015610b0457600080fd5b506105077f77884798208df1e64f70968be41ef2d3211ec53613397ca59313416813df088881565b610585610b3a366004613093565b611a14565b348015610b4b57600080fd5b50610507611b6e565b348015610b6057600080fd5b50610585610b6f366004613172565b611c21565b348015610b8057600080fd5b506105076c14dd185ada5b99d4995dd85c99609a1b81565b348015610ba457600080fd5b5060008051602061353383398151915260005260026020527f38eb620cc1355554391bc9d0c224a726d66edf7c3eb2d67af166d6e81fe4456254610507565b348015610bef57600080fd5b506105077f9f1de481f899d76889aa8a2b44cc7b604d42691aa93d4ba618a1a1fd439f505081565b348015610c2357600080fd5b506105077fe10074dceffb75f13bf0ce50145afd35182d63796823f1280ce40e01c19109e781565b348015610c5757600080fd5b5060026020527fc77500645c8acbf6137a0e04460f85359483e3b08db30285cca2957db0c90904547fca55256fa8d248a334a7cb4ca72d9aec97700e81b3784b43c388b9a8d0e2b8e3547f77884798208df1e64f70968be41ef2d3211ec53613397ca59313416813df08886000527fba5524e01c8da4a46e6f70a8f29f4dc63d299ad01d2f01ccae9666cb85f6f7c05460408051938452602084019290925290820152606001610511565b348015610d0e57600080fd5b506105077111dbdd995c9b985b98d950dbdb9d1c9858dd60721b81565b348015610d3757600080fd5b506105076000805160206135fa83398151915281565b348015610d5957600080fd5b5061050760008051602061357383398151915281565b348015610d7b57600080fd5b50610585610d8a366004613048565b611eca565b348015610d9b57600080fd5b50610585611f11565b348015610db057600080fd5b50610585610dbf36600461300b565b611f25565b348015610dd057600080fd5b506000805160206134b383398151915260005260026020527f673c817b54d53459fa437d6b11ba6445e517d8dc789e6b815675ff10e6aeeb5d54610507565b348015610e1b57600080fd5b50609d546106f1906001600160a01b031681565b348015610e3b57600080fd5b50610585610e4a366004613048565b611f62565b348015610e5b57600080fd5b506105077f9346226931826838eedd13d9677fa33551e7c81f604b171ef3fac355458da9aa81565b348015610e8f57600080fd5b5061050769456e7653746f7261676560b01b81565b348015610eb057600080fd5b50610585610ebf366004613048565b611fa9565b348015610ed057600080fd5b50610585610edf36600461323f565b611ff0565b348015610ef057600080fd5b50610585610eff36600461300b565b612164565b348015610f1057600080fd5b50610585610f1f366004613270565b6121ba565b348015610f3057600080fd5b50606b546001600160a01b03166106f1565b348015610f4e57600080fd5b5060026020527fe6679752544e0074a1c9bec26775117fd40d133ce43cc443fb5f2512ebd7c2af547f89fc1b873e9ef0767cf41719c69200c3f894c43d73c49eb537ca60ec0a2d14bc547fa04c555f883d07b5c5fe85061c632b6bfec696b67afa8ad5e9773b89249984b2547fdd5a41a7fc01f5c6d30816b17f638d6531625f1e1eaa599673ab2f6079f2dd9d6000527f286a480438fb7eb9f5465ec436e6de865b71284ceb2be7ff1cea98f38c26308254604080519485526020850193909352918301526060820152608001610511565b34801561102c57600080fd5b5061050761271081565b34801561104257600080fd5b506109d061105136600461323f565b61228f565b34801561106257600080fd5b50611076611071366004613048565b6123ee565b60405161051191906132ea565b34801561108f57600080fd5b506105076c42616c6c6f7453746f7261676560981b81565b3480156110b357600080fd5b506105856110c2366004613048565b612490565b3480156110d357600080fd5b506105856110e236600461300b565b6124d7565b3480156110f357600080fd5b506105077f6c6f69f426081752a5d3e73746599acd2a4cb145d5de4203ca1e3473b281680b81565b34801561112757600080fd5b50610507611136366004613048565b60009081526006602052604090205490565b34801561115457600080fd5b50610585611163366004613076565b612514565b34801561117457600080fd5b5061050760008051602061353383398151915281565b34801561119657600080fd5b5061119f6125bc565b60408051928352602083019190915201610511565b3480156111c057600080fd5b506105076000805160206134d383398151915281565b3480156111e257600080fd5b506105856111f1366004613270565b612626565b34801561120257600080fd5b50610507611211366004613048565b60009081526002602052604090205490565b34801561122f57600080fd5b5061058561123e36600461300b565b6126fd565b34801561124f57600080fd5b506105076000805160206134b383398151915281565b34801561127157600080fd5b50611076611280366004613048565b61274c565b34801561129157600080fd5b506105076000805160206134f383398151915281565b3480156112b357600080fd5b506105077fb38b2c133e931937bd95337c65c8aefa7040ed64bbed732e3e29a4944c65747381565b3480156112e757600080fd5b506105077fc9e15e34073efbcd0328740feaf503caac9124b55b38c73d1a97b53da80a2f6081565b34801561131b57600080fd5b506105077f04f7b94450bbcad85f37ea47cd1062728f884bb2040e357738f8fd53056134bc81565b34801561134f57600080fd5b5060008051602061361a83398151915260005260026020527f38ddebd16170f2328156b72e009b30efe6e74a1b9fc4389d05f50efa2a42c6f854610507565b34801561139a57600080fd5b5061119f612769565b3480156113af57600080fd5b506105076113be366004613048565b60009081526001602052604090205490565b3480156113dc57600080fd5b506105856113eb36600461300b565b6127cf565b3480156113fc57600080fd5b5061058561140b366004613048565b612812565b34801561141c57600080fd5b5061058561142b366004613076565b612859565b34801561143c57600080fd5b5061050760008051602061361a83398151915281565b34801561145e57600080fd5b5061058561146d366004613048565b6128cf565b34801561147e57600080fd5b5061058561148d366004613048565b612916565b905090565b336114a061295d565b6001600160a01b0316146114cf5760405162461bcd60e51b81526004016114c6906132fd565b60405180910390fd5b61271081836114de868861333a565b6114e8919061333a565b6114f2919061333a565b1461153f5760405162461bcd60e51b815260206004820152601e60248201527f57726f6e6720726577617264206469737472756274696f6e20726174696f000060448201526064016114c6565b6115697fc9e15e34073efbcd0328740feaf503caac9124b55b38c73d1a97b53da80a2f608561297d565b6115937f9f1de481f899d76889aa8a2b44cc7b604d42691aa93d4ba618a1a1fd439f50508461297d565b6115bd7f9346226931826838eedd13d9677fa33551e7c81f604b171ef3fac355458da9aa8361297d565b6115e77fdd5a41a7fc01f5c6d30816b17f638d6531625f1e1eaa599673ab2f6079f2dd9d8261297d565b50505050565b336115f661295d565b6001600160a01b03161461161c5760405162461bcd60e51b81526004016114c6906132fd565b61162a61146d826020015190565b50565b3361163661295d565b6001600160a01b03161461165c5760405162461bcd60e51b81526004016114c6906132fd565b61162a61140b826020015190565b3361167361295d565b6001600160a01b0316146116995760405162461bcd60e51b81526004016114c6906132fd565b6116c37f1d36f8ce53f59e624857e1d8dc7932d19981a2ea1b8faa4eb8ff843fc3e5a2788561297d565b6116ed7fb38b2c133e931937bd95337c65c8aefa7040ed64bbed732e3e29a4944c6574738461297d565b6117177f77884798208df1e64f70968be41ef2d3211ec53613397ca59313416813df08888361297d565b6115e760008051602061361a8339815191528261297d565b3361173861295d565b6001600160a01b03161461175e5760405162461bcd60e51b81526004016114c6906132fd565b61162a6110c2826020015190565b3361177561295d565b6001600160a01b03161461179b5760405162461bcd60e51b81526004016114c6906132fd565b61162a6000805160206135738339815191528261297d565b336117bc61295d565b6001600160a01b0316146117e25760405162461bcd60e51b81526004016114c6906132fd565b61162a61148d826020015190565b336117f961295d565b6001600160a01b03161461181f5760405162461bcd60e51b81526004016114c6906132fd565b61162a6000805160206135338339815191528261297d565b3361184061295d565b6001600160a01b0316146118665760405162461bcd60e51b81526004016114c6906132fd565b61162a6000805160206134f38339815191528261297d565b6118866129cc565b6118a181604051806020016040528060008152506000612a26565b6040516001600160a01b038216907f2a74bbaf99cf25202925743012bb137262ab3432c821e1ba94ddbecac2ea970890600090a250565b336118e161295d565b6001600160a01b0316146119075760405162461bcd60e51b81526004016114c6906132fd565b61162a610817826020015190565b3361191e61295d565b6001600160a01b0316146119445760405162461bcd60e51b81526004016114c6906132fd565b61162a610e4a826020015190565b3361195b61295d565b6001600160a01b0316146119815760405162461bcd60e51b81526004016114c6906132fd565b61162a610d8a826020015190565b3361199861295d565b6001600160a01b0316146119be5760405162461bcd60e51b81526004016114c6906132fd565b61162a610858826020015190565b336119d561295d565b6001600160a01b0316146119fb5760405162461bcd60e51b81526004016114c6906132fd565b60208101516040820151611a0f8282612626565b505050565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003611aa15760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b60648201526084016114c6565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611aea600080516020613593833981519152546001600160a01b031690565b6001600160a01b031614611b555760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b60648201526084016114c6565b611b5e82612b91565b611b6a82826001612a26565b5050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611c0e5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016114c6565b5060008051602061359383398151915290565b603854610100900460ff1615808015611c415750603854600160ff909116105b80611c5b5750303b158015611c5b575060385460ff166001145b611cbe5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016114c6565b6038805460ff191660011790558015611ce1576038805461ff0019166101001790555b600080516020613593833981519152546001600160a01b03166001600160a01b0316846001600160a01b031603611d705760405162461bcd60e51b815260206004820152602d60248201527f72656769737472792073686f756c64206e6f742062652073616d65206173206960448201526c36b83632b6b2b73a30ba34b7b760991b60648201526084016114c6565b611d78612bc0565b611d8184612514565b60005b8251811015611e7e576000611dbe858381518110611da457611da4613352565b602002602001015160009081526002602052604090205490565b9050838281518110611dd257611dd2613352565b60200260200101516000141580611de857508015155b611e275760405162461bcd60e51b815260206004820152601060248201526f696e76616c6964207661726961626c6560801b60448201526064016114c6565b80600003611e6b57611e6b858381518110611e4457611e44613352565b6020026020010151858481518110611e5e57611e5e613352565b602002602001015161297d565b5080611e7681613368565b915050611d84565b5080156115e7576038805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b33611ed361295d565b6001600160a01b031614611ef95760405162461bcd60e51b81526004016114c6906132fd565b61162a60008051602061361a8339815191528261297d565b611f196129cc565b611f236000612bef565b565b33611f2e61295d565b6001600160a01b031614611f545760405162461bcd60e51b81526004016114c6906132fd565b61162a6107d7826020015190565b33611f6b61295d565b6001600160a01b031614611f915760405162461bcd60e51b81526004016114c6906132fd565b61162a6000805160206135da8339815191528261297d565b33611fb261295d565b6001600160a01b031614611fd85760405162461bcd60e51b81526004016114c6906132fd565b61162a6000805160206135138339815191528261297d565b6000805160206135da833981519152820361200e57611b6a81611915565b7fe10074dceffb75f13bf0ce50145afd35182d63796823f1280ce40e01c19109e7820361203e57611b6a816119cc565b7f6c6f69f426081752a5d3e73746599acd2a4cb145d5de4203ca1e3473b281680b820361206e57611b6a816127cf565b6000805160206134d3833981519152820361208c57611b6a816115ed565b6000805160206134f383398151915282036120aa57611b6a8161198f565b60008051602061353383398151915282036120c857611b6a816118d8565b60008051602061355383398151915282036120e657611b6a8161162d565b7f9b2e0c7fdae148f225bae7deb92d7e7bd24bb77edb12956e8fa7204900dd8a22820361211657611b6a81612164565b7f04f7b94450bbcad85f37ea47cd1062728f884bb2040e357738f8fd53056134bc820361214657611b6a816126fd565b60008051602061361a8339815191528203611b6a57611b6a81611952565b3361216d61295d565b6001600160a01b0316146121935760405162461bcd60e51b81526004016114c6906132fd565b60208101516040820151606083015160808401516121b384848484611497565b5050505050565b336121c361295d565b6001600160a01b0316146121e95760405162461bcd60e51b81526004016114c6906132fd565b8082111561225f5760405162461bcd60e51b815260206004820152603e60248201527f4d696e696d756d207374616b696e67206d75737420626520736d616c6c65722060448201527f616e6420657175616c207468616e206d6178696d756d207374616b696e67000060648201526084016114c6565b6122776000805160206135738339815191528361297d565b611b6a6000805160206134b38339815191528261297d565b60007f9b2e0c7fdae148f225bae7deb92d7e7bd24bb77edb12956e8fa7204900dd8a22830361231a57600080600080858060200190518101906122d29190613381565b9350935093509350612710818385876122eb919061333a565b6122f5919061333a565b6122ff919061333a565b146123115760009450505050506123e8565b505050506123e4565b7f6c6f69f426081752a5d3e73746599acd2a4cb145d5de4203ca1e3473b281680b83148061236757507fe10074dceffb75f13bf0ce50145afd35182d63796823f1280ce40e01c19109e783145b156123a1576000808380602001905181019061238391906133b7565b915091508082111561239a576000925050506123e8565b50506123e4565b6000805160206134f383398151915283036123e4576000828060200190518101906123cc91906133db565b90506103e88110156123e25760009150506123e8565b505b5060015b92915050565b600081815260036020526040902080546060919061240b906133f4565b80601f0160208091040260200160405190810160405280929190818152602001828054612437906133f4565b80156124845780601f1061245957610100808354040283529160200191612484565b820191906000526020600020905b81548152906001019060200180831161246757829003601f168201915b50505050509050919050565b3361249961295d565b6001600160a01b0316146124bf5760405162461bcd60e51b81526004016114c6906132fd565b61162a6000805160206135fa8339815191528261297d565b336124e061295d565b6001600160a01b0316146125065760405162461bcd60e51b81526004016114c6906132fd565b61162a610ebf826020015190565b61251c6129cc565b6001600160a01b0381166125725760405162461bcd60e51b815260206004820152601a60248201527f416464726573732073686f756c64206265206e6f6e2d7a65726f00000000000060448201526064016114c6565b609d80546001600160a01b0319166001600160a01b0383169081179091556040517f278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd90600090a250565b60026020527fd7a437163435b79030132855d9146a10e26570f1767311219983e0c430cfa690546000805160206134b383398151915260009081527f673c817b54d53459fa437d6b11ba6445e517d8dc789e6b815675ff10e6aeeb5d54909182915b915091509091565b3361262f61295d565b6001600160a01b0316146126555760405162461bcd60e51b81526004016114c6906132fd565b808211156126cd576040805162461bcd60e51b81526020600482015260248101919091527f4d696e696d756d206475726174696f6e206d75737420626520736d616c6c657260448201527f20616e6420657175616c207468616e206d6178696d756d206475726174696f6e60648201526084016114c6565b6126e56000805160206135fa8339815191528361297d565b611b6a6000805160206135138339815191528261297d565b3361270661295d565b6001600160a01b03161461272c5760405162461bcd60e51b81526004016114c6906132fd565b60208101516040820151606083015160808401516121b38484848461166a565b600081815260056020526040902080546060919061240b906133f4565b60026020527f8e8c0bef3b4557d74ba01e0433728c574990153e35cacae69cb26f23a8a401755460008051602061351383398151915260009081527f5a1993193b8890c75d83157b2c2c12b144545357ab60021d7429856434de051b549091829161261e565b336127d861295d565b6001600160a01b0316146127fe5760405162461bcd60e51b81526004016114c6906132fd565b60208101516040820151611a0f82826121ba565b3361281b61295d565b6001600160a01b0316146128415760405162461bcd60e51b81526004016114c6906132fd565b61162a6000805160206135538339815191528261297d565b6128616129cc565b6001600160a01b0381166128c65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016114c6565b61162a81612bef565b336128d861295d565b6001600160a01b0316146128fe5760405162461bcd60e51b81526004016114c6906132fd565b61162a6000805160206134d38339815191528261297d565b3361291f61295d565b6001600160a01b0316146129455760405162461bcd60e51b81526004016114c6906132fd565b61162a6000805160206134b38339815191528261297d565b60006114927111dbdd995c9b985b98d950dbdb9d1c9858dd60721b612c41565b6000828152600260205260409020819055817f58d7c10adfd5778016889c15d422d57f2975b9292415b54fe5d8cd4241200612826040516129c091815260200190565b60405180910390a25050565b606b546001600160a01b03163314611f235760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016114c6565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615612a5957611a0f83612caf565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612ab3575060408051601f3d908101601f19168201909252612ab0918101906133db565b60015b612b165760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b60648201526084016114c6565b6000805160206135938339815191528114612b855760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b60648201526084016114c6565b50611a0f838383612d4b565b33612b9a61295d565b6001600160a01b03161461162a5760405162461bcd60e51b81526004016114c6906132fd565b603854610100900460ff16612be75760405162461bcd60e51b81526004016114c69061342e565b611f23612d70565b606b80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b609d54604051630d2020dd60e01b8152600481018390526000916001600160a01b031690630d2020dd90602401602060405180830381865afa158015612c8b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123e89190613479565b6001600160a01b0381163b612d1c5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016114c6565b60008051602061359383398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b612d5483612da0565b600082511180612d615750805b15611a0f576115e78383612de0565b603854610100900460ff16612d975760405162461bcd60e51b81526004016114c69061342e565b611f2333612bef565b612da981612caf565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060612e0583836040518060600160405280602781526020016135b360279139612e0c565b9392505050565b60606001600160a01b0384163b612e745760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016114c6565b600080856001600160a01b031685604051612e8f9190613496565b600060405180830381855af49150503d8060008114612eca576040519150601f19603f3d011682016040523d82523d6000602084013e612ecf565b606091505b5091509150612edf828286612ee9565b9695505050505050565b60608315612ef8575081612e05565b825115612f085782518084602001fd5b8160405162461bcd60e51b81526004016114c691906132ea565b60008060008060808587031215612f3857600080fd5b5050823594602084013594506040840135936060013592509050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715612f9357612f93612f54565b604052919050565b600082601f830112612fac57600080fd5b813567ffffffffffffffff811115612fc657612fc6612f54565b612fd9601f8201601f1916602001612f6a565b818152846020838601011115612fee57600080fd5b816020850160208301376000918101602001919091529392505050565b60006020828403121561301d57600080fd5b813567ffffffffffffffff81111561303457600080fd5b61304084828501612f9b565b949350505050565b60006020828403121561305a57600080fd5b5035919050565b6001600160a01b038116811461162a57600080fd5b60006020828403121561308857600080fd5b8135612e0581613061565b600080604083850312156130a657600080fd5b82356130b181613061565b9150602083013567ffffffffffffffff8111156130cd57600080fd5b6130d985828601612f9b565b9150509250929050565b600067ffffffffffffffff8211156130fd576130fd612f54565b5060051b60200190565b600082601f83011261311857600080fd5b8135602061312d613128836130e3565b612f6a565b82815260059290921b8401810191818101908684111561314c57600080fd5b8286015b848110156131675780358352918301918301613150565b509695505050505050565b60008060006060848603121561318757600080fd5b833561319281613061565b925060208481013567ffffffffffffffff808211156131b057600080fd5b818701915087601f8301126131c457600080fd5b81356131d2613128826130e3565b81815260059190911b8301840190848101908a8311156131f157600080fd5b938501935b8285101561320f578435825293850193908501906131f6565b96505050604087013592508083111561322757600080fd5b505061323586828701613107565b9150509250925092565b6000806040838503121561325257600080fd5b82359150602083013567ffffffffffffffff8111156130cd57600080fd5b6000806040838503121561328357600080fd5b50508035926020909101359150565b60005b838110156132ad578181015183820152602001613295565b838111156115e75750506000910152565b600081518084526132d6816020860160208601613292565b601f01601f19169290920160200192915050565b602081526000612e0560208301846132be565b6020808252600d908201526c2737902832b936b4b9b9b4b7b760991b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6000821982111561334d5761334d613324565b500190565b634e487b7160e01b600052603260045260246000fd5b60006001820161337a5761337a613324565b5060010190565b6000806000806080858703121561339757600080fd5b505082516020840151604085015160609095015191969095509092509050565b600080604083850312156133ca57600080fd5b505080516020909101519092909150565b6000602082840312156133ed57600080fd5b5051919050565b600181811c9082168061340857607f821691505b60208210810361342857634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60006020828403121561348b57600080fd5b8151612e0581613061565b600082516134a8818460208701613292565b919091019291505056fe18ad4415ef4a621ce1a136395c51ab6c3712bb2e24b79d526059925cea58dcb8829561ab7af084b7efc6600518d2df79b8d95f3f4c3a550f54f8f7ec7d2b80578086da5becff4dfac91a3105821b361078d2d4abba0ccc2401b974cf0dcf05c10c4fbe9dc9de15dd7c0d064975ee1a2f2f9b954fa0e65d4f6cddba94884bdc3e89dd490ecaf395283ed4ff2fd9557ca767fc425dce063451a9b0da6d72f600c3be90e461bbdb9a95a694f7796912ea04244caf7f5b60ad7ded17e16821d3e44c0b09c9badbbeb6c813a598ee910770a39ccda797a1940439bb6e47fc6c87548b360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65642a268972a70c8c688b62366bdfdd9bb09cf19d3e5b6e7e7bb158e671ffdcedd2c69fc6b7d0efc934fd5a3581c7253a7107a952526bb6dbcd814ef8d8dae1f44a7c1150f0e1a39ff55552d52764f97e6c387e2a247e1df344369f122c4254be2fa2646970667358221220baa5bc7f1d90c081dacd5776c70f2fa86c00d6ba71186486575b670769a2c9f864736f6c634300080e0033",
}

// EnvStorageImpABI is the input ABI used to generate the binding from.
// Deprecated: Use EnvStorageImpMetaData.ABI instead.
var EnvStorageImpABI = EnvStorageImpMetaData.ABI

// Deprecated: Use EnvStorageImpMetaData.Sigs instead.
// EnvStorageImpFuncSigs maps the 4-byte function signature to its string representation.
var EnvStorageImpFuncSigs = EnvStorageImpMetaData.Sigs

// EnvStorageImpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnvStorageImpMetaData.Bin instead.
var EnvStorageImpBin = EnvStorageImpMetaData.Bin

// DeployEnvStorageImp deploys a new Ethereum contract, binding an instance of EnvStorageImp to it.
func DeployEnvStorageImp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnvStorageImp, error) {
	parsed, err := EnvStorageImpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnvStorageImpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnvStorageImp{EnvStorageImpCaller: EnvStorageImpCaller{contract: contract}, EnvStorageImpTransactor: EnvStorageImpTransactor{contract: contract}, EnvStorageImpFilterer: EnvStorageImpFilterer{contract: contract}}, nil
}

// EnvStorageImp is an auto generated Go binding around an Ethereum contract.
type EnvStorageImp struct {
	EnvStorageImpCaller     // Read-only binding to the contract
	EnvStorageImpTransactor // Write-only binding to the contract
	EnvStorageImpFilterer   // Log filterer for contract events
}

// EnvStorageImpCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnvStorageImpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnvStorageImpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnvStorageImpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnvStorageImpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnvStorageImpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnvStorageImpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnvStorageImpSession struct {
	Contract     *EnvStorageImp    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnvStorageImpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnvStorageImpCallerSession struct {
	Contract *EnvStorageImpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnvStorageImpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnvStorageImpTransactorSession struct {
	Contract     *EnvStorageImpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnvStorageImpRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnvStorageImpRaw struct {
	Contract *EnvStorageImp // Generic contract binding to access the raw methods on
}

// EnvStorageImpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnvStorageImpCallerRaw struct {
	Contract *EnvStorageImpCaller // Generic read-only contract binding to access the raw methods on
}

// EnvStorageImpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnvStorageImpTransactorRaw struct {
	Contract *EnvStorageImpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnvStorageImp creates a new instance of EnvStorageImp, bound to a specific deployed contract.
func NewEnvStorageImp(address common.Address, backend bind.ContractBackend) (*EnvStorageImp, error) {
	contract, err := bindEnvStorageImp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImp{EnvStorageImpCaller: EnvStorageImpCaller{contract: contract}, EnvStorageImpTransactor: EnvStorageImpTransactor{contract: contract}, EnvStorageImpFilterer: EnvStorageImpFilterer{contract: contract}}, nil
}

// NewEnvStorageImpCaller creates a new read-only instance of EnvStorageImp, bound to a specific deployed contract.
func NewEnvStorageImpCaller(address common.Address, caller bind.ContractCaller) (*EnvStorageImpCaller, error) {
	contract, err := bindEnvStorageImp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpCaller{contract: contract}, nil
}

// NewEnvStorageImpTransactor creates a new write-only instance of EnvStorageImp, bound to a specific deployed contract.
func NewEnvStorageImpTransactor(address common.Address, transactor bind.ContractTransactor) (*EnvStorageImpTransactor, error) {
	contract, err := bindEnvStorageImp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpTransactor{contract: contract}, nil
}

// NewEnvStorageImpFilterer creates a new log filterer instance of EnvStorageImp, bound to a specific deployed contract.
func NewEnvStorageImpFilterer(address common.Address, filterer bind.ContractFilterer) (*EnvStorageImpFilterer, error) {
	contract, err := bindEnvStorageImp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpFilterer{contract: contract}, nil
}

// bindEnvStorageImp binds a generic wrapper to an already deployed contract.
func bindEnvStorageImp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnvStorageImpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnvStorageImp *EnvStorageImpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnvStorageImp.Contract.EnvStorageImpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnvStorageImp *EnvStorageImpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.EnvStorageImpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnvStorageImp *EnvStorageImpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.EnvStorageImpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnvStorageImp *EnvStorageImpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnvStorageImp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnvStorageImp *EnvStorageImpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnvStorageImp *EnvStorageImpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.contract.Transact(opts, method, params...)
}

// BALLOTDURATIONMAXNAME is a free data retrieval call binding the contract method 0x0015a73b.
//
// Solidity: function BALLOT_DURATION_MAX_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BALLOTDURATIONMAXNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BALLOT_DURATION_MAX_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTDURATIONMAXNAME is a free data retrieval call binding the contract method 0x0015a73b.
//
// Solidity: function BALLOT_DURATION_MAX_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BALLOTDURATIONMAXNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BALLOTDURATIONMAXNAME(&_EnvStorageImp.CallOpts)
}

// BALLOTDURATIONMAXNAME is a free data retrieval call binding the contract method 0x0015a73b.
//
// Solidity: function BALLOT_DURATION_MAX_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BALLOTDURATIONMAXNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BALLOTDURATIONMAXNAME(&_EnvStorageImp.CallOpts)
}

// BALLOTDURATIONMINMAXNAME is a free data retrieval call binding the contract method 0x656e3052.
//
// Solidity: function BALLOT_DURATION_MIN_MAX_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BALLOTDURATIONMINMAXNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BALLOT_DURATION_MIN_MAX_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTDURATIONMINMAXNAME is a free data retrieval call binding the contract method 0x656e3052.
//
// Solidity: function BALLOT_DURATION_MIN_MAX_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BALLOTDURATIONMINMAXNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BALLOTDURATIONMINMAXNAME(&_EnvStorageImp.CallOpts)
}

// BALLOTDURATIONMINMAXNAME is a free data retrieval call binding the contract method 0x656e3052.
//
// Solidity: function BALLOT_DURATION_MIN_MAX_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BALLOTDURATIONMINMAXNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BALLOTDURATIONMINMAXNAME(&_EnvStorageImp.CallOpts)
}

// BALLOTDURATIONMINNAME is a free data retrieval call binding the contract method 0x6d583ca7.
//
// Solidity: function BALLOT_DURATION_MIN_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BALLOTDURATIONMINNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BALLOT_DURATION_MIN_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTDURATIONMINNAME is a free data retrieval call binding the contract method 0x6d583ca7.
//
// Solidity: function BALLOT_DURATION_MIN_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BALLOTDURATIONMINNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BALLOTDURATIONMINNAME(&_EnvStorageImp.CallOpts)
}

// BALLOTDURATIONMINNAME is a free data retrieval call binding the contract method 0x6d583ca7.
//
// Solidity: function BALLOT_DURATION_MIN_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BALLOTDURATIONMINNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BALLOTDURATIONMINNAME(&_EnvStorageImp.CallOpts)
}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BALLOTSTORAGENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BALLOT_STORAGE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BALLOTSTORAGENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BALLOTSTORAGENAME(&_EnvStorageImp.CallOpts)
}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BALLOTSTORAGENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BALLOTSTORAGENAME(&_EnvStorageImp.CallOpts)
}

// BASEFEEMAXCHANGERATENAME is a free data retrieval call binding the contract method 0xc42a0abc.
//
// Solidity: function BASE_FEE_MAX_CHANGE_RATE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BASEFEEMAXCHANGERATENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BASE_FEE_MAX_CHANGE_RATE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BASEFEEMAXCHANGERATENAME is a free data retrieval call binding the contract method 0xc42a0abc.
//
// Solidity: function BASE_FEE_MAX_CHANGE_RATE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BASEFEEMAXCHANGERATENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BASEFEEMAXCHANGERATENAME(&_EnvStorageImp.CallOpts)
}

// BASEFEEMAXCHANGERATENAME is a free data retrieval call binding the contract method 0xc42a0abc.
//
// Solidity: function BASE_FEE_MAX_CHANGE_RATE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BASEFEEMAXCHANGERATENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BASEFEEMAXCHANGERATENAME(&_EnvStorageImp.CallOpts)
}

// BLOCKSPERNAME is a free data retrieval call binding the contract method 0x3f35c8fe.
//
// Solidity: function BLOCKS_PER_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BLOCKSPERNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BLOCKS_PER_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKSPERNAME is a free data retrieval call binding the contract method 0x3f35c8fe.
//
// Solidity: function BLOCKS_PER_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BLOCKSPERNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKSPERNAME(&_EnvStorageImp.CallOpts)
}

// BLOCKSPERNAME is a free data retrieval call binding the contract method 0x3f35c8fe.
//
// Solidity: function BLOCKS_PER_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BLOCKSPERNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKSPERNAME(&_EnvStorageImp.CallOpts)
}

// BLOCKCREATIONTIMENAME is a free data retrieval call binding the contract method 0xc0b4fe15.
//
// Solidity: function BLOCK_CREATION_TIME_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BLOCKCREATIONTIMENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BLOCK_CREATION_TIME_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKCREATIONTIMENAME is a free data retrieval call binding the contract method 0xc0b4fe15.
//
// Solidity: function BLOCK_CREATION_TIME_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BLOCKCREATIONTIMENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKCREATIONTIMENAME(&_EnvStorageImp.CallOpts)
}

// BLOCKCREATIONTIMENAME is a free data retrieval call binding the contract method 0xc0b4fe15.
//
// Solidity: function BLOCK_CREATION_TIME_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BLOCKCREATIONTIMENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKCREATIONTIMENAME(&_EnvStorageImp.CallOpts)
}

// BLOCKGASLIMITNAME is a free data retrieval call binding the contract method 0x238737b6.
//
// Solidity: function BLOCK_GASLIMIT_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BLOCKGASLIMITNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BLOCK_GASLIMIT_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKGASLIMITNAME is a free data retrieval call binding the contract method 0x238737b6.
//
// Solidity: function BLOCK_GASLIMIT_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BLOCKGASLIMITNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKGASLIMITNAME(&_EnvStorageImp.CallOpts)
}

// BLOCKGASLIMITNAME is a free data retrieval call binding the contract method 0x238737b6.
//
// Solidity: function BLOCK_GASLIMIT_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BLOCKGASLIMITNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKGASLIMITNAME(&_EnvStorageImp.CallOpts)
}

// BLOCKREWARDAMOUNTNAME is a free data retrieval call binding the contract method 0xa9b629b2.
//
// Solidity: function BLOCK_REWARD_AMOUNT_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BLOCKREWARDAMOUNTNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BLOCK_REWARD_AMOUNT_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREWARDAMOUNTNAME is a free data retrieval call binding the contract method 0xa9b629b2.
//
// Solidity: function BLOCK_REWARD_AMOUNT_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BLOCKREWARDAMOUNTNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKREWARDAMOUNTNAME(&_EnvStorageImp.CallOpts)
}

// BLOCKREWARDAMOUNTNAME is a free data retrieval call binding the contract method 0xa9b629b2.
//
// Solidity: function BLOCK_REWARD_AMOUNT_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BLOCKREWARDAMOUNTNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKREWARDAMOUNTNAME(&_EnvStorageImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME is a free data retrieval call binding the contract method 0xc6713baf.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME is a free data retrieval call binding the contract method 0xc6713baf.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME(&_EnvStorageImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME is a free data retrieval call binding the contract method 0xc6713baf.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME(&_EnvStorageImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME is a free data retrieval call binding the contract method 0x7b2bfb01.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME is a free data retrieval call binding the contract method 0x7b2bfb01.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME(&_EnvStorageImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME is a free data retrieval call binding the contract method 0x7b2bfb01.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME(&_EnvStorageImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONMAINTENANCENAME is a free data retrieval call binding the contract method 0x46946416.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BLOCKREWARDDISTRIBUTIONMAINTENANCENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREWARDDISTRIBUTIONMAINTENANCENAME is a free data retrieval call binding the contract method 0x46946416.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BLOCKREWARDDISTRIBUTIONMAINTENANCENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKREWARDDISTRIBUTIONMAINTENANCENAME(&_EnvStorageImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONMAINTENANCENAME is a free data retrieval call binding the contract method 0x46946416.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BLOCKREWARDDISTRIBUTIONMAINTENANCENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKREWARDDISTRIBUTIONMAINTENANCENAME(&_EnvStorageImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONMETHODNAME is a free data retrieval call binding the contract method 0x278bb12a.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_METHOD_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BLOCKREWARDDISTRIBUTIONMETHODNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BLOCK_REWARD_DISTRIBUTION_METHOD_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREWARDDISTRIBUTIONMETHODNAME is a free data retrieval call binding the contract method 0x278bb12a.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_METHOD_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BLOCKREWARDDISTRIBUTIONMETHODNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKREWARDDISTRIBUTIONMETHODNAME(&_EnvStorageImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONMETHODNAME is a free data retrieval call binding the contract method 0x278bb12a.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_METHOD_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BLOCKREWARDDISTRIBUTIONMETHODNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKREWARDDISTRIBUTIONMETHODNAME(&_EnvStorageImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME is a free data retrieval call binding the contract method 0x6167eb45.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME is a free data retrieval call binding the contract method 0x6167eb45.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME(&_EnvStorageImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME is a free data retrieval call binding the contract method 0x6167eb45.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME(&_EnvStorageImp.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCaller) DENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpSession) DENOMINATOR() (*big.Int, error) {
	return _EnvStorageImp.Contract.DENOMINATOR(&_EnvStorageImp.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) DENOMINATOR() (*big.Int, error) {
	return _EnvStorageImp.Contract.DENOMINATOR(&_EnvStorageImp.CallOpts)
}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) ECOSYSTEMNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "ECOSYSTEM_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) ECOSYSTEMNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.ECOSYSTEMNAME(&_EnvStorageImp.CallOpts)
}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) ECOSYSTEMNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.ECOSYSTEMNAME(&_EnvStorageImp.CallOpts)
}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) ENVSTORAGENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "ENV_STORAGE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) ENVSTORAGENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.ENVSTORAGENAME(&_EnvStorageImp.CallOpts)
}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) ENVSTORAGENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.ENVSTORAGENAME(&_EnvStorageImp.CallOpts)
}

// GASLIMITANDBASEFEENAME is a free data retrieval call binding the contract method 0xc7d3da34.
//
// Solidity: function GASLIMIT_AND_BASE_FEE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) GASLIMITANDBASEFEENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "GASLIMIT_AND_BASE_FEE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GASLIMITANDBASEFEENAME is a free data retrieval call binding the contract method 0xc7d3da34.
//
// Solidity: function GASLIMIT_AND_BASE_FEE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) GASLIMITANDBASEFEENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.GASLIMITANDBASEFEENAME(&_EnvStorageImp.CallOpts)
}

// GASLIMITANDBASEFEENAME is a free data retrieval call binding the contract method 0xc7d3da34.
//
// Solidity: function GASLIMIT_AND_BASE_FEE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) GASLIMITANDBASEFEENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.GASLIMITANDBASEFEENAME(&_EnvStorageImp.CallOpts)
}

// GASTARGETPERCENTAGENAME is a free data retrieval call binding the contract method 0x4d273e28.
//
// Solidity: function GAS_TARGET_PERCENTAGE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) GASTARGETPERCENTAGENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "GAS_TARGET_PERCENTAGE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GASTARGETPERCENTAGENAME is a free data retrieval call binding the contract method 0x4d273e28.
//
// Solidity: function GAS_TARGET_PERCENTAGE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) GASTARGETPERCENTAGENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.GASTARGETPERCENTAGENAME(&_EnvStorageImp.CallOpts)
}

// GASTARGETPERCENTAGENAME is a free data retrieval call binding the contract method 0x4d273e28.
//
// Solidity: function GAS_TARGET_PERCENTAGE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) GASTARGETPERCENTAGENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.GASTARGETPERCENTAGENAME(&_EnvStorageImp.CallOpts)
}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) GOVNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "GOV_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) GOVNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.GOVNAME(&_EnvStorageImp.CallOpts)
}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) GOVNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.GOVNAME(&_EnvStorageImp.CallOpts)
}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) MAINTENANCENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "MAINTENANCE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) MAINTENANCENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.MAINTENANCENAME(&_EnvStorageImp.CallOpts)
}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) MAINTENANCENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.MAINTENANCENAME(&_EnvStorageImp.CallOpts)
}

// MAXBASEFEENAME is a free data retrieval call binding the contract method 0xf38ecf47.
//
// Solidity: function MAX_BASE_FEE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) MAXBASEFEENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "MAX_BASE_FEE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAXBASEFEENAME is a free data retrieval call binding the contract method 0xf38ecf47.
//
// Solidity: function MAX_BASE_FEE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) MAXBASEFEENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.MAXBASEFEENAME(&_EnvStorageImp.CallOpts)
}

// MAXBASEFEENAME is a free data retrieval call binding the contract method 0xf38ecf47.
//
// Solidity: function MAX_BASE_FEE_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) MAXBASEFEENAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.MAXBASEFEENAME(&_EnvStorageImp.CallOpts)
}

// MAXIDLEBLOCKINTERVALNAME is a free data retrieval call binding the contract method 0xb128f880.
//
// Solidity: function MAX_IDLE_BLOCK_INTERVAL_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) MAXIDLEBLOCKINTERVALNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "MAX_IDLE_BLOCK_INTERVAL_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAXIDLEBLOCKINTERVALNAME is a free data retrieval call binding the contract method 0xb128f880.
//
// Solidity: function MAX_IDLE_BLOCK_INTERVAL_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) MAXIDLEBLOCKINTERVALNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.MAXIDLEBLOCKINTERVALNAME(&_EnvStorageImp.CallOpts)
}

// MAXIDLEBLOCKINTERVALNAME is a free data retrieval call binding the contract method 0xb128f880.
//
// Solidity: function MAX_IDLE_BLOCK_INTERVAL_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) MAXIDLEBLOCKINTERVALNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.MAXIDLEBLOCKINTERVALNAME(&_EnvStorageImp.CallOpts)
}

// MAXPRIORITYFEEPERGASNAME is a free data retrieval call binding the contract method 0x38294419.
//
// Solidity: function MAX_PRIORITY_FEE_PER_GAS_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) MAXPRIORITYFEEPERGASNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "MAX_PRIORITY_FEE_PER_GAS_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAXPRIORITYFEEPERGASNAME is a free data retrieval call binding the contract method 0x38294419.
//
// Solidity: function MAX_PRIORITY_FEE_PER_GAS_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) MAXPRIORITYFEEPERGASNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.MAXPRIORITYFEEPERGASNAME(&_EnvStorageImp.CallOpts)
}

// MAXPRIORITYFEEPERGASNAME is a free data retrieval call binding the contract method 0x38294419.
//
// Solidity: function MAX_PRIORITY_FEE_PER_GAS_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) MAXPRIORITYFEEPERGASNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.MAXPRIORITYFEEPERGASNAME(&_EnvStorageImp.CallOpts)
}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) REWARDPOOLNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "REWARD_POOL_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) REWARDPOOLNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.REWARDPOOLNAME(&_EnvStorageImp.CallOpts)
}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) REWARDPOOLNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.REWARDPOOLNAME(&_EnvStorageImp.CallOpts)
}

// STAKINGMAXNAME is a free data retrieval call binding the contract method 0xc00ace6c.
//
// Solidity: function STAKING_MAX_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) STAKINGMAXNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "STAKING_MAX_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMAXNAME is a free data retrieval call binding the contract method 0xc00ace6c.
//
// Solidity: function STAKING_MAX_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) STAKINGMAXNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.STAKINGMAXNAME(&_EnvStorageImp.CallOpts)
}

// STAKINGMAXNAME is a free data retrieval call binding the contract method 0xc00ace6c.
//
// Solidity: function STAKING_MAX_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) STAKINGMAXNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.STAKINGMAXNAME(&_EnvStorageImp.CallOpts)
}

// STAKINGMINMAXNAME is a free data retrieval call binding the contract method 0xa6868b7d.
//
// Solidity: function STAKING_MIN_MAX_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) STAKINGMINMAXNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "STAKING_MIN_MAX_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMINMAXNAME is a free data retrieval call binding the contract method 0xa6868b7d.
//
// Solidity: function STAKING_MIN_MAX_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) STAKINGMINMAXNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.STAKINGMINMAXNAME(&_EnvStorageImp.CallOpts)
}

// STAKINGMINMAXNAME is a free data retrieval call binding the contract method 0xa6868b7d.
//
// Solidity: function STAKING_MIN_MAX_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) STAKINGMINMAXNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.STAKINGMINMAXNAME(&_EnvStorageImp.CallOpts)
}

// STAKINGMINNAME is a free data retrieval call binding the contract method 0x6fde207a.
//
// Solidity: function STAKING_MIN_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) STAKINGMINNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "STAKING_MIN_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMINNAME is a free data retrieval call binding the contract method 0x6fde207a.
//
// Solidity: function STAKING_MIN_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) STAKINGMINNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.STAKINGMINNAME(&_EnvStorageImp.CallOpts)
}

// STAKINGMINNAME is a free data retrieval call binding the contract method 0x6fde207a.
//
// Solidity: function STAKING_MIN_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) STAKINGMINNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.STAKINGMINNAME(&_EnvStorageImp.CallOpts)
}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) STAKINGNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "STAKING_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) STAKINGNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.STAKINGNAME(&_EnvStorageImp.CallOpts)
}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) STAKINGNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.STAKINGNAME(&_EnvStorageImp.CallOpts)
}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) STAKINGREWARDNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "STAKING_REWARD_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) STAKINGREWARDNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.STAKINGREWARDNAME(&_EnvStorageImp.CallOpts)
}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) STAKINGREWARDNAME() ([32]byte, error) {
	return _EnvStorageImp.Contract.STAKINGREWARDNAME(&_EnvStorageImp.CallOpts)
}

// CheckVariableCondition is a free data retrieval call binding the contract method 0x9801bff9.
//
// Solidity: function checkVariableCondition(bytes32 envKey, bytes envVal) pure returns(bool)
func (_EnvStorageImp *EnvStorageImpCaller) CheckVariableCondition(opts *bind.CallOpts, envKey [32]byte, envVal []byte) (bool, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "checkVariableCondition", envKey, envVal)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckVariableCondition is a free data retrieval call binding the contract method 0x9801bff9.
//
// Solidity: function checkVariableCondition(bytes32 envKey, bytes envVal) pure returns(bool)
func (_EnvStorageImp *EnvStorageImpSession) CheckVariableCondition(envKey [32]byte, envVal []byte) (bool, error) {
	return _EnvStorageImp.Contract.CheckVariableCondition(&_EnvStorageImp.CallOpts, envKey, envVal)
}

// CheckVariableCondition is a free data retrieval call binding the contract method 0x9801bff9.
//
// Solidity: function checkVariableCondition(bytes32 envKey, bytes envVal) pure returns(bool)
func (_EnvStorageImp *EnvStorageImpCallerSession) CheckVariableCondition(envKey [32]byte, envVal []byte) (bool, error) {
	return _EnvStorageImp.Contract.CheckVariableCondition(&_EnvStorageImp.CallOpts, envKey, envVal)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 h) view returns(address)
func (_EnvStorageImp *EnvStorageImpCaller) GetAddress(opts *bind.CallOpts, h [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getAddress", h)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 h) view returns(address)
func (_EnvStorageImp *EnvStorageImpSession) GetAddress(h [32]byte) (common.Address, error) {
	return _EnvStorageImp.Contract.GetAddress(&_EnvStorageImp.CallOpts, h)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 h) view returns(address)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetAddress(h [32]byte) (common.Address, error) {
	return _EnvStorageImp.Contract.GetAddress(&_EnvStorageImp.CallOpts, h)
}

// GetBallotDurationMax is a free data retrieval call binding the contract method 0x1b27e01b.
//
// Solidity: function getBallotDurationMax() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetBallotDurationMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getBallotDurationMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBallotDurationMax is a free data retrieval call binding the contract method 0x1b27e01b.
//
// Solidity: function getBallotDurationMax() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetBallotDurationMax() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetBallotDurationMax(&_EnvStorageImp.CallOpts)
}

// GetBallotDurationMax is a free data retrieval call binding the contract method 0x1b27e01b.
//
// Solidity: function getBallotDurationMax() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetBallotDurationMax() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetBallotDurationMax(&_EnvStorageImp.CallOpts)
}

// GetBallotDurationMin is a free data retrieval call binding the contract method 0x33be496e.
//
// Solidity: function getBallotDurationMin() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetBallotDurationMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getBallotDurationMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBallotDurationMin is a free data retrieval call binding the contract method 0x33be496e.
//
// Solidity: function getBallotDurationMin() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetBallotDurationMin() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetBallotDurationMin(&_EnvStorageImp.CallOpts)
}

// GetBallotDurationMin is a free data retrieval call binding the contract method 0x33be496e.
//
// Solidity: function getBallotDurationMin() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetBallotDurationMin() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetBallotDurationMin(&_EnvStorageImp.CallOpts)
}

// GetBallotDurationMinMax is a free data retrieval call binding the contract method 0xcdb9f643.
//
// Solidity: function getBallotDurationMinMax() view returns(uint256, uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetBallotDurationMinMax(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getBallotDurationMinMax")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetBallotDurationMinMax is a free data retrieval call binding the contract method 0xcdb9f643.
//
// Solidity: function getBallotDurationMinMax() view returns(uint256, uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetBallotDurationMinMax() (*big.Int, *big.Int, error) {
	return _EnvStorageImp.Contract.GetBallotDurationMinMax(&_EnvStorageImp.CallOpts)
}

// GetBallotDurationMinMax is a free data retrieval call binding the contract method 0xcdb9f643.
//
// Solidity: function getBallotDurationMinMax() view returns(uint256, uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetBallotDurationMinMax() (*big.Int, *big.Int, error) {
	return _EnvStorageImp.Contract.GetBallotDurationMinMax(&_EnvStorageImp.CallOpts)
}

// GetBlockCreationTime is a free data retrieval call binding the contract method 0x33e31184.
//
// Solidity: function getBlockCreationTime() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetBlockCreationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getBlockCreationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockCreationTime is a free data retrieval call binding the contract method 0x33e31184.
//
// Solidity: function getBlockCreationTime() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetBlockCreationTime() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetBlockCreationTime(&_EnvStorageImp.CallOpts)
}

// GetBlockCreationTime is a free data retrieval call binding the contract method 0x33e31184.
//
// Solidity: function getBlockCreationTime() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetBlockCreationTime() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetBlockCreationTime(&_EnvStorageImp.CallOpts)
}

// GetBlockRewardAmount is a free data retrieval call binding the contract method 0x5dba8c4a.
//
// Solidity: function getBlockRewardAmount() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetBlockRewardAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getBlockRewardAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockRewardAmount is a free data retrieval call binding the contract method 0x5dba8c4a.
//
// Solidity: function getBlockRewardAmount() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetBlockRewardAmount() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetBlockRewardAmount(&_EnvStorageImp.CallOpts)
}

// GetBlockRewardAmount is a free data retrieval call binding the contract method 0x5dba8c4a.
//
// Solidity: function getBlockRewardAmount() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetBlockRewardAmount() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetBlockRewardAmount(&_EnvStorageImp.CallOpts)
}

// GetBlockRewardDistributionMethod is a free data retrieval call binding the contract method 0x8f6fe725.
//
// Solidity: function getBlockRewardDistributionMethod() view returns(uint256, uint256, uint256, uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetBlockRewardDistributionMethod(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getBlockRewardDistributionMethod")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetBlockRewardDistributionMethod is a free data retrieval call binding the contract method 0x8f6fe725.
//
// Solidity: function getBlockRewardDistributionMethod() view returns(uint256, uint256, uint256, uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetBlockRewardDistributionMethod() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _EnvStorageImp.Contract.GetBlockRewardDistributionMethod(&_EnvStorageImp.CallOpts)
}

// GetBlockRewardDistributionMethod is a free data retrieval call binding the contract method 0x8f6fe725.
//
// Solidity: function getBlockRewardDistributionMethod() view returns(uint256, uint256, uint256, uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetBlockRewardDistributionMethod() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _EnvStorageImp.Contract.GetBlockRewardDistributionMethod(&_EnvStorageImp.CallOpts)
}

// GetBlocksPer is a free data retrieval call binding the contract method 0x3690057a.
//
// Solidity: function getBlocksPer() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetBlocksPer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getBlocksPer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlocksPer is a free data retrieval call binding the contract method 0x3690057a.
//
// Solidity: function getBlocksPer() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetBlocksPer() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetBlocksPer(&_EnvStorageImp.CallOpts)
}

// GetBlocksPer is a free data retrieval call binding the contract method 0x3690057a.
//
// Solidity: function getBlocksPer() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetBlocksPer() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetBlocksPer(&_EnvStorageImp.CallOpts)
}

// GetBoolean is a free data retrieval call binding the contract method 0x3848207a.
//
// Solidity: function getBoolean(bytes32 h) view returns(bool)
func (_EnvStorageImp *EnvStorageImpCaller) GetBoolean(opts *bind.CallOpts, h [32]byte) (bool, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getBoolean", h)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBoolean is a free data retrieval call binding the contract method 0x3848207a.
//
// Solidity: function getBoolean(bytes32 h) view returns(bool)
func (_EnvStorageImp *EnvStorageImpSession) GetBoolean(h [32]byte) (bool, error) {
	return _EnvStorageImp.Contract.GetBoolean(&_EnvStorageImp.CallOpts, h)
}

// GetBoolean is a free data retrieval call binding the contract method 0x3848207a.
//
// Solidity: function getBoolean(bytes32 h) view returns(bool)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetBoolean(h [32]byte) (bool, error) {
	return _EnvStorageImp.Contract.GetBoolean(&_EnvStorageImp.CallOpts, h)
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(bytes32 h) view returns(bytes)
func (_EnvStorageImp *EnvStorageImpCaller) GetBytes(opts *bind.CallOpts, h [32]byte) ([]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getBytes", h)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(bytes32 h) view returns(bytes)
func (_EnvStorageImp *EnvStorageImpSession) GetBytes(h [32]byte) ([]byte, error) {
	return _EnvStorageImp.Contract.GetBytes(&_EnvStorageImp.CallOpts, h)
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(bytes32 h) view returns(bytes)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetBytes(h [32]byte) ([]byte, error) {
	return _EnvStorageImp.Contract.GetBytes(&_EnvStorageImp.CallOpts, h)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 h) view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) GetBytes32(opts *bind.CallOpts, h [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getBytes32", h)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 h) view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) GetBytes32(h [32]byte) ([32]byte, error) {
	return _EnvStorageImp.Contract.GetBytes32(&_EnvStorageImp.CallOpts, h)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 h) view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetBytes32(h [32]byte) ([32]byte, error) {
	return _EnvStorageImp.Contract.GetBytes32(&_EnvStorageImp.CallOpts, h)
}

// GetGasLimitAndBaseFee is a free data retrieval call binding the contract method 0x67d1a2e0.
//
// Solidity: function getGasLimitAndBaseFee() view returns(uint256, uint256, uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetGasLimitAndBaseFee(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getGasLimitAndBaseFee")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetGasLimitAndBaseFee is a free data retrieval call binding the contract method 0x67d1a2e0.
//
// Solidity: function getGasLimitAndBaseFee() view returns(uint256, uint256, uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetGasLimitAndBaseFee() (*big.Int, *big.Int, *big.Int, error) {
	return _EnvStorageImp.Contract.GetGasLimitAndBaseFee(&_EnvStorageImp.CallOpts)
}

// GetGasLimitAndBaseFee is a free data retrieval call binding the contract method 0x67d1a2e0.
//
// Solidity: function getGasLimitAndBaseFee() view returns(uint256, uint256, uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetGasLimitAndBaseFee() (*big.Int, *big.Int, *big.Int, error) {
	return _EnvStorageImp.Contract.GetGasLimitAndBaseFee(&_EnvStorageImp.CallOpts)
}

// GetInt is a free data retrieval call binding the contract method 0xdc97d962.
//
// Solidity: function getInt(bytes32 h) view returns(int256)
func (_EnvStorageImp *EnvStorageImpCaller) GetInt(opts *bind.CallOpts, h [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getInt", h)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInt is a free data retrieval call binding the contract method 0xdc97d962.
//
// Solidity: function getInt(bytes32 h) view returns(int256)
func (_EnvStorageImp *EnvStorageImpSession) GetInt(h [32]byte) (*big.Int, error) {
	return _EnvStorageImp.Contract.GetInt(&_EnvStorageImp.CallOpts, h)
}

// GetInt is a free data retrieval call binding the contract method 0xdc97d962.
//
// Solidity: function getInt(bytes32 h) view returns(int256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetInt(h [32]byte) (*big.Int, error) {
	return _EnvStorageImp.Contract.GetInt(&_EnvStorageImp.CallOpts, h)
}

// GetMaxBaseFee is a free data retrieval call binding the contract method 0xc95437be.
//
// Solidity: function getMaxBaseFee() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetMaxBaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getMaxBaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxBaseFee is a free data retrieval call binding the contract method 0xc95437be.
//
// Solidity: function getMaxBaseFee() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetMaxBaseFee() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetMaxBaseFee(&_EnvStorageImp.CallOpts)
}

// GetMaxBaseFee is a free data retrieval call binding the contract method 0xc95437be.
//
// Solidity: function getMaxBaseFee() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetMaxBaseFee() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetMaxBaseFee(&_EnvStorageImp.CallOpts)
}

// GetMaxIdleBlockInterval is a free data retrieval call binding the contract method 0x185582f1.
//
// Solidity: function getMaxIdleBlockInterval() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetMaxIdleBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getMaxIdleBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxIdleBlockInterval is a free data retrieval call binding the contract method 0x185582f1.
//
// Solidity: function getMaxIdleBlockInterval() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetMaxIdleBlockInterval() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetMaxIdleBlockInterval(&_EnvStorageImp.CallOpts)
}

// GetMaxIdleBlockInterval is a free data retrieval call binding the contract method 0x185582f1.
//
// Solidity: function getMaxIdleBlockInterval() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetMaxIdleBlockInterval() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetMaxIdleBlockInterval(&_EnvStorageImp.CallOpts)
}

// GetMaxPriorityFeePerGas is a free data retrieval call binding the contract method 0x2b2eaa92.
//
// Solidity: function getMaxPriorityFeePerGas() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetMaxPriorityFeePerGas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getMaxPriorityFeePerGas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPriorityFeePerGas is a free data retrieval call binding the contract method 0x2b2eaa92.
//
// Solidity: function getMaxPriorityFeePerGas() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetMaxPriorityFeePerGas() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetMaxPriorityFeePerGas(&_EnvStorageImp.CallOpts)
}

// GetMaxPriorityFeePerGas is a free data retrieval call binding the contract method 0x2b2eaa92.
//
// Solidity: function getMaxPriorityFeePerGas() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetMaxPriorityFeePerGas() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetMaxPriorityFeePerGas(&_EnvStorageImp.CallOpts)
}

// GetStakingMax is a free data retrieval call binding the contract method 0x737c59b8.
//
// Solidity: function getStakingMax() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetStakingMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getStakingMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingMax is a free data retrieval call binding the contract method 0x737c59b8.
//
// Solidity: function getStakingMax() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetStakingMax() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetStakingMax(&_EnvStorageImp.CallOpts)
}

// GetStakingMax is a free data retrieval call binding the contract method 0x737c59b8.
//
// Solidity: function getStakingMax() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetStakingMax() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetStakingMax(&_EnvStorageImp.CallOpts)
}

// GetStakingMin is a free data retrieval call binding the contract method 0x076cd77f.
//
// Solidity: function getStakingMin() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetStakingMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getStakingMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingMin is a free data retrieval call binding the contract method 0x076cd77f.
//
// Solidity: function getStakingMin() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetStakingMin() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetStakingMin(&_EnvStorageImp.CallOpts)
}

// GetStakingMin is a free data retrieval call binding the contract method 0x076cd77f.
//
// Solidity: function getStakingMin() view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetStakingMin() (*big.Int, error) {
	return _EnvStorageImp.Contract.GetStakingMin(&_EnvStorageImp.CallOpts)
}

// GetStakingMinMax is a free data retrieval call binding the contract method 0xaba01fee.
//
// Solidity: function getStakingMinMax() view returns(uint256, uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetStakingMinMax(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getStakingMinMax")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStakingMinMax is a free data retrieval call binding the contract method 0xaba01fee.
//
// Solidity: function getStakingMinMax() view returns(uint256, uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetStakingMinMax() (*big.Int, *big.Int, error) {
	return _EnvStorageImp.Contract.GetStakingMinMax(&_EnvStorageImp.CallOpts)
}

// GetStakingMinMax is a free data retrieval call binding the contract method 0xaba01fee.
//
// Solidity: function getStakingMinMax() view returns(uint256, uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetStakingMinMax() (*big.Int, *big.Int, error) {
	return _EnvStorageImp.Contract.GetStakingMinMax(&_EnvStorageImp.CallOpts)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(bytes32 h) view returns(string)
func (_EnvStorageImp *EnvStorageImpCaller) GetString(opts *bind.CallOpts, h [32]byte) (string, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getString", h)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(bytes32 h) view returns(string)
func (_EnvStorageImp *EnvStorageImpSession) GetString(h [32]byte) (string, error) {
	return _EnvStorageImp.Contract.GetString(&_EnvStorageImp.CallOpts, h)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(bytes32 h) view returns(string)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetString(h [32]byte) (string, error) {
	return _EnvStorageImp.Contract.GetString(&_EnvStorageImp.CallOpts, h)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 h) view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCaller) GetUint(opts *bind.CallOpts, h [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "getUint", h)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 h) view returns(uint256)
func (_EnvStorageImp *EnvStorageImpSession) GetUint(h [32]byte) (*big.Int, error) {
	return _EnvStorageImp.Contract.GetUint(&_EnvStorageImp.CallOpts, h)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 h) view returns(uint256)
func (_EnvStorageImp *EnvStorageImpCallerSession) GetUint(h [32]byte) (*big.Int, error) {
	return _EnvStorageImp.Contract.GetUint(&_EnvStorageImp.CallOpts, h)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnvStorageImp *EnvStorageImpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnvStorageImp *EnvStorageImpSession) Owner() (common.Address, error) {
	return _EnvStorageImp.Contract.Owner(&_EnvStorageImp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnvStorageImp *EnvStorageImpCallerSession) Owner() (common.Address, error) {
	return _EnvStorageImp.Contract.Owner(&_EnvStorageImp.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpSession) ProxiableUUID() ([32]byte, error) {
	return _EnvStorageImp.Contract.ProxiableUUID(&_EnvStorageImp.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EnvStorageImp *EnvStorageImpCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EnvStorageImp.Contract.ProxiableUUID(&_EnvStorageImp.CallOpts)
}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_EnvStorageImp *EnvStorageImpCaller) Reg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnvStorageImp.contract.Call(opts, &out, "reg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_EnvStorageImp *EnvStorageImpSession) Reg() (common.Address, error) {
	return _EnvStorageImp.Contract.Reg(&_EnvStorageImp.CallOpts)
}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_EnvStorageImp *EnvStorageImpCallerSession) Reg() (common.Address, error) {
	return _EnvStorageImp.Contract.Reg(&_EnvStorageImp.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x539927be.
//
// Solidity: function initialize(address _registry, bytes32[] names, uint256[] infos) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, names [][32]byte, infos []*big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "initialize", _registry, names, infos)
}

// Initialize is a paid mutator transaction binding the contract method 0x539927be.
//
// Solidity: function initialize(address _registry, bytes32[] names, uint256[] infos) returns()
func (_EnvStorageImp *EnvStorageImpSession) Initialize(_registry common.Address, names [][32]byte, infos []*big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.Initialize(&_EnvStorageImp.TransactOpts, _registry, names, infos)
}

// Initialize is a paid mutator transaction binding the contract method 0x539927be.
//
// Solidity: function initialize(address _registry, bytes32[] names, uint256[] infos) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) Initialize(_registry common.Address, names [][32]byte, infos []*big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.Initialize(&_EnvStorageImp.TransactOpts, _registry, names, infos)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnvStorageImp *EnvStorageImpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnvStorageImp *EnvStorageImpSession) RenounceOwnership() (*types.Transaction, error) {
	return _EnvStorageImp.Contract.RenounceOwnership(&_EnvStorageImp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EnvStorageImp.Contract.RenounceOwnership(&_EnvStorageImp.TransactOpts)
}

// SetBallotDurationMax is a paid mutator transaction binding the contract method 0x852cf38f.
//
// Solidity: function setBallotDurationMax(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBallotDurationMax(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBallotDurationMax", _value)
}

// SetBallotDurationMax is a paid mutator transaction binding the contract method 0x852cf38f.
//
// Solidity: function setBallotDurationMax(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBallotDurationMax(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBallotDurationMax(&_EnvStorageImp.TransactOpts, _value)
}

// SetBallotDurationMax is a paid mutator transaction binding the contract method 0x852cf38f.
//
// Solidity: function setBallotDurationMax(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBallotDurationMax(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBallotDurationMax(&_EnvStorageImp.TransactOpts, _value)
}

// SetBallotDurationMaxByBytes is a paid mutator transaction binding the contract method 0xa5d6a581.
//
// Solidity: function setBallotDurationMaxByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBallotDurationMaxByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBallotDurationMaxByBytes", _value)
}

// SetBallotDurationMaxByBytes is a paid mutator transaction binding the contract method 0xa5d6a581.
//
// Solidity: function setBallotDurationMaxByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBallotDurationMaxByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBallotDurationMaxByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetBallotDurationMaxByBytes is a paid mutator transaction binding the contract method 0xa5d6a581.
//
// Solidity: function setBallotDurationMaxByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBallotDurationMaxByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBallotDurationMaxByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetBallotDurationMin is a paid mutator transaction binding the contract method 0xa36f259b.
//
// Solidity: function setBallotDurationMin(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBallotDurationMin(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBallotDurationMin", _value)
}

// SetBallotDurationMin is a paid mutator transaction binding the contract method 0xa36f259b.
//
// Solidity: function setBallotDurationMin(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBallotDurationMin(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBallotDurationMin(&_EnvStorageImp.TransactOpts, _value)
}

// SetBallotDurationMin is a paid mutator transaction binding the contract method 0xa36f259b.
//
// Solidity: function setBallotDurationMin(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBallotDurationMin(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBallotDurationMin(&_EnvStorageImp.TransactOpts, _value)
}

// SetBallotDurationMinByBytes is a paid mutator transaction binding the contract method 0x124cea37.
//
// Solidity: function setBallotDurationMinByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBallotDurationMinByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBallotDurationMinByBytes", _value)
}

// SetBallotDurationMinByBytes is a paid mutator transaction binding the contract method 0x124cea37.
//
// Solidity: function setBallotDurationMinByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBallotDurationMinByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBallotDurationMinByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetBallotDurationMinByBytes is a paid mutator transaction binding the contract method 0x124cea37.
//
// Solidity: function setBallotDurationMinByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBallotDurationMinByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBallotDurationMinByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetBallotDurationMinMax is a paid mutator transaction binding the contract method 0xb4d6bd3b.
//
// Solidity: function setBallotDurationMinMax(uint256 _min, uint256 _max) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBallotDurationMinMax(opts *bind.TransactOpts, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBallotDurationMinMax", _min, _max)
}

// SetBallotDurationMinMax is a paid mutator transaction binding the contract method 0xb4d6bd3b.
//
// Solidity: function setBallotDurationMinMax(uint256 _min, uint256 _max) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBallotDurationMinMax(_min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBallotDurationMinMax(&_EnvStorageImp.TransactOpts, _min, _max)
}

// SetBallotDurationMinMax is a paid mutator transaction binding the contract method 0xb4d6bd3b.
//
// Solidity: function setBallotDurationMinMax(uint256 _min, uint256 _max) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBallotDurationMinMax(_min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBallotDurationMinMax(&_EnvStorageImp.TransactOpts, _min, _max)
}

// SetBallotDurationMinMaxByBytes is a paid mutator transaction binding the contract method 0x45b5ec29.
//
// Solidity: function setBallotDurationMinMaxByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBallotDurationMinMaxByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBallotDurationMinMaxByBytes", _value)
}

// SetBallotDurationMinMaxByBytes is a paid mutator transaction binding the contract method 0x45b5ec29.
//
// Solidity: function setBallotDurationMinMaxByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBallotDurationMinMaxByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBallotDurationMinMaxByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetBallotDurationMinMaxByBytes is a paid mutator transaction binding the contract method 0x45b5ec29.
//
// Solidity: function setBallotDurationMinMaxByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBallotDurationMinMaxByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBallotDurationMinMaxByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlockCreationTime is a paid mutator transaction binding the contract method 0x3305508e.
//
// Solidity: function setBlockCreationTime(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBlockCreationTime(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBlockCreationTime", _value)
}

// SetBlockCreationTime is a paid mutator transaction binding the contract method 0x3305508e.
//
// Solidity: function setBlockCreationTime(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBlockCreationTime(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlockCreationTime(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlockCreationTime is a paid mutator transaction binding the contract method 0x3305508e.
//
// Solidity: function setBlockCreationTime(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBlockCreationTime(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlockCreationTime(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlockCreationTimeByBytes is a paid mutator transaction binding the contract method 0x44b89914.
//
// Solidity: function setBlockCreationTimeByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBlockCreationTimeByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBlockCreationTimeByBytes", _value)
}

// SetBlockCreationTimeByBytes is a paid mutator transaction binding the contract method 0x44b89914.
//
// Solidity: function setBlockCreationTimeByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBlockCreationTimeByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlockCreationTimeByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlockCreationTimeByBytes is a paid mutator transaction binding the contract method 0x44b89914.
//
// Solidity: function setBlockCreationTimeByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBlockCreationTimeByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlockCreationTimeByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlockRewardAmount is a paid mutator transaction binding the contract method 0x2ed19cd5.
//
// Solidity: function setBlockRewardAmount(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBlockRewardAmount(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBlockRewardAmount", _value)
}

// SetBlockRewardAmount is a paid mutator transaction binding the contract method 0x2ed19cd5.
//
// Solidity: function setBlockRewardAmount(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBlockRewardAmount(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlockRewardAmount(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlockRewardAmount is a paid mutator transaction binding the contract method 0x2ed19cd5.
//
// Solidity: function setBlockRewardAmount(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBlockRewardAmount(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlockRewardAmount(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlockRewardAmountByBytes is a paid mutator transaction binding the contract method 0x3d4c65f3.
//
// Solidity: function setBlockRewardAmountByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBlockRewardAmountByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBlockRewardAmountByBytes", _value)
}

// SetBlockRewardAmountByBytes is a paid mutator transaction binding the contract method 0x3d4c65f3.
//
// Solidity: function setBlockRewardAmountByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBlockRewardAmountByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlockRewardAmountByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlockRewardAmountByBytes is a paid mutator transaction binding the contract method 0x3d4c65f3.
//
// Solidity: function setBlockRewardAmountByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBlockRewardAmountByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlockRewardAmountByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlockRewardDistributionMethod is a paid mutator transaction binding the contract method 0x0add66dd.
//
// Solidity: function setBlockRewardDistributionMethod(uint256 _block_producer, uint256 _staking_reward, uint256 _ecofund, uint256 _maintenance) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBlockRewardDistributionMethod(opts *bind.TransactOpts, _block_producer *big.Int, _staking_reward *big.Int, _ecofund *big.Int, _maintenance *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBlockRewardDistributionMethod", _block_producer, _staking_reward, _ecofund, _maintenance)
}

// SetBlockRewardDistributionMethod is a paid mutator transaction binding the contract method 0x0add66dd.
//
// Solidity: function setBlockRewardDistributionMethod(uint256 _block_producer, uint256 _staking_reward, uint256 _ecofund, uint256 _maintenance) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBlockRewardDistributionMethod(_block_producer *big.Int, _staking_reward *big.Int, _ecofund *big.Int, _maintenance *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlockRewardDistributionMethod(&_EnvStorageImp.TransactOpts, _block_producer, _staking_reward, _ecofund, _maintenance)
}

// SetBlockRewardDistributionMethod is a paid mutator transaction binding the contract method 0x0add66dd.
//
// Solidity: function setBlockRewardDistributionMethod(uint256 _block_producer, uint256 _staking_reward, uint256 _ecofund, uint256 _maintenance) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBlockRewardDistributionMethod(_block_producer *big.Int, _staking_reward *big.Int, _ecofund *big.Int, _maintenance *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlockRewardDistributionMethod(&_EnvStorageImp.TransactOpts, _block_producer, _staking_reward, _ecofund, _maintenance)
}

// SetBlockRewardDistributionMethodByBytes is a paid mutator transaction binding the contract method 0x8c8b887e.
//
// Solidity: function setBlockRewardDistributionMethodByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBlockRewardDistributionMethodByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBlockRewardDistributionMethodByBytes", _value)
}

// SetBlockRewardDistributionMethodByBytes is a paid mutator transaction binding the contract method 0x8c8b887e.
//
// Solidity: function setBlockRewardDistributionMethodByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBlockRewardDistributionMethodByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlockRewardDistributionMethodByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlockRewardDistributionMethodByBytes is a paid mutator transaction binding the contract method 0x8c8b887e.
//
// Solidity: function setBlockRewardDistributionMethodByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBlockRewardDistributionMethodByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlockRewardDistributionMethodByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlocksPer is a paid mutator transaction binding the contract method 0x79e85859.
//
// Solidity: function setBlocksPer(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBlocksPer(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBlocksPer", _value)
}

// SetBlocksPer is a paid mutator transaction binding the contract method 0x79e85859.
//
// Solidity: function setBlocksPer(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBlocksPer(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlocksPer(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlocksPer is a paid mutator transaction binding the contract method 0x79e85859.
//
// Solidity: function setBlocksPer(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBlocksPer(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlocksPer(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlocksPerByBytes is a paid mutator transaction binding the contract method 0x3e8daafe.
//
// Solidity: function setBlocksPerByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetBlocksPerByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setBlocksPerByBytes", _value)
}

// SetBlocksPerByBytes is a paid mutator transaction binding the contract method 0x3e8daafe.
//
// Solidity: function setBlocksPerByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetBlocksPerByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlocksPerByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetBlocksPerByBytes is a paid mutator transaction binding the contract method 0x3e8daafe.
//
// Solidity: function setBlocksPerByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetBlocksPerByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetBlocksPerByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetGasLimitAndBaseFee is a paid mutator transaction binding the contract method 0x0fca11d2.
//
// Solidity: function setGasLimitAndBaseFee(uint256 _block_GasLimit, uint256 _baseFeeMaxChangeRate, uint256 _gasTargetPercentage, uint256 _maxBaseFee) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetGasLimitAndBaseFee(opts *bind.TransactOpts, _block_GasLimit *big.Int, _baseFeeMaxChangeRate *big.Int, _gasTargetPercentage *big.Int, _maxBaseFee *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setGasLimitAndBaseFee", _block_GasLimit, _baseFeeMaxChangeRate, _gasTargetPercentage, _maxBaseFee)
}

// SetGasLimitAndBaseFee is a paid mutator transaction binding the contract method 0x0fca11d2.
//
// Solidity: function setGasLimitAndBaseFee(uint256 _block_GasLimit, uint256 _baseFeeMaxChangeRate, uint256 _gasTargetPercentage, uint256 _maxBaseFee) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetGasLimitAndBaseFee(_block_GasLimit *big.Int, _baseFeeMaxChangeRate *big.Int, _gasTargetPercentage *big.Int, _maxBaseFee *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetGasLimitAndBaseFee(&_EnvStorageImp.TransactOpts, _block_GasLimit, _baseFeeMaxChangeRate, _gasTargetPercentage, _maxBaseFee)
}

// SetGasLimitAndBaseFee is a paid mutator transaction binding the contract method 0x0fca11d2.
//
// Solidity: function setGasLimitAndBaseFee(uint256 _block_GasLimit, uint256 _baseFeeMaxChangeRate, uint256 _gasTargetPercentage, uint256 _maxBaseFee) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetGasLimitAndBaseFee(_block_GasLimit *big.Int, _baseFeeMaxChangeRate *big.Int, _gasTargetPercentage *big.Int, _maxBaseFee *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetGasLimitAndBaseFee(&_EnvStorageImp.TransactOpts, _block_GasLimit, _baseFeeMaxChangeRate, _gasTargetPercentage, _maxBaseFee)
}

// SetGasLimitAndBaseFeeByBytes is a paid mutator transaction binding the contract method 0xbe33732a.
//
// Solidity: function setGasLimitAndBaseFeeByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetGasLimitAndBaseFeeByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setGasLimitAndBaseFeeByBytes", _value)
}

// SetGasLimitAndBaseFeeByBytes is a paid mutator transaction binding the contract method 0xbe33732a.
//
// Solidity: function setGasLimitAndBaseFeeByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetGasLimitAndBaseFeeByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetGasLimitAndBaseFeeByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetGasLimitAndBaseFeeByBytes is a paid mutator transaction binding the contract method 0xbe33732a.
//
// Solidity: function setGasLimitAndBaseFeeByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetGasLimitAndBaseFeeByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetGasLimitAndBaseFeeByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetMaxBaseFee is a paid mutator transaction binding the contract method 0x6fe13177.
//
// Solidity: function setMaxBaseFee(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetMaxBaseFee(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setMaxBaseFee", _value)
}

// SetMaxBaseFee is a paid mutator transaction binding the contract method 0x6fe13177.
//
// Solidity: function setMaxBaseFee(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetMaxBaseFee(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetMaxBaseFee(&_EnvStorageImp.TransactOpts, _value)
}

// SetMaxBaseFee is a paid mutator transaction binding the contract method 0x6fe13177.
//
// Solidity: function setMaxBaseFee(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetMaxBaseFee(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetMaxBaseFee(&_EnvStorageImp.TransactOpts, _value)
}

// SetMaxBaseFeeByBytes is a paid mutator transaction binding the contract method 0x408d79cf.
//
// Solidity: function setMaxBaseFeeByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetMaxBaseFeeByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setMaxBaseFeeByBytes", _value)
}

// SetMaxBaseFeeByBytes is a paid mutator transaction binding the contract method 0x408d79cf.
//
// Solidity: function setMaxBaseFeeByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetMaxBaseFeeByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetMaxBaseFeeByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetMaxBaseFeeByBytes is a paid mutator transaction binding the contract method 0x408d79cf.
//
// Solidity: function setMaxBaseFeeByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetMaxBaseFeeByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetMaxBaseFeeByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetMaxIdleBlockInterval is a paid mutator transaction binding the contract method 0xf6fd7129.
//
// Solidity: function setMaxIdleBlockInterval(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetMaxIdleBlockInterval(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setMaxIdleBlockInterval", _value)
}

// SetMaxIdleBlockInterval is a paid mutator transaction binding the contract method 0xf6fd7129.
//
// Solidity: function setMaxIdleBlockInterval(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetMaxIdleBlockInterval(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetMaxIdleBlockInterval(&_EnvStorageImp.TransactOpts, _value)
}

// SetMaxIdleBlockInterval is a paid mutator transaction binding the contract method 0xf6fd7129.
//
// Solidity: function setMaxIdleBlockInterval(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetMaxIdleBlockInterval(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetMaxIdleBlockInterval(&_EnvStorageImp.TransactOpts, _value)
}

// SetMaxIdleBlockIntervalByBytes is a paid mutator transaction binding the contract method 0x0b90a39a.
//
// Solidity: function setMaxIdleBlockIntervalByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetMaxIdleBlockIntervalByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setMaxIdleBlockIntervalByBytes", _value)
}

// SetMaxIdleBlockIntervalByBytes is a paid mutator transaction binding the contract method 0x0b90a39a.
//
// Solidity: function setMaxIdleBlockIntervalByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetMaxIdleBlockIntervalByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetMaxIdleBlockIntervalByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetMaxIdleBlockIntervalByBytes is a paid mutator transaction binding the contract method 0x0b90a39a.
//
// Solidity: function setMaxIdleBlockIntervalByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetMaxIdleBlockIntervalByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetMaxIdleBlockIntervalByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetMaxPriorityFeePerGas is a paid mutator transaction binding the contract method 0xec3df879.
//
// Solidity: function setMaxPriorityFeePerGas(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetMaxPriorityFeePerGas(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setMaxPriorityFeePerGas", _value)
}

// SetMaxPriorityFeePerGas is a paid mutator transaction binding the contract method 0xec3df879.
//
// Solidity: function setMaxPriorityFeePerGas(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetMaxPriorityFeePerGas(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetMaxPriorityFeePerGas(&_EnvStorageImp.TransactOpts, _value)
}

// SetMaxPriorityFeePerGas is a paid mutator transaction binding the contract method 0xec3df879.
//
// Solidity: function setMaxPriorityFeePerGas(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetMaxPriorityFeePerGas(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetMaxPriorityFeePerGas(&_EnvStorageImp.TransactOpts, _value)
}

// SetMaxPriorityFeePerGasByBytes is a paid mutator transaction binding the contract method 0x0fc238bf.
//
// Solidity: function setMaxPriorityFeePerGasByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetMaxPriorityFeePerGasByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setMaxPriorityFeePerGasByBytes", _value)
}

// SetMaxPriorityFeePerGasByBytes is a paid mutator transaction binding the contract method 0x0fc238bf.
//
// Solidity: function setMaxPriorityFeePerGasByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetMaxPriorityFeePerGasByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetMaxPriorityFeePerGasByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetMaxPriorityFeePerGasByBytes is a paid mutator transaction binding the contract method 0x0fc238bf.
//
// Solidity: function setMaxPriorityFeePerGasByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetMaxPriorityFeePerGasByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetMaxPriorityFeePerGasByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetRegistry(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setRegistry", _addr)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetRegistry(_addr common.Address) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetRegistry(&_EnvStorageImp.TransactOpts, _addr)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetRegistry(_addr common.Address) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetRegistry(&_EnvStorageImp.TransactOpts, _addr)
}

// SetStakingMax is a paid mutator transaction binding the contract method 0xf8214591.
//
// Solidity: function setStakingMax(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetStakingMax(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setStakingMax", _value)
}

// SetStakingMax is a paid mutator transaction binding the contract method 0xf8214591.
//
// Solidity: function setStakingMax(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetStakingMax(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetStakingMax(&_EnvStorageImp.TransactOpts, _value)
}

// SetStakingMax is a paid mutator transaction binding the contract method 0xf8214591.
//
// Solidity: function setStakingMax(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetStakingMax(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetStakingMax(&_EnvStorageImp.TransactOpts, _value)
}

// SetStakingMaxByBytes is a paid mutator transaction binding the contract method 0x2eccd832.
//
// Solidity: function setStakingMaxByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetStakingMaxByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setStakingMaxByBytes", _value)
}

// SetStakingMaxByBytes is a paid mutator transaction binding the contract method 0x2eccd832.
//
// Solidity: function setStakingMaxByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetStakingMaxByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetStakingMaxByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetStakingMaxByBytes is a paid mutator transaction binding the contract method 0x2eccd832.
//
// Solidity: function setStakingMaxByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetStakingMaxByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetStakingMaxByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetStakingMin is a paid mutator transaction binding the contract method 0x2eb57c65.
//
// Solidity: function setStakingMin(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetStakingMin(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setStakingMin", _value)
}

// SetStakingMin is a paid mutator transaction binding the contract method 0x2eb57c65.
//
// Solidity: function setStakingMin(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetStakingMin(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetStakingMin(&_EnvStorageImp.TransactOpts, _value)
}

// SetStakingMin is a paid mutator transaction binding the contract method 0x2eb57c65.
//
// Solidity: function setStakingMin(uint256 _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetStakingMin(_value *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetStakingMin(&_EnvStorageImp.TransactOpts, _value)
}

// SetStakingMinByBytes is a paid mutator transaction binding the contract method 0x71c6960d.
//
// Solidity: function setStakingMinByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetStakingMinByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setStakingMinByBytes", _value)
}

// SetStakingMinByBytes is a paid mutator transaction binding the contract method 0x71c6960d.
//
// Solidity: function setStakingMinByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetStakingMinByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetStakingMinByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetStakingMinByBytes is a paid mutator transaction binding the contract method 0x71c6960d.
//
// Solidity: function setStakingMinByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetStakingMinByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetStakingMinByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetStakingMinMax is a paid mutator transaction binding the contract method 0x8d5cdd7e.
//
// Solidity: function setStakingMinMax(uint256 _min, uint256 _max) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetStakingMinMax(opts *bind.TransactOpts, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setStakingMinMax", _min, _max)
}

// SetStakingMinMax is a paid mutator transaction binding the contract method 0x8d5cdd7e.
//
// Solidity: function setStakingMinMax(uint256 _min, uint256 _max) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetStakingMinMax(_min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetStakingMinMax(&_EnvStorageImp.TransactOpts, _min, _max)
}

// SetStakingMinMax is a paid mutator transaction binding the contract method 0x8d5cdd7e.
//
// Solidity: function setStakingMinMax(uint256 _min, uint256 _max) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetStakingMinMax(_min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetStakingMinMax(&_EnvStorageImp.TransactOpts, _min, _max)
}

// SetStakingMinMaxByBytes is a paid mutator transaction binding the contract method 0xe078869f.
//
// Solidity: function setStakingMinMaxByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetStakingMinMaxByBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setStakingMinMaxByBytes", _value)
}

// SetStakingMinMaxByBytes is a paid mutator transaction binding the contract method 0xe078869f.
//
// Solidity: function setStakingMinMaxByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetStakingMinMaxByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetStakingMinMaxByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetStakingMinMaxByBytes is a paid mutator transaction binding the contract method 0xe078869f.
//
// Solidity: function setStakingMinMaxByBytes(bytes _value) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetStakingMinMaxByBytes(_value []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetStakingMinMaxByBytes(&_EnvStorageImp.TransactOpts, _value)
}

// SetVariable is a paid mutator transaction binding the contract method 0x88c28019.
//
// Solidity: function setVariable(bytes32 envKey, bytes envVal) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) SetVariable(opts *bind.TransactOpts, envKey [32]byte, envVal []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "setVariable", envKey, envVal)
}

// SetVariable is a paid mutator transaction binding the contract method 0x88c28019.
//
// Solidity: function setVariable(bytes32 envKey, bytes envVal) returns()
func (_EnvStorageImp *EnvStorageImpSession) SetVariable(envKey [32]byte, envVal []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetVariable(&_EnvStorageImp.TransactOpts, envKey, envVal)
}

// SetVariable is a paid mutator transaction binding the contract method 0x88c28019.
//
// Solidity: function setVariable(bytes32 envKey, bytes envVal) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) SetVariable(envKey [32]byte, envVal []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.SetVariable(&_EnvStorageImp.TransactOpts, envKey, envVal)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnvStorageImp *EnvStorageImpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.TransferOwnership(&_EnvStorageImp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.TransferOwnership(&_EnvStorageImp.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EnvStorageImp *EnvStorageImpTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EnvStorageImp *EnvStorageImpSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.UpgradeTo(&_EnvStorageImp.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.UpgradeTo(&_EnvStorageImp.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EnvStorageImp *EnvStorageImpTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EnvStorageImp.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EnvStorageImp *EnvStorageImpSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.UpgradeToAndCall(&_EnvStorageImp.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EnvStorageImp *EnvStorageImpTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EnvStorageImp.Contract.UpgradeToAndCall(&_EnvStorageImp.TransactOpts, newImplementation, data)
}

// EnvStorageImpAddressVarableChangedIterator is returned from FilterAddressVarableChanged and is used to iterate over the raw logs and unpacked data for AddressVarableChanged events raised by the EnvStorageImp contract.
type EnvStorageImpAddressVarableChangedIterator struct {
	Event *EnvStorageImpAddressVarableChanged // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpAddressVarableChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpAddressVarableChanged)
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
		it.Event = new(EnvStorageImpAddressVarableChanged)
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
func (it *EnvStorageImpAddressVarableChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpAddressVarableChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpAddressVarableChanged represents a AddressVarableChanged event raised by the EnvStorageImp contract.
type EnvStorageImpAddressVarableChanged struct {
	Name  [32]byte
	Value common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddressVarableChanged is a free log retrieval operation binding the contract event 0x5ba696c1849fe23c6b280f9f6057cb4574c818e0cd542032eed6e211f1e02464.
//
// Solidity: event AddressVarableChanged(bytes32 indexed _name, address _value)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterAddressVarableChanged(opts *bind.FilterOpts, _name [][32]byte) (*EnvStorageImpAddressVarableChangedIterator, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "AddressVarableChanged", _nameRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpAddressVarableChangedIterator{contract: _EnvStorageImp.contract, event: "AddressVarableChanged", logs: logs, sub: sub}, nil
}

// WatchAddressVarableChanged is a free log subscription operation binding the contract event 0x5ba696c1849fe23c6b280f9f6057cb4574c818e0cd542032eed6e211f1e02464.
//
// Solidity: event AddressVarableChanged(bytes32 indexed _name, address _value)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchAddressVarableChanged(opts *bind.WatchOpts, sink chan<- *EnvStorageImpAddressVarableChanged, _name [][32]byte) (event.Subscription, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "AddressVarableChanged", _nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpAddressVarableChanged)
				if err := _EnvStorageImp.contract.UnpackLog(event, "AddressVarableChanged", log); err != nil {
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

// ParseAddressVarableChanged is a log parse operation binding the contract event 0x5ba696c1849fe23c6b280f9f6057cb4574c818e0cd542032eed6e211f1e02464.
//
// Solidity: event AddressVarableChanged(bytes32 indexed _name, address _value)
func (_EnvStorageImp *EnvStorageImpFilterer) ParseAddressVarableChanged(log types.Log) (*EnvStorageImpAddressVarableChanged, error) {
	event := new(EnvStorageImpAddressVarableChanged)
	if err := _EnvStorageImp.contract.UnpackLog(event, "AddressVarableChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the EnvStorageImp contract.
type EnvStorageImpAdminChangedIterator struct {
	Event *EnvStorageImpAdminChanged // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpAdminChanged)
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
		it.Event = new(EnvStorageImpAdminChanged)
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
func (it *EnvStorageImpAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpAdminChanged represents a AdminChanged event raised by the EnvStorageImp contract.
type EnvStorageImpAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*EnvStorageImpAdminChangedIterator, error) {

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpAdminChangedIterator{contract: _EnvStorageImp.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *EnvStorageImpAdminChanged) (event.Subscription, error) {

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpAdminChanged)
				if err := _EnvStorageImp.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EnvStorageImp *EnvStorageImpFilterer) ParseAdminChanged(log types.Log) (*EnvStorageImpAdminChanged, error) {
	event := new(EnvStorageImpAdminChanged)
	if err := _EnvStorageImp.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the EnvStorageImp contract.
type EnvStorageImpBeaconUpgradedIterator struct {
	Event *EnvStorageImpBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpBeaconUpgraded)
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
		it.Event = new(EnvStorageImpBeaconUpgraded)
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
func (it *EnvStorageImpBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpBeaconUpgraded represents a BeaconUpgraded event raised by the EnvStorageImp contract.
type EnvStorageImpBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*EnvStorageImpBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpBeaconUpgradedIterator{contract: _EnvStorageImp.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *EnvStorageImpBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpBeaconUpgraded)
				if err := _EnvStorageImp.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EnvStorageImp *EnvStorageImpFilterer) ParseBeaconUpgraded(log types.Log) (*EnvStorageImpBeaconUpgraded, error) {
	event := new(EnvStorageImpBeaconUpgraded)
	if err := _EnvStorageImp.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpBytes32VarableChangedIterator is returned from FilterBytes32VarableChanged and is used to iterate over the raw logs and unpacked data for Bytes32VarableChanged events raised by the EnvStorageImp contract.
type EnvStorageImpBytes32VarableChangedIterator struct {
	Event *EnvStorageImpBytes32VarableChanged // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpBytes32VarableChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpBytes32VarableChanged)
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
		it.Event = new(EnvStorageImpBytes32VarableChanged)
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
func (it *EnvStorageImpBytes32VarableChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpBytes32VarableChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpBytes32VarableChanged represents a Bytes32VarableChanged event raised by the EnvStorageImp contract.
type EnvStorageImpBytes32VarableChanged struct {
	Name  [32]byte
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBytes32VarableChanged is a free log retrieval operation binding the contract event 0xde32aec28e018560c62b0217a8e7cfdc094cda68f1fc7ecd4ed833ee1d4862c9.
//
// Solidity: event Bytes32VarableChanged(bytes32 indexed _name, bytes32 _value)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterBytes32VarableChanged(opts *bind.FilterOpts, _name [][32]byte) (*EnvStorageImpBytes32VarableChangedIterator, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "Bytes32VarableChanged", _nameRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpBytes32VarableChangedIterator{contract: _EnvStorageImp.contract, event: "Bytes32VarableChanged", logs: logs, sub: sub}, nil
}

// WatchBytes32VarableChanged is a free log subscription operation binding the contract event 0xde32aec28e018560c62b0217a8e7cfdc094cda68f1fc7ecd4ed833ee1d4862c9.
//
// Solidity: event Bytes32VarableChanged(bytes32 indexed _name, bytes32 _value)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchBytes32VarableChanged(opts *bind.WatchOpts, sink chan<- *EnvStorageImpBytes32VarableChanged, _name [][32]byte) (event.Subscription, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "Bytes32VarableChanged", _nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpBytes32VarableChanged)
				if err := _EnvStorageImp.contract.UnpackLog(event, "Bytes32VarableChanged", log); err != nil {
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

// ParseBytes32VarableChanged is a log parse operation binding the contract event 0xde32aec28e018560c62b0217a8e7cfdc094cda68f1fc7ecd4ed833ee1d4862c9.
//
// Solidity: event Bytes32VarableChanged(bytes32 indexed _name, bytes32 _value)
func (_EnvStorageImp *EnvStorageImpFilterer) ParseBytes32VarableChanged(log types.Log) (*EnvStorageImpBytes32VarableChanged, error) {
	event := new(EnvStorageImpBytes32VarableChanged)
	if err := _EnvStorageImp.contract.UnpackLog(event, "Bytes32VarableChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpBytesVarableChangedIterator is returned from FilterBytesVarableChanged and is used to iterate over the raw logs and unpacked data for BytesVarableChanged events raised by the EnvStorageImp contract.
type EnvStorageImpBytesVarableChangedIterator struct {
	Event *EnvStorageImpBytesVarableChanged // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpBytesVarableChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpBytesVarableChanged)
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
		it.Event = new(EnvStorageImpBytesVarableChanged)
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
func (it *EnvStorageImpBytesVarableChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpBytesVarableChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpBytesVarableChanged represents a BytesVarableChanged event raised by the EnvStorageImp contract.
type EnvStorageImpBytesVarableChanged struct {
	Name  [32]byte
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBytesVarableChanged is a free log retrieval operation binding the contract event 0x871b63105583c40db6d530b9340369142c5ae265e41a501d300455015dc1283a.
//
// Solidity: event BytesVarableChanged(bytes32 indexed _name, bytes _value)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterBytesVarableChanged(opts *bind.FilterOpts, _name [][32]byte) (*EnvStorageImpBytesVarableChangedIterator, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "BytesVarableChanged", _nameRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpBytesVarableChangedIterator{contract: _EnvStorageImp.contract, event: "BytesVarableChanged", logs: logs, sub: sub}, nil
}

// WatchBytesVarableChanged is a free log subscription operation binding the contract event 0x871b63105583c40db6d530b9340369142c5ae265e41a501d300455015dc1283a.
//
// Solidity: event BytesVarableChanged(bytes32 indexed _name, bytes _value)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchBytesVarableChanged(opts *bind.WatchOpts, sink chan<- *EnvStorageImpBytesVarableChanged, _name [][32]byte) (event.Subscription, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "BytesVarableChanged", _nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpBytesVarableChanged)
				if err := _EnvStorageImp.contract.UnpackLog(event, "BytesVarableChanged", log); err != nil {
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

// ParseBytesVarableChanged is a log parse operation binding the contract event 0x871b63105583c40db6d530b9340369142c5ae265e41a501d300455015dc1283a.
//
// Solidity: event BytesVarableChanged(bytes32 indexed _name, bytes _value)
func (_EnvStorageImp *EnvStorageImpFilterer) ParseBytesVarableChanged(log types.Log) (*EnvStorageImpBytesVarableChanged, error) {
	event := new(EnvStorageImpBytesVarableChanged)
	if err := _EnvStorageImp.contract.UnpackLog(event, "BytesVarableChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EnvStorageImp contract.
type EnvStorageImpInitializedIterator struct {
	Event *EnvStorageImpInitialized // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpInitialized)
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
		it.Event = new(EnvStorageImpInitialized)
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
func (it *EnvStorageImpInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpInitialized represents a Initialized event raised by the EnvStorageImp contract.
type EnvStorageImpInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterInitialized(opts *bind.FilterOpts) (*EnvStorageImpInitializedIterator, error) {

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpInitializedIterator{contract: _EnvStorageImp.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EnvStorageImpInitialized) (event.Subscription, error) {

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpInitialized)
				if err := _EnvStorageImp.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EnvStorageImp *EnvStorageImpFilterer) ParseInitialized(log types.Log) (*EnvStorageImpInitialized, error) {
	event := new(EnvStorageImpInitialized)
	if err := _EnvStorageImp.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpIntVarableChangedIterator is returned from FilterIntVarableChanged and is used to iterate over the raw logs and unpacked data for IntVarableChanged events raised by the EnvStorageImp contract.
type EnvStorageImpIntVarableChangedIterator struct {
	Event *EnvStorageImpIntVarableChanged // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpIntVarableChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpIntVarableChanged)
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
		it.Event = new(EnvStorageImpIntVarableChanged)
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
func (it *EnvStorageImpIntVarableChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpIntVarableChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpIntVarableChanged represents a IntVarableChanged event raised by the EnvStorageImp contract.
type EnvStorageImpIntVarableChanged struct {
	Name  [32]byte
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterIntVarableChanged is a free log retrieval operation binding the contract event 0x3b274b11383ad4ca069e655ed60df9e4307b7a0f82ba2245a3cb9789ab814bf6.
//
// Solidity: event IntVarableChanged(bytes32 indexed _name, int256 _value)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterIntVarableChanged(opts *bind.FilterOpts, _name [][32]byte) (*EnvStorageImpIntVarableChangedIterator, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "IntVarableChanged", _nameRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpIntVarableChangedIterator{contract: _EnvStorageImp.contract, event: "IntVarableChanged", logs: logs, sub: sub}, nil
}

// WatchIntVarableChanged is a free log subscription operation binding the contract event 0x3b274b11383ad4ca069e655ed60df9e4307b7a0f82ba2245a3cb9789ab814bf6.
//
// Solidity: event IntVarableChanged(bytes32 indexed _name, int256 _value)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchIntVarableChanged(opts *bind.WatchOpts, sink chan<- *EnvStorageImpIntVarableChanged, _name [][32]byte) (event.Subscription, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "IntVarableChanged", _nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpIntVarableChanged)
				if err := _EnvStorageImp.contract.UnpackLog(event, "IntVarableChanged", log); err != nil {
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

// ParseIntVarableChanged is a log parse operation binding the contract event 0x3b274b11383ad4ca069e655ed60df9e4307b7a0f82ba2245a3cb9789ab814bf6.
//
// Solidity: event IntVarableChanged(bytes32 indexed _name, int256 _value)
func (_EnvStorageImp *EnvStorageImpFilterer) ParseIntVarableChanged(log types.Log) (*EnvStorageImpIntVarableChanged, error) {
	event := new(EnvStorageImpIntVarableChanged)
	if err := _EnvStorageImp.contract.UnpackLog(event, "IntVarableChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EnvStorageImp contract.
type EnvStorageImpOwnershipTransferredIterator struct {
	Event *EnvStorageImpOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpOwnershipTransferred)
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
		it.Event = new(EnvStorageImpOwnershipTransferred)
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
func (it *EnvStorageImpOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpOwnershipTransferred represents a OwnershipTransferred event raised by the EnvStorageImp contract.
type EnvStorageImpOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EnvStorageImpOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpOwnershipTransferredIterator{contract: _EnvStorageImp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EnvStorageImpOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpOwnershipTransferred)
				if err := _EnvStorageImp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EnvStorageImp *EnvStorageImpFilterer) ParseOwnershipTransferred(log types.Log) (*EnvStorageImpOwnershipTransferred, error) {
	event := new(EnvStorageImpOwnershipTransferred)
	if err := _EnvStorageImp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpSetRegistryIterator is returned from FilterSetRegistry and is used to iterate over the raw logs and unpacked data for SetRegistry events raised by the EnvStorageImp contract.
type EnvStorageImpSetRegistryIterator struct {
	Event *EnvStorageImpSetRegistry // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpSetRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpSetRegistry)
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
		it.Event = new(EnvStorageImpSetRegistry)
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
func (it *EnvStorageImpSetRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpSetRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpSetRegistry represents a SetRegistry event raised by the EnvStorageImp contract.
type EnvStorageImpSetRegistry struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetRegistry is a free log retrieval operation binding the contract event 0x278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd.
//
// Solidity: event SetRegistry(address indexed addr)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterSetRegistry(opts *bind.FilterOpts, addr []common.Address) (*EnvStorageImpSetRegistryIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "SetRegistry", addrRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpSetRegistryIterator{contract: _EnvStorageImp.contract, event: "SetRegistry", logs: logs, sub: sub}, nil
}

// WatchSetRegistry is a free log subscription operation binding the contract event 0x278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd.
//
// Solidity: event SetRegistry(address indexed addr)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchSetRegistry(opts *bind.WatchOpts, sink chan<- *EnvStorageImpSetRegistry, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "SetRegistry", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpSetRegistry)
				if err := _EnvStorageImp.contract.UnpackLog(event, "SetRegistry", log); err != nil {
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

// ParseSetRegistry is a log parse operation binding the contract event 0x278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd.
//
// Solidity: event SetRegistry(address indexed addr)
func (_EnvStorageImp *EnvStorageImpFilterer) ParseSetRegistry(log types.Log) (*EnvStorageImpSetRegistry, error) {
	event := new(EnvStorageImpSetRegistry)
	if err := _EnvStorageImp.contract.UnpackLog(event, "SetRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpStringVarableChangedIterator is returned from FilterStringVarableChanged and is used to iterate over the raw logs and unpacked data for StringVarableChanged events raised by the EnvStorageImp contract.
type EnvStorageImpStringVarableChangedIterator struct {
	Event *EnvStorageImpStringVarableChanged // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpStringVarableChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpStringVarableChanged)
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
		it.Event = new(EnvStorageImpStringVarableChanged)
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
func (it *EnvStorageImpStringVarableChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpStringVarableChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpStringVarableChanged represents a StringVarableChanged event raised by the EnvStorageImp contract.
type EnvStorageImpStringVarableChanged struct {
	Name  [32]byte
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStringVarableChanged is a free log retrieval operation binding the contract event 0x43849fbc1e66a5d2cd0ace04916e2e6c8f314c1904d104cad97150428a370a41.
//
// Solidity: event StringVarableChanged(bytes32 indexed _name, string _value)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterStringVarableChanged(opts *bind.FilterOpts, _name [][32]byte) (*EnvStorageImpStringVarableChangedIterator, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "StringVarableChanged", _nameRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpStringVarableChangedIterator{contract: _EnvStorageImp.contract, event: "StringVarableChanged", logs: logs, sub: sub}, nil
}

// WatchStringVarableChanged is a free log subscription operation binding the contract event 0x43849fbc1e66a5d2cd0ace04916e2e6c8f314c1904d104cad97150428a370a41.
//
// Solidity: event StringVarableChanged(bytes32 indexed _name, string _value)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchStringVarableChanged(opts *bind.WatchOpts, sink chan<- *EnvStorageImpStringVarableChanged, _name [][32]byte) (event.Subscription, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "StringVarableChanged", _nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpStringVarableChanged)
				if err := _EnvStorageImp.contract.UnpackLog(event, "StringVarableChanged", log); err != nil {
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

// ParseStringVarableChanged is a log parse operation binding the contract event 0x43849fbc1e66a5d2cd0ace04916e2e6c8f314c1904d104cad97150428a370a41.
//
// Solidity: event StringVarableChanged(bytes32 indexed _name, string _value)
func (_EnvStorageImp *EnvStorageImpFilterer) ParseStringVarableChanged(log types.Log) (*EnvStorageImpStringVarableChanged, error) {
	event := new(EnvStorageImpStringVarableChanged)
	if err := _EnvStorageImp.contract.UnpackLog(event, "StringVarableChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpUintVarableChangedIterator is returned from FilterUintVarableChanged and is used to iterate over the raw logs and unpacked data for UintVarableChanged events raised by the EnvStorageImp contract.
type EnvStorageImpUintVarableChangedIterator struct {
	Event *EnvStorageImpUintVarableChanged // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpUintVarableChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpUintVarableChanged)
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
		it.Event = new(EnvStorageImpUintVarableChanged)
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
func (it *EnvStorageImpUintVarableChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpUintVarableChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpUintVarableChanged represents a UintVarableChanged event raised by the EnvStorageImp contract.
type EnvStorageImpUintVarableChanged struct {
	Name  [32]byte
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUintVarableChanged is a free log retrieval operation binding the contract event 0x58d7c10adfd5778016889c15d422d57f2975b9292415b54fe5d8cd4241200612.
//
// Solidity: event UintVarableChanged(bytes32 indexed _name, uint256 _value)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterUintVarableChanged(opts *bind.FilterOpts, _name [][32]byte) (*EnvStorageImpUintVarableChangedIterator, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "UintVarableChanged", _nameRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpUintVarableChangedIterator{contract: _EnvStorageImp.contract, event: "UintVarableChanged", logs: logs, sub: sub}, nil
}

// WatchUintVarableChanged is a free log subscription operation binding the contract event 0x58d7c10adfd5778016889c15d422d57f2975b9292415b54fe5d8cd4241200612.
//
// Solidity: event UintVarableChanged(bytes32 indexed _name, uint256 _value)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchUintVarableChanged(opts *bind.WatchOpts, sink chan<- *EnvStorageImpUintVarableChanged, _name [][32]byte) (event.Subscription, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "UintVarableChanged", _nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpUintVarableChanged)
				if err := _EnvStorageImp.contract.UnpackLog(event, "UintVarableChanged", log); err != nil {
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

// ParseUintVarableChanged is a log parse operation binding the contract event 0x58d7c10adfd5778016889c15d422d57f2975b9292415b54fe5d8cd4241200612.
//
// Solidity: event UintVarableChanged(bytes32 indexed _name, uint256 _value)
func (_EnvStorageImp *EnvStorageImpFilterer) ParseUintVarableChanged(log types.Log) (*EnvStorageImpUintVarableChanged, error) {
	event := new(EnvStorageImpUintVarableChanged)
	if err := _EnvStorageImp.contract.UnpackLog(event, "UintVarableChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpUpgradeImplementationIterator is returned from FilterUpgradeImplementation and is used to iterate over the raw logs and unpacked data for UpgradeImplementation events raised by the EnvStorageImp contract.
type EnvStorageImpUpgradeImplementationIterator struct {
	Event *EnvStorageImpUpgradeImplementation // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpUpgradeImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpUpgradeImplementation)
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
		it.Event = new(EnvStorageImpUpgradeImplementation)
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
func (it *EnvStorageImpUpgradeImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpUpgradeImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpUpgradeImplementation represents a UpgradeImplementation event raised by the EnvStorageImp contract.
type EnvStorageImpUpgradeImplementation struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgradeImplementation is a free log retrieval operation binding the contract event 0x2a74bbaf99cf25202925743012bb137262ab3432c821e1ba94ddbecac2ea9708.
//
// Solidity: event UpgradeImplementation(address indexed implementation)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterUpgradeImplementation(opts *bind.FilterOpts, implementation []common.Address) (*EnvStorageImpUpgradeImplementationIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "UpgradeImplementation", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpUpgradeImplementationIterator{contract: _EnvStorageImp.contract, event: "UpgradeImplementation", logs: logs, sub: sub}, nil
}

// WatchUpgradeImplementation is a free log subscription operation binding the contract event 0x2a74bbaf99cf25202925743012bb137262ab3432c821e1ba94ddbecac2ea9708.
//
// Solidity: event UpgradeImplementation(address indexed implementation)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchUpgradeImplementation(opts *bind.WatchOpts, sink chan<- *EnvStorageImpUpgradeImplementation, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "UpgradeImplementation", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpUpgradeImplementation)
				if err := _EnvStorageImp.contract.UnpackLog(event, "UpgradeImplementation", log); err != nil {
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

// ParseUpgradeImplementation is a log parse operation binding the contract event 0x2a74bbaf99cf25202925743012bb137262ab3432c821e1ba94ddbecac2ea9708.
//
// Solidity: event UpgradeImplementation(address indexed implementation)
func (_EnvStorageImp *EnvStorageImpFilterer) ParseUpgradeImplementation(log types.Log) (*EnvStorageImpUpgradeImplementation, error) {
	event := new(EnvStorageImpUpgradeImplementation)
	if err := _EnvStorageImp.contract.UnpackLog(event, "UpgradeImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EnvStorageImp contract.
type EnvStorageImpUpgradedIterator struct {
	Event *EnvStorageImpUpgraded // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpUpgraded)
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
		it.Event = new(EnvStorageImpUpgraded)
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
func (it *EnvStorageImpUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpUpgraded represents a Upgraded event raised by the EnvStorageImp contract.
type EnvStorageImpUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EnvStorageImpUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpUpgradedIterator{contract: _EnvStorageImp.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EnvStorageImpUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpUpgraded)
				if err := _EnvStorageImp.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EnvStorageImp *EnvStorageImpFilterer) ParseUpgraded(log types.Log) (*EnvStorageImpUpgraded, error) {
	event := new(EnvStorageImpUpgraded)
	if err := _EnvStorageImp.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnvStorageImpVarableChangedIterator is returned from FilterVarableChanged and is used to iterate over the raw logs and unpacked data for VarableChanged events raised by the EnvStorageImp contract.
type EnvStorageImpVarableChangedIterator struct {
	Event *EnvStorageImpVarableChanged // Event containing the contract specifics and raw log

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
func (it *EnvStorageImpVarableChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnvStorageImpVarableChanged)
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
		it.Event = new(EnvStorageImpVarableChanged)
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
func (it *EnvStorageImpVarableChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnvStorageImpVarableChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnvStorageImpVarableChanged represents a VarableChanged event raised by the EnvStorageImp contract.
type EnvStorageImpVarableChanged struct {
	Name  [32]byte
	Type  *big.Int
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVarableChanged is a free log retrieval operation binding the contract event 0x60841501cd68d1c6414ed8f47a579de24c7e2f9da7e97f10c4932196d92df402.
//
// Solidity: event VarableChanged(bytes32 indexed _name, uint256 indexed _type, string _value)
func (_EnvStorageImp *EnvStorageImpFilterer) FilterVarableChanged(opts *bind.FilterOpts, _name [][32]byte, _type []*big.Int) (*EnvStorageImpVarableChangedIterator, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}
	var _typeRule []interface{}
	for _, _typeItem := range _type {
		_typeRule = append(_typeRule, _typeItem)
	}

	logs, sub, err := _EnvStorageImp.contract.FilterLogs(opts, "VarableChanged", _nameRule, _typeRule)
	if err != nil {
		return nil, err
	}
	return &EnvStorageImpVarableChangedIterator{contract: _EnvStorageImp.contract, event: "VarableChanged", logs: logs, sub: sub}, nil
}

// WatchVarableChanged is a free log subscription operation binding the contract event 0x60841501cd68d1c6414ed8f47a579de24c7e2f9da7e97f10c4932196d92df402.
//
// Solidity: event VarableChanged(bytes32 indexed _name, uint256 indexed _type, string _value)
func (_EnvStorageImp *EnvStorageImpFilterer) WatchVarableChanged(opts *bind.WatchOpts, sink chan<- *EnvStorageImpVarableChanged, _name [][32]byte, _type []*big.Int) (event.Subscription, error) {

	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}
	var _typeRule []interface{}
	for _, _typeItem := range _type {
		_typeRule = append(_typeRule, _typeItem)
	}

	logs, sub, err := _EnvStorageImp.contract.WatchLogs(opts, "VarableChanged", _nameRule, _typeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnvStorageImpVarableChanged)
				if err := _EnvStorageImp.contract.UnpackLog(event, "VarableChanged", log); err != nil {
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

// ParseVarableChanged is a log parse operation binding the contract event 0x60841501cd68d1c6414ed8f47a579de24c7e2f9da7e97f10c4932196d92df402.
//
// Solidity: event VarableChanged(bytes32 indexed _name, uint256 indexed _type, string _value)
func (_EnvStorageImp *EnvStorageImpFilterer) ParseVarableChanged(log types.Log) (*EnvStorageImpVarableChanged, error) {
	event := new(EnvStorageImpVarableChanged)
	if err := _EnvStorageImp.contract.UnpackLog(event, "VarableChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
