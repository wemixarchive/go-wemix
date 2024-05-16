// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../interface/IStaking.sol";
import "../interface/IGov.sol";
import "../GovChecker.sol";

abstract contract AGov is GovChecker, IGov {
	uint public modifiedBlock;

	// For voting member
	mapping(uint256 => address) internal voters;
	mapping(address => uint256) public voterIdx;
	uint256 internal memberLength;

	// For reward member
	mapping(uint256 => address) internal rewards;
	mapping(address => uint256) public rewardIdx;

	//For staking member
	mapping(uint256 => address) internal stakers;
	mapping(address => uint256) public stakerIdx;

	//For a node duplicate check
	// mainnet value is here
	// mapping(bytes32=>bool) internal checkNodeInfo;
	mapping(bytes => bool) internal checkNodeName;
	mapping(bytes => bool) internal checkNodeEnode;
	mapping(bytes32 => bool) internal checkNodeIpPort;

	// For enode
	struct Node {
		bytes name;
		bytes enode;
		bytes ip;
		uint port;
	}

	mapping(uint256 => Node) internal nodes;
	mapping(address => uint256) internal nodeIdxFromMember;
	mapping(uint256 => address) internal nodeToMember;
	uint256 internal nodeLength;

	// For ballot
	uint256 public ballotLength;
	uint256 public voteLength;
	uint256 internal ballotInVoting;

	function isReward(address addr) public view override returns (bool) {
		return (rewardIdx[addr] != 0);
	}
	function isVoter(address addr) public view override returns (bool) {
		return (voterIdx[addr] != 0);
	}
	function isStaker(address addr) public view override returns (bool) {
		return (stakerIdx[addr] != 0);
	}
	function isMember(address addr) public view override returns (bool) {
		return (isStaker(addr) || isVoter(addr));
	}
	function getMember(uint256 idx) public view override returns (address) {
		return stakers[idx];
	}
	function getMemberLength() public view override returns (uint256) {
		return memberLength;
	}
	function getReward(uint256 idx) public view override returns (address) {
		return rewards[idx];
	}
	function getNodeIdxFromMember(address addr) public view override returns (uint256) {
		return nodeIdxFromMember[addr];
	}
	function getMemberFromNodeIdx(uint256 idx) public view override returns (address) {
		return nodeToMember[idx];
	}
	function getNodeLength() public view override returns (uint256) {
		return nodeLength;
	}
	//====NxtMeta=====/
	function getVoter(uint256 idx) public view override returns (address) {
		return voters[idx];
	}

	function getNode(uint256 idx) public view override returns (bytes memory name, bytes memory enode, bytes memory ip, uint port) {
		return (nodes[idx].name, nodes[idx].enode, nodes[idx].ip, nodes[idx].port);
	}

	function getBallotInVoting() public view override returns (uint256) {
		return ballotInVoting;
	}
}
