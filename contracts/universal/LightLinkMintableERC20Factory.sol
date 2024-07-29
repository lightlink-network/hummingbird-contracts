// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import {LightLinkMintableERC20} from "../universal/LightLinkMintableERC20.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";

/// @custom:proxied
/// @custom:predeployed 0x4200000000000000000000000000000000000012
/// @title LightLinkMintableERC20Factory
/// @notice LightLinkMintableERC20Factory is a factory contract that generates LightLinkMintableERC20
///         contracts on the network it's deployed to. Simplifies the deployment process for users
///         who may be less familiar with deploying smart contracts. Designed to be backwards
///         compatible with the older StandardL2ERC20Factory contract.
contract LightLinkMintableERC20Factory is Initializable {
    /// @custom:spacer LightLinkMintableERC20Factory's initializer slot spacing
    /// @notice Spacer to avoid packing into the initializer slot
    bytes30 private spacer_0_2_30;

    /// @notice Address of the StandardBridge on this chain.
    /// @custom:network-specific
    address public bridge;

    /// @notice Reserve extra slots in the storage layout for future upgrades.
    ///         A gap size of 49 was chosen here, so that the first slot used in a child contract
    ///         would be 1 plus a multiple of 50.
    uint256[49] private __gap;

    /// @custom:legacy
    /// @notice Emitted whenever a new LightLinkMintableERC20 is created. Legacy version of the newer
    ///         LightLinkMintableERC20Created event. We recommend relying on that event instead.
    /// @param remoteToken Address of the token on the remote chain.
    /// @param localToken  Address of the created token on the local chain.
    event StandardL2TokenCreated(
        address indexed remoteToken,
        address indexed localToken
    );

    /// @notice Emitted whenever a new LightLinkMintableERC20 is created.
    /// @param localToken  Address of the created token on the local chain.
    /// @param remoteToken Address of the corresponding token on the remote chain.
    /// @param deployer    Address of the account that deployed the token.
    event LightLinkMintableERC20Created(
        address indexed localToken,
        address indexed remoteToken,
        address deployer
    );

    /// @notice The semver MUST be bumped any time that there is a change in
    ///         the LightLinkMintableERC20 token contract since this contract
    ///         is responsible for deploying LightLinkMintableERC20 contracts.
    /// @notice Semantic version.
    /// @custom:semver 1.9.0
    string public constant version = "1.9.0";

    /// @notice Constructs the LightLinkMintableERC20Factory contract.
    constructor() {
        initialize({_bridge: address(0)});
    }

    /// @notice Initializes the contract.
    /// @param _bridge Address of the StandardBridge on this chain.
    function initialize(address _bridge) public initializer {
        bridge = _bridge;
    }

    /// @notice Getter function for the address of the StandardBridge on this chain.
    ///         Public getter is legacy and will be removed in the future. Use `bridge` instead.
    /// @return Address of the StandardBridge on this chain.
    /// @custom:legacy
    function BRIDGE() external view returns (address) {
        return bridge;
    }

    /// @custom:legacy
    /// @notice Creates an instance of the LightLinkMintableERC20 contract. Legacy version of the
    ///         newer createLightLinkMintableERC20 function, which has a more intuitive name.
    /// @param _remoteToken Address of the token on the remote chain.
    /// @param _name        ERC20 name.
    /// @param _symbol      ERC20 symbol.
    /// @return Address of the newly created token.
    function createStandardL2Token(
        address _remoteToken,
        string memory _name,
        string memory _symbol
    ) external returns (address) {
        return createLightLinkMintableERC20(_remoteToken, _name, _symbol);
    }

    /// @notice Creates an instance of the LightLinkMintableERC20 contract.
    /// @param _remoteToken Address of the token on the remote chain.
    /// @param _name        ERC20 name.
    /// @param _symbol      ERC20 symbol.
    /// @return Address of the newly created token.
    function createLightLinkMintableERC20(
        address _remoteToken,
        string memory _name,
        string memory _symbol
    ) public returns (address) {
        return
            createLightLinkMintableERC20WithDecimals(
                _remoteToken,
                _name,
                _symbol,
                18
            );
    }

    /// @notice Creates an instance of the LightLinkMintableERC20 contract, with specified decimals.
    /// @param _remoteToken Address of the token on the remote chain.
    /// @param _name        ERC20 name.
    /// @param _symbol      ERC20 symbol.
    /// @param _decimals    ERC20 decimals
    /// @return Address of the newly created token.
    function createLightLinkMintableERC20WithDecimals(
        address _remoteToken,
        string memory _name,
        string memory _symbol,
        uint8 _decimals
    ) public returns (address) {
        require(
            _remoteToken != address(0),
            "LightLinkMintableERC20Factory: must provide remote token address"
        );

        bytes32 salt = keccak256(
            abi.encode(_remoteToken, _name, _symbol, _decimals)
        );
        address localToken = address(
            new LightLinkMintableERC20{salt: salt}(
                bridge,
                _remoteToken,
                _name,
                _symbol,
                _decimals
            )
        );

        // Emit the old event too for legacy support.
        emit StandardL2TokenCreated(_remoteToken, localToken);

        // Emit the updated event. The arguments here differ from the legacy event, but
        // are consistent with the ordering used in StandardBridge events.
        emit LightLinkMintableERC20Created(
            localToken,
            _remoteToken,
            msg.sender
        );

        return localToken;
    }
}
