// SPDX-License-Identifier: UNLICENSED
// LightLink Hummingbird v0.0.3
// TODO: use single version
pragma solidity ^0.8.0;

// UUPS
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "./interfaces/IRLPReader.sol";
import "blobstream-contracts/src/lib/verifier/DAVerifier.sol";
import "blobstream-contracts/src/IDAOracle.sol";
import "./interfaces/ICanonicalStateChain.sol";
import "./lib/Lib_RLPEncode.sol";

import "hardhat/console.sol";

// TODO: remove this in production
// hardhat console
import "hardhat/console.sol";

contract ChainOracle is UUPSUpgradeable, OwnableUpgradeable {
    ICanonicalStateChain public canonicalStateChain;
    IDAOracle public daOracle;
    IRLPReader public rlpReader;

    struct ShareRange {
        uint256 start;
        uint256 end;
    }

    // L2 Header
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

    // Supported transaction types

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

    mapping(bytes32 => bytes[]) public shares;
    mapping(bytes32 => L2Header) public headers;
    mapping(bytes32 => DepositTx) public transactions;

    // a special mapping of sharekey to rblock
    mapping(bytes32 => bytes32) private _sharekeyToRblock;
    mapping(bytes32 => bytes32) public headerToRblock;

    function _authorizeUpgrade(address) internal override onlyOwner {}

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

    // provideShares loads some shares that were uploaded to the Data Availability layer. It verifies the shares
    // are included in a given rblock (bundle) and stores them in the contract.
    // @param _rblock The rblock (bundle) that the shares are related to.
    // @param _proof The proof that the shares are available and part of the rblocks dataroot commitment.
    function provideShares(
        bytes32 _rblock,
        SharesProof memory _proof
    ) public returns (bytes32) {
        // 1. Load the rblock (bundle) from the canonical state chain.
        ICanonicalStateChain.Header memory rHead = canonicalStateChain.headers(
            _rblock
        );
        require(rHead.epoch > 0, "rblock not found");
        require(
            rHead.celestiaHeight == _proof.attestationProof.tuple.height,
            "rblock height mismatch"
        );

        // 2. verify shares are valid
        DAVerifier.verifySharesToDataRootTupleRoot(
            daOracle,
            _proof,
            _proof.attestationProof.tuple.dataRoot
        );

        // 3. create a share by hashing the rblock and shares
        bytes32 shareKey = ShareKey(_rblock, _proof.data);

        // 4. store the shares
        shares[shareKey] = _proof.data;

        // 5. store the sharekey to rblock
        _sharekeyToRblock[shareKey] = _rblock;

        return shareKey;
    }

    function provideHeader(
        bytes32 _shareKey,
        ShareRange[] calldata _range
    ) public returns (bytes32) {
        bytes[] storage shareData = shares[_shareKey];
        require(shareData.length > 0, "share not found");

        console.log(">> Shares found");

        // 1. Decode the RLP header.
        L2Header memory header = decodeRLPHeader(
            extractData(shareData, _range)
        );
        console.log(">> Header decoded");
        require(header.number > 0, "header number is 0");

        // 2. Hash the header.
        bytes32 headerHash = hashHeader(header);
        console.log(">> Header hashed");

        // 3. Store the header.
        require(headers[headerHash].number == 0, "header already exists");
        headers[headerHash] = header;
        console.log(">> Header stored");

        // 4. Store the header to rblock
        headerToRblock[headerHash] = _sharekeyToRblock[_shareKey];
        console.log(">> Header to rblock stored");

        return headerHash;
    }

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

    function ShareKey(
        bytes32 _rblock,
        bytes[] memory _shareData
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(_rblock, _shareData));
    }

    // extractData extracts the data from the shares using the range.
    // TODO: Move to a library
    function extractData(
        bytes[] memory raw,
        ShareRange[] memory ranges
    ) public pure returns (bytes memory) {
        bytes memory data;

        for (uint i = 0; i < ranges.length; i++) {
            ShareRange memory r = ranges[i];

            // Ensure that the range is valid for the corresponding raw data
            require(r.end <= raw[i].length, "Invalid range");

            // Concatenating the specified range of bytes
            for (uint j = r.start; j < r.end; j++) {
                data = abi.encodePacked(data, raw[i][j]);
            }
        }

        return data;
    }

    // decodeRLPHHeader decodes an RLP header into the Header struct.
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

    // HashHeader hashes an Ethereum header in the same way that it is hashed on Ethereum.
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
            uint8 v,
            bytes32 r,
            bytes32 s
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

    uint256[50] private __gap;
}
