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

// CanonicalStateChainCelestiaPointer is an auto generated low-level Go binding around an user-defined struct.
type CanonicalStateChainCelestiaPointer struct {
	Height     uint64
	ShareStart *big.Int
	ShareLen   uint16
}

// CanonicalStateChainHeader is an auto generated low-level Go binding around an user-defined struct.
type CanonicalStateChainHeader struct {
	Epoch            uint64
	L2Height         uint64
	PrevHash         [32]byte
	StateRoot        [32]byte
	CelestiaPointers []CanonicalStateChainCelestiaPointer
}

// CanonicalStateChainMetaData contains all meta data concerning the CanonicalStateChain contract.
var CanonicalStateChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenge\",\"type\":\"address\"}],\"name\":\"ChallengeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"publisher\",\"type\":\"address\"}],\"name\":\"PublisherChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"RolledBack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l2Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"shareStart\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"shareLen\",\"type\":\"uint16\"}],\"internalType\":\"structCanonicalStateChain.CelestiaPointer[]\",\"name\":\"celestiaPointers\",\"type\":\"tuple[]\"}],\"internalType\":\"structCanonicalStateChain.Header\",\"name\":\"_header\",\"type\":\"tuple\"}],\"name\":\"calculateHeaderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainHead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHead\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l2Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"shareStart\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"shareLen\",\"type\":\"uint16\"}],\"internalType\":\"structCanonicalStateChain.CelestiaPointer[]\",\"name\":\"celestiaPointers\",\"type\":\"tuple[]\"}],\"internalType\":\"structCanonicalStateChain.Header\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getHeaderByHash\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l2Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"shareStart\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"shareLen\",\"type\":\"uint16\"}],\"internalType\":\"structCanonicalStateChain.CelestiaPointer[]\",\"name\":\"celestiaPointers\",\"type\":\"tuple[]\"}],\"internalType\":\"structCanonicalStateChain.Header\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getHeaderByNum\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l2Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"shareStart\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"shareLen\",\"type\":\"uint16\"}],\"internalType\":\"structCanonicalStateChain.CelestiaPointer[]\",\"name\":\"celestiaPointers\",\"type\":\"tuple[]\"}],\"internalType\":\"structCanonicalStateChain.Header\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"headerMetadata\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"publisher\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_publisher\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l2Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"shareStart\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"shareLen\",\"type\":\"uint16\"}],\"internalType\":\"structCanonicalStateChain.CelestiaPointer[]\",\"name\":\"celestiaPointers\",\"type\":\"tuple[]\"}],\"internalType\":\"structCanonicalStateChain.Header\",\"name\":\"_header\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPointers\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publisher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l2Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"shareStart\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"shareLen\",\"type\":\"uint16\"}],\"internalType\":\"structCanonicalStateChain.CelestiaPointer[]\",\"name\":\"celestiaPointers\",\"type\":\"tuple[]\"}],\"internalType\":\"structCanonicalStateChain.Header\",\"name\":\"_header\",\"type\":\"tuple\"}],\"name\":\"pushBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_blockHash\",\"type\":\"bytes32\"}],\"name\":\"rollback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_challenge\",\"type\":\"address\"}],\"name\":\"setChallengeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_maxPointers\",\"type\":\"uint8\"}],\"name\":\"setMaxPointers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_publisher\",\"type\":\"address\"}],\"name\":\"setPublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CanonicalStateChain *CanonicalStateChainCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CanonicalStateChain *CanonicalStateChainSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CanonicalStateChain.Contract.UPGRADEINTERFACEVERSION(&_CanonicalStateChain.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CanonicalStateChain *CanonicalStateChainCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CanonicalStateChain.Contract.UPGRADEINTERFACEVERSION(&_CanonicalStateChain.CallOpts)
}

// CalculateHeaderHash is a free data retrieval call binding the contract method 0x44114703.
//
// Solidity: function calculateHeaderHash((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]) _header) pure returns(bytes32)
func (_CanonicalStateChain *CanonicalStateChainCaller) CalculateHeaderHash(opts *bind.CallOpts, _header CanonicalStateChainHeader) ([32]byte, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "calculateHeaderHash", _header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateHeaderHash is a free data retrieval call binding the contract method 0x44114703.
//
// Solidity: function calculateHeaderHash((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]) _header) pure returns(bytes32)
func (_CanonicalStateChain *CanonicalStateChainSession) CalculateHeaderHash(_header CanonicalStateChainHeader) ([32]byte, error) {
	return _CanonicalStateChain.Contract.CalculateHeaderHash(&_CanonicalStateChain.CallOpts, _header)
}

// CalculateHeaderHash is a free data retrieval call binding the contract method 0x44114703.
//
// Solidity: function calculateHeaderHash((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]) _header) pure returns(bytes32)
func (_CanonicalStateChain *CanonicalStateChainCallerSession) CalculateHeaderHash(_header CanonicalStateChainHeader) ([32]byte, error) {
	return _CanonicalStateChain.Contract.CalculateHeaderHash(&_CanonicalStateChain.CallOpts, _header)
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

// GetHead is a free data retrieval call binding the contract method 0xdc281aff.
//
// Solidity: function getHead() view returns((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]))
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
// Solidity: function getHead() view returns((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]))
func (_CanonicalStateChain *CanonicalStateChainSession) GetHead() (CanonicalStateChainHeader, error) {
	return _CanonicalStateChain.Contract.GetHead(&_CanonicalStateChain.CallOpts)
}

// GetHead is a free data retrieval call binding the contract method 0xdc281aff.
//
// Solidity: function getHead() view returns((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]))
func (_CanonicalStateChain *CanonicalStateChainCallerSession) GetHead() (CanonicalStateChainHeader, error) {
	return _CanonicalStateChain.Contract.GetHead(&_CanonicalStateChain.CallOpts)
}

// GetHeaderByHash is a free data retrieval call binding the contract method 0xb76971ce.
//
// Solidity: function getHeaderByHash(bytes32 _hash) view returns((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]))
func (_CanonicalStateChain *CanonicalStateChainCaller) GetHeaderByHash(opts *bind.CallOpts, _hash [32]byte) (CanonicalStateChainHeader, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "getHeaderByHash", _hash)

	if err != nil {
		return *new(CanonicalStateChainHeader), err
	}

	out0 := *abi.ConvertType(out[0], new(CanonicalStateChainHeader)).(*CanonicalStateChainHeader)

	return out0, err

}

// GetHeaderByHash is a free data retrieval call binding the contract method 0xb76971ce.
//
// Solidity: function getHeaderByHash(bytes32 _hash) view returns((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]))
func (_CanonicalStateChain *CanonicalStateChainSession) GetHeaderByHash(_hash [32]byte) (CanonicalStateChainHeader, error) {
	return _CanonicalStateChain.Contract.GetHeaderByHash(&_CanonicalStateChain.CallOpts, _hash)
}

// GetHeaderByHash is a free data retrieval call binding the contract method 0xb76971ce.
//
// Solidity: function getHeaderByHash(bytes32 _hash) view returns((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]))
func (_CanonicalStateChain *CanonicalStateChainCallerSession) GetHeaderByHash(_hash [32]byte) (CanonicalStateChainHeader, error) {
	return _CanonicalStateChain.Contract.GetHeaderByHash(&_CanonicalStateChain.CallOpts, _hash)
}

// GetHeaderByNum is a free data retrieval call binding the contract method 0x7255f37e.
//
// Solidity: function getHeaderByNum(uint256 _index) view returns((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]))
func (_CanonicalStateChain *CanonicalStateChainCaller) GetHeaderByNum(opts *bind.CallOpts, _index *big.Int) (CanonicalStateChainHeader, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "getHeaderByNum", _index)

	if err != nil {
		return *new(CanonicalStateChainHeader), err
	}

	out0 := *abi.ConvertType(out[0], new(CanonicalStateChainHeader)).(*CanonicalStateChainHeader)

	return out0, err

}

// GetHeaderByNum is a free data retrieval call binding the contract method 0x7255f37e.
//
// Solidity: function getHeaderByNum(uint256 _index) view returns((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]))
func (_CanonicalStateChain *CanonicalStateChainSession) GetHeaderByNum(_index *big.Int) (CanonicalStateChainHeader, error) {
	return _CanonicalStateChain.Contract.GetHeaderByNum(&_CanonicalStateChain.CallOpts, _index)
}

// GetHeaderByNum is a free data retrieval call binding the contract method 0x7255f37e.
//
// Solidity: function getHeaderByNum(uint256 _index) view returns((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]))
func (_CanonicalStateChain *CanonicalStateChainCallerSession) GetHeaderByNum(_index *big.Int) (CanonicalStateChainHeader, error) {
	return _CanonicalStateChain.Contract.GetHeaderByNum(&_CanonicalStateChain.CallOpts, _index)
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

// MaxPointers is a free data retrieval call binding the contract method 0x691bc74a.
//
// Solidity: function maxPointers() view returns(uint8)
func (_CanonicalStateChain *CanonicalStateChainCaller) MaxPointers(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "maxPointers")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MaxPointers is a free data retrieval call binding the contract method 0x691bc74a.
//
// Solidity: function maxPointers() view returns(uint8)
func (_CanonicalStateChain *CanonicalStateChainSession) MaxPointers() (uint8, error) {
	return _CanonicalStateChain.Contract.MaxPointers(&_CanonicalStateChain.CallOpts)
}

// MaxPointers is a free data retrieval call binding the contract method 0x691bc74a.
//
// Solidity: function maxPointers() view returns(uint8)
func (_CanonicalStateChain *CanonicalStateChainCallerSession) MaxPointers() (uint8, error) {
	return _CanonicalStateChain.Contract.MaxPointers(&_CanonicalStateChain.CallOpts)
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

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CanonicalStateChain *CanonicalStateChainCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CanonicalStateChain.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CanonicalStateChain *CanonicalStateChainSession) ProxiableUUID() ([32]byte, error) {
	return _CanonicalStateChain.Contract.ProxiableUUID(&_CanonicalStateChain.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CanonicalStateChain *CanonicalStateChainCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CanonicalStateChain.Contract.ProxiableUUID(&_CanonicalStateChain.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x41d3ecd8.
//
// Solidity: function initialize(address _publisher, (uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]) _header) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactor) Initialize(opts *bind.TransactOpts, _publisher common.Address, _header CanonicalStateChainHeader) (*types.Transaction, error) {
	return _CanonicalStateChain.contract.Transact(opts, "initialize", _publisher, _header)
}

// Initialize is a paid mutator transaction binding the contract method 0x41d3ecd8.
//
// Solidity: function initialize(address _publisher, (uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]) _header) returns()
func (_CanonicalStateChain *CanonicalStateChainSession) Initialize(_publisher common.Address, _header CanonicalStateChainHeader) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.Initialize(&_CanonicalStateChain.TransactOpts, _publisher, _header)
}

// Initialize is a paid mutator transaction binding the contract method 0x41d3ecd8.
//
// Solidity: function initialize(address _publisher, (uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]) _header) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactorSession) Initialize(_publisher common.Address, _header CanonicalStateChainHeader) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.Initialize(&_CanonicalStateChain.TransactOpts, _publisher, _header)
}

// PushBlock is a paid mutator transaction binding the contract method 0x8f74b739.
//
// Solidity: function pushBlock((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]) _header) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactor) PushBlock(opts *bind.TransactOpts, _header CanonicalStateChainHeader) (*types.Transaction, error) {
	return _CanonicalStateChain.contract.Transact(opts, "pushBlock", _header)
}

// PushBlock is a paid mutator transaction binding the contract method 0x8f74b739.
//
// Solidity: function pushBlock((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]) _header) returns()
func (_CanonicalStateChain *CanonicalStateChainSession) PushBlock(_header CanonicalStateChainHeader) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.PushBlock(&_CanonicalStateChain.TransactOpts, _header)
}

// PushBlock is a paid mutator transaction binding the contract method 0x8f74b739.
//
// Solidity: function pushBlock((uint64,uint64,bytes32,bytes32,(uint64,uint24,uint16)[]) _header) returns()
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

// Rollback is a paid mutator transaction binding the contract method 0x34c9bca5.
//
// Solidity: function rollback(uint256 _blockNumber, bytes32 _blockHash) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactor) Rollback(opts *bind.TransactOpts, _blockNumber *big.Int, _blockHash [32]byte) (*types.Transaction, error) {
	return _CanonicalStateChain.contract.Transact(opts, "rollback", _blockNumber, _blockHash)
}

// Rollback is a paid mutator transaction binding the contract method 0x34c9bca5.
//
// Solidity: function rollback(uint256 _blockNumber, bytes32 _blockHash) returns()
func (_CanonicalStateChain *CanonicalStateChainSession) Rollback(_blockNumber *big.Int, _blockHash [32]byte) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.Rollback(&_CanonicalStateChain.TransactOpts, _blockNumber, _blockHash)
}

// Rollback is a paid mutator transaction binding the contract method 0x34c9bca5.
//
// Solidity: function rollback(uint256 _blockNumber, bytes32 _blockHash) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactorSession) Rollback(_blockNumber *big.Int, _blockHash [32]byte) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.Rollback(&_CanonicalStateChain.TransactOpts, _blockNumber, _blockHash)
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

// SetMaxPointers is a paid mutator transaction binding the contract method 0x6d1005c5.
//
// Solidity: function setMaxPointers(uint8 _maxPointers) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactor) SetMaxPointers(opts *bind.TransactOpts, _maxPointers uint8) (*types.Transaction, error) {
	return _CanonicalStateChain.contract.Transact(opts, "setMaxPointers", _maxPointers)
}

// SetMaxPointers is a paid mutator transaction binding the contract method 0x6d1005c5.
//
// Solidity: function setMaxPointers(uint8 _maxPointers) returns()
func (_CanonicalStateChain *CanonicalStateChainSession) SetMaxPointers(_maxPointers uint8) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.SetMaxPointers(&_CanonicalStateChain.TransactOpts, _maxPointers)
}

// SetMaxPointers is a paid mutator transaction binding the contract method 0x6d1005c5.
//
// Solidity: function setMaxPointers(uint8 _maxPointers) returns()
func (_CanonicalStateChain *CanonicalStateChainTransactorSession) SetMaxPointers(_maxPointers uint8) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.SetMaxPointers(&_CanonicalStateChain.TransactOpts, _maxPointers)
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

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CanonicalStateChain *CanonicalStateChainTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CanonicalStateChain.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CanonicalStateChain *CanonicalStateChainSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.UpgradeToAndCall(&_CanonicalStateChain.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CanonicalStateChain *CanonicalStateChainTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CanonicalStateChain.Contract.UpgradeToAndCall(&_CanonicalStateChain.TransactOpts, newImplementation, data)
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

// CanonicalStateChainInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CanonicalStateChain contract.
type CanonicalStateChainInitializedIterator struct {
	Event *CanonicalStateChainInitialized // Event containing the contract specifics and raw log

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
func (it *CanonicalStateChainInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanonicalStateChainInitialized)
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
		it.Event = new(CanonicalStateChainInitialized)
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
func (it *CanonicalStateChainInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanonicalStateChainInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanonicalStateChainInitialized represents a Initialized event raised by the CanonicalStateChain contract.
type CanonicalStateChainInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CanonicalStateChain *CanonicalStateChainFilterer) FilterInitialized(opts *bind.FilterOpts) (*CanonicalStateChainInitializedIterator, error) {

	logs, sub, err := _CanonicalStateChain.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CanonicalStateChainInitializedIterator{contract: _CanonicalStateChain.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CanonicalStateChain *CanonicalStateChainFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CanonicalStateChainInitialized) (event.Subscription, error) {

	logs, sub, err := _CanonicalStateChain.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanonicalStateChainInitialized)
				if err := _CanonicalStateChain.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CanonicalStateChain *CanonicalStateChainFilterer) ParseInitialized(log types.Log) (*CanonicalStateChainInitialized, error) {
	event := new(CanonicalStateChainInitialized)
	if err := _CanonicalStateChain.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// CanonicalStateChainUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CanonicalStateChain contract.
type CanonicalStateChainUpgradedIterator struct {
	Event *CanonicalStateChainUpgraded // Event containing the contract specifics and raw log

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
func (it *CanonicalStateChainUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanonicalStateChainUpgraded)
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
		it.Event = new(CanonicalStateChainUpgraded)
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
func (it *CanonicalStateChainUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanonicalStateChainUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanonicalStateChainUpgraded represents a Upgraded event raised by the CanonicalStateChain contract.
type CanonicalStateChainUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CanonicalStateChain *CanonicalStateChainFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CanonicalStateChainUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CanonicalStateChain.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CanonicalStateChainUpgradedIterator{contract: _CanonicalStateChain.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CanonicalStateChain *CanonicalStateChainFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CanonicalStateChainUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CanonicalStateChain.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanonicalStateChainUpgraded)
				if err := _CanonicalStateChain.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CanonicalStateChain *CanonicalStateChainFilterer) ParseUpgraded(log types.Log) (*CanonicalStateChainUpgraded, error) {
	event := new(CanonicalStateChainUpgraded)
	if err := _CanonicalStateChain.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
