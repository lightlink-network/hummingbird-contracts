// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

/// @title CoreProxy
/// @author LightLink Hummingbird
/// @custom:version v1.1.0-beta
/// @notice The core proxy contract for the Hummingbird protocol.
contract CoreProxy is ERC1967Proxy {
    /// @notice CoreProxy constructor
    /// @param _logic - The address of the logic contract
    /// @param _data - The data to be used for the proxy contract
    constructor(
        address _logic,
        bytes memory _data
    ) ERC1967Proxy(_logic, _data) {}
}
