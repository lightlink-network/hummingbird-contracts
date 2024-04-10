// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.5.10 <0.9.0;

import "./lib/Lib_RLPReader.sol";

/// @title RLPReader
/// @author LightLink Hummingbird
/// @custom:version v1.0.0-alpha
/// @notice A contract that reads RLP encoded data.
contract RLPReader {
    using Lib_RLPReader for bytes;
    using Lib_RLPReader for uint;
    using Lib_RLPReader for Lib_RLPReader.RLPItem;
    using Lib_RLPReader for Lib_RLPReader.Iterator;

    /// @notice Returns true if the item is a list.
    /// @param item - The RLP encoded data.
    /// @return True if the item is a list.
    function isList(bytes memory item) public pure returns (bool) {
        Lib_RLPReader.RLPItem memory rlpItem = item.toRlpItem();
        return rlpItem.isList();
    }

    /// @notice Returns the length of the item.
    /// @param item - The RLP encoded data.
    /// @return The length of the item.
    function itemLength(bytes memory item) public pure returns (uint) {
        uint memPtr;
        assembly {
            memPtr := add(item, 0x20)
        }

        return _itemLength(memPtr);
    }

    /// @notice Returns the length of the RLP encoded data.
    /// @param item - The RLP encoded data.
    /// @return The length of the RLP encoded data.
    function rlpLen(bytes memory item) public pure returns (uint) {
        Lib_RLPReader.RLPItem memory rlpItem = item.toRlpItem();
        return rlpItem.rlpLen();
    }

    /// @notice Returns the location of the payload.
    /// @param item - The RLP encoded data.
    /// @return payloadMemPtr - The memory pointer to the payload.
    /// @return payloadLen - The length of the payload.
    /// @return itemMemPtr - The memory pointer to the item.
    function payloadLocation(
        bytes memory item
    )
        public
        pure
        returns (uint payloadMemPtr, uint payloadLen, uint itemMemPtr)
    {
        Lib_RLPReader.RLPItem memory rlpItem = item.toRlpItem();
        (uint memPtr, uint len) = rlpItem.payloadLocation();
        return (memPtr, len, rlpItem.memPtr);
    }

    /// @notice Returns the number of items in the list.
    /// @param item - The RLP encoded data.
    /// @return The number of items in the list.
    function numItems(bytes memory item) public pure returns (uint) {
        Lib_RLPReader.RLPItem[] memory rlpItem = item.toRlpItem().toList();
        return rlpItem.length;
    }

    /// @notice Returns the keccak256 hash of the RLP encoded data.
    /// @param item - The RLP encoded data.
    /// @return The keccak256 hash of the RLP encoded data.
    function rlpBytesKeccak256(
        bytes memory item
    ) public pure returns (bytes32) {
        Lib_RLPReader.RLPItem memory rlpItem = item.toRlpItem();
        return rlpItem.rlpBytesKeccak256();
    }

    /// @notice Returns the keccak256 hash of the payload.
    /// @param item - The RLP encoded data.
    /// @return The keccak256 hash of the payload.
    function payloadKeccak256(bytes memory item) public pure returns (bytes32) {
        Lib_RLPReader.RLPItem memory rlpItem = item.toRlpItem();
        return rlpItem.payloadKeccak256();
    }

    /// @notice Returns the RLP encoded data.
    /// @param item - The RLP encoded data.
    /// @return The RLP encoded data.
    function toRlpBytes(bytes memory item) public pure returns (bytes memory) {
        Lib_RLPReader.RLPItem memory rlpItem = item.toRlpItem();
        return rlpItem.toRlpBytes();
    }

    /// @notice Returns the RLP encoded data.
    /// @param item - The RLP encoded data.
    /// @return The RLP encoded data.
    function toBytes(bytes memory item) public pure returns (bytes memory) {
        Lib_RLPReader.RLPItem memory rlpItem = item.toRlpItem();
        return rlpItem.toBytes();
    }

    /// @notice Returns the RLP encoded data as a uint.
    /// @param item - The RLP encoded data.
    /// @return The RLP encoded data as a uint.
    function toUint(bytes memory item) public pure returns (uint) {
        Lib_RLPReader.RLPItem memory rlpItem = item.toRlpItem();
        return rlpItem.toUint();
    }

    /// @notice Returns the RLP encoded data as a uint or reverts.
    /// @param item - The RLP encoded data.
    /// @return The RLP encoded data as a uint.
    function toUintStrict(bytes memory item) public pure returns (uint) {
        Lib_RLPReader.RLPItem memory rlpItem = item.toRlpItem();
        return rlpItem.toUintStrict();
    }

    /// @notice Returns the RLP encoded data as an address.
    /// @param item - The RLP encoded data.
    /// @return The RLP encoded data as an address.
    function toAddress(bytes memory item) public pure returns (address) {
        Lib_RLPReader.RLPItem memory rlpItem = item.toRlpItem();
        return rlpItem.toAddress();
    }

    /// @notice Returns the RLP encoded data as a boolean.
    /// @param item - The RLP encoded data.
    /// @return The RLP encoded data as a boolean.
    function toBoolean(bytes memory item) public pure returns (bool) {
        Lib_RLPReader.RLPItem memory rlpItem = item.toRlpItem();
        return rlpItem.toBoolean();
    }

    /// @notice Returns the RLP encoded data as a string.
    /// @param item - The RLP encoded data.
    /// @return The RLP encoded data as a string.
    function bytesToString(
        bytes memory item
    ) public pure returns (string memory) {
        Lib_RLPReader.RLPItem memory rlpItem = item.toRlpItem();
        return string(rlpItem.toBytes());
    }

    /// @notice Returns an iterator for the RLP encoded data.
    /// @param item - The RLP encoded data.å
    function toIterator(bytes memory item) public pure {
        // we just care that this does not revert
        item.toRlpItem().iterator();
    }

    /// @notice Returns the nested iteration of the RLP encoded data.
    /// @param item - The RLP encoded data.
    /// @return The nested iteration of the RLP encoded data.
    function nestedIteration(
        bytes memory item
    ) public pure returns (string memory) {
        Lib_RLPReader.Iterator memory iter = item.toRlpItem().iterator();
        Lib_RLPReader.Iterator memory subIter = iter.next().iterator();
        string memory result = string(subIter.next().toBytes());

        require(!iter.hasNext());
        require(!subIter.hasNext());

        return result;
    }

    /// @notice Returns the block header from the RLP encoded data.
    /// @param rlpHeader - The RLP encoded block header.
    /// @return parentHash - The parent hash.
    /// @return sha3Uncles - The sha3 uncles.
    /// @return coinbase - The coinbase address.
    /// @return stateRoot - The state root.
    /// @return transactionsRoot - The transactions root.
    /// @return receiptsRoot - The receipts root.
    /// @return difficulty - The difficulty.
    /// @return number - The block number.
    /// @return gasLimit - The gas limit.
    /// @return gasUsed - The gas used.
    /// @return timestamp - The timestamp.
    /// @return nonce - The nonce.
    function toBlockHeader(
        bytes memory rlpHeader
    )
        public
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
        )
    {
        Lib_RLPReader.Iterator memory it = rlpHeader.toRlpItem().iterator();
        uint idx;
        while (it.hasNext()) {
            if (idx == 0) parentHash = bytes32(it.next().toUint());
            else if (idx == 1) sha3Uncles = bytes32(it.next().toUint());
            else if (idx == 2) coinbase = it.next().toAddress();
            else if (idx == 3) stateRoot = bytes32(it.next().toUint());
            else if (idx == 4) transactionsRoot = bytes32(it.next().toUint());
            else if (idx == 5) receiptsRoot = bytes32(it.next().toUint());
            else if (idx == 7) difficulty = it.next().toUint();
            else if (idx == 8) number = it.next().toUint();
            else if (idx == 9) gasLimit = it.next().toUint();
            else if (idx == 10) gasUsed = it.next().toUint();
            else if (idx == 11) timestamp = it.next().toUint();
            else if (idx == 14) nonce = it.next().toUint();
            else it.next();

            idx++;
        }
    }

    /// @notice Returns the legacy transaction from the RLP encoded data.
    /// @param rlpTx - The RLP encoded transaction.
    /// @return nonce - The nonce.
    /// @return gasPrice - The gas price.
    /// @return gasLimit - The gas limit.
    /// @return to - The recipient address.
    /// @return value - The value.
    /// @return data - The data.
    /// @return v - The v value.
    /// @return r - The r value.
    /// @return s - The s value.
    function toLegacyTx(
        bytes memory rlpTx
    )
        public
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
        )
    {
        Lib_RLPReader.Iterator memory it = rlpTx.toRlpItem().iterator();
        uint idx;
        while (it.hasNext()) {
            if (idx == 0) nonce = it.next().toUint();
            else if (idx == 1) gasPrice = it.next().toUint();
            else if (idx == 2) gasLimit = it.next().toUint();
            else if (idx == 3) to = it.next().toAddress();
            else if (idx == 4) value = it.next().toUint();
            else if (idx == 5) data = it.next().toBytes();
            else if (idx == 6) v = it.next().toUint();
            else if (idx == 7) r = it.next().toUint();
            else if (idx == 8) s = it.next().toUint();
            else it.next();
            idx++;
        }
        return (nonce, gasPrice, gasLimit, to, value, data, v, r, s);
    }

    /// @notice Returns the deposit transaction from the RLP encoded data.
    /// @param rlpTx - The RLP encoded transaction.
    /// @return chainId - The chain ID.
    /// @return nonce - The nonce.
    /// @return gasPrice - The gas price.
    /// @return gasLimit - The gas limit.
    /// @return to - The recipient address.
    /// @return value - The value.
    /// @return data - The data.
    /// @return v - The v value.
    /// @return r - The r value.
    /// @return s - The s value.
    function toDepositTx(
        bytes memory rlpTx
    )
        public
        pure
        returns (
            uint chainId,
            uint nonce,
            uint gasPrice,
            uint gasLimit,
            address to,
            uint value,
            bytes memory data,
            uint v,
            uint r,
            uint s
        )
    {
        Lib_RLPReader.Iterator memory it = rlpTx.toRlpItem().iterator();
        uint idx;
        while (it.hasNext()) {
            if (idx == 0) chainId = it.next().toUint();
            else if (idx == 1) nonce = it.next().toUint();
            else if (idx == 2) gasPrice = it.next().toUint();
            else if (idx == 3) gasLimit = it.next().toUint();
            else if (idx == 4) to = it.next().toAddress();
            else if (idx == 5) value = it.next().toUint();
            else if (idx == 6) data = it.next().toBytes();
            else if (idx == 7) v = it.next().toUint();
            else if (idx == 8) r = it.next().toUint();
            else if (idx == 9) s = it.next().toUint();
            else it.next();

            idx++;
        }
    }

    /* custom destructuring */

    function customDestructure(
        bytes memory item
    ) public pure returns (address, bool, uint) {
        // first three elements follow the return types in order. Ignore the rest
        Lib_RLPReader.RLPItem[] memory items = item.toRlpItem().toList();
        return (items[0].toAddress(), items[1].toBoolean(), items[2].toUint());
    }

    function customNestedDestructure(
        bytes memory item
    ) public pure returns (address, uint) {
        Lib_RLPReader.RLPItem[] memory items = item.toRlpItem().toList();
        items = items[0].toList();
        return (items[0].toAddress(), items[1].toUint());
    }

    // expects [[bytes, bytes]]
    function customNestedDestructureKeccak(
        bytes memory item
    ) public pure returns (bytes32, bytes32) {
        Lib_RLPReader.RLPItem[] memory items = item.toRlpItem().toList();
        items = items[0].toList();
        return (items[0].payloadKeccak256(), items[1].payloadKeccak256());
    }

    function customNestedToRlpBytes(
        bytes memory item
    ) public pure returns (bytes memory) {
        Lib_RLPReader.RLPItem[] memory items = item.toRlpItem().toList();
        return items[0].toRlpBytes();
    }

    /* Copied verbatim from the reader contract due to scope */
    uint8 constant STRING_SHORT_START = 0x80;
    uint8 constant STRING_LONG_START = 0xb8;
    uint8 constant LIST_SHORT_START = 0xc0;
    uint8 constant LIST_LONG_START = 0xf8;

    function _itemLength(uint memPtr) private pure returns (uint) {
        uint itemLen;
        uint byte0;
        assembly {
            byte0 := byte(0, mload(memPtr))
        }

        if (byte0 < STRING_SHORT_START) itemLen = 1;
        else if (byte0 < STRING_LONG_START)
            itemLen = byte0 - STRING_SHORT_START + 1;
        else if (byte0 < LIST_SHORT_START) {
            assembly {
                let byteLen := sub(byte0, 0xb7) // # of bytes the actual length is
                memPtr := add(memPtr, 1) // skip over the first byte

                /* 32 byte word size */
                let dataLen := div(mload(memPtr), exp(256, sub(32, byteLen))) // right shifting to get the len
                itemLen := add(dataLen, add(byteLen, 1))
            }
        } else if (byte0 < LIST_LONG_START) {
            itemLen = byte0 - LIST_SHORT_START + 1;
        } else {
            assembly {
                let byteLen := sub(byte0, 0xf7)
                memPtr := add(memPtr, 1)

                let dataLen := div(mload(memPtr), exp(256, sub(32, byteLen))) // right shifting to the correct length
                itemLen := add(dataLen, add(byteLen, 1))
            }
        }

        return itemLen;
    }
}
