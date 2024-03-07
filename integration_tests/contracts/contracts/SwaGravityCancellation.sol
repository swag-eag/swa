pragma solidity ^0.6.6;

contract SwaGravityCancellation {

    event __SwaCancelSendToEvmChain(address indexed sender, uint256 id);

    // Cancel a send to chain transaction considering if it hasnt been batched yet.
    function cancelTransaction(uint256 id) public {
        emit __SwaCancelSendToEvmChain(msg.sender, id);
    }
}
