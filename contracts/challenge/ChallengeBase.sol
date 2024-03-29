// SPDX-License-Identifier: MIT
// LightLink Hummingbird v0.1.1

pragma solidity ^0.8.0;

// UUPS
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import "../interfaces/ITreasury.sol";
import "../interfaces/ICanonicalStateChain.sol";
import "./mips/IMipsChallenge.sol";
import "blobstream-contracts/src/IDAOracle.sol";
import "../interfaces/IChainOracle.sol";

// ChallengeBase is the base contract for all challenges.
// It contains the global variables for challenge period, fee, and reward.
//
// - The owner can set the challenge period, fee, and reward. Thus is expected
//   to be the DAO Governance contract.

contract ChallengeBase is
    UUPSUpgradeable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable
{
    uint256 public challengeWindow; // Maximum age of a block that can be challenged.
    uint256 public challengePeriod; // The period of time that a challenge is open for.
    uint256 public challengeFee; // The fee required to make a challenge.
    uint256 public challengeReward; // The reward for successfully challenging a block.

    address public defender; // The address of the defender.

    IChainOracle public chainOracle; // The address of the chain oracle.
    ITreasury public treasury; // The address of the treasury to pay out challenges.
    ICanonicalStateChain public chain; // The address of the canonical state chain.
    IDAOracle public daOracle; // The address of the data availability oracle.
    IMipsChallenge public mipsChallenge; // The address of the MIPS challenge contract.

    function _authorizeUpgrade(address) internal override onlyOwner {}

    function __ChallengeBase_init(
        address _treasury,
        address _chain,
        address _daOracle,
        address _mipsChallenge,
        address _chainOracle
    ) internal {
        __UUPSUpgradeable_init();
        __Ownable_init(msg.sender);
        __ReentrancyGuard_init();

        challengeWindow = 3 days;
        challengePeriod = 2 days;
        challengeFee = 1.5 ether;
        challengeReward = 0.2 ether; // unused.

        treasury = ITreasury(_treasury); // TODO: remove
        chain = ICanonicalStateChain(_chain);
        daOracle = IDAOracle(_daOracle);
        mipsChallenge = IMipsChallenge(_mipsChallenge);
        chainOracle = IChainOracle(_chainOracle);
    }

    function _isTargetWithinChallengeWindow(
        uint256 _index
    ) internal view returns (bool) {
        return
            block.timestamp <=
            chain.headerMetadata(chain.chain(_index)).timestamp +
                challengeWindow;
    }

    // mustBeWithinChallengeWindow ensures that the block is within the challenge window.
    // It is used to prevent challenges on blocks that are too old.
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

    // mustBeCanonical ensures that the block is in the chain.
    modifier mustBeCanonical(uint256 index) {
        require(index <= chain.chainHead(), "block not in the chain yet");
        _;
    }

    modifier requireChallengeFee() {
        require(msg.value >= challengeFee, "challenge fee not paid");
        _;
    }

    // setters
    function setChallengeWindow(uint256 _challengeWindow) external onlyOwner {
        challengeWindow = _challengeWindow;
    }

    function setChallengePeriod(uint256 _challengePeriod) external onlyOwner {
        challengePeriod = _challengePeriod;
    }

    function setChallengeFee(uint256 _challengeFee) external onlyOwner {
        challengeFee = _challengeFee;
    }

    function setChallengeReward(uint256 _challengeReward) external onlyOwner {
        challengeReward = _challengeReward;
    }

    function setDefender(address _defender) external onlyOwner {
        defender = _defender;
    }

    // gap
    uint256[50] private __gap;
}
