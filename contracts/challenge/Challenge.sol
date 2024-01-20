// SPDX-License-Identifier: UNLICENSED
// LightLink Hummingbird v0.0.1

pragma solidity ^0.8.0;

import "./ChallengeHeader.sol";
import "./ChallengeDataAvailability.sol";
import "./ChallengeExecution.sol";

// no constructor
contract Challenge is
    ChallengeHeader,
    ChallengeDataAvailability,
    ChallengeExecution
{
    function initialize(
        address _treasury,
        address _chain,
        address _daOracle,
        address _mipsChallenge
    ) public initializer {
        __ChallengeBase_init(_treasury, _chain, _daOracle, _mipsChallenge);

        __ChallengeHeader_init();
    }
}
