// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import {Script} from "../../../lib/forge-std/src/Script.sol";
import {console2 as console} from "../../../lib/forge-std/src/console2.sol";
import {stdJson} from "../../../lib/forge-std/src/StdJson.sol";
import {Executables} from "scripts/forge/Executables.sol";
import {Process} from "scripts/forge/libraries/Process.sol";
import {Chains} from "scripts/forge/Chains.sol";
import {Config, Fork, ForkUtils} from "scripts/forge/Config.sol";

/// @title DeployConfig
/// @notice Represents the configuration required to deploy the system. It is expected
///         to read the file from JSON. A future improvement would be to have fallback
///         values if they are not defined in the JSON themselves.
contract DeployConfig is Script {
    using stdJson for string;
    using ForkUtils for Fork;

    /// @notice Represents an unset offset value, as opposed to 0, which denotes no-offset.
    uint256 constant NULL_OFFSET = type(uint256).max;

    string internal _json;

    uint256 public l2ChainID;
    address public proxyAdminOwner;
    address public l1CrossDomainMessengerProxy;
    address public l1StandardBridgeProxy;

    function read(string memory _path) public {
        console.log("DeployConfig: reading file %s", _path);
        try vm.readFile(_path) returns (string memory data) {
            _json = data;
        } catch {
            require(
                false,
                string.concat("Cannot find deploy config file at ", _path)
            );
        }

        l2ChainID = stdJson.readUint(_json, "$.l2ChainID");
        proxyAdminOwner = stdJson.readAddress(_json, "$.proxyAdminOwner");
        l1CrossDomainMessengerProxy = stdJson.readAddress(
            _json,
            "$.l1CrossDomainMessengerProxy"
        );
        l1StandardBridgeProxy = stdJson.readAddress(
            _json,
            "$.l1StandardBridgeProxy"
        );
    }

    function fork() public view returns (Fork fork_) {
        // let env var take precedence
        fork_ = Config.fork();
        if (fork_ == Fork.NONE) {
            // Will revert if no deploy config can be found either.
            fork_ = latestGenesisFork();
            console.log(
                "DeployConfig: using deploy config fork: %s",
                fork_.toString()
            );
        } else {
            console.log(
                "DeployConfig: using env var fork: %s",
                fork_.toString()
            );
        }
    }

    function latestGenesisFork() internal pure returns (Fork) {
        return Fork.NONE;
    }

    function _getBlockByTag(string memory _tag) internal returns (bytes32) {
        string[] memory cmd = new string[](3);
        cmd[0] = Executables.bash;
        cmd[1] = "-c";
        cmd[2] = string.concat(
            "cast block ",
            _tag,
            " --json | ",
            Executables.jq,
            " -r .hash"
        );
        bytes memory res = Process.run(cmd);
        return abi.decode(res, (bytes32));
    }

    function _readOr(
        string memory json,
        string memory key,
        bool defaultValue
    ) internal view returns (bool) {
        return vm.keyExistsJson(json, key) ? json.readBool(key) : defaultValue;
    }

    function _readOr(
        string memory json,
        string memory key,
        uint256 defaultValue
    ) internal view returns (uint256) {
        return
            (vm.keyExistsJson(json, key) && !_isNull(json, key))
                ? json.readUint(key)
                : defaultValue;
    }

    function _readOr(
        string memory json,
        string memory key,
        address defaultValue
    ) internal view returns (address) {
        return
            vm.keyExistsJson(json, key) ? json.readAddress(key) : defaultValue;
    }

    function _isNull(
        string memory json,
        string memory key
    ) internal pure returns (bool) {
        string memory value = json.readString(key);
        return (keccak256(bytes(value)) == keccak256(bytes("null")));
    }

    function _readOr(
        string memory json,
        string memory key,
        string memory defaultValue
    ) internal view returns (string memory) {
        return vm.keyExists(json, key) ? json.readString(key) : defaultValue;
    }
}
