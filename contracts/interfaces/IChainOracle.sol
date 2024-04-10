// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "blobstream-contracts/src/lib/verifier/DAVerifier.sol";

interface IChainOracle {
    /// @notice The shares range struct represents the range of shares in a block.
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
    ) external returns (bytes32);

    /// @notice Decodes the shares into an L2 header and stores it
    ///         in the contract.
    /// @param _shareKey - The share key that the header is related to.
    /// @param _range - The range of the shares that contain the header.
    /// @return The hash of the header.
    function provideHeader(
        bytes32 _shareKey,
        ShareRange[] calldata _range
    ) external returns (bytes32);

    /// @notice Decodes the shares into a transaction and stores it
    ///         in the contract.
    /// @param _shareKey - The share key that the transaction is related to.
    /// @param _range - The range of the shares that contain the transaction.
    /// @return The hash of the transaction.
    function provideLegacyTx(
        bytes32 _shareKey,
        ShareRange[] calldata _range
    ) external returns (bytes32);

    /// @notice Calulates the share key from the rblock and share data.
    /// @param _rblock - The rblock that the shares are related to.
    /// @param _shareData - The share data.
    /// @return The share key.
    function ShareKey(
        bytes32 _rblock,
        bytes[] memory _shareData
    ) external pure returns (bytes32);

    /// @notice Stores shares that are provided to the contract.
    function shares(bytes32 _key) external view returns (bytes[] memory);

    /// @notice Stores headers that are provided to the contract.
    function headers(
        bytes32 _headerHash
    ) external view returns (L2Header memory);

    /// @notice Stores transactions that are provided to the contract.
    function transactions(
        bytes32 _txHash
    ) external view returns (DepositTx memory);

    /// @notice Stores the header to rblock mapping.
    function headerToRblock(
        bytes32 _headerHash
    ) external view returns (bytes32);

    /// @notice Returns the header for a given header hash.
    /// @param _headerHash - The hash of the header.
    /// @return The header.
    function getHeader(
        bytes32 _headerHash
    ) external view returns (L2Header memory);

    /// @notice Returns the transaction for a given transaction hash.
    /// @param _txHash - The hash of the transaction.
    /// @return The transaction.
    function getTransaction(
        bytes32 _txHash
    ) external view returns (DepositTx memory);
}
