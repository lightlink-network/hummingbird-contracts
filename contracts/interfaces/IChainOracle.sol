// SPDX-License-Identifier: MIT
// LightLink Hummingbird v0.2.0

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

    function shares(bytes32 _key) external view returns (bytes[] memory);

    function headers(
        bytes32 _headerHash
    ) external view returns (L2Header memory);

    function transactions(
        bytes32 _txHash
    ) external view returns (DepositTx memory);

    function headerToRblock(
        bytes32 _headerHash
    ) external view returns (bytes32);

    function getHeader(
        bytes32 _headerHash
    ) external view returns (L2Header memory);

    function getTransaction(
        bytes32 _txHash
    ) external view returns (DepositTx memory);
}
