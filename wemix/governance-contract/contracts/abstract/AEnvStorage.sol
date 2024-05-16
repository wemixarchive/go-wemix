// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../GovChecker.sol";
import "../storage/EternalStorage.sol";

abstract contract AEnvStorage is EternalStorage, GovChecker {
	event StringVarableChanged(bytes32 indexed _name, string _value);

	event UintVarableChanged(bytes32 indexed _name, uint _value);

	event IntVarableChanged(bytes32 indexed _name, int _value);

	event AddressVarableChanged(bytes32 indexed _name, address _value);

	event Bytes32VarableChanged(bytes32 indexed _name, bytes32 _value);

	event BytesVarableChanged(bytes32 indexed _name, bytes _value);

	event VarableChanged(bytes32 indexed _name, uint256 indexed _type, string _value);

	/**
	 * @dev Allows the owner to set a value for a int variable
	 * @param h The keccak256 hash of the variable name
	 * @param v The value to be stored
	 */
	function setInt(bytes32 h, int256 v) internal {
		_setInt(h, v);
		emit IntVarableChanged(h, v);
	}

	/**
	 * @dev Allows the owner to set a value for a boolean variable
	 * @param h The keccak256 hash of the variable name
	 * @param v The value to be stored
	 */
	function setUint(bytes32 h, uint256 v) internal {
		_setUint(h, v);
		emit UintVarableChanged(h, v);
	}

	/**
	 * @dev Allows the owner to set a value for a address variable
	 * @param h The keccak256 hash of the variable name
	 * @param v The value to be stored
	 */
	function setAddress(bytes32 h, address v) internal {
		_setAddress(h, v);
		emit AddressVarableChanged(h, v);
	}

	/**
	 * @dev Allows the owner to set a value for a string variable
	 * @param h The keccak256 hash of the variable name
	 * @param v The value to be stored
	 */
	function setString(bytes32 h, string memory v) internal {
		_setString(h, v);
		emit StringVarableChanged(h, v);
	}

	/**
	 * @dev Allows the owner to set a value for a bytes variable
	 * @param h The keccak256 hash of the variable name
	 * @param v The value to be stored
	 */
	function setBytes(bytes32 h, bytes memory v) internal {
		_setBytes(h, v);
		emit BytesVarableChanged(h, v);
	}

	/**
	 * @dev Allows the owner to set a value for a bytes32 variable
	 * @param h The keccak256 hash of the variable name
	 * @param v The value to be stored
	 */
	function setBytes32(bytes32 h, bytes32 v) internal {
		_setBytes32(h, v);
		emit Bytes32VarableChanged(h, v);
	}

	// mapping(bytes32 => Variable) internal s;

	// /**
	// * @dev Add a new variable .
	// * @param _h The keccak256 hash of the variable name
	// * @param _t The type of value
	// * @param _v The value to be stored
	// */
	// function _add(bytes32 _h, uint256 _t, string _v) internal  {
	//     require(s[_h]._name == "","not found");
	//     s[_h] = Variable(_h,_t,_v);
	//     emit VarableAdded(_h,_t,_v);
	// }
	// /**
	// * @dev Update a new variable .
	// * @param _h The keccak256 hash of the variable name
	// * @param _t The type of value
	// * @param _v The value to be stored
	// */
	// function _change(bytes32 _h, uint256 _t, string _v) internal {
	//     require(s[_h]._name == _h,"not found");
	//     Variable storage v = s[_h];
	//     v._name = _h;
	//     v._type = _t;
	//     v._value = _v;
	//     emit VarableChanged(_h,_t,_v);

	// }
	// /**
	// * @dev Get the type & value stored of a string variable by the hash name
	// * @param _h The keccak256 hash of the variable name
	// */
	// function get(bytes32 _h) public view returns (uint256 varType, string varVal){
	//     //require(s[_h]._name == _h,"not found");
	//     return (s[_h]._type, s[_h]._value);
	// }
	// /**
	// * @dev Get the type stored of a string variable by the hash name
	// * @param _h The keccak256 hash of the variable name
	// */
	// function getType(bytes32 _h) public view returns (uint256){
	//     require(s[_h]._name == _h,"not found");
	//     return s[_h]._type;
	// }
	// /**
	// * @dev Get the value stored of a string variable by the hash name
	// * @param _h The keccak256 hash of the variable name
	// */
	// function getValue(bytes32 _h) public view returns (string){
	//     require(s[_h]._name == _h,"not found");
	//     return s[_h]._value;
	// }
}
