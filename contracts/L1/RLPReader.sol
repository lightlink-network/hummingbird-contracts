// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.5.10 <0.9.0;

import "../libraries/Lib_RLPReader.sol";

/// @title RLPReader
/// @author LightLink Hummingbird
/// @custom:version v1.1.0-beta
/// @notice A contract that reads RLP encoded data.
contract RLPReader {
    using Lib_RLPReader for bytes;
    using Lib_RLPReader for uint;
    using Lib_RLPReader for Lib_RLPReader.RLPItem;
    using Lib_RLPReader for Lib_RLPReader.Iterator;

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
}
