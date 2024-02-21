// SPDX-License-Identifier: UNLICENSED
// LightLink Hummingbird v0.0.3

pragma solidity ^0.8.0;

import "./ChallengeHeader.sol";
import "./ChallengeDataAvailability.sol";
import "./ChallengeL2Header.sol";
import "./ChallengeL2Tx.sol";

// no constructor
contract Challenge is
    ChallengeHeader,
    ChallengeDataAvailability,
    ChallengeL2Header,
    ChallengeL2Tx
{
    function initialize(
        address _treasury,
        address _chain,
        address _daOracle,
        address _mipsChallenge,
        address _chainOracle
    ) public initializer {
        __ChallengeBase_init(
            _treasury,
            _chain,
            _daOracle,
            _mipsChallenge,
            _chainOracle
        );

        __ChallengeHeader_init();
    }
}
