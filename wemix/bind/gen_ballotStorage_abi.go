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

// BallotStorageMetaData contains all meta data concerning the BallotStorage contract.
var BallotStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_imp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"5c60da1b": "implementation()",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161077238038061077283398101604081905261002f91610326565b604080516020810190915260008152819061006b60017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd61034f565b60008051602061072b8339815191521461008757610087610374565b6100938282600061009b565b505050610405565b6100a4836100d1565b6000825111806100b15750805b156100cc576100ca838361011160201b61008b1760201c565b505b505050565b6100da8161013d565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060610136838360405180606001604052806027815260200161074b602791396101fd565b9392505050565b610150816102db60201b6100b71760201c565b6101b75760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b806101dc60008051602061072b83398151915260001b6102ea60201b6100c61760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606001600160a01b0384163b6102655760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016101ae565b600080856001600160a01b03168560405161028091906103b6565b600060405180830381855af49150503d80600081146102bb576040519150601f19603f3d011682016040523d82523d6000602084013e6102c0565b606091505b5090925090506102d18282866102ed565b9695505050505050565b6001600160a01b03163b151590565b90565b606083156102fc575081610136565b82511561030c5782518084602001fd5b8160405162461bcd60e51b81526004016101ae91906103d2565b60006020828403121561033857600080fd5b81516001600160a01b038116811461013657600080fd5b60008282101561036f57634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052600160045260246000fd5b60005b838110156103a557818101518382015260200161038d565b838111156100ca5750506000910152565b600082516103c881846020870161038a565b9190910192915050565b60208152600082518060208401526103f181604085016020870161038a565b601f01601f19169190910160400192915050565b610317806104146000396000f3fe6080604052600436106100225760003560e01c80635c60da1b1461003957610031565b366100315761002f61006a565b005b61002f61006a565b34801561004557600080fd5b5061004e61007c565b6040516001600160a01b03909116815260200160405180910390f35b61007a6100756100c9565b6100fc565b565b60006100866100c9565b905090565b60606100b083836040518060600160405280602781526020016102bb60279139610120565b9392505050565b6001600160a01b03163b151590565b90565b60006100867f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b3660008037600080366000845af43d6000803e80801561011b573d6000f35b3d6000fd5b60606001600160a01b0384163b61018d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084015b60405180910390fd5b600080856001600160a01b0316856040516101a8919061026b565b600060405180830381855af49150503d80600081146101e3576040519150601f19603f3d011682016040523d82523d6000602084013e6101e8565b606091505b50915091506101f8828286610202565b9695505050505050565b606083156102115750816100b0565b8251156102215782518084602001fd5b8160405162461bcd60e51b81526004016101849190610287565b60005b8381101561025657818101518382015260200161023e565b83811115610265576000848401525b50505050565b6000825161027d81846020870161023b565b9190910192915050565b60208152600082518060208401526102a681604085016020870161023b565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220ee902ae1ebb19baf0a3fcf922eb15a905dba5a97ac0a363ad62a66403d3450e564736f6c634300080e0033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
}

// BallotStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use BallotStorageMetaData.ABI instead.
var BallotStorageABI = BallotStorageMetaData.ABI

// Deprecated: Use BallotStorageMetaData.Sigs instead.
// BallotStorageFuncSigs maps the 4-byte function signature to its string representation.
var BallotStorageFuncSigs = BallotStorageMetaData.Sigs

// BallotStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BallotStorageMetaData.Bin instead.
var BallotStorageBin = BallotStorageMetaData.Bin

// DeployBallotStorage deploys a new Ethereum contract, binding an instance of BallotStorage to it.
func DeployBallotStorage(auth *bind.TransactOpts, backend bind.ContractBackend, _imp common.Address) (common.Address, *types.Transaction, *BallotStorage, error) {
	parsed, err := BallotStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BallotStorageBin), backend, _imp)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BallotStorage{BallotStorageCaller: BallotStorageCaller{contract: contract}, BallotStorageTransactor: BallotStorageTransactor{contract: contract}, BallotStorageFilterer: BallotStorageFilterer{contract: contract}}, nil
}

// BallotStorage is an auto generated Go binding around an Ethereum contract.
type BallotStorage struct {
	BallotStorageCaller     // Read-only binding to the contract
	BallotStorageTransactor // Write-only binding to the contract
	BallotStorageFilterer   // Log filterer for contract events
}

// BallotStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type BallotStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BallotStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BallotStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BallotStorageSession struct {
	Contract     *BallotStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BallotStorageCallerSession struct {
	Contract *BallotStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BallotStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BallotStorageTransactorSession struct {
	Contract     *BallotStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BallotStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type BallotStorageRaw struct {
	Contract *BallotStorage // Generic contract binding to access the raw methods on
}

// BallotStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BallotStorageCallerRaw struct {
	Contract *BallotStorageCaller // Generic read-only contract binding to access the raw methods on
}

// BallotStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BallotStorageTransactorRaw struct {
	Contract *BallotStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBallotStorage creates a new instance of BallotStorage, bound to a specific deployed contract.
func NewBallotStorage(address common.Address, backend bind.ContractBackend) (*BallotStorage, error) {
	contract, err := bindBallotStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BallotStorage{BallotStorageCaller: BallotStorageCaller{contract: contract}, BallotStorageTransactor: BallotStorageTransactor{contract: contract}, BallotStorageFilterer: BallotStorageFilterer{contract: contract}}, nil
}

// NewBallotStorageCaller creates a new read-only instance of BallotStorage, bound to a specific deployed contract.
func NewBallotStorageCaller(address common.Address, caller bind.ContractCaller) (*BallotStorageCaller, error) {
	contract, err := bindBallotStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BallotStorageCaller{contract: contract}, nil
}

// NewBallotStorageTransactor creates a new write-only instance of BallotStorage, bound to a specific deployed contract.
func NewBallotStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*BallotStorageTransactor, error) {
	contract, err := bindBallotStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BallotStorageTransactor{contract: contract}, nil
}

// NewBallotStorageFilterer creates a new log filterer instance of BallotStorage, bound to a specific deployed contract.
func NewBallotStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*BallotStorageFilterer, error) {
	contract, err := bindBallotStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BallotStorageFilterer{contract: contract}, nil
}

// bindBallotStorage binds a generic wrapper to an already deployed contract.
func bindBallotStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BallotStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BallotStorage *BallotStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BallotStorage.Contract.BallotStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BallotStorage *BallotStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotStorage.Contract.BallotStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BallotStorage *BallotStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BallotStorage.Contract.BallotStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BallotStorage *BallotStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BallotStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BallotStorage *BallotStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BallotStorage *BallotStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BallotStorage.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_BallotStorage *BallotStorageCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BallotStorage.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_BallotStorage *BallotStorageSession) Implementation() (common.Address, error) {
	return _BallotStorage.Contract.Implementation(&_BallotStorage.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_BallotStorage *BallotStorageCallerSession) Implementation() (common.Address, error) {
	return _BallotStorage.Contract.Implementation(&_BallotStorage.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BallotStorage *BallotStorageTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BallotStorage.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BallotStorage *BallotStorageSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BallotStorage.Contract.Fallback(&_BallotStorage.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BallotStorage *BallotStorageTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BallotStorage.Contract.Fallback(&_BallotStorage.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BallotStorage *BallotStorageTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotStorage.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BallotStorage *BallotStorageSession) Receive() (*types.Transaction, error) {
	return _BallotStorage.Contract.Receive(&_BallotStorage.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BallotStorage *BallotStorageTransactorSession) Receive() (*types.Transaction, error) {
	return _BallotStorage.Contract.Receive(&_BallotStorage.TransactOpts)
}

// BallotStorageAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the BallotStorage contract.
type BallotStorageAdminChangedIterator struct {
	Event *BallotStorageAdminChanged // Event containing the contract specifics and raw log

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
func (it *BallotStorageAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageAdminChanged)
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
		it.Event = new(BallotStorageAdminChanged)
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
func (it *BallotStorageAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageAdminChanged represents a AdminChanged event raised by the BallotStorage contract.
type BallotStorageAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BallotStorage *BallotStorageFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BallotStorageAdminChangedIterator, error) {

	logs, sub, err := _BallotStorage.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BallotStorageAdminChangedIterator{contract: _BallotStorage.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BallotStorage *BallotStorageFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BallotStorageAdminChanged) (event.Subscription, error) {

	logs, sub, err := _BallotStorage.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageAdminChanged)
				if err := _BallotStorage.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_BallotStorage *BallotStorageFilterer) ParseAdminChanged(log types.Log) (*BallotStorageAdminChanged, error) {
	event := new(BallotStorageAdminChanged)
	if err := _BallotStorage.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the BallotStorage contract.
type BallotStorageBeaconUpgradedIterator struct {
	Event *BallotStorageBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BallotStorageBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageBeaconUpgraded)
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
		it.Event = new(BallotStorageBeaconUpgraded)
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
func (it *BallotStorageBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageBeaconUpgraded represents a BeaconUpgraded event raised by the BallotStorage contract.
type BallotStorageBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BallotStorage *BallotStorageFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BallotStorageBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BallotStorage.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BallotStorageBeaconUpgradedIterator{contract: _BallotStorage.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BallotStorage *BallotStorageFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BallotStorageBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BallotStorage.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageBeaconUpgraded)
				if err := _BallotStorage.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_BallotStorage *BallotStorageFilterer) ParseBeaconUpgraded(log types.Log) (*BallotStorageBeaconUpgraded, error) {
	event := new(BallotStorageBeaconUpgraded)
	if err := _BallotStorage.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BallotStorage contract.
type BallotStorageUpgradedIterator struct {
	Event *BallotStorageUpgraded // Event containing the contract specifics and raw log

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
func (it *BallotStorageUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageUpgraded)
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
		it.Event = new(BallotStorageUpgraded)
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
func (it *BallotStorageUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageUpgraded represents a Upgraded event raised by the BallotStorage contract.
type BallotStorageUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BallotStorage *BallotStorageFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BallotStorageUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BallotStorage.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BallotStorageUpgradedIterator{contract: _BallotStorage.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BallotStorage *BallotStorageFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BallotStorageUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BallotStorage.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageUpgraded)
				if err := _BallotStorage.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BallotStorage *BallotStorageFilterer) ParseUpgraded(log types.Log) (*BallotStorageUpgraded, error) {
	event := new(BallotStorageUpgraded)
	if err := _BallotStorage.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageImpMetaData contains all meta data concerning the BallotStorageImp contract.
var BallotStorageImpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ballotId\",\"type\":\"uint256\"}],\"name\":\"BallotCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ballotId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ballotType\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"BallotCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ballotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"BallotFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ballotId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"BallotStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ballotId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updatedBy\",\"type\":\"address\"}],\"name\":\"BallotUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"}],\"name\":\"SetPrevBallotStorage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"voteid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ballotId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decision\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ECOSYSTEM_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ENV_STORAGE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAINTENANCE_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_POOL_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_REWARD_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ballotId\",\"type\":\"uint256\"}],\"name\":\"cancelBallot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ballotType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newGovernanceAddress\",\"type\":\"address\"}],\"name\":\"createBallotForAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashing\",\"type\":\"uint256\"}],\"name\":\"createBallotForExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ballotType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oldStakerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newStakerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newVoterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newRewardAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_newNodeName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_newNodeId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_newNodeIp\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_newNodePort\",\"type\":\"uint256\"}],\"name\":\"createBallotForMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ballotType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_envVariableName\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_envVariableType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_envVariableValue\",\"type\":\"bytes\"}],\"name\":\"createBallotForVariable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ballotId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_decision\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"}],\"name\":\"createVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ballotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ballotState\",\"type\":\"uint256\"}],\"name\":\"finalizeBallot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getBallotAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newGovernanceAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getBallotBasic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ballotType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"totalVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"powerOfAccepts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"powerOfRejects\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFinalized\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBallotCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getBallotForExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashing\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getBallotMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"oldStakerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newStakerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newVoterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newRewardAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"newNodeName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newNodeId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newNodeIp\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"newNodePort\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getBallotPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getBallotState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ballotType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFinalized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getBallotVariable\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"envVariableName\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"envVariableType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"envVariableValue\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getBallotVotingInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"powerOfAccepts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"powerOfRejects\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxVotingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinVotingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPreviousBallotStorage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteId\",\"type\":\"uint256\"}],\"name\":\"getVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"voteId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ballotId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"decision\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint56\",\"name\":\"_ballotId\",\"type\":\"uint56\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"hasAlreadyVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reg\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setPreviousBallotStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ballotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"}],\"name\":\"startBallot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ballotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"updateBallotDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ballotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockAmount\",\"type\":\"uint256\"}],\"name\":\"updateBallotMemberLockAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ballotId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_memo\",\"type\":\"bytes\"}],\"name\":\"updateBallotMemo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImp\",\"type\":\"address\"}],\"name\":\"upgradeBallotStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9986e4b9": "BALLOT_STORAGE_NAME()",
		"34125c84": "ECOSYSTEM_NAME()",
		"7bf46530": "ENV_STORAGE_NAME()",
		"6c78d2cf": "GOV_NAME()",
		"4bd1ed76": "MAINTENANCE_NAME()",
		"2f40992e": "REWARD_POOL_NAME()",
		"1e0cba0d": "STAKING_NAME()",
		"5a731cca": "STAKING_REWARD_NAME()",
		"155ca224": "cancelBallot(uint256)",
		"0a3a63fe": "createBallotForAddress(uint256,uint256,uint256,address,address)",
		"22640859": "createBallotForExit(uint256,uint256,uint256)",
		"daacbb95": "createBallotForMember(uint256,uint256,uint256,address,address,address,address,address,bytes,bytes,bytes,uint256)",
		"4a57823e": "createBallotForVariable(uint256,uint256,uint256,address,bytes32,uint256,bytes)",
		"96462b9c": "createVote(uint256,uint256,address,uint256,uint256)",
		"a91e59ba": "finalizeBallot(uint256,uint256)",
		"7efa9ae3": "getBallotAddress(uint256)",
		"02b385fb": "getBallotBasic(uint256)",
		"b4741495": "getBallotCount()",
		"8c7be692": "getBallotForExit(uint256)",
		"73df4e01": "getBallotMember(uint256)",
		"09970688": "getBallotPeriod(uint256)",
		"688ca5b2": "getBallotState(uint256)",
		"1d940da2": "getBallotVariable(uint256)",
		"56ba988e": "getBallotVotingInfo(uint256)",
		"ce04b9d4": "getMaxVotingDuration()",
		"1c150171": "getMinVotingDuration()",
		"b23c676c": "getPreviousBallotStorage()",
		"557ed1ba": "getTime()",
		"5a55c1f0": "getVote(uint256)",
		"f680e555": "hasAlreadyVoted(uint56,address)",
		"c4d66de8": "initialize(address)",
		"6c57f5a9": "isDisabled()",
		"8da5cb5b": "owner()",
		"52d1902d": "proxiableUUID()",
		"738fdd1a": "reg()",
		"715018a6": "renounceOwnership()",
		"2a74f38c": "setPreviousBallotStorage(address)",
		"a91ee0dc": "setRegistry(address)",
		"c0b6f186": "startBallot(uint256,uint256,uint256)",
		"f2fde38b": "transferOwnership(address)",
		"656bc633": "updateBallotDuration(uint256,uint256)",
		"72d0ec92": "updateBallotMemberLockAmount(uint256,uint256)",
		"bce0dbc1": "updateBallotMemo(uint256,bytes)",
		"f51a88e2": "upgradeBallotStorage(address)",
		"3659cfe6": "upgradeTo(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000e8565b600054610100900460ff1615620000935760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015620000e6576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6080516144aa6200012060003960008181611147015281816111870152818161132d0152818161136d015261140001526144aa6000f3fe60806040526004361061027d5760003560e01c80636c78d2cf1161014f578063a91e59ba116100c1578063c4d66de81161007a578063c4d66de814610872578063ce04b9d414610892578063daacbb95146108a7578063f2fde38b146108c7578063f51a88e2146108e7578063f680e5551461090757600080fd5b8063a91e59ba146107bf578063a91ee0dc146107df578063b23c676c146107ff578063b47414951461081d578063bce0dbc114610832578063c0b6f1861461085257600080fd5b80637bf46530116101135780637bf46530146106b95780637efa9ae3146106da5780638c7be692146107135780638da5cb5b1461075d57806396462b9c1461077b5780639986e4b91461079b57600080fd5b80636c78d2cf146105ee578063715018a61461061757806372d0ec921461062c578063738fdd1a1461064c57806373df4e011461068457600080fd5b80633659cfe6116101f357806356ba988e116101ac57806356ba988e146104cd5780635a55c1f0146104ed5780635a731cca14610548578063656bc6331461056c578063688ca5b21461058c5780636c57f5a9146105c957600080fd5b80633659cfe6146104305780634a57823e146104505780634bd1ed76146104705780634f1ef2861461049257806352d1902d146104a5578063557ed1ba146104ba57600080fd5b80631d940da2116102455780631d940da2146103625780631e0cba0d1461039157806322640859146103af5780632a74f38c146103cf5780632f40992e146103ef57806334125c841461041057600080fd5b806302b385fb1461028257806309970688146102c25780630a3a63fe146102fd578063155ca2241461032b5780631c1501711461034d575b600080fd5b34801561028e57600080fd5b506102a261029d366004613acc565b610958565b6040516102b99b9a99989796959493929190613b3d565b60405180910390f35b3480156102ce57600080fd5b506102e26102dd366004613acc565b610b23565b604080519384526020840192909252908201526060016102b9565b34801561030957600080fd5b5061031d610318366004613bd1565b610c83565b6040519081526020016102b9565b34801561033757600080fd5b5061034b610346366004613acc565b610d9b565b005b34801561035957600080fd5b5061031d610efc565b34801561036e57600080fd5b5061038261037d366004613acc565b610f6c565b6040516102b993929190613c27565b34801561039d57600080fd5b5061031d665374616b696e6760c81b81565b3480156103bb57600080fd5b5061034b6103ca366004613c4f565b611024565b3480156103db57600080fd5b5061034b6103ea366004613c7b565b6110a3565b3480156103fb57600080fd5b5061031d6914995dd85c99141bdbdb60b21b81565b34801561041c57600080fd5b5061031d6845636f73797374656d60b81b81565b34801561043c57600080fd5b5061034b61044b366004613c7b565b61113d565b34801561045c57600080fd5b5061031d61046b366004613d3b565b61121c565b34801561047c57600080fd5b5061031d6a4d61696e74656e616e636560a81b81565b61034b6104a0366004613dbb565b611323565b3480156104b157600080fd5b5061031d6113f3565b3480156104c657600080fd5b504261031d565b3480156104d957600080fd5b506102e26104e8366004613acc565b6114a6565b3480156104f957600080fd5b5061050d610508366004613acc565b6115f7565b6040805196875260208701959095526001600160a01b03909316938501939093526060840152608083019190915260a082015260c0016102b9565b34801561055457600080fd5b5061031d6c14dd185ada5b99d4995dd85c99609a1b81565b34801561057857600080fd5b5061034b610587366004613e0b565b6116c5565b34801561059857600080fd5b506105ac6105a7366004613acc565b6118d8565b6040805193845260208401929092521515908201526060016102b9565b3480156105d557600080fd5b506105de611a2e565b60405190151581526020016102b9565b3480156105fa57600080fd5b5061031d7111dbdd995c9b985b98d950dbdb9d1c9858dd60721b81565b34801561062357600080fd5b5061034b611a52565b34801561063857600080fd5b5061034b610647366004613e0b565b611a66565b34801561065857600080fd5b5060655461066c906001600160a01b031681565b6040516001600160a01b0390911681526020016102b9565b34801561069057600080fd5b506106a461069f366004613acc565b611c00565b6040516102b999989796959493929190613e2d565b3480156106c557600080fd5b5061031d69456e7653746f7261676560b01b81565b3480156106e657600080fd5b5061066c6106f5366004613acc565b6000908152606860205260409020600101546001600160a01b031690565b34801561071f57600080fd5b5061074861072e366004613acc565b6000908152606e6020526040902080546001909101549091565b604080519283526020830191909152016102b9565b34801561076957600080fd5b506033546001600160a01b031661066c565b34801561078757600080fd5b5061034b610796366004613eac565b611e0d565b3480156107a757600080fd5b5061031d6c42616c6c6f7453746f7261676560981b81565b3480156107cb57600080fd5b5061034b6107da366004613e0b565b6120f8565b3480156107eb57600080fd5b5061034b6107fa366004613c7b565b612276565b34801561080b57600080fd5b50606c546001600160a01b031661066c565b34801561082957600080fd5b50606d5461031d565b34801561083e57600080fd5b5061034b61084d366004613ef5565b61231e565b34801561085e57600080fd5b5061034b61086d366004613c4f565b612463565b34801561087e57600080fd5b5061034b61088d366004613c7b565b612643565b34801561089e57600080fd5b5061031d61275d565b3480156108b357600080fd5b5061034b6108c2366004613f26565b6127a4565b3480156108d357600080fd5b5061034b6108e2366004613c7b565b612ab4565b3480156108f357600080fd5b5061034b610902366004613c7b565b612b2a565b34801561091357600080fd5b506105de610922366004614040565b66ffffffffffffff82166000908152606b602090815260408083206001600160a01b038516845290915290205460ff1692915050565b60008060008060606000806000806000806000606660008e815260200190815260200160002060405180610180016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020016005820180546109f190614085565b80601f0160208091040260200160405190810160405280929190818152602001828054610a1d90614085565b8015610a6a5780601f10610a3f57610100808354040283529160200191610a6a565b820191906000526020600020905b815481529060010190602001808311610a4d57829003601f168201915b5050505050815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a820160009054906101000a900460ff16151515158152602001600b82015481525050905080602001519b5080604001519a5080606001519950806080015198508060a0015197508060c0015196508060e00151955080610100015194508061012001519350806101400151925080610160015191505091939597999b90929496989a50565b6000818152606660209081526040808320815161018081018352815481526001820154938101939093526002810154918301919091526003810154606083015260048101546001600160a01b031660808301526005810180548493849384939192909160a0840191610b9490614085565b80601f0160208091040260200160405190810160405280929190818152602001828054610bc090614085565b8015610c0d5780601f10610be257610100808354040283529160200191610c0d565b820191906000526020600020905b815481529060010190602001808311610bf057829003601f168201915b5050505050815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a820160009054906101000a900460ff16151515158152602001600b82015481525050905080602001519350806040015192508061016001519150509193909250565b600033610c8e612b4a565b6001600160a01b031614610cbd5760405162461bcd60e51b8152600401610cb4906140bf565b60405180910390fd5b610cc5612b6a565b6001600160a01b0316306001600160a01b031614610cf55760405162461bcd60e51b8152600401610cb4906140e6565b60048514610d155760405162461bcd60e51b8152600401610cb49061410b565b6001600160a01b038216610d3b5760405162461bcd60e51b8152600401610cb490614138565b610d4786868686612b85565b506040805180820182528681526001600160a01b0392831660208083019182526000898152606890915292909220905181559051600190910180546001600160a01b03191691909216179055509192915050565b8033610da5612b4a565b6001600160a01b03161480610dd357506000818152606660205260409020600401546001600160a01b031633145b610def5760405162461bcd60e51b8152600401610cb4906140bf565b610df7612b6a565b6001600160a01b0316306001600160a01b031614610e275760405162461bcd60e51b8152600401610cb4906140e6565b6000828152606660205260409020548214610e545760405162461bcd60e51b8152600401610cb490614163565b6000828152606660205260409020600a015460ff1615610e865760405162461bcd60e51b8152600401610cb49061418f565b600160008381526066602052604090206009015414610eb75760405162461bcd60e51b8152600401610cb4906141ba565b60008281526066602052604090206005600982015560405183907fd5e541d004c50564e5e05fcbc6be2916c68d817507693dc3774c69dde4ce13dc90600090a2505050565b6000610f06612e40565b6001600160a01b03166333be496e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f6791906141e3565b905090565b600081815260696020526040902060018101546002820154600383018054929391926060929190610f9c90614085565b80601f0160208091040260200160405190810160405280929190818152602001828054610fc890614085565b80156110155780601f10610fea57610100808354040283529160200191611015565b820191906000526020600020905b815481529060010190602001808311610ff857829003601f168201915b50505050509150509193909250565b3361102d612b4a565b6001600160a01b0316146110535760405162461bcd60e51b8152600401610cb4906140bf565b61105b612b6a565b6001600160a01b0316306001600160a01b03161461108b5760405162461bcd60e51b8152600401610cb4906140e6565b6000928352606e602052604090922090815560010155565b6110ab612e58565b6001600160a01b0381166110f35760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b6044820152606401610cb4565b606c80546001600160a01b0319166001600160a01b0383169081179091556040517f3d809312b6e303291a93b307c7ddbd0960c094f5f0fb4e3ba0758775013edeb390600090a250565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036111855760405162461bcd60e51b8152600401610cb4906141fc565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166111ce60008051602061442e833981519152546001600160a01b031690565b6001600160a01b0316146111f45760405162461bcd60e51b8152600401610cb490614248565b6111fd81612eb2565b6040805160008082526020820190925261121991839190612eba565b50565b600033611227612b4a565b6001600160a01b03161461124d5760405162461bcd60e51b8152600401610cb4906140bf565b611255612b6a565b6001600160a01b0316306001600160a01b0316146112855760405162461bcd60e51b8152600401610cb4906140e6565b61129087858461302a565b6112ac5760405162461bcd60e51b8152600401610cb490614138565b6112b888888888612b85565b6040805160808101825289815260208082018781528284018781526060840187815260008e81526069855295909520845181559151600183015551600282015592518051929384939092611313926003850192910190613a33565b50999a9950505050505050505050565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361136b5760405162461bcd60e51b8152600401610cb4906141fc565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166113b460008051602061442e833981519152546001600160a01b031690565b6001600160a01b0316146113da5760405162461bcd60e51b8152600401610cb490614248565b6113e382612eb2565b6113ef82826001612eba565b5050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146114935760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610cb4565b5060008051602061442e83398151915290565b6000818152606660209081526040808320815161018081018352815481526001820154938101939093526002810154918301919091526003810154606083015260048101546001600160a01b031660808301526005810180548493849384939192909160a084019161151790614085565b80601f016020809104026020016040519081016040528092919081815260200182805461154390614085565b80156115905780601f1061156557610100808354040283529160200191611590565b820191906000526020600020905b81548152906001019060200180831161157357829003601f168201915b505050918352505060068201546020820152600782015460408201526008820154606082015260098201546080820152600a82015460ff16151560a0820152600b9091015460c09182015281015160e0820151610100909201519097919650945092505050565b6000818152606a60205260408120548190819081908190819087146116535760405162461bcd60e51b81526020600482015260126024820152711b9bdd08195e1a5cdd1959081d9bdd19525960721b6044820152606401610cb4565b50505060009384525050606a6020908152604092839020835160c0810185528154808252600183015493820184905260028301546001600160a01b031695820186905260038301546060830181905260048401546080840181905260059094015460a090930183905290969395945092565b81336116cf612b4a565b6001600160a01b031614806116fd57506000818152606660205260409020600401546001600160a01b031633145b6117195760405162461bcd60e51b8152600401610cb4906140bf565b611721612b6a565b6001600160a01b0316306001600160a01b0316146117515760405162461bcd60e51b8152600401610cb4906140e6565b818061175b610efc565b11156117a95760405162461bcd60e51b815260206004820152601c60248201527f556e646572206d696e2076616c7565206f6620206475726174696f6e000000006044820152606401610cb4565b6117b161275d565b8111156118005760405162461bcd60e51b815260206004820152601a60248201527f4f766572206d61782076616c7565206f66206475726174696f6e0000000000006044820152606401610cb4565b600084815260666020526040902054841461182d5760405162461bcd60e51b8152600401610cb490614163565b6000848152606660205260409020600a015460ff161561185f5760405162461bcd60e51b8152600401610cb49061418f565b6001600085815260666020526040902060090154146118905760405162461bcd60e51b8152600401610cb4906141ba565b600084815260666020526040808220600b810186905590519091339187917ff0855c0e0ad9b8a162b87f2e4e07d4b2a3f0a45126b15ff4a78b217ad19a901a91a35050505050565b6000818152606660209081526040808320815161018081018352815481526001820154938101939093526002810154918301919091526003810154606083015260048101546001600160a01b031660808301526005810180548493849384939192909160a084019161194990614085565b80601f016020809104026020016040519081016040528092919081815260200182805461197590614085565b80156119c25780601f10611997576101008083540402835291602001916119c2565b820191906000526020600020905b8154815290600101906020018083116119a557829003601f168201915b50505091835250506006820154602082015260078201546040820152600882015460608083019190915260098301546080830152600a83015460ff16151560a0830152600b9092015460c090910152810151610120820151610140909201519097919650945092505050565b6000611a38612b6a565b6001600160a01b0316306001600160a01b03161415905090565b611a5a612e58565b611a64600061310a565b565b33611a6f612b4a565b6001600160a01b031614611a955760405162461bcd60e51b8152600401610cb4906140bf565b611a9d612b6a565b6001600160a01b0316306001600160a01b031614611acd5760405162461bcd60e51b8152600401610cb4906140e6565b6000828152606660205260409020548214611afa5760405162461bcd60e51b8152600401610cb490614163565b6000828152606760205260409020548214611b575760405162461bcd60e51b815260206004820152601860248201527f6e6f7420657869737465642042616c6c6f744d656d62657200000000000000006044820152606401610cb4565b6000828152606660205260409020600a015460ff1615611b895760405162461bcd60e51b8152600401610cb49061418f565b600160008381526066602052604090206009015414611bba5760405162461bcd60e51b8152600401610cb4906141ba565b6000828152606760205260408082206009810184905590519091339185917ff0855c0e0ad9b8a162b87f2e4e07d4b2a3f0a45126b15ff4a78b217ad19a901a91a3505050565b600081815260676020526040812060018101546002820154600383015460048401546005850180546001600160a01b039586169794861696938616959092169360609384938493839291611c5390614085565b80601f0160208091040260200160405190810160405280929190818152602001828054611c7f90614085565b8015611ccc5780601f10611ca157610100808354040283529160200191611ccc565b820191906000526020600020905b815481529060010190602001808311611caf57829003601f168201915b50505050509550806006018054611ce290614085565b80601f0160208091040260200160405190810160405280929190818152602001828054611d0e90614085565b8015611d5b5780601f10611d3057610100808354040283529160200191611d5b565b820191906000526020600020905b815481529060010190602001808311611d3e57829003601f168201915b50505050509450806007018054611d7190614085565b80601f0160208091040260200160405190810160405280929190818152602001828054611d9d90614085565b8015611dea5780601f10611dbf57610100808354040283529160200191611dea565b820191906000526020600020905b815481529060010190602001808311611dcd57829003601f168201915b505050505093508060080154925080600901549150509193959799909294969850565b33611e16612b4a565b6001600160a01b031614611e3c5760405162461bcd60e51b8152600401610cb4906140bf565b611e44612b6a565b6001600160a01b0316306001600160a01b031614611e745760405162461bcd60e51b8152600401610cb4906140e6565b6001821480611e835750600282145b611ec25760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b2103232b1b4b9b4b7b760811b6044820152606401610cb4565b6000848152606660205260409020548414611eef5760405162461bcd60e51b8152600401610cb490614163565b6000858152606a6020526040902054859003611f465760405162461bcd60e51b8152602060048201526016602482015275185b1c9958591e48195e1a5cdd1959081d9bdd19525960521b6044820152606401610cb4565b6000848152606b602090815260408083206001600160a01b038716845290915290205460ff1615611fa95760405162461bcd60e51b815260206004820152600d60248201526c185b1c9958591e481d9bdd1959609a1b6044820152606401610cb4565b6002600085815260666020526040902060090154146120015760405162461bcd60e51b81526020600482015260146024820152734e6f7420496e50726f677265737320537461746560601b6044820152606401610cb4565b6040518060c00160405280868152602001858152602001846001600160a01b0316815260200183815260200182815260200161203a4290565b90526000868152606a602090815260409182902083518155908301516001820155908201516002820180546001600160a01b0319166001600160a01b03909216919091179055606082015160038201556080820151600482015560a0909101516005909101556120ac8484848461315c565b826001600160a01b031684867f41df84b3b467b06744e40c92613c666324e7c640ce0a41ec06efdf602d367606856040516120e991815260200190565b60405180910390a45050505050565b33612101612b4a565b6001600160a01b0316146121275760405162461bcd60e51b8152600401610cb4906140bf565b61212f612b6a565b6001600160a01b0316306001600160a01b03161461215f5760405162461bcd60e51b8152600401610cb4906140e6565b600082815260666020526040902054821461218c5760405162461bcd60e51b8152600401610cb490614163565b6000828152606660205260409020600a015460ff16156121be5760405162461bcd60e51b8152600401610cb49061418f565b60038114806121cd5750600481145b6122105760405162461bcd60e51b8152602060048201526014602482015273496e76616c69642042616c6c6f7420537461746560601b6044820152606401610cb4565b6000828152606660205260409081902060098101839055600a8101805460ff19166001179055905183907fdc921f027328d7238b58d77649ebcbd0c8b1c494c66ba53dfe53e0de65f6dd9f906122699085815260200190565b60405180910390a2505050565b61227e612e58565b6001600160a01b0381166122d45760405162461bcd60e51b815260206004820152601a60248201527f416464726573732073686f756c64206265206e6f6e2d7a65726f0000000000006044820152606401610cb4565b606580546001600160a01b0319166001600160a01b0383169081179091556040517f278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd90600090a250565b8133612328612b4a565b6001600160a01b0316148061235657506000818152606660205260409020600401546001600160a01b031633145b6123725760405162461bcd60e51b8152600401610cb4906140bf565b61237a612b6a565b6001600160a01b0316306001600160a01b0316146123aa5760405162461bcd60e51b8152600401610cb4906140e6565b60008381526066602052604090205483146123d75760405162461bcd60e51b8152600401610cb490614163565b6000838152606660205260409020600a015460ff16156124095760405162461bcd60e51b8152600401610cb49061418f565b60008381526066602090815260409091208351909161242f916005840191860190613a33565b50604051339085907ff0855c0e0ad9b8a162b87f2e4e07d4b2a3f0a45126b15ff4a78b217ad19a901a90600090a350505050565b3361246c612b4a565b6001600160a01b0316146124925760405162461bcd60e51b8152600401610cb4906140bf565b61249a612b6a565b6001600160a01b0316306001600160a01b0316146124ca5760405162461bcd60e51b8152600401610cb4906140e6565b81816000821180156124dc5750600081115b61251c5760405162461bcd60e51b815260206004820152601160248201527007374617274206f7220656e64206973203607c1b6044820152606401610cb4565b81811161255a5760405162461bcd60e51b815260206004820152600c60248201526b1cdd185c9d080f8f48195b9960a21b6044820152606401610cb4565b60008581526066602052604090205485146125875760405162461bcd60e51b8152600401610cb490614163565b6000858152606660205260409020600a015460ff16156125b95760405162461bcd60e51b8152600401610cb49061418f565b6001600086815260666020526040902060090154146125ea5760405162461bcd60e51b8152600401610cb4906141ba565b600085815260666020526040902060018101859055600280820185905560098201556040518490869088907fd9938a514dab5cdce149a77493f694bd70c24a4833a278fd4d86fbdf859099c590600090a4505050505050565b600054610100900460ff16158080156126635750600054600160ff909116105b8061267d5750303b15801561267d575060005460ff166001145b6126e05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610cb4565b6000805460ff191660011790558015612703576000805461ff0019166101001790555b61270b6132c4565b61271482612276565b80156113ef576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6000612767612e40565b6001600160a01b0316631b27e01b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f43573d6000803e3d6000fd5b336127ad612b4a565b6001600160a01b0316146127d35760405162461bcd60e51b8152600401610cb4906140bf565b6127db612b6a565b6001600160a01b0316306001600160a01b03161461280b5760405162461bcd60e51b8152600401610cb4906140e6565b61281c8b89898989898989896132f3565b6128385760405162461bcd60e51b8152600401610cb490614138565b6128448c8c8c8c612b85565b6128be6040518061014001604052806000815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160608152602001606081526020016060815260200160008152602001600081525090565b8c8160000181815250508881602001906001600160a01b031690816001600160a01b0316815250508781604001906001600160a01b031690816001600160a01b0316815250508681606001906001600160a01b031690816001600160a01b0316815250508581608001906001600160a01b031690816001600160a01b031681525050848160a00181905250838160c00181905250828160e00181905250818161010001818152505080606760008f81526020019081526020016000206000820151816000015560208201518160010160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160020160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060608201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060808201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060a0820151816005019080519060200190612a54929190613a33565b5060c08201518051612a70916006840191602090910190613a33565b5060e08201518051612a8c916007840191602090910190613a33565b5061010082015160088201556101209091015160099091015550505050505050505050505050565b612abc612e58565b6001600160a01b038116612b215760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610cb4565b6112198161310a565b612b32612e58565b6001600160a01b03811615611219576111fd81612eb2565b6000610f677111dbdd995c9b985b98d950dbdb9d1c9858dd60721b613752565b6000610f676c42616c6c6f7453746f7261676560981b613752565b8180612b8f610efc565b1115612bdd5760405162461bcd60e51b815260206004820152601c60248201527f556e646572206d696e2076616c7565206f6620206475726174696f6e000000006044820152606401610cb4565b612be561275d565b811115612c345760405162461bcd60e51b815260206004820152601a60248201527f4f766572206d61782076616c7565206f66206475726174696f6e0000000000006044820152606401610cb4565b600085815260666020526040902054859003612c8b5760405162461bcd60e51b8152602060048201526016602482015275105b1c9958591e48195e1a5cdd19590818985b1b1bdd60521b6044820152606401610cb4565b612cfa6040518061018001604052806000815260200160008152602001600081526020016000815260200160006001600160a01b031681526020016060815260200160008152602001600081526020016000815260200160008152602001600015158152602001600081525090565b858152606081018590526001600160a01b0383166080820152600161012082015260006101408201819052610160820185905286815260666020908152604091829020835181558184015160018201559183015160028301556060830151600383015560808301516004830180546001600160a01b0319166001600160a01b0390921691909117905560a08301518051849392612d9e926005850192910190613a33565b5060c0820151600682015560e0820151600782015561010082015160088201556101208201516009820155610140820151600a8201805460ff191691151591909117905561016090910151600b90910155606d54612dfd906001614294565b606d556040516001600160a01b03841690869088907fd1ba591c76ef71222e2d30b8277758713cc6eef1de29efaf98a716744ac2420b90600090a4505050505050565b6000610f6769456e7653746f7261676560b01b613752565b6033546001600160a01b03163314611a645760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610cb4565b611219612e58565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615612ef257612eed836137c6565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612f4c575060408051601f3d908101601f19168201909252612f49918101906141e3565b60015b612faf5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610cb4565b60008051602061442e833981519152811461301e5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610cb4565b50612eed838383613862565b60006005841461304c5760405162461bcd60e51b8152600401610cb49061410b565b826130a35760405162461bcd60e51b815260206004820152602160248201527f496e76616c696420656e7669726f6e6d656e74207661726961626c65206e616d6044820152606560f81b6064820152608401610cb4565b60008251116130ff5760405162461bcd60e51b815260206004820152602260248201527f496e76616c696420656e7669726f6e6d656e74207661726961626c652076616c604482015261756560f01b6064820152608401610cb4565b5060015b9392505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600182148061316b5750600282145b6131aa5760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b2103232b1b4b9b4b7b760811b6044820152606401610cb4565b60008481526066602052604090205484146131d75760405162461bcd60e51b8152600401610cb490614163565b6000848152606b602090815260408083206001600160a01b038716845290915290205460ff161561323a5760405162461bcd60e51b815260206004820152600d60248201526c185b1c9958591e481d9bdd1959609a1b6044820152606401610cb4565b6000848152606660209081526040808320606b83528184206001600160a01b0388168552909252909120805460ff19166001908117909155600682015461328091614294565b6006820155600183036132a75781816007015461329d9190614294565b60078201556132bd565b8181600801546132b79190614294565b60088201555b5050505050565b600054610100900460ff166132eb5760405162461bcd60e51b8152600401610cb4906142ba565b611a6461388d565b600060018a10158015613307575060038a11155b6133235760405162461bcd60e51b8152600401610cb49061410b565b60028a036134e2576001600160a01b0389166133515760405162461bcd60e51b8152600401610cb490614305565b6001600160a01b038816156133785760405162461bcd60e51b8152600401610cb49061433c565b6001600160a01b0387161561339f5760405162461bcd60e51b8152600401610cb490614373565b6001600160a01b038616156133c65760405162461bcd60e51b8152600401610cb4906143aa565b84511561340d5760405162461bcd60e51b8152602060048201526015602482015274496e76616c6964206e6577206e6f6465206e616d6560581b6044820152606401610cb4565b8351156134525760405162461bcd60e51b8152602060048201526013602482015272125b9d985b1a59081b995dc81b9bd919481a59606a1b6044820152606401610cb4565b8251156134975760405162461bcd60e51b81526020600482015260136024820152720496e76616c6964206e6577206e6f646520495606c1b6044820152606401610cb4565b81156134dd5760405162461bcd60e51b8152602060048201526015602482015274125b9d985b1a59081b995dc81b9bd91948141bdc9d605a1b6044820152606401610cb4565b613742565b600085511161352b5760405162461bcd60e51b8152602060048201526015602482015274496e76616c6964206e6577206e6f6465206e616d6560581b6044820152606401610cb4565b83516040146135725760405162461bcd60e51b8152602060048201526013602482015272125b9d985b1a59081b995dc81b9bd919481a59606a1b6044820152606401610cb4565b60008351116135b95760405162461bcd60e51b81526020600482015260136024820152720496e76616c6964206e6577206e6f646520495606c1b6044820152606401610cb4565b600082116136015760405162461bcd60e51b8152602060048201526015602482015274125b9d985b1a59081b995dc81b9bd91948141bdc9d605a1b6044820152606401610cb4565b60018a036136a2576001600160a01b038916156136305760405162461bcd60e51b8152600401610cb490614305565b6001600160a01b0388166136565760405162461bcd60e51b8152600401610cb49061433c565b6001600160a01b03871661367c5760405162461bcd60e51b8152600401610cb490614373565b6001600160a01b0386166134dd5760405162461bcd60e51b8152600401610cb4906143aa565b60038a03613742576001600160a01b0389166136d05760405162461bcd60e51b8152600401610cb490614305565b6001600160a01b0388166136f65760405162461bcd60e51b8152600401610cb49061433c565b6001600160a01b03871661371c5760405162461bcd60e51b8152600401610cb490614373565b6001600160a01b0386166137425760405162461bcd60e51b8152600401610cb4906143aa565b5060019998505050505050505050565b606554604051630d2020dd60e01b8152600481018390526000916001600160a01b031690630d2020dd90602401602060405180830381865afa15801561379c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137c091906143e1565b92915050565b6001600160a01b0381163b6138335760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610cb4565b60008051602061442e83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b61386b836138bd565b6000825111806138785750805b15612eed5761388783836138fd565b50505050565b600054610100900460ff166138b45760405162461bcd60e51b8152600401610cb4906142ba565b611a643361310a565b6138c6816137c6565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060613103838360405180606001604052806027815260200161444e6027913960606001600160a01b0384163b6139855760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610cb4565b600080856001600160a01b0316856040516139a091906143fe565b600060405180830381855af49150503d80600081146139db576040519150601f19603f3d011682016040523d82523d6000602084013e6139e0565b606091505b50915091506139f08282866139fa565b9695505050505050565b60608315613a09575081613103565b825115613a195782518084602001fd5b8160405162461bcd60e51b8152600401610cb4919061441a565b828054613a3f90614085565b90600052602060002090601f016020900481019282613a615760008555613aa7565b82601f10613a7a57805160ff1916838001178555613aa7565b82800160010185558215613aa7579182015b82811115613aa7578251825591602001919060010190613a8c565b50613ab3929150613ab7565b5090565b5b80821115613ab35760008155600101613ab8565b600060208284031215613ade57600080fd5b5035919050565b60005b83811015613b00578181015183820152602001613ae8565b838111156138875750506000910152565b60008151808452613b29816020860160208601613ae5565b601f01601f19169290920160200192915050565b60006101608d83528c60208401528b604084015260018060a01b038b166060840152806080840152613b718184018b613b11565b60a0840199909952505060c081019590955260e085019390935261010084019190915215156101208301526101409091015295945050505050565b6001600160a01b038116811461121957600080fd5b8035613bcc81613bac565b919050565b600080600080600060a08688031215613be957600080fd5b8535945060208601359350604086013592506060860135613c0981613bac565b91506080860135613c1981613bac565b809150509295509295909350565b838152826020820152606060408201526000613c466060830184613b11565b95945050505050565b600080600060608486031215613c6457600080fd5b505081359360208301359350604090920135919050565b600060208284031215613c8d57600080fd5b813561310381613bac565b634e487b7160e01b600052604160045260246000fd5b600082601f830112613cbf57600080fd5b813567ffffffffffffffff80821115613cda57613cda613c98565b604051601f8301601f19908116603f01168101908282118183101715613d0257613d02613c98565b81604052838152866020858801011115613d1b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600080600080600060e0888a031215613d5657600080fd5b8735965060208801359550604088013594506060880135613d7681613bac565b93506080880135925060a0880135915060c088013567ffffffffffffffff811115613da057600080fd5b613dac8a828b01613cae565b91505092959891949750929550565b60008060408385031215613dce57600080fd5b8235613dd981613bac565b9150602083013567ffffffffffffffff811115613df557600080fd5b613e0185828601613cae565b9150509250929050565b60008060408385031215613e1e57600080fd5b50508035926020909101359150565b6001600160a01b038a81168252898116602083015288811660408301528716606082015261012060808201819052600090613e6a83820189613b11565b905082810360a0840152613e7e8188613b11565b905082810360c0840152613e928187613b11565b60e084019590955250506101000152979650505050505050565b600080600080600060a08688031215613ec457600080fd5b85359450602086013593506040860135613edd81613bac565b94979396509394606081013594506080013592915050565b60008060408385031215613f0857600080fd5b82359150602083013567ffffffffffffffff811115613df557600080fd5b6000806000806000806000806000806000806101808d8f031215613f4957600080fd5b8c359b5060208d01359a5060408d01359950613f6760608e01613bc1565b9850613f7560808e01613bc1565b9750613f8360a08e01613bc1565b9650613f9160c08e01613bc1565b9550613f9f60e08e01613bc1565b945067ffffffffffffffff6101008e01351115613fbb57600080fd5b613fcc8e6101008f01358f01613cae565b935067ffffffffffffffff6101208e01351115613fe857600080fd5b613ff98e6101208f01358f01613cae565b925067ffffffffffffffff6101408e0135111561401557600080fd5b6140268e6101408f01358f01613cae565b91506101608d013590509295989b509295989b509295989b565b6000806040838503121561405357600080fd5b823566ffffffffffffff8116811461406a57600080fd5b9150602083013561407a81613bac565b809150509250929050565b600181811c9082168061409957607f821691505b6020821081036140b957634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252600d908201526c2737902832b936b4b9b9b4b7b760991b604082015260600190565b6020808252600b908201526a125cc8111a5cd8589b195960aa1b604082015260600190565b602080825260139082015272496e76616c69642042616c6c6f74205479706560681b604082015260600190565b60208082526011908201527024b73b30b634b2102830b930b6b2ba32b960791b604082015260600190565b6020808252601290820152711b9bdd08195e1a5cdd19590810985b1b1bdd60721b604082015260600190565b602080825260119082015270185b1c9958591e48199a5b985b1a5e9959607a1b604082015260600190565b6020808252600f908201526e4e6f7420526561647920537461746560881b604082015260600190565b6000602082840312156141f557600080fd5b5051919050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b600082198211156142b557634e487b7160e01b600052601160045260246000fd5b500190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b6020808252601a908201527f496e76616c6964206f6c64207374616b65722061646472657373000000000000604082015260600190565b6020808252601a908201527f496e76616c6964206e6577207374616b65722061646472657373000000000000604082015260600190565b60208082526019908201527f496e76616c6964206e657720766f746572206164647265737300000000000000604082015260600190565b6020808252601a908201527f496e76616c6964206e6577207265776172642061646472657373000000000000604082015260600190565b6000602082840312156143f357600080fd5b815161310381613bac565b60008251614410818460208701613ae5565b9190910192915050565b6020815260006131036020830184613b1156fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220d596f8c32db0e2a432c61e3e481cfa3a0849a6a96ced7f94574bcc375400012264736f6c634300080e0033",
}

// BallotStorageImpABI is the input ABI used to generate the binding from.
// Deprecated: Use BallotStorageImpMetaData.ABI instead.
var BallotStorageImpABI = BallotStorageImpMetaData.ABI

// Deprecated: Use BallotStorageImpMetaData.Sigs instead.
// BallotStorageImpFuncSigs maps the 4-byte function signature to its string representation.
var BallotStorageImpFuncSigs = BallotStorageImpMetaData.Sigs

// BallotStorageImpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BallotStorageImpMetaData.Bin instead.
var BallotStorageImpBin = BallotStorageImpMetaData.Bin

// DeployBallotStorageImp deploys a new Ethereum contract, binding an instance of BallotStorageImp to it.
func DeployBallotStorageImp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BallotStorageImp, error) {
	parsed, err := BallotStorageImpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BallotStorageImpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BallotStorageImp{BallotStorageImpCaller: BallotStorageImpCaller{contract: contract}, BallotStorageImpTransactor: BallotStorageImpTransactor{contract: contract}, BallotStorageImpFilterer: BallotStorageImpFilterer{contract: contract}}, nil
}

// BallotStorageImp is an auto generated Go binding around an Ethereum contract.
type BallotStorageImp struct {
	BallotStorageImpCaller     // Read-only binding to the contract
	BallotStorageImpTransactor // Write-only binding to the contract
	BallotStorageImpFilterer   // Log filterer for contract events
}

// BallotStorageImpCaller is an auto generated read-only Go binding around an Ethereum contract.
type BallotStorageImpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotStorageImpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BallotStorageImpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotStorageImpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BallotStorageImpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotStorageImpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BallotStorageImpSession struct {
	Contract     *BallotStorageImp // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotStorageImpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BallotStorageImpCallerSession struct {
	Contract *BallotStorageImpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BallotStorageImpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BallotStorageImpTransactorSession struct {
	Contract     *BallotStorageImpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BallotStorageImpRaw is an auto generated low-level Go binding around an Ethereum contract.
type BallotStorageImpRaw struct {
	Contract *BallotStorageImp // Generic contract binding to access the raw methods on
}

// BallotStorageImpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BallotStorageImpCallerRaw struct {
	Contract *BallotStorageImpCaller // Generic read-only contract binding to access the raw methods on
}

// BallotStorageImpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BallotStorageImpTransactorRaw struct {
	Contract *BallotStorageImpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBallotStorageImp creates a new instance of BallotStorageImp, bound to a specific deployed contract.
func NewBallotStorageImp(address common.Address, backend bind.ContractBackend) (*BallotStorageImp, error) {
	contract, err := bindBallotStorageImp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImp{BallotStorageImpCaller: BallotStorageImpCaller{contract: contract}, BallotStorageImpTransactor: BallotStorageImpTransactor{contract: contract}, BallotStorageImpFilterer: BallotStorageImpFilterer{contract: contract}}, nil
}

// NewBallotStorageImpCaller creates a new read-only instance of BallotStorageImp, bound to a specific deployed contract.
func NewBallotStorageImpCaller(address common.Address, caller bind.ContractCaller) (*BallotStorageImpCaller, error) {
	contract, err := bindBallotStorageImp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpCaller{contract: contract}, nil
}

// NewBallotStorageImpTransactor creates a new write-only instance of BallotStorageImp, bound to a specific deployed contract.
func NewBallotStorageImpTransactor(address common.Address, transactor bind.ContractTransactor) (*BallotStorageImpTransactor, error) {
	contract, err := bindBallotStorageImp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpTransactor{contract: contract}, nil
}

// NewBallotStorageImpFilterer creates a new log filterer instance of BallotStorageImp, bound to a specific deployed contract.
func NewBallotStorageImpFilterer(address common.Address, filterer bind.ContractFilterer) (*BallotStorageImpFilterer, error) {
	contract, err := bindBallotStorageImp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpFilterer{contract: contract}, nil
}

// bindBallotStorageImp binds a generic wrapper to an already deployed contract.
func bindBallotStorageImp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BallotStorageImpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BallotStorageImp *BallotStorageImpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BallotStorageImp.Contract.BallotStorageImpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BallotStorageImp *BallotStorageImpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.BallotStorageImpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BallotStorageImp *BallotStorageImpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.BallotStorageImpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BallotStorageImp *BallotStorageImpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BallotStorageImp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BallotStorageImp *BallotStorageImpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BallotStorageImp *BallotStorageImpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.contract.Transact(opts, method, params...)
}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCaller) BALLOTSTORAGENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "BALLOT_STORAGE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpSession) BALLOTSTORAGENAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.BALLOTSTORAGENAME(&_BallotStorageImp.CallOpts)
}

// BALLOTSTORAGENAME is a free data retrieval call binding the contract method 0x9986e4b9.
//
// Solidity: function BALLOT_STORAGE_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCallerSession) BALLOTSTORAGENAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.BALLOTSTORAGENAME(&_BallotStorageImp.CallOpts)
}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCaller) ECOSYSTEMNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "ECOSYSTEM_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpSession) ECOSYSTEMNAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.ECOSYSTEMNAME(&_BallotStorageImp.CallOpts)
}

// ECOSYSTEMNAME is a free data retrieval call binding the contract method 0x34125c84.
//
// Solidity: function ECOSYSTEM_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCallerSession) ECOSYSTEMNAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.ECOSYSTEMNAME(&_BallotStorageImp.CallOpts)
}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCaller) ENVSTORAGENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "ENV_STORAGE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpSession) ENVSTORAGENAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.ENVSTORAGENAME(&_BallotStorageImp.CallOpts)
}

// ENVSTORAGENAME is a free data retrieval call binding the contract method 0x7bf46530.
//
// Solidity: function ENV_STORAGE_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCallerSession) ENVSTORAGENAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.ENVSTORAGENAME(&_BallotStorageImp.CallOpts)
}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCaller) GOVNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "GOV_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpSession) GOVNAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.GOVNAME(&_BallotStorageImp.CallOpts)
}

// GOVNAME is a free data retrieval call binding the contract method 0x6c78d2cf.
//
// Solidity: function GOV_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCallerSession) GOVNAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.GOVNAME(&_BallotStorageImp.CallOpts)
}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCaller) MAINTENANCENAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "MAINTENANCE_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpSession) MAINTENANCENAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.MAINTENANCENAME(&_BallotStorageImp.CallOpts)
}

// MAINTENANCENAME is a free data retrieval call binding the contract method 0x4bd1ed76.
//
// Solidity: function MAINTENANCE_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCallerSession) MAINTENANCENAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.MAINTENANCENAME(&_BallotStorageImp.CallOpts)
}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCaller) REWARDPOOLNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "REWARD_POOL_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpSession) REWARDPOOLNAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.REWARDPOOLNAME(&_BallotStorageImp.CallOpts)
}

// REWARDPOOLNAME is a free data retrieval call binding the contract method 0x2f40992e.
//
// Solidity: function REWARD_POOL_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCallerSession) REWARDPOOLNAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.REWARDPOOLNAME(&_BallotStorageImp.CallOpts)
}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCaller) STAKINGNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "STAKING_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpSession) STAKINGNAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.STAKINGNAME(&_BallotStorageImp.CallOpts)
}

// STAKINGNAME is a free data retrieval call binding the contract method 0x1e0cba0d.
//
// Solidity: function STAKING_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCallerSession) STAKINGNAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.STAKINGNAME(&_BallotStorageImp.CallOpts)
}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCaller) STAKINGREWARDNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "STAKING_REWARD_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpSession) STAKINGREWARDNAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.STAKINGREWARDNAME(&_BallotStorageImp.CallOpts)
}

// STAKINGREWARDNAME is a free data retrieval call binding the contract method 0x5a731cca.
//
// Solidity: function STAKING_REWARD_NAME() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCallerSession) STAKINGREWARDNAME() ([32]byte, error) {
	return _BallotStorageImp.Contract.STAKINGREWARDNAME(&_BallotStorageImp.CallOpts)
}

// GetBallotAddress is a free data retrieval call binding the contract method 0x7efa9ae3.
//
// Solidity: function getBallotAddress(uint256 _id) view returns(address newGovernanceAddress)
func (_BallotStorageImp *BallotStorageImpCaller) GetBallotAddress(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getBallotAddress", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBallotAddress is a free data retrieval call binding the contract method 0x7efa9ae3.
//
// Solidity: function getBallotAddress(uint256 _id) view returns(address newGovernanceAddress)
func (_BallotStorageImp *BallotStorageImpSession) GetBallotAddress(_id *big.Int) (common.Address, error) {
	return _BallotStorageImp.Contract.GetBallotAddress(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotAddress is a free data retrieval call binding the contract method 0x7efa9ae3.
//
// Solidity: function getBallotAddress(uint256 _id) view returns(address newGovernanceAddress)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetBallotAddress(_id *big.Int) (common.Address, error) {
	return _BallotStorageImp.Contract.GetBallotAddress(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotBasic is a free data retrieval call binding the contract method 0x02b385fb.
//
// Solidity: function getBallotBasic(uint256 _id) view returns(uint256 startTime, uint256 endTime, uint256 ballotType, address creator, bytes memo, uint256 totalVoters, uint256 powerOfAccepts, uint256 powerOfRejects, uint256 state, bool isFinalized, uint256 duration)
func (_BallotStorageImp *BallotStorageImpCaller) GetBallotBasic(opts *bind.CallOpts, _id *big.Int) (struct {
	StartTime      *big.Int
	EndTime        *big.Int
	BallotType     *big.Int
	Creator        common.Address
	Memo           []byte
	TotalVoters    *big.Int
	PowerOfAccepts *big.Int
	PowerOfRejects *big.Int
	State          *big.Int
	IsFinalized    bool
	Duration       *big.Int
}, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getBallotBasic", _id)

	outstruct := new(struct {
		StartTime      *big.Int
		EndTime        *big.Int
		BallotType     *big.Int
		Creator        common.Address
		Memo           []byte
		TotalVoters    *big.Int
		PowerOfAccepts *big.Int
		PowerOfRejects *big.Int
		State          *big.Int
		IsFinalized    bool
		Duration       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BallotType = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Creator = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Memo = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.TotalVoters = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PowerOfAccepts = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.PowerOfRejects = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.State = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.IsFinalized = *abi.ConvertType(out[9], new(bool)).(*bool)
	outstruct.Duration = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBallotBasic is a free data retrieval call binding the contract method 0x02b385fb.
//
// Solidity: function getBallotBasic(uint256 _id) view returns(uint256 startTime, uint256 endTime, uint256 ballotType, address creator, bytes memo, uint256 totalVoters, uint256 powerOfAccepts, uint256 powerOfRejects, uint256 state, bool isFinalized, uint256 duration)
func (_BallotStorageImp *BallotStorageImpSession) GetBallotBasic(_id *big.Int) (struct {
	StartTime      *big.Int
	EndTime        *big.Int
	BallotType     *big.Int
	Creator        common.Address
	Memo           []byte
	TotalVoters    *big.Int
	PowerOfAccepts *big.Int
	PowerOfRejects *big.Int
	State          *big.Int
	IsFinalized    bool
	Duration       *big.Int
}, error) {
	return _BallotStorageImp.Contract.GetBallotBasic(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotBasic is a free data retrieval call binding the contract method 0x02b385fb.
//
// Solidity: function getBallotBasic(uint256 _id) view returns(uint256 startTime, uint256 endTime, uint256 ballotType, address creator, bytes memo, uint256 totalVoters, uint256 powerOfAccepts, uint256 powerOfRejects, uint256 state, bool isFinalized, uint256 duration)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetBallotBasic(_id *big.Int) (struct {
	StartTime      *big.Int
	EndTime        *big.Int
	BallotType     *big.Int
	Creator        common.Address
	Memo           []byte
	TotalVoters    *big.Int
	PowerOfAccepts *big.Int
	PowerOfRejects *big.Int
	State          *big.Int
	IsFinalized    bool
	Duration       *big.Int
}, error) {
	return _BallotStorageImp.Contract.GetBallotBasic(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() view returns(uint256)
func (_BallotStorageImp *BallotStorageImpCaller) GetBallotCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getBallotCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() view returns(uint256)
func (_BallotStorageImp *BallotStorageImpSession) GetBallotCount() (*big.Int, error) {
	return _BallotStorageImp.Contract.GetBallotCount(&_BallotStorageImp.CallOpts)
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() view returns(uint256)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetBallotCount() (*big.Int, error) {
	return _BallotStorageImp.Contract.GetBallotCount(&_BallotStorageImp.CallOpts)
}

// GetBallotForExit is a free data retrieval call binding the contract method 0x8c7be692.
//
// Solidity: function getBallotForExit(uint256 _id) view returns(uint256 unlockAmount, uint256 slashing)
func (_BallotStorageImp *BallotStorageImpCaller) GetBallotForExit(opts *bind.CallOpts, _id *big.Int) (struct {
	UnlockAmount *big.Int
	Slashing     *big.Int
}, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getBallotForExit", _id)

	outstruct := new(struct {
		UnlockAmount *big.Int
		Slashing     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UnlockAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slashing = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBallotForExit is a free data retrieval call binding the contract method 0x8c7be692.
//
// Solidity: function getBallotForExit(uint256 _id) view returns(uint256 unlockAmount, uint256 slashing)
func (_BallotStorageImp *BallotStorageImpSession) GetBallotForExit(_id *big.Int) (struct {
	UnlockAmount *big.Int
	Slashing     *big.Int
}, error) {
	return _BallotStorageImp.Contract.GetBallotForExit(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotForExit is a free data retrieval call binding the contract method 0x8c7be692.
//
// Solidity: function getBallotForExit(uint256 _id) view returns(uint256 unlockAmount, uint256 slashing)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetBallotForExit(_id *big.Int) (struct {
	UnlockAmount *big.Int
	Slashing     *big.Int
}, error) {
	return _BallotStorageImp.Contract.GetBallotForExit(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotMember is a free data retrieval call binding the contract method 0x73df4e01.
//
// Solidity: function getBallotMember(uint256 _id) view returns(address oldStakerAddress, address newStakerAddress, address newVoterAddress, address newRewardAddress, bytes newNodeName, bytes newNodeId, bytes newNodeIp, uint256 newNodePort, uint256 lockAmount)
func (_BallotStorageImp *BallotStorageImpCaller) GetBallotMember(opts *bind.CallOpts, _id *big.Int) (struct {
	OldStakerAddress common.Address
	NewStakerAddress common.Address
	NewVoterAddress  common.Address
	NewRewardAddress common.Address
	NewNodeName      []byte
	NewNodeId        []byte
	NewNodeIp        []byte
	NewNodePort      *big.Int
	LockAmount       *big.Int
}, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getBallotMember", _id)

	outstruct := new(struct {
		OldStakerAddress common.Address
		NewStakerAddress common.Address
		NewVoterAddress  common.Address
		NewRewardAddress common.Address
		NewNodeName      []byte
		NewNodeId        []byte
		NewNodeIp        []byte
		NewNodePort      *big.Int
		LockAmount       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OldStakerAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NewStakerAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NewVoterAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.NewRewardAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.NewNodeName = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.NewNodeId = *abi.ConvertType(out[5], new([]byte)).(*[]byte)
	outstruct.NewNodeIp = *abi.ConvertType(out[6], new([]byte)).(*[]byte)
	outstruct.NewNodePort = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.LockAmount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBallotMember is a free data retrieval call binding the contract method 0x73df4e01.
//
// Solidity: function getBallotMember(uint256 _id) view returns(address oldStakerAddress, address newStakerAddress, address newVoterAddress, address newRewardAddress, bytes newNodeName, bytes newNodeId, bytes newNodeIp, uint256 newNodePort, uint256 lockAmount)
func (_BallotStorageImp *BallotStorageImpSession) GetBallotMember(_id *big.Int) (struct {
	OldStakerAddress common.Address
	NewStakerAddress common.Address
	NewVoterAddress  common.Address
	NewRewardAddress common.Address
	NewNodeName      []byte
	NewNodeId        []byte
	NewNodeIp        []byte
	NewNodePort      *big.Int
	LockAmount       *big.Int
}, error) {
	return _BallotStorageImp.Contract.GetBallotMember(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotMember is a free data retrieval call binding the contract method 0x73df4e01.
//
// Solidity: function getBallotMember(uint256 _id) view returns(address oldStakerAddress, address newStakerAddress, address newVoterAddress, address newRewardAddress, bytes newNodeName, bytes newNodeId, bytes newNodeIp, uint256 newNodePort, uint256 lockAmount)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetBallotMember(_id *big.Int) (struct {
	OldStakerAddress common.Address
	NewStakerAddress common.Address
	NewVoterAddress  common.Address
	NewRewardAddress common.Address
	NewNodeName      []byte
	NewNodeId        []byte
	NewNodeIp        []byte
	NewNodePort      *big.Int
	LockAmount       *big.Int
}, error) {
	return _BallotStorageImp.Contract.GetBallotMember(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotPeriod is a free data retrieval call binding the contract method 0x09970688.
//
// Solidity: function getBallotPeriod(uint256 _id) view returns(uint256 startTime, uint256 endTime, uint256 duration)
func (_BallotStorageImp *BallotStorageImpCaller) GetBallotPeriod(opts *bind.CallOpts, _id *big.Int) (struct {
	StartTime *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getBallotPeriod", _id)

	outstruct := new(struct {
		StartTime *big.Int
		EndTime   *big.Int
		Duration  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBallotPeriod is a free data retrieval call binding the contract method 0x09970688.
//
// Solidity: function getBallotPeriod(uint256 _id) view returns(uint256 startTime, uint256 endTime, uint256 duration)
func (_BallotStorageImp *BallotStorageImpSession) GetBallotPeriod(_id *big.Int) (struct {
	StartTime *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _BallotStorageImp.Contract.GetBallotPeriod(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotPeriod is a free data retrieval call binding the contract method 0x09970688.
//
// Solidity: function getBallotPeriod(uint256 _id) view returns(uint256 startTime, uint256 endTime, uint256 duration)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetBallotPeriod(_id *big.Int) (struct {
	StartTime *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _BallotStorageImp.Contract.GetBallotPeriod(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotState is a free data retrieval call binding the contract method 0x688ca5b2.
//
// Solidity: function getBallotState(uint256 _id) view returns(uint256 ballotType, uint256 state, bool isFinalized)
func (_BallotStorageImp *BallotStorageImpCaller) GetBallotState(opts *bind.CallOpts, _id *big.Int) (struct {
	BallotType  *big.Int
	State       *big.Int
	IsFinalized bool
}, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getBallotState", _id)

	outstruct := new(struct {
		BallotType  *big.Int
		State       *big.Int
		IsFinalized bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BallotType = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.State = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsFinalized = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetBallotState is a free data retrieval call binding the contract method 0x688ca5b2.
//
// Solidity: function getBallotState(uint256 _id) view returns(uint256 ballotType, uint256 state, bool isFinalized)
func (_BallotStorageImp *BallotStorageImpSession) GetBallotState(_id *big.Int) (struct {
	BallotType  *big.Int
	State       *big.Int
	IsFinalized bool
}, error) {
	return _BallotStorageImp.Contract.GetBallotState(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotState is a free data retrieval call binding the contract method 0x688ca5b2.
//
// Solidity: function getBallotState(uint256 _id) view returns(uint256 ballotType, uint256 state, bool isFinalized)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetBallotState(_id *big.Int) (struct {
	BallotType  *big.Int
	State       *big.Int
	IsFinalized bool
}, error) {
	return _BallotStorageImp.Contract.GetBallotState(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotVariable is a free data retrieval call binding the contract method 0x1d940da2.
//
// Solidity: function getBallotVariable(uint256 _id) view returns(bytes32 envVariableName, uint256 envVariableType, bytes envVariableValue)
func (_BallotStorageImp *BallotStorageImpCaller) GetBallotVariable(opts *bind.CallOpts, _id *big.Int) (struct {
	EnvVariableName  [32]byte
	EnvVariableType  *big.Int
	EnvVariableValue []byte
}, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getBallotVariable", _id)

	outstruct := new(struct {
		EnvVariableName  [32]byte
		EnvVariableType  *big.Int
		EnvVariableValue []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EnvVariableName = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.EnvVariableType = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EnvVariableValue = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetBallotVariable is a free data retrieval call binding the contract method 0x1d940da2.
//
// Solidity: function getBallotVariable(uint256 _id) view returns(bytes32 envVariableName, uint256 envVariableType, bytes envVariableValue)
func (_BallotStorageImp *BallotStorageImpSession) GetBallotVariable(_id *big.Int) (struct {
	EnvVariableName  [32]byte
	EnvVariableType  *big.Int
	EnvVariableValue []byte
}, error) {
	return _BallotStorageImp.Contract.GetBallotVariable(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotVariable is a free data retrieval call binding the contract method 0x1d940da2.
//
// Solidity: function getBallotVariable(uint256 _id) view returns(bytes32 envVariableName, uint256 envVariableType, bytes envVariableValue)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetBallotVariable(_id *big.Int) (struct {
	EnvVariableName  [32]byte
	EnvVariableType  *big.Int
	EnvVariableValue []byte
}, error) {
	return _BallotStorageImp.Contract.GetBallotVariable(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotVotingInfo is a free data retrieval call binding the contract method 0x56ba988e.
//
// Solidity: function getBallotVotingInfo(uint256 _id) view returns(uint256 totalVoters, uint256 powerOfAccepts, uint256 powerOfRejects)
func (_BallotStorageImp *BallotStorageImpCaller) GetBallotVotingInfo(opts *bind.CallOpts, _id *big.Int) (struct {
	TotalVoters    *big.Int
	PowerOfAccepts *big.Int
	PowerOfRejects *big.Int
}, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getBallotVotingInfo", _id)

	outstruct := new(struct {
		TotalVoters    *big.Int
		PowerOfAccepts *big.Int
		PowerOfRejects *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalVoters = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PowerOfAccepts = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PowerOfRejects = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBallotVotingInfo is a free data retrieval call binding the contract method 0x56ba988e.
//
// Solidity: function getBallotVotingInfo(uint256 _id) view returns(uint256 totalVoters, uint256 powerOfAccepts, uint256 powerOfRejects)
func (_BallotStorageImp *BallotStorageImpSession) GetBallotVotingInfo(_id *big.Int) (struct {
	TotalVoters    *big.Int
	PowerOfAccepts *big.Int
	PowerOfRejects *big.Int
}, error) {
	return _BallotStorageImp.Contract.GetBallotVotingInfo(&_BallotStorageImp.CallOpts, _id)
}

// GetBallotVotingInfo is a free data retrieval call binding the contract method 0x56ba988e.
//
// Solidity: function getBallotVotingInfo(uint256 _id) view returns(uint256 totalVoters, uint256 powerOfAccepts, uint256 powerOfRejects)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetBallotVotingInfo(_id *big.Int) (struct {
	TotalVoters    *big.Int
	PowerOfAccepts *big.Int
	PowerOfRejects *big.Int
}, error) {
	return _BallotStorageImp.Contract.GetBallotVotingInfo(&_BallotStorageImp.CallOpts, _id)
}

// GetMaxVotingDuration is a free data retrieval call binding the contract method 0xce04b9d4.
//
// Solidity: function getMaxVotingDuration() view returns(uint256)
func (_BallotStorageImp *BallotStorageImpCaller) GetMaxVotingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getMaxVotingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxVotingDuration is a free data retrieval call binding the contract method 0xce04b9d4.
//
// Solidity: function getMaxVotingDuration() view returns(uint256)
func (_BallotStorageImp *BallotStorageImpSession) GetMaxVotingDuration() (*big.Int, error) {
	return _BallotStorageImp.Contract.GetMaxVotingDuration(&_BallotStorageImp.CallOpts)
}

// GetMaxVotingDuration is a free data retrieval call binding the contract method 0xce04b9d4.
//
// Solidity: function getMaxVotingDuration() view returns(uint256)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetMaxVotingDuration() (*big.Int, error) {
	return _BallotStorageImp.Contract.GetMaxVotingDuration(&_BallotStorageImp.CallOpts)
}

// GetMinVotingDuration is a free data retrieval call binding the contract method 0x1c150171.
//
// Solidity: function getMinVotingDuration() view returns(uint256)
func (_BallotStorageImp *BallotStorageImpCaller) GetMinVotingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getMinVotingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinVotingDuration is a free data retrieval call binding the contract method 0x1c150171.
//
// Solidity: function getMinVotingDuration() view returns(uint256)
func (_BallotStorageImp *BallotStorageImpSession) GetMinVotingDuration() (*big.Int, error) {
	return _BallotStorageImp.Contract.GetMinVotingDuration(&_BallotStorageImp.CallOpts)
}

// GetMinVotingDuration is a free data retrieval call binding the contract method 0x1c150171.
//
// Solidity: function getMinVotingDuration() view returns(uint256)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetMinVotingDuration() (*big.Int, error) {
	return _BallotStorageImp.Contract.GetMinVotingDuration(&_BallotStorageImp.CallOpts)
}

// GetPreviousBallotStorage is a free data retrieval call binding the contract method 0xb23c676c.
//
// Solidity: function getPreviousBallotStorage() view returns(address)
func (_BallotStorageImp *BallotStorageImpCaller) GetPreviousBallotStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getPreviousBallotStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPreviousBallotStorage is a free data retrieval call binding the contract method 0xb23c676c.
//
// Solidity: function getPreviousBallotStorage() view returns(address)
func (_BallotStorageImp *BallotStorageImpSession) GetPreviousBallotStorage() (common.Address, error) {
	return _BallotStorageImp.Contract.GetPreviousBallotStorage(&_BallotStorageImp.CallOpts)
}

// GetPreviousBallotStorage is a free data retrieval call binding the contract method 0xb23c676c.
//
// Solidity: function getPreviousBallotStorage() view returns(address)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetPreviousBallotStorage() (common.Address, error) {
	return _BallotStorageImp.Contract.GetPreviousBallotStorage(&_BallotStorageImp.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint256)
func (_BallotStorageImp *BallotStorageImpCaller) GetTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint256)
func (_BallotStorageImp *BallotStorageImpSession) GetTime() (*big.Int, error) {
	return _BallotStorageImp.Contract.GetTime(&_BallotStorageImp.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint256)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetTime() (*big.Int, error) {
	return _BallotStorageImp.Contract.GetTime(&_BallotStorageImp.CallOpts)
}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 _voteId) view returns(uint256 voteId, uint256 ballotId, address voter, uint256 decision, uint256 power, uint256 time)
func (_BallotStorageImp *BallotStorageImpCaller) GetVote(opts *bind.CallOpts, _voteId *big.Int) (struct {
	VoteId   *big.Int
	BallotId *big.Int
	Voter    common.Address
	Decision *big.Int
	Power    *big.Int
	Time     *big.Int
}, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "getVote", _voteId)

	outstruct := new(struct {
		VoteId   *big.Int
		BallotId *big.Int
		Voter    common.Address
		Decision *big.Int
		Power    *big.Int
		Time     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VoteId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BallotId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Voter = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Decision = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Power = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 _voteId) view returns(uint256 voteId, uint256 ballotId, address voter, uint256 decision, uint256 power, uint256 time)
func (_BallotStorageImp *BallotStorageImpSession) GetVote(_voteId *big.Int) (struct {
	VoteId   *big.Int
	BallotId *big.Int
	Voter    common.Address
	Decision *big.Int
	Power    *big.Int
	Time     *big.Int
}, error) {
	return _BallotStorageImp.Contract.GetVote(&_BallotStorageImp.CallOpts, _voteId)
}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 _voteId) view returns(uint256 voteId, uint256 ballotId, address voter, uint256 decision, uint256 power, uint256 time)
func (_BallotStorageImp *BallotStorageImpCallerSession) GetVote(_voteId *big.Int) (struct {
	VoteId   *big.Int
	BallotId *big.Int
	Voter    common.Address
	Decision *big.Int
	Power    *big.Int
	Time     *big.Int
}, error) {
	return _BallotStorageImp.Contract.GetVote(&_BallotStorageImp.CallOpts, _voteId)
}

// HasAlreadyVoted is a free data retrieval call binding the contract method 0xf680e555.
//
// Solidity: function hasAlreadyVoted(uint56 _ballotId, address _voter) view returns(bool)
func (_BallotStorageImp *BallotStorageImpCaller) HasAlreadyVoted(opts *bind.CallOpts, _ballotId *big.Int, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "hasAlreadyVoted", _ballotId, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAlreadyVoted is a free data retrieval call binding the contract method 0xf680e555.
//
// Solidity: function hasAlreadyVoted(uint56 _ballotId, address _voter) view returns(bool)
func (_BallotStorageImp *BallotStorageImpSession) HasAlreadyVoted(_ballotId *big.Int, _voter common.Address) (bool, error) {
	return _BallotStorageImp.Contract.HasAlreadyVoted(&_BallotStorageImp.CallOpts, _ballotId, _voter)
}

// HasAlreadyVoted is a free data retrieval call binding the contract method 0xf680e555.
//
// Solidity: function hasAlreadyVoted(uint56 _ballotId, address _voter) view returns(bool)
func (_BallotStorageImp *BallotStorageImpCallerSession) HasAlreadyVoted(_ballotId *big.Int, _voter common.Address) (bool, error) {
	return _BallotStorageImp.Contract.HasAlreadyVoted(&_BallotStorageImp.CallOpts, _ballotId, _voter)
}

// IsDisabled is a free data retrieval call binding the contract method 0x6c57f5a9.
//
// Solidity: function isDisabled() view returns(bool)
func (_BallotStorageImp *BallotStorageImpCaller) IsDisabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "isDisabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDisabled is a free data retrieval call binding the contract method 0x6c57f5a9.
//
// Solidity: function isDisabled() view returns(bool)
func (_BallotStorageImp *BallotStorageImpSession) IsDisabled() (bool, error) {
	return _BallotStorageImp.Contract.IsDisabled(&_BallotStorageImp.CallOpts)
}

// IsDisabled is a free data retrieval call binding the contract method 0x6c57f5a9.
//
// Solidity: function isDisabled() view returns(bool)
func (_BallotStorageImp *BallotStorageImpCallerSession) IsDisabled() (bool, error) {
	return _BallotStorageImp.Contract.IsDisabled(&_BallotStorageImp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BallotStorageImp *BallotStorageImpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BallotStorageImp *BallotStorageImpSession) Owner() (common.Address, error) {
	return _BallotStorageImp.Contract.Owner(&_BallotStorageImp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BallotStorageImp *BallotStorageImpCallerSession) Owner() (common.Address, error) {
	return _BallotStorageImp.Contract.Owner(&_BallotStorageImp.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpSession) ProxiableUUID() ([32]byte, error) {
	return _BallotStorageImp.Contract.ProxiableUUID(&_BallotStorageImp.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BallotStorageImp *BallotStorageImpCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BallotStorageImp.Contract.ProxiableUUID(&_BallotStorageImp.CallOpts)
}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_BallotStorageImp *BallotStorageImpCaller) Reg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BallotStorageImp.contract.Call(opts, &out, "reg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_BallotStorageImp *BallotStorageImpSession) Reg() (common.Address, error) {
	return _BallotStorageImp.Contract.Reg(&_BallotStorageImp.CallOpts)
}

// Reg is a free data retrieval call binding the contract method 0x738fdd1a.
//
// Solidity: function reg() view returns(address)
func (_BallotStorageImp *BallotStorageImpCallerSession) Reg() (common.Address, error) {
	return _BallotStorageImp.Contract.Reg(&_BallotStorageImp.CallOpts)
}

// CancelBallot is a paid mutator transaction binding the contract method 0x155ca224.
//
// Solidity: function cancelBallot(uint256 _ballotId) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) CancelBallot(opts *bind.TransactOpts, _ballotId *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "cancelBallot", _ballotId)
}

// CancelBallot is a paid mutator transaction binding the contract method 0x155ca224.
//
// Solidity: function cancelBallot(uint256 _ballotId) returns()
func (_BallotStorageImp *BallotStorageImpSession) CancelBallot(_ballotId *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.CancelBallot(&_BallotStorageImp.TransactOpts, _ballotId)
}

// CancelBallot is a paid mutator transaction binding the contract method 0x155ca224.
//
// Solidity: function cancelBallot(uint256 _ballotId) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) CancelBallot(_ballotId *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.CancelBallot(&_BallotStorageImp.TransactOpts, _ballotId)
}

// CreateBallotForAddress is a paid mutator transaction binding the contract method 0x0a3a63fe.
//
// Solidity: function createBallotForAddress(uint256 _id, uint256 _ballotType, uint256 _duration, address _creator, address _newGovernanceAddress) returns(uint256)
func (_BallotStorageImp *BallotStorageImpTransactor) CreateBallotForAddress(opts *bind.TransactOpts, _id *big.Int, _ballotType *big.Int, _duration *big.Int, _creator common.Address, _newGovernanceAddress common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "createBallotForAddress", _id, _ballotType, _duration, _creator, _newGovernanceAddress)
}

// CreateBallotForAddress is a paid mutator transaction binding the contract method 0x0a3a63fe.
//
// Solidity: function createBallotForAddress(uint256 _id, uint256 _ballotType, uint256 _duration, address _creator, address _newGovernanceAddress) returns(uint256)
func (_BallotStorageImp *BallotStorageImpSession) CreateBallotForAddress(_id *big.Int, _ballotType *big.Int, _duration *big.Int, _creator common.Address, _newGovernanceAddress common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.CreateBallotForAddress(&_BallotStorageImp.TransactOpts, _id, _ballotType, _duration, _creator, _newGovernanceAddress)
}

// CreateBallotForAddress is a paid mutator transaction binding the contract method 0x0a3a63fe.
//
// Solidity: function createBallotForAddress(uint256 _id, uint256 _ballotType, uint256 _duration, address _creator, address _newGovernanceAddress) returns(uint256)
func (_BallotStorageImp *BallotStorageImpTransactorSession) CreateBallotForAddress(_id *big.Int, _ballotType *big.Int, _duration *big.Int, _creator common.Address, _newGovernanceAddress common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.CreateBallotForAddress(&_BallotStorageImp.TransactOpts, _id, _ballotType, _duration, _creator, _newGovernanceAddress)
}

// CreateBallotForExit is a paid mutator transaction binding the contract method 0x22640859.
//
// Solidity: function createBallotForExit(uint256 _id, uint256 _unlockAmount, uint256 _slashing) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) CreateBallotForExit(opts *bind.TransactOpts, _id *big.Int, _unlockAmount *big.Int, _slashing *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "createBallotForExit", _id, _unlockAmount, _slashing)
}

// CreateBallotForExit is a paid mutator transaction binding the contract method 0x22640859.
//
// Solidity: function createBallotForExit(uint256 _id, uint256 _unlockAmount, uint256 _slashing) returns()
func (_BallotStorageImp *BallotStorageImpSession) CreateBallotForExit(_id *big.Int, _unlockAmount *big.Int, _slashing *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.CreateBallotForExit(&_BallotStorageImp.TransactOpts, _id, _unlockAmount, _slashing)
}

// CreateBallotForExit is a paid mutator transaction binding the contract method 0x22640859.
//
// Solidity: function createBallotForExit(uint256 _id, uint256 _unlockAmount, uint256 _slashing) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) CreateBallotForExit(_id *big.Int, _unlockAmount *big.Int, _slashing *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.CreateBallotForExit(&_BallotStorageImp.TransactOpts, _id, _unlockAmount, _slashing)
}

// CreateBallotForMember is a paid mutator transaction binding the contract method 0xdaacbb95.
//
// Solidity: function createBallotForMember(uint256 _id, uint256 _ballotType, uint256 _duration, address _creator, address _oldStakerAddress, address _newStakerAddress, address _newVoterAddress, address _newRewardAddress, bytes _newNodeName, bytes _newNodeId, bytes _newNodeIp, uint256 _newNodePort) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) CreateBallotForMember(opts *bind.TransactOpts, _id *big.Int, _ballotType *big.Int, _duration *big.Int, _creator common.Address, _oldStakerAddress common.Address, _newStakerAddress common.Address, _newVoterAddress common.Address, _newRewardAddress common.Address, _newNodeName []byte, _newNodeId []byte, _newNodeIp []byte, _newNodePort *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "createBallotForMember", _id, _ballotType, _duration, _creator, _oldStakerAddress, _newStakerAddress, _newVoterAddress, _newRewardAddress, _newNodeName, _newNodeId, _newNodeIp, _newNodePort)
}

// CreateBallotForMember is a paid mutator transaction binding the contract method 0xdaacbb95.
//
// Solidity: function createBallotForMember(uint256 _id, uint256 _ballotType, uint256 _duration, address _creator, address _oldStakerAddress, address _newStakerAddress, address _newVoterAddress, address _newRewardAddress, bytes _newNodeName, bytes _newNodeId, bytes _newNodeIp, uint256 _newNodePort) returns()
func (_BallotStorageImp *BallotStorageImpSession) CreateBallotForMember(_id *big.Int, _ballotType *big.Int, _duration *big.Int, _creator common.Address, _oldStakerAddress common.Address, _newStakerAddress common.Address, _newVoterAddress common.Address, _newRewardAddress common.Address, _newNodeName []byte, _newNodeId []byte, _newNodeIp []byte, _newNodePort *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.CreateBallotForMember(&_BallotStorageImp.TransactOpts, _id, _ballotType, _duration, _creator, _oldStakerAddress, _newStakerAddress, _newVoterAddress, _newRewardAddress, _newNodeName, _newNodeId, _newNodeIp, _newNodePort)
}

// CreateBallotForMember is a paid mutator transaction binding the contract method 0xdaacbb95.
//
// Solidity: function createBallotForMember(uint256 _id, uint256 _ballotType, uint256 _duration, address _creator, address _oldStakerAddress, address _newStakerAddress, address _newVoterAddress, address _newRewardAddress, bytes _newNodeName, bytes _newNodeId, bytes _newNodeIp, uint256 _newNodePort) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) CreateBallotForMember(_id *big.Int, _ballotType *big.Int, _duration *big.Int, _creator common.Address, _oldStakerAddress common.Address, _newStakerAddress common.Address, _newVoterAddress common.Address, _newRewardAddress common.Address, _newNodeName []byte, _newNodeId []byte, _newNodeIp []byte, _newNodePort *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.CreateBallotForMember(&_BallotStorageImp.TransactOpts, _id, _ballotType, _duration, _creator, _oldStakerAddress, _newStakerAddress, _newVoterAddress, _newRewardAddress, _newNodeName, _newNodeId, _newNodeIp, _newNodePort)
}

// CreateBallotForVariable is a paid mutator transaction binding the contract method 0x4a57823e.
//
// Solidity: function createBallotForVariable(uint256 _id, uint256 _ballotType, uint256 _duration, address _creator, bytes32 _envVariableName, uint256 _envVariableType, bytes _envVariableValue) returns(uint256)
func (_BallotStorageImp *BallotStorageImpTransactor) CreateBallotForVariable(opts *bind.TransactOpts, _id *big.Int, _ballotType *big.Int, _duration *big.Int, _creator common.Address, _envVariableName [32]byte, _envVariableType *big.Int, _envVariableValue []byte) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "createBallotForVariable", _id, _ballotType, _duration, _creator, _envVariableName, _envVariableType, _envVariableValue)
}

// CreateBallotForVariable is a paid mutator transaction binding the contract method 0x4a57823e.
//
// Solidity: function createBallotForVariable(uint256 _id, uint256 _ballotType, uint256 _duration, address _creator, bytes32 _envVariableName, uint256 _envVariableType, bytes _envVariableValue) returns(uint256)
func (_BallotStorageImp *BallotStorageImpSession) CreateBallotForVariable(_id *big.Int, _ballotType *big.Int, _duration *big.Int, _creator common.Address, _envVariableName [32]byte, _envVariableType *big.Int, _envVariableValue []byte) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.CreateBallotForVariable(&_BallotStorageImp.TransactOpts, _id, _ballotType, _duration, _creator, _envVariableName, _envVariableType, _envVariableValue)
}

// CreateBallotForVariable is a paid mutator transaction binding the contract method 0x4a57823e.
//
// Solidity: function createBallotForVariable(uint256 _id, uint256 _ballotType, uint256 _duration, address _creator, bytes32 _envVariableName, uint256 _envVariableType, bytes _envVariableValue) returns(uint256)
func (_BallotStorageImp *BallotStorageImpTransactorSession) CreateBallotForVariable(_id *big.Int, _ballotType *big.Int, _duration *big.Int, _creator common.Address, _envVariableName [32]byte, _envVariableType *big.Int, _envVariableValue []byte) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.CreateBallotForVariable(&_BallotStorageImp.TransactOpts, _id, _ballotType, _duration, _creator, _envVariableName, _envVariableType, _envVariableValue)
}

// CreateVote is a paid mutator transaction binding the contract method 0x96462b9c.
//
// Solidity: function createVote(uint256 _voteId, uint256 _ballotId, address _voter, uint256 _decision, uint256 _power) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) CreateVote(opts *bind.TransactOpts, _voteId *big.Int, _ballotId *big.Int, _voter common.Address, _decision *big.Int, _power *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "createVote", _voteId, _ballotId, _voter, _decision, _power)
}

// CreateVote is a paid mutator transaction binding the contract method 0x96462b9c.
//
// Solidity: function createVote(uint256 _voteId, uint256 _ballotId, address _voter, uint256 _decision, uint256 _power) returns()
func (_BallotStorageImp *BallotStorageImpSession) CreateVote(_voteId *big.Int, _ballotId *big.Int, _voter common.Address, _decision *big.Int, _power *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.CreateVote(&_BallotStorageImp.TransactOpts, _voteId, _ballotId, _voter, _decision, _power)
}

// CreateVote is a paid mutator transaction binding the contract method 0x96462b9c.
//
// Solidity: function createVote(uint256 _voteId, uint256 _ballotId, address _voter, uint256 _decision, uint256 _power) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) CreateVote(_voteId *big.Int, _ballotId *big.Int, _voter common.Address, _decision *big.Int, _power *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.CreateVote(&_BallotStorageImp.TransactOpts, _voteId, _ballotId, _voter, _decision, _power)
}

// FinalizeBallot is a paid mutator transaction binding the contract method 0xa91e59ba.
//
// Solidity: function finalizeBallot(uint256 _ballotId, uint256 _ballotState) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) FinalizeBallot(opts *bind.TransactOpts, _ballotId *big.Int, _ballotState *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "finalizeBallot", _ballotId, _ballotState)
}

// FinalizeBallot is a paid mutator transaction binding the contract method 0xa91e59ba.
//
// Solidity: function finalizeBallot(uint256 _ballotId, uint256 _ballotState) returns()
func (_BallotStorageImp *BallotStorageImpSession) FinalizeBallot(_ballotId *big.Int, _ballotState *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.FinalizeBallot(&_BallotStorageImp.TransactOpts, _ballotId, _ballotState)
}

// FinalizeBallot is a paid mutator transaction binding the contract method 0xa91e59ba.
//
// Solidity: function finalizeBallot(uint256 _ballotId, uint256 _ballotState) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) FinalizeBallot(_ballotId *big.Int, _ballotState *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.FinalizeBallot(&_BallotStorageImp.TransactOpts, _ballotId, _ballotState)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_BallotStorageImp *BallotStorageImpSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.Initialize(&_BallotStorageImp.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.Initialize(&_BallotStorageImp.TransactOpts, _registry)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BallotStorageImp *BallotStorageImpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BallotStorageImp *BallotStorageImpSession) RenounceOwnership() (*types.Transaction, error) {
	return _BallotStorageImp.Contract.RenounceOwnership(&_BallotStorageImp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BallotStorageImp.Contract.RenounceOwnership(&_BallotStorageImp.TransactOpts)
}

// SetPreviousBallotStorage is a paid mutator transaction binding the contract method 0x2a74f38c.
//
// Solidity: function setPreviousBallotStorage(address _address) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) SetPreviousBallotStorage(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "setPreviousBallotStorage", _address)
}

// SetPreviousBallotStorage is a paid mutator transaction binding the contract method 0x2a74f38c.
//
// Solidity: function setPreviousBallotStorage(address _address) returns()
func (_BallotStorageImp *BallotStorageImpSession) SetPreviousBallotStorage(_address common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.SetPreviousBallotStorage(&_BallotStorageImp.TransactOpts, _address)
}

// SetPreviousBallotStorage is a paid mutator transaction binding the contract method 0x2a74f38c.
//
// Solidity: function setPreviousBallotStorage(address _address) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) SetPreviousBallotStorage(_address common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.SetPreviousBallotStorage(&_BallotStorageImp.TransactOpts, _address)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) SetRegistry(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "setRegistry", _addr)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_BallotStorageImp *BallotStorageImpSession) SetRegistry(_addr common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.SetRegistry(&_BallotStorageImp.TransactOpts, _addr)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _addr) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) SetRegistry(_addr common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.SetRegistry(&_BallotStorageImp.TransactOpts, _addr)
}

// StartBallot is a paid mutator transaction binding the contract method 0xc0b6f186.
//
// Solidity: function startBallot(uint256 _ballotId, uint256 _startTime, uint256 _endTime) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) StartBallot(opts *bind.TransactOpts, _ballotId *big.Int, _startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "startBallot", _ballotId, _startTime, _endTime)
}

// StartBallot is a paid mutator transaction binding the contract method 0xc0b6f186.
//
// Solidity: function startBallot(uint256 _ballotId, uint256 _startTime, uint256 _endTime) returns()
func (_BallotStorageImp *BallotStorageImpSession) StartBallot(_ballotId *big.Int, _startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.StartBallot(&_BallotStorageImp.TransactOpts, _ballotId, _startTime, _endTime)
}

// StartBallot is a paid mutator transaction binding the contract method 0xc0b6f186.
//
// Solidity: function startBallot(uint256 _ballotId, uint256 _startTime, uint256 _endTime) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) StartBallot(_ballotId *big.Int, _startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.StartBallot(&_BallotStorageImp.TransactOpts, _ballotId, _startTime, _endTime)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BallotStorageImp *BallotStorageImpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.TransferOwnership(&_BallotStorageImp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.TransferOwnership(&_BallotStorageImp.TransactOpts, newOwner)
}

// UpdateBallotDuration is a paid mutator transaction binding the contract method 0x656bc633.
//
// Solidity: function updateBallotDuration(uint256 _ballotId, uint256 _duration) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) UpdateBallotDuration(opts *bind.TransactOpts, _ballotId *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "updateBallotDuration", _ballotId, _duration)
}

// UpdateBallotDuration is a paid mutator transaction binding the contract method 0x656bc633.
//
// Solidity: function updateBallotDuration(uint256 _ballotId, uint256 _duration) returns()
func (_BallotStorageImp *BallotStorageImpSession) UpdateBallotDuration(_ballotId *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.UpdateBallotDuration(&_BallotStorageImp.TransactOpts, _ballotId, _duration)
}

// UpdateBallotDuration is a paid mutator transaction binding the contract method 0x656bc633.
//
// Solidity: function updateBallotDuration(uint256 _ballotId, uint256 _duration) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) UpdateBallotDuration(_ballotId *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.UpdateBallotDuration(&_BallotStorageImp.TransactOpts, _ballotId, _duration)
}

// UpdateBallotMemberLockAmount is a paid mutator transaction binding the contract method 0x72d0ec92.
//
// Solidity: function updateBallotMemberLockAmount(uint256 _ballotId, uint256 _lockAmount) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) UpdateBallotMemberLockAmount(opts *bind.TransactOpts, _ballotId *big.Int, _lockAmount *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "updateBallotMemberLockAmount", _ballotId, _lockAmount)
}

// UpdateBallotMemberLockAmount is a paid mutator transaction binding the contract method 0x72d0ec92.
//
// Solidity: function updateBallotMemberLockAmount(uint256 _ballotId, uint256 _lockAmount) returns()
func (_BallotStorageImp *BallotStorageImpSession) UpdateBallotMemberLockAmount(_ballotId *big.Int, _lockAmount *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.UpdateBallotMemberLockAmount(&_BallotStorageImp.TransactOpts, _ballotId, _lockAmount)
}

// UpdateBallotMemberLockAmount is a paid mutator transaction binding the contract method 0x72d0ec92.
//
// Solidity: function updateBallotMemberLockAmount(uint256 _ballotId, uint256 _lockAmount) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) UpdateBallotMemberLockAmount(_ballotId *big.Int, _lockAmount *big.Int) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.UpdateBallotMemberLockAmount(&_BallotStorageImp.TransactOpts, _ballotId, _lockAmount)
}

// UpdateBallotMemo is a paid mutator transaction binding the contract method 0xbce0dbc1.
//
// Solidity: function updateBallotMemo(uint256 _ballotId, bytes _memo) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) UpdateBallotMemo(opts *bind.TransactOpts, _ballotId *big.Int, _memo []byte) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "updateBallotMemo", _ballotId, _memo)
}

// UpdateBallotMemo is a paid mutator transaction binding the contract method 0xbce0dbc1.
//
// Solidity: function updateBallotMemo(uint256 _ballotId, bytes _memo) returns()
func (_BallotStorageImp *BallotStorageImpSession) UpdateBallotMemo(_ballotId *big.Int, _memo []byte) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.UpdateBallotMemo(&_BallotStorageImp.TransactOpts, _ballotId, _memo)
}

// UpdateBallotMemo is a paid mutator transaction binding the contract method 0xbce0dbc1.
//
// Solidity: function updateBallotMemo(uint256 _ballotId, bytes _memo) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) UpdateBallotMemo(_ballotId *big.Int, _memo []byte) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.UpdateBallotMemo(&_BallotStorageImp.TransactOpts, _ballotId, _memo)
}

// UpgradeBallotStorage is a paid mutator transaction binding the contract method 0xf51a88e2.
//
// Solidity: function upgradeBallotStorage(address newImp) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) UpgradeBallotStorage(opts *bind.TransactOpts, newImp common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "upgradeBallotStorage", newImp)
}

// UpgradeBallotStorage is a paid mutator transaction binding the contract method 0xf51a88e2.
//
// Solidity: function upgradeBallotStorage(address newImp) returns()
func (_BallotStorageImp *BallotStorageImpSession) UpgradeBallotStorage(newImp common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.UpgradeBallotStorage(&_BallotStorageImp.TransactOpts, newImp)
}

// UpgradeBallotStorage is a paid mutator transaction binding the contract method 0xf51a88e2.
//
// Solidity: function upgradeBallotStorage(address newImp) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) UpgradeBallotStorage(newImp common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.UpgradeBallotStorage(&_BallotStorageImp.TransactOpts, newImp)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BallotStorageImp *BallotStorageImpTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BallotStorageImp *BallotStorageImpSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.UpgradeTo(&_BallotStorageImp.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.UpgradeTo(&_BallotStorageImp.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BallotStorageImp *BallotStorageImpTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BallotStorageImp.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BallotStorageImp *BallotStorageImpSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.UpgradeToAndCall(&_BallotStorageImp.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BallotStorageImp *BallotStorageImpTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BallotStorageImp.Contract.UpgradeToAndCall(&_BallotStorageImp.TransactOpts, newImplementation, data)
}

// BallotStorageImpAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the BallotStorageImp contract.
type BallotStorageImpAdminChangedIterator struct {
	Event *BallotStorageImpAdminChanged // Event containing the contract specifics and raw log

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
func (it *BallotStorageImpAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageImpAdminChanged)
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
		it.Event = new(BallotStorageImpAdminChanged)
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
func (it *BallotStorageImpAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageImpAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageImpAdminChanged represents a AdminChanged event raised by the BallotStorageImp contract.
type BallotStorageImpAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BallotStorageImp *BallotStorageImpFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BallotStorageImpAdminChangedIterator, error) {

	logs, sub, err := _BallotStorageImp.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpAdminChangedIterator{contract: _BallotStorageImp.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BallotStorageImp *BallotStorageImpFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BallotStorageImpAdminChanged) (event.Subscription, error) {

	logs, sub, err := _BallotStorageImp.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageImpAdminChanged)
				if err := _BallotStorageImp.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_BallotStorageImp *BallotStorageImpFilterer) ParseAdminChanged(log types.Log) (*BallotStorageImpAdminChanged, error) {
	event := new(BallotStorageImpAdminChanged)
	if err := _BallotStorageImp.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageImpBallotCanceledIterator is returned from FilterBallotCanceled and is used to iterate over the raw logs and unpacked data for BallotCanceled events raised by the BallotStorageImp contract.
type BallotStorageImpBallotCanceledIterator struct {
	Event *BallotStorageImpBallotCanceled // Event containing the contract specifics and raw log

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
func (it *BallotStorageImpBallotCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageImpBallotCanceled)
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
		it.Event = new(BallotStorageImpBallotCanceled)
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
func (it *BallotStorageImpBallotCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageImpBallotCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageImpBallotCanceled represents a BallotCanceled event raised by the BallotStorageImp contract.
type BallotStorageImpBallotCanceled struct {
	BallotId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBallotCanceled is a free log retrieval operation binding the contract event 0xd5e541d004c50564e5e05fcbc6be2916c68d817507693dc3774c69dde4ce13dc.
//
// Solidity: event BallotCanceled(uint256 indexed ballotId)
func (_BallotStorageImp *BallotStorageImpFilterer) FilterBallotCanceled(opts *bind.FilterOpts, ballotId []*big.Int) (*BallotStorageImpBallotCanceledIterator, error) {

	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}

	logs, sub, err := _BallotStorageImp.contract.FilterLogs(opts, "BallotCanceled", ballotIdRule)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpBallotCanceledIterator{contract: _BallotStorageImp.contract, event: "BallotCanceled", logs: logs, sub: sub}, nil
}

// WatchBallotCanceled is a free log subscription operation binding the contract event 0xd5e541d004c50564e5e05fcbc6be2916c68d817507693dc3774c69dde4ce13dc.
//
// Solidity: event BallotCanceled(uint256 indexed ballotId)
func (_BallotStorageImp *BallotStorageImpFilterer) WatchBallotCanceled(opts *bind.WatchOpts, sink chan<- *BallotStorageImpBallotCanceled, ballotId []*big.Int) (event.Subscription, error) {

	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}

	logs, sub, err := _BallotStorageImp.contract.WatchLogs(opts, "BallotCanceled", ballotIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageImpBallotCanceled)
				if err := _BallotStorageImp.contract.UnpackLog(event, "BallotCanceled", log); err != nil {
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

// ParseBallotCanceled is a log parse operation binding the contract event 0xd5e541d004c50564e5e05fcbc6be2916c68d817507693dc3774c69dde4ce13dc.
//
// Solidity: event BallotCanceled(uint256 indexed ballotId)
func (_BallotStorageImp *BallotStorageImpFilterer) ParseBallotCanceled(log types.Log) (*BallotStorageImpBallotCanceled, error) {
	event := new(BallotStorageImpBallotCanceled)
	if err := _BallotStorageImp.contract.UnpackLog(event, "BallotCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageImpBallotCreatedIterator is returned from FilterBallotCreated and is used to iterate over the raw logs and unpacked data for BallotCreated events raised by the BallotStorageImp contract.
type BallotStorageImpBallotCreatedIterator struct {
	Event *BallotStorageImpBallotCreated // Event containing the contract specifics and raw log

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
func (it *BallotStorageImpBallotCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageImpBallotCreated)
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
		it.Event = new(BallotStorageImpBallotCreated)
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
func (it *BallotStorageImpBallotCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageImpBallotCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageImpBallotCreated represents a BallotCreated event raised by the BallotStorageImp contract.
type BallotStorageImpBallotCreated struct {
	BallotId   *big.Int
	BallotType *big.Int
	Creator    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBallotCreated is a free log retrieval operation binding the contract event 0xd1ba591c76ef71222e2d30b8277758713cc6eef1de29efaf98a716744ac2420b.
//
// Solidity: event BallotCreated(uint256 indexed ballotId, uint256 indexed ballotType, address indexed creator)
func (_BallotStorageImp *BallotStorageImpFilterer) FilterBallotCreated(opts *bind.FilterOpts, ballotId []*big.Int, ballotType []*big.Int, creator []common.Address) (*BallotStorageImpBallotCreatedIterator, error) {

	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}
	var ballotTypeRule []interface{}
	for _, ballotTypeItem := range ballotType {
		ballotTypeRule = append(ballotTypeRule, ballotTypeItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _BallotStorageImp.contract.FilterLogs(opts, "BallotCreated", ballotIdRule, ballotTypeRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpBallotCreatedIterator{contract: _BallotStorageImp.contract, event: "BallotCreated", logs: logs, sub: sub}, nil
}

// WatchBallotCreated is a free log subscription operation binding the contract event 0xd1ba591c76ef71222e2d30b8277758713cc6eef1de29efaf98a716744ac2420b.
//
// Solidity: event BallotCreated(uint256 indexed ballotId, uint256 indexed ballotType, address indexed creator)
func (_BallotStorageImp *BallotStorageImpFilterer) WatchBallotCreated(opts *bind.WatchOpts, sink chan<- *BallotStorageImpBallotCreated, ballotId []*big.Int, ballotType []*big.Int, creator []common.Address) (event.Subscription, error) {

	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}
	var ballotTypeRule []interface{}
	for _, ballotTypeItem := range ballotType {
		ballotTypeRule = append(ballotTypeRule, ballotTypeItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _BallotStorageImp.contract.WatchLogs(opts, "BallotCreated", ballotIdRule, ballotTypeRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageImpBallotCreated)
				if err := _BallotStorageImp.contract.UnpackLog(event, "BallotCreated", log); err != nil {
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

// ParseBallotCreated is a log parse operation binding the contract event 0xd1ba591c76ef71222e2d30b8277758713cc6eef1de29efaf98a716744ac2420b.
//
// Solidity: event BallotCreated(uint256 indexed ballotId, uint256 indexed ballotType, address indexed creator)
func (_BallotStorageImp *BallotStorageImpFilterer) ParseBallotCreated(log types.Log) (*BallotStorageImpBallotCreated, error) {
	event := new(BallotStorageImpBallotCreated)
	if err := _BallotStorageImp.contract.UnpackLog(event, "BallotCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageImpBallotFinalizedIterator is returned from FilterBallotFinalized and is used to iterate over the raw logs and unpacked data for BallotFinalized events raised by the BallotStorageImp contract.
type BallotStorageImpBallotFinalizedIterator struct {
	Event *BallotStorageImpBallotFinalized // Event containing the contract specifics and raw log

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
func (it *BallotStorageImpBallotFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageImpBallotFinalized)
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
		it.Event = new(BallotStorageImpBallotFinalized)
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
func (it *BallotStorageImpBallotFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageImpBallotFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageImpBallotFinalized represents a BallotFinalized event raised by the BallotStorageImp contract.
type BallotStorageImpBallotFinalized struct {
	BallotId *big.Int
	State    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBallotFinalized is a free log retrieval operation binding the contract event 0xdc921f027328d7238b58d77649ebcbd0c8b1c494c66ba53dfe53e0de65f6dd9f.
//
// Solidity: event BallotFinalized(uint256 indexed ballotId, uint256 state)
func (_BallotStorageImp *BallotStorageImpFilterer) FilterBallotFinalized(opts *bind.FilterOpts, ballotId []*big.Int) (*BallotStorageImpBallotFinalizedIterator, error) {

	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}

	logs, sub, err := _BallotStorageImp.contract.FilterLogs(opts, "BallotFinalized", ballotIdRule)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpBallotFinalizedIterator{contract: _BallotStorageImp.contract, event: "BallotFinalized", logs: logs, sub: sub}, nil
}

// WatchBallotFinalized is a free log subscription operation binding the contract event 0xdc921f027328d7238b58d77649ebcbd0c8b1c494c66ba53dfe53e0de65f6dd9f.
//
// Solidity: event BallotFinalized(uint256 indexed ballotId, uint256 state)
func (_BallotStorageImp *BallotStorageImpFilterer) WatchBallotFinalized(opts *bind.WatchOpts, sink chan<- *BallotStorageImpBallotFinalized, ballotId []*big.Int) (event.Subscription, error) {

	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}

	logs, sub, err := _BallotStorageImp.contract.WatchLogs(opts, "BallotFinalized", ballotIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageImpBallotFinalized)
				if err := _BallotStorageImp.contract.UnpackLog(event, "BallotFinalized", log); err != nil {
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

// ParseBallotFinalized is a log parse operation binding the contract event 0xdc921f027328d7238b58d77649ebcbd0c8b1c494c66ba53dfe53e0de65f6dd9f.
//
// Solidity: event BallotFinalized(uint256 indexed ballotId, uint256 state)
func (_BallotStorageImp *BallotStorageImpFilterer) ParseBallotFinalized(log types.Log) (*BallotStorageImpBallotFinalized, error) {
	event := new(BallotStorageImpBallotFinalized)
	if err := _BallotStorageImp.contract.UnpackLog(event, "BallotFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageImpBallotStartedIterator is returned from FilterBallotStarted and is used to iterate over the raw logs and unpacked data for BallotStarted events raised by the BallotStorageImp contract.
type BallotStorageImpBallotStartedIterator struct {
	Event *BallotStorageImpBallotStarted // Event containing the contract specifics and raw log

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
func (it *BallotStorageImpBallotStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageImpBallotStarted)
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
		it.Event = new(BallotStorageImpBallotStarted)
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
func (it *BallotStorageImpBallotStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageImpBallotStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageImpBallotStarted represents a BallotStarted event raised by the BallotStorageImp contract.
type BallotStorageImpBallotStarted struct {
	BallotId  *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBallotStarted is a free log retrieval operation binding the contract event 0xd9938a514dab5cdce149a77493f694bd70c24a4833a278fd4d86fbdf859099c5.
//
// Solidity: event BallotStarted(uint256 indexed ballotId, uint256 indexed startTime, uint256 indexed endTime)
func (_BallotStorageImp *BallotStorageImpFilterer) FilterBallotStarted(opts *bind.FilterOpts, ballotId []*big.Int, startTime []*big.Int, endTime []*big.Int) (*BallotStorageImpBallotStartedIterator, error) {

	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}
	var startTimeRule []interface{}
	for _, startTimeItem := range startTime {
		startTimeRule = append(startTimeRule, startTimeItem)
	}
	var endTimeRule []interface{}
	for _, endTimeItem := range endTime {
		endTimeRule = append(endTimeRule, endTimeItem)
	}

	logs, sub, err := _BallotStorageImp.contract.FilterLogs(opts, "BallotStarted", ballotIdRule, startTimeRule, endTimeRule)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpBallotStartedIterator{contract: _BallotStorageImp.contract, event: "BallotStarted", logs: logs, sub: sub}, nil
}

// WatchBallotStarted is a free log subscription operation binding the contract event 0xd9938a514dab5cdce149a77493f694bd70c24a4833a278fd4d86fbdf859099c5.
//
// Solidity: event BallotStarted(uint256 indexed ballotId, uint256 indexed startTime, uint256 indexed endTime)
func (_BallotStorageImp *BallotStorageImpFilterer) WatchBallotStarted(opts *bind.WatchOpts, sink chan<- *BallotStorageImpBallotStarted, ballotId []*big.Int, startTime []*big.Int, endTime []*big.Int) (event.Subscription, error) {

	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}
	var startTimeRule []interface{}
	for _, startTimeItem := range startTime {
		startTimeRule = append(startTimeRule, startTimeItem)
	}
	var endTimeRule []interface{}
	for _, endTimeItem := range endTime {
		endTimeRule = append(endTimeRule, endTimeItem)
	}

	logs, sub, err := _BallotStorageImp.contract.WatchLogs(opts, "BallotStarted", ballotIdRule, startTimeRule, endTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageImpBallotStarted)
				if err := _BallotStorageImp.contract.UnpackLog(event, "BallotStarted", log); err != nil {
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

// ParseBallotStarted is a log parse operation binding the contract event 0xd9938a514dab5cdce149a77493f694bd70c24a4833a278fd4d86fbdf859099c5.
//
// Solidity: event BallotStarted(uint256 indexed ballotId, uint256 indexed startTime, uint256 indexed endTime)
func (_BallotStorageImp *BallotStorageImpFilterer) ParseBallotStarted(log types.Log) (*BallotStorageImpBallotStarted, error) {
	event := new(BallotStorageImpBallotStarted)
	if err := _BallotStorageImp.contract.UnpackLog(event, "BallotStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageImpBallotUpdatedIterator is returned from FilterBallotUpdated and is used to iterate over the raw logs and unpacked data for BallotUpdated events raised by the BallotStorageImp contract.
type BallotStorageImpBallotUpdatedIterator struct {
	Event *BallotStorageImpBallotUpdated // Event containing the contract specifics and raw log

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
func (it *BallotStorageImpBallotUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageImpBallotUpdated)
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
		it.Event = new(BallotStorageImpBallotUpdated)
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
func (it *BallotStorageImpBallotUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageImpBallotUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageImpBallotUpdated represents a BallotUpdated event raised by the BallotStorageImp contract.
type BallotStorageImpBallotUpdated struct {
	BallotId  *big.Int
	UpdatedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBallotUpdated is a free log retrieval operation binding the contract event 0xf0855c0e0ad9b8a162b87f2e4e07d4b2a3f0a45126b15ff4a78b217ad19a901a.
//
// Solidity: event BallotUpdated(uint256 indexed ballotId, address indexed updatedBy)
func (_BallotStorageImp *BallotStorageImpFilterer) FilterBallotUpdated(opts *bind.FilterOpts, ballotId []*big.Int, updatedBy []common.Address) (*BallotStorageImpBallotUpdatedIterator, error) {

	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}
	var updatedByRule []interface{}
	for _, updatedByItem := range updatedBy {
		updatedByRule = append(updatedByRule, updatedByItem)
	}

	logs, sub, err := _BallotStorageImp.contract.FilterLogs(opts, "BallotUpdated", ballotIdRule, updatedByRule)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpBallotUpdatedIterator{contract: _BallotStorageImp.contract, event: "BallotUpdated", logs: logs, sub: sub}, nil
}

// WatchBallotUpdated is a free log subscription operation binding the contract event 0xf0855c0e0ad9b8a162b87f2e4e07d4b2a3f0a45126b15ff4a78b217ad19a901a.
//
// Solidity: event BallotUpdated(uint256 indexed ballotId, address indexed updatedBy)
func (_BallotStorageImp *BallotStorageImpFilterer) WatchBallotUpdated(opts *bind.WatchOpts, sink chan<- *BallotStorageImpBallotUpdated, ballotId []*big.Int, updatedBy []common.Address) (event.Subscription, error) {

	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}
	var updatedByRule []interface{}
	for _, updatedByItem := range updatedBy {
		updatedByRule = append(updatedByRule, updatedByItem)
	}

	logs, sub, err := _BallotStorageImp.contract.WatchLogs(opts, "BallotUpdated", ballotIdRule, updatedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageImpBallotUpdated)
				if err := _BallotStorageImp.contract.UnpackLog(event, "BallotUpdated", log); err != nil {
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

// ParseBallotUpdated is a log parse operation binding the contract event 0xf0855c0e0ad9b8a162b87f2e4e07d4b2a3f0a45126b15ff4a78b217ad19a901a.
//
// Solidity: event BallotUpdated(uint256 indexed ballotId, address indexed updatedBy)
func (_BallotStorageImp *BallotStorageImpFilterer) ParseBallotUpdated(log types.Log) (*BallotStorageImpBallotUpdated, error) {
	event := new(BallotStorageImpBallotUpdated)
	if err := _BallotStorageImp.contract.UnpackLog(event, "BallotUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageImpBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the BallotStorageImp contract.
type BallotStorageImpBeaconUpgradedIterator struct {
	Event *BallotStorageImpBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BallotStorageImpBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageImpBeaconUpgraded)
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
		it.Event = new(BallotStorageImpBeaconUpgraded)
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
func (it *BallotStorageImpBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageImpBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageImpBeaconUpgraded represents a BeaconUpgraded event raised by the BallotStorageImp contract.
type BallotStorageImpBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BallotStorageImp *BallotStorageImpFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BallotStorageImpBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BallotStorageImp.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpBeaconUpgradedIterator{contract: _BallotStorageImp.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BallotStorageImp *BallotStorageImpFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BallotStorageImpBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BallotStorageImp.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageImpBeaconUpgraded)
				if err := _BallotStorageImp.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_BallotStorageImp *BallotStorageImpFilterer) ParseBeaconUpgraded(log types.Log) (*BallotStorageImpBeaconUpgraded, error) {
	event := new(BallotStorageImpBeaconUpgraded)
	if err := _BallotStorageImp.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageImpInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BallotStorageImp contract.
type BallotStorageImpInitializedIterator struct {
	Event *BallotStorageImpInitialized // Event containing the contract specifics and raw log

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
func (it *BallotStorageImpInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageImpInitialized)
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
		it.Event = new(BallotStorageImpInitialized)
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
func (it *BallotStorageImpInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageImpInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageImpInitialized represents a Initialized event raised by the BallotStorageImp contract.
type BallotStorageImpInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BallotStorageImp *BallotStorageImpFilterer) FilterInitialized(opts *bind.FilterOpts) (*BallotStorageImpInitializedIterator, error) {

	logs, sub, err := _BallotStorageImp.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpInitializedIterator{contract: _BallotStorageImp.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BallotStorageImp *BallotStorageImpFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BallotStorageImpInitialized) (event.Subscription, error) {

	logs, sub, err := _BallotStorageImp.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageImpInitialized)
				if err := _BallotStorageImp.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BallotStorageImp *BallotStorageImpFilterer) ParseInitialized(log types.Log) (*BallotStorageImpInitialized, error) {
	event := new(BallotStorageImpInitialized)
	if err := _BallotStorageImp.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageImpOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BallotStorageImp contract.
type BallotStorageImpOwnershipTransferredIterator struct {
	Event *BallotStorageImpOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BallotStorageImpOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageImpOwnershipTransferred)
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
		it.Event = new(BallotStorageImpOwnershipTransferred)
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
func (it *BallotStorageImpOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageImpOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageImpOwnershipTransferred represents a OwnershipTransferred event raised by the BallotStorageImp contract.
type BallotStorageImpOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BallotStorageImp *BallotStorageImpFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BallotStorageImpOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BallotStorageImp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpOwnershipTransferredIterator{contract: _BallotStorageImp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BallotStorageImp *BallotStorageImpFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BallotStorageImpOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BallotStorageImp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageImpOwnershipTransferred)
				if err := _BallotStorageImp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BallotStorageImp *BallotStorageImpFilterer) ParseOwnershipTransferred(log types.Log) (*BallotStorageImpOwnershipTransferred, error) {
	event := new(BallotStorageImpOwnershipTransferred)
	if err := _BallotStorageImp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageImpSetPrevBallotStorageIterator is returned from FilterSetPrevBallotStorage and is used to iterate over the raw logs and unpacked data for SetPrevBallotStorage events raised by the BallotStorageImp contract.
type BallotStorageImpSetPrevBallotStorageIterator struct {
	Event *BallotStorageImpSetPrevBallotStorage // Event containing the contract specifics and raw log

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
func (it *BallotStorageImpSetPrevBallotStorageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageImpSetPrevBallotStorage)
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
		it.Event = new(BallotStorageImpSetPrevBallotStorage)
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
func (it *BallotStorageImpSetPrevBallotStorageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageImpSetPrevBallotStorageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageImpSetPrevBallotStorage represents a SetPrevBallotStorage event raised by the BallotStorageImp contract.
type BallotStorageImpSetPrevBallotStorage struct {
	Previous common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPrevBallotStorage is a free log retrieval operation binding the contract event 0x3d809312b6e303291a93b307c7ddbd0960c094f5f0fb4e3ba0758775013edeb3.
//
// Solidity: event SetPrevBallotStorage(address indexed previous)
func (_BallotStorageImp *BallotStorageImpFilterer) FilterSetPrevBallotStorage(opts *bind.FilterOpts, previous []common.Address) (*BallotStorageImpSetPrevBallotStorageIterator, error) {

	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}

	logs, sub, err := _BallotStorageImp.contract.FilterLogs(opts, "SetPrevBallotStorage", previousRule)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpSetPrevBallotStorageIterator{contract: _BallotStorageImp.contract, event: "SetPrevBallotStorage", logs: logs, sub: sub}, nil
}

// WatchSetPrevBallotStorage is a free log subscription operation binding the contract event 0x3d809312b6e303291a93b307c7ddbd0960c094f5f0fb4e3ba0758775013edeb3.
//
// Solidity: event SetPrevBallotStorage(address indexed previous)
func (_BallotStorageImp *BallotStorageImpFilterer) WatchSetPrevBallotStorage(opts *bind.WatchOpts, sink chan<- *BallotStorageImpSetPrevBallotStorage, previous []common.Address) (event.Subscription, error) {

	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}

	logs, sub, err := _BallotStorageImp.contract.WatchLogs(opts, "SetPrevBallotStorage", previousRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageImpSetPrevBallotStorage)
				if err := _BallotStorageImp.contract.UnpackLog(event, "SetPrevBallotStorage", log); err != nil {
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

// ParseSetPrevBallotStorage is a log parse operation binding the contract event 0x3d809312b6e303291a93b307c7ddbd0960c094f5f0fb4e3ba0758775013edeb3.
//
// Solidity: event SetPrevBallotStorage(address indexed previous)
func (_BallotStorageImp *BallotStorageImpFilterer) ParseSetPrevBallotStorage(log types.Log) (*BallotStorageImpSetPrevBallotStorage, error) {
	event := new(BallotStorageImpSetPrevBallotStorage)
	if err := _BallotStorageImp.contract.UnpackLog(event, "SetPrevBallotStorage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageImpSetRegistryIterator is returned from FilterSetRegistry and is used to iterate over the raw logs and unpacked data for SetRegistry events raised by the BallotStorageImp contract.
type BallotStorageImpSetRegistryIterator struct {
	Event *BallotStorageImpSetRegistry // Event containing the contract specifics and raw log

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
func (it *BallotStorageImpSetRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageImpSetRegistry)
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
		it.Event = new(BallotStorageImpSetRegistry)
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
func (it *BallotStorageImpSetRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageImpSetRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageImpSetRegistry represents a SetRegistry event raised by the BallotStorageImp contract.
type BallotStorageImpSetRegistry struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetRegistry is a free log retrieval operation binding the contract event 0x278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd.
//
// Solidity: event SetRegistry(address indexed addr)
func (_BallotStorageImp *BallotStorageImpFilterer) FilterSetRegistry(opts *bind.FilterOpts, addr []common.Address) (*BallotStorageImpSetRegistryIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BallotStorageImp.contract.FilterLogs(opts, "SetRegistry", addrRule)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpSetRegistryIterator{contract: _BallotStorageImp.contract, event: "SetRegistry", logs: logs, sub: sub}, nil
}

// WatchSetRegistry is a free log subscription operation binding the contract event 0x278c70ced5f3e0e5eeb385b5ff9cb735748ba00a625147e66065ed48fc1562cd.
//
// Solidity: event SetRegistry(address indexed addr)
func (_BallotStorageImp *BallotStorageImpFilterer) WatchSetRegistry(opts *bind.WatchOpts, sink chan<- *BallotStorageImpSetRegistry, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BallotStorageImp.contract.WatchLogs(opts, "SetRegistry", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageImpSetRegistry)
				if err := _BallotStorageImp.contract.UnpackLog(event, "SetRegistry", log); err != nil {
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
func (_BallotStorageImp *BallotStorageImpFilterer) ParseSetRegistry(log types.Log) (*BallotStorageImpSetRegistry, error) {
	event := new(BallotStorageImpSetRegistry)
	if err := _BallotStorageImp.contract.UnpackLog(event, "SetRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageImpUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BallotStorageImp contract.
type BallotStorageImpUpgradedIterator struct {
	Event *BallotStorageImpUpgraded // Event containing the contract specifics and raw log

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
func (it *BallotStorageImpUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageImpUpgraded)
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
		it.Event = new(BallotStorageImpUpgraded)
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
func (it *BallotStorageImpUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageImpUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageImpUpgraded represents a Upgraded event raised by the BallotStorageImp contract.
type BallotStorageImpUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BallotStorageImp *BallotStorageImpFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BallotStorageImpUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BallotStorageImp.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpUpgradedIterator{contract: _BallotStorageImp.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BallotStorageImp *BallotStorageImpFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BallotStorageImpUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BallotStorageImp.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageImpUpgraded)
				if err := _BallotStorageImp.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BallotStorageImp *BallotStorageImpFilterer) ParseUpgraded(log types.Log) (*BallotStorageImpUpgraded, error) {
	event := new(BallotStorageImpUpgraded)
	if err := _BallotStorageImp.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BallotStorageImpVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the BallotStorageImp contract.
type BallotStorageImpVotedIterator struct {
	Event *BallotStorageImpVoted // Event containing the contract specifics and raw log

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
func (it *BallotStorageImpVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotStorageImpVoted)
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
		it.Event = new(BallotStorageImpVoted)
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
func (it *BallotStorageImpVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotStorageImpVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotStorageImpVoted represents a Voted event raised by the BallotStorageImp contract.
type BallotStorageImpVoted struct {
	Voteid   *big.Int
	BallotId *big.Int
	Voter    common.Address
	Decision *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x41df84b3b467b06744e40c92613c666324e7c640ce0a41ec06efdf602d367606.
//
// Solidity: event Voted(uint256 indexed voteid, uint256 indexed ballotId, address indexed voter, uint256 decision)
func (_BallotStorageImp *BallotStorageImpFilterer) FilterVoted(opts *bind.FilterOpts, voteid []*big.Int, ballotId []*big.Int, voter []common.Address) (*BallotStorageImpVotedIterator, error) {

	var voteidRule []interface{}
	for _, voteidItem := range voteid {
		voteidRule = append(voteidRule, voteidItem)
	}
	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _BallotStorageImp.contract.FilterLogs(opts, "Voted", voteidRule, ballotIdRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &BallotStorageImpVotedIterator{contract: _BallotStorageImp.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x41df84b3b467b06744e40c92613c666324e7c640ce0a41ec06efdf602d367606.
//
// Solidity: event Voted(uint256 indexed voteid, uint256 indexed ballotId, address indexed voter, uint256 decision)
func (_BallotStorageImp *BallotStorageImpFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *BallotStorageImpVoted, voteid []*big.Int, ballotId []*big.Int, voter []common.Address) (event.Subscription, error) {

	var voteidRule []interface{}
	for _, voteidItem := range voteid {
		voteidRule = append(voteidRule, voteidItem)
	}
	var ballotIdRule []interface{}
	for _, ballotIdItem := range ballotId {
		ballotIdRule = append(ballotIdRule, ballotIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _BallotStorageImp.contract.WatchLogs(opts, "Voted", voteidRule, ballotIdRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotStorageImpVoted)
				if err := _BallotStorageImp.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x41df84b3b467b06744e40c92613c666324e7c640ce0a41ec06efdf602d367606.
//
// Solidity: event Voted(uint256 indexed voteid, uint256 indexed ballotId, address indexed voter, uint256 decision)
func (_BallotStorageImp *BallotStorageImpFilterer) ParseVoted(log types.Log) (*BallotStorageImpVoted, error) {
	event := new(BallotStorageImpVoted)
	if err := _BallotStorageImp.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
