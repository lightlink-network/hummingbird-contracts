// SPDX-License-Identifier: MIT
// LightLink Hummingbird v0.1.1

pragma solidity ^0.8.0;

interface ITreasury {
  function makePayment(address payable _beneficary, uint256 _amount) external returns (int256);
  function disposableFunds() external view returns (uint256);
}