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

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_imp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"5c60da1b": "implementation()",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161077238038061077283398101604081905261002f91610326565b604080516020810190915260008152819061006b60017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd61034f565b60008051602061072b8339815191521461008757610087610374565b6100938282600061009b565b505050610405565b6100a4836100d1565b6000825111806100b15750805b156100cc576100ca838361011160201b61008b1760201c565b505b505050565b6100da8161013d565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060610136838360405180606001604052806027815260200161074b602791396101fd565b9392505050565b610150816102db60201b6100b71760201c565b6101b75760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b806101dc60008051602061072b83398151915260001b6102ea60201b6100c61760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606001600160a01b0384163b6102655760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016101ae565b600080856001600160a01b03168560405161028091906103b6565b600060405180830381855af49150503d80600081146102bb576040519150601f19603f3d011682016040523d82523d6000602084013e6102c0565b606091505b5090925090506102d18282866102ed565b9695505050505050565b6001600160a01b03163b151590565b90565b606083156102fc575081610136565b82511561030c5782518084602001fd5b8160405162461bcd60e51b81526004016101ae91906103d2565b60006020828403121561033857600080fd5b81516001600160a01b038116811461013657600080fd5b60008282101561036f57634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052600160045260246000fd5b60005b838110156103a557818101518382015260200161038d565b838111156100ca5750506000910152565b600082516103c881846020870161038a565b9190910192915050565b60208152600082518060208401526103f181604085016020870161038a565b601f01601f19169190910160400192915050565b610317806104146000396000f3fe6080604052600436106100225760003560e01c80635c60da1b1461003957610031565b366100315761002f61006a565b005b61002f61006a565b34801561004557600080fd5b5061004e61007c565b6040516001600160a01b03909116815260200160405180910390f35b61007a6100756100c9565b6100fc565b565b60006100866100c9565b905090565b60606100b083836040518060600160405280602781526020016102bb60279139610120565b9392505050565b6001600160a01b03163b151590565b90565b60006100867f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b3660008037600080366000845af43d6000803e80801561011b573d6000f35b3d6000fd5b60606001600160a01b0384163b61018d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084015b60405180910390fd5b600080856001600160a01b0316856040516101a8919061026b565b600060405180830381855af49150503d80600081146101e3576040519150601f19603f3d011682016040523d82523d6000602084013e6101e8565b606091505b50915091506101f8828286610202565b9695505050505050565b606083156102115750816100b0565b8251156102215782518084602001fd5b8160405162461bcd60e51b81526004016101849190610287565b60005b8381101561025657818101518382015260200161023e565b83811115610265576000848401525b50505050565b6000825161027d81846020870161023b565b9190910192915050565b60208152600082518060208401526102a681604085016020870161023b565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212204c6da138e4aadf72357da99a86f99d8734efaebb1c66d7fcc80efea9db3acd3264736f6c634300080e0033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Deprecated: Use StakingMetaData.Sigs instead.
// StakingFuncSigs maps the 4-byte function signature to its string representation.
var StakingFuncSigs = StakingMetaData.Sigs

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend, _imp common.Address) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend, _imp)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Staking *StakingCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Staking *StakingSession) Implementation() (common.Address, error) {
	return _Staking.Contract.Implementation(&_Staking.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Staking *StakingCallerSession) Implementation() (common.Address, error) {
	return _Staking.Contract.Implementation(&_Staking.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Staking *StakingTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Staking.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Staking *StakingSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Staking.Contract.Fallback(&_Staking.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Staking *StakingTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Staking.Contract.Fallback(&_Staking.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactorSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// StakingAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Staking contract.
type StakingAdminChangedIterator struct {
	Event *StakingAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingAdminChanged)
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
		it.Event = new(StakingAdminChanged)
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
func (it *StakingAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingAdminChanged represents a AdminChanged event raised by the Staking contract.
type StakingAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Staking *StakingFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*StakingAdminChangedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &StakingAdminChangedIterator{contract: _Staking.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Staking *StakingFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingAdminChanged)
				if err := _Staking.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Staking *StakingFilterer) ParseAdminChanged(log types.Log) (*StakingAdminChanged, error) {
	event := new(StakingAdminChanged)
	if err := _Staking.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Staking contract.
type StakingBeaconUpgradedIterator struct {
	Event *StakingBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *StakingBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingBeaconUpgraded)
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
		it.Event = new(StakingBeaconUpgraded)
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
func (it *StakingBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingBeaconUpgraded represents a BeaconUpgraded event raised by the Staking contract.
type StakingBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Staking *StakingFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*StakingBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &StakingBeaconUpgradedIterator{contract: _Staking.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Staking *StakingFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *StakingBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingBeaconUpgraded)
				if err := _Staking.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Staking *StakingFilterer) ParseBeaconUpgraded(log types.Log) (*StakingBeaconUpgraded, error) {
	event := new(StakingBeaconUpgraded)
	if err := _Staking.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Staking contract.
type StakingUpgradedIterator struct {
	Event *StakingUpgraded // Event containing the contract specifics and raw log

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
func (it *StakingUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUpgraded)
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
		it.Event = new(StakingUpgraded)
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
func (it *StakingUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUpgraded represents a Upgraded event raised by the Staking contract.
type StakingUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Staking *StakingFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StakingUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StakingUpgradedIterator{contract: _Staking.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Staking *StakingFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StakingUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUpgraded)
				if err := _Staking.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Staking *StakingFilterer) ParseUpgraded(log types.Log) (*StakingUpgraded, error) {
	event := new(StakingUpgraded)
	if err := _Staking.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpMetaData contains all meta data concerning the StakingImp contract.
var StakingImpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ncp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ncpTotalLocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotalLocked\",\"type\":\"uint256\"}],\"name\":\"DelegateStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ncp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ncpTotalLocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotalLocked\",\"type\":\"uint256\"}],\"name\":\"DelegateUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ncp\",\"type\":\"address\"}],\"name\":\"NCPAddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Revoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"TransferLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ECOSYSTEM_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ENV_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAINTENANCE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_POOL_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_REWARD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"availableBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"calcVotingWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"factor\",\"type\":\"uint32\"}],\"name\":\"calcVotingWeightWithScaleFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ncp\",\"type\":\"address\"}],\"name\":\"delegateDepositAndLockMore\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ncp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegateUnlockAndWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ncp\",\"type\":\"address\"}],\"name\":\"getRatioOfUserBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalLockedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRevoked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"}],\"name\":\"lockMore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"lockedBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ncpStaking\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reg\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ncpStaking\",\"type\":\"address\"}],\"name\":\"setNCPStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ext\",\"type\":\"uint256\"}],\"name\":\"transferLocked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImp\",\"type\":\"address\"}],\"name\":\"upgradeStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ncp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ncp\",\"type\":\"address\"}],\"name\":\"userTotalBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"9986e4b9": "BALLOT_STORAGE_NAME()",
		"34125c84": "ECOSYSTEM_NAME()",
		"7bf46530": "ENV_STORAGE_NAME()",
		"6c78d2cf": "GOV_NAME()",
		"4bd1ed76": "MAINTENANCE_NAME()",
		"2f40992e": "REWARD_POOL_NAME()",
		"1e0cba0d": "STAKING_NAME()",
		"5a731cca": "STAKING_REWARD_NAME()",
		"25d998bb": "availableBalanceOf(address)",
		"70a08231": "balanceOf(address)",
		"884d97a7": "calcVotingWeight(address)",
		"7d77a0eb": "calcVotingWeightWithScaleFactor(address,uint32)",
		"bac4f338": "delegateDepositAndLockMore(address)",
		"f6931822": "delegateUnlockAndWithdraw(address,uint256)",
		"d0e30db0": "deposit()",
		"2b0b9c5e": "getRatioOfUserBalance(address)",
		"d5c25890": "getTotalLockedBalance()",
		"c0d91eaf": "init(address,bytes)",
		"2bc9ed02": "isRevoked()",
		"282d3fdf": "lock(address,uint256)",
		"9667e76a": "lockMore(uint256)",
		"59355736": "lockedBalanceOf(address)",
		"f1b8aa1d": "ncpStaking()",
		"8da5cb5b": "owner()",
		"52d1902d": "proxiableUUID()",
		"738fdd1a": "reg()",
		"715018a6": "renounceOwnership()",
		"b6549f75": "revoke()",
		"7f2f4c06": "setNCPStaking(address)",
		"a91ee0dc": "setRegistry(address)",
		"92a950b6": "transferLocked(address,uint256,uint256)",
		"f2fde38b": "transferOwnership(address)",
		"7eee288d": "unlock(address,uint256)",
		"f3f63080": "upgradeStaking(address)",
		"3659cfe6": "upgradeTo(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
		"12853615": "userBalanceOf(address,address)",
		"193468ac": "userTotalBalanceOf(address)",
		"2e1a7d4d": "withdraw(uint256)",
	},
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000e8565b600054610100900460ff1615620000935760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015620000e6576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6080516135176200012060003960008181610be601528181610c2601528181610cc501528181610d050152610d9401526135176000f3fe6080604052600436106102345760003560e01c8063738fdd1a1161012e578063a91ee0dc116100ab578063d5c258901161006f578063d5c25890146106b3578063f1b8aa1d146106c8578063f2fde38b146106e8578063f3f6308014610708578063f69318221461072857600080fd5b8063a91ee0dc14610643578063b6549f7514610663578063bac4f33814610678578063c0d91eaf1461068b578063d0e30db0146106ab57600080fd5b8063884d97a7116100f2578063884d97a7146105a15780638da5cb5b146105c157806392a950b6146105df5780639667e76a146105ff5780639986e4b91461061f57600080fd5b8063738fdd1a146104e85780637bf46530146105205780637d77a0eb146105415780637eee288d146105615780637f2f4c061461058157600080fd5b806334125c84116101bc5780635935573611610180578063593557361461041a5780635a731cca146104505780636c78d2cf1461047457806370a082311461049d578063715018a6146104d357600080fd5b806334125c84146103905780633659cfe6146103b05780634bd1ed76146103d05780634f1ef286146103f257806352d1902d1461040557600080fd5b8063282d3fdf11610203578063282d3fdf146102ea5780632b0b9c5e1461030c5780632bc9ed021461032c5780632e1a7d4d1461034f5780632f40992e1461036f57600080fd5b80631285361514610243578063193468ac146102765780631e0cba0d146102ac57806325d998bb146102ca57600080fd5b3661023e57600080fd5b600080fd5b34801561024f57600080fd5b5061026361025e366004612e48565b610748565b6040519081526020015b60405180910390f35b34801561028257600080fd5b50610263610291366004612e81565b6001600160a01b0316600090815260ce602052604090205490565b3480156102b857600080fd5b50610263665374616b696e6760c81b81565b3480156102d657600080fd5b506102636102e5366004612e81565b610775565b3480156102f657600080fd5b5061030a610305366004612e9e565b6107a3565b005b34801561031857600080fd5b50610263610327366004612e81565b6107e9565b34801561033857600080fd5b50609b5460ff16604051901515815260200161026d565b34801561035b57600080fd5b5061030a61036a366004612eca565b610845565b34801561037b57600080fd5b506102636914995dd85c99141bdbdb60b21b81565b34801561039c57600080fd5b506102636845636f73797374656d60b81b81565b3480156103bc57600080fd5b5061030a6103cb366004612e81565b610bdc565b3480156103dc57600080fd5b506102636a4d61696e74656e616e636560a81b81565b61030a610400366004612f2a565b610cbb565b34801561041157600080fd5b50610263610d87565b34801561042657600080fd5b50610263610435366004612e81565b6001600160a01b031660009081526099602052604090205490565b34801561045c57600080fd5b506102636c14dd185ada5b99d4995dd85c99609a1b81565b34801561048057600080fd5b506102637111dbdd995c9b985b98d950dbdb9d1c9858dd60721b81565b3480156104a957600080fd5b506102636104b8366004612e81565b6001600160a01b031660009081526098602052604090205490565b3480156104df57600080fd5b5061030a610e3a565b3480156104f457600080fd5b50606554610508906001600160a01b031681565b6040516001600160a01b03909116815260200161026d565b34801561052c57600080fd5b5061026369456e7653746f7261676560b01b81565b34801561054d57600080fd5b5061026361055c366004612fd2565b610e4e565b34801561056d57600080fd5b5061030a61057c366004612e9e565b610ec2565b34801561058d57600080fd5b5061030a61059c366004612e81565b610efb565b3480156105ad57600080fd5b506102636105bc366004612e81565b610fa3565b3480156105cd57600080fd5b506033546001600160a01b0316610508565b3480156105eb57600080fd5b5061030a6105fa366004613009565b610fb0565b34801561060b57600080fd5b5061030a61061a366004612eca565b6112e6565b34801561062b57600080fd5b506102636c42616c6c6f7453746f7261676560981b81565b34801561064f57600080fd5b5061030a61065e366004612e81565b6114e1565b34801561066f57600080fd5b5061030a611589565b61030a610686366004612e81565b611693565b34801561069757600080fd5b5061030a6106a6366004612f2a565b611a85565b61030a611c36565b3480156106bf57600080fd5b50609a54610263565b3480156106d457600080fd5b5060cf54610508906001600160a01b031681565b3480156106f457600080fd5b5061030a610703366004612e81565b612000565b34801561071457600080fd5b5061030a610723366004612e81565b612076565b34801561073457600080fd5b5061030a610743366004612e9e565b612096565b6001600160a01b03808316600090815260cd60209081526040808320938516835292905220545b92915050565b6001600160a01b038116600090815260996020908152604080832054609890925282205461076f9190613054565b336107ac61249b565b6001600160a01b0316146107db5760405162461bcd60e51b81526004016107d29061306b565b60405180910390fd5b6107e582826124c0565b5050565b6001600160a01b038116600090815260ce60209081526040808320546099909252822054811580610818575080155b15610827575060009392505050565b80610833836064613092565b61083d91906130b1565b949350505050565b6002606654036108675760405162461bcd60e51b81526004016107d2906130d3565b6002606655609b5460ff161561088f5760405162461bcd60e51b81526004016107d29061310a565b600081116108af5760405162461bcd60e51b81526004016107d29061312e565b60006108b9612735565b6001600160a01b031663076cd77f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061091a919061316f565b905060006109288383613188565b33600090815260ce602090815260408083205460999092529091205461094e9190613054565b106109615761095d338461274d565b5060015b61096a33610775565b8311156109895760405162461bcd60e51b81526004016107d2906131a0565b336000908152609860205260409020546109a4908490613054565b3360009081526098602052604090205560cf546001600160a01b0316158015906109cb5750805b15610add5760cf546040516000916001600160a01b03169085908381818185875af1925050503d8060008114610a1d576040519150601f19603f3d011682016040523d82523d6000602084013e610a22565b606091505b5050905080610a735760405162461bcd60e51b815260206004820152601e60248201527f5472616e7366657220746f204e4350207374616b696e67206661696c6564000060448201526064016107d2565b60cf546040516306aa67f960e01b8152600481018690523360248201526001600160a01b03909116906306aa67f990604401600060405180830381600087803b158015610abf57600080fd5b505af1158015610ad3573d6000803e3d6000fd5b5050505050610b77565b604051600090339085908381818185875af1925050503d8060008114610b1f576040519150601f19603f3d011682016040523d82523d6000602084013e610b24565b606091505b5050905080610b755760405162461bcd60e51b815260206004820152601960248201527f5472616e7366657220746f2073656e646572206661696c65640000000000000060448201526064016107d2565b505b336000818152609860205260409020547f204fccf0d92ed8d48f204adb39b2e81e92bad0dedb93f5716ca9478cfb57de00908590610bb484610775565b6040805193845260208401929092529082015260600160405180910390a25050600160665550565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003610c245760405162461bcd60e51b81526004016107d2906131f4565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610c6d60008051602061349b833981519152546001600160a01b031690565b6001600160a01b031614610c935760405162461bcd60e51b81526004016107d290613240565b610c9c81612896565b60408051600080825260208201909252610cb89183919061289e565b50565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003610d035760405162461bcd60e51b81526004016107d2906131f4565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610d4c60008051602061349b833981519152546001600160a01b031690565b6001600160a01b031614610d725760405162461bcd60e51b81526004016107d290613240565b610d7b82612896565b6107e58282600161289e565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e275760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016107d2565b5060008051602061349b83398151915290565b610e42612a09565b610e4c6000612a63565b565b6001600160a01b0382166000908152609960205260408120541580610e77575063ffffffff8216155b15610e845750600061076f565b609a546001600160a01b038416600090815260996020526040902054610eb19063ffffffff851690613092565b610ebb91906130b1565b9392505050565b33610ecb61249b565b6001600160a01b031614610ef15760405162461bcd60e51b81526004016107d29061306b565b6107e5828261274d565b610f03612a09565b6001600160a01b038116610f595760405162461bcd60e51b815260206004820152601e60248201527f4e43505374616b696e6720697320746865207a65726f2061646472657373000060448201526064016107d2565b60cf80546001600160a01b0319166001600160a01b0383169081179091556040517ffd5754300dde6066eda4fabd23616c1d560a3360c85c0716c46e00649bdeeddf90600090a250565b600061076f826064610e4e565b33610fb961249b565b6001600160a01b031614610fdf5760405162461bcd60e51b81526004016107d29061306b565b6000610ff4661390d4115e1a5d60ca1b612ab5565b90506110008484610ec2565b6001600160a01b038416600090815260986020526040902054611024908490613054565b6001600160a01b038516600090815260986020526040812091909155611048612b23565b6001600160a01b03811660009081526098602052604090205490915061106f908590613188565b6001600160a01b03808316600090815260986020908152604080832094909455918816815260ce909152908120546110a79085613054565b90506110b38682610ec2565b6001600160a01b03861660009081526099602090815260408083205460ce909252909120548110156111665760405162461bcd60e51b815260206004820152605060248201527f7472616e73666572656442616c616e6365206d7573742062652067726561746560448201527f72207468616e206f7220657175616c20746f205f6c6f636b656455736572426160648201526f3630b731b2aa37a721a82a37ba30b61760811b608482015260a4016107d2565b6111708782610ec2565b6001600160a01b038716600090815260986020526040902054611194908290613054565b6001600160a01b0388811660008181526098602090815260408083209590955560ce9052839020549251638408bdb160e01b81526004810191909152602481018490526044810192909252851690638408bdb19083906064016000604051808303818588803b15801561120657600080fd5b505af115801561121a573d6000803e3d6000fd5b505060cf546001600160a01b0316159250611268915050576001600160a01b03808816600090815260ce6020908152604080832083905560cd825280832060cf549094168352929052908120555b6001600160a01b0387167f2caed32a519a1fd89486d3ffe06202febb5ed51534d571dbab93058545a29e2461129d8389613188565b6001600160a01b038a166000908152609860205260409020546112bf8b610775565b6040805193845260208401929092529082015260600160405180910390a250505050505050565b6112ee61249b565b604051636f1e853360e01b81523360048201526001600160a01b039190911690636f1e853390602401602060405180830381865afa158015611334573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611358919061328c565b6113745760405162461bcd60e51b81526004016107d29061306b565b61137e33826124c0565b60cf546001600160a01b031615610cb85760cf5460405163bbc2611360e01b81523360048201526000916001600160a01b03169063bbc2611390602401602060405180830381865afa1580156113d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113fc919061316f565b60cf54604051631069f3b560e01b8152600481018390523360248201529192506001600160a01b031690631069f3b59060440160a060405180830381865afa15801561144c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061147091906132ae565b60800151156107e55760cf546040516301008e9960e61b8152600481018490523360248201526001600160a01b0390911690634023a64090604401600060405180830381600087803b1580156114c557600080fd5b505af11580156114d9573d6000803e3d6000fd5b505050505050565b6114e9612a09565b6001600160a01b03811661153f5760405162461bcd60e51b815260206004820152601a60248201527f416464726573732073686f756c64206265206e6f6e2d7a65726f00000000000060448201526064016107d2565b606580546001600160a01b0319166001600160a01b0383169081179091556040517f278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd90600090a250565b611591612a09565b609b5460ff16156115b45760405162461bcd60e51b81526004016107d29061310a565b60006115c86033546001600160a01b031690565b905047806116065760405162461bcd60e51b815260206004820152600b60248201526a062616c616e6365203d20360ac1b60448201526064016107d2565b6040516001600160a01b0383169082156108fc029083906000818181858888f1935050505015801561163c573d6000803e3d6000fd5b50609b805460ff191660011790556040516001600160a01b038316907f713b90881ad62c4fa8ab6bd9197fa86481fc0c11b2edba60026514281b2dbac4906116879084815260200190565b60405180910390a25050565b6002606654036116b55760405162461bcd60e51b81526004016107d2906130d3565b6002606655609b5460ff16156116dd5760405162461bcd60e51b81526004016107d29061310a565b60cf546001600160a01b031633146117075760405162461bcd60e51b81526004016107d29061331e565b600034116117275760405162461bcd60e51b81526004016107d29061336d565b61172f61249b565b60405163288c314960e21b81526001600160a01b038381166004830152919091169063a230c52490602401602060405180830381865afa158015611777573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061179b919061328c565b6117e05760405162461bcd60e51b81526020600482015260166024820152752721a81039b437bab63210313290309036b2b6b132b960511b60448201526064016107d2565b6001600160a01b0381166000908152609860205260409020543490611806908290613188565b6001600160a01b03831660009081526098602052604081209190915561182a612735565b6001600160a01b031663076cd77f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611867573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061188b919061316f565b90506000611897612735565b6001600160a01b031663737c59b86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f8919061316f565b6001600160a01b038516600090815260996020526040902054909150821180159061194757506001600160a01b0384166000908152609960205260409020548190611944908590613188565b11155b6119935760405162461bcd60e51b815260206004820152601f60248201527f757365722073686f756c6420626520696e207374616b696e672072616e67650060448201526064016107d2565b61199d84846124c0565b6001600160a01b038416600090815260cd602090815260408083203384529091529020546119cc908490613188565b6001600160a01b038516600081815260cd6020908152604080832033845282528083209490945591815260ce9091522054611a08908490613188565b6001600160a01b038516600081815260ce6020908152604080832085905560cd825280832033808552908352928190205481518981529283019590955281019390935290917f74cfc20f0e6d14384c3a60820d3e814f6979d009cdbb43db27fa56fe475172fd9060600160405180910390a3505060016066555050565b600054610100900460ff1615808015611aa55750600054600160ff909116105b80611abf5750303b158015611abf575060005460ff166001145b611b225760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016107d2565b6000805460ff191660011790558015611b45576000805461ff0019166101001790555b6000609a55609b805460ff19169055611b5c612b3a565b611b64612b69565b611b6d836114e1565b815115611beb57600080600080602086019150855182611b8d9190613188565b90505b80821015611be65781519350611ba7602083613188565b9150808210611bb557600080fd5b81519250611bc4602083613188565b6001600160a01b03851660009081526098602052604090208490559150611b90565b505050505b8015611c31576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b600260665403611c585760405162461bcd60e51b81526004016107d2906130d3565b6002606655609b5460ff1615611c805760405162461bcd60e51b81526004016107d29061310a565b60003411611ca05760405162461bcd60e51b81526004016107d29061336d565b33600090815260986020526040902054611cbb903490613188565b33600090815260986020526040902055611cd361249b565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015611d19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d3d919061328c565b15611f9e576000611d4c612735565b6001600160a01b031663076cd77f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dad919061316f565b3360009081526099602052604090205490915081118015611def575033600090815260996020526040902054611de39082613054565b611dec33610775565b10155b15611f9c5760cf546001600160a01b031615611f785760cf5460405163bbc2611360e01b81523360048201526000916001600160a01b03169063bbc2611390602401602060405180830381865afa158015611e4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e72919061316f565b60cf54604051631069f3b560e01b8152600481018390523360248201529192506001600160a01b031690631069f3b59060440160a060405180830381865afa158015611ec2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ee691906132ae565b6080015115611f765760cf54336000908152609960205260409020546001600160a01b0390911690634023a64090611f1e9085613054565b6040516001600160e01b031960e084901b1681526004810191909152336024820152604401600060405180830381600087803b158015611f5d57600080fd5b505af1158015611f71573d6000803e3d6000fd5b505050505b505b33600081815260996020526040902054611f9c9190611f979084613054565b6124c0565b505b336000818152609860205260409020547fb4caaf29adda3eefee3ad552a8e85058589bf834c7466cae4ee58787f70589ed903490611fdb84610775565b6040805193845260208401929092529082015260600160405180910390a26001606655565b612008612a09565b6001600160a01b03811661206d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016107d2565b610cb881612a63565b61207e612a09565b6001600160a01b03811615610cb857610c9c81612896565b6002606654036120b85760405162461bcd60e51b81526004016107d2906130d3565b6002606655609b5460ff16156120e05760405162461bcd60e51b81526004016107d29061310a565b60cf546001600160a01b0316331461210a5760405162461bcd60e51b81526004016107d29061331e565b6000811161212a5760405162461bcd60e51b81526004016107d29061312e565b61213261249b565b60405163288c314960e21b81526001600160a01b038481166004830152919091169063a230c52490602401602060405180830381865afa15801561217a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061219e919061328c565b6121e35760405162461bcd60e51b81526020600482015260166024820152752721a81039b437bab63210313290309036b2b6b132b960511b60448201526064016107d2565b8060006121ee612735565b6001600160a01b031663076cd77f6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561222b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061224f919061316f565b6001600160a01b038516600090815260cd6020908152604080832033845290915290205490915082118015906122ae575080826122a1866001600160a01b031660009081526099602052604090205490565b6122ab9190613054565b10155b6122ca5760405162461bcd60e51b81526004016107d2906131a0565b6122d4848361274d565b6001600160a01b0384166000908152609860205260409020546122f8908390613054565b6001600160a01b03851660009081526098602090815260408083209390935560cd815282822033835290522054612330908390613054565b6001600160a01b038516600081815260cd6020908152604080832033845282528083209490945591815260ce909152205461236c908390613054565b6001600160a01b03858116600090815260ce60205260408082209390935560cf54925190929091169084908381818185875af1925050503d80600081146123cf576040519150601f19603f3d011682016040523d82523d6000602084013e6123d4565b606091505b50509050806124255760405162461bcd60e51b815260206004820152601e60248201527f5472616e7366657220746f204e4350207374616b696e67206661696c6564000060448201526064016107d2565b6001600160a01b038516600081815260ce602090815260408083205460cd835281842033808652908452938290205482518981529384019190915282820152517f03d2bb70c6ccc49d68a465a06edffb976961cf8930888658ca2339fa62b8bda29181900360600190a350506001606655505050565b60006124bb7111dbdd995c9b985b98d950dbdb9d1c9858dd60721b612ab5565b905090565b806000036124cc575050565b6001600160a01b03821660009081526098602052604090205481111561254d5760405162461bcd60e51b815260206004820152603060248201527f4c6f636b20616d6f756e742073686f756c6420626520657175616c206f72206c60448201526f657373207468616e2062616c616e636560801b60648201526084016107d2565b8061255783610775565b10156125b55760405162461bcd60e51b815260206004820152602760248201527f496e73756666696369656e742062616c616e636520746861742063616e206265604482015266081b1bd8dad95960ca1b60648201526084016107d2565b60006125bf612735565b6001600160a01b031663737c59b86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156125fc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612620919061316f565b6001600160a01b038416600090815260996020526040902054909150612647908390613188565b6001600160a01b03841660009081526099602052604090208190558110156126bb5760405162461bcd60e51b815260206004820152602160248201527f4c6f636b65642062616c616e6365206973206c6172676572207468616e206d616044820152600f60fb1b60648201526084016107d2565b81609a546126c99190613188565b609a556001600160a01b0383166000818152609860205260409020547f44cebfefa4561bee5b61d675ccfd8dc9969fff9cc15e7a4eccccd62af94f9c1190849061271287610775565b6040805193845260208401929092529082015260600160405180910390a2505050565b60006124bb69456e7653746f7261676560b01b612ab5565b80600003612759575050565b6001600160a01b0382166000908152609960205260409020548111156127e75760405162461bcd60e51b815260206004820152603960248201527f556e6c6f636b20616d6f756e742073686f756c6420626520657175616c206f7260448201527f206c657373207468616e2062616c616e6365206c6f636b65640000000000000060648201526084016107d2565b6001600160a01b03821660009081526099602052604090205461280b908290613054565b6001600160a01b038316600090815260996020526040902055609a54612832908290613054565b609a556001600160a01b0382166000818152609860205260409020547f5245d528087a96a64f4589a764f00061e4671eab90cb1e019b1a5b24b2e4c2a890839061287b86610775565b60408051938452602084019290925290820152606001611687565b610cb8612a09565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156128d157611c3183612b98565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561292b575060408051601f3d908101601f191682019092526129289181019061316f565b60015b61298e5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b60648201526084016107d2565b60008051602061349b83398151915281146129fd5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b60648201526084016107d2565b50611c31838383612c34565b6033546001600160a01b03163314610e4c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016107d2565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b606554604051630d2020dd60e01b8152600481018390526000916001600160a01b031690630d2020dd90602401602060405180830381865afa158015612aff573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076f91906133b7565b60006124bb6845636f73797374656d60b81b612ab5565b600054610100900460ff16612b615760405162461bcd60e51b81526004016107d2906133d4565b610e4c612c5f565b600054610100900460ff16612b905760405162461bcd60e51b81526004016107d2906133d4565b610e4c612c8d565b6001600160a01b0381163b612c055760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016107d2565b60008051602061349b83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b612c3d83612cbd565b600082511180612c4a5750805b15611c3157612c598383612cfd565b50505050565b600054610100900460ff16612c865760405162461bcd60e51b81526004016107d2906133d4565b6001606655565b600054610100900460ff16612cb45760405162461bcd60e51b81526004016107d2906133d4565b610e4c33612a63565b612cc681612b98565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060610ebb83836040518060600160405280602781526020016134bb6027913960606001600160a01b0384163b612d855760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016107d2565b600080856001600160a01b031685604051612da0919061344b565b600060405180830381855af49150503d8060008114612ddb576040519150601f19603f3d011682016040523d82523d6000602084013e612de0565b606091505b5091509150612df0828286612dfa565b9695505050505050565b60608315612e09575081610ebb565b825115612e195782518084602001fd5b8160405162461bcd60e51b81526004016107d29190613467565b6001600160a01b0381168114610cb857600080fd5b60008060408385031215612e5b57600080fd5b8235612e6681612e33565b91506020830135612e7681612e33565b809150509250929050565b600060208284031215612e9357600080fd5b8135610ebb81612e33565b60008060408385031215612eb157600080fd5b8235612ebc81612e33565b946020939093013593505050565b600060208284031215612edc57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715612f2257612f22612ee3565b604052919050565b60008060408385031215612f3d57600080fd5b8235612f4881612e33565b915060208381013567ffffffffffffffff80821115612f6657600080fd5b818601915086601f830112612f7a57600080fd5b813581811115612f8c57612f8c612ee3565b612f9e601f8201601f19168501612ef9565b91508082528784828501011115612fb457600080fd5b80848401858401376000848284010152508093505050509250929050565b60008060408385031215612fe557600080fd5b8235612ff081612e33565b9150602083013563ffffffff81168114612e7657600080fd5b60008060006060848603121561301e57600080fd5b833561302981612e33565b95602085013595506040909401359392505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156130665761306661303e565b500390565b6020808252600d908201526c2737902832b936b4b9b9b4b7b760991b604082015260600190565b60008160001904831182151516156130ac576130ac61303e565b500290565b6000826130ce57634e487b7160e01b600052601260045260246000fd5b500490565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b6020808252600a9082015269125cc81c995d9bdad95960b21b604082015260600190565b60208082526021908201527f416d6f756e742073686f756c6420626520626967676572207468616e207a65726040820152606f60f81b606082015260800190565b60006020828403121561318157600080fd5b5051919050565b6000821982111561319b5761319b61303e565b500190565b60208082526034908201527f576974686472617720616d6f756e742073686f756c6420626520657175616c206040820152736f72206c657373207468616e2062616c616e636560601b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b60006020828403121561329e57600080fd5b81518015158114610ebb57600080fd5b600060a082840312156132c057600080fd5b60405160a0810181811067ffffffffffffffff821117156132e3576132e3612ee3565b806040525082518152602083015160208201526040830151604082015260608301516060820152608083015160808201528091505092915050565b6020808252602f908201527f4f6e6c79204e43505374616b696e6720636f6e74726163742063616e2063616c60408201526e36103a3434b990333ab731ba34b7b760891b606082015260800190565b6020808252602a908201527f4465706f73697420616d6f756e742073686f756c642062652067726561746572604082015269207468616e207a65726f60b01b606082015260800190565b6000602082840312156133c957600080fd5b8151610ebb81612e33565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60005b8381101561343a578181015183820152602001613422565b83811115612c595750506000910152565b6000825161345d81846020870161341f565b9190910192915050565b602081526000825180602084015261348681604085016020870161341f565b601f01601f1916919091016040019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122013c16b0bae42160552b3be738d1a01aeee0a159c353d7f57cf558d5df1a3758564736f6c634300080e0033",
}

// StakingImpABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingImpMetaData.ABI instead.
var StakingImpABI = StakingImpMetaData.ABI

// Deprecated: Use StakingImpMetaData.Sigs instead.
// StakingImpFuncSigs maps the 4-byte function signature to its string representation.
var StakingImpFuncSigs = StakingImpMetaData.Sigs

// StakingImpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingImpMetaData.Bin instead.
var StakingImpBin = StakingImpMetaData.Bin

// DeployStakingImp deploys a new Ethereum contract, binding an instance of StakingImp to it.
func DeployStakingImp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StakingImp, error) {
	parsed, err := StakingImpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingImpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingImp{StakingImpCaller: StakingImpCaller{contract: contract}, StakingImpTransactor: StakingImpTransactor{contract: contract}, StakingImpFilterer: StakingImpFilterer{contract: contract}}, nil
}

// StakingImp is an auto generated Go binding around an Ethereum contract.
type StakingImp struct {
	StakingImpCaller     // Read-only binding to the contract
	StakingImpTransactor // Write-only binding to the contract
	StakingImpFilterer   // Log filterer for contract events
}

// StakingImpCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingImpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingImpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingImpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingImpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingImpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingImpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingImpSession struct {
	Contract     *StakingImp       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingImpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingImpCallerSession struct {
	Contract *StakingImpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StakingImpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingImpTransactorSession struct {
	Contract     *StakingImpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StakingImpRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingImpRaw struct {
	Contract *StakingImp // Generic contract binding to access the raw methods on
}

// StakingImpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingImpCallerRaw struct {
	Contract *StakingImpCaller // Generic read-only contract binding to access the raw methods on
}

// StakingImpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingImpTransactorRaw struct {
	Contract *StakingImpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingImp creates a new instance of StakingImp, bound to a specific deployed contract.
func NewStakingImp(address common.Address, backend bind.ContractBackend) (*StakingImp, error) {
	contract, err := bindStakingImp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingImp{StakingImpCaller: StakingImpCaller{contract: contract}, StakingImpTransactor: StakingImpTransactor{contract: contract}, StakingImpFilterer: StakingImpFilterer{contract: contract}}, nil
}

// NewStakingImpCaller creates a new read-only instance of StakingImp, bound to a specific deployed contract.
func NewStakingImpCaller(address common.Address, caller bind.ContractCaller) (*StakingImpCaller, error) {
	contract, err := bindStakingImp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingImpCaller{contract: contract}, nil
}

// NewStakingImpTransactor creates a new write-only instance of StakingImp, bound to a specific deployed contract.
func NewStakingImpTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingImpTransactor, error) {
	contract, err := bindStakingImp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingImpTransactor{contract: contract}, nil
}

// NewStakingImpFilterer creates a new log filterer instance of StakingImp, bound to a specific deployed contract.
func NewStakingImpFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingImpFilterer, error) {
	contract, err := bindStakingImp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingImpFilterer{contract: contract}, nil
}

// bindStakingImp binds a generic wrapper to an already deployed contract.
func bindStakingImp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingImpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingImp *StakingImpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingImp.Contract.StakingImpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingImp *StakingImpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingImp.Contract.StakingImpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingImp *StakingImpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingImp.Contract.StakingImpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingImp *StakingImpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingImp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingImp *StakingImpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingImp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingImp *StakingImpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingImp.Contract.contract.Transact(opts, method, params...)
}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCaller) BALLOTSTORAGENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "BALLOT_STORAGE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_StakingImp *StakingImpSession) BALLOTSTORAGENAME() ([32]byte, error) {
	return _StakingImp.Contract.BALLOTSTORAGENAME(&_StakingImp.CallOpts)
}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCallerSession) BALLOTSTORAGENAME() ([32]byte, error) {
	return _StakingImp.Contract.BALLOTSTORAGENAME(&_StakingImp.CallOpts)
}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCaller) ECOSYSTEMNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "ECOSYSTEM_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_StakingImp *StakingImpSession) ECOSYSTEMNAME() ([32]byte, error) {
	return _StakingImp.Contract.ECOSYSTEMNAME(&_StakingImp.CallOpts)
}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCallerSession) ECOSYSTEMNAME() ([32]byte, error) {
	return _StakingImp.Contract.ECOSYSTEMNAME(&_StakingImp.CallOpts)
}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCaller) ENVSTORAGENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "ENV_STORAGE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_StakingImp *StakingImpSession) ENVSTORAGENAME() ([32]byte, error) {
	return _StakingImp.Contract.ENVSTORAGENAME(&_StakingImp.CallOpts)
}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCallerSession) ENVSTORAGENAME() ([32]byte, error) {
	return _StakingImp.Contract.ENVSTORAGENAME(&_StakingImp.CallOpts)
}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCaller) GOVNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "GOV_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_StakingImp *StakingImpSession) GOVNAME() ([32]byte, error) {
	return _StakingImp.Contract.GOVNAME(&_StakingImp.CallOpts)
}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCallerSession) GOVNAME() ([32]byte, error) {
	return _StakingImp.Contract.GOVNAME(&_StakingImp.CallOpts)
}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCaller) MAINTENANCENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "MAINTENANCE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_StakingImp *StakingImpSession) MAINTENANCENAME() ([32]byte, error) {
	return _StakingImp.Contract.MAINTENANCENAME(&_StakingImp.CallOpts)
}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCallerSession) MAINTENANCENAME() ([32]byte, error) {
	return _StakingImp.Contract.MAINTENANCENAME(&_StakingImp.CallOpts)
}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCaller) REWARDPOOLNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "REWARD_POOL_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_StakingImp *StakingImpSession) REWARDPOOLNAME() ([32]byte, error) {
	return _StakingImp.Contract.REWARDPOOLNAME(&_StakingImp.CallOpts)
}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCallerSession) REWARDPOOLNAME() ([32]byte, error) {
	return _StakingImp.Contract.REWARDPOOLNAME(&_StakingImp.CallOpts)
}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCaller) STAKINGNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "STAKING_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_StakingImp *StakingImpSession) STAKINGNAME() ([32]byte, error) {
	return _StakingImp.Contract.STAKINGNAME(&_StakingImp.CallOpts)
}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCallerSession) STAKINGNAME() ([32]byte, error) {
	return _StakingImp.Contract.STAKINGNAME(&_StakingImp.CallOpts)
}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCaller) STAKINGREWARDNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "STAKING_REWARD_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_StakingImp *StakingImpSession) STAKINGREWARDNAME() ([32]byte, error) {
	return _StakingImp.Contract.STAKINGREWARDNAME(&_StakingImp.CallOpts)
}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_StakingImp *StakingImpCallerSession) STAKINGREWARDNAME() ([32]byte, error) {
	return _StakingImp.Contract.STAKINGREWARDNAME(&_StakingImp.CallOpts)
}

// AvailableBalanceOf is a free data retrieval call binding the contract method 0x25d998bb.
//
// Solidity: function availableBalanceOf(address payee) view returns(uint256)
func (_StakingImp *StakingImpCaller) AvailableBalanceOf(opts *bind.CallOpts, payee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "availableBalanceOf", payee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableBalanceOf is a free data retrieval call binding the contract method 0x25d998bb.
//
// Solidity: function availableBalanceOf(address payee) view returns(uint256)
func (_StakingImp *StakingImpSession) AvailableBalanceOf(payee common.Address) (*big.Int, error) {
	return _StakingImp.Contract.AvailableBalanceOf(&_StakingImp.CallOpts, payee)
}

// AvailableBalanceOf is a free data retrieval call binding the contract method 0x25d998bb.
//
// Solidity: function availableBalanceOf(address payee) view returns(uint256)
func (_StakingImp *StakingImpCallerSession) AvailableBalanceOf(payee common.Address) (*big.Int, error) {
	return _StakingImp.Contract.AvailableBalanceOf(&_StakingImp.CallOpts, payee)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address payee) view returns(uint256)
func (_StakingImp *StakingImpCaller) BalanceOf(opts *bind.CallOpts, payee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "balanceOf", payee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address payee) view returns(uint256)
func (_StakingImp *StakingImpSession) BalanceOf(payee common.Address) (*big.Int, error) {
	return _StakingImp.Contract.BalanceOf(&_StakingImp.CallOpts, payee)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address payee) view returns(uint256)
func (_StakingImp *StakingImpCallerSession) BalanceOf(payee common.Address) (*big.Int, error) {
	return _StakingImp.Contract.BalanceOf(&_StakingImp.CallOpts, payee)
}

// CalcVotingWeight is a free data retrieval call binding the contract method 0x884d97a7.
//
// Solidity: function calcVotingWeight(address payee) view returns(uint256)
func (_StakingImp *StakingImpCaller) CalcVotingWeight(opts *bind.CallOpts, payee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "calcVotingWeight", payee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcVotingWeight is a free data retrieval call binding the contract method 0x884d97a7.
//
// Solidity: function calcVotingWeight(address payee) view returns(uint256)
func (_StakingImp *StakingImpSession) CalcVotingWeight(payee common.Address) (*big.Int, error) {
	return _StakingImp.Contract.CalcVotingWeight(&_StakingImp.CallOpts, payee)
}

// CalcVotingWeight is a free data retrieval call binding the contract method 0x884d97a7.
//
// Solidity: function calcVotingWeight(address payee) view returns(uint256)
func (_StakingImp *StakingImpCallerSession) CalcVotingWeight(payee common.Address) (*big.Int, error) {
	return _StakingImp.Contract.CalcVotingWeight(&_StakingImp.CallOpts, payee)
}

// CalcVotingWeightWithScaleFactor is a free data retrieval call binding the contract method 0x7d77a0eb.
//
// Solidity: function calcVotingWeightWithScaleFactor(address payee, uint32 factor) view returns(uint256)
func (_StakingImp *StakingImpCaller) CalcVotingWeightWithScaleFactor(opts *bind.CallOpts, payee common.Address, factor uint32) (*big.Int, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "calcVotingWeightWithScaleFactor", payee, factor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcVotingWeightWithScaleFactor is a free data retrieval call binding the contract method 0x7d77a0eb.
//
// Solidity: function calcVotingWeightWithScaleFactor(address payee, uint32 factor) view returns(uint256)
func (_StakingImp *StakingImpSession) CalcVotingWeightWithScaleFactor(payee common.Address, factor uint32) (*big.Int, error) {
	return _StakingImp.Contract.CalcVotingWeightWithScaleFactor(&_StakingImp.CallOpts, payee, factor)
}

// CalcVotingWeightWithScaleFactor is a free data retrieval call binding the contract method 0x7d77a0eb.
//
// Solidity: function calcVotingWeightWithScaleFactor(address payee, uint32 factor) view returns(uint256)
func (_StakingImp *StakingImpCallerSession) CalcVotingWeightWithScaleFactor(payee common.Address, factor uint32) (*big.Int, error) {
	return _StakingImp.Contract.CalcVotingWeightWithScaleFactor(&_StakingImp.CallOpts, payee, factor)
}

// GetRatioOfUserBalance is a free data retrieval call binding the contract method 0x2b0b9c5e.
//
// Solidity: function getRatioOfUserBalance(address ncp) view returns(uint256)
func (_StakingImp *StakingImpCaller) GetRatioOfUserBalance(opts *bind.CallOpts, ncp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "getRatioOfUserBalance", ncp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRatioOfUserBalance is a free data retrieval call binding the contract method 0x2b0b9c5e.
//
// Solidity: function getRatioOfUserBalance(address ncp) view returns(uint256)
func (_StakingImp *StakingImpSession) GetRatioOfUserBalance(ncp common.Address) (*big.Int, error) {
	return _StakingImp.Contract.GetRatioOfUserBalance(&_StakingImp.CallOpts, ncp)
}

// GetRatioOfUserBalance is a free data retrieval call binding the contract method 0x2b0b9c5e.
//
// Solidity: function getRatioOfUserBalance(address ncp) view returns(uint256)
func (_StakingImp *StakingImpCallerSession) GetRatioOfUserBalance(ncp common.Address) (*big.Int, error) {
	return _StakingImp.Contract.GetRatioOfUserBalance(&_StakingImp.CallOpts, ncp)
}

// GetTotalLockedBalance is a free data retrieval call binding the contract method 0xd5c25890.
//
// Solidity: function getTotalLockedBalance() view returns(uint256)
func (_StakingImp *StakingImpCaller) GetTotalLockedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "getTotalLockedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalLockedBalance is a free data retrieval call binding the contract method 0xd5c25890.
//
// Solidity: function getTotalLockedBalance() view returns(uint256)
func (_StakingImp *StakingImpSession) GetTotalLockedBalance() (*big.Int, error) {
	return _StakingImp.Contract.GetTotalLockedBalance(&_StakingImp.CallOpts)
}

// GetTotalLockedBalance is a free data retrieval call binding the contract method 0xd5c25890.
//
// Solidity: function getTotalLockedBalance() view returns(uint256)
func (_StakingImp *StakingImpCallerSession) GetTotalLockedBalance() (*big.Int, error) {
	return _StakingImp.Contract.GetTotalLockedBalance(&_StakingImp.CallOpts)
}

// IsRevoked is a free data retrieval call binding the contract method 0x2bc9ed02.
//
// Solidity: function isRevoked() view returns(bool)
func (_StakingImp *StakingImpCaller) IsRevoked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "isRevoked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRevoked is a free data retrieval call binding the contract method 0x2bc9ed02.
//
// Solidity: function isRevoked() view returns(bool)
func (_StakingImp *StakingImpSession) IsRevoked() (bool, error) {
	return _StakingImp.Contract.IsRevoked(&_StakingImp.CallOpts)
}

// IsRevoked is a free data retrieval call binding the contract method 0x2bc9ed02.
//
// Solidity: function isRevoked() view returns(bool)
func (_StakingImp *StakingImpCallerSession) IsRevoked() (bool, error) {
	return _StakingImp.Contract.IsRevoked(&_StakingImp.CallOpts)
}

// LockedBalanceOf is a free data retrieval call binding the contract method 0x59355736.
//
// Solidity: function lockedBalanceOf(address payee) view returns(uint256)
func (_StakingImp *StakingImpCaller) LockedBalanceOf(opts *bind.CallOpts, payee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "lockedBalanceOf", payee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedBalanceOf is a free data retrieval call binding the contract method 0x59355736.
//
// Solidity: function lockedBalanceOf(address payee) view returns(uint256)
func (_StakingImp *StakingImpSession) LockedBalanceOf(payee common.Address) (*big.Int, error) {
	return _StakingImp.Contract.LockedBalanceOf(&_StakingImp.CallOpts, payee)
}

// LockedBalanceOf is a free data retrieval call binding the contract method 0x59355736.
//
// Solidity: function lockedBalanceOf(address payee) view returns(uint256)
func (_StakingImp *StakingImpCallerSession) LockedBalanceOf(payee common.Address) (*big.Int, error) {
	return _StakingImp.Contract.LockedBalanceOf(&_StakingImp.CallOpts, payee)
}

// NcpStaking is a free data retrieval call binding the contract method 0xf1b8aa1d.
//
// Solidity: function ncpStaking() view returns(address)
func (_StakingImp *StakingImpCaller) NcpStaking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "ncpStaking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NcpStaking is a free data retrieval call binding the contract method 0xf1b8aa1d.
//
// Solidity: function ncpStaking() view returns(address)
func (_StakingImp *StakingImpSession) NcpStaking() (common.Address, error) {
	return _StakingImp.Contract.NcpStaking(&_StakingImp.CallOpts)
}

// NcpStaking is a free data retrieval call binding the contract method 0xf1b8aa1d.
//
// Solidity: function ncpStaking() view returns(address)
func (_StakingImp *StakingImpCallerSession) NcpStaking() (common.Address, error) {
	return _StakingImp.Contract.NcpStaking(&_StakingImp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingImp *StakingImpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingImp *StakingImpSession) Owner() (common.Address, error) {
	return _StakingImp.Contract.Owner(&_StakingImp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingImp *StakingImpCallerSession) Owner() (common.Address, error) {
	return _StakingImp.Contract.Owner(&_StakingImp.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingImp *StakingImpCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingImp *StakingImpSession) ProxiableUUID() ([32]byte, error) {
	return _StakingImp.Contract.ProxiableUUID(&_StakingImp.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingImp *StakingImpCallerSession) ProxiableUUID() ([32]byte, error) {
	return _StakingImp.Contract.ProxiableUUID(&_StakingImp.CallOpts)
}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_StakingImp *StakingImpCaller) Reg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "reg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_StakingImp *StakingImpSession) Reg() (common.Address, error) {
	return _StakingImp.Contract.Reg(&_StakingImp.CallOpts)
}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_StakingImp *StakingImpCallerSession) Reg() (common.Address, error) {
	return _StakingImp.Contract.Reg(&_StakingImp.CallOpts)
}

// UserBalanceOf is a free data retrieval call binding the contract method 0x12853615.
//
// Solidity: function userBalanceOf(address ncp, address user) view returns(uint256)
func (_StakingImp *StakingImpCaller) UserBalanceOf(opts *bind.CallOpts, ncp common.Address, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "userBalanceOf", ncp, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBalanceOf is a free data retrieval call binding the contract method 0x12853615.
//
// Solidity: function userBalanceOf(address ncp, address user) view returns(uint256)
func (_StakingImp *StakingImpSession) UserBalanceOf(ncp common.Address, user common.Address) (*big.Int, error) {
	return _StakingImp.Contract.UserBalanceOf(&_StakingImp.CallOpts, ncp, user)
}

// UserBalanceOf is a free data retrieval call binding the contract method 0x12853615.
//
// Solidity: function userBalanceOf(address ncp, address user) view returns(uint256)
func (_StakingImp *StakingImpCallerSession) UserBalanceOf(ncp common.Address, user common.Address) (*big.Int, error) {
	return _StakingImp.Contract.UserBalanceOf(&_StakingImp.CallOpts, ncp, user)
}

// UserTotalBalanceOf is a free data retrieval call binding the contract method 0x193468ac.
//
// Solidity: function userTotalBalanceOf(address ncp) view returns(uint256)
func (_StakingImp *StakingImpCaller) UserTotalBalanceOf(opts *bind.CallOpts, ncp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingImp.contract.Call(opts, &out, "userTotalBalanceOf", ncp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTotalBalanceOf is a free data retrieval call binding the contract method 0x193468ac.
//
// Solidity: function userTotalBalanceOf(address ncp) view returns(uint256)
func (_StakingImp *StakingImpSession) UserTotalBalanceOf(ncp common.Address) (*big.Int, error) {
	return _StakingImp.Contract.UserTotalBalanceOf(&_StakingImp.CallOpts, ncp)
}

// UserTotalBalanceOf is a free data retrieval call binding the contract method 0x193468ac.
//
// Solidity: function userTotalBalanceOf(address ncp) view returns(uint256)
func (_StakingImp *StakingImpCallerSession) UserTotalBalanceOf(ncp common.Address) (*big.Int, error) {
	return _StakingImp.Contract.UserTotalBalanceOf(&_StakingImp.CallOpts, ncp)
}

// DelegateDepositAndLockMore is a paid mutator transaction binding the contract method 0xbac4f338.
//
// Solidity: function delegateDepositAndLockMore(address ncp) payable returns()
func (_StakingImp *StakingImpTransactor) DelegateDepositAndLockMore(opts *bind.TransactOpts, ncp common.Address) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "delegateDepositAndLockMore", ncp)
}

// DelegateDepositAndLockMore is a paid mutator transaction binding the contract method 0xbac4f338.
//
// Solidity: function delegateDepositAndLockMore(address ncp) payable returns()
func (_StakingImp *StakingImpSession) DelegateDepositAndLockMore(ncp common.Address) (*types.Transaction, error) {
	return _StakingImp.Contract.DelegateDepositAndLockMore(&_StakingImp.TransactOpts, ncp)
}

// DelegateDepositAndLockMore is a paid mutator transaction binding the contract method 0xbac4f338.
//
// Solidity: function delegateDepositAndLockMore(address ncp) payable returns()
func (_StakingImp *StakingImpTransactorSession) DelegateDepositAndLockMore(ncp common.Address) (*types.Transaction, error) {
	return _StakingImp.Contract.DelegateDepositAndLockMore(&_StakingImp.TransactOpts, ncp)
}

// DelegateUnlockAndWithdraw is a paid mutator transaction binding the contract method 0xf6931822.
//
// Solidity: function delegateUnlockAndWithdraw(address ncp, uint256 amount) returns()
func (_StakingImp *StakingImpTransactor) DelegateUnlockAndWithdraw(opts *bind.TransactOpts, ncp common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "delegateUnlockAndWithdraw", ncp, amount)
}

// DelegateUnlockAndWithdraw is a paid mutator transaction binding the contract method 0xf6931822.
//
// Solidity: function delegateUnlockAndWithdraw(address ncp, uint256 amount) returns()
func (_StakingImp *StakingImpSession) DelegateUnlockAndWithdraw(ncp common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingImp.Contract.DelegateUnlockAndWithdraw(&_StakingImp.TransactOpts, ncp, amount)
}

// DelegateUnlockAndWithdraw is a paid mutator transaction binding the contract method 0xf6931822.
//
// Solidity: function delegateUnlockAndWithdraw(address ncp, uint256 amount) returns()
func (_StakingImp *StakingImpTransactorSession) DelegateUnlockAndWithdraw(ncp common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingImp.Contract.DelegateUnlockAndWithdraw(&_StakingImp.TransactOpts, ncp, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_StakingImp *StakingImpTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_StakingImp *StakingImpSession) Deposit() (*types.Transaction, error) {
	return _StakingImp.Contract.Deposit(&_StakingImp.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_StakingImp *StakingImpTransactorSession) Deposit() (*types.Transaction, error) {
	return _StakingImp.Contract.Deposit(&_StakingImp.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xc0d91eaf.
//
// Solidity: function init(address registry, bytes data) returns()
func (_StakingImp *StakingImpTransactor) Init(opts *bind.TransactOpts, registry common.Address, data []byte) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "init", registry, data)
}

// Init is a paid mutator transaction binding the contract method 0xc0d91eaf.
//
// Solidity: function init(address registry, bytes data) returns()
func (_StakingImp *StakingImpSession) Init(registry common.Address, data []byte) (*types.Transaction, error) {
	return _StakingImp.Contract.Init(&_StakingImp.TransactOpts, registry, data)
}

// Init is a paid mutator transaction binding the contract method 0xc0d91eaf.
//
// Solidity: function init(address registry, bytes data) returns()
func (_StakingImp *StakingImpTransactorSession) Init(registry common.Address, data []byte) (*types.Transaction, error) {
	return _StakingImp.Contract.Init(&_StakingImp.TransactOpts, registry, data)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address payee, uint256 lockAmount) returns()
func (_StakingImp *StakingImpTransactor) Lock(opts *bind.TransactOpts, payee common.Address, lockAmount *big.Int) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "lock", payee, lockAmount)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address payee, uint256 lockAmount) returns()
func (_StakingImp *StakingImpSession) Lock(payee common.Address, lockAmount *big.Int) (*types.Transaction, error) {
	return _StakingImp.Contract.Lock(&_StakingImp.TransactOpts, payee, lockAmount)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address payee, uint256 lockAmount) returns()
func (_StakingImp *StakingImpTransactorSession) Lock(payee common.Address, lockAmount *big.Int) (*types.Transaction, error) {
	return _StakingImp.Contract.Lock(&_StakingImp.TransactOpts, payee, lockAmount)
}

// LockMore is a paid mutator transaction binding the contract method 0x9667e76a.
//
// Solidity: function lockMore(uint256 lockAmount) returns()
func (_StakingImp *StakingImpTransactor) LockMore(opts *bind.TransactOpts, lockAmount *big.Int) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "lockMore", lockAmount)
}

// LockMore is a paid mutator transaction binding the contract method 0x9667e76a.
//
// Solidity: function lockMore(uint256 lockAmount) returns()
func (_StakingImp *StakingImpSession) LockMore(lockAmount *big.Int) (*types.Transaction, error) {
	return _StakingImp.Contract.LockMore(&_StakingImp.TransactOpts, lockAmount)
}

// LockMore is a paid mutator transaction binding the contract method 0x9667e76a.
//
// Solidity: function lockMore(uint256 lockAmount) returns()
func (_StakingImp *StakingImpTransactorSession) LockMore(lockAmount *big.Int) (*types.Transaction, error) {
	return _StakingImp.Contract.LockMore(&_StakingImp.TransactOpts, lockAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingImp *StakingImpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingImp *StakingImpSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingImp.Contract.RenounceOwnership(&_StakingImp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingImp *StakingImpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingImp.Contract.RenounceOwnership(&_StakingImp.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_StakingImp *StakingImpTransactor) Revoke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "revoke")
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_StakingImp *StakingImpSession) Revoke() (*types.Transaction, error) {
	return _StakingImp.Contract.Revoke(&_StakingImp.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_StakingImp *StakingImpTransactorSession) Revoke() (*types.Transaction, error) {
	return _StakingImp.Contract.Revoke(&_StakingImp.TransactOpts)
}

// SetNCPStaking is a paid mutator transaction binding the contract method 0x7f2f4c06.
//
// Solidity: function setNCPStaking(address _ncpStaking) returns()
func (_StakingImp *StakingImpTransactor) SetNCPStaking(opts *bind.TransactOpts, _ncpStaking common.Address) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "setNCPStaking", _ncpStaking)
}

// SetNCPStaking is a paid mutator transaction binding the contract method 0x7f2f4c06.
//
// Solidity: function setNCPStaking(address _ncpStaking) returns()
func (_StakingImp *StakingImpSession) SetNCPStaking(_ncpStaking common.Address) (*types.Transaction, error) {
	return _StakingImp.Contract.SetNCPStaking(&_StakingImp.TransactOpts, _ncpStaking)
}

// SetNCPStaking is a paid mutator transaction binding the contract method 0x7f2f4c06.
//
// Solidity: function setNCPStaking(address _ncpStaking) returns()
func (_StakingImp *StakingImpTransactorSession) SetNCPStaking(_ncpStaking common.Address) (*types.Transaction, error) {
	return _StakingImp.Contract.SetNCPStaking(&_StakingImp.TransactOpts, _ncpStaking)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_StakingImp *StakingImpTransactor) SetRegistry(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "setRegistry", _addr)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_StakingImp *StakingImpSession) SetRegistry(_addr common.Address) (*types.Transaction, error) {
	return _StakingImp.Contract.SetRegistry(&_StakingImp.TransactOpts, _addr)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_StakingImp *StakingImpTransactorSession) SetRegistry(_addr common.Address) (*types.Transaction, error) {
	return _StakingImp.Contract.SetRegistry(&_StakingImp.TransactOpts, _addr)
}

// TransferLocked is a paid mutator transaction binding the contract method 0x92a950b6.
//
// Solidity: function transferLocked(address from, uint256 slashing, uint256 ext) returns()
func (_StakingImp *StakingImpTransactor) TransferLocked(opts *bind.TransactOpts, from common.Address, slashing *big.Int, ext *big.Int) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "transferLocked", from, slashing, ext)
}

// TransferLocked is a paid mutator transaction binding the contract method 0x92a950b6.
//
// Solidity: function transferLocked(address from, uint256 slashing, uint256 ext) returns()
func (_StakingImp *StakingImpSession) TransferLocked(from common.Address, slashing *big.Int, ext *big.Int) (*types.Transaction, error) {
	return _StakingImp.Contract.TransferLocked(&_StakingImp.TransactOpts, from, slashing, ext)
}

// TransferLocked is a paid mutator transaction binding the contract method 0x92a950b6.
//
// Solidity: function transferLocked(address from, uint256 slashing, uint256 ext) returns()
func (_StakingImp *StakingImpTransactorSession) TransferLocked(from common.Address, slashing *big.Int, ext *big.Int) (*types.Transaction, error) {
	return _StakingImp.Contract.TransferLocked(&_StakingImp.TransactOpts, from, slashing, ext)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingImp *StakingImpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingImp *StakingImpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingImp.Contract.TransferOwnership(&_StakingImp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingImp *StakingImpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingImp.Contract.TransferOwnership(&_StakingImp.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address payee, uint256 unlockAmount) returns()
func (_StakingImp *StakingImpTransactor) Unlock(opts *bind.TransactOpts, payee common.Address, unlockAmount *big.Int) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "unlock", payee, unlockAmount)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address payee, uint256 unlockAmount) returns()
func (_StakingImp *StakingImpSession) Unlock(payee common.Address, unlockAmount *big.Int) (*types.Transaction, error) {
	return _StakingImp.Contract.Unlock(&_StakingImp.TransactOpts, payee, unlockAmount)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address payee, uint256 unlockAmount) returns()
func (_StakingImp *StakingImpTransactorSession) Unlock(payee common.Address, unlockAmount *big.Int) (*types.Transaction, error) {
	return _StakingImp.Contract.Unlock(&_StakingImp.TransactOpts, payee, unlockAmount)
}

// UpgradeStaking is a paid mutator transaction binding the contract method 0xf3f63080.
//
// Solidity: function upgradeStaking(address newImp) returns()
func (_StakingImp *StakingImpTransactor) UpgradeStaking(opts *bind.TransactOpts, newImp common.Address) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "upgradeStaking", newImp)
}

// UpgradeStaking is a paid mutator transaction binding the contract method 0xf3f63080.
//
// Solidity: function upgradeStaking(address newImp) returns()
func (_StakingImp *StakingImpSession) UpgradeStaking(newImp common.Address) (*types.Transaction, error) {
	return _StakingImp.Contract.UpgradeStaking(&_StakingImp.TransactOpts, newImp)
}

// UpgradeStaking is a paid mutator transaction binding the contract method 0xf3f63080.
//
// Solidity: function upgradeStaking(address newImp) returns()
func (_StakingImp *StakingImpTransactorSession) UpgradeStaking(newImp common.Address) (*types.Transaction, error) {
	return _StakingImp.Contract.UpgradeStaking(&_StakingImp.TransactOpts, newImp)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StakingImp *StakingImpTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StakingImp *StakingImpSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StakingImp.Contract.UpgradeTo(&_StakingImp.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StakingImp *StakingImpTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StakingImp.Contract.UpgradeTo(&_StakingImp.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingImp *StakingImpTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingImp *StakingImpSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingImp.Contract.UpgradeToAndCall(&_StakingImp.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingImp *StakingImpTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingImp.Contract.UpgradeToAndCall(&_StakingImp.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_StakingImp *StakingImpTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakingImp.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_StakingImp *StakingImpSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _StakingImp.Contract.Withdraw(&_StakingImp.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_StakingImp *StakingImpTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _StakingImp.Contract.Withdraw(&_StakingImp.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingImp *StakingImpTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingImp.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingImp *StakingImpSession) Receive() (*types.Transaction, error) {
	return _StakingImp.Contract.Receive(&_StakingImp.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingImp *StakingImpTransactorSession) Receive() (*types.Transaction, error) {
	return _StakingImp.Contract.Receive(&_StakingImp.TransactOpts)
}

// StakingImpAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the StakingImp contract.
type StakingImpAdminChangedIterator struct {
	Event *StakingImpAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingImpAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpAdminChanged)
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
		it.Event = new(StakingImpAdminChanged)
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
func (it *StakingImpAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpAdminChanged represents a AdminChanged event raised by the StakingImp contract.
type StakingImpAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StakingImp *StakingImpFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*StakingImpAdminChangedIterator, error) {

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &StakingImpAdminChangedIterator{contract: _StakingImp.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StakingImp *StakingImpFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingImpAdminChanged) (event.Subscription, error) {

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpAdminChanged)
				if err := _StakingImp.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_StakingImp *StakingImpFilterer) ParseAdminChanged(log types.Log) (*StakingImpAdminChanged, error) {
	event := new(StakingImpAdminChanged)
	if err := _StakingImp.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the StakingImp contract.
type StakingImpBeaconUpgradedIterator struct {
	Event *StakingImpBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *StakingImpBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpBeaconUpgraded)
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
		it.Event = new(StakingImpBeaconUpgraded)
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
func (it *StakingImpBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpBeaconUpgraded represents a BeaconUpgraded event raised by the StakingImp contract.
type StakingImpBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StakingImp *StakingImpFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*StakingImpBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &StakingImpBeaconUpgradedIterator{contract: _StakingImp.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StakingImp *StakingImpFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *StakingImpBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpBeaconUpgraded)
				if err := _StakingImp.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_StakingImp *StakingImpFilterer) ParseBeaconUpgraded(log types.Log) (*StakingImpBeaconUpgraded, error) {
	event := new(StakingImpBeaconUpgraded)
	if err := _StakingImp.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpDelegateStakedIterator is returned from FilterDelegateStaked and is used to iterate over the raw logs and unpacked data for DelegateStaked events raised by the StakingImp contract.
type StakingImpDelegateStakedIterator struct {
	Event *StakingImpDelegateStaked // Event containing the contract specifics and raw log

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
func (it *StakingImpDelegateStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpDelegateStaked)
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
		it.Event = new(StakingImpDelegateStaked)
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
func (it *StakingImpDelegateStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpDelegateStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpDelegateStaked represents a DelegateStaked event raised by the StakingImp contract.
type StakingImpDelegateStaked struct {
	Payee           common.Address
	Amount          *big.Int
	Ncp             common.Address
	NcpTotalLocked  *big.Int
	UserTotalLocked *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateStaked is a free log retrieval operation binding the contract event 0x74cfc20f0e6d14384c3a60820d3e814f6979d009cdbb43db27fa56fe475172fd.
//
// Solidity: event DelegateStaked(address indexed payee, uint256 amount, address indexed ncp, uint256 ncpTotalLocked, uint256 userTotalLocked)
func (_StakingImp *StakingImpFilterer) FilterDelegateStaked(opts *bind.FilterOpts, payee []common.Address, ncp []common.Address) (*StakingImpDelegateStakedIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var ncpRule []interface{}
	for _, ncpItem := range ncp {
		ncpRule = append(ncpRule, ncpItem)
	}

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "DelegateStaked", payeeRule, ncpRule)
	if err != nil {
		return nil, err
	}
	return &StakingImpDelegateStakedIterator{contract: _StakingImp.contract, event: "DelegateStaked", logs: logs, sub: sub}, nil
}

// WatchDelegateStaked is a free log subscription operation binding the contract event 0x74cfc20f0e6d14384c3a60820d3e814f6979d009cdbb43db27fa56fe475172fd.
//
// Solidity: event DelegateStaked(address indexed payee, uint256 amount, address indexed ncp, uint256 ncpTotalLocked, uint256 userTotalLocked)
func (_StakingImp *StakingImpFilterer) WatchDelegateStaked(opts *bind.WatchOpts, sink chan<- *StakingImpDelegateStaked, payee []common.Address, ncp []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var ncpRule []interface{}
	for _, ncpItem := range ncp {
		ncpRule = append(ncpRule, ncpItem)
	}

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "DelegateStaked", payeeRule, ncpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpDelegateStaked)
				if err := _StakingImp.contract.UnpackLog(event, "DelegateStaked", log); err != nil {
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

// ParseDelegateStaked is a log parse operation binding the contract event 0x74cfc20f0e6d14384c3a60820d3e814f6979d009cdbb43db27fa56fe475172fd.
//
// Solidity: event DelegateStaked(address indexed payee, uint256 amount, address indexed ncp, uint256 ncpTotalLocked, uint256 userTotalLocked)
func (_StakingImp *StakingImpFilterer) ParseDelegateStaked(log types.Log) (*StakingImpDelegateStaked, error) {
	event := new(StakingImpDelegateStaked)
	if err := _StakingImp.contract.UnpackLog(event, "DelegateStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpDelegateUnstakedIterator is returned from FilterDelegateUnstaked and is used to iterate over the raw logs and unpacked data for DelegateUnstaked events raised by the StakingImp contract.
type StakingImpDelegateUnstakedIterator struct {
	Event *StakingImpDelegateUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingImpDelegateUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpDelegateUnstaked)
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
		it.Event = new(StakingImpDelegateUnstaked)
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
func (it *StakingImpDelegateUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpDelegateUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpDelegateUnstaked represents a DelegateUnstaked event raised by the StakingImp contract.
type StakingImpDelegateUnstaked struct {
	Payee           common.Address
	Amount          *big.Int
	Ncp             common.Address
	NcpTotalLocked  *big.Int
	UserTotalLocked *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateUnstaked is a free log retrieval operation binding the contract event 0x03d2bb70c6ccc49d68a465a06edffb976961cf8930888658ca2339fa62b8bda2.
//
// Solidity: event DelegateUnstaked(address indexed payee, uint256 amount, address indexed ncp, uint256 ncpTotalLocked, uint256 userTotalLocked)
func (_StakingImp *StakingImpFilterer) FilterDelegateUnstaked(opts *bind.FilterOpts, payee []common.Address, ncp []common.Address) (*StakingImpDelegateUnstakedIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var ncpRule []interface{}
	for _, ncpItem := range ncp {
		ncpRule = append(ncpRule, ncpItem)
	}

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "DelegateUnstaked", payeeRule, ncpRule)
	if err != nil {
		return nil, err
	}
	return &StakingImpDelegateUnstakedIterator{contract: _StakingImp.contract, event: "DelegateUnstaked", logs: logs, sub: sub}, nil
}

// WatchDelegateUnstaked is a free log subscription operation binding the contract event 0x03d2bb70c6ccc49d68a465a06edffb976961cf8930888658ca2339fa62b8bda2.
//
// Solidity: event DelegateUnstaked(address indexed payee, uint256 amount, address indexed ncp, uint256 ncpTotalLocked, uint256 userTotalLocked)
func (_StakingImp *StakingImpFilterer) WatchDelegateUnstaked(opts *bind.WatchOpts, sink chan<- *StakingImpDelegateUnstaked, payee []common.Address, ncp []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var ncpRule []interface{}
	for _, ncpItem := range ncp {
		ncpRule = append(ncpRule, ncpItem)
	}

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "DelegateUnstaked", payeeRule, ncpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpDelegateUnstaked)
				if err := _StakingImp.contract.UnpackLog(event, "DelegateUnstaked", log); err != nil {
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

// ParseDelegateUnstaked is a log parse operation binding the contract event 0x03d2bb70c6ccc49d68a465a06edffb976961cf8930888658ca2339fa62b8bda2.
//
// Solidity: event DelegateUnstaked(address indexed payee, uint256 amount, address indexed ncp, uint256 ncpTotalLocked, uint256 userTotalLocked)
func (_StakingImp *StakingImpFilterer) ParseDelegateUnstaked(log types.Log) (*StakingImpDelegateUnstaked, error) {
	event := new(StakingImpDelegateUnstaked)
	if err := _StakingImp.contract.UnpackLog(event, "DelegateUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingImp contract.
type StakingImpInitializedIterator struct {
	Event *StakingImpInitialized // Event containing the contract specifics and raw log

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
func (it *StakingImpInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpInitialized)
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
		it.Event = new(StakingImpInitialized)
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
func (it *StakingImpInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpInitialized represents a Initialized event raised by the StakingImp contract.
type StakingImpInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingImp *StakingImpFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingImpInitializedIterator, error) {

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingImpInitializedIterator{contract: _StakingImp.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingImp *StakingImpFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingImpInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpInitialized)
				if err := _StakingImp.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakingImp *StakingImpFilterer) ParseInitialized(log types.Log) (*StakingImpInitialized, error) {
	event := new(StakingImpInitialized)
	if err := _StakingImp.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the StakingImp contract.
type StakingImpLockedIterator struct {
	Event *StakingImpLocked // Event containing the contract specifics and raw log

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
func (it *StakingImpLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpLocked)
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
		it.Event = new(StakingImpLocked)
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
func (it *StakingImpLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpLocked represents a Locked event raised by the StakingImp contract.
type StakingImpLocked struct {
	Payee     common.Address
	Amount    *big.Int
	Total     *big.Int
	Available *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x44cebfefa4561bee5b61d675ccfd8dc9969fff9cc15e7a4eccccd62af94f9c11.
//
// Solidity: event Locked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) FilterLocked(opts *bind.FilterOpts, payee []common.Address) (*StakingImpLockedIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "Locked", payeeRule)
	if err != nil {
		return nil, err
	}
	return &StakingImpLockedIterator{contract: _StakingImp.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x44cebfefa4561bee5b61d675ccfd8dc9969fff9cc15e7a4eccccd62af94f9c11.
//
// Solidity: event Locked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *StakingImpLocked, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "Locked", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpLocked)
				if err := _StakingImp.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x44cebfefa4561bee5b61d675ccfd8dc9969fff9cc15e7a4eccccd62af94f9c11.
//
// Solidity: event Locked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) ParseLocked(log types.Log) (*StakingImpLocked, error) {
	event := new(StakingImpLocked)
	if err := _StakingImp.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpNCPAddrChangedIterator is returned from FilterNCPAddrChanged and is used to iterate over the raw logs and unpacked data for NCPAddrChanged events raised by the StakingImp contract.
type StakingImpNCPAddrChangedIterator struct {
	Event *StakingImpNCPAddrChanged // Event containing the contract specifics and raw log

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
func (it *StakingImpNCPAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpNCPAddrChanged)
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
		it.Event = new(StakingImpNCPAddrChanged)
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
func (it *StakingImpNCPAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpNCPAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpNCPAddrChanged represents a NCPAddrChanged event raised by the StakingImp contract.
type StakingImpNCPAddrChanged struct {
	Ncp common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNCPAddrChanged is a free log retrieval operation binding the contract event 0xfd5754300dde6066eda4fabd23616c1d560a3360c85c0716c46e00649bdeeddf.
//
// Solidity: event NCPAddrChanged(address indexed ncp)
func (_StakingImp *StakingImpFilterer) FilterNCPAddrChanged(opts *bind.FilterOpts, ncp []common.Address) (*StakingImpNCPAddrChangedIterator, error) {

	var ncpRule []interface{}
	for _, ncpItem := range ncp {
		ncpRule = append(ncpRule, ncpItem)
	}

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "NCPAddrChanged", ncpRule)
	if err != nil {
		return nil, err
	}
	return &StakingImpNCPAddrChangedIterator{contract: _StakingImp.contract, event: "NCPAddrChanged", logs: logs, sub: sub}, nil
}

// WatchNCPAddrChanged is a free log subscription operation binding the contract event 0xfd5754300dde6066eda4fabd23616c1d560a3360c85c0716c46e00649bdeeddf.
//
// Solidity: event NCPAddrChanged(address indexed ncp)
func (_StakingImp *StakingImpFilterer) WatchNCPAddrChanged(opts *bind.WatchOpts, sink chan<- *StakingImpNCPAddrChanged, ncp []common.Address) (event.Subscription, error) {

	var ncpRule []interface{}
	for _, ncpItem := range ncp {
		ncpRule = append(ncpRule, ncpItem)
	}

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "NCPAddrChanged", ncpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpNCPAddrChanged)
				if err := _StakingImp.contract.UnpackLog(event, "NCPAddrChanged", log); err != nil {
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

// ParseNCPAddrChanged is a log parse operation binding the contract event 0xfd5754300dde6066eda4fabd23616c1d560a3360c85c0716c46e00649bdeeddf.
//
// Solidity: event NCPAddrChanged(address indexed ncp)
func (_StakingImp *StakingImpFilterer) ParseNCPAddrChanged(log types.Log) (*StakingImpNCPAddrChanged, error) {
	event := new(StakingImpNCPAddrChanged)
	if err := _StakingImp.contract.UnpackLog(event, "NCPAddrChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingImp contract.
type StakingImpOwnershipTransferredIterator struct {
	Event *StakingImpOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingImpOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpOwnershipTransferred)
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
		it.Event = new(StakingImpOwnershipTransferred)
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
func (it *StakingImpOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpOwnershipTransferred represents a OwnershipTransferred event raised by the StakingImp contract.
type StakingImpOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingImp *StakingImpFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingImpOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingImpOwnershipTransferredIterator{contract: _StakingImp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingImp *StakingImpFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingImpOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpOwnershipTransferred)
				if err := _StakingImp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingImp *StakingImpFilterer) ParseOwnershipTransferred(log types.Log) (*StakingImpOwnershipTransferred, error) {
	event := new(StakingImpOwnershipTransferred)
	if err := _StakingImp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpRevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the StakingImp contract.
type StakingImpRevokedIterator struct {
	Event *StakingImpRevoked // Event containing the contract specifics and raw log

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
func (it *StakingImpRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpRevoked)
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
		it.Event = new(StakingImpRevoked)
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
func (it *StakingImpRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpRevoked represents a Revoked event raised by the StakingImp contract.
type StakingImpRevoked struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0x713b90881ad62c4fa8ab6bd9197fa86481fc0c11b2edba60026514281b2dbac4.
//
// Solidity: event Revoked(address indexed owner, uint256 amount)
func (_StakingImp *StakingImpFilterer) FilterRevoked(opts *bind.FilterOpts, owner []common.Address) (*StakingImpRevokedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "Revoked", ownerRule)
	if err != nil {
		return nil, err
	}
	return &StakingImpRevokedIterator{contract: _StakingImp.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0x713b90881ad62c4fa8ab6bd9197fa86481fc0c11b2edba60026514281b2dbac4.
//
// Solidity: event Revoked(address indexed owner, uint256 amount)
func (_StakingImp *StakingImpFilterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *StakingImpRevoked, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "Revoked", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpRevoked)
				if err := _StakingImp.contract.UnpackLog(event, "Revoked", log); err != nil {
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

// ParseRevoked is a log parse operation binding the contract event 0x713b90881ad62c4fa8ab6bd9197fa86481fc0c11b2edba60026514281b2dbac4.
//
// Solidity: event Revoked(address indexed owner, uint256 amount)
func (_StakingImp *StakingImpFilterer) ParseRevoked(log types.Log) (*StakingImpRevoked, error) {
	event := new(StakingImpRevoked)
	if err := _StakingImp.contract.UnpackLog(event, "Revoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpSetRegistryIterator is returned from FilterSetRegistry and is used to iterate over the raw logs and unpacked data for SetRegistry events raised by the StakingImp contract.
type StakingImpSetRegistryIterator struct {
	Event *StakingImpSetRegistry // Event containing the contract specifics and raw log

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
func (it *StakingImpSetRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpSetRegistry)
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
		it.Event = new(StakingImpSetRegistry)
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
func (it *StakingImpSetRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpSetRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpSetRegistry represents a SetRegistry event raised by the StakingImp contract.
type StakingImpSetRegistry struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetRegistry is a free log retrieval operation binding the contract event 0x278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd.
//
// Solidity: event SetRegistry(address indexed addr)
func (_StakingImp *StakingImpFilterer) FilterSetRegistry(opts *bind.FilterOpts, addr []common.Address) (*StakingImpSetRegistryIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "SetRegistry", addrRule)
	if err != nil {
		return nil, err
	}
	return &StakingImpSetRegistryIterator{contract: _StakingImp.contract, event: "SetRegistry", logs: logs, sub: sub}, nil
}

// WatchSetRegistry is a free log subscription operation binding the contract event 0x278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd.
//
// Solidity: event SetRegistry(address indexed addr)
func (_StakingImp *StakingImpFilterer) WatchSetRegistry(opts *bind.WatchOpts, sink chan<- *StakingImpSetRegistry, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "SetRegistry", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpSetRegistry)
				if err := _StakingImp.contract.UnpackLog(event, "SetRegistry", log); err != nil {
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
func (_StakingImp *StakingImpFilterer) ParseSetRegistry(log types.Log) (*StakingImpSetRegistry, error) {
	event := new(StakingImpSetRegistry)
	if err := _StakingImp.contract.UnpackLog(event, "SetRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the StakingImp contract.
type StakingImpStakedIterator struct {
	Event *StakingImpStaked // Event containing the contract specifics and raw log

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
func (it *StakingImpStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpStaked)
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
		it.Event = new(StakingImpStaked)
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
func (it *StakingImpStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpStaked represents a Staked event raised by the StakingImp contract.
type StakingImpStaked struct {
	Payee     common.Address
	Amount    *big.Int
	Total     *big.Int
	Available *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xb4caaf29adda3eefee3ad552a8e85058589bf834c7466cae4ee58787f70589ed.
//
// Solidity: event Staked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) FilterStaked(opts *bind.FilterOpts, payee []common.Address) (*StakingImpStakedIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "Staked", payeeRule)
	if err != nil {
		return nil, err
	}
	return &StakingImpStakedIterator{contract: _StakingImp.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xb4caaf29adda3eefee3ad552a8e85058589bf834c7466cae4ee58787f70589ed.
//
// Solidity: event Staked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingImpStaked, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "Staked", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpStaked)
				if err := _StakingImp.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0xb4caaf29adda3eefee3ad552a8e85058589bf834c7466cae4ee58787f70589ed.
//
// Solidity: event Staked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) ParseStaked(log types.Log) (*StakingImpStaked, error) {
	event := new(StakingImpStaked)
	if err := _StakingImp.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpTransferLockedIterator is returned from FilterTransferLocked and is used to iterate over the raw logs and unpacked data for TransferLocked events raised by the StakingImp contract.
type StakingImpTransferLockedIterator struct {
	Event *StakingImpTransferLocked // Event containing the contract specifics and raw log

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
func (it *StakingImpTransferLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpTransferLocked)
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
		it.Event = new(StakingImpTransferLocked)
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
func (it *StakingImpTransferLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpTransferLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpTransferLocked represents a TransferLocked event raised by the StakingImp contract.
type StakingImpTransferLocked struct {
	Payee     common.Address
	Amount    *big.Int
	Total     *big.Int
	Available *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferLocked is a free log retrieval operation binding the contract event 0x2caed32a519a1fd89486d3ffe06202febb5ed51534d571dbab93058545a29e24.
//
// Solidity: event TransferLocked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) FilterTransferLocked(opts *bind.FilterOpts, payee []common.Address) (*StakingImpTransferLockedIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "TransferLocked", payeeRule)
	if err != nil {
		return nil, err
	}
	return &StakingImpTransferLockedIterator{contract: _StakingImp.contract, event: "TransferLocked", logs: logs, sub: sub}, nil
}

// WatchTransferLocked is a free log subscription operation binding the contract event 0x2caed32a519a1fd89486d3ffe06202febb5ed51534d571dbab93058545a29e24.
//
// Solidity: event TransferLocked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) WatchTransferLocked(opts *bind.WatchOpts, sink chan<- *StakingImpTransferLocked, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "TransferLocked", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpTransferLocked)
				if err := _StakingImp.contract.UnpackLog(event, "TransferLocked", log); err != nil {
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

// ParseTransferLocked is a log parse operation binding the contract event 0x2caed32a519a1fd89486d3ffe06202febb5ed51534d571dbab93058545a29e24.
//
// Solidity: event TransferLocked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) ParseTransferLocked(log types.Log) (*StakingImpTransferLocked, error) {
	event := new(StakingImpTransferLocked)
	if err := _StakingImp.contract.UnpackLog(event, "TransferLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the StakingImp contract.
type StakingImpUnlockedIterator struct {
	Event *StakingImpUnlocked // Event containing the contract specifics and raw log

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
func (it *StakingImpUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpUnlocked)
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
		it.Event = new(StakingImpUnlocked)
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
func (it *StakingImpUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpUnlocked represents a Unlocked event raised by the StakingImp contract.
type StakingImpUnlocked struct {
	Payee     common.Address
	Amount    *big.Int
	Total     *big.Int
	Available *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x5245d528087a96a64f4589a764f00061e4671eab90cb1e019b1a5b24b2e4c2a8.
//
// Solidity: event Unlocked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) FilterUnlocked(opts *bind.FilterOpts, payee []common.Address) (*StakingImpUnlockedIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "Unlocked", payeeRule)
	if err != nil {
		return nil, err
	}
	return &StakingImpUnlockedIterator{contract: _StakingImp.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x5245d528087a96a64f4589a764f00061e4671eab90cb1e019b1a5b24b2e4c2a8.
//
// Solidity: event Unlocked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *StakingImpUnlocked, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "Unlocked", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpUnlocked)
				if err := _StakingImp.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0x5245d528087a96a64f4589a764f00061e4671eab90cb1e019b1a5b24b2e4c2a8.
//
// Solidity: event Unlocked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) ParseUnlocked(log types.Log) (*StakingImpUnlocked, error) {
	event := new(StakingImpUnlocked)
	if err := _StakingImp.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the StakingImp contract.
type StakingImpUnstakedIterator struct {
	Event *StakingImpUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingImpUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpUnstaked)
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
		it.Event = new(StakingImpUnstaked)
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
func (it *StakingImpUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpUnstaked represents a Unstaked event raised by the StakingImp contract.
type StakingImpUnstaked struct {
	Payee     common.Address
	Amount    *big.Int
	Total     *big.Int
	Available *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x204fccf0d92ed8d48f204adb39b2e81e92bad0dedb93f5716ca9478cfb57de00.
//
// Solidity: event Unstaked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) FilterUnstaked(opts *bind.FilterOpts, payee []common.Address) (*StakingImpUnstakedIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "Unstaked", payeeRule)
	if err != nil {
		return nil, err
	}
	return &StakingImpUnstakedIterator{contract: _StakingImp.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x204fccf0d92ed8d48f204adb39b2e81e92bad0dedb93f5716ca9478cfb57de00.
//
// Solidity: event Unstaked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakingImpUnstaked, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "Unstaked", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpUnstaked)
				if err := _StakingImp.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x204fccf0d92ed8d48f204adb39b2e81e92bad0dedb93f5716ca9478cfb57de00.
//
// Solidity: event Unstaked(address indexed payee, uint256 amount, uint256 total, uint256 available)
func (_StakingImp *StakingImpFilterer) ParseUnstaked(log types.Log) (*StakingImpUnstaked, error) {
	event := new(StakingImpUnstaked)
	if err := _StakingImp.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingImpUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StakingImp contract.
type StakingImpUpgradedIterator struct {
	Event *StakingImpUpgraded // Event containing the contract specifics and raw log

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
func (it *StakingImpUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingImpUpgraded)
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
		it.Event = new(StakingImpUpgraded)
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
func (it *StakingImpUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingImpUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingImpUpgraded represents a Upgraded event raised by the StakingImp contract.
type StakingImpUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakingImp *StakingImpFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StakingImpUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakingImp.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StakingImpUpgradedIterator{contract: _StakingImp.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakingImp *StakingImpFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StakingImpUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakingImp.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingImpUpgraded)
				if err := _StakingImp.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_StakingImp *StakingImpFilterer) ParseUpgraded(log types.Log) (*StakingImpUpgraded, error) {
	event := new(StakingImpUpgraded)
	if err := _StakingImp.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
