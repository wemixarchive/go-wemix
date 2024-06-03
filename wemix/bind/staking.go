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
	Bin: "0x608060405234801561001057600080fd5b5060405161070238038061070283398101604081905261002f91610304565b80604051806020016040528060008152506100528282600061005a60201b60201c565b5050506103a8565b61006383610090565b6000825111806100705750805b1561008b5761008983836100d060201b61008b1760201c565b505b505050565b610099816100fc565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606100f583836040518060600160405280602781526020016106db602791396101ce565b9392505050565b61010f8161024760201b6100b71760201c565b6101765760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b806101ad7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61025660201b6100c61760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b6060600080856001600160a01b0316856040516101eb9190610359565b600060405180830381855af49150503d8060008114610226576040519150601f19603f3d011682016040523d82523d6000602084013e61022b565b606091505b50909250905061023d86838387610259565b9695505050505050565b6001600160a01b03163b151590565b90565b606083156102c85782516000036102c1576001600160a01b0385163b6102c15760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161016d565b50816102d2565b6102d283836102da565b949350505050565b8151156102ea5781518083602001fd5b8060405162461bcd60e51b815260040161016d9190610375565b60006020828403121561031657600080fd5b81516001600160a01b03811681146100f557600080fd5b60005b83811015610348578181015183820152602001610330565b838111156100895750506000910152565b6000825161036b81846020870161032d565b9190910192915050565b602081526000825180602084015261039481604085016020870161032d565b601f01601f19169190910160400192915050565b610324806103b76000396000f3fe6080604052600436106100225760003560e01c80635c60da1b1461003957610031565b366100315761002f61006a565b005b61002f61006a565b34801561004557600080fd5b5061004e61007c565b6040516001600160a01b03909116815260200160405180910390f35b61007a6100756100c9565b6100fc565b565b60006100866100c9565b905090565b60606100b083836040518060600160405280602781526020016102c860279139610120565b9392505050565b6001600160a01b03163b151590565b90565b60006100867f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b3660008037600080366000845af43d6000803e80801561011b573d6000f35b3d6000fd5b6060600080856001600160a01b03168560405161013d9190610278565b600060405180830381855af49150503d8060008114610178576040519150601f19603f3d011682016040523d82523d6000602084013e61017d565b606091505b509150915061018e86838387610198565b9695505050505050565b6060831561020c578251600003610205576001600160a01b0385163b6102055760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064015b60405180910390fd5b5081610216565b610216838361021e565b949350505050565b81511561022e5781518083602001fd5b8060405162461bcd60e51b81526004016101fc9190610294565b60005b8381101561026357818101518382015260200161024b565b83811115610272576000848401525b50505050565b6000825161028a818460208701610248565b9190910192915050565b60208152600082518060208401526102b3816040850160208701610248565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122078f5f98dd9a2b222be6b1f84538c386c7aed8e448103840ee2a4a12f219cc94d64736f6c634300080e0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

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
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000e7565b600054610100900460ff1615620000935760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811614620000e5576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6080516134c56200011f60003960008181610bcc01528181610c0c01528181610ca801528181610ce80152610d7701526134c56000f3fe6080604052600436106102345760003560e01c8063738fdd1a1161012e578063a91ee0dc116100ab578063d5c258901161006f578063d5c25890146106b3578063f1b8aa1d146106c8578063f2fde38b146106e8578063f3f6308014610708578063f69318221461072857600080fd5b8063a91ee0dc14610643578063b6549f7514610663578063bac4f33814610678578063c0d91eaf1461068b578063d0e30db0146106ab57600080fd5b8063884d97a7116100f2578063884d97a7146105a15780638da5cb5b146105c157806392a950b6146105df5780639667e76a146105ff5780639986e4b91461061f57600080fd5b8063738fdd1a146104e85780637bf46530146105205780637d77a0eb146105415780637eee288d146105615780637f2f4c061461058157600080fd5b806334125c84116101bc5780635935573611610180578063593557361461041a5780635a731cca146104505780636c78d2cf1461047457806370a082311461049d578063715018a6146104d357600080fd5b806334125c84146103905780633659cfe6146103b05780634bd1ed76146103d05780634f1ef286146103f257806352d1902d1461040557600080fd5b8063282d3fdf11610203578063282d3fdf146102ea5780632b0b9c5e1461030c5780632bc9ed021461032c5780632e1a7d4d1461034f5780632f40992e1461036f57600080fd5b80631285361514610243578063193468ac146102765780631e0cba0d146102ac57806325d998bb146102ca57600080fd5b3661023e57600080fd5b600080fd5b34801561024f57600080fd5b5061026361025e366004612e2d565b610748565b6040519081526020015b60405180910390f35b34801561028257600080fd5b50610263610291366004612e66565b6001600160a01b0316600090815260ce602052604090205490565b3480156102b857600080fd5b50610263665374616b696e6760c81b81565b3480156102d657600080fd5b506102636102e5366004612e66565b610775565b3480156102f657600080fd5b5061030a610305366004612e83565b6107a3565b005b34801561031857600080fd5b50610263610327366004612e66565b6107e9565b34801561033857600080fd5b50609b5460ff16604051901515815260200161026d565b34801561035b57600080fd5b5061030a61036a366004612eaf565b610845565b34801561037b57600080fd5b506102636914995dd85c99141bdbdb60b21b81565b34801561039c57600080fd5b506102636845636f73797374656d60b81b81565b3480156103bc57600080fd5b5061030a6103cb366004612e66565b610bc2565b3480156103dc57600080fd5b506102636a4d61696e74656e616e636560a81b81565b61030a610400366004612f0f565b610c9e565b34801561041157600080fd5b50610263610d6a565b34801561042657600080fd5b50610263610435366004612e66565b6001600160a01b031660009081526099602052604090205490565b34801561045c57600080fd5b506102636c14dd185ada5b99d4995dd85c99609a1b81565b34801561048057600080fd5b506102637111dbdd995c9b985b98d950dbdb9d1c9858dd60721b81565b3480156104a957600080fd5b506102636104b8366004612e66565b6001600160a01b031660009081526098602052604090205490565b3480156104df57600080fd5b5061030a610e1d565b3480156104f457600080fd5b50606554610508906001600160a01b031681565b6040516001600160a01b03909116815260200161026d565b34801561052c57600080fd5b5061026369456e7653746f7261676560b01b81565b34801561054d57600080fd5b5061026361055c366004612fb7565b610e31565b34801561056d57600080fd5b5061030a61057c366004612e83565b610ea5565b34801561058d57600080fd5b5061030a61059c366004612e66565b610ede565b3480156105ad57600080fd5b506102636105bc366004612e66565b610f86565b3480156105cd57600080fd5b506033546001600160a01b0316610508565b3480156105eb57600080fd5b5061030a6105fa366004612fee565b610f93565b34801561060b57600080fd5b5061030a61061a366004612eaf565b6112c9565b34801561062b57600080fd5b506102636c42616c6c6f7453746f7261676560981b81565b34801561064f57600080fd5b5061030a61065e366004612e66565b6114c4565b34801561066f57600080fd5b5061030a61156c565b61030a610686366004612e66565b611676565b34801561069757600080fd5b5061030a6106a6366004612f0f565b611a4b565b61030a611bfc565b3480156106bf57600080fd5b50609a54610263565b3480156106d457600080fd5b5060cf54610508906001600160a01b031681565b3480156106f457600080fd5b5061030a610703366004612e66565b611faa565b34801561071457600080fd5b5061030a610723366004612e66565b612020565b34801561073457600080fd5b5061030a610743366004612e83565b612040565b6001600160a01b03808316600090815260cd60209081526040808320938516835292905220545b92915050565b6001600160a01b038116600090815260996020908152604080832054609890925282205461076f9190613039565b336107ac612427565b6001600160a01b0316146107db5760405162461bcd60e51b81526004016107d290613050565b60405180910390fd5b6107e5828261244c565b5050565b6001600160a01b038116600090815260ce60209081526040808320546099909252822054811580610818575080155b15610827575060009392505050565b80610833836064613077565b61083d9190613096565b949350505050565b61084d6126c1565b609b5460ff16156108705760405162461bcd60e51b81526004016107d2906130b8565b600081116108905760405162461bcd60e51b81526004016107d2906130dc565b600061089a61271a565b6001600160a01b031663076cd77f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fb919061311d565b905060006109098383613136565b33600090815260ce602090815260408083205460999092529091205461092f9190613039565b106109425761093e3384612732565b5060015b61094b33610775565b83111561096a5760405162461bcd60e51b81526004016107d29061314e565b33600090815260986020526040902054610985908490613039565b3360009081526098602052604090205560cf546001600160a01b0316158015906109ac5750805b15610abe5760cf546040516000916001600160a01b03169085908381818185875af1925050503d80600081146109fe576040519150601f19603f3d011682016040523d82523d6000602084013e610a03565b606091505b5050905080610a545760405162461bcd60e51b815260206004820152601e60248201527f5472616e7366657220746f204e4350207374616b696e67206661696c6564000060448201526064016107d2565b60cf546040516306aa67f960e01b8152600481018690523360248201526001600160a01b03909116906306aa67f990604401600060405180830381600087803b158015610aa057600080fd5b505af1158015610ab4573d6000803e3d6000fd5b5050505050610b58565b604051600090339085908381818185875af1925050503d8060008114610b00576040519150601f19603f3d011682016040523d82523d6000602084013e610b05565b606091505b5050905080610b565760405162461bcd60e51b815260206004820152601960248201527f5472616e7366657220746f2073656e646572206661696c65640000000000000060448201526064016107d2565b505b336000818152609860205260409020547f204fccf0d92ed8d48f204adb39b2e81e92bad0dedb93f5716ca9478cfb57de00908590610b9584610775565b6040805193845260208401929092529082015260600160405180910390a25050610bbf6001606655565b50565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003610c0a5760405162461bcd60e51b81526004016107d2906131a2565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610c53600080516020613449833981519152546001600160a01b031690565b6001600160a01b031614610c795760405162461bcd60e51b81526004016107d2906131ee565b610c8281612882565b60408051600080825260208201909252610bbf9183919061288a565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003610ce65760405162461bcd60e51b81526004016107d2906131a2565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610d2f600080516020613449833981519152546001600160a01b031690565b6001600160a01b031614610d555760405162461bcd60e51b81526004016107d2906131ee565b610d5e82612882565b6107e58282600161288a565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e0a5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016107d2565b5060008051602061344983398151915290565b610e256129f5565b610e2f6000612a4f565b565b6001600160a01b0382166000908152609960205260408120541580610e5a575063ffffffff8216155b15610e675750600061076f565b609a546001600160a01b038416600090815260996020526040902054610e949063ffffffff851690613077565b610e9e9190613096565b9392505050565b33610eae612427565b6001600160a01b031614610ed45760405162461bcd60e51b81526004016107d290613050565b6107e58282612732565b610ee66129f5565b6001600160a01b038116610f3c5760405162461bcd60e51b815260206004820152601e60248201527f4e43505374616b696e6720697320746865207a65726f2061646472657373000060448201526064016107d2565b60cf80546001600160a01b0319166001600160a01b0383169081179091556040517ffd5754300dde6066eda4fabd23616c1d560a3360c85c0716c46e00649bdeeddf90600090a250565b600061076f826064610e31565b33610f9c612427565b6001600160a01b031614610fc25760405162461bcd60e51b81526004016107d290613050565b6000610fd7661390d4115e1a5d60ca1b612aa1565b9050610fe38484610ea5565b6001600160a01b038416600090815260986020526040902054611007908490613039565b6001600160a01b03851660009081526098602052604081209190915561102b612b0f565b6001600160a01b038116600090815260986020526040902054909150611052908590613136565b6001600160a01b03808316600090815260986020908152604080832094909455918816815260ce9091529081205461108a9085613039565b90506110968682610ea5565b6001600160a01b03861660009081526099602090815260408083205460ce909252909120548110156111495760405162461bcd60e51b815260206004820152605060248201527f7472616e73666572656442616c616e6365206d7573742062652067726561746560448201527f72207468616e206f7220657175616c20746f205f6c6f636b656455736572426160648201526f3630b731b2aa37a721a82a37ba30b61760811b608482015260a4016107d2565b6111538782610ea5565b6001600160a01b038716600090815260986020526040902054611177908290613039565b6001600160a01b0388811660008181526098602090815260408083209590955560ce9052839020549251638408bdb160e01b81526004810191909152602481018490526044810192909252851690638408bdb19083906064016000604051808303818588803b1580156111e957600080fd5b505af11580156111fd573d6000803e3d6000fd5b505060cf546001600160a01b031615925061124b915050576001600160a01b03808816600090815260ce6020908152604080832083905560cd825280832060cf549094168352929052908120555b6001600160a01b0387167f2caed32a519a1fd89486d3ffe06202febb5ed51534d571dbab93058545a29e246112808389613136565b6001600160a01b038a166000908152609860205260409020546112a28b610775565b6040805193845260208401929092529082015260600160405180910390a250505050505050565b6112d1612427565b604051636f1e853360e01b81523360048201526001600160a01b039190911690636f1e853390602401602060405180830381865afa158015611317573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061133b919061323a565b6113575760405162461bcd60e51b81526004016107d290613050565b611361338261244c565b60cf546001600160a01b031615610bbf5760cf5460405163bbc2611360e01b81523360048201526000916001600160a01b03169063bbc2611390602401602060405180830381865afa1580156113bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113df919061311d565b60cf54604051631069f3b560e01b8152600481018390523360248201529192506001600160a01b031690631069f3b59060440160a060405180830381865afa15801561142f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611453919061325c565b60800151156107e55760cf546040516301008e9960e61b8152600481018490523360248201526001600160a01b0390911690634023a64090604401600060405180830381600087803b1580156114a857600080fd5b505af11580156114bc573d6000803e3d6000fd5b505050505050565b6114cc6129f5565b6001600160a01b0381166115225760405162461bcd60e51b815260206004820152601a60248201527f416464726573732073686f756c64206265206e6f6e2d7a65726f00000000000060448201526064016107d2565b606580546001600160a01b0319166001600160a01b0383169081179091556040517f278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd90600090a250565b6115746129f5565b609b5460ff16156115975760405162461bcd60e51b81526004016107d2906130b8565b60006115ab6033546001600160a01b031690565b905047806115e95760405162461bcd60e51b815260206004820152600b60248201526a062616c616e6365203d20360ac1b60448201526064016107d2565b6040516001600160a01b0383169082156108fc029083906000818181858888f1935050505015801561161f573d6000803e3d6000fd5b50609b805460ff191660011790556040516001600160a01b038316907f713b90881ad62c4fa8ab6bd9197fa86481fc0c11b2edba60026514281b2dbac49061166a9084815260200190565b60405180910390a25050565b61167e6126c1565b609b5460ff16156116a15760405162461bcd60e51b81526004016107d2906130b8565b60cf546001600160a01b031633146116cb5760405162461bcd60e51b81526004016107d2906132cc565b600034116116eb5760405162461bcd60e51b81526004016107d29061331b565b6116f3612427565b60405163288c314960e21b81526001600160a01b038381166004830152919091169063a230c52490602401602060405180830381865afa15801561173b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061175f919061323a565b6117a45760405162461bcd60e51b81526020600482015260166024820152752721a81039b437bab63210313290309036b2b6b132b960511b60448201526064016107d2565b6001600160a01b03811660009081526098602052604090205434906117ca908290613136565b6001600160a01b0383166000908152609860205260408120919091556117ee61271a565b6001600160a01b031663076cd77f6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561182b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061184f919061311d565b9050600061185b61271a565b6001600160a01b031663737c59b86040518163ffffffff1660e01b8152600401602060405180830381865afa158015611898573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118bc919061311d565b6001600160a01b038516600090815260996020526040902054909150821180159061190b57506001600160a01b0384166000908152609960205260409020548190611908908590613136565b11155b6119575760405162461bcd60e51b815260206004820152601f60248201527f757365722073686f756c6420626520696e207374616b696e672072616e67650060448201526064016107d2565b611961848461244c565b6001600160a01b038416600090815260cd60209081526040808320338452909152902054611990908490613136565b6001600160a01b038516600081815260cd6020908152604080832033845282528083209490945591815260ce90915220546119cc908490613136565b6001600160a01b038516600081815260ce6020908152604080832085905560cd825280832033808552908352928190205481518981529283019590955281019390935290917f74cfc20f0e6d14384c3a60820d3e814f6979d009cdbb43db27fa56fe475172fd9060600160405180910390a3505050610bbf6001606655565b600054610100900460ff1615808015611a6b5750600054600160ff909116105b80611a855750303b158015611a85575060005460ff166001145b611ae85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016107d2565b6000805460ff191660011790558015611b0b576000805461ff0019166101001790555b6000609a55609b805460ff19169055611b22612b26565b611b2a612b55565b611b33836114c4565b815115611bb157600080600080602086019150855182611b539190613136565b90505b80821015611bac5781519350611b6d602083613136565b9150808210611b7b57600080fd5b81519250611b8a602083613136565b6001600160a01b03851660009081526098602052604090208490559150611b56565b505050505b8015611bf7576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b611c046126c1565b609b5460ff1615611c275760405162461bcd60e51b81526004016107d2906130b8565b60003411611c475760405162461bcd60e51b81526004016107d29061331b565b33600090815260986020526040902054611c62903490613136565b33600090815260986020526040902055611c7a612427565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c52490602401602060405180830381865afa158015611cc0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ce4919061323a565b15611f45576000611cf361271a565b6001600160a01b031663076cd77f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d54919061311d565b3360009081526099602052604090205490915081118015611d96575033600090815260996020526040902054611d8a9082613039565b611d9333610775565b10155b15611f435760cf546001600160a01b031615611f1f5760cf5460405163bbc2611360e01b81523360048201526000916001600160a01b03169063bbc2611390602401602060405180830381865afa158015611df5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e19919061311d565b60cf54604051631069f3b560e01b8152600481018390523360248201529192506001600160a01b031690631069f3b59060440160a060405180830381865afa158015611e69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e8d919061325c565b6080015115611f1d5760cf54336000908152609960205260409020546001600160a01b0390911690634023a64090611ec59085613039565b6040516001600160e01b031960e084901b1681526004810191909152336024820152604401600060405180830381600087803b158015611f0457600080fd5b505af1158015611f18573d6000803e3d6000fd5b505050505b505b33600081815260996020526040902054611f439190611f3e9084613039565b61244c565b505b336000818152609860205260409020547fb4caaf29adda3eefee3ad552a8e85058589bf834c7466cae4ee58787f70589ed903490611f8284610775565b6040805193845260208401929092529082015260600160405180910390a2610e2f6001606655565b611fb26129f5565b6001600160a01b0381166120175760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016107d2565b610bbf81612a4f565b6120286129f5565b6001600160a01b03811615610bbf57610c8281612882565b6120486126c1565b609b5460ff161561206b5760405162461bcd60e51b81526004016107d2906130b8565b60cf546001600160a01b031633146120955760405162461bcd60e51b81526004016107d2906132cc565b600081116120b55760405162461bcd60e51b81526004016107d2906130dc565b6120bd612427565b60405163288c314960e21b81526001600160a01b038481166004830152919091169063a230c52490602401602060405180830381865afa158015612105573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612129919061323a565b61216e5760405162461bcd60e51b81526020600482015260166024820152752721a81039b437bab63210313290309036b2b6b132b960511b60448201526064016107d2565b80600061217961271a565b6001600160a01b031663076cd77f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121da919061311d565b6001600160a01b038516600090815260cd6020908152604080832033845290915290205490915082118015906122395750808261222c866001600160a01b031660009081526099602052604090205490565b6122369190613039565b10155b6122555760405162461bcd60e51b81526004016107d29061314e565b61225f8483612732565b6001600160a01b038416600090815260986020526040902054612283908390613039565b6001600160a01b03851660009081526098602090815260408083209390935560cd8152828220338352905220546122bb908390613039565b6001600160a01b038516600081815260cd6020908152604080832033845282528083209490945591815260ce90915220546122f7908390613039565b6001600160a01b03858116600090815260ce60205260408082209390935560cf54925190929091169084908381818185875af1925050503d806000811461235a576040519150601f19603f3d011682016040523d82523d6000602084013e61235f565b606091505b50509050806123b05760405162461bcd60e51b815260206004820152601e60248201527f5472616e7366657220746f204e4350207374616b696e67206661696c6564000060448201526064016107d2565b6001600160a01b038516600081815260ce602090815260408083205460cd835281842033808652908452938290205482518981529384019190915282820152517f03d2bb70c6ccc49d68a465a06edffb976961cf8930888658ca2339fa62b8bda29181900360600190a35050506107e56001606655565b60006124477111dbdd995c9b985b98d950dbdb9d1c9858dd60721b612aa1565b905090565b80600003612458575050565b6001600160a01b0382166000908152609860205260409020548111156124d95760405162461bcd60e51b815260206004820152603060248201527f4c6f636b20616d6f756e742073686f756c6420626520657175616c206f72206c60448201526f657373207468616e2062616c616e636560801b60648201526084016107d2565b806124e383610775565b10156125415760405162461bcd60e51b815260206004820152602760248201527f496e73756666696369656e742062616c616e636520746861742063616e206265604482015266081b1bd8dad95960ca1b60648201526084016107d2565b600061254b61271a565b6001600160a01b031663737c59b86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612588573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125ac919061311d565b6001600160a01b0384166000908152609960205260409020549091506125d3908390613136565b6001600160a01b03841660009081526099602052604090208190558110156126475760405162461bcd60e51b815260206004820152602160248201527f4c6f636b65642062616c616e6365206973206c6172676572207468616e206d616044820152600f60fb1b60648201526084016107d2565b81609a546126559190613136565b609a556001600160a01b0383166000818152609860205260409020547f44cebfefa4561bee5b61d675ccfd8dc9969fff9cc15e7a4eccccd62af94f9c1190849061269e87610775565b6040805193845260208401929092529082015260600160405180910390a2505050565b6002606654036127135760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016107d2565b6002606655565b600061244769456e7653746f7261676560b01b612aa1565b8060000361273e575050565b6001600160a01b0382166000908152609960205260409020548111156127cc5760405162461bcd60e51b815260206004820152603960248201527f556e6c6f636b20616d6f756e742073686f756c6420626520657175616c206f7260448201527f206c657373207468616e2062616c616e6365206c6f636b65640000000000000060648201526084016107d2565b6001600160a01b0382166000908152609960205260409020546127f0908290613039565b6001600160a01b038316600090815260996020526040902055609a54612817908290613039565b609a556001600160a01b0382166000818152609860205260409020547f5245d528087a96a64f4589a764f00061e4671eab90cb1e019b1a5b24b2e4c2a890839061286086610775565b6040805193845260208401929092529082015260600161166a565b6001606655565b610bbf6129f5565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156128bd57611bf783612b84565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612917575060408051601f3d908101601f191682019092526129149181019061311d565b60015b61297a5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b60648201526084016107d2565b60008051602061344983398151915281146129e95760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b60648201526084016107d2565b50611bf7838383612c20565b6033546001600160a01b03163314610e2f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016107d2565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b606554604051630d2020dd60e01b8152600481018390526000916001600160a01b031690630d2020dd90602401602060405180830381865afa158015612aeb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076f9190613365565b60006124476845636f73797374656d60b81b612aa1565b600054610100900460ff16612b4d5760405162461bcd60e51b81526004016107d290613382565b610e2f612c4b565b600054610100900460ff16612b7c5760405162461bcd60e51b81526004016107d290613382565b610e2f612c72565b6001600160a01b0381163b612bf15760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016107d2565b60008051602061344983398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b612c2983612ca2565b600082511180612c365750805b15611bf757612c458383612ce2565b50505050565b600054610100900460ff1661287b5760405162461bcd60e51b81526004016107d290613382565b600054610100900460ff16612c995760405162461bcd60e51b81526004016107d290613382565b610e2f33612a4f565b612cab81612b84565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060610e9e8383604051806060016040528060278152602001613469602791396060600080856001600160a01b031685604051612d1f91906133f9565b600060405180830381855af49150503d8060008114612d5a576040519150601f19603f3d011682016040523d82523d6000602084013e612d5f565b606091505b5091509150612d7086838387612d7a565b9695505050505050565b60608315612de9578251600003612de2576001600160a01b0385163b612de25760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016107d2565b508161083d565b61083d8383815115612dfe5781518083602001fd5b8060405162461bcd60e51b81526004016107d29190613415565b6001600160a01b0381168114610bbf57600080fd5b60008060408385031215612e4057600080fd5b8235612e4b81612e18565b91506020830135612e5b81612e18565b809150509250929050565b600060208284031215612e7857600080fd5b8135610e9e81612e18565b60008060408385031215612e9657600080fd5b8235612ea181612e18565b946020939093013593505050565b600060208284031215612ec157600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715612f0757612f07612ec8565b604052919050565b60008060408385031215612f2257600080fd5b8235612f2d81612e18565b915060208381013567ffffffffffffffff80821115612f4b57600080fd5b818601915086601f830112612f5f57600080fd5b813581811115612f7157612f71612ec8565b612f83601f8201601f19168501612ede565b91508082528784828501011115612f9957600080fd5b80848401858401376000848284010152508093505050509250929050565b60008060408385031215612fca57600080fd5b8235612fd581612e18565b9150602083013563ffffffff81168114612e5b57600080fd5b60008060006060848603121561300357600080fd5b833561300e81612e18565b95602085013595506040909401359392505050565b634e487b7160e01b600052601160045260246000fd5b60008282101561304b5761304b613023565b500390565b6020808252600d908201526c2737902832b936b4b9b9b4b7b760991b604082015260600190565b600081600019048311821515161561309157613091613023565b500290565b6000826130b357634e487b7160e01b600052601260045260246000fd5b500490565b6020808252600a9082015269125cc81c995d9bdad95960b21b604082015260600190565b60208082526021908201527f416d6f756e742073686f756c6420626520626967676572207468616e207a65726040820152606f60f81b606082015260800190565b60006020828403121561312f57600080fd5b5051919050565b6000821982111561314957613149613023565b500190565b60208082526034908201527f576974686472617720616d6f756e742073686f756c6420626520657175616c206040820152736f72206c657373207468616e2062616c616e636560601b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b60006020828403121561324c57600080fd5b81518015158114610e9e57600080fd5b600060a0828403121561326e57600080fd5b60405160a0810181811067ffffffffffffffff8211171561329157613291612ec8565b806040525082518152602083015160208201526040830151604082015260608301516060820152608083015160808201528091505092915050565b6020808252602f908201527f4f6e6c79204e43505374616b696e6720636f6e74726163742063616e2063616c60408201526e36103a3434b990333ab731ba34b7b760891b606082015260800190565b6020808252602a908201527f4465706f73697420616d6f756e742073686f756c642062652067726561746572604082015269207468616e207a65726f60b01b606082015260800190565b60006020828403121561337757600080fd5b8151610e9e81612e18565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60005b838110156133e85781810151838201526020016133d0565b83811115612c455750506000910152565b6000825161340b8184602087016133cd565b9190910192915050565b60208152600082518060208401526134348160408501602087016133cd565b601f01601f1916919091016040019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220cbfd73c1953b63341112520dc0b4ec15de379a6ca84968b465bda2cf2a7ed96d64736f6c634300080e0033",
}

// StakingImpABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingImpMetaData.ABI instead.
var StakingImpABI = StakingImpMetaData.ABI

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
