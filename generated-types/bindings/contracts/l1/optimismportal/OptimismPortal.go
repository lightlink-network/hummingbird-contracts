// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimismportal

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
	_ = abi.ConvertType
)

// OptimismPortalMetaData contains all meta data concerning the OptimismPortal contract.
var OptimismPortalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"opaqueData\",\"type\":\"bytes\"}],\"name\":\"TransactionDeposited\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_isCreation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_byteCount\",\"type\":\"uint64\"}],\"name\":\"minimumGasLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paused_\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// OptimismPortalABI is the input ABI used to generate the binding from.
// Deprecated: Use OptimismPortalMetaData.ABI instead.
var OptimismPortalABI = OptimismPortalMetaData.ABI

// OptimismPortal is an auto generated Go binding around an Ethereum contract.
type OptimismPortal struct {
	OptimismPortalCaller     // Read-only binding to the contract
	OptimismPortalTransactor // Write-only binding to the contract
	OptimismPortalFilterer   // Log filterer for contract events
}

// OptimismPortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptimismPortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismPortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptimismPortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismPortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptimismPortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismPortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptimismPortalSession struct {
	Contract     *OptimismPortal   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OptimismPortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptimismPortalCallerSession struct {
	Contract *OptimismPortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OptimismPortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptimismPortalTransactorSession struct {
	Contract     *OptimismPortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OptimismPortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptimismPortalRaw struct {
	Contract *OptimismPortal // Generic contract binding to access the raw methods on
}

// OptimismPortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptimismPortalCallerRaw struct {
	Contract *OptimismPortalCaller // Generic read-only contract binding to access the raw methods on
}

// OptimismPortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptimismPortalTransactorRaw struct {
	Contract *OptimismPortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptimismPortal creates a new instance of OptimismPortal, bound to a specific deployed contract.
func NewOptimismPortal(address common.Address, backend bind.ContractBackend) (*OptimismPortal, error) {
	contract, err := bindOptimismPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimismPortal{OptimismPortalCaller: OptimismPortalCaller{contract: contract}, OptimismPortalTransactor: OptimismPortalTransactor{contract: contract}, OptimismPortalFilterer: OptimismPortalFilterer{contract: contract}}, nil
}

// NewOptimismPortalCaller creates a new read-only instance of OptimismPortal, bound to a specific deployed contract.
func NewOptimismPortalCaller(address common.Address, caller bind.ContractCaller) (*OptimismPortalCaller, error) {
	contract, err := bindOptimismPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismPortalCaller{contract: contract}, nil
}

// NewOptimismPortalTransactor creates a new write-only instance of OptimismPortal, bound to a specific deployed contract.
func NewOptimismPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimismPortalTransactor, error) {
	contract, err := bindOptimismPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismPortalTransactor{contract: contract}, nil
}

// NewOptimismPortalFilterer creates a new log filterer instance of OptimismPortal, bound to a specific deployed contract.
func NewOptimismPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimismPortalFilterer, error) {
	contract, err := bindOptimismPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimismPortalFilterer{contract: contract}, nil
}

// bindOptimismPortal binds a generic wrapper to an already deployed contract.
func bindOptimismPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimismPortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimismPortal *OptimismPortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismPortal.Contract.OptimismPortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimismPortal *OptimismPortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismPortal.Contract.OptimismPortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimismPortal *OptimismPortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismPortal.Contract.OptimismPortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimismPortal *OptimismPortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismPortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimismPortal *OptimismPortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismPortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimismPortal *OptimismPortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismPortal.Contract.contract.Transact(opts, method, params...)
}

// MinimumGasLimit is a free data retrieval call binding the contract method 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (_OptimismPortal *OptimismPortalCaller) MinimumGasLimit(opts *bind.CallOpts, _byteCount uint64) (uint64, error) {
	var out []interface{}
	err := _OptimismPortal.contract.Call(opts, &out, "minimumGasLimit", _byteCount)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinimumGasLimit is a free data retrieval call binding the contract method 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (_OptimismPortal *OptimismPortalSession) MinimumGasLimit(_byteCount uint64) (uint64, error) {
	return _OptimismPortal.Contract.MinimumGasLimit(&_OptimismPortal.CallOpts, _byteCount)
}

// MinimumGasLimit is a free data retrieval call binding the contract method 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (_OptimismPortal *OptimismPortalCallerSession) MinimumGasLimit(_byteCount uint64) (uint64, error) {
	return _OptimismPortal.Contract.MinimumGasLimit(&_OptimismPortal.CallOpts, _byteCount)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool paused_)
func (_OptimismPortal *OptimismPortalCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OptimismPortal.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool paused_)
func (_OptimismPortal *OptimismPortalSession) Paused() (bool, error) {
	return _OptimismPortal.Contract.Paused(&_OptimismPortal.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool paused_)
func (_OptimismPortal *OptimismPortalCallerSession) Paused() (bool, error) {
	return _OptimismPortal.Contract.Paused(&_OptimismPortal.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OptimismPortal *OptimismPortalCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OptimismPortal.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OptimismPortal *OptimismPortalSession) Version() (string, error) {
	return _OptimismPortal.Contract.Version(&_OptimismPortal.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OptimismPortal *OptimismPortalCallerSession) Version() (string, error) {
	return _OptimismPortal.Contract.Version(&_OptimismPortal.CallOpts)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0xe9e05c42.
//
// Solidity: function depositTransaction(address _to, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_OptimismPortal *OptimismPortalTransactor) DepositTransaction(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _OptimismPortal.contract.Transact(opts, "depositTransaction", _to, _value, _gasLimit, _isCreation, _data)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0xe9e05c42.
//
// Solidity: function depositTransaction(address _to, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_OptimismPortal *OptimismPortalSession) DepositTransaction(_to common.Address, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _OptimismPortal.Contract.DepositTransaction(&_OptimismPortal.TransactOpts, _to, _value, _gasLimit, _isCreation, _data)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0xe9e05c42.
//
// Solidity: function depositTransaction(address _to, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_OptimismPortal *OptimismPortalTransactorSession) DepositTransaction(_to common.Address, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _OptimismPortal.Contract.DepositTransaction(&_OptimismPortal.TransactOpts, _to, _value, _gasLimit, _isCreation, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OptimismPortal *OptimismPortalTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismPortal.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OptimismPortal *OptimismPortalSession) Initialize() (*types.Transaction, error) {
	return _OptimismPortal.Contract.Initialize(&_OptimismPortal.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OptimismPortal *OptimismPortalTransactorSession) Initialize() (*types.Transaction, error) {
	return _OptimismPortal.Contract.Initialize(&_OptimismPortal.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OptimismPortal *OptimismPortalTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismPortal.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OptimismPortal *OptimismPortalSession) Receive() (*types.Transaction, error) {
	return _OptimismPortal.Contract.Receive(&_OptimismPortal.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OptimismPortal *OptimismPortalTransactorSession) Receive() (*types.Transaction, error) {
	return _OptimismPortal.Contract.Receive(&_OptimismPortal.TransactOpts)
}

// OptimismPortalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OptimismPortal contract.
type OptimismPortalInitializedIterator struct {
	Event *OptimismPortalInitialized // Event containing the contract specifics and raw log

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
func (it *OptimismPortalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismPortalInitialized)
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
		it.Event = new(OptimismPortalInitialized)
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
func (it *OptimismPortalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimismPortalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimismPortalInitialized represents a Initialized event raised by the OptimismPortal contract.
type OptimismPortalInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OptimismPortal *OptimismPortalFilterer) FilterInitialized(opts *bind.FilterOpts) (*OptimismPortalInitializedIterator, error) {

	logs, sub, err := _OptimismPortal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OptimismPortalInitializedIterator{contract: _OptimismPortal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OptimismPortal *OptimismPortalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OptimismPortalInitialized) (event.Subscription, error) {

	logs, sub, err := _OptimismPortal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimismPortalInitialized)
				if err := _OptimismPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OptimismPortal *OptimismPortalFilterer) ParseInitialized(log types.Log) (*OptimismPortalInitialized, error) {
	event := new(OptimismPortalInitialized)
	if err := _OptimismPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptimismPortalTransactionDepositedIterator is returned from FilterTransactionDeposited and is used to iterate over the raw logs and unpacked data for TransactionDeposited events raised by the OptimismPortal contract.
type OptimismPortalTransactionDepositedIterator struct {
	Event *OptimismPortalTransactionDeposited // Event containing the contract specifics and raw log

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
func (it *OptimismPortalTransactionDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismPortalTransactionDeposited)
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
		it.Event = new(OptimismPortalTransactionDeposited)
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
func (it *OptimismPortalTransactionDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimismPortalTransactionDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimismPortalTransactionDeposited represents a TransactionDeposited event raised by the OptimismPortal contract.
type OptimismPortalTransactionDeposited struct {
	From       common.Address
	To         common.Address
	Version    *big.Int
	OpaqueData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionDeposited is a free log retrieval operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_OptimismPortal *OptimismPortalFilterer) FilterTransactionDeposited(opts *bind.FilterOpts, from []common.Address, to []common.Address, version []*big.Int) (*OptimismPortalTransactionDepositedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _OptimismPortal.contract.FilterLogs(opts, "TransactionDeposited", fromRule, toRule, versionRule)
	if err != nil {
		return nil, err
	}
	return &OptimismPortalTransactionDepositedIterator{contract: _OptimismPortal.contract, event: "TransactionDeposited", logs: logs, sub: sub}, nil
}

// WatchTransactionDeposited is a free log subscription operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_OptimismPortal *OptimismPortalFilterer) WatchTransactionDeposited(opts *bind.WatchOpts, sink chan<- *OptimismPortalTransactionDeposited, from []common.Address, to []common.Address, version []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _OptimismPortal.contract.WatchLogs(opts, "TransactionDeposited", fromRule, toRule, versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimismPortalTransactionDeposited)
				if err := _OptimismPortal.contract.UnpackLog(event, "TransactionDeposited", log); err != nil {
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

// ParseTransactionDeposited is a log parse operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_OptimismPortal *OptimismPortalFilterer) ParseTransactionDeposited(log types.Log) (*OptimismPortalTransactionDeposited, error) {
	event := new(OptimismPortalTransactionDeposited)
	if err := _OptimismPortal.contract.UnpackLog(event, "TransactionDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
