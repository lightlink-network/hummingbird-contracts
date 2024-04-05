// SPDX-License-Identifier: MIT
// LightLink Hummingbird v0.2.0

pragma solidity ^0.8.0;

import "./ChallengeHeader.sol";
import "./ChallengeDataAvailability.sol";
import "./ChallengeL2Header.sol";

// Challenge is the entry point for all validity challenges.
//
// Currently available challenges:
// - DataAvailability – Verifies data root inclusion on Celestia
// - Header – Verifies the validity of a block header
//
// Challenge mechanisms allow for the verification of rollup validity, with
// invalid blocks causing a rollback. Challenges require a fee, incentivizing
// valid challenges and discouraging frivolous ones, while compensating
// defenders for their costs. Challenges must be made within a specified time
// window post-block publication, with late challenges being rejected. The
// rules for fees and timing are outlined in ChallengeBase.sol.

// The following challenges are currently disabled on mainnet, until celestia is ready:
// ChallengeL2Tx
contract Challenge is
    ChallengeHeader,
    ChallengeDataAvailability,
    ChallengeL2Header
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
