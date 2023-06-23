//SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

contract ContractFactory {
    address[] public deployedContracts;

    event ContractCreated(address indexed newContract, address indexed creator);

    function createContract() public {
        address newContract = address(new MyContract(msg.sender));
        deployedContracts.push(newContract);
        emit ContractCreated(newContract, msg.sender);
    }

    function getDeployedContracts() public view returns (address[] memory) {
        return deployedContracts;
    }
}

contract MyContract {
    address public creator;

    constructor(address _creator) {
        creator = _creator;
    }

    // Add more functions and logic to the contract as needed
}
