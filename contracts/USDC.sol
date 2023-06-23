//SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

import "github.com/OpenZeppelin/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";
import "github.com/OpenZeppelin/openzeppelin-contracts/contracts/token/ERC20/ERC20Detailed.sol";
import "github.com/OpenZeppelin/openzeppelin-contracts/contracts/token/ERC20/ERC20Mintable.sol";

contract USDCKovan is ERC20, ERC20Detailed, ERC20Mintable {
    
    string NAME = "USDCKovan";
    string SYMBOL = "USDC";
    uint8  DECIMALS = 6;

    constructor() 
        public 
        ERC20Detailed(NAME, SYMBOL, DECIMALS)
        ERC20Mintable()
    {
    }
}