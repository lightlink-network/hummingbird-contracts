// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainoracle

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

// AttestationProof is an auto generated low-level Go binding around an user-defined struct.
type AttestationProof struct {
	TupleRootNonce *big.Int
	Tuple          DataRootTuple
	Proof          BinaryMerkleProof
}

// BinaryMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type BinaryMerkleProof struct {
	SideNodes [][32]byte
	Key       *big.Int
	NumLeaves *big.Int
}

// ChainOracleL2Header is an auto generated low-level Go binding around an user-defined struct.
type ChainOracleL2Header struct {
	ParentHash       [32]byte
	UncleHash        [32]byte
	Beneficiary      common.Address
	StateRoot        [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	LogsBloom        [32]byte
	Difficulty       *big.Int
	Number           *big.Int
	GasLimit         *big.Int
	GasUsed          *big.Int
	Timestamp        *big.Int
	ExtraData        [32]byte
	MixHash          [32]byte
	Nonce            *big.Int
}

// ChainOracleShareRange is an auto generated low-level Go binding around an user-defined struct.
type ChainOracleShareRange struct {
	Start *big.Int
	End   *big.Int
}

// DataRootTuple is an auto generated low-level Go binding around an user-defined struct.
type DataRootTuple struct {
	Height   *big.Int
	DataRoot [32]byte
}

// Namespace is an auto generated low-level Go binding around an user-defined struct.
type Namespace struct {
	Version [1]byte
	Id      [28]byte
}

// NamespaceMerkleMultiproof is an auto generated low-level Go binding around an user-defined struct.
type NamespaceMerkleMultiproof struct {
	BeginKey  *big.Int
	EndKey    *big.Int
	SideNodes []NamespaceNode
}

// NamespaceNode is an auto generated low-level Go binding around an user-defined struct.
type NamespaceNode struct {
	Min    Namespace
	Max    Namespace
	Digest [32]byte
}

// SharesProof is an auto generated low-level Go binding around an user-defined struct.
type SharesProof struct {
	Data             [][]byte
	ShareProofs      []NamespaceMerkleMultiproof
	Namespace        Namespace
	RowRoots         []NamespaceNode
	RowProofs        []BinaryMerkleProof
	AttestationProof AttestationProof
}

// ChainOracleMetaData contains all meta data concerning the ChainOracle contract.
var ChainOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_canonicalStateChain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rlpReader\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_rblock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"_shareData\",\"type\":\"bytes[]\"}],\"name\":\"ShareKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canonicalStateChain\",\"outputs\":[{\"internalType\":\"contractICanonicalStateChain\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daOracle\",\"outputs\":[{\"internalType\":\"contractIDAOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"decodeRLPHeader\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"uncleHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"logsBloom\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"extraData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structChainOracle.L2Header\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_shareData\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structChainOracle.ShareRange[]\",\"name\":\"_shareRanges\",\"type\":\"tuple[]\"}],\"name\":\"extractData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"uncleHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"logsBloom\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"extraData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structChainOracle.L2Header\",\"name\":\"_header\",\"type\":\"tuple\"}],\"name\":\"hashHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"headers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"uncleHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"logsBloom\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"extraData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_shareKey\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structChainOracle.ShareRange[]\",\"name\":\"_range\",\"type\":\"tuple[]\"}],\"name\":\"provideHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_rblock\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"namespace\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"provideShares\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rlpReader\",\"outputs\":[{\"internalType\":\"contractIRLPReader\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"shares\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ChainOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use ChainOracleMetaData.ABI instead.
var ChainOracleABI = ChainOracleMetaData.ABI

// ChainOracle is an auto generated Go binding around an Ethereum contract.
type ChainOracle struct {
	ChainOracleCaller     // Read-only binding to the contract
	ChainOracleTransactor // Write-only binding to the contract
	ChainOracleFilterer   // Log filterer for contract events
}

// ChainOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainOracleSession struct {
	Contract     *ChainOracle      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChainOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainOracleCallerSession struct {
	Contract *ChainOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ChainOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainOracleTransactorSession struct {
	Contract     *ChainOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ChainOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainOracleRaw struct {
	Contract *ChainOracle // Generic contract binding to access the raw methods on
}

// ChainOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainOracleCallerRaw struct {
	Contract *ChainOracleCaller // Generic read-only contract binding to access the raw methods on
}

// ChainOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainOracleTransactorRaw struct {
	Contract *ChainOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainOracle creates a new instance of ChainOracle, bound to a specific deployed contract.
func NewChainOracle(address common.Address, backend bind.ContractBackend) (*ChainOracle, error) {
	contract, err := bindChainOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainOracle{ChainOracleCaller: ChainOracleCaller{contract: contract}, ChainOracleTransactor: ChainOracleTransactor{contract: contract}, ChainOracleFilterer: ChainOracleFilterer{contract: contract}}, nil
}

// NewChainOracleCaller creates a new read-only instance of ChainOracle, bound to a specific deployed contract.
func NewChainOracleCaller(address common.Address, caller bind.ContractCaller) (*ChainOracleCaller, error) {
	contract, err := bindChainOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainOracleCaller{contract: contract}, nil
}

// NewChainOracleTransactor creates a new write-only instance of ChainOracle, bound to a specific deployed contract.
func NewChainOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainOracleTransactor, error) {
	contract, err := bindChainOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainOracleTransactor{contract: contract}, nil
}

// NewChainOracleFilterer creates a new log filterer instance of ChainOracle, bound to a specific deployed contract.
func NewChainOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainOracleFilterer, error) {
	contract, err := bindChainOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainOracleFilterer{contract: contract}, nil
}

// bindChainOracle binds a generic wrapper to an already deployed contract.
func bindChainOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChainOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainOracle *ChainOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainOracle.Contract.ChainOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainOracle *ChainOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainOracle.Contract.ChainOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainOracle *ChainOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainOracle.Contract.ChainOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainOracle *ChainOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainOracle *ChainOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainOracle *ChainOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainOracle.Contract.contract.Transact(opts, method, params...)
}

// ShareKey is a free data retrieval call binding the contract method 0x40a2ab6e.
//
// Solidity: function ShareKey(bytes32 _rblock, bytes[] _shareData) pure returns(bytes32)
func (_ChainOracle *ChainOracleCaller) ShareKey(opts *bind.CallOpts, _rblock [32]byte, _shareData [][]byte) ([32]byte, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "ShareKey", _rblock, _shareData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ShareKey is a free data retrieval call binding the contract method 0x40a2ab6e.
//
// Solidity: function ShareKey(bytes32 _rblock, bytes[] _shareData) pure returns(bytes32)
func (_ChainOracle *ChainOracleSession) ShareKey(_rblock [32]byte, _shareData [][]byte) ([32]byte, error) {
	return _ChainOracle.Contract.ShareKey(&_ChainOracle.CallOpts, _rblock, _shareData)
}

// ShareKey is a free data retrieval call binding the contract method 0x40a2ab6e.
//
// Solidity: function ShareKey(bytes32 _rblock, bytes[] _shareData) pure returns(bytes32)
func (_ChainOracle *ChainOracleCallerSession) ShareKey(_rblock [32]byte, _shareData [][]byte) ([32]byte, error) {
	return _ChainOracle.Contract.ShareKey(&_ChainOracle.CallOpts, _rblock, _shareData)
}

// CanonicalStateChain is a free data retrieval call binding the contract method 0x8c69fa5d.
//
// Solidity: function canonicalStateChain() view returns(address)
func (_ChainOracle *ChainOracleCaller) CanonicalStateChain(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "canonicalStateChain")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CanonicalStateChain is a free data retrieval call binding the contract method 0x8c69fa5d.
//
// Solidity: function canonicalStateChain() view returns(address)
func (_ChainOracle *ChainOracleSession) CanonicalStateChain() (common.Address, error) {
	return _ChainOracle.Contract.CanonicalStateChain(&_ChainOracle.CallOpts)
}

// CanonicalStateChain is a free data retrieval call binding the contract method 0x8c69fa5d.
//
// Solidity: function canonicalStateChain() view returns(address)
func (_ChainOracle *ChainOracleCallerSession) CanonicalStateChain() (common.Address, error) {
	return _ChainOracle.Contract.CanonicalStateChain(&_ChainOracle.CallOpts)
}

// DaOracle is a free data retrieval call binding the contract method 0xee223c02.
//
// Solidity: function daOracle() view returns(address)
func (_ChainOracle *ChainOracleCaller) DaOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "daOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaOracle is a free data retrieval call binding the contract method 0xee223c02.
//
// Solidity: function daOracle() view returns(address)
func (_ChainOracle *ChainOracleSession) DaOracle() (common.Address, error) {
	return _ChainOracle.Contract.DaOracle(&_ChainOracle.CallOpts)
}

// DaOracle is a free data retrieval call binding the contract method 0xee223c02.
//
// Solidity: function daOracle() view returns(address)
func (_ChainOracle *ChainOracleCallerSession) DaOracle() (common.Address, error) {
	return _ChainOracle.Contract.DaOracle(&_ChainOracle.CallOpts)
}

// DecodeRLPHeader is a free data retrieval call binding the contract method 0x26410af6.
//
// Solidity: function decodeRLPHeader(bytes _data) view returns((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,uint256))
func (_ChainOracle *ChainOracleCaller) DecodeRLPHeader(opts *bind.CallOpts, _data []byte) (ChainOracleL2Header, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "decodeRLPHeader", _data)

	if err != nil {
		return *new(ChainOracleL2Header), err
	}

	out0 := *abi.ConvertType(out[0], new(ChainOracleL2Header)).(*ChainOracleL2Header)

	return out0, err

}

// DecodeRLPHeader is a free data retrieval call binding the contract method 0x26410af6.
//
// Solidity: function decodeRLPHeader(bytes _data) view returns((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,uint256))
func (_ChainOracle *ChainOracleSession) DecodeRLPHeader(_data []byte) (ChainOracleL2Header, error) {
	return _ChainOracle.Contract.DecodeRLPHeader(&_ChainOracle.CallOpts, _data)
}

// DecodeRLPHeader is a free data retrieval call binding the contract method 0x26410af6.
//
// Solidity: function decodeRLPHeader(bytes _data) view returns((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,uint256))
func (_ChainOracle *ChainOracleCallerSession) DecodeRLPHeader(_data []byte) (ChainOracleL2Header, error) {
	return _ChainOracle.Contract.DecodeRLPHeader(&_ChainOracle.CallOpts, _data)
}

// ExtractData is a free data retrieval call binding the contract method 0xacd23ff9.
//
// Solidity: function extractData(bytes[] _shareData, (uint256,uint256)[] _shareRanges) pure returns(bytes)
func (_ChainOracle *ChainOracleCaller) ExtractData(opts *bind.CallOpts, _shareData [][]byte, _shareRanges []ChainOracleShareRange) ([]byte, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "extractData", _shareData, _shareRanges)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ExtractData is a free data retrieval call binding the contract method 0xacd23ff9.
//
// Solidity: function extractData(bytes[] _shareData, (uint256,uint256)[] _shareRanges) pure returns(bytes)
func (_ChainOracle *ChainOracleSession) ExtractData(_shareData [][]byte, _shareRanges []ChainOracleShareRange) ([]byte, error) {
	return _ChainOracle.Contract.ExtractData(&_ChainOracle.CallOpts, _shareData, _shareRanges)
}

// ExtractData is a free data retrieval call binding the contract method 0xacd23ff9.
//
// Solidity: function extractData(bytes[] _shareData, (uint256,uint256)[] _shareRanges) pure returns(bytes)
func (_ChainOracle *ChainOracleCallerSession) ExtractData(_shareData [][]byte, _shareRanges []ChainOracleShareRange) ([]byte, error) {
	return _ChainOracle.Contract.ExtractData(&_ChainOracle.CallOpts, _shareData, _shareRanges)
}

// HashHeader is a free data retrieval call binding the contract method 0xf391deb2.
//
// Solidity: function hashHeader((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,uint256) _header) pure returns(bytes32)
func (_ChainOracle *ChainOracleCaller) HashHeader(opts *bind.CallOpts, _header ChainOracleL2Header) ([32]byte, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "hashHeader", _header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashHeader is a free data retrieval call binding the contract method 0xf391deb2.
//
// Solidity: function hashHeader((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,uint256) _header) pure returns(bytes32)
func (_ChainOracle *ChainOracleSession) HashHeader(_header ChainOracleL2Header) ([32]byte, error) {
	return _ChainOracle.Contract.HashHeader(&_ChainOracle.CallOpts, _header)
}

// HashHeader is a free data retrieval call binding the contract method 0xf391deb2.
//
// Solidity: function hashHeader((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,uint256) _header) pure returns(bytes32)
func (_ChainOracle *ChainOracleCallerSession) HashHeader(_header ChainOracleL2Header) ([32]byte, error) {
	return _ChainOracle.Contract.HashHeader(&_ChainOracle.CallOpts, _header)
}

// Headers is a free data retrieval call binding the contract method 0x9e7f2700.
//
// Solidity: function headers(bytes32 ) view returns(bytes32 parentHash, bytes32 uncleHash, address beneficiary, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, bytes32 logsBloom, uint256 difficulty, uint256 number, uint256 gasLimit, uint256 gasUsed, uint256 timestamp, bytes32 extraData, bytes32 mixHash, uint256 nonce)
func (_ChainOracle *ChainOracleCaller) Headers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ParentHash       [32]byte
	UncleHash        [32]byte
	Beneficiary      common.Address
	StateRoot        [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	LogsBloom        [32]byte
	Difficulty       *big.Int
	Number           *big.Int
	GasLimit         *big.Int
	GasUsed          *big.Int
	Timestamp        *big.Int
	ExtraData        [32]byte
	MixHash          [32]byte
	Nonce            *big.Int
}, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "headers", arg0)

	outstruct := new(struct {
		ParentHash       [32]byte
		UncleHash        [32]byte
		Beneficiary      common.Address
		StateRoot        [32]byte
		TransactionsRoot [32]byte
		ReceiptsRoot     [32]byte
		LogsBloom        [32]byte
		Difficulty       *big.Int
		Number           *big.Int
		GasLimit         *big.Int
		GasUsed          *big.Int
		Timestamp        *big.Int
		ExtraData        [32]byte
		MixHash          [32]byte
		Nonce            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ParentHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.UncleHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Beneficiary = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.StateRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.TransactionsRoot = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.ReceiptsRoot = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.LogsBloom = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	outstruct.Difficulty = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Number = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.GasLimit = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.GasUsed = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.ExtraData = *abi.ConvertType(out[12], new([32]byte)).(*[32]byte)
	outstruct.MixHash = *abi.ConvertType(out[13], new([32]byte)).(*[32]byte)
	outstruct.Nonce = *abi.ConvertType(out[14], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Headers is a free data retrieval call binding the contract method 0x9e7f2700.
//
// Solidity: function headers(bytes32 ) view returns(bytes32 parentHash, bytes32 uncleHash, address beneficiary, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, bytes32 logsBloom, uint256 difficulty, uint256 number, uint256 gasLimit, uint256 gasUsed, uint256 timestamp, bytes32 extraData, bytes32 mixHash, uint256 nonce)
func (_ChainOracle *ChainOracleSession) Headers(arg0 [32]byte) (struct {
	ParentHash       [32]byte
	UncleHash        [32]byte
	Beneficiary      common.Address
	StateRoot        [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	LogsBloom        [32]byte
	Difficulty       *big.Int
	Number           *big.Int
	GasLimit         *big.Int
	GasUsed          *big.Int
	Timestamp        *big.Int
	ExtraData        [32]byte
	MixHash          [32]byte
	Nonce            *big.Int
}, error) {
	return _ChainOracle.Contract.Headers(&_ChainOracle.CallOpts, arg0)
}

// Headers is a free data retrieval call binding the contract method 0x9e7f2700.
//
// Solidity: function headers(bytes32 ) view returns(bytes32 parentHash, bytes32 uncleHash, address beneficiary, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, bytes32 logsBloom, uint256 difficulty, uint256 number, uint256 gasLimit, uint256 gasUsed, uint256 timestamp, bytes32 extraData, bytes32 mixHash, uint256 nonce)
func (_ChainOracle *ChainOracleCallerSession) Headers(arg0 [32]byte) (struct {
	ParentHash       [32]byte
	UncleHash        [32]byte
	Beneficiary      common.Address
	StateRoot        [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	LogsBloom        [32]byte
	Difficulty       *big.Int
	Number           *big.Int
	GasLimit         *big.Int
	GasUsed          *big.Int
	Timestamp        *big.Int
	ExtraData        [32]byte
	MixHash          [32]byte
	Nonce            *big.Int
}, error) {
	return _ChainOracle.Contract.Headers(&_ChainOracle.CallOpts, arg0)
}

// RlpReader is a free data retrieval call binding the contract method 0x0a9b3a32.
//
// Solidity: function rlpReader() view returns(address)
func (_ChainOracle *ChainOracleCaller) RlpReader(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "rlpReader")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RlpReader is a free data retrieval call binding the contract method 0x0a9b3a32.
//
// Solidity: function rlpReader() view returns(address)
func (_ChainOracle *ChainOracleSession) RlpReader() (common.Address, error) {
	return _ChainOracle.Contract.RlpReader(&_ChainOracle.CallOpts)
}

// RlpReader is a free data retrieval call binding the contract method 0x0a9b3a32.
//
// Solidity: function rlpReader() view returns(address)
func (_ChainOracle *ChainOracleCallerSession) RlpReader() (common.Address, error) {
	return _ChainOracle.Contract.RlpReader(&_ChainOracle.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0x263d5f11.
//
// Solidity: function shares(bytes32 , uint256 ) view returns(bytes)
func (_ChainOracle *ChainOracleCaller) Shares(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "shares", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0x263d5f11.
//
// Solidity: function shares(bytes32 , uint256 ) view returns(bytes)
func (_ChainOracle *ChainOracleSession) Shares(arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	return _ChainOracle.Contract.Shares(&_ChainOracle.CallOpts, arg0, arg1)
}

// Shares is a free data retrieval call binding the contract method 0x263d5f11.
//
// Solidity: function shares(bytes32 , uint256 ) view returns(bytes)
func (_ChainOracle *ChainOracleCallerSession) Shares(arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	return _ChainOracle.Contract.Shares(&_ChainOracle.CallOpts, arg0, arg1)
}

// ProvideHeader is a paid mutator transaction binding the contract method 0xd82d0abe.
//
// Solidity: function provideHeader(bytes32 _shareKey, (uint256,uint256)[] _range) returns(bytes32)
func (_ChainOracle *ChainOracleTransactor) ProvideHeader(opts *bind.TransactOpts, _shareKey [32]byte, _range []ChainOracleShareRange) (*types.Transaction, error) {
	return _ChainOracle.contract.Transact(opts, "provideHeader", _shareKey, _range)
}

// ProvideHeader is a paid mutator transaction binding the contract method 0xd82d0abe.
//
// Solidity: function provideHeader(bytes32 _shareKey, (uint256,uint256)[] _range) returns(bytes32)
func (_ChainOracle *ChainOracleSession) ProvideHeader(_shareKey [32]byte, _range []ChainOracleShareRange) (*types.Transaction, error) {
	return _ChainOracle.Contract.ProvideHeader(&_ChainOracle.TransactOpts, _shareKey, _range)
}

// ProvideHeader is a paid mutator transaction binding the contract method 0xd82d0abe.
//
// Solidity: function provideHeader(bytes32 _shareKey, (uint256,uint256)[] _range) returns(bytes32)
func (_ChainOracle *ChainOracleTransactorSession) ProvideHeader(_shareKey [32]byte, _range []ChainOracleShareRange) (*types.Transaction, error) {
	return _ChainOracle.Contract.ProvideHeader(&_ChainOracle.TransactOpts, _shareKey, _range)
}

// ProvideShares is a paid mutator transaction binding the contract method 0x2c97b5b3.
//
// Solidity: function provideShares(bytes32 _rblock, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes32)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _proof) returns(bytes32)
func (_ChainOracle *ChainOracleTransactor) ProvideShares(opts *bind.TransactOpts, _rblock [32]byte, _proof SharesProof) (*types.Transaction, error) {
	return _ChainOracle.contract.Transact(opts, "provideShares", _rblock, _proof)
}

// ProvideShares is a paid mutator transaction binding the contract method 0x2c97b5b3.
//
// Solidity: function provideShares(bytes32 _rblock, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes32)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _proof) returns(bytes32)
func (_ChainOracle *ChainOracleSession) ProvideShares(_rblock [32]byte, _proof SharesProof) (*types.Transaction, error) {
	return _ChainOracle.Contract.ProvideShares(&_ChainOracle.TransactOpts, _rblock, _proof)
}

// ProvideShares is a paid mutator transaction binding the contract method 0x2c97b5b3.
//
// Solidity: function provideShares(bytes32 _rblock, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes32)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _proof) returns(bytes32)
func (_ChainOracle *ChainOracleTransactorSession) ProvideShares(_rblock [32]byte, _proof SharesProof) (*types.Transaction, error) {
	return _ChainOracle.Contract.ProvideShares(&_ChainOracle.TransactOpts, _rblock, _proof)
}
