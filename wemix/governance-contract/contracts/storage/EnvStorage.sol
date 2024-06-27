// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";

contract EnvStorage is ERC1967Proxy, IBeacon {
	constructor(address _implementation) ERC1967Proxy(_implementation, "") {}

	function implementation() external view override returns (address) {
		return _implementation();
	}
}
