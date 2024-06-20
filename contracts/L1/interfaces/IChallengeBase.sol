// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

interface IChallengeBase {
    /// @return The total time in seconds for a block to be finalized.
    function finalizationSeconds() external view returns (uint256);
}
