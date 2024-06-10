// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package enterpriseportal

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

// EnterprisePortalMetaData contains all meta data concerning the EnterprisePortal contract.
var EnterprisePortalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"publicAddress\",\"type\":\"address\"}],\"name\":\"EnterpriseCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"}],\"name\":\"EnterpriseRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasDeducted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"whitelistedContract\",\"type\":\"address\"}],\"name\":\"WhitelistedContractAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"whitelistedContract\",\"type\":\"address\"}],\"name\":\"WhitelistedContractRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"whitelistedContract\",\"type\":\"address\"}],\"name\":\"addWhitelistedContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_publicAddress\",\"type\":\"address\"}],\"name\":\"createEnterprise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deductGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"}],\"name\":\"getEnterpriseDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"isContractWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextEnterpriseId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"}],\"name\":\"removeEnterprise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"whitelistedContract\",\"type\":\"address\"}],\"name\":\"removeWhitelistedContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// EnterprisePortalABI is the input ABI used to generate the binding from.
// Deprecated: Use EnterprisePortalMetaData.ABI instead.
var EnterprisePortalABI = EnterprisePortalMetaData.ABI

// EnterprisePortal is an auto generated Go binding around an Ethereum contract.
type EnterprisePortal struct {
	EnterprisePortalCaller     // Read-only binding to the contract
	EnterprisePortalTransactor // Write-only binding to the contract
	EnterprisePortalFilterer   // Log filterer for contract events
}

// EnterprisePortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnterprisePortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterprisePortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnterprisePortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterprisePortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnterprisePortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterprisePortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnterprisePortalSession struct {
	Contract     *EnterprisePortal // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnterprisePortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnterprisePortalCallerSession struct {
	Contract *EnterprisePortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EnterprisePortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnterprisePortalTransactorSession struct {
	Contract     *EnterprisePortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EnterprisePortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnterprisePortalRaw struct {
	Contract *EnterprisePortal // Generic contract binding to access the raw methods on
}

// EnterprisePortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnterprisePortalCallerRaw struct {
	Contract *EnterprisePortalCaller // Generic read-only contract binding to access the raw methods on
}

// EnterprisePortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnterprisePortalTransactorRaw struct {
	Contract *EnterprisePortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnterprisePortal creates a new instance of EnterprisePortal, bound to a specific deployed contract.
func NewEnterprisePortal(address common.Address, backend bind.ContractBackend) (*EnterprisePortal, error) {
	contract, err := bindEnterprisePortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnterprisePortal{EnterprisePortalCaller: EnterprisePortalCaller{contract: contract}, EnterprisePortalTransactor: EnterprisePortalTransactor{contract: contract}, EnterprisePortalFilterer: EnterprisePortalFilterer{contract: contract}}, nil
}

// NewEnterprisePortalCaller creates a new read-only instance of EnterprisePortal, bound to a specific deployed contract.
func NewEnterprisePortalCaller(address common.Address, caller bind.ContractCaller) (*EnterprisePortalCaller, error) {
	contract, err := bindEnterprisePortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnterprisePortalCaller{contract: contract}, nil
}

// NewEnterprisePortalTransactor creates a new write-only instance of EnterprisePortal, bound to a specific deployed contract.
func NewEnterprisePortalTransactor(address common.Address, transactor bind.ContractTransactor) (*EnterprisePortalTransactor, error) {
	contract, err := bindEnterprisePortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnterprisePortalTransactor{contract: contract}, nil
}

// NewEnterprisePortalFilterer creates a new log filterer instance of EnterprisePortal, bound to a specific deployed contract.
func NewEnterprisePortalFilterer(address common.Address, filterer bind.ContractFilterer) (*EnterprisePortalFilterer, error) {
	contract, err := bindEnterprisePortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnterprisePortalFilterer{contract: contract}, nil
}

// bindEnterprisePortal binds a generic wrapper to an already deployed contract.
func bindEnterprisePortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnterprisePortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterprisePortal *EnterprisePortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterprisePortal.Contract.EnterprisePortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterprisePortal *EnterprisePortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.EnterprisePortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterprisePortal *EnterprisePortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.EnterprisePortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterprisePortal *EnterprisePortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterprisePortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterprisePortal *EnterprisePortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterprisePortal *EnterprisePortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EnterprisePortal *EnterprisePortalCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EnterprisePortal.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EnterprisePortal *EnterprisePortalSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _EnterprisePortal.Contract.UPGRADEINTERFACEVERSION(&_EnterprisePortal.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EnterprisePortal *EnterprisePortalCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _EnterprisePortal.Contract.UPGRADEINTERFACEVERSION(&_EnterprisePortal.CallOpts)
}

// GetEnterpriseDetails is a free data retrieval call binding the contract method 0xb4906dd4.
//
// Solidity: function getEnterpriseDetails(uint256 enterpriseId) view returns(uint256, address, uint256)
func (_EnterprisePortal *EnterprisePortalCaller) GetEnterpriseDetails(opts *bind.CallOpts, enterpriseId *big.Int) (*big.Int, common.Address, *big.Int, error) {
	var out []interface{}
	err := _EnterprisePortal.contract.Call(opts, &out, "getEnterpriseDetails", enterpriseId)

	if err != nil {
		return *new(*big.Int), *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetEnterpriseDetails is a free data retrieval call binding the contract method 0xb4906dd4.
//
// Solidity: function getEnterpriseDetails(uint256 enterpriseId) view returns(uint256, address, uint256)
func (_EnterprisePortal *EnterprisePortalSession) GetEnterpriseDetails(enterpriseId *big.Int) (*big.Int, common.Address, *big.Int, error) {
	return _EnterprisePortal.Contract.GetEnterpriseDetails(&_EnterprisePortal.CallOpts, enterpriseId)
}

// GetEnterpriseDetails is a free data retrieval call binding the contract method 0xb4906dd4.
//
// Solidity: function getEnterpriseDetails(uint256 enterpriseId) view returns(uint256, address, uint256)
func (_EnterprisePortal *EnterprisePortalCallerSession) GetEnterpriseDetails(enterpriseId *big.Int) (*big.Int, common.Address, *big.Int, error) {
	return _EnterprisePortal.Contract.GetEnterpriseDetails(&_EnterprisePortal.CallOpts, enterpriseId)
}

// IsContractWhitelisted is a free data retrieval call binding the contract method 0xa2751e09.
//
// Solidity: function isContractWhitelisted(uint256 enterpriseId, address contractAddress) view returns(bool)
func (_EnterprisePortal *EnterprisePortalCaller) IsContractWhitelisted(opts *bind.CallOpts, enterpriseId *big.Int, contractAddress common.Address) (bool, error) {
	var out []interface{}
	err := _EnterprisePortal.contract.Call(opts, &out, "isContractWhitelisted", enterpriseId, contractAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContractWhitelisted is a free data retrieval call binding the contract method 0xa2751e09.
//
// Solidity: function isContractWhitelisted(uint256 enterpriseId, address contractAddress) view returns(bool)
func (_EnterprisePortal *EnterprisePortalSession) IsContractWhitelisted(enterpriseId *big.Int, contractAddress common.Address) (bool, error) {
	return _EnterprisePortal.Contract.IsContractWhitelisted(&_EnterprisePortal.CallOpts, enterpriseId, contractAddress)
}

// IsContractWhitelisted is a free data retrieval call binding the contract method 0xa2751e09.
//
// Solidity: function isContractWhitelisted(uint256 enterpriseId, address contractAddress) view returns(bool)
func (_EnterprisePortal *EnterprisePortalCallerSession) IsContractWhitelisted(enterpriseId *big.Int, contractAddress common.Address) (bool, error) {
	return _EnterprisePortal.Contract.IsContractWhitelisted(&_EnterprisePortal.CallOpts, enterpriseId, contractAddress)
}

// NextEnterpriseId is a free data retrieval call binding the contract method 0xe7ae3131.
//
// Solidity: function nextEnterpriseId() view returns(uint256)
func (_EnterprisePortal *EnterprisePortalCaller) NextEnterpriseId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnterprisePortal.contract.Call(opts, &out, "nextEnterpriseId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextEnterpriseId is a free data retrieval call binding the contract method 0xe7ae3131.
//
// Solidity: function nextEnterpriseId() view returns(uint256)
func (_EnterprisePortal *EnterprisePortalSession) NextEnterpriseId() (*big.Int, error) {
	return _EnterprisePortal.Contract.NextEnterpriseId(&_EnterprisePortal.CallOpts)
}

// NextEnterpriseId is a free data retrieval call binding the contract method 0xe7ae3131.
//
// Solidity: function nextEnterpriseId() view returns(uint256)
func (_EnterprisePortal *EnterprisePortalCallerSession) NextEnterpriseId() (*big.Int, error) {
	return _EnterprisePortal.Contract.NextEnterpriseId(&_EnterprisePortal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnterprisePortal *EnterprisePortalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterprisePortal.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnterprisePortal *EnterprisePortalSession) Owner() (common.Address, error) {
	return _EnterprisePortal.Contract.Owner(&_EnterprisePortal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnterprisePortal *EnterprisePortalCallerSession) Owner() (common.Address, error) {
	return _EnterprisePortal.Contract.Owner(&_EnterprisePortal.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EnterprisePortal *EnterprisePortalCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnterprisePortal.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EnterprisePortal *EnterprisePortalSession) ProxiableUUID() ([32]byte, error) {
	return _EnterprisePortal.Contract.ProxiableUUID(&_EnterprisePortal.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EnterprisePortal *EnterprisePortalCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EnterprisePortal.Contract.ProxiableUUID(&_EnterprisePortal.CallOpts)
}

// AddGas is a paid mutator transaction binding the contract method 0x4d0b16b5.
//
// Solidity: function addGas(uint256 enterpriseId, uint256 amount) returns()
func (_EnterprisePortal *EnterprisePortalTransactor) AddGas(opts *bind.TransactOpts, enterpriseId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _EnterprisePortal.contract.Transact(opts, "addGas", enterpriseId, amount)
}

// AddGas is a paid mutator transaction binding the contract method 0x4d0b16b5.
//
// Solidity: function addGas(uint256 enterpriseId, uint256 amount) returns()
func (_EnterprisePortal *EnterprisePortalSession) AddGas(enterpriseId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.AddGas(&_EnterprisePortal.TransactOpts, enterpriseId, amount)
}

// AddGas is a paid mutator transaction binding the contract method 0x4d0b16b5.
//
// Solidity: function addGas(uint256 enterpriseId, uint256 amount) returns()
func (_EnterprisePortal *EnterprisePortalTransactorSession) AddGas(enterpriseId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.AddGas(&_EnterprisePortal.TransactOpts, enterpriseId, amount)
}

// AddWhitelistedContract is a paid mutator transaction binding the contract method 0x25b02a9e.
//
// Solidity: function addWhitelistedContract(uint256 enterpriseId, address whitelistedContract) returns()
func (_EnterprisePortal *EnterprisePortalTransactor) AddWhitelistedContract(opts *bind.TransactOpts, enterpriseId *big.Int, whitelistedContract common.Address) (*types.Transaction, error) {
	return _EnterprisePortal.contract.Transact(opts, "addWhitelistedContract", enterpriseId, whitelistedContract)
}

// AddWhitelistedContract is a paid mutator transaction binding the contract method 0x25b02a9e.
//
// Solidity: function addWhitelistedContract(uint256 enterpriseId, address whitelistedContract) returns()
func (_EnterprisePortal *EnterprisePortalSession) AddWhitelistedContract(enterpriseId *big.Int, whitelistedContract common.Address) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.AddWhitelistedContract(&_EnterprisePortal.TransactOpts, enterpriseId, whitelistedContract)
}

// AddWhitelistedContract is a paid mutator transaction binding the contract method 0x25b02a9e.
//
// Solidity: function addWhitelistedContract(uint256 enterpriseId, address whitelistedContract) returns()
func (_EnterprisePortal *EnterprisePortalTransactorSession) AddWhitelistedContract(enterpriseId *big.Int, whitelistedContract common.Address) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.AddWhitelistedContract(&_EnterprisePortal.TransactOpts, enterpriseId, whitelistedContract)
}

// CreateEnterprise is a paid mutator transaction binding the contract method 0xd25de18b.
//
// Solidity: function createEnterprise(address _publicAddress) returns()
func (_EnterprisePortal *EnterprisePortalTransactor) CreateEnterprise(opts *bind.TransactOpts, _publicAddress common.Address) (*types.Transaction, error) {
	return _EnterprisePortal.contract.Transact(opts, "createEnterprise", _publicAddress)
}

// CreateEnterprise is a paid mutator transaction binding the contract method 0xd25de18b.
//
// Solidity: function createEnterprise(address _publicAddress) returns()
func (_EnterprisePortal *EnterprisePortalSession) CreateEnterprise(_publicAddress common.Address) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.CreateEnterprise(&_EnterprisePortal.TransactOpts, _publicAddress)
}

// CreateEnterprise is a paid mutator transaction binding the contract method 0xd25de18b.
//
// Solidity: function createEnterprise(address _publicAddress) returns()
func (_EnterprisePortal *EnterprisePortalTransactorSession) CreateEnterprise(_publicAddress common.Address) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.CreateEnterprise(&_EnterprisePortal.TransactOpts, _publicAddress)
}

// DeductGas is a paid mutator transaction binding the contract method 0x2605aeaa.
//
// Solidity: function deductGas(uint256 enterpriseId, uint256 amount) returns()
func (_EnterprisePortal *EnterprisePortalTransactor) DeductGas(opts *bind.TransactOpts, enterpriseId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _EnterprisePortal.contract.Transact(opts, "deductGas", enterpriseId, amount)
}

// DeductGas is a paid mutator transaction binding the contract method 0x2605aeaa.
//
// Solidity: function deductGas(uint256 enterpriseId, uint256 amount) returns()
func (_EnterprisePortal *EnterprisePortalSession) DeductGas(enterpriseId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.DeductGas(&_EnterprisePortal.TransactOpts, enterpriseId, amount)
}

// DeductGas is a paid mutator transaction binding the contract method 0x2605aeaa.
//
// Solidity: function deductGas(uint256 enterpriseId, uint256 amount) returns()
func (_EnterprisePortal *EnterprisePortalTransactorSession) DeductGas(enterpriseId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.DeductGas(&_EnterprisePortal.TransactOpts, enterpriseId, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_EnterprisePortal *EnterprisePortalTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterprisePortal.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_EnterprisePortal *EnterprisePortalSession) Initialize() (*types.Transaction, error) {
	return _EnterprisePortal.Contract.Initialize(&_EnterprisePortal.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_EnterprisePortal *EnterprisePortalTransactorSession) Initialize() (*types.Transaction, error) {
	return _EnterprisePortal.Contract.Initialize(&_EnterprisePortal.TransactOpts)
}

// RemoveEnterprise is a paid mutator transaction binding the contract method 0x7aaeeb5f.
//
// Solidity: function removeEnterprise(uint256 enterpriseId) returns()
func (_EnterprisePortal *EnterprisePortalTransactor) RemoveEnterprise(opts *bind.TransactOpts, enterpriseId *big.Int) (*types.Transaction, error) {
	return _EnterprisePortal.contract.Transact(opts, "removeEnterprise", enterpriseId)
}

// RemoveEnterprise is a paid mutator transaction binding the contract method 0x7aaeeb5f.
//
// Solidity: function removeEnterprise(uint256 enterpriseId) returns()
func (_EnterprisePortal *EnterprisePortalSession) RemoveEnterprise(enterpriseId *big.Int) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.RemoveEnterprise(&_EnterprisePortal.TransactOpts, enterpriseId)
}

// RemoveEnterprise is a paid mutator transaction binding the contract method 0x7aaeeb5f.
//
// Solidity: function removeEnterprise(uint256 enterpriseId) returns()
func (_EnterprisePortal *EnterprisePortalTransactorSession) RemoveEnterprise(enterpriseId *big.Int) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.RemoveEnterprise(&_EnterprisePortal.TransactOpts, enterpriseId)
}

// RemoveWhitelistedContract is a paid mutator transaction binding the contract method 0xae4bc351.
//
// Solidity: function removeWhitelistedContract(uint256 enterpriseId, address whitelistedContract) returns()
func (_EnterprisePortal *EnterprisePortalTransactor) RemoveWhitelistedContract(opts *bind.TransactOpts, enterpriseId *big.Int, whitelistedContract common.Address) (*types.Transaction, error) {
	return _EnterprisePortal.contract.Transact(opts, "removeWhitelistedContract", enterpriseId, whitelistedContract)
}

// RemoveWhitelistedContract is a paid mutator transaction binding the contract method 0xae4bc351.
//
// Solidity: function removeWhitelistedContract(uint256 enterpriseId, address whitelistedContract) returns()
func (_EnterprisePortal *EnterprisePortalSession) RemoveWhitelistedContract(enterpriseId *big.Int, whitelistedContract common.Address) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.RemoveWhitelistedContract(&_EnterprisePortal.TransactOpts, enterpriseId, whitelistedContract)
}

// RemoveWhitelistedContract is a paid mutator transaction binding the contract method 0xae4bc351.
//
// Solidity: function removeWhitelistedContract(uint256 enterpriseId, address whitelistedContract) returns()
func (_EnterprisePortal *EnterprisePortalTransactorSession) RemoveWhitelistedContract(enterpriseId *big.Int, whitelistedContract common.Address) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.RemoveWhitelistedContract(&_EnterprisePortal.TransactOpts, enterpriseId, whitelistedContract)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnterprisePortal *EnterprisePortalTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterprisePortal.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnterprisePortal *EnterprisePortalSession) RenounceOwnership() (*types.Transaction, error) {
	return _EnterprisePortal.Contract.RenounceOwnership(&_EnterprisePortal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnterprisePortal *EnterprisePortalTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EnterprisePortal.Contract.RenounceOwnership(&_EnterprisePortal.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnterprisePortal *EnterprisePortalTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EnterprisePortal.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnterprisePortal *EnterprisePortalSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.TransferOwnership(&_EnterprisePortal.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnterprisePortal *EnterprisePortalTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.TransferOwnership(&_EnterprisePortal.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EnterprisePortal *EnterprisePortalTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EnterprisePortal.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EnterprisePortal *EnterprisePortalSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.UpgradeToAndCall(&_EnterprisePortal.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EnterprisePortal *EnterprisePortalTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EnterprisePortal.Contract.UpgradeToAndCall(&_EnterprisePortal.TransactOpts, newImplementation, data)
}

// EnterprisePortalEnterpriseCreatedIterator is returned from FilterEnterpriseCreated and is used to iterate over the raw logs and unpacked data for EnterpriseCreated events raised by the EnterprisePortal contract.
type EnterprisePortalEnterpriseCreatedIterator struct {
	Event *EnterprisePortalEnterpriseCreated // Event containing the contract specifics and raw log

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
func (it *EnterprisePortalEnterpriseCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterprisePortalEnterpriseCreated)
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
		it.Event = new(EnterprisePortalEnterpriseCreated)
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
func (it *EnterprisePortalEnterpriseCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterprisePortalEnterpriseCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterprisePortalEnterpriseCreated represents a EnterpriseCreated event raised by the EnterprisePortal contract.
type EnterprisePortalEnterpriseCreated struct {
	EnterpriseId  *big.Int
	PublicAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEnterpriseCreated is a free log retrieval operation binding the contract event 0xe32c44b8d8aa648af9461aa9eadff5cc889895b0fd67b143482ba7b68c5a0ca8.
//
// Solidity: event EnterpriseCreated(uint256 indexed enterpriseId, address indexed publicAddress)
func (_EnterprisePortal *EnterprisePortalFilterer) FilterEnterpriseCreated(opts *bind.FilterOpts, enterpriseId []*big.Int, publicAddress []common.Address) (*EnterprisePortalEnterpriseCreatedIterator, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}
	var publicAddressRule []interface{}
	for _, publicAddressItem := range publicAddress {
		publicAddressRule = append(publicAddressRule, publicAddressItem)
	}

	logs, sub, err := _EnterprisePortal.contract.FilterLogs(opts, "EnterpriseCreated", enterpriseIdRule, publicAddressRule)
	if err != nil {
		return nil, err
	}
	return &EnterprisePortalEnterpriseCreatedIterator{contract: _EnterprisePortal.contract, event: "EnterpriseCreated", logs: logs, sub: sub}, nil
}

// WatchEnterpriseCreated is a free log subscription operation binding the contract event 0xe32c44b8d8aa648af9461aa9eadff5cc889895b0fd67b143482ba7b68c5a0ca8.
//
// Solidity: event EnterpriseCreated(uint256 indexed enterpriseId, address indexed publicAddress)
func (_EnterprisePortal *EnterprisePortalFilterer) WatchEnterpriseCreated(opts *bind.WatchOpts, sink chan<- *EnterprisePortalEnterpriseCreated, enterpriseId []*big.Int, publicAddress []common.Address) (event.Subscription, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}
	var publicAddressRule []interface{}
	for _, publicAddressItem := range publicAddress {
		publicAddressRule = append(publicAddressRule, publicAddressItem)
	}

	logs, sub, err := _EnterprisePortal.contract.WatchLogs(opts, "EnterpriseCreated", enterpriseIdRule, publicAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterprisePortalEnterpriseCreated)
				if err := _EnterprisePortal.contract.UnpackLog(event, "EnterpriseCreated", log); err != nil {
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

// ParseEnterpriseCreated is a log parse operation binding the contract event 0xe32c44b8d8aa648af9461aa9eadff5cc889895b0fd67b143482ba7b68c5a0ca8.
//
// Solidity: event EnterpriseCreated(uint256 indexed enterpriseId, address indexed publicAddress)
func (_EnterprisePortal *EnterprisePortalFilterer) ParseEnterpriseCreated(log types.Log) (*EnterprisePortalEnterpriseCreated, error) {
	event := new(EnterprisePortalEnterpriseCreated)
	if err := _EnterprisePortal.contract.UnpackLog(event, "EnterpriseCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterprisePortalEnterpriseRemovedIterator is returned from FilterEnterpriseRemoved and is used to iterate over the raw logs and unpacked data for EnterpriseRemoved events raised by the EnterprisePortal contract.
type EnterprisePortalEnterpriseRemovedIterator struct {
	Event *EnterprisePortalEnterpriseRemoved // Event containing the contract specifics and raw log

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
func (it *EnterprisePortalEnterpriseRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterprisePortalEnterpriseRemoved)
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
		it.Event = new(EnterprisePortalEnterpriseRemoved)
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
func (it *EnterprisePortalEnterpriseRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterprisePortalEnterpriseRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterprisePortalEnterpriseRemoved represents a EnterpriseRemoved event raised by the EnterprisePortal contract.
type EnterprisePortalEnterpriseRemoved struct {
	EnterpriseId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEnterpriseRemoved is a free log retrieval operation binding the contract event 0x7595592180477b9469ce7148e840767b1943640315934888bd6d8e7ecaa57460.
//
// Solidity: event EnterpriseRemoved(uint256 indexed enterpriseId)
func (_EnterprisePortal *EnterprisePortalFilterer) FilterEnterpriseRemoved(opts *bind.FilterOpts, enterpriseId []*big.Int) (*EnterprisePortalEnterpriseRemovedIterator, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}

	logs, sub, err := _EnterprisePortal.contract.FilterLogs(opts, "EnterpriseRemoved", enterpriseIdRule)
	if err != nil {
		return nil, err
	}
	return &EnterprisePortalEnterpriseRemovedIterator{contract: _EnterprisePortal.contract, event: "EnterpriseRemoved", logs: logs, sub: sub}, nil
}

// WatchEnterpriseRemoved is a free log subscription operation binding the contract event 0x7595592180477b9469ce7148e840767b1943640315934888bd6d8e7ecaa57460.
//
// Solidity: event EnterpriseRemoved(uint256 indexed enterpriseId)
func (_EnterprisePortal *EnterprisePortalFilterer) WatchEnterpriseRemoved(opts *bind.WatchOpts, sink chan<- *EnterprisePortalEnterpriseRemoved, enterpriseId []*big.Int) (event.Subscription, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}

	logs, sub, err := _EnterprisePortal.contract.WatchLogs(opts, "EnterpriseRemoved", enterpriseIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterprisePortalEnterpriseRemoved)
				if err := _EnterprisePortal.contract.UnpackLog(event, "EnterpriseRemoved", log); err != nil {
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

// ParseEnterpriseRemoved is a log parse operation binding the contract event 0x7595592180477b9469ce7148e840767b1943640315934888bd6d8e7ecaa57460.
//
// Solidity: event EnterpriseRemoved(uint256 indexed enterpriseId)
func (_EnterprisePortal *EnterprisePortalFilterer) ParseEnterpriseRemoved(log types.Log) (*EnterprisePortalEnterpriseRemoved, error) {
	event := new(EnterprisePortalEnterpriseRemoved)
	if err := _EnterprisePortal.contract.UnpackLog(event, "EnterpriseRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterprisePortalGasAddedIterator is returned from FilterGasAdded and is used to iterate over the raw logs and unpacked data for GasAdded events raised by the EnterprisePortal contract.
type EnterprisePortalGasAddedIterator struct {
	Event *EnterprisePortalGasAdded // Event containing the contract specifics and raw log

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
func (it *EnterprisePortalGasAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterprisePortalGasAdded)
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
		it.Event = new(EnterprisePortalGasAdded)
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
func (it *EnterprisePortalGasAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterprisePortalGasAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterprisePortalGasAdded represents a GasAdded event raised by the EnterprisePortal contract.
type EnterprisePortalGasAdded struct {
	EnterpriseId *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGasAdded is a free log retrieval operation binding the contract event 0x1bae0ffe945f42799e53fe678eb5ed2c4d1c33d42ee06931ea5d74aaacb2011d.
//
// Solidity: event GasAdded(uint256 indexed enterpriseId, uint256 amount)
func (_EnterprisePortal *EnterprisePortalFilterer) FilterGasAdded(opts *bind.FilterOpts, enterpriseId []*big.Int) (*EnterprisePortalGasAddedIterator, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}

	logs, sub, err := _EnterprisePortal.contract.FilterLogs(opts, "GasAdded", enterpriseIdRule)
	if err != nil {
		return nil, err
	}
	return &EnterprisePortalGasAddedIterator{contract: _EnterprisePortal.contract, event: "GasAdded", logs: logs, sub: sub}, nil
}

// WatchGasAdded is a free log subscription operation binding the contract event 0x1bae0ffe945f42799e53fe678eb5ed2c4d1c33d42ee06931ea5d74aaacb2011d.
//
// Solidity: event GasAdded(uint256 indexed enterpriseId, uint256 amount)
func (_EnterprisePortal *EnterprisePortalFilterer) WatchGasAdded(opts *bind.WatchOpts, sink chan<- *EnterprisePortalGasAdded, enterpriseId []*big.Int) (event.Subscription, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}

	logs, sub, err := _EnterprisePortal.contract.WatchLogs(opts, "GasAdded", enterpriseIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterprisePortalGasAdded)
				if err := _EnterprisePortal.contract.UnpackLog(event, "GasAdded", log); err != nil {
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

// ParseGasAdded is a log parse operation binding the contract event 0x1bae0ffe945f42799e53fe678eb5ed2c4d1c33d42ee06931ea5d74aaacb2011d.
//
// Solidity: event GasAdded(uint256 indexed enterpriseId, uint256 amount)
func (_EnterprisePortal *EnterprisePortalFilterer) ParseGasAdded(log types.Log) (*EnterprisePortalGasAdded, error) {
	event := new(EnterprisePortalGasAdded)
	if err := _EnterprisePortal.contract.UnpackLog(event, "GasAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterprisePortalGasDeductedIterator is returned from FilterGasDeducted and is used to iterate over the raw logs and unpacked data for GasDeducted events raised by the EnterprisePortal contract.
type EnterprisePortalGasDeductedIterator struct {
	Event *EnterprisePortalGasDeducted // Event containing the contract specifics and raw log

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
func (it *EnterprisePortalGasDeductedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterprisePortalGasDeducted)
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
		it.Event = new(EnterprisePortalGasDeducted)
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
func (it *EnterprisePortalGasDeductedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterprisePortalGasDeductedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterprisePortalGasDeducted represents a GasDeducted event raised by the EnterprisePortal contract.
type EnterprisePortalGasDeducted struct {
	EnterpriseId *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGasDeducted is a free log retrieval operation binding the contract event 0x7bcaec5d68e62be9724911d03486d3007cea9f870b477fcfe7d834234730cda0.
//
// Solidity: event GasDeducted(uint256 indexed enterpriseId, uint256 amount)
func (_EnterprisePortal *EnterprisePortalFilterer) FilterGasDeducted(opts *bind.FilterOpts, enterpriseId []*big.Int) (*EnterprisePortalGasDeductedIterator, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}

	logs, sub, err := _EnterprisePortal.contract.FilterLogs(opts, "GasDeducted", enterpriseIdRule)
	if err != nil {
		return nil, err
	}
	return &EnterprisePortalGasDeductedIterator{contract: _EnterprisePortal.contract, event: "GasDeducted", logs: logs, sub: sub}, nil
}

// WatchGasDeducted is a free log subscription operation binding the contract event 0x7bcaec5d68e62be9724911d03486d3007cea9f870b477fcfe7d834234730cda0.
//
// Solidity: event GasDeducted(uint256 indexed enterpriseId, uint256 amount)
func (_EnterprisePortal *EnterprisePortalFilterer) WatchGasDeducted(opts *bind.WatchOpts, sink chan<- *EnterprisePortalGasDeducted, enterpriseId []*big.Int) (event.Subscription, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}

	logs, sub, err := _EnterprisePortal.contract.WatchLogs(opts, "GasDeducted", enterpriseIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterprisePortalGasDeducted)
				if err := _EnterprisePortal.contract.UnpackLog(event, "GasDeducted", log); err != nil {
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

// ParseGasDeducted is a log parse operation binding the contract event 0x7bcaec5d68e62be9724911d03486d3007cea9f870b477fcfe7d834234730cda0.
//
// Solidity: event GasDeducted(uint256 indexed enterpriseId, uint256 amount)
func (_EnterprisePortal *EnterprisePortalFilterer) ParseGasDeducted(log types.Log) (*EnterprisePortalGasDeducted, error) {
	event := new(EnterprisePortalGasDeducted)
	if err := _EnterprisePortal.contract.UnpackLog(event, "GasDeducted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterprisePortalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EnterprisePortal contract.
type EnterprisePortalInitializedIterator struct {
	Event *EnterprisePortalInitialized // Event containing the contract specifics and raw log

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
func (it *EnterprisePortalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterprisePortalInitialized)
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
		it.Event = new(EnterprisePortalInitialized)
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
func (it *EnterprisePortalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterprisePortalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterprisePortalInitialized represents a Initialized event raised by the EnterprisePortal contract.
type EnterprisePortalInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EnterprisePortal *EnterprisePortalFilterer) FilterInitialized(opts *bind.FilterOpts) (*EnterprisePortalInitializedIterator, error) {

	logs, sub, err := _EnterprisePortal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EnterprisePortalInitializedIterator{contract: _EnterprisePortal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EnterprisePortal *EnterprisePortalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EnterprisePortalInitialized) (event.Subscription, error) {

	logs, sub, err := _EnterprisePortal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterprisePortalInitialized)
				if err := _EnterprisePortal.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EnterprisePortal *EnterprisePortalFilterer) ParseInitialized(log types.Log) (*EnterprisePortalInitialized, error) {
	event := new(EnterprisePortalInitialized)
	if err := _EnterprisePortal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterprisePortalOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EnterprisePortal contract.
type EnterprisePortalOwnershipTransferredIterator struct {
	Event *EnterprisePortalOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EnterprisePortalOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterprisePortalOwnershipTransferred)
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
		it.Event = new(EnterprisePortalOwnershipTransferred)
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
func (it *EnterprisePortalOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterprisePortalOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterprisePortalOwnershipTransferred represents a OwnershipTransferred event raised by the EnterprisePortal contract.
type EnterprisePortalOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EnterprisePortal *EnterprisePortalFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EnterprisePortalOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EnterprisePortal.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EnterprisePortalOwnershipTransferredIterator{contract: _EnterprisePortal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EnterprisePortal *EnterprisePortalFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EnterprisePortalOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EnterprisePortal.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterprisePortalOwnershipTransferred)
				if err := _EnterprisePortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EnterprisePortal *EnterprisePortalFilterer) ParseOwnershipTransferred(log types.Log) (*EnterprisePortalOwnershipTransferred, error) {
	event := new(EnterprisePortalOwnershipTransferred)
	if err := _EnterprisePortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterprisePortalUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EnterprisePortal contract.
type EnterprisePortalUpgradedIterator struct {
	Event *EnterprisePortalUpgraded // Event containing the contract specifics and raw log

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
func (it *EnterprisePortalUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterprisePortalUpgraded)
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
		it.Event = new(EnterprisePortalUpgraded)
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
func (it *EnterprisePortalUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterprisePortalUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterprisePortalUpgraded represents a Upgraded event raised by the EnterprisePortal contract.
type EnterprisePortalUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EnterprisePortal *EnterprisePortalFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EnterprisePortalUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnterprisePortal.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EnterprisePortalUpgradedIterator{contract: _EnterprisePortal.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EnterprisePortal *EnterprisePortalFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EnterprisePortalUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnterprisePortal.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterprisePortalUpgraded)
				if err := _EnterprisePortal.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_EnterprisePortal *EnterprisePortalFilterer) ParseUpgraded(log types.Log) (*EnterprisePortalUpgraded, error) {
	event := new(EnterprisePortalUpgraded)
	if err := _EnterprisePortal.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterprisePortalWhitelistedContractAddedIterator is returned from FilterWhitelistedContractAdded and is used to iterate over the raw logs and unpacked data for WhitelistedContractAdded events raised by the EnterprisePortal contract.
type EnterprisePortalWhitelistedContractAddedIterator struct {
	Event *EnterprisePortalWhitelistedContractAdded // Event containing the contract specifics and raw log

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
func (it *EnterprisePortalWhitelistedContractAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterprisePortalWhitelistedContractAdded)
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
		it.Event = new(EnterprisePortalWhitelistedContractAdded)
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
func (it *EnterprisePortalWhitelistedContractAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterprisePortalWhitelistedContractAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterprisePortalWhitelistedContractAdded represents a WhitelistedContractAdded event raised by the EnterprisePortal contract.
type EnterprisePortalWhitelistedContractAdded struct {
	EnterpriseId        *big.Int
	WhitelistedContract common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedContractAdded is a free log retrieval operation binding the contract event 0x1e617678f5241bb68eb4c54388bd575389cd61ba253c6ef9403c933ae87551bc.
//
// Solidity: event WhitelistedContractAdded(uint256 indexed enterpriseId, address indexed whitelistedContract)
func (_EnterprisePortal *EnterprisePortalFilterer) FilterWhitelistedContractAdded(opts *bind.FilterOpts, enterpriseId []*big.Int, whitelistedContract []common.Address) (*EnterprisePortalWhitelistedContractAddedIterator, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}
	var whitelistedContractRule []interface{}
	for _, whitelistedContractItem := range whitelistedContract {
		whitelistedContractRule = append(whitelistedContractRule, whitelistedContractItem)
	}

	logs, sub, err := _EnterprisePortal.contract.FilterLogs(opts, "WhitelistedContractAdded", enterpriseIdRule, whitelistedContractRule)
	if err != nil {
		return nil, err
	}
	return &EnterprisePortalWhitelistedContractAddedIterator{contract: _EnterprisePortal.contract, event: "WhitelistedContractAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedContractAdded is a free log subscription operation binding the contract event 0x1e617678f5241bb68eb4c54388bd575389cd61ba253c6ef9403c933ae87551bc.
//
// Solidity: event WhitelistedContractAdded(uint256 indexed enterpriseId, address indexed whitelistedContract)
func (_EnterprisePortal *EnterprisePortalFilterer) WatchWhitelistedContractAdded(opts *bind.WatchOpts, sink chan<- *EnterprisePortalWhitelistedContractAdded, enterpriseId []*big.Int, whitelistedContract []common.Address) (event.Subscription, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}
	var whitelistedContractRule []interface{}
	for _, whitelistedContractItem := range whitelistedContract {
		whitelistedContractRule = append(whitelistedContractRule, whitelistedContractItem)
	}

	logs, sub, err := _EnterprisePortal.contract.WatchLogs(opts, "WhitelistedContractAdded", enterpriseIdRule, whitelistedContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterprisePortalWhitelistedContractAdded)
				if err := _EnterprisePortal.contract.UnpackLog(event, "WhitelistedContractAdded", log); err != nil {
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

// ParseWhitelistedContractAdded is a log parse operation binding the contract event 0x1e617678f5241bb68eb4c54388bd575389cd61ba253c6ef9403c933ae87551bc.
//
// Solidity: event WhitelistedContractAdded(uint256 indexed enterpriseId, address indexed whitelistedContract)
func (_EnterprisePortal *EnterprisePortalFilterer) ParseWhitelistedContractAdded(log types.Log) (*EnterprisePortalWhitelistedContractAdded, error) {
	event := new(EnterprisePortalWhitelistedContractAdded)
	if err := _EnterprisePortal.contract.UnpackLog(event, "WhitelistedContractAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterprisePortalWhitelistedContractRemovedIterator is returned from FilterWhitelistedContractRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedContractRemoved events raised by the EnterprisePortal contract.
type EnterprisePortalWhitelistedContractRemovedIterator struct {
	Event *EnterprisePortalWhitelistedContractRemoved // Event containing the contract specifics and raw log

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
func (it *EnterprisePortalWhitelistedContractRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterprisePortalWhitelistedContractRemoved)
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
		it.Event = new(EnterprisePortalWhitelistedContractRemoved)
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
func (it *EnterprisePortalWhitelistedContractRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterprisePortalWhitelistedContractRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterprisePortalWhitelistedContractRemoved represents a WhitelistedContractRemoved event raised by the EnterprisePortal contract.
type EnterprisePortalWhitelistedContractRemoved struct {
	EnterpriseId        *big.Int
	WhitelistedContract common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedContractRemoved is a free log retrieval operation binding the contract event 0xf411dfd94bda56cc600b43a44f8fd60805d7fe1d5bf0bb11e4a6735cf031f333.
//
// Solidity: event WhitelistedContractRemoved(uint256 indexed enterpriseId, address indexed whitelistedContract)
func (_EnterprisePortal *EnterprisePortalFilterer) FilterWhitelistedContractRemoved(opts *bind.FilterOpts, enterpriseId []*big.Int, whitelistedContract []common.Address) (*EnterprisePortalWhitelistedContractRemovedIterator, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}
	var whitelistedContractRule []interface{}
	for _, whitelistedContractItem := range whitelistedContract {
		whitelistedContractRule = append(whitelistedContractRule, whitelistedContractItem)
	}

	logs, sub, err := _EnterprisePortal.contract.FilterLogs(opts, "WhitelistedContractRemoved", enterpriseIdRule, whitelistedContractRule)
	if err != nil {
		return nil, err
	}
	return &EnterprisePortalWhitelistedContractRemovedIterator{contract: _EnterprisePortal.contract, event: "WhitelistedContractRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedContractRemoved is a free log subscription operation binding the contract event 0xf411dfd94bda56cc600b43a44f8fd60805d7fe1d5bf0bb11e4a6735cf031f333.
//
// Solidity: event WhitelistedContractRemoved(uint256 indexed enterpriseId, address indexed whitelistedContract)
func (_EnterprisePortal *EnterprisePortalFilterer) WatchWhitelistedContractRemoved(opts *bind.WatchOpts, sink chan<- *EnterprisePortalWhitelistedContractRemoved, enterpriseId []*big.Int, whitelistedContract []common.Address) (event.Subscription, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}
	var whitelistedContractRule []interface{}
	for _, whitelistedContractItem := range whitelistedContract {
		whitelistedContractRule = append(whitelistedContractRule, whitelistedContractItem)
	}

	logs, sub, err := _EnterprisePortal.contract.WatchLogs(opts, "WhitelistedContractRemoved", enterpriseIdRule, whitelistedContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterprisePortalWhitelistedContractRemoved)
				if err := _EnterprisePortal.contract.UnpackLog(event, "WhitelistedContractRemoved", log); err != nil {
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

// ParseWhitelistedContractRemoved is a log parse operation binding the contract event 0xf411dfd94bda56cc600b43a44f8fd60805d7fe1d5bf0bb11e4a6735cf031f333.
//
// Solidity: event WhitelistedContractRemoved(uint256 indexed enterpriseId, address indexed whitelistedContract)
func (_EnterprisePortal *EnterprisePortalFilterer) ParseWhitelistedContractRemoved(log types.Log) (*EnterprisePortalWhitelistedContractRemoved, error) {
	event := new(EnterprisePortalWhitelistedContractRemoved)
	if err := _EnterprisePortal.contract.UnpackLog(event, "WhitelistedContractRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
