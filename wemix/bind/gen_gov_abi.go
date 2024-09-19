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
	Bin: "0x608060405234801561001057600080fd5b5060405161077238038061077283398101604081905261002f91610326565b604080516020810190915260008152819061006b60017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd61034f565b60008051602061072b8339815191521461008757610087610374565b6100938282600061009b565b505050610405565b6100a4836100d1565b6000825111806100b15750805b156100cc576100ca838361011160201b61008b1760201c565b505b505050565b6100da8161013d565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060610136838360405180606001604052806027815260200161074b602791396101fd565b9392505050565b610150816102db60201b6100b71760201c565b6101b75760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b806101dc60008051602061072b83398151915260001b6102ea60201b6100c61760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606001600160a01b0384163b6102655760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016101ae565b600080856001600160a01b03168560405161028091906103b6565b600060405180830381855af49150503d80600081146102bb576040519150601f19603f3d011682016040523d82523d6000602084013e6102c0565b606091505b5090925090506102d18282866102ed565b9695505050505050565b6001600160a01b03163b151590565b90565b606083156102fc575081610136565b82511561030c5782518084602001fd5b8160405162461bcd60e51b81526004016101ae91906103d2565b60006020828403121561033857600080fd5b81516001600160a01b038116811461013657600080fd5b60008282101561036f57634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052600160045260246000fd5b60005b838110156103a557818101518382015260200161038d565b838111156100ca5750506000910152565b600082516103c881846020870161038a565b9190910192915050565b60208152600082518060208401526103f181604085016020870161038a565b601f01601f19169190910160400192915050565b610317806104146000396000f3fe6080604052600436106100225760003560e01c80635c60da1b1461003957610031565b366100315761002f61006a565b005b61002f61006a565b34801561004557600080fd5b5061004e61007c565b6040516001600160a01b03909116815260200160405180910390f35b61007a6100756100c9565b6100fc565b565b60006100866100c9565b905090565b60606100b083836040518060600160405280602781526020016102bb60279139610120565b9392505050565b6001600160a01b03163b151590565b90565b60006100867f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b3660008037600080366000845af43d6000803e80801561011b573d6000f35b3d6000fd5b60606001600160a01b0384163b61018d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084015b60405180910390fd5b600080856001600160a01b0316856040516101a8919061026b565b600060405180830381855af49150503d80600081146101e3576040519150601f19603f3d011682016040523d82523d6000602084013e6101e8565b606091505b50915091506101f8828286610202565b9695505050505050565b606083156102115750816100b0565b8251156102215782518084602001fd5b8160405162461bcd60e51b81526004016101849190610287565b60005b8381101561025657818101518382015260200161023e565b83811115610265576000848401525b50505050565b6000825161027d81846020870161023b565b9190910192915050565b60208152600082518060208401526102a681604085016020870161023b565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212205c5f915dcf6f7660b1c747da8970454451eff69d1682eba92b7b9f4409ff356864736f6c634300080e0033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"envName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"envType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"envVal\",\"type\":\"bytes\"}],\"name\":\"EnvChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"calldatas\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"FailReturnValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"GovDataMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"MemberAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVoter\",\"type\":\"address\"}],\"name\":\"MemberChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"MemberRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"MemberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ballotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"NotApplicable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"SetProposalTimePeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_DURATION_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BALLOT_DURATION_MIN_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BALLOT_DURATION_MIN_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BALLOT_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_FEE_MAX_CHANGE_RATE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCKS_PER_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_CREATION_TIME_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_GASLIMIT_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_AMOUNT_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_METHOD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ECOSYSTEM_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ENV_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GASLIMIT_AND_BASE_FEE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAS_TARGET_PERCENTAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAINTENANCE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BASE_FEE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_IDLE_BLOCK_INTERVAL_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PRIORITY_FEE_PER_GAS_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_POOL_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MIN_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MIN_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_REWARD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"port\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"internalType\":\"structGovImp.MemberInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"addProposalToAddMember\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"envName\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"envType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"envVal\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"addProposalToChangeEnv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"addProposalToChangeGov\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"port\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"internalType\":\"structGovImp.MemberInfo\",\"name\":\"newInfo\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"oldStaker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashing\",\"type\":\"uint256\"}],\"name\":\"addProposalToChangeMember\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"addProposalToExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashing\",\"type\":\"uint256\"}],\"name\":\"addProposalToRemoveMember\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballotLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkUnfinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeEndedVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBallotInVoting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxStaking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxVotingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getMemberFromNodeIdx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMemberLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStaking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinVotingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"port\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getNodeIdxFromMember\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getStakerAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getVoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"port\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"oldModifiedBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"}],\"name\":\"initMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initOnce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isStaker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastAddProposalTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldGov\",\"type\":\"address\"}],\"name\":\"migrateFromLegacy\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modifiedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposal_time_period\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reg\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"setProposalTimePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approval\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voterIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
		"404e585d": "addProposalToExecute(address,bytes,bytes,uint256)",
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
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000e8565b600054610100900460ff1615620000935760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015620000e6576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b608051618bd9620001046000396000612c4c0152618bd96000f3fe6080604052600436106104725760003560e01c806372016f751161024a578063af6af2ff11610139578063cec5b622116100b6578063e27bdaef1161007a578063e27bdaef14610fcc578063e75235b814610fec578063e9523fb514611001578063f2fde38b14611017578063f38ecf471461103757600080fd5b8063cec5b62214610f29578063d07bff0c14610f56578063d6f9cfce14610f8c578063d965ea0014610fa2578063de09b37714610fb757600080fd5b8063c6713baf116100fd578063c6713baf14610e56578063c7d3da3414610e8a578063c9d27afe14610ebe578063ce04b9d414610ede578063ce6a54ff14610ef357600080fd5b8063af6af2ff14610d71578063b128f88014610d86578063c00ace6c14610dba578063c0b4fe1514610dee578063c42a0abc14610e2257600080fd5b8063a0c12683116101c7578063a8915a3e1161018b578063a8915a3e14610c9a578063a91ee0dc14610cba578063a9b629b214610cda578063aaf0dd3614610d0e578063ab3545e514610d3b57600080fd5b8063a0c1268314610bc1578063a230c52414610bee578063a6868b7d14610c0e578063a7771ee314610c42578063a78a818814610c7a57600080fd5b8063894f51111161020e578063894f511114610b295780638d39e33a14610b495780638da5cb5b14610b69578063918f867414610b875780639986e4b914610b9d57600080fd5b806372016f7514610a89578063738fdd1a14610a9e5780637b2bfb0114610abe5780637bf4653014610af25780637d10dd1b14610b1357600080fd5b8063404e585d116103665780635aaa4040116102e35780636d583ca7116102a75780636d583ca7146109b45780636f1e8533146109e85780636f6de96d14610a205780636fde207a14610a40578063715018a614610a7457600080fd5b80635aaa4040146108f95780636167eb451461090e578063656e3052146109425780636ba99181146109765780636c78d2cf1461098b57600080fd5b80634d5ce0381161032a5780634d5ce038146108625780634f0f4aa9146108825780634f1ef286146108b257806352d1902d146108c05780635a731cca146108d557600080fd5b8063404e585d146107a557806340690353146107b857806346946416146107d85780634bd1ed761461080c5780634d273e281461082e57600080fd5b8063278bb12a116103f45780633659cfe6116103b85780633659cfe6146106dd57806336e83d83146106fd578063382944191461071d578063397e38e7146107515780633f35c8fe1461077157600080fd5b8063278bb12a146106325780632f40992e146106665780633310569c1461068757806334125c841461069d578063351bacda146106bd57600080fd5b80631c12b0301161043b5780631c12b030146105705780631c150171146105955780631c4b774b146105aa5780631e0cba0d146105e0578063238737b6146105fe57600080fd5b806215a73b146104775780630efa4909146104be578063139d9dd3146104de57806315bf6b4d1461050b57806316fbe83114610559575b600080fd5b34801561048357600080fd5b506104ab7f0c4fbe9dc9de15dd7c0d064975ee1a2f2f9b954fa0e65d4f6cddba94884bdc3e81565b6040519081526020015b60405180910390f35b3480156104ca57600080fd5b506104ab6104d9366004617b48565b61106b565b3480156104ea57600080fd5b506104ab6104f9366004617ba0565b60ab6020526000908152604090205481565b34801561051757600080fd5b50610541610526366004617bbd565b6000908152607360205260409020546001600160a01b031690565b6040516001600160a01b0390911681526020016104b5565b34801561056557600080fd5b5061056e611405565b005b34801561057c57600080fd5b50610585611744565b60405190151581526020016104b5565b3480156105a157600080fd5b506104ab611795565b3480156105b657600080fd5b506105416105c5366004617bbd565b6000908152606a60205260409020546001600160a01b031690565b3480156105ec57600080fd5b506104ab665374616b696e6760c81b81565b34801561060a57600080fd5b506104ab7f1d36f8ce53f59e624857e1d8dc7932d19981a2ea1b8faa4eb8ff843fc3e5a27881565b34801561063e57600080fd5b506104ab7f9b2e0c7fdae148f225bae7deb92d7e7bd24bb77edb12956e8fa7204900dd8a2281565b34801561067257600080fd5b506104ab6914995dd85c99141bdbdb60b21b81565b34801561069357600080fd5b506104ab60aa5481565b3480156106a957600080fd5b506104ab6845636f73797374656d60b81b81565b3480156106c957600080fd5b5061056e6106d8366004617bd6565b611805565b3480156106e957600080fd5b5061056e6106f8366004617ba0565b611e85565b34801561070957600080fd5b506104ab610718366004617d3c565b611ebe565b34801561072957600080fd5b506104ab7fbe90e461bbdb9a95a694f7796912ea04244caf7f5b60ad7ded17e16821d3e44c81565b34801561075d57600080fd5b5061056e61076c366004617d78565b6123d6565b34801561077d57600080fd5b506104ab7f2a268972a70c8c688b62366bdfdd9bb09cf19d3e5b6e7e7bb158e671ffdcedd281565b61056e6107b3366004617dba565b6124d1565b3480156107c457600080fd5b506104ab6107d3366004617e37565b6127b1565b3480156107e457600080fd5b506104ab7fdd5a41a7fc01f5c6d30816b17f638d6531625f1e1eaa599673ab2f6079f2dd9d81565b34801561081857600080fd5b506104ab6a4d61696e74656e616e636560a81b81565b34801561083a57600080fd5b506104ab7f77884798208df1e64f70968be41ef2d3211ec53613397ca59313416813df088881565b34801561086e57600080fd5b5061058561087d366004617ba0565b612a45565b34801561088e57600080fd5b506108a261089d366004617bbd565b612a62565b6040516104b59493929190617f0d565b61056e6106f8366004617f58565b3480156108cc57600080fd5b506104ab612c3f565b3480156108e157600080fd5b506104ab6c14dd185ada5b99d4995dd85c99609a1b81565b34801561090557600080fd5b506104ab612cf2565b34801561091a57600080fd5b506104ab7f9f1de481f899d76889aa8a2b44cc7b604d42691aa93d4ba618a1a1fd439f505081565b34801561094e57600080fd5b506104ab7fe10074dceffb75f13bf0ce50145afd35182d63796823f1280ce40e01c19109e781565b34801561098257600080fd5b5061056e612d39565b34801561099757600080fd5b506104ab7111dbdd995c9b985b98d950dbdb9d1c9858dd60721b81565b3480156109c057600080fd5b506104ab7fc69fc6b7d0efc934fd5a3581c7253a7107a952526bb6dbcd814ef8d8dae1f44a81565b3480156109f457600080fd5b50610585610a03366004617ba0565b6001600160a01b03166000908152606d6020526040902054151590565b348015610a2c57600080fd5b50610541610a3b366004617ba0565b612e27565b348015610a4c57600080fd5b506104ab7f0b09c9badbbeb6c813a598ee910770a39ccda797a1940439bb6e47fc6c87548b81565b348015610a8057600080fd5b5061056e612e96565b348015610a9557600080fd5b506074546104ab565b348015610aaa57600080fd5b50606554610541906001600160a01b031681565b348015610aca57600080fd5b506104ab7f9346226931826838eedd13d9677fa33551e7c81f604b171ef3fac355458da9aa81565b348015610afe57600080fd5b506104ab69456e7653746f7261676560b01b81565b348015610b1f57600080fd5b506104ab60665481565b348015610b3557600080fd5b506104ab610b44366004617fa7565b612eaa565b348015610b5557600080fd5b506104ab610b64366004617ba0565b613208565b348015610b7557600080fd5b506033546001600160a01b0316610541565b348015610b9357600080fd5b506104ab61271081565b348015610ba957600080fd5b506104ab6c42616c6c6f7453746f7261676560981b81565b348015610bcd57600080fd5b506104ab610bdc366004617ba0565b606d6020526000908152604090205481565b348015610bfa57600080fd5b50610585610c09366004617ba0565b613a9a565b348015610c1a57600080fd5b506104ab7f6c6f69f426081752a5d3e73746599acd2a4cb145d5de4203ca1e3473b281680b81565b348015610c4e57600080fd5b50610585610c5d366004617ba0565b6001600160a01b0316600090815260686020526040902054151590565b348015610c8657600080fd5b506104ab610c9536600461801b565b613add565b348015610ca657600080fd5b5061056e610cb5366004618079565b614239565b348015610cc657600080fd5b5061056e610cd5366004617ba0565b6146a2565b348015610ce657600080fd5b506104ab7f89dd490ecaf395283ed4ff2fd9557ca767fc425dce063451a9b0da6d72f600c381565b348015610d1a57600080fd5b506104ab610d29366004617ba0565b606b6020526000908152604090205481565b348015610d4757600080fd5b50610541610d56366004617bbd565b6000908152606c60205260409020546001600160a01b031690565b348015610d7d57600080fd5b506104ab61474a565b348015610d9257600080fd5b506104ab7f829561ab7af084b7efc6600518d2df79b8d95f3f4c3a550f54f8f7ec7d2b805781565b348015610dc657600080fd5b506104ab7f18ad4415ef4a621ce1a136395c51ab6c3712bb2e24b79d526059925cea58dcb881565b348015610dfa57600080fd5b506104ab7f8086da5becff4dfac91a3105821b361078d2d4abba0ccc2401b974cf0dcf05c181565b348015610e2e57600080fd5b506104ab7fb38b2c133e931937bd95337c65c8aefa7040ed64bbed732e3e29a4944c65747381565b348015610e6257600080fd5b506104ab7fc9e15e34073efbcd0328740feaf503caac9124b55b38c73d1a97b53da80a2f6081565b348015610e9657600080fd5b506104ab7f04f7b94450bbcad85f37ea47cd1062728f884bb2040e357738f8fd53056134bc81565b348015610eca57600080fd5b5061056e610ed9366004618133565b614791565b348015610eea57600080fd5b506104ab614986565b348015610eff57600080fd5b506104ab610f0e366004617ba0565b6001600160a01b031660009081526072602052604090205490565b348015610f3557600080fd5b506104ab610f44366004617ba0565b60686020526000908152604090205481565b348015610f6257600080fd5b50610541610f71366004617bbd565b6000908152606760205260409020546001600160a01b031690565b348015610f9857600080fd5b506104ab60755481565b348015610fae57600080fd5b506069546104ab565b348015610fc357600080fd5b506077546104ab565b348015610fd857600080fd5b5061056e610fe7366004617bbd565b6149cd565b348015610ff857600080fd5b506113896104ab565b34801561100d57600080fd5b506104ab60765481565b34801561102357600080fd5b5061056e611032366004617ba0565b614a53565b34801561104357600080fd5b506104ab7f7c1150f0e1a39ff55552d52764f97e6c387e2a247e1df344369f122c4254be2f81565b6000611075614acc565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa1580156110bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110df9190618163565b6111045760405162461bcd60e51b81526004016110fb90618180565b60405180910390fd5b600061110f33612e27565b60aa546001600160a01b038216600090815260ab60205260409020549192509061113990426181bd565b10156111575760405162461bcd60e51b81526004016110fb906181d4565b600061116233612e27565b905061116c612cf2565b61117582614aec565b11158015611192575061118661474a565b61118f82614aec565b10155b6111ae5760405162461bcd60e51b81526004016110fb9061820b565b6001600160a01b0386166112045760405162461bcd60e51b815260206004820152601d60248201527f496d706c656d656e746174696f6e2063616e6e6f74206265207a65726f00000060448201526064016110fb565b600080516020618b1d833981519152546001600160a01b03166001600160a01b0316866001600160a01b0316036112755760405162461bcd60e51b815260206004820152601560248201527453616d6520636f6e7472616374206164647265737360581b60448201526064016110fb565b856001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156112cf575060408051601f3d908101601f191682019092526112cc91810190618242565b60015b6112eb5760405162461bcd60e51b81526004016110fb9061825b565b600080516020618b1d83398151915281146113185760405162461bcd60e51b81526004016110fb906182a9565b506075546113279060016182f2565b9250611331614b63565b6001600160a01b0316630a3a63fe607554600161134e91906182f2565b60046040516001600160e01b031960e085901b16815260048101929092526024820152604481018790523360648201526001600160a01b038916608482015260a4016020604051808303816000875af11580156113af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113d39190618242565b506113de8386614b7e565b5060758290556001600160a01b0316600090815260ab602052604090204290559392505050565b600054600290610100900460ff16158015611427575060005460ff8083169116105b6114435760405162461bcd60e51b81526004016110fb9061830a565b6000805461ffff191660ff83161761010017905561145f614be9565b60005b6069548110156117105760008181526071602052604080822081516080810190925280548290829061149390618358565b80601f01602080910402602001604051908101604052809291908181526020018280546114bf90618358565b801561150c5780601f106114e15761010080835404028352916020019161150c565b820191906000526020600020905b8154815290600101906020018083116114ef57829003601f168201915b5050505050815260200160018201805461152590618358565b80601f016020809104026020016040519081016040528092919081815260200182805461155190618358565b801561159e5780601f106115735761010080835404028352916020019161159e565b820191906000526020600020905b81548152906001019060200180831161158157829003601f168201915b505050505081526020016002820180546115b790618358565b80601f01602080910402602001604051908101604052809291908181526020018280546115e390618358565b80156116305780601f1061160557610100808354040283529160200191611630565b820191906000526020600020905b81548152906001019060200180831161161357829003601f168201915b5050505050815260200160038201548152505090506001606e826000015160405161165b919061838c565b90815260405160209181900382018120805460ff191693151593909317909255820151600191606f9161168d9161838c565b908152602001604051809103902060006101000a81548160ff021916908315150217905550600160706000836040015184606001516040516020016116d39291906183a8565b60408051808303601f19018152918152815160209283012083529082019290925201600020805460ff191691151591909117905550600101611462565b506000805461ff001916905560405160ff82168152600080516020618b3d833981519152906020015b60405180910390a150565b600060775460001461178f57600061175d607754614c43565b50915050600061176e607754614ccb565b50915050600119820161178c574281101561178c5760009250505090565b50505b50600190565b600061179f614d46565b6001600160a01b03166333be496e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118009190618242565b905090565b600054610100900460ff16158080156118255750600054600160ff909116105b8061183f5750303b15801561183f575060005460ff166001145b61185b5760405162461bcd60e51b81526004016110fb9061830a565b6000805460ff19166001179055801561187e576000805461ff0019166101001790555b611886614d5e565b61188e614d8d565b611897846146a2565b4360665560006118a5614dbc565b90506118af61474a565b84101580156118c55750836118c2612cf2565b10155b6118e15760405162461bcd60e51b81526004016110fb906183ca565b600080600060608060606000806000905060008060208d0191508c518261190891906182f2565b90505b80821015611e3657815199506119226020836182f2565b915080821061193057600080fd5b8151985061193f6020836182f2565b915080821061194d57600080fd5b8151975061195c6020836182f2565b915080821061196a57600080fd5b8196508651602061197b91906182f2565b61198590836182f2565b915080821061199357600080fd5b819550855160206119a491906182f2565b6119ae90836182f2565b91508082106119bc57600080fd5b819450845160206119cd91906182f2565b6119d790836182f2565b91508082106119e557600080fd5b815193506119f46020836182f2565b9150611a016001846182f2565b9250611a0c8a613a9a565b158015611a1f5750611a1d89613a9a565b155b8015611a315750611a2f88612a45565b155b611a6e5760405162461bcd60e51b815260206004820152600e60248201526d20b63932b0b23c9036b2b6b132b960911b60448201526064016110fb565b886067600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606860008b6001600160a01b03166001600160a01b031681526020019081526020016000208190555087606a600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606b60008a6001600160a01b03166001600160a01b031681526020019081526020016000208190555089606c600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606d60008c6001600160a01b03166001600160a01b0316815260200190815260200160002081905550886001600160a01b03168a6001600160a01b03167f6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba660405160405180910390a36040516325d998bb60e01b81526001600160a01b038b811660048301528f91908d16906325d998bb90602401602060405180830381865afa158015611c1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c419190618242565b1015611c865760405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e74207374616b696e6760601b60448201526064016110fb565b611c9287878787614dd1565b611cae5760405162461bcd60e51b81526004016110fb906183f7565b611cb88a8f614e80565b600083815260716020908152604090912088519091611cdb9183918b01906178f2565b508651611cf190600183019060208a01906178f2565b508551611d0790600283019060208901906178f2565b508481600301819055506001606e89604051611d23919061838c565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f90611d55908a9061838c565b908152602001604051809103902060006101000a81548160ff0219169083151502179055506001607060008888604051602001611d939291906183a8565b60405160208183030381529060405280519060200120815260200190815260200160002060006101000a81548160ff0219169083151502179055508a6073600086815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555083607260008d6001600160a01b03166001600160a01b03168152602001908152602001600020819055505061190b565b50506069819055607455505086159550611e7f945050505050576000805461ff001916905560405160018152600080516020618b3d833981519152906020015b60405180910390a15b50505050565b60405162461bcd60e51b815260206004820152600e60248201526d496e76616c69642061636365737360901b60448201526064016110fb565b6000611ec8614acc565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015611f0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f329190618163565b611f4e5760405162461bcd60e51b81526004016110fb90618180565b6000611f5933612e27565b60aa546001600160a01b038216600090815260ab602052604090205491925090611f8390426181bd565b1015611fa15760405162461bcd60e51b81526004016110fb906181d4565b6000611fac33612e27565b9050611fb6612cf2565b611fbf82614aec565b11158015611fdc5750611fd061474a565b611fd982614aec565b10155b611ff85760405162461bcd60e51b81526004016110fb9061820b565b602084015184906001600160a01b03166120445760405162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b2103b37ba32b960991b60448201526064016110fb565b60008160600151511161208d5760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964206e6f6465206e616d6560781b60448201526064016110fb565b60008160a0015151116120d45760405162461bcd60e51b815260206004820152600f60248201526e0496e76616c6964206e6f646520495608c1b60448201526064016110fb565b60008160c001511161211c5760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a59081b9bd919481c1bdc9d607a1b60448201526064016110fb565b6000816080015151116121665760405162461bcd60e51b8152602060048201526012602482015271496e76616c6964206e6f646520656e6f646560701b60448201526064016110fb565b600081610100015151116121ab5760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964206d656d6f60a01b60448201526064016110fb565b6000816101200151116121f35760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b210323ab930ba34b7b760811b60448201526064016110fb565b6121fb61474a565b8160e00151101580156122195750612211612cf2565b8160e0015111155b61225b5760405162461bcd60e51b8152602060048201526013602482015272125b9d985b1a59081b1bd8dac8105b5bdd5b9d606a1b60448201526064016110fb565b845161226690613a9a565b15801561227b5750845161227990612a45565b155b6122b85760405162461bcd60e51b815260206004820152600e60248201526d20b63932b0b23c9036b2b6b132b960911b60448201526064016110fb565b84602001516001600160a01b031685600001516001600160a01b03161480156122fa575084604001516001600160a01b031685600001516001600160a01b0316145b61233c5760405162461bcd60e51b815260206004820152601360248201527229ba30b5b2b91034b9903737ba103b37ba32b960691b60448201526064016110fb565b612358856060015186608001518760a001518860c00151614dd1565b6123745760405162461bcd60e51b81526004016110fb906183f7565b6075546123829060016182f2565b935061239384600133600089614ebf565b6123a1848660e00151614f62565b6123b084866101000151614b7e565b505060758290556001600160a01b0316600090815260ab60205260409020429055919050565b600054610100900460ff16158080156123f65750600054600160ff909116105b806124105750303b158015612410575060005460ff166001145b61242c5760405162461bcd60e51b81526004016110fb9061830a565b6000805460ff19166001179055801561244f576000805461ff0019166101001790555b612457614d5e565b61245f614d8d565b612468846146a2565b606683905561247682614a53565b60405133907fab2db0a6f442428b686ffa80eadcaabe7d5ee00049c6ae888a237edd3238d85690600090a28015611e7f576000805461ff001916905560405160018152600080516020618b3d83398151915290602001611e76565b6124d9614acc565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa15801561251f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125439190618163565b61255f5760405162461bcd60e51b81526004016110fb90618180565b600061256a33612e27565b60aa546001600160a01b038216600090815260ab60205260409020549192509061259490426181bd565b10156125b25760405162461bcd60e51b81526004016110fb906181d4565b60006125bd33612e27565b90506125c7612cf2565b6125d082614aec565b111580156125ed57506125e161474a565b6125ea82614aec565b10155b6126095760405162461bcd60e51b81526004016110fb9061820b565b6001600160a01b0386166126575760405162461bcd60e51b81526020600482015260156024820152747461726765742063616e6e6f74206265207a65726f60581b60448201526064016110fb565b3334156126fd576000816001600160a01b0316600060405160006040518083038185875af1925050503d80600081146126ac576040519150601f19603f3d011682016040523d82523d6000602084013e6126b1565b606091505b50509050806126fb5760405162461bcd60e51b815260206004820152601660248201527563726561746f72206973206e6f742070617961626c6560501b60448201526064016110fb565b505b6000607554600161270e91906182f2565b9050612718614b63565b6001600160a01b031663aea83e4082600688868d348e6040518863ffffffff1660e01b81526004016127509796959493929190618425565b600060405180830381600087803b15801561276a57600080fd5b505af115801561277e573d6000803e3d6000fd5b5050505061278c8187614b7e565b60755550506001600160a01b0316600090815260ab6020526040902042905550505050565b60006127bb614acc565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015612801573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128259190618163565b6128415760405162461bcd60e51b81526004016110fb90618180565b600061284c33612e27565b60aa546001600160a01b038216600090815260ab60205260409020549192509061287690426181bd565b10156128945760405162461bcd60e51b81526004016110fb906181d4565b600061289f33612e27565b90506128a9612cf2565b6128b282614aec565b111580156128cf57506128c361474a565b6128cc82614aec565b10155b6128eb5760405162461bcd60e51b81526004016110fb9061820b565b866001111580156128fd575060098711155b6129385760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964207479706560a01b60448201526064016110fb565b6129428887614fa0565b61297e5760405162461bcd60e51b815260206004820152600d60248201526c496e76616c69642076616c756560981b60448201526064016110fb565b60755461298c9060016182f2565b9250612996614b63565b6001600160a01b0316634a57823e84600587338d8d8d6040518863ffffffff1660e01b81526004016129ce9796959493929190618475565b6020604051808303816000875af11580156129ed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a119190618242565b50612a1c8386614b7e565b5060758290556001600160a01b0316600090815260ab6020526040902042905595945050505050565b6001600160a01b03166000908152606b6020526040902054151590565b600081815260716020526040812060038101548154606093849384939192909160018301916002840191908490612a9890618358565b80601f0160208091040260200160405190810160405280929190818152602001828054612ac490618358565b8015612b115780601f10612ae657610100808354040283529160200191612b11565b820191906000526020600020905b815481529060010190602001808311612af457829003601f168201915b50505050509350828054612b2490618358565b80601f0160208091040260200160405190810160405280929190818152602001828054612b5090618358565b8015612b9d5780601f10612b7257610100808354040283529160200191612b9d565b820191906000526020600020905b815481529060010190602001808311612b8057829003601f168201915b50505050509250818054612bb090618358565b80601f0160208091040260200160405190810160405280929190818152602001828054612bdc90618358565b8015612c295780601f10612bfe57610100808354040283529160200191612c29565b820191906000526020600020905b815481529060010190602001808311612c0c57829003601f168201915b5050505050915093509350935093509193509193565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612cdf5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016110fb565b50600080516020618b1d83398151915290565b6000612cfc614d46565b6001600160a01b031663737c59b86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117dc573d6000803e3d6000fd5b612d41614acc565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015612d87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dab9190618163565b612dc75760405162461bcd60e51b81526004016110fb90618180565b612dcf611744565b15612e125760405162461bcd60e51b8152602060048201526013602482015272159bdd1a5b99c81a5cc81b9bdd08195b991959606a1b60448201526064016110fb565b607754612e2090600461501f565b6000607755565b6001600160a01b0381166000908152606d602052604081205415612e49575090565b6001600160a01b03821660009081526068602052604090205415612e9157506001600160a01b038082166000908152606860209081526040808320548352606c909152902054165b919050565b612e9e614be9565b612ea8600061505d565b565b6000612eb4614acc565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015612efa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f1e9190618163565b612f3a5760405162461bcd60e51b81526004016110fb90618180565b6000612f4533612e27565b60aa546001600160a01b038216600090815260ab602052604090205491925090612f6f90426181bd565b1015612f8d5760405162461bcd60e51b81526004016110fb906181d4565b6000612f9833612e27565b9050612fa2612cf2565b612fab82614aec565b11158015612fc85750612fbc61474a565b612fc582614aec565b10155b612fe45760405162461bcd60e51b81526004016110fb9061820b565b6001600160a01b03891661302c5760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b60448201526064016110fb565b61303589613a9a565b61306e5760405162461bcd60e51b815260206004820152600a6024820152692737b716b6b2b6b132b960b11b60448201526064016110fb565b600161307960695490565b116130c65760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f742072656d6f7665206120736f6c65206d656d626572000000000060448201526064016110fb565b876130d08a614aec565b10156131315760405162461bcd60e51b815260206004820152602a60248201527f496e73756666696369656e742062616c616e636520746861742063616e206265604482015269103ab73637b1b5b2b21760b11b60648201526084016110fb565b60755461313f9060016182f2565b604080516101408101825260008082526020808301829052828401829052835182815280820185526060840152835182815280820185526080840152835182815290810190935260a082019290925260c081019190915260e081018a9052610100810189905261012081018890529093506131be846002338d85614ebf565b6131c8848a614f62565b6131d28489614b7e565b6131dd8487876150af565b505060758290556001600160a01b0316600090815260ab602052604090204290559695505050505050565b60008054610100900460ff16158080156132295750600054600160ff909116105b806132435750303b158015613243575060005460ff166001145b61325f5760405162461bcd60e51b81526004016110fb9061830a565b6000805460ff191660011790558015613282576000805461ff0019166101001790555b61328a614d5e565b613292614d8d565b60008390506132fc816001600160a01b031663738fdd1a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156132d8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd591906184b4565b43606681905550613368816001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613344573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103291906184b4565b60015b816001600160a01b031663d965ea006040518163ffffffff1660e01b8152600401602060405180830381865afa1580156133a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133cd9190618242565b81116138bd5760405163ab3545e560e01b8152600481018290526001600160a01b0383169063ab3545e590602401602060405180830381865afa158015613418573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061343c91906184b4565b6000828152606c6020908152604080832080546001600160a01b0319166001600160a01b039586169081179091558352606d909152908190208390555163341effc360e21b8152600481018390529083169063d07bff0c90602401602060405180830381865afa1580156134b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134d891906184b4565b600082815260676020908152604080832080546001600160a01b0319166001600160a01b03958616908117909155835260689091529081902083905551631c4b774b60e01b81526004810183905290831690631c4b774b90602401602060405180830381865afa158015613550573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061357491906184b4565b6000828152606a6020908152604080832080546001600160a01b0319166001600160a01b03959095169485179055928252606b8152828220849055606984905582516080810184526060808252918101829052928301819052820152604051634f0f4aa960e01b8152600481018390526001600160a01b03841690634f0f4aa990602401600060405180830381865afa158015613615573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261363d9190810190618521565b606085018190526040850182905260208501839052838552613662939291908561511e565b6136ae5760405162461bcd60e51b815260206004820152601760248201527f6e6f646520696e666f206973206475706c69636174656400000000000000000060448201526064016110fb565b6001606e82600001516040516136c4919061838c565b90815260405160209181900382018120805460ff191693151593909317909255820151600191606f916136f69161838c565b908152602001604051809103902060006101000a81548160ff0219169083151502179055506001607060008360400151846060015160405160200161373c9291906183a8565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff1916941515949094179093558483526071815291208251805184936137929284929101906178f2565b5060208281015180516137ab92600185019201906178f2565b50604082015180516137c79160028401916020909101906178f2565b50606091909101516003909101556000828152606c6020818152604080842080546001600160a01b039081168652607284528286208890558786528154607385529583902080546001600160a01b031916968216969096179095556074879055929091529054905163139d9dd360e01b815290821660048201529084169063139d9dd390602401602060405180830381865afa15801561386b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061388f9190618242565b6000838152606c60209081526040808320546001600160a01b0316835260ab9091529020555060010161336b565b50806001600160a01b0316633310569c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156138fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139209190618242565b60aa81905550806001600160a01b031663d6f9cfce6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613964573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139889190618242565b607581905550806001600160a01b031663e9523fb56040518163ffffffff1660e01b8152600401602060405180830381865afa1580156139cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139f09190618242565b607681905550806001600160a01b031663de09b3776040518163ffffffff1660e01b8152600401602060405180830381865afa158015613a34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a589190618242565b60775550600091508015613a94576000805461ff001916905560405160018152600080516020618b3d8339815191529060200160405180910390a15b50919050565b6001600160a01b0381166000908152606d6020526040812054151580613ad757506001600160a01b03821660009081526068602052604090205415155b92915050565b6000613ae7614acc565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015613b2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b519190618163565b613b6d5760405162461bcd60e51b81526004016110fb90618180565b6000613b7833612e27565b60aa546001600160a01b038216600090815260ab602052604090205491925090613ba290426181bd565b1015613bc05760405162461bcd60e51b81526004016110fb906181d4565b6000613bcb33612e27565b9050613bd5612cf2565b613bde82614aec565b11158015613bfb5750613bef61474a565b613bf882614aec565b10155b613c175760405162461bcd60e51b81526004016110fb9061820b565b602087015187906001600160a01b0316613c635760405162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b2103b37ba32b960991b60448201526064016110fb565b600081606001515111613cac5760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964206e6f6465206e616d6560781b60448201526064016110fb565b60008160a001515111613cf35760405162461bcd60e51b815260206004820152600f60248201526e0496e76616c6964206e6f646520495608c1b60448201526064016110fb565b60008160c0015111613d3b5760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a59081b9bd919481c1bdc9d607a1b60448201526064016110fb565b600081608001515111613d855760405162461bcd60e51b8152602060048201526012602482015271496e76616c6964206e6f646520656e6f646560701b60448201526064016110fb565b60008161010001515111613dca5760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964206d656d6f60a01b60448201526064016110fb565b600081610120015111613e125760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b210323ab930ba34b7b760811b60448201526064016110fb565b613e1a61474a565b8160e0015110158015613e385750613e30612cf2565b8160e0015111155b613e7a5760405162461bcd60e51b8152602060048201526013602482015272125b9d985b1a59081b1bd8dac8105b5bdd5b9d606a1b60448201526064016110fb565b6001600160a01b038716613ec65760405162461bcd60e51b8152602060048201526013602482015272496e76616c6964206f6c64204164647265737360681b60448201526064016110fb565b613ecf87613a9a565b613f085760405162461bcd60e51b815260206004820152600a6024820152692737b716b6b2b6b132b960b11b60448201526064016110fb565b6020808901516001600160a01b038981166000908152606d8452604080822054825260679094529290922054821691161480613f595750866001600160a01b031688602001516001600160a01b0316145b80613f835750613f6c8860200151613a9a565b158015613f835750613f818860200151612a45565b155b801561400557506040808901516001600160a01b038981166000908152606d6020908152848220548252606a90529290922054821691161480613fdb5750866001600160a01b031688604001516001600160a01b0316145b806140055750613fee8860400151613a9a565b15801561400557506140038860400151612a45565b155b6140445760405162461bcd60e51b815260206004820152601060248201526f20b63932b0b23c90309036b2b6b132b960811b60448201526064016110fb565b336001600160a01b038816148015614068575087516001600160a01b038881169116145b156140bd5785158015614079575084155b6140b85760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a59081c1c9bdc1bdcd85b60821b60448201526064016110fb565b61416c565b87516001600160a01b0388811691161461416c576140d961474a565b6140e386886182f2565b111561416c5760405162461bcd60e51b815260206004820152604c60248201527f496e76616c696420616d6f756e743a2028756e6c6f636b416d6f756e74202b2060448201527f736c617368696e6729206d75737420626520657175616c206f72206c6f77207460648201526b6f206d696e5374616b696e6760a01b608482015260a4016110fb565b60755461417a9060016182f2565b935061418a846003338a8c614ebf565b614198848960e00151614f62565b6141a784896101000151614b7e565b6141b28487876150af565b6075849055336001600160a01b0388161480156141db575087516001600160a01b038881169116145b156142155760006141eb85614ccb565b925050506142058542834261420091906182f2565b615247565b61421385600360018061528c565b505b50506001600160a01b0316600090815260ab60205260409020429055949350505050565b600054610100900460ff16158080156142595750600054600160ff909116105b806142735750303b158015614273575060005460ff166001145b61428f5760405162461bcd60e51b81526004016110fb9061830a565b6000805460ff1916600117905580156142b2576000805461ff0019166101001790555b6142ba614d5e565b6142c2614d8d565b6142cb876146a2565b6142d361474a565b86101580156142e95750856142e6612cf2565b10155b6143055760405162461bcd60e51b81526004016110fb906183ca565b600061430f614dbc565b6040516325d998bb60e01b815233600482015290915087906001600160a01b038316906325d998bb90602401602060405180830381865afa158015614358573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061437c9190618242565b10156143c15760405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e74207374616b696e6760601b60448201526064016110fb565b60405163282d3fdf60e01b8152336004820152602481018890526001600160a01b0382169063282d3fdf90604401600060405180830381600087803b15801561440957600080fd5b505af115801561441d573d6000803e3d6000fd5b5050600160698190557f6bee784efeb983674392298ab585b22866bedf00ebb0eea949d1e66f3f50e71d8054336001600160a01b0319918216811790925560008281526068602090815260408083208690557ff585789965ba69220d5ce3dc1b444eb22ff546f2650694fef8fafe9c26560af98054851686179055606b82528083208690557fdcf345d7f6a8deb7427d0fee62009fa15100353a1c666b51bb5387b25addcfa98054909416909417909255606d825291822083905560748390559190526071815288517f169c6be1b0e6ab5de76b532e587a77340130ac65c5591db02be822dcf1dc0ed69350614518925083918a01906178f2565b50855161452e90600183019060208901906178f2565b50845161454490600283019060208801906178f2565b508381600301819055506001606e88604051614560919061838c565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f9061459290899061838c565b908152602001604051809103902060006101000a81548160ff02191690831515021790555060016070600087876040516020016145d09291906183a8565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff1916941515949094179093556074543380855260728352838520829055908452607390915281832080546001600160a01b03191682179055436066559051909182917f6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba69190a350508015614699576000805461ff001916905560405160018152600080516020618b3d8339815191529060200160405180910390a15b50505050505050565b6146aa614be9565b6001600160a01b0381166147005760405162461bcd60e51b815260206004820152601a60248201527f416464726573732073686f756c64206265206e6f6e2d7a65726f00000000000060448201526064016110fb565b606580546001600160a01b0319166001600160a01b0383169081179091556040517f278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd90600090a250565b6000614754614d46565b6001600160a01b031663076cd77f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117dc573d6000803e3d6000fd5b6002607854036147e35760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016110fb565b60026078556147f0614acc565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015614836573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061485a9190618163565b6148765760405162461bcd60e51b81526004016110fb90618180565b600061488133612e27565b905061488b612cf2565b61489482614aec565b111580156148b157506148a561474a565b6148ae82614aec565b10155b6148cd5760405162461bcd60e51b81526004016110fb9061820b565b6148d5611744565b61490b5760405162461bcd60e51b8152602060048201526007602482015266115e1c1a5c995960ca1b60448201526064016110fb565b6000614916846153bf565b905061492284846154d2565b60008061492e866155ac565b9250925050600061493e61138990565b9050808310158061494f5750808210155b80614964575061495f82846182f2565b612710145b15614978576149788785848611600061528c565b505060016078555050505050565b6000614990614d46565b6001600160a01b0316631b27e01b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117dc573d6000803e3d6000fd5b6149d5614be9565b610e108110614a1e5760405162461bcd60e51b81526020600482015260156024820152746e6577506572696f6420697320746f6f206c6f6e6760581b60448201526064016110fb565b60aa8190556040518181527f17c6f1d1ce638844b664872f5c6eecb7d150ec0c41187d7f85826a656ee7946f90602001611739565b614a5b614be9565b6001600160a01b038116614ac05760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016110fb565b614ac98161505d565b50565b60006118007111dbdd995c9b985b98d950dbdb9d1c9858dd60721b6155e6565b6000614af6614dbc565b604051632c9aab9b60e11b81526001600160a01b03848116600483015291909116906359355736906024015b602060405180830381865afa158015614b3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ad79190618242565b60006118006c42616c6c6f7453746f7261676560981b6155e6565b614b86614b63565b6001600160a01b031663bce0dbc183836040518363ffffffff1660e01b8152600401614bb39291906185b0565b600060405180830381600087803b158015614bcd57600080fd5b505af1158015614be1573d6000803e3d6000fd5b505050505050565b6033546001600160a01b03163314612ea85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016110fb565b6000806000614c50614b63565b6001600160a01b031663688ca5b2856040518263ffffffff1660e01b8152600401614c7d91815260200190565b606060405180830381865afa158015614c9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614cbe91906185c9565b9250925092509193909250565b6000806000614cd8614b63565b6001600160a01b03166309970688856040518263ffffffff1660e01b8152600401614d0591815260200190565b606060405180830381865afa158015614d22573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614cbe91906185f7565b600061180069456e7653746f7261676560b01b6155e6565b600054610100900460ff16614d855760405162461bcd60e51b81526004016110fb90618625565b612ea8615654565b600054610100900460ff16614db45760405162461bcd60e51b81526004016110fb90618625565b612ea8615682565b6000611800665374616b696e6760c81b6155e6565b604051600190606f90614de590869061838c565b9081526040519081900360200190205460ff1615614e01575060005b606e85604051614e11919061838c565b9081526040519081900360200190205460ff1615614e2d575060005b60008383604051602001614e429291906183a8565b60408051601f1981840301815291815281516020928301206000818152607090935291205490915060ff1615614e7757600091505b50949350505050565b614e88614dbc565b60405163282d3fdf60e01b81526001600160a01b03848116600483015260248201849052919091169063282d3fdf90604401614bb3565b614ec7614b63565b6001600160a01b031663daacbb95868684610120015187878760000151886020015189604001518a606001518b608001518c60a001518d60c001516040518d63ffffffff1660e01b8152600401614f299c9b9a99989796959493929190618670565b600060405180830381600087803b158015614f4357600080fd5b505af1158015614f57573d6000803e3d6000fd5b505050505050505050565b614f6a614b63565b604051633968764960e11b815260048101849052602481018390526001600160a01b0391909116906372d0ec9290604401614bb3565b6000614faa614d46565b6001600160a01b0316639801bff984846040518363ffffffff1660e01b8152600401614fd79291906185b0565b602060405180830381865afa158015614ff4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906150189190618163565b9392505050565b615027614b63565b60405163548f2cdd60e11b815260048101849052602481018390526001600160a01b03919091169063a91e59ba90604401614bb3565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6150b7614b63565b604051632264085960e01b81526004810185905260248101849052604481018390526001600160a01b0391909116906322640859906064015b600060405180830381600087803b15801561510a57600080fd5b505af1158015614699573d6000803e3d6000fd5b6000600190508480519060200120826020015180519060200120141580156151655750606f85604051615151919061838c565b9081526040519081900360200190205460ff165b1561516e575060005b8580519060200120826000015180519060200120141580156151af5750606e8660405161519b919061838c565b9081526040519081900360200190205460ff165b156151b8575060005b600084846040516020016151cd9291906183a8565b60405160208183030381529060405280519060200120905080836040015184606001516040516020016152019291906183a8565b6040516020818303038152906040528051906020012014158015615233575060008181526070602052604090205460ff165b1561523d57600091505b5095945050505050565b61524f614b63565b60405163605b78c360e11b81526004810185905260248101849052604481018390526001600160a01b03919091169063c0b6f186906064016150f0565b6004821561530c5750600360001984016152b6576152a9856156b2565b6152b1575060045b6153a3565b600284036152c7576152b185615bbc565b600384036152d9576152a985836160fc565b600484036152ea576152b185616696565b600584036152fb576152b185616758565b600684036152b1576152b18561689c565b600684036153a357600061531e614b63565b60405163060e6a9960e01b8152600481018890529091506000906001600160a01b0383169063060e6a9990602401600060405180830381865afa158015615369573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615391919081019061870d565b509150506153a08288836169ef565b50505b6153ad858261501f565b816153b85760006077555b5050505050565b60008060006153cd84614c43565b5090925090506001810361546b57607754156153fb5760405162461bcd60e51b81526004016110fb9061875c565b600061540685614ccb565b92505050615412611795565b811015615435576154308542615426611795565b61420090426182f2565b615460565b8061543e614986565b1015615451576154308542615426614986565b615460854261420084826182f2565b5060778490556154cb565b600281036154995760775484146154945760405162461bcd60e51b81526004016110fb9061875c565b6154cb565b60405162461bcd60e51b8152602060048201526007602482015266115e1c1a5c995960ca1b60448201526064016110fb565b5092915050565b600060765460016154e391906182f2565b905060006154f033612e27565b905060006154fd60695490565b6155099061271061879f565b905060008461551957600261551c565b60015b9050615526614b63565b6040516325918ae760e21b815260048101869052602481018890526001600160a01b038581166044830152606482018490526084820185905291909116906396462b9c9060a401600060405180830381600087803b15801561558757600080fd5b505af115801561559b573d6000803e3d6000fd5b505050607694909455505050505050565b60008060006155b9614b63565b6001600160a01b03166356ba988e856040518263ffffffff1660e01b8152600401614d0591815260200190565b606554604051630d2020dd60e01b8152600481018390526000916001600160a01b031690630d2020dd90602401602060405180830381865afa158015615630573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ad791906184b4565b600054610100900460ff1661567b5760405162461bcd60e51b81526004016110fb90618625565b6001607855565b600054610100900460ff166156a95760405162461bcd60e51b81526004016110fb90618625565b612ea83361505d565b60006156bf826001616b21565b6000806000806000806000806156d48a616c3d565b98509850985098509850985098509850506156ee88613a9a565b1561574c5789600080516020618b5d8339815191526040516157349060208082526010908201526f20b63932b0b23c90309036b2b6b132b960811b604082015260600190565b60405180910390a25060009998505050505050505050565b61575586612a45565b1561579d5789600080516020618b5d8339815191526040516157349060208082526012908201527120b63932b0b23c9030903932bbb0b93232b960711b604082015260600190565b6157a561474a565b8110806157b85750806157b6612cf2565b105b156157d95789600080516020618b5d833981519152604051615734906183ca565b806157e389616ce4565b10156158055789600080516020618b5d833981519152604051615734906187c1565b866001600160a01b0316886001600160a01b0316141580156158395750856001600160a01b0316886001600160a01b031614155b156158855789600080516020618b5d83398151915260405161573490602080825260169082015275496e76616c6964206d656d626572206164647265737360501b604082015260600190565b61588f8882614e80565b600060695460016158a091906182f2565b9050876067600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606860008a6001600160a01b03166001600160a01b031681526020019081526020016000208190555086606a600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606b6000896001600160a01b03166001600160a01b031681526020019081526020016000208190555088606c600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606d60008b6001600160a01b03166001600160a01b0316815260200190815260200160002081905550600060745460016159d991906182f2565b60008181526071602090815260409091208951929350916159ff918391908b01906178f2565b508651615a1590600183019060208a01906178f2565b508551615a2b90600283019060208901906178f2565b508481600301819055506001606e89604051615a47919061838c565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f90615a79908a9061838c565b908152602001604051809103902060006101000a81548160ff0219169083151502179055506001607060008888604051602001615ab79291906183a8565b60405160208183030381529060405280519060200120815260200190815260200160002060006101000a81548160ff0219169083151502179055508a6073600084815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081607260008d6001600160a01b03166001600160a01b0316815260200190815260200160002081905550826069819055508160748190555043606681905550896001600160a01b03168b6001600160a01b03167f6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba660405160405180910390a35060019c9b505050505050505050505050565b615bc7816002616b21565b6000615bd282616c3d565b50505050505050509050615be581613a9a565b615c3a5781600080516020618b5d833981519152604051615c2e906020808252601490820152732737ba1030b63932b0b23c90309036b2b6b132b960611b604082015260600190565b60405180910390a25050565b6001600160a01b038082166000818152606d6020818152604080842054606954808652606c8452828620548287526067855283872054606a86529387205497909652939092529094928316939083169216908414615da257606980546000908152606c6020818152604080842080546001600160a01b038c8116808852606d80875285892080549e8a529787528589209087529b821688528488209c909c559486905581546001600160a01b03199081169092558954821690851617909855888452606b8083528185205486548652606a80855283872080548a881689528487528589208054948a529287528589209487528716808952858920939093559087905580548b16905581548a16179055978352606880825288842054945484526067808352898520805489861687528385528b872080548989529386528c8820949095529094168086529985205592905580548616905580549094168517909355919291615e1f565b606980546000908152606c6020908152604080832080546001600160a01b03199081169091556001600160a01b038a81168552606d845282852085905585548552606a84528285208054831690558681168552606b8452828520859055945484526067835281842080549091169055928516825260689052908120555b6001606954615e2e91906181bd565b6069556000848152607160205260408082209051909190606f90615e569060018501906188a1565b908152604051908190036020018120805492151560ff1990931692909217909155600090606e90615e889084906188a1565b90815260405160209181900382018120805460ff19169315159390931790925560038301546000926070928492615ec4926002880192016188ad565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff1916941515949094179093556074546001600160a01b038a16845260729091529120541461601e576001600160a01b038087166000908152607260209081526040808320546074548452607383528184205460719093529220805492985092169550829190615f5f90618358565b615f6a929190617976565b5060745460009081526071602052604090206001908101805491830191615f9090618358565b615f9b929190617976565b5060745460009081526071602052604090206002908101805491830191615fc190618358565b615fcc929190617976565b506074546000908152607160209081526040808320600390810154908501558783526073825280832080546001600160a01b0319166001600160a01b0389169081179091558352607290915290208590555b60748054600090815260736020908152604080832080546001600160a01b03191690556001600160a01b038a1683526072825280832083905592548252607190529081209061606d82826179f1565b61607b6001830160006179f1565b6160896002830160006179f1565b6003820160009055505060016074546160a291906181bd565b607455436066556160b38787616d1e565b826001600160a01b0316866001600160a01b03167faa91016c21c52c58ac64f23f71bbe75becc9ada603e18ee671d09ff15492d1c160405160405180910390a350505050505050565b60008161610e5761610e836003616b21565b60008060008060008060008060006161258c616c3d565b98509850985098509850985098509850985061614089613a9a565b6161ac578b600080516020618b5d833981519152604051616192906020808252601b908201527f4f6c642061646472657373206973206e6f742061206d656d6265720000000000604082015260600190565b60405180910390a260009950505050505050505050613ad7565b6161bf8c8c8b8b8b8b8b8b8b8b8b616eec565b6161d55760009950505050505050505050613ad7565b6001600160a01b03808a166000818152606d6020526040902054918a1614616244576000818152606c6020908152604080832080546001600160a01b0319166001600160a01b038e81169182179092558452606d909252808320849055908c1682528120556162448983614e80565b6001600160a01b038a166000908152607260209081526040808320548084526071909252808320905191929091606e9061627f9084906188a1565b908152604051908190036020018120805492151560ff1990931692909217909155600090606f906162b49060018501906188a1565b90815260405160209181900382018120805460ff191693151593909317909255600383015460009260709284926162f0926002880192016188ad565b60408051808303601f1901815291815281516020928301208352828201939093529101600020805460ff1916921515929092179091558851616337918391908b01906178f2565b50865161634d90600183019060208a01906178f2565b50855161636390600283019060208901906178f2565b506003810185905543606655604051600190606e90616383908b9061838c565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f906163b5908a9061838c565b908152602001604051809103902060006101000a81548160ff02191690831515021790555060016070600088886040516020016163f39291906183a8565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff191694151594909417909355858352606a90529020546001600160a01b0390811691508916811461648f576000838152606a6020908152604080832080546001600160a01b0319166001600160a01b038e81169182179092558452606b90925280832086905590831682528120555b506000828152606760205260409020546001600160a01b03908116908a16811461654157896067600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606860008c6001600160a01b03166001600160a01b0316815260200190815260200160002081905550600060686000836001600160a01b03166001600160a01b03168152602001908152602001600020819055505b50896001600160a01b03168b6001600160a01b03161461664157896073600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080607260008c6001600160a01b03166001600160a01b03168152602001908152602001600020819055506000607260008d6001600160a01b03166001600160a01b03168152602001908152602001600020819055506165f28e8c616d1e565b886001600160a01b03168a6001600160a01b03168c6001600160a01b03167f15f4d750630db473a85edd9d47c500527a2648cc5e676f39645e52790cf07be060405160405180910390a4616682565b896001600160a01b03168b6001600160a01b03167f1feee1b4fcb797c62645da41c5c6edd5f91d4291de0054da625c42b823594c1f60405160405180910390a35b5060019d9c50505050505050505050505050565b6166a1816004616b21565b60006166ab614b63565b6001600160a01b0316637efa9ae3836040518263ffffffff1660e01b81526004016166d891815260200190565b602060405180830381865afa1580156166f5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061671991906184b4565b90506001600160a01b038116156167545761673381617486565b6040805160008082526020820190925261674f91839190617514565b436066555b5050565b616763816005616b21565b6000806000616770614b63565b6001600160a01b0316631d940da2856040518263ffffffff1660e01b815260040161679d91815260200190565b600060405180830381865afa1580156167ba573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526167e291908101906188c6565b92509250925060006167f2614d46565b6040516388c2801960e01b81529091506001600160a01b038216906388c280199061682390879086906004016185b0565b600060405180830381600087803b15801561683d57600080fd5b505af1158015616851573d6000803e3d6000fd5b50504360665550506040517f701c16c2519cdb79aaac423a84733590e3510d9552055b6ad6908f0ab12b6c299061688d908690869086906188ff565b60405180910390a15050505050565b6168a7816006616b21565b60006168b1614b63565b90506000806000836001600160a01b031663060e6a99866040518263ffffffff1660e01b81526004016168e691815260200190565b600060405180830381865afa158015616903573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261692b919081019061870d565b925092509250600080846001600160a01b0316848460405161694d919061838c565b60006040518083038185875af1925050503d806000811461698a576040519150601f19603f3d011682016040523d82523d6000602084013e61698f565b606091505b509150915043606681905550846001600160a01b03168215157f57508e122d767c940c66fc3762dd79760bb534bd2f7c4902ae678853f7dec9af8686856040516169db93929190618927565b60405180910390a381614699576146998688865b806000036169fc57505050565b6040516302b385fb60e01b8152600481018390526000906001600160a01b038516906302b385fb90602401600060405180830381865afa158015616a44573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616a6c9190810190618952565b505050505050509350505050600080826001600160a01b03168460405160006040518083038185875af1925050503d8060008114616ac6576040519150601f19603f3d011682016040523d82523d6000602084013e616acb565b606091505b509150915081614be157826001600160a01b0316857f2c57f69ad84301745a13b9cc61ef37ef145f0db4f64480f09ccd360cc093dc668684604051616b119291906185b0565b60405180910390a3505050505050565b600080616b2d84614c43565b5091509150828214616b775760405162461bcd60e51b8152602060048201526013602482015272496e76616c696420766f74696e67207479706560681b60448201526064016110fb565b60028114616bbe5760405162461bcd60e51b8152602060048201526014602482015273496e76616c696420766f74696e6720737461746560601b60448201526064016110fb565b600080616bca866155ac565b9250925050616bd861138990565b82101580616be857506113898110155b80616bfd5750616bf881836182f2565b612710145b614be15760405162461bcd60e51b8152602060048201526011602482015270139bdd081e595d08199a5b985b1a5e9959607a1b60448201526064016110fb565b6000806000806060806060600080616c53614b63565b6001600160a01b03166373df4e018b6040518263ffffffff1660e01b8152600401616c8091815260200190565b600060405180830381865afa158015616c9d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616cc59190810190618a07565b9850985098509850985098509850985098509193959799909294969850565b6000616cee614dbc565b6040516325d998bb60e01b81526001600160a01b03848116600483015291909116906325d998bb90602401614b22565b600080616d2a846175fb565b91509150616d3661474a565b616d4082846182f2565b1115616dd05760405162461bcd60e51b815260206004820152605360248201527f6d696e5374616b696e672076616c7565206d757374206265206772656174657260448201527f207468616e206f7220657175616c20746f207468652073756d206f6620756e6c6064820152726f636b416d6f756e742c20736c617368696e6760681b608482015260a4016110fb565b6000616dda614dbc565b604051632c9aab9b60e11b81526001600160a01b038681166004830152919250600091831690635935573690602401602060405180830381865afa158015616e26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616e4a9190618242565b90506000616e5661474a565b616e6090836181bd565b905084821115616ee257616e74868661767c565b604051634954a85b60e11b81526001600160a01b03878116600483015260248201869052604482018390528416906392a950b690606401600060405180830381600087803b158015616ec557600080fd5b505af1158015616ed9573d6000803e3d6000fd5b50505050614699565b614699868361767c565b60008a616efe57616efe8c6003616b21565b616f078a613a9a565b616f69578b600080516020618b5d833981519152604051616f59906020808252601b908201527f4f6c642061646472657373206973206e6f742061206d656d6265720000000000604082015260600190565b60405180910390a2506000617477565b6001600160a01b03808b166000818152606d6020526040902054918b16146170e057616f948a613a9a565b15616ff9578c600080516020618b5d833981519152604051616fe7906020808252601f908201527f6e6577206164647265737320697320616c72656164792061206d656d62657200604082015260600190565b60405180910390a26000915050617477565b886001600160a01b03168a6001600160a01b03161415801561702d5750876001600160a01b03168a6001600160a01b031614155b15617078578c600080516020618b5d833981519152604051616fe790602080825260159082015274496e76616c696420766f746572206164647265737360581b604082015260600190565b61708061474a565b831080617093575082617091612cf2565b105b156170b4578c600080516020618b5d833981519152604051616fe7906183ca565b826170be8b616ce4565b10156170e0578c600080516020618b5d833981519152604051616fe7906187c1565b6001600160a01b038b16600090815260726020908152604080832054808452607190925280832081516080810190925280549293928290829061712290618358565b80601f016020809104026020016040519081016040528092919081815260200182805461714e90618358565b801561719b5780601f106171705761010080835404028352916020019161719b565b820191906000526020600020905b81548152906001019060200180831161717e57829003601f168201915b505050505081526020016001820180546171b490618358565b80601f01602080910402602001604051908101604052809291908181526020018280546171e090618358565b801561722d5780601f106172025761010080835404028352916020019161722d565b820191906000526020600020905b81548152906001019060200180831161721057829003601f168201915b5050505050815260200160028201805461724690618358565b80601f016020809104026020016040519081016040528092919081815260200182805461727290618358565b80156172bf5780601f10617294576101008083540402835291602001916172bf565b820191906000526020600020905b8154815290600101906020018083116172a257829003601f168201915b5050505050815260200160038201548152505090506172e1898989898561511e565b617315578e600080516020618b5d833981519152604051617301906183f7565b60405180910390a260009350505050617477565b506000828152606a60205260409020546001600160a01b03908116908d8116908b16148015906173575750896001600160a01b0316816001600160a01b031614155b801561737657506173678a613a9a565b8061737657506173768a612a45565b156173c2578e600080516020618b5d83398151915260405161730190602080825260169082015275496e76616c696420726577617264206164647265737360501b604082015260600190565b506000828152606760205260409020546001600160a01b03908116908d8116908c161480159061740457508a6001600160a01b0316816001600160a01b031614155b801561742357506174148b613a9a565b8061742357506174238b612a45565b1561746f578e600080516020618b5d83398151915260405161730190602080825260169082015275496e76616c696420766f74657273206164647265737360501b604082015260600190565b506001925050505b9b9a5050505050505050505050565b61748e614acc565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa1580156174d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906174f89190618163565b614ac95760405162461bcd60e51b81526004016110fb90618180565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561754c57617547836176bb565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156175a6575060408051601f3d908101601f191682019092526175a391810190618242565b60015b6175c25760405162461bcd60e51b81526004016110fb9061825b565b600080516020618b1d83398151915281146175ef5760405162461bcd60e51b81526004016110fb906182a9565b50617547838383617757565b600080617606614b63565b6001600160a01b0316638c7be692846040518263ffffffff1660e01b815260040161763391815260200190565b6040805180830381865afa15801561764f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906176739190618ae5565b91509150915091565b617684614dbc565b604051637eee288d60e01b81526001600160a01b038481166004830152602482018490529190911690637eee288d90604401614bb3565b6001600160a01b0381163b6177285760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016110fb565b600080516020618b1d83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6177608361777c565b60008251118061776d5750805b1561754757611e7f83836177bc565b617785816176bb565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606150188383604051806060016040528060278152602001618b7d6027913960606001600160a01b0384163b6178445760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016110fb565b600080856001600160a01b03168560405161785f919061838c565b600060405180830381855af49150503d806000811461789a576040519150601f19603f3d011682016040523d82523d6000602084013e61789f565b606091505b50915091506178af8282866178b9565b9695505050505050565b606083156178c8575081615018565b8251156178d85782518084602001fd5b8160405162461bcd60e51b81526004016110fb9190618b09565b8280546178fe90618358565b90600052602060002090601f0160209004810192826179205760008555617966565b82601f1061793957805160ff1916838001178555617966565b82800160010185558215617966579182015b8281111561796657825182559160200191906001019061794b565b50617972929150617a27565b5090565b82805461798290618358565b90600052602060002090601f0160209004810192826179a45760008555617966565b82601f106179b55780548555617966565b8280016001018555821561796657600052602060002091601f016020900482015b828111156179665782548255916001019190600101906179d6565b5080546179fd90618358565b6000825580601f10617a0d575050565b601f016020900490600052602060002090810190614ac991905b5b808211156179725760008155600101617a28565b6001600160a01b0381168114614ac957600080fd5b8035612e9181617a3c565b634e487b7160e01b600052604160045260246000fd5b60405161014081016001600160401b0381118282101715617a9557617a95617a5c565b60405290565b604051601f8201601f191681016001600160401b0381118282101715617ac357617ac3617a5c565b604052919050565b60006001600160401b03821115617ae457617ae4617a5c565b50601f01601f191660200190565b600082601f830112617b0357600080fd5b8135617b16617b1182617acb565b617a9b565b818152846020838601011115617b2b57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600060608486031215617b5d57600080fd5b8335617b6881617a3c565b925060208401356001600160401b03811115617b8357600080fd5b617b8f86828701617af2565b925050604084013590509250925092565b600060208284031215617bb257600080fd5b813561501881617a3c565b600060208284031215617bcf57600080fd5b5035919050565b600080600060608486031215617beb57600080fd5b8335617bf681617a3c565b92506020840135915060408401356001600160401b03811115617c1857600080fd5b617c2486828701617af2565b9150509250925092565b60006101408284031215617c4157600080fd5b617c49617a72565b9050617c5482617a51565b8152617c6260208301617a51565b6020820152617c7360408301617a51565b604082015260608201356001600160401b0380821115617c9257600080fd5b617c9e85838601617af2565b60608401526080840135915080821115617cb757600080fd5b617cc385838601617af2565b608084015260a0840135915080821115617cdc57600080fd5b617ce885838601617af2565b60a084015260c084013560c084015260e084013560e084015261010091508184013581811115617d1757600080fd5b617d2386828701617af2565b8385015250505061012080830135818301525092915050565b600060208284031215617d4e57600080fd5b81356001600160401b03811115617d6457600080fd5b617d7084828501617c2e565b949350505050565b600080600060608486031215617d8d57600080fd5b8335617d9881617a3c565b9250602084013591506040840135617daf81617a3c565b809150509250925092565b60008060008060808587031215617dd057600080fd5b8435617ddb81617a3c565b935060208501356001600160401b0380821115617df757600080fd5b617e0388838901617af2565b94506040870135915080821115617e1957600080fd5b50617e2687828801617af2565b949793965093946060013593505050565b600080600080600060a08688031215617e4f57600080fd5b853594506020860135935060408601356001600160401b0380821115617e7457600080fd5b617e8089838a01617af2565b94506060880135915080821115617e9657600080fd5b50617ea388828901617af2565b95989497509295608001359392505050565b60005b83811015617ed0578181015183820152602001617eb8565b83811115611e7f5750506000910152565b60008151808452617ef9816020860160208601617eb5565b601f01601f19169290920160200192915050565b608081526000617f206080830187617ee1565b8281036020840152617f328187617ee1565b90508281036040840152617f468186617ee1565b91505082606083015295945050505050565b60008060408385031215617f6b57600080fd5b8235617f7681617a3c565b915060208301356001600160401b03811115617f9157600080fd5b617f9d85828601617af2565b9150509250929050565b60008060008060008060c08789031215617fc057600080fd5b8635617fcb81617a3c565b95506020870135945060408701356001600160401b03811115617fed57600080fd5b617ff989828a01617af2565b945050606087013592506080870135915060a087013590509295509295509295565b6000806000806080858703121561803157600080fd5b84356001600160401b0381111561804757600080fd5b61805387828801617c2e565b945050602085013561806481617a3c565b93969395505050506040820135916060013590565b60008060008060008060c0878903121561809257600080fd5b863561809d81617a3c565b95506020870135945060408701356001600160401b03808211156180c057600080fd5b6180cc8a838b01617af2565b955060608901359150808211156180e257600080fd5b6180ee8a838b01617af2565b9450608089013591508082111561810457600080fd5b5061811189828a01617af2565b92505060a087013590509295509295509295565b8015158114614ac957600080fd5b6000806040838503121561814657600080fd5b82359150602083013561815881618125565b809150509250929050565b60006020828403121561817557600080fd5b815161501881618125565b6020808252600d908201526c2737902832b936b4b9b9b4b7b760991b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6000828210156181cf576181cf6181a7565b500390565b6020808252601d908201527f43616e6e6f74206164642070726f706f73616c20746f6f206561726c79000000604082015260600190565b60208082526017908201527f496e76616c6964207374616b696e672062616c616e6365000000000000000000604082015260600190565b60006020828403121561825457600080fd5b5051919050565b6020808252602e908201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960408201526d6f6e206973206e6f74205555505360901b606082015260800190565b60208082526029908201527f45524331393637557067726164653a20756e737570706f727465642070726f786040820152681a58589b195555525160ba1b606082015260800190565b60008219821115618305576183056181a7565b500190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b600181811c9082168061836c57607f821691505b602082108103613a9457634e487b7160e01b600052602260045260246000fd5b6000825161839e818460208701617eb5565b9190910192915050565b600083516183ba818460208801617eb5565b9190910191825250602001919050565b602080825260139082015272125b9d985b1a59081b1bd8dac8185b5bdd5b9d606a1b604082015260600190565b6020808252601490820152734475706c696361746564206e6f646520696e666f60601b604082015260600190565b878152866020820152856040820152600060018060a01b0380871660608401528086166080840152508360a083015260e060c083015261846860e0830184617ee1565b9998505050505050505050565b87815286602082015285604082015260018060a01b03851660608201528360808201528260a082015260e060c0820152600061846860e0830184617ee1565b6000602082840312156184c657600080fd5b815161501881617a3c565b8051612e9181617a3c565b600082601f8301126184ed57600080fd5b81516184fb617b1182617acb565b81815284602083860101111561851057600080fd5b617d70826020830160208701617eb5565b6000806000806080858703121561853757600080fd5b84516001600160401b038082111561854e57600080fd5b61855a888389016184dc565b9550602087015191508082111561857057600080fd5b61857c888389016184dc565b9450604087015191508082111561859257600080fd5b5061859f878288016184dc565b606096909601519497939650505050565b828152604060208201526000617d706040830184617ee1565b6000806000606084860312156185de57600080fd5b83519250602084015191506040840151617daf81618125565b60008060006060848603121561860c57600080fd5b8351925060208401519150604084015190509250925092565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b8c8152602081018c9052604081018b90526001600160a01b038a81166060830152898116608083015288811660a083015287811660c0830152861660e08201526000610180806101008401526186c881840188617ee1565b90508281036101208401526186dd8187617ee1565b90508281036101408401526186f28186617ee1565b915050826101608301529d9c50505050505050505050505050565b60008060006060848603121561872257600080fd5b835161872d81617a3c565b6020850151604086015191945092506001600160401b0381111561875057600080fd5b617c24868287016184dc565b60208082526023908201527f4e6f7720696e20766f74696e67207769746820646966666572656e742062616c6040820152621b1bdd60ea1b606082015260800190565b6000826187bc57634e487b7160e01b600052601260045260246000fd5b500490565b60208082526027908201527f496e73756666696369656e742062616c616e636520746861742063616e206265604082015266081b1bd8dad95960ca1b606082015260800190565b8054600090600181811c908083168061882257607f831692505b6020808410820361884357634e487b7160e01b600052602260045260246000fd5b818015618857576001811461886857618895565b60ff19861689528489019650618895565b60008881526020902060005b8681101561888d5781548b820152908501908301618874565b505084890196505b50505050505092915050565b60006150188284618808565b60006188b98285618808565b9283525050602001919050565b6000806000606084860312156188db57600080fd5b835192506020840151915060408401516001600160401b0381111561875057600080fd5b83815282602082015260606040820152600061891e6060830184617ee1565b95945050505050565b8381526060602082015260006189406060830185617ee1565b82810360408401526178af8185617ee1565b60008060008060008060008060008060006101608c8e03121561897457600080fd5b8b519a5060208c0151995060408c0151985060608c015161899481617a3c565b60808d01519098506001600160401b038111156189b057600080fd5b6189bc8e828f016184dc565b97505060a08c0151955060c08c0151945060e08c015193506101008c015192506101208c01516189eb81618125565b809250506101408c015190509295989b509295989b9093969950565b60008060008060008060008060006101208a8c031215618a2657600080fd5b618a2f8a6184d1565b9850618a3d60208b016184d1565b9750618a4b60408b016184d1565b9650618a5960608b016184d1565b955060808a01516001600160401b0380821115618a7557600080fd5b618a818d838e016184dc565b965060a08c0151915080821115618a9757600080fd5b618aa38d838e016184dc565b955060c08c0151915080821115618ab957600080fd5b50618ac68c828d016184dc565b93505060e08a015191506101008a015190509295985092959850929598565b60008060408385031215618af857600080fd5b505080516020909101519092909150565b6020815260006150186020830184617ee156fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249885e7f4987c0698db47045ad8cea110b51138f0eecbd94915842328cf6c3dc97d416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212209dfea10b5c3f59723a267d8b27b3a510120d9c4c19c0de8eb23e161703dede8364736f6c634300080e0033",
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

// AddProposalToExecute is a paid mutator transaction binding the contract method 0x404e585d.
//
// Solidity: function addProposalToExecute(address _target, bytes _calldata, bytes _memo, uint256 _duration) payable returns()
func (_GovImp *GovImpTransactor) AddProposalToExecute(opts *bind.TransactOpts, _target common.Address, _calldata []byte, _memo []byte, _duration *big.Int) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "addProposalToExecute", _target, _calldata, _memo, _duration)
}

// AddProposalToExecute is a paid mutator transaction binding the contract method 0x404e585d.
//
// Solidity: function addProposalToExecute(address _target, bytes _calldata, bytes _memo, uint256 _duration) payable returns()
func (_GovImp *GovImpSession) AddProposalToExecute(_target common.Address, _calldata []byte, _memo []byte, _duration *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.AddProposalToExecute(&_GovImp.TransactOpts, _target, _calldata, _memo, _duration)
}

// AddProposalToExecute is a paid mutator transaction binding the contract method 0x404e585d.
//
// Solidity: function addProposalToExecute(address _target, bytes _calldata, bytes _memo, uint256 _duration) payable returns()
func (_GovImp *GovImpTransactorSession) AddProposalToExecute(_target common.Address, _calldata []byte, _memo []byte, _duration *big.Int) (*types.Transaction, error) {
	return _GovImp.Contract.AddProposalToExecute(&_GovImp.TransactOpts, _target, _calldata, _memo, _duration)
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

// GovImpExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GovImp contract.
type GovImpExecutedIterator struct {
	Event *GovImpExecuted // Event containing the contract specifics and raw log

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
func (it *GovImpExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpExecuted)
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
		it.Event = new(GovImpExecuted)
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
func (it *GovImpExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpExecuted represents a Executed event raised by the GovImp contract.
type GovImpExecuted struct {
	Success    bool
	To         common.Address
	Value      *big.Int
	Calldatas  []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x57508e122d767c940c66fc3762dd79760bb534bd2f7c4902ae678853f7dec9af.
//
// Solidity: event Executed(bool indexed success, address indexed to, uint256 value, bytes calldatas, bytes returnData)
func (_GovImp *GovImpFilterer) FilterExecuted(opts *bind.FilterOpts, success []bool, to []common.Address) (*GovImpExecutedIterator, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "Executed", successRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GovImpExecutedIterator{contract: _GovImp.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x57508e122d767c940c66fc3762dd79760bb534bd2f7c4902ae678853f7dec9af.
//
// Solidity: event Executed(bool indexed success, address indexed to, uint256 value, bytes calldatas, bytes returnData)
func (_GovImp *GovImpFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GovImpExecuted, success []bool, to []common.Address) (event.Subscription, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "Executed", successRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpExecuted)
				if err := _GovImp.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x57508e122d767c940c66fc3762dd79760bb534bd2f7c4902ae678853f7dec9af.
//
// Solidity: event Executed(bool indexed success, address indexed to, uint256 value, bytes calldatas, bytes returnData)
func (_GovImp *GovImpFilterer) ParseExecuted(log types.Log) (*GovImpExecuted, error) {
	event := new(GovImpExecuted)
	if err := _GovImp.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovImpFailReturnValueIterator is returned from FilterFailReturnValue and is used to iterate over the raw logs and unpacked data for FailReturnValue events raised by the GovImp contract.
type GovImpFailReturnValueIterator struct {
	Event *GovImpFailReturnValue // Event containing the contract specifics and raw log

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
func (it *GovImpFailReturnValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovImpFailReturnValue)
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
		it.Event = new(GovImpFailReturnValue)
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
func (it *GovImpFailReturnValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovImpFailReturnValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovImpFailReturnValue represents a FailReturnValue event raised by the GovImp contract.
type GovImpFailReturnValue struct {
	BallotIdx *big.Int
	Creator   common.Address
	Value     *big.Int
	Result    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFailReturnValue is a free log retrieval operation binding the contract event 0x2c57f69ad84301745a13b9cc61ef37ef145f0db4f64480f09ccd360cc093dc66.
//
// Solidity: event FailReturnValue(uint256 indexed ballotIdx, address indexed creator, uint256 value, bytes result)
func (_GovImp *GovImpFilterer) FilterFailReturnValue(opts *bind.FilterOpts, ballotIdx []*big.Int, creator []common.Address) (*GovImpFailReturnValueIterator, error) {

	var ballotIdxRule []interface{}
	for _, ballotIdxItem := range ballotIdx {
		ballotIdxRule = append(ballotIdxRule, ballotIdxItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _GovImp.contract.FilterLogs(opts, "FailReturnValue", ballotIdxRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &GovImpFailReturnValueIterator{contract: _GovImp.contract, event: "FailReturnValue", logs: logs, sub: sub}, nil
}

// WatchFailReturnValue is a free log subscription operation binding the contract event 0x2c57f69ad84301745a13b9cc61ef37ef145f0db4f64480f09ccd360cc093dc66.
//
// Solidity: event FailReturnValue(uint256 indexed ballotIdx, address indexed creator, uint256 value, bytes result)
func (_GovImp *GovImpFilterer) WatchFailReturnValue(opts *bind.WatchOpts, sink chan<- *GovImpFailReturnValue, ballotIdx []*big.Int, creator []common.Address) (event.Subscription, error) {

	var ballotIdxRule []interface{}
	for _, ballotIdxItem := range ballotIdx {
		ballotIdxRule = append(ballotIdxRule, ballotIdxItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _GovImp.contract.WatchLogs(opts, "FailReturnValue", ballotIdxRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovImpFailReturnValue)
				if err := _GovImp.contract.UnpackLog(event, "FailReturnValue", log); err != nil {
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

// ParseFailReturnValue is a log parse operation binding the contract event 0x2c57f69ad84301745a13b9cc61ef37ef145f0db4f64480f09ccd360cc093dc66.
//
// Solidity: event FailReturnValue(uint256 indexed ballotIdx, address indexed creator, uint256 value, bytes result)
func (_GovImp *GovImpFilterer) ParseFailReturnValue(log types.Log) (*GovImpFailReturnValue, error) {
	event := new(GovImpFailReturnValue)
	if err := _GovImp.contract.UnpackLog(event, "FailReturnValue", log); err != nil {
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
