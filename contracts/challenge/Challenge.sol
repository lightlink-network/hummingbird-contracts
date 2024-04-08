// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ChallengeHeader.sol";
import "./ChallengeDataAvailability.sol";
import "./ChallengeL2Header.sol";

/// @custom:proxied
/// @title  Challenge
/// @author LightLink Hummingbird
/// @custom:version v1.0.0-alpha
/// @notice Challenge is the entry point for all validity challenges.
///         Challenge mechanisms allow for the verification of rollup
///         validity, with invalid blocks causing a rollback.
///         Challenges require a fee, incentivizing valid challenges
///         and discouraging frivolous ones, while compensating
///         defenders for their costs.
///
///         Challenges must be made within a specified time window
///         post-block publication, with late challenges being
///         rejected. The rules for fees and timing are outlined in
///         ChallengeBase.sol.
contract Challenge is
    ChallengeHeader,
    ChallengeDataAvailability,
    ChallengeL2Header
{
    /// @notice Initializes the Challenge contract.
    /// @param _chain - The address of the chain contract.
    /// @param _daOracle - The address of the data availability oracle.
    /// @param _mipsChallenge - The address of the MIPS challenge contract.
    /// @param _chainOracle - The address of the chain oracle contract.
    function initialize(
        address _chain,
        address _daOracle,
        address _mipsChallenge,
        address _chainOracle
    ) public initializer {
        __ChallengeBase_init(_chain, _daOracle, _mipsChallenge, _chainOracle);

        __ChallengeHeader_init();
    }
}
