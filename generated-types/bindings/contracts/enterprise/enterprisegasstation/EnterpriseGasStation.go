// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package enterprisegasstation

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

// EnterpriseGasStationMetaData contains all meta data concerning the EnterpriseGasStation contract.
var EnterpriseGasStationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costInTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"GasPlanAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"GasPlanStatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"costInTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"addGasPlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enterprisePortalContract\",\"outputs\":[{\"internalType\":\"contractEnterprisePortal\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gasPlans\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costInTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_enterprisePortalContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_llToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reserveWallet\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"llToken\",\"outputs\":[{\"internalType\":\"contractLLERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextPlanId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"enterpriseId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"}],\"name\":\"purchaseGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"setGasPlanStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// EnterpriseGasStationABI is the input ABI used to generate the binding from.
// Deprecated: Use EnterpriseGasStationMetaData.ABI instead.
var EnterpriseGasStationABI = EnterpriseGasStationMetaData.ABI

// EnterpriseGasStation is an auto generated Go binding around an Ethereum contract.
type EnterpriseGasStation struct {
	EnterpriseGasStationCaller     // Read-only binding to the contract
	EnterpriseGasStationTransactor // Write-only binding to the contract
	EnterpriseGasStationFilterer   // Log filterer for contract events
}

// EnterpriseGasStationCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnterpriseGasStationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseGasStationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnterpriseGasStationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseGasStationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnterpriseGasStationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseGasStationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnterpriseGasStationSession struct {
	Contract     *EnterpriseGasStation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EnterpriseGasStationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnterpriseGasStationCallerSession struct {
	Contract *EnterpriseGasStationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// EnterpriseGasStationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnterpriseGasStationTransactorSession struct {
	Contract     *EnterpriseGasStationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// EnterpriseGasStationRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnterpriseGasStationRaw struct {
	Contract *EnterpriseGasStation // Generic contract binding to access the raw methods on
}

// EnterpriseGasStationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnterpriseGasStationCallerRaw struct {
	Contract *EnterpriseGasStationCaller // Generic read-only contract binding to access the raw methods on
}

// EnterpriseGasStationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnterpriseGasStationTransactorRaw struct {
	Contract *EnterpriseGasStationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnterpriseGasStation creates a new instance of EnterpriseGasStation, bound to a specific deployed contract.
func NewEnterpriseGasStation(address common.Address, backend bind.ContractBackend) (*EnterpriseGasStation, error) {
	contract, err := bindEnterpriseGasStation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnterpriseGasStation{EnterpriseGasStationCaller: EnterpriseGasStationCaller{contract: contract}, EnterpriseGasStationTransactor: EnterpriseGasStationTransactor{contract: contract}, EnterpriseGasStationFilterer: EnterpriseGasStationFilterer{contract: contract}}, nil
}

// NewEnterpriseGasStationCaller creates a new read-only instance of EnterpriseGasStation, bound to a specific deployed contract.
func NewEnterpriseGasStationCaller(address common.Address, caller bind.ContractCaller) (*EnterpriseGasStationCaller, error) {
	contract, err := bindEnterpriseGasStation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseGasStationCaller{contract: contract}, nil
}

// NewEnterpriseGasStationTransactor creates a new write-only instance of EnterpriseGasStation, bound to a specific deployed contract.
func NewEnterpriseGasStationTransactor(address common.Address, transactor bind.ContractTransactor) (*EnterpriseGasStationTransactor, error) {
	contract, err := bindEnterpriseGasStation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseGasStationTransactor{contract: contract}, nil
}

// NewEnterpriseGasStationFilterer creates a new log filterer instance of EnterpriseGasStation, bound to a specific deployed contract.
func NewEnterpriseGasStationFilterer(address common.Address, filterer bind.ContractFilterer) (*EnterpriseGasStationFilterer, error) {
	contract, err := bindEnterpriseGasStation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnterpriseGasStationFilterer{contract: contract}, nil
}

// bindEnterpriseGasStation binds a generic wrapper to an already deployed contract.
func bindEnterpriseGasStation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnterpriseGasStationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterpriseGasStation *EnterpriseGasStationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterpriseGasStation.Contract.EnterpriseGasStationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterpriseGasStation *EnterpriseGasStationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.EnterpriseGasStationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterpriseGasStation *EnterpriseGasStationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.EnterpriseGasStationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterpriseGasStation *EnterpriseGasStationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterpriseGasStation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterpriseGasStation *EnterpriseGasStationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterpriseGasStation *EnterpriseGasStationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EnterpriseGasStation *EnterpriseGasStationCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EnterpriseGasStation.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EnterpriseGasStation *EnterpriseGasStationSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _EnterpriseGasStation.Contract.UPGRADEINTERFACEVERSION(&_EnterpriseGasStation.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EnterpriseGasStation *EnterpriseGasStationCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _EnterpriseGasStation.Contract.UPGRADEINTERFACEVERSION(&_EnterpriseGasStation.CallOpts)
}

// EnterprisePortalContract is a free data retrieval call binding the contract method 0xa3a5c00a.
//
// Solidity: function enterprisePortalContract() view returns(address)
func (_EnterpriseGasStation *EnterpriseGasStationCaller) EnterprisePortalContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseGasStation.contract.Call(opts, &out, "enterprisePortalContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnterprisePortalContract is a free data retrieval call binding the contract method 0xa3a5c00a.
//
// Solidity: function enterprisePortalContract() view returns(address)
func (_EnterpriseGasStation *EnterpriseGasStationSession) EnterprisePortalContract() (common.Address, error) {
	return _EnterpriseGasStation.Contract.EnterprisePortalContract(&_EnterpriseGasStation.CallOpts)
}

// EnterprisePortalContract is a free data retrieval call binding the contract method 0xa3a5c00a.
//
// Solidity: function enterprisePortalContract() view returns(address)
func (_EnterpriseGasStation *EnterpriseGasStationCallerSession) EnterprisePortalContract() (common.Address, error) {
	return _EnterpriseGasStation.Contract.EnterprisePortalContract(&_EnterpriseGasStation.CallOpts)
}

// GasPlans is a free data retrieval call binding the contract method 0x2d8fff2f.
//
// Solidity: function gasPlans(uint256 ) view returns(uint256 id, uint256 costInTokens, uint256 gasAmount, bool isActive)
func (_EnterpriseGasStation *EnterpriseGasStationCaller) GasPlans(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id           *big.Int
	CostInTokens *big.Int
	GasAmount    *big.Int
	IsActive     bool
}, error) {
	var out []interface{}
	err := _EnterpriseGasStation.contract.Call(opts, &out, "gasPlans", arg0)

	outstruct := new(struct {
		Id           *big.Int
		CostInTokens *big.Int
		GasAmount    *big.Int
		IsActive     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CostInTokens = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.GasAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// GasPlans is a free data retrieval call binding the contract method 0x2d8fff2f.
//
// Solidity: function gasPlans(uint256 ) view returns(uint256 id, uint256 costInTokens, uint256 gasAmount, bool isActive)
func (_EnterpriseGasStation *EnterpriseGasStationSession) GasPlans(arg0 *big.Int) (struct {
	Id           *big.Int
	CostInTokens *big.Int
	GasAmount    *big.Int
	IsActive     bool
}, error) {
	return _EnterpriseGasStation.Contract.GasPlans(&_EnterpriseGasStation.CallOpts, arg0)
}

// GasPlans is a free data retrieval call binding the contract method 0x2d8fff2f.
//
// Solidity: function gasPlans(uint256 ) view returns(uint256 id, uint256 costInTokens, uint256 gasAmount, bool isActive)
func (_EnterpriseGasStation *EnterpriseGasStationCallerSession) GasPlans(arg0 *big.Int) (struct {
	Id           *big.Int
	CostInTokens *big.Int
	GasAmount    *big.Int
	IsActive     bool
}, error) {
	return _EnterpriseGasStation.Contract.GasPlans(&_EnterpriseGasStation.CallOpts, arg0)
}

// LlToken is a free data retrieval call binding the contract method 0xbe4d9f66.
//
// Solidity: function llToken() view returns(address)
func (_EnterpriseGasStation *EnterpriseGasStationCaller) LlToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseGasStation.contract.Call(opts, &out, "llToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LlToken is a free data retrieval call binding the contract method 0xbe4d9f66.
//
// Solidity: function llToken() view returns(address)
func (_EnterpriseGasStation *EnterpriseGasStationSession) LlToken() (common.Address, error) {
	return _EnterpriseGasStation.Contract.LlToken(&_EnterpriseGasStation.CallOpts)
}

// LlToken is a free data retrieval call binding the contract method 0xbe4d9f66.
//
// Solidity: function llToken() view returns(address)
func (_EnterpriseGasStation *EnterpriseGasStationCallerSession) LlToken() (common.Address, error) {
	return _EnterpriseGasStation.Contract.LlToken(&_EnterpriseGasStation.CallOpts)
}

// NextPlanId is a free data retrieval call binding the contract method 0x5f8d26b2.
//
// Solidity: function nextPlanId() view returns(uint256)
func (_EnterpriseGasStation *EnterpriseGasStationCaller) NextPlanId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnterpriseGasStation.contract.Call(opts, &out, "nextPlanId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextPlanId is a free data retrieval call binding the contract method 0x5f8d26b2.
//
// Solidity: function nextPlanId() view returns(uint256)
func (_EnterpriseGasStation *EnterpriseGasStationSession) NextPlanId() (*big.Int, error) {
	return _EnterpriseGasStation.Contract.NextPlanId(&_EnterpriseGasStation.CallOpts)
}

// NextPlanId is a free data retrieval call binding the contract method 0x5f8d26b2.
//
// Solidity: function nextPlanId() view returns(uint256)
func (_EnterpriseGasStation *EnterpriseGasStationCallerSession) NextPlanId() (*big.Int, error) {
	return _EnterpriseGasStation.Contract.NextPlanId(&_EnterpriseGasStation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnterpriseGasStation *EnterpriseGasStationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseGasStation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnterpriseGasStation *EnterpriseGasStationSession) Owner() (common.Address, error) {
	return _EnterpriseGasStation.Contract.Owner(&_EnterpriseGasStation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnterpriseGasStation *EnterpriseGasStationCallerSession) Owner() (common.Address, error) {
	return _EnterpriseGasStation.Contract.Owner(&_EnterpriseGasStation.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EnterpriseGasStation *EnterpriseGasStationCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnterpriseGasStation.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EnterpriseGasStation *EnterpriseGasStationSession) ProxiableUUID() ([32]byte, error) {
	return _EnterpriseGasStation.Contract.ProxiableUUID(&_EnterpriseGasStation.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EnterpriseGasStation *EnterpriseGasStationCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EnterpriseGasStation.Contract.ProxiableUUID(&_EnterpriseGasStation.CallOpts)
}

// ReserveWallet is a free data retrieval call binding the contract method 0xd72b11bd.
//
// Solidity: function reserveWallet() view returns(address)
func (_EnterpriseGasStation *EnterpriseGasStationCaller) ReserveWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseGasStation.contract.Call(opts, &out, "reserveWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReserveWallet is a free data retrieval call binding the contract method 0xd72b11bd.
//
// Solidity: function reserveWallet() view returns(address)
func (_EnterpriseGasStation *EnterpriseGasStationSession) ReserveWallet() (common.Address, error) {
	return _EnterpriseGasStation.Contract.ReserveWallet(&_EnterpriseGasStation.CallOpts)
}

// ReserveWallet is a free data retrieval call binding the contract method 0xd72b11bd.
//
// Solidity: function reserveWallet() view returns(address)
func (_EnterpriseGasStation *EnterpriseGasStationCallerSession) ReserveWallet() (common.Address, error) {
	return _EnterpriseGasStation.Contract.ReserveWallet(&_EnterpriseGasStation.CallOpts)
}

// AddGasPlan is a paid mutator transaction binding the contract method 0x6240d47f.
//
// Solidity: function addGasPlan(uint256 costInTokens, uint256 gasAmount, bool isActive) returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactor) AddGasPlan(opts *bind.TransactOpts, costInTokens *big.Int, gasAmount *big.Int, isActive bool) (*types.Transaction, error) {
	return _EnterpriseGasStation.contract.Transact(opts, "addGasPlan", costInTokens, gasAmount, isActive)
}

// AddGasPlan is a paid mutator transaction binding the contract method 0x6240d47f.
//
// Solidity: function addGasPlan(uint256 costInTokens, uint256 gasAmount, bool isActive) returns()
func (_EnterpriseGasStation *EnterpriseGasStationSession) AddGasPlan(costInTokens *big.Int, gasAmount *big.Int, isActive bool) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.AddGasPlan(&_EnterpriseGasStation.TransactOpts, costInTokens, gasAmount, isActive)
}

// AddGasPlan is a paid mutator transaction binding the contract method 0x6240d47f.
//
// Solidity: function addGasPlan(uint256 costInTokens, uint256 gasAmount, bool isActive) returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactorSession) AddGasPlan(costInTokens *big.Int, gasAmount *big.Int, isActive bool) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.AddGasPlan(&_EnterpriseGasStation.TransactOpts, costInTokens, gasAmount, isActive)
}

// ClaimGas is a paid mutator transaction binding the contract method 0xc0785ab8.
//
// Solidity: function claimGas(uint256 enterpriseId, uint256 amount) returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactor) ClaimGas(opts *bind.TransactOpts, enterpriseId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _EnterpriseGasStation.contract.Transact(opts, "claimGas", enterpriseId, amount)
}

// ClaimGas is a paid mutator transaction binding the contract method 0xc0785ab8.
//
// Solidity: function claimGas(uint256 enterpriseId, uint256 amount) returns()
func (_EnterpriseGasStation *EnterpriseGasStationSession) ClaimGas(enterpriseId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.ClaimGas(&_EnterpriseGasStation.TransactOpts, enterpriseId, amount)
}

// ClaimGas is a paid mutator transaction binding the contract method 0xc0785ab8.
//
// Solidity: function claimGas(uint256 enterpriseId, uint256 amount) returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactorSession) ClaimGas(enterpriseId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.ClaimGas(&_EnterpriseGasStation.TransactOpts, enterpriseId, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _enterprisePortalContract, address _llToken, address _reserveWallet) returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactor) Initialize(opts *bind.TransactOpts, _enterprisePortalContract common.Address, _llToken common.Address, _reserveWallet common.Address) (*types.Transaction, error) {
	return _EnterpriseGasStation.contract.Transact(opts, "initialize", _enterprisePortalContract, _llToken, _reserveWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _enterprisePortalContract, address _llToken, address _reserveWallet) returns()
func (_EnterpriseGasStation *EnterpriseGasStationSession) Initialize(_enterprisePortalContract common.Address, _llToken common.Address, _reserveWallet common.Address) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.Initialize(&_EnterpriseGasStation.TransactOpts, _enterprisePortalContract, _llToken, _reserveWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _enterprisePortalContract, address _llToken, address _reserveWallet) returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactorSession) Initialize(_enterprisePortalContract common.Address, _llToken common.Address, _reserveWallet common.Address) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.Initialize(&_EnterpriseGasStation.TransactOpts, _enterprisePortalContract, _llToken, _reserveWallet)
}

// PurchaseGas is a paid mutator transaction binding the contract method 0x7dd82068.
//
// Solidity: function purchaseGas(uint256 enterpriseId, uint256 planId) returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactor) PurchaseGas(opts *bind.TransactOpts, enterpriseId *big.Int, planId *big.Int) (*types.Transaction, error) {
	return _EnterpriseGasStation.contract.Transact(opts, "purchaseGas", enterpriseId, planId)
}

// PurchaseGas is a paid mutator transaction binding the contract method 0x7dd82068.
//
// Solidity: function purchaseGas(uint256 enterpriseId, uint256 planId) returns()
func (_EnterpriseGasStation *EnterpriseGasStationSession) PurchaseGas(enterpriseId *big.Int, planId *big.Int) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.PurchaseGas(&_EnterpriseGasStation.TransactOpts, enterpriseId, planId)
}

// PurchaseGas is a paid mutator transaction binding the contract method 0x7dd82068.
//
// Solidity: function purchaseGas(uint256 enterpriseId, uint256 planId) returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactorSession) PurchaseGas(enterpriseId *big.Int, planId *big.Int) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.PurchaseGas(&_EnterpriseGasStation.TransactOpts, enterpriseId, planId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseGasStation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnterpriseGasStation *EnterpriseGasStationSession) RenounceOwnership() (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.RenounceOwnership(&_EnterpriseGasStation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.RenounceOwnership(&_EnterpriseGasStation.TransactOpts)
}

// SetGasPlanStatus is a paid mutator transaction binding the contract method 0x9296749e.
//
// Solidity: function setGasPlanStatus(uint256 planId, bool isActive) returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactor) SetGasPlanStatus(opts *bind.TransactOpts, planId *big.Int, isActive bool) (*types.Transaction, error) {
	return _EnterpriseGasStation.contract.Transact(opts, "setGasPlanStatus", planId, isActive)
}

// SetGasPlanStatus is a paid mutator transaction binding the contract method 0x9296749e.
//
// Solidity: function setGasPlanStatus(uint256 planId, bool isActive) returns()
func (_EnterpriseGasStation *EnterpriseGasStationSession) SetGasPlanStatus(planId *big.Int, isActive bool) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.SetGasPlanStatus(&_EnterpriseGasStation.TransactOpts, planId, isActive)
}

// SetGasPlanStatus is a paid mutator transaction binding the contract method 0x9296749e.
//
// Solidity: function setGasPlanStatus(uint256 planId, bool isActive) returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactorSession) SetGasPlanStatus(planId *big.Int, isActive bool) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.SetGasPlanStatus(&_EnterpriseGasStation.TransactOpts, planId, isActive)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EnterpriseGasStation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnterpriseGasStation *EnterpriseGasStationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.TransferOwnership(&_EnterpriseGasStation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.TransferOwnership(&_EnterpriseGasStation.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EnterpriseGasStation.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EnterpriseGasStation *EnterpriseGasStationSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.UpgradeToAndCall(&_EnterpriseGasStation.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EnterpriseGasStation *EnterpriseGasStationTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EnterpriseGasStation.Contract.UpgradeToAndCall(&_EnterpriseGasStation.TransactOpts, newImplementation, data)
}

// EnterpriseGasStationGasClaimedIterator is returned from FilterGasClaimed and is used to iterate over the raw logs and unpacked data for GasClaimed events raised by the EnterpriseGasStation contract.
type EnterpriseGasStationGasClaimedIterator struct {
	Event *EnterpriseGasStationGasClaimed // Event containing the contract specifics and raw log

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
func (it *EnterpriseGasStationGasClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseGasStationGasClaimed)
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
		it.Event = new(EnterpriseGasStationGasClaimed)
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
func (it *EnterpriseGasStationGasClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseGasStationGasClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseGasStationGasClaimed represents a GasClaimed event raised by the EnterpriseGasStation contract.
type EnterpriseGasStationGasClaimed struct {
	EnterpriseId *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGasClaimed is a free log retrieval operation binding the contract event 0xd52e35f4f279be24b2b7d68e02535f0e1b775658f223700cda44dda590de4196.
//
// Solidity: event GasClaimed(uint256 indexed enterpriseId, uint256 amount)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) FilterGasClaimed(opts *bind.FilterOpts, enterpriseId []*big.Int) (*EnterpriseGasStationGasClaimedIterator, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}

	logs, sub, err := _EnterpriseGasStation.contract.FilterLogs(opts, "GasClaimed", enterpriseIdRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseGasStationGasClaimedIterator{contract: _EnterpriseGasStation.contract, event: "GasClaimed", logs: logs, sub: sub}, nil
}

// WatchGasClaimed is a free log subscription operation binding the contract event 0xd52e35f4f279be24b2b7d68e02535f0e1b775658f223700cda44dda590de4196.
//
// Solidity: event GasClaimed(uint256 indexed enterpriseId, uint256 amount)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) WatchGasClaimed(opts *bind.WatchOpts, sink chan<- *EnterpriseGasStationGasClaimed, enterpriseId []*big.Int) (event.Subscription, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}

	logs, sub, err := _EnterpriseGasStation.contract.WatchLogs(opts, "GasClaimed", enterpriseIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseGasStationGasClaimed)
				if err := _EnterpriseGasStation.contract.UnpackLog(event, "GasClaimed", log); err != nil {
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

// ParseGasClaimed is a log parse operation binding the contract event 0xd52e35f4f279be24b2b7d68e02535f0e1b775658f223700cda44dda590de4196.
//
// Solidity: event GasClaimed(uint256 indexed enterpriseId, uint256 amount)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) ParseGasClaimed(log types.Log) (*EnterpriseGasStationGasClaimed, error) {
	event := new(EnterpriseGasStationGasClaimed)
	if err := _EnterpriseGasStation.contract.UnpackLog(event, "GasClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseGasStationGasPlanAddedIterator is returned from FilterGasPlanAdded and is used to iterate over the raw logs and unpacked data for GasPlanAdded events raised by the EnterpriseGasStation contract.
type EnterpriseGasStationGasPlanAddedIterator struct {
	Event *EnterpriseGasStationGasPlanAdded // Event containing the contract specifics and raw log

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
func (it *EnterpriseGasStationGasPlanAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseGasStationGasPlanAdded)
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
		it.Event = new(EnterpriseGasStationGasPlanAdded)
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
func (it *EnterpriseGasStationGasPlanAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseGasStationGasPlanAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseGasStationGasPlanAdded represents a GasPlanAdded event raised by the EnterpriseGasStation contract.
type EnterpriseGasStationGasPlanAdded struct {
	PlanId       *big.Int
	CostInTokens *big.Int
	GasAmount    *big.Int
	IsActive     bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGasPlanAdded is a free log retrieval operation binding the contract event 0x5e17482cfe25c63b3933f0c830f5e0591f63b1a8308f86962b6854fb2705ac1d.
//
// Solidity: event GasPlanAdded(uint256 indexed planId, uint256 costInTokens, uint256 gasAmount, bool isActive)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) FilterGasPlanAdded(opts *bind.FilterOpts, planId []*big.Int) (*EnterpriseGasStationGasPlanAddedIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _EnterpriseGasStation.contract.FilterLogs(opts, "GasPlanAdded", planIdRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseGasStationGasPlanAddedIterator{contract: _EnterpriseGasStation.contract, event: "GasPlanAdded", logs: logs, sub: sub}, nil
}

// WatchGasPlanAdded is a free log subscription operation binding the contract event 0x5e17482cfe25c63b3933f0c830f5e0591f63b1a8308f86962b6854fb2705ac1d.
//
// Solidity: event GasPlanAdded(uint256 indexed planId, uint256 costInTokens, uint256 gasAmount, bool isActive)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) WatchGasPlanAdded(opts *bind.WatchOpts, sink chan<- *EnterpriseGasStationGasPlanAdded, planId []*big.Int) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _EnterpriseGasStation.contract.WatchLogs(opts, "GasPlanAdded", planIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseGasStationGasPlanAdded)
				if err := _EnterpriseGasStation.contract.UnpackLog(event, "GasPlanAdded", log); err != nil {
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

// ParseGasPlanAdded is a log parse operation binding the contract event 0x5e17482cfe25c63b3933f0c830f5e0591f63b1a8308f86962b6854fb2705ac1d.
//
// Solidity: event GasPlanAdded(uint256 indexed planId, uint256 costInTokens, uint256 gasAmount, bool isActive)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) ParseGasPlanAdded(log types.Log) (*EnterpriseGasStationGasPlanAdded, error) {
	event := new(EnterpriseGasStationGasPlanAdded)
	if err := _EnterpriseGasStation.contract.UnpackLog(event, "GasPlanAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseGasStationGasPlanStatusUpdatedIterator is returned from FilterGasPlanStatusUpdated and is used to iterate over the raw logs and unpacked data for GasPlanStatusUpdated events raised by the EnterpriseGasStation contract.
type EnterpriseGasStationGasPlanStatusUpdatedIterator struct {
	Event *EnterpriseGasStationGasPlanStatusUpdated // Event containing the contract specifics and raw log

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
func (it *EnterpriseGasStationGasPlanStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseGasStationGasPlanStatusUpdated)
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
		it.Event = new(EnterpriseGasStationGasPlanStatusUpdated)
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
func (it *EnterpriseGasStationGasPlanStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseGasStationGasPlanStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseGasStationGasPlanStatusUpdated represents a GasPlanStatusUpdated event raised by the EnterpriseGasStation contract.
type EnterpriseGasStationGasPlanStatusUpdated struct {
	PlanId   *big.Int
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGasPlanStatusUpdated is a free log retrieval operation binding the contract event 0x8996afa63e5458ff6d2aaab9a21ad269cd87e9a2c37e34ccd3b5610039864448.
//
// Solidity: event GasPlanStatusUpdated(uint256 indexed planId, bool isActive)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) FilterGasPlanStatusUpdated(opts *bind.FilterOpts, planId []*big.Int) (*EnterpriseGasStationGasPlanStatusUpdatedIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _EnterpriseGasStation.contract.FilterLogs(opts, "GasPlanStatusUpdated", planIdRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseGasStationGasPlanStatusUpdatedIterator{contract: _EnterpriseGasStation.contract, event: "GasPlanStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchGasPlanStatusUpdated is a free log subscription operation binding the contract event 0x8996afa63e5458ff6d2aaab9a21ad269cd87e9a2c37e34ccd3b5610039864448.
//
// Solidity: event GasPlanStatusUpdated(uint256 indexed planId, bool isActive)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) WatchGasPlanStatusUpdated(opts *bind.WatchOpts, sink chan<- *EnterpriseGasStationGasPlanStatusUpdated, planId []*big.Int) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _EnterpriseGasStation.contract.WatchLogs(opts, "GasPlanStatusUpdated", planIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseGasStationGasPlanStatusUpdated)
				if err := _EnterpriseGasStation.contract.UnpackLog(event, "GasPlanStatusUpdated", log); err != nil {
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

// ParseGasPlanStatusUpdated is a log parse operation binding the contract event 0x8996afa63e5458ff6d2aaab9a21ad269cd87e9a2c37e34ccd3b5610039864448.
//
// Solidity: event GasPlanStatusUpdated(uint256 indexed planId, bool isActive)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) ParseGasPlanStatusUpdated(log types.Log) (*EnterpriseGasStationGasPlanStatusUpdated, error) {
	event := new(EnterpriseGasStationGasPlanStatusUpdated)
	if err := _EnterpriseGasStation.contract.UnpackLog(event, "GasPlanStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseGasStationGasPurchasedIterator is returned from FilterGasPurchased and is used to iterate over the raw logs and unpacked data for GasPurchased events raised by the EnterpriseGasStation contract.
type EnterpriseGasStationGasPurchasedIterator struct {
	Event *EnterpriseGasStationGasPurchased // Event containing the contract specifics and raw log

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
func (it *EnterpriseGasStationGasPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseGasStationGasPurchased)
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
		it.Event = new(EnterpriseGasStationGasPurchased)
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
func (it *EnterpriseGasStationGasPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseGasStationGasPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseGasStationGasPurchased represents a GasPurchased event raised by the EnterpriseGasStation contract.
type EnterpriseGasStationGasPurchased struct {
	EnterpriseId *big.Int
	PlanId       *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGasPurchased is a free log retrieval operation binding the contract event 0x51643e89a7b412e7c4ff69e8e8246bcb91f01f8446bdaee54d0baff46a301ce4.
//
// Solidity: event GasPurchased(uint256 indexed enterpriseId, uint256 indexed planId, uint256 amount)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) FilterGasPurchased(opts *bind.FilterOpts, enterpriseId []*big.Int, planId []*big.Int) (*EnterpriseGasStationGasPurchasedIterator, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}
	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _EnterpriseGasStation.contract.FilterLogs(opts, "GasPurchased", enterpriseIdRule, planIdRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseGasStationGasPurchasedIterator{contract: _EnterpriseGasStation.contract, event: "GasPurchased", logs: logs, sub: sub}, nil
}

// WatchGasPurchased is a free log subscription operation binding the contract event 0x51643e89a7b412e7c4ff69e8e8246bcb91f01f8446bdaee54d0baff46a301ce4.
//
// Solidity: event GasPurchased(uint256 indexed enterpriseId, uint256 indexed planId, uint256 amount)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) WatchGasPurchased(opts *bind.WatchOpts, sink chan<- *EnterpriseGasStationGasPurchased, enterpriseId []*big.Int, planId []*big.Int) (event.Subscription, error) {

	var enterpriseIdRule []interface{}
	for _, enterpriseIdItem := range enterpriseId {
		enterpriseIdRule = append(enterpriseIdRule, enterpriseIdItem)
	}
	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _EnterpriseGasStation.contract.WatchLogs(opts, "GasPurchased", enterpriseIdRule, planIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseGasStationGasPurchased)
				if err := _EnterpriseGasStation.contract.UnpackLog(event, "GasPurchased", log); err != nil {
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

// ParseGasPurchased is a log parse operation binding the contract event 0x51643e89a7b412e7c4ff69e8e8246bcb91f01f8446bdaee54d0baff46a301ce4.
//
// Solidity: event GasPurchased(uint256 indexed enterpriseId, uint256 indexed planId, uint256 amount)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) ParseGasPurchased(log types.Log) (*EnterpriseGasStationGasPurchased, error) {
	event := new(EnterpriseGasStationGasPurchased)
	if err := _EnterpriseGasStation.contract.UnpackLog(event, "GasPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseGasStationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EnterpriseGasStation contract.
type EnterpriseGasStationInitializedIterator struct {
	Event *EnterpriseGasStationInitialized // Event containing the contract specifics and raw log

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
func (it *EnterpriseGasStationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseGasStationInitialized)
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
		it.Event = new(EnterpriseGasStationInitialized)
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
func (it *EnterpriseGasStationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseGasStationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseGasStationInitialized represents a Initialized event raised by the EnterpriseGasStation contract.
type EnterpriseGasStationInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) FilterInitialized(opts *bind.FilterOpts) (*EnterpriseGasStationInitializedIterator, error) {

	logs, sub, err := _EnterpriseGasStation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EnterpriseGasStationInitializedIterator{contract: _EnterpriseGasStation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EnterpriseGasStationInitialized) (event.Subscription, error) {

	logs, sub, err := _EnterpriseGasStation.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseGasStationInitialized)
				if err := _EnterpriseGasStation.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) ParseInitialized(log types.Log) (*EnterpriseGasStationInitialized, error) {
	event := new(EnterpriseGasStationInitialized)
	if err := _EnterpriseGasStation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseGasStationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EnterpriseGasStation contract.
type EnterpriseGasStationOwnershipTransferredIterator struct {
	Event *EnterpriseGasStationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EnterpriseGasStationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseGasStationOwnershipTransferred)
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
		it.Event = new(EnterpriseGasStationOwnershipTransferred)
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
func (it *EnterpriseGasStationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseGasStationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseGasStationOwnershipTransferred represents a OwnershipTransferred event raised by the EnterpriseGasStation contract.
type EnterpriseGasStationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EnterpriseGasStationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EnterpriseGasStation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseGasStationOwnershipTransferredIterator{contract: _EnterpriseGasStation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EnterpriseGasStationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EnterpriseGasStation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseGasStationOwnershipTransferred)
				if err := _EnterpriseGasStation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) ParseOwnershipTransferred(log types.Log) (*EnterpriseGasStationOwnershipTransferred, error) {
	event := new(EnterpriseGasStationOwnershipTransferred)
	if err := _EnterpriseGasStation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseGasStationUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EnterpriseGasStation contract.
type EnterpriseGasStationUpgradedIterator struct {
	Event *EnterpriseGasStationUpgraded // Event containing the contract specifics and raw log

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
func (it *EnterpriseGasStationUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseGasStationUpgraded)
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
		it.Event = new(EnterpriseGasStationUpgraded)
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
func (it *EnterpriseGasStationUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseGasStationUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseGasStationUpgraded represents a Upgraded event raised by the EnterpriseGasStation contract.
type EnterpriseGasStationUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EnterpriseGasStationUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnterpriseGasStation.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseGasStationUpgradedIterator{contract: _EnterpriseGasStation.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EnterpriseGasStationUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnterpriseGasStation.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseGasStationUpgraded)
				if err := _EnterpriseGasStation.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_EnterpriseGasStation *EnterpriseGasStationFilterer) ParseUpgraded(log types.Log) (*EnterpriseGasStationUpgraded, error) {
	event := new(EnterpriseGasStationUpgraded)
	if err := _EnterpriseGasStation.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
