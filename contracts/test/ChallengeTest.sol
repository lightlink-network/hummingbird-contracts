// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ChallengeHeaderTest.sol";

// Expose internal functions for testing

contract ChallengeTest is ChallengeHeaderTest {
    /// @notice Initializes the Challenge contract.
    /// @param _chain - The address of the chain contract.
    /// @param _daOracle - The address of the data availability oracle.
    /// @param _chainOracle - The address of the chain oracle contract.
    function initialize(
        address _chain,
        address _daOracle,
        address _chainOracle
    ) public initializer {
        __ChallengeBase_init(_chain, _daOracle, _chainOracle);

        __ChallengeHeader_init();
    }
}
