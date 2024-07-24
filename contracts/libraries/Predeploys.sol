// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title Predeploys
/// @notice Contains constant addresses for protocol contracts that are pre-deployed to the L2 system.
//          This excludes the preinstalls (non-protocol contracts).
library Predeploys {
    /// @notice Number of predeploy-namespace addresses reserved for protocol usage.
    uint256 internal constant PREDEPLOY_COUNT = 2048;

    /// @notice Address of the L2CrossDomainMessenger predeploy.
    address internal constant L2_CROSS_DOMAIN_MESSENGER =
        0x4200000000000000000000000000000000000007;

    /// @notice Address of the L2StandardBridge predeploy.
    address internal constant L2_STANDARD_BRIDGE =
        0x4200000000000000000000000000000000000010;

    /// @notice Address of the LightLinkMintableERC20Factory predeploy.
    address internal constant LIGHTLINK_MINTABLE_ERC20_FACTORY =
        0x4200000000000000000000000000000000000012;

    /// @notice Address of the L1Block predeploy.
    address internal constant L1_BLOCK_ATTRIBUTES =
        0x4200000000000000000000000000000000000015;

    /// @notice Address of the L2ToL1MessagePasser predeploy.
    address internal constant L2_TO_L1_MESSAGE_PASSER =
        0x4200000000000000000000000000000000000016;

    /// @notice Address of the ProxyAdmin predeploy.
    address internal constant PROXY_ADMIN =
        0x4200000000000000000000000000000000000018;

    /// @notice Returns the name of the predeploy at the given address.
    function getName(address _addr) internal pure returns (string memory out_) {
        require(
            isPredeployNamespace(_addr),
            "Predeploys: address must be a predeploy"
        );
        if (_addr == L2_CROSS_DOMAIN_MESSENGER) return "L2CrossDomainMessenger";
        if (_addr == L2_STANDARD_BRIDGE) return "L2StandardBridge";
        if (_addr == LIGHTLINK_MINTABLE_ERC20_FACTORY)
            return "LightLinkMintableERC20Factory";
        if (_addr == L1_BLOCK_ATTRIBUTES) return "L1Block";
        if (_addr == L2_TO_L1_MESSAGE_PASSER) return "L2ToL1MessagePasser";
        if (_addr == PROXY_ADMIN) return "ProxyAdmin";
        revert("Predeploys: unnamed predeploy");
    }

    /// @notice Returns true if the predeploy is not proxied.
    function notProxied(address _addr) internal pure returns (bool) {
        return false;
    }

    /// @notice Returns true if the address is a defined predeploy that is embedded into new OP-Stack chains.
    function isSupportedPredeploy(address _addr) internal pure returns (bool) {
        return
            _addr == L2_CROSS_DOMAIN_MESSENGER ||
            _addr == L2_STANDARD_BRIDGE ||
            _addr == LIGHTLINK_MINTABLE_ERC20_FACTORY ||
            _addr == L1_BLOCK_ATTRIBUTES ||
            _addr == L2_TO_L1_MESSAGE_PASSER ||
            _addr == PROXY_ADMIN;
    }

    function isPredeployNamespace(address _addr) internal pure returns (bool) {
        return
            uint160(_addr) >> 11 ==
            uint160(0x4200000000000000000000000000000000000000) >> 11;
    }

    /// @notice Function to compute the expected address of the predeploy implementation
    ///         in the genesis state.
    function predeployToCodeNamespace(
        address _addr
    ) internal pure returns (address) {
        require(
            isPredeployNamespace(_addr),
            "Predeploys: can only derive code-namespace address for predeploy addresses"
        );
        return
            address(
                uint160(
                    (uint256(uint160(_addr)) & 0xffff) |
                        uint256(
                            uint160(0xc0D3C0d3C0d3C0D3c0d3C0d3c0D3C0d3c0d30000)
                        )
                )
            );
    }
}
