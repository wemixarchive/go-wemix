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

// GovImpMemberInfo is an auto generated low-level Go binding around an user-defined struct.
type GovImpMemberInfo struct {
	Staker     common.Address
	Voter      common.Address
	Reward     common.Address
	Name       []byte
	Enode      []byte
	Ip         []byte
	Port       *big.Int
	LockAmount *big.Int
	Memo       []byte
	Duration   *big.Int
}

// GovMetaData contains all meta data concerning the Gov contract.
var GovMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_imp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"5c60da1b": "implementation()",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161077238038061077283398101604081905261002f91610326565b604080516020810190915260008152819061006b60017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd61034f565b60008051602061072b8339815191521461008757610087610374565b6100938282600061009b565b505050610405565b6100a4836100d1565b6000825111806100b15750805b156100cc576100ca838361011160201b61008b1760201c565b505b505050565b6100da8161013d565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060610136838360405180606001604052806027815260200161074b602791396101fd565b9392505050565b610150816102db60201b6100b71760201c565b6101b75760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b806101dc60008051602061072b83398151915260001b6102ea60201b6100c61760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606001600160a01b0384163b6102655760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016101ae565b600080856001600160a01b03168560405161028091906103b6565b600060405180830381855af49150503d80600081146102bb576040519150601f19603f3d011682016040523d82523d6000602084013e6102c0565b606091505b5090925090506102d18282866102ed565b9695505050505050565b6001600160a01b03163b151590565b90565b606083156102fc575081610136565b82511561030c5782518084602001fd5b8160405162461bcd60e51b81526004016101ae91906103d2565b60006020828403121561033857600080fd5b81516001600160a01b038116811461013657600080fd5b60008282101561036f57634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052600160045260246000fd5b60005b838110156103a557818101518382015260200161038d565b838111156100ca5750506000910152565b600082516103c881846020870161038a565b9190910192915050565b60208152600082518060208401526103f181604085016020870161038a565b601f01601f19169190910160400192915050565b610317806104146000396000f3fe6080604052600436106100225760003560e01c80635c60da1b1461003957610031565b366100315761002f61006a565b005b61002f61006a565b34801561004557600080fd5b5061004e61007c565b6040516001600160a01b03909116815260200160405180910390f35b61007a6100756100c9565b6100fc565b565b60006100866100c9565b905090565b60606100b083836040518060600160405280602781526020016102bb60279139610120565b9392505050565b6001600160a01b03163b151590565b90565b60006100867f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b3660008037600080366000845af43d6000803e80801561011b573d6000f35b3d6000fd5b60606001600160a01b0384163b61018d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084015b60405180910390fd5b600080856001600160a01b0316856040516101a8919061026b565b600060405180830381855af49150503d80600081146101e3576040519150601f19603f3d011682016040523d82523d6000602084013e6101e8565b606091505b50915091506101f8828286610202565b9695505050505050565b606083156102115750816100b0565b8251156102215782518084602001fd5b8160405162461bcd60e51b81526004016101849190610287565b60005b8381101561025657818101518382015260200161023e565b83811115610265576000848401525b50505050565b6000825161027d81846020870161023b565b9190910192915050565b60208152600082518060208401526102a681604085016020870161023b565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212208728a0da67e5fc45b834821dc2e96583496c6ee88c819583331569902d33549964736f6c634300080e0033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
}

// GovABI is the input ABI used to generate the binding from.
// Deprecated: Use GovMetaData.ABI instead.
var GovABI = GovMetaData.ABI

// Deprecated: Use GovMetaData.Sigs instead.
// GovFuncSigs maps the 4-byte function signature to its string representation.
var GovFuncSigs = GovMetaData.Sigs

// GovBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovMetaData.Bin instead.
var GovBin = GovMetaData.Bin

// DeployGov deploys a new Ethereum contract, binding an instance of Gov to it.
func DeployGov(auth *bind.TransactOpts, backend bind.ContractBackend, _imp common.Address) (common.Address, *types.Transaction, *Gov, error) {
	parsed, err := GovMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovBin), backend, _imp)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gov{GovCaller: GovCaller{contract: contract}, GovTransactor: GovTransactor{contract: contract}, GovFilterer: GovFilterer{contract: contract}}, nil
}

// Gov is an auto generated Go binding around an Ethereum contract.
type Gov struct {
	GovCaller     // Read-only binding to the contract
	GovTransactor // Write-only binding to the contract
	GovFilterer   // Log filterer for contract events
}

// GovCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovSession struct {
	Contract     *Gov              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovCallerSession struct {
	Contract *GovCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GovTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovTransactorSession struct {
	Contract     *GovTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovRaw struct {
	Contract *Gov // Generic contract binding to access the raw methods on
}

// GovCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovCallerRaw struct {
	Contract *GovCaller // Generic read-only contract binding to access the raw methods on
}

// GovTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovTransactorRaw struct {
	Contract *GovTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGov creates a new instance of Gov, bound to a specific deployed contract.
func NewGov(address common.Address, backend bind.ContractBackend) (*Gov, error) {
	contract, err := bindGov(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gov{GovCaller: GovCaller{contract: contract}, GovTransactor: GovTransactor{contract: contract}, GovFilterer: GovFilterer{contract: contract}}, nil
}

// NewGovCaller creates a new read-only instance of Gov, bound to a specific deployed contract.
func NewGovCaller(address common.Address, caller bind.ContractCaller) (*GovCaller, error) {
	contract, err := bindGov(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovCaller{contract: contract}, nil
}

// NewGovTransactor creates a new write-only instance of Gov, bound to a specific deployed contract.
func NewGovTransactor(address common.Address, transactor bind.ContractTransactor) (*GovTransactor, error) {
	contract, err := bindGov(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovTransactor{contract: contract}, nil
}

// NewGovFilterer creates a new log filterer instance of Gov, bound to a specific deployed contract.
func NewGovFilterer(address common.Address, filterer bind.ContractFilterer) (*GovFilterer, error) {
	contract, err := bindGov(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovFilterer{contract: contract}, nil
}

// bindGov binds a generic wrapper to an already deployed contract.
func bindGov(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gov *GovRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gov.Contract.GovCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gov *GovRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gov.Contract.GovTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gov *GovRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gov.Contract.GovTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gov *GovCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gov.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gov *GovTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gov.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gov *GovTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gov.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Gov *GovCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Gov *GovSession) Implementation() (common.Address, error) {
	return _Gov.Contract.Implementation(&_Gov.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Gov *GovCallerSession) Implementation() (common.Address, error) {
	return _Gov.Contract.Implementation(&_Gov.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Gov *GovTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Gov.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Gov *GovSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Gov.Contract.Fallback(&_Gov.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Gov *GovTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Gov.Contract.Fallback(&_Gov.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gov *GovTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gov.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gov *GovSession) Receive() (*types.Transaction, error) {
	return _Gov.Contract.Receive(&_Gov.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gov *GovTransactorSession) Receive() (*types.Transaction, error) {
	return _Gov.Contract.Receive(&_Gov.TransactOpts)
}

// GovAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Gov contract.
type GovAdminChangedIterator struct {
	Event *GovAdminChanged // Event containing the contract specifics and raw log

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
func (it *GovAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovAdminChanged)
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
		it.Event = new(GovAdminChanged)
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
func (it *GovAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovAdminChanged represents a AdminChanged event raised by the Gov contract.
type GovAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Gov *GovFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*GovAdminChangedIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &GovAdminChangedIterator{contract: _Gov.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Gov *GovFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *GovAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovAdminChanged)
				if err := _Gov.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Gov *GovFilterer) ParseAdminChanged(log types.Log) (*GovAdminChanged, error) {
	event := new(GovAdminChanged)
	if err := _Gov.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Gov contract.
type GovBeaconUpgradedIterator struct {
	Event *GovBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *GovBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovBeaconUpgraded)
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
		it.Event = new(GovBeaconUpgraded)
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
func (it *GovBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovBeaconUpgraded represents a BeaconUpgraded event raised by the Gov contract.
type GovBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Gov *GovFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*GovBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Gov.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &GovBeaconUpgradedIterator{contract: _Gov.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Gov *GovFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *GovBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Gov.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovBeaconUpgraded)
				if err := _Gov.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Gov *GovFilterer) ParseBeaconUpgraded(log types.Log) (*GovBeaconUpgraded, error) {
	event := new(GovBeaconUpgraded)
	if err := _Gov.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Gov contract.
type GovUpgradedIterator struct {
	Event *GovUpgraded // Event containing the contract specifics and raw log

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
func (it *GovUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovUpgraded)
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
		it.Event = new(GovUpgraded)
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
func (it *GovUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovUpgraded represents a Upgraded event raised by the Gov contract.
type GovUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Gov *GovFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GovUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Gov.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GovUpgradedIterator{contract: _Gov.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Gov *GovFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GovUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Gov.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovUpgraded)
				if err := _Gov.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Gov *GovFilterer) ParseUpgraded(log types.Log) (*GovUpgraded, error) {
	event := new(GovUpgraded)
	if err := _Gov.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpMetaData contains all meta data concerning the GovImp contract.
var GovImpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"envName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"envType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"envVal\",\"type\":\"bytes\"}],\"name\":\"EnvChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"GovDataMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"MemberAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVoter\",\"type\":\"address\"}],\"name\":\"MemberChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"MemberRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"MemberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ballotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"NotApplicable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"SetProposalTimePeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_DURATION_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BALLOT_DURATION_MIN_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BALLOT_DURATION_MIN_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BALLOT_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_FEE_MAX_CHANGE_RATE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCKS_PER_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_CREATION_TIME_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_GASLIMIT_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_AMOUNT_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_METHOD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ECOSYSTEM_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ENV_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GASLIMIT_AND_BASE_FEE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAS_TARGET_PERCENTAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAINTENANCE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BASE_FEE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_IDLE_BLOCK_INTERVAL_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PRIORITY_FEE_PER_GAS_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_POOL_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MIN_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MIN_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_REWARD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"port\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"internalType\":\"structGovImp.MemberInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"addProposalToAddMember\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"envName\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"envType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"envVal\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"addProposalToChangeEnv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"addProposalToChangeGov\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"port\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"internalType\":\"structGovImp.MemberInfo\",\"name\":\"newInfo\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"oldStaker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashing\",\"type\":\"uint256\"}],\"name\":\"addProposalToChangeMember\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashing\",\"type\":\"uint256\"}],\"name\":\"addProposalToRemoveMember\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballotLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkUnfinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeEndedVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBallotInVoting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxStaking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxVotingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getMemberFromNodeIdx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMemberLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStaking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinVotingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"port\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getNodeIdxFromMember\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getStakerAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getVoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"port\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"oldModifiedBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"}],\"name\":\"initMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initOnce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isStaker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastAddProposalTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldGov\",\"type\":\"address\"}],\"name\":\"migrateFromLegacy\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modifiedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposal_time_period\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reg\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"setProposalTimePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approval\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voterIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
		"36e83d83": "addProposalToAddMember((address,address,address,bytes,bytes,bytes,uint256,uint256,bytes,uint256))",
		"40690353": "addProposalToChangeEnv(bytes32,uint256,bytes,bytes,uint256)",
		"0efa4909": "addProposalToChangeGov(address,bytes,uint256)",
		"a78a8188": "addProposalToChangeMember((address,address,address,bytes,bytes,bytes,uint256,uint256,bytes,uint256),address,uint256,uint256)",
		"894f5111": "addProposalToRemoveMember(address,uint256,bytes,uint256,uint256,uint256)",
		"d6f9cfce": "ballotLength()",
		"1c12b030": "checkUnfinalized()",
		"6ba99181": "finalizeEndedVote()",
		"de09b377": "getBallotInVoting()",
		"5aaa4040": "getMaxStaking()",
		"ce04b9d4": "getMaxVotingDuration()",
		"ab3545e5": "getMember(uint256)",
		"15bf6b4d": "getMemberFromNodeIdx(uint256)",
		"d965ea00": "getMemberLength()",
		"af6af2ff": "getMinStaking()",
		"1c150171": "getMinVotingDuration()",
		"4f0f4aa9": "getNode(uint256)",
		"ce6a54ff": "getNodeIdxFromMember(address)",
		"72016f75": "getNodeLength()",
		"1c4b774b": "getReward(uint256)",
		"6f6de96d": "getStakerAddr(address)",
		"e75235b8": "getThreshold()",
		"d07bff0c": "getVoter(uint256)",
		"a8915a3e": "init(address,uint256,bytes,bytes,bytes,uint256)",
		"397e38e7": "initMigration(address,uint256,address)",
		"351bacda": "initOnce(address,uint256,bytes)",
		"a230c524": "isMember(address)",
		"4d5ce038": "isReward(address)",
		"6f1e8533": "isStaker(address)",
		"a7771ee3": "isVoter(address)",
		"139d9dd3": "lastAddProposalTime(address)",
		"8d39e33a": "migrateFromLegacy(address)",
		"7d10dd1b": "modifiedBlock()",
		"8da5cb5b": "owner()",
		"3310569c": "proposal_time_period()",
		"52d1902d": "proxiableUUID()",
		"16fbe831": "reInit()",
		"738fdd1a": "reg()",
		"715018a6": "renounceOwnership()",
		"aaf0dd36": "rewardIdx(address)",
		"e27bdaef": "setProposalTimePeriod(uint256)",
		"a91ee0dc": "setRegistry(address)",
		"a0c12683": "stakerIdx(address)",
		"f2fde38b": "transferOwnership(address)",
		"3659cfe6": "upgradeTo(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
		"c9d27afe": "vote(uint256,bool)",
		"e9523fb5": "voteLength()",
		"cec5b622": "voterIdx(address)",
	},
	Bin: "0x60a06040523060601b6080523480156200001857600080fd5b506200002362000029565b620000eb565b600054610100900460ff1615620000965760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015620000e9576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60805160601c6187816200010a60003960006129c801526187816000f3fe6080604052600436106104675760003560e01c806372016f751161024a578063af6af2ff11610139578063cec5b622116100b6578063e27bdaef1161007a578063e27bdaef14610f98578063e75235b814610fb8578063e9523fb514610fcd578063f2fde38b14610fe3578063f38ecf471461100357610467565b8063cec5b62214610ef5578063d07bff0c14610f22578063d6f9cfce14610f58578063d965ea0014610f6e578063de09b37714610f8357610467565b8063c6713baf116100fd578063c6713baf14610e22578063c7d3da3414610e56578063c9d27afe14610e8a578063ce04b9d414610eaa578063ce6a54ff14610ebf57610467565b8063af6af2ff14610d3d578063b128f88014610d52578063c00ace6c14610d86578063c0b4fe1514610dba578063c42a0abc14610dee57610467565b8063a0c12683116101c7578063a8915a3e1161018b578063a8915a3e14610c66578063a91ee0dc14610c86578063a9b629b214610ca6578063aaf0dd3614610cda578063ab3545e514610d0757610467565b8063a0c1268314610b8d578063a230c52414610bba578063a6868b7d14610bda578063a7771ee314610c0e578063a78a818814610c4657610467565b8063894f51111161020e578063894f511114610af55780638d39e33a14610b155780638da5cb5b14610b35578063918f867414610b535780639986e4b914610b6957610467565b806372016f7514610a55578063738fdd1a14610a6a5780637b2bfb0114610a8a5780637bf4653014610abe5780637d10dd1b14610adf57610467565b80633f35c8fe116103665780635aaa4040116102e35780636d583ca7116102a75780636d583ca7146109805780636f1e8533146109b45780636f6de96d146109ec5780636fde207a14610a0c578063715018a614610a4057610467565b80635aaa4040146108c55780636167eb45146108da578063656e30521461090e5780636ba99181146109425780636c78d2cf1461095757610467565b80634d5ce0381161032a5780634d5ce0381461082e5780634f0f4aa91461084e5780634f1ef2861461087e57806352d1902d1461088c5780635a731cca146108a157610467565b80633f35c8fe14610750578063406903531461078457806346946416146107a45780634bd1ed76146107d85780634d273e28146107fa57610467565b8063238737b6116103f4578063351bacda116103b8578063351bacda1461069c5780633659cfe6146106bc57806336e83d83146106dc57806338294419146106fc578063397e38e71461073057610467565b8063238737b6146105dd578063278bb12a146106115780632f40992e146106455780633310569c1461066657806334125c841461067c57610467565b806316fbe8311161043b57806316fbe831146105385780631c12b0301461054f5780631c150171146105745780631c4b774b146105895780631e0cba0d146105bf57610467565b806215a73b1461046c5780630efa4909146104b3578063139d9dd3146104d357806315bf6b4d14610500575b600080fd5b34801561047857600080fd5b506104a07f0c4fbe9dc9de15dd7c0d064975ee1a2f2f9b954fa0e65d4f6cddba94884bdc3e81565b6040519081526020015b60405180910390f35b3480156104bf57600080fd5b506104a06104ce366004617b08565b611037565b3480156104df57600080fd5b506104a06104ee3660046179a9565b60ab6020526000908152604090205481565b34801561050c57600080fd5b5061052061051b366004617f1a565b6113f0565b6040516001600160a01b0390911681526020016104aa565b34801561054457600080fd5b5061054d61140e565b005b34801561055b57600080fd5b50610564611752565b60405190151581526020016104aa565b34801561058057600080fd5b506104a06117a6565b34801561059557600080fd5b506105206105a4366004617f1a565b6000908152606a60205260409020546001600160a01b031690565b3480156105cb57600080fd5b506104a0665374616b696e6760c81b81565b3480156105e957600080fd5b506104a07f1d36f8ce53f59e624857e1d8dc7932d19981a2ea1b8faa4eb8ff843fc3e5a27881565b34801561061d57600080fd5b506104a07f9b2e0c7fdae148f225bae7deb92d7e7bd24bb77edb12956e8fa7204900dd8a2281565b34801561065157600080fd5b506104a06914995dd85c99141bdbdb60b21b81565b34801561067257600080fd5b506104a060aa5481565b34801561068857600080fd5b506104a06845636f73797374656d60b81b81565b3480156106a857600080fd5b5061054d6106b7366004617b9f565b611825565b3480156106c857600080fd5b5061054d6106d73660046179a9565b611eb4565b3480156106e857600080fd5b506104a06106f7366004617e8c565b611eed565b34801561070857600080fd5b506104a07fbe90e461bbdb9a95a694f7796912ea04244caf7f5b60ad7ded17e16821d3e44c81565b34801561073c57600080fd5b5061054d61074b366004617b5e565b612414565b34801561075c57600080fd5b506104a07f2a268972a70c8c688b62366bdfdd9bb09cf19d3e5b6e7e7bb158e671ffdcedd281565b34801561079057600080fd5b506104a061079f366004617d86565b61250f565b3480156107b057600080fd5b506104a07fdd5a41a7fc01f5c6d30816b17f638d6531625f1e1eaa599673ab2f6079f2dd9d81565b3480156107e457600080fd5b506104a06a4d61696e74656e616e636560a81b81565b34801561080657600080fd5b506104a07f77884798208df1e64f70968be41ef2d3211ec53613397ca59313416813df088881565b34801561083a57600080fd5b506105646108493660046179a9565b6127c1565b34801561085a57600080fd5b5061086e610869366004617f1a565b6127de565b6040516104aa9493929190618146565b61054d6106d7366004617abb565b34801561089857600080fd5b506104a06129bb565b3480156108ad57600080fd5b506104a06c14dd185ada5b99d4995dd85c99609a1b81565b3480156108d157600080fd5b506104a0612a6e565b3480156108e657600080fd5b506104a07f9f1de481f899d76889aa8a2b44cc7b604d42691aa93d4ba618a1a1fd439f505081565b34801561091a57600080fd5b506104a07fe10074dceffb75f13bf0ce50145afd35182d63796823f1280ce40e01c19109e781565b34801561094e57600080fd5b5061054d612ab0565b34801561096357600080fd5b506104a07111dbdd995c9b985b98d950dbdb9d1c9858dd60721b81565b34801561098c57600080fd5b506104a07fc69fc6b7d0efc934fd5a3581c7253a7107a952526bb6dbcd814ef8d8dae1f44a81565b3480156109c057600080fd5b506105646109cf3660046179a9565b6001600160a01b03166000908152606d6020526040902054151590565b3480156109f857600080fd5b50610520610a073660046179a9565b612bad565b348015610a1857600080fd5b506104a07f0b09c9badbbeb6c813a598ee910770a39ccda797a1940439bb6e47fc6c87548b81565b348015610a4c57600080fd5b5061054d612c1c565b348015610a6157600080fd5b506074546104a0565b348015610a7657600080fd5b50606554610520906001600160a01b031681565b348015610a9657600080fd5b506104a07f9346226931826838eedd13d9677fa33551e7c81f604b171ef3fac355458da9aa81565b348015610aca57600080fd5b506104a069456e7653746f7261676560b01b81565b348015610aeb57600080fd5b506104a060665481565b348015610b0157600080fd5b506104a0610b10366004617c9d565b612c30565b348015610b2157600080fd5b506104a0610b303660046179a9565b612fad565b348015610b4157600080fd5b506033546001600160a01b0316610520565b348015610b5f57600080fd5b506104a061271081565b348015610b7557600080fd5b506104a06c42616c6c6f7453746f7261676560981b81565b348015610b9957600080fd5b506104a0610ba83660046179a9565b606d6020526000908152604090205481565b348015610bc657600080fd5b50610564610bd53660046179a9565b6138f5565b348015610be657600080fd5b506104a07f6c6f69f426081752a5d3e73746599acd2a4cb145d5de4203ca1e3473b281680b81565b348015610c1a57600080fd5b50610564610c293660046179a9565b6001600160a01b0316600090815260686020526040902054151590565b348015610c5257600080fd5b506104a0610c61366004617ebe565b613938565b348015610c7257600080fd5b5061054d610c81366004617bf5565b6140b3565b348015610c9257600080fd5b5061054d610ca13660046179a9565b61452b565b348015610cb257600080fd5b506104a07f89dd490ecaf395283ed4ff2fd9557ca767fc425dce063451a9b0da6d72f600c381565b348015610ce657600080fd5b506104a0610cf53660046179a9565b606b6020526000908152604090205481565b348015610d1357600080fd5b50610520610d22366004617f1a565b6000908152606c60205260409020546001600160a01b031690565b348015610d4957600080fd5b506104a06145d3565b348015610d5e57600080fd5b506104a07f829561ab7af084b7efc6600518d2df79b8d95f3f4c3a550f54f8f7ec7d2b805781565b348015610d9257600080fd5b506104a07f18ad4415ef4a621ce1a136395c51ab6c3712bb2e24b79d526059925cea58dcb881565b348015610dc657600080fd5b506104a07f8086da5becff4dfac91a3105821b361078d2d4abba0ccc2401b974cf0dcf05c181565b348015610dfa57600080fd5b506104a07fb38b2c133e931937bd95337c65c8aefa7040ed64bbed732e3e29a4944c65747381565b348015610e2e57600080fd5b506104a07fc9e15e34073efbcd0328740feaf503caac9124b55b38c73d1a97b53da80a2f6081565b348015610e6257600080fd5b506104a07f04f7b94450bbcad85f37ea47cd1062728f884bb2040e357738f8fd53056134bc81565b348015610e9657600080fd5b5061054d610ea5366004617f32565b614615565b348015610eb657600080fd5b506104a061481a565b348015610ecb57600080fd5b506104a0610eda3660046179a9565b6001600160a01b031660009081526072602052604090205490565b348015610f0157600080fd5b506104a0610f103660046179a9565b60686020526000908152604090205481565b348015610f2e57600080fd5b50610520610f3d366004617f1a565b6000908152606760205260409020546001600160a01b031690565b348015610f6457600080fd5b506104a060755481565b348015610f7a57600080fd5b506069546104a0565b348015610f8f57600080fd5b506077546104a0565b348015610fa457600080fd5b5061054d610fb3366004617f1a565b61485c565b348015610fc457600080fd5b506113896104a0565b348015610fd957600080fd5b506104a060765481565b348015610fef57600080fd5b5061054d610ffe3660046179a9565b6148e2565b34801561100f57600080fd5b506104a07f7c1150f0e1a39ff55552d52764f97e6c387e2a247e1df344369f122c4254be2f81565b600061104161495b565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b15801561108257600080fd5b505afa158015611096573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110ba9190617d0f565b6110df5760405162461bcd60e51b81526004016110d6906182b9565b60405180910390fd5b60006110ea33610a07565b60aa546001600160a01b038216600090815260ab60205260409020549192509061111490426185fd565b10156111325760405162461bcd60e51b81526004016110d69061844e565b600061113d33610a07565b9050611147612a6e565b6111508261497b565b1115801561116d57506111616145d3565b61116a8261497b565b10155b6111895760405162461bcd60e51b81526004016110d6906183cc565b6001600160a01b0386166111df5760405162461bcd60e51b815260206004820152601d60248201527f496d706c656d656e746174696f6e2063616e6e6f74206265207a65726f00000060448201526064016110d6565b6000805160206186c5833981519152546001600160a01b03878116911614156112425760405162461bcd60e51b815260206004820152601560248201527453616d6520636f6e7472616374206164647265737360581b60448201526064016110d6565b856001600160a01b03166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561127b57600080fd5b505afa9250505080156112ab575060408051601f3d908101601f191682019092526112a891810190617d2b565b60015b6112c75760405162461bcd60e51b81526004016110d69061830e565b6000805160206186c583398151915281146112f45760405162461bcd60e51b81526004016110d690618222565b506075546113039060016185c5565b925061130d614a01565b6001600160a01b0316630a3a63fe607554600161132a91906185c5565b60046040516001600160e01b031960e085901b16815260048101929092526024820152604481018790523360648201526001600160a01b038916608482015260a401602060405180830381600087803b15801561138657600080fd5b505af115801561139a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113be9190617d2b565b506113c98386614a1c565b5060758290556001600160a01b0316600090815260ab602052604090204290559392505050565b6000818152607360205260409020546001600160a01b03165b919050565b600054600290610100900460ff16158015611430575060005460ff8083169116105b61144c5760405162461bcd60e51b81526004016110d69061826b565b6000805461010060ff1990911660ff84161761ff00191617905561146e614a87565b60015b606954811161171e576000818152607160205260408082208151608081019092528054829082906114a190618640565b80601f01602080910402602001604051908101604052809291908181526020018280546114cd90618640565b801561151a5780601f106114ef5761010080835404028352916020019161151a565b820191906000526020600020905b8154815290600101906020018083116114fd57829003601f168201915b5050505050815260200160018201805461153390618640565b80601f016020809104026020016040519081016040528092919081815260200182805461155f90618640565b80156115ac5780601f10611581576101008083540402835291602001916115ac565b820191906000526020600020905b81548152906001019060200180831161158f57829003601f168201915b505050505081526020016002820180546115c590618640565b80601f01602080910402602001604051908101604052809291908181526020018280546115f190618640565b801561163e5780601f106116135761010080835404028352916020019161163e565b820191906000526020600020905b81548152906001019060200180831161162157829003601f168201915b5050505050815260200160038201548152505090506001606e826000015160405161166991906180a2565b90815260405160209181900382018120805460ff191693151593909317909255820151600191606f9161169b916180a2565b908152602001604051809103902060006101000a81548160ff021916908315150217905550600160706000836040015184606001516040516020016116e19291906180be565b60408051808303601f19018152918152815160209283012083529082019290925201600020805460ff191691151591909117905550600101611471565b506000805461ff001916905560405160ff821681526000805160206186e5833981519152906020015b60405180910390a150565b600060775460001461179f57600061176b607754614ae1565b50915050600061177c607754614b78565b50915050600282141561179c574281101561179c576000925050506117a3565b50505b5060015b90565b60006117b0614c02565b6001600160a01b03166333be496e6040518163ffffffff1660e01b815260040160206040518083038186803b1580156117e857600080fd5b505afa1580156117fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118209190617d2b565b905090565b600054610100900460ff16158080156118455750600054600160ff909116105b8061185f5750303b15801561185f575060005460ff166001145b61187b5760405162461bcd60e51b81526004016110d69061826b565b6000805460ff19166001179055801561189e576000805461ff0019166101001790555b6118a6614c1a565b6118ae614c49565b6118b78461452b565b4360665560006118c5614c78565b90506118cf6145d3565b84101580156118e55750836118e2612a6e565b10155b6119015760405162461bcd60e51b81526004016110d69061835c565b600080600060608060606000806000905060008060208d0191508c518261192891906185c5565b90505b80821015611e6557815199506119426020836185c5565b915080821061195057600080fd5b8151985061195f6020836185c5565b915080821061196d57600080fd5b8151975061197c6020836185c5565b915080821061198a57600080fd5b8196508651602061199b91906185c5565b6119a590836185c5565b91508082106119b357600080fd5b819550855160206119c491906185c5565b6119ce90836185c5565b91508082106119dc57600080fd5b819450845160206119ed91906185c5565b6119f790836185c5565b9150808210611a0557600080fd5b81519350611a146020836185c5565b9150611a216001846185c5565b9250611a2c8a6138f5565b158015611a3f5750611a3d896138f5565b155b8015611a515750611a4f886127c1565b155b611a8e5760405162461bcd60e51b815260206004820152600e60248201526d20b63932b0b23c9036b2b6b132b960911b60448201526064016110d6565b886067600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606860008b6001600160a01b03166001600160a01b031681526020019081526020016000208190555087606a600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606b60008a6001600160a01b03166001600160a01b031681526020019081526020016000208190555089606c600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606d60008c6001600160a01b03166001600160a01b0316815260200190815260200160002081905550886001600160a01b03168a6001600160a01b03167f6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba660405160405180910390a36040516325d998bb60e01b81526001600160a01b038b811660048301528f91908d16906325d998bb9060240160206040518083038186803b158015611c3857600080fd5b505afa158015611c4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c709190617d2b565b1015611cb55760405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e74207374616b696e6760601b60448201526064016110d6565b611cc187878787614c8d565b611cdd5760405162461bcd60e51b81526004016110d6906182e0565b611ce78a8f614d3c565b600083815260716020908152604090912088519091611d0a9183918b019061769b565b508651611d2090600183019060208a019061769b565b508551611d36906002830190602089019061769b565b508481600301819055506001606e89604051611d5291906180a2565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f90611d84908a906180a2565b908152602001604051809103902060006101000a81548160ff0219169083151502179055506001607060008888604051602001611dc29291906180be565b60405160208183030381529060405280519060200120815260200190815260200160002060006101000a81548160ff0219169083151502179055508a6073600086815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555083607260008d6001600160a01b03166001600160a01b03168152602001908152602001600020819055505061192b565b50506069819055607455505086159550611eae945050505050576000805461ff0019169055604051600181526000805160206186e5833981519152906020015b60405180910390a15b50505050565b60405162461bcd60e51b815260206004820152600e60248201526d496e76616c69642061636365737360901b60448201526064016110d6565b6000611ef761495b565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b158015611f3857600080fd5b505afa158015611f4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f709190617d0f565b611f8c5760405162461bcd60e51b81526004016110d6906182b9565b6000611f9733610a07565b60aa546001600160a01b038216600090815260ab602052604090205491925090611fc190426185fd565b1015611fdf5760405162461bcd60e51b81526004016110d69061844e565b6000611fea33610a07565b9050611ff4612a6e565b611ffd8261497b565b1115801561201a575061200e6145d3565b6120178261497b565b10155b6120365760405162461bcd60e51b81526004016110d6906183cc565b602084015184906001600160a01b03166120825760405162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b2103b37ba32b960991b60448201526064016110d6565b6000816060015151116120cb5760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964206e6f6465206e616d6560781b60448201526064016110d6565b60008160a0015151116121125760405162461bcd60e51b815260206004820152600f60248201526e0496e76616c6964206e6f646520495608c1b60448201526064016110d6565b60008160c001511161215a5760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a59081b9bd919481c1bdc9d607a1b60448201526064016110d6565b6000816080015151116121a45760405162461bcd60e51b8152602060048201526012602482015271496e76616c6964206e6f646520656e6f646560701b60448201526064016110d6565b600081610100015151116121e95760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964206d656d6f60a01b60448201526064016110d6565b6000816101200151116122315760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b210323ab930ba34b7b760811b60448201526064016110d6565b6122396145d3565b8160e0015110158015612257575061224f612a6e565b8160e0015111155b6122995760405162461bcd60e51b8152602060048201526013602482015272125b9d985b1a59081b1bd8dac8105b5bdd5b9d606a1b60448201526064016110d6565b84516122a4906138f5565b1580156122b9575084516122b7906127c1565b155b6122f65760405162461bcd60e51b815260206004820152600e60248201526d20b63932b0b23c9036b2b6b132b960911b60448201526064016110d6565b84602001516001600160a01b031685600001516001600160a01b0316148015612338575084604001516001600160a01b031685600001516001600160a01b0316145b61237a5760405162461bcd60e51b815260206004820152601360248201527229ba30b5b2b91034b9903737ba103b37ba32b960691b60448201526064016110d6565b612396856060015186608001518760a001518860c00151614c8d565b6123b25760405162461bcd60e51b81526004016110d6906182e0565b6075546123c09060016185c5565b93506123d184600133600089614d7b565b6123df848660e00151614e1e565b6123ee84866101000151614a1c565b505060758290556001600160a01b0316600090815260ab60205260409020429055919050565b600054610100900460ff16158080156124345750600054600160ff909116105b8061244e5750303b15801561244e575060005460ff166001145b61246a5760405162461bcd60e51b81526004016110d69061826b565b6000805460ff19166001179055801561248d576000805461ff0019166101001790555b612495614c1a565b61249d614c49565b6124a68461452b565b60668390556124b4826148e2565b60405133907fab2db0a6f442428b686ffa80eadcaabe7d5ee00049c6ae888a237edd3238d85690600090a28015611eae576000805461ff0019169055604051600181526000805160206186e583398151915290602001611ea5565b600061251961495b565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b15801561255a57600080fd5b505afa15801561256e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125929190617d0f565b6125ae5760405162461bcd60e51b81526004016110d6906182b9565b60006125b933610a07565b60aa546001600160a01b038216600090815260ab6020526040902054919250906125e390426185fd565b10156126015760405162461bcd60e51b81526004016110d69061844e565b600061260c33610a07565b9050612616612a6e565b61261f8261497b565b1115801561263c57506126306145d3565b6126398261497b565b10155b6126585760405162461bcd60e51b81526004016110d6906183cc565b8660011115801561266a575060098711155b6126a55760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964207479706560a01b60448201526064016110d6565b6126af8887614e5c565b6126eb5760405162461bcd60e51b815260206004820152600d60248201526c496e76616c69642076616c756560981b60448201526064016110d6565b6075546126f99060016185c5565b9250612703614a01565b6001600160a01b0316634a57823e84600587338d8d8d6040518863ffffffff1660e01b815260040161273b9796959493929190618522565b602060405180830381600087803b15801561275557600080fd5b505af1158015612769573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061278d9190617d2b565b506127988386614a1c565b5060758290556001600160a01b0316600090815260ab6020526040902042905595945050505050565b6001600160a01b03166000908152606b6020526040902054151590565b60008181526071602052604081206003810154815460609384938493919290916001830191600284019190849061281490618640565b80601f016020809104026020016040519081016040528092919081815260200182805461284090618640565b801561288d5780601f106128625761010080835404028352916020019161288d565b820191906000526020600020905b81548152906001019060200180831161287057829003601f168201915b505050505093508280546128a090618640565b80601f01602080910402602001604051908101604052809291908181526020018280546128cc90618640565b80156129195780601f106128ee57610100808354040283529160200191612919565b820191906000526020600020905b8154815290600101906020018083116128fc57829003601f168201915b5050505050925081805461292c90618640565b80601f016020809104026020016040519081016040528092919081815260200182805461295890618640565b80156129a55780601f1061297a576101008083540402835291602001916129a5565b820191906000526020600020905b81548152906001019060200180831161298857829003601f168201915b5050505050915093509350935093509193509193565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612a5b5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016110d6565b506000805160206186c583398151915290565b6000612a78614c02565b6001600160a01b031663737c59b86040518163ffffffff1660e01b815260040160206040518083038186803b1580156117e857600080fd5b612ab861495b565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b158015612af957600080fd5b505afa158015612b0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b319190617d0f565b612b4d5760405162461bcd60e51b81526004016110d6906182b9565b612b55611752565b15612b985760405162461bcd60e51b8152602060048201526013602482015272159bdd1a5b99c81a5cc81b9bdd08195b991959606a1b60448201526064016110d6565b607754612ba6906004614eea565b6000607755565b6001600160a01b0381166000908152606d602052604081205415612bd2575080611409565b6001600160a01b0382166000908152606860205260409020541561140957506001600160a01b039081166000908152606860209081526040808320548352606c9091529020541690565b612c24614a87565b612c2e6000614f28565b565b6000612c3a61495b565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b158015612c7b57600080fd5b505afa158015612c8f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cb39190617d0f565b612ccf5760405162461bcd60e51b81526004016110d6906182b9565b6000612cda33610a07565b60aa546001600160a01b038216600090815260ab602052604090205491925090612d0490426185fd565b1015612d225760405162461bcd60e51b81526004016110d69061844e565b6000612d2d33610a07565b9050612d37612a6e565b612d408261497b565b11158015612d5d5750612d516145d3565b612d5a8261497b565b10155b612d795760405162461bcd60e51b81526004016110d6906183cc565b6001600160a01b038916612dc15760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b60448201526064016110d6565b6001600160a01b0389166000908152606d6020526040902054612e135760405162461bcd60e51b815260206004820152600a6024820152692737b716b6b2b6b132b960b11b60448201526064016110d6565b6001612e1e60695490565b11612e6b5760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f742072656d6f7665206120736f6c65206d656d626572000000000060448201526064016110d6565b87612e758a61497b565b1015612ed65760405162461bcd60e51b815260206004820152602a60248201527f496e73756666696369656e742062616c616e636520746861742063616e206265604482015269103ab73637b1b5b2b21760b11b60648201526084016110d6565b607554612ee49060016185c5565b604080516101408101825260008082526020808301829052828401829052835182815280820185526060840152835182815280820185526080840152835182815290810190935260a082019290925260c081019190915260e081018a905261010081018990526101208101889052909350612f63846002338d85614d7b565b612f6d848a614e1e565b612f778489614a1c565b612f82848787614f7a565b505060758290556001600160a01b0316600090815260ab602052604090204290559695505050505050565b60008054610100900460ff1615808015612fce5750600054600160ff909116105b80612fe85750303b158015612fe8575060005460ff166001145b6130045760405162461bcd60e51b81526004016110d69061826b565b6000805460ff191660011790558015613027576000805461ff0019166101001790555b61302f614c1a565b613037614c49565b60008390506130b0816001600160a01b031663738fdd1a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561307857600080fd5b505afa15801561308c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca191906179c5565b4360668190555061312b816001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156130f357600080fd5b505afa158015613107573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ffe91906179c5565b60015b816001600160a01b031663d965ea006040518163ffffffff1660e01b815260040160206040518083038186803b15801561316757600080fd5b505afa15801561317b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061319f9190617d2b565b81116136dc5760405163ab3545e560e01b8152600481018290526001600160a01b0383169063ab3545e59060240160206040518083038186803b1580156131e557600080fd5b505afa1580156131f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061321d91906179c5565b6000828152606c6020908152604080832080546001600160a01b0319166001600160a01b03958616179081905584168352606d909152908190208390555163341effc360e21b8152600481018390529083169063d07bff0c9060240160206040518083038186803b15801561329157600080fd5b505afa1580156132a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132c991906179c5565b600082815260676020908152604080832080546001600160a01b0319166001600160a01b0395861617908190558416835260689091529081902083905551631c4b774b60e01b81526004810183905290831690631c4b774b9060240160206040518083038186803b15801561333d57600080fd5b505afa158015613351573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061337591906179c5565b6000828152606a6020908152604080832080546001600160a01b0319166001600160a01b0395861617908190559093168252606b8152828220849055606984905582516080810184526060808252918101829052928301819052820152604051634f0f4aa960e01b8152600481018390526001600160a01b03841690634f0f4aa99060240160006040518083038186803b15801561341257600080fd5b505afa158015613426573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261344e9190810190617e01565b60608501819052604085018290526020850183905283855261347293929190614c8d565b6134be5760405162461bcd60e51b815260206004820152601760248201527f6e6f646520696e666f206973206475706c69636174656400000000000000000060448201526064016110d6565b6001606e82600001516040516134d491906180a2565b90815260405160209181900382018120805460ff191693151593909317909255820151600191606f91613506916180a2565b908152602001604051809103902060006101000a81548160ff0219169083151502179055506001607060008360400151846060015160405160200161354c9291906180be565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff1916941515949094179093558483526071815291208251805184936135a292849291019061769b565b5060208281015180516135bb926001850192019061769b565b50604082015180516135d791600284019160209091019061769b565b50606091909101516003909101556000828152606c6020818152604080842080546001600160a01b039081168652607284528286208890558786528154607385529583902080546001600160a01b031916968216969096179095556074879055929091529054905163139d9dd360e01b815290821660048201529084169063139d9dd39060240160206040518083038186803b15801561367657600080fd5b505afa15801561368a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136ae9190617d2b565b6000838152606c60209081526040808320546001600160a01b0316835260ab9091529020555060010161312e565b50806001600160a01b0316633310569c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561371657600080fd5b505afa15801561372a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061374e9190617d2b565b60aa81905550806001600160a01b031663d6f9cfce6040518163ffffffff1660e01b815260040160206040518083038186803b15801561378d57600080fd5b505afa1580156137a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137c59190617d2b565b607581905550806001600160a01b031663e9523fb56040518163ffffffff1660e01b815260040160206040518083038186803b15801561380457600080fd5b505afa158015613818573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061383c9190617d2b565b607681905550806001600160a01b031663de09b3776040518163ffffffff1660e01b815260040160206040518083038186803b15801561387b57600080fd5b505afa15801561388f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138b39190617d2b565b607755506000915080156138ef576000805461ff0019169055604051600181526000805160206186e58339815191529060200160405180910390a15b50919050565b6001600160a01b0381166000908152606d602052604081205415158061393257506001600160a01b03821660009081526068602052604090205415155b92915050565b600061394261495b565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b15801561398357600080fd5b505afa158015613997573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139bb9190617d0f565b6139d75760405162461bcd60e51b81526004016110d6906182b9565b60006139e233610a07565b60aa546001600160a01b038216600090815260ab602052604090205491925090613a0c90426185fd565b1015613a2a5760405162461bcd60e51b81526004016110d69061844e565b6000613a3533610a07565b9050613a3f612a6e565b613a488261497b565b11158015613a655750613a596145d3565b613a628261497b565b10155b613a815760405162461bcd60e51b81526004016110d6906183cc565b602087015187906001600160a01b0316613acd5760405162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b2103b37ba32b960991b60448201526064016110d6565b600081606001515111613b165760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964206e6f6465206e616d6560781b60448201526064016110d6565b60008160a001515111613b5d5760405162461bcd60e51b815260206004820152600f60248201526e0496e76616c6964206e6f646520495608c1b60448201526064016110d6565b60008160c0015111613ba55760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a59081b9bd919481c1bdc9d607a1b60448201526064016110d6565b600081608001515111613bef5760405162461bcd60e51b8152602060048201526012602482015271496e76616c6964206e6f646520656e6f646560701b60448201526064016110d6565b60008161010001515111613c345760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964206d656d6f60a01b60448201526064016110d6565b600081610120015111613c7c5760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b210323ab930ba34b7b760811b60448201526064016110d6565b613c846145d3565b8160e0015110158015613ca25750613c9a612a6e565b8160e0015111155b613ce45760405162461bcd60e51b8152602060048201526013602482015272125b9d985b1a59081b1bd8dac8105b5bdd5b9d606a1b60448201526064016110d6565b6001600160a01b038716613d305760405162461bcd60e51b8152602060048201526013602482015272496e76616c6964206f6c64204164647265737360681b60448201526064016110d6565b6001600160a01b0387166000908152606d6020526040902054613d825760405162461bcd60e51b815260206004820152600a6024820152692737b716b6b2b6b132b960b11b60448201526064016110d6565b6020808901516001600160a01b038981166000908152606d8452604080822054825260679094529290922054821691161480613dd35750866001600160a01b031688602001516001600160a01b0316145b80613dfd5750613de688602001516138f5565b158015613dfd5750613dfb88602001516127c1565b155b8015613e7f57506040808901516001600160a01b038981166000908152606d6020908152848220548252606a90529290922054821691161480613e555750866001600160a01b031688604001516001600160a01b0316145b80613e7f5750613e6888604001516138f5565b158015613e7f5750613e7d88604001516127c1565b155b613ebe5760405162461bcd60e51b815260206004820152601060248201526f20b63932b0b23c90309036b2b6b132b960811b60448201526064016110d6565b336001600160a01b038816148015613ee2575087516001600160a01b038881169116145b15613f375785158015613ef3575084155b613f325760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a59081c1c9bdc1bdcd85b60821b60448201526064016110d6565b613fe6565b87516001600160a01b03888116911614613fe657613f536145d3565b613f5d86886185c5565b1115613fe65760405162461bcd60e51b815260206004820152604c60248201527f496e76616c696420616d6f756e743a2028756e6c6f636b416d6f756e74202b2060448201527f736c617368696e6729206d75737420626520657175616c206f72206c6f77207460648201526b6f206d696e5374616b696e6760a01b608482015260a4016110d6565b607554613ff49060016185c5565b9350614004846003338a8c614d7b565b614012848960e00151614e1e565b61402184896101000151614a1c565b61402c848787614f7a565b6075849055336001600160a01b038816148015614055575087516001600160a01b038881169116145b1561408f57600061406585614b78565b9250505061407f8542834261407a91906185c5565b614fe9565b61408d85600360018061502e565b505b50506001600160a01b0316600090815260ab60205260409020429055949350505050565b600054610100900460ff16158080156140d35750600054600160ff909116105b806140ed5750303b1580156140ed575060005460ff166001145b6141095760405162461bcd60e51b81526004016110d69061826b565b6000805460ff19166001179055801561412c576000805461ff0019166101001790555b614134614c1a565b61413c614c49565b6141458761452b565b61414d6145d3565b8610158015614163575085614160612a6e565b10155b61417f5760405162461bcd60e51b81526004016110d69061835c565b6000614189614c78565b6040516325d998bb60e01b815233600482015290915087906001600160a01b038316906325d998bb9060240160206040518083038186803b1580156141cd57600080fd5b505afa1580156141e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906142059190617d2b565b101561424a5760405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e74207374616b696e6760601b60448201526064016110d6565b60405163282d3fdf60e01b8152336004820152602481018890526001600160a01b0382169063282d3fdf90604401600060405180830381600087803b15801561429257600080fd5b505af11580156142a6573d6000803e3d6000fd5b5050600160698190557f6bee784efeb983674392298ab585b22866bedf00ebb0eea949d1e66f3f50e71d8054336001600160a01b0319918216811790925560008281526068602090815260408083208690557ff585789965ba69220d5ce3dc1b444eb22ff546f2650694fef8fafe9c26560af98054851686179055606b82528083208690557fdcf345d7f6a8deb7427d0fee62009fa15100353a1c666b51bb5387b25addcfa98054909416909417909255606d825291822083905560748390559190526071815288517f169c6be1b0e6ab5de76b532e587a77340130ac65c5591db02be822dcf1dc0ed693506143a1925083918a019061769b565b5085516143b7906001830190602089019061769b565b5084516143cd906002830190602088019061769b565b508381600301819055506001606e886040516143e991906180a2565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f9061441b9089906180a2565b908152602001604051809103902060006101000a81548160ff02191690831515021790555060016070600087876040516020016144599291906180be565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff1916941515949094179093556074543380855260728352838520829055908452607390915281832080546001600160a01b03191682179055436066559051909182917f6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba69190a350508015614522576000805461ff0019169055604051600181526000805160206186e58339815191529060200160405180910390a15b50505050505050565b614533614a87565b6001600160a01b0381166145895760405162461bcd60e51b815260206004820152601a60248201527f416464726573732073686f756c64206265206e6f6e2d7a65726f00000000000060448201526064016110d6565b606580546001600160a01b0319166001600160a01b0383169081179091556040517f278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd90600090a250565b60006145dd614c02565b6001600160a01b031663076cd77f6040518163ffffffff1660e01b815260040160206040518083038186803b1580156117e857600080fd5b600260785414156146685760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016110d6565b600260785561467561495b565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b1580156146b657600080fd5b505afa1580156146ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906146ee9190617d0f565b61470a5760405162461bcd60e51b81526004016110d6906182b9565b600061471533610a07565b905061471f612a6e565b6147288261497b565b1115801561474557506147396145d3565b6147428261497b565b10155b6147615760405162461bcd60e51b81526004016110d6906183cc565b614769611752565b61479f5760405162461bcd60e51b8152602060048201526007602482015266115e1c1a5c995960ca1b60448201526064016110d6565b60006147aa846150bd565b90506147b684846151d2565b6000806147c2866152ac565b925092505060006147d261138990565b905080831015806147e35750808210155b806147f857506147f382846185c5565b612710145b1561480c5761480c8785848611600061502e565b505060016078555050505050565b6000614824614c02565b6001600160a01b0316631b27e01b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156117e857600080fd5b614864614a87565b610e1081106148ad5760405162461bcd60e51b81526020600482015260156024820152746e6577506572696f6420697320746f6f206c6f6e6760581b60448201526064016110d6565b60aa8190556040518181527f17c6f1d1ce638844b664872f5c6eecb7d150ec0c41187d7f85826a656ee7946f90602001611747565b6148ea614a87565b6001600160a01b03811661494f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016110d6565b61495881614f28565b50565b60006118207111dbdd995c9b985b98d950dbdb9d1c9858dd60721b6152e6565b6000614985614c78565b604051632c9aab9b60e11b81526001600160a01b03848116600483015291909116906359355736906024015b60206040518083038186803b1580156149c957600080fd5b505afa1580156149dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139329190617d2b565b60006118206c42616c6c6f7453746f7261676560981b6152e6565b614a24614a01565b6001600160a01b031663bce0dbc183836040518363ffffffff1660e01b8152600401614a51929190618105565b600060405180830381600087803b158015614a6b57600080fd5b505af1158015614a7f573d6000803e3d6000fd5b505050505050565b6033546001600160a01b03163314612c2e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016110d6565b6000806000614aee614a01565b6001600160a01b031663688ca5b2856040518263ffffffff1660e01b8152600401614b1b91815260200190565b60606040518083038186803b158015614b3357600080fd5b505afa158015614b47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614b6b9190617f84565b9250925092509193909250565b6000806000614b85614a01565b6001600160a01b03166309970688856040518263ffffffff1660e01b8152600401614bb291815260200190565b60606040518083038186803b158015614bca57600080fd5b505afa158015614bde573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614b6b9190617fb1565b600061182069456e7653746f7261676560b01b6152e6565b600054610100900460ff16614c415760405162461bcd60e51b81526004016110d690618403565b612c2e615363565b600054610100900460ff16614c705760405162461bcd60e51b81526004016110d690618403565b612c2e615391565b6000611820665374616b696e6760c81b6152e6565b604051600190606f90614ca19086906180a2565b9081526040519081900360200190205460ff1615614cbd575060005b606e85604051614ccd91906180a2565b9081526040519081900360200190205460ff1615614ce9575060005b60008383604051602001614cfe9291906180be565b60408051601f1981840301815291815281516020928301206000818152607090935291205490915060ff1615614d3357600091505b50949350505050565b614d44614c78565b60405163282d3fdf60e01b81526001600160a01b03848116600483015260248201849052919091169063282d3fdf90604401614a51565b614d83614a01565b6001600160a01b031663daacbb95868684610120015187878760000151886020015189604001518a606001518b608001518c60a001518d60c001516040518d63ffffffff1660e01b8152600401614de59c9b9a99989796959493929190618485565b600060405180830381600087803b158015614dff57600080fd5b505af1158015614e13573d6000803e3d6000fd5b505050505050505050565b614e26614a01565b604051633968764960e11b815260048101849052602481018390526001600160a01b0391909116906372d0ec9290604401614a51565b6000614e66614c02565b6001600160a01b0316639801bff984846040518363ffffffff1660e01b8152600401614e93929190618105565b60206040518083038186803b158015614eab57600080fd5b505afa158015614ebf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ee39190617d0f565b9392505050565b614ef2614a01565b60405163548f2cdd60e11b815260048101849052602481018390526001600160a01b03919091169063a91e59ba90604401614a51565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b614f82614a01565b604051632264085960e01b81526004810185905260248101849052604481018390526001600160a01b0391909116906322640859906064015b600060405180830381600087803b158015614fd557600080fd5b505af1158015614522573d6000803e3d6000fd5b614ff1614a01565b60405163605b78c360e11b81526004810185905260248101849052604481018390526001600160a01b03919091169063c0b6f18690606401614fbb565b600482156150a15750600360018414156150585761504b856153c1565b615053575060045b6150a1565b600284141561506a57615053856158f9565b600384141561507d5761504b8583615f9b565b600484141561508f5761505385616503565b60058414156150a1576150a1856165d4565b6150ab8582614eea565b816150b65760006077555b5050505050565b60008060006150cb84614ae1565b509092509050600181141561516a57607754156150fa5760405162461bcd60e51b81526004016110d690618389565b600061510585614b78565b925050506151116117a6565b8110156151345761512f85426151256117a6565b61407a90426185c5565b61515f565b8061513d61481a565b10156151505761512f854261512561481a565b61515f854261407a84826185c5565b5060778490556151cb565b60028114156151995760775484146151945760405162461bcd60e51b81526004016110d690618389565b6151cb565b60405162461bcd60e51b8152602060048201526007602482015266115e1c1a5c995960ca1b60448201526064016110d6565b5092915050565b600060765460016151e391906185c5565b905060006151f033612bad565b905060006151fd60695490565b615209906127106185dd565b905060008461521957600261521c565b60015b9050615226614a01565b6040516325918ae760e21b815260048101869052602481018890526001600160a01b038581166044830152606482018490526084820185905291909116906396462b9c9060a401600060405180830381600087803b15801561528757600080fd5b505af115801561529b573d6000803e3d6000fd5b505050607694909455505050505050565b60008060006152b9614a01565b6001600160a01b03166356ba988e856040518263ffffffff1660e01b8152600401614bb291815260200190565b606554604051630d2020dd60e01b8152600481018390526000916001600160a01b031690630d2020dd9060240160206040518083038186803b15801561532b57600080fd5b505afa15801561533f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061393291906179c5565b600054610100900460ff1661538a5760405162461bcd60e51b81526004016110d690618403565b6001607855565b600054610100900460ff166153b85760405162461bcd60e51b81526004016110d690618403565b612c2e33614f28565b60006153cf8260015b616727565b6000806000806000806000806153e48a616843565b98509850985098509850985098509850506153fe886138f5565b1561545d57896000805160206187058339815191526040516154449060208082526010908201526f20b63932b0b23c90309036b2b6b132b960811b604082015260600190565b60405180910390a2600098505050505050505050611409565b615466866127c1565b156154ae57896000805160206187058339815191526040516154449060208082526012908201527120b63932b0b23c9030903932bbb0b93232b960711b604082015260600190565b6154b66145d3565b8110806154c95750806154c7612a6e565b105b156154ea57896000805160206187058339815191526040516154449061835c565b806154f4896168f9565b10156155165789600080516020618705833981519152604051615444906181db565b866001600160a01b0316886001600160a01b03161415801561554a5750856001600160a01b0316886001600160a01b031614155b15615596578960008051602061870583398151915260405161544490602080825260169082015275496e76616c6964206d656d626572206164647265737360501b604082015260600190565b6155a285858585614c8d565b6155c25789600080516020618705833981519152604051615444906182e0565b6155cc8882614d3c565b600060695460016155dd91906185c5565b9050876067600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606860008a6001600160a01b03166001600160a01b031681526020019081526020016000208190555086606a600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606b6000896001600160a01b03166001600160a01b031681526020019081526020016000208190555088606c600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606d60008b6001600160a01b03166001600160a01b03168152602001908152602001600020819055506000607454600161571691906185c5565b600081815260716020908152604090912089519293509161573c918391908b019061769b565b50865161575290600183019060208a019061769b565b508551615768906002830190602089019061769b565b508481600301819055506001606e8960405161578491906180a2565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f906157b6908a906180a2565b908152602001604051809103902060006101000a81548160ff02191690831515021790555060016070600088886040516020016157f49291906180be565b60405160208183030381529060405280519060200120815260200190815260200160002060006101000a81548160ff0219169083151502179055508a6073600084815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081607260008d6001600160a01b03166001600160a01b0316815260200190815260200160002081905550826069819055508160748190555043606681905550896001600160a01b03168b6001600160a01b03167f6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba660405160405180910390a35060019c9b505050505050505050505050565b6159048160026153ca565b600061590f82616843565b50505050505050509050615922816138f5565b615979578160008051602061870583398151915260405161596b906020808252601490820152732737ba1030b63932b0b23c90309036b2b6b132b960611b604082015260600190565b60405180910390a250614958565b6001600160a01b0381166000908152606d602052604090205415806159b457506001600160a01b038116600090815260726020526040902054155b156159fe578160008051602061870583398151915260405161596b906020808252601490820152732737ba1030b63932b0b23c90309036b2b6b132b960611b604082015260600190565b6001615a0960695490565b11615a5c578160008051602061870583398151915260405161596b906020808252601b908201527f43616e6e6f742072656d6f7665206120736f6c65206d656d6265720000000000604082015260600190565b6001600160a01b038082166000908152606d60209081526040808320548084526067835281842054606a9093529220546069549293918216929116908314615c10576069546000908152606c602090815260408083205486845281842080546001600160a01b0319166001600160a01b03928316908117909155808552606d84528285208890559085168452606b9092529091205480615b355760405162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e4caeec2e4c840d2dcc8caf60631b60448201526064016110d6565b6069546000908152606a602090815260408083205484845281842080546001600160a01b0319166001600160a01b03928316908117909155808552606b8452828520869055908816845260689092529091205480615bcb5760405162461bcd60e51b8152602060048201526013602482015272092dcecc2d8d2c840ecdee8cae440d2dcc8caf606b1b60448201526064016110d6565b60695460009081526067602090815260408083205484845281842080546001600160a01b0319166001600160a01b039092169182179055835260689091529020555050505b606980546000908152606c6020908152604080832080546001600160a01b03199081169091556001600160a01b038981168552606d845282852085905585548552606a84528285208054831690558681168552606b84528285208590558554855260678452828520805490921690915586168352606890915281205554615c99906001906185fd565b6069556001600160a01b03841660009081526072602052604090205480615cf75760405162461bcd60e51b8152602060048201526012602482015271092dcecc2d8d2c840dcdec8ca40d2dcc8caf60731b60448201526064016110d6565b6000818152607160205260408082209051909190606f90615d1c9060018501906180e0565b908152604051908190036020018120805492151560ff1990931692909217909155600090606e90615d4e9084906180e0565b90815260405160209181900382018120805460ff19169315159390931790925560038301546000926070928492615d8a926002880192016180ec565b60408051808303601f19018152918152815160209283012083529082019290925201600020805460ff19169115159190911790556074548214615ebd57607454600090815260736020908152604080832054607190925290912080546001600160a01b0390921691839190615dfe90618640565b615e0992919061771f565b5060745460009081526071602052604090206001908101805491840191615e2f90618640565b615e3a92919061771f565b5060745460009081526071602052604090206002908101805491840191615e6090618640565b615e6b92919061771f565b506074546000908152607160209081526040808320600390810154908601558583526073825280832080546001600160a01b039095166001600160a01b03199095168517905592825260729052208290555b60748054600090815260736020908152604080832080546001600160a01b03191690556001600160a01b038a16835260728252808320839055925482526071905290812090615f0c828261779a565b615f1a60018301600061779a565b615f2860028301600061779a565b600382016000905550506001607454615f4191906185fd565b60745543606655615f528787616933565b836001600160a01b0316866001600160a01b03167faa91016c21c52c58ac64f23f71bbe75becc9ada603e18ee671d09ff15492d1c160405160405180910390a350505050505050565b600081615fad57615fad8360036153ca565b6000806000806000806000806000615fc48c616843565b985098509850985098509850985098509850615fdf896138f5565b616019578b600080516020618705833981519152604051615fff906181a4565b60405180910390a260009950505050505050505050613932565b61602c8c8c8b8b8b8b8b8b8b8b8b616b10565b6160425760009950505050505050505050613932565b6001600160a01b03808a166000818152606d6020526040902054918a16146160b1576000818152606c6020908152604080832080546001600160a01b0319166001600160a01b038e81169182179092558452606d909252808320849055908c1682528120556160b18983614d3c565b6001600160a01b038a166000908152607260209081526040808320548084526071909252808320905191929091606e906160ec9084906180e0565b908152604051908190036020018120805492151560ff1990931692909217909155600090606f906161219060018501906180e0565b90815260405160209181900382018120805460ff1916931515939093179092556003830154600092607092849261615d926002880192016180ec565b60408051808303601f1901815291815281516020928301208352828201939093529101600020805460ff19169215159290921790915588516161a4918391908b019061769b565b5086516161ba90600183019060208a019061769b565b5085516161d0906002830190602089019061769b565b506003810185905543606655604051600190606e906161f0908b906180a2565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f90616222908a906180a2565b908152602001604051809103902060006101000a81548160ff02191690831515021790555060016070600088886040516020016162609291906180be565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff191694151594909417909355858352606a90529020546001600160a01b039081169150891681146162fc576000838152606a6020908152604080832080546001600160a01b0319166001600160a01b038e81169182179092558452606b90925280832086905590831682528120555b506000828152606760205260409020546001600160a01b03908116908a1681146163ae57896067600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606860008c6001600160a01b03166001600160a01b0316815260200190815260200160002081905550600060686000836001600160a01b03166001600160a01b03168152602001908152602001600020819055505b50896001600160a01b03168b6001600160a01b0316146164ae57896073600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080607260008c6001600160a01b03166001600160a01b03168152602001908152602001600020819055506000607260008d6001600160a01b03166001600160a01b031681526020019081526020016000208190555061645f8e8c616933565b886001600160a01b03168a6001600160a01b03168c6001600160a01b03167f15f4d750630db473a85edd9d47c500527a2648cc5e676f39645e52790cf07be060405160405180910390a46164ef565b896001600160a01b03168b6001600160a01b03167f1feee1b4fcb797c62645da41c5c6edd5f91d4291de0054da625c42b823594c1f60405160405180910390a35b5060019d9c50505050505050505050505050565b61650e8160046153ca565b6000616518614a01565b6001600160a01b0316637efa9ae3836040518263ffffffff1660e01b815260040161654591815260200190565b60206040518083038186803b15801561655d57600080fd5b505afa158015616571573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061659591906179c5565b90506001600160a01b038116156165d0576165af816170d4565b604080516000808252602082019092526165cb91839190617171565b436066555b5050565b6165df8160056153ca565b60008060006165ec614a01565b6001600160a01b0316631d940da2856040518263ffffffff1660e01b815260040161661991815260200190565b60006040518083038186803b15801561663157600080fd5b505afa158015616645573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261666d9190810190617d43565b925092509250600061667d614c02565b6040516388c2801960e01b81529091506001600160a01b038216906388c28019906166ae9087908690600401618105565b600060405180830381600087803b1580156166c857600080fd5b505af11580156166dc573d6000803e3d6000fd5b50504360665550506040517f701c16c2519cdb79aaac423a84733590e3510d9552055b6ad6908f0ab12b6c29906167189086908690869061811e565b60405180910390a15050505050565b60008061673384614ae1565b509150915082821461677d5760405162461bcd60e51b8152602060048201526013602482015272496e76616c696420766f74696e67207479706560681b60448201526064016110d6565b600281146167c45760405162461bcd60e51b8152602060048201526014602482015273496e76616c696420766f74696e6720737461746560601b60448201526064016110d6565b6000806167d0866152ac565b92509250506167de61138990565b821015806167ee57506113898110155b8061680357506167fe81836185c5565b612710145b614a7f5760405162461bcd60e51b8152602060048201526011602482015270139bdd081e595d08199a5b985b1a5e9959607a1b60448201526064016110d6565b6000806000806060806060600080616859614a01565b6001600160a01b03166373df4e018b6040518263ffffffff1660e01b815260040161688691815260200190565b60006040518083038186803b15801561689e57600080fd5b505afa1580156168b2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526168da91908101906179e1565b9850985098509850985098509850985098509193959799909294969850565b6000616903614c78565b6040516325d998bb60e01b81526001600160a01b03848116600483015291909116906325d998bb906024016149b1565b60008061693f8461726c565b9150915061694b6145d3565b61695582846185c5565b11156169e55760405162461bcd60e51b815260206004820152605360248201527f6d696e5374616b696e672076616c7565206d757374206265206772656174657260448201527f207468616e206f7220657175616c20746f207468652073756d206f6620756e6c6064820152726f636b416d6f756e742c20736c617368696e6760681b608482015260a4016110d6565b60006169ef614c78565b604051632c9aab9b60e11b81526001600160a01b03868116600483015291925060009183169063593557369060240160206040518083038186803b158015616a3657600080fd5b505afa158015616a4a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616a6e9190617d2b565b90506000616a7a6145d3565b616a8490836185fd565b905084821115616b0657616a9886866172fc565b604051634954a85b60e11b81526001600160a01b03878116600483015260248201869052604482018390528416906392a950b690606401600060405180830381600087803b158015616ae957600080fd5b505af1158015616afd573d6000803e3d6000fd5b50505050614522565b61452286836172fc565b60008a616b2257616b228c60036153ca565b616b2b8a6138f5565b616b5b578b600080516020618705833981519152604051616b4b906181a4565b60405180910390a25060006170c5565b6001600160a01b038a166000908152606d60205260409020541580616b9657506001600160a01b038a16600090815260726020526040902054155b15616bb7578b600080516020618705833981519152604051616b4b906181a4565b6001600160a01b03808b166000818152606d6020526040902054918b1614616d2e57616be28a6138f5565b15616c47578c600080516020618705833981519152604051616c35906020808252601f908201527f6e6577206164647265737320697320616c72656164792061206d656d62657200604082015260600190565b60405180910390a260009150506170c5565b886001600160a01b03168a6001600160a01b031614158015616c7b5750876001600160a01b03168a6001600160a01b031614155b15616cc6578c600080516020618705833981519152604051616c3590602080825260159082015274496e76616c696420766f746572206164647265737360581b604082015260600190565b616cce6145d3565b831080616ce1575082616cdf612a6e565b105b15616d02578c600080516020618705833981519152604051616c359061835c565b82616d0c8b6168f9565b1015616d2e578c600080516020618705833981519152604051616c35906181db565b6001600160a01b038b166000908152607260209081526040808320548084526071909252808320815160808101909252805492939282908290616d7090618640565b80601f0160208091040260200160405190810160405280929190818152602001828054616d9c90618640565b8015616de95780601f10616dbe57610100808354040283529160200191616de9565b820191906000526020600020905b815481529060010190602001808311616dcc57829003601f168201915b50505050508152602001600182018054616e0290618640565b80601f0160208091040260200160405190810160405280929190818152602001828054616e2e90618640565b8015616e7b5780601f10616e5057610100808354040283529160200191616e7b565b820191906000526020600020905b815481529060010190602001808311616e5e57829003601f168201915b50505050508152602001600282018054616e9490618640565b80601f0160208091040260200160405190810160405280929190818152602001828054616ec090618640565b8015616f0d5780601f10616ee257610100808354040283529160200191616f0d565b820191906000526020600020905b815481529060010190602001808311616ef057829003601f168201915b505050505081526020016003820154815250509050616f2f898989898561733b565b616f63578e600080516020618705833981519152604051616f4f906182e0565b60405180910390a2600093505050506170c5565b506000828152606a60205260409020546001600160a01b03908116908d8116908b1614801590616fa55750896001600160a01b0316816001600160a01b031614155b8015616fc45750616fb58a6138f5565b80616fc45750616fc48a6127c1565b15617010578e600080516020618705833981519152604051616f4f90602080825260169082015275496e76616c696420726577617264206164647265737360501b604082015260600190565b506000828152606760205260409020546001600160a01b03908116908d8116908c161480159061705257508a6001600160a01b0316816001600160a01b031614155b801561707157506170628b6138f5565b8061707157506170718b6127c1565b156170bd578e600080516020618705833981519152604051616f4f90602080825260169082015275496e76616c696420766f74657273206164647265737360501b604082015260600190565b506001925050505b9b9a5050505050505050505050565b6170dc61495b565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b15801561711d57600080fd5b505afa158015617131573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906171559190617d0f565b6149585760405162461bcd60e51b81526004016110d6906182b9565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156171a9576171a483617464565b617267565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156171e257600080fd5b505afa925050508015617212575060408051601f3d908101601f1916820190925261720f91810190617d2b565b60015b61722e5760405162461bcd60e51b81526004016110d69061830e565b6000805160206186c5833981519152811461725b5760405162461bcd60e51b81526004016110d690618222565b50617267838383617500565b505050565b600080617277614a01565b6001600160a01b0316638c7be692846040518263ffffffff1660e01b81526004016172a491815260200190565b604080518083038186803b1580156172bb57600080fd5b505afa1580156172cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906172f39190617f61565b91509150915091565b617304614c78565b604051637eee288d60e01b81526001600160a01b038481166004830152602482018490529190911690637eee288d90604401614a51565b6000600190508480519060200120826020015180519060200120141580156173825750606f8560405161736e91906180a2565b9081526040519081900360200190205460ff165b1561738b575060005b8580519060200120826000015180519060200120141580156173cc5750606e866040516173b891906180a2565b9081526040519081900360200190205460ff165b156173d5575060005b600084846040516020016173ea9291906180be565b604051602081830303815290604052805190602001209050808360400151846060015160405160200161741e9291906180be565b6040516020818303038152906040528051906020012014158015617450575060008181526070602052604090205460ff165b1561745a57600091505b5095945050505050565b6001600160a01b0381163b6174d15760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016110d6565b6000805160206186c583398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b61750983617525565b6000825111806175165750805b1561726757611eae8383617565565b61752e81617464565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060614ee383836040518060600160405280602781526020016187256027913960606001600160a01b0384163b6175ed5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016110d6565b600080856001600160a01b03168560405161760891906180a2565b600060405180830381855af49150503d8060008114617643576040519150601f19603f3d011682016040523d82523d6000602084013e617648565b606091505b5091509150617658828286617662565b9695505050505050565b60608315617671575081614ee3565b8251156176815782518084602001fd5b8160405162461bcd60e51b81526004016110d69190618191565b8280546176a790618640565b90600052602060002090601f0160209004810192826176c9576000855561770f565b82601f106176e257805160ff191683800117855561770f565b8280016001018555821561770f579182015b8281111561770f5782518255916020019190600101906176f4565b5061771b9291506177d2565b5090565b82805461772b90618640565b90600052602060002090601f01602090048101928261774d576000855561770f565b82601f1061775e578054855561770f565b8280016001018555821561770f57600052602060002091601f016020900482015b8281111561770f57825482559160010191906001019061777f565b5080546177a690618640565b6000825580601f106177b85750614958565b601f01602090049060005260206000209081019061495891905b5b8082111561771b57600081556001016177d3565b8035611409816186a1565b8051611409816186a1565b600082601f83011261780d578081fd5b813561782061781b8261859e565b61856e565b818152846020838601011115617834578283fd5b816020850160208301379081016020019190915292915050565b600082601f83011261785e578081fd5b815161786c61781b8261859e565b818152846020838601011115617880578283fd5b617891826020830160208701618614565b949350505050565b60006101408083850312156178ac578182fd5b6178b58161856e565b9150506178c1826177e7565b81526178cf602083016177e7565b60208201526178e0604083016177e7565b604082015260608201356001600160401b03808211156178ff57600080fd5b61790b858386016177fd565b6060840152608084013591508082111561792457600080fd5b617930858386016177fd565b608084015260a084013591508082111561794957600080fd5b617955858386016177fd565b60a084015260c084013560c084015260e084013560e08401526101009150818401358181111561798457600080fd5b617990868287016177fd565b8385015250505061012080830135818301525092915050565b6000602082840312156179ba578081fd5b8135614ee3816186a1565b6000602082840312156179d6578081fd5b8151614ee3816186a1565b60008060008060008060008060006101208a8c0312156179ff578485fd5b617a088a6177f2565b9850617a1660208b016177f2565b9750617a2460408b016177f2565b9650617a3260608b016177f2565b955060808a01516001600160401b0380821115617a4d578687fd5b617a598d838e0161784e565b965060a08c0151915080821115617a6e578586fd5b617a7a8d838e0161784e565b955060c08c0151915080821115617a8f578485fd5b50617a9c8c828d0161784e565b93505060e08a015191506101008a015190509295985092959850929598565b60008060408385031215617acd578182fd5b8235617ad8816186a1565b915060208301356001600160401b03811115617af2578182fd5b617afe858286016177fd565b9150509250929050565b600080600060608486031215617b1c578283fd5b8335617b27816186a1565b925060208401356001600160401b03811115617b41578283fd5b617b4d868287016177fd565b925050604084013590509250925092565b600080600060608486031215617b72578283fd5b8335617b7d816186a1565b9250602084013591506040840135617b94816186a1565b809150509250925092565b600080600060608486031215617bb3578283fd5b8335617bbe816186a1565b92506020840135915060408401356001600160401b03811115617bdf578182fd5b617beb868287016177fd565b9150509250925092565b60008060008060008060c08789031215617c0d578384fd5b8635617c18816186a1565b95506020870135945060408701356001600160401b0380821115617c3a578586fd5b617c468a838b016177fd565b95506060890135915080821115617c5b578384fd5b617c678a838b016177fd565b94506080890135915080821115617c7c578384fd5b50617c8989828a016177fd565b92505060a087013590509295509295509295565b60008060008060008060c08789031215617cb5578384fd5b8635617cc0816186a1565b95506020870135945060408701356001600160401b03811115617ce1578485fd5b617ced89828a016177fd565b945050606087013592506080870135915060a087013590509295509295509295565b600060208284031215617d20578081fd5b8151614ee3816186b6565b600060208284031215617d3c578081fd5b5051919050565b600080600060608486031215617d57578081fd5b835192506020840151915060408401516001600160401b03811115617d7a578182fd5b617beb8682870161784e565b600080600080600060a08688031215617d9d578283fd5b853594506020860135935060408601356001600160401b0380821115617dc1578485fd5b617dcd89838a016177fd565b94506060880135915080821115617de2578283fd5b50617def888289016177fd565b95989497509295608001359392505050565b60008060008060808587031215617e16578182fd5b84516001600160401b0380821115617e2c578384fd5b617e388883890161784e565b95506020870151915080821115617e4d578384fd5b617e598883890161784e565b94506040870151915080821115617e6e578384fd5b50617e7b8782880161784e565b606096909601519497939650505050565b600060208284031215617e9d578081fd5b81356001600160401b03811115617eb2578182fd5b61789184828501617899565b60008060008060808587031215617ed3578182fd5b84356001600160401b03811115617ee8578283fd5b617ef487828801617899565b9450506020850135617f05816186a1565b93969395505050506040820135916060013590565b600060208284031215617f2b578081fd5b5035919050565b60008060408385031215617f44578182fd5b823591506020830135617f56816186b6565b809150509250929050565b60008060408385031215617f73578182fd5b505080516020909101519092909150565b600080600060608486031215617f98578081fd5b83519250602084015191506040840151617b94816186b6565b600080600060608486031215617fc5578081fd5b8351925060208401519150604084015190509250925092565b60008151808452617ff6816020860160208601618614565b601f01601f19169290920160200192915050565b80546000906002810460018083168061802457607f831692505b602080841082141561804457634e487b7160e01b86526022600452602486fd5b818015618058576001811461806957618096565b60ff19861689528489019650618096565b60008881526020902060005b8681101561808e5781548b820152908501908301618075565b505084890196505b50505050505092915050565b600082516180b4818460208701618614565b9190910192915050565b600083516180d0818460208801618614565b9190910191825250602001919050565b6000614ee3828461800a565b60006180f8828561800a565b9283525050602001919050565b6000838252604060208301526178916040830184617fde565b60008482528360208301526060604083015261813d6060830184617fde565b95945050505050565b6000608082526181596080830187617fde565b828103602084015261816b8187617fde565b9050828103604084015261817f8186617fde565b91505082606083015295945050505050565b600060208252614ee36020830184617fde565b6020808252601b908201527f4f6c642061646472657373206973206e6f742061206d656d6265720000000000604082015260600190565b60208082526027908201527f496e73756666696369656e742062616c616e636520746861742063616e206265604082015266081b1bd8dad95960ca1b606082015260800190565b60208082526029908201527f45524331393637557067726164653a20756e737570706f727465642070726f786040820152681a58589b195555525160ba1b606082015260800190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6020808252600d908201526c2737902832b936b4b9b9b4b7b760991b604082015260600190565b6020808252601490820152734475706c696361746564206e6f646520696e666f60601b604082015260600190565b6020808252602e908201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960408201526d6f6e206973206e6f74205555505360901b606082015260800190565b602080825260139082015272125b9d985b1a59081b1bd8dac8185b5bdd5b9d606a1b604082015260600190565b60208082526023908201527f4e6f7720696e20766f74696e67207769746820646966666572656e742062616c6040820152621b1bdd60ea1b606082015260800190565b60208082526017908201527f496e76616c6964207374616b696e672062616c616e6365000000000000000000604082015260600190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b6020808252601d908201527f43616e6e6f74206164642070726f706f73616c20746f6f206561726c79000000604082015260600190565b8c8152602081018c9052604081018b90526001600160a01b038a81166060830152898116608083015288811660a083015287811660c0830152861660e08201526000610180806101008401526184dd81840188617fde565b90508281036101208401526184f28187617fde565b90508281036101408401526185078186617fde565b915050826101608301529d9c50505050505050505050505050565b600088825287602083015286604083015260018060a01b03861660608301528460808301528360a083015260e060c083015261856160e0830184617fde565b9998505050505050505050565b604051601f8201601f191681016001600160401b03811182821017156185965761859661868b565b604052919050565b60006001600160401b038211156185b7576185b761868b565b50601f01601f191660200190565b600082198211156185d8576185d8618675565b500190565b6000826185f857634e487b7160e01b81526012600452602481fd5b500490565b60008282101561860f5761860f618675565b500390565b60005b8381101561862f578181015183820152602001618617565b83811115611eae5750506000910152565b60028104600182168061865457607f821691505b602082108114156138ef57634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461495857600080fd5b801515811461495857600080fdfe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249885e7f4987c0698db47045ad8cea110b51138f0eecbd94915842328cf6c3dc97d416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220eb7cc0f68b66dc052002dff0485799467c5a0459f4a63b6a80dcfc5f6b431d8d64736f6c63430008020033",
}

// GovImpABI is the input ABI used to generate the binding from.
// Deprecated: Use GovImpMetaData.ABI instead.
var GovImpABI = GovImpMetaData.ABI

// Deprecated: Use GovImpMetaData.Sigs instead.
// GovImpFuncSigs maps the 4-byte function signature to its string representation.
var GovImpFuncSigs = GovImpMetaData.Sigs

// GovImpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovImpMetaData.Bin instead.
var GovImpBin = GovImpMetaData.Bin

// DeployGovImp deploys a new Ethereum contract, binding an instance of GovImp to it.
func DeployGovImp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GovImp, error) {
	parsed, err := GovImpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovImpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovImp{GovImpCaller: GovImpCaller{contract: contract}, GovImpTransactor: GovImpTransactor{contract: contract}, GovImpFilterer: GovImpFilterer{contract: contract}}, nil
}

// GovImp is an auto generated Go binding around an Ethereum contract.
type GovImp struct {
	GovImpCaller     // Read-only binding to the contract
	GovImpTransactor // Write-only binding to the contract
	GovImpFilterer   // Log filterer for contract events
}

// GovImpCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovImpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovImpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovImpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovImpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovImpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovImpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovImpSession struct {
	Contract     *GovImp           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovImpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovImpCallerSession struct {
	Contract *GovImpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GovImpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovImpTransactorSession struct {
	Contract     *GovImpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovImpRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovImpRaw struct {
	Contract *GovImp // Generic contract binding to access the raw methods on
}

// GovImpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovImpCallerRaw struct {
	Contract *GovImpCaller // Generic read-only contract binding to access the raw methods on
}

// GovImpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovImpTransactorRaw struct {
	Contract *GovImpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovImp creates a new instance of GovImp, bound to a specific deployed contract.
func NewGovImp(address common.Address, backend bind.ContractBackend) (*GovImp, error) {
	contract, err := bindGovImp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovImp{GovImpCaller: GovImpCaller{contract: contract}, GovImpTransactor: GovImpTransactor{contract: contract}, GovImpFilterer: GovImpFilterer{contract: contract}}, nil
}

// NewGovImpCaller creates a new read-only instance of GovImp, bound to a specific deployed contract.
func NewGovImpCaller(address common.Address, caller bind.ContractCaller) (*GovImpCaller, error) {
	contract, err := bindGovImp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovImpCaller{contract: contract}, nil
}

// NewGovImpTransactor creates a new write-only instance of GovImp, bound to a specific deployed contract.
func NewGovImpTransactor(address common.Address, transactor bind.ContractTransactor) (*GovImpTransactor, error) {
	contract, err := bindGovImp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovImpTransactor{contract: contract}, nil
}

// NewGovImpFilterer creates a new log filterer instance of GovImp, bound to a specific deployed contract.
func NewGovImpFilterer(address common.Address, filterer bind.ContractFilterer) (*GovImpFilterer, error) {
	contract, err := bindGovImp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovImpFilterer{contract: contract}, nil
}

// bindGovImp binds a generic wrapper to an already deployed contract.
func bindGovImp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovImpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovImp *GovImpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovImp.Contract.GovImpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovImp *GovImpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovImp.Contract.GovImpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovImp *GovImpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovImp.Contract.GovImpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovImp *GovImpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovImp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovImp *GovImpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovImp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovImp *GovImpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovImp.Contract.contract.Transact(opts, method, params...)
}

// BALLOTDURATIONMAXNAME is a free data retrieval call binding the contract method 0x0015a73b.
//
// Solidity: function BALLOT_DURATION_MAX_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BALLOTDURATIONMAXNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BALLOT_DURATION_MAX_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTDURATIONMAXNAME is a free data retrieval call binding the contract method 0x0015a73b.
//
// Solidity: function BALLOT_DURATION_MAX_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BALLOTDURATIONMAXNAME() ([32]byte, error) {
	return _GovImp.Contract.BALLOTDURATIONMAXNAME(&_GovImp.CallOpts)
}

// BALLOTDURATIONMAXNAME is a free data retrieval call binding the contract method 0x0015a73b.
//
// Solidity: function BALLOT_DURATION_MAX_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BALLOTDURATIONMAXNAME() ([32]byte, error) {
	return _GovImp.Contract.BALLOTDURATIONMAXNAME(&_GovImp.CallOpts)
}

// BALLOTDURATIONMINMAXNAME is a free data retrieval call binding the contract method 0x656e3052.
//
// Solidity: function BALLOT_DURATION_MIN_MAX_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BALLOTDURATIONMINMAXNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BALLOT_DURATION_MIN_MAX_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTDURATIONMINMAXNAME is a free data retrieval call binding the contract method 0x656e3052.
//
// Solidity: function BALLOT_DURATION_MIN_MAX_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BALLOTDURATIONMINMAXNAME() ([32]byte, error) {
	return _GovImp.Contract.BALLOTDURATIONMINMAXNAME(&_GovImp.CallOpts)
}

// BALLOTDURATIONMINMAXNAME is a free data retrieval call binding the contract method 0x656e3052.
//
// Solidity: function BALLOT_DURATION_MIN_MAX_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BALLOTDURATIONMINMAXNAME() ([32]byte, error) {
	return _GovImp.Contract.BALLOTDURATIONMINMAXNAME(&_GovImp.CallOpts)
}

// BALLOTDURATIONMINNAME is a free data retrieval call binding the contract method 0x6d583ca7.
//
// Solidity: function BALLOT_DURATION_MIN_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BALLOTDURATIONMINNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BALLOT_DURATION_MIN_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTDURATIONMINNAME is a free data retrieval call binding the contract method 0x6d583ca7.
//
// Solidity: function BALLOT_DURATION_MIN_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BALLOTDURATIONMINNAME() ([32]byte, error) {
	return _GovImp.Contract.BALLOTDURATIONMINNAME(&_GovImp.CallOpts)
}

// BALLOTDURATIONMINNAME is a free data retrieval call binding the contract method 0x6d583ca7.
//
// Solidity: function BALLOT_DURATION_MIN_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BALLOTDURATIONMINNAME() ([32]byte, error) {
	return _GovImp.Contract.BALLOTDURATIONMINNAME(&_GovImp.CallOpts)
}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BALLOTSTORAGENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BALLOT_STORAGE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BALLOTSTORAGENAME() ([32]byte, error) {
	return _GovImp.Contract.BALLOTSTORAGENAME(&_GovImp.CallOpts)
}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BALLOTSTORAGENAME() ([32]byte, error) {
	return _GovImp.Contract.BALLOTSTORAGENAME(&_GovImp.CallOpts)
}

// BASEFEEMAXCHANGERATENAME is a free data retrieval call binding the contract method 0xc42a0abc.
//
// Solidity: function BASE_FEE_MAX_CHANGE_RATE_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BASEFEEMAXCHANGERATENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BASE_FEE_MAX_CHANGE_RATE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BASEFEEMAXCHANGERATENAME is a free data retrieval call binding the contract method 0xc42a0abc.
//
// Solidity: function BASE_FEE_MAX_CHANGE_RATE_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BASEFEEMAXCHANGERATENAME() ([32]byte, error) {
	return _GovImp.Contract.BASEFEEMAXCHANGERATENAME(&_GovImp.CallOpts)
}

// BASEFEEMAXCHANGERATENAME is a free data retrieval call binding the contract method 0xc42a0abc.
//
// Solidity: function BASE_FEE_MAX_CHANGE_RATE_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BASEFEEMAXCHANGERATENAME() ([32]byte, error) {
	return _GovImp.Contract.BASEFEEMAXCHANGERATENAME(&_GovImp.CallOpts)
}

// BLOCKSPERNAME is a free data retrieval call binding the contract method 0x3f35c8fe.
//
// Solidity: function BLOCKS_PER_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BLOCKSPERNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BLOCKS_PER_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKSPERNAME is a free data retrieval call binding the contract method 0x3f35c8fe.
//
// Solidity: function BLOCKS_PER_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BLOCKSPERNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKSPERNAME(&_GovImp.CallOpts)
}

// BLOCKSPERNAME is a free data retrieval call binding the contract method 0x3f35c8fe.
//
// Solidity: function BLOCKS_PER_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BLOCKSPERNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKSPERNAME(&_GovImp.CallOpts)
}

// BLOCKCREATIONTIMENAME is a free data retrieval call binding the contract method 0xc0b4fe15.
//
// Solidity: function BLOCK_CREATION_TIME_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BLOCKCREATIONTIMENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BLOCK_CREATION_TIME_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKCREATIONTIMENAME is a free data retrieval call binding the contract method 0xc0b4fe15.
//
// Solidity: function BLOCK_CREATION_TIME_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BLOCKCREATIONTIMENAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKCREATIONTIMENAME(&_GovImp.CallOpts)
}

// BLOCKCREATIONTIMENAME is a free data retrieval call binding the contract method 0xc0b4fe15.
//
// Solidity: function BLOCK_CREATION_TIME_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BLOCKCREATIONTIMENAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKCREATIONTIMENAME(&_GovImp.CallOpts)
}

// BLOCKGASLIMITNAME is a free data retrieval call binding the contract method 0x238737b6.
//
// Solidity: function BLOCK_GASLIMIT_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BLOCKGASLIMITNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BLOCK_GASLIMIT_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKGASLIMITNAME is a free data retrieval call binding the contract method 0x238737b6.
//
// Solidity: function BLOCK_GASLIMIT_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BLOCKGASLIMITNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKGASLIMITNAME(&_GovImp.CallOpts)
}

// BLOCKGASLIMITNAME is a free data retrieval call binding the contract method 0x238737b6.
//
// Solidity: function BLOCK_GASLIMIT_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BLOCKGASLIMITNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKGASLIMITNAME(&_GovImp.CallOpts)
}

// BLOCKREWARDAMOUNTNAME is a free data retrieval call binding the contract method 0xa9b629b2.
//
// Solidity: function BLOCK_REWARD_AMOUNT_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BLOCKREWARDAMOUNTNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BLOCK_REWARD_AMOUNT_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREWARDAMOUNTNAME is a free data retrieval call binding the contract method 0xa9b629b2.
//
// Solidity: function BLOCK_REWARD_AMOUNT_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BLOCKREWARDAMOUNTNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDAMOUNTNAME(&_GovImp.CallOpts)
}

// BLOCKREWARDAMOUNTNAME is a free data retrieval call binding the contract method 0xa9b629b2.
//
// Solidity: function BLOCK_REWARD_AMOUNT_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BLOCKREWARDAMOUNTNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDAMOUNTNAME(&_GovImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME is a free data retrieval call binding the contract method 0xc6713baf.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME is a free data retrieval call binding the contract method 0xc6713baf.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME(&_GovImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME is a free data retrieval call binding the contract method 0xc6713baf.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDDISTRIBUTIONBLOCKPRODUCERNAME(&_GovImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME is a free data retrieval call binding the contract method 0x7b2bfb01.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME is a free data retrieval call binding the contract method 0x7b2bfb01.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME(&_GovImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME is a free data retrieval call binding the contract method 0x7b2bfb01.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDDISTRIBUTIONECOSYSTEMNAME(&_GovImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONMAINTENANCENAME is a free data retrieval call binding the contract method 0x46946416.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BLOCKREWARDDISTRIBUTIONMAINTENANCENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREWARDDISTRIBUTIONMAINTENANCENAME is a free data retrieval call binding the contract method 0x46946416.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BLOCKREWARDDISTRIBUTIONMAINTENANCENAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDDISTRIBUTIONMAINTENANCENAME(&_GovImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONMAINTENANCENAME is a free data retrieval call binding the contract method 0x46946416.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BLOCKREWARDDISTRIBUTIONMAINTENANCENAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDDISTRIBUTIONMAINTENANCENAME(&_GovImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONMETHODNAME is a free data retrieval call binding the contract method 0x278bb12a.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_METHOD_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BLOCKREWARDDISTRIBUTIONMETHODNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BLOCK_REWARD_DISTRIBUTION_METHOD_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREWARDDISTRIBUTIONMETHODNAME is a free data retrieval call binding the contract method 0x278bb12a.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_METHOD_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BLOCKREWARDDISTRIBUTIONMETHODNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDDISTRIBUTIONMETHODNAME(&_GovImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONMETHODNAME is a free data retrieval call binding the contract method 0x278bb12a.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_METHOD_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BLOCKREWARDDISTRIBUTIONMETHODNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDDISTRIBUTIONMETHODNAME(&_GovImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME is a free data retrieval call binding the contract method 0x6167eb45.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME is a free data retrieval call binding the contract method 0x6167eb45.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME(&_GovImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME is a free data retrieval call binding the contract method 0x6167eb45.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDDISTRIBUTIONSTAKINGREWARDNAME(&_GovImp.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_GovImp *GovImpCaller) DENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_GovImp *GovImpSession) DENOMINATOR() (*big.Int, error) {
	return _GovImp.Contract.DENOMINATOR(&_GovImp.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_GovImp *GovImpCallerSession) DENOMINATOR() (*big.Int, error) {
	return _GovImp.Contract.DENOMINATOR(&_GovImp.CallOpts)
}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) ECOSYSTEMNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "ECOSYSTEM_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) ECOSYSTEMNAME() ([32]byte, error) {
	return _GovImp.Contract.ECOSYSTEMNAME(&_GovImp.CallOpts)
}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) ECOSYSTEMNAME() ([32]byte, error) {
	return _GovImp.Contract.ECOSYSTEMNAME(&_GovImp.CallOpts)
}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) ENVSTORAGENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "ENV_STORAGE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) ENVSTORAGENAME() ([32]byte, error) {
	return _GovImp.Contract.ENVSTORAGENAME(&_GovImp.CallOpts)
}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) ENVSTORAGENAME() ([32]byte, error) {
	return _GovImp.Contract.ENVSTORAGENAME(&_GovImp.CallOpts)
}

// GASLIMITANDBASEFEENAME is a free data retrieval call binding the contract method 0xc7d3da34.
//
// Solidity: function GASLIMIT_AND_BASE_FEE_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) GASLIMITANDBASEFEENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "GASLIMIT_AND_BASE_FEE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GASLIMITANDBASEFEENAME is a free data retrieval call binding the contract method 0xc7d3da34.
//
// Solidity: function GASLIMIT_AND_BASE_FEE_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) GASLIMITANDBASEFEENAME() ([32]byte, error) {
	return _GovImp.Contract.GASLIMITANDBASEFEENAME(&_GovImp.CallOpts)
}

// GASLIMITANDBASEFEENAME is a free data retrieval call binding the contract method 0xc7d3da34.
//
// Solidity: function GASLIMIT_AND_BASE_FEE_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) GASLIMITANDBASEFEENAME() ([32]byte, error) {
	return _GovImp.Contract.GASLIMITANDBASEFEENAME(&_GovImp.CallOpts)
}

// GASTARGETPERCENTAGENAME is a free data retrieval call binding the contract method 0x4d273e28.
//
// Solidity: function GAS_TARGET_PERCENTAGE_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) GASTARGETPERCENTAGENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "GAS_TARGET_PERCENTAGE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GASTARGETPERCENTAGENAME is a free data retrieval call binding the contract method 0x4d273e28.
//
// Solidity: function GAS_TARGET_PERCENTAGE_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) GASTARGETPERCENTAGENAME() ([32]byte, error) {
	return _GovImp.Contract.GASTARGETPERCENTAGENAME(&_GovImp.CallOpts)
}

// GASTARGETPERCENTAGENAME is a free data retrieval call binding the contract method 0x4d273e28.
//
// Solidity: function GAS_TARGET_PERCENTAGE_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) GASTARGETPERCENTAGENAME() ([32]byte, error) {
	return _GovImp.Contract.GASTARGETPERCENTAGENAME(&_GovImp.CallOpts)
}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) GOVNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "GOV_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) GOVNAME() ([32]byte, error) {
	return _GovImp.Contract.GOVNAME(&_GovImp.CallOpts)
}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) GOVNAME() ([32]byte, error) {
	return _GovImp.Contract.GOVNAME(&_GovImp.CallOpts)
}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) MAINTENANCENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "MAINTENANCE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) MAINTENANCENAME() ([32]byte, error) {
	return _GovImp.Contract.MAINTENANCENAME(&_GovImp.CallOpts)
}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) MAINTENANCENAME() ([32]byte, error) {
	return _GovImp.Contract.MAINTENANCENAME(&_GovImp.CallOpts)
}

// MAXBASEFEENAME is a free data retrieval call binding the contract method 0xf38ecf47.
//
// Solidity: function MAX_BASE_FEE_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) MAXBASEFEENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "MAX_BASE_FEE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAXBASEFEENAME is a free data retrieval call binding the contract method 0xf38ecf47.
//
// Solidity: function MAX_BASE_FEE_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) MAXBASEFEENAME() ([32]byte, error) {
	return _GovImp.Contract.MAXBASEFEENAME(&_GovImp.CallOpts)
}

// MAXBASEFEENAME is a free data retrieval call binding the contract method 0xf38ecf47.
//
// Solidity: function MAX_BASE_FEE_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) MAXBASEFEENAME() ([32]byte, error) {
	return _GovImp.Contract.MAXBASEFEENAME(&_GovImp.CallOpts)
}

// MAXIDLEBLOCKINTERVALNAME is a free data retrieval call binding the contract method 0xb128f880.
//
// Solidity: function MAX_IDLE_BLOCK_INTERVAL_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) MAXIDLEBLOCKINTERVALNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "MAX_IDLE_BLOCK_INTERVAL_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAXIDLEBLOCKINTERVALNAME is a free data retrieval call binding the contract method 0xb128f880.
//
// Solidity: function MAX_IDLE_BLOCK_INTERVAL_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) MAXIDLEBLOCKINTERVALNAME() ([32]byte, error) {
	return _GovImp.Contract.MAXIDLEBLOCKINTERVALNAME(&_GovImp.CallOpts)
}

// MAXIDLEBLOCKINTERVALNAME is a free data retrieval call binding the contract method 0xb128f880.
//
// Solidity: function MAX_IDLE_BLOCK_INTERVAL_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) MAXIDLEBLOCKINTERVALNAME() ([32]byte, error) {
	return _GovImp.Contract.MAXIDLEBLOCKINTERVALNAME(&_GovImp.CallOpts)
}

// MAXPRIORITYFEEPERGASNAME is a free data retrieval call binding the contract method 0x38294419.
//
// Solidity: function MAX_PRIORITY_FEE_PER_GAS_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) MAXPRIORITYFEEPERGASNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "MAX_PRIORITY_FEE_PER_GAS_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAXPRIORITYFEEPERGASNAME is a free data retrieval call binding the contract method 0x38294419.
//
// Solidity: function MAX_PRIORITY_FEE_PER_GAS_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) MAXPRIORITYFEEPERGASNAME() ([32]byte, error) {
	return _GovImp.Contract.MAXPRIORITYFEEPERGASNAME(&_GovImp.CallOpts)
}

// MAXPRIORITYFEEPERGASNAME is a free data retrieval call binding the contract method 0x38294419.
//
// Solidity: function MAX_PRIORITY_FEE_PER_GAS_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) MAXPRIORITYFEEPERGASNAME() ([32]byte, error) {
	return _GovImp.Contract.MAXPRIORITYFEEPERGASNAME(&_GovImp.CallOpts)
}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) REWARDPOOLNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "REWARD_POOL_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) REWARDPOOLNAME() ([32]byte, error) {
	return _GovImp.Contract.REWARDPOOLNAME(&_GovImp.CallOpts)
}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) REWARDPOOLNAME() ([32]byte, error) {
	return _GovImp.Contract.REWARDPOOLNAME(&_GovImp.CallOpts)
}

// STAKINGMAXNAME is a free data retrieval call binding the contract method 0xc00ace6c.
//
// Solidity: function STAKING_MAX_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) STAKINGMAXNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "STAKING_MAX_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMAXNAME is a free data retrieval call binding the contract method 0xc00ace6c.
//
// Solidity: function STAKING_MAX_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) STAKINGMAXNAME() ([32]byte, error) {
	return _GovImp.Contract.STAKINGMAXNAME(&_GovImp.CallOpts)
}

// STAKINGMAXNAME is a free data retrieval call binding the contract method 0xc00ace6c.
//
// Solidity: function STAKING_MAX_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) STAKINGMAXNAME() ([32]byte, error) {
	return _GovImp.Contract.STAKINGMAXNAME(&_GovImp.CallOpts)
}

// STAKINGMINMAXNAME is a free data retrieval call binding the contract method 0xa6868b7d.
//
// Solidity: function STAKING_MIN_MAX_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) STAKINGMINMAXNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "STAKING_MIN_MAX_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMINMAXNAME is a free data retrieval call binding the contract method 0xa6868b7d.
//
// Solidity: function STAKING_MIN_MAX_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) STAKINGMINMAXNAME() ([32]byte, error) {
	return _GovImp.Contract.STAKINGMINMAXNAME(&_GovImp.CallOpts)
}

// STAKINGMINMAXNAME is a free data retrieval call binding the contract method 0xa6868b7d.
//
// Solidity: function STAKING_MIN_MAX_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) STAKINGMINMAXNAME() ([32]byte, error) {
	return _GovImp.Contract.STAKINGMINMAXNAME(&_GovImp.CallOpts)
}

// STAKINGMINNAME is a free data retrieval call binding the contract method 0x6fde207a.
//
// Solidity: function STAKING_MIN_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) STAKINGMINNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "STAKING_MIN_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMINNAME is a free data retrieval call binding the contract method 0x6fde207a.
//
// Solidity: function STAKING_MIN_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) STAKINGMINNAME() ([32]byte, error) {
	return _GovImp.Contract.STAKINGMINNAME(&_GovImp.CallOpts)
}

// STAKINGMINNAME is a free data retrieval call binding the contract method 0x6fde207a.
//
// Solidity: function STAKING_MIN_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) STAKINGMINNAME() ([32]byte, error) {
	return _GovImp.Contract.STAKINGMINNAME(&_GovImp.CallOpts)
}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) STAKINGNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "STAKING_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) STAKINGNAME() ([32]byte, error) {
	return _GovImp.Contract.STAKINGNAME(&_GovImp.CallOpts)
}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) STAKINGNAME() ([32]byte, error) {
	return _GovImp.Contract.STAKINGNAME(&_GovImp.CallOpts)
}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) STAKINGREWARDNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "STAKING_REWARD_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) STAKINGREWARDNAME() ([32]byte, error) {
	return _GovImp.Contract.STAKINGREWARDNAME(&_GovImp.CallOpts)
}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) STAKINGREWARDNAME() ([32]byte, error) {
	return _GovImp.Contract.STAKINGREWARDNAME(&_GovImp.CallOpts)
}

// BallotLength is a free data retrieval call binding the contract method 0xd6f9cfce.
//
// Solidity: function ballotLength() view returns(uint256)
func (_GovImp *GovImpCaller) BallotLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "ballotLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BallotLength is a free data retrieval call binding the contract method 0xd6f9cfce.
//
// Solidity: function ballotLength() view returns(uint256)
func (_GovImp *GovImpSession) BallotLength() (*big.Int, error) {
	return _GovImp.Contract.BallotLength(&_GovImp.CallOpts)
}

// BallotLength is a free data retrieval call binding the contract method 0xd6f9cfce.
//
// Solidity: function ballotLength() view returns(uint256)
func (_GovImp *GovImpCallerSession) BallotLength() (*big.Int, error) {
	return _GovImp.Contract.BallotLength(&_GovImp.CallOpts)
}

// CheckUnfinalized is a free data retrieval call binding the contract method 0x1c12b030.
//
// Solidity: function checkUnfinalized() view returns(bool)
func (_GovImp *GovImpCaller) CheckUnfinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "checkUnfinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckUnfinalized is a free data retrieval call binding the contract method 0x1c12b030.
//
// Solidity: function checkUnfinalized() view returns(bool)
func (_GovImp *GovImpSession) CheckUnfinalized() (bool, error) {
	return _GovImp.Contract.CheckUnfinalized(&_GovImp.CallOpts)
}

// CheckUnfinalized is a free data retrieval call binding the contract method 0x1c12b030.
//
// Solidity: function checkUnfinalized() view returns(bool)
func (_GovImp *GovImpCallerSession) CheckUnfinalized() (bool, error) {
	return _GovImp.Contract.CheckUnfinalized(&_GovImp.CallOpts)
}

// GetBallotInVoting is a free data retrieval call binding the contract method 0xde09b377.
//
// Solidity: function getBallotInVoting() view returns(uint256)
func (_GovImp *GovImpCaller) GetBallotInVoting(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getBallotInVoting")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBallotInVoting is a free data retrieval call binding the contract method 0xde09b377.
//
// Solidity: function getBallotInVoting() view returns(uint256)
func (_GovImp *GovImpSession) GetBallotInVoting() (*big.Int, error) {
	return _GovImp.Contract.GetBallotInVoting(&_GovImp.CallOpts)
}

// GetBallotInVoting is a free data retrieval call binding the contract method 0xde09b377.
//
// Solidity: function getBallotInVoting() view returns(uint256)
func (_GovImp *GovImpCallerSession) GetBallotInVoting() (*big.Int, error) {
	return _GovImp.Contract.GetBallotInVoting(&_GovImp.CallOpts)
}

// GetMaxStaking is a free data retrieval call binding the contract method 0x5aaa4040.
//
// Solidity: function getMaxStaking() view returns(uint256)
func (_GovImp *GovImpCaller) GetMaxStaking(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getMaxStaking")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxStaking is a free data retrieval call binding the contract method 0x5aaa4040.
//
// Solidity: function getMaxStaking() view returns(uint256)
func (_GovImp *GovImpSession) GetMaxStaking() (*big.Int, error) {
	return _GovImp.Contract.GetMaxStaking(&_GovImp.CallOpts)
}

// GetMaxStaking is a free data retrieval call binding the contract method 0x5aaa4040.
//
// Solidity: function getMaxStaking() view returns(uint256)
func (_GovImp *GovImpCallerSession) GetMaxStaking() (*big.Int, error) {
	return _GovImp.Contract.GetMaxStaking(&_GovImp.CallOpts)
}

// GetMaxVotingDuration is a free data retrieval call binding the contract method 0xce04b9d4.
//
// Solidity: function getMaxVotingDuration() view returns(uint256)
func (_GovImp *GovImpCaller) GetMaxVotingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getMaxVotingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxVotingDuration is a free data retrieval call binding the contract method 0xce04b9d4.
//
// Solidity: function getMaxVotingDuration() view returns(uint256)
func (_GovImp *GovImpSession) GetMaxVotingDuration() (*big.Int, error) {
	return _GovImp.Contract.GetMaxVotingDuration(&_GovImp.CallOpts)
}

// GetMaxVotingDuration is a free data retrieval call binding the contract method 0xce04b9d4.
//
// Solidity: function getMaxVotingDuration() view returns(uint256)
func (_GovImp *GovImpCallerSession) GetMaxVotingDuration() (*big.Int, error) {
	return _GovImp.Contract.GetMaxVotingDuration(&_GovImp.CallOpts)
}

// GetMember is a free data retrieval call binding the contract method 0xab3545e5.
//
// Solidity: function getMember(uint256 idx) view returns(address)
func (_GovImp *GovImpCaller) GetMember(opts *bind.CallOpts, idx *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getMember", idx)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMember is a free data retrieval call binding the contract method 0xab3545e5.
//
// Solidity: function getMember(uint256 idx) view returns(address)
func (_GovImp *GovImpSession) GetMember(idx *big.Int) (common.Address, error) {
	return _GovImp.Contract.GetMember(&_GovImp.CallOpts, idx)
}

// GetMember is a free data retrieval call binding the contract method 0xab3545e5.
//
// Solidity: function getMember(uint256 idx) view returns(address)
func (_GovImp *GovImpCallerSession) GetMember(idx *big.Int) (common.Address, error) {
	return _GovImp.Contract.GetMember(&_GovImp.CallOpts, idx)
}

// GetMemberFromNodeIdx is a free data retrieval call binding the contract method 0x15bf6b4d.
//
// Solidity: function getMemberFromNodeIdx(uint256 idx) view returns(address)
func (_GovImp *GovImpCaller) GetMemberFromNodeIdx(opts *bind.CallOpts, idx *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getMemberFromNodeIdx", idx)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMemberFromNodeIdx is a free data retrieval call binding the contract method 0x15bf6b4d.
//
// Solidity: function getMemberFromNodeIdx(uint256 idx) view returns(address)
func (_GovImp *GovImpSession) GetMemberFromNodeIdx(idx *big.Int) (common.Address, error) {
	return _GovImp.Contract.GetMemberFromNodeIdx(&_GovImp.CallOpts, idx)
}

// GetMemberFromNodeIdx is a free data retrieval call binding the contract method 0x15bf6b4d.
//
// Solidity: function getMemberFromNodeIdx(uint256 idx) view returns(address)
func (_GovImp *GovImpCallerSession) GetMemberFromNodeIdx(idx *big.Int) (common.Address, error) {
	return _GovImp.Contract.GetMemberFromNodeIdx(&_GovImp.CallOpts, idx)
}

// GetMemberLength is a free data retrieval call binding the contract method 0xd965ea00.
//
// Solidity: function getMemberLength() view returns(uint256)
func (_GovImp *GovImpCaller) GetMemberLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getMemberLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMemberLength is a free data retrieval call binding the contract method 0xd965ea00.
//
// Solidity: function getMemberLength() view returns(uint256)
func (_GovImp *GovImpSession) GetMemberLength() (*big.Int, error) {
	return _GovImp.Contract.GetMemberLength(&_GovImp.CallOpts)
}

// GetMemberLength is a free data retrieval call binding the contract method 0xd965ea00.
//
// Solidity: function getMemberLength() view returns(uint256)
func (_GovImp *GovImpCallerSession) GetMemberLength() (*big.Int, error) {
	return _GovImp.Contract.GetMemberLength(&_GovImp.CallOpts)
}

// GetMinStaking is a free data retrieval call binding the contract method 0xaf6af2ff.
//
// Solidity: function getMinStaking() view returns(uint256)
func (_GovImp *GovImpCaller) GetMinStaking(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getMinStaking")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinStaking is a free data retrieval call binding the contract method 0xaf6af2ff.
//
// Solidity: function getMinStaking() view returns(uint256)
func (_GovImp *GovImpSession) GetMinStaking() (*big.Int, error) {
	return _GovImp.Contract.GetMinStaking(&_GovImp.CallOpts)
}

// GetMinStaking is a free data retrieval call binding the contract method 0xaf6af2ff.
//
// Solidity: function getMinStaking() view returns(uint256)
func (_GovImp *GovImpCallerSession) GetMinStaking() (*big.Int, error) {
	return _GovImp.Contract.GetMinStaking(&_GovImp.CallOpts)
}

// GetMinVotingDuration is a free data retrieval call binding the contract method 0x1c150171.
//
// Solidity: function getMinVotingDuration() view returns(uint256)
func (_GovImp *GovImpCaller) GetMinVotingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getMinVotingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinVotingDuration is a free data retrieval call binding the contract method 0x1c150171.
//
// Solidity: function getMinVotingDuration() view returns(uint256)
func (_GovImp *GovImpSession) GetMinVotingDuration() (*big.Int, error) {
	return _GovImp.Contract.GetMinVotingDuration(&_GovImp.CallOpts)
}

// GetMinVotingDuration is a free data retrieval call binding the contract method 0x1c150171.
//
// Solidity: function getMinVotingDuration() view returns(uint256)
func (_GovImp *GovImpCallerSession) GetMinVotingDuration() (*big.Int, error) {
	return _GovImp.Contract.GetMinVotingDuration(&_GovImp.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 idx) view returns(bytes name, bytes enode, bytes ip, uint256 port)
func (_GovImp *GovImpCaller) GetNode(opts *bind.CallOpts, idx *big.Int) (struct {
	Name  []byte
	Enode []byte
	Ip    []byte
	Port  *big.Int
}, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getNode", idx)

	outstruct := new(struct {
		Name  []byte
		Enode []byte
		Ip    []byte
		Port  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Enode = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Ip = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Port = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 idx) view returns(bytes name, bytes enode, bytes ip, uint256 port)
func (_GovImp *GovImpSession) GetNode(idx *big.Int) (struct {
	Name  []byte
	Enode []byte
	Ip    []byte
	Port  *big.Int
}, error) {
	return _GovImp.Contract.GetNode(&_GovImp.CallOpts, idx)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 idx) view returns(bytes name, bytes enode, bytes ip, uint256 port)
func (_GovImp *GovImpCallerSession) GetNode(idx *big.Int) (struct {
	Name  []byte
	Enode []byte
	Ip    []byte
	Port  *big.Int
}, error) {
	return _GovImp.Contract.GetNode(&_GovImp.CallOpts, idx)
}

// GetNodeIdxFromMember is a free data retrieval call binding the contract method 0xce6a54ff.
//
// Solidity: function getNodeIdxFromMember(address addr) view returns(uint256)
func (_GovImp *GovImpCaller) GetNodeIdxFromMember(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getNodeIdxFromMember", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeIdxFromMember is a free data retrieval call binding the contract method 0xce6a54ff.
//
// Solidity: function getNodeIdxFromMember(address addr) view returns(uint256)
func (_GovImp *GovImpSession) GetNodeIdxFromMember(addr common.Address) (*big.Int, error) {
	return _GovImp.Contract.GetNodeIdxFromMember(&_GovImp.CallOpts, addr)
}

// GetNodeIdxFromMember is a free data retrieval call binding the contract method 0xce6a54ff.
//
// Solidity: function getNodeIdxFromMember(address addr) view returns(uint256)
func (_GovImp *GovImpCallerSession) GetNodeIdxFromMember(addr common.Address) (*big.Int, error) {
	return _GovImp.Contract.GetNodeIdxFromMember(&_GovImp.CallOpts, addr)
}

// GetNodeLength is a free data retrieval call binding the contract method 0x72016f75.
//
// Solidity: function getNodeLength() view returns(uint256)
func (_GovImp *GovImpCaller) GetNodeLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getNodeLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeLength is a free data retrieval call binding the contract method 0x72016f75.
//
// Solidity: function getNodeLength() view returns(uint256)
func (_GovImp *GovImpSession) GetNodeLength() (*big.Int, error) {
	return _GovImp.Contract.GetNodeLength(&_GovImp.CallOpts)
}

// GetNodeLength is a free data retrieval call binding the contract method 0x72016f75.
//
// Solidity: function getNodeLength() view returns(uint256)
func (_GovImp *GovImpCallerSession) GetNodeLength() (*big.Int, error) {
	return _GovImp.Contract.GetNodeLength(&_GovImp.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 idx) view returns(address)
func (_GovImp *GovImpCaller) GetReward(opts *bind.CallOpts, idx *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getReward", idx)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 idx) view returns(address)
func (_GovImp *GovImpSession) GetReward(idx *big.Int) (common.Address, error) {
	return _GovImp.Contract.GetReward(&_GovImp.CallOpts, idx)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 idx) view returns(address)
func (_GovImp *GovImpCallerSession) GetReward(idx *big.Int) (common.Address, error) {
	return _GovImp.Contract.GetReward(&_GovImp.CallOpts, idx)
}

// GetStakerAddr is a free data retrieval call binding the contract method 0x6f6de96d.
//
// Solidity: function getStakerAddr(address _addr) view returns(address staker)
func (_GovImp *GovImpCaller) GetStakerAddr(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getStakerAddr", _addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakerAddr is a free data retrieval call binding the contract method 0x6f6de96d.
//
// Solidity: function getStakerAddr(address _addr) view returns(address staker)
func (_GovImp *GovImpSession) GetStakerAddr(_addr common.Address) (common.Address, error) {
	return _GovImp.Contract.GetStakerAddr(&_GovImp.CallOpts, _addr)
}

// GetStakerAddr is a free data retrieval call binding the contract method 0x6f6de96d.
//
// Solidity: function getStakerAddr(address _addr) view returns(address staker)
func (_GovImp *GovImpCallerSession) GetStakerAddr(_addr common.Address) (common.Address, error) {
	return _GovImp.Contract.GetStakerAddr(&_GovImp.CallOpts, _addr)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() pure returns(uint256)
func (_GovImp *GovImpCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() pure returns(uint256)
func (_GovImp *GovImpSession) GetThreshold() (*big.Int, error) {
	return _GovImp.Contract.GetThreshold(&_GovImp.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() pure returns(uint256)
func (_GovImp *GovImpCallerSession) GetThreshold() (*big.Int, error) {
	return _GovImp.Contract.GetThreshold(&_GovImp.CallOpts)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 idx) view returns(address)
func (_GovImp *GovImpCaller) GetVoter(opts *bind.CallOpts, idx *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "getVoter", idx)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 idx) view returns(address)
func (_GovImp *GovImpSession) GetVoter(idx *big.Int) (common.Address, error) {
	return _GovImp.Contract.GetVoter(&_GovImp.CallOpts, idx)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 idx) view returns(address)
func (_GovImp *GovImpCallerSession) GetVoter(idx *big.Int) (common.Address, error) {
	return _GovImp.Contract.GetVoter(&_GovImp.CallOpts, idx)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address addr) view returns(bool)
func (_GovImp *GovImpCaller) IsMember(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "isMember", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address addr) view returns(bool)
func (_GovImp *GovImpSession) IsMember(addr common.Address) (bool, error) {
	return _GovImp.Contract.IsMember(&_GovImp.CallOpts, addr)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address addr) view returns(bool)
func (_GovImp *GovImpCallerSession) IsMember(addr common.Address) (bool, error) {
	return _GovImp.Contract.IsMember(&_GovImp.CallOpts, addr)
}

// IsReward is a free data retrieval call binding the contract method 0x4d5ce038.
//
// Solidity: function isReward(address addr) view returns(bool)
func (_GovImp *GovImpCaller) IsReward(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "isReward", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReward is a free data retrieval call binding the contract method 0x4d5ce038.
//
// Solidity: function isReward(address addr) view returns(bool)
func (_GovImp *GovImpSession) IsReward(addr common.Address) (bool, error) {
	return _GovImp.Contract.IsReward(&_GovImp.CallOpts, addr)
}

// IsReward is a free data retrieval call binding the contract method 0x4d5ce038.
//
// Solidity: function isReward(address addr) view returns(bool)
func (_GovImp *GovImpCallerSession) IsReward(addr common.Address) (bool, error) {
	return _GovImp.Contract.IsReward(&_GovImp.CallOpts, addr)
}

// IsStaker is a free data retrieval call binding the contract method 0x6f1e8533.
//
// Solidity: function isStaker(address addr) view returns(bool)
func (_GovImp *GovImpCaller) IsStaker(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "isStaker", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaker is a free data retrieval call binding the contract method 0x6f1e8533.
//
// Solidity: function isStaker(address addr) view returns(bool)
func (_GovImp *GovImpSession) IsStaker(addr common.Address) (bool, error) {
	return _GovImp.Contract.IsStaker(&_GovImp.CallOpts, addr)
}

// IsStaker is a free data retrieval call binding the contract method 0x6f1e8533.
//
// Solidity: function isStaker(address addr) view returns(bool)
func (_GovImp *GovImpCallerSession) IsStaker(addr common.Address) (bool, error) {
	return _GovImp.Contract.IsStaker(&_GovImp.CallOpts, addr)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address addr) view returns(bool)
func (_GovImp *GovImpCaller) IsVoter(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "isVoter", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address addr) view returns(bool)
func (_GovImp *GovImpSession) IsVoter(addr common.Address) (bool, error) {
	return _GovImp.Contract.IsVoter(&_GovImp.CallOpts, addr)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address addr) view returns(bool)
func (_GovImp *GovImpCallerSession) IsVoter(addr common.Address) (bool, error) {
	return _GovImp.Contract.IsVoter(&_GovImp.CallOpts, addr)
}

// LastAddProposalTime is a free data retrieval call binding the contract method 0x139d9dd3.
//
// Solidity: function lastAddProposalTime(address ) view returns(uint256)
func (_GovImp *GovImpCaller) LastAddProposalTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "lastAddProposalTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastAddProposalTime is a free data retrieval call binding the contract method 0x139d9dd3.
//
// Solidity: function lastAddProposalTime(address ) view returns(uint256)
func (_GovImp *GovImpSession) LastAddProposalTime(arg0 common.Address) (*big.Int, error) {
	return _GovImp.Contract.LastAddProposalTime(&_GovImp.CallOpts, arg0)
}

// LastAddProposalTime is a free data retrieval call binding the contract method 0x139d9dd3.
//
// Solidity: function lastAddProposalTime(address ) view returns(uint256)
func (_GovImp *GovImpCallerSession) LastAddProposalTime(arg0 common.Address) (*big.Int, error) {
	return _GovImp.Contract.LastAddProposalTime(&_GovImp.CallOpts, arg0)
}

// ModifiedBlock is a free data retrieval call binding the contract method 0x7d10dd1b.
//
// Solidity: function modifiedBlock() view returns(uint256)
func (_GovImp *GovImpCaller) ModifiedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "modifiedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModifiedBlock is a free data retrieval call binding the contract method 0x7d10dd1b.
//
// Solidity: function modifiedBlock() view returns(uint256)
func (_GovImp *GovImpSession) ModifiedBlock() (*big.Int, error) {
	return _GovImp.Contract.ModifiedBlock(&_GovImp.CallOpts)
}

// ModifiedBlock is a free data retrieval call binding the contract method 0x7d10dd1b.
//
// Solidity: function modifiedBlock() view returns(uint256)
func (_GovImp *GovImpCallerSession) ModifiedBlock() (*big.Int, error) {
	return _GovImp.Contract.ModifiedBlock(&_GovImp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GovImp *GovImpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GovImp *GovImpSession) Owner() (common.Address, error) {
	return _GovImp.Contract.Owner(&_GovImp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GovImp *GovImpCallerSession) Owner() (common.Address, error) {
	return _GovImp.Contract.Owner(&_GovImp.CallOpts)
}

// ProposalTimePeriod is a free data retrieval call binding the contract method 0x3310569c.
//
// Solidity: function proposal_time_period() view returns(uint256)
func (_GovImp *GovImpCaller) ProposalTimePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "proposal_time_period")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalTimePeriod is a free data retrieval call binding the contract method 0x3310569c.
//
// Solidity: function proposal_time_period() view returns(uint256)
func (_GovImp *GovImpSession) ProposalTimePeriod() (*big.Int, error) {
	return _GovImp.Contract.ProposalTimePeriod(&_GovImp.CallOpts)
}

// ProposalTimePeriod is a free data retrieval call binding the contract method 0x3310569c.
//
// Solidity: function proposal_time_period() view returns(uint256)
func (_GovImp *GovImpCallerSession) ProposalTimePeriod() (*big.Int, error) {
	return _GovImp.Contract.ProposalTimePeriod(&_GovImp.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GovImp *GovImpCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GovImp *GovImpSession) ProxiableUUID() ([32]byte, error) {
	return _GovImp.Contract.ProxiableUUID(&_GovImp.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GovImp *GovImpCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GovImp.Contract.ProxiableUUID(&_GovImp.CallOpts)
}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_GovImp *GovImpCaller) Reg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "reg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_GovImp *GovImpSession) Reg() (common.Address, error) {
	return _GovImp.Contract.Reg(&_GovImp.CallOpts)
}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_GovImp *GovImpCallerSession) Reg() (common.Address, error) {
	return _GovImp.Contract.Reg(&_GovImp.CallOpts)
}

// RewardIdx is a free data retrieval call binding the contract method 0xaaf0dd36.
//
// Solidity: function rewardIdx(address ) view returns(uint256)
func (_GovImp *GovImpCaller) RewardIdx(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "rewardIdx", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardIdx is a free data retrieval call binding the contract method 0xaaf0dd36.
//
// Solidity: function rewardIdx(address ) view returns(uint256)
func (_GovImp *GovImpSession) RewardIdx(arg0 common.Address) (*big.Int, error) {
	return _GovImp.Contract.RewardIdx(&_GovImp.CallOpts, arg0)
}

// RewardIdx is a free data retrieval call binding the contract method 0xaaf0dd36.
//
// Solidity: function rewardIdx(address ) view returns(uint256)
func (_GovImp *GovImpCallerSession) RewardIdx(arg0 common.Address) (*big.Int, error) {
	return _GovImp.Contract.RewardIdx(&_GovImp.CallOpts, arg0)
}

// StakerIdx is a free data retrieval call binding the contract method 0xa0c12683.
//
// Solidity: function stakerIdx(address ) view returns(uint256)
func (_GovImp *GovImpCaller) StakerIdx(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "stakerIdx", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerIdx is a free data retrieval call binding the contract method 0xa0c12683.
//
// Solidity: function stakerIdx(address ) view returns(uint256)
func (_GovImp *GovImpSession) StakerIdx(arg0 common.Address) (*big.Int, error) {
	return _GovImp.Contract.StakerIdx(&_GovImp.CallOpts, arg0)
}

// StakerIdx is a free data retrieval call binding the contract method 0xa0c12683.
//
// Solidity: function stakerIdx(address ) view returns(uint256)
func (_GovImp *GovImpCallerSession) StakerIdx(arg0 common.Address) (*big.Int, error) {
	return _GovImp.Contract.StakerIdx(&_GovImp.CallOpts, arg0)
}

// VoteLength is a free data retrieval call binding the contract method 0xe9523fb5.
//
// Solidity: function voteLength() view returns(uint256)
func (_GovImp *GovImpCaller) VoteLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "voteLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteLength is a free data retrieval call binding the contract method 0xe9523fb5.
//
// Solidity: function voteLength() view returns(uint256)
func (_GovImp *GovImpSession) VoteLength() (*big.Int, error) {
	return _GovImp.Contract.VoteLength(&_GovImp.CallOpts)
}

// VoteLength is a free data retrieval call binding the contract method 0xe9523fb5.
//
// Solidity: function voteLength() view returns(uint256)
func (_GovImp *GovImpCallerSession) VoteLength() (*big.Int, error) {
	return _GovImp.Contract.VoteLength(&_GovImp.CallOpts)
}

// VoterIdx is a free data retrieval call binding the contract method 0xcec5b622.
//
// Solidity: function voterIdx(address ) view returns(uint256)
func (_GovImp *GovImpCaller) VoterIdx(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "voterIdx", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoterIdx is a free data retrieval call binding the contract method 0xcec5b622.
//
// Solidity: function voterIdx(address ) view returns(uint256)
func (_GovImp *GovImpSession) VoterIdx(arg0 common.Address) (*big.Int, error) {
	return _GovImp.Contract.VoterIdx(&_GovImp.CallOpts, arg0)
}

// VoterIdx is a free data retrieval call binding the contract method 0xcec5b622.
//
// Solidity: function voterIdx(address ) view returns(uint256)
func (_GovImp *GovImpCallerSession) VoterIdx(arg0 common.Address) (*big.Int, error) {
	return _GovImp.Contract.VoterIdx(&_GovImp.CallOpts, arg0)
}

// AddProposalToAddMember is a paid mutator transaction binding the contract method 0x36e83d83.
//
// Solidity: function addProposalToAddMember((address,address,address,bytes,bytes,bytes,uint256,uint256,bytes,uint256) info) returns(uint256 ballotIdx)
func (_GovImp *GovImpTransactor) AddProposalToAddMember(opts *bind.TransactOpts, info GovImpMemberInfo) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "addProposalToAddMember", info)
}

// AddProposalToAddMember is a paid mutator transaction binding the contract method 0x36e83d83.
//
// Solidity: function addProposalToAddMember((address,address,address,bytes,bytes,bytes,uint256,uint256,bytes,uint256) info) returns(uint256 ballotIdx)
func (_GovImp *GovImpSession) AddProposalToAddMember(info GovImpMemberInfo) (*types.Transaction, error) {
	return _GovImp.Contract.AddProposalToAddMember(&_GovImp.TransactOpts, info)
}

// AddProposalToAddMember is a paid mutator transaction binding the contract method 0x36e83d83.
//
// Solidity: function addProposalToAddMember((address,address,address,bytes,bytes,bytes,uint256,uint256,bytes,uint256) info) returns(uint256 ballotIdx)
func (_GovImp *GovImpTransactorSession) AddProposalToAddMember(info GovImpMemberInfo) (*types.Transaction, error) {
	return _GovImp.Contract.AddProposalToAddMember(&_GovImp.TransactOpts, info)
}

// AddProposalToChangeEnv is a paid mutator transaction binding the contract method 0x40690353.
//
// Solidity: function addProposalToChangeEnv(bytes32 envName, uint256 envType, bytes envVal, bytes memo, uint256 duration) returns(uint256 ballotIdx)
func (_GovImp *GovImpTransactor) AddProposalToChangeEnv(opts *bind.TransactOpts, envName [32]byte, envType *big.Int, envVal []byte, memo []byte, duration *big.Int) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "addProposalToChangeEnv", envName, envType, envVal, memo, duration)
}

// AddProposalToChangeEnv is a paid mutator transaction binding the contract method 0x40690353.
//
// Solidity: function addProposalToChangeEnv(bytes32 envName, uint256 envType, bytes envVal, bytes memo, uint256 duration) returns(uint256 ballotIdx)
func (_GovImp *GovImpSession) AddProposalToChangeEnv(envName [32]byte, envType *big.Int, envVal []byte, memo []byte, duration *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.AddProposalToChangeEnv(&_GovImp.TransactOpts, envName, envType, envVal, memo, duration)
}

// AddProposalToChangeEnv is a paid mutator transaction binding the contract method 0x40690353.
//
// Solidity: function addProposalToChangeEnv(bytes32 envName, uint256 envType, bytes envVal, bytes memo, uint256 duration) returns(uint256 ballotIdx)
func (_GovImp *GovImpTransactorSession) AddProposalToChangeEnv(envName [32]byte, envType *big.Int, envVal []byte, memo []byte, duration *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.AddProposalToChangeEnv(&_GovImp.TransactOpts, envName, envType, envVal, memo, duration)
}

// AddProposalToChangeGov is a paid mutator transaction binding the contract method 0x0efa4909.
//
// Solidity: function addProposalToChangeGov(address newGovAddr, bytes memo, uint256 duration) returns(uint256 ballotIdx)
func (_GovImp *GovImpTransactor) AddProposalToChangeGov(opts *bind.TransactOpts, newGovAddr common.Address, memo []byte, duration *big.Int) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "addProposalToChangeGov", newGovAddr, memo, duration)
}

// AddProposalToChangeGov is a paid mutator transaction binding the contract method 0x0efa4909.
//
// Solidity: function addProposalToChangeGov(address newGovAddr, bytes memo, uint256 duration) returns(uint256 ballotIdx)
func (_GovImp *GovImpSession) AddProposalToChangeGov(newGovAddr common.Address, memo []byte, duration *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.AddProposalToChangeGov(&_GovImp.TransactOpts, newGovAddr, memo, duration)
}

// AddProposalToChangeGov is a paid mutator transaction binding the contract method 0x0efa4909.
//
// Solidity: function addProposalToChangeGov(address newGovAddr, bytes memo, uint256 duration) returns(uint256 ballotIdx)
func (_GovImp *GovImpTransactorSession) AddProposalToChangeGov(newGovAddr common.Address, memo []byte, duration *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.AddProposalToChangeGov(&_GovImp.TransactOpts, newGovAddr, memo, duration)
}

// AddProposalToChangeMember is a paid mutator transaction binding the contract method 0xa78a8188.
//
// Solidity: function addProposalToChangeMember((address,address,address,bytes,bytes,bytes,uint256,uint256,bytes,uint256) newInfo, address oldStaker, uint256 unlockAmount, uint256 slashing) returns(uint256 ballotIdx)
func (_GovImp *GovImpTransactor) AddProposalToChangeMember(opts *bind.TransactOpts, newInfo GovImpMemberInfo, oldStaker common.Address, unlockAmount *big.Int, slashing *big.Int) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "addProposalToChangeMember", newInfo, oldStaker, unlockAmount, slashing)
}

// AddProposalToChangeMember is a paid mutator transaction binding the contract method 0xa78a8188.
//
// Solidity: function addProposalToChangeMember((address,address,address,bytes,bytes,bytes,uint256,uint256,bytes,uint256) newInfo, address oldStaker, uint256 unlockAmount, uint256 slashing) returns(uint256 ballotIdx)
func (_GovImp *GovImpSession) AddProposalToChangeMember(newInfo GovImpMemberInfo, oldStaker common.Address, unlockAmount *big.Int, slashing *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.AddProposalToChangeMember(&_GovImp.TransactOpts, newInfo, oldStaker, unlockAmount, slashing)
}

// AddProposalToChangeMember is a paid mutator transaction binding the contract method 0xa78a8188.
//
// Solidity: function addProposalToChangeMember((address,address,address,bytes,bytes,bytes,uint256,uint256,bytes,uint256) newInfo, address oldStaker, uint256 unlockAmount, uint256 slashing) returns(uint256 ballotIdx)
func (_GovImp *GovImpTransactorSession) AddProposalToChangeMember(newInfo GovImpMemberInfo, oldStaker common.Address, unlockAmount *big.Int, slashing *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.AddProposalToChangeMember(&_GovImp.TransactOpts, newInfo, oldStaker, unlockAmount, slashing)
}

// AddProposalToRemoveMember is a paid mutator transaction binding the contract method 0x894f5111.
//
// Solidity: function addProposalToRemoveMember(address staker, uint256 lockAmount, bytes memo, uint256 duration, uint256 unlockAmount, uint256 slashing) returns(uint256 ballotIdx)
func (_GovImp *GovImpTransactor) AddProposalToRemoveMember(opts *bind.TransactOpts, staker common.Address, lockAmount *big.Int, memo []byte, duration *big.Int, unlockAmount *big.Int, slashing *big.Int) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "addProposalToRemoveMember", staker, lockAmount, memo, duration, unlockAmount, slashing)
}

// AddProposalToRemoveMember is a paid mutator transaction binding the contract method 0x894f5111.
//
// Solidity: function addProposalToRemoveMember(address staker, uint256 lockAmount, bytes memo, uint256 duration, uint256 unlockAmount, uint256 slashing) returns(uint256 ballotIdx)
func (_GovImp *GovImpSession) AddProposalToRemoveMember(staker common.Address, lockAmount *big.Int, memo []byte, duration *big.Int, unlockAmount *big.Int, slashing *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.AddProposalToRemoveMember(&_GovImp.TransactOpts, staker, lockAmount, memo, duration, unlockAmount, slashing)
}

// AddProposalToRemoveMember is a paid mutator transaction binding the contract method 0x894f5111.
//
// Solidity: function addProposalToRemoveMember(address staker, uint256 lockAmount, bytes memo, uint256 duration, uint256 unlockAmount, uint256 slashing) returns(uint256 ballotIdx)
func (_GovImp *GovImpTransactorSession) AddProposalToRemoveMember(staker common.Address, lockAmount *big.Int, memo []byte, duration *big.Int, unlockAmount *big.Int, slashing *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.AddProposalToRemoveMember(&_GovImp.TransactOpts, staker, lockAmount, memo, duration, unlockAmount, slashing)
}

// FinalizeEndedVote is a paid mutator transaction binding the contract method 0x6ba99181.
//
// Solidity: function finalizeEndedVote() returns()
func (_GovImp *GovImpTransactor) FinalizeEndedVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "finalizeEndedVote")
}

// FinalizeEndedVote is a paid mutator transaction binding the contract method 0x6ba99181.
//
// Solidity: function finalizeEndedVote() returns()
func (_GovImp *GovImpSession) FinalizeEndedVote() (*types.Transaction, error) {
	return _GovImp.Contract.FinalizeEndedVote(&_GovImp.TransactOpts)
}

// FinalizeEndedVote is a paid mutator transaction binding the contract method 0x6ba99181.
//
// Solidity: function finalizeEndedVote() returns()
func (_GovImp *GovImpTransactorSession) FinalizeEndedVote() (*types.Transaction, error) {
	return _GovImp.Contract.FinalizeEndedVote(&_GovImp.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xa8915a3e.
//
// Solidity: function init(address registry, uint256 lockAmount, bytes name, bytes enode, bytes ip, uint256 port) returns()
func (_GovImp *GovImpTransactor) Init(opts *bind.TransactOpts, registry common.Address, lockAmount *big.Int, name []byte, enode []byte, ip []byte, port *big.Int) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "init", registry, lockAmount, name, enode, ip, port)
}

// Init is a paid mutator transaction binding the contract method 0xa8915a3e.
//
// Solidity: function init(address registry, uint256 lockAmount, bytes name, bytes enode, bytes ip, uint256 port) returns()
func (_GovImp *GovImpSession) Init(registry common.Address, lockAmount *big.Int, name []byte, enode []byte, ip []byte, port *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.Init(&_GovImp.TransactOpts, registry, lockAmount, name, enode, ip, port)
}

// Init is a paid mutator transaction binding the contract method 0xa8915a3e.
//
// Solidity: function init(address registry, uint256 lockAmount, bytes name, bytes enode, bytes ip, uint256 port) returns()
func (_GovImp *GovImpTransactorSession) Init(registry common.Address, lockAmount *big.Int, name []byte, enode []byte, ip []byte, port *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.Init(&_GovImp.TransactOpts, registry, lockAmount, name, enode, ip, port)
}

// InitMigration is a paid mutator transaction binding the contract method 0x397e38e7.
//
// Solidity: function initMigration(address registry, uint256 oldModifiedBlock, address oldOwner) returns()
func (_GovImp *GovImpTransactor) InitMigration(opts *bind.TransactOpts, registry common.Address, oldModifiedBlock *big.Int, oldOwner common.Address) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "initMigration", registry, oldModifiedBlock, oldOwner)
}

// InitMigration is a paid mutator transaction binding the contract method 0x397e38e7.
//
// Solidity: function initMigration(address registry, uint256 oldModifiedBlock, address oldOwner) returns()
func (_GovImp *GovImpSession) InitMigration(registry common.Address, oldModifiedBlock *big.Int, oldOwner common.Address) (*types.Transaction, error) {
	return _GovImp.Contract.InitMigration(&_GovImp.TransactOpts, registry, oldModifiedBlock, oldOwner)
}

// InitMigration is a paid mutator transaction binding the contract method 0x397e38e7.
//
// Solidity: function initMigration(address registry, uint256 oldModifiedBlock, address oldOwner) returns()
func (_GovImp *GovImpTransactorSession) InitMigration(registry common.Address, oldModifiedBlock *big.Int, oldOwner common.Address) (*types.Transaction, error) {
	return _GovImp.Contract.InitMigration(&_GovImp.TransactOpts, registry, oldModifiedBlock, oldOwner)
}

// InitOnce is a paid mutator transaction binding the contract method 0x351bacda.
//
// Solidity: function initOnce(address registry, uint256 lockAmount, bytes data) returns()
func (_GovImp *GovImpTransactor) InitOnce(opts *bind.TransactOpts, registry common.Address, lockAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "initOnce", registry, lockAmount, data)
}

// InitOnce is a paid mutator transaction binding the contract method 0x351bacda.
//
// Solidity: function initOnce(address registry, uint256 lockAmount, bytes data) returns()
func (_GovImp *GovImpSession) InitOnce(registry common.Address, lockAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _GovImp.Contract.InitOnce(&_GovImp.TransactOpts, registry, lockAmount, data)
}

// InitOnce is a paid mutator transaction binding the contract method 0x351bacda.
//
// Solidity: function initOnce(address registry, uint256 lockAmount, bytes data) returns()
func (_GovImp *GovImpTransactorSession) InitOnce(registry common.Address, lockAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _GovImp.Contract.InitOnce(&_GovImp.TransactOpts, registry, lockAmount, data)
}

// MigrateFromLegacy is a paid mutator transaction binding the contract method 0x8d39e33a.
//
// Solidity: function migrateFromLegacy(address oldGov) returns(int256)
func (_GovImp *GovImpTransactor) MigrateFromLegacy(opts *bind.TransactOpts, oldGov common.Address) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "migrateFromLegacy", oldGov)
}

// MigrateFromLegacy is a paid mutator transaction binding the contract method 0x8d39e33a.
//
// Solidity: function migrateFromLegacy(address oldGov) returns(int256)
func (_GovImp *GovImpSession) MigrateFromLegacy(oldGov common.Address) (*types.Transaction, error) {
	return _GovImp.Contract.MigrateFromLegacy(&_GovImp.TransactOpts, oldGov)
}

// MigrateFromLegacy is a paid mutator transaction binding the contract method 0x8d39e33a.
//
// Solidity: function migrateFromLegacy(address oldGov) returns(int256)
func (_GovImp *GovImpTransactorSession) MigrateFromLegacy(oldGov common.Address) (*types.Transaction, error) {
	return _GovImp.Contract.MigrateFromLegacy(&_GovImp.TransactOpts, oldGov)
}

// ReInit is a paid mutator transaction binding the contract method 0x16fbe831.
//
// Solidity: function reInit() returns()
func (_GovImp *GovImpTransactor) ReInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "reInit")
}

// ReInit is a paid mutator transaction binding the contract method 0x16fbe831.
//
// Solidity: function reInit() returns()
func (_GovImp *GovImpSession) ReInit() (*types.Transaction, error) {
	return _GovImp.Contract.ReInit(&_GovImp.TransactOpts)
}

// ReInit is a paid mutator transaction binding the contract method 0x16fbe831.
//
// Solidity: function reInit() returns()
func (_GovImp *GovImpTransactorSession) ReInit() (*types.Transaction, error) {
	return _GovImp.Contract.ReInit(&_GovImp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GovImp *GovImpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GovImp *GovImpSession) RenounceOwnership() (*types.Transaction, error) {
	return _GovImp.Contract.RenounceOwnership(&_GovImp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GovImp *GovImpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GovImp.Contract.RenounceOwnership(&_GovImp.TransactOpts)
}

// SetProposalTimePeriod is a paid mutator transaction binding the contract method 0xe27bdaef.
//
// Solidity: function setProposalTimePeriod(uint256 newPeriod) returns()
func (_GovImp *GovImpTransactor) SetProposalTimePeriod(opts *bind.TransactOpts, newPeriod *big.Int) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "setProposalTimePeriod", newPeriod)
}

// SetProposalTimePeriod is a paid mutator transaction binding the contract method 0xe27bdaef.
//
// Solidity: function setProposalTimePeriod(uint256 newPeriod) returns()
func (_GovImp *GovImpSession) SetProposalTimePeriod(newPeriod *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.SetProposalTimePeriod(&_GovImp.TransactOpts, newPeriod)
}

// SetProposalTimePeriod is a paid mutator transaction binding the contract method 0xe27bdaef.
//
// Solidity: function setProposalTimePeriod(uint256 newPeriod) returns()
func (_GovImp *GovImpTransactorSession) SetProposalTimePeriod(newPeriod *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.SetProposalTimePeriod(&_GovImp.TransactOpts, newPeriod)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_GovImp *GovImpTransactor) SetRegistry(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "setRegistry", _addr)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_GovImp *GovImpSession) SetRegistry(_addr common.Address) (*types.Transaction, error) {
	return _GovImp.Contract.SetRegistry(&_GovImp.TransactOpts, _addr)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_GovImp *GovImpTransactorSession) SetRegistry(_addr common.Address) (*types.Transaction, error) {
	return _GovImp.Contract.SetRegistry(&_GovImp.TransactOpts, _addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GovImp *GovImpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GovImp *GovImpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GovImp.Contract.TransferOwnership(&_GovImp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GovImp *GovImpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GovImp.Contract.TransferOwnership(&_GovImp.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address ) returns()
func (_GovImp *GovImpTransactor) UpgradeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "upgradeTo", arg0)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address ) returns()
func (_GovImp *GovImpSession) UpgradeTo(arg0 common.Address) (*types.Transaction, error) {
	return _GovImp.Contract.UpgradeTo(&_GovImp.TransactOpts, arg0)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address ) returns()
func (_GovImp *GovImpTransactorSession) UpgradeTo(arg0 common.Address) (*types.Transaction, error) {
	return _GovImp.Contract.UpgradeTo(&_GovImp.TransactOpts, arg0)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address , bytes ) payable returns()
func (_GovImp *GovImpTransactor) UpgradeToAndCall(opts *bind.TransactOpts, arg0 common.Address, arg1 []byte) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "upgradeToAndCall", arg0, arg1)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address , bytes ) payable returns()
func (_GovImp *GovImpSession) UpgradeToAndCall(arg0 common.Address, arg1 []byte) (*types.Transaction, error) {
	return _GovImp.Contract.UpgradeToAndCall(&_GovImp.TransactOpts, arg0, arg1)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address , bytes ) payable returns()
func (_GovImp *GovImpTransactorSession) UpgradeToAndCall(arg0 common.Address, arg1 []byte) (*types.Transaction, error) {
	return _GovImp.Contract.UpgradeToAndCall(&_GovImp.TransactOpts, arg0, arg1)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 ballotIdx, bool approval) returns()
func (_GovImp *GovImpTransactor) Vote(opts *bind.TransactOpts, ballotIdx *big.Int, approval bool) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "vote", ballotIdx, approval)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 ballotIdx, bool approval) returns()
func (_GovImp *GovImpSession) Vote(ballotIdx *big.Int, approval bool) (*types.Transaction, error) {
	return _GovImp.Contract.Vote(&_GovImp.TransactOpts, ballotIdx, approval)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 ballotIdx, bool approval) returns()
func (_GovImp *GovImpTransactorSession) Vote(ballotIdx *big.Int, approval bool) (*types.Transaction, error) {
	return _GovImp.Contract.Vote(&_GovImp.TransactOpts, ballotIdx, approval)
}

// GovImpAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the GovImp contract.
type GovImpAdminChangedIterator struct {
	Event *GovImpAdminChanged // Event containing the contract specifics and raw log

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
func (it *GovImpAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpAdminChanged)
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
		it.Event = new(GovImpAdminChanged)
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
func (it *GovImpAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpAdminChanged represents a AdminChanged event raised by the GovImp contract.
type GovImpAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GovImp *GovImpFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*GovImpAdminChangedIterator, error) {

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &GovImpAdminChangedIterator{contract: _GovImp.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GovImp *GovImpFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *GovImpAdminChanged) (event.Subscription, error) {

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpAdminChanged)
				if err := _GovImp.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_GovImp *GovImpFilterer) ParseAdminChanged(log types.Log) (*GovImpAdminChanged, error) {
	event := new(GovImpAdminChanged)
	if err := _GovImp.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the GovImp contract.
type GovImpBeaconUpgradedIterator struct {
	Event *GovImpBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *GovImpBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpBeaconUpgraded)
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
		it.Event = new(GovImpBeaconUpgraded)
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
func (it *GovImpBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpBeaconUpgraded represents a BeaconUpgraded event raised by the GovImp contract.
type GovImpBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GovImp *GovImpFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*GovImpBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &GovImpBeaconUpgradedIterator{contract: _GovImp.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GovImp *GovImpFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *GovImpBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpBeaconUpgraded)
				if err := _GovImp.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_GovImp *GovImpFilterer) ParseBeaconUpgraded(log types.Log) (*GovImpBeaconUpgraded, error) {
	event := new(GovImpBeaconUpgraded)
	if err := _GovImp.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpEnvChangedIterator is returned from FilterEnvChanged and is used to iterate over the raw logs and unpacked data for EnvChanged events raised by the GovImp contract.
type GovImpEnvChangedIterator struct {
	Event *GovImpEnvChanged // Event containing the contract specifics and raw log

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
func (it *GovImpEnvChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpEnvChanged)
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
		it.Event = new(GovImpEnvChanged)
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
func (it *GovImpEnvChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpEnvChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpEnvChanged represents a EnvChanged event raised by the GovImp contract.
type GovImpEnvChanged struct {
	EnvName [32]byte
	EnvType *big.Int
	EnvVal  []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEnvChanged is a free log retrieval operation binding the contract event 0x701c16c2519cdb79aaac423a84733590e3510d9552055b6ad6908f0ab12b6c29.
//
// Solidity: event EnvChanged(bytes32 envName, uint256 envType, bytes envVal)
func (_GovImp *GovImpFilterer) FilterEnvChanged(opts *bind.FilterOpts) (*GovImpEnvChangedIterator, error) {

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "EnvChanged")
	if err != nil {
		return nil, err
	}
	return &GovImpEnvChangedIterator{contract: _GovImp.contract, event: "EnvChanged", logs: logs, sub: sub}, nil
}

// WatchEnvChanged is a free log subscription operation binding the contract event 0x701c16c2519cdb79aaac423a84733590e3510d9552055b6ad6908f0ab12b6c29.
//
// Solidity: event EnvChanged(bytes32 envName, uint256 envType, bytes envVal)
func (_GovImp *GovImpFilterer) WatchEnvChanged(opts *bind.WatchOpts, sink chan<- *GovImpEnvChanged) (event.Subscription, error) {

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "EnvChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpEnvChanged)
				if err := _GovImp.contract.UnpackLog(event, "EnvChanged", log); err != nil {
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

// ParseEnvChanged is a log parse operation binding the contract event 0x701c16c2519cdb79aaac423a84733590e3510d9552055b6ad6908f0ab12b6c29.
//
// Solidity: event EnvChanged(bytes32 envName, uint256 envType, bytes envVal)
func (_GovImp *GovImpFilterer) ParseEnvChanged(log types.Log) (*GovImpEnvChanged, error) {
	event := new(GovImpEnvChanged)
	if err := _GovImp.contract.UnpackLog(event, "EnvChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpGovDataMigratedIterator is returned from FilterGovDataMigrated and is used to iterate over the raw logs and unpacked data for GovDataMigrated events raised by the GovImp contract.
type GovImpGovDataMigratedIterator struct {
	Event *GovImpGovDataMigrated // Event containing the contract specifics and raw log

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
func (it *GovImpGovDataMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpGovDataMigrated)
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
		it.Event = new(GovImpGovDataMigrated)
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
func (it *GovImpGovDataMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpGovDataMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpGovDataMigrated represents a GovDataMigrated event raised by the GovImp contract.
type GovImpGovDataMigrated struct {
	From common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGovDataMigrated is a free log retrieval operation binding the contract event 0xab2db0a6f442428b686ffa80eadcaabe7d5ee00049c6ae888a237edd3238d856.
//
// Solidity: event GovDataMigrated(address indexed from)
func (_GovImp *GovImpFilterer) FilterGovDataMigrated(opts *bind.FilterOpts, from []common.Address) (*GovImpGovDataMigratedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "GovDataMigrated", fromRule)
	if err != nil {
		return nil, err
	}
	return &GovImpGovDataMigratedIterator{contract: _GovImp.contract, event: "GovDataMigrated", logs: logs, sub: sub}, nil
}

// WatchGovDataMigrated is a free log subscription operation binding the contract event 0xab2db0a6f442428b686ffa80eadcaabe7d5ee00049c6ae888a237edd3238d856.
//
// Solidity: event GovDataMigrated(address indexed from)
func (_GovImp *GovImpFilterer) WatchGovDataMigrated(opts *bind.WatchOpts, sink chan<- *GovImpGovDataMigrated, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "GovDataMigrated", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpGovDataMigrated)
				if err := _GovImp.contract.UnpackLog(event, "GovDataMigrated", log); err != nil {
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

// ParseGovDataMigrated is a log parse operation binding the contract event 0xab2db0a6f442428b686ffa80eadcaabe7d5ee00049c6ae888a237edd3238d856.
//
// Solidity: event GovDataMigrated(address indexed from)
func (_GovImp *GovImpFilterer) ParseGovDataMigrated(log types.Log) (*GovImpGovDataMigrated, error) {
	event := new(GovImpGovDataMigrated)
	if err := _GovImp.contract.UnpackLog(event, "GovDataMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GovImp contract.
type GovImpInitializedIterator struct {
	Event *GovImpInitialized // Event containing the contract specifics and raw log

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
func (it *GovImpInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpInitialized)
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
		it.Event = new(GovImpInitialized)
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
func (it *GovImpInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpInitialized represents a Initialized event raised by the GovImp contract.
type GovImpInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GovImp *GovImpFilterer) FilterInitialized(opts *bind.FilterOpts) (*GovImpInitializedIterator, error) {

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GovImpInitializedIterator{contract: _GovImp.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GovImp *GovImpFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GovImpInitialized) (event.Subscription, error) {

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpInitialized)
				if err := _GovImp.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GovImp *GovImpFilterer) ParseInitialized(log types.Log) (*GovImpInitialized, error) {
	event := new(GovImpInitialized)
	if err := _GovImp.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpMemberAddedIterator is returned from FilterMemberAdded and is used to iterate over the raw logs and unpacked data for MemberAdded events raised by the GovImp contract.
type GovImpMemberAddedIterator struct {
	Event *GovImpMemberAdded // Event containing the contract specifics and raw log

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
func (it *GovImpMemberAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpMemberAdded)
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
		it.Event = new(GovImpMemberAdded)
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
func (it *GovImpMemberAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpMemberAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpMemberAdded represents a MemberAdded event raised by the GovImp contract.
type GovImpMemberAdded struct {
	Addr  common.Address
	Voter common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMemberAdded is a free log retrieval operation binding the contract event 0x6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba6.
//
// Solidity: event MemberAdded(address indexed addr, address indexed voter)
func (_GovImp *GovImpFilterer) FilterMemberAdded(opts *bind.FilterOpts, addr []common.Address, voter []common.Address) (*GovImpMemberAddedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "MemberAdded", addrRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &GovImpMemberAddedIterator{contract: _GovImp.contract, event: "MemberAdded", logs: logs, sub: sub}, nil
}

// WatchMemberAdded is a free log subscription operation binding the contract event 0x6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba6.
//
// Solidity: event MemberAdded(address indexed addr, address indexed voter)
func (_GovImp *GovImpFilterer) WatchMemberAdded(opts *bind.WatchOpts, sink chan<- *GovImpMemberAdded, addr []common.Address, voter []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "MemberAdded", addrRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpMemberAdded)
				if err := _GovImp.contract.UnpackLog(event, "MemberAdded", log); err != nil {
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

// ParseMemberAdded is a log parse operation binding the contract event 0x6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba6.
//
// Solidity: event MemberAdded(address indexed addr, address indexed voter)
func (_GovImp *GovImpFilterer) ParseMemberAdded(log types.Log) (*GovImpMemberAdded, error) {
	event := new(GovImpMemberAdded)
	if err := _GovImp.contract.UnpackLog(event, "MemberAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpMemberChangedIterator is returned from FilterMemberChanged and is used to iterate over the raw logs and unpacked data for MemberChanged events raised by the GovImp contract.
type GovImpMemberChangedIterator struct {
	Event *GovImpMemberChanged // Event containing the contract specifics and raw log

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
func (it *GovImpMemberChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpMemberChanged)
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
		it.Event = new(GovImpMemberChanged)
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
func (it *GovImpMemberChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpMemberChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpMemberChanged represents a MemberChanged event raised by the GovImp contract.
type GovImpMemberChanged struct {
	OldAddr  common.Address
	NewAddr  common.Address
	NewVoter common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMemberChanged is a free log retrieval operation binding the contract event 0x15f4d750630db473a85edd9d47c500527a2648cc5e676f39645e52790cf07be0.
//
// Solidity: event MemberChanged(address indexed oldAddr, address indexed newAddr, address indexed newVoter)
func (_GovImp *GovImpFilterer) FilterMemberChanged(opts *bind.FilterOpts, oldAddr []common.Address, newAddr []common.Address, newVoter []common.Address) (*GovImpMemberChangedIterator, error) {

	var oldAddrRule []interface{}
	for _, oldAddrItem := range oldAddr {
		oldAddrRule = append(oldAddrRule, oldAddrItem)
	}
	var newAddrRule []interface{}
	for _, newAddrItem := range newAddr {
		newAddrRule = append(newAddrRule, newAddrItem)
	}
	var newVoterRule []interface{}
	for _, newVoterItem := range newVoter {
		newVoterRule = append(newVoterRule, newVoterItem)
	}

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "MemberChanged", oldAddrRule, newAddrRule, newVoterRule)
	if err != nil {
		return nil, err
	}
	return &GovImpMemberChangedIterator{contract: _GovImp.contract, event: "MemberChanged", logs: logs, sub: sub}, nil
}

// WatchMemberChanged is a free log subscription operation binding the contract event 0x15f4d750630db473a85edd9d47c500527a2648cc5e676f39645e52790cf07be0.
//
// Solidity: event MemberChanged(address indexed oldAddr, address indexed newAddr, address indexed newVoter)
func (_GovImp *GovImpFilterer) WatchMemberChanged(opts *bind.WatchOpts, sink chan<- *GovImpMemberChanged, oldAddr []common.Address, newAddr []common.Address, newVoter []common.Address) (event.Subscription, error) {

	var oldAddrRule []interface{}
	for _, oldAddrItem := range oldAddr {
		oldAddrRule = append(oldAddrRule, oldAddrItem)
	}
	var newAddrRule []interface{}
	for _, newAddrItem := range newAddr {
		newAddrRule = append(newAddrRule, newAddrItem)
	}
	var newVoterRule []interface{}
	for _, newVoterItem := range newVoter {
		newVoterRule = append(newVoterRule, newVoterItem)
	}

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "MemberChanged", oldAddrRule, newAddrRule, newVoterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpMemberChanged)
				if err := _GovImp.contract.UnpackLog(event, "MemberChanged", log); err != nil {
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

// ParseMemberChanged is a log parse operation binding the contract event 0x15f4d750630db473a85edd9d47c500527a2648cc5e676f39645e52790cf07be0.
//
// Solidity: event MemberChanged(address indexed oldAddr, address indexed newAddr, address indexed newVoter)
func (_GovImp *GovImpFilterer) ParseMemberChanged(log types.Log) (*GovImpMemberChanged, error) {
	event := new(GovImpMemberChanged)
	if err := _GovImp.contract.UnpackLog(event, "MemberChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpMemberRemovedIterator is returned from FilterMemberRemoved and is used to iterate over the raw logs and unpacked data for MemberRemoved events raised by the GovImp contract.
type GovImpMemberRemovedIterator struct {
	Event *GovImpMemberRemoved // Event containing the contract specifics and raw log

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
func (it *GovImpMemberRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpMemberRemoved)
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
		it.Event = new(GovImpMemberRemoved)
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
func (it *GovImpMemberRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpMemberRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpMemberRemoved represents a MemberRemoved event raised by the GovImp contract.
type GovImpMemberRemoved struct {
	Addr  common.Address
	Voter common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMemberRemoved is a free log retrieval operation binding the contract event 0xaa91016c21c52c58ac64f23f71bbe75becc9ada603e18ee671d09ff15492d1c1.
//
// Solidity: event MemberRemoved(address indexed addr, address indexed voter)
func (_GovImp *GovImpFilterer) FilterMemberRemoved(opts *bind.FilterOpts, addr []common.Address, voter []common.Address) (*GovImpMemberRemovedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "MemberRemoved", addrRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &GovImpMemberRemovedIterator{contract: _GovImp.contract, event: "MemberRemoved", logs: logs, sub: sub}, nil
}

// WatchMemberRemoved is a free log subscription operation binding the contract event 0xaa91016c21c52c58ac64f23f71bbe75becc9ada603e18ee671d09ff15492d1c1.
//
// Solidity: event MemberRemoved(address indexed addr, address indexed voter)
func (_GovImp *GovImpFilterer) WatchMemberRemoved(opts *bind.WatchOpts, sink chan<- *GovImpMemberRemoved, addr []common.Address, voter []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "MemberRemoved", addrRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpMemberRemoved)
				if err := _GovImp.contract.UnpackLog(event, "MemberRemoved", log); err != nil {
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

// ParseMemberRemoved is a log parse operation binding the contract event 0xaa91016c21c52c58ac64f23f71bbe75becc9ada603e18ee671d09ff15492d1c1.
//
// Solidity: event MemberRemoved(address indexed addr, address indexed voter)
func (_GovImp *GovImpFilterer) ParseMemberRemoved(log types.Log) (*GovImpMemberRemoved, error) {
	event := new(GovImpMemberRemoved)
	if err := _GovImp.contract.UnpackLog(event, "MemberRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpMemberUpdatedIterator is returned from FilterMemberUpdated and is used to iterate over the raw logs and unpacked data for MemberUpdated events raised by the GovImp contract.
type GovImpMemberUpdatedIterator struct {
	Event *GovImpMemberUpdated // Event containing the contract specifics and raw log

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
func (it *GovImpMemberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpMemberUpdated)
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
		it.Event = new(GovImpMemberUpdated)
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
func (it *GovImpMemberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpMemberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpMemberUpdated represents a MemberUpdated event raised by the GovImp contract.
type GovImpMemberUpdated struct {
	Addr  common.Address
	Voter common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMemberUpdated is a free log retrieval operation binding the contract event 0x1feee1b4fcb797c62645da41c5c6edd5f91d4291de0054da625c42b823594c1f.
//
// Solidity: event MemberUpdated(address indexed addr, address indexed voter)
func (_GovImp *GovImpFilterer) FilterMemberUpdated(opts *bind.FilterOpts, addr []common.Address, voter []common.Address) (*GovImpMemberUpdatedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "MemberUpdated", addrRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &GovImpMemberUpdatedIterator{contract: _GovImp.contract, event: "MemberUpdated", logs: logs, sub: sub}, nil
}

// WatchMemberUpdated is a free log subscription operation binding the contract event 0x1feee1b4fcb797c62645da41c5c6edd5f91d4291de0054da625c42b823594c1f.
//
// Solidity: event MemberUpdated(address indexed addr, address indexed voter)
func (_GovImp *GovImpFilterer) WatchMemberUpdated(opts *bind.WatchOpts, sink chan<- *GovImpMemberUpdated, addr []common.Address, voter []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "MemberUpdated", addrRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpMemberUpdated)
				if err := _GovImp.contract.UnpackLog(event, "MemberUpdated", log); err != nil {
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

// ParseMemberUpdated is a log parse operation binding the contract event 0x1feee1b4fcb797c62645da41c5c6edd5f91d4291de0054da625c42b823594c1f.
//
// Solidity: event MemberUpdated(address indexed addr, address indexed voter)
func (_GovImp *GovImpFilterer) ParseMemberUpdated(log types.Log) (*GovImpMemberUpdated, error) {
	event := new(GovImpMemberUpdated)
	if err := _GovImp.contract.UnpackLog(event, "MemberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpNotApplicableIterator is returned from FilterNotApplicable and is used to iterate over the raw logs and unpacked data for NotApplicable events raised by the GovImp contract.
type GovImpNotApplicableIterator struct {
	Event *GovImpNotApplicable // Event containing the contract specifics and raw log

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
func (it *GovImpNotApplicableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpNotApplicable)
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
		it.Event = new(GovImpNotApplicable)
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
func (it *GovImpNotApplicableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpNotApplicableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpNotApplicable represents a NotApplicable event raised by the GovImp contract.
type GovImpNotApplicable struct {
	BallotId *big.Int
	Reason   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNotApplicable is a free log retrieval operation binding the contract event 0x85e7f4987c0698db47045ad8cea110b51138f0eecbd94915842328cf6c3dc97d.
//
// Solidity: event NotApplicable(uint256 indexed ballotId, string reason)
func (_GovImp *GovImpFilterer) FilterNotApplicable(opts *bind.FilterOpts, ballotId []*big.Int) (*GovImpNotApplicableIterator, error) {

	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "NotApplicable", ballotIdRule)
	if err != nil {
		return nil, err
	}
	return &GovImpNotApplicableIterator{contract: _GovImp.contract, event: "NotApplicable", logs: logs, sub: sub}, nil
}

// WatchNotApplicable is a free log subscription operation binding the contract event 0x85e7f4987c0698db47045ad8cea110b51138f0eecbd94915842328cf6c3dc97d.
//
// Solidity: event NotApplicable(uint256 indexed ballotId, string reason)
func (_GovImp *GovImpFilterer) WatchNotApplicable(opts *bind.WatchOpts, sink chan<- *GovImpNotApplicable, ballotId []*big.Int) (event.Subscription, error) {

	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "NotApplicable", ballotIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpNotApplicable)
				if err := _GovImp.contract.UnpackLog(event, "NotApplicable", log); err != nil {
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

// ParseNotApplicable is a log parse operation binding the contract event 0x85e7f4987c0698db47045ad8cea110b51138f0eecbd94915842328cf6c3dc97d.
//
// Solidity: event NotApplicable(uint256 indexed ballotId, string reason)
func (_GovImp *GovImpFilterer) ParseNotApplicable(log types.Log) (*GovImpNotApplicable, error) {
	event := new(GovImpNotApplicable)
	if err := _GovImp.contract.UnpackLog(event, "NotApplicable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GovImp contract.
type GovImpOwnershipTransferredIterator struct {
	Event *GovImpOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GovImpOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpOwnershipTransferred)
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
		it.Event = new(GovImpOwnershipTransferred)
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
func (it *GovImpOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpOwnershipTransferred represents a OwnershipTransferred event raised by the GovImp contract.
type GovImpOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GovImp *GovImpFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GovImpOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GovImpOwnershipTransferredIterator{contract: _GovImp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GovImp *GovImpFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovImpOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpOwnershipTransferred)
				if err := _GovImp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GovImp *GovImpFilterer) ParseOwnershipTransferred(log types.Log) (*GovImpOwnershipTransferred, error) {
	event := new(GovImpOwnershipTransferred)
	if err := _GovImp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpSetProposalTimePeriodIterator is returned from FilterSetProposalTimePeriod and is used to iterate over the raw logs and unpacked data for SetProposalTimePeriod events raised by the GovImp contract.
type GovImpSetProposalTimePeriodIterator struct {
	Event *GovImpSetProposalTimePeriod // Event containing the contract specifics and raw log

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
func (it *GovImpSetProposalTimePeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpSetProposalTimePeriod)
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
		it.Event = new(GovImpSetProposalTimePeriod)
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
func (it *GovImpSetProposalTimePeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpSetProposalTimePeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpSetProposalTimePeriod represents a SetProposalTimePeriod event raised by the GovImp contract.
type GovImpSetProposalTimePeriod struct {
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetProposalTimePeriod is a free log retrieval operation binding the contract event 0x17c6f1d1ce638844b664872f5c6eecb7d150ec0c41187d7f85826a656ee7946f.
//
// Solidity: event SetProposalTimePeriod(uint256 newPeriod)
func (_GovImp *GovImpFilterer) FilterSetProposalTimePeriod(opts *bind.FilterOpts) (*GovImpSetProposalTimePeriodIterator, error) {

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "SetProposalTimePeriod")
	if err != nil {
		return nil, err
	}
	return &GovImpSetProposalTimePeriodIterator{contract: _GovImp.contract, event: "SetProposalTimePeriod", logs: logs, sub: sub}, nil
}

// WatchSetProposalTimePeriod is a free log subscription operation binding the contract event 0x17c6f1d1ce638844b664872f5c6eecb7d150ec0c41187d7f85826a656ee7946f.
//
// Solidity: event SetProposalTimePeriod(uint256 newPeriod)
func (_GovImp *GovImpFilterer) WatchSetProposalTimePeriod(opts *bind.WatchOpts, sink chan<- *GovImpSetProposalTimePeriod) (event.Subscription, error) {

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "SetProposalTimePeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpSetProposalTimePeriod)
				if err := _GovImp.contract.UnpackLog(event, "SetProposalTimePeriod", log); err != nil {
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

// ParseSetProposalTimePeriod is a log parse operation binding the contract event 0x17c6f1d1ce638844b664872f5c6eecb7d150ec0c41187d7f85826a656ee7946f.
//
// Solidity: event SetProposalTimePeriod(uint256 newPeriod)
func (_GovImp *GovImpFilterer) ParseSetProposalTimePeriod(log types.Log) (*GovImpSetProposalTimePeriod, error) {
	event := new(GovImpSetProposalTimePeriod)
	if err := _GovImp.contract.UnpackLog(event, "SetProposalTimePeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpSetRegistryIterator is returned from FilterSetRegistry and is used to iterate over the raw logs and unpacked data for SetRegistry events raised by the GovImp contract.
type GovImpSetRegistryIterator struct {
	Event *GovImpSetRegistry // Event containing the contract specifics and raw log

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
func (it *GovImpSetRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpSetRegistry)
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
		it.Event = new(GovImpSetRegistry)
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
func (it *GovImpSetRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpSetRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpSetRegistry represents a SetRegistry event raised by the GovImp contract.
type GovImpSetRegistry struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetRegistry is a free log retrieval operation binding the contract event 0x278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd.
//
// Solidity: event SetRegistry(address indexed addr)
func (_GovImp *GovImpFilterer) FilterSetRegistry(opts *bind.FilterOpts, addr []common.Address) (*GovImpSetRegistryIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "SetRegistry", addrRule)
	if err != nil {
		return nil, err
	}
	return &GovImpSetRegistryIterator{contract: _GovImp.contract, event: "SetRegistry", logs: logs, sub: sub}, nil
}

// WatchSetRegistry is a free log subscription operation binding the contract event 0x278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd.
//
// Solidity: event SetRegistry(address indexed addr)
func (_GovImp *GovImpFilterer) WatchSetRegistry(opts *bind.WatchOpts, sink chan<- *GovImpSetRegistry, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "SetRegistry", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpSetRegistry)
				if err := _GovImp.contract.UnpackLog(event, "SetRegistry", log); err != nil {
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
func (_GovImp *GovImpFilterer) ParseSetRegistry(log types.Log) (*GovImpSetRegistry, error) {
	event := new(GovImpSetRegistry)
	if err := _GovImp.contract.UnpackLog(event, "SetRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GovImp contract.
type GovImpUpgradedIterator struct {
	Event *GovImpUpgraded // Event containing the contract specifics and raw log

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
func (it *GovImpUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpUpgraded)
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
		it.Event = new(GovImpUpgraded)
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
func (it *GovImpUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpUpgraded represents a Upgraded event raised by the GovImp contract.
type GovImpUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GovImp *GovImpFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GovImpUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GovImpUpgradedIterator{contract: _GovImp.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GovImp *GovImpFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GovImpUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpUpgraded)
				if err := _GovImp.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GovImp *GovImpFilterer) ParseUpgraded(log types.Log) (*GovImpUpgraded, error) {
	event := new(GovImpUpgraded)
	if err := _GovImp.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
