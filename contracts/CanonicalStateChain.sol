// SPDX-License-Identifier: UNLICENSED
// LightLink Hummingbird v0.1.0

// TODO: use single version
pragma solidity ^0.8.0;

// UUPS
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

// The Canonical State Chain (CSC) can be considered the source of truth for
// the layer two chain. All layer two blocks will eventually be bundled up by
// the hummingbird publisher and published to the CSC. The chain can only be
// rolled back by the challenge contract, assuming the block is within the
// challenge/pending window.
//
// - The chain is append only, and the publisher is the only address that can
//   add new blocks to the chain.
// - The owner of the contract can replace the publisher address, and is
//   expected to be the DAO Governance contract.

contract CanonicalStateChain is UUPSUpgradeable, OwnableUpgradeable {
    struct Header {
        bytes32 prevHash; // PrevHash is the hash of the previous block bundle.
        uint32 epoch; // Epoch refers to a block number on the Ethereum blockchain.
        uint32 l2Height; // L2Height is the index of the Last L2 Block in this bundle.
        // Pointer to the blocks contents on celestia.
        // See `Span` from https://docs.celestia.org/developers/blobstream-offchain#defining-a-chain
        uint32 celestiaHeight;
        uint32 celestiaShareStart;
        uint32 celestiaShareLen;
    }

    struct HeaderMetadata {
        address publisher;
        uint32 timestamp;
    }

    event BlockAdded(uint256 indexed blockNumber);
    event RolledBack(uint256 indexed blockNumber);
    event PublisherChanged(address indexed publisher);
    event ChallengeChanged(address indexed challenge);

    // publisher is the verified address that can add new blocks to the chain.
    // This address can be replaced by the owner of the contract, (expected to be
    // the rollup contract).
    address public publisher;

    // Challenge is the address of the challenge contract. This contract can
    // rollback the chain after a successful challenge is made.
    address public challenge;

    // Rollup Blockchain.
    uint256 public chainHead; // The index of the last block in the chain.
    mapping(bytes32 => Header) public headers; // block hash => header
    mapping(bytes32 => HeaderMetadata) public headerMetadata; // block hash => metadata
    mapping(uint256 => bytes32) public chain; // block number => block hash

    function _authorizeUpgrade(address) internal override onlyOwner {}

    function initialize(
        address _publisher,
        Header memory _header
    ) public initializer {
        __Ownable_init(msg.sender);
        publisher = _publisher;

        // Add the genesis block.
        bytes32 _hash = hash(_header);
        headers[_hash] = _header;
        chain[0] = _hash;
    }

    // pushBlock optimistically pushes block headers to the canonical chain.
    // The only fields that are checked are the epoch and prevHash.
    // Other fields are optimistically assumed to be correct, however they can be
    // challenged and rolled back via challenge contract.
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

        // check that the block is not already in the chain.
        bytes32 _hash = hash(_header);
        require(headers[_hash].epoch == 0, "block already exists");

        // Add the block to the chain.
        chainHead++;
        headers[_hash] = _header;
        chain[chainHead] = _hash;

        // Save the metadata.
        headerMetadata[_hash] = HeaderMetadata(
            msg.sender,
            uint32(block.timestamp)
        );

        emit BlockAdded(chainHead);
    }

    function hash(Header memory _header) internal pure returns (bytes32) {
        return keccak256(abi.encode(_header));
    }

    function getBlock(uint256 _index) public view returns (Header memory) {
        return headers[chain[_index]];
    }

    function getHead() public view returns (Header memory) {
        return headers[chain[chainHead]];
    }

    // Rollback reverts the chain to a previous state, It can only be called by the challenge
    // contract.
    function rollback(uint256 _blockNumber) external {
        require(
            msg.sender == challenge,
            "only challenge contract can rollback chain"
        );
        require(
            _blockNumber < chainHead,
            "block number must be less than chain head"
        );
        chainHead = _blockNumber;
        emit RolledBack(_blockNumber);
    }

    function setPublisher(address _publisher) external onlyOwner {
        publisher = _publisher;
        emit PublisherChanged(_publisher);
    }

    function setChallengeContract(address _challenge) external onlyOwner {
        challenge = _challenge;
        emit ChallengeChanged(_challenge);
    }

    uint256[50] private __gap;
}
