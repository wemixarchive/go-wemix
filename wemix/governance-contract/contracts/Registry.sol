// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./interface/IRegistry.sol";

/**
 * @title Registry
 * @dev Registry Contract used to set domain and permission
 *      The contracts used by the permissioned user in ShowMe references permission in this contract.
 *      Not only contract address but also general address can be set in this contract.
 *      Owner should set domain and permission.
 */
contract Registry is Ownable, IRegistry {
	// "Wemix Registry"
	uint public magic = 0x57656d6978205265676973747279;
	uint public modifiedBlock;

	mapping(bytes32 => address) public contracts;
	mapping(bytes32 => mapping(address => bool)) public permissions;

	event SetContractDomain(address setter, bytes32 indexed name, address indexed addr);
	event SetPermission(bytes32 indexed _contract, address indexed granted, bool status);

	constructor() Ownable() {}

	/*
	 * @dev Function to set contract(can be general address) domain
	 *      Only owner can use this function
	 * @param _name name
	 * @param _addr address
	 * @return A boolean that indicates if the operation was successful.
	 */
	function setContractDomain(bytes32 _name, address _addr) public onlyOwner returns (bool success) {
		require(_addr != address(0x0), "address should be non-zero");
		contracts[_name] = _addr;
		modifiedBlock = block.number;

		emit SetContractDomain(msg.sender, _name, _addr);
		return true;
	}

	/*
	 * @dev Function to get contract(can be general address) address
	 *      Anyone can use this function
	 * @param _name _name
	 * @return An address of the _name
	 */
	function getContractAddress(bytes32 _name) public view override returns (address addr) {
		require(contracts[_name] != address(0x0), "address should be non-zero");
		return contracts[_name];
	}

	/*
	 * @dev Function to set permission on contract
	 *      using modifier 'permissioned' references mapping variable 'permissions'
	 *      Only owner can use this function
	 * @param _contract contract name
	 * @param _granted granted address
	 * @param _status true = can use, false = cannot use. default is false
	 * @return A boolean that indicates if the operation was successful.
	 */
	function setPermission(bytes32 _contract, address _granted, bool _status) public onlyOwner returns (bool success) {
		require(_granted != address(0x0), "address should be non-zero");
		permissions[_contract][_granted] = _status;
		modifiedBlock = block.number;

		emit SetPermission(_contract, _granted, _status);
		return true;
	}

	/*
	 * @dev Function to get permission on contract
	 *      using modifier 'permissioned' references mapping variable 'permissions'
	 * @param _contract contract name
	 * @param _granted granted address
	 * @return permission result
	 */
	function getPermission(bytes32 _contract, address _granted) public view returns (bool found) {
		return permissions[_contract][_granted];
	}
}
