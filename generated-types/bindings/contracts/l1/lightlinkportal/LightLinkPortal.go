// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lightlinkportal

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

// LightLinkPortalMetaData contains all meta data concerning the LightLinkPortal contract.
var LightLinkPortalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContentLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasEstimation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataRemainder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidHeader\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LargeCalldata\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonReentrant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCustomGasToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SmallGasLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedString\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"opaqueData\",\"type\":\"bytes\"}],\"name\":\"TransactionDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"WithdrawalProven\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_isCreation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20Transaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_isCreation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donateETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.WithdrawalTransaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"finalizeWithdrawalTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"finalizedWithdrawals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICanonicalStateChain\",\"name\":\"_l2Oracle\",\"type\":\"address\"},{\"internalType\":\"contractIChallengeBase\",\"name\":\"_challenge\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"isOutputFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Oracle\",\"outputs\":[{\"internalType\":\"contractICanonicalStateChain\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_byteCount\",\"type\":\"uint64\"}],\"name\":\"minimumGasLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"prevBaseFee\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"prevBoughtGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"prevBlockNum\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.WithdrawalTransaction\",\"name\":\"_tx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"latestBlockhash\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.OutputRootProof\",\"name\":\"_outputRootProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_withdrawalProof\",\"type\":\"bytes[]\"}],\"name\":\"proveWithdrawalTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"provenWithdrawals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2OutputIndex\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resourceConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maxResourceLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"elasticityMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"baseFeeMaxChangeDenominator\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"minimumBaseFee\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"systemTxMaxGas\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"maximumBaseFee\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_symbol\",\"type\":\"bytes32\"}],\"name\":\"setGasPayingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LightLinkPortalABI is the input ABI used to generate the binding from.
// Deprecated: Use LightLinkPortalMetaData.ABI instead.
var LightLinkPortalABI = LightLinkPortalMetaData.ABI

// LightLinkPortal is an auto generated Go binding around an Ethereum contract.
type LightLinkPortal struct {
	LightLinkPortalCaller     // Read-only binding to the contract
	LightLinkPortalTransactor // Write-only binding to the contract
	LightLinkPortalFilterer   // Log filterer for contract events
}

// LightLinkPortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type LightLinkPortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightLinkPortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LightLinkPortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightLinkPortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LightLinkPortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightLinkPortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LightLinkPortalSession struct {
	Contract     *LightLinkPortal  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LightLinkPortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LightLinkPortalCallerSession struct {
	Contract *LightLinkPortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LightLinkPortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LightLinkPortalTransactorSession struct {
	Contract     *LightLinkPortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LightLinkPortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type LightLinkPortalRaw struct {
	Contract *LightLinkPortal // Generic contract binding to access the raw methods on
}

// LightLinkPortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LightLinkPortalCallerRaw struct {
	Contract *LightLinkPortalCaller // Generic read-only contract binding to access the raw methods on
}

// LightLinkPortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LightLinkPortalTransactorRaw struct {
	Contract *LightLinkPortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLightLinkPortal creates a new instance of LightLinkPortal, bound to a specific deployed contract.
func NewLightLinkPortal(address common.Address, backend bind.ContractBackend) (*LightLinkPortal, error) {
	contract, err := bindLightLinkPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LightLinkPortal{LightLinkPortalCaller: LightLinkPortalCaller{contract: contract}, LightLinkPortalTransactor: LightLinkPortalTransactor{contract: contract}, LightLinkPortalFilterer: LightLinkPortalFilterer{contract: contract}}, nil
}

// NewLightLinkPortalCaller creates a new read-only instance of LightLinkPortal, bound to a specific deployed contract.
func NewLightLinkPortalCaller(address common.Address, caller bind.ContractCaller) (*LightLinkPortalCaller, error) {
	contract, err := bindLightLinkPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LightLinkPortalCaller{contract: contract}, nil
}

// NewLightLinkPortalTransactor creates a new write-only instance of LightLinkPortal, bound to a specific deployed contract.
func NewLightLinkPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*LightLinkPortalTransactor, error) {
	contract, err := bindLightLinkPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LightLinkPortalTransactor{contract: contract}, nil
}

// NewLightLinkPortalFilterer creates a new log filterer instance of LightLinkPortal, bound to a specific deployed contract.
func NewLightLinkPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*LightLinkPortalFilterer, error) {
	contract, err := bindLightLinkPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LightLinkPortalFilterer{contract: contract}, nil
}

// bindLightLinkPortal binds a generic wrapper to an already deployed contract.
func bindLightLinkPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LightLinkPortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightLinkPortal *LightLinkPortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightLinkPortal.Contract.LightLinkPortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightLinkPortal *LightLinkPortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.LightLinkPortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightLinkPortal *LightLinkPortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.LightLinkPortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightLinkPortal *LightLinkPortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightLinkPortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightLinkPortal *LightLinkPortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightLinkPortal *LightLinkPortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_LightLinkPortal *LightLinkPortalCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightLinkPortal.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_LightLinkPortal *LightLinkPortalSession) Balance() (*big.Int, error) {
	return _LightLinkPortal.Contract.Balance(&_LightLinkPortal.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_LightLinkPortal *LightLinkPortalCallerSession) Balance() (*big.Int, error) {
	return _LightLinkPortal.Contract.Balance(&_LightLinkPortal.CallOpts)
}

// FinalizedWithdrawals is a free data retrieval call binding the contract method 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (_LightLinkPortal *LightLinkPortalCaller) FinalizedWithdrawals(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _LightLinkPortal.contract.Call(opts, &out, "finalizedWithdrawals", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FinalizedWithdrawals is a free data retrieval call binding the contract method 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (_LightLinkPortal *LightLinkPortalSession) FinalizedWithdrawals(arg0 [32]byte) (bool, error) {
	return _LightLinkPortal.Contract.FinalizedWithdrawals(&_LightLinkPortal.CallOpts, arg0)
}

// FinalizedWithdrawals is a free data retrieval call binding the contract method 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (_LightLinkPortal *LightLinkPortalCallerSession) FinalizedWithdrawals(arg0 [32]byte) (bool, error) {
	return _LightLinkPortal.Contract.FinalizedWithdrawals(&_LightLinkPortal.CallOpts, arg0)
}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 _l2OutputIndex) view returns(bool)
func (_LightLinkPortal *LightLinkPortalCaller) IsOutputFinalized(opts *bind.CallOpts, _l2OutputIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _LightLinkPortal.contract.Call(opts, &out, "isOutputFinalized", _l2OutputIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 _l2OutputIndex) view returns(bool)
func (_LightLinkPortal *LightLinkPortalSession) IsOutputFinalized(_l2OutputIndex *big.Int) (bool, error) {
	return _LightLinkPortal.Contract.IsOutputFinalized(&_LightLinkPortal.CallOpts, _l2OutputIndex)
}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 _l2OutputIndex) view returns(bool)
func (_LightLinkPortal *LightLinkPortalCallerSession) IsOutputFinalized(_l2OutputIndex *big.Int) (bool, error) {
	return _LightLinkPortal.Contract.IsOutputFinalized(&_LightLinkPortal.CallOpts, _l2OutputIndex)
}

// L2Oracle is a free data retrieval call binding the contract method 0x9b5f694a.
//
// Solidity: function l2Oracle() view returns(address)
func (_LightLinkPortal *LightLinkPortalCaller) L2Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightLinkPortal.contract.Call(opts, &out, "l2Oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Oracle is a free data retrieval call binding the contract method 0x9b5f694a.
//
// Solidity: function l2Oracle() view returns(address)
func (_LightLinkPortal *LightLinkPortalSession) L2Oracle() (common.Address, error) {
	return _LightLinkPortal.Contract.L2Oracle(&_LightLinkPortal.CallOpts)
}

// L2Oracle is a free data retrieval call binding the contract method 0x9b5f694a.
//
// Solidity: function l2Oracle() view returns(address)
func (_LightLinkPortal *LightLinkPortalCallerSession) L2Oracle() (common.Address, error) {
	return _LightLinkPortal.Contract.L2Oracle(&_LightLinkPortal.CallOpts)
}

// L2Sender is a free data retrieval call binding the contract method 0x9bf62d82.
//
// Solidity: function l2Sender() view returns(address)
func (_LightLinkPortal *LightLinkPortalCaller) L2Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightLinkPortal.contract.Call(opts, &out, "l2Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Sender is a free data retrieval call binding the contract method 0x9bf62d82.
//
// Solidity: function l2Sender() view returns(address)
func (_LightLinkPortal *LightLinkPortalSession) L2Sender() (common.Address, error) {
	return _LightLinkPortal.Contract.L2Sender(&_LightLinkPortal.CallOpts)
}

// L2Sender is a free data retrieval call binding the contract method 0x9bf62d82.
//
// Solidity: function l2Sender() view returns(address)
func (_LightLinkPortal *LightLinkPortalCallerSession) L2Sender() (common.Address, error) {
	return _LightLinkPortal.Contract.L2Sender(&_LightLinkPortal.CallOpts)
}

// MinimumGasLimit is a free data retrieval call binding the contract method 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (_LightLinkPortal *LightLinkPortalCaller) MinimumGasLimit(opts *bind.CallOpts, _byteCount uint64) (uint64, error) {
	var out []interface{}
	err := _LightLinkPortal.contract.Call(opts, &out, "minimumGasLimit", _byteCount)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinimumGasLimit is a free data retrieval call binding the contract method 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (_LightLinkPortal *LightLinkPortalSession) MinimumGasLimit(_byteCount uint64) (uint64, error) {
	return _LightLinkPortal.Contract.MinimumGasLimit(&_LightLinkPortal.CallOpts, _byteCount)
}

// MinimumGasLimit is a free data retrieval call binding the contract method 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (_LightLinkPortal *LightLinkPortalCallerSession) MinimumGasLimit(_byteCount uint64) (uint64, error) {
	return _LightLinkPortal.Contract.MinimumGasLimit(&_LightLinkPortal.CallOpts, _byteCount)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightLinkPortal *LightLinkPortalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightLinkPortal.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightLinkPortal *LightLinkPortalSession) Owner() (common.Address, error) {
	return _LightLinkPortal.Contract.Owner(&_LightLinkPortal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightLinkPortal *LightLinkPortalCallerSession) Owner() (common.Address, error) {
	return _LightLinkPortal.Contract.Owner(&_LightLinkPortal.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum)
func (_LightLinkPortal *LightLinkPortalCaller) Params(opts *bind.CallOpts) (struct {
	PrevBaseFee   *big.Int
	PrevBoughtGas uint64
	PrevBlockNum  uint64
}, error) {
	var out []interface{}
	err := _LightLinkPortal.contract.Call(opts, &out, "params")

	outstruct := new(struct {
		PrevBaseFee   *big.Int
		PrevBoughtGas uint64
		PrevBlockNum  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PrevBaseFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PrevBoughtGas = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.PrevBlockNum = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum)
func (_LightLinkPortal *LightLinkPortalSession) Params() (struct {
	PrevBaseFee   *big.Int
	PrevBoughtGas uint64
	PrevBlockNum  uint64
}, error) {
	return _LightLinkPortal.Contract.Params(&_LightLinkPortal.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum)
func (_LightLinkPortal *LightLinkPortalCallerSession) Params() (struct {
	PrevBaseFee   *big.Int
	PrevBoughtGas uint64
	PrevBlockNum  uint64
}, error) {
	return _LightLinkPortal.Contract.Params(&_LightLinkPortal.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LightLinkPortal *LightLinkPortalCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LightLinkPortal.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LightLinkPortal *LightLinkPortalSession) Paused() (bool, error) {
	return _LightLinkPortal.Contract.Paused(&_LightLinkPortal.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LightLinkPortal *LightLinkPortalCallerSession) Paused() (bool, error) {
	return _LightLinkPortal.Contract.Paused(&_LightLinkPortal.CallOpts)
}

// ProvenWithdrawals is a free data retrieval call binding the contract method 0xe965084c.
//
// Solidity: function provenWithdrawals(bytes32 ) view returns(bytes32 outputRoot, uint128 timestamp, uint128 l2OutputIndex)
func (_LightLinkPortal *LightLinkPortalCaller) ProvenWithdrawals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2OutputIndex *big.Int
}, error) {
	var out []interface{}
	err := _LightLinkPortal.contract.Call(opts, &out, "provenWithdrawals", arg0)

	outstruct := new(struct {
		OutputRoot    [32]byte
		Timestamp     *big.Int
		L2OutputIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutputRoot = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.L2OutputIndex = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProvenWithdrawals is a free data retrieval call binding the contract method 0xe965084c.
//
// Solidity: function provenWithdrawals(bytes32 ) view returns(bytes32 outputRoot, uint128 timestamp, uint128 l2OutputIndex)
func (_LightLinkPortal *LightLinkPortalSession) ProvenWithdrawals(arg0 [32]byte) (struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2OutputIndex *big.Int
}, error) {
	return _LightLinkPortal.Contract.ProvenWithdrawals(&_LightLinkPortal.CallOpts, arg0)
}

// ProvenWithdrawals is a free data retrieval call binding the contract method 0xe965084c.
//
// Solidity: function provenWithdrawals(bytes32 ) view returns(bytes32 outputRoot, uint128 timestamp, uint128 l2OutputIndex)
func (_LightLinkPortal *LightLinkPortalCallerSession) ProvenWithdrawals(arg0 [32]byte) (struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2OutputIndex *big.Int
}, error) {
	return _LightLinkPortal.Contract.ProvenWithdrawals(&_LightLinkPortal.CallOpts, arg0)
}

// ResourceConfig is a free data retrieval call binding the contract method 0xcc731b02.
//
// Solidity: function resourceConfig() view returns(uint32 maxResourceLimit, uint8 elasticityMultiplier, uint8 baseFeeMaxChangeDenominator, uint32 minimumBaseFee, uint32 systemTxMaxGas, uint128 maximumBaseFee)
func (_LightLinkPortal *LightLinkPortalCaller) ResourceConfig(opts *bind.CallOpts) (struct {
	MaxResourceLimit            uint32
	ElasticityMultiplier        uint8
	BaseFeeMaxChangeDenominator uint8
	MinimumBaseFee              uint32
	SystemTxMaxGas              uint32
	MaximumBaseFee              *big.Int
}, error) {
	var out []interface{}
	err := _LightLinkPortal.contract.Call(opts, &out, "resourceConfig")

	outstruct := new(struct {
		MaxResourceLimit            uint32
		ElasticityMultiplier        uint8
		BaseFeeMaxChangeDenominator uint8
		MinimumBaseFee              uint32
		SystemTxMaxGas              uint32
		MaximumBaseFee              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxResourceLimit = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ElasticityMultiplier = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.BaseFeeMaxChangeDenominator = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.MinimumBaseFee = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.SystemTxMaxGas = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.MaximumBaseFee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ResourceConfig is a free data retrieval call binding the contract method 0xcc731b02.
//
// Solidity: function resourceConfig() view returns(uint32 maxResourceLimit, uint8 elasticityMultiplier, uint8 baseFeeMaxChangeDenominator, uint32 minimumBaseFee, uint32 systemTxMaxGas, uint128 maximumBaseFee)
func (_LightLinkPortal *LightLinkPortalSession) ResourceConfig() (struct {
	MaxResourceLimit            uint32
	ElasticityMultiplier        uint8
	BaseFeeMaxChangeDenominator uint8
	MinimumBaseFee              uint32
	SystemTxMaxGas              uint32
	MaximumBaseFee              *big.Int
}, error) {
	return _LightLinkPortal.Contract.ResourceConfig(&_LightLinkPortal.CallOpts)
}

// ResourceConfig is a free data retrieval call binding the contract method 0xcc731b02.
//
// Solidity: function resourceConfig() view returns(uint32 maxResourceLimit, uint8 elasticityMultiplier, uint8 baseFeeMaxChangeDenominator, uint32 minimumBaseFee, uint32 systemTxMaxGas, uint128 maximumBaseFee)
func (_LightLinkPortal *LightLinkPortalCallerSession) ResourceConfig() (struct {
	MaxResourceLimit            uint32
	ElasticityMultiplier        uint8
	BaseFeeMaxChangeDenominator uint8
	MinimumBaseFee              uint32
	SystemTxMaxGas              uint32
	MaximumBaseFee              *big.Int
}, error) {
	return _LightLinkPortal.Contract.ResourceConfig(&_LightLinkPortal.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_LightLinkPortal *LightLinkPortalCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LightLinkPortal.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_LightLinkPortal *LightLinkPortalSession) Version() (string, error) {
	return _LightLinkPortal.Contract.Version(&_LightLinkPortal.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_LightLinkPortal *LightLinkPortalCallerSession) Version() (string, error) {
	return _LightLinkPortal.Contract.Version(&_LightLinkPortal.CallOpts)
}

// DepositERC20Transaction is a paid mutator transaction binding the contract method 0x149f2f22.
//
// Solidity: function depositERC20Transaction(address _to, uint256 _mint, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) returns()
func (_LightLinkPortal *LightLinkPortalTransactor) DepositERC20Transaction(opts *bind.TransactOpts, _to common.Address, _mint *big.Int, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _LightLinkPortal.contract.Transact(opts, "depositERC20Transaction", _to, _mint, _value, _gasLimit, _isCreation, _data)
}

// DepositERC20Transaction is a paid mutator transaction binding the contract method 0x149f2f22.
//
// Solidity: function depositERC20Transaction(address _to, uint256 _mint, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) returns()
func (_LightLinkPortal *LightLinkPortalSession) DepositERC20Transaction(_to common.Address, _mint *big.Int, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.DepositERC20Transaction(&_LightLinkPortal.TransactOpts, _to, _mint, _value, _gasLimit, _isCreation, _data)
}

// DepositERC20Transaction is a paid mutator transaction binding the contract method 0x149f2f22.
//
// Solidity: function depositERC20Transaction(address _to, uint256 _mint, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) returns()
func (_LightLinkPortal *LightLinkPortalTransactorSession) DepositERC20Transaction(_to common.Address, _mint *big.Int, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.DepositERC20Transaction(&_LightLinkPortal.TransactOpts, _to, _mint, _value, _gasLimit, _isCreation, _data)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0xe9e05c42.
//
// Solidity: function depositTransaction(address _to, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_LightLinkPortal *LightLinkPortalTransactor) DepositTransaction(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _LightLinkPortal.contract.Transact(opts, "depositTransaction", _to, _value, _gasLimit, _isCreation, _data)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0xe9e05c42.
//
// Solidity: function depositTransaction(address _to, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_LightLinkPortal *LightLinkPortalSession) DepositTransaction(_to common.Address, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.DepositTransaction(&_LightLinkPortal.TransactOpts, _to, _value, _gasLimit, _isCreation, _data)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0xe9e05c42.
//
// Solidity: function depositTransaction(address _to, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_LightLinkPortal *LightLinkPortalTransactorSession) DepositTransaction(_to common.Address, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.DepositTransaction(&_LightLinkPortal.TransactOpts, _to, _value, _gasLimit, _isCreation, _data)
}

// DonateETH is a paid mutator transaction binding the contract method 0x8b4c40b0.
//
// Solidity: function donateETH() payable returns()
func (_LightLinkPortal *LightLinkPortalTransactor) DonateETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightLinkPortal.contract.Transact(opts, "donateETH")
}

// DonateETH is a paid mutator transaction binding the contract method 0x8b4c40b0.
//
// Solidity: function donateETH() payable returns()
func (_LightLinkPortal *LightLinkPortalSession) DonateETH() (*types.Transaction, error) {
	return _LightLinkPortal.Contract.DonateETH(&_LightLinkPortal.TransactOpts)
}

// DonateETH is a paid mutator transaction binding the contract method 0x8b4c40b0.
//
// Solidity: function donateETH() payable returns()
func (_LightLinkPortal *LightLinkPortalTransactorSession) DonateETH() (*types.Transaction, error) {
	return _LightLinkPortal.Contract.DonateETH(&_LightLinkPortal.TransactOpts)
}

// FinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x8c3152e9.
//
// Solidity: function finalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx) returns()
func (_LightLinkPortal *LightLinkPortalTransactor) FinalizeWithdrawalTransaction(opts *bind.TransactOpts, _tx TypesWithdrawalTransaction) (*types.Transaction, error) {
	return _LightLinkPortal.contract.Transact(opts, "finalizeWithdrawalTransaction", _tx)
}

// FinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x8c3152e9.
//
// Solidity: function finalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx) returns()
func (_LightLinkPortal *LightLinkPortalSession) FinalizeWithdrawalTransaction(_tx TypesWithdrawalTransaction) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.FinalizeWithdrawalTransaction(&_LightLinkPortal.TransactOpts, _tx)
}

// FinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x8c3152e9.
//
// Solidity: function finalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx) returns()
func (_LightLinkPortal *LightLinkPortalTransactorSession) FinalizeWithdrawalTransaction(_tx TypesWithdrawalTransaction) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.FinalizeWithdrawalTransaction(&_LightLinkPortal.TransactOpts, _tx)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _l2Oracle, address _challenge) returns()
func (_LightLinkPortal *LightLinkPortalTransactor) Initialize(opts *bind.TransactOpts, _l2Oracle common.Address, _challenge common.Address) (*types.Transaction, error) {
	return _LightLinkPortal.contract.Transact(opts, "initialize", _l2Oracle, _challenge)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _l2Oracle, address _challenge) returns()
func (_LightLinkPortal *LightLinkPortalSession) Initialize(_l2Oracle common.Address, _challenge common.Address) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.Initialize(&_LightLinkPortal.TransactOpts, _l2Oracle, _challenge)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _l2Oracle, address _challenge) returns()
func (_LightLinkPortal *LightLinkPortalTransactorSession) Initialize(_l2Oracle common.Address, _challenge common.Address) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.Initialize(&_LightLinkPortal.TransactOpts, _l2Oracle, _challenge)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LightLinkPortal *LightLinkPortalTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightLinkPortal.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LightLinkPortal *LightLinkPortalSession) Pause() (*types.Transaction, error) {
	return _LightLinkPortal.Contract.Pause(&_LightLinkPortal.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LightLinkPortal *LightLinkPortalTransactorSession) Pause() (*types.Transaction, error) {
	return _LightLinkPortal.Contract.Pause(&_LightLinkPortal.TransactOpts)
}

// ProveWithdrawalTransaction is a paid mutator transaction binding the contract method 0x4870496f.
//
// Solidity: function proveWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx, uint256 _l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _withdrawalProof) returns()
func (_LightLinkPortal *LightLinkPortalTransactor) ProveWithdrawalTransaction(opts *bind.TransactOpts, _tx TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _LightLinkPortal.contract.Transact(opts, "proveWithdrawalTransaction", _tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
}

// ProveWithdrawalTransaction is a paid mutator transaction binding the contract method 0x4870496f.
//
// Solidity: function proveWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx, uint256 _l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _withdrawalProof) returns()
func (_LightLinkPortal *LightLinkPortalSession) ProveWithdrawalTransaction(_tx TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.ProveWithdrawalTransaction(&_LightLinkPortal.TransactOpts, _tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
}

// ProveWithdrawalTransaction is a paid mutator transaction binding the contract method 0x4870496f.
//
// Solidity: function proveWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx, uint256 _l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _withdrawalProof) returns()
func (_LightLinkPortal *LightLinkPortalTransactorSession) ProveWithdrawalTransaction(_tx TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.ProveWithdrawalTransaction(&_LightLinkPortal.TransactOpts, _tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightLinkPortal *LightLinkPortalTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightLinkPortal.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightLinkPortal *LightLinkPortalSession) RenounceOwnership() (*types.Transaction, error) {
	return _LightLinkPortal.Contract.RenounceOwnership(&_LightLinkPortal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightLinkPortal *LightLinkPortalTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LightLinkPortal.Contract.RenounceOwnership(&_LightLinkPortal.TransactOpts)
}

// SetGasPayingToken is a paid mutator transaction binding the contract method 0x71cfaa3f.
//
// Solidity: function setGasPayingToken(address _token, uint8 _decimals, bytes32 _name, bytes32 _symbol) returns()
func (_LightLinkPortal *LightLinkPortalTransactor) SetGasPayingToken(opts *bind.TransactOpts, _token common.Address, _decimals uint8, _name [32]byte, _symbol [32]byte) (*types.Transaction, error) {
	return _LightLinkPortal.contract.Transact(opts, "setGasPayingToken", _token, _decimals, _name, _symbol)
}

// SetGasPayingToken is a paid mutator transaction binding the contract method 0x71cfaa3f.
//
// Solidity: function setGasPayingToken(address _token, uint8 _decimals, bytes32 _name, bytes32 _symbol) returns()
func (_LightLinkPortal *LightLinkPortalSession) SetGasPayingToken(_token common.Address, _decimals uint8, _name [32]byte, _symbol [32]byte) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.SetGasPayingToken(&_LightLinkPortal.TransactOpts, _token, _decimals, _name, _symbol)
}

// SetGasPayingToken is a paid mutator transaction binding the contract method 0x71cfaa3f.
//
// Solidity: function setGasPayingToken(address _token, uint8 _decimals, bytes32 _name, bytes32 _symbol) returns()
func (_LightLinkPortal *LightLinkPortalTransactorSession) SetGasPayingToken(_token common.Address, _decimals uint8, _name [32]byte, _symbol [32]byte) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.SetGasPayingToken(&_LightLinkPortal.TransactOpts, _token, _decimals, _name, _symbol)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightLinkPortal *LightLinkPortalTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LightLinkPortal.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightLinkPortal *LightLinkPortalSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.TransferOwnership(&_LightLinkPortal.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightLinkPortal *LightLinkPortalTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LightLinkPortal.Contract.TransferOwnership(&_LightLinkPortal.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LightLinkPortal *LightLinkPortalTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightLinkPortal.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LightLinkPortal *LightLinkPortalSession) Unpause() (*types.Transaction, error) {
	return _LightLinkPortal.Contract.Unpause(&_LightLinkPortal.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LightLinkPortal *LightLinkPortalTransactorSession) Unpause() (*types.Transaction, error) {
	return _LightLinkPortal.Contract.Unpause(&_LightLinkPortal.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LightLinkPortal *LightLinkPortalTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightLinkPortal.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LightLinkPortal *LightLinkPortalSession) Receive() (*types.Transaction, error) {
	return _LightLinkPortal.Contract.Receive(&_LightLinkPortal.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LightLinkPortal *LightLinkPortalTransactorSession) Receive() (*types.Transaction, error) {
	return _LightLinkPortal.Contract.Receive(&_LightLinkPortal.TransactOpts)
}

// LightLinkPortalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LightLinkPortal contract.
type LightLinkPortalInitializedIterator struct {
	Event *LightLinkPortalInitialized // Event containing the contract specifics and raw log

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
func (it *LightLinkPortalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightLinkPortalInitialized)
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
		it.Event = new(LightLinkPortalInitialized)
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
func (it *LightLinkPortalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightLinkPortalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightLinkPortalInitialized represents a Initialized event raised by the LightLinkPortal contract.
type LightLinkPortalInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LightLinkPortal *LightLinkPortalFilterer) FilterInitialized(opts *bind.FilterOpts) (*LightLinkPortalInitializedIterator, error) {

	logs, sub, err := _LightLinkPortal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LightLinkPortalInitializedIterator{contract: _LightLinkPortal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LightLinkPortal *LightLinkPortalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LightLinkPortalInitialized) (event.Subscription, error) {

	logs, sub, err := _LightLinkPortal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightLinkPortalInitialized)
				if err := _LightLinkPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LightLinkPortal *LightLinkPortalFilterer) ParseInitialized(log types.Log) (*LightLinkPortalInitialized, error) {
	event := new(LightLinkPortalInitialized)
	if err := _LightLinkPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightLinkPortalOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LightLinkPortal contract.
type LightLinkPortalOwnershipTransferredIterator struct {
	Event *LightLinkPortalOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LightLinkPortalOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightLinkPortalOwnershipTransferred)
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
		it.Event = new(LightLinkPortalOwnershipTransferred)
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
func (it *LightLinkPortalOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightLinkPortalOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightLinkPortalOwnershipTransferred represents a OwnershipTransferred event raised by the LightLinkPortal contract.
type LightLinkPortalOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LightLinkPortal *LightLinkPortalFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LightLinkPortalOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LightLinkPortal.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LightLinkPortalOwnershipTransferredIterator{contract: _LightLinkPortal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LightLinkPortal *LightLinkPortalFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LightLinkPortalOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LightLinkPortal.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightLinkPortalOwnershipTransferred)
				if err := _LightLinkPortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LightLinkPortal *LightLinkPortalFilterer) ParseOwnershipTransferred(log types.Log) (*LightLinkPortalOwnershipTransferred, error) {
	event := new(LightLinkPortalOwnershipTransferred)
	if err := _LightLinkPortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightLinkPortalPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the LightLinkPortal contract.
type LightLinkPortalPauseIterator struct {
	Event *LightLinkPortalPause // Event containing the contract specifics and raw log

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
func (it *LightLinkPortalPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightLinkPortalPause)
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
		it.Event = new(LightLinkPortalPause)
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
func (it *LightLinkPortalPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightLinkPortalPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightLinkPortalPause represents a Pause event raised by the LightLinkPortal contract.
type LightLinkPortalPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_LightLinkPortal *LightLinkPortalFilterer) FilterPause(opts *bind.FilterOpts) (*LightLinkPortalPauseIterator, error) {

	logs, sub, err := _LightLinkPortal.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &LightLinkPortalPauseIterator{contract: _LightLinkPortal.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_LightLinkPortal *LightLinkPortalFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *LightLinkPortalPause) (event.Subscription, error) {

	logs, sub, err := _LightLinkPortal.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightLinkPortalPause)
				if err := _LightLinkPortal.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_LightLinkPortal *LightLinkPortalFilterer) ParsePause(log types.Log) (*LightLinkPortalPause, error) {
	event := new(LightLinkPortalPause)
	if err := _LightLinkPortal.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightLinkPortalTransactionDepositedIterator is returned from FilterTransactionDeposited and is used to iterate over the raw logs and unpacked data for TransactionDeposited events raised by the LightLinkPortal contract.
type LightLinkPortalTransactionDepositedIterator struct {
	Event *LightLinkPortalTransactionDeposited // Event containing the contract specifics and raw log

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
func (it *LightLinkPortalTransactionDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightLinkPortalTransactionDeposited)
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
		it.Event = new(LightLinkPortalTransactionDeposited)
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
func (it *LightLinkPortalTransactionDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightLinkPortalTransactionDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightLinkPortalTransactionDeposited represents a TransactionDeposited event raised by the LightLinkPortal contract.
type LightLinkPortalTransactionDeposited struct {
	From       common.Address
	To         common.Address
	Version    *big.Int
	OpaqueData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionDeposited is a free log retrieval operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_LightLinkPortal *LightLinkPortalFilterer) FilterTransactionDeposited(opts *bind.FilterOpts, from []common.Address, to []common.Address, version []*big.Int) (*LightLinkPortalTransactionDepositedIterator, error) {

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

	logs, sub, err := _LightLinkPortal.contract.FilterLogs(opts, "TransactionDeposited", fromRule, toRule, versionRule)
	if err != nil {
		return nil, err
	}
	return &LightLinkPortalTransactionDepositedIterator{contract: _LightLinkPortal.contract, event: "TransactionDeposited", logs: logs, sub: sub}, nil
}

// WatchTransactionDeposited is a free log subscription operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_LightLinkPortal *LightLinkPortalFilterer) WatchTransactionDeposited(opts *bind.WatchOpts, sink chan<- *LightLinkPortalTransactionDeposited, from []common.Address, to []common.Address, version []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _LightLinkPortal.contract.WatchLogs(opts, "TransactionDeposited", fromRule, toRule, versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightLinkPortalTransactionDeposited)
				if err := _LightLinkPortal.contract.UnpackLog(event, "TransactionDeposited", log); err != nil {
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
func (_LightLinkPortal *LightLinkPortalFilterer) ParseTransactionDeposited(log types.Log) (*LightLinkPortalTransactionDeposited, error) {
	event := new(LightLinkPortalTransactionDeposited)
	if err := _LightLinkPortal.contract.UnpackLog(event, "TransactionDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightLinkPortalUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the LightLinkPortal contract.
type LightLinkPortalUnpauseIterator struct {
	Event *LightLinkPortalUnpause // Event containing the contract specifics and raw log

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
func (it *LightLinkPortalUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightLinkPortalUnpause)
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
		it.Event = new(LightLinkPortalUnpause)
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
func (it *LightLinkPortalUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightLinkPortalUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightLinkPortalUnpause represents a Unpause event raised by the LightLinkPortal contract.
type LightLinkPortalUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_LightLinkPortal *LightLinkPortalFilterer) FilterUnpause(opts *bind.FilterOpts) (*LightLinkPortalUnpauseIterator, error) {

	logs, sub, err := _LightLinkPortal.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &LightLinkPortalUnpauseIterator{contract: _LightLinkPortal.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_LightLinkPortal *LightLinkPortalFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *LightLinkPortalUnpause) (event.Subscription, error) {

	logs, sub, err := _LightLinkPortal.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightLinkPortalUnpause)
				if err := _LightLinkPortal.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_LightLinkPortal *LightLinkPortalFilterer) ParseUnpause(log types.Log) (*LightLinkPortalUnpause, error) {
	event := new(LightLinkPortalUnpause)
	if err := _LightLinkPortal.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightLinkPortalWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the LightLinkPortal contract.
type LightLinkPortalWithdrawalFinalizedIterator struct {
	Event *LightLinkPortalWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *LightLinkPortalWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightLinkPortalWithdrawalFinalized)
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
		it.Event = new(LightLinkPortalWithdrawalFinalized)
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
func (it *LightLinkPortalWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightLinkPortalWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightLinkPortalWithdrawalFinalized represents a WithdrawalFinalized event raised by the LightLinkPortal contract.
type LightLinkPortalWithdrawalFinalized struct {
	WithdrawalHash [32]byte
	Success        bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_LightLinkPortal *LightLinkPortalFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, withdrawalHash [][32]byte) (*LightLinkPortalWithdrawalFinalizedIterator, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}

	logs, sub, err := _LightLinkPortal.contract.FilterLogs(opts, "WithdrawalFinalized", withdrawalHashRule)
	if err != nil {
		return nil, err
	}
	return &LightLinkPortalWithdrawalFinalizedIterator{contract: _LightLinkPortal.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_LightLinkPortal *LightLinkPortalFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *LightLinkPortalWithdrawalFinalized, withdrawalHash [][32]byte) (event.Subscription, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}

	logs, sub, err := _LightLinkPortal.contract.WatchLogs(opts, "WithdrawalFinalized", withdrawalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightLinkPortalWithdrawalFinalized)
				if err := _LightLinkPortal.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_LightLinkPortal *LightLinkPortalFilterer) ParseWithdrawalFinalized(log types.Log) (*LightLinkPortalWithdrawalFinalized, error) {
	event := new(LightLinkPortalWithdrawalFinalized)
	if err := _LightLinkPortal.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightLinkPortalWithdrawalProvenIterator is returned from FilterWithdrawalProven and is used to iterate over the raw logs and unpacked data for WithdrawalProven events raised by the LightLinkPortal contract.
type LightLinkPortalWithdrawalProvenIterator struct {
	Event *LightLinkPortalWithdrawalProven // Event containing the contract specifics and raw log

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
func (it *LightLinkPortalWithdrawalProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightLinkPortalWithdrawalProven)
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
		it.Event = new(LightLinkPortalWithdrawalProven)
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
func (it *LightLinkPortalWithdrawalProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightLinkPortalWithdrawalProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightLinkPortalWithdrawalProven represents a WithdrawalProven event raised by the LightLinkPortal contract.
type LightLinkPortalWithdrawalProven struct {
	WithdrawalHash [32]byte
	From           common.Address
	To             common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalProven is a free log retrieval operation binding the contract event 0x67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f62.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (_LightLinkPortal *LightLinkPortalFilterer) FilterWithdrawalProven(opts *bind.FilterOpts, withdrawalHash [][32]byte, from []common.Address, to []common.Address) (*LightLinkPortalWithdrawalProvenIterator, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LightLinkPortal.contract.FilterLogs(opts, "WithdrawalProven", withdrawalHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LightLinkPortalWithdrawalProvenIterator{contract: _LightLinkPortal.contract, event: "WithdrawalProven", logs: logs, sub: sub}, nil
}

// WatchWithdrawalProven is a free log subscription operation binding the contract event 0x67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f62.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (_LightLinkPortal *LightLinkPortalFilterer) WatchWithdrawalProven(opts *bind.WatchOpts, sink chan<- *LightLinkPortalWithdrawalProven, withdrawalHash [][32]byte, from []common.Address, to []common.Address) (event.Subscription, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LightLinkPortal.contract.WatchLogs(opts, "WithdrawalProven", withdrawalHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightLinkPortalWithdrawalProven)
				if err := _LightLinkPortal.contract.UnpackLog(event, "WithdrawalProven", log); err != nil {
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

// ParseWithdrawalProven is a log parse operation binding the contract event 0x67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f62.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (_LightLinkPortal *LightLinkPortalFilterer) ParseWithdrawalProven(log types.Log) (*LightLinkPortalWithdrawalProven, error) {
	event := new(LightLinkPortalWithdrawalProven)
	if err := _LightLinkPortal.contract.UnpackLog(event, "WithdrawalProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
