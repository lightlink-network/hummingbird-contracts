// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package canonicalstatechain

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

// CanonicalStateChainHeader is an auto generated low-level Go binding around an user-defined struct.
type CanonicalStateChainHeader struct {
	Epoch              uint64
	L2Height           uint64
	PrevHash           [32]byte
	TxRoot             [32]byte
	BlockRoot          [32]byte
	StateRoot          [32]byte
	CelestiaHeight     uint64
	CelestiaShareStart uint64
	CelestiaShareLen   uint64
}

// CanonicalStateChainMetaData contains all meta data concerning the CanonicalStateChain contract.
var CanonicalStateChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_publisher\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l2Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"celestiaHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"celestiaShareStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"celestiaShareLen\",\"type\":\"uint64\"}],\"internalType\":\"structCanonicalStateChain.Header\",\"name\":\"_header\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenge\",\"type\":\"address\"}],\"name\":\"ChallengeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"publisher\",\"type\":\"address\"}],\"name\":\"PublisherChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"RolledBack\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainHead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l2Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"celestiaHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"celestiaShareStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"celestiaShareLen\",\"type\":\"uint64\"}],\"internalType\":\"structCanonicalStateChain.Header\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHead\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l2Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"celestiaHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"celestiaShareStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"celestiaShareLen\",\"type\":\"uint64\"}],\"internalType\":\"structCanonicalStateChain.Header\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"headerMetadata\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"publisher\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"headers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l2Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"celestiaHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"celestiaShareStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"celestiaShareLen\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publisher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l2Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"celestiaHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"celestiaShareStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"celestiaShareLen\",\"type\":\"uint64\"}],\"internalType\":\"structCanonicalStateChain.Header\",\"name\":\"_header\",\"type\":\"tuple\"}],\"name\":\"pushBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"rollback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_challenge\",\"type\":\"address\"}],\"name\":\"setChallengeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_publisher\",\"type\":\"address\"}],\"name\":\"setPublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CanonicalStateChainABI is the input ABI used to generate the binding from.
// Deprecated: Use CanonicalStateChainMetaData.ABI instead.
var CanonicalStateChainABI = CanonicalStateChainMetaData.ABI

// CanonicalStateChain is an auto generated Go binding around an Ethereum contract.
type CanonicalStateChain struct {
	CanonicalStateChainCaller     // Read-only binding to the contract
	CanonicalStateChainTransactor // Write-only binding to the contract
	CanonicalStateChainFilterer   // Log filterer for contract events
}

// CanonicalStateChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type CanonicalStateChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalStateChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CanonicalStateChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalStateChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CanonicalStateChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalStateChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CanonicalStateChainSession struct {
	Contract     *CanonicalStateChain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CanonicalStateChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CanonicalStateChainCallerSession struct {
	Contract *CanonicalStateChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CanonicalStateChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CanonicalStateChainTransactorSession struct {
	Contract     *CanonicalStateChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CanonicalStateChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type CanonicalStateChainRaw struct {
	Contract *CanonicalStateChain // Generic contract binding to access the raw methods on
}

// CanonicalStateChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CanonicalStateChainCallerRaw struct {
	Contract *CanonicalStateChainCaller // Generic read-only contract binding to access the raw methods on
}

// CanonicalStateChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CanonicalStateChainTransactorRaw struct {
	Contract *CanonicalStateChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCanonicalStateChain creates a new instance of CanonicalStateChain, bound to a specific deployed contract.
func NewCanonicalStateChain(address common.Address, backend bind.ContractBackend) (*CanonicalStateChain, error) {
	contract, err := bindCanonicalStateChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CanonicalStateChain{CanonicalStateChainCaller: CanonicalStateChainCaller{contract: contract}, CanonicalStateChainTransactor: CanonicalStateChainTransactor{contract: contract}, CanonicalStateChainFilterer: CanonicalStateChainFilterer{contract: contract}}, nil
}

// NewCanonicalStateChainCaller creates a new read-only instance of CanonicalStateChain, bound to a specific deployed contract.
func NewCanonicalStateChainCaller(address common.Address, caller bind.ContractCaller) (*CanonicalStateChainCaller, error) {
	contract, err := bindCanonicalStateChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CanonicalStateChainCaller{contract: contract}, nil
}

// NewCanonicalStateChainTransactor creates a new write-only instance of CanonicalStateChain, bound to a specific deployed contract.
func NewCanonicalStateChainTransactor(address common.Address, transactor bind.ContractTransactor) (*CanonicalStateChainTransactor, error) {
	contract, err := bindCanonicalStateChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CanonicalStateChainTransactor{contract: contract}, nil
}

// NewCanonicalStateChainFilterer creates a new log filterer instance of CanonicalStateChain, bound to a specific deployed contract.
func NewCanonicalStateChainFilterer(address common.Address, filterer bind.ContractFilterer) (*CanonicalStateChainFilterer, error) {
	contract, err := bindCanonicalStateChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CanonicalStateChainFilterer{contract: contract}, nil
}

// bindCanonicalStateChain binds a generic wrapper to an already deployed contract.
func bindCanonicalStateChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CanonicalStateChainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanonicalStateChain *CanonicalStateChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CanonicalStateChain.Contract.CanonicalStateChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanonicalStateChain *CanonicalStateChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.CanonicalStateChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanonicalStateChain *CanonicalStateChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.CanonicalStateChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanonicalStateChain *CanonicalStateChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CanonicalStateChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanonicalStateChain *CanonicalStateChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanonicalStateChain *CanonicalStateChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.contract.Transact(opts, method, params...)
}

// Chain is a free data retrieval call binding the contract method 0x5852cc0c.
//
// Solidity: function chain(uint256 ) view returns(bytes32)
func (_CanonicalStateChain *CanonicalStateChainCaller) Chain(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "chain", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Chain is a free data retrieval call binding the contract method 0x5852cc0c.
//
// Solidity: function chain(uint256 ) view returns(bytes32)
func (_CanonicalStateChain *CanonicalStateChainSession) Chain(arg0 *big.Int) ([32]byte, error) {
	return _CanonicalStateChain.Contract.Chain(&_CanonicalStateChain.CallOpts, arg0)
}

// Chain is a free data retrieval call binding the contract method 0x5852cc0c.
//
// Solidity: function chain(uint256 ) view returns(bytes32)
func (_CanonicalStateChain *CanonicalStateChainCallerSession) Chain(arg0 *big.Int) ([32]byte, error) {
	return _CanonicalStateChain.Contract.Chain(&_CanonicalStateChain.CallOpts, arg0)
}

// ChainHead is a free data retrieval call binding the contract method 0x008f51c6.
//
// Solidity: function chainHead() view returns(uint256)
func (_CanonicalStateChain *CanonicalStateChainCaller) ChainHead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "chainHead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainHead is a free data retrieval call binding the contract method 0x008f51c6.
//
// Solidity: function chainHead() view returns(uint256)
func (_CanonicalStateChain *CanonicalStateChainSession) ChainHead() (*big.Int, error) {
	return _CanonicalStateChain.Contract.ChainHead(&_CanonicalStateChain.CallOpts)
}

// ChainHead is a free data retrieval call binding the contract method 0x008f51c6.
//
// Solidity: function chainHead() view returns(uint256)
func (_CanonicalStateChain *CanonicalStateChainCallerSession) ChainHead() (*big.Int, error) {
	return _CanonicalStateChain.Contract.ChainHead(&_CanonicalStateChain.CallOpts)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(address)
func (_CanonicalStateChain *CanonicalStateChainCaller) Challenge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "challenge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(address)
func (_CanonicalStateChain *CanonicalStateChainSession) Challenge() (common.Address, error) {
	return _CanonicalStateChain.Contract.Challenge(&_CanonicalStateChain.CallOpts)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(address)
func (_CanonicalStateChain *CanonicalStateChainCallerSession) Challenge() (common.Address, error) {
	return _CanonicalStateChain.Contract.Challenge(&_CanonicalStateChain.CallOpts)
}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 _index) view returns((uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64))
func (_CanonicalStateChain *CanonicalStateChainCaller) GetBlock(opts *bind.CallOpts, _index *big.Int) (CanonicalStateChainHeader, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "getBlock", _index)

	if err != nil {
		return *new(CanonicalStateChainHeader), err
	}

	out0 := *abi.ConvertType(out[0], new(CanonicalStateChainHeader)).(*CanonicalStateChainHeader)

	return out0, err

}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 _index) view returns((uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64))
func (_CanonicalStateChain *CanonicalStateChainSession) GetBlock(_index *big.Int) (CanonicalStateChainHeader, error) {
	return _CanonicalStateChain.Contract.GetBlock(&_CanonicalStateChain.CallOpts, _index)
}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 _index) view returns((uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64))
func (_CanonicalStateChain *CanonicalStateChainCallerSession) GetBlock(_index *big.Int) (CanonicalStateChainHeader, error) {
	return _CanonicalStateChain.Contract.GetBlock(&_CanonicalStateChain.CallOpts, _index)
}

// GetHead is a free data retrieval call binding the contract method 0xdc281aff.
//
// Solidity: function getHead() view returns((uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64))
func (_CanonicalStateChain *CanonicalStateChainCaller) GetHead(opts *bind.CallOpts) (CanonicalStateChainHeader, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "getHead")

	if err != nil {
		return *new(CanonicalStateChainHeader), err
	}

	out0 := *abi.ConvertType(out[0], new(CanonicalStateChainHeader)).(*CanonicalStateChainHeader)

	return out0, err

}

// GetHead is a free data retrieval call binding the contract method 0xdc281aff.
//
// Solidity: function getHead() view returns((uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64))
func (_CanonicalStateChain *CanonicalStateChainSession) GetHead() (CanonicalStateChainHeader, error) {
	return _CanonicalStateChain.Contract.GetHead(&_CanonicalStateChain.CallOpts)
}

// GetHead is a free data retrieval call binding the contract method 0xdc281aff.
//
// Solidity: function getHead() view returns((uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64))
func (_CanonicalStateChain *CanonicalStateChainCallerSession) GetHead() (CanonicalStateChainHeader, error) {
	return _CanonicalStateChain.Contract.GetHead(&_CanonicalStateChain.CallOpts)
}

// HeaderMetadata is a free data retrieval call binding the contract method 0x28a8d0e4.
//
// Solidity: function headerMetadata(bytes32 ) view returns(uint64 timestamp, address publisher)
func (_CanonicalStateChain *CanonicalStateChainCaller) HeaderMetadata(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Timestamp uint64
	Publisher common.Address
}, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "headerMetadata", arg0)

	outstruct := new(struct {
		Timestamp uint64
		Publisher common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Publisher = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// HeaderMetadata is a free data retrieval call binding the contract method 0x28a8d0e4.
//
// Solidity: function headerMetadata(bytes32 ) view returns(uint64 timestamp, address publisher)
func (_CanonicalStateChain *CanonicalStateChainSession) HeaderMetadata(arg0 [32]byte) (struct {
	Timestamp uint64
	Publisher common.Address
}, error) {
	return _CanonicalStateChain.Contract.HeaderMetadata(&_CanonicalStateChain.CallOpts, arg0)
}

// HeaderMetadata is a free data retrieval call binding the contract method 0x28a8d0e4.
//
// Solidity: function headerMetadata(bytes32 ) view returns(uint64 timestamp, address publisher)
func (_CanonicalStateChain *CanonicalStateChainCallerSession) HeaderMetadata(arg0 [32]byte) (struct {
	Timestamp uint64
	Publisher common.Address
}, error) {
	return _CanonicalStateChain.Contract.HeaderMetadata(&_CanonicalStateChain.CallOpts, arg0)
}

// Headers is a free data retrieval call binding the contract method 0x9e7f2700.
//
// Solidity: function headers(bytes32 ) view returns(uint64 epoch, uint64 l2Height, bytes32 prevHash, bytes32 txRoot, bytes32 blockRoot, bytes32 stateRoot, uint64 celestiaHeight, uint64 celestiaShareStart, uint64 celestiaShareLen)
func (_CanonicalStateChain *CanonicalStateChainCaller) Headers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Epoch              uint64
	L2Height           uint64
	PrevHash           [32]byte
	TxRoot             [32]byte
	BlockRoot          [32]byte
	StateRoot          [32]byte
	CelestiaHeight     uint64
	CelestiaShareStart uint64
	CelestiaShareLen   uint64
}, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "headers", arg0)

	outstruct := new(struct {
		Epoch              uint64
		L2Height           uint64
		PrevHash           [32]byte
		TxRoot             [32]byte
		BlockRoot          [32]byte
		StateRoot          [32]byte
		CelestiaHeight     uint64
		CelestiaShareStart uint64
		CelestiaShareLen   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.L2Height = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.PrevHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.TxRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.BlockRoot = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.StateRoot = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.CelestiaHeight = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.CelestiaShareStart = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.CelestiaShareLen = *abi.ConvertType(out[8], new(uint64)).(*uint64)

	return *outstruct, err

}

// Headers is a free data retrieval call binding the contract method 0x9e7f2700.
//
// Solidity: function headers(bytes32 ) view returns(uint64 epoch, uint64 l2Height, bytes32 prevHash, bytes32 txRoot, bytes32 blockRoot, bytes32 stateRoot, uint64 celestiaHeight, uint64 celestiaShareStart, uint64 celestiaShareLen)
func (_CanonicalStateChain *CanonicalStateChainSession) Headers(arg0 [32]byte) (struct {
	Epoch              uint64
	L2Height           uint64
	PrevHash           [32]byte
	TxRoot             [32]byte
	BlockRoot          [32]byte
	StateRoot          [32]byte
	CelestiaHeight     uint64
	CelestiaShareStart uint64
	CelestiaShareLen   uint64
}, error) {
	return _CanonicalStateChain.Contract.Headers(&_CanonicalStateChain.CallOpts, arg0)
}

// Headers is a free data retrieval call binding the contract method 0x9e7f2700.
//
// Solidity: function headers(bytes32 ) view returns(uint64 epoch, uint64 l2Height, bytes32 prevHash, bytes32 txRoot, bytes32 blockRoot, bytes32 stateRoot, uint64 celestiaHeight, uint64 celestiaShareStart, uint64 celestiaShareLen)
func (_CanonicalStateChain *CanonicalStateChainCallerSession) Headers(arg0 [32]byte) (struct {
	Epoch              uint64
	L2Height           uint64
	PrevHash           [32]byte
	TxRoot             [32]byte
	BlockRoot          [32]byte
	StateRoot          [32]byte
	CelestiaHeight     uint64
	CelestiaShareStart uint64
	CelestiaShareLen   uint64
}, error) {
	return _CanonicalStateChain.Contract.Headers(&_CanonicalStateChain.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CanonicalStateChain *CanonicalStateChainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CanonicalStateChain *CanonicalStateChainSession) Owner() (common.Address, error) {
	return _CanonicalStateChain.Contract.Owner(&_CanonicalStateChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CanonicalStateChain *CanonicalStateChainCallerSession) Owner() (common.Address, error) {
	return _CanonicalStateChain.Contract.Owner(&_CanonicalStateChain.CallOpts)
}

// Publisher is a free data retrieval call binding the contract method 0x8c72c54e.
//
// Solidity: function publisher() view returns(address)
func (_CanonicalStateChain *CanonicalStateChainCaller) Publisher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "publisher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Publisher is a free data retrieval call binding the contract method 0x8c72c54e.
//
// Solidity: function publisher() view returns(address)
func (_CanonicalStateChain *CanonicalStateChainSession) Publisher() (common.Address, error) {
	return _CanonicalStateChain.Contract.Publisher(&_CanonicalStateChain.CallOpts)
}

// Publisher is a free data retrieval call binding the contract method 0x8c72c54e.
//
// Solidity: function publisher() view returns(address)
func (_CanonicalStateChain *CanonicalStateChainCallerSession) Publisher() (common.Address, error) {
	return _CanonicalStateChain.Contract.Publisher(&_CanonicalStateChain.CallOpts)
}

// PushBlock is a paid mutator transaction binding the contract method 0xad0516dc.
//
// Solidity: function pushBlock((uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64) _header) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactor) PushBlock(opts *bind.TransactOpts, _header CanonicalStateChainHeader) (*types.Transaction, error) {
	return _CanonicalStateChain.contract.Transact(opts, "pushBlock", _header)
}

// PushBlock is a paid mutator transaction binding the contract method 0xad0516dc.
//
// Solidity: function pushBlock((uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64) _header) returns()
func (_CanonicalStateChain *CanonicalStateChainSession) PushBlock(_header CanonicalStateChainHeader) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.PushBlock(&_CanonicalStateChain.TransactOpts, _header)
}

// PushBlock is a paid mutator transaction binding the contract method 0xad0516dc.
//
// Solidity: function pushBlock((uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64) _header) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactorSession) PushBlock(_header CanonicalStateChainHeader) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.PushBlock(&_CanonicalStateChain.TransactOpts, _header)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CanonicalStateChain *CanonicalStateChainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanonicalStateChain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CanonicalStateChain *CanonicalStateChainSession) RenounceOwnership() (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.RenounceOwnership(&_CanonicalStateChain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CanonicalStateChain *CanonicalStateChainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.RenounceOwnership(&_CanonicalStateChain.TransactOpts)
}

// Rollback is a paid mutator transaction binding the contract method 0x0da9da20.
//
// Solidity: function rollback(uint256 _blockNumber) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactor) Rollback(opts *bind.TransactOpts, _blockNumber *big.Int) (*types.Transaction, error) {
	return _CanonicalStateChain.contract.Transact(opts, "rollback", _blockNumber)
}

// Rollback is a paid mutator transaction binding the contract method 0x0da9da20.
//
// Solidity: function rollback(uint256 _blockNumber) returns()
func (_CanonicalStateChain *CanonicalStateChainSession) Rollback(_blockNumber *big.Int) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.Rollback(&_CanonicalStateChain.TransactOpts, _blockNumber)
}

// Rollback is a paid mutator transaction binding the contract method 0x0da9da20.
//
// Solidity: function rollback(uint256 _blockNumber) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactorSession) Rollback(_blockNumber *big.Int) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.Rollback(&_CanonicalStateChain.TransactOpts, _blockNumber)
}

// SetChallengeContract is a paid mutator transaction binding the contract method 0xb37256b9.
//
// Solidity: function setChallengeContract(address _challenge) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactor) SetChallengeContract(opts *bind.TransactOpts, _challenge common.Address) (*types.Transaction, error) {
	return _CanonicalStateChain.contract.Transact(opts, "setChallengeContract", _challenge)
}

// SetChallengeContract is a paid mutator transaction binding the contract method 0xb37256b9.
//
// Solidity: function setChallengeContract(address _challenge) returns()
func (_CanonicalStateChain *CanonicalStateChainSession) SetChallengeContract(_challenge common.Address) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.SetChallengeContract(&_CanonicalStateChain.TransactOpts, _challenge)
}

// SetChallengeContract is a paid mutator transaction binding the contract method 0xb37256b9.
//
// Solidity: function setChallengeContract(address _challenge) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactorSession) SetChallengeContract(_challenge common.Address) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.SetChallengeContract(&_CanonicalStateChain.TransactOpts, _challenge)
}

// SetPublisher is a paid mutator transaction binding the contract method 0xcab63661.
//
// Solidity: function setPublisher(address _publisher) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactor) SetPublisher(opts *bind.TransactOpts, _publisher common.Address) (*types.Transaction, error) {
	return _CanonicalStateChain.contract.Transact(opts, "setPublisher", _publisher)
}

// SetPublisher is a paid mutator transaction binding the contract method 0xcab63661.
//
// Solidity: function setPublisher(address _publisher) returns()
func (_CanonicalStateChain *CanonicalStateChainSession) SetPublisher(_publisher common.Address) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.SetPublisher(&_CanonicalStateChain.TransactOpts, _publisher)
}

// SetPublisher is a paid mutator transaction binding the contract method 0xcab63661.
//
// Solidity: function setPublisher(address _publisher) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactorSession) SetPublisher(_publisher common.Address) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.SetPublisher(&_CanonicalStateChain.TransactOpts, _publisher)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CanonicalStateChain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CanonicalStateChain *CanonicalStateChainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.TransferOwnership(&_CanonicalStateChain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.TransferOwnership(&_CanonicalStateChain.TransactOpts, newOwner)
}

// CanonicalStateChainBlockAddedIterator is returned from FilterBlockAdded and is used to iterate over the raw logs and unpacked data for BlockAdded events raised by the CanonicalStateChain contract.
type CanonicalStateChainBlockAddedIterator struct {
	Event *CanonicalStateChainBlockAdded // Event containing the contract specifics and raw log

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
func (it *CanonicalStateChainBlockAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanonicalStateChainBlockAdded)
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
		it.Event = new(CanonicalStateChainBlockAdded)
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
func (it *CanonicalStateChainBlockAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanonicalStateChainBlockAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanonicalStateChainBlockAdded represents a BlockAdded event raised by the CanonicalStateChain contract.
type CanonicalStateChainBlockAdded struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockAdded is a free log retrieval operation binding the contract event 0xa37f9fb2f8e66e6e5746e84c33d55fc62d920182d22358f2adc6855d3ac4d437.
//
// Solidity: event BlockAdded(uint256 indexed blockNumber)
func (_CanonicalStateChain *CanonicalStateChainFilterer) FilterBlockAdded(opts *bind.FilterOpts, blockNumber []*big.Int) (*CanonicalStateChainBlockAddedIterator, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _CanonicalStateChain.contract.FilterLogs(opts, "BlockAdded", blockNumberRule)
	if err != nil {
		return nil, err
	}
	return &CanonicalStateChainBlockAddedIterator{contract: _CanonicalStateChain.contract, event: "BlockAdded", logs: logs, sub: sub}, nil
}

// WatchBlockAdded is a free log subscription operation binding the contract event 0xa37f9fb2f8e66e6e5746e84c33d55fc62d920182d22358f2adc6855d3ac4d437.
//
// Solidity: event BlockAdded(uint256 indexed blockNumber)
func (_CanonicalStateChain *CanonicalStateChainFilterer) WatchBlockAdded(opts *bind.WatchOpts, sink chan<- *CanonicalStateChainBlockAdded, blockNumber []*big.Int) (event.Subscription, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _CanonicalStateChain.contract.WatchLogs(opts, "BlockAdded", blockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanonicalStateChainBlockAdded)
				if err := _CanonicalStateChain.contract.UnpackLog(event, "BlockAdded", log); err != nil {
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

// ParseBlockAdded is a log parse operation binding the contract event 0xa37f9fb2f8e66e6e5746e84c33d55fc62d920182d22358f2adc6855d3ac4d437.
//
// Solidity: event BlockAdded(uint256 indexed blockNumber)
func (_CanonicalStateChain *CanonicalStateChainFilterer) ParseBlockAdded(log types.Log) (*CanonicalStateChainBlockAdded, error) {
	event := new(CanonicalStateChainBlockAdded)
	if err := _CanonicalStateChain.contract.UnpackLog(event, "BlockAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CanonicalStateChainChallengeChangedIterator is returned from FilterChallengeChanged and is used to iterate over the raw logs and unpacked data for ChallengeChanged events raised by the CanonicalStateChain contract.
type CanonicalStateChainChallengeChangedIterator struct {
	Event *CanonicalStateChainChallengeChanged // Event containing the contract specifics and raw log

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
func (it *CanonicalStateChainChallengeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanonicalStateChainChallengeChanged)
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
		it.Event = new(CanonicalStateChainChallengeChanged)
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
func (it *CanonicalStateChainChallengeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanonicalStateChainChallengeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanonicalStateChainChallengeChanged represents a ChallengeChanged event raised by the CanonicalStateChain contract.
type CanonicalStateChainChallengeChanged struct {
	Challenge common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChallengeChanged is a free log retrieval operation binding the contract event 0xe06eac444661557e3ac16a5251a66b82c3f985c3e3b15eac7ea4b4fac6eeac2c.
//
// Solidity: event ChallengeChanged(address indexed challenge)
func (_CanonicalStateChain *CanonicalStateChainFilterer) FilterChallengeChanged(opts *bind.FilterOpts, challenge []common.Address) (*CanonicalStateChainChallengeChangedIterator, error) {

	var challengeRule []interface{}
	for _, challengeItem := range challenge {
		challengeRule = append(challengeRule, challengeItem)
	}

	logs, sub, err := _CanonicalStateChain.contract.FilterLogs(opts, "ChallengeChanged", challengeRule)
	if err != nil {
		return nil, err
	}
	return &CanonicalStateChainChallengeChangedIterator{contract: _CanonicalStateChain.contract, event: "ChallengeChanged", logs: logs, sub: sub}, nil
}

// WatchChallengeChanged is a free log subscription operation binding the contract event 0xe06eac444661557e3ac16a5251a66b82c3f985c3e3b15eac7ea4b4fac6eeac2c.
//
// Solidity: event ChallengeChanged(address indexed challenge)
func (_CanonicalStateChain *CanonicalStateChainFilterer) WatchChallengeChanged(opts *bind.WatchOpts, sink chan<- *CanonicalStateChainChallengeChanged, challenge []common.Address) (event.Subscription, error) {

	var challengeRule []interface{}
	for _, challengeItem := range challenge {
		challengeRule = append(challengeRule, challengeItem)
	}

	logs, sub, err := _CanonicalStateChain.contract.WatchLogs(opts, "ChallengeChanged", challengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanonicalStateChainChallengeChanged)
				if err := _CanonicalStateChain.contract.UnpackLog(event, "ChallengeChanged", log); err != nil {
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

// ParseChallengeChanged is a log parse operation binding the contract event 0xe06eac444661557e3ac16a5251a66b82c3f985c3e3b15eac7ea4b4fac6eeac2c.
//
// Solidity: event ChallengeChanged(address indexed challenge)
func (_CanonicalStateChain *CanonicalStateChainFilterer) ParseChallengeChanged(log types.Log) (*CanonicalStateChainChallengeChanged, error) {
	event := new(CanonicalStateChainChallengeChanged)
	if err := _CanonicalStateChain.contract.UnpackLog(event, "ChallengeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CanonicalStateChainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CanonicalStateChain contract.
type CanonicalStateChainOwnershipTransferredIterator struct {
	Event *CanonicalStateChainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CanonicalStateChainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanonicalStateChainOwnershipTransferred)
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
		it.Event = new(CanonicalStateChainOwnershipTransferred)
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
func (it *CanonicalStateChainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanonicalStateChainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanonicalStateChainOwnershipTransferred represents a OwnershipTransferred event raised by the CanonicalStateChain contract.
type CanonicalStateChainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CanonicalStateChain *CanonicalStateChainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CanonicalStateChainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CanonicalStateChain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CanonicalStateChainOwnershipTransferredIterator{contract: _CanonicalStateChain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CanonicalStateChain *CanonicalStateChainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CanonicalStateChainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CanonicalStateChain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanonicalStateChainOwnershipTransferred)
				if err := _CanonicalStateChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CanonicalStateChain *CanonicalStateChainFilterer) ParseOwnershipTransferred(log types.Log) (*CanonicalStateChainOwnershipTransferred, error) {
	event := new(CanonicalStateChainOwnershipTransferred)
	if err := _CanonicalStateChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CanonicalStateChainPublisherChangedIterator is returned from FilterPublisherChanged and is used to iterate over the raw logs and unpacked data for PublisherChanged events raised by the CanonicalStateChain contract.
type CanonicalStateChainPublisherChangedIterator struct {
	Event *CanonicalStateChainPublisherChanged // Event containing the contract specifics and raw log

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
func (it *CanonicalStateChainPublisherChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanonicalStateChainPublisherChanged)
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
		it.Event = new(CanonicalStateChainPublisherChanged)
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
func (it *CanonicalStateChainPublisherChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanonicalStateChainPublisherChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanonicalStateChainPublisherChanged represents a PublisherChanged event raised by the CanonicalStateChain contract.
type CanonicalStateChainPublisherChanged struct {
	Publisher common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPublisherChanged is a free log retrieval operation binding the contract event 0x55eb99d77b0e1ed261c0a8d11f026f811b8af01455a2b45189bcc87b93dfbbb7.
//
// Solidity: event PublisherChanged(address indexed publisher)
func (_CanonicalStateChain *CanonicalStateChainFilterer) FilterPublisherChanged(opts *bind.FilterOpts, publisher []common.Address) (*CanonicalStateChainPublisherChangedIterator, error) {

	var publisherRule []interface{}
	for _, publisherItem := range publisher {
		publisherRule = append(publisherRule, publisherItem)
	}

	logs, sub, err := _CanonicalStateChain.contract.FilterLogs(opts, "PublisherChanged", publisherRule)
	if err != nil {
		return nil, err
	}
	return &CanonicalStateChainPublisherChangedIterator{contract: _CanonicalStateChain.contract, event: "PublisherChanged", logs: logs, sub: sub}, nil
}

// WatchPublisherChanged is a free log subscription operation binding the contract event 0x55eb99d77b0e1ed261c0a8d11f026f811b8af01455a2b45189bcc87b93dfbbb7.
//
// Solidity: event PublisherChanged(address indexed publisher)
func (_CanonicalStateChain *CanonicalStateChainFilterer) WatchPublisherChanged(opts *bind.WatchOpts, sink chan<- *CanonicalStateChainPublisherChanged, publisher []common.Address) (event.Subscription, error) {

	var publisherRule []interface{}
	for _, publisherItem := range publisher {
		publisherRule = append(publisherRule, publisherItem)
	}

	logs, sub, err := _CanonicalStateChain.contract.WatchLogs(opts, "PublisherChanged", publisherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanonicalStateChainPublisherChanged)
				if err := _CanonicalStateChain.contract.UnpackLog(event, "PublisherChanged", log); err != nil {
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

// ParsePublisherChanged is a log parse operation binding the contract event 0x55eb99d77b0e1ed261c0a8d11f026f811b8af01455a2b45189bcc87b93dfbbb7.
//
// Solidity: event PublisherChanged(address indexed publisher)
func (_CanonicalStateChain *CanonicalStateChainFilterer) ParsePublisherChanged(log types.Log) (*CanonicalStateChainPublisherChanged, error) {
	event := new(CanonicalStateChainPublisherChanged)
	if err := _CanonicalStateChain.contract.UnpackLog(event, "PublisherChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CanonicalStateChainRolledBackIterator is returned from FilterRolledBack and is used to iterate over the raw logs and unpacked data for RolledBack events raised by the CanonicalStateChain contract.
type CanonicalStateChainRolledBackIterator struct {
	Event *CanonicalStateChainRolledBack // Event containing the contract specifics and raw log

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
func (it *CanonicalStateChainRolledBackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanonicalStateChainRolledBack)
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
		it.Event = new(CanonicalStateChainRolledBack)
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
func (it *CanonicalStateChainRolledBackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanonicalStateChainRolledBackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanonicalStateChainRolledBack represents a RolledBack event raised by the CanonicalStateChain contract.
type CanonicalStateChainRolledBack struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRolledBack is a free log retrieval operation binding the contract event 0xbd56a1ce5e71ef906a2c86c43372d012f8ab2422ff19bfdba9b686ac0936f86f.
//
// Solidity: event RolledBack(uint256 indexed blockNumber)
func (_CanonicalStateChain *CanonicalStateChainFilterer) FilterRolledBack(opts *bind.FilterOpts, blockNumber []*big.Int) (*CanonicalStateChainRolledBackIterator, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _CanonicalStateChain.contract.FilterLogs(opts, "RolledBack", blockNumberRule)
	if err != nil {
		return nil, err
	}
	return &CanonicalStateChainRolledBackIterator{contract: _CanonicalStateChain.contract, event: "RolledBack", logs: logs, sub: sub}, nil
}

// WatchRolledBack is a free log subscription operation binding the contract event 0xbd56a1ce5e71ef906a2c86c43372d012f8ab2422ff19bfdba9b686ac0936f86f.
//
// Solidity: event RolledBack(uint256 indexed blockNumber)
func (_CanonicalStateChain *CanonicalStateChainFilterer) WatchRolledBack(opts *bind.WatchOpts, sink chan<- *CanonicalStateChainRolledBack, blockNumber []*big.Int) (event.Subscription, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _CanonicalStateChain.contract.WatchLogs(opts, "RolledBack", blockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanonicalStateChainRolledBack)
				if err := _CanonicalStateChain.contract.UnpackLog(event, "RolledBack", log); err != nil {
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

// ParseRolledBack is a log parse operation binding the contract event 0xbd56a1ce5e71ef906a2c86c43372d012f8ab2422ff19bfdba9b686ac0936f86f.
//
// Solidity: event RolledBack(uint256 indexed blockNumber)
func (_CanonicalStateChain *CanonicalStateChainFilterer) ParseRolledBack(log types.Log) (*CanonicalStateChainRolledBack, error) {
	event := new(CanonicalStateChainRolledBack)
	if err := _CanonicalStateChain.contract.UnpackLog(event, "RolledBack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
