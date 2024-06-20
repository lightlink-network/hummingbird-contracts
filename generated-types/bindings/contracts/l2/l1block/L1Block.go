// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1block

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

// L1BlockMetaData contains all meta data concerning the L1Block contract.
var L1BlockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NotDepositor\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"}],\"name\":\"GasPayingTokenSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSITOR_ACCOUNT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseFeeScalar\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"basefee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batcherHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blobBaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blobBaseFeeScalar\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasPayingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasPayingTokenName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasPayingTokenSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCustomGasToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1FeeOverhead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1FeeScalar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_symbol\",\"type\":\"bytes32\"}],\"name\":\"setGasPayingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_basefee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_batcherHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l1FeeOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l1FeeScalar\",\"type\":\"uint256\"}],\"name\":\"setL1BlockValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setL1BlockValuesEcotone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// L1BlockABI is the input ABI used to generate the binding from.
// Deprecated: Use L1BlockMetaData.ABI instead.
var L1BlockABI = L1BlockMetaData.ABI

// L1Block is an auto generated Go binding around an Ethereum contract.
type L1Block struct {
	L1BlockCaller     // Read-only binding to the contract
	L1BlockTransactor // Write-only binding to the contract
	L1BlockFilterer   // Log filterer for contract events
}

// L1BlockCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1BlockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1BlockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1BlockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1BlockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1BlockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1BlockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1BlockSession struct {
	Contract     *L1Block          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1BlockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1BlockCallerSession struct {
	Contract *L1BlockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// L1BlockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1BlockTransactorSession struct {
	Contract     *L1BlockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// L1BlockRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1BlockRaw struct {
	Contract *L1Block // Generic contract binding to access the raw methods on
}

// L1BlockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1BlockCallerRaw struct {
	Contract *L1BlockCaller // Generic read-only contract binding to access the raw methods on
}

// L1BlockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1BlockTransactorRaw struct {
	Contract *L1BlockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1Block creates a new instance of L1Block, bound to a specific deployed contract.
func NewL1Block(address common.Address, backend bind.ContractBackend) (*L1Block, error) {
	contract, err := bindL1Block(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1Block{L1BlockCaller: L1BlockCaller{contract: contract}, L1BlockTransactor: L1BlockTransactor{contract: contract}, L1BlockFilterer: L1BlockFilterer{contract: contract}}, nil
}

// NewL1BlockCaller creates a new read-only instance of L1Block, bound to a specific deployed contract.
func NewL1BlockCaller(address common.Address, caller bind.ContractCaller) (*L1BlockCaller, error) {
	contract, err := bindL1Block(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1BlockCaller{contract: contract}, nil
}

// NewL1BlockTransactor creates a new write-only instance of L1Block, bound to a specific deployed contract.
func NewL1BlockTransactor(address common.Address, transactor bind.ContractTransactor) (*L1BlockTransactor, error) {
	contract, err := bindL1Block(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1BlockTransactor{contract: contract}, nil
}

// NewL1BlockFilterer creates a new log filterer instance of L1Block, bound to a specific deployed contract.
func NewL1BlockFilterer(address common.Address, filterer bind.ContractFilterer) (*L1BlockFilterer, error) {
	contract, err := bindL1Block(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1BlockFilterer{contract: contract}, nil
}

// bindL1Block binds a generic wrapper to an already deployed contract.
func bindL1Block(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1BlockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1Block *L1BlockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1Block.Contract.L1BlockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1Block *L1BlockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Block.Contract.L1BlockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1Block *L1BlockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1Block.Contract.L1BlockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1Block *L1BlockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1Block.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1Block *L1BlockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Block.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1Block *L1BlockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1Block.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() pure returns(address addr_)
func (_L1Block *L1BlockCaller) DEPOSITORACCOUNT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "DEPOSITOR_ACCOUNT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() pure returns(address addr_)
func (_L1Block *L1BlockSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _L1Block.Contract.DEPOSITORACCOUNT(&_L1Block.CallOpts)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() pure returns(address addr_)
func (_L1Block *L1BlockCallerSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _L1Block.Contract.DEPOSITORACCOUNT(&_L1Block.CallOpts)
}

// BaseFeeScalar is a free data retrieval call binding the contract method 0xc5985918.
//
// Solidity: function baseFeeScalar() view returns(uint32)
func (_L1Block *L1BlockCaller) BaseFeeScalar(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "baseFeeScalar")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BaseFeeScalar is a free data retrieval call binding the contract method 0xc5985918.
//
// Solidity: function baseFeeScalar() view returns(uint32)
func (_L1Block *L1BlockSession) BaseFeeScalar() (uint32, error) {
	return _L1Block.Contract.BaseFeeScalar(&_L1Block.CallOpts)
}

// BaseFeeScalar is a free data retrieval call binding the contract method 0xc5985918.
//
// Solidity: function baseFeeScalar() view returns(uint32)
func (_L1Block *L1BlockCallerSession) BaseFeeScalar() (uint32, error) {
	return _L1Block.Contract.BaseFeeScalar(&_L1Block.CallOpts)
}

// Basefee is a free data retrieval call binding the contract method 0x5cf24969.
//
// Solidity: function basefee() view returns(uint256)
func (_L1Block *L1BlockCaller) Basefee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "basefee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Basefee is a free data retrieval call binding the contract method 0x5cf24969.
//
// Solidity: function basefee() view returns(uint256)
func (_L1Block *L1BlockSession) Basefee() (*big.Int, error) {
	return _L1Block.Contract.Basefee(&_L1Block.CallOpts)
}

// Basefee is a free data retrieval call binding the contract method 0x5cf24969.
//
// Solidity: function basefee() view returns(uint256)
func (_L1Block *L1BlockCallerSession) Basefee() (*big.Int, error) {
	return _L1Block.Contract.Basefee(&_L1Block.CallOpts)
}

// BatcherHash is a free data retrieval call binding the contract method 0xe81b2c6d.
//
// Solidity: function batcherHash() view returns(bytes32)
func (_L1Block *L1BlockCaller) BatcherHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "batcherHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatcherHash is a free data retrieval call binding the contract method 0xe81b2c6d.
//
// Solidity: function batcherHash() view returns(bytes32)
func (_L1Block *L1BlockSession) BatcherHash() ([32]byte, error) {
	return _L1Block.Contract.BatcherHash(&_L1Block.CallOpts)
}

// BatcherHash is a free data retrieval call binding the contract method 0xe81b2c6d.
//
// Solidity: function batcherHash() view returns(bytes32)
func (_L1Block *L1BlockCallerSession) BatcherHash() ([32]byte, error) {
	return _L1Block.Contract.BatcherHash(&_L1Block.CallOpts)
}

// BlobBaseFee is a free data retrieval call binding the contract method 0xf8206140.
//
// Solidity: function blobBaseFee() view returns(uint256)
func (_L1Block *L1BlockCaller) BlobBaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "blobBaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlobBaseFee is a free data retrieval call binding the contract method 0xf8206140.
//
// Solidity: function blobBaseFee() view returns(uint256)
func (_L1Block *L1BlockSession) BlobBaseFee() (*big.Int, error) {
	return _L1Block.Contract.BlobBaseFee(&_L1Block.CallOpts)
}

// BlobBaseFee is a free data retrieval call binding the contract method 0xf8206140.
//
// Solidity: function blobBaseFee() view returns(uint256)
func (_L1Block *L1BlockCallerSession) BlobBaseFee() (*big.Int, error) {
	return _L1Block.Contract.BlobBaseFee(&_L1Block.CallOpts)
}

// BlobBaseFeeScalar is a free data retrieval call binding the contract method 0x68d5dca6.
//
// Solidity: function blobBaseFeeScalar() view returns(uint32)
func (_L1Block *L1BlockCaller) BlobBaseFeeScalar(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "blobBaseFeeScalar")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BlobBaseFeeScalar is a free data retrieval call binding the contract method 0x68d5dca6.
//
// Solidity: function blobBaseFeeScalar() view returns(uint32)
func (_L1Block *L1BlockSession) BlobBaseFeeScalar() (uint32, error) {
	return _L1Block.Contract.BlobBaseFeeScalar(&_L1Block.CallOpts)
}

// BlobBaseFeeScalar is a free data retrieval call binding the contract method 0x68d5dca6.
//
// Solidity: function blobBaseFeeScalar() view returns(uint32)
func (_L1Block *L1BlockCallerSession) BlobBaseFeeScalar() (uint32, error) {
	return _L1Block.Contract.BlobBaseFeeScalar(&_L1Block.CallOpts)
}

// GasPayingToken is a free data retrieval call binding the contract method 0x4397dfef.
//
// Solidity: function gasPayingToken() view returns(address addr_, uint8 decimals_)
func (_L1Block *L1BlockCaller) GasPayingToken(opts *bind.CallOpts) (struct {
	Addr     common.Address
	Decimals uint8
}, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "gasPayingToken")

	outstruct := new(struct {
		Addr     common.Address
		Decimals uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Decimals = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// GasPayingToken is a free data retrieval call binding the contract method 0x4397dfef.
//
// Solidity: function gasPayingToken() view returns(address addr_, uint8 decimals_)
func (_L1Block *L1BlockSession) GasPayingToken() (struct {
	Addr     common.Address
	Decimals uint8
}, error) {
	return _L1Block.Contract.GasPayingToken(&_L1Block.CallOpts)
}

// GasPayingToken is a free data retrieval call binding the contract method 0x4397dfef.
//
// Solidity: function gasPayingToken() view returns(address addr_, uint8 decimals_)
func (_L1Block *L1BlockCallerSession) GasPayingToken() (struct {
	Addr     common.Address
	Decimals uint8
}, error) {
	return _L1Block.Contract.GasPayingToken(&_L1Block.CallOpts)
}

// GasPayingTokenName is a free data retrieval call binding the contract method 0xd8444715.
//
// Solidity: function gasPayingTokenName() view returns(string name_)
func (_L1Block *L1BlockCaller) GasPayingTokenName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "gasPayingTokenName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GasPayingTokenName is a free data retrieval call binding the contract method 0xd8444715.
//
// Solidity: function gasPayingTokenName() view returns(string name_)
func (_L1Block *L1BlockSession) GasPayingTokenName() (string, error) {
	return _L1Block.Contract.GasPayingTokenName(&_L1Block.CallOpts)
}

// GasPayingTokenName is a free data retrieval call binding the contract method 0xd8444715.
//
// Solidity: function gasPayingTokenName() view returns(string name_)
func (_L1Block *L1BlockCallerSession) GasPayingTokenName() (string, error) {
	return _L1Block.Contract.GasPayingTokenName(&_L1Block.CallOpts)
}

// GasPayingTokenSymbol is a free data retrieval call binding the contract method 0x550fcdc9.
//
// Solidity: function gasPayingTokenSymbol() view returns(string symbol_)
func (_L1Block *L1BlockCaller) GasPayingTokenSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "gasPayingTokenSymbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GasPayingTokenSymbol is a free data retrieval call binding the contract method 0x550fcdc9.
//
// Solidity: function gasPayingTokenSymbol() view returns(string symbol_)
func (_L1Block *L1BlockSession) GasPayingTokenSymbol() (string, error) {
	return _L1Block.Contract.GasPayingTokenSymbol(&_L1Block.CallOpts)
}

// GasPayingTokenSymbol is a free data retrieval call binding the contract method 0x550fcdc9.
//
// Solidity: function gasPayingTokenSymbol() view returns(string symbol_)
func (_L1Block *L1BlockCallerSession) GasPayingTokenSymbol() (string, error) {
	return _L1Block.Contract.GasPayingTokenSymbol(&_L1Block.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x09bd5a60.
//
// Solidity: function hash() view returns(bytes32)
func (_L1Block *L1BlockCaller) Hash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "hash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x09bd5a60.
//
// Solidity: function hash() view returns(bytes32)
func (_L1Block *L1BlockSession) Hash() ([32]byte, error) {
	return _L1Block.Contract.Hash(&_L1Block.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x09bd5a60.
//
// Solidity: function hash() view returns(bytes32)
func (_L1Block *L1BlockCallerSession) Hash() ([32]byte, error) {
	return _L1Block.Contract.Hash(&_L1Block.CallOpts)
}

// IsCustomGasToken is a free data retrieval call binding the contract method 0x21326849.
//
// Solidity: function isCustomGasToken() view returns(bool)
func (_L1Block *L1BlockCaller) IsCustomGasToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "isCustomGasToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCustomGasToken is a free data retrieval call binding the contract method 0x21326849.
//
// Solidity: function isCustomGasToken() view returns(bool)
func (_L1Block *L1BlockSession) IsCustomGasToken() (bool, error) {
	return _L1Block.Contract.IsCustomGasToken(&_L1Block.CallOpts)
}

// IsCustomGasToken is a free data retrieval call binding the contract method 0x21326849.
//
// Solidity: function isCustomGasToken() view returns(bool)
func (_L1Block *L1BlockCallerSession) IsCustomGasToken() (bool, error) {
	return _L1Block.Contract.IsCustomGasToken(&_L1Block.CallOpts)
}

// L1FeeOverhead is a free data retrieval call binding the contract method 0x8b239f73.
//
// Solidity: function l1FeeOverhead() view returns(uint256)
func (_L1Block *L1BlockCaller) L1FeeOverhead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "l1FeeOverhead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1FeeOverhead is a free data retrieval call binding the contract method 0x8b239f73.
//
// Solidity: function l1FeeOverhead() view returns(uint256)
func (_L1Block *L1BlockSession) L1FeeOverhead() (*big.Int, error) {
	return _L1Block.Contract.L1FeeOverhead(&_L1Block.CallOpts)
}

// L1FeeOverhead is a free data retrieval call binding the contract method 0x8b239f73.
//
// Solidity: function l1FeeOverhead() view returns(uint256)
func (_L1Block *L1BlockCallerSession) L1FeeOverhead() (*big.Int, error) {
	return _L1Block.Contract.L1FeeOverhead(&_L1Block.CallOpts)
}

// L1FeeScalar is a free data retrieval call binding the contract method 0x9e8c4966.
//
// Solidity: function l1FeeScalar() view returns(uint256)
func (_L1Block *L1BlockCaller) L1FeeScalar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "l1FeeScalar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1FeeScalar is a free data retrieval call binding the contract method 0x9e8c4966.
//
// Solidity: function l1FeeScalar() view returns(uint256)
func (_L1Block *L1BlockSession) L1FeeScalar() (*big.Int, error) {
	return _L1Block.Contract.L1FeeScalar(&_L1Block.CallOpts)
}

// L1FeeScalar is a free data retrieval call binding the contract method 0x9e8c4966.
//
// Solidity: function l1FeeScalar() view returns(uint256)
func (_L1Block *L1BlockCallerSession) L1FeeScalar() (*big.Int, error) {
	return _L1Block.Contract.L1FeeScalar(&_L1Block.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint64)
func (_L1Block *L1BlockCaller) Number(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "number")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint64)
func (_L1Block *L1BlockSession) Number() (uint64, error) {
	return _L1Block.Contract.Number(&_L1Block.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint64)
func (_L1Block *L1BlockCallerSession) Number() (uint64, error) {
	return _L1Block.Contract.Number(&_L1Block.CallOpts)
}

// SequenceNumber is a free data retrieval call binding the contract method 0x64ca23ef.
//
// Solidity: function sequenceNumber() view returns(uint64)
func (_L1Block *L1BlockCaller) SequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "sequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SequenceNumber is a free data retrieval call binding the contract method 0x64ca23ef.
//
// Solidity: function sequenceNumber() view returns(uint64)
func (_L1Block *L1BlockSession) SequenceNumber() (uint64, error) {
	return _L1Block.Contract.SequenceNumber(&_L1Block.CallOpts)
}

// SequenceNumber is a free data retrieval call binding the contract method 0x64ca23ef.
//
// Solidity: function sequenceNumber() view returns(uint64)
func (_L1Block *L1BlockCallerSession) SequenceNumber() (uint64, error) {
	return _L1Block.Contract.SequenceNumber(&_L1Block.CallOpts)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint64)
func (_L1Block *L1BlockCaller) Timestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "timestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint64)
func (_L1Block *L1BlockSession) Timestamp() (uint64, error) {
	return _L1Block.Contract.Timestamp(&_L1Block.CallOpts)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint64)
func (_L1Block *L1BlockCallerSession) Timestamp() (uint64, error) {
	return _L1Block.Contract.Timestamp(&_L1Block.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_L1Block *L1BlockCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L1Block.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_L1Block *L1BlockSession) Version() (string, error) {
	return _L1Block.Contract.Version(&_L1Block.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_L1Block *L1BlockCallerSession) Version() (string, error) {
	return _L1Block.Contract.Version(&_L1Block.CallOpts)
}

// SetGasPayingToken is a paid mutator transaction binding the contract method 0x71cfaa3f.
//
// Solidity: function setGasPayingToken(address _token, uint8 _decimals, bytes32 _name, bytes32 _symbol) returns()
func (_L1Block *L1BlockTransactor) SetGasPayingToken(opts *bind.TransactOpts, _token common.Address, _decimals uint8, _name [32]byte, _symbol [32]byte) (*types.Transaction, error) {
	return _L1Block.contract.Transact(opts, "setGasPayingToken", _token, _decimals, _name, _symbol)
}

// SetGasPayingToken is a paid mutator transaction binding the contract method 0x71cfaa3f.
//
// Solidity: function setGasPayingToken(address _token, uint8 _decimals, bytes32 _name, bytes32 _symbol) returns()
func (_L1Block *L1BlockSession) SetGasPayingToken(_token common.Address, _decimals uint8, _name [32]byte, _symbol [32]byte) (*types.Transaction, error) {
	return _L1Block.Contract.SetGasPayingToken(&_L1Block.TransactOpts, _token, _decimals, _name, _symbol)
}

// SetGasPayingToken is a paid mutator transaction binding the contract method 0x71cfaa3f.
//
// Solidity: function setGasPayingToken(address _token, uint8 _decimals, bytes32 _name, bytes32 _symbol) returns()
func (_L1Block *L1BlockTransactorSession) SetGasPayingToken(_token common.Address, _decimals uint8, _name [32]byte, _symbol [32]byte) (*types.Transaction, error) {
	return _L1Block.Contract.SetGasPayingToken(&_L1Block.TransactOpts, _token, _decimals, _name, _symbol)
}

// SetL1BlockValues is a paid mutator transaction binding the contract method 0x015d8eb9.
//
// Solidity: function setL1BlockValues(uint64 _number, uint64 _timestamp, uint256 _basefee, bytes32 _hash, uint64 _sequenceNumber, bytes32 _batcherHash, uint256 _l1FeeOverhead, uint256 _l1FeeScalar) returns()
func (_L1Block *L1BlockTransactor) SetL1BlockValues(opts *bind.TransactOpts, _number uint64, _timestamp uint64, _basefee *big.Int, _hash [32]byte, _sequenceNumber uint64, _batcherHash [32]byte, _l1FeeOverhead *big.Int, _l1FeeScalar *big.Int) (*types.Transaction, error) {
	return _L1Block.contract.Transact(opts, "setL1BlockValues", _number, _timestamp, _basefee, _hash, _sequenceNumber, _batcherHash, _l1FeeOverhead, _l1FeeScalar)
}

// SetL1BlockValues is a paid mutator transaction binding the contract method 0x015d8eb9.
//
// Solidity: function setL1BlockValues(uint64 _number, uint64 _timestamp, uint256 _basefee, bytes32 _hash, uint64 _sequenceNumber, bytes32 _batcherHash, uint256 _l1FeeOverhead, uint256 _l1FeeScalar) returns()
func (_L1Block *L1BlockSession) SetL1BlockValues(_number uint64, _timestamp uint64, _basefee *big.Int, _hash [32]byte, _sequenceNumber uint64, _batcherHash [32]byte, _l1FeeOverhead *big.Int, _l1FeeScalar *big.Int) (*types.Transaction, error) {
	return _L1Block.Contract.SetL1BlockValues(&_L1Block.TransactOpts, _number, _timestamp, _basefee, _hash, _sequenceNumber, _batcherHash, _l1FeeOverhead, _l1FeeScalar)
}

// SetL1BlockValues is a paid mutator transaction binding the contract method 0x015d8eb9.
//
// Solidity: function setL1BlockValues(uint64 _number, uint64 _timestamp, uint256 _basefee, bytes32 _hash, uint64 _sequenceNumber, bytes32 _batcherHash, uint256 _l1FeeOverhead, uint256 _l1FeeScalar) returns()
func (_L1Block *L1BlockTransactorSession) SetL1BlockValues(_number uint64, _timestamp uint64, _basefee *big.Int, _hash [32]byte, _sequenceNumber uint64, _batcherHash [32]byte, _l1FeeOverhead *big.Int, _l1FeeScalar *big.Int) (*types.Transaction, error) {
	return _L1Block.Contract.SetL1BlockValues(&_L1Block.TransactOpts, _number, _timestamp, _basefee, _hash, _sequenceNumber, _batcherHash, _l1FeeOverhead, _l1FeeScalar)
}

// SetL1BlockValuesEcotone is a paid mutator transaction binding the contract method 0x440a5e20.
//
// Solidity: function setL1BlockValuesEcotone() returns()
func (_L1Block *L1BlockTransactor) SetL1BlockValuesEcotone(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Block.contract.Transact(opts, "setL1BlockValuesEcotone")
}

// SetL1BlockValuesEcotone is a paid mutator transaction binding the contract method 0x440a5e20.
//
// Solidity: function setL1BlockValuesEcotone() returns()
func (_L1Block *L1BlockSession) SetL1BlockValuesEcotone() (*types.Transaction, error) {
	return _L1Block.Contract.SetL1BlockValuesEcotone(&_L1Block.TransactOpts)
}

// SetL1BlockValuesEcotone is a paid mutator transaction binding the contract method 0x440a5e20.
//
// Solidity: function setL1BlockValuesEcotone() returns()
func (_L1Block *L1BlockTransactorSession) SetL1BlockValuesEcotone() (*types.Transaction, error) {
	return _L1Block.Contract.SetL1BlockValuesEcotone(&_L1Block.TransactOpts)
}

// L1BlockGasPayingTokenSetIterator is returned from FilterGasPayingTokenSet and is used to iterate over the raw logs and unpacked data for GasPayingTokenSet events raised by the L1Block contract.
type L1BlockGasPayingTokenSetIterator struct {
	Event *L1BlockGasPayingTokenSet // Event containing the contract specifics and raw log

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
func (it *L1BlockGasPayingTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1BlockGasPayingTokenSet)
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
		it.Event = new(L1BlockGasPayingTokenSet)
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
func (it *L1BlockGasPayingTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1BlockGasPayingTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1BlockGasPayingTokenSet represents a GasPayingTokenSet event raised by the L1Block contract.
type L1BlockGasPayingTokenSet struct {
	Token    common.Address
	Decimals uint8
	Name     [32]byte
	Symbol   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGasPayingTokenSet is a free log retrieval operation binding the contract event 0x10e43c4d58f3ef4edae7c1ca2e7f02d46b2cadbcc046737038527ed8486ffeb0.
//
// Solidity: event GasPayingTokenSet(address indexed token, uint8 indexed decimals, bytes32 name, bytes32 symbol)
func (_L1Block *L1BlockFilterer) FilterGasPayingTokenSet(opts *bind.FilterOpts, token []common.Address, decimals []uint8) (*L1BlockGasPayingTokenSetIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var decimalsRule []interface{}
	for _, decimalsItem := range decimals {
		decimalsRule = append(decimalsRule, decimalsItem)
	}

	logs, sub, err := _L1Block.contract.FilterLogs(opts, "GasPayingTokenSet", tokenRule, decimalsRule)
	if err != nil {
		return nil, err
	}
	return &L1BlockGasPayingTokenSetIterator{contract: _L1Block.contract, event: "GasPayingTokenSet", logs: logs, sub: sub}, nil
}

// WatchGasPayingTokenSet is a free log subscription operation binding the contract event 0x10e43c4d58f3ef4edae7c1ca2e7f02d46b2cadbcc046737038527ed8486ffeb0.
//
// Solidity: event GasPayingTokenSet(address indexed token, uint8 indexed decimals, bytes32 name, bytes32 symbol)
func (_L1Block *L1BlockFilterer) WatchGasPayingTokenSet(opts *bind.WatchOpts, sink chan<- *L1BlockGasPayingTokenSet, token []common.Address, decimals []uint8) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var decimalsRule []interface{}
	for _, decimalsItem := range decimals {
		decimalsRule = append(decimalsRule, decimalsItem)
	}

	logs, sub, err := _L1Block.contract.WatchLogs(opts, "GasPayingTokenSet", tokenRule, decimalsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1BlockGasPayingTokenSet)
				if err := _L1Block.contract.UnpackLog(event, "GasPayingTokenSet", log); err != nil {
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

// ParseGasPayingTokenSet is a log parse operation binding the contract event 0x10e43c4d58f3ef4edae7c1ca2e7f02d46b2cadbcc046737038527ed8486ffeb0.
//
// Solidity: event GasPayingTokenSet(address indexed token, uint8 indexed decimals, bytes32 name, bytes32 symbol)
func (_L1Block *L1BlockFilterer) ParseGasPayingTokenSet(log types.Log) (*L1BlockGasPayingTokenSet, error) {
	event := new(L1BlockGasPayingTokenSet)
	if err := _L1Block.contract.UnpackLog(event, "GasPayingTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
