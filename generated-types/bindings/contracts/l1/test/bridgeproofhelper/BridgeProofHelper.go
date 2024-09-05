// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgeproofhelper

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

// TypesOutputRootProof is an auto generated low-level Go binding around an user-defined struct.
type TypesOutputRootProof struct {
	Version                  [32]byte
	StateRoot                [32]byte
	MessagePasserStorageRoot [32]byte
	LatestBlockhash          [32]byte
}

// TypesWithdrawalTransaction is an auto generated low-level Go binding around an user-defined struct.
type TypesWithdrawalTransaction struct {
	Nonce    *big.Int
	Sender   common.Address
	Target   common.Address
	Value    *big.Int
	GasLimit *big.Int
	Data     []byte
}

// BridgeProofHelperMetaData contains all meta data concerning the BridgeProofHelper contract.
var BridgeProofHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ContentLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataRemainder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidHeader\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"latestBlockhash\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.OutputRootProof\",\"name\":\"outputProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"withdrawalProof\",\"type\":\"bytes[]\"}],\"name\":\"checkProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.WithdrawalTransaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"hashWithdrawalTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// BridgeProofHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeProofHelperMetaData.ABI instead.
var BridgeProofHelperABI = BridgeProofHelperMetaData.ABI

// BridgeProofHelper is an auto generated Go binding around an Ethereum contract.
type BridgeProofHelper struct {
	BridgeProofHelperCaller     // Read-only binding to the contract
	BridgeProofHelperTransactor // Write-only binding to the contract
	BridgeProofHelperFilterer   // Log filterer for contract events
}

// BridgeProofHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeProofHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeProofHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeProofHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeProofHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeProofHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeProofHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeProofHelperSession struct {
	Contract     *BridgeProofHelper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BridgeProofHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeProofHelperCallerSession struct {
	Contract *BridgeProofHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BridgeProofHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeProofHelperTransactorSession struct {
	Contract     *BridgeProofHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BridgeProofHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeProofHelperRaw struct {
	Contract *BridgeProofHelper // Generic contract binding to access the raw methods on
}

// BridgeProofHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeProofHelperCallerRaw struct {
	Contract *BridgeProofHelperCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeProofHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeProofHelperTransactorRaw struct {
	Contract *BridgeProofHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeProofHelper creates a new instance of BridgeProofHelper, bound to a specific deployed contract.
func NewBridgeProofHelper(address common.Address, backend bind.ContractBackend) (*BridgeProofHelper, error) {
	contract, err := bindBridgeProofHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeProofHelper{BridgeProofHelperCaller: BridgeProofHelperCaller{contract: contract}, BridgeProofHelperTransactor: BridgeProofHelperTransactor{contract: contract}, BridgeProofHelperFilterer: BridgeProofHelperFilterer{contract: contract}}, nil
}

// NewBridgeProofHelperCaller creates a new read-only instance of BridgeProofHelper, bound to a specific deployed contract.
func NewBridgeProofHelperCaller(address common.Address, caller bind.ContractCaller) (*BridgeProofHelperCaller, error) {
	contract, err := bindBridgeProofHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeProofHelperCaller{contract: contract}, nil
}

// NewBridgeProofHelperTransactor creates a new write-only instance of BridgeProofHelper, bound to a specific deployed contract.
func NewBridgeProofHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeProofHelperTransactor, error) {
	contract, err := bindBridgeProofHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeProofHelperTransactor{contract: contract}, nil
}

// NewBridgeProofHelperFilterer creates a new log filterer instance of BridgeProofHelper, bound to a specific deployed contract.
func NewBridgeProofHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeProofHelperFilterer, error) {
	contract, err := bindBridgeProofHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeProofHelperFilterer{contract: contract}, nil
}

// bindBridgeProofHelper binds a generic wrapper to an already deployed contract.
func bindBridgeProofHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeProofHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeProofHelper *BridgeProofHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeProofHelper.Contract.BridgeProofHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeProofHelper *BridgeProofHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeProofHelper.Contract.BridgeProofHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeProofHelper *BridgeProofHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeProofHelper.Contract.BridgeProofHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeProofHelper *BridgeProofHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeProofHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeProofHelper *BridgeProofHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeProofHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeProofHelper *BridgeProofHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeProofHelper.Contract.contract.Transact(opts, method, params...)
}

// CheckProof is a free data retrieval call binding the contract method 0x8e59013c.
//
// Solidity: function checkProof(bytes32 outputRoot, (bytes32,bytes32,bytes32,bytes32) outputProof, bytes32 withdrawalHash, bytes[] withdrawalProof) pure returns(bool)
func (_BridgeProofHelper *BridgeProofHelperCaller) CheckProof(opts *bind.CallOpts, outputRoot [32]byte, outputProof TypesOutputRootProof, withdrawalHash [32]byte, withdrawalProof [][]byte) (bool, error) {
	var out []interface{}
	err := _BridgeProofHelper.contract.Call(opts, &out, "checkProof", outputRoot, outputProof, withdrawalHash, withdrawalProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckProof is a free data retrieval call binding the contract method 0x8e59013c.
//
// Solidity: function checkProof(bytes32 outputRoot, (bytes32,bytes32,bytes32,bytes32) outputProof, bytes32 withdrawalHash, bytes[] withdrawalProof) pure returns(bool)
func (_BridgeProofHelper *BridgeProofHelperSession) CheckProof(outputRoot [32]byte, outputProof TypesOutputRootProof, withdrawalHash [32]byte, withdrawalProof [][]byte) (bool, error) {
	return _BridgeProofHelper.Contract.CheckProof(&_BridgeProofHelper.CallOpts, outputRoot, outputProof, withdrawalHash, withdrawalProof)
}

// CheckProof is a free data retrieval call binding the contract method 0x8e59013c.
//
// Solidity: function checkProof(bytes32 outputRoot, (bytes32,bytes32,bytes32,bytes32) outputProof, bytes32 withdrawalHash, bytes[] withdrawalProof) pure returns(bool)
func (_BridgeProofHelper *BridgeProofHelperCallerSession) CheckProof(outputRoot [32]byte, outputProof TypesOutputRootProof, withdrawalHash [32]byte, withdrawalProof [][]byte) (bool, error) {
	return _BridgeProofHelper.Contract.CheckProof(&_BridgeProofHelper.CallOpts, outputRoot, outputProof, withdrawalHash, withdrawalProof)
}

// HashWithdrawalTx is a free data retrieval call binding the contract method 0xcac2647a.
//
// Solidity: function hashWithdrawalTx((uint256,address,address,uint256,uint256,bytes) _tx) pure returns(bytes32)
func (_BridgeProofHelper *BridgeProofHelperCaller) HashWithdrawalTx(opts *bind.CallOpts, _tx TypesWithdrawalTransaction) ([32]byte, error) {
	var out []interface{}
	err := _BridgeProofHelper.contract.Call(opts, &out, "hashWithdrawalTx", _tx)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashWithdrawalTx is a free data retrieval call binding the contract method 0xcac2647a.
//
// Solidity: function hashWithdrawalTx((uint256,address,address,uint256,uint256,bytes) _tx) pure returns(bytes32)
func (_BridgeProofHelper *BridgeProofHelperSession) HashWithdrawalTx(_tx TypesWithdrawalTransaction) ([32]byte, error) {
	return _BridgeProofHelper.Contract.HashWithdrawalTx(&_BridgeProofHelper.CallOpts, _tx)
}

// HashWithdrawalTx is a free data retrieval call binding the contract method 0xcac2647a.
//
// Solidity: function hashWithdrawalTx((uint256,address,address,uint256,uint256,bytes) _tx) pure returns(bytes32)
func (_BridgeProofHelper *BridgeProofHelperCallerSession) HashWithdrawalTx(_tx TypesWithdrawalTransaction) ([32]byte, error) {
	return _BridgeProofHelper.Contract.HashWithdrawalTx(&_BridgeProofHelper.CallOpts, _tx)
}
