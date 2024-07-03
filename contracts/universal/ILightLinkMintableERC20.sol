// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

/// @title ILightLinkMintableERC20
/// @notice This interface is available on the LightLinkMintableERC20 contract.
///         We declare it as a separate interface so that it can be used in
///         custom implementations of LightLinkMintableERC20.
interface ILightLinkMintableERC20 is IERC165 {
    function remoteToken() external view returns (address);

    function bridge() external returns (address);

    function mint(address _to, uint256 _amount) external;

    function burn(address _from, uint256 _amount) external;
}

/// @custom:legacy
/// @title ILegacyMintableERC20
/// @notice This interface was available on the legacy L2StandardERC20 contract.
///         It remains available on the LightLinkMintableERC20 contract for
///         backwards compatibility.
interface ILegacyMintableERC20 is IERC165 {
    function l1Token() external view returns (address);

    function mint(address _to, uint256 _amount) external;

    function burn(address _from, uint256 _amount) external;
}