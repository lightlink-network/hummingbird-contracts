pragma solidity ^0.8.0;

import "./interfaces/IRLPReader.sol";
import "blobstream-contracts/src/lib/verifier/DAVerifier.sol";
import "blobstream-contracts/src/IDAOracle.sol";
import "./interfaces/ICanonicalStateChain.sol";

// hardhat console
import "hardhat/console.sol";

contract ChainOracle {

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

    mapping(bytes32 => bytes[]) public shares;
    mapping(bytes32 => L2Header) public headers;

    constructor(address _canonicalStateChain, address _daOracle, address _rlpReader) {
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
        require(rHead.celestiaHeight == _proof.attestationProof.tuple.height, "rblock height mismatch");
        
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
        return shareKey;
    }

    function provideHeader(bytes32 _shareKey, ShareRange[] calldata _range) public  returns (bytes32) {
        bytes[] storage shareData = shares[_shareKey];
        require(shareData.length > 0, "share not found");
        
        // 1. Decode the RLP header.
        L2Header memory header = decodeRLPHeader(extractData(shareData, _range));
        require(header.number > 0, "header number is 0");

        // 2. Hash the header.
        bytes32 headerHash = hashHeader(header);
 
        // 3. Store the header.
        require(headers[headerHash].number == 0, "header already exists");
        headers[headerHash] = header;

        return headerHash;
    }

    function ShareKey(bytes32 _rblock, bytes[] memory _shareData) public pure returns (bytes32) {
      return keccak256(abi.encode(_rblock, _shareData));
    }

    // extractData extracts the data from the shares using the range.
    // TODO: Move to a library
      function extractData(bytes[] memory raw, ShareRange[] memory ranges) public pure returns (bytes memory) {
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
        // 1. Decode the RLP header.
        RLPItem[] memory decodedHeader = rlpReader.readList(
            rlpReader.toRLPItem(_data)
        );

        require(decodedHeader.length == 15, "invalid header length");

        // 2. Create a header struct.
        L2Header memory header;

        // 3. Decode the header.
        header.parentHash = rlpReader.readBytes32(decodedHeader[0]);
        header.uncleHash = rlpReader.readBytes32(decodedHeader[1]);
        header.beneficiary = rlpReader.readAddress(decodedHeader[2]);
        header.stateRoot = rlpReader.readBytes32(decodedHeader[3]);
        header.transactionsRoot = rlpReader.readBytes32(decodedHeader[4]);
        header.receiptsRoot = rlpReader.readBytes32(decodedHeader[5]);
        header.logsBloom = rlpReader.readBytes32(decodedHeader[6]);
        header.difficulty = rlpReader.readUint256(decodedHeader[7]);
        header.number = rlpReader.readUint256(decodedHeader[8]);
        header.gasLimit = rlpReader.readUint256(decodedHeader[9]);
        header.gasUsed = rlpReader.readUint256(decodedHeader[10]);
        header.timestamp = rlpReader.readUint256(decodedHeader[11]);
        header.extraData = rlpReader.readBytes32(decodedHeader[12]);
        header.mixHash = rlpReader.readBytes32(decodedHeader[13]);
        header.nonce = rlpReader.readUint256(decodedHeader[14]);

        return header;
    }

    function hashHeader(L2Header memory _header) public pure returns (bytes32) {
        return keccak256(
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