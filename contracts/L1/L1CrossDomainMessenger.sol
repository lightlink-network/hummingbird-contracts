// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import { Predeploys } from "contracts/libraries/Predeploys.sol";
import { OptimismPortal } from "contracts/L1/OptimismPortal.sol";
import { CrossDomainMessenger } from "contracts/universal/CrossDomainMessenger.sol";
import { ISemver } from "contracts/universal/ISemver.sol";

/// @custom:proxied
/// @title L1CrossDomainMessenger
/// @notice The L1CrossDomainMessenger is a message passing interface between L1 and L2 responsible
///         for sending and receiving data on the L1 side. Users are encouraged to use this
///         interface instead of interacting with lower-level contracts directly.
contract L1CrossDomainMessenger is CrossDomainMessenger, ISemver {
    /// @notice Contract of the OptimismPortal.
    /// @custom:network-specific
    OptimismPortal public portal;

    /// @notice Semantic version.
    /// @custom:semver 0.1.4
    string public constant version = "0.1.4";

    /// @notice Constructs the L1CrossDomainMessenger contract.
    constructor() CrossDomainMessenger() {
        initialize({_portal: OptimismPortal(payable(address(0))) });
    }

    /// @notice Initializes the contract.
    /// @param _portal Contract of the OptimismPortal contract on this network.
    function initialize(OptimismPortal _portal) public initializer {
        portal = _portal;
        __CrossDomainMessenger_init({ _otherMessenger: CrossDomainMessenger(Predeploys.L2_CROSS_DOMAIN_MESSENGER) });
    }

    /// @notice Getter function for the OptimismPortal contract on this chain.
    ///         Public getter is legacy and will be removed in the future. Use `portal()` instead.
    /// @return Contract of the OptimismPortal on this chain.
    /// @custom:legacy
    function PORTAL() external view returns (OptimismPortal) {
        return portal;
    }

    /// @inheritdoc CrossDomainMessenger
    function _sendMessage(address _to, uint64 _gasLimit, uint256 _value, bytes memory _data) internal override {
        portal.depositTransaction{ value: _value }({
            _to: _to,
            _value: _value,
            _gasLimit: _gasLimit,
            _isCreation: false,
            _data: _data
        });
    }

    /// @inheritdoc CrossDomainMessenger
    function _isOtherMessenger() internal view override returns (bool) {
        return msg.sender == address(portal); // TODO: add back later portal.l2Sender() == address(otherMessenger);
    }

    /// @inheritdoc CrossDomainMessenger
    function _isUnsafeTarget(address _target) internal view override returns (bool) {
        return _target == address(this) || _target == address(portal);
    }

    /// @inheritdoc CrossDomainMessenger
    function paused() public view override returns (bool) {
        // TODO: Implement pause
        return false;
    }
}