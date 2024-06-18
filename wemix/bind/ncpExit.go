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

// NCPExitMetaData contains all meta data concerning the NCPExit contract.
var NCPExitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_imp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"5c60da1b": "implementation()",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161077238038061077283398101604081905261002f91610326565b604080516020810190915260008152819061006b60017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd61034f565b60008051602061072b8339815191521461008757610087610374565b6100938282600061009b565b505050610405565b6100a4836100d1565b6000825111806100b15750805b156100cc576100ca838361011160201b61008b1760201c565b505b505050565b6100da8161013d565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060610136838360405180606001604052806027815260200161074b602791396101fd565b9392505050565b610150816102db60201b6100b71760201c565b6101b75760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b806101dc60008051602061072b83398151915260001b6102ea60201b6100c61760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606001600160a01b0384163b6102655760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016101ae565b600080856001600160a01b03168560405161028091906103b6565b600060405180830381855af49150503d80600081146102bb576040519150601f19603f3d011682016040523d82523d6000602084013e6102c0565b606091505b5090925090506102d18282866102ed565b9695505050505050565b6001600160a01b03163b151590565b90565b606083156102fc575081610136565b82511561030c5782518084602001fd5b8160405162461bcd60e51b81526004016101ae91906103d2565b60006020828403121561033857600080fd5b81516001600160a01b038116811461013657600080fd5b60008282101561036f57634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052600160045260246000fd5b60005b838110156103a557818101518382015260200161038d565b838111156100ca5750506000910152565b600082516103c881846020870161038a565b9190910192915050565b60208152600082518060208401526103f181604085016020870161038a565b601f01601f19169190910160400192915050565b610317806104146000396000f3fe6080604052600436106100225760003560e01c80635c60da1b1461003957610031565b366100315761002f61006a565b005b61002f61006a565b34801561004557600080fd5b5061004e61007c565b6040516001600160a01b03909116815260200160405180910390f35b61007a6100756100c9565b6100fc565b565b60006100866100c9565b905090565b60606100b083836040518060600160405280602781526020016102bb60279139610120565b9392505050565b6001600160a01b03163b151590565b90565b60006100867f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b3660008037600080366000845af43d6000803e80801561011b573d6000f35b3d6000fd5b60606001600160a01b0384163b61018d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084015b60405180910390fd5b600080856001600160a01b0316856040516101a8919061026b565b600060405180830381855af49150503d80600081146101e3576040519150601f19603f3d011682016040523d82523d6000602084013e6101e8565b606091505b50915091506101f8828286610202565b9695505050505050565b606083156102115750816100b0565b8251156102215782518084602001fd5b8160405162461bcd60e51b81526004016101849190610287565b60005b8381101561025657818101518382015260200161023e565b83811115610265576000848401525b50505050565b6000825161027d81846020870161023b565b9190910192915050565b60208152600082518060208401526102a681604085016020870161023b565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220e4a9b37c89f90cd2690d55474893040990839e6d9630f2ca907c1c29365285c264736f6c634300080e0033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
}

// NCPExitABI is the input ABI used to generate the binding from.
// Deprecated: Use NCPExitMetaData.ABI instead.
var NCPExitABI = NCPExitMetaData.ABI

// Deprecated: Use NCPExitMetaData.Sigs instead.
// NCPExitFuncSigs maps the 4-byte function signature to its string representation.
var NCPExitFuncSigs = NCPExitMetaData.Sigs

// NCPExitBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NCPExitMetaData.Bin instead.
var NCPExitBin = NCPExitMetaData.Bin

// DeployNCPExit deploys a new Ethereum contract, binding an instance of NCPExit to it.
func DeployNCPExit(auth *bind.TransactOpts, backend bind.ContractBackend, _imp common.Address) (common.Address, *types.Transaction, *NCPExit, error) {
	parsed, err := NCPExitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NCPExitBin), backend, _imp)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NCPExit{NCPExitCaller: NCPExitCaller{contract: contract}, NCPExitTransactor: NCPExitTransactor{contract: contract}, NCPExitFilterer: NCPExitFilterer{contract: contract}}, nil
}

// NCPExit is an auto generated Go binding around an Ethereum contract.
type NCPExit struct {
	NCPExitCaller     // Read-only binding to the contract
	NCPExitTransactor // Write-only binding to the contract
	NCPExitFilterer   // Log filterer for contract events
}

// NCPExitCaller is an auto generated read-only Go binding around an Ethereum contract.
type NCPExitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NCPExitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NCPExitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NCPExitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NCPExitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NCPExitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NCPExitSession struct {
	Contract     *NCPExit          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NCPExitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NCPExitCallerSession struct {
	Contract *NCPExitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// NCPExitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NCPExitTransactorSession struct {
	Contract     *NCPExitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NCPExitRaw is an auto generated low-level Go binding around an Ethereum contract.
type NCPExitRaw struct {
	Contract *NCPExit // Generic contract binding to access the raw methods on
}

// NCPExitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NCPExitCallerRaw struct {
	Contract *NCPExitCaller // Generic read-only contract binding to access the raw methods on
}

// NCPExitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NCPExitTransactorRaw struct {
	Contract *NCPExitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNCPExit creates a new instance of NCPExit, bound to a specific deployed contract.
func NewNCPExit(address common.Address, backend bind.ContractBackend) (*NCPExit, error) {
	contract, err := bindNCPExit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NCPExit{NCPExitCaller: NCPExitCaller{contract: contract}, NCPExitTransactor: NCPExitTransactor{contract: contract}, NCPExitFilterer: NCPExitFilterer{contract: contract}}, nil
}

// NewNCPExitCaller creates a new read-only instance of NCPExit, bound to a specific deployed contract.
func NewNCPExitCaller(address common.Address, caller bind.ContractCaller) (*NCPExitCaller, error) {
	contract, err := bindNCPExit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NCPExitCaller{contract: contract}, nil
}

// NewNCPExitTransactor creates a new write-only instance of NCPExit, bound to a specific deployed contract.
func NewNCPExitTransactor(address common.Address, transactor bind.ContractTransactor) (*NCPExitTransactor, error) {
	contract, err := bindNCPExit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NCPExitTransactor{contract: contract}, nil
}

// NewNCPExitFilterer creates a new log filterer instance of NCPExit, bound to a specific deployed contract.
func NewNCPExitFilterer(address common.Address, filterer bind.ContractFilterer) (*NCPExitFilterer, error) {
	contract, err := bindNCPExit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NCPExitFilterer{contract: contract}, nil
}

// bindNCPExit binds a generic wrapper to an already deployed contract.
func bindNCPExit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NCPExitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NCPExit *NCPExitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NCPExit.Contract.NCPExitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NCPExit *NCPExitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NCPExit.Contract.NCPExitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NCPExit *NCPExitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NCPExit.Contract.NCPExitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NCPExit *NCPExitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NCPExit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NCPExit *NCPExitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NCPExit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NCPExit *NCPExitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NCPExit.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_NCPExit *NCPExitCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NCPExit.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_NCPExit *NCPExitSession) Implementation() (common.Address, error) {
	return _NCPExit.Contract.Implementation(&_NCPExit.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_NCPExit *NCPExitCallerSession) Implementation() (common.Address, error) {
	return _NCPExit.Contract.Implementation(&_NCPExit.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NCPExit *NCPExitTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _NCPExit.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NCPExit *NCPExitSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NCPExit.Contract.Fallback(&_NCPExit.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NCPExit *NCPExitTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NCPExit.Contract.Fallback(&_NCPExit.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NCPExit *NCPExitTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NCPExit.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NCPExit *NCPExitSession) Receive() (*types.Transaction, error) {
	return _NCPExit.Contract.Receive(&_NCPExit.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NCPExit *NCPExitTransactorSession) Receive() (*types.Transaction, error) {
	return _NCPExit.Contract.Receive(&_NCPExit.TransactOpts)
}

// NCPExitAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the NCPExit contract.
type NCPExitAdminChangedIterator struct {
	Event *NCPExitAdminChanged // Event containing the contract specifics and raw log

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
func (it *NCPExitAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NCPExitAdminChanged)
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
		it.Event = new(NCPExitAdminChanged)
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
func (it *NCPExitAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NCPExitAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NCPExitAdminChanged represents a AdminChanged event raised by the NCPExit contract.
type NCPExitAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NCPExit *NCPExitFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*NCPExitAdminChangedIterator, error) {

	logs, sub, err := _NCPExit.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &NCPExitAdminChangedIterator{contract: _NCPExit.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NCPExit *NCPExitFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *NCPExitAdminChanged) (event.Subscription, error) {

	logs, sub, err := _NCPExit.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NCPExitAdminChanged)
				if err := _NCPExit.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_NCPExit *NCPExitFilterer) ParseAdminChanged(log types.Log) (*NCPExitAdminChanged, error) {
	event := new(NCPExitAdminChanged)
	if err := _NCPExit.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NCPExitBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the NCPExit contract.
type NCPExitBeaconUpgradedIterator struct {
	Event *NCPExitBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *NCPExitBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NCPExitBeaconUpgraded)
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
		it.Event = new(NCPExitBeaconUpgraded)
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
func (it *NCPExitBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NCPExitBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NCPExitBeaconUpgraded represents a BeaconUpgraded event raised by the NCPExit contract.
type NCPExitBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NCPExit *NCPExitFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*NCPExitBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NCPExit.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &NCPExitBeaconUpgradedIterator{contract: _NCPExit.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NCPExit *NCPExitFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *NCPExitBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NCPExit.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NCPExitBeaconUpgraded)
				if err := _NCPExit.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_NCPExit *NCPExitFilterer) ParseBeaconUpgraded(log types.Log) (*NCPExitBeaconUpgraded, error) {
	event := new(NCPExitBeaconUpgraded)
	if err := _NCPExit.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NCPExitUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NCPExit contract.
type NCPExitUpgradedIterator struct {
	Event *NCPExitUpgraded // Event containing the contract specifics and raw log

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
func (it *NCPExitUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NCPExitUpgraded)
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
		it.Event = new(NCPExitUpgraded)
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
func (it *NCPExitUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NCPExitUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NCPExitUpgraded represents a Upgraded event raised by the NCPExit contract.
type NCPExitUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NCPExit *NCPExitFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NCPExitUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NCPExit.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NCPExitUpgradedIterator{contract: _NCPExit.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NCPExit *NCPExitFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NCPExitUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NCPExit.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NCPExitUpgraded)
				if err := _NCPExit.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_NCPExit *NCPExitFilterer) ParseUpgraded(log types.Log) (*NCPExitUpgraded, error) {
	event := new(NCPExitUpgraded)
	if err := _NCPExit.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NCPExitImpMetaData contains all meta data concerning the NCPExitImp contract.
var NCPExitImpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ECOSYSTEM_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ENV_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAINTENANCE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_POOL_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_REWARD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"exitNcp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedUserBalanceToNCPTotal\",\"type\":\"uint256\"}],\"name\":\"depositExitAmount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"exitNcp\",\"type\":\"address\"}],\"name\":\"getAvailableAmountForAdministrator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"exitNcp\",\"type\":\"address\"}],\"name\":\"getLockedUserBalanceToNCPTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reg\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdministrator\",\"type\":\"address\"}],\"name\":\"setAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdministratorSetter\",\"type\":\"address\"}],\"name\":\"setAdministratorSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImp\",\"type\":\"address\"}],\"name\":\"upgradeNCPExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"exitNcp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawForAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"exitNcp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exitUser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawForUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"9986e4b9": "BALLOT_STORAGE_NAME()",
		"34125c84": "ECOSYSTEM_NAME()",
		"7bf46530": "ENV_STORAGE_NAME()",
		"6c78d2cf": "GOV_NAME()",
		"4bd1ed76": "MAINTENANCE_NAME()",
		"2f40992e": "REWARD_POOL_NAME()",
		"1e0cba0d": "STAKING_NAME()",
		"5a731cca": "STAKING_REWARD_NAME()",
		"8408bdb1": "depositExitAmount(address,uint256,uint256)",
		"dd38f90b": "getAvailableAmountForAdministrator(address)",
		"805c5ccc": "getLockedUserBalanceToNCPTotal(address)",
		"c4d66de8": "initialize(address)",
		"8da5cb5b": "owner()",
		"52d1902d": "proxiableUUID()",
		"738fdd1a": "reg()",
		"715018a6": "renounceOwnership()",
		"df8089ef": "setAdministrator(address)",
		"8c56c8c3": "setAdministratorSetter(address)",
		"a91ee0dc": "setRegistry(address)",
		"f2fde38b": "transferOwnership(address)",
		"1083fc6e": "upgradeNCPExit(address)",
		"3659cfe6": "upgradeTo(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
		"00fc5701": "withdrawForAdministrator(address,uint256,address)",
		"b007a5ad": "withdrawForUser(address,address,uint256)",
	},
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100e2565b600054610100900460ff161561008e5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156100e0576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b608051611a67610119600039600081816105ee0152818161062e015281816106ae015281816106ee01526107810152611a676000f3fe6080604052600436106101695760003560e01c8063738fdd1a116100d15780639986e4b91161008a578063c4d66de811610064578063c4d66de814610439578063dd38f90b14610459578063df8089ef14610479578063f2fde38b1461049957600080fd5b80639986e4b9146103d5578063a91ee0dc146103f9578063b007a5ad1461041957600080fd5b8063738fdd1a146102f55780637bf465301461032d578063805c5ccc1461034e5780638408bdb1146103845780638c56c8c3146103975780638da5cb5b146103b757600080fd5b80634bd1ed76116101235780634bd1ed76146102495780634f1ef2861461026b57806352d1902d1461027e5780635a731cca146102935780636c78d2cf146102b7578063715018a6146102e057600080fd5b8062fc5701146101755780631083fc6e146101975780631e0cba0d146101b75780632f40992e146101e857806334125c84146102095780633659cfe61461022957600080fd5b3661017057005b600080fd5b34801561018157600080fd5b50610195610190366004611588565b6104b9565b005b3480156101a357600080fd5b506101956101b23660046115ca565b6105a5565b3480156101c357600080fd5b506101d5665374616b696e6760c81b81565b6040519081526020015b60405180910390f35b3480156101f457600080fd5b506101d56914995dd85c99141bdbdb60b21b81565b34801561021557600080fd5b506101d56845636f73797374656d60b81b81565b34801561023557600080fd5b506101956102443660046115ca565b6105e4565b34801561025557600080fd5b506101d56a4d61696e74656e616e636560a81b81565b6101956102793660046115fd565b6106a4565b34801561028a57600080fd5b506101d5610774565b34801561029f57600080fd5b506101d56c14dd185ada5b99d4995dd85c99609a1b81565b3480156102c357600080fd5b506101d57111dbdd995c9b985b98d950dbdb9d1c9858dd60721b81565b3480156102ec57600080fd5b50610195610827565b34801561030157600080fd5b50606554610315906001600160a01b031681565b6040516001600160a01b0390911681526020016101df565b34801561033957600080fd5b506101d569456e7653746f7261676560b01b81565b34801561035a57600080fd5b506101d56103693660046115ca565b6001600160a01b03166000908152609b602052604090205490565b6101956103923660046116c1565b61083b565b3480156103a357600080fd5b506101956103b23660046115ca565b610929565b3480156103c357600080fd5b506033546001600160a01b0316610315565b3480156103e157600080fd5b506101d56c42616c6c6f7453746f7261676560981b81565b34801561040557600080fd5b506101956104143660046115ca565b61099b565b34801561042557600080fd5b506101956104343660046116f6565b610a13565b34801561044557600080fd5b506101956104543660046115ca565b610c10565b34801561046557600080fd5b506101d56104743660046115ca565b610d5f565b34801561048557600080fd5b506101956104943660046115ca565b610dc1565b3480156104a557600080fd5b506101956104b43660046115ca565b610e33565b6002606654036104e45760405162461bcd60e51b81526004016104db90611737565b60405180910390fd5b60026066556098546001600160a01b031633146105135760405162461bcd60e51b81526004016104db9061176e565b6001600160a01b0383166000908152609b6020908152604080832054609a909252909120548391610543916117b8565b101561054e57600080fd5b6001600160a01b0383166000908152609a60205260409020546105729083906117b8565b6001600160a01b038085166000908152609a602052604090209190915561059b90821683610ea9565b5050600160665550565b6105ad610fc7565b6001600160a01b038116156105e1576105c581611021565b604080516000808252602082019092526105e191839190611029565b50565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361062c5760405162461bcd60e51b81526004016104db906117dd565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166106756000805160206119eb833981519152546001600160a01b031690565b6001600160a01b03161461069b5760405162461bcd60e51b81526004016104db90611829565b6105c581611021565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036106ec5760405162461bcd60e51b81526004016104db906117dd565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166107356000805160206119eb833981519152546001600160a01b031690565b6001600160a01b03161461075b5760405162461bcd60e51b81526004016104db90611829565b61076482611021565b61077082826001611029565b5050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108145760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016104db565b506000805160206119eb83398151915290565b61082f610fc7565b6108396000611194565b565b60026066540361085d5760405162461bcd60e51b81526004016104db90611737565b600260665561086a6111e6565b6001600160a01b0316336001600160a01b0316146108f05760405162461bcd60e51b815260206004820152603860248201527f4f6e6c7920676f7665726e616e6365207374616b696e6720636f6e747261637460448201527f2063616e2063616c6c20746869732066756e6374696f6e2e000000000000000060648201526084016104db565b3482146108fc57600080fd5b6001600160a01b039092166000908152609a6020908152604080832093909355609b905220556001606655565b6099546001600160a01b031633146109535760405162461bcd60e51b81526004016104db90611875565b6001600160a01b0381166109795760405162461bcd60e51b81526004016104db906118b7565b609980546001600160a01b0319166001600160a01b0392909216919091179055565b6109a3610fc7565b6001600160a01b0381166109c95760405162461bcd60e51b81526004016104db906118b7565b606580546001600160a01b0319166001600160a01b0383169081179091556040517f278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd90600090a250565b600260665403610a355760405162461bcd60e51b81526004016104db90611737565b6002606655610a426111e6565b6001600160a01b031663f1b8aa1d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a7f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa391906118ee565b6001600160a01b0316336001600160a01b031614610b135760405162461bcd60e51b815260206004820152602760248201527f4f6e6c79204e63705374616b696e672063616e2063616c6c20746869732066756044820152663731ba34b7b71760c91b60648201526084016104db565b6001600160a01b0383166000908152609b6020526040902054811115610b935760405162461bcd60e51b815260206004820152602f60248201527f5f6c6f636b65645573657242616c616e6365546f4e4350546f74616c5b65786960448201526e1d1398dc17480f8f48185b5bdd5b9d608a1b60648201526084016104db565b6001600160a01b0383166000908152609a6020526040902054610bb79082906117b8565b6001600160a01b0384166000908152609a6020908152604080832093909355609b90522054610be79082906117b8565b6001600160a01b038085166000908152609b602052604090209190915561059b90831682610ea9565b600054610100900460ff1615808015610c305750600054600160ff909116105b80610c4a5750303b158015610c4a575060005460ff166001145b610cad5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016104db565b6000805460ff191660011790558015610cd0576000805461ff0019166101001790555b610cd8611200565b610ce061122f565b610ce98261099b565b603354609880546001600160a01b039092166001600160a01b03199283168117909155609980549092161790558015610770576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6098546000906001600160a01b03163314610d8c5760405162461bcd60e51b81526004016104db9061176e565b6001600160a01b0382166000908152609b6020908152604080832054609a90925290912054610dbb91906117b8565b92915050565b6099546001600160a01b03163314610deb5760405162461bcd60e51b81526004016104db90611875565b6001600160a01b038116610e115760405162461bcd60e51b81526004016104db906118b7565b609880546001600160a01b0319166001600160a01b0392909216919091179055565b610e3b610fc7565b6001600160a01b038116610ea05760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016104db565b6105e181611194565b80471015610ef95760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e636500000060448201526064016104db565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114610f46576040519150601f19603f3d011682016040523d82523d6000602084013e610f4b565b606091505b5050905080610fc25760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d6179206861766520726576657274656400000000000060648201526084016104db565b505050565b6033546001600160a01b031633146108395760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104db565b6105e1610fc7565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561105c57610fc28361125e565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156110b6575060408051601f3d908101601f191682019092526110b39181019061190b565b60015b6111195760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b60648201526084016104db565b6000805160206119eb83398151915281146111885760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b60648201526084016104db565b50610fc28383836112fa565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006111fb665374616b696e6760c81b611325565b905090565b600054610100900460ff166112275760405162461bcd60e51b81526004016104db90611924565b610839611393565b600054610100900460ff166112565760405162461bcd60e51b81526004016104db90611924565b6108396113c3565b6001600160a01b0381163b6112cb5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016104db565b6000805160206119eb83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b611303836113f1565b6000825111806113105750805b15610fc25761131f8383611431565b50505050565b606554604051630d2020dd60e01b8152600481018390526000916001600160a01b031690630d2020dd90602401602060405180830381865afa15801561136f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dbb91906118ee565b600054610100900460ff166113ba5760405162461bcd60e51b81526004016104db90611924565b61083933611194565b600054610100900460ff166113ea5760405162461bcd60e51b81526004016104db90611924565b6001606655565b6113fa8161125e565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606114568383604051806060016040528060278152602001611a0b6027913961145d565b9392505050565b60606001600160a01b0384163b6114c55760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016104db565b600080856001600160a01b0316856040516114e0919061199b565b600060405180830381855af49150503d806000811461151b576040519150601f19603f3d011682016040523d82523d6000602084013e611520565b606091505b509150915061153082828661153a565b9695505050505050565b60608315611549575081611456565b8251156115595782518084602001fd5b8160405162461bcd60e51b81526004016104db91906119b7565b6001600160a01b03811681146105e157600080fd5b60008060006060848603121561159d57600080fd5b83356115a881611573565b92506020840135915060408401356115bf81611573565b809150509250925092565b6000602082840312156115dc57600080fd5b813561145681611573565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561161057600080fd5b823561161b81611573565b9150602083013567ffffffffffffffff8082111561163857600080fd5b818501915085601f83011261164c57600080fd5b81358181111561165e5761165e6115e7565b604051601f8201601f19908116603f01168101908382118183101715611686576116866115e7565b8160405282815288602084870101111561169f57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b6000806000606084860312156116d657600080fd5b83356116e181611573565b95602085013595506040909401359392505050565b60008060006060848603121561170b57600080fd5b833561171681611573565b9250602084013561172681611573565b929592945050506040919091013590565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b6020808252602a908201527f4f6e6c792041646d696e6973747261746f722063616e2063616c6c207468697360408201526910333ab731ba34b7b71760b11b606082015260800190565b6000828210156117d857634e487b7160e01b600052601160045260246000fd5b500390565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b60208082526022908201527f43616c6c6572206973206e6f742041646d696e6973747261746f725365747465604082015261391760f11b606082015260800190565b6020808252601a908201527f416464726573732073686f756c64206265206e6f6e2d7a65726f000000000000604082015260600190565b60006020828403121561190057600080fd5b815161145681611573565b60006020828403121561191d57600080fd5b5051919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60005b8381101561198a578181015183820152602001611972565b8381111561131f5750506000910152565b600082516119ad81846020870161196f565b9190910192915050565b60208152600082518060208401526119d681604085016020870161196f565b601f01601f1916919091016040019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220b00f4e5c886400503c92d7ed21f3804a28036c7ccf34340b73b7a47ece983e6764736f6c634300080e0033",
}

// NCPExitImpABI is the input ABI used to generate the binding from.
// Deprecated: Use NCPExitImpMetaData.ABI instead.
var NCPExitImpABI = NCPExitImpMetaData.ABI

// Deprecated: Use NCPExitImpMetaData.Sigs instead.
// NCPExitImpFuncSigs maps the 4-byte function signature to its string representation.
var NCPExitImpFuncSigs = NCPExitImpMetaData.Sigs

// NCPExitImpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NCPExitImpMetaData.Bin instead.
var NCPExitImpBin = NCPExitImpMetaData.Bin

// DeployNCPExitImp deploys a new Ethereum contract, binding an instance of NCPExitImp to it.
func DeployNCPExitImp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NCPExitImp, error) {
	parsed, err := NCPExitImpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NCPExitImpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NCPExitImp{NCPExitImpCaller: NCPExitImpCaller{contract: contract}, NCPExitImpTransactor: NCPExitImpTransactor{contract: contract}, NCPExitImpFilterer: NCPExitImpFilterer{contract: contract}}, nil
}

// NCPExitImp is an auto generated Go binding around an Ethereum contract.
type NCPExitImp struct {
	NCPExitImpCaller     // Read-only binding to the contract
	NCPExitImpTransactor // Write-only binding to the contract
	NCPExitImpFilterer   // Log filterer for contract events
}

// NCPExitImpCaller is an auto generated read-only Go binding around an Ethereum contract.
type NCPExitImpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NCPExitImpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NCPExitImpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NCPExitImpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NCPExitImpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NCPExitImpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NCPExitImpSession struct {
	Contract     *NCPExitImp       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NCPExitImpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NCPExitImpCallerSession struct {
	Contract *NCPExitImpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NCPExitImpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NCPExitImpTransactorSession struct {
	Contract     *NCPExitImpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NCPExitImpRaw is an auto generated low-level Go binding around an Ethereum contract.
type NCPExitImpRaw struct {
	Contract *NCPExitImp // Generic contract binding to access the raw methods on
}

// NCPExitImpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NCPExitImpCallerRaw struct {
	Contract *NCPExitImpCaller // Generic read-only contract binding to access the raw methods on
}

// NCPExitImpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NCPExitImpTransactorRaw struct {
	Contract *NCPExitImpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNCPExitImp creates a new instance of NCPExitImp, bound to a specific deployed contract.
func NewNCPExitImp(address common.Address, backend bind.ContractBackend) (*NCPExitImp, error) {
	contract, err := bindNCPExitImp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NCPExitImp{NCPExitImpCaller: NCPExitImpCaller{contract: contract}, NCPExitImpTransactor: NCPExitImpTransactor{contract: contract}, NCPExitImpFilterer: NCPExitImpFilterer{contract: contract}}, nil
}

// NewNCPExitImpCaller creates a new read-only instance of NCPExitImp, bound to a specific deployed contract.
func NewNCPExitImpCaller(address common.Address, caller bind.ContractCaller) (*NCPExitImpCaller, error) {
	contract, err := bindNCPExitImp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NCPExitImpCaller{contract: contract}, nil
}

// NewNCPExitImpTransactor creates a new write-only instance of NCPExitImp, bound to a specific deployed contract.
func NewNCPExitImpTransactor(address common.Address, transactor bind.ContractTransactor) (*NCPExitImpTransactor, error) {
	contract, err := bindNCPExitImp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NCPExitImpTransactor{contract: contract}, nil
}

// NewNCPExitImpFilterer creates a new log filterer instance of NCPExitImp, bound to a specific deployed contract.
func NewNCPExitImpFilterer(address common.Address, filterer bind.ContractFilterer) (*NCPExitImpFilterer, error) {
	contract, err := bindNCPExitImp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NCPExitImpFilterer{contract: contract}, nil
}

// bindNCPExitImp binds a generic wrapper to an already deployed contract.
func bindNCPExitImp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NCPExitImpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NCPExitImp *NCPExitImpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NCPExitImp.Contract.NCPExitImpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NCPExitImp *NCPExitImpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NCPExitImp.Contract.NCPExitImpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NCPExitImp *NCPExitImpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NCPExitImp.Contract.NCPExitImpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NCPExitImp *NCPExitImpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NCPExitImp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NCPExitImp *NCPExitImpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NCPExitImp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NCPExitImp *NCPExitImpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NCPExitImp.Contract.contract.Transact(opts, method, params...)
}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCaller) BALLOTSTORAGENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NCPExitImp.contract.Call(opts, &out, "BALLOT_STORAGE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpSession) BALLOTSTORAGENAME() ([32]byte, error) {
	return _NCPExitImp.Contract.BALLOTSTORAGENAME(&_NCPExitImp.CallOpts)
}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCallerSession) BALLOTSTORAGENAME() ([32]byte, error) {
	return _NCPExitImp.Contract.BALLOTSTORAGENAME(&_NCPExitImp.CallOpts)
}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCaller) ECOSYSTEMNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NCPExitImp.contract.Call(opts, &out, "ECOSYSTEM_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpSession) ECOSYSTEMNAME() ([32]byte, error) {
	return _NCPExitImp.Contract.ECOSYSTEMNAME(&_NCPExitImp.CallOpts)
}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCallerSession) ECOSYSTEMNAME() ([32]byte, error) {
	return _NCPExitImp.Contract.ECOSYSTEMNAME(&_NCPExitImp.CallOpts)
}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCaller) ENVSTORAGENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NCPExitImp.contract.Call(opts, &out, "ENV_STORAGE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpSession) ENVSTORAGENAME() ([32]byte, error) {
	return _NCPExitImp.Contract.ENVSTORAGENAME(&_NCPExitImp.CallOpts)
}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCallerSession) ENVSTORAGENAME() ([32]byte, error) {
	return _NCPExitImp.Contract.ENVSTORAGENAME(&_NCPExitImp.CallOpts)
}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCaller) GOVNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NCPExitImp.contract.Call(opts, &out, "GOV_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpSession) GOVNAME() ([32]byte, error) {
	return _NCPExitImp.Contract.GOVNAME(&_NCPExitImp.CallOpts)
}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCallerSession) GOVNAME() ([32]byte, error) {
	return _NCPExitImp.Contract.GOVNAME(&_NCPExitImp.CallOpts)
}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCaller) MAINTENANCENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NCPExitImp.contract.Call(opts, &out, "MAINTENANCE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpSession) MAINTENANCENAME() ([32]byte, error) {
	return _NCPExitImp.Contract.MAINTENANCENAME(&_NCPExitImp.CallOpts)
}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCallerSession) MAINTENANCENAME() ([32]byte, error) {
	return _NCPExitImp.Contract.MAINTENANCENAME(&_NCPExitImp.CallOpts)
}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCaller) REWARDPOOLNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NCPExitImp.contract.Call(opts, &out, "REWARD_POOL_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpSession) REWARDPOOLNAME() ([32]byte, error) {
	return _NCPExitImp.Contract.REWARDPOOLNAME(&_NCPExitImp.CallOpts)
}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCallerSession) REWARDPOOLNAME() ([32]byte, error) {
	return _NCPExitImp.Contract.REWARDPOOLNAME(&_NCPExitImp.CallOpts)
}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCaller) STAKINGNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NCPExitImp.contract.Call(opts, &out, "STAKING_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpSession) STAKINGNAME() ([32]byte, error) {
	return _NCPExitImp.Contract.STAKINGNAME(&_NCPExitImp.CallOpts)
}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCallerSession) STAKINGNAME() ([32]byte, error) {
	return _NCPExitImp.Contract.STAKINGNAME(&_NCPExitImp.CallOpts)
}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCaller) STAKINGREWARDNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NCPExitImp.contract.Call(opts, &out, "STAKING_REWARD_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpSession) STAKINGREWARDNAME() ([32]byte, error) {
	return _NCPExitImp.Contract.STAKINGREWARDNAME(&_NCPExitImp.CallOpts)
}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCallerSession) STAKINGREWARDNAME() ([32]byte, error) {
	return _NCPExitImp.Contract.STAKINGREWARDNAME(&_NCPExitImp.CallOpts)
}

// GetAvailableAmountForAdministrator is a free data retrieval call binding the contract method 0xdd38f90b.
//
// Solidity: function getAvailableAmountForAdministrator(address exitNcp) view returns(uint256)
func (_NCPExitImp *NCPExitImpCaller) GetAvailableAmountForAdministrator(opts *bind.CallOpts, exitNcp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NCPExitImp.contract.Call(opts, &out, "getAvailableAmountForAdministrator", exitNcp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableAmountForAdministrator is a free data retrieval call binding the contract method 0xdd38f90b.
//
// Solidity: function getAvailableAmountForAdministrator(address exitNcp) view returns(uint256)
func (_NCPExitImp *NCPExitImpSession) GetAvailableAmountForAdministrator(exitNcp common.Address) (*big.Int, error) {
	return _NCPExitImp.Contract.GetAvailableAmountForAdministrator(&_NCPExitImp.CallOpts, exitNcp)
}

// GetAvailableAmountForAdministrator is a free data retrieval call binding the contract method 0xdd38f90b.
//
// Solidity: function getAvailableAmountForAdministrator(address exitNcp) view returns(uint256)
func (_NCPExitImp *NCPExitImpCallerSession) GetAvailableAmountForAdministrator(exitNcp common.Address) (*big.Int, error) {
	return _NCPExitImp.Contract.GetAvailableAmountForAdministrator(&_NCPExitImp.CallOpts, exitNcp)
}

// GetLockedUserBalanceToNCPTotal is a free data retrieval call binding the contract method 0x805c5ccc.
//
// Solidity: function getLockedUserBalanceToNCPTotal(address exitNcp) view returns(uint256)
func (_NCPExitImp *NCPExitImpCaller) GetLockedUserBalanceToNCPTotal(opts *bind.CallOpts, exitNcp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NCPExitImp.contract.Call(opts, &out, "getLockedUserBalanceToNCPTotal", exitNcp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedUserBalanceToNCPTotal is a free data retrieval call binding the contract method 0x805c5ccc.
//
// Solidity: function getLockedUserBalanceToNCPTotal(address exitNcp) view returns(uint256)
func (_NCPExitImp *NCPExitImpSession) GetLockedUserBalanceToNCPTotal(exitNcp common.Address) (*big.Int, error) {
	return _NCPExitImp.Contract.GetLockedUserBalanceToNCPTotal(&_NCPExitImp.CallOpts, exitNcp)
}

// GetLockedUserBalanceToNCPTotal is a free data retrieval call binding the contract method 0x805c5ccc.
//
// Solidity: function getLockedUserBalanceToNCPTotal(address exitNcp) view returns(uint256)
func (_NCPExitImp *NCPExitImpCallerSession) GetLockedUserBalanceToNCPTotal(exitNcp common.Address) (*big.Int, error) {
	return _NCPExitImp.Contract.GetLockedUserBalanceToNCPTotal(&_NCPExitImp.CallOpts, exitNcp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NCPExitImp *NCPExitImpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NCPExitImp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NCPExitImp *NCPExitImpSession) Owner() (common.Address, error) {
	return _NCPExitImp.Contract.Owner(&_NCPExitImp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NCPExitImp *NCPExitImpCallerSession) Owner() (common.Address, error) {
	return _NCPExitImp.Contract.Owner(&_NCPExitImp.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NCPExitImp.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NCPExitImp *NCPExitImpSession) ProxiableUUID() ([32]byte, error) {
	return _NCPExitImp.Contract.ProxiableUUID(&_NCPExitImp.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NCPExitImp *NCPExitImpCallerSession) ProxiableUUID() ([32]byte, error) {
	return _NCPExitImp.Contract.ProxiableUUID(&_NCPExitImp.CallOpts)
}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_NCPExitImp *NCPExitImpCaller) Reg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NCPExitImp.contract.Call(opts, &out, "reg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_NCPExitImp *NCPExitImpSession) Reg() (common.Address, error) {
	return _NCPExitImp.Contract.Reg(&_NCPExitImp.CallOpts)
}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_NCPExitImp *NCPExitImpCallerSession) Reg() (common.Address, error) {
	return _NCPExitImp.Contract.Reg(&_NCPExitImp.CallOpts)
}

// DepositExitAmount is a paid mutator transaction binding the contract method 0x8408bdb1.
//
// Solidity: function depositExitAmount(address exitNcp, uint256 totalAmount, uint256 lockedUserBalanceToNCPTotal) payable returns()
func (_NCPExitImp *NCPExitImpTransactor) DepositExitAmount(opts *bind.TransactOpts, exitNcp common.Address, totalAmount *big.Int, lockedUserBalanceToNCPTotal *big.Int) (*types.Transaction, error) {
	return _NCPExitImp.contract.Transact(opts, "depositExitAmount", exitNcp, totalAmount, lockedUserBalanceToNCPTotal)
}

// DepositExitAmount is a paid mutator transaction binding the contract method 0x8408bdb1.
//
// Solidity: function depositExitAmount(address exitNcp, uint256 totalAmount, uint256 lockedUserBalanceToNCPTotal) payable returns()
func (_NCPExitImp *NCPExitImpSession) DepositExitAmount(exitNcp common.Address, totalAmount *big.Int, lockedUserBalanceToNCPTotal *big.Int) (*types.Transaction, error) {
	return _NCPExitImp.Contract.DepositExitAmount(&_NCPExitImp.TransactOpts, exitNcp, totalAmount, lockedUserBalanceToNCPTotal)
}

// DepositExitAmount is a paid mutator transaction binding the contract method 0x8408bdb1.
//
// Solidity: function depositExitAmount(address exitNcp, uint256 totalAmount, uint256 lockedUserBalanceToNCPTotal) payable returns()
func (_NCPExitImp *NCPExitImpTransactorSession) DepositExitAmount(exitNcp common.Address, totalAmount *big.Int, lockedUserBalanceToNCPTotal *big.Int) (*types.Transaction, error) {
	return _NCPExitImp.Contract.DepositExitAmount(&_NCPExitImp.TransactOpts, exitNcp, totalAmount, lockedUserBalanceToNCPTotal)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_NCPExitImp *NCPExitImpTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _NCPExitImp.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_NCPExitImp *NCPExitImpSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.Initialize(&_NCPExitImp.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_NCPExitImp *NCPExitImpTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.Initialize(&_NCPExitImp.TransactOpts, _registry)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NCPExitImp *NCPExitImpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NCPExitImp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NCPExitImp *NCPExitImpSession) RenounceOwnership() (*types.Transaction, error) {
	return _NCPExitImp.Contract.RenounceOwnership(&_NCPExitImp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NCPExitImp *NCPExitImpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NCPExitImp.Contract.RenounceOwnership(&_NCPExitImp.TransactOpts)
}

// SetAdministrator is a paid mutator transaction binding the contract method 0xdf8089ef.
//
// Solidity: function setAdministrator(address _newAdministrator) returns()
func (_NCPExitImp *NCPExitImpTransactor) SetAdministrator(opts *bind.TransactOpts, _newAdministrator common.Address) (*types.Transaction, error) {
	return _NCPExitImp.contract.Transact(opts, "setAdministrator", _newAdministrator)
}

// SetAdministrator is a paid mutator transaction binding the contract method 0xdf8089ef.
//
// Solidity: function setAdministrator(address _newAdministrator) returns()
func (_NCPExitImp *NCPExitImpSession) SetAdministrator(_newAdministrator common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.SetAdministrator(&_NCPExitImp.TransactOpts, _newAdministrator)
}

// SetAdministrator is a paid mutator transaction binding the contract method 0xdf8089ef.
//
// Solidity: function setAdministrator(address _newAdministrator) returns()
func (_NCPExitImp *NCPExitImpTransactorSession) SetAdministrator(_newAdministrator common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.SetAdministrator(&_NCPExitImp.TransactOpts, _newAdministrator)
}

// SetAdministratorSetter is a paid mutator transaction binding the contract method 0x8c56c8c3.
//
// Solidity: function setAdministratorSetter(address _newAdministratorSetter) returns()
func (_NCPExitImp *NCPExitImpTransactor) SetAdministratorSetter(opts *bind.TransactOpts, _newAdministratorSetter common.Address) (*types.Transaction, error) {
	return _NCPExitImp.contract.Transact(opts, "setAdministratorSetter", _newAdministratorSetter)
}

// SetAdministratorSetter is a paid mutator transaction binding the contract method 0x8c56c8c3.
//
// Solidity: function setAdministratorSetter(address _newAdministratorSetter) returns()
func (_NCPExitImp *NCPExitImpSession) SetAdministratorSetter(_newAdministratorSetter common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.SetAdministratorSetter(&_NCPExitImp.TransactOpts, _newAdministratorSetter)
}

// SetAdministratorSetter is a paid mutator transaction binding the contract method 0x8c56c8c3.
//
// Solidity: function setAdministratorSetter(address _newAdministratorSetter) returns()
func (_NCPExitImp *NCPExitImpTransactorSession) SetAdministratorSetter(_newAdministratorSetter common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.SetAdministratorSetter(&_NCPExitImp.TransactOpts, _newAdministratorSetter)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_NCPExitImp *NCPExitImpTransactor) SetRegistry(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _NCPExitImp.contract.Transact(opts, "setRegistry", _addr)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_NCPExitImp *NCPExitImpSession) SetRegistry(_addr common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.SetRegistry(&_NCPExitImp.TransactOpts, _addr)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_NCPExitImp *NCPExitImpTransactorSession) SetRegistry(_addr common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.SetRegistry(&_NCPExitImp.TransactOpts, _addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NCPExitImp *NCPExitImpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NCPExitImp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NCPExitImp *NCPExitImpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.TransferOwnership(&_NCPExitImp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NCPExitImp *NCPExitImpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.TransferOwnership(&_NCPExitImp.TransactOpts, newOwner)
}

// UpgradeNCPExit is a paid mutator transaction binding the contract method 0x1083fc6e.
//
// Solidity: function upgradeNCPExit(address newImp) returns()
func (_NCPExitImp *NCPExitImpTransactor) UpgradeNCPExit(opts *bind.TransactOpts, newImp common.Address) (*types.Transaction, error) {
	return _NCPExitImp.contract.Transact(opts, "upgradeNCPExit", newImp)
}

// UpgradeNCPExit is a paid mutator transaction binding the contract method 0x1083fc6e.
//
// Solidity: function upgradeNCPExit(address newImp) returns()
func (_NCPExitImp *NCPExitImpSession) UpgradeNCPExit(newImp common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.UpgradeNCPExit(&_NCPExitImp.TransactOpts, newImp)
}

// UpgradeNCPExit is a paid mutator transaction binding the contract method 0x1083fc6e.
//
// Solidity: function upgradeNCPExit(address newImp) returns()
func (_NCPExitImp *NCPExitImpTransactorSession) UpgradeNCPExit(newImp common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.UpgradeNCPExit(&_NCPExitImp.TransactOpts, newImp)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NCPExitImp *NCPExitImpTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _NCPExitImp.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NCPExitImp *NCPExitImpSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.UpgradeTo(&_NCPExitImp.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NCPExitImp *NCPExitImpTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.UpgradeTo(&_NCPExitImp.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NCPExitImp *NCPExitImpTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NCPExitImp.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NCPExitImp *NCPExitImpSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NCPExitImp.Contract.UpgradeToAndCall(&_NCPExitImp.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NCPExitImp *NCPExitImpTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NCPExitImp.Contract.UpgradeToAndCall(&_NCPExitImp.TransactOpts, newImplementation, data)
}

// WithdrawForAdministrator is a paid mutator transaction binding the contract method 0x00fc5701.
//
// Solidity: function withdrawForAdministrator(address exitNcp, uint256 amount, address to) returns()
func (_NCPExitImp *NCPExitImpTransactor) WithdrawForAdministrator(opts *bind.TransactOpts, exitNcp common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _NCPExitImp.contract.Transact(opts, "withdrawForAdministrator", exitNcp, amount, to)
}

// WithdrawForAdministrator is a paid mutator transaction binding the contract method 0x00fc5701.
//
// Solidity: function withdrawForAdministrator(address exitNcp, uint256 amount, address to) returns()
func (_NCPExitImp *NCPExitImpSession) WithdrawForAdministrator(exitNcp common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.WithdrawForAdministrator(&_NCPExitImp.TransactOpts, exitNcp, amount, to)
}

// WithdrawForAdministrator is a paid mutator transaction binding the contract method 0x00fc5701.
//
// Solidity: function withdrawForAdministrator(address exitNcp, uint256 amount, address to) returns()
func (_NCPExitImp *NCPExitImpTransactorSession) WithdrawForAdministrator(exitNcp common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _NCPExitImp.Contract.WithdrawForAdministrator(&_NCPExitImp.TransactOpts, exitNcp, amount, to)
}

// WithdrawForUser is a paid mutator transaction binding the contract method 0xb007a5ad.
//
// Solidity: function withdrawForUser(address exitNcp, address exitUser, uint256 amount) returns()
func (_NCPExitImp *NCPExitImpTransactor) WithdrawForUser(opts *bind.TransactOpts, exitNcp common.Address, exitUser common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NCPExitImp.contract.Transact(opts, "withdrawForUser", exitNcp, exitUser, amount)
}

// WithdrawForUser is a paid mutator transaction binding the contract method 0xb007a5ad.
//
// Solidity: function withdrawForUser(address exitNcp, address exitUser, uint256 amount) returns()
func (_NCPExitImp *NCPExitImpSession) WithdrawForUser(exitNcp common.Address, exitUser common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NCPExitImp.Contract.WithdrawForUser(&_NCPExitImp.TransactOpts, exitNcp, exitUser, amount)
}

// WithdrawForUser is a paid mutator transaction binding the contract method 0xb007a5ad.
//
// Solidity: function withdrawForUser(address exitNcp, address exitUser, uint256 amount) returns()
func (_NCPExitImp *NCPExitImpTransactorSession) WithdrawForUser(exitNcp common.Address, exitUser common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NCPExitImp.Contract.WithdrawForUser(&_NCPExitImp.TransactOpts, exitNcp, exitUser, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NCPExitImp *NCPExitImpTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NCPExitImp.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NCPExitImp *NCPExitImpSession) Receive() (*types.Transaction, error) {
	return _NCPExitImp.Contract.Receive(&_NCPExitImp.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NCPExitImp *NCPExitImpTransactorSession) Receive() (*types.Transaction, error) {
	return _NCPExitImp.Contract.Receive(&_NCPExitImp.TransactOpts)
}

// NCPExitImpAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the NCPExitImp contract.
type NCPExitImpAdminChangedIterator struct {
	Event *NCPExitImpAdminChanged // Event containing the contract specifics and raw log

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
func (it *NCPExitImpAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NCPExitImpAdminChanged)
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
		it.Event = new(NCPExitImpAdminChanged)
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
func (it *NCPExitImpAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NCPExitImpAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NCPExitImpAdminChanged represents a AdminChanged event raised by the NCPExitImp contract.
type NCPExitImpAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NCPExitImp *NCPExitImpFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*NCPExitImpAdminChangedIterator, error) {

	logs, sub, err := _NCPExitImp.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &NCPExitImpAdminChangedIterator{contract: _NCPExitImp.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NCPExitImp *NCPExitImpFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *NCPExitImpAdminChanged) (event.Subscription, error) {

	logs, sub, err := _NCPExitImp.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NCPExitImpAdminChanged)
				if err := _NCPExitImp.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_NCPExitImp *NCPExitImpFilterer) ParseAdminChanged(log types.Log) (*NCPExitImpAdminChanged, error) {
	event := new(NCPExitImpAdminChanged)
	if err := _NCPExitImp.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NCPExitImpBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the NCPExitImp contract.
type NCPExitImpBeaconUpgradedIterator struct {
	Event *NCPExitImpBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *NCPExitImpBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NCPExitImpBeaconUpgraded)
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
		it.Event = new(NCPExitImpBeaconUpgraded)
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
func (it *NCPExitImpBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NCPExitImpBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NCPExitImpBeaconUpgraded represents a BeaconUpgraded event raised by the NCPExitImp contract.
type NCPExitImpBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NCPExitImp *NCPExitImpFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*NCPExitImpBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NCPExitImp.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &NCPExitImpBeaconUpgradedIterator{contract: _NCPExitImp.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NCPExitImp *NCPExitImpFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *NCPExitImpBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NCPExitImp.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NCPExitImpBeaconUpgraded)
				if err := _NCPExitImp.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_NCPExitImp *NCPExitImpFilterer) ParseBeaconUpgraded(log types.Log) (*NCPExitImpBeaconUpgraded, error) {
	event := new(NCPExitImpBeaconUpgraded)
	if err := _NCPExitImp.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NCPExitImpInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NCPExitImp contract.
type NCPExitImpInitializedIterator struct {
	Event *NCPExitImpInitialized // Event containing the contract specifics and raw log

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
func (it *NCPExitImpInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NCPExitImpInitialized)
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
		it.Event = new(NCPExitImpInitialized)
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
func (it *NCPExitImpInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NCPExitImpInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NCPExitImpInitialized represents a Initialized event raised by the NCPExitImp contract.
type NCPExitImpInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NCPExitImp *NCPExitImpFilterer) FilterInitialized(opts *bind.FilterOpts) (*NCPExitImpInitializedIterator, error) {

	logs, sub, err := _NCPExitImp.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NCPExitImpInitializedIterator{contract: _NCPExitImp.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NCPExitImp *NCPExitImpFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NCPExitImpInitialized) (event.Subscription, error) {

	logs, sub, err := _NCPExitImp.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NCPExitImpInitialized)
				if err := _NCPExitImp.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NCPExitImp *NCPExitImpFilterer) ParseInitialized(log types.Log) (*NCPExitImpInitialized, error) {
	event := new(NCPExitImpInitialized)
	if err := _NCPExitImp.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NCPExitImpOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NCPExitImp contract.
type NCPExitImpOwnershipTransferredIterator struct {
	Event *NCPExitImpOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NCPExitImpOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NCPExitImpOwnershipTransferred)
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
		it.Event = new(NCPExitImpOwnershipTransferred)
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
func (it *NCPExitImpOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NCPExitImpOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NCPExitImpOwnershipTransferred represents a OwnershipTransferred event raised by the NCPExitImp contract.
type NCPExitImpOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NCPExitImp *NCPExitImpFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NCPExitImpOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NCPExitImp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NCPExitImpOwnershipTransferredIterator{contract: _NCPExitImp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NCPExitImp *NCPExitImpFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NCPExitImpOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NCPExitImp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NCPExitImpOwnershipTransferred)
				if err := _NCPExitImp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NCPExitImp *NCPExitImpFilterer) ParseOwnershipTransferred(log types.Log) (*NCPExitImpOwnershipTransferred, error) {
	event := new(NCPExitImpOwnershipTransferred)
	if err := _NCPExitImp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NCPExitImpSetRegistryIterator is returned from FilterSetRegistry and is used to iterate over the raw logs and unpacked data for SetRegistry events raised by the NCPExitImp contract.
type NCPExitImpSetRegistryIterator struct {
	Event *NCPExitImpSetRegistry // Event containing the contract specifics and raw log

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
func (it *NCPExitImpSetRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NCPExitImpSetRegistry)
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
		it.Event = new(NCPExitImpSetRegistry)
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
func (it *NCPExitImpSetRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NCPExitImpSetRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NCPExitImpSetRegistry represents a SetRegistry event raised by the NCPExitImp contract.
type NCPExitImpSetRegistry struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetRegistry is a free log retrieval operation binding the contract event 0x278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd.
//
// Solidity: event SetRegistry(address indexed addr)
func (_NCPExitImp *NCPExitImpFilterer) FilterSetRegistry(opts *bind.FilterOpts, addr []common.Address) (*NCPExitImpSetRegistryIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _NCPExitImp.contract.FilterLogs(opts, "SetRegistry", addrRule)
	if err != nil {
		return nil, err
	}
	return &NCPExitImpSetRegistryIterator{contract: _NCPExitImp.contract, event: "SetRegistry", logs: logs, sub: sub}, nil
}

// WatchSetRegistry is a free log subscription operation binding the contract event 0x278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd.
//
// Solidity: event SetRegistry(address indexed addr)
func (_NCPExitImp *NCPExitImpFilterer) WatchSetRegistry(opts *bind.WatchOpts, sink chan<- *NCPExitImpSetRegistry, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _NCPExitImp.contract.WatchLogs(opts, "SetRegistry", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NCPExitImpSetRegistry)
				if err := _NCPExitImp.contract.UnpackLog(event, "SetRegistry", log); err != nil {
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
func (_NCPExitImp *NCPExitImpFilterer) ParseSetRegistry(log types.Log) (*NCPExitImpSetRegistry, error) {
	event := new(NCPExitImpSetRegistry)
	if err := _NCPExitImp.contract.UnpackLog(event, "SetRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NCPExitImpUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NCPExitImp contract.
type NCPExitImpUpgradedIterator struct {
	Event *NCPExitImpUpgraded // Event containing the contract specifics and raw log

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
func (it *NCPExitImpUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NCPExitImpUpgraded)
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
		it.Event = new(NCPExitImpUpgraded)
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
func (it *NCPExitImpUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NCPExitImpUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NCPExitImpUpgraded represents a Upgraded event raised by the NCPExitImp contract.
type NCPExitImpUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NCPExitImp *NCPExitImpFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NCPExitImpUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NCPExitImp.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NCPExitImpUpgradedIterator{contract: _NCPExitImp.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NCPExitImp *NCPExitImpFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NCPExitImpUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NCPExitImp.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NCPExitImpUpgraded)
				if err := _NCPExitImp.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_NCPExitImp *NCPExitImpFilterer) ParseUpgraded(log types.Log) (*NCPExitImpUpgraded, error) {
	event := new(NCPExitImpUpgraded)
	if err := _NCPExitImp.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
