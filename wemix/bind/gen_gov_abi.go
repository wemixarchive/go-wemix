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
	Bin: "0x608060405234801561001057600080fd5b5060405161077238038061077283398101604081905261002f91610326565b604080516020810190915260008152819061006b60017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd61039e565b60008051602061072b83398151915214610087576100876103ef565b6100938282600061009b565b505050610405565b6100a4836100d1565b6000825111806100b15750805b156100cc576100ca838361011160201b61008b1760201c565b505b505050565b6100da8161013d565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060610136838360405180606001604052806027815260200161074b602791396101fd565b9392505050565b610150816102db60201b6100b71760201c565b6101b75760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b806101dc60008051602061072b83398151915260001b6102ea60201b6100c61760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606001600160a01b0384163b6102655760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016101ae565b600080856001600160a01b031685604051610280919061034f565b600060405180830381855af49150503d80600081146102bb576040519150601f19603f3d011682016040523d82523d6000602084013e6102c0565b606091505b5090925090506102d18282866102ed565b9695505050505050565b6001600160a01b03163b151590565b90565b606083156102fc575081610136565b82511561030c5782518084602001fd5b8160405162461bcd60e51b81526004016101ae919061036b565b60006020828403121561033857600080fd5b81516001600160a01b038116811461013657600080fd5b600082516103618184602087016103c3565b9190910192915050565b602081526000825180602084015261038a8160408501602087016103c3565b601f01601f19169190910160400192915050565b6000828210156103be57634e487b7160e01b600052601160045260246000fd5b500390565b60005b838110156103de5781810151838201526020016103c6565b838111156100ca5750506000910152565b634e487b7160e01b600052600160045260246000fd5b610317806104146000396000f3fe6080604052600436106100225760003560e01c80635c60da1b1461003957610031565b366100315761002f61006a565b005b61002f61006a565b34801561004557600080fd5b5061004e61007c565b6040516001600160a01b03909116815260200160405180910390f35b61007a6100756100c9565b6100fc565b565b60006100866100c9565b905090565b60606100b083836040518060600160405280602781526020016102bb60279139610120565b9392505050565b6001600160a01b03163b151590565b90565b60006100867f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b3660008037600080366000845af43d6000803e80801561011b573d6000f35b3d6000fd5b60606001600160a01b0384163b61018d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084015b60405180910390fd5b600080856001600160a01b0316856040516101a8919061023b565b600060405180830381855af49150503d80600081146101e3576040519150601f19603f3d011682016040523d82523d6000602084013e6101e8565b606091505b50915091506101f8828286610202565b9695505050505050565b606083156102115750816100b0565b8251156102215782518084602001fd5b8160405162461bcd60e51b81526004016101849190610257565b6000825161024d81846020870161028a565b9190910192915050565b602081526000825180602084015261027681604085016020870161028a565b601f01601f19169190910160400192915050565b60005b838110156102a557818101518382015260200161028d565b838111156102b4576000848401525b5050505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212203c9a20da5aa52c7b0291aa3d5fe1e7f3a44f26cfd6cdbbe90d752ce6055b52cb64736f6c63430008060033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
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
	Bin: "0x60a06040523060601b6080523480156200001857600080fd5b506200002362000029565b620000eb565b600054610100900460ff1615620000965760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015620000e9576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60805160601c6187bb6200010a60003960006129b701526187bb6000f3fe6080604052600436106104675760003560e01c806372016f751161024a578063af6af2ff11610139578063cec5b622116100b6578063e27bdaef1161007a578063e27bdaef14610fae578063e75235b814610fce578063e9523fb514610fe3578063f2fde38b14610ff9578063f38ecf471461101957600080fd5b8063cec5b62214610f0b578063d07bff0c14610f38578063d6f9cfce14610f6e578063d965ea0014610f84578063de09b37714610f9957600080fd5b8063c6713baf116100fd578063c6713baf14610e38578063c7d3da3414610e6c578063c9d27afe14610ea0578063ce04b9d414610ec0578063ce6a54ff14610ed557600080fd5b8063af6af2ff14610d53578063b128f88014610d68578063c00ace6c14610d9c578063c0b4fe1514610dd0578063c42a0abc14610e0457600080fd5b8063a0c12683116101c7578063a8915a3e1161018b578063a8915a3e14610c7c578063a91ee0dc14610c9c578063a9b629b214610cbc578063aaf0dd3614610cf0578063ab3545e514610d1d57600080fd5b8063a0c1268314610ba3578063a230c52414610bd0578063a6868b7d14610bf0578063a7771ee314610c24578063a78a818814610c5c57600080fd5b8063894f51111161020e578063894f511114610b0b5780638d39e33a14610b2b5780638da5cb5b14610b4b578063918f867414610b695780639986e4b914610b7f57600080fd5b806372016f7514610a6b578063738fdd1a14610a805780637b2bfb0114610aa05780637bf4653014610ad45780637d10dd1b14610af557600080fd5b80633f35c8fe116103665780635aaa4040116102e35780636d583ca7116102a75780636d583ca7146109965780636f1e8533146109ca5780636f6de96d14610a025780636fde207a14610a22578063715018a614610a5657600080fd5b80635aaa4040146108db5780636167eb45146108f0578063656e3052146109245780636ba99181146109585780636c78d2cf1461096d57600080fd5b80634d5ce0381161032a5780634d5ce038146108445780634f0f4aa9146108645780634f1ef2861461089457806352d1902d146108a25780635a731cca146108b757600080fd5b80633f35c8fe14610766578063406903531461079a57806346946416146107ba5780634bd1ed76146107ee5780634d273e281461081057600080fd5b8063238737b6116103f4578063351bacda116103b8578063351bacda146106b25780633659cfe6146106d257806336e83d83146106f25780633829441914610712578063397e38e71461074657600080fd5b8063238737b6146105f3578063278bb12a146106275780632f40992e1461065b5780633310569c1461067c57806334125c841461069257600080fd5b806316fbe8311161043b57806316fbe8311461054e5780631c12b030146105655780631c1501711461058a5780631c4b774b1461059f5780631e0cba0d146105d557600080fd5b806215a73b1461046c5780630efa4909146104b3578063139d9dd3146104d357806315bf6b4d14610500575b600080fd5b34801561047857600080fd5b506104a07f0c4fbe9dc9de15dd7c0d064975ee1a2f2f9b954fa0e65d4f6cddba94884bdc3e81565b6040519081526020015b60405180910390f35b3480156104bf57600080fd5b506104a06104ce366004617af6565b61104d565b3480156104df57600080fd5b506104a06104ee36600461798f565b60ab6020526000908152604090205481565b34801561050c57600080fd5b5061053661051b366004617f22565b6000908152607360205260409020546001600160a01b031690565b6040516001600160a01b0390911681526020016104aa565b34801561055a57600080fd5b50610563611406565b005b34801561057157600080fd5b5061057a611744565b60405190151581526020016104aa565b34801561059657600080fd5b506104a0611795565b3480156105ab57600080fd5b506105366105ba366004617f22565b6000908152606a60205260409020546001600160a01b031690565b3480156105e157600080fd5b506104a0665374616b696e6760c81b81565b3480156105ff57600080fd5b506104a07f1d36f8ce53f59e624857e1d8dc7932d19981a2ea1b8faa4eb8ff843fc3e5a27881565b34801561063357600080fd5b506104a07f9b2e0c7fdae148f225bae7deb92d7e7bd24bb77edb12956e8fa7204900dd8a2281565b34801561066757600080fd5b506104a06914995dd85c99141bdbdb60b21b81565b34801561068857600080fd5b506104a060aa5481565b34801561069e57600080fd5b506104a06845636f73797374656d60b81b81565b3480156106be57600080fd5b506105636106cd366004617b90565b611814565b3480156106de57600080fd5b506105636106ed36600461798f565b611ea3565b3480156106fe57600080fd5b506104a061070d366004617e90565b611edc565b34801561071e57600080fd5b506104a07fbe90e461bbdb9a95a694f7796912ea04244caf7f5b60ad7ded17e16821d3e44c81565b34801561075257600080fd5b50610563610761366004617b4e565b612403565b34801561077257600080fd5b506104a07f2a268972a70c8c688b62366bdfdd9bb09cf19d3e5b6e7e7bb158e671ffdcedd281565b3480156107a657600080fd5b506104a06107b5366004617d83565b6124fe565b3480156107c657600080fd5b506104a07fdd5a41a7fc01f5c6d30816b17f638d6531625f1e1eaa599673ab2f6079f2dd9d81565b3480156107fa57600080fd5b506104a06a4d61696e74656e616e636560a81b81565b34801561081c57600080fd5b506104a07f77884798208df1e64f70968be41ef2d3211ec53613397ca59313416813df088881565b34801561085057600080fd5b5061057a61085f36600461798f565b6127b0565b34801561087057600080fd5b5061088461087f366004617f22565b6127cd565b6040516104aa9493929190618155565b6105636106ed366004617aa7565b3480156108ae57600080fd5b506104a06129aa565b3480156108c357600080fd5b506104a06c14dd185ada5b99d4995dd85c99609a1b81565b3480156108e757600080fd5b506104a0612a5d565b3480156108fc57600080fd5b506104a07f9f1de481f899d76889aa8a2b44cc7b604d42691aa93d4ba618a1a1fd439f505081565b34801561093057600080fd5b506104a07fe10074dceffb75f13bf0ce50145afd35182d63796823f1280ce40e01c19109e781565b34801561096457600080fd5b50610563612a9f565b34801561097957600080fd5b506104a07111dbdd995c9b985b98d950dbdb9d1c9858dd60721b81565b3480156109a257600080fd5b506104a07fc69fc6b7d0efc934fd5a3581c7253a7107a952526bb6dbcd814ef8d8dae1f44a81565b3480156109d657600080fd5b5061057a6109e536600461798f565b6001600160a01b03166000908152606d6020526040902054151590565b348015610a0e57600080fd5b50610536610a1d36600461798f565b612b9c565b348015610a2e57600080fd5b506104a07f0b09c9badbbeb6c813a598ee910770a39ccda797a1940439bb6e47fc6c87548b81565b348015610a6257600080fd5b50610563612c0b565b348015610a7757600080fd5b506074546104a0565b348015610a8c57600080fd5b50606554610536906001600160a01b031681565b348015610aac57600080fd5b506104a07f9346226931826838eedd13d9677fa33551e7c81f604b171ef3fac355458da9aa81565b348015610ae057600080fd5b506104a069456e7653746f7261676560b01b81565b348015610b0157600080fd5b506104a060665481565b348015610b1757600080fd5b506104a0610b26366004617c94565b612c1f565b348015610b3757600080fd5b506104a0610b4636600461798f565b612f9c565b348015610b5757600080fd5b506033546001600160a01b0316610536565b348015610b7557600080fd5b506104a061271081565b348015610b8b57600080fd5b506104a06c42616c6c6f7453746f7261676560981b81565b348015610baf57600080fd5b506104a0610bbe36600461798f565b606d6020526000908152604090205481565b348015610bdc57600080fd5b5061057a610beb36600461798f565b6138e1565b348015610bfc57600080fd5b506104a07f6c6f69f426081752a5d3e73746599acd2a4cb145d5de4203ca1e3473b281680b81565b348015610c3057600080fd5b5061057a610c3f36600461798f565b6001600160a01b0316600090815260686020526040902054151590565b348015610c6857600080fd5b506104a0610c77366004617ec4565b613924565b348015610c8857600080fd5b50610563610c97366004617be8565b61409f565b348015610ca857600080fd5b50610563610cb736600461798f565b614517565b348015610cc857600080fd5b506104a07f89dd490ecaf395283ed4ff2fd9557ca767fc425dce063451a9b0da6d72f600c381565b348015610cfc57600080fd5b506104a0610d0b36600461798f565b606b6020526000908152604090205481565b348015610d2957600080fd5b50610536610d38366004617f22565b6000908152606c60205260409020546001600160a01b031690565b348015610d5f57600080fd5b506104a06145bf565b348015610d7457600080fd5b506104a07f829561ab7af084b7efc6600518d2df79b8d95f3f4c3a550f54f8f7ec7d2b805781565b348015610da857600080fd5b506104a07f18ad4415ef4a621ce1a136395c51ab6c3712bb2e24b79d526059925cea58dcb881565b348015610ddc57600080fd5b506104a07f8086da5becff4dfac91a3105821b361078d2d4abba0ccc2401b974cf0dcf05c181565b348015610e1057600080fd5b506104a07fb38b2c133e931937bd95337c65c8aefa7040ed64bbed732e3e29a4944c65747381565b348015610e4457600080fd5b506104a07fc9e15e34073efbcd0328740feaf503caac9124b55b38c73d1a97b53da80a2f6081565b348015610e7857600080fd5b506104a07f04f7b94450bbcad85f37ea47cd1062728f884bb2040e357738f8fd53056134bc81565b348015610eac57600080fd5b50610563610ebb366004617f3b565b614601565b348015610ecc57600080fd5b506104a0614806565b348015610ee157600080fd5b506104a0610ef036600461798f565b6001600160a01b031660009081526072602052604090205490565b348015610f1757600080fd5b506104a0610f2636600461798f565b60686020526000908152604090205481565b348015610f4457600080fd5b50610536610f53366004617f22565b6000908152606760205260409020546001600160a01b031690565b348015610f7a57600080fd5b506104a060755481565b348015610f9057600080fd5b506069546104a0565b348015610fa557600080fd5b506077546104a0565b348015610fba57600080fd5b50610563610fc9366004617f22565b614848565b348015610fda57600080fd5b506113896104a0565b348015610fef57600080fd5b506104a060765481565b34801561100557600080fd5b5061056361101436600461798f565b6148ce565b34801561102557600080fd5b506104a07f7c1150f0e1a39ff55552d52764f97e6c387e2a247e1df344369f122c4254be2f81565b6000611057614947565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b15801561109857600080fd5b505afa1580156110ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d09190617d08565b6110f55760405162461bcd60e51b81526004016110ec906182c8565b60405180910390fd5b600061110033612b9c565b60aa546001600160a01b038216600090815260ab60205260409020549192509061112a9042618637565b10156111485760405162461bcd60e51b81526004016110ec9061845d565b600061115333612b9c565b905061115d612a5d565b61116682614967565b1115801561118357506111776145bf565b61118082614967565b10155b61119f5760405162461bcd60e51b81526004016110ec906183db565b6001600160a01b0386166111f55760405162461bcd60e51b815260206004820152601d60248201527f496d706c656d656e746174696f6e2063616e6e6f74206265207a65726f00000060448201526064016110ec565b6000805160206186ff833981519152546001600160a01b03878116911614156112585760405162461bcd60e51b815260206004820152601560248201527453616d6520636f6e7472616374206164647265737360581b60448201526064016110ec565b856001600160a01b03166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561129157600080fd5b505afa9250505080156112c1575060408051601f3d908101601f191682019092526112be91810190617d25565b60015b6112dd5760405162461bcd60e51b81526004016110ec9061831d565b6000805160206186ff833981519152811461130a5760405162461bcd60e51b81526004016110ec90618231565b506075546113199060016185fd565b92506113236149ed565b6001600160a01b0316630a3a63fe607554600161134091906185fd565b60046040516001600160e01b031960e085901b16815260048101929092526024820152604481018790523360648201526001600160a01b038916608482015260a401602060405180830381600087803b15801561139c57600080fd5b505af11580156113b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113d49190617d25565b506113df8386614a08565b5060758290556001600160a01b0316600090815260ab602052604090204290559392505050565b600054600290610100900460ff16158015611428575060005460ff8083169116105b6114445760405162461bcd60e51b81526004016110ec9061827a565b6000805461ffff191660ff831617610100179055611460614a73565b60015b6069548111611710576000818152607160205260408082208151608081019092528054829082906114939061867a565b80601f01602080910402602001604051908101604052809291908181526020018280546114bf9061867a565b801561150c5780601f106114e15761010080835404028352916020019161150c565b820191906000526020600020905b8154815290600101906020018083116114ef57829003601f168201915b505050505081526020016001820180546115259061867a565b80601f01602080910402602001604051908101604052809291908181526020018280546115519061867a565b801561159e5780601f106115735761010080835404028352916020019161159e565b820191906000526020600020905b81548152906001019060200180831161158157829003601f168201915b505050505081526020016002820180546115b79061867a565b80601f01602080910402602001604051908101604052809291908181526020018280546115e39061867a565b80156116305780601f1061160557610100808354040283529160200191611630565b820191906000526020600020905b81548152906001019060200180831161161357829003601f168201915b5050505050815260200160038201548152505090506001606e826000015160405161165b91906180b1565b90815260405160209181900382018120805460ff191693151593909317909255820151600191606f9161168d916180b1565b908152602001604051809103902060006101000a81548160ff021916908315150217905550600160706000836040015184606001516040516020016116d39291906180cd565b60408051808303601f19018152918152815160209283012083529082019290925201600020805460ff191691151591909117905550600101611463565b506000805461ff001916905560405160ff8216815260008051602061871f833981519152906020015b60405180910390a150565b600060775460001461178f57600061175d607754614acd565b50915050600061176e607754614b64565b50915050600282141561178c574281101561178c5760009250505090565b50505b50600190565b600061179f614bee565b6001600160a01b03166333be496e6040518163ffffffff1660e01b815260040160206040518083038186803b1580156117d757600080fd5b505afa1580156117eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061180f9190617d25565b905090565b600054610100900460ff16158080156118345750600054600160ff909116105b8061184e5750303b15801561184e575060005460ff166001145b61186a5760405162461bcd60e51b81526004016110ec9061827a565b6000805460ff19166001179055801561188d576000805461ff0019166101001790555b611895614c06565b61189d614c35565b6118a684614517565b4360665560006118b4614c64565b90506118be6145bf565b84101580156118d45750836118d1612a5d565b10155b6118f05760405162461bcd60e51b81526004016110ec9061836b565b600080600060608060606000806000905060008060208d0191508c518261191791906185fd565b90505b80821015611e5457815199506119316020836185fd565b915080821061193f57600080fd5b8151985061194e6020836185fd565b915080821061195c57600080fd5b8151975061196b6020836185fd565b915080821061197957600080fd5b8196508651602061198a91906185fd565b61199490836185fd565b91508082106119a257600080fd5b819550855160206119b391906185fd565b6119bd90836185fd565b91508082106119cb57600080fd5b819450845160206119dc91906185fd565b6119e690836185fd565b91508082106119f457600080fd5b81519350611a036020836185fd565b9150611a106001846185fd565b9250611a1b8a6138e1565b158015611a2e5750611a2c896138e1565b155b8015611a405750611a3e886127b0565b155b611a7d5760405162461bcd60e51b815260206004820152600e60248201526d20b63932b0b23c9036b2b6b132b960911b60448201526064016110ec565b886067600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606860008b6001600160a01b03166001600160a01b031681526020019081526020016000208190555087606a600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606b60008a6001600160a01b03166001600160a01b031681526020019081526020016000208190555089606c600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606d60008c6001600160a01b03166001600160a01b0316815260200190815260200160002081905550886001600160a01b03168a6001600160a01b03167f6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba660405160405180910390a36040516325d998bb60e01b81526001600160a01b038b811660048301528f91908d16906325d998bb9060240160206040518083038186803b158015611c2757600080fd5b505afa158015611c3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c5f9190617d25565b1015611ca45760405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e74207374616b696e6760601b60448201526064016110ec565b611cb087878787614c79565b611ccc5760405162461bcd60e51b81526004016110ec906182ef565b611cd68a8f614d28565b600083815260716020908152604090912088519091611cf99183918b019061767e565b508651611d0f90600183019060208a019061767e565b508551611d25906002830190602089019061767e565b508481600301819055506001606e89604051611d4191906180b1565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f90611d73908a906180b1565b908152602001604051809103902060006101000a81548160ff0219169083151502179055506001607060008888604051602001611db19291906180cd565b60405160208183030381529060405280519060200120815260200190815260200160002060006101000a81548160ff0219169083151502179055508a6073600086815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555083607260008d6001600160a01b03166001600160a01b03168152602001908152602001600020819055505061191a565b50506069819055607455505086159550611e9d945050505050576000805461ff00191690556040516001815260008051602061871f833981519152906020015b60405180910390a15b50505050565b60405162461bcd60e51b815260206004820152600e60248201526d496e76616c69642061636365737360901b60448201526064016110ec565b6000611ee6614947565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b158015611f2757600080fd5b505afa158015611f3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f5f9190617d08565b611f7b5760405162461bcd60e51b81526004016110ec906182c8565b6000611f8633612b9c565b60aa546001600160a01b038216600090815260ab602052604090205491925090611fb09042618637565b1015611fce5760405162461bcd60e51b81526004016110ec9061845d565b6000611fd933612b9c565b9050611fe3612a5d565b611fec82614967565b111580156120095750611ffd6145bf565b61200682614967565b10155b6120255760405162461bcd60e51b81526004016110ec906183db565b602084015184906001600160a01b03166120715760405162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b2103b37ba32b960991b60448201526064016110ec565b6000816060015151116120ba5760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964206e6f6465206e616d6560781b60448201526064016110ec565b60008160a0015151116121015760405162461bcd60e51b815260206004820152600f60248201526e0496e76616c6964206e6f646520495608c1b60448201526064016110ec565b60008160c00151116121495760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a59081b9bd919481c1bdc9d607a1b60448201526064016110ec565b6000816080015151116121935760405162461bcd60e51b8152602060048201526012602482015271496e76616c6964206e6f646520656e6f646560701b60448201526064016110ec565b600081610100015151116121d85760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964206d656d6f60a01b60448201526064016110ec565b6000816101200151116122205760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b210323ab930ba34b7b760811b60448201526064016110ec565b6122286145bf565b8160e0015110158015612246575061223e612a5d565b8160e0015111155b6122885760405162461bcd60e51b8152602060048201526013602482015272125b9d985b1a59081b1bd8dac8105b5bdd5b9d606a1b60448201526064016110ec565b8451612293906138e1565b1580156122a8575084516122a6906127b0565b155b6122e55760405162461bcd60e51b815260206004820152600e60248201526d20b63932b0b23c9036b2b6b132b960911b60448201526064016110ec565b84602001516001600160a01b031685600001516001600160a01b0316148015612327575084604001516001600160a01b031685600001516001600160a01b0316145b6123695760405162461bcd60e51b815260206004820152601360248201527229ba30b5b2b91034b9903737ba103b37ba32b960691b60448201526064016110ec565b612385856060015186608001518760a001518860c00151614c79565b6123a15760405162461bcd60e51b81526004016110ec906182ef565b6075546123af9060016185fd565b93506123c084600133600089614d67565b6123ce848660e00151614e0a565b6123dd84866101000151614a08565b505060758290556001600160a01b0316600090815260ab60205260409020429055919050565b600054610100900460ff16158080156124235750600054600160ff909116105b8061243d5750303b15801561243d575060005460ff166001145b6124595760405162461bcd60e51b81526004016110ec9061827a565b6000805460ff19166001179055801561247c576000805461ff0019166101001790555b612484614c06565b61248c614c35565b61249584614517565b60668390556124a3826148ce565b60405133907fab2db0a6f442428b686ffa80eadcaabe7d5ee00049c6ae888a237edd3238d85690600090a28015611e9d576000805461ff00191690556040516001815260008051602061871f83398151915290602001611e94565b6000612508614947565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b15801561254957600080fd5b505afa15801561255d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125819190617d08565b61259d5760405162461bcd60e51b81526004016110ec906182c8565b60006125a833612b9c565b60aa546001600160a01b038216600090815260ab6020526040902054919250906125d29042618637565b10156125f05760405162461bcd60e51b81526004016110ec9061845d565b60006125fb33612b9c565b9050612605612a5d565b61260e82614967565b1115801561262b575061261f6145bf565b61262882614967565b10155b6126475760405162461bcd60e51b81526004016110ec906183db565b86600111158015612659575060098711155b6126945760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964207479706560a01b60448201526064016110ec565b61269e8887614e48565b6126da5760405162461bcd60e51b815260206004820152600d60248201526c496e76616c69642076616c756560981b60448201526064016110ec565b6075546126e89060016185fd565b92506126f26149ed565b6001600160a01b0316634a57823e84600587338d8d8d6040518863ffffffff1660e01b815260040161272a9796959493929190618531565b602060405180830381600087803b15801561274457600080fd5b505af1158015612758573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061277c9190617d25565b506127878386614a08565b5060758290556001600160a01b0316600090815260ab6020526040902042905595945050505050565b6001600160a01b03166000908152606b6020526040902054151590565b6000818152607160205260408120600381015481546060938493849391929091600183019160028401919084906128039061867a565b80601f016020809104026020016040519081016040528092919081815260200182805461282f9061867a565b801561287c5780601f106128515761010080835404028352916020019161287c565b820191906000526020600020905b81548152906001019060200180831161285f57829003601f168201915b5050505050935082805461288f9061867a565b80601f01602080910402602001604051908101604052809291908181526020018280546128bb9061867a565b80156129085780601f106128dd57610100808354040283529160200191612908565b820191906000526020600020905b8154815290600101906020018083116128eb57829003601f168201915b5050505050925081805461291b9061867a565b80601f01602080910402602001604051908101604052809291908181526020018280546129479061867a565b80156129945780601f1061296957610100808354040283529160200191612994565b820191906000526020600020905b81548152906001019060200180831161297757829003601f168201915b5050505050915093509350935093509193509193565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612a4a5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016110ec565b506000805160206186ff83398151915290565b6000612a67614bee565b6001600160a01b031663737c59b86040518163ffffffff1660e01b815260040160206040518083038186803b1580156117d757600080fd5b612aa7614947565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b158015612ae857600080fd5b505afa158015612afc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b209190617d08565b612b3c5760405162461bcd60e51b81526004016110ec906182c8565b612b44611744565b15612b875760405162461bcd60e51b8152602060048201526013602482015272159bdd1a5b99c81a5cc81b9bdd08195b991959606a1b60448201526064016110ec565b607754612b95906004614ed6565b6000607755565b6001600160a01b0381166000908152606d602052604081205415612bbe575090565b6001600160a01b03821660009081526068602052604090205415612c0657506001600160a01b038082166000908152606860209081526040808320548352606c909152902054165b919050565b612c13614a73565b612c1d6000614f14565b565b6000612c29614947565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b158015612c6a57600080fd5b505afa158015612c7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ca29190617d08565b612cbe5760405162461bcd60e51b81526004016110ec906182c8565b6000612cc933612b9c565b60aa546001600160a01b038216600090815260ab602052604090205491925090612cf39042618637565b1015612d115760405162461bcd60e51b81526004016110ec9061845d565b6000612d1c33612b9c565b9050612d26612a5d565b612d2f82614967565b11158015612d4c5750612d406145bf565b612d4982614967565b10155b612d685760405162461bcd60e51b81526004016110ec906183db565b6001600160a01b038916612db05760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b60448201526064016110ec565b6001600160a01b0389166000908152606d6020526040902054612e025760405162461bcd60e51b815260206004820152600a6024820152692737b716b6b2b6b132b960b11b60448201526064016110ec565b6001612e0d60695490565b11612e5a5760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f742072656d6f7665206120736f6c65206d656d626572000000000060448201526064016110ec565b87612e648a614967565b1015612ec55760405162461bcd60e51b815260206004820152602a60248201527f496e73756666696369656e742062616c616e636520746861742063616e206265604482015269103ab73637b1b5b2b21760b11b60648201526084016110ec565b607554612ed39060016185fd565b604080516101408101825260008082526020808301829052828401829052835182815280820185526060840152835182815280820185526080840152835182815290810190935260a082019290925260c081019190915260e081018a905261010081018990526101208101889052909350612f52846002338d85614d67565b612f5c848a614e0a565b612f668489614a08565b612f71848787614f66565b505060758290556001600160a01b0316600090815260ab602052604090204290559695505050505050565b60008054610100900460ff1615808015612fbd5750600054600160ff909116105b80612fd75750303b158015612fd7575060005460ff166001145b612ff35760405162461bcd60e51b81526004016110ec9061827a565b6000805460ff191660011790558015613016576000805461ff0019166101001790555b61301e614c06565b613026614c35565b600083905061309f816001600160a01b031663738fdd1a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561306757600080fd5b505afa15801561307b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb791906179ac565b4360668190555061311a816001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156130e257600080fd5b505afa1580156130f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061101491906179ac565b60015b816001600160a01b031663d965ea006040518163ffffffff1660e01b815260040160206040518083038186803b15801561315657600080fd5b505afa15801561316a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061318e9190617d25565b81116136c85760405163ab3545e560e01b8152600481018290526001600160a01b0383169063ab3545e59060240160206040518083038186803b1580156131d457600080fd5b505afa1580156131e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061320c91906179ac565b6000828152606c6020908152604080832080546001600160a01b0319166001600160a01b039586169081179091558352606d909152908190208390555163341effc360e21b8152600481018390529083169063d07bff0c9060240160206040518083038186803b15801561327f57600080fd5b505afa158015613293573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132b791906179ac565b600082815260676020908152604080832080546001600160a01b0319166001600160a01b03958616908117909155835260689091529081902083905551631c4b774b60e01b81526004810183905290831690631c4b774b9060240160206040518083038186803b15801561332a57600080fd5b505afa15801561333e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061336291906179ac565b6000828152606a6020908152604080832080546001600160a01b0319166001600160a01b03959095169485179055928252606b8152828220849055606984905582516080810184526060808252918101829052928301819052820152604051634f0f4aa960e01b8152600481018390526001600160a01b03841690634f0f4aa99060240160006040518083038186803b1580156133fe57600080fd5b505afa158015613412573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261343a9190810190617e01565b60608501819052604085018290526020850183905283855261345e93929190614c79565b6134aa5760405162461bcd60e51b815260206004820152601760248201527f6e6f646520696e666f206973206475706c69636174656400000000000000000060448201526064016110ec565b6001606e82600001516040516134c091906180b1565b90815260405160209181900382018120805460ff191693151593909317909255820151600191606f916134f2916180b1565b908152602001604051809103902060006101000a81548160ff021916908315150217905550600160706000836040015184606001516040516020016135389291906180cd565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff19169415159490941790935584835260718152912082518051849361358e92849291019061767e565b5060208281015180516135a7926001850192019061767e565b50604082015180516135c391600284019160209091019061767e565b50606091909101516003909101556000828152606c6020818152604080842080546001600160a01b039081168652607284528286208890558786528154607385529583902080546001600160a01b031916968216969096179095556074879055929091529054905163139d9dd360e01b815290821660048201529084169063139d9dd39060240160206040518083038186803b15801561366257600080fd5b505afa158015613676573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061369a9190617d25565b6000838152606c60209081526040808320546001600160a01b0316835260ab9091529020555060010161311d565b50806001600160a01b0316633310569c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561370257600080fd5b505afa158015613716573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061373a9190617d25565b60aa81905550806001600160a01b031663d6f9cfce6040518163ffffffff1660e01b815260040160206040518083038186803b15801561377957600080fd5b505afa15801561378d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137b19190617d25565b607581905550806001600160a01b031663e9523fb56040518163ffffffff1660e01b815260040160206040518083038186803b1580156137f057600080fd5b505afa158015613804573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138289190617d25565b607681905550806001600160a01b031663de09b3776040518163ffffffff1660e01b815260040160206040518083038186803b15801561386757600080fd5b505afa15801561387b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061389f9190617d25565b607755506000915080156138db576000805461ff00191690556040516001815260008051602061871f8339815191529060200160405180910390a15b50919050565b6001600160a01b0381166000908152606d602052604081205415158061391e57506001600160a01b03821660009081526068602052604090205415155b92915050565b600061392e614947565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b15801561396f57600080fd5b505afa158015613983573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139a79190617d08565b6139c35760405162461bcd60e51b81526004016110ec906182c8565b60006139ce33612b9c565b60aa546001600160a01b038216600090815260ab6020526040902054919250906139f89042618637565b1015613a165760405162461bcd60e51b81526004016110ec9061845d565b6000613a2133612b9c565b9050613a2b612a5d565b613a3482614967565b11158015613a515750613a456145bf565b613a4e82614967565b10155b613a6d5760405162461bcd60e51b81526004016110ec906183db565b602087015187906001600160a01b0316613ab95760405162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b2103b37ba32b960991b60448201526064016110ec565b600081606001515111613b025760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964206e6f6465206e616d6560781b60448201526064016110ec565b60008160a001515111613b495760405162461bcd60e51b815260206004820152600f60248201526e0496e76616c6964206e6f646520495608c1b60448201526064016110ec565b60008160c0015111613b915760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a59081b9bd919481c1bdc9d607a1b60448201526064016110ec565b600081608001515111613bdb5760405162461bcd60e51b8152602060048201526012602482015271496e76616c6964206e6f646520656e6f646560701b60448201526064016110ec565b60008161010001515111613c205760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964206d656d6f60a01b60448201526064016110ec565b600081610120015111613c685760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b210323ab930ba34b7b760811b60448201526064016110ec565b613c706145bf565b8160e0015110158015613c8e5750613c86612a5d565b8160e0015111155b613cd05760405162461bcd60e51b8152602060048201526013602482015272125b9d985b1a59081b1bd8dac8105b5bdd5b9d606a1b60448201526064016110ec565b6001600160a01b038716613d1c5760405162461bcd60e51b8152602060048201526013602482015272496e76616c6964206f6c64204164647265737360681b60448201526064016110ec565b6001600160a01b0387166000908152606d6020526040902054613d6e5760405162461bcd60e51b815260206004820152600a6024820152692737b716b6b2b6b132b960b11b60448201526064016110ec565b6020808901516001600160a01b038981166000908152606d8452604080822054825260679094529290922054821691161480613dbf5750866001600160a01b031688602001516001600160a01b0316145b80613de95750613dd288602001516138e1565b158015613de95750613de788602001516127b0565b155b8015613e6b57506040808901516001600160a01b038981166000908152606d6020908152848220548252606a90529290922054821691161480613e415750866001600160a01b031688604001516001600160a01b0316145b80613e6b5750613e5488604001516138e1565b158015613e6b5750613e6988604001516127b0565b155b613eaa5760405162461bcd60e51b815260206004820152601060248201526f20b63932b0b23c90309036b2b6b132b960811b60448201526064016110ec565b336001600160a01b038816148015613ece575087516001600160a01b038881169116145b15613f235785158015613edf575084155b613f1e5760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a59081c1c9bdc1bdcd85b60821b60448201526064016110ec565b613fd2565b87516001600160a01b03888116911614613fd257613f3f6145bf565b613f4986886185fd565b1115613fd25760405162461bcd60e51b815260206004820152604c60248201527f496e76616c696420616d6f756e743a2028756e6c6f636b416d6f756e74202b2060448201527f736c617368696e6729206d75737420626520657175616c206f72206c6f77207460648201526b6f206d696e5374616b696e6760a01b608482015260a4016110ec565b607554613fe09060016185fd565b9350613ff0846003338a8c614d67565b613ffe848960e00151614e0a565b61400d84896101000151614a08565b614018848787614f66565b6075849055336001600160a01b038816148015614041575087516001600160a01b038881169116145b1561407b57600061405185614b64565b9250505061406b8542834261406691906185fd565b614fd5565b61407985600360018061501a565b505b50506001600160a01b0316600090815260ab60205260409020429055949350505050565b600054610100900460ff16158080156140bf5750600054600160ff909116105b806140d95750303b1580156140d9575060005460ff166001145b6140f55760405162461bcd60e51b81526004016110ec9061827a565b6000805460ff191660011790558015614118576000805461ff0019166101001790555b614120614c06565b614128614c35565b61413187614517565b6141396145bf565b861015801561414f57508561414c612a5d565b10155b61416b5760405162461bcd60e51b81526004016110ec9061836b565b6000614175614c64565b6040516325d998bb60e01b815233600482015290915087906001600160a01b038316906325d998bb9060240160206040518083038186803b1580156141b957600080fd5b505afa1580156141cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141f19190617d25565b10156142365760405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e74207374616b696e6760601b60448201526064016110ec565b60405163282d3fdf60e01b8152336004820152602481018890526001600160a01b0382169063282d3fdf90604401600060405180830381600087803b15801561427e57600080fd5b505af1158015614292573d6000803e3d6000fd5b5050600160698190557f6bee784efeb983674392298ab585b22866bedf00ebb0eea949d1e66f3f50e71d8054336001600160a01b0319918216811790925560008281526068602090815260408083208690557ff585789965ba69220d5ce3dc1b444eb22ff546f2650694fef8fafe9c26560af98054851686179055606b82528083208690557fdcf345d7f6a8deb7427d0fee62009fa15100353a1c666b51bb5387b25addcfa98054909416909417909255606d825291822083905560748390559190526071815288517f169c6be1b0e6ab5de76b532e587a77340130ac65c5591db02be822dcf1dc0ed6935061438d925083918a019061767e565b5085516143a3906001830190602089019061767e565b5084516143b9906002830190602088019061767e565b508381600301819055506001606e886040516143d591906180b1565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f906144079089906180b1565b908152602001604051809103902060006101000a81548160ff02191690831515021790555060016070600087876040516020016144459291906180cd565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff1916941515949094179093556074543380855260728352838520829055908452607390915281832080546001600160a01b03191682179055436066559051909182917f6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba69190a35050801561450e576000805461ff00191690556040516001815260008051602061871f8339815191529060200160405180910390a15b50505050505050565b61451f614a73565b6001600160a01b0381166145755760405162461bcd60e51b815260206004820152601a60248201527f416464726573732073686f756c64206265206e6f6e2d7a65726f00000000000060448201526064016110ec565b606580546001600160a01b0319166001600160a01b0383169081179091556040517f278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd90600090a250565b60006145c9614bee565b6001600160a01b031663076cd77f6040518163ffffffff1660e01b815260040160206040518083038186803b1580156117d757600080fd5b600260785414156146545760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016110ec565b6002607855614661614947565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b1580156146a257600080fd5b505afa1580156146b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906146da9190617d08565b6146f65760405162461bcd60e51b81526004016110ec906182c8565b600061470133612b9c565b905061470b612a5d565b61471482614967565b1115801561473157506147256145bf565b61472e82614967565b10155b61474d5760405162461bcd60e51b81526004016110ec906183db565b614755611744565b61478b5760405162461bcd60e51b8152602060048201526007602482015266115e1c1a5c995960ca1b60448201526064016110ec565b6000614796846150a9565b90506147a284846151be565b6000806147ae86615298565b925092505060006147be61138990565b905080831015806147cf5750808210155b806147e457506147df82846185fd565b612710145b156147f8576147f88785848611600061501a565b505060016078555050505050565b6000614810614bee565b6001600160a01b0316631b27e01b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156117d757600080fd5b614850614a73565b610e1081106148995760405162461bcd60e51b81526020600482015260156024820152746e6577506572696f6420697320746f6f206c6f6e6760581b60448201526064016110ec565b60aa8190556040518181527f17c6f1d1ce638844b664872f5c6eecb7d150ec0c41187d7f85826a656ee7946f90602001611739565b6148d6614a73565b6001600160a01b03811661493b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016110ec565b61494481614f14565b50565b600061180f7111dbdd995c9b985b98d950dbdb9d1c9858dd60721b6152d2565b6000614971614c64565b604051632c9aab9b60e11b81526001600160a01b03848116600483015291909116906359355736906024015b60206040518083038186803b1580156149b557600080fd5b505afa1580156149c9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061391e9190617d25565b600061180f6c42616c6c6f7453746f7261676560981b6152d2565b614a106149ed565b6001600160a01b031663bce0dbc183836040518363ffffffff1660e01b8152600401614a3d929190618114565b600060405180830381600087803b158015614a5757600080fd5b505af1158015614a6b573d6000803e3d6000fd5b505050505050565b6033546001600160a01b03163314612c1d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016110ec565b6000806000614ada6149ed565b6001600160a01b031663688ca5b2856040518263ffffffff1660e01b8152600401614b0791815260200190565b60606040518083038186803b158015614b1f57600080fd5b505afa158015614b33573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614b579190617f8f565b9250925092509193909250565b6000806000614b716149ed565b6001600160a01b03166309970688856040518263ffffffff1660e01b8152600401614b9e91815260200190565b60606040518083038186803b158015614bb657600080fd5b505afa158015614bca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614b579190617fbd565b600061180f69456e7653746f7261676560b01b6152d2565b600054610100900460ff16614c2d5760405162461bcd60e51b81526004016110ec90618412565b612c1d61534f565b600054610100900460ff16614c5c5760405162461bcd60e51b81526004016110ec90618412565b612c1d61537d565b600061180f665374616b696e6760c81b6152d2565b604051600190606f90614c8d9086906180b1565b9081526040519081900360200190205460ff1615614ca9575060005b606e85604051614cb991906180b1565b9081526040519081900360200190205460ff1615614cd5575060005b60008383604051602001614cea9291906180cd565b60408051601f1981840301815291815281516020928301206000818152607090935291205490915060ff1615614d1f57600091505b50949350505050565b614d30614c64565b60405163282d3fdf60e01b81526001600160a01b03848116600483015260248201849052919091169063282d3fdf90604401614a3d565b614d6f6149ed565b6001600160a01b031663daacbb95868684610120015187878760000151886020015189604001518a606001518b608001518c60a001518d60c001516040518d63ffffffff1660e01b8152600401614dd19c9b9a99989796959493929190618494565b600060405180830381600087803b158015614deb57600080fd5b505af1158015614dff573d6000803e3d6000fd5b505050505050505050565b614e126149ed565b604051633968764960e11b815260048101849052602481018390526001600160a01b0391909116906372d0ec9290604401614a3d565b6000614e52614bee565b6001600160a01b0316639801bff984846040518363ffffffff1660e01b8152600401614e7f929190618114565b60206040518083038186803b158015614e9757600080fd5b505afa158015614eab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ecf9190617d08565b9392505050565b614ede6149ed565b60405163548f2cdd60e11b815260048101849052602481018390526001600160a01b03919091169063a91e59ba90604401614a3d565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b614f6e6149ed565b604051632264085960e01b81526004810185905260248101849052604481018390526001600160a01b0391909116906322640859906064015b600060405180830381600087803b158015614fc157600080fd5b505af115801561450e573d6000803e3d6000fd5b614fdd6149ed565b60405163605b78c360e11b81526004810185905260248101849052604481018390526001600160a01b03919091169063c0b6f18690606401614fa7565b6004821561508d57506003600184141561504457615037856153ad565b61503f575060045b61508d565b60028414156150565761503f856158e3565b6003841415615069576150378583615f83565b600484141561507b5761503f856164eb565b600584141561508d5761508d856165bc565b6150978582614ed6565b816150a25760006077555b5050505050565b60008060006150b784614acd565b509092509050600181141561515657607754156150e65760405162461bcd60e51b81526004016110ec90618398565b60006150f185614b64565b925050506150fd611795565b8110156151205761511b8542615111611795565b61406690426185fd565b61514b565b80615129614806565b101561513c5761511b8542615111614806565b61514b854261406684826185fd565b5060778490556151b7565b60028114156151855760775484146151805760405162461bcd60e51b81526004016110ec90618398565b6151b7565b60405162461bcd60e51b8152602060048201526007602482015266115e1c1a5c995960ca1b60448201526064016110ec565b5092915050565b600060765460016151cf91906185fd565b905060006151dc33612b9c565b905060006151e960695490565b6151f590612710618615565b9050600084615205576002615208565b60015b90506152126149ed565b6040516325918ae760e21b815260048101869052602481018890526001600160a01b038581166044830152606482018490526084820185905291909116906396462b9c9060a401600060405180830381600087803b15801561527357600080fd5b505af1158015615287573d6000803e3d6000fd5b505050607694909455505050505050565b60008060006152a56149ed565b6001600160a01b03166356ba988e856040518263ffffffff1660e01b8152600401614b9e91815260200190565b606554604051630d2020dd60e01b8152600481018390526000916001600160a01b031690630d2020dd9060240160206040518083038186803b15801561531757600080fd5b505afa15801561532b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061391e91906179ac565b600054610100900460ff166153765760405162461bcd60e51b81526004016110ec90618412565b6001607855565b600054610100900460ff166153a45760405162461bcd60e51b81526004016110ec90618412565b612c1d33614f14565b60006153ba82600161670f565b6000806000806000806000806153cf8a61682b565b98509850985098509850985098509850506153e9886138e1565b15615447578960008051602061873f83398151915260405161542f9060208082526010908201526f20b63932b0b23c90309036b2b6b132b960811b604082015260600190565b60405180910390a25060009998505050505050505050565b615450866127b0565b15615498578960008051602061873f83398151915260405161542f9060208082526012908201527120b63932b0b23c9030903932bbb0b93232b960711b604082015260600190565b6154a06145bf565b8110806154b35750806154b1612a5d565b105b156154d4578960008051602061873f83398151915260405161542f9061836b565b806154de896168e1565b1015615500578960008051602061873f83398151915260405161542f906181ea565b866001600160a01b0316886001600160a01b0316141580156155345750856001600160a01b0316886001600160a01b031614155b15615580578960008051602061873f83398151915260405161542f90602080825260169082015275496e76616c6964206d656d626572206164647265737360501b604082015260600190565b61558c85858585614c79565b6155ac578960008051602061873f83398151915260405161542f906182ef565b6155b68882614d28565b600060695460016155c791906185fd565b9050876067600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606860008a6001600160a01b03166001600160a01b031681526020019081526020016000208190555086606a600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606b6000896001600160a01b03166001600160a01b031681526020019081526020016000208190555088606c600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606d60008b6001600160a01b03166001600160a01b03168152602001908152602001600020819055506000607454600161570091906185fd565b6000818152607160209081526040909120895192935091615726918391908b019061767e565b50865161573c90600183019060208a019061767e565b508551615752906002830190602089019061767e565b508481600301819055506001606e8960405161576e91906180b1565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f906157a0908a906180b1565b908152602001604051809103902060006101000a81548160ff02191690831515021790555060016070600088886040516020016157de9291906180cd565b60405160208183030381529060405280519060200120815260200190815260200160002060006101000a81548160ff0219169083151502179055508a6073600084815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081607260008d6001600160a01b03166001600160a01b0316815260200190815260200160002081905550826069819055508160748190555043606681905550896001600160a01b03168b6001600160a01b03167f6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba660405160405180910390a35060019c9b505050505050505050505050565b6158ee81600261670f565b60006158f98261682b565b5050505050505050905061590c816138e1565b615961578160008051602061873f833981519152604051615955906020808252601490820152732737ba1030b63932b0b23c90309036b2b6b132b960611b604082015260600190565b60405180910390a25050565b6001600160a01b0381166000908152606d6020526040902054158061599c57506001600160a01b038116600090815260726020526040902054155b156159e6578160008051602061873f833981519152604051615955906020808252601490820152732737ba1030b63932b0b23c90309036b2b6b132b960611b604082015260600190565b60016159f160695490565b11615a44578160008051602061873f833981519152604051615955906020808252601b908201527f43616e6e6f742072656d6f7665206120736f6c65206d656d6265720000000000604082015260600190565b6001600160a01b038082166000908152606d60209081526040808320548084526067835281842054606a9093529220546069549293918216929116908314615bf8576069546000908152606c602090815260408083205486845281842080546001600160a01b0319166001600160a01b03928316908117909155808552606d84528285208890559085168452606b9092529091205480615b1d5760405162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e4caeec2e4c840d2dcc8caf60631b60448201526064016110ec565b6069546000908152606a602090815260408083205484845281842080546001600160a01b0319166001600160a01b03928316908117909155808552606b8452828520869055908816845260689092529091205480615bb35760405162461bcd60e51b8152602060048201526013602482015272092dcecc2d8d2c840ecdee8cae440d2dcc8caf606b1b60448201526064016110ec565b60695460009081526067602090815260408083205484845281842080546001600160a01b0319166001600160a01b039092169182179055835260689091529020555050505b606980546000908152606c6020908152604080832080546001600160a01b03199081169091556001600160a01b038981168552606d845282852085905585548552606a84528285208054831690558681168552606b84528285208590558554855260678452828520805490921690915586168352606890915281205554615c8190600190618637565b6069556001600160a01b03841660009081526072602052604090205480615cdf5760405162461bcd60e51b8152602060048201526012602482015271092dcecc2d8d2c840dcdec8ca40d2dcc8caf60731b60448201526064016110ec565b6000818152607160205260408082209051909190606f90615d049060018501906180ef565b908152604051908190036020018120805492151560ff1990931692909217909155600090606e90615d369084906180ef565b90815260405160209181900382018120805460ff19169315159390931790925560038301546000926070928492615d72926002880192016180fb565b60408051808303601f19018152918152815160209283012083529082019290925201600020805460ff19169115159190911790556074548214615ea557607454600090815260736020908152604080832054607190925290912080546001600160a01b0390921691839190615de69061867a565b615df1929190617702565b5060745460009081526071602052604090206001908101805491840191615e179061867a565b615e22929190617702565b5060745460009081526071602052604090206002908101805491840191615e489061867a565b615e53929190617702565b506074546000908152607160209081526040808320600390810154908601558583526073825280832080546001600160a01b039095166001600160a01b03199095168517905592825260729052208290555b60748054600090815260736020908152604080832080546001600160a01b03191690556001600160a01b038a16835260728252808320839055925482526071905290812090615ef4828261777d565b615f0260018301600061777d565b615f1060028301600061777d565b600382016000905550506001607454615f299190618637565b60745543606655615f3a878761691b565b836001600160a01b0316866001600160a01b03167faa91016c21c52c58ac64f23f71bbe75becc9ada603e18ee671d09ff15492d1c160405160405180910390a350505050505050565b600081615f9557615f9583600361670f565b6000806000806000806000806000615fac8c61682b565b985098509850985098509850985098509850615fc7896138e1565b616001578b60008051602061873f833981519152604051615fe7906181b3565b60405180910390a26000995050505050505050505061391e565b6160148c8c8b8b8b8b8b8b8b8b8b616af8565b61602a576000995050505050505050505061391e565b6001600160a01b03808a166000818152606d6020526040902054918a1614616099576000818152606c6020908152604080832080546001600160a01b0319166001600160a01b038e81169182179092558452606d909252808320849055908c1682528120556160998983614d28565b6001600160a01b038a166000908152607260209081526040808320548084526071909252808320905191929091606e906160d49084906180ef565b908152604051908190036020018120805492151560ff1990931692909217909155600090606f906161099060018501906180ef565b90815260405160209181900382018120805460ff19169315159390931790925560038301546000926070928492616145926002880192016180fb565b60408051808303601f1901815291815281516020928301208352828201939093529101600020805460ff191692151592909217909155885161618c918391908b019061767e565b5086516161a290600183019060208a019061767e565b5085516161b8906002830190602089019061767e565b506003810185905543606655604051600190606e906161d8908b906180b1565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f9061620a908a906180b1565b908152602001604051809103902060006101000a81548160ff02191690831515021790555060016070600088886040516020016162489291906180cd565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff191694151594909417909355858352606a90529020546001600160a01b039081169150891681146162e4576000838152606a6020908152604080832080546001600160a01b0319166001600160a01b038e81169182179092558452606b90925280832086905590831682528120555b506000828152606760205260409020546001600160a01b03908116908a16811461639657896067600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606860008c6001600160a01b03166001600160a01b0316815260200190815260200160002081905550600060686000836001600160a01b03166001600160a01b03168152602001908152602001600020819055505b50896001600160a01b03168b6001600160a01b03161461649657896073600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080607260008c6001600160a01b03166001600160a01b03168152602001908152602001600020819055506000607260008d6001600160a01b03166001600160a01b03168152602001908152602001600020819055506164478e8c61691b565b886001600160a01b03168a6001600160a01b03168c6001600160a01b03167f15f4d750630db473a85edd9d47c500527a2648cc5e676f39645e52790cf07be060405160405180910390a46164d7565b896001600160a01b03168b6001600160a01b03167f1feee1b4fcb797c62645da41c5c6edd5f91d4291de0054da625c42b823594c1f60405160405180910390a35b5060019d9c50505050505050505050505050565b6164f681600461670f565b60006165006149ed565b6001600160a01b0316637efa9ae3836040518263ffffffff1660e01b815260040161652d91815260200190565b60206040518083038186803b15801561654557600080fd5b505afa158015616559573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061657d91906179ac565b90506001600160a01b038116156165b857616597816170bc565b604080516000808252602082019092526165b391839190617159565b436066555b5050565b6165c781600561670f565b60008060006165d46149ed565b6001600160a01b0316631d940da2856040518263ffffffff1660e01b815260040161660191815260200190565b60006040518083038186803b15801561661957600080fd5b505afa15801561662d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526166559190810190617d3e565b9250925092506000616665614bee565b6040516388c2801960e01b81529091506001600160a01b038216906388c28019906166969087908690600401618114565b600060405180830381600087803b1580156166b057600080fd5b505af11580156166c4573d6000803e3d6000fd5b50504360665550506040517f701c16c2519cdb79aaac423a84733590e3510d9552055b6ad6908f0ab12b6c29906167009086908690869061812d565b60405180910390a15050505050565b60008061671b84614acd565b50915091508282146167655760405162461bcd60e51b8152602060048201526013602482015272496e76616c696420766f74696e67207479706560681b60448201526064016110ec565b600281146167ac5760405162461bcd60e51b8152602060048201526014602482015273496e76616c696420766f74696e6720737461746560601b60448201526064016110ec565b6000806167b886615298565b92509250506167c661138990565b821015806167d657506113898110155b806167eb57506167e681836185fd565b612710145b614a6b5760405162461bcd60e51b8152602060048201526011602482015270139bdd081e595d08199a5b985b1a5e9959607a1b60448201526064016110ec565b60008060008060608060606000806168416149ed565b6001600160a01b03166373df4e018b6040518263ffffffff1660e01b815260040161686e91815260200190565b60006040518083038186803b15801561688657600080fd5b505afa15801561689a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526168c291908101906179c9565b9850985098509850985098509850985098509193959799909294969850565b60006168eb614c64565b6040516325d998bb60e01b81526001600160a01b03848116600483015291909116906325d998bb9060240161499d565b6000806169278461724f565b915091506169336145bf565b61693d82846185fd565b11156169cd5760405162461bcd60e51b815260206004820152605360248201527f6d696e5374616b696e672076616c7565206d757374206265206772656174657260448201527f207468616e206f7220657175616c20746f207468652073756d206f6620756e6c6064820152726f636b416d6f756e742c20736c617368696e6760681b608482015260a4016110ec565b60006169d7614c64565b604051632c9aab9b60e11b81526001600160a01b03868116600483015291925060009183169063593557369060240160206040518083038186803b158015616a1e57600080fd5b505afa158015616a32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616a569190617d25565b90506000616a626145bf565b616a6c9083618637565b905084821115616aee57616a8086866172df565b604051634954a85b60e11b81526001600160a01b03878116600483015260248201869052604482018390528416906392a950b690606401600060405180830381600087803b158015616ad157600080fd5b505af1158015616ae5573d6000803e3d6000fd5b5050505061450e565b61450e86836172df565b60008a616b0a57616b0a8c600361670f565b616b138a6138e1565b616b43578b60008051602061873f833981519152604051616b33906181b3565b60405180910390a25060006170ad565b6001600160a01b038a166000908152606d60205260409020541580616b7e57506001600160a01b038a16600090815260726020526040902054155b15616b9f578b60008051602061873f833981519152604051616b33906181b3565b6001600160a01b03808b166000818152606d6020526040902054918b1614616d1657616bca8a6138e1565b15616c2f578c60008051602061873f833981519152604051616c1d906020808252601f908201527f6e6577206164647265737320697320616c72656164792061206d656d62657200604082015260600190565b60405180910390a260009150506170ad565b886001600160a01b03168a6001600160a01b031614158015616c635750876001600160a01b03168a6001600160a01b031614155b15616cae578c60008051602061873f833981519152604051616c1d90602080825260159082015274496e76616c696420766f746572206164647265737360581b604082015260600190565b616cb66145bf565b831080616cc9575082616cc7612a5d565b105b15616cea578c60008051602061873f833981519152604051616c1d9061836b565b82616cf48b6168e1565b1015616d16578c60008051602061873f833981519152604051616c1d906181ea565b6001600160a01b038b166000908152607260209081526040808320548084526071909252808320815160808101909252805492939282908290616d589061867a565b80601f0160208091040260200160405190810160405280929190818152602001828054616d849061867a565b8015616dd15780601f10616da657610100808354040283529160200191616dd1565b820191906000526020600020905b815481529060010190602001808311616db457829003601f168201915b50505050508152602001600182018054616dea9061867a565b80601f0160208091040260200160405190810160405280929190818152602001828054616e169061867a565b8015616e635780601f10616e3857610100808354040283529160200191616e63565b820191906000526020600020905b815481529060010190602001808311616e4657829003601f168201915b50505050508152602001600282018054616e7c9061867a565b80601f0160208091040260200160405190810160405280929190818152602001828054616ea89061867a565b8015616ef55780601f10616eca57610100808354040283529160200191616ef5565b820191906000526020600020905b815481529060010190602001808311616ed857829003601f168201915b505050505081526020016003820154815250509050616f17898989898561731e565b616f4b578e60008051602061873f833981519152604051616f37906182ef565b60405180910390a2600093505050506170ad565b506000828152606a60205260409020546001600160a01b03908116908d8116908b1614801590616f8d5750896001600160a01b0316816001600160a01b031614155b8015616fac5750616f9d8a6138e1565b80616fac5750616fac8a6127b0565b15616ff8578e60008051602061873f833981519152604051616f3790602080825260169082015275496e76616c696420726577617264206164647265737360501b604082015260600190565b506000828152606760205260409020546001600160a01b03908116908d8116908c161480159061703a57508a6001600160a01b0316816001600160a01b031614155b8015617059575061704a8b6138e1565b8061705957506170598b6127b0565b156170a5578e60008051602061873f833981519152604051616f3790602080825260169082015275496e76616c696420766f74657273206164647265737360501b604082015260600190565b506001925050505b9b9a5050505050505050505050565b6170c4614947565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b15801561710557600080fd5b505afa158015617119573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061713d9190617d08565b6149445760405162461bcd60e51b81526004016110ec906182c8565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156171915761718c83617447565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156171ca57600080fd5b505afa9250505080156171fa575060408051601f3d908101601f191682019092526171f791810190617d25565b60015b6172165760405162461bcd60e51b81526004016110ec9061831d565b6000805160206186ff83398151915281146172435760405162461bcd60e51b81526004016110ec90618231565b5061718c8383836174e3565b60008061725a6149ed565b6001600160a01b0316638c7be692846040518263ffffffff1660e01b815260040161728791815260200190565b604080518083038186803b15801561729e57600080fd5b505afa1580156172b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906172d69190617f6b565b91509150915091565b6172e7614c64565b604051637eee288d60e01b81526001600160a01b038481166004830152602482018490529190911690637eee288d90604401614a3d565b6000600190508480519060200120826020015180519060200120141580156173655750606f8560405161735191906180b1565b9081526040519081900360200190205460ff165b1561736e575060005b8580519060200120826000015180519060200120141580156173af5750606e8660405161739b91906180b1565b9081526040519081900360200190205460ff165b156173b8575060005b600084846040516020016173cd9291906180cd565b60405160208183030381529060405280519060200120905080836040015184606001516040516020016174019291906180cd565b6040516020818303038152906040528051906020012014158015617433575060008181526070602052604090205460ff165b1561743d57600091505b5095945050505050565b6001600160a01b0381163b6174b45760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016110ec565b6000805160206186ff83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6174ec83617508565b6000825111806174f95750805b1561718c57611e9d8383617548565b61751181617447565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060614ecf838360405180606001604052806027815260200161875f6027913960606001600160a01b0384163b6175d05760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016110ec565b600080856001600160a01b0316856040516175eb91906180b1565b600060405180830381855af49150503d8060008114617626576040519150601f19603f3d011682016040523d82523d6000602084013e61762b565b606091505b509150915061763b828286617645565b9695505050505050565b60608315617654575081614ecf565b8251156176645782518084602001fd5b8160405162461bcd60e51b81526004016110ec91906181a0565b82805461768a9061867a565b90600052602060002090601f0160209004810192826176ac57600085556176f2565b82601f106176c557805160ff19168380011785556176f2565b828001600101855582156176f2579182015b828111156176f25782518255916020019190600101906176d7565b506176fe9291506177b3565b5090565b82805461770e9061867a565b90600052602060002090601f01602090048101928261773057600085556176f2565b82601f1061774157805485556176f2565b828001600101855582156176f257600052602060002091601f016020900482015b828111156176f2578254825591600101919060010190617762565b5080546177899061867a565b6000825580601f10617799575050565b601f01602090049060005260206000209081019061494491905b5b808211156176fe57600081556001016177b4565b8035612c06816186db565b8051612c06816186db565b600082601f8301126177ef57600080fd5b81356178026177fd826185d6565b6185a6565b81815284602083860101111561781757600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261784557600080fd5b81516178536177fd826185d6565b81815284602083860101111561786857600080fd5b61787982602083016020870161864e565b949350505050565b6000610140828403121561789457600080fd5b61789c61857d565b90506178a7826177c8565b81526178b5602083016177c8565b60208201526178c6604083016177c8565b604082015260608201356001600160401b03808211156178e557600080fd5b6178f1858386016177de565b6060840152608084013591508082111561790a57600080fd5b617916858386016177de565b608084015260a084013591508082111561792f57600080fd5b61793b858386016177de565b60a084015260c084013560c084015260e084013560e08401526101009150818401358181111561796a57600080fd5b617976868287016177de565b8385015250505061012080830135818301525092915050565b6000602082840312156179a157600080fd5b8135614ecf816186db565b6000602082840312156179be57600080fd5b8151614ecf816186db565b60008060008060008060008060006101208a8c0312156179e857600080fd5b6179f18a6177d3565b98506179ff60208b016177d3565b9750617a0d60408b016177d3565b9650617a1b60608b016177d3565b955060808a01516001600160401b0380821115617a3757600080fd5b617a438d838e01617834565b965060a08c0151915080821115617a5957600080fd5b617a658d838e01617834565b955060c08c0151915080821115617a7b57600080fd5b50617a888c828d01617834565b93505060e08a015191506101008a015190509295985092959850929598565b60008060408385031215617aba57600080fd5b8235617ac5816186db565b915060208301356001600160401b03811115617ae057600080fd5b617aec858286016177de565b9150509250929050565b600080600060608486031215617b0b57600080fd5b8335617b16816186db565b925060208401356001600160401b03811115617b3157600080fd5b617b3d868287016177de565b925050604084013590509250925092565b600080600060608486031215617b6357600080fd5b8335617b6e816186db565b9250602084013591506040840135617b85816186db565b809150509250925092565b600080600060608486031215617ba557600080fd5b8335617bb0816186db565b92506020840135915060408401356001600160401b03811115617bd257600080fd5b617bde868287016177de565b9150509250925092565b60008060008060008060c08789031215617c0157600080fd5b8635617c0c816186db565b95506020870135945060408701356001600160401b0380821115617c2f57600080fd5b617c3b8a838b016177de565b95506060890135915080821115617c5157600080fd5b617c5d8a838b016177de565b94506080890135915080821115617c7357600080fd5b50617c8089828a016177de565b92505060a087013590509295509295509295565b60008060008060008060c08789031215617cad57600080fd5b8635617cb8816186db565b95506020870135945060408701356001600160401b03811115617cda57600080fd5b617ce689828a016177de565b945050606087013592506080870135915060a087013590509295509295509295565b600060208284031215617d1a57600080fd5b8151614ecf816186f0565b600060208284031215617d3757600080fd5b5051919050565b600080600060608486031215617d5357600080fd5b835192506020840151915060408401516001600160401b03811115617d7757600080fd5b617bde86828701617834565b600080600080600060a08688031215617d9b57600080fd5b853594506020860135935060408601356001600160401b0380821115617dc057600080fd5b617dcc89838a016177de565b94506060880135915080821115617de257600080fd5b50617def888289016177de565b95989497509295608001359392505050565b60008060008060808587031215617e1757600080fd5b84516001600160401b0380821115617e2e57600080fd5b617e3a88838901617834565b95506020870151915080821115617e5057600080fd5b617e5c88838901617834565b94506040870151915080821115617e7257600080fd5b50617e7f87828801617834565b606096909601519497939650505050565b600060208284031215617ea257600080fd5b81356001600160401b03811115617eb857600080fd5b61787984828501617881565b60008060008060808587031215617eda57600080fd5b84356001600160401b03811115617ef057600080fd5b617efc87828801617881565b9450506020850135617f0d816186db565b93969395505050506040820135916060013590565b600060208284031215617f3457600080fd5b5035919050565b60008060408385031215617f4e57600080fd5b823591506020830135617f60816186f0565b809150509250929050565b60008060408385031215617f7e57600080fd5b505080516020909101519092909150565b600080600060608486031215617fa457600080fd5b83519250602084015191506040840151617b85816186f0565b600080600060608486031215617fd257600080fd5b8351925060208401519150604084015190509250925092565b6000815180845261800381602086016020860161864e565b601f01601f19169290920160200192915050565b8054600090600181811c908083168061803157607f831692505b602080841082141561805357634e487b7160e01b600052602260045260246000fd5b8180156180675760018114618078576180a5565b60ff198616895284890196506180a5565b60008881526020902060005b8681101561809d5781548b820152908501908301618084565b505084890196505b50505050505092915050565b600082516180c381846020870161864e565b9190910192915050565b600083516180df81846020880161864e565b9190910191825250602001919050565b6000614ecf8284618017565b60006181078285618017565b9283525050602001919050565b8281526040602082015260006178796040830184617feb565b83815282602082015260606040820152600061814c6060830184617feb565b95945050505050565b6080815260006181686080830187617feb565b828103602084015261817a8187617feb565b9050828103604084015261818e8186617feb565b91505082606083015295945050505050565b602081526000614ecf6020830184617feb565b6020808252601b908201527f4f6c642061646472657373206973206e6f742061206d656d6265720000000000604082015260600190565b60208082526027908201527f496e73756666696369656e742062616c616e636520746861742063616e206265604082015266081b1bd8dad95960ca1b606082015260800190565b60208082526029908201527f45524331393637557067726164653a20756e737570706f727465642070726f786040820152681a58589b195555525160ba1b606082015260800190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6020808252600d908201526c2737902832b936b4b9b9b4b7b760991b604082015260600190565b6020808252601490820152734475706c696361746564206e6f646520696e666f60601b604082015260600190565b6020808252602e908201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960408201526d6f6e206973206e6f74205555505360901b606082015260800190565b602080825260139082015272125b9d985b1a59081b1bd8dac8185b5bdd5b9d606a1b604082015260600190565b60208082526023908201527f4e6f7720696e20766f74696e67207769746820646966666572656e742062616c6040820152621b1bdd60ea1b606082015260800190565b60208082526017908201527f496e76616c6964207374616b696e672062616c616e6365000000000000000000604082015260600190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b6020808252601d908201527f43616e6e6f74206164642070726f706f73616c20746f6f206561726c79000000604082015260600190565b8c8152602081018c9052604081018b90526001600160a01b038a81166060830152898116608083015288811660a083015287811660c0830152861660e08201526000610180806101008401526184ec81840188617feb565b90508281036101208401526185018187617feb565b90508281036101408401526185168186617feb565b915050826101608301529d9c50505050505050505050505050565b87815286602082015285604082015260018060a01b03851660608201528360808201528260a082015260e060c0820152600061857060e0830184617feb565b9998505050505050505050565b60405161014081016001600160401b03811182821017156185a0576185a06186c5565b60405290565b604051601f8201601f191681016001600160401b03811182821017156185ce576185ce6186c5565b604052919050565b60006001600160401b038211156185ef576185ef6186c5565b50601f01601f191660200190565b60008219821115618610576186106186af565b500190565b60008261863257634e487b7160e01b600052601260045260246000fd5b500490565b600082821015618649576186496186af565b500390565b60005b83811015618669578181015183820152602001618651565b83811115611e9d5750506000910152565b600181811c9082168061868e57607f821691505b602082108114156138db57634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461494457600080fd5b801515811461494457600080fdfe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249885e7f4987c0698db47045ad8cea110b51138f0eecbd94915842328cf6c3dc97d416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220945c6a3e5d9b990fc60629c530b7df5bd7cf0e3e9d0c588fd5fd82a04f3b2fda64736f6c63430008060033",
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
