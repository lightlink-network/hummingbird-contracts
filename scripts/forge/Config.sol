// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Vm, VmSafe} from "../../lib/forge-std/src/Vm.sol";

/// @notice Enum representing different ways of outputting genesis allocs.
/// @custom:value NONE    No output, used in internal tests.
/// @custom:value LATEST  Output allocs only for latest fork.
/// @custom:value ALL     Output allocs for all intermediary forks.
enum OutputMode {
    NONE,
    LATEST,
    ALL
}

library OutputModeUtils {
    function toString(OutputMode _mode) internal pure returns (string memory) {
        if (_mode == OutputMode.NONE) {
            return "none";
        } else if (_mode == OutputMode.LATEST) {
            return "latest";
        } else if (_mode == OutputMode.ALL) {
            return "all";
        } else {
            return "unknown";
        }
    }
}

/// @notice Enum of forks available for selection when generating genesis allocs.
enum Fork {
    NONE
}

Fork constant LATEST_FORK = Fork.NONE;

library ForkUtils {
    function toString(Fork _fork) internal pure returns (string memory) {
        if (_fork == Fork.NONE) {
            return "none";
        } else {
            return "unknown";
        }
    }
}

/// @title Config
/// @notice Contains all env var based config. Add any new env var parsing to this file
///         to ensure that all config is in a single place.
library Config {
    /// @notice Foundry cheatcode VM.
    Vm private constant vm =
        Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    /// @notice Returns the path on the local filesystem where the deployment artifact is
    ///         written to disk after doing a deployment.
    function deploymentOutfile() internal view returns (string memory _env) {
        _env = vm.envOr(
            "DEPLOYMENT_OUTFILE",
            string.concat(
                vm.projectRoot(),
                "/deployments/",
                vm.toString(block.chainid),
                "-deploy.json"
            )
        );
    }

    /// @notice Returns the path on the local filesystem where the deploy config is
    function deployConfigPath() internal view returns (string memory _env) {
        if (vm.isContext(VmSafe.ForgeContext.TestGroup)) {
            _env = string.concat(
                vm.projectRoot(),
                "/deploy-config/hardhat.json"
            );
        } else {
            _env = vm.envOr("DEPLOY_CONFIG_PATH", string(""));
            require(
                bytes(_env).length > 0,
                "Config: must set DEPLOY_CONFIG_PATH to filesystem path of deploy config"
            );
        }
    }

    /// @notice Returns the chainid from the EVM context or the value of the CHAIN_ID env var as
    ///         an override.
    function chainID() internal view returns (uint256 _env) {
        _env = vm.envOr("CHAIN_ID", block.chainid);
    }

    /// @notice Returns the value of the env var CONTRACT_ADDRESSES_PATH which is a JSON key/value
    ///         pair of contract names and their addresses. Each key/value pair is passed to `save`
    ///         which then backs the `getAddress` function.
    function contractAddressesPath()
        internal
        view
        returns (string memory _env)
    {
        _env = vm.envOr("CONTRACT_ADDRESSES_PATH", string(""));
    }

    /// @notice Returns the path that the state dump file should be written to or read from
    ///         on the local filesystem.
    function stateDumpPath(
        string memory _suffix
    ) internal view returns (string memory _env) {
        _env = vm.envOr(
            "STATE_DUMP_PATH",
            string.concat(
                vm.projectRoot(),
                "/deployments/state-dump-",
                vm.toString(block.chainid),
                _suffix,
                ".json"
            )
        );
    }

    /// @notice Returns the OutputMode for genesis allocs generation.
    ///         It reads the mode from the environment variable OUTPUT_MODE.
    ///         If it is unset, OutputMode.ALL is returned.
    function outputMode() internal view returns (OutputMode) {
        string memory modeStr = vm.envOr("OUTPUT_MODE", string("latest"));
        bytes32 modeHash = keccak256(bytes(modeStr));
        if (modeHash == keccak256(bytes("none"))) {
            return OutputMode.NONE;
        } else if (modeHash == keccak256(bytes("latest"))) {
            return OutputMode.LATEST;
        } else if (modeHash == keccak256(bytes("all"))) {
            return OutputMode.ALL;
        } else {
            revert(string.concat("Config: unknown output mode: ", modeStr));
        }
    }

    /// @notice Returns the latest fork to use for genesis allocs generation.
    ///         It reads the fork from the environment variable FORK. If it is
    ///         unset, NONE is returned.
    ///         If set to the special value "latest", the latest fork is returned.
    function fork() internal view returns (Fork) {
        string memory forkStr = vm.envOr("FORK", string(""));
        if (bytes(forkStr).length == 0) {
            return Fork.NONE;
        }
        bytes32 forkHash = keccak256(bytes(forkStr));
        if (forkHash == keccak256(bytes("latest"))) {
            return LATEST_FORK;
        } else {
            revert(string.concat("Config: unknown fork: ", forkStr));
        }
    }
}
