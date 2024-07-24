// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Script} from "../../../lib/forge-std/src/Script.sol";
import {console} from "../../../lib/forge-std/src/console.sol";

import {Config} from "scripts/forge/Config.sol";
import {DeployConfig} from "scripts/forge/deploy/DeployConfig.s.sol";
import {Executables} from "scripts/forge/Executables.sol";
import {Artifacts} from "scripts/forge/Artifacts.s.sol";

/// @title Deployer
/// @author tynes
/// @notice A contract that can make deploying and interacting with deployments easy.
abstract contract Deployer is Script, Artifacts {
    DeployConfig public constant cfg =
        DeployConfig(
            address(
                uint160(
                    uint256(keccak256(abi.encode("lightlink.deployconfig")))
                )
            )
        );

    /// @notice Sets up the artifacts contract.
    function setUp() public virtual override {
        Artifacts.setUp();

        console.log("Commit hash: %s", Executables.gitCommitHash());

        vm.etch(
            address(cfg),
            vm.getDeployedCode("DeployConfig.s.sol:DeployConfig")
        );
        vm.label(address(cfg), "DeployConfig");
        vm.allowCheatcodes(address(cfg));
        cfg.read(Config.deployConfigPath());
    }
}
