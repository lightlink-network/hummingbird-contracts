// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import {Script} from "../../lib/forge-std/src/Script.sol";
import {console2 as console} from "../../lib/forge-std/src/console2.sol";
import {Deployer} from "scripts/forge/deploy/Deployer.sol";
import {Config, OutputMode, OutputModeUtils, Fork, ForkUtils, LATEST_FORK} from "scripts/forge/Config.sol";
import {Process} from "scripts/forge/libraries/Process.sol";
import {EIP1967Helper} from "scripts/forge/libraries/EIP1967Helper.sol";
import {L2CrossDomainMessenger} from "../../contracts/L2/L2CrossDomainMessenger.sol";
import {L1Block} from "../../contracts/L2/L1Block.sol";
import {L2StandardBridge} from "../../contracts/L2/L2StandardBridge.sol";
import {LightLinkMintableERC20Factory} from "../../contracts/universal/LightLinkMintableERC20Factory.sol";
import {L1CrossDomainMessenger} from "../../contracts/L1/L1CrossDomainMessenger.sol";
import {L1StandardBridge} from "../../contracts/L1/L1StandardBridge.sol";
import {Predeploys} from "../../contracts/libraries/Predeploys.sol";
import {Proxy} from "../../contracts/universal/Proxy.sol";
import {ProxyAdmin} from "../../contracts/universal/ProxyAdmin.sol";

interface IInitializable {
    function initialize(address _addr) external;
}

struct L1Dependencies {
    address payable l1CrossDomainMessengerProxy;
    address payable l1StandardBridgeProxy;
}

/// @title L2Genesis
/// @notice Generates the genesis state for the L2 network.
///         The following safety invariants are used when setting state:
///         1. `vm.getDeployedBytecode` can only be used with `vm.etch` when there are no side
///         effects in the constructor and no immutables in the bytecode.
///         2. A contract must be deployed using the `new` syntax if there are immutables in the code.
///         Any other side effects from the init code besides setting the immutables must be cleaned up afterwards.
contract L2Genesis is Deployer {
    using ForkUtils for Fork;
    using OutputModeUtils for OutputMode;

    /// @notice The address of the deployer account.
    address internal deployer;

    /// @notice Sets up the script and ensures the deployer account is used to make calls.
    function setUp() public override {
        deployer = makeAddr("deployer");
        super.setUp();
    }

    function artifactDependencies()
        internal
        view
        returns (L1Dependencies memory l1Dependencies_)
    {
        return
            L1Dependencies({
                l1CrossDomainMessengerProxy: payable(
                    cfg.l1CrossDomainMessengerProxy()
                ),
                l1StandardBridgeProxy: payable(cfg.l1StandardBridgeProxy())
            });
    }

    /// @notice The alloc object is sorted numerically by address.
    ///         Sets the precompiles, proxies, and the implementation accounts to be `vm.dumpState`
    ///         to generate a L2 genesis alloc.
    function runWithStateDump() public {
        runWithOptions(Config.outputMode(), cfg.fork(), artifactDependencies());
    }

    /// @notice Build the L2 genesis.
    function runWithOptions(
        OutputMode _mode,
        Fork _fork,
        L1Dependencies memory _l1Dependencies
    ) public {
        console.log(
            "L2Genesis: outputMode: %s, fork: %s",
            _mode.toString(),
            _fork.toString()
        );
        vm.startPrank(deployer);
        vm.chainId(cfg.l2ChainID());

        setPredeployProxies();
        setPredeployImplementations(_l1Dependencies);

        vm.stopPrank();

        if (writeForkGenesisAllocs(_fork, Fork.NONE, _mode)) {
            return;
        }
    }

    function writeForkGenesisAllocs(
        Fork _latest,
        Fork _current,
        OutputMode _mode
    ) internal returns (bool isLatest_) {
        if (
            _mode == OutputMode.ALL ||
            (_latest == _current && _mode == OutputMode.LATEST)
        ) {
            string memory suffix = string.concat("-", _current.toString());
            writeGenesisAllocs(Config.stateDumpPath(suffix));
        }
        if (_latest == _current) {
            isLatest_ = true;
        }
    }

    /// @notice Set up the accounts that correspond to the predeploys.
    ///         The Proxy bytecode should be set. All proxied predeploys should have
    ///         the 1967 admin slot set to the ProxyAdmin predeploy. All defined predeploys
    ///         should have their implementations set.
    ///         Warning: the predeploy accounts have contract code, but 0 nonce value.
    function setPredeployProxies() public {
        console.log("Setting Predeploy proxies");
        bytes memory code = vm.getDeployedCode("Proxy.sol:Proxy");
        uint160 prefix = uint160(0x420) << 148;

        console.log(
            "Setting proxy deployed bytecode for addresses in range %s through %s",
            address(prefix | uint160(0)),
            address(prefix | uint160(Predeploys.PREDEPLOY_COUNT - 1))
        );
        for (uint256 i = 0; i < Predeploys.PREDEPLOY_COUNT; i++) {
            address addr = address(prefix | uint160(i));
            if (Predeploys.notProxied(addr)) {
                console.log("Skipping proxy at %s", addr);
                continue;
            }

            vm.etch(addr, code);
            EIP1967Helper.setAdmin(addr, Predeploys.PROXY_ADMIN);

            if (Predeploys.isSupportedPredeploy(addr)) {
                address implementation = Predeploys.predeployToCodeNamespace(
                    addr
                );
                console.log(
                    "Setting proxy %s implementation: %s",
                    addr,
                    implementation
                );
                EIP1967Helper.setImplementation(addr, implementation);
            }
        }
    }

    /// @notice Sets all the implementations for the predeploy proxies. For contracts without proxies,
    ///      sets the deployed bytecode at their expected predeploy address.
    ///      LEGACY_ERC20_ETH and L1_MESSAGE_SENDER are deprecated and are not set.
    function setPredeployImplementations(
        L1Dependencies memory _l1Dependencies
    ) internal {
        console.log(
            "Setting predeploy implementations with L1 contract dependencies:"
        );
        console.log(
            "- L1CrossDomainMessengerProxy: %s",
            _l1Dependencies.l1CrossDomainMessengerProxy
        );
        console.log(
            "- L1StandardBridgeProxy: %s",
            _l1Dependencies.l1StandardBridgeProxy
        );

        // TODO: revert func signatures to original
        // setL2CrossDomainMessenger(_l1Dependencies.l1CrossDomainMessengerProxy); // 7
        // setL2StandardBridge(_l1Dependencies.l1StandardBridgeProxy); // 10
        setLightLinkMintableERC20Factory(); // 12
        setL1Block(); // 15
        setL2ToL1MessagePasser(); // 16
        setProxyAdmin(); // 18
    }

    function setProxyAdmin() public {
        // Note the ProxyAdmin implementation itself is behind a proxy that owns itself.
        address impl = _setImplementationCode(Predeploys.PROXY_ADMIN);

        bytes32 _ownerSlot = bytes32(0);

        // there is no initialize() function, so we just set the storage manually.
        vm.store(
            Predeploys.PROXY_ADMIN,
            _ownerSlot,
            bytes32(uint256(uint160(cfg.proxyAdminOwner())))
        );
        // update the proxy to not be uninitialized (although not standard initialize pattern)
        vm.store(
            impl,
            _ownerSlot,
            bytes32(uint256(uint160(cfg.proxyAdminOwner())))
        );
    }

    function setL2ToL1MessagePasser() public {
        _setImplementationCode(Predeploys.L2_TO_L1_MESSAGE_PASSER);
    }

    // /// @notice This predeploy is following the safety invariant #1.
    // function setL2CrossDomainMessenger(
    //     address payable _l1CrossDomainMessengerProxy
    // ) public {
    //     address impl = _setImplementationCode(
    //         Predeploys.L2_CROSS_DOMAIN_MESSENGER
    //     );

    //     L2CrossDomainMessenger(impl).initialize({
    //         _l1CrossDomainMessenger: L1CrossDomainMessenger(address(0))
    //     });

    //     L2CrossDomainMessenger(Predeploys.L2_CROSS_DOMAIN_MESSENGER)
    //         .initialize({
    //             _l1CrossDomainMessenger: L1CrossDomainMessenger(
    //                 _l1CrossDomainMessengerProxy
    //             )
    //         });
    // }

    // /// @notice This predeploy is following the safety invariant #1.
    // function setL2StandardBridge(
    //     address payable _l1StandardBridgeProxy
    // ) public {
    //     address impl = _setImplementationCode(Predeploys.L2_STANDARD_BRIDGE);

    //     L2StandardBridge(payable(impl)).initialize({
    //         _otherBridge: L1StandardBridge(payable(address(0)))
    //     });

    //     L2StandardBridge(payable(Predeploys.L2_STANDARD_BRIDGE)).initialize({
    //         _otherBridge: L1StandardBridge(_l1StandardBridgeProxy)
    //     });
    // }

    /// @notice This predeploy is following the safety invariant #1.
    function setLightLinkMintableERC20Factory() public {
        address impl = _setImplementationCode(
            Predeploys.LIGHTLINK_MINTABLE_ERC20_FACTORY
        );

        LightLinkMintableERC20Factory(impl).initialize({_bridge: address(0)});

        LightLinkMintableERC20Factory(
            Predeploys.LIGHTLINK_MINTABLE_ERC20_FACTORY
        ).initialize({_bridge: Predeploys.L2_STANDARD_BRIDGE});
    }

    /// @notice This predeploy is following the safety invariant #1.
    function setL1Block() public {
        _setImplementationCode(Predeploys.L1_BLOCK_ATTRIBUTES);
        // Note: L1 block attributes are set to 0.
        // Before the first user-tx the state is overwritten with actual L1 attributes.
    }

    /// @notice Sets the bytecode in state
    function _setImplementationCode(address _addr) internal returns (address) {
        string memory cname = Predeploys.getName(_addr);
        address impl = Predeploys.predeployToCodeNamespace(_addr);
        console.log("Setting %s implementation at: %s", cname, impl);
        vm.etch(impl, vm.getDeployedCode(string.concat(cname, ".sol:", cname)));
        return impl;
    }

    /// @notice Writes the genesis allocs, i.e. the state dump, to disk
    function writeGenesisAllocs(string memory _path) public {
        /// Reset so its not included state dump
        vm.etch(address(cfg), "");
        vm.etch(msg.sender, "");
        vm.resetNonce(msg.sender);
        vm.deal(msg.sender, 0);

        vm.deal(deployer, 0);
        vm.resetNonce(deployer);

        console.log("Writing state dump to: %s", _path);
        vm.dumpState(_path);
        //sortJsonByKeys(_path);
    }

    /// @notice Sorts the allocs by address
    function sortJsonByKeys(string memory _path) internal {
        string[] memory commands = new string[](3);
        commands[0] = "bash";
        commands[1] = "-c";
        commands[2] = string.concat(
            "cat <<< $(jq -S '.' ",
            _path,
            ") > ",
            _path
        );
        Process.run(commands);
    }
}
