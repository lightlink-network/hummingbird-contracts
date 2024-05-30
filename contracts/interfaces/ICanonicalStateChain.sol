// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ICanonicalStateChain {
    /// @notice The header struct represents a block header in the rollup chain.
    /// @param epoch - Epoch refers to a block number on the Ethereum blockchain.
    /// @param l2Height - L2Height is the index of the Last L2 Block in this bundle.
    /// @param prevHash - PrevHash is the hash of the previous block bundle.
    /// @param stateRoot - The Stateroot after applying all the blocks in the Bundle.
    /// @param celestiaPointers - Pointer to the blocks contents on celestia.
    /// See `Span` from https://docs.celestia.org/developers/blobstream-offchain#defining-a-chain
    struct Header {
        uint64 epoch;
        uint64 l2Height;
        bytes32 prevHash;
        bytes32 stateRoot;
        bytes32 shareRoot;
        CelestiaPointer[] celestiaPointers;
    }

    /// @notice A pointer to a shares on Celestia.
    /// @param height - The height of the block on Celestia.
    /// @param shareStart - The start index of shares in block on Celestia.
    /// @param shareLen - The length of the shares in block on Celestia.
    struct CelestiaPointer {
        uint64 height;
        uint24 shareStart;
        uint16 shareLen;
    }

    /// @notice The metadata of a block header.
    /// @param timestamp - The timestamp the block was added.
    /// @param publisher - The address of the publisher that added the block.
    struct HeaderMetadata {
        uint64 timestamp;
        address publisher;
    }

    /// @notice Emitted when a new block is added to the chain.
    /// @param blockNumber - The block number of the new block.
    event BlockAdded(uint256 indexed blockNumber);

    /// @notice Emitted when the chain is rolled back.
    /// @param blockNumber - The block number the chain was rolled back to.
    event Rolledback(uint256 indexed blockNumber);

    /// @notice Emitted when the publisher address is changed.
    /// @param publisher - The new publisher address.
    event PublisherChanged(address indexed publisher);

    /// @notice The address of the publisher. Publisher is the verified address
    ///         that can add new blocks to the chain. This address can be
    ///         replaced by the owner of the contract, (expected to be the
    ///         rollup contract).
    /// @return The address of the publisher.
    function publisher() external view returns (address);

    /// @notice The address of the challenge contract. Challenge is the address
    ///         of the challenge contract. This contract can rollback the chain
    ///         after a successful challenge is made.
    /// @return The address of the challenge contract.
    function challenge() external view returns (address);

    /// @notice The index of the last block in the chain.
    /// @return The index of the last block in the chain.
    function chainHead() external view returns (uint256);

    /// @notice The canonical chain of block headers.
    /// @return The block header.
    function headers(bytes32) external view returns (Header memory);

    /// @notice Returns the block header by hash.
    /// @return The block header.
    function getHeaderByHash(bytes32) external view returns (Header memory);

    /// @notice The metadata of a block header.
    /// @return The metadata of a block header.
    function headerMetadata(
        bytes32
    ) external view returns (HeaderMetadata memory);

    /// @notice Returns the block hash by number.
    /// @return The block hash.
    function chain(uint256) external view returns (bytes32);

    /// @notice Optimistically pushes block headers to the canonical chain.
    ///         The only fields that are checked are the epoch and prevHash.
    /// @param _header - The block header to push.
    function pushBlock(Header calldata _header) external;

    /// @notice Returns the hash of a block header.
    /// @param _header - The block header to hash.
    function hash(Header memory _header) external pure returns (bytes32);

    /// @notice Returns the hash of a block header.
    /// @return The hash of the block header.
    function getHeaderByNum(
        uint256 _index
    ) external view returns (Header memory);

    /// @notice Returns the header of the last block in the chain.
    /// @return The header of the last block in the chain.
    function getHead() external view returns (Header memory);

    /// @notice Rolls back the chain to a previous block number. Reverts
    ///         the chain to a previous state, It can only be called by
    ///         the challenge contract.
    function rollback(uint256 _blockNumber, bytes32 _blockhash) external;

    /// @notice Sets the publisher address.
    /// @param _publisher - The new publisher address.
    function setPublisher(address _publisher) external;
}
