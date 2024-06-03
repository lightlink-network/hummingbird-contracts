// SPDX-License-Identifier: MIT
pragma solidity 0.8.22; // TODO: use single version

import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "./interfaces/IRLPReader.sol";
import "blobstream-contracts/src/lib/verifier/DAVerifier.sol";
import "blobstream-contracts/src/IDAOracle.sol";
import "./interfaces/ICanonicalStateChain.sol";
import "./lib/Lib_RLPEncode.sol";

/// @custom:proxied
/// @title ChainOracle
/// @author LightLink Hummingbird
/// @custom:version v1.1.0-beta
/// @notice This contract enables any user to directly upload valid Layer 2 blocks, from
///         the data availability layer, on to Layer 1. Once loaded, the headers and
///         transactions can be fetched from the ChainOracle by their respective hashes.
///         This mechanism is crucial for the other challenges listed below.
///
///         Data is loaded in two parts:
///         1. Celestia shares are loaded, along with the required merkle proofs and
///            validator attestations.
///         2. Stored shares can then be decoded into Layer 2 headers and transactions.
contract ChainOracle is UUPSUpgradeable, OwnableUpgradeable {
    /// @notice The Canonical State Chain contract.
    ICanonicalStateChain public canonicalStateChain;

    /// @notice The Data Availability Oracle contract.
    IDAOracle public daOracle;

    /// @notice The RLP Reader contract.
    IRLPReader public rlpReader;

    /// @notice The SharesProof struct.
    /// @param start - The start index of the shares in the block.
    /// @param end - The end index of the shares in the block.
    struct ShareRange {
        uint256 start;
        uint256 end;
    }

    /// @notice An L2 Header.
    /// @param parentHash - The hash of the parent block.
    /// @param uncleHash - The hash of the uncle block.
    /// @param beneficiary - The address of the beneficiary.
    /// @param stateRoot - The state root hash.
    /// @param transactionsRoot - The transactions root hash.
    /// @param receiptsRoot - The receipts root hash.
    /// @param logsBloom - The logs bloom filter.
    /// @param difficulty - The difficulty of the block.
    /// @param number - The block number.
    /// @param gasLimit - The gas limit of the block.
    /// @param gasUsed - The gas used in the block.
    /// @param timestamp - The timestamp of the block.
    /// @param extraData - The extra data of the block.
    /// @param mixHash - The mix hash of the block.
    /// @param nonce - The nonce of the block.
    struct L2Header {
        bytes32 parentHash;
        bytes32 uncleHash;
        address beneficiary;
        bytes32 stateRoot;
        bytes32 transactionsRoot;
        bytes32 receiptsRoot;
        bytes logsBloom;
        uint256 difficulty;
        uint256 number;
        uint256 gasLimit;
        uint256 gasUsed;
        uint256 timestamp;
        bytes extraData;
        bytes32 mixHash;
        uint256 nonce;
    }

    /// @notice A Legacy Transaction.
    /// @param nonce - The nonce of the transaction.
    /// @param gasPrice - The gas price of the transaction.
    /// @param gas - The gas limit of the transaction.
    /// @param to - The address of the recipient.
    /// @param value - The value of the transaction.
    /// @param data - The data of the transaction.
    /// @param r - The r value of the signature.
    /// @param s - The s value of the signature.
    /// @param v - The v value of the signature.
    struct LegacyTx {
        uint64 nonce;
        uint256 gasPrice;
        uint64 gas;
        address to;
        uint256 value;
        bytes data;
        uint256 r;
        uint256 s;
        uint256 v;
    }

    /// @notice A Deposit Transaction.
    /// @param chainId - The chain ID of the transaction.
    /// @param nonce - The nonce of the transaction.
    /// @param gasPrice - The gas price of the transaction.
    /// @param gas - The gas limit of the transaction.
    /// @param to - The address of the recipient.
    /// @param value - The value of the transaction.
    /// @param data - The data of the transaction.
    /// @param r - The r value of the signature.
    /// @param s - The s value of the signature.
    /// @param v - The v value of the signature.
    struct DepositTx {
        uint256 chainId;
        uint64 nonce;
        uint256 gasPrice;
        uint64 gas;
        address to;
        uint256 value;
        bytes data;
        uint256 r;
        uint256 s;
        uint256 v;
    }

    /// @notice Stores shares that are provided to the contract.
    mapping(bytes32 => bytes[]) public shares;

    /// @notice Stores headers that are provided to the contract.
    mapping(bytes32 => L2Header) public headers;

    /// @notice Stores transactions that are provided to the contract.
    mapping(bytes32 => DepositTx) public transactions;

    /// @notice Stores the sharekey to rblock mapping.
    /// @dev a special mapping of sharekey to rblock
    mapping(bytes32 => bytes32) private _sharekeyToRblock;

    /// @notice Stores the header to rblock mapping.
    mapping(bytes32 => bytes32) public headerToRblock;

    /// @notice This function is a special internal function that's part of
    ///         the UUPS upgradeable contract's lifecycle. When you want to
    ///         upgrade the contract to a new version, _authorizeUpgrade is
    ///         called to check whether the upgrade is authorized, thus
    ///         preventing anyone from just upgrading the contract.
    /// @dev Only the owner can call this function.
    function _authorizeUpgrade(address) internal override onlyOwner {}

    /// @notice Initializes the contract with the canonical state chain, the data
    ///         availability oracle, and the RLP reader.
    /// @param _canonicalStateChain - The address of the canonical state chain.
    /// @param _daOracle - The address of the data availability oracle.
    /// @param _rlpReader - The address of the RLP reader.
    function initialize(
        address _canonicalStateChain,
        address _daOracle,
        address _rlpReader
    ) public initializer {
        __Ownable_init(msg.sender);
        canonicalStateChain = ICanonicalStateChain(_canonicalStateChain);
        daOracle = IDAOracle(_daOracle);
        rlpReader = IRLPReader(_rlpReader);
    }

    /// @notice Loads some shares that were uploaded to the Data
    ///         Availability layer. It verifies the shares are included in a
    ///         given rblock (bundle) and stores them in the contract.
    /// @param _rblock - The rblock (bundle) that the shares are related to.
    /// @param _proof - The proof that the shares are available and part of the
    ///               rblocks dataroot commitment.
    /// @return The share key that the shares are stored under.
    function provideShares(
        bytes32 _rblock,
        uint8 _pointer,
        SharesProof memory _proof
    ) public returns (bytes32) {
        // 1. Load the rblock (bundle) from the canonical state chain.
        ICanonicalStateChain.Header memory rHead = canonicalStateChain
            .getHeaderByHash(_rblock);
        require(rHead.epoch > 0, "rblock not found");
        require(
            rHead.celestiaPointers[_pointer].height ==
                _proof.attestationProof.tuple.height,
            "rblock height mismatch"
        );

        // 2. verify shares are valid
        (bool verified, ) = DAVerifier.verifySharesToDataRootTupleRoot(
            daOracle,
            _proof,
            _proof.attestationProof.tuple.dataRoot
        );
        require(verified, "shares not verified");

        (uint256 squaresize, ) = DAVerifier.computeSquareSizeFromRowProof(
            _proof.rowProofs[0]
        );

        // check that the share index is within the celestia pointer range.
        uint64 shareStart = rHead.celestiaPointers[_pointer].shareStart;
        uint64 shareEnd = shareStart +
            rHead.celestiaPointers[_pointer].shareLen;
        uint256 shareIndexInRow = _proof.shareProofs[0].beginKey;
        uint256 shareIndexInRowMajorOrder = shareIndexInRow +
            squaresize *
            _proof.rowProofs[0].key;
        require(
            shareIndexInRowMajorOrder >= shareStart &&
                shareIndexInRowMajorOrder < shareEnd,
            "provided share must be within the celestia pointer range"
        );

        // 3. create a share by hashing the rblock and shares
        bytes32 shareKey = ShareKey(_rblock, _proof.data);

        // 4. store the shares
        shares[shareKey] = _proof.data;

        // 5. store the sharekey to rblock
        _sharekeyToRblock[shareKey] = _rblock;

        return shareKey;
    }

    /// @notice Decodes the shares into an L2 header and stores it
    ///         in the contract.
    /// @param _shareKey - The share key that the header is related to.
    /// @param _range - The range of the shares that contain the header.
    /// @return The hash of the header.
    function provideHeader(
        bytes32 _shareKey,
        ShareRange[] calldata _range
    ) public returns (bytes32) {
        bytes[] storage shareData = shares[_shareKey];
        require(shareData.length > 0, "share not found");

        // 1. Decode the RLP header.
        L2Header memory header = decodeRLPHeader(
            extractData(shareData, _range)
        );
        require(header.number > 0, "header number is 0");

        // 2. Hash the header.
        bytes32 headerHash = hashHeader(header);

        // 3. Store the header.
        require(headers[headerHash].number == 0, "header already exists");
        headers[headerHash] = header;

        // 4. Store the header to rblock
        headerToRblock[headerHash] = _sharekeyToRblock[_shareKey];

        return headerHash;
    }

    /// @notice Decodes the shares into a transaction and stores it
    ///         in the contract.
    /// @param _shareKey - The share key that the transaction is related to.
    /// @param _range - The range of the shares that contain the transaction.
    /// @return The hash of the transaction.
    function provideLegacyTx(
        bytes32 _shareKey,
        ShareRange[] calldata _range
    ) public returns (bytes32) {
        bytes[] storage shareData = shares[_shareKey];
        require(shareData.length > 0, "share not found");

        // 1. Extract the RLP transaction from the shares.
        bytes memory rlpTx = extractData(shareData, _range);

        // 2. Decode the RLP transaction.
        LegacyTx memory ltx = decodeLegacyTx(rlpTx);

        // 3. Hash the transaction.
        bytes32 txHash = keccak256(rlpTx);

        // 4. Store the transaction.
        require(transactions[txHash].nonce == 0, "transaction already exists");
        transactions[txHash] = DepositTx({
            chainId: 0,
            nonce: ltx.nonce,
            gasPrice: ltx.gasPrice,
            gas: ltx.gas,
            to: ltx.to,
            value: ltx.value,
            data: ltx.data,
            r: ltx.r,
            s: ltx.s,
            v: ltx.v
        });

        return txHash;
    }

    /// @notice Calulates the share key from the rblock and share data.
    /// @param _rblock - The rblock that the shares are related to.
    /// @param _shareData - The share data.
    /// @return The share key.
    function ShareKey(
        bytes32 _rblock,
        bytes[] memory _shareData
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(_rblock, _shareData));
    }

    /// TODO: Move to a library
    /// @notice Extracts the data from the shares using the range.
    /// @param raw - The raw data.
    /// @param ranges - The ranges of the data.
    /// @return The extracted data.
    function extractData(
        bytes[] memory raw,
        ShareRange[] memory ranges
    ) public pure returns (bytes memory) {
        // figure out the length of the data
        uint256 length = 0;
        for (uint i = 0; i < ranges.length; i++) {
            ShareRange memory r = ranges[i];
            length += r.end - r.start;
        }

        // copy the data using the ranges
        bytes memory data = new bytes(length);
        uint256 index = 0;
        for (uint i = 0; i < ranges.length; i++) {
            ShareRange memory r = ranges[i];

            // Ensure that the range is valid for the corresponding raw data
            require(r.end <= raw[i].length, "Invalid range");

            for (uint j = r.start; j < r.end; j++) {
                data[index] = raw[i][j];
                index++;
            }
        }

        return data;
    }

    /// @notice Decodes an RLP header into the Header struct.
    /// @param _data - The RLP encoded header.
    /// @return The decoded header.
    function decodeRLPHeader(
        bytes memory _data
    ) public view returns (L2Header memory) {
        (
            bytes32 parentHash,
            bytes32 sha3Uncles,
            address coinbase,
            bytes32 stateRoot,
            bytes32 transactionsRoot,
            bytes32 receiptsRoot,
            uint difficulty,
            uint number,
            uint gasLimit,
            uint gasUsed,
            uint timestamp,
            uint nonce
        ) = rlpReader.toBlockHeader(_data);
        L2Header memory header = L2Header({
            parentHash: parentHash,
            uncleHash: sha3Uncles,
            beneficiary: coinbase,
            stateRoot: stateRoot,
            transactionsRoot: transactionsRoot,
            receiptsRoot: receiptsRoot,
            logsBloom: bytes(
                abi.encodePacked(
                    bytes32(0),
                    bytes32(0),
                    bytes32(0),
                    bytes32(0),
                    bytes32(0),
                    bytes32(0),
                    bytes32(0),
                    bytes32(0)
                )
            ),
            difficulty: difficulty,
            number: number,
            gasLimit: gasLimit,
            gasUsed: gasUsed,
            timestamp: timestamp,
            extraData: bytes(""),
            mixHash: bytes32(0),
            nonce: nonce
        });
        return header;
    }

    /// @notice Hashes an Ethereum header in the same way that it is hashed on Ethereum.
    /// @param _header - The header to hash.
    /// @return The hash of the header.
    function hashHeader(L2Header memory _header) public pure returns (bytes32) {
        bytes[] memory list = new bytes[](15);
        list[0] = RLPEncode.encodeBytes(abi.encodePacked(_header.parentHash));
        list[1] = RLPEncode.encodeBytes(abi.encodePacked(_header.uncleHash));
        list[2] = RLPEncode.encodeAddress(_header.beneficiary);
        list[3] = RLPEncode.encodeBytes(abi.encodePacked(_header.stateRoot));
        list[4] = RLPEncode.encodeBytes(
            abi.encodePacked(_header.transactionsRoot)
        );
        list[5] = RLPEncode.encodeBytes(abi.encodePacked(_header.receiptsRoot));
        list[6] = RLPEncode.encodeBytes(_header.logsBloom);
        list[7] = RLPEncode.encodeUint(_header.difficulty);
        list[8] = RLPEncode.encodeUint(_header.number);
        list[9] = RLPEncode.encodeUint(_header.gasLimit);
        list[10] = RLPEncode.encodeUint(_header.gasUsed);
        list[11] = RLPEncode.encodeUint(_header.timestamp);
        list[12] = RLPEncode.encodeBytes(_header.extraData);
        list[13] = RLPEncode.encodeBytes(abi.encodePacked(_header.mixHash));
        list[14] = RLPEncode.encodeUint(_header.nonce);
        return keccak256(RLPEncode.encodeList(list));
    }

    /// @notice Decodes a legacy transaction from RLP encoded data.
    /// @param _data - The RLP encoded transaction.
    /// @return The decoded transaction.
    function decodeLegacyTx(
        bytes memory _data
    ) public view returns (LegacyTx memory) {
        (
            uint nonce,
            uint gasPrice,
            uint gasLimit,
            address to,
            uint value,
            bytes memory data,
            uint v,
            uint r,
            uint s
        ) = rlpReader.toLegacyTx(_data);

        LegacyTx memory ltx = LegacyTx({
            nonce: uint64(nonce),
            gasPrice: gasPrice,
            gas: uint64(gasLimit),
            to: to,
            value: value,
            data: data,
            r: uint256(r),
            s: uint256(s),
            v: v
        });

        return ltx;
    }

    /// @notice Decodes a deposit transaction from RLP encoded data.
    /// @param _data - The RLP encoded transaction.
    /// @return The decoded transaction.
    function decodeDepositTx(
        bytes memory _data
    ) public view returns (DepositTx memory) {
        (
            uint256 chainId,
            uint nonce,
            uint gasPrice,
            uint gasLimit,
            address to,
            uint value,
            bytes memory data,
            uint8 v,
            bytes32 r,
            bytes32 s
        ) = rlpReader.toDepositTx(_data);
        DepositTx memory dtx = DepositTx({
            chainId: chainId,
            nonce: uint64(nonce),
            gasPrice: gasPrice,
            gas: uint64(gasLimit),
            to: to,
            value: value,
            data: data,
            r: uint256(r),
            s: uint256(s),
            v: v
        });
        return dtx;
    }

    /// @notice Returns the header for a given header hash.
    /// @param _headerHash - The hash of the header.
    /// @return The header.
    function getHeader(
        bytes32 _headerHash
    ) public view returns (L2Header memory) {
        return headers[_headerHash];
    }

    /// @notice Returns the transaction for a given transaction hash.
    /// @param _txHash - The hash of the transaction.
    /// @return The transaction.
    function getTransaction(
        bytes32 _txHash
    ) public view returns (DepositTx memory) {
        return transactions[_txHash];
    }

    /// @notice Sets the RLPReader contract address.
    function setRLPReader(address _rlpReader) public onlyOwner {
        rlpReader = IRLPReader(_rlpReader);
    }

    /// @notice Sets the data availability oracle contract address.
    /// @param _daOracle The new data availability oracle address.
    /// @dev Only the owner can call this function.
    function setDAOracle(address _daOracle) public onlyOwner {
        daOracle = IDAOracle(_daOracle);
    }

    uint256[50] private __gap;
}
