// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/// @custom:proxied
/// @title Canonical State Chain
/// @author LightLink Hummingbird
/// @custom:version v1.0.0-alpha
/// @notice The Canonical State Chain (CSC) is the source of truth for the layer two chain.
///         All layer two blocks will eventually be bundled up by the hummingbird publisher
///         and published to the CSC. The chain can only be rolled back by the challenge contract,
///         assuming the block is within the challenge/pending window.
///
///         The chain is append only, and the publisher is the only address that can add new blocks
///         to the chain. The owner of the contract can replace the publisher address, and is expected
///         to be the DAO Governance contract.
contract CanonicalStateChain is UUPSUpgradeable, OwnableUpgradeable {
    /// @notice The header of a L1 rollup block.
    /// @param epoch - Refers to a block number on the Ethereum blockchain
    /// @param l2Height - The index of the Last L2 Block in this bundle.
    /// @param prevHash - The hash of the previous block bundle.
    /// @param stateRoot - The Stateroot after applying all the blocks in the Bundle.
    /// @param celestiaPointers - Pointer to the blocks contents on celestia.
    /// See `Span` from https://docs.celestia.org/developers/blobstream-offchain#defining-a-chain
    struct Header {
        uint64 epoch;
        uint64 l2Height;
        bytes32 prevHash;
        bytes32 stateRoot;
        CelestiaPointer[] celestiaPointers;
    }

    /// @notice A pointer to a block on Celestia.
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
    event RolledBack(uint256 indexed blockNumber);

    /// @notice Emitted when the publisher address is changed.
    /// @param publisher - The new publisher address.
    event PublisherChanged(address indexed publisher);

    /// @notice Emitted when the challenge contract address is changed.
    /// @param challenge - The new challenge contract address.
    event ChallengeChanged(address indexed challenge);

    /// @notice The address of the publisher. The publisher is the only address
    ///         that can add new blocks to the chain.
    address public publisher;

    /// @notice The challenge contract is the address of the challenge contract.
    /// This contract can rollback the chain after a successful challenge is made.
    address public challenge;

    /// @notice The index of the last block in the chain.
    uint256 public chainHead;

    /// @notice The maximum number of celestia pointers a block can have.
    uint8 public maxPointers;

    /// @notice The canonical chain of block headers.
    mapping(bytes32 => Header) private headers;

    /// @notice The metadata of a block header.
    mapping(bytes32 => HeaderMetadata) public headerMetadata;

    /// @notice The canonical chain of block hashes.
    mapping(uint256 => bytes32) public chain;

    /// @notice This function is a special internal function that's part of
    ///         the UUPS upgradeable contract's lifecycle. When you want to
    ///         upgrade the contract to a new version, _authorizeUpgrade is
    ///         called to check whether the upgrade is authorized, thus
    ///         preventing anyone from just upgrading the contract.
    /// @dev Only the owner can call this function.
    function _authorizeUpgrade(address) internal override onlyOwner {}

    /// @notice Initializes the contract with the publisher and the genesis block.
    /// @param _publisher - The address of the publisher.
    /// @param _header - The genesis block header.
    function initialize(
        address _publisher,
        Header memory _header
    ) public initializer {
        __Ownable_init(msg.sender);
        publisher = _publisher;

        // Add the genesis block.
        bytes32 _hash = calculateHeaderHash(_header);
        headers[_hash] = _header;
        chain[0] = _hash;

        maxPointers = 7;
    }

    /// @notice Optimistically pushes block headers to the canonical chain.
    ///         The only fields that are checked are the epoch, prevHash and that the
    ///         block has atleast one celestia pointer. Other fields are optimistically
    ///         assumed to be correct, however they can be challenged and rolled back
    ///         via challenge contract.
    /// @param _header - The block header to add.
    function pushBlock(Header calldata _header) external {
        require(msg.sender == publisher, "only publisher can add blocks");

        // Check that the epoch is greater than the previous epoch.
        require(
            _header.epoch > headers[chain[chainHead]].epoch,
            "epoch must be greater than previous epoch"
        );
        require(
            _header.prevHash == chain[chainHead],
            "prevHash must be the previous block hash"
        );
        require(
            _header.celestiaPointers.length > 0,
            "block must have atleast one celestia pointer"
        );
        require(
            _header.celestiaPointers.length <= maxPointers,
            "block has too many celestia pointers"
        );

        // check that the block is not already in the chain.
        bytes32 _hash = calculateHeaderHash(_header);
        require(headers[_hash].epoch == 0, "block already exists");

        // Add the block to the chain.
        chainHead++;
        headers[_hash] = _header;
        chain[chainHead] = _hash;

        // Save the metadata.
        headerMetadata[_hash] = HeaderMetadata(
            uint64(block.timestamp),
            msg.sender
        );

        emit BlockAdded(chainHead);
    }

    /// @notice Returns the hash of a block header.
    /// @param _header - The block header to hash.
    /// @return The hash of the block header.
    function calculateHeaderHash(
        Header memory _header
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(_header));
    }

    /// @notice Returns the block header at a given block number.
    /// @param _index - The block number to get the header for.
    /// @return The block header.
    function getHeaderByNum(
        uint256 _index
    ) public view returns (Header memory) {
        return headers[chain[_index]];
    }

    /// @notice Returns the block header at the head of the chain.
    /// @return The block header.
    function getHead() public view returns (Header memory) {
        return headers[chain[chainHead]];
    }

    /// @notice Rolls back the chain to a previous block number. This function can only be
    ///         called by the challenge contract.
    /// @param _blockNumber - The block number to roll back to.
    /// @param _blockHash - The hash the block being purged.
    function rollback(uint256 _blockNumber, bytes32 _blockHash) external {
        require(
            msg.sender == challenge,
            "only challenge contract can rollback chain"
        );
        require(
            _blockNumber < chainHead,
            "block number must be less than chain head"
        );
        require(
            chain[_blockNumber + 1] == _blockHash,
            "block hash must match block number"
        );
        chainHead = _blockNumber;
        emit RolledBack(_blockNumber);
    }

    /// @notice Sets the publisher address.
    /// @param _publisher - The new publisher address.
    /// @dev Only the owner can call this function.
    function setPublisher(address _publisher) external onlyOwner {
        publisher = _publisher;
        emit PublisherChanged(_publisher);
    }

    /// @notice Sets the challenge contract address.
    /// @param _challenge - The new challenge contract address.
    /// @dev Only the owner can call this function.
    function setChallengeContract(address _challenge) external onlyOwner {
        challenge = _challenge;
        emit ChallengeChanged(_challenge);
    }

    /// @notice Returns the block header hash at a given block number.
    /// @param _hash - The hash of the block header to get.
    function getHeaderByHash(
        bytes32 _hash
    ) public view returns (Header memory) {
        return headers[_hash];
    }

    /// @notice Sets the maximum number of celestia pointers a block can have.
    /// @param _maxPointers - The new maximum number of celestia pointers.
    /// @dev Only the owner can call this function.
    function setMaxPointers(uint8 _maxPointers) external onlyOwner {
        maxPointers = _maxPointers;
    }

    uint256[50] private __gap;
}
