// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import "../interfaces/IChainOracle.sol";
import "../interfaces/ICanonicalStateChain.sol";
import "blobstream-contracts/src/IDAOracle.sol";

/// @title  ChallengeBase
/// @author LightLink Hummingbird
/// @custom:version v1.1.0-beta
/// @notice ChallengeBase is the base contract for all challenges. It contains
///         the global variables for challenge period, fee, and reward. The
///         owner can set the challenge period, fee, and reward. Thus is
///         expected to be the DAO Governance contract.
contract ChallengeBase is
    UUPSUpgradeable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable
{
    /// @notice Maximum age of a block that can be challenged.
    uint256 public challengeWindow;

    /// @notice The period of time that a challenge is open for.
    uint256 public challengePeriod;

    /// @notice The fee required to make a challenge.
    uint256 public challengeFee;

    /// @notice The reward for successfully challenging a block.
    uint256 public challengeReward;

    /// @notice The address of the defender.
    address public defender;

    /// @notice The address of the chain oracle.
    IChainOracle public chainOracle;

    /// @notice The address of the canonical state chain.
    ICanonicalStateChain public chain;

    /// @notice The namespace used for data availability.
    Namespace public daNamespace;

    /// @notice The address of the data availability oracle.
    IDAOracle public daOracle;

    /// @notice This function is a special internal function that's part of
    ///         the UUPS upgradeable contract's lifecycle. When you want to
    ///         upgrade the contract to a new version, _authorizeUpgrade is
    ///         called to check whether the upgrade is authorized, thus
    ///         preventing anyone from just upgrading the contract.
    /// @dev Only the owner can call this function.
    function _authorizeUpgrade(address) internal override onlyOwner {}

    /// @notice Initializes the contract with the chain, daOracle,
    ///         and chainOracle addresses.
    /// @param _chain The address of the canonical state chain.
    /// @param _daOracle The address of the data availability oracle.
    /// @param _chainOracle The address of the chain oracle.
    function __ChallengeBase_init(
        address _chain,
        address _daOracle,
        address _chainOracle
    ) internal {
        __UUPSUpgradeable_init();
        __Ownable_init(msg.sender);
        __ReentrancyGuard_init();

        challengeWindow = 3 days;
        challengePeriod = 2 days;
        challengeFee = 1.5 ether;
        challengeReward = 0.2 ether; // unused.

        chain = ICanonicalStateChain(_chain);
        daOracle = IDAOracle(_daOracle);
        chainOracle = IChainOracle(_chainOracle);
    }

    /// @notice Ensures that the block is within the challenge window. It
    ///         is used to prevent challenges on blocks that are too old.
    /// @param index The index of the block to check.
    modifier mustBeWithinChallengeWindow(uint256 index) {
        require(index != 0, "cannot challenge genesis block");
        require(
            block.timestamp <=
                chain.headerMetadata(chain.chain(index)).timestamp +
                    challengeWindow,
            "block is too old to challenge"
        );
        _;
    }

    /// @notice Ensures that the block is within the chain.
    /// @param index The index of the block to check.
    modifier mustBeCanonical(uint256 index) {
        require(index <= chain.chainHead(), "block not in the chain yet");
        _;
    }

    /// @notice Ensures the challenger has paid the challenge fee.
    modifier requireChallengeFee() {
        require(msg.value == challengeFee, "challenge fee not paid");
        _;
    }

    /// @return The total time in seconds for a block to be finalized.
    function finalizationSeconds() external view returns (uint256) {
        return challengePeriod + challengeWindow;
    }

    /// @notice Sets the challenge window time in seconds.
    /// @param _challengeWindow The new challenge window time.
    /// @dev Only the owner can call this function.
    function setChallengeWindow(uint256 _challengeWindow) external onlyOwner {
        require(
            _challengeWindow >= 12 hours && _challengeWindow <= 3 weeks,
            "challenge window must be between 12 hours and 3 weeks"
        );
        challengeWindow = _challengeWindow;
    }

    /// @notice Sets the challenge period time in seconds.
    /// @param _challengePeriod The new challenge period time.
    /// @dev Only the owner can call this function.
    function setChallengePeriod(uint256 _challengePeriod) external onlyOwner {
        require(
            _challengePeriod >= 12 hours && _challengePeriod <= 3 weeks,
            "challenge period must be between 12 hours and 3 weeks"
        );
        challengePeriod = _challengePeriod;
    }

    /// @notice Sets the challenge fee in wei.
    /// @param _challengeFee The new challenge fee.
    /// @dev Only the owner can call this function.
    function setChallengeFee(uint256 _challengeFee) external onlyOwner {
        require(
            _challengeFee >= 0.01 ether && _challengeFee <= 10 ether,
            "challenge fee must be between 0.01 ether and 10 ether"
        );
        challengeFee = _challengeFee;
    }

    /// @notice Sets the challenge reward in wei.
    /// @param _challengeReward The new challenge reward.
    /// @dev Only the owner can call this function.
    function setChallengeReward(uint256 _challengeReward) external onlyOwner {
        require(
            _challengeReward >= 0.01 ether && _challengeReward <= 10 ether,
            "challenge reward must be between 0.01 ether and 10 ether"
        );
        challengeReward = _challengeReward;
    }

    /// @notice Sets the defender address.
    /// @param _defender The new defender address.
    /// @dev Only the owner can call this function.
    function setDefender(address _defender) external onlyOwner {
        require(_defender != address(0), "defender cannot be the zero address");
        defender = _defender;
    }

    /// @notice Sets the namespace.
    /// @param _namespace The new namespace.
    /// @dev Only the owner can call this function.
    function setDANamespace(Namespace memory _namespace) external onlyOwner {
        daNamespace = _namespace;
    }

    /// @notice Sets the data availability oracle address.
    /// @param _daOracle The new data availability oracle address.
    /// @dev Only the owner can call this function.
    function setDAOracle(address _daOracle) external onlyOwner {
        require(_daOracle != address(0), "daOracle cannot be the zero address");
        daOracle = IDAOracle(_daOracle);
    }

    uint256[50] private __gap;
}
