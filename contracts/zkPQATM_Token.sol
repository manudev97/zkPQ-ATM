// SPDX-License-Identifier: MIT
// Compatible with OpenZeppelin Contracts ^5.0.0
pragma solidity >=0.8.20 <0.9.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Pausable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract zkPQATM_Token is ERC20, ERC20Pausable, Ownable {
    string public greeting = "Deploy token ZKATM";
    address public atmAddress;

    constructor(address initialOwner)
        ERC20("zkPQATM_Token", "ZKATM")
        Ownable(initialOwner)
    {
        _mint(msg.sender, 1000000 * 10 ** decimals());
    }
    
    function updateATMAddress(address _atm) external onlyOwner {
        atmAddress = _atm;
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function mint(address to, uint256 amount) public onlyATM {
        _mint(to, amount);
    }

    // The following functions are overrides required by Solidity.

    function _update(address from, address to, uint256 value)
        internal
        override(ERC20, ERC20Pausable)
    {
        super._update(from, to, value);
    }

    modifier onlyATM(){
        require(msg.sender == atmAddress, "only_atm");
        _;
    }
}
