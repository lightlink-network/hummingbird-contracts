// SPDX-License-Identifier: MIT
// LightLink Hummingbird v0.1.1

pragma solidity ^0.8.0;

interface IRLPReader {
    function isList(bytes memory item) external pure returns (bool);

    function itemLength(bytes memory item) external pure returns (uint);

    function rlpLen(bytes memory item) external pure returns (uint);

    function payloadLocation(
        bytes memory item
    )
        external
        pure
        returns (uint payloadMemPtr, uint payloadLen, uint itemMemPtr);

    function numItems(bytes memory item) external pure returns (uint);

    function rlpBytesKeccak256(
        bytes memory item
    ) external pure returns (bytes32);

    function payloadKeccak256(
        bytes memory item
    ) external pure returns (bytes32);

    function toRlpBytes(bytes memory item) external pure returns (bytes memory);

    function toBytes(bytes memory item) external pure returns (bytes memory);

    function toUint(bytes memory item) external pure returns (uint);

    function toUintStrict(bytes memory item) external pure returns (uint);

    function toAddress(bytes memory item) external pure returns (address);

    function toBoolean(bytes memory item) external pure returns (bool);

    function bytesToString(
        bytes memory item
    ) external pure returns (string memory);

    function toIterator(bytes memory item) external pure;

    function nestedIteration(
        bytes memory item
    ) external pure returns (string memory);

    function toBlockHeader(
        bytes memory rlpHeader
    )
        external
        pure
        returns (
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
        );

    function toLegacyTx(
        bytes memory rlpTx
    )
        external
        pure
        returns (
            uint nonce,
            uint gasPrice,
            uint gasLimit,
            address to,
            uint value,
            bytes memory data,
            uint v,
            uint r,
            uint s
        );

    function toDepositTx(
        bytes memory rlpTx
    )
        external
        pure
        returns (
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
        );
}
