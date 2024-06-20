// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iblobstreamx

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

// IBlobstreamXMetaData contains all meta data concerning the IBlobstreamX contract.
var IBlobstreamXMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ContractFrozen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataCommitmentNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LatestHeaderNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetBlockNotInRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedBlockMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedHeaderNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proofNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"startBlock\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"endBlock\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"dataCommitment\",\"type\":\"bytes32\"}],\"name\":\"DataCommitmentStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"}],\"name\":\"HeadUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"trustedBlock\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trustedHeader\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"targetBlock\",\"type\":\"uint64\"}],\"name\":\"HeaderRangeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"trustedBlock\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trustedHeader\",\"type\":\"bytes32\"}],\"name\":\"NextHeaderRequested\",\"type\":\"event\"}]",
}

// IBlobstreamXABI is the input ABI used to generate the binding from.
// Deprecated: Use IBlobstreamXMetaData.ABI instead.
var IBlobstreamXABI = IBlobstreamXMetaData.ABI

// IBlobstreamX is an auto generated Go binding around an Ethereum contract.
type IBlobstreamX struct {
	IBlobstreamXCaller     // Read-only binding to the contract
	IBlobstreamXTransactor // Write-only binding to the contract
	IBlobstreamXFilterer   // Log filterer for contract events
}

// IBlobstreamXCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBlobstreamXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlobstreamXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBlobstreamXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlobstreamXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBlobstreamXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlobstreamXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBlobstreamXSession struct {
	Contract     *IBlobstreamX     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBlobstreamXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBlobstreamXCallerSession struct {
	Contract *IBlobstreamXCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IBlobstreamXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBlobstreamXTransactorSession struct {
	Contract     *IBlobstreamXTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IBlobstreamXRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBlobstreamXRaw struct {
	Contract *IBlobstreamX // Generic contract binding to access the raw methods on
}

// IBlobstreamXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBlobstreamXCallerRaw struct {
	Contract *IBlobstreamXCaller // Generic read-only contract binding to access the raw methods on
}

// IBlobstreamXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBlobstreamXTransactorRaw struct {
	Contract *IBlobstreamXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBlobstreamX creates a new instance of IBlobstreamX, bound to a specific deployed contract.
func NewIBlobstreamX(address common.Address, backend bind.ContractBackend) (*IBlobstreamX, error) {
	contract, err := bindIBlobstreamX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBlobstreamX{IBlobstreamXCaller: IBlobstreamXCaller{contract: contract}, IBlobstreamXTransactor: IBlobstreamXTransactor{contract: contract}, IBlobstreamXFilterer: IBlobstreamXFilterer{contract: contract}}, nil
}

// NewIBlobstreamXCaller creates a new read-only instance of IBlobstreamX, bound to a specific deployed contract.
func NewIBlobstreamXCaller(address common.Address, caller bind.ContractCaller) (*IBlobstreamXCaller, error) {
	contract, err := bindIBlobstreamX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBlobstreamXCaller{contract: contract}, nil
}

// NewIBlobstreamXTransactor creates a new write-only instance of IBlobstreamX, bound to a specific deployed contract.
func NewIBlobstreamXTransactor(address common.Address, transactor bind.ContractTransactor) (*IBlobstreamXTransactor, error) {
	contract, err := bindIBlobstreamX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBlobstreamXTransactor{contract: contract}, nil
}

// NewIBlobstreamXFilterer creates a new log filterer instance of IBlobstreamX, bound to a specific deployed contract.
func NewIBlobstreamXFilterer(address common.Address, filterer bind.ContractFilterer) (*IBlobstreamXFilterer, error) {
	contract, err := bindIBlobstreamX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBlobstreamXFilterer{contract: contract}, nil
}

// bindIBlobstreamX binds a generic wrapper to an already deployed contract.
func bindIBlobstreamX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBlobstreamXMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlobstreamX *IBlobstreamXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlobstreamX.Contract.IBlobstreamXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlobstreamX *IBlobstreamXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlobstreamX.Contract.IBlobstreamXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlobstreamX *IBlobstreamXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlobstreamX.Contract.IBlobstreamXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlobstreamX *IBlobstreamXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlobstreamX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlobstreamX *IBlobstreamXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlobstreamX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlobstreamX *IBlobstreamXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlobstreamX.Contract.contract.Transact(opts, method, params...)
}

// IBlobstreamXDataCommitmentStoredIterator is returned from FilterDataCommitmentStored and is used to iterate over the raw logs and unpacked data for DataCommitmentStored events raised by the IBlobstreamX contract.
type IBlobstreamXDataCommitmentStoredIterator struct {
	Event *IBlobstreamXDataCommitmentStored // Event containing the contract specifics and raw log

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
func (it *IBlobstreamXDataCommitmentStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlobstreamXDataCommitmentStored)
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
		it.Event = new(IBlobstreamXDataCommitmentStored)
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
func (it *IBlobstreamXDataCommitmentStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlobstreamXDataCommitmentStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlobstreamXDataCommitmentStored represents a DataCommitmentStored event raised by the IBlobstreamX contract.
type IBlobstreamXDataCommitmentStored struct {
	ProofNonce     *big.Int
	StartBlock     uint64
	EndBlock       uint64
	DataCommitment [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDataCommitmentStored is a free log retrieval operation binding the contract event 0x34dd3689f5bd77a60a3ff2e09483dcab032fa2f1fd7227af3e24bed21beab1cb.
//
// Solidity: event DataCommitmentStored(uint256 proofNonce, uint64 indexed startBlock, uint64 indexed endBlock, bytes32 indexed dataCommitment)
func (_IBlobstreamX *IBlobstreamXFilterer) FilterDataCommitmentStored(opts *bind.FilterOpts, startBlock []uint64, endBlock []uint64, dataCommitment [][32]byte) (*IBlobstreamXDataCommitmentStoredIterator, error) {

	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}
	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}
	var dataCommitmentRule []interface{}
	for _, dataCommitmentItem := range dataCommitment {
		dataCommitmentRule = append(dataCommitmentRule, dataCommitmentItem)
	}

	logs, sub, err := _IBlobstreamX.contract.FilterLogs(opts, "DataCommitmentStored", startBlockRule, endBlockRule, dataCommitmentRule)
	if err != nil {
		return nil, err
	}
	return &IBlobstreamXDataCommitmentStoredIterator{contract: _IBlobstreamX.contract, event: "DataCommitmentStored", logs: logs, sub: sub}, nil
}

// WatchDataCommitmentStored is a free log subscription operation binding the contract event 0x34dd3689f5bd77a60a3ff2e09483dcab032fa2f1fd7227af3e24bed21beab1cb.
//
// Solidity: event DataCommitmentStored(uint256 proofNonce, uint64 indexed startBlock, uint64 indexed endBlock, bytes32 indexed dataCommitment)
func (_IBlobstreamX *IBlobstreamXFilterer) WatchDataCommitmentStored(opts *bind.WatchOpts, sink chan<- *IBlobstreamXDataCommitmentStored, startBlock []uint64, endBlock []uint64, dataCommitment [][32]byte) (event.Subscription, error) {

	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}
	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}
	var dataCommitmentRule []interface{}
	for _, dataCommitmentItem := range dataCommitment {
		dataCommitmentRule = append(dataCommitmentRule, dataCommitmentItem)
	}

	logs, sub, err := _IBlobstreamX.contract.WatchLogs(opts, "DataCommitmentStored", startBlockRule, endBlockRule, dataCommitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlobstreamXDataCommitmentStored)
				if err := _IBlobstreamX.contract.UnpackLog(event, "DataCommitmentStored", log); err != nil {
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

// ParseDataCommitmentStored is a log parse operation binding the contract event 0x34dd3689f5bd77a60a3ff2e09483dcab032fa2f1fd7227af3e24bed21beab1cb.
//
// Solidity: event DataCommitmentStored(uint256 proofNonce, uint64 indexed startBlock, uint64 indexed endBlock, bytes32 indexed dataCommitment)
func (_IBlobstreamX *IBlobstreamXFilterer) ParseDataCommitmentStored(log types.Log) (*IBlobstreamXDataCommitmentStored, error) {
	event := new(IBlobstreamXDataCommitmentStored)
	if err := _IBlobstreamX.contract.UnpackLog(event, "DataCommitmentStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBlobstreamXHeadUpdateIterator is returned from FilterHeadUpdate and is used to iterate over the raw logs and unpacked data for HeadUpdate events raised by the IBlobstreamX contract.
type IBlobstreamXHeadUpdateIterator struct {
	Event *IBlobstreamXHeadUpdate // Event containing the contract specifics and raw log

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
func (it *IBlobstreamXHeadUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlobstreamXHeadUpdate)
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
		it.Event = new(IBlobstreamXHeadUpdate)
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
func (it *IBlobstreamXHeadUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlobstreamXHeadUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlobstreamXHeadUpdate represents a HeadUpdate event raised by the IBlobstreamX contract.
type IBlobstreamXHeadUpdate struct {
	BlockNumber uint64
	HeaderHash  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterHeadUpdate is a free log retrieval operation binding the contract event 0x292f5abc3167175400fca463fa99530cda826ec53ec5eb1f3a2776006dacd75d.
//
// Solidity: event HeadUpdate(uint64 blockNumber, bytes32 headerHash)
func (_IBlobstreamX *IBlobstreamXFilterer) FilterHeadUpdate(opts *bind.FilterOpts) (*IBlobstreamXHeadUpdateIterator, error) {

	logs, sub, err := _IBlobstreamX.contract.FilterLogs(opts, "HeadUpdate")
	if err != nil {
		return nil, err
	}
	return &IBlobstreamXHeadUpdateIterator{contract: _IBlobstreamX.contract, event: "HeadUpdate", logs: logs, sub: sub}, nil
}

// WatchHeadUpdate is a free log subscription operation binding the contract event 0x292f5abc3167175400fca463fa99530cda826ec53ec5eb1f3a2776006dacd75d.
//
// Solidity: event HeadUpdate(uint64 blockNumber, bytes32 headerHash)
func (_IBlobstreamX *IBlobstreamXFilterer) WatchHeadUpdate(opts *bind.WatchOpts, sink chan<- *IBlobstreamXHeadUpdate) (event.Subscription, error) {

	logs, sub, err := _IBlobstreamX.contract.WatchLogs(opts, "HeadUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlobstreamXHeadUpdate)
				if err := _IBlobstreamX.contract.UnpackLog(event, "HeadUpdate", log); err != nil {
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

// ParseHeadUpdate is a log parse operation binding the contract event 0x292f5abc3167175400fca463fa99530cda826ec53ec5eb1f3a2776006dacd75d.
//
// Solidity: event HeadUpdate(uint64 blockNumber, bytes32 headerHash)
func (_IBlobstreamX *IBlobstreamXFilterer) ParseHeadUpdate(log types.Log) (*IBlobstreamXHeadUpdate, error) {
	event := new(IBlobstreamXHeadUpdate)
	if err := _IBlobstreamX.contract.UnpackLog(event, "HeadUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBlobstreamXHeaderRangeRequestedIterator is returned from FilterHeaderRangeRequested and is used to iterate over the raw logs and unpacked data for HeaderRangeRequested events raised by the IBlobstreamX contract.
type IBlobstreamXHeaderRangeRequestedIterator struct {
	Event *IBlobstreamXHeaderRangeRequested // Event containing the contract specifics and raw log

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
func (it *IBlobstreamXHeaderRangeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlobstreamXHeaderRangeRequested)
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
		it.Event = new(IBlobstreamXHeaderRangeRequested)
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
func (it *IBlobstreamXHeaderRangeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlobstreamXHeaderRangeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlobstreamXHeaderRangeRequested represents a HeaderRangeRequested event raised by the IBlobstreamX contract.
type IBlobstreamXHeaderRangeRequested struct {
	TrustedBlock  uint64
	TrustedHeader [32]byte
	TargetBlock   uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterHeaderRangeRequested is a free log retrieval operation binding the contract event 0xc8c2769e8252a98c11f30f2aa43a6721c1f3c046617bea717ae0999c6507dcfe.
//
// Solidity: event HeaderRangeRequested(uint64 indexed trustedBlock, bytes32 indexed trustedHeader, uint64 indexed targetBlock)
func (_IBlobstreamX *IBlobstreamXFilterer) FilterHeaderRangeRequested(opts *bind.FilterOpts, trustedBlock []uint64, trustedHeader [][32]byte, targetBlock []uint64) (*IBlobstreamXHeaderRangeRequestedIterator, error) {

	var trustedBlockRule []interface{}
	for _, trustedBlockItem := range trustedBlock {
		trustedBlockRule = append(trustedBlockRule, trustedBlockItem)
	}
	var trustedHeaderRule []interface{}
	for _, trustedHeaderItem := range trustedHeader {
		trustedHeaderRule = append(trustedHeaderRule, trustedHeaderItem)
	}
	var targetBlockRule []interface{}
	for _, targetBlockItem := range targetBlock {
		targetBlockRule = append(targetBlockRule, targetBlockItem)
	}

	logs, sub, err := _IBlobstreamX.contract.FilterLogs(opts, "HeaderRangeRequested", trustedBlockRule, trustedHeaderRule, targetBlockRule)
	if err != nil {
		return nil, err
	}
	return &IBlobstreamXHeaderRangeRequestedIterator{contract: _IBlobstreamX.contract, event: "HeaderRangeRequested", logs: logs, sub: sub}, nil
}

// WatchHeaderRangeRequested is a free log subscription operation binding the contract event 0xc8c2769e8252a98c11f30f2aa43a6721c1f3c046617bea717ae0999c6507dcfe.
//
// Solidity: event HeaderRangeRequested(uint64 indexed trustedBlock, bytes32 indexed trustedHeader, uint64 indexed targetBlock)
func (_IBlobstreamX *IBlobstreamXFilterer) WatchHeaderRangeRequested(opts *bind.WatchOpts, sink chan<- *IBlobstreamXHeaderRangeRequested, trustedBlock []uint64, trustedHeader [][32]byte, targetBlock []uint64) (event.Subscription, error) {

	var trustedBlockRule []interface{}
	for _, trustedBlockItem := range trustedBlock {
		trustedBlockRule = append(trustedBlockRule, trustedBlockItem)
	}
	var trustedHeaderRule []interface{}
	for _, trustedHeaderItem := range trustedHeader {
		trustedHeaderRule = append(trustedHeaderRule, trustedHeaderItem)
	}
	var targetBlockRule []interface{}
	for _, targetBlockItem := range targetBlock {
		targetBlockRule = append(targetBlockRule, targetBlockItem)
	}

	logs, sub, err := _IBlobstreamX.contract.WatchLogs(opts, "HeaderRangeRequested", trustedBlockRule, trustedHeaderRule, targetBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlobstreamXHeaderRangeRequested)
				if err := _IBlobstreamX.contract.UnpackLog(event, "HeaderRangeRequested", log); err != nil {
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

// ParseHeaderRangeRequested is a log parse operation binding the contract event 0xc8c2769e8252a98c11f30f2aa43a6721c1f3c046617bea717ae0999c6507dcfe.
//
// Solidity: event HeaderRangeRequested(uint64 indexed trustedBlock, bytes32 indexed trustedHeader, uint64 indexed targetBlock)
func (_IBlobstreamX *IBlobstreamXFilterer) ParseHeaderRangeRequested(log types.Log) (*IBlobstreamXHeaderRangeRequested, error) {
	event := new(IBlobstreamXHeaderRangeRequested)
	if err := _IBlobstreamX.contract.UnpackLog(event, "HeaderRangeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBlobstreamXNextHeaderRequestedIterator is returned from FilterNextHeaderRequested and is used to iterate over the raw logs and unpacked data for NextHeaderRequested events raised by the IBlobstreamX contract.
type IBlobstreamXNextHeaderRequestedIterator struct {
	Event *IBlobstreamXNextHeaderRequested // Event containing the contract specifics and raw log

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
func (it *IBlobstreamXNextHeaderRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlobstreamXNextHeaderRequested)
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
		it.Event = new(IBlobstreamXNextHeaderRequested)
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
func (it *IBlobstreamXNextHeaderRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlobstreamXNextHeaderRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlobstreamXNextHeaderRequested represents a NextHeaderRequested event raised by the IBlobstreamX contract.
type IBlobstreamXNextHeaderRequested struct {
	TrustedBlock  uint64
	TrustedHeader [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNextHeaderRequested is a free log retrieval operation binding the contract event 0x76ede8001988e3e0a3ef39c130769d418d94626cd8f53e08b45541d15526f8cc.
//
// Solidity: event NextHeaderRequested(uint64 indexed trustedBlock, bytes32 indexed trustedHeader)
func (_IBlobstreamX *IBlobstreamXFilterer) FilterNextHeaderRequested(opts *bind.FilterOpts, trustedBlock []uint64, trustedHeader [][32]byte) (*IBlobstreamXNextHeaderRequestedIterator, error) {

	var trustedBlockRule []interface{}
	for _, trustedBlockItem := range trustedBlock {
		trustedBlockRule = append(trustedBlockRule, trustedBlockItem)
	}
	var trustedHeaderRule []interface{}
	for _, trustedHeaderItem := range trustedHeader {
		trustedHeaderRule = append(trustedHeaderRule, trustedHeaderItem)
	}

	logs, sub, err := _IBlobstreamX.contract.FilterLogs(opts, "NextHeaderRequested", trustedBlockRule, trustedHeaderRule)
	if err != nil {
		return nil, err
	}
	return &IBlobstreamXNextHeaderRequestedIterator{contract: _IBlobstreamX.contract, event: "NextHeaderRequested", logs: logs, sub: sub}, nil
}

// WatchNextHeaderRequested is a free log subscription operation binding the contract event 0x76ede8001988e3e0a3ef39c130769d418d94626cd8f53e08b45541d15526f8cc.
//
// Solidity: event NextHeaderRequested(uint64 indexed trustedBlock, bytes32 indexed trustedHeader)
func (_IBlobstreamX *IBlobstreamXFilterer) WatchNextHeaderRequested(opts *bind.WatchOpts, sink chan<- *IBlobstreamXNextHeaderRequested, trustedBlock []uint64, trustedHeader [][32]byte) (event.Subscription, error) {

	var trustedBlockRule []interface{}
	for _, trustedBlockItem := range trustedBlock {
		trustedBlockRule = append(trustedBlockRule, trustedBlockItem)
	}
	var trustedHeaderRule []interface{}
	for _, trustedHeaderItem := range trustedHeader {
		trustedHeaderRule = append(trustedHeaderRule, trustedHeaderItem)
	}

	logs, sub, err := _IBlobstreamX.contract.WatchLogs(opts, "NextHeaderRequested", trustedBlockRule, trustedHeaderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlobstreamXNextHeaderRequested)
				if err := _IBlobstreamX.contract.UnpackLog(event, "NextHeaderRequested", log); err != nil {
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

// ParseNextHeaderRequested is a log parse operation binding the contract event 0x76ede8001988e3e0a3ef39c130769d418d94626cd8f53e08b45541d15526f8cc.
//
// Solidity: event NextHeaderRequested(uint64 indexed trustedBlock, bytes32 indexed trustedHeader)
func (_IBlobstreamX *IBlobstreamXFilterer) ParseNextHeaderRequested(log types.Log) (*IBlobstreamXNextHeaderRequested, error) {
	event := new(IBlobstreamXNextHeaderRequested)
	if err := _IBlobstreamX.contract.UnpackLog(event, "NextHeaderRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
