// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import {Predeploys} from "../libraries/Predeploys.sol";
import {Constants} from "../libraries/Constants.sol";
import {LightLinkPortal} from "./LightLinkPortal.sol";
import {CrossDomainMessenger} from "../libraries/CrossDomainMessenger.sol";

/// @custom:proxied
/// @title L1CrossDomainMessenger
/// @notice The L1CrossDomainMessenger is a message passing interface between L1 and L2 responsible
///         for sending and receiving data on the L1 side. Users are encouraged to use this
///         interface instead of interacting with lower-level contracts directly.
contract L1CrossDomainMessenger is CrossDomainMessenger {
    /// @notice Contract of the LightLinkPortal.
    /// @custom:network-specific
    LightLinkPortal public portal;

    /// @notice Semantic version.
    /// @custom:semver 2.4.0
    string public constant version = "2.4.0";

    /// @notice Constructs the L1CrossDomainMessenger contract.
    constructor() CrossDomainMessenger() {
        initialize({_portal: LightLinkPortal(payable(address(0)))});
    }

    /// @notice Initializes the contract.
    /// @param _portal Contract of the LightLinkPortal contract on this network.
    function initialize(LightLinkPortal _portal) public initializer {
        portal = _portal;
        __CrossDomainMessenger_init({
            _otherMessenger: CrossDomainMessenger(
                Predeploys.L2_CROSS_DOMAIN_MESSENGER
            )
        });
    }

    /// @inheritdoc CrossDomainMessenger
    function gasPayingToken()
        internal
        view
        override
        returns (address _addr, uint8 _decimals)
    {
        // (addr_, decimals_) = systemConfig.gasPayingToken();
        // TODO: Uncomment the above line when the gas paying token is implemented.
        _addr = Constants.ETHER;
        _decimals = 18;
    }

    /// @notice Getter function for the LightLinkPortal contract on this chain.
    ///         Public getter is legacy and will be removed in the future. Use `portal()` instead.
    /// @return Contract of the LightLinkPortal on this chain.
    /// @custom:legacy
    function PORTAL() external view returns (LightLinkPortal) {
        return portal;
    }

    /// @inheritdoc CrossDomainMessenger
    function _sendMessage(
        address _to,
        uint64 _gasLimit,
        uint256 _value,
        bytes memory _data
    ) internal override {
        portal.depositTransaction{value: _value}({
            _to: _to,
            _value: _value,
            _gasLimit: _gasLimit,
            _isCreation: false,
            _data: _data
        });
    }

    /// @inheritdoc CrossDomainMessenger
    function _isOtherMessenger() internal view override returns (bool) {
        return
            msg.sender == address(portal) &&
            portal.l2Sender() == address(otherMessenger);
    }

    /// @inheritdoc CrossDomainMessenger
    function _isUnsafeTarget(
        address _target
    ) internal view override returns (bool) {
        return _target == address(this) || _target == address(portal);
    }

    /// @inheritdoc CrossDomainMessenger
    function paused() public view override returns (bool) {
        return false;
    }
}
