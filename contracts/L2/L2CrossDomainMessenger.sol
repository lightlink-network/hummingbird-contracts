// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import {AddressAliasHelper} from "../libraries/AddressAliasHelper.sol";
// import {Predeploys} from "../libraries/Predeploys.sol";
import {CrossDomainMessenger} from "../universal/CrossDomainMessenger.sol";
import {L2ToL1MessagePasser} from "./L2ToL1MessagePasser.sol";
import {Constants} from "../libraries/Constants.sol";
import {L1Block} from "./L1Block.sol";

/// @custom:proxied
/// @custom:predeploy 0x4200000000000000000000000000000000000007
/// @title L2CrossDomainMessenger
/// @notice The L2CrossDomainMessenger is a high-level interface for message passing between L1 and
///         L2 on the L2 side. Users are generally encouraged to use this contract instead of lower
///         level message passing contracts.
contract L2CrossDomainMessenger is CrossDomainMessenger {
    /// @custom:semver 2.1.0
    string public constant version = "2.1.0";
    address l2ToL1MessagePasser;
    address l1BlockAttributes;

    /// @notice Constructs the L2CrossDomainMessenger contract.
    constructor() CrossDomainMessenger() {
        initialize({
            _l1CrossDomainMessenger: CrossDomainMessenger(address(0)),
            _l2ToL1MessagePasser: address(0),
            _l1BlockAttributes: address(0)
        });
    }

    /// @notice Initializer.
    /// @param _l1CrossDomainMessenger L1CrossDomainMessenger contract on the other network.
    function initialize(
        CrossDomainMessenger _l1CrossDomainMessenger,
        address _l2ToL1MessagePasser,
        address _l1BlockAttributes
    ) public initializer {
        __CrossDomainMessenger_init({_otherMessenger: _l1CrossDomainMessenger});
        l2ToL1MessagePasser = _l2ToL1MessagePasser;
        l1BlockAttributes = _l1BlockAttributes;
    }

    /// @notice Getter for the remote messenger.
    ///         Public getter is legacy and will be removed in the future. Use `otherMessenger()` instead.
    /// @return L1CrossDomainMessenger contract.
    /// @custom:legacy
    function l1CrossDomainMessenger()
        public
        view
        returns (CrossDomainMessenger)
    {
        return otherMessenger;
    }

    /// @inheritdoc CrossDomainMessenger
    function _sendMessage(
        address _to,
        uint64 _gasLimit,
        uint256 _value,
        bytes memory _data
    ) internal override {
        L2ToL1MessagePasser(payable(l2ToL1MessagePasser)).initiateWithdrawal{
            value: _value
        }(_to, _gasLimit, _data);
    }

    /// @inheritdoc CrossDomainMessenger
    function gasPayingToken()
        internal
        view
        override
        returns (address addr_, uint8 decimals_)
    {
        (addr_, decimals_) = L1Block(l1BlockAttributes).gasPayingToken();
    }

    /// @inheritdoc CrossDomainMessenger
    function _isOtherMessenger() internal view override returns (bool) {
        return
            AddressAliasHelper.undoL1ToL2Alias(msg.sender) ==
            address(otherMessenger);
    }

    /// @inheritdoc CrossDomainMessenger
    function _isUnsafeTarget(
        address _target
    ) internal view override returns (bool) {
        return
            _target == address(this) || _target == address(l2ToL1MessagePasser);
    }
}
