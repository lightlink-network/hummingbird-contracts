pragma solidity ^0.8.0;

import "blobstream-contracts/src/lib/verifier/DAVerifier.sol";
import "blobstream-contracts/src/IDAOracle.sol";
import "./interfaces/ICanonicalStateChain.sol";
import "./lib/Lib_RLPReader.sol";

// TODO: Add support for loading more types: Transaction, Receipt, Log, etc

// ChainLoader is a contract that can be used to load chain data from the Data Availability layer.
// Verifing the data is available and contained with an rblocks body.
//
// Typically data is only loaded as part of a correctness challenge, but this contract can be used to load data
// for any purpose.
//
// While this contract doesnt not verify the data correctness, it does verify that the data is:
// A) Part of full set that is publically available
// B) The content of a submitted rblock (bundle).
// C) RLP decodable into the requested format (e.g. header, transaction, etc)
//
// Loading chain data takes 2 steps: First load the shares (with the share proofs).
// Finally decode and load the data from those shares, e.g. via loadHeader, loadTransaction, etc.
//
// ! Beware the rblock referenced may not be canonical, and may be a fork.
// ! You cannot load data from the genesis rblock.
// ! Loaded data may not be correct.
contract ChainLoader {
    IDAOracle public daOracle;
    ICanonicalStateChain public canonicalStateChain;

    // L2 Header
    struct L2Header {
        bytes32 parentHash;
        bytes32 uncleHash;
        address beneficiary;
        bytes32 stateRoot;
        bytes32 transactionsRoot;
        bytes32 receiptsRoot;
        bytes32 logsBloom;
        uint256 difficulty;
        uint256 number;
        uint256 gasLimit;
        uint256 gasUsed;
        uint256 timestamp;
        bytes32 extraData;
        bytes32 mixHash;
        uint256 nonce;
    }

    struct ShareRange {
        uint64 start; // The index of the first byte in the share.
        uint64 end; // The index of the last byte in the share. (Eclusive)
    }

    // Verified shares that have been stored. The key is the keccack256(abi.encode(keccack256(rblock), share))
    mapping(bytes32 => bytes[]) public shares;

    // Verified headers that have been stored. The key is the hash of the header.
    mapping(bytes32 => L2Header) public headers;

    // loadShares loads some shares that were uploaded to the Data Availability layer. It verifies the shares
    // are included in a given rblock (bundle) and stores them in the contract.
    // @param _rblock The rblock (bundle) that the shares are related to.
    // @param _proof The proof that the shares are available and part of the rblocks dataroot commitment.
    function loadShares(
        bytes32 _rblock,
        SharesProof memory _proof
    ) public returns (bytes32) {
        ICanonicalStateChain.Header memory rHead = canonicalStateChain.headers(
            _rblock
        );

        // 1. Load the rblock (bundle) from the canonical state chain.
        require(rHead.epoch > 0, "rblock not found");

        // 2. verify shares are valid
        DAVerifier.verifySharesToDataRootTupleRoot(
            daOracle,
            _proof,
            rHead.celestiaDataRoot
        );

        // 3. create a share by hashing the rblock and shares
        bytes32 shareKey = hash(_rblock, hash(_proof.data));

        // 4. store the shares
        shares[shareKey] = _proof.data;
        return shareKey;
    }

    // loadHeader loads a header from the shares that have been stored in the contract.
    // @param _rblock The rblock (bundle) that the shares are related to.
    // @param _shareHash The hash of the shares that container the encoded header.
    // @param _shareRange The range in those shares that contain the encoded header.
    function loadHeader(
        bytes32 _rblock,
        bytes32 _shareHash,
        ShareRange[] memory _shareRanges
    ) public returns (bytes32) {
        // 1. construct the share key
        bytes32 shareKey = hash(_rblock, _shareHash);

        // 2. load the shares
        bytes[] storage shareData = shares[shareKey];
        require(shareData.length > 0, "shares not found");

        // 3. extract the data from the shares using the range.
        bytes memory data = extractData(shareData, _shareRanges);

        // 4. decode the header
        L2Header memory header = decodeRLPHeader(data);

        // 5. hash the header
        bytes32 headerHash = hash(header);

        // 6. store the header
        headers[headerHash] = header;

        return headerHash;
    }

    // extractData extracts the data from the shares using the range.
    // TODO: Move to a library
    function extractData(
        bytes[] memory _shareData,
        ShareRange[] memory _shareRanges
    ) public pure returns (bytes memory) {
        // 0. Calculate the total length of the data.
        uint totalLength = 0;
        for (uint i = 0; i < _shareRanges.length; i++) {
            totalLength += _shareRanges[i].end - _shareRanges[i].start;
        }

        // 1. Create a buffer to store the data.
        bytes memory data = new bytes(totalLength);

        // 2. Loop over the shares.
        for (uint i = 0; i < _shareRanges.length; i++) {
            // 3. Get the share data.
            bytes memory share = _shareData[i];

            // 4. Loop over the range.
            for (
                uint j = _shareRanges[i].start;
                j <= _shareRanges[i].end;
                j++
            ) {
                // 5. Copy the data from the share to the buffer.
                data[j] = share[j];
            }
        }

        return data;
    }

    // decodeRLPHHeader decodes an RLP header into the Header struct.
    function decodeRLPHeader(
        bytes memory _data
    ) public pure returns (L2Header memory) {
        // 1. Decode the RLP header.
        Lib_RLPReader.RLPItem[] memory decodedHeader = Lib_RLPReader.readList(
            _data
        );

        // 2. Create a header struct.
        L2Header memory header;

        // 3. Decode the header.
        header.parentHash = Lib_RLPReader.readBytes32(decodedHeader[0]);
        header.uncleHash = Lib_RLPReader.readBytes32(decodedHeader[1]);
        header.beneficiary = Lib_RLPReader.readAddress(decodedHeader[2]);
        header.stateRoot = Lib_RLPReader.readBytes32(decodedHeader[3]);
        header.transactionsRoot = Lib_RLPReader.readBytes32(decodedHeader[4]);
        header.receiptsRoot = Lib_RLPReader.readBytes32(decodedHeader[5]);
        header.logsBloom = Lib_RLPReader.readBytes32(decodedHeader[6]);
        header.difficulty = Lib_RLPReader.readUint256(decodedHeader[7]);
        header.number = Lib_RLPReader.readUint256(decodedHeader[8]);
        header.gasLimit = Lib_RLPReader.readUint256(decodedHeader[9]);
        header.gasUsed = Lib_RLPReader.readUint256(decodedHeader[10]);
        header.timestamp = Lib_RLPReader.readUint256(decodedHeader[11]);
        header.extraData = Lib_RLPReader.readBytes32(decodedHeader[12]);
        header.mixHash = Lib_RLPReader.readBytes32(decodedHeader[13]);
        header.nonce = Lib_RLPReader.readUint256(decodedHeader[14]);

        return header;
    }

    // helpers for hashing

    function hash(bytes[] memory _shareData) internal pure returns (bytes32) {
        return keccak256(abi.encode(_shareData));
    }

    function hash(
        bytes32 _rblock,
        bytes32 _shareHash
    ) internal pure returns (bytes32) {
        return keccak256(abi.encode(_rblock, _shareHash));
    }

    function hash(L2Header memory _header) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    _header.parentHash,
                    _header.uncleHash,
                    _header.beneficiary,
                    _header.stateRoot,
                    _header.transactionsRoot,
                    _header.receiptsRoot,
                    _header.logsBloom,
                    _header.difficulty,
                    _header.number,
                    _header.gasLimit,
                    _header.gasUsed,
                    _header.timestamp,
                    _header.extraData,
                    _header.mixHash,
                    _header.nonce
                )
            );
    }
}
