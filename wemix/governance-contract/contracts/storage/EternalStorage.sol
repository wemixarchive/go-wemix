// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title EternalStorage
 * @dev An ownable contract that can be used as a storage where the variables
 * are stored in a set of mappings indexed by hash names.
 */
contract EternalStorage {
	struct Storage {
		mapping(bytes32 => bool) _bool;
		mapping(bytes32 => int256) _int;
		mapping(bytes32 => uint256) _uint;
		mapping(bytes32 => string) _string;
		mapping(bytes32 => address) _address;
		mapping(bytes32 => bytes) _bytes;
		mapping(bytes32 => bytes32) _bytes32;
	}

	Storage internal s;

	/**
	 * @dev Get the value stored of a boolean variable by the hash name
	 * @param h The keccak256 hash of the variable name
	 */
	function getBoolean(bytes32 h) public view returns (bool) {
		return s._bool[h];
	}

	/**
	 * @dev Get the value stored of a int variable by the hash name
	 * @param h The keccak256 hash of the variable name
	 */
	function getInt(bytes32 h) public view returns (int) {
		return s._int[h];
	}

	/**
	 * @dev Get the value stored of a uint variable by the hash name
	 * @param h The keccak256 hash of the variable name
	 */
	function getUint(bytes32 h) public view returns (uint256) {
		return s._uint[h];
	}

	/**
	 * @dev Get the value stored of a address variable by the hash name
	 * @param h The keccak256 hash of the variable name
	 */
	function getAddress(bytes32 h) public view returns (address) {
		return s._address[h];
	}

	/**
	 * @dev Get the value stored of a string variable by the hash name
	 * @param h The keccak256 hash of the variable name
	 */
	function getString(bytes32 h) public view returns (string memory) {
		return s._string[h];
	}

	/**
	 * @dev Get the value stored of a bytes variable by the hash name
	 * @param h The keccak256 hash of the variable name
	 */
	function getBytes(bytes32 h) public view returns (bytes memory) {
		return s._bytes[h];
	}

	/**
	 * @dev Get the value stored of a bytes variable by the hash name
	 * @param h The keccak256 hash of the variable name
	 */
	function getBytes32(bytes32 h) public view returns (bytes32) {
		return s._bytes32[h];
	}

	/**
	 * @dev Allows the owner to set a value for a boolean variable
	 * @param h The keccak256 hash of the variable name
	 * @param v The value to be stored
	 */
	function _setBoolean(bytes32 h, bool v) internal {
		s._bool[h] = v;
	}

	/**
	 * @dev Allows the owner to set a value for a int variable
	 * @param h The keccak256 hash of the variable name
	 * @param v The value to be stored
	 */
	function _setInt(bytes32 h, int256 v) internal {
		s._int[h] = v;
	}

	/**
	 * @dev Allows the owner to set a value for a boolean variable
	 * @param h The keccak256 hash of the variable name
	 * @param v The value to be stored
	 */
	function _setUint(bytes32 h, uint256 v) internal {
		s._uint[h] = v;
	}

	/**
	 * @dev Allows the owner to set a value for a address variable
	 * @param h The keccak256 hash of the variable name
	 * @param v The value to be stored
	 */
	function _setAddress(bytes32 h, address v) internal {
		s._address[h] = v;
	}

	/**
	 * @dev Allows the owner to set a value for a string variable
	 * @param h The keccak256 hash of the variable name
	 * @param v The value to be stored
	 */
	function _setString(bytes32 h, string memory v) internal {
		s._string[h] = v;
	}

	/**
	 * @dev Allows the owner to set a value for a bytes variable
	 * @param h The keccak256 hash of the variable name
	 * @param v The value to be stored
	 */
	function _setBytes(bytes32 h, bytes memory v) internal {
		s._bytes[h] = v;
	}

	/**
	 * @dev Allows the owner to set a value for a bytes32 variable
	 * @param h The keccak256 hash of the variable name
	 * @param v The value to be stored
	 */
	function _setBytes32(bytes32 h, bytes32 v) internal {
		s._bytes32[h] = v;
	}

	/**
	 * @dev This empty reserved space is put in place to allow future versions to add new
	 * variables without shifting down storage in the inheritance chain.
	 * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
	 */
	uint256[49] private __gap;
}
