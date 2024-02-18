pragma solidity ^0.8.0;

import "blobstream-contracts/src/lib/verifier/DAVerifier.sol";

interface IChainOracle {
    // Structures (These often go in separate interfaces in larger projects)
    struct ShareRange {
        uint256 start;
        uint256 end;
    }

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

    // External Functions
    function provideShares(
        bytes32 _rblock,
        SharesProof memory _proof
    ) external returns (bytes32);

    function provideHeader(
        bytes32 _shareKey,
        ShareRange[] calldata _range
    ) external returns (bytes32);

    function provideLegacyTx(
        bytes32 _shareKey,
        ShareRange[] calldata _range
    ) external returns (bytes32);

    function ShareKey(
        bytes32 _rblock,
        bytes[] memory _shareData
    ) external pure returns (bytes32);

    // It's up to you whether you want to include `extractData`, `decodeRLPHeader`, etc.
    // in the interface or keep them as internal functions within the contract.
}
