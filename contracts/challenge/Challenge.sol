pragma solidity ^0.8.0;

import "./ChallengeHeader.sol";
import "./ChallengeDataAvailability.sol";
import "./ChallengeExecution.sol";

contract Challenge is
    ChallengeHeader,
    ChallengeDataAvailability,
    ChallengeExecution
{
    constructor(
        address _treasury,
        address _chain,
        address _daOracle,
        address _mipsChallenge
    ) ChallengeBase(_treasury, _chain, _daOracle, _mipsChallenge) {}
}
