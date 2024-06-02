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

// ChainOracleDepositTx is an auto generated low-level Go binding around an user-defined struct.
type ChainOracleDepositTx struct {
	ChainId  *big.Int
	Nonce    uint64
	GasPrice *big.Int
	Gas      uint64
	To       common.Address
	Value    *big.Int
	Data     []byte
	R        *big.Int
	S        *big.Int
	V        *big.Int
}

// ChainOracleL2Header is an auto generated low-level Go binding around an user-defined struct.
type ChainOracleL2Header struct {
	ParentHash       [32]byte
	UncleHash        [32]byte
	Beneficiary      common.Address
	StateRoot        [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	LogsBloom        []byte
	Difficulty       *big.Int
	Number           *big.Int
	GasLimit         *big.Int
	GasUsed          *big.Int
	Timestamp        *big.Int
	ExtraData        []byte
	MixHash          [32]byte
	Nonce            *big.Int
}

// ChainOracleLegacyTx is an auto generated low-level Go binding around an user-defined struct.
type ChainOracleLegacyTx struct {
	Nonce    uint64
	GasPrice *big.Int
	Gas      uint64
	To       common.Address
	Value    *big.Int
	Data     []byte
	R        *big.Int
	S        *big.Int
	V        *big.Int
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_rblock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"_shareData\",\"type\":\"bytes[]\"}],\"name\":\"ShareKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canonicalStateChain\",\"outputs\":[{\"internalType\":\"contractICanonicalStateChain\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daOracle\",\"outputs\":[{\"internalType\":\"contractIDAOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"decodeDepositTx\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"internalType\":\"structChainOracle.DepositTx\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"decodeLegacyTx\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"internalType\":\"structChainOracle.LegacyTx\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"decodeRLPHeader\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"uncleHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"logsBloom\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structChainOracle.L2Header\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"raw\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structChainOracle.ShareRange[]\",\"name\":\"ranges\",\"type\":\"tuple[]\"}],\"name\":\"extractData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_headerHash\",\"type\":\"bytes32\"}],\"name\":\"getHeader\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"uncleHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"logsBloom\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structChainOracle.L2Header\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"getTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"internalType\":\"structChainOracle.DepositTx\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"uncleHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"logsBloom\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structChainOracle.L2Header\",\"name\":\"_header\",\"type\":\"tuple\"}],\"name\":\"hashHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"headerToRblock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"headers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"uncleHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"logsBloom\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_canonicalStateChain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rlpReader\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_shareKey\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structChainOracle.ShareRange[]\",\"name\":\"_range\",\"type\":\"tuple[]\"}],\"name\":\"provideHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_shareKey\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structChainOracle.ShareRange[]\",\"name\":\"_range\",\"type\":\"tuple[]\"}],\"name\":\"provideLegacyTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_rblock\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_pointer\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"namespace\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"provideShares\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rlpReader\",\"outputs\":[{\"internalType\":\"contractIRLPReader\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rlpReader\",\"type\":\"address\"}],\"name\":\"setRLPReader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"shares\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ChainOracle *ChainOracleCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ChainOracle *ChainOracleSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ChainOracle.Contract.UPGRADEINTERFACEVERSION(&_ChainOracle.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ChainOracle *ChainOracleCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ChainOracle.Contract.UPGRADEINTERFACEVERSION(&_ChainOracle.CallOpts)
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

// DecodeDepositTx is a free data retrieval call binding the contract method 0x7f2596fe.
//
// Solidity: function decodeDepositTx(bytes _data) view returns((uint256,uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256))
func (_ChainOracle *ChainOracleCaller) DecodeDepositTx(opts *bind.CallOpts, _data []byte) (ChainOracleDepositTx, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "decodeDepositTx", _data)

	if err != nil {
		return *new(ChainOracleDepositTx), err
	}

	out0 := *abi.ConvertType(out[0], new(ChainOracleDepositTx)).(*ChainOracleDepositTx)

	return out0, err

}

// DecodeDepositTx is a free data retrieval call binding the contract method 0x7f2596fe.
//
// Solidity: function decodeDepositTx(bytes _data) view returns((uint256,uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256))
func (_ChainOracle *ChainOracleSession) DecodeDepositTx(_data []byte) (ChainOracleDepositTx, error) {
	return _ChainOracle.Contract.DecodeDepositTx(&_ChainOracle.CallOpts, _data)
}

// DecodeDepositTx is a free data retrieval call binding the contract method 0x7f2596fe.
//
// Solidity: function decodeDepositTx(bytes _data) view returns((uint256,uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256))
func (_ChainOracle *ChainOracleCallerSession) DecodeDepositTx(_data []byte) (ChainOracleDepositTx, error) {
	return _ChainOracle.Contract.DecodeDepositTx(&_ChainOracle.CallOpts, _data)
}

// DecodeLegacyTx is a free data retrieval call binding the contract method 0xbfcbf73e.
//
// Solidity: function decodeLegacyTx(bytes _data) view returns((uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256))
func (_ChainOracle *ChainOracleCaller) DecodeLegacyTx(opts *bind.CallOpts, _data []byte) (ChainOracleLegacyTx, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "decodeLegacyTx", _data)

	if err != nil {
		return *new(ChainOracleLegacyTx), err
	}

	out0 := *abi.ConvertType(out[0], new(ChainOracleLegacyTx)).(*ChainOracleLegacyTx)

	return out0, err

}

// DecodeLegacyTx is a free data retrieval call binding the contract method 0xbfcbf73e.
//
// Solidity: function decodeLegacyTx(bytes _data) view returns((uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256))
func (_ChainOracle *ChainOracleSession) DecodeLegacyTx(_data []byte) (ChainOracleLegacyTx, error) {
	return _ChainOracle.Contract.DecodeLegacyTx(&_ChainOracle.CallOpts, _data)
}

// DecodeLegacyTx is a free data retrieval call binding the contract method 0xbfcbf73e.
//
// Solidity: function decodeLegacyTx(bytes _data) view returns((uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256))
func (_ChainOracle *ChainOracleCallerSession) DecodeLegacyTx(_data []byte) (ChainOracleLegacyTx, error) {
	return _ChainOracle.Contract.DecodeLegacyTx(&_ChainOracle.CallOpts, _data)
}

// DecodeRLPHeader is a free data retrieval call binding the contract method 0x26410af6.
//
// Solidity: function decodeRLPHeader(bytes _data) view returns((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes32,uint256))
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
// Solidity: function decodeRLPHeader(bytes _data) view returns((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes32,uint256))
func (_ChainOracle *ChainOracleSession) DecodeRLPHeader(_data []byte) (ChainOracleL2Header, error) {
	return _ChainOracle.Contract.DecodeRLPHeader(&_ChainOracle.CallOpts, _data)
}

// DecodeRLPHeader is a free data retrieval call binding the contract method 0x26410af6.
//
// Solidity: function decodeRLPHeader(bytes _data) view returns((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes32,uint256))
func (_ChainOracle *ChainOracleCallerSession) DecodeRLPHeader(_data []byte) (ChainOracleL2Header, error) {
	return _ChainOracle.Contract.DecodeRLPHeader(&_ChainOracle.CallOpts, _data)
}

// ExtractData is a free data retrieval call binding the contract method 0xacd23ff9.
//
// Solidity: function extractData(bytes[] raw, (uint256,uint256)[] ranges) pure returns(bytes)
func (_ChainOracle *ChainOracleCaller) ExtractData(opts *bind.CallOpts, raw [][]byte, ranges []ChainOracleShareRange) ([]byte, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "extractData", raw, ranges)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ExtractData is a free data retrieval call binding the contract method 0xacd23ff9.
//
// Solidity: function extractData(bytes[] raw, (uint256,uint256)[] ranges) pure returns(bytes)
func (_ChainOracle *ChainOracleSession) ExtractData(raw [][]byte, ranges []ChainOracleShareRange) ([]byte, error) {
	return _ChainOracle.Contract.ExtractData(&_ChainOracle.CallOpts, raw, ranges)
}

// ExtractData is a free data retrieval call binding the contract method 0xacd23ff9.
//
// Solidity: function extractData(bytes[] raw, (uint256,uint256)[] ranges) pure returns(bytes)
func (_ChainOracle *ChainOracleCallerSession) ExtractData(raw [][]byte, ranges []ChainOracleShareRange) ([]byte, error) {
	return _ChainOracle.Contract.ExtractData(&_ChainOracle.CallOpts, raw, ranges)
}

// GetHeader is a free data retrieval call binding the contract method 0xb9615878.
//
// Solidity: function getHeader(bytes32 _headerHash) view returns((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes32,uint256))
func (_ChainOracle *ChainOracleCaller) GetHeader(opts *bind.CallOpts, _headerHash [32]byte) (ChainOracleL2Header, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "getHeader", _headerHash)

	if err != nil {
		return *new(ChainOracleL2Header), err
	}

	out0 := *abi.ConvertType(out[0], new(ChainOracleL2Header)).(*ChainOracleL2Header)

	return out0, err

}

// GetHeader is a free data retrieval call binding the contract method 0xb9615878.
//
// Solidity: function getHeader(bytes32 _headerHash) view returns((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes32,uint256))
func (_ChainOracle *ChainOracleSession) GetHeader(_headerHash [32]byte) (ChainOracleL2Header, error) {
	return _ChainOracle.Contract.GetHeader(&_ChainOracle.CallOpts, _headerHash)
}

// GetHeader is a free data retrieval call binding the contract method 0xb9615878.
//
// Solidity: function getHeader(bytes32 _headerHash) view returns((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes32,uint256))
func (_ChainOracle *ChainOracleCallerSession) GetHeader(_headerHash [32]byte) (ChainOracleL2Header, error) {
	return _ChainOracle.Contract.GetHeader(&_ChainOracle.CallOpts, _headerHash)
}

// GetTransaction is a free data retrieval call binding the contract method 0x4aae13ca.
//
// Solidity: function getTransaction(bytes32 _txHash) view returns((uint256,uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256))
func (_ChainOracle *ChainOracleCaller) GetTransaction(opts *bind.CallOpts, _txHash [32]byte) (ChainOracleDepositTx, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "getTransaction", _txHash)

	if err != nil {
		return *new(ChainOracleDepositTx), err
	}

	out0 := *abi.ConvertType(out[0], new(ChainOracleDepositTx)).(*ChainOracleDepositTx)

	return out0, err

}

// GetTransaction is a free data retrieval call binding the contract method 0x4aae13ca.
//
// Solidity: function getTransaction(bytes32 _txHash) view returns((uint256,uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256))
func (_ChainOracle *ChainOracleSession) GetTransaction(_txHash [32]byte) (ChainOracleDepositTx, error) {
	return _ChainOracle.Contract.GetTransaction(&_ChainOracle.CallOpts, _txHash)
}

// GetTransaction is a free data retrieval call binding the contract method 0x4aae13ca.
//
// Solidity: function getTransaction(bytes32 _txHash) view returns((uint256,uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256))
func (_ChainOracle *ChainOracleCallerSession) GetTransaction(_txHash [32]byte) (ChainOracleDepositTx, error) {
	return _ChainOracle.Contract.GetTransaction(&_ChainOracle.CallOpts, _txHash)
}

// HashHeader is a free data retrieval call binding the contract method 0xefd0f984.
//
// Solidity: function hashHeader((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes32,uint256) _header) pure returns(bytes32)
func (_ChainOracle *ChainOracleCaller) HashHeader(opts *bind.CallOpts, _header ChainOracleL2Header) ([32]byte, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "hashHeader", _header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashHeader is a free data retrieval call binding the contract method 0xefd0f984.
//
// Solidity: function hashHeader((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes32,uint256) _header) pure returns(bytes32)
func (_ChainOracle *ChainOracleSession) HashHeader(_header ChainOracleL2Header) ([32]byte, error) {
	return _ChainOracle.Contract.HashHeader(&_ChainOracle.CallOpts, _header)
}

// HashHeader is a free data retrieval call binding the contract method 0xefd0f984.
//
// Solidity: function hashHeader((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes32,uint256) _header) pure returns(bytes32)
func (_ChainOracle *ChainOracleCallerSession) HashHeader(_header ChainOracleL2Header) ([32]byte, error) {
	return _ChainOracle.Contract.HashHeader(&_ChainOracle.CallOpts, _header)
}

// HeaderToRblock is a free data retrieval call binding the contract method 0x7afdc391.
//
// Solidity: function headerToRblock(bytes32 ) view returns(bytes32)
func (_ChainOracle *ChainOracleCaller) HeaderToRblock(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "headerToRblock", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HeaderToRblock is a free data retrieval call binding the contract method 0x7afdc391.
//
// Solidity: function headerToRblock(bytes32 ) view returns(bytes32)
func (_ChainOracle *ChainOracleSession) HeaderToRblock(arg0 [32]byte) ([32]byte, error) {
	return _ChainOracle.Contract.HeaderToRblock(&_ChainOracle.CallOpts, arg0)
}

// HeaderToRblock is a free data retrieval call binding the contract method 0x7afdc391.
//
// Solidity: function headerToRblock(bytes32 ) view returns(bytes32)
func (_ChainOracle *ChainOracleCallerSession) HeaderToRblock(arg0 [32]byte) ([32]byte, error) {
	return _ChainOracle.Contract.HeaderToRblock(&_ChainOracle.CallOpts, arg0)
}

// Headers is a free data retrieval call binding the contract method 0x9e7f2700.
//
// Solidity: function headers(bytes32 ) view returns(bytes32 parentHash, bytes32 uncleHash, address beneficiary, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, bytes logsBloom, uint256 difficulty, uint256 number, uint256 gasLimit, uint256 gasUsed, uint256 timestamp, bytes extraData, bytes32 mixHash, uint256 nonce)
func (_ChainOracle *ChainOracleCaller) Headers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ParentHash       [32]byte
	UncleHash        [32]byte
	Beneficiary      common.Address
	StateRoot        [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	LogsBloom        []byte
	Difficulty       *big.Int
	Number           *big.Int
	GasLimit         *big.Int
	GasUsed          *big.Int
	Timestamp        *big.Int
	ExtraData        []byte
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
		LogsBloom        []byte
		Difficulty       *big.Int
		Number           *big.Int
		GasLimit         *big.Int
		GasUsed          *big.Int
		Timestamp        *big.Int
		ExtraData        []byte
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
	outstruct.LogsBloom = *abi.ConvertType(out[6], new([]byte)).(*[]byte)
	outstruct.Difficulty = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Number = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.GasLimit = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.GasUsed = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.ExtraData = *abi.ConvertType(out[12], new([]byte)).(*[]byte)
	outstruct.MixHash = *abi.ConvertType(out[13], new([32]byte)).(*[32]byte)
	outstruct.Nonce = *abi.ConvertType(out[14], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Headers is a free data retrieval call binding the contract method 0x9e7f2700.
//
// Solidity: function headers(bytes32 ) view returns(bytes32 parentHash, bytes32 uncleHash, address beneficiary, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, bytes logsBloom, uint256 difficulty, uint256 number, uint256 gasLimit, uint256 gasUsed, uint256 timestamp, bytes extraData, bytes32 mixHash, uint256 nonce)
func (_ChainOracle *ChainOracleSession) Headers(arg0 [32]byte) (struct {
	ParentHash       [32]byte
	UncleHash        [32]byte
	Beneficiary      common.Address
	StateRoot        [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	LogsBloom        []byte
	Difficulty       *big.Int
	Number           *big.Int
	GasLimit         *big.Int
	GasUsed          *big.Int
	Timestamp        *big.Int
	ExtraData        []byte
	MixHash          [32]byte
	Nonce            *big.Int
}, error) {
	return _ChainOracle.Contract.Headers(&_ChainOracle.CallOpts, arg0)
}

// Headers is a free data retrieval call binding the contract method 0x9e7f2700.
//
// Solidity: function headers(bytes32 ) view returns(bytes32 parentHash, bytes32 uncleHash, address beneficiary, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, bytes logsBloom, uint256 difficulty, uint256 number, uint256 gasLimit, uint256 gasUsed, uint256 timestamp, bytes extraData, bytes32 mixHash, uint256 nonce)
func (_ChainOracle *ChainOracleCallerSession) Headers(arg0 [32]byte) (struct {
	ParentHash       [32]byte
	UncleHash        [32]byte
	Beneficiary      common.Address
	StateRoot        [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	LogsBloom        []byte
	Difficulty       *big.Int
	Number           *big.Int
	GasLimit         *big.Int
	GasUsed          *big.Int
	Timestamp        *big.Int
	ExtraData        []byte
	MixHash          [32]byte
	Nonce            *big.Int
}, error) {
	return _ChainOracle.Contract.Headers(&_ChainOracle.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainOracle *ChainOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainOracle *ChainOracleSession) Owner() (common.Address, error) {
	return _ChainOracle.Contract.Owner(&_ChainOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainOracle *ChainOracleCallerSession) Owner() (common.Address, error) {
	return _ChainOracle.Contract.Owner(&_ChainOracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ChainOracle *ChainOracleCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ChainOracle *ChainOracleSession) ProxiableUUID() ([32]byte, error) {
	return _ChainOracle.Contract.ProxiableUUID(&_ChainOracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ChainOracle *ChainOracleCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ChainOracle.Contract.ProxiableUUID(&_ChainOracle.CallOpts)
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

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions(bytes32 ) view returns(uint256 chainId, uint64 nonce, uint256 gasPrice, uint64 gas, address to, uint256 value, bytes data, uint256 r, uint256 s, uint256 v)
func (_ChainOracle *ChainOracleCaller) Transactions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ChainId  *big.Int
	Nonce    uint64
	GasPrice *big.Int
	Gas      uint64
	To       common.Address
	Value    *big.Int
	Data     []byte
	R        *big.Int
	S        *big.Int
	V        *big.Int
}, error) {
	var out []interface{}
	err := _ChainOracle.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		ChainId  *big.Int
		Nonce    uint64
		GasPrice *big.Int
		Gas      uint64
		To       common.Address
		Value    *big.Int
		Data     []byte
		R        *big.Int
		S        *big.Int
		V        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Nonce = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.GasPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Gas = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.To = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[6], new([]byte)).(*[]byte)
	outstruct.R = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.S = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.V = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions(bytes32 ) view returns(uint256 chainId, uint64 nonce, uint256 gasPrice, uint64 gas, address to, uint256 value, bytes data, uint256 r, uint256 s, uint256 v)
func (_ChainOracle *ChainOracleSession) Transactions(arg0 [32]byte) (struct {
	ChainId  *big.Int
	Nonce    uint64
	GasPrice *big.Int
	Gas      uint64
	To       common.Address
	Value    *big.Int
	Data     []byte
	R        *big.Int
	S        *big.Int
	V        *big.Int
}, error) {
	return _ChainOracle.Contract.Transactions(&_ChainOracle.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions(bytes32 ) view returns(uint256 chainId, uint64 nonce, uint256 gasPrice, uint64 gas, address to, uint256 value, bytes data, uint256 r, uint256 s, uint256 v)
func (_ChainOracle *ChainOracleCallerSession) Transactions(arg0 [32]byte) (struct {
	ChainId  *big.Int
	Nonce    uint64
	GasPrice *big.Int
	Gas      uint64
	To       common.Address
	Value    *big.Int
	Data     []byte
	R        *big.Int
	S        *big.Int
	V        *big.Int
}, error) {
	return _ChainOracle.Contract.Transactions(&_ChainOracle.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _canonicalStateChain, address _daOracle, address _rlpReader) returns()
func (_ChainOracle *ChainOracleTransactor) Initialize(opts *bind.TransactOpts, _canonicalStateChain common.Address, _daOracle common.Address, _rlpReader common.Address) (*types.Transaction, error) {
	return _ChainOracle.contract.Transact(opts, "initialize", _canonicalStateChain, _daOracle, _rlpReader)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _canonicalStateChain, address _daOracle, address _rlpReader) returns()
func (_ChainOracle *ChainOracleSession) Initialize(_canonicalStateChain common.Address, _daOracle common.Address, _rlpReader common.Address) (*types.Transaction, error) {
	return _ChainOracle.Contract.Initialize(&_ChainOracle.TransactOpts, _canonicalStateChain, _daOracle, _rlpReader)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _canonicalStateChain, address _daOracle, address _rlpReader) returns()
func (_ChainOracle *ChainOracleTransactorSession) Initialize(_canonicalStateChain common.Address, _daOracle common.Address, _rlpReader common.Address) (*types.Transaction, error) {
	return _ChainOracle.Contract.Initialize(&_ChainOracle.TransactOpts, _canonicalStateChain, _daOracle, _rlpReader)
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

// ProvideLegacyTx is a paid mutator transaction binding the contract method 0x20a5acaa.
//
// Solidity: function provideLegacyTx(bytes32 _shareKey, (uint256,uint256)[] _range) returns(bytes32)
func (_ChainOracle *ChainOracleTransactor) ProvideLegacyTx(opts *bind.TransactOpts, _shareKey [32]byte, _range []ChainOracleShareRange) (*types.Transaction, error) {
	return _ChainOracle.contract.Transact(opts, "provideLegacyTx", _shareKey, _range)
}

// ProvideLegacyTx is a paid mutator transaction binding the contract method 0x20a5acaa.
//
// Solidity: function provideLegacyTx(bytes32 _shareKey, (uint256,uint256)[] _range) returns(bytes32)
func (_ChainOracle *ChainOracleSession) ProvideLegacyTx(_shareKey [32]byte, _range []ChainOracleShareRange) (*types.Transaction, error) {
	return _ChainOracle.Contract.ProvideLegacyTx(&_ChainOracle.TransactOpts, _shareKey, _range)
}

// ProvideLegacyTx is a paid mutator transaction binding the contract method 0x20a5acaa.
//
// Solidity: function provideLegacyTx(bytes32 _shareKey, (uint256,uint256)[] _range) returns(bytes32)
func (_ChainOracle *ChainOracleTransactorSession) ProvideLegacyTx(_shareKey [32]byte, _range []ChainOracleShareRange) (*types.Transaction, error) {
	return _ChainOracle.Contract.ProvideLegacyTx(&_ChainOracle.TransactOpts, _shareKey, _range)
}

// ProvideShares is a paid mutator transaction binding the contract method 0x8ac0606b.
//
// Solidity: function provideShares(bytes32 _rblock, uint8 _pointer, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes32)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _proof) returns(bytes32)
func (_ChainOracle *ChainOracleTransactor) ProvideShares(opts *bind.TransactOpts, _rblock [32]byte, _pointer uint8, _proof SharesProof) (*types.Transaction, error) {
	return _ChainOracle.contract.Transact(opts, "provideShares", _rblock, _pointer, _proof)
}

// ProvideShares is a paid mutator transaction binding the contract method 0x8ac0606b.
//
// Solidity: function provideShares(bytes32 _rblock, uint8 _pointer, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes32)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _proof) returns(bytes32)
func (_ChainOracle *ChainOracleSession) ProvideShares(_rblock [32]byte, _pointer uint8, _proof SharesProof) (*types.Transaction, error) {
	return _ChainOracle.Contract.ProvideShares(&_ChainOracle.TransactOpts, _rblock, _pointer, _proof)
}

// ProvideShares is a paid mutator transaction binding the contract method 0x8ac0606b.
//
// Solidity: function provideShares(bytes32 _rblock, uint8 _pointer, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes32)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _proof) returns(bytes32)
func (_ChainOracle *ChainOracleTransactorSession) ProvideShares(_rblock [32]byte, _pointer uint8, _proof SharesProof) (*types.Transaction, error) {
	return _ChainOracle.Contract.ProvideShares(&_ChainOracle.TransactOpts, _rblock, _pointer, _proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainOracle *ChainOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainOracle *ChainOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChainOracle.Contract.RenounceOwnership(&_ChainOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainOracle *ChainOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChainOracle.Contract.RenounceOwnership(&_ChainOracle.TransactOpts)
}

// SetRLPReader is a paid mutator transaction binding the contract method 0xa8567593.
//
// Solidity: function setRLPReader(address _rlpReader) returns()
func (_ChainOracle *ChainOracleTransactor) SetRLPReader(opts *bind.TransactOpts, _rlpReader common.Address) (*types.Transaction, error) {
	return _ChainOracle.contract.Transact(opts, "setRLPReader", _rlpReader)
}

// SetRLPReader is a paid mutator transaction binding the contract method 0xa8567593.
//
// Solidity: function setRLPReader(address _rlpReader) returns()
func (_ChainOracle *ChainOracleSession) SetRLPReader(_rlpReader common.Address) (*types.Transaction, error) {
	return _ChainOracle.Contract.SetRLPReader(&_ChainOracle.TransactOpts, _rlpReader)
}

// SetRLPReader is a paid mutator transaction binding the contract method 0xa8567593.
//
// Solidity: function setRLPReader(address _rlpReader) returns()
func (_ChainOracle *ChainOracleTransactorSession) SetRLPReader(_rlpReader common.Address) (*types.Transaction, error) {
	return _ChainOracle.Contract.SetRLPReader(&_ChainOracle.TransactOpts, _rlpReader)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainOracle *ChainOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChainOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainOracle *ChainOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChainOracle.Contract.TransferOwnership(&_ChainOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainOracle *ChainOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChainOracle.Contract.TransferOwnership(&_ChainOracle.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChainOracle *ChainOracleTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChainOracle.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChainOracle *ChainOracleSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChainOracle.Contract.UpgradeToAndCall(&_ChainOracle.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChainOracle *ChainOracleTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChainOracle.Contract.UpgradeToAndCall(&_ChainOracle.TransactOpts, newImplementation, data)
}

// ChainOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ChainOracle contract.
type ChainOracleInitializedIterator struct {
	Event *ChainOracleInitialized // Event containing the contract specifics and raw log

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
func (it *ChainOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainOracleInitialized)
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
		it.Event = new(ChainOracleInitialized)
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
func (it *ChainOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainOracleInitialized represents a Initialized event raised by the ChainOracle contract.
type ChainOracleInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ChainOracle *ChainOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*ChainOracleInitializedIterator, error) {

	logs, sub, err := _ChainOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ChainOracleInitializedIterator{contract: _ChainOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ChainOracle *ChainOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ChainOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _ChainOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainOracleInitialized)
				if err := _ChainOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ChainOracle *ChainOracleFilterer) ParseInitialized(log types.Log) (*ChainOracleInitialized, error) {
	event := new(ChainOracleInitialized)
	if err := _ChainOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChainOracle contract.
type ChainOracleOwnershipTransferredIterator struct {
	Event *ChainOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChainOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainOracleOwnershipTransferred)
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
		it.Event = new(ChainOracleOwnershipTransferred)
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
func (it *ChainOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainOracleOwnershipTransferred represents a OwnershipTransferred event raised by the ChainOracle contract.
type ChainOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChainOracle *ChainOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChainOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChainOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChainOracleOwnershipTransferredIterator{contract: _ChainOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChainOracle *ChainOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChainOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChainOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainOracleOwnershipTransferred)
				if err := _ChainOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ChainOracle *ChainOracleFilterer) ParseOwnershipTransferred(log types.Log) (*ChainOracleOwnershipTransferred, error) {
	event := new(ChainOracleOwnershipTransferred)
	if err := _ChainOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainOracleUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ChainOracle contract.
type ChainOracleUpgradedIterator struct {
	Event *ChainOracleUpgraded // Event containing the contract specifics and raw log

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
func (it *ChainOracleUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainOracleUpgraded)
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
		it.Event = new(ChainOracleUpgraded)
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
func (it *ChainOracleUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainOracleUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainOracleUpgraded represents a Upgraded event raised by the ChainOracle contract.
type ChainOracleUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ChainOracle *ChainOracleFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ChainOracleUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ChainOracle.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ChainOracleUpgradedIterator{contract: _ChainOracle.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ChainOracle *ChainOracleFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ChainOracleUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ChainOracle.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainOracleUpgraded)
				if err := _ChainOracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ChainOracle *ChainOracleFilterer) ParseUpgraded(log types.Log) (*ChainOracleUpgraded, error) {
	event := new(ChainOracleUpgraded)
	if err := _ChainOracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
