pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../interfaces/ITreasury.sol";
import "../interfaces/ICanonicalStateChain.sol";
import "./mips/IMipsChallenge.sol";
import "blobstream-contracts/src/IDAOracle.sol";

contract ChallengeBase is Ownable {
    uint256 challengeWindow; // Maximum age of a block that can be challenged.
    uint256 challengePeriod; // The period of time that a challenge is open for.
    uint256 challengeFee; // The fee required to make a challenge.
    uint256 challengeReward; // The reward for successfully challenging a block.

    address defender; // The address of the defender.

    ITreasury public treasury; // The address of the treasury to pay out challenges.
    ICanonicalStateChain public chain; // The address of the canonical state chain.
    IDAOracle public daOracle; // The address of the data availability oracle.
    IMipsChallenge public mipsChallenge; // The address of the MIPS challenge contract.

    constructor(
        address _treasury,
        address _chain,
        address _daOracle,
        address _mipsChallenge
    ) Ownable(msg.sender) {
        challengeWindow = 1 days;
        challengePeriod = 1 days;
        challengeFee = 0.1 ether;
        challengeReward = 0.2 ether;

        treasury = ITreasury(_treasury);
        chain = ICanonicalStateChain(_chain);
        daOracle = IDAOracle(_daOracle);
        mipsChallenge = IMipsChallenge(_mipsChallenge);
    }

    function _isTargetWithinChallengeWindow(
        uint256 _index
    ) internal view returns (bool) {
        return
            block.timestamp <=
            chain.headerMetadata(chain.chain(_index)).timestamp +
                challengeWindow;
    }

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

    modifier mustBeCanonical(uint256 index) {
        require(index <= chain.chainHead(), "block not in the chain yet");
        _;
    }

    modifier requireChallengeFee() {
        require(msg.value >= challengeFee, "insufficient challenge fee");
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
}
