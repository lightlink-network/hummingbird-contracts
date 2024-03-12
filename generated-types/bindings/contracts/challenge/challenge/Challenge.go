// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challenge

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

// BinaryMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type BinaryMerkleProof struct {
	SideNodes [][32]byte
	Key       *big.Int
	NumLeaves *big.Int
}

// ChallengeDataAvailabilityChallengeDAProof is an auto generated low-level Go binding around an user-defined struct.
type ChallengeDataAvailabilityChallengeDAProof struct {
	RootNonce     *big.Int
	DataRootTuple DataRootTuple
	Proof         BinaryMerkleProof
}

// ChallengeL2HeaderL2HeaderPointer is an auto generated low-level Go binding around an user-defined struct.
type ChallengeL2HeaderL2HeaderPointer struct {
	Rblock [32]byte
	Number *big.Int
}

// DataRootTuple is an auto generated low-level Go binding around an user-defined struct.
type DataRootTuple struct {
	Height   *big.Int
	DataRoot [32]byte
}

// ChallengeMetaData contains all meta data concerning the Challenge contract.
var ChallengeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_blockHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_blockIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_expiry\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumChallengeDataAvailability.ChallengeDAStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"ChallengeDAUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_blockIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"enumChallengeHeader.InvalidHeaderReason\",\"name\":\"_reason\",\"type\":\"uint8\"}],\"name\":\"InvalidHeader\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"challengeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2Number\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rblock\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumChallengeL2Header.L2HeaderChallengeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"L2HeaderChallengeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chain\",\"outputs\":[{\"internalType\":\"contractICanonicalStateChain\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainOracle\",\"outputs\":[{\"internalType\":\"contractIChainOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockIndex\",\"type\":\"uint256\"}],\"name\":\"challengeDataRootInclusion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rblockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2Num\",\"type\":\"uint256\"}],\"name\":\"challengeL2Header\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"daChallenges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"enumChallengeDataAvailability.ChallengeDAStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daOracle\",\"outputs\":[{\"internalType\":\"contractIDAOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"dataRootTuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structChallengeDataAvailability.ChallengeDAProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"defendDataRootInclusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_headerHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_headerPrevHash\",\"type\":\"bytes32\"}],\"name\":\"defendL2Header\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mipsChallenge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_chainOracle\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockIndex\",\"type\":\"uint256\"}],\"name\":\"invalidateHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_rblockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2Num\",\"type\":\"uint256\"}],\"name\":\"l2HeaderChallengeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"l2HeaderChallenges\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rblock\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeL2Header.L2HeaderPointer\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rblock\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeL2Header.L2HeaderPointer\",\"name\":\"prevHeader\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"challengeEnd\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"enumChallengeL2Header.L2HeaderChallengeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mipsChallenge\",\"outputs\":[{\"internalType\":\"contractIMipsChallenge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_challengeFee\",\"type\":\"uint256\"}],\"name\":\"setChallengeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_challengePeriod\",\"type\":\"uint256\"}],\"name\":\"setChallengePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_challengeReward\",\"type\":\"uint256\"}],\"name\":\"setChallengeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_challengeWindow\",\"type\":\"uint256\"}],\"name\":\"setChallengeWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defender\",\"type\":\"address\"}],\"name\":\"setDefender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_blockhash\",\"type\":\"bytes32\"}],\"name\":\"settleDataRootInclusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"}],\"name\":\"settleL2HeaderChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"contractITreasury\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ChallengeABI is the input ABI used to generate the binding from.
// Deprecated: Use ChallengeMetaData.ABI instead.
var ChallengeABI = ChallengeMetaData.ABI

// Challenge is an auto generated Go binding around an Ethereum contract.
type Challenge struct {
	ChallengeCaller     // Read-only binding to the contract
	ChallengeTransactor // Write-only binding to the contract
	ChallengeFilterer   // Log filterer for contract events
}

// ChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeSession struct {
	Contract     *Challenge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeCallerSession struct {
	Contract *ChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeTransactorSession struct {
	Contract     *ChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeRaw struct {
	Contract *Challenge // Generic contract binding to access the raw methods on
}

// ChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeCallerRaw struct {
	Contract *ChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeTransactorRaw struct {
	Contract *ChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallenge creates a new instance of Challenge, bound to a specific deployed contract.
func NewChallenge(address common.Address, backend bind.ContractBackend) (*Challenge, error) {
	contract, err := bindChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Challenge{ChallengeCaller: ChallengeCaller{contract: contract}, ChallengeTransactor: ChallengeTransactor{contract: contract}, ChallengeFilterer: ChallengeFilterer{contract: contract}}, nil
}

// NewChallengeCaller creates a new read-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeCaller(address common.Address, caller bind.ContractCaller) (*ChallengeCaller, error) {
	contract, err := bindChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeCaller{contract: contract}, nil
}

// NewChallengeTransactor creates a new write-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeTransactor, error) {
	contract, err := bindChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTransactor{contract: contract}, nil
}

// NewChallengeFilterer creates a new log filterer instance of Challenge, bound to a specific deployed contract.
func NewChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeFilterer, error) {
	contract, err := bindChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeFilterer{contract: contract}, nil
}

// bindChallenge binds a generic wrapper to an already deployed contract.
func bindChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChallengeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.ChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Challenge *ChallengeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Challenge *ChallengeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Challenge.Contract.UPGRADEINTERFACEVERSION(&_Challenge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Challenge *ChallengeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Challenge.Contract.UPGRADEINTERFACEVERSION(&_Challenge.CallOpts)
}

// Chain is a free data retrieval call binding the contract method 0xc763e5a1.
//
// Solidity: function chain() view returns(address)
func (_Challenge *ChallengeCaller) Chain(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "chain")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chain is a free data retrieval call binding the contract method 0xc763e5a1.
//
// Solidity: function chain() view returns(address)
func (_Challenge *ChallengeSession) Chain() (common.Address, error) {
	return _Challenge.Contract.Chain(&_Challenge.CallOpts)
}

// Chain is a free data retrieval call binding the contract method 0xc763e5a1.
//
// Solidity: function chain() view returns(address)
func (_Challenge *ChallengeCallerSession) Chain() (common.Address, error) {
	return _Challenge.Contract.Chain(&_Challenge.CallOpts)
}

// ChainOracle is a free data retrieval call binding the contract method 0xbfcf4495.
//
// Solidity: function chainOracle() view returns(address)
func (_Challenge *ChallengeCaller) ChainOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "chainOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChainOracle is a free data retrieval call binding the contract method 0xbfcf4495.
//
// Solidity: function chainOracle() view returns(address)
func (_Challenge *ChallengeSession) ChainOracle() (common.Address, error) {
	return _Challenge.Contract.ChainOracle(&_Challenge.CallOpts)
}

// ChainOracle is a free data retrieval call binding the contract method 0xbfcf4495.
//
// Solidity: function chainOracle() view returns(address)
func (_Challenge *ChallengeCallerSession) ChainOracle() (common.Address, error) {
	return _Challenge.Contract.ChainOracle(&_Challenge.CallOpts)
}

// ChallengeFee is a free data retrieval call binding the contract method 0x1bd8f9ca.
//
// Solidity: function challengeFee() view returns(uint256)
func (_Challenge *ChallengeCaller) ChallengeFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "challengeFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeFee is a free data retrieval call binding the contract method 0x1bd8f9ca.
//
// Solidity: function challengeFee() view returns(uint256)
func (_Challenge *ChallengeSession) ChallengeFee() (*big.Int, error) {
	return _Challenge.Contract.ChallengeFee(&_Challenge.CallOpts)
}

// ChallengeFee is a free data retrieval call binding the contract method 0x1bd8f9ca.
//
// Solidity: function challengeFee() view returns(uint256)
func (_Challenge *ChallengeCallerSession) ChallengeFee() (*big.Int, error) {
	return _Challenge.Contract.ChallengeFee(&_Challenge.CallOpts)
}

// ChallengePeriod is a free data retrieval call binding the contract method 0xf3f480d9.
//
// Solidity: function challengePeriod() view returns(uint256)
func (_Challenge *ChallengeCaller) ChallengePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "challengePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengePeriod is a free data retrieval call binding the contract method 0xf3f480d9.
//
// Solidity: function challengePeriod() view returns(uint256)
func (_Challenge *ChallengeSession) ChallengePeriod() (*big.Int, error) {
	return _Challenge.Contract.ChallengePeriod(&_Challenge.CallOpts)
}

// ChallengePeriod is a free data retrieval call binding the contract method 0xf3f480d9.
//
// Solidity: function challengePeriod() view returns(uint256)
func (_Challenge *ChallengeCallerSession) ChallengePeriod() (*big.Int, error) {
	return _Challenge.Contract.ChallengePeriod(&_Challenge.CallOpts)
}

// ChallengeReward is a free data retrieval call binding the contract method 0x3ea0c15e.
//
// Solidity: function challengeReward() view returns(uint256)
func (_Challenge *ChallengeCaller) ChallengeReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "challengeReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeReward is a free data retrieval call binding the contract method 0x3ea0c15e.
//
// Solidity: function challengeReward() view returns(uint256)
func (_Challenge *ChallengeSession) ChallengeReward() (*big.Int, error) {
	return _Challenge.Contract.ChallengeReward(&_Challenge.CallOpts)
}

// ChallengeReward is a free data retrieval call binding the contract method 0x3ea0c15e.
//
// Solidity: function challengeReward() view returns(uint256)
func (_Challenge *ChallengeCallerSession) ChallengeReward() (*big.Int, error) {
	return _Challenge.Contract.ChallengeReward(&_Challenge.CallOpts)
}

// ChallengeWindow is a free data retrieval call binding the contract method 0x861a1412.
//
// Solidity: function challengeWindow() view returns(uint256)
func (_Challenge *ChallengeCaller) ChallengeWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "challengeWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeWindow is a free data retrieval call binding the contract method 0x861a1412.
//
// Solidity: function challengeWindow() view returns(uint256)
func (_Challenge *ChallengeSession) ChallengeWindow() (*big.Int, error) {
	return _Challenge.Contract.ChallengeWindow(&_Challenge.CallOpts)
}

// ChallengeWindow is a free data retrieval call binding the contract method 0x861a1412.
//
// Solidity: function challengeWindow() view returns(uint256)
func (_Challenge *ChallengeCallerSession) ChallengeWindow() (*big.Int, error) {
	return _Challenge.Contract.ChallengeWindow(&_Challenge.CallOpts)
}

// DaChallenges is a free data retrieval call binding the contract method 0x113e70fb.
//
// Solidity: function daChallenges(bytes32 ) view returns(uint256 blockIndex, address challenger, uint256 expiry, uint8 status)
func (_Challenge *ChallengeCaller) DaChallenges(opts *bind.CallOpts, arg0 [32]byte) (struct {
	BlockIndex *big.Int
	Challenger common.Address
	Expiry     *big.Int
	Status     uint8
}, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "daChallenges", arg0)

	outstruct := new(struct {
		BlockIndex *big.Int
		Challenger common.Address
		Expiry     *big.Int
		Status     uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockIndex = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Challenger = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Expiry = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// DaChallenges is a free data retrieval call binding the contract method 0x113e70fb.
//
// Solidity: function daChallenges(bytes32 ) view returns(uint256 blockIndex, address challenger, uint256 expiry, uint8 status)
func (_Challenge *ChallengeSession) DaChallenges(arg0 [32]byte) (struct {
	BlockIndex *big.Int
	Challenger common.Address
	Expiry     *big.Int
	Status     uint8
}, error) {
	return _Challenge.Contract.DaChallenges(&_Challenge.CallOpts, arg0)
}

// DaChallenges is a free data retrieval call binding the contract method 0x113e70fb.
//
// Solidity: function daChallenges(bytes32 ) view returns(uint256 blockIndex, address challenger, uint256 expiry, uint8 status)
func (_Challenge *ChallengeCallerSession) DaChallenges(arg0 [32]byte) (struct {
	BlockIndex *big.Int
	Challenger common.Address
	Expiry     *big.Int
	Status     uint8
}, error) {
	return _Challenge.Contract.DaChallenges(&_Challenge.CallOpts, arg0)
}

// DaOracle is a free data retrieval call binding the contract method 0xee223c02.
//
// Solidity: function daOracle() view returns(address)
func (_Challenge *ChallengeCaller) DaOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "daOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaOracle is a free data retrieval call binding the contract method 0xee223c02.
//
// Solidity: function daOracle() view returns(address)
func (_Challenge *ChallengeSession) DaOracle() (common.Address, error) {
	return _Challenge.Contract.DaOracle(&_Challenge.CallOpts)
}

// DaOracle is a free data retrieval call binding the contract method 0xee223c02.
//
// Solidity: function daOracle() view returns(address)
func (_Challenge *ChallengeCallerSession) DaOracle() (common.Address, error) {
	return _Challenge.Contract.DaOracle(&_Challenge.CallOpts)
}

// Defender is a free data retrieval call binding the contract method 0x7f4c91c5.
//
// Solidity: function defender() view returns(address)
func (_Challenge *ChallengeCaller) Defender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "defender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Defender is a free data retrieval call binding the contract method 0x7f4c91c5.
//
// Solidity: function defender() view returns(address)
func (_Challenge *ChallengeSession) Defender() (common.Address, error) {
	return _Challenge.Contract.Defender(&_Challenge.CallOpts)
}

// Defender is a free data retrieval call binding the contract method 0x7f4c91c5.
//
// Solidity: function defender() view returns(address)
func (_Challenge *ChallengeCallerSession) Defender() (common.Address, error) {
	return _Challenge.Contract.Defender(&_Challenge.CallOpts)
}

// L2HeaderChallengeHash is a free data retrieval call binding the contract method 0xfa8e8de2.
//
// Solidity: function l2HeaderChallengeHash(bytes32 _rblockHash, uint256 _l2Num) pure returns(bytes32)
func (_Challenge *ChallengeCaller) L2HeaderChallengeHash(opts *bind.CallOpts, _rblockHash [32]byte, _l2Num *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "l2HeaderChallengeHash", _rblockHash, _l2Num)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2HeaderChallengeHash is a free data retrieval call binding the contract method 0xfa8e8de2.
//
// Solidity: function l2HeaderChallengeHash(bytes32 _rblockHash, uint256 _l2Num) pure returns(bytes32)
func (_Challenge *ChallengeSession) L2HeaderChallengeHash(_rblockHash [32]byte, _l2Num *big.Int) ([32]byte, error) {
	return _Challenge.Contract.L2HeaderChallengeHash(&_Challenge.CallOpts, _rblockHash, _l2Num)
}

// L2HeaderChallengeHash is a free data retrieval call binding the contract method 0xfa8e8de2.
//
// Solidity: function l2HeaderChallengeHash(bytes32 _rblockHash, uint256 _l2Num) pure returns(bytes32)
func (_Challenge *ChallengeCallerSession) L2HeaderChallengeHash(_rblockHash [32]byte, _l2Num *big.Int) ([32]byte, error) {
	return _Challenge.Contract.L2HeaderChallengeHash(&_Challenge.CallOpts, _rblockHash, _l2Num)
}

// L2HeaderChallenges is a free data retrieval call binding the contract method 0x6da802c8.
//
// Solidity: function l2HeaderChallenges(bytes32 ) view returns((bytes32,uint256) header, (bytes32,uint256) prevHeader, uint256 challengeEnd, address challenger, uint8 status)
func (_Challenge *ChallengeCaller) L2HeaderChallenges(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Header       ChallengeL2HeaderL2HeaderPointer
	PrevHeader   ChallengeL2HeaderL2HeaderPointer
	ChallengeEnd *big.Int
	Challenger   common.Address
	Status       uint8
}, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "l2HeaderChallenges", arg0)

	outstruct := new(struct {
		Header       ChallengeL2HeaderL2HeaderPointer
		PrevHeader   ChallengeL2HeaderL2HeaderPointer
		ChallengeEnd *big.Int
		Challenger   common.Address
		Status       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Header = *abi.ConvertType(out[0], new(ChallengeL2HeaderL2HeaderPointer)).(*ChallengeL2HeaderL2HeaderPointer)
	outstruct.PrevHeader = *abi.ConvertType(out[1], new(ChallengeL2HeaderL2HeaderPointer)).(*ChallengeL2HeaderL2HeaderPointer)
	outstruct.ChallengeEnd = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Challenger = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return *outstruct, err

}

// L2HeaderChallenges is a free data retrieval call binding the contract method 0x6da802c8.
//
// Solidity: function l2HeaderChallenges(bytes32 ) view returns((bytes32,uint256) header, (bytes32,uint256) prevHeader, uint256 challengeEnd, address challenger, uint8 status)
func (_Challenge *ChallengeSession) L2HeaderChallenges(arg0 [32]byte) (struct {
	Header       ChallengeL2HeaderL2HeaderPointer
	PrevHeader   ChallengeL2HeaderL2HeaderPointer
	ChallengeEnd *big.Int
	Challenger   common.Address
	Status       uint8
}, error) {
	return _Challenge.Contract.L2HeaderChallenges(&_Challenge.CallOpts, arg0)
}

// L2HeaderChallenges is a free data retrieval call binding the contract method 0x6da802c8.
//
// Solidity: function l2HeaderChallenges(bytes32 ) view returns((bytes32,uint256) header, (bytes32,uint256) prevHeader, uint256 challengeEnd, address challenger, uint8 status)
func (_Challenge *ChallengeCallerSession) L2HeaderChallenges(arg0 [32]byte) (struct {
	Header       ChallengeL2HeaderL2HeaderPointer
	PrevHeader   ChallengeL2HeaderL2HeaderPointer
	ChallengeEnd *big.Int
	Challenger   common.Address
	Status       uint8
}, error) {
	return _Challenge.Contract.L2HeaderChallenges(&_Challenge.CallOpts, arg0)
}

// MipsChallenge is a free data retrieval call binding the contract method 0xdca384be.
//
// Solidity: function mipsChallenge() view returns(address)
func (_Challenge *ChallengeCaller) MipsChallenge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "mipsChallenge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MipsChallenge is a free data retrieval call binding the contract method 0xdca384be.
//
// Solidity: function mipsChallenge() view returns(address)
func (_Challenge *ChallengeSession) MipsChallenge() (common.Address, error) {
	return _Challenge.Contract.MipsChallenge(&_Challenge.CallOpts)
}

// MipsChallenge is a free data retrieval call binding the contract method 0xdca384be.
//
// Solidity: function mipsChallenge() view returns(address)
func (_Challenge *ChallengeCallerSession) MipsChallenge() (common.Address, error) {
	return _Challenge.Contract.MipsChallenge(&_Challenge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Challenge *ChallengeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Challenge *ChallengeSession) Owner() (common.Address, error) {
	return _Challenge.Contract.Owner(&_Challenge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Challenge *ChallengeCallerSession) Owner() (common.Address, error) {
	return _Challenge.Contract.Owner(&_Challenge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Challenge *ChallengeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Challenge *ChallengeSession) ProxiableUUID() ([32]byte, error) {
	return _Challenge.Contract.ProxiableUUID(&_Challenge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Challenge *ChallengeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Challenge.Contract.ProxiableUUID(&_Challenge.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Challenge *ChallengeCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Challenge *ChallengeSession) Treasury() (common.Address, error) {
	return _Challenge.Contract.Treasury(&_Challenge.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Challenge *ChallengeCallerSession) Treasury() (common.Address, error) {
	return _Challenge.Contract.Treasury(&_Challenge.CallOpts)
}

// ChallengeDataRootInclusion is a paid mutator transaction binding the contract method 0x7739f135.
//
// Solidity: function challengeDataRootInclusion(uint256 _blockIndex) payable returns(uint256)
func (_Challenge *ChallengeTransactor) ChallengeDataRootInclusion(opts *bind.TransactOpts, _blockIndex *big.Int) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "challengeDataRootInclusion", _blockIndex)
}

// ChallengeDataRootInclusion is a paid mutator transaction binding the contract method 0x7739f135.
//
// Solidity: function challengeDataRootInclusion(uint256 _blockIndex) payable returns(uint256)
func (_Challenge *ChallengeSession) ChallengeDataRootInclusion(_blockIndex *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeDataRootInclusion(&_Challenge.TransactOpts, _blockIndex)
}

// ChallengeDataRootInclusion is a paid mutator transaction binding the contract method 0x7739f135.
//
// Solidity: function challengeDataRootInclusion(uint256 _blockIndex) payable returns(uint256)
func (_Challenge *ChallengeTransactorSession) ChallengeDataRootInclusion(_blockIndex *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeDataRootInclusion(&_Challenge.TransactOpts, _blockIndex)
}

// ChallengeL2Header is a paid mutator transaction binding the contract method 0x5ae45d8b.
//
// Solidity: function challengeL2Header(uint256 _rblockNum, uint256 _l2Num) payable returns(bytes32)
func (_Challenge *ChallengeTransactor) ChallengeL2Header(opts *bind.TransactOpts, _rblockNum *big.Int, _l2Num *big.Int) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "challengeL2Header", _rblockNum, _l2Num)
}

// ChallengeL2Header is a paid mutator transaction binding the contract method 0x5ae45d8b.
//
// Solidity: function challengeL2Header(uint256 _rblockNum, uint256 _l2Num) payable returns(bytes32)
func (_Challenge *ChallengeSession) ChallengeL2Header(_rblockNum *big.Int, _l2Num *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeL2Header(&_Challenge.TransactOpts, _rblockNum, _l2Num)
}

// ChallengeL2Header is a paid mutator transaction binding the contract method 0x5ae45d8b.
//
// Solidity: function challengeL2Header(uint256 _rblockNum, uint256 _l2Num) payable returns(bytes32)
func (_Challenge *ChallengeTransactorSession) ChallengeL2Header(_rblockNum *big.Int, _l2Num *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeL2Header(&_Challenge.TransactOpts, _rblockNum, _l2Num)
}

// DefendDataRootInclusion is a paid mutator transaction binding the contract method 0x2fc72885.
//
// Solidity: function defendDataRootInclusion(bytes32 _blockHash, (uint256,(uint256,bytes32),(bytes32[],uint256,uint256)) _proof) returns()
func (_Challenge *ChallengeTransactor) DefendDataRootInclusion(opts *bind.TransactOpts, _blockHash [32]byte, _proof ChallengeDataAvailabilityChallengeDAProof) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "defendDataRootInclusion", _blockHash, _proof)
}

// DefendDataRootInclusion is a paid mutator transaction binding the contract method 0x2fc72885.
//
// Solidity: function defendDataRootInclusion(bytes32 _blockHash, (uint256,(uint256,bytes32),(bytes32[],uint256,uint256)) _proof) returns()
func (_Challenge *ChallengeSession) DefendDataRootInclusion(_blockHash [32]byte, _proof ChallengeDataAvailabilityChallengeDAProof) (*types.Transaction, error) {
	return _Challenge.Contract.DefendDataRootInclusion(&_Challenge.TransactOpts, _blockHash, _proof)
}

// DefendDataRootInclusion is a paid mutator transaction binding the contract method 0x2fc72885.
//
// Solidity: function defendDataRootInclusion(bytes32 _blockHash, (uint256,(uint256,bytes32),(bytes32[],uint256,uint256)) _proof) returns()
func (_Challenge *ChallengeTransactorSession) DefendDataRootInclusion(_blockHash [32]byte, _proof ChallengeDataAvailabilityChallengeDAProof) (*types.Transaction, error) {
	return _Challenge.Contract.DefendDataRootInclusion(&_Challenge.TransactOpts, _blockHash, _proof)
}

// DefendL2Header is a paid mutator transaction binding the contract method 0x0200501d.
//
// Solidity: function defendL2Header(bytes32 _challengeHash, bytes32 _headerHash, bytes32 _headerPrevHash) returns()
func (_Challenge *ChallengeTransactor) DefendL2Header(opts *bind.TransactOpts, _challengeHash [32]byte, _headerHash [32]byte, _headerPrevHash [32]byte) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "defendL2Header", _challengeHash, _headerHash, _headerPrevHash)
}

// DefendL2Header is a paid mutator transaction binding the contract method 0x0200501d.
//
// Solidity: function defendL2Header(bytes32 _challengeHash, bytes32 _headerHash, bytes32 _headerPrevHash) returns()
func (_Challenge *ChallengeSession) DefendL2Header(_challengeHash [32]byte, _headerHash [32]byte, _headerPrevHash [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.DefendL2Header(&_Challenge.TransactOpts, _challengeHash, _headerHash, _headerPrevHash)
}

// DefendL2Header is a paid mutator transaction binding the contract method 0x0200501d.
//
// Solidity: function defendL2Header(bytes32 _challengeHash, bytes32 _headerHash, bytes32 _headerPrevHash) returns()
func (_Challenge *ChallengeTransactorSession) DefendL2Header(_challengeHash [32]byte, _headerHash [32]byte, _headerPrevHash [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.DefendL2Header(&_Challenge.TransactOpts, _challengeHash, _headerHash, _headerPrevHash)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _treasury, address _chain, address _daOracle, address _mipsChallenge, address _chainOracle) returns()
func (_Challenge *ChallengeTransactor) Initialize(opts *bind.TransactOpts, _treasury common.Address, _chain common.Address, _daOracle common.Address, _mipsChallenge common.Address, _chainOracle common.Address) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "initialize", _treasury, _chain, _daOracle, _mipsChallenge, _chainOracle)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _treasury, address _chain, address _daOracle, address _mipsChallenge, address _chainOracle) returns()
func (_Challenge *ChallengeSession) Initialize(_treasury common.Address, _chain common.Address, _daOracle common.Address, _mipsChallenge common.Address, _chainOracle common.Address) (*types.Transaction, error) {
	return _Challenge.Contract.Initialize(&_Challenge.TransactOpts, _treasury, _chain, _daOracle, _mipsChallenge, _chainOracle)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _treasury, address _chain, address _daOracle, address _mipsChallenge, address _chainOracle) returns()
func (_Challenge *ChallengeTransactorSession) Initialize(_treasury common.Address, _chain common.Address, _daOracle common.Address, _mipsChallenge common.Address, _chainOracle common.Address) (*types.Transaction, error) {
	return _Challenge.Contract.Initialize(&_Challenge.TransactOpts, _treasury, _chain, _daOracle, _mipsChallenge, _chainOracle)
}

// InvalidateHeader is a paid mutator transaction binding the contract method 0x5dade412.
//
// Solidity: function invalidateHeader(uint256 _blockIndex) returns()
func (_Challenge *ChallengeTransactor) InvalidateHeader(opts *bind.TransactOpts, _blockIndex *big.Int) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "invalidateHeader", _blockIndex)
}

// InvalidateHeader is a paid mutator transaction binding the contract method 0x5dade412.
//
// Solidity: function invalidateHeader(uint256 _blockIndex) returns()
func (_Challenge *ChallengeSession) InvalidateHeader(_blockIndex *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.InvalidateHeader(&_Challenge.TransactOpts, _blockIndex)
}

// InvalidateHeader is a paid mutator transaction binding the contract method 0x5dade412.
//
// Solidity: function invalidateHeader(uint256 _blockIndex) returns()
func (_Challenge *ChallengeTransactorSession) InvalidateHeader(_blockIndex *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.InvalidateHeader(&_Challenge.TransactOpts, _blockIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Challenge *ChallengeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Challenge *ChallengeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Challenge.Contract.RenounceOwnership(&_Challenge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Challenge *ChallengeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Challenge.Contract.RenounceOwnership(&_Challenge.TransactOpts)
}

// SetChallengeFee is a paid mutator transaction binding the contract method 0x35bf82f6.
//
// Solidity: function setChallengeFee(uint256 _challengeFee) returns()
func (_Challenge *ChallengeTransactor) SetChallengeFee(opts *bind.TransactOpts, _challengeFee *big.Int) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "setChallengeFee", _challengeFee)
}

// SetChallengeFee is a paid mutator transaction binding the contract method 0x35bf82f6.
//
// Solidity: function setChallengeFee(uint256 _challengeFee) returns()
func (_Challenge *ChallengeSession) SetChallengeFee(_challengeFee *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.SetChallengeFee(&_Challenge.TransactOpts, _challengeFee)
}

// SetChallengeFee is a paid mutator transaction binding the contract method 0x35bf82f6.
//
// Solidity: function setChallengeFee(uint256 _challengeFee) returns()
func (_Challenge *ChallengeTransactorSession) SetChallengeFee(_challengeFee *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.SetChallengeFee(&_Challenge.TransactOpts, _challengeFee)
}

// SetChallengePeriod is a paid mutator transaction binding the contract method 0x5d475fdd.
//
// Solidity: function setChallengePeriod(uint256 _challengePeriod) returns()
func (_Challenge *ChallengeTransactor) SetChallengePeriod(opts *bind.TransactOpts, _challengePeriod *big.Int) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "setChallengePeriod", _challengePeriod)
}

// SetChallengePeriod is a paid mutator transaction binding the contract method 0x5d475fdd.
//
// Solidity: function setChallengePeriod(uint256 _challengePeriod) returns()
func (_Challenge *ChallengeSession) SetChallengePeriod(_challengePeriod *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.SetChallengePeriod(&_Challenge.TransactOpts, _challengePeriod)
}

// SetChallengePeriod is a paid mutator transaction binding the contract method 0x5d475fdd.
//
// Solidity: function setChallengePeriod(uint256 _challengePeriod) returns()
func (_Challenge *ChallengeTransactorSession) SetChallengePeriod(_challengePeriod *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.SetChallengePeriod(&_Challenge.TransactOpts, _challengePeriod)
}

// SetChallengeReward is a paid mutator transaction binding the contract method 0x7d3020ad.
//
// Solidity: function setChallengeReward(uint256 _challengeReward) returns()
func (_Challenge *ChallengeTransactor) SetChallengeReward(opts *bind.TransactOpts, _challengeReward *big.Int) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "setChallengeReward", _challengeReward)
}

// SetChallengeReward is a paid mutator transaction binding the contract method 0x7d3020ad.
//
// Solidity: function setChallengeReward(uint256 _challengeReward) returns()
func (_Challenge *ChallengeSession) SetChallengeReward(_challengeReward *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.SetChallengeReward(&_Challenge.TransactOpts, _challengeReward)
}

// SetChallengeReward is a paid mutator transaction binding the contract method 0x7d3020ad.
//
// Solidity: function setChallengeReward(uint256 _challengeReward) returns()
func (_Challenge *ChallengeTransactorSession) SetChallengeReward(_challengeReward *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.SetChallengeReward(&_Challenge.TransactOpts, _challengeReward)
}

// SetChallengeWindow is a paid mutator transaction binding the contract method 0x01c1aa0d.
//
// Solidity: function setChallengeWindow(uint256 _challengeWindow) returns()
func (_Challenge *ChallengeTransactor) SetChallengeWindow(opts *bind.TransactOpts, _challengeWindow *big.Int) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "setChallengeWindow", _challengeWindow)
}

// SetChallengeWindow is a paid mutator transaction binding the contract method 0x01c1aa0d.
//
// Solidity: function setChallengeWindow(uint256 _challengeWindow) returns()
func (_Challenge *ChallengeSession) SetChallengeWindow(_challengeWindow *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.SetChallengeWindow(&_Challenge.TransactOpts, _challengeWindow)
}

// SetChallengeWindow is a paid mutator transaction binding the contract method 0x01c1aa0d.
//
// Solidity: function setChallengeWindow(uint256 _challengeWindow) returns()
func (_Challenge *ChallengeTransactorSession) SetChallengeWindow(_challengeWindow *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.SetChallengeWindow(&_Challenge.TransactOpts, _challengeWindow)
}

// SetDefender is a paid mutator transaction binding the contract method 0x163a7177.
//
// Solidity: function setDefender(address _defender) returns()
func (_Challenge *ChallengeTransactor) SetDefender(opts *bind.TransactOpts, _defender common.Address) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "setDefender", _defender)
}

// SetDefender is a paid mutator transaction binding the contract method 0x163a7177.
//
// Solidity: function setDefender(address _defender) returns()
func (_Challenge *ChallengeSession) SetDefender(_defender common.Address) (*types.Transaction, error) {
	return _Challenge.Contract.SetDefender(&_Challenge.TransactOpts, _defender)
}

// SetDefender is a paid mutator transaction binding the contract method 0x163a7177.
//
// Solidity: function setDefender(address _defender) returns()
func (_Challenge *ChallengeTransactorSession) SetDefender(_defender common.Address) (*types.Transaction, error) {
	return _Challenge.Contract.SetDefender(&_Challenge.TransactOpts, _defender)
}

// SettleDataRootInclusion is a paid mutator transaction binding the contract method 0x5bba0ea9.
//
// Solidity: function settleDataRootInclusion(bytes32 _blockhash) returns()
func (_Challenge *ChallengeTransactor) SettleDataRootInclusion(opts *bind.TransactOpts, _blockhash [32]byte) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "settleDataRootInclusion", _blockhash)
}

// SettleDataRootInclusion is a paid mutator transaction binding the contract method 0x5bba0ea9.
//
// Solidity: function settleDataRootInclusion(bytes32 _blockhash) returns()
func (_Challenge *ChallengeSession) SettleDataRootInclusion(_blockhash [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.SettleDataRootInclusion(&_Challenge.TransactOpts, _blockhash)
}

// SettleDataRootInclusion is a paid mutator transaction binding the contract method 0x5bba0ea9.
//
// Solidity: function settleDataRootInclusion(bytes32 _blockhash) returns()
func (_Challenge *ChallengeTransactorSession) SettleDataRootInclusion(_blockhash [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.SettleDataRootInclusion(&_Challenge.TransactOpts, _blockhash)
}

// SettleL2HeaderChallenge is a paid mutator transaction binding the contract method 0xf8a22c6c.
//
// Solidity: function settleL2HeaderChallenge(bytes32 _challengeHash) returns()
func (_Challenge *ChallengeTransactor) SettleL2HeaderChallenge(opts *bind.TransactOpts, _challengeHash [32]byte) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "settleL2HeaderChallenge", _challengeHash)
}

// SettleL2HeaderChallenge is a paid mutator transaction binding the contract method 0xf8a22c6c.
//
// Solidity: function settleL2HeaderChallenge(bytes32 _challengeHash) returns()
func (_Challenge *ChallengeSession) SettleL2HeaderChallenge(_challengeHash [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.SettleL2HeaderChallenge(&_Challenge.TransactOpts, _challengeHash)
}

// SettleL2HeaderChallenge is a paid mutator transaction binding the contract method 0xf8a22c6c.
//
// Solidity: function settleL2HeaderChallenge(bytes32 _challengeHash) returns()
func (_Challenge *ChallengeTransactorSession) SettleL2HeaderChallenge(_challengeHash [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.SettleL2HeaderChallenge(&_Challenge.TransactOpts, _challengeHash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Challenge *ChallengeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Challenge *ChallengeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Challenge.Contract.TransferOwnership(&_Challenge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Challenge *ChallengeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Challenge.Contract.TransferOwnership(&_Challenge.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Challenge *ChallengeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Challenge *ChallengeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Challenge.Contract.UpgradeToAndCall(&_Challenge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Challenge *ChallengeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Challenge.Contract.UpgradeToAndCall(&_Challenge.TransactOpts, newImplementation, data)
}

// ChallengeChallengeDAUpdateIterator is returned from FilterChallengeDAUpdate and is used to iterate over the raw logs and unpacked data for ChallengeDAUpdate events raised by the Challenge contract.
type ChallengeChallengeDAUpdateIterator struct {
	Event *ChallengeChallengeDAUpdate // Event containing the contract specifics and raw log

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
func (it *ChallengeChallengeDAUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeChallengeDAUpdate)
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
		it.Event = new(ChallengeChallengeDAUpdate)
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
func (it *ChallengeChallengeDAUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeChallengeDAUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeChallengeDAUpdate represents a ChallengeDAUpdate event raised by the Challenge contract.
type ChallengeChallengeDAUpdate struct {
	BlockHash  [32]byte
	BlockIndex *big.Int
	Expiry     *big.Int
	Status     uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallengeDAUpdate is a free log retrieval operation binding the contract event 0xcdbf857f75b46d9b94fe7c990b2b0457f168aef7b089a0cb5d4e88074f2d88e8.
//
// Solidity: event ChallengeDAUpdate(bytes32 indexed _blockHash, uint256 indexed _blockIndex, uint256 _expiry, uint8 indexed _status)
func (_Challenge *ChallengeFilterer) FilterChallengeDAUpdate(opts *bind.FilterOpts, _blockHash [][32]byte, _blockIndex []*big.Int, _status []uint8) (*ChallengeChallengeDAUpdateIterator, error) {

	var _blockHashRule []interface{}
	for _, _blockHashItem := range _blockHash {
		_blockHashRule = append(_blockHashRule, _blockHashItem)
	}
	var _blockIndexRule []interface{}
	for _, _blockIndexItem := range _blockIndex {
		_blockIndexRule = append(_blockIndexRule, _blockIndexItem)
	}

	var _statusRule []interface{}
	for _, _statusItem := range _status {
		_statusRule = append(_statusRule, _statusItem)
	}

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "ChallengeDAUpdate", _blockHashRule, _blockIndexRule, _statusRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeChallengeDAUpdateIterator{contract: _Challenge.contract, event: "ChallengeDAUpdate", logs: logs, sub: sub}, nil
}

// WatchChallengeDAUpdate is a free log subscription operation binding the contract event 0xcdbf857f75b46d9b94fe7c990b2b0457f168aef7b089a0cb5d4e88074f2d88e8.
//
// Solidity: event ChallengeDAUpdate(bytes32 indexed _blockHash, uint256 indexed _blockIndex, uint256 _expiry, uint8 indexed _status)
func (_Challenge *ChallengeFilterer) WatchChallengeDAUpdate(opts *bind.WatchOpts, sink chan<- *ChallengeChallengeDAUpdate, _blockHash [][32]byte, _blockIndex []*big.Int, _status []uint8) (event.Subscription, error) {

	var _blockHashRule []interface{}
	for _, _blockHashItem := range _blockHash {
		_blockHashRule = append(_blockHashRule, _blockHashItem)
	}
	var _blockIndexRule []interface{}
	for _, _blockIndexItem := range _blockIndex {
		_blockIndexRule = append(_blockIndexRule, _blockIndexItem)
	}

	var _statusRule []interface{}
	for _, _statusItem := range _status {
		_statusRule = append(_statusRule, _statusItem)
	}

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "ChallengeDAUpdate", _blockHashRule, _blockIndexRule, _statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeChallengeDAUpdate)
				if err := _Challenge.contract.UnpackLog(event, "ChallengeDAUpdate", log); err != nil {
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

// ParseChallengeDAUpdate is a log parse operation binding the contract event 0xcdbf857f75b46d9b94fe7c990b2b0457f168aef7b089a0cb5d4e88074f2d88e8.
//
// Solidity: event ChallengeDAUpdate(bytes32 indexed _blockHash, uint256 indexed _blockIndex, uint256 _expiry, uint8 indexed _status)
func (_Challenge *ChallengeFilterer) ParseChallengeDAUpdate(log types.Log) (*ChallengeChallengeDAUpdate, error) {
	event := new(ChallengeChallengeDAUpdate)
	if err := _Challenge.contract.UnpackLog(event, "ChallengeDAUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Challenge contract.
type ChallengeInitializedIterator struct {
	Event *ChallengeInitialized // Event containing the contract specifics and raw log

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
func (it *ChallengeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeInitialized)
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
		it.Event = new(ChallengeInitialized)
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
func (it *ChallengeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeInitialized represents a Initialized event raised by the Challenge contract.
type ChallengeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Challenge *ChallengeFilterer) FilterInitialized(opts *bind.FilterOpts) (*ChallengeInitializedIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ChallengeInitializedIterator{contract: _Challenge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Challenge *ChallengeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ChallengeInitialized) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeInitialized)
				if err := _Challenge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Challenge *ChallengeFilterer) ParseInitialized(log types.Log) (*ChallengeInitialized, error) {
	event := new(ChallengeInitialized)
	if err := _Challenge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeInvalidHeaderIterator is returned from FilterInvalidHeader and is used to iterate over the raw logs and unpacked data for InvalidHeader events raised by the Challenge contract.
type ChallengeInvalidHeaderIterator struct {
	Event *ChallengeInvalidHeader // Event containing the contract specifics and raw log

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
func (it *ChallengeInvalidHeaderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeInvalidHeader)
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
		it.Event = new(ChallengeInvalidHeader)
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
func (it *ChallengeInvalidHeaderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeInvalidHeaderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeInvalidHeader represents a InvalidHeader event raised by the Challenge contract.
type ChallengeInvalidHeader struct {
	BlockIndex *big.Int
	Hash       [32]byte
	Reason     uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvalidHeader is a free log retrieval operation binding the contract event 0xea46f8ad2711844c28d6aa0fe8ed10b1ac38bdcdc6df7ba3b8f3bfc35232f31b.
//
// Solidity: event InvalidHeader(uint256 indexed _blockIndex, bytes32 indexed _hash, uint8 indexed _reason)
func (_Challenge *ChallengeFilterer) FilterInvalidHeader(opts *bind.FilterOpts, _blockIndex []*big.Int, _hash [][32]byte, _reason []uint8) (*ChallengeInvalidHeaderIterator, error) {

	var _blockIndexRule []interface{}
	for _, _blockIndexItem := range _blockIndex {
		_blockIndexRule = append(_blockIndexRule, _blockIndexItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _reasonRule []interface{}
	for _, _reasonItem := range _reason {
		_reasonRule = append(_reasonRule, _reasonItem)
	}

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "InvalidHeader", _blockIndexRule, _hashRule, _reasonRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeInvalidHeaderIterator{contract: _Challenge.contract, event: "InvalidHeader", logs: logs, sub: sub}, nil
}

// WatchInvalidHeader is a free log subscription operation binding the contract event 0xea46f8ad2711844c28d6aa0fe8ed10b1ac38bdcdc6df7ba3b8f3bfc35232f31b.
//
// Solidity: event InvalidHeader(uint256 indexed _blockIndex, bytes32 indexed _hash, uint8 indexed _reason)
func (_Challenge *ChallengeFilterer) WatchInvalidHeader(opts *bind.WatchOpts, sink chan<- *ChallengeInvalidHeader, _blockIndex []*big.Int, _hash [][32]byte, _reason []uint8) (event.Subscription, error) {

	var _blockIndexRule []interface{}
	for _, _blockIndexItem := range _blockIndex {
		_blockIndexRule = append(_blockIndexRule, _blockIndexItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _reasonRule []interface{}
	for _, _reasonItem := range _reason {
		_reasonRule = append(_reasonRule, _reasonItem)
	}

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "InvalidHeader", _blockIndexRule, _hashRule, _reasonRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeInvalidHeader)
				if err := _Challenge.contract.UnpackLog(event, "InvalidHeader", log); err != nil {
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

// ParseInvalidHeader is a log parse operation binding the contract event 0xea46f8ad2711844c28d6aa0fe8ed10b1ac38bdcdc6df7ba3b8f3bfc35232f31b.
//
// Solidity: event InvalidHeader(uint256 indexed _blockIndex, bytes32 indexed _hash, uint8 indexed _reason)
func (_Challenge *ChallengeFilterer) ParseInvalidHeader(log types.Log) (*ChallengeInvalidHeader, error) {
	event := new(ChallengeInvalidHeader)
	if err := _Challenge.contract.UnpackLog(event, "InvalidHeader", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeL2HeaderChallengeUpdateIterator is returned from FilterL2HeaderChallengeUpdate and is used to iterate over the raw logs and unpacked data for L2HeaderChallengeUpdate events raised by the Challenge contract.
type ChallengeL2HeaderChallengeUpdateIterator struct {
	Event *ChallengeL2HeaderChallengeUpdate // Event containing the contract specifics and raw log

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
func (it *ChallengeL2HeaderChallengeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeL2HeaderChallengeUpdate)
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
		it.Event = new(ChallengeL2HeaderChallengeUpdate)
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
func (it *ChallengeL2HeaderChallengeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeL2HeaderChallengeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeL2HeaderChallengeUpdate represents a L2HeaderChallengeUpdate event raised by the Challenge contract.
type ChallengeL2HeaderChallengeUpdate struct {
	ChallengeHash [32]byte
	L2Number      *big.Int
	Rblock        [32]byte
	Expiry        *big.Int
	Status        uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterL2HeaderChallengeUpdate is a free log retrieval operation binding the contract event 0x89389219f04af163105506312f99d6ca034af96d3ee1128efc6a9619ee2aeec2.
//
// Solidity: event L2HeaderChallengeUpdate(bytes32 indexed challengeHash, uint256 indexed l2Number, bytes32 rblock, uint256 expiry, uint8 indexed status)
func (_Challenge *ChallengeFilterer) FilterL2HeaderChallengeUpdate(opts *bind.FilterOpts, challengeHash [][32]byte, l2Number []*big.Int, status []uint8) (*ChallengeL2HeaderChallengeUpdateIterator, error) {

	var challengeHashRule []interface{}
	for _, challengeHashItem := range challengeHash {
		challengeHashRule = append(challengeHashRule, challengeHashItem)
	}
	var l2NumberRule []interface{}
	for _, l2NumberItem := range l2Number {
		l2NumberRule = append(l2NumberRule, l2NumberItem)
	}

	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "L2HeaderChallengeUpdate", challengeHashRule, l2NumberRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeL2HeaderChallengeUpdateIterator{contract: _Challenge.contract, event: "L2HeaderChallengeUpdate", logs: logs, sub: sub}, nil
}

// WatchL2HeaderChallengeUpdate is a free log subscription operation binding the contract event 0x89389219f04af163105506312f99d6ca034af96d3ee1128efc6a9619ee2aeec2.
//
// Solidity: event L2HeaderChallengeUpdate(bytes32 indexed challengeHash, uint256 indexed l2Number, bytes32 rblock, uint256 expiry, uint8 indexed status)
func (_Challenge *ChallengeFilterer) WatchL2HeaderChallengeUpdate(opts *bind.WatchOpts, sink chan<- *ChallengeL2HeaderChallengeUpdate, challengeHash [][32]byte, l2Number []*big.Int, status []uint8) (event.Subscription, error) {

	var challengeHashRule []interface{}
	for _, challengeHashItem := range challengeHash {
		challengeHashRule = append(challengeHashRule, challengeHashItem)
	}
	var l2NumberRule []interface{}
	for _, l2NumberItem := range l2Number {
		l2NumberRule = append(l2NumberRule, l2NumberItem)
	}

	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "L2HeaderChallengeUpdate", challengeHashRule, l2NumberRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeL2HeaderChallengeUpdate)
				if err := _Challenge.contract.UnpackLog(event, "L2HeaderChallengeUpdate", log); err != nil {
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

// ParseL2HeaderChallengeUpdate is a log parse operation binding the contract event 0x89389219f04af163105506312f99d6ca034af96d3ee1128efc6a9619ee2aeec2.
//
// Solidity: event L2HeaderChallengeUpdate(bytes32 indexed challengeHash, uint256 indexed l2Number, bytes32 rblock, uint256 expiry, uint8 indexed status)
func (_Challenge *ChallengeFilterer) ParseL2HeaderChallengeUpdate(log types.Log) (*ChallengeL2HeaderChallengeUpdate, error) {
	event := new(ChallengeL2HeaderChallengeUpdate)
	if err := _Challenge.contract.UnpackLog(event, "L2HeaderChallengeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Challenge contract.
type ChallengeOwnershipTransferredIterator struct {
	Event *ChallengeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChallengeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeOwnershipTransferred)
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
		it.Event = new(ChallengeOwnershipTransferred)
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
func (it *ChallengeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeOwnershipTransferred represents a OwnershipTransferred event raised by the Challenge contract.
type ChallengeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Challenge *ChallengeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChallengeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeOwnershipTransferredIterator{contract: _Challenge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Challenge *ChallengeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChallengeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeOwnershipTransferred)
				if err := _Challenge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Challenge *ChallengeFilterer) ParseOwnershipTransferred(log types.Log) (*ChallengeOwnershipTransferred, error) {
	event := new(ChallengeOwnershipTransferred)
	if err := _Challenge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Challenge contract.
type ChallengeUpgradedIterator struct {
	Event *ChallengeUpgraded // Event containing the contract specifics and raw log

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
func (it *ChallengeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeUpgraded)
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
		it.Event = new(ChallengeUpgraded)
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
func (it *ChallengeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeUpgraded represents a Upgraded event raised by the Challenge contract.
type ChallengeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Challenge *ChallengeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ChallengeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeUpgradedIterator{contract: _Challenge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Challenge *ChallengeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ChallengeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeUpgraded)
				if err := _Challenge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Challenge *ChallengeFilterer) ParseUpgraded(log types.Log) (*ChallengeUpgraded, error) {
	event := new(ChallengeUpgraded)
	if err := _Challenge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
