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
	Bin: "0x608060405234801561001057600080fd5b5060405161070238038061070283398101604081905261002f91610304565b80604051806020016040528060008152506100528282600061005a60201b60201c565b5050506103a8565b61006383610090565b6000825111806100705750805b1561008b5761008983836100d060201b61008b1760201c565b505b505050565b610099816100fc565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606100f583836040518060600160405280602781526020016106db602791396101ce565b9392505050565b61010f8161024760201b6100b71760201c565b6101765760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b806101ad7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61025660201b6100c61760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b6060600080856001600160a01b0316856040516101eb9190610359565b600060405180830381855af49150503d8060008114610226576040519150601f19603f3d011682016040523d82523d6000602084013e61022b565b606091505b50909250905061023d86838387610259565b9695505050505050565b6001600160a01b03163b151590565b90565b606083156102c85782516000036102c1576001600160a01b0385163b6102c15760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161016d565b50816102d2565b6102d283836102da565b949350505050565b8151156102ea5781518083602001fd5b8060405162461bcd60e51b815260040161016d9190610375565b60006020828403121561031657600080fd5b81516001600160a01b03811681146100f557600080fd5b60005b83811015610348578181015183820152602001610330565b838111156100895750506000910152565b6000825161036b81846020870161032d565b9190910192915050565b602081526000825180602084015261039481604085016020870161032d565b601f01601f19169190910160400192915050565b610324806103b76000396000f3fe6080604052600436106100225760003560e01c80635c60da1b1461003957610031565b366100315761002f61006a565b005b61002f61006a565b34801561004557600080fd5b5061004e61007c565b6040516001600160a01b03909116815260200160405180910390f35b61007a6100756100c9565b6100fc565b565b60006100866100c9565b905090565b60606100b083836040518060600160405280602781526020016102c860279139610120565b9392505050565b6001600160a01b03163b151590565b90565b60006100867f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b3660008037600080366000845af43d6000803e80801561011b573d6000f35b3d6000fd5b6060600080856001600160a01b03168560405161013d9190610278565b600060405180830381855af49150503d8060008114610178576040519150601f19603f3d011682016040523d82523d6000602084013e61017d565b606091505b509150915061018e86838387610198565b9695505050505050565b6060831561020c578251600003610205576001600160a01b0385163b6102055760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064015b60405180910390fd5b5081610216565b610216838361021e565b949350505050565b81511561022e5781518083602001fd5b8060405162461bcd60e51b81526004016101fc9190610294565b60005b8381101561026357818101518382015260200161024b565b83811115610272576000848401525b50505050565b6000825161028a818460208701610248565b9190910192915050565b60208152600082518060208401526102b3816040850160208701610248565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212202bc8194e40b4067bdacc5e803dad2b88ec6e4472b072326b764054e68c4a96d864736f6c634300080e0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
}

// GovABI is the input ABI used to generate the binding from.
// Deprecated: Use GovMetaData.ABI instead.
var GovABI = GovMetaData.ABI

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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"envName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"envType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"envVal\",\"type\":\"bytes\"}],\"name\":\"EnvChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"GovDataMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"MemberAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVoter\",\"type\":\"address\"}],\"name\":\"MemberChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"MemberRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"MemberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ballotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"NotApplicable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"SetProposalTimePeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_DURATION_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BALLOT_DURATION_MIN_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BALLOT_DURATION_MIN_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BALLOT_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_FEE_MAX_CHANGE_RATE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCKS_PER_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_CREATION_TIME_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_GASLIMIT_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_AMOUNT_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_MAINTANANCE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_METHOD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ECOSYSTEM_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ENV_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GASLIMIT_AND_BASE_FEE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAS_TARGET_PERCENTAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAINTENANCE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BASE_FEE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_IDLE_BLOCK_INTERVAL_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PRIORITY_FEE_PER_GAS_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_POOL_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MIN_MAX_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MIN_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_REWARD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"port\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"internalType\":\"structGovImp.MemberInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"addProposalToAddMember\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"envName\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"envType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"envVal\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"addProposalToChangeEnv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"addProposalToChangeGov\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"port\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"internalType\":\"structGovImp.MemberInfo\",\"name\":\"newInfo\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"oldStaker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashing\",\"type\":\"uint256\"}],\"name\":\"addProposalToChangeMember\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashing\",\"type\":\"uint256\"}],\"name\":\"addProposalToRemoveMember\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballotLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkUnfinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeEndedVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBallotInVoting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxStaking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxVotingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getMemberFromNodeIdx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMemberLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStaking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinVotingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"port\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getNodeIdxFromMember\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getStakerAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getVoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"port\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"oldModifiedBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"}],\"name\":\"initMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initOnce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isStaker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastAddProposalTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldGov\",\"type\":\"address\"}],\"name\":\"migrateFromLegacy\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modifiedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposal_time_period\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reg\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"setProposalTimePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotIdx\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approval\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voterIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000e8565b600054610100900460ff1615620000935760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015620000e6576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6080516185d86200012060003960008181611e6501528181611ea5015281816129eb01528181612a2b0152612aa701526185d86000f3fe6080604052600436106104675760003560e01c806372016f751161024a578063af6af2ff11610139578063cec5b622116100b6578063e27bdaef1161007a578063e27bdaef14610fb3578063e75235b814610fd3578063e9523fb514610fe8578063f2fde38b14610ffe578063f38ecf471461101e57600080fd5b8063cec5b62214610f10578063d07bff0c14610f3d578063d6f9cfce14610f73578063d965ea0014610f89578063de09b37714610f9e57600080fd5b8063c6713baf116100fd578063c6713baf14610e3d578063c7d3da3414610e71578063c9d27afe14610ea5578063ce04b9d414610ec5578063ce6a54ff14610eda57600080fd5b8063af6af2ff14610d58578063b128f88014610d6d578063c00ace6c14610da1578063c0b4fe1514610dd5578063c42a0abc14610e0957600080fd5b8063a0c12683116101c7578063a8915a3e1161018b578063a8915a3e14610c81578063a91ee0dc14610ca1578063a9b629b214610cc1578063aaf0dd3614610cf5578063ab3545e514610d2257600080fd5b8063a0c1268314610ba8578063a230c52414610bd5578063a6868b7d14610bf5578063a7771ee314610c29578063a78a818814610c6157600080fd5b8063894f51111161020e578063894f511114610b105780638d39e33a14610b305780638da5cb5b14610b50578063918f867414610b6e5780639986e4b914610b8457600080fd5b806372016f7514610a70578063738fdd1a14610a855780637b2bfb0114610aa55780637bf4653014610ad95780637d10dd1b14610afa57600080fd5b8063397e38e7116103665780635aaa4040116102e35780636d583ca7116102a75780636d583ca71461099b5780636f1e8533146109cf5780636f6de96d14610a075780636fde207a14610a27578063715018a614610a5b57600080fd5b80635aaa4040146108e05780636167eb45146108f5578063656e3052146109295780636ba991811461095d5780636c78d2cf1461097257600080fd5b80634d5ce0381161032a5780634d5ce038146108445780634f0f4aa9146108645780634f1ef2861461089457806352d1902d146108a75780635a731cca146108bc57600080fd5b8063397e38e71461077a5780633f35c8fe1461079a57806340690353146107ce5780634bd1ed76146107ee5780634d273e281461081057600080fd5b80631e0cba0d116103f457806334125c84116103b857806334125c84146106c6578063351bacda146106e65780633659cfe61461070657806336e83d8314610726578063382944191461074657600080fd5b80631e0cba0d14610609578063238737b614610627578063278bb12a1461065b5780632f40992e1461068f5780633310569c146106b057600080fd5b806315bf6b4d1161043b57806315bf6b4d1461053457806316fbe831146105825780631c12b030146105995780631c150171146105be5780631c4b774b146105d357600080fd5b806215a73b1461046c5780630b1d39b8146104b35780630efa4909146104e7578063139d9dd314610507575b600080fd5b34801561047857600080fd5b506104a07f0c4fbe9dc9de15dd7c0d064975ee1a2f2f9b954fa0e65d4f6cddba94884bdc3e81565b6040519081526020015b60405180910390f35b3480156104bf57600080fd5b506104a07fdd5a41a7fc01f5c6d30816b17f638d6531625f1e1eaa599673ab2f6079f2dd9d81565b3480156104f357600080fd5b506104a061050236600461769a565b611052565b34801561051357600080fd5b506104a06105223660046176f2565b60ab6020526000908152604090205481565b34801561054057600080fd5b5061056a61054f36600461770f565b6000908152607360205260409020546001600160a01b031690565b6040516001600160a01b0390911681526020016104aa565b34801561058e57600080fd5b506105976113db565b005b3480156105a557600080fd5b506105ae61171a565b60405190151581526020016104aa565b3480156105ca57600080fd5b506104a061176b565b3480156105df57600080fd5b5061056a6105ee36600461770f565b6000908152606a60205260409020546001600160a01b031690565b34801561061557600080fd5b506104a0665374616b696e6760c81b81565b34801561063357600080fd5b506104a07f1d36f8ce53f59e624857e1d8dc7932d19981a2ea1b8faa4eb8ff843fc3e5a27881565b34801561066757600080fd5b506104a07f9b2e0c7fdae148f225bae7deb92d7e7bd24bb77edb12956e8fa7204900dd8a2281565b34801561069b57600080fd5b506104a06914995dd85c99141bdbdb60b21b81565b3480156106bc57600080fd5b506104a060aa5481565b3480156106d257600080fd5b506104a06845636f73797374656d60b81b81565b3480156106f257600080fd5b50610597610701366004617728565b6117db565b34801561071257600080fd5b506105976107213660046176f2565b611e5b565b34801561073257600080fd5b506104a061074136600461788e565b611f40565b34801561075257600080fd5b506104a07fbe90e461bbdb9a95a694f7796912ea04244caf7f5b60ad7ded17e16821d3e44c81565b34801561078657600080fd5b506105976107953660046178c2565b612458565b3480156107a657600080fd5b506104a07f2a268972a70c8c688b62366bdfdd9bb09cf19d3e5b6e7e7bb158e671ffdcedd281565b3480156107da57600080fd5b506104a06107e9366004617904565b612553565b3480156107fa57600080fd5b506104a06a4d61696e74656e616e636560a81b81565b34801561081c57600080fd5b506104a07f77884798208df1e64f70968be41ef2d3211ec53613397ca59313416813df088881565b34801561085057600080fd5b506105ae61085f3660046176f2565b6127e7565b34801561087057600080fd5b5061088461087f36600461770f565b612804565b6040516104aa94939291906179da565b6105976108a2366004617a25565b6129e1565b3480156108b357600080fd5b506104a0612a9a565b3480156108c857600080fd5b506104a06c14dd185ada5b99d4995dd85c99609a1b81565b3480156108ec57600080fd5b506104a0612b4d565b34801561090157600080fd5b506104a07f9f1de481f899d76889aa8a2b44cc7b604d42691aa93d4ba618a1a1fd439f505081565b34801561093557600080fd5b506104a07fe10074dceffb75f13bf0ce50145afd35182d63796823f1280ce40e01c19109e781565b34801561096957600080fd5b50610597612b94565b34801561097e57600080fd5b506104a07111dbdd995c9b985b98d950dbdb9d1c9858dd60721b81565b3480156109a757600080fd5b506104a07fc69fc6b7d0efc934fd5a3581c7253a7107a952526bb6dbcd814ef8d8dae1f44a81565b3480156109db57600080fd5b506105ae6109ea3660046176f2565b6001600160a01b03166000908152606d6020526040902054151590565b348015610a1357600080fd5b5061056a610a223660046176f2565b612c82565b348015610a3357600080fd5b506104a07f0b09c9badbbeb6c813a598ee910770a39ccda797a1940439bb6e47fc6c87548b81565b348015610a6757600080fd5b50610597612cf1565b348015610a7c57600080fd5b506074546104a0565b348015610a9157600080fd5b5060655461056a906001600160a01b031681565b348015610ab157600080fd5b506104a07f9346226931826838eedd13d9677fa33551e7c81f604b171ef3fac355458da9aa81565b348015610ae557600080fd5b506104a069456e7653746f7261676560b01b81565b348015610b0657600080fd5b506104a060665481565b348015610b1c57600080fd5b506104a0610b2b366004617a74565b612d05565b348015610b3c57600080fd5b506104a0610b4b3660046176f2565b613063565b348015610b5c57600080fd5b506033546001600160a01b031661056a565b348015610b7a57600080fd5b506104a061271081565b348015610b9057600080fd5b506104a06c42616c6c6f7453746f7261676560981b81565b348015610bb457600080fd5b506104a0610bc33660046176f2565b606d6020526000908152604090205481565b348015610be157600080fd5b506105ae610bf03660046176f2565b6138f5565b348015610c0157600080fd5b506104a07f6c6f69f426081752a5d3e73746599acd2a4cb145d5de4203ca1e3473b281680b81565b348015610c3557600080fd5b506105ae610c443660046176f2565b6001600160a01b0316600090815260686020526040902054151590565b348015610c6d57600080fd5b506104a0610c7c366004617ae8565b613938565b348015610c8d57600080fd5b50610597610c9c366004617b46565b614094565b348015610cad57600080fd5b50610597610cbc3660046176f2565b6144fd565b348015610ccd57600080fd5b506104a07f89dd490ecaf395283ed4ff2fd9557ca767fc425dce063451a9b0da6d72f600c381565b348015610d0157600080fd5b506104a0610d103660046176f2565b606b6020526000908152604090205481565b348015610d2e57600080fd5b5061056a610d3d36600461770f565b6000908152606c60205260409020546001600160a01b031690565b348015610d6457600080fd5b506104a06145a5565b348015610d7957600080fd5b506104a07f829561ab7af084b7efc6600518d2df79b8d95f3f4c3a550f54f8f7ec7d2b805781565b348015610dad57600080fd5b506104a07f18ad4415ef4a621ce1a136395c51ab6c3712bb2e24b79d526059925cea58dcb881565b348015610de157600080fd5b506104a07f8086da5becff4dfac91a3105821b361078d2d4abba0ccc2401b974cf0dcf05c181565b348015610e1557600080fd5b506104a07fb38b2c133e931937bd95337c65c8aefa7040ed64bbed732e3e29a4944c65747381565b348015610e4957600080fd5b506104a07fc9e15e34073efbcd0328740feaf503caac9124b55b38c73d1a97b53da80a2f6081565b348015610e7d57600080fd5b506104a07f04f7b94450bbcad85f37ea47cd1062728f884bb2040e357738f8fd53056134bc81565b348015610eb157600080fd5b50610597610ec0366004617c00565b6145ec565b348015610ed157600080fd5b506104a0614793565b348015610ee657600080fd5b506104a0610ef53660046176f2565b6001600160a01b031660009081526072602052604090205490565b348015610f1c57600080fd5b506104a0610f2b3660046176f2565b60686020526000908152604090205481565b348015610f4957600080fd5b5061056a610f5836600461770f565b6000908152606760205260409020546001600160a01b031690565b348015610f7f57600080fd5b506104a060755481565b348015610f9557600080fd5b506069546104a0565b348015610faa57600080fd5b506077546104a0565b348015610fbf57600080fd5b50610597610fce36600461770f565b6147da565b348015610fdf57600080fd5b506113896104a0565b348015610ff457600080fd5b506104a060765481565b34801561100a57600080fd5b506105976110193660046176f2565b614860565b34801561102a57600080fd5b506104a07f7c1150f0e1a39ff55552d52764f97e6c387e2a247e1df344369f122c4254be2f81565b600061105c6148d6565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa1580156110a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110c69190617c30565b6110eb5760405162461bcd60e51b81526004016110e290617c4d565b60405180910390fd5b60006110f633612c82565b60aa546001600160a01b038216600090815260ab6020526040902054919250906111209042617c8a565b101561113e5760405162461bcd60e51b81526004016110e290617ca1565b600061114933612c82565b9050611153612b4d565b61115c826148f6565b11158015611179575061116d6145a5565b611176826148f6565b10155b6111955760405162461bcd60e51b81526004016110e290617cd8565b6001600160a01b0386166111eb5760405162461bcd60e51b815260206004820152601d60248201527f496d706c656d656e746174696f6e2063616e6e6f74206265207a65726f00000060448201526064016110e2565b6111f361496d565b6001600160a01b0316866001600160a01b03160361124b5760405162461bcd60e51b815260206004820152601560248201527453616d6520636f6e7472616374206164647265737360581b60448201526064016110e2565b856001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156112a5575060408051601f3d908101601f191682019092526112a291810190617d0f565b60015b6112c15760405162461bcd60e51b81526004016110e290617d28565b60008051602061851c83398151915281146112ee5760405162461bcd60e51b81526004016110e290617d76565b506075546112fd906001617dbf565b9250611307614989565b6001600160a01b0316630a3a63fe60755460016113249190617dbf565b60046040516001600160e01b031960e085901b16815260048101929092526024820152604481018790523360648201526001600160a01b038916608482015260a4016020604051808303816000875af1158015611385573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113a99190617d0f565b506113b483866149a4565b5060758290556001600160a01b0316600090815260ab602052604090204290559392505050565b600054600290610100900460ff161580156113fd575060005460ff8083169116105b6114195760405162461bcd60e51b81526004016110e290617dd7565b6000805461ffff191660ff831617610100179055611435614a0f565b60005b6069548110156116e65760008181526071602052604080822081516080810190925280548290829061146990617e25565b80601f016020809104026020016040519081016040528092919081815260200182805461149590617e25565b80156114e25780601f106114b7576101008083540402835291602001916114e2565b820191906000526020600020905b8154815290600101906020018083116114c557829003601f168201915b505050505081526020016001820180546114fb90617e25565b80601f016020809104026020016040519081016040528092919081815260200182805461152790617e25565b80156115745780601f1061154957610100808354040283529160200191611574565b820191906000526020600020905b81548152906001019060200180831161155757829003601f168201915b5050505050815260200160028201805461158d90617e25565b80601f01602080910402602001604051908101604052809291908181526020018280546115b990617e25565b80156116065780601f106115db57610100808354040283529160200191611606565b820191906000526020600020905b8154815290600101906020018083116115e957829003601f168201915b5050505050815260200160038201548152505090506001606e82600001516040516116319190617e59565b90815260405160209181900382018120805460ff191693151593909317909255820151600191606f9161166391617e59565b908152602001604051809103902060006101000a81548160ff021916908315150217905550600160706000836040015184606001516040516020016116a9929190617e75565b60408051808303601f19018152918152815160209283012083529082019290925201600020805460ff191691151591909117905550600101611438565b506000805461ff001916905560405160ff8216815260008051602061853c833981519152906020015b60405180910390a150565b6000607754600014611765576000611733607754614a69565b509150506000611744607754614af1565b50915050600119820161176257428110156117625760009250505090565b50505b50600190565b6000611775614b6c565b6001600160a01b03166333be496e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117d69190617d0f565b905090565b600054610100900460ff16158080156117fb5750600054600160ff909116105b806118155750303b158015611815575060005460ff166001145b6118315760405162461bcd60e51b81526004016110e290617dd7565b6000805460ff191660011790558015611854576000805461ff0019166101001790555b61185c614b84565b611864614bb3565b61186d846144fd565b43606655600061187b614be2565b90506118856145a5565b841015801561189b575083611898612b4d565b10155b6118b75760405162461bcd60e51b81526004016110e290617e97565b600080600060608060606000806000905060008060208d0191508c51826118de9190617dbf565b90505b80821015611e0c57815199506118f8602083617dbf565b915080821061190657600080fd5b81519850611915602083617dbf565b915080821061192357600080fd5b81519750611932602083617dbf565b915080821061194057600080fd5b819650865160206119519190617dbf565b61195b9083617dbf565b915080821061196957600080fd5b8195508551602061197a9190617dbf565b6119849083617dbf565b915080821061199257600080fd5b819450845160206119a39190617dbf565b6119ad9083617dbf565b91508082106119bb57600080fd5b815193506119ca602083617dbf565b91506119d7600184617dbf565b92506119e28a6138f5565b1580156119f557506119f3896138f5565b155b8015611a075750611a05886127e7565b155b611a445760405162461bcd60e51b815260206004820152600e60248201526d20b63932b0b23c9036b2b6b132b960911b60448201526064016110e2565b886067600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606860008b6001600160a01b03166001600160a01b031681526020019081526020016000208190555087606a600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606b60008a6001600160a01b03166001600160a01b031681526020019081526020016000208190555089606c600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606d60008c6001600160a01b03166001600160a01b0316815260200190815260200160002081905550886001600160a01b03168a6001600160a01b03167f6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba660405160405180910390a36040516325d998bb60e01b81526001600160a01b038b811660048301528f91908d16906325d998bb90602401602060405180830381865afa158015611bf3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c179190617d0f565b1015611c5c5760405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e74207374616b696e6760601b60448201526064016110e2565b611c6887878787614bf7565b611c845760405162461bcd60e51b81526004016110e290617ec4565b611c8e8a8f614ca6565b600083815260716020908152604090912088519091611cb19183918b0190617444565b508651611cc790600183019060208a0190617444565b508551611cdd9060028301906020890190617444565b508481600301819055506001606e89604051611cf99190617e59565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f90611d2b908a90617e59565b908152602001604051809103902060006101000a81548160ff0219169083151502179055506001607060008888604051602001611d69929190617e75565b60405160208183030381529060405280519060200120815260200190815260200160002060006101000a81548160ff0219169083151502179055508a6073600086815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555083607260008d6001600160a01b03166001600160a01b0316815260200190815260200160002081905550506118e1565b50506069819055607455505086159550611e55945050505050576000805461ff00191690556040516001815260008051602061853c833981519152906020015b60405180910390a15b50505050565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003611ea35760405162461bcd60e51b81526004016110e290617ef2565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611ed561496d565b6001600160a01b031614611efb5760405162461bcd60e51b81526004016110e290617f3e565b611f0481614ce5565b611f3d8160005b6040519080825280601f01601f191660200182016040528015611f35576020820181803683370190505b506000614d73565b50565b6000611f4a6148d6565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015611f90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fb49190617c30565b611fd05760405162461bcd60e51b81526004016110e290617c4d565b6000611fdb33612c82565b60aa546001600160a01b038216600090815260ab6020526040902054919250906120059042617c8a565b10156120235760405162461bcd60e51b81526004016110e290617ca1565b600061202e33612c82565b9050612038612b4d565b612041826148f6565b1115801561205e57506120526145a5565b61205b826148f6565b10155b61207a5760405162461bcd60e51b81526004016110e290617cd8565b602084015184906001600160a01b03166120c65760405162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b2103b37ba32b960991b60448201526064016110e2565b60008160600151511161210f5760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964206e6f6465206e616d6560781b60448201526064016110e2565b60008160a0015151116121565760405162461bcd60e51b815260206004820152600f60248201526e0496e76616c6964206e6f646520495608c1b60448201526064016110e2565b60008160c001511161219e5760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a59081b9bd919481c1bdc9d607a1b60448201526064016110e2565b6000816080015151116121e85760405162461bcd60e51b8152602060048201526012602482015271496e76616c6964206e6f646520656e6f646560701b60448201526064016110e2565b6000816101000151511161222d5760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964206d656d6f60a01b60448201526064016110e2565b6000816101200151116122755760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b210323ab930ba34b7b760811b60448201526064016110e2565b61227d6145a5565b8160e001511015801561229b5750612293612b4d565b8160e0015111155b6122dd5760405162461bcd60e51b8152602060048201526013602482015272125b9d985b1a59081b1bd8dac8105b5bdd5b9d606a1b60448201526064016110e2565b84516122e8906138f5565b1580156122fd575084516122fb906127e7565b155b61233a5760405162461bcd60e51b815260206004820152600e60248201526d20b63932b0b23c9036b2b6b132b960911b60448201526064016110e2565b84602001516001600160a01b031685600001516001600160a01b031614801561237c575084604001516001600160a01b031685600001516001600160a01b0316145b6123be5760405162461bcd60e51b815260206004820152601360248201527229ba30b5b2b91034b9903737ba103b37ba32b960691b60448201526064016110e2565b6123da856060015186608001518760a001518860c00151614bf7565b6123f65760405162461bcd60e51b81526004016110e290617ec4565b607554612404906001617dbf565b935061241584600133600089614e5a565b612423848660e00151614efd565b612432848661010001516149a4565b505060758290556001600160a01b0316600090815260ab60205260409020429055919050565b600054610100900460ff16158080156124785750600054600160ff909116105b806124925750303b158015612492575060005460ff166001145b6124ae5760405162461bcd60e51b81526004016110e290617dd7565b6000805460ff1916600117905580156124d1576000805461ff0019166101001790555b6124d9614b84565b6124e1614bb3565b6124ea846144fd565b60668390556124f882614860565b60405133907fab2db0a6f442428b686ffa80eadcaabe7d5ee00049c6ae888a237edd3238d85690600090a28015611e55576000805461ff00191690556040516001815260008051602061853c83398151915290602001611e4c565b600061255d6148d6565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa1580156125a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125c79190617c30565b6125e35760405162461bcd60e51b81526004016110e290617c4d565b60006125ee33612c82565b60aa546001600160a01b038216600090815260ab6020526040902054919250906126189042617c8a565b10156126365760405162461bcd60e51b81526004016110e290617ca1565b600061264133612c82565b905061264b612b4d565b612654826148f6565b1115801561267157506126656145a5565b61266e826148f6565b10155b61268d5760405162461bcd60e51b81526004016110e290617cd8565b8660011115801561269f575060098711155b6126da5760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964207479706560a01b60448201526064016110e2565b6126e48887614f3b565b6127205760405162461bcd60e51b815260206004820152600d60248201526c496e76616c69642076616c756560981b60448201526064016110e2565b60755461272e906001617dbf565b9250612738614989565b6001600160a01b0316634a57823e84600587338d8d8d6040518863ffffffff1660e01b81526004016127709796959493929190617f8a565b6020604051808303816000875af115801561278f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127b39190617d0f565b506127be83866149a4565b5060758290556001600160a01b0316600090815260ab6020526040902042905595945050505050565b6001600160a01b03166000908152606b6020526040902054151590565b60008181526071602052604081206003810154815460609384938493919290916001830191600284019190849061283a90617e25565b80601f016020809104026020016040519081016040528092919081815260200182805461286690617e25565b80156128b35780601f10612888576101008083540402835291602001916128b3565b820191906000526020600020905b81548152906001019060200180831161289657829003601f168201915b505050505093508280546128c690617e25565b80601f01602080910402602001604051908101604052809291908181526020018280546128f290617e25565b801561293f5780601f106129145761010080835404028352916020019161293f565b820191906000526020600020905b81548152906001019060200180831161292257829003601f168201915b5050505050925081805461295290617e25565b80601f016020809104026020016040519081016040528092919081815260200182805461297e90617e25565b80156129cb5780601f106129a0576101008083540402835291602001916129cb565b820191906000526020600020905b8154815290600101906020018083116129ae57829003601f168201915b5050505050915093509350935093509193509193565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003612a295760405162461bcd60e51b81526004016110e290617ef2565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612a5b61496d565b6001600160a01b031614612a815760405162461bcd60e51b81526004016110e290617f3e565b612a8a82614ce5565b612a9682826001614d73565b5050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612b3a5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016110e2565b5060008051602061851c83398151915290565b6000612b57614b6c565b6001600160a01b031663737c59b86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117b2573d6000803e3d6000fd5b612b9c6148d6565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015612be2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c069190617c30565b612c225760405162461bcd60e51b81526004016110e290617c4d565b612c2a61171a565b15612c6d5760405162461bcd60e51b8152602060048201526013602482015272159bdd1a5b99c81a5cc81b9bdd08195b991959606a1b60448201526064016110e2565b607754612c7b906004614fba565b6000607755565b6001600160a01b0381166000908152606d602052604081205415612ca4575090565b6001600160a01b03821660009081526068602052604090205415612cec57506001600160a01b038082166000908152606860209081526040808320548352606c909152902054165b919050565b612cf9614a0f565b612d036000614ff8565b565b6000612d0f6148d6565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015612d55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d799190617c30565b612d955760405162461bcd60e51b81526004016110e290617c4d565b6000612da033612c82565b60aa546001600160a01b038216600090815260ab602052604090205491925090612dca9042617c8a565b1015612de85760405162461bcd60e51b81526004016110e290617ca1565b6000612df333612c82565b9050612dfd612b4d565b612e06826148f6565b11158015612e235750612e176145a5565b612e20826148f6565b10155b612e3f5760405162461bcd60e51b81526004016110e290617cd8565b6001600160a01b038916612e875760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b60448201526064016110e2565b612e90896138f5565b612ec95760405162461bcd60e51b815260206004820152600a6024820152692737b716b6b2b6b132b960b11b60448201526064016110e2565b6001612ed460695490565b11612f215760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f742072656d6f7665206120736f6c65206d656d626572000000000060448201526064016110e2565b87612f2b8a6148f6565b1015612f8c5760405162461bcd60e51b815260206004820152602a60248201527f496e73756666696369656e742062616c616e636520746861742063616e206265604482015269103ab73637b1b5b2b21760b11b60648201526084016110e2565b607554612f9a906001617dbf565b604080516101408101825260008082526020808301829052828401829052835182815280820185526060840152835182815280820185526080840152835182815290810190935260a082019290925260c081019190915260e081018a905261010081018990526101208101889052909350613019846002338d85614e5a565b613023848a614efd565b61302d84896149a4565b61303884878761504a565b505060758290556001600160a01b0316600090815260ab602052604090204290559695505050505050565b60008054610100900460ff16158080156130845750600054600160ff909116105b8061309e5750303b15801561309e575060005460ff166001145b6130ba5760405162461bcd60e51b81526004016110e290617dd7565b6000805460ff1916600117905580156130dd576000805461ff0019166101001790555b6130e5614b84565b6130ed614bb3565b6000839050613157816001600160a01b031663738fdd1a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613133573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cbc9190617fd6565b436066819055506131c3816001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561319f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110199190617fd6565b60015b816001600160a01b031663d965ea006040518163ffffffff1660e01b8152600401602060405180830381865afa158015613204573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132289190617d0f565b81116137185760405163ab3545e560e01b8152600481018290526001600160a01b0383169063ab3545e590602401602060405180830381865afa158015613273573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132979190617fd6565b6000828152606c6020908152604080832080546001600160a01b0319166001600160a01b039586169081179091558352606d909152908190208390555163341effc360e21b8152600481018390529083169063d07bff0c90602401602060405180830381865afa15801561330f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133339190617fd6565b600082815260676020908152604080832080546001600160a01b0319166001600160a01b03958616908117909155835260689091529081902083905551631c4b774b60e01b81526004810183905290831690631c4b774b90602401602060405180830381865afa1580156133ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133cf9190617fd6565b6000828152606a6020908152604080832080546001600160a01b0319166001600160a01b03959095169485179055928252606b8152828220849055606984905582516080810184526060808252918101829052928301819052820152604051634f0f4aa960e01b8152600481018390526001600160a01b03841690634f0f4aa990602401600060405180830381865afa158015613470573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526134989190810190618043565b6060850181905260408501829052602085018390528385526134bd93929190856150b9565b6135095760405162461bcd60e51b815260206004820152601760248201527f6e6f646520696e666f206973206475706c69636174656400000000000000000060448201526064016110e2565b6001606e826000015160405161351f9190617e59565b90815260405160209181900382018120805460ff191693151593909317909255820151600191606f9161355191617e59565b908152602001604051809103902060006101000a81548160ff02191690831515021790555060016070600083604001518460600151604051602001613597929190617e75565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff1916941515949094179093558483526071815291208251805184936135ed928492910190617444565b5060208281015180516136069260018501920190617444565b5060408201518051613622916002840191602090910190617444565b50606091909101516003909101556000828152606c6020818152604080842080546001600160a01b039081168652607284528286208890558786528154607385529583902080546001600160a01b031916968216969096179095556074879055929091529054905163139d9dd360e01b815290821660048201529084169063139d9dd390602401602060405180830381865afa1580156136c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136ea9190617d0f565b6000838152606c60209081526040808320546001600160a01b0316835260ab909152902055506001016131c6565b50806001600160a01b0316633310569c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613757573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061377b9190617d0f565b60aa81905550806001600160a01b031663d6f9cfce6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156137bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137e39190617d0f565b607581905550806001600160a01b031663e9523fb56040518163ffffffff1660e01b8152600401602060405180830381865afa158015613827573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061384b9190617d0f565b607681905550806001600160a01b031663de09b3776040518163ffffffff1660e01b8152600401602060405180830381865afa15801561388f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138b39190617d0f565b607755506000915080156138ef576000805461ff00191690556040516001815260008051602061853c8339815191529060200160405180910390a15b50919050565b6001600160a01b0381166000908152606d602052604081205415158061393257506001600160a01b03821660009081526068602052604090205415155b92915050565b60006139426148d6565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015613988573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139ac9190617c30565b6139c85760405162461bcd60e51b81526004016110e290617c4d565b60006139d333612c82565b60aa546001600160a01b038216600090815260ab6020526040902054919250906139fd9042617c8a565b1015613a1b5760405162461bcd60e51b81526004016110e290617ca1565b6000613a2633612c82565b9050613a30612b4d565b613a39826148f6565b11158015613a565750613a4a6145a5565b613a53826148f6565b10155b613a725760405162461bcd60e51b81526004016110e290617cd8565b602087015187906001600160a01b0316613abe5760405162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b2103b37ba32b960991b60448201526064016110e2565b600081606001515111613b075760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964206e6f6465206e616d6560781b60448201526064016110e2565b60008160a001515111613b4e5760405162461bcd60e51b815260206004820152600f60248201526e0496e76616c6964206e6f646520495608c1b60448201526064016110e2565b60008160c0015111613b965760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a59081b9bd919481c1bdc9d607a1b60448201526064016110e2565b600081608001515111613be05760405162461bcd60e51b8152602060048201526012602482015271496e76616c6964206e6f646520656e6f646560701b60448201526064016110e2565b60008161010001515111613c255760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964206d656d6f60a01b60448201526064016110e2565b600081610120015111613c6d5760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b210323ab930ba34b7b760811b60448201526064016110e2565b613c756145a5565b8160e0015110158015613c935750613c8b612b4d565b8160e0015111155b613cd55760405162461bcd60e51b8152602060048201526013602482015272125b9d985b1a59081b1bd8dac8105b5bdd5b9d606a1b60448201526064016110e2565b6001600160a01b038716613d215760405162461bcd60e51b8152602060048201526013602482015272496e76616c6964206f6c64204164647265737360681b60448201526064016110e2565b613d2a876138f5565b613d635760405162461bcd60e51b815260206004820152600a6024820152692737b716b6b2b6b132b960b11b60448201526064016110e2565b6020808901516001600160a01b038981166000908152606d8452604080822054825260679094529290922054821691161480613db45750866001600160a01b031688602001516001600160a01b0316145b80613dde5750613dc788602001516138f5565b158015613dde5750613ddc88602001516127e7565b155b8015613e6057506040808901516001600160a01b038981166000908152606d6020908152848220548252606a90529290922054821691161480613e365750866001600160a01b031688604001516001600160a01b0316145b80613e605750613e4988604001516138f5565b158015613e605750613e5e88604001516127e7565b155b613e9f5760405162461bcd60e51b815260206004820152601060248201526f20b63932b0b23c90309036b2b6b132b960811b60448201526064016110e2565b336001600160a01b038816148015613ec3575087516001600160a01b038881169116145b15613f185785158015613ed4575084155b613f135760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a59081c1c9bdc1bdcd85b60821b60448201526064016110e2565b613fc7565b87516001600160a01b03888116911614613fc757613f346145a5565b613f3e8688617dbf565b1115613fc75760405162461bcd60e51b815260206004820152604c60248201527f496e76616c696420616d6f756e743a2028756e6c6f636b416d6f756e74202b2060448201527f736c617368696e6729206d75737420626520657175616c206f72206c6f77207460648201526b6f206d696e5374616b696e6760a01b608482015260a4016110e2565b607554613fd5906001617dbf565b9350613fe5846003338a8c614e5a565b613ff3848960e00151614efd565b614002848961010001516149a4565b61400d84878761504a565b6075849055336001600160a01b038816148015614036575087516001600160a01b038881169116145b1561407057600061404685614af1565b925050506140608542834261405b9190617dbf565b6151e2565b61406e856003600180615227565b505b50506001600160a01b0316600090815260ab60205260409020429055949350505050565b600054610100900460ff16158080156140b45750600054600160ff909116105b806140ce5750303b1580156140ce575060005460ff166001145b6140ea5760405162461bcd60e51b81526004016110e290617dd7565b6000805460ff19166001179055801561410d576000805461ff0019166101001790555b614115614b84565b61411d614bb3565b614126876144fd565b61412e6145a5565b8610158015614144575085614141612b4d565b10155b6141605760405162461bcd60e51b81526004016110e290617e97565b600061416a614be2565b6040516325d998bb60e01b815233600482015290915087906001600160a01b038316906325d998bb90602401602060405180830381865afa1580156141b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141d79190617d0f565b101561421c5760405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e74207374616b696e6760601b60448201526064016110e2565b60405163282d3fdf60e01b8152336004820152602481018890526001600160a01b0382169063282d3fdf90604401600060405180830381600087803b15801561426457600080fd5b505af1158015614278573d6000803e3d6000fd5b5050600160698190557f6bee784efeb983674392298ab585b22866bedf00ebb0eea949d1e66f3f50e71d8054336001600160a01b0319918216811790925560008281526068602090815260408083208690557ff585789965ba69220d5ce3dc1b444eb22ff546f2650694fef8fafe9c26560af98054851686179055606b82528083208690557fdcf345d7f6a8deb7427d0fee62009fa15100353a1c666b51bb5387b25addcfa98054909416909417909255606d825291822083905560748390559190526071815288517f169c6be1b0e6ab5de76b532e587a77340130ac65c5591db02be822dcf1dc0ed69350614373925083918a0190617444565b5085516143899060018301906020890190617444565b50845161439f9060028301906020880190617444565b508381600301819055506001606e886040516143bb9190617e59565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f906143ed908990617e59565b908152602001604051809103902060006101000a81548160ff021916908315150217905550600160706000878760405160200161442b929190617e75565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff1916941515949094179093556074543380855260728352838520829055908452607390915281832080546001600160a01b03191682179055436066559051909182917f6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba69190a3505080156144f4576000805461ff00191690556040516001815260008051602061853c8339815191529060200160405180910390a15b50505050505050565b614505614a0f565b6001600160a01b03811661455b5760405162461bcd60e51b815260206004820152601a60248201527f416464726573732073686f756c64206265206e6f6e2d7a65726f00000000000060448201526064016110e2565b606580546001600160a01b0319166001600160a01b0383169081179091556040517f278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd90600090a250565b60006145af614b6c565b6001600160a01b031663076cd77f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117b2573d6000803e3d6000fd5b6145f46152b2565b6145fc6148d6565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015614642573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906146669190617c30565b6146825760405162461bcd60e51b81526004016110e290617c4d565b600061468d33612c82565b9050614697612b4d565b6146a0826148f6565b111580156146bd57506146b16145a5565b6146ba826148f6565b10155b6146d95760405162461bcd60e51b81526004016110e290617cd8565b6146e161171a565b6147175760405162461bcd60e51b8152602060048201526007602482015266115e1c1a5c995960ca1b60448201526064016110e2565b60006147228461530b565b905061472e848461541e565b60008061473a866154f8565b9250925050600061474a61138990565b9050808310158061475b5750808210155b80614770575061476b8284617dbf565b612710145b156147845761478487858486116000615227565b5050505050612a966001607855565b600061479d614b6c565b6001600160a01b0316631b27e01b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117b2573d6000803e3d6000fd5b6147e2614a0f565b610e10811061482b5760405162461bcd60e51b81526020600482015260156024820152746e6577506572696f6420697320746f6f206c6f6e6760581b60448201526064016110e2565b60aa8190556040518181527f17c6f1d1ce638844b664872f5c6eecb7d150ec0c41187d7f85826a656ee7946f9060200161170f565b614868614a0f565b6001600160a01b0381166148cd5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016110e2565b611f3d81614ff8565b60006117d67111dbdd995c9b985b98d950dbdb9d1c9858dd60721b615539565b6000614900614be2565b604051632c9aab9b60e11b81526001600160a01b03848116600483015291909116906359355736906024015b602060405180830381865afa158015614949573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139329190617d0f565b60008051602061851c833981519152546001600160a01b031690565b60006117d66c42616c6c6f7453746f7261676560981b615539565b6149ac614989565b6001600160a01b031663bce0dbc183836040518363ffffffff1660e01b81526004016149d99291906180d2565b600060405180830381600087803b1580156149f357600080fd5b505af1158015614a07573d6000803e3d6000fd5b505050505050565b6033546001600160a01b03163314612d035760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016110e2565b6000806000614a76614989565b6001600160a01b031663688ca5b2856040518263ffffffff1660e01b8152600401614aa391815260200190565b606060405180830381865afa158015614ac0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ae491906180eb565b9250925092509193909250565b6000806000614afe614989565b6001600160a01b03166309970688856040518263ffffffff1660e01b8152600401614b2b91815260200190565b606060405180830381865afa158015614b48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ae49190618119565b60006117d669456e7653746f7261676560b01b615539565b600054610100900460ff16614bab5760405162461bcd60e51b81526004016110e290618147565b612d036155a7565b600054610100900460ff16614bda5760405162461bcd60e51b81526004016110e290618147565b612d036155ce565b60006117d6665374616b696e6760c81b615539565b604051600190606f90614c0b908690617e59565b9081526040519081900360200190205460ff1615614c27575060005b606e85604051614c379190617e59565b9081526040519081900360200190205460ff1615614c53575060005b60008383604051602001614c68929190617e75565b60408051601f1981840301815291815281516020928301206000818152607090935291205490915060ff1615614c9d57600091505b50949350505050565b614cae614be2565b60405163282d3fdf60e01b81526001600160a01b03848116600483015260248201849052919091169063282d3fdf906044016149d9565b614ced6148d6565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015614d33573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614d579190617c30565b611f3d5760405162461bcd60e51b81526004016110e290617c4d565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615614dab57614da6836155fe565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015614e05575060408051601f3d908101601f19168201909252614e0291810190617d0f565b60015b614e215760405162461bcd60e51b81526004016110e290617d28565b60008051602061851c8339815191528114614e4e5760405162461bcd60e51b81526004016110e290617d76565b50614da683838361569a565b614e62614989565b6001600160a01b031663daacbb95868684610120015187878760000151886020015189604001518a606001518b608001518c60a001518d60c001516040518d63ffffffff1660e01b8152600401614ec49c9b9a99989796959493929190618192565b600060405180830381600087803b158015614ede57600080fd5b505af1158015614ef2573d6000803e3d6000fd5b505050505050505050565b614f05614989565b604051633968764960e11b815260048101849052602481018390526001600160a01b0391909116906372d0ec92906044016149d9565b6000614f45614b6c565b6001600160a01b0316639801bff984846040518363ffffffff1660e01b8152600401614f729291906180d2565b602060405180830381865afa158015614f8f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614fb39190617c30565b9392505050565b614fc2614989565b60405163548f2cdd60e11b815260048101849052602481018390526001600160a01b03919091169063a91e59ba906044016149d9565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b615052614989565b604051632264085960e01b81526004810185905260248101849052604481018390526001600160a01b0391909116906322640859906064015b600060405180830381600087803b1580156150a557600080fd5b505af11580156144f4573d6000803e3d6000fd5b6000600190508480519060200120826020015180519060200120141580156151005750606f856040516150ec9190617e59565b9081526040519081900360200190205460ff165b15615109575060005b85805190602001208260000151805190602001201415801561514a5750606e866040516151369190617e59565b9081526040519081900360200190205460ff165b15615153575060005b60008484604051602001615168929190617e75565b604051602081830303815290604052805190602001209050808360400151846060015160405160200161519c929190617e75565b60405160208183030381529060405280519060200120141580156151ce575060008181526070602052604090205460ff165b156151d857600091505b5095945050505050565b6151ea614989565b60405163605b78c360e11b81526004810185905260248101849052604481018390526001600160a01b03919091169063c0b6f1869060640161508b565b6004821561529657506003600019840161525157615244856156bf565b61524c575060045b615296565b600284036152625761524c85615bc9565b60038403615274576152448583616109565b600484036152855761524c856166a3565b600584036152965761529685616753565b6152a08582614fba565b816152ab5760006077555b5050505050565b6002607854036153045760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016110e2565b6002607855565b600080600061531984614a69565b509092509050600181036153b757607754156153475760405162461bcd60e51b81526004016110e29061822f565b600061535285614af1565b9250505061535e61176b565b8110156153815761537c854261537261176b565b61405b9042617dbf565b6153ac565b8061538a614793565b101561539d5761537c8542615372614793565b6153ac854261405b8482617dbf565b506077849055615417565b600281036153e55760775484146153e05760405162461bcd60e51b81526004016110e29061822f565b615417565b60405162461bcd60e51b8152602060048201526007602482015266115e1c1a5c995960ca1b60448201526064016110e2565b5092915050565b6000607654600161542f9190617dbf565b9050600061543c33612c82565b9050600061544960695490565b61545590612710618272565b9050600084615465576002615468565b60015b9050615472614989565b6040516325918ae760e21b815260048101869052602481018890526001600160a01b038581166044830152606482018490526084820185905291909116906396462b9c9060a401600060405180830381600087803b1580156154d357600080fd5b505af11580156154e7573d6000803e3d6000fd5b505050607694909455505050505050565b6000806000615505614989565b6001600160a01b03166356ba988e856040518263ffffffff1660e01b8152600401614b2b91815260200190565b6001607855565b606554604051630d2020dd60e01b8152600481018390526000916001600160a01b031690630d2020dd90602401602060405180830381865afa158015615583573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139329190617fd6565b600054610100900460ff166155325760405162461bcd60e51b81526004016110e290618147565b600054610100900460ff166155f55760405162461bcd60e51b81526004016110e290618147565b612d0333614ff8565b6001600160a01b0381163b61566b5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016110e2565b60008051602061851c83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6156a383616897565b6000825111806156b05750805b15614da657611e5583836168d7565b60006156cc8260016168fc565b6000806000806000806000806156e18a616a18565b98509850985098509850985098509850506156fb886138f5565b15615759578960008051602061855c8339815191526040516157419060208082526010908201526f20b63932b0b23c90309036b2b6b132b960811b604082015260600190565b60405180910390a25060009998505050505050505050565b615762866127e7565b156157aa578960008051602061855c8339815191526040516157419060208082526012908201527120b63932b0b23c9030903932bbb0b93232b960711b604082015260600190565b6157b26145a5565b8110806157c55750806157c3612b4d565b105b156157e6578960008051602061855c83398151915260405161574190617e97565b806157f089616abf565b1015615812578960008051602061855c83398151915260405161574190618294565b866001600160a01b0316886001600160a01b0316141580156158465750856001600160a01b0316886001600160a01b031614155b15615892578960008051602061855c83398151915260405161574190602080825260169082015275496e76616c6964206d656d626572206164647265737360501b604082015260600190565b61589c8882614ca6565b600060695460016158ad9190617dbf565b9050876067600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606860008a6001600160a01b03166001600160a01b031681526020019081526020016000208190555086606a600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606b6000896001600160a01b03166001600160a01b031681526020019081526020016000208190555088606c600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606d60008b6001600160a01b03166001600160a01b0316815260200190815260200160002081905550600060745460016159e69190617dbf565b6000818152607160209081526040909120895192935091615a0c918391908b0190617444565b508651615a2290600183019060208a0190617444565b508551615a389060028301906020890190617444565b508481600301819055506001606e89604051615a549190617e59565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f90615a86908a90617e59565b908152602001604051809103902060006101000a81548160ff0219169083151502179055506001607060008888604051602001615ac4929190617e75565b60405160208183030381529060405280519060200120815260200190815260200160002060006101000a81548160ff0219169083151502179055508a6073600084815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081607260008d6001600160a01b03166001600160a01b0316815260200190815260200160002081905550826069819055508160748190555043606681905550896001600160a01b03168b6001600160a01b03167f6a2af11b2d73f347f9d5840aea46899e17609730b5cd91bd9c312098038acba660405160405180910390a35060019c9b505050505050505050505050565b615bd48160026168fc565b6000615bdf82616a18565b50505050505050509050615bf2816138f5565b615c47578160008051602061855c833981519152604051615c3b906020808252601490820152732737ba1030b63932b0b23c90309036b2b6b132b960611b604082015260600190565b60405180910390a25050565b6001600160a01b038082166000818152606d6020818152604080842054606954808652606c8452828620548287526067855283872054606a86529387205497909652939092529094928316939083169216908414615daf57606980546000908152606c6020818152604080842080546001600160a01b038c8116808852606d80875285892080549e8a529787528589209087529b821688528488209c909c559486905581546001600160a01b03199081169092558954821690851617909855888452606b8083528185205486548652606a80855283872080548a881689528487528589208054948a529287528589209487528716808952858920939093559087905580548b16905581548a16179055978352606880825288842054945484526067808352898520805489861687528385528b872080548989529386528c8820949095529094168086529985205592905580548616905580549094168517909355919291615e2c565b606980546000908152606c6020908152604080832080546001600160a01b03199081169091556001600160a01b038a81168552606d845282852085905585548552606a84528285208054831690558681168552606b8452828520859055945484526067835281842080549091169055928516825260689052908120555b6001606954615e3b9190617c8a565b6069556000848152607160205260408082209051909190606f90615e63906001850190618374565b908152604051908190036020018120805492151560ff1990931692909217909155600090606e90615e95908490618374565b90815260405160209181900382018120805460ff19169315159390931790925560038301546000926070928492615ed192600288019201618380565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff1916941515949094179093556074546001600160a01b038a16845260729091529120541461602b576001600160a01b038087166000908152607260209081526040808320546074548452607383528184205460719093529220805492985092169550829190615f6c90617e25565b615f779291906174c8565b5060745460009081526071602052604090206001908101805491830191615f9d90617e25565b615fa89291906174c8565b5060745460009081526071602052604090206002908101805491830191615fce90617e25565b615fd99291906174c8565b506074546000908152607160209081526040808320600390810154908501558783526073825280832080546001600160a01b0319166001600160a01b0389169081179091558352607290915290208590555b60748054600090815260736020908152604080832080546001600160a01b03191690556001600160a01b038a1683526072825280832083905592548252607190529081209061607a8282617543565b616088600183016000617543565b616096600283016000617543565b6003820160009055505060016074546160af9190617c8a565b607455436066556160c08787616af9565b826001600160a01b0316866001600160a01b03167faa91016c21c52c58ac64f23f71bbe75becc9ada603e18ee671d09ff15492d1c160405160405180910390a350505050505050565b60008161611b5761611b8360036168fc565b60008060008060008060008060006161328c616a18565b98509850985098509850985098509850985061614d896138f5565b6161b9578b60008051602061855c83398151915260405161619f906020808252601b908201527f4f6c642061646472657373206973206e6f742061206d656d6265720000000000604082015260600190565b60405180910390a260009950505050505050505050613932565b6161cc8c8c8b8b8b8b8b8b8b8b8b616cc7565b6161e25760009950505050505050505050613932565b6001600160a01b03808a166000818152606d6020526040902054918a1614616251576000818152606c6020908152604080832080546001600160a01b0319166001600160a01b038e81169182179092558452606d909252808320849055908c1682528120556162518983614ca6565b6001600160a01b038a166000908152607260209081526040808320548084526071909252808320905191929091606e9061628c908490618374565b908152604051908190036020018120805492151560ff1990931692909217909155600090606f906162c1906001850190618374565b90815260405160209181900382018120805460ff191693151593909317909255600383015460009260709284926162fd92600288019201618380565b60408051808303601f1901815291815281516020928301208352828201939093529101600020805460ff1916921515929092179091558851616344918391908b0190617444565b50865161635a90600183019060208a0190617444565b5085516163709060028301906020890190617444565b506003810185905543606655604051600190606e90616390908b90617e59565b908152604051908190036020018120805492151560ff1990931692909217909155600190606f906163c2908a90617e59565b908152602001604051809103902060006101000a81548160ff0219169083151502179055506001607060008888604051602001616400929190617e75565b60408051808303601f1901815291815281516020928301208352828201939093529082016000908120805460ff191694151594909417909355858352606a90529020546001600160a01b0390811691508916811461649c576000838152606a6020908152604080832080546001600160a01b0319166001600160a01b038e81169182179092558452606b90925280832086905590831682528120555b506000828152606760205260409020546001600160a01b03908116908a16811461654e57896067600085815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082606860008c6001600160a01b03166001600160a01b0316815260200190815260200160002081905550600060686000836001600160a01b03166001600160a01b03168152602001908152602001600020819055505b50896001600160a01b03168b6001600160a01b03161461664e57896073600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080607260008c6001600160a01b03166001600160a01b03168152602001908152602001600020819055506000607260008d6001600160a01b03166001600160a01b03168152602001908152602001600020819055506165ff8e8c616af9565b886001600160a01b03168a6001600160a01b03168c6001600160a01b03167f15f4d750630db473a85edd9d47c500527a2648cc5e676f39645e52790cf07be060405160405180910390a461668f565b896001600160a01b03168b6001600160a01b03167f1feee1b4fcb797c62645da41c5c6edd5f91d4291de0054da625c42b823594c1f60405160405180910390a35b5060019d9c50505050505050505050505050565b6166ae8160046168fc565b60006166b8614989565b6001600160a01b0316637efa9ae3836040518263ffffffff1660e01b81526004016166e591815260200190565b602060405180830381865afa158015616702573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906167269190617fd6565b90506001600160a01b03811615612a965761674081614ce5565b61674b816000611f0b565b436066555050565b61675e8160056168fc565b600080600061676b614989565b6001600160a01b0316631d940da2856040518263ffffffff1660e01b815260040161679891815260200190565b600060405180830381865afa1580156167b5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526167dd9190810190618399565b92509250925060006167ed614b6c565b6040516388c2801960e01b81529091506001600160a01b038216906388c280199061681e90879086906004016180d2565b600060405180830381600087803b15801561683857600080fd5b505af115801561684c573d6000803e3d6000fd5b50504360665550506040517f701c16c2519cdb79aaac423a84733590e3510d9552055b6ad6908f0ab12b6c2990616888908690869086906183de565b60405180910390a15050505050565b6168a0816155fe565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060614fb3838360405180606001604052806027815260200161857c60279139617261565b60008061690884614a69565b50915091508282146169525760405162461bcd60e51b8152602060048201526013602482015272496e76616c696420766f74696e67207479706560681b60448201526064016110e2565b600281146169995760405162461bcd60e51b8152602060048201526014602482015273496e76616c696420766f74696e6720737461746560601b60448201526064016110e2565b6000806169a5866154f8565b92509250506169b361138990565b821015806169c357506113898110155b806169d857506169d38183617dbf565b612710145b614a075760405162461bcd60e51b8152602060048201526011602482015270139bdd081e595d08199a5b985b1a5e9959607a1b60448201526064016110e2565b6000806000806060806060600080616a2e614989565b6001600160a01b03166373df4e018b6040518263ffffffff1660e01b8152600401616a5b91815260200190565b600060405180830381865afa158015616a78573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616aa09190810190618406565b9850985098509850985098509850985098509193959799909294969850565b6000616ac9614be2565b6040516325d998bb60e01b81526001600160a01b03848116600483015291909116906325d998bb9060240161492c565b600080616b05846172d9565b91509150616b116145a5565b616b1b8284617dbf565b1115616bab5760405162461bcd60e51b815260206004820152605360248201527f6d696e5374616b696e672076616c7565206d757374206265206772656174657260448201527f207468616e206f7220657175616c20746f207468652073756d206f6620756e6c6064820152726f636b416d6f756e742c20736c617368696e6760681b608482015260a4016110e2565b6000616bb5614be2565b604051632c9aab9b60e11b81526001600160a01b038681166004830152919250600091831690635935573690602401602060405180830381865afa158015616c01573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616c259190617d0f565b90506000616c316145a5565b616c3b9083617c8a565b905084821115616cbd57616c4f868661735a565b604051634954a85b60e11b81526001600160a01b03878116600483015260248201869052604482018390528416906392a950b690606401600060405180830381600087803b158015616ca057600080fd5b505af1158015616cb4573d6000803e3d6000fd5b505050506144f4565b6144f4868361735a565b60008a616cd957616cd98c60036168fc565b616ce28a6138f5565b616d44578b60008051602061855c833981519152604051616d34906020808252601b908201527f4f6c642061646472657373206973206e6f742061206d656d6265720000000000604082015260600190565b60405180910390a2506000617252565b6001600160a01b03808b166000818152606d6020526040902054918b1614616ebb57616d6f8a6138f5565b15616dd4578c60008051602061855c833981519152604051616dc2906020808252601f908201527f6e6577206164647265737320697320616c72656164792061206d656d62657200604082015260600190565b60405180910390a26000915050617252565b886001600160a01b03168a6001600160a01b031614158015616e085750876001600160a01b03168a6001600160a01b031614155b15616e53578c60008051602061855c833981519152604051616dc290602080825260159082015274496e76616c696420766f746572206164647265737360581b604082015260600190565b616e5b6145a5565b831080616e6e575082616e6c612b4d565b105b15616e8f578c60008051602061855c833981519152604051616dc290617e97565b82616e998b616abf565b1015616ebb578c60008051602061855c833981519152604051616dc290618294565b6001600160a01b038b166000908152607260209081526040808320548084526071909252808320815160808101909252805492939282908290616efd90617e25565b80601f0160208091040260200160405190810160405280929190818152602001828054616f2990617e25565b8015616f765780601f10616f4b57610100808354040283529160200191616f76565b820191906000526020600020905b815481529060010190602001808311616f5957829003601f168201915b50505050508152602001600182018054616f8f90617e25565b80601f0160208091040260200160405190810160405280929190818152602001828054616fbb90617e25565b80156170085780601f10616fdd57610100808354040283529160200191617008565b820191906000526020600020905b815481529060010190602001808311616feb57829003601f168201915b5050505050815260200160028201805461702190617e25565b80601f016020809104026020016040519081016040528092919081815260200182805461704d90617e25565b801561709a5780601f1061706f5761010080835404028352916020019161709a565b820191906000526020600020905b81548152906001019060200180831161707d57829003601f168201915b5050505050815260200160038201548152505090506170bc89898989856150b9565b6170f0578e60008051602061855c8339815191526040516170dc90617ec4565b60405180910390a260009350505050617252565b506000828152606a60205260409020546001600160a01b03908116908d8116908b16148015906171325750896001600160a01b0316816001600160a01b031614155b801561715157506171428a6138f5565b8061715157506171518a6127e7565b1561719d578e60008051602061855c8339815191526040516170dc90602080825260169082015275496e76616c696420726577617264206164647265737360501b604082015260600190565b506000828152606760205260409020546001600160a01b03908116908d8116908c16148015906171df57508a6001600160a01b0316816001600160a01b031614155b80156171fe57506171ef8b6138f5565b806171fe57506171fe8b6127e7565b1561724a578e60008051602061855c8339815191526040516170dc90602080825260169082015275496e76616c696420766f74657273206164647265737360501b604082015260600190565b506001925050505b9b9a5050505050505050505050565b6060600080856001600160a01b03168560405161727e9190617e59565b600060405180830381855af49150503d80600081146172b9576040519150601f19603f3d011682016040523d82523d6000602084013e6172be565b606091505b50915091506172cf86838387617399565b9695505050505050565b6000806172e4614989565b6001600160a01b0316638c7be692846040518263ffffffff1660e01b815260040161731191815260200190565b6040805180830381865afa15801561732d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061735191906184e4565b91509150915091565b617362614be2565b604051637eee288d60e01b81526001600160a01b038481166004830152602482018490529190911690637eee288d906044016149d9565b60608315617408578251600003617401576001600160a01b0385163b6174015760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016110e2565b5081617412565b617412838361741a565b949350505050565b81511561742a5781518083602001fd5b8060405162461bcd60e51b81526004016110e29190618508565b82805461745090617e25565b90600052602060002090601f01602090048101928261747257600085556174b8565b82601f1061748b57805160ff19168380011785556174b8565b828001600101855582156174b8579182015b828111156174b857825182559160200191906001019061749d565b506174c4929150617579565b5090565b8280546174d490617e25565b90600052602060002090601f0160209004810192826174f657600085556174b8565b82601f1061750757805485556174b8565b828001600101855582156174b857600052602060002091601f016020900482015b828111156174b8578254825591600101919060010190617528565b50805461754f90617e25565b6000825580601f1061755f575050565b601f016020900490600052602060002090810190611f3d91905b5b808211156174c4576000815560010161757a565b6001600160a01b0381168114611f3d57600080fd5b8035612cec8161758e565b634e487b7160e01b600052604160045260246000fd5b60405161014081016001600160401b03811182821017156175e7576175e76175ae565b60405290565b604051601f8201601f191681016001600160401b0381118282101715617615576176156175ae565b604052919050565b60006001600160401b03821115617636576176366175ae565b50601f01601f191660200190565b600082601f83011261765557600080fd5b81356176686176638261761d565b6175ed565b81815284602083860101111561767d57600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000606084860312156176af57600080fd5b83356176ba8161758e565b925060208401356001600160401b038111156176d557600080fd5b6176e186828701617644565b925050604084013590509250925092565b60006020828403121561770457600080fd5b8135614fb38161758e565b60006020828403121561772157600080fd5b5035919050565b60008060006060848603121561773d57600080fd5b83356177488161758e565b92506020840135915060408401356001600160401b0381111561776a57600080fd5b61777686828701617644565b9150509250925092565b6000610140828403121561779357600080fd5b61779b6175c4565b90506177a6826175a3565b81526177b4602083016175a3565b60208201526177c5604083016175a3565b604082015260608201356001600160401b03808211156177e457600080fd5b6177f085838601617644565b6060840152608084013591508082111561780957600080fd5b61781585838601617644565b608084015260a084013591508082111561782e57600080fd5b61783a85838601617644565b60a084015260c084013560c084015260e084013560e08401526101009150818401358181111561786957600080fd5b61787586828701617644565b8385015250505061012080830135818301525092915050565b6000602082840312156178a057600080fd5b81356001600160401b038111156178b657600080fd5b61741284828501617780565b6000806000606084860312156178d757600080fd5b83356178e28161758e565b92506020840135915060408401356178f98161758e565b809150509250925092565b600080600080600060a0868803121561791c57600080fd5b853594506020860135935060408601356001600160401b038082111561794157600080fd5b61794d89838a01617644565b9450606088013591508082111561796357600080fd5b5061797088828901617644565b95989497509295608001359392505050565b60005b8381101561799d578181015183820152602001617985565b83811115611e555750506000910152565b600081518084526179c6816020860160208601617982565b601f01601f19169290920160200192915050565b6080815260006179ed60808301876179ae565b82810360208401526179ff81876179ae565b90508281036040840152617a1381866179ae565b91505082606083015295945050505050565b60008060408385031215617a3857600080fd5b8235617a438161758e565b915060208301356001600160401b03811115617a5e57600080fd5b617a6a85828601617644565b9150509250929050565b60008060008060008060c08789031215617a8d57600080fd5b8635617a988161758e565b95506020870135945060408701356001600160401b03811115617aba57600080fd5b617ac689828a01617644565b945050606087013592506080870135915060a087013590509295509295509295565b60008060008060808587031215617afe57600080fd5b84356001600160401b03811115617b1457600080fd5b617b2087828801617780565b9450506020850135617b318161758e565b93969395505050506040820135916060013590565b60008060008060008060c08789031215617b5f57600080fd5b8635617b6a8161758e565b95506020870135945060408701356001600160401b0380821115617b8d57600080fd5b617b998a838b01617644565b95506060890135915080821115617baf57600080fd5b617bbb8a838b01617644565b94506080890135915080821115617bd157600080fd5b50617bde89828a01617644565b92505060a087013590509295509295509295565b8015158114611f3d57600080fd5b60008060408385031215617c1357600080fd5b823591506020830135617c2581617bf2565b809150509250929050565b600060208284031215617c4257600080fd5b8151614fb381617bf2565b6020808252600d908201526c2737902832b936b4b9b9b4b7b760991b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b600082821015617c9c57617c9c617c74565b500390565b6020808252601d908201527f43616e6e6f74206164642070726f706f73616c20746f6f206561726c79000000604082015260600190565b60208082526017908201527f496e76616c6964207374616b696e672062616c616e6365000000000000000000604082015260600190565b600060208284031215617d2157600080fd5b5051919050565b6020808252602e908201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960408201526d6f6e206973206e6f74205555505360901b606082015260800190565b60208082526029908201527f45524331393637557067726164653a20756e737570706f727465642070726f786040820152681a58589b195555525160ba1b606082015260800190565b60008219821115617dd257617dd2617c74565b500190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b600181811c90821680617e3957607f821691505b6020821081036138ef57634e487b7160e01b600052602260045260246000fd5b60008251617e6b818460208701617982565b9190910192915050565b60008351617e87818460208801617982565b9190910191825250602001919050565b602080825260139082015272125b9d985b1a59081b1bd8dac8185b5bdd5b9d606a1b604082015260600190565b6020808252601490820152734475706c696361746564206e6f646520696e666f60601b604082015260600190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b87815286602082015285604082015260018060a01b03851660608201528360808201528260a082015260e060c08201526000617fc960e08301846179ae565b9998505050505050505050565b600060208284031215617fe857600080fd5b8151614fb38161758e565b8051612cec8161758e565b600082601f83011261800f57600080fd5b815161801d6176638261761d565b81815284602083860101111561803257600080fd5b617412826020830160208701617982565b6000806000806080858703121561805957600080fd5b84516001600160401b038082111561807057600080fd5b61807c88838901617ffe565b9550602087015191508082111561809257600080fd5b61809e88838901617ffe565b945060408701519150808211156180b457600080fd5b506180c187828801617ffe565b606096909601519497939650505050565b82815260406020820152600061741260408301846179ae565b60008060006060848603121561810057600080fd5b835192506020840151915060408401516178f981617bf2565b60008060006060848603121561812e57600080fd5b8351925060208401519150604084015190509250925092565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b8c8152602081018c9052604081018b90526001600160a01b038a81166060830152898116608083015288811660a083015287811660c0830152861660e08201526000610180806101008401526181ea818401886179ae565b90508281036101208401526181ff81876179ae565b905082810361014084015261821481866179ae565b915050826101608301529d9c50505050505050505050505050565b60208082526023908201527f4e6f7720696e20766f74696e67207769746820646966666572656e742062616c6040820152621b1bdd60ea1b606082015260800190565b60008261828f57634e487b7160e01b600052601260045260246000fd5b500490565b60208082526027908201527f496e73756666696369656e742062616c616e636520746861742063616e206265604082015266081b1bd8dad95960ca1b606082015260800190565b8054600090600181811c90808316806182f557607f831692505b6020808410820361831657634e487b7160e01b600052602260045260246000fd5b81801561832a576001811461833b57618368565b60ff19861689528489019650618368565b60008881526020902060005b868110156183605781548b820152908501908301618347565b505084890196505b50505050505092915050565b6000614fb382846182db565b600061838c82856182db565b9283525050602001919050565b6000806000606084860312156183ae57600080fd5b835192506020840151915060408401516001600160401b038111156183d257600080fd5b61777686828701617ffe565b8381528260208201526060604082015260006183fd60608301846179ae565b95945050505050565b60008060008060008060008060006101208a8c03121561842557600080fd5b61842e8a617ff3565b985061843c60208b01617ff3565b975061844a60408b01617ff3565b965061845860608b01617ff3565b955060808a01516001600160401b038082111561847457600080fd5b6184808d838e01617ffe565b965060a08c015191508082111561849657600080fd5b6184a28d838e01617ffe565b955060c08c01519150808211156184b857600080fd5b506184c58c828d01617ffe565b93505060e08a015191506101008a015190509295985092959850929598565b600080604083850312156184f757600080fd5b505080516020909101519092909150565b602081526000614fb360208301846179ae56fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249885e7f4987c0698db47045ad8cea110b51138f0eecbd94915842328cf6c3dc97d416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220215d25ede0ee3be33a028eb08f73a6744f79eb175781e9ee1cc1fc1a6c1065d764736f6c634300080e0033",
}

// GovImpABI is the input ABI used to generate the binding from.
// Deprecated: Use GovImpMetaData.ABI instead.
var GovImpABI = GovImpMetaData.ABI

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

// BLOCKREWARDDISTRIBUTIONMAINTANANCENAME is a free data retrieval call binding the contract method 0x0b1d39b8.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_MAINTANANCE_NAME() view returns(bytes32)
func (_GovImp *GovImpCaller) BLOCKREWARDDISTRIBUTIONMAINTANANCENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovImp.contract.Call(opts, &out, "BLOCK_REWARD_DISTRIBUTION_MAINTANANCE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREWARDDISTRIBUTIONMAINTANANCENAME is a free data retrieval call binding the contract method 0x0b1d39b8.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_MAINTANANCE_NAME() view returns(bytes32)
func (_GovImp *GovImpSession) BLOCKREWARDDISTRIBUTIONMAINTANANCENAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDDISTRIBUTIONMAINTANANCENAME(&_GovImp.CallOpts)
}

// BLOCKREWARDDISTRIBUTIONMAINTANANCENAME is a free data retrieval call binding the contract method 0x0b1d39b8.
//
// Solidity: function BLOCK_REWARD_DISTRIBUTION_MAINTANANCE_NAME() view returns(bytes32)
func (_GovImp *GovImpCallerSession) BLOCKREWARDDISTRIBUTIONMAINTANANCENAME() ([32]byte, error) {
	return _GovImp.Contract.BLOCKREWARDDISTRIBUTIONMAINTANANCENAME(&_GovImp.CallOpts)
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
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GovImp *GovImpTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GovImp *GovImpSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GovImp.Contract.UpgradeTo(&_GovImp.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GovImp *GovImpTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GovImp.Contract.UpgradeTo(&_GovImp.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GovImp *GovImpTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GovImp.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GovImp *GovImpSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GovImp.Contract.UpgradeToAndCall(&_GovImp.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GovImp *GovImpTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GovImp.Contract.UpgradeToAndCall(&_GovImp.TransactOpts, newImplementation, data)
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
