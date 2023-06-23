//SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

contract SelfDestructContract {
    address public owner;

    constructor() {
        owner = msg.sender;
    }

    function destroyContract() public {
        require(msg.sender == owner, "Only the owner can destroy the contract.");
        selfdestruct(payable(owner));
    }
}
