pragma solidity 0.8.22;

// PingPong is a test contract to test the cross-domain messaging functionality.
contract PingPong {
    event Pong(address sender, string message);

    function ping(string memory _message) external {
        emit Pong(msg.sender, _message);
    }
}
