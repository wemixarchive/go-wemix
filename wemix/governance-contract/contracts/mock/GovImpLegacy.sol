// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// =============================================================================
// TEST-ONLY MOCK — DO NOT DEPLOY TO PRODUCTION.
//
// GovImpLegacy is a byte-for-byte copy of GovImp.sol *before* the CertiK W1G
// fixes (W1G-01/02/03/04) were applied. It exists so the Go test suite can:
//   1. deploy the UNFIXED implementation behind the real Gov proxy,
//   2. build up a realistic on-chain ledger (multiple stakers, some with a
//      self-separated staker/voter/reward), and
//   3. reproduce each vulnerability on that ledger (the "red" proof),
//   4. then upgrade the proxy to the fixed GovImp and show the same attack is
//      now blocked while normal operation keeps working (the "green" proof).
//
// The storage layout is IDENTICAL to GovImp (same AGov base, same trailing
// state vars + __gap), which is what makes the proxy upgrade state-preserving.
// Only the four fix sites differ from GovImp — they are marked "LEGACY:" below.
// =============================================================================

import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "../abstract/BallotEnums.sol";
import "../abstract/EnvConstants.sol";
import "../abstract/AGov.sol";

import "../interface/IBallotStorage.sol";
import "../interface/IEnvStorage.sol";
import "../interface/IStaking.sol";

import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";

contract GovImpLegacy is AGov, ReentrancyGuardUpgradeable, BallotEnums, EnvConstants, UUPSUpgradeable {
    enum VariableTypes {
        Invalid,
        Int,
        Uint,
        Uint2,
        Uint3,
        Uint4,
        Address,
        Bytes32,
        Bytes,
        String
    }

    address constant ZERO = address(0);

    event MemberAdded(address indexed addr, address indexed voter);
    event MemberRemoved(address indexed addr, address indexed voter);
    event MemberChanged(address indexed oldAddr, address indexed newAddr, address indexed newVoter);
    event EnvChanged(bytes32 envName, uint256 envType, bytes envVal);
    event MemberUpdated(address indexed addr, address indexed voter);
    event NotApplicable(uint256 indexed ballotId, string reason);

    event SetProposalTimePeriod(uint256 newPeriod);
    event GovDataMigrated(address indexed from);

    struct MemberInfo {
        address staker;
        address voter; // voter
        address reward;
        bytes name;
        bytes enode;
        bytes ip;
        uint256 port;
        uint256 lockAmount;
        bytes memo;
        uint256 duration;
    }

    modifier checkLockedAmount() {
        address staker = getStakerAddr(_msgSender());
        require(lockedBalanceOf(staker) <= getMaxStaking() && lockedBalanceOf(staker) >= getMinStaking(), "Invalid staking balance");
        _;
    }

    modifier checkTimePeriod() {
        address staker = getStakerAddr(_msgSender());
        require((block.timestamp - lastAddProposalTime[staker]) >= proposal_time_period, "Cannot add proposal too early");
        _;
        lastAddProposalTime[staker] = block.timestamp;
    }

    modifier checkMemberInfo(MemberInfo memory info) {
        require(info.voter != ZERO, "Invalid voter");
        require(info.name.length > 0, "Invalid node name");
        require(info.ip.length > 0, "Invalid node IP");
        require(info.port > 0, "Invalid node port");
        require(info.enode.length > 0, "Invalid node enode");
        require(info.memo.length > 0, "Invalid memo");
        require(info.duration > 0, "Invalid duration");
        require(info.lockAmount >= getMinStaking() && info.lockAmount <= getMaxStaking(), "Invalid lock Amount");
        _;
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function init(address registry, uint256 lockAmount, bytes memory name, bytes memory enode, bytes memory ip, uint port) public initializer {
        __ReentrancyGuard_init();
        __Ownable_init();
        setRegistry(registry);

        require(lockAmount >= getMinStaking() && getMaxStaking() >= lockAmount, "Invalid lock amount");

        // Lock
        IStaking staking = IStaking(getStakingAddress());
        require(staking.availableBalanceOf(msg.sender) >= lockAmount, "Insufficient staking");
        staking.lock(msg.sender, lockAmount);

        // Add voting member
        memberLength = 1;
        voters[memberLength] = msg.sender;
        voterIdx[msg.sender] = memberLength;

        // Add reward member
        rewards[memberLength] = msg.sender;
        rewardIdx[msg.sender] = memberLength;

        stakers[memberLength] = msg.sender;
        stakerIdx[msg.sender] = memberLength;

        // Add node
        nodeLength = 1;
        Node storage node = nodes[nodeLength];
        node.name = name;
        node.enode = enode;
        node.ip = ip;
        node.port = port;
        checkNodeName[name] = true;
        checkNodeEnode[enode] = true;

        checkNodeIpPort[keccak256(abi.encodePacked(ip, port))] = true;

        nodeIdxFromMember[msg.sender] = nodeLength;
        nodeToMember[nodeLength] = msg.sender;

        modifiedBlock = block.number;
        emit MemberAdded(msg.sender, msg.sender);
    }

    function initOnce(address registry, uint256 lockAmount, bytes memory data) public initializer {
        __ReentrancyGuard_init();
        __Ownable_init();
        setRegistry(registry);

        modifiedBlock = block.number;

        // Lock
        IStaking staking = IStaking(getStakingAddress());

        require(lockAmount >= getMinStaking() && getMaxStaking() >= lockAmount, "Invalid lock amount");

        address staker;
        address voter;
        address reward;
        bytes memory name;
        bytes memory enode;
        bytes memory ip;
        uint port;
        uint idx = 0;

        uint ix;
        uint eix;
        assembly {
            ix := add(data, 0x20)
        }
        eix = ix + data.length;
        while (ix < eix) {
            assembly {
                staker := mload(ix)
            }
            ix += 0x20;
            require(ix < eix);

            assembly {
                voter := mload(ix)
            }
            ix += 0x20;
            require(ix < eix);

            assembly {
                reward := mload(ix)
            }
            ix += 0x20;
            require(ix < eix);

            assembly {
                name := ix
            }
            ix += 0x20 + name.length;
            require(ix < eix);

            assembly {
                enode := ix
            }
            ix += 0x20 + enode.length;
            require(ix < eix);

            assembly {
                ip := ix
            }
            ix += 0x20 + ip.length;
            require(ix < eix);

            assembly {
                port := mload(ix)
            }
            ix += 0x20;

            idx += 1;
            require(!isMember(staker) && !isMember(voter) && !isReward(reward), "Already member");
            voters[idx] = voter;
            voterIdx[voter] = idx;
            rewards[idx] = reward;
            rewardIdx[reward] = idx;
            stakers[idx] = staker;
            stakerIdx[staker] = idx;
            emit MemberAdded(staker, voter); // staker, voter

            require(staking.availableBalanceOf(staker) >= lockAmount, "Insufficient staking");

            require(checkNodeInfoAdd(name, enode, ip, port), "Duplicated node info");

            lock(staker, lockAmount);

            Node storage node = nodes[idx];
            node.name = name;
            node.enode = enode;
            node.ip = ip;
            node.port = port;
            checkNodeName[name] = true;
            checkNodeEnode[enode] = true;
            checkNodeIpPort[keccak256(abi.encodePacked(ip, port))] = true;

            nodeToMember[idx] = staker;
            nodeIdxFromMember[staker] = idx;
        }
        memberLength = idx;
        nodeLength = idx;
    }

    function addProposalToAddMember(
        MemberInfo memory info
    ) external onlyGovMem checkTimePeriod checkLockedAmount checkMemberInfo(info) returns (uint256 ballotIdx) {
        require(!isMember(info.staker) && !isReward(info.staker), "Already member");
        require(info.staker == info.voter && info.staker == info.reward, "Staker is not voter");
        require(checkNodeInfoAdd(info.name, info.enode, info.ip, info.port), "Duplicated node info");
        ballotIdx = ballotLength + 1;
        createBallotForMember(
            ballotIdx, // ballot id
            uint256(BallotTypes.MemberAdd), // ballot type
            msg.sender, // creator
            ZERO, // old staker address
            info
        );
        updateBallotLock(ballotIdx, info.lockAmount);
        updateBallotMemo(ballotIdx, info.memo);
        ballotLength = ballotIdx;
    }

    function addProposalToRemoveMember(
        address staker,
        uint256 lockAmount,
        bytes memory memo,
        uint256 duration,
        uint256 unlockAmount,
        uint256 slashing
    ) external onlyGovMem checkTimePeriod checkLockedAmount returns (uint256 ballotIdx) {
        require(staker != ZERO, "Invalid address");
        // LEGACY (pre W1G-04): isMember admits voter-only addresses (stakerIdx==0).
        require(isMember(staker), "Non-member");
        require(getMemberLength() > 1, "Cannot remove a sole member");
        require(lockedBalanceOf(staker) >= lockAmount, "Insufficient balance that can be unlocked.");
        ballotIdx = ballotLength + 1;

        MemberInfo memory info = MemberInfo(
            ZERO, // new staker address
            ZERO,
            ZERO,
            new bytes(0), // new name
            new bytes(0), // new enode
            new bytes(0), // new ip
            0, // new port
            lockAmount,
            memo,
            duration
        );
        createBallotForMember(
            ballotIdx, // ballot id
            uint256(BallotTypes.MemberRemoval), // ballot type
            msg.sender,
            staker,
            info
        );
        updateBallotLock(ballotIdx, lockAmount);
        updateBallotMemo(ballotIdx, memo);
        createBallotForExit(ballotIdx, unlockAmount, slashing);
        ballotLength = ballotIdx;
    }

    function addProposalToChangeMember(
        MemberInfo memory newInfo,
        address oldStaker,
        uint256 unlockAmount,
        uint256 slashing
    ) external onlyGovMem checkTimePeriod checkLockedAmount checkMemberInfo(newInfo) returns (uint256 ballotIdx) {
        require(oldStaker != ZERO, "Invalid old Address");
        // LEGACY (pre W1G-04): isMember admits voter-only addresses (stakerIdx==0).
        require(isMember(oldStaker), "Non-member");

        require(
            (voters[stakerIdx[oldStaker]] == newInfo.voter ||
                newInfo.voter == oldStaker ||
                ((!isMember(newInfo.voter)) && !isReward(newInfo.voter))) &&
                (rewards[stakerIdx[oldStaker]] == newInfo.reward ||
                    newInfo.reward == oldStaker ||
                    ((!isMember(newInfo.reward)) && !isReward(newInfo.reward))),
            "Already a member"
        );
        // For exit
        if (msg.sender == oldStaker && oldStaker == newInfo.staker) {
            require(unlockAmount == 0 && slashing == 0, "Invalid proposal");
        } else if (oldStaker != newInfo.staker /* && msg.sender != oldStaker */) {
            require(unlockAmount + slashing <= getMinStaking(), "Invalid amount: (unlockAmount + slashing) must be equal or low to minStaking");
        }

        ballotIdx = ballotLength + 1;
        createBallotForMember(
            ballotIdx, // ballot id
            uint256(BallotTypes.MemberChange), // ballot type
            msg.sender, // creator
            oldStaker, // old staker address
            newInfo
        );
        updateBallotLock(ballotIdx, newInfo.lockAmount);
        updateBallotMemo(ballotIdx, newInfo.memo);
        createBallotForExit(ballotIdx, unlockAmount, slashing);
        ballotLength = ballotIdx;
        if (msg.sender == oldStaker && oldStaker == newInfo.staker) {
            (, , uint256 duration) = getBallotPeriod(ballotIdx);
            startBallot(ballotIdx, block.timestamp, block.timestamp + duration);
            finalizeVote(ballotIdx, uint256(BallotTypes.MemberChange), true, true);
        }
    }

    function addProposalToChangeGov(
        address newGovAddr,
        bytes memory memo,
        uint256 duration
    ) external onlyGovMem checkTimePeriod checkLockedAmount returns (uint256 ballotIdx) {
        require(newGovAddr != ZERO, "Implementation cannot be zero");
        require(newGovAddr != _getImplementation(), "Same contract address");
        try IERC1822Proxiable(newGovAddr).proxiableUUID() returns (bytes32 slot) {
            require(slot == _IMPLEMENTATION_SLOT, "ERC1967Upgrade: unsupported proxiableUUID");
        } catch {
            revert("ERC1967Upgrade: new implementation is not UUPS");
        }
        ballotIdx = ballotLength + 1;
        IBallotStorage(getBallotStorageAddress()).createBallotForAddress(
            ballotLength + 1, // ballot id
            uint256(BallotTypes.GovernanceChange), // ballot type
            duration,
            msg.sender, // creator
            newGovAddr // new governance address
        );
        updateBallotMemo(ballotIdx, memo);
        ballotLength = ballotIdx;
    }

    function addProposalToChangeEnv(
        bytes32 envName,
        uint256 envType,
        bytes memory envVal,
        bytes memory memo,
        uint256 duration
    ) external onlyGovMem checkTimePeriod checkLockedAmount returns (uint256 ballotIdx) {
        require(uint256(VariableTypes.Int) <= envType && envType <= uint256(VariableTypes.String), "Invalid type");
        require(checkVariableCondition(envName, envVal), "Invalid value");

        ballotIdx = ballotLength + 1;
        IBallotStorage(getBallotStorageAddress()).createBallotForVariable(
            ballotIdx, // ballot id
            uint256(BallotTypes.EnvValChange), // ballot type
            duration,
            msg.sender, // creator
            envName, // env name
            envType, // env type
            envVal // env value
        );
        updateBallotMemo(ballotIdx, memo);
        ballotLength = ballotIdx;
    }

    function vote(uint256 ballotIdx, bool approval) external nonReentrant onlyGovMem checkLockedAmount {
        require(checkUnfinalized(), "Expired");

        uint256 ballotType = checkVotable(ballotIdx);
        createVote(ballotIdx, approval);
        (, uint256 accept, uint256 reject) = getBallotVotingInfo(ballotIdx);
        uint256 threshold = getThreshold();
        if (accept >= threshold || reject >= threshold || (accept + reject) == 10000) {
            finalizeVote(ballotIdx, ballotType, accept > reject, false);
        }
    }

    function getMinStaking() public view returns (uint256) {
        return IEnvStorage(getEnvStorageAddress()).getStakingMin();
    }

    function getMaxStaking() public view returns (uint256) {
        return IEnvStorage(getEnvStorageAddress()).getStakingMax();
    }

    function getMinVotingDuration() public view returns (uint256) {
        return IEnvStorage(getEnvStorageAddress()).getBallotDurationMin();
    }

    function getMaxVotingDuration() public view returns (uint256) {
        return IEnvStorage(getEnvStorageAddress()).getBallotDurationMax();
    }

    function getThreshold() public pure returns (uint256) {
        return 5001;
    } // 50.01% from 5001 of 10000

    function checkUnfinalized() public view returns (bool) {
        if (ballotInVoting != 0) {
            (, uint256 state, ) = getBallotState(ballotInVoting);
            (, uint256 endTime, ) = getBallotPeriod(ballotInVoting);
            if (state == uint256(BallotStates.InProgress)) {
                if (endTime < block.timestamp) return false;
            }
        }
        return true;
    }

    function finalizeEndedVote() public onlyGovMem {
        require(!checkUnfinalized(), "Voting is not ended");
        finalizeBallot(ballotInVoting, uint256(BallotStates.Rejected));
        ballotInVoting = 0;
    }

    function checkVotable(uint256 ballotIdx) private returns (uint256) {
        (uint256 ballotType, uint256 state, ) = getBallotState(ballotIdx);
        if (state == uint256(BallotStates.Ready)) {
            require(ballotInVoting == 0, "Now in voting with different ballot");
            (, , uint256 duration) = getBallotPeriod(ballotIdx);
            if (duration < getMinVotingDuration()) {
                startBallot(ballotIdx, block.timestamp, block.timestamp + getMinVotingDuration());
            } else if (getMaxVotingDuration() < duration) {
                startBallot(ballotIdx, block.timestamp, block.timestamp + getMaxVotingDuration());
            } else {
                startBallot(ballotIdx, block.timestamp, block.timestamp + duration);
            }
            ballotInVoting = ballotIdx;
        } else if (state == uint256(BallotStates.InProgress)) {
            require(ballotIdx == ballotInVoting, "Now in voting with different ballot");
        } else {
            revert("Expired");
        }
        return ballotType;
    }

    function createVote(uint256 ballotIdx, bool approval) private {
        uint256 voteIdx = voteLength + 1;
        address staker = getStakerAddr(msg.sender);
        uint256 weight = 10000 / getMemberLength();
        uint256 decision = approval ? uint256(DecisionTypes.Accept) : uint256(DecisionTypes.Reject);
        IBallotStorage(getBallotStorageAddress()).createVote(voteIdx, ballotIdx, staker, decision, weight);
        voteLength = voteIdx;
    }

    function finalizeVote(uint256 ballotIdx, uint256 ballotType, bool isAccepted, bool self) private {
        uint256 ballotState = uint256(BallotStates.Rejected);
        if (isAccepted) {
            ballotState = uint256(BallotStates.Accepted);

            if (ballotType == uint256(BallotTypes.MemberAdd)) {
                if (!addMember(ballotIdx)) {
                    ballotState = uint256(BallotStates.Rejected);
                }
            } else if (ballotType == uint256(BallotTypes.MemberRemoval)) {
                removeMember(ballotIdx);
            } else if (ballotType == uint256(BallotTypes.MemberChange)) {
                if (!changeMember(ballotIdx, self)) {
                    ballotState = uint256(BallotStates.Rejected);
                }
            } else if (ballotType == uint256(BallotTypes.GovernanceChange)) {
                changeGov(ballotIdx);
            } else if (ballotType == uint256(BallotTypes.EnvValChange)) {
                applyEnv(ballotIdx);
            }
        }
        finalizeBallot(ballotIdx, ballotState);
        if (!self) ballotInVoting = 0;
    }

    function fromValidBallot(uint256 ballotIdx, uint256 targetType) private view {
        (uint256 ballotType, uint256 state, ) = getBallotState(ballotIdx);
        require(ballotType == targetType, "Invalid voting type");
        require(state == uint(BallotStates.InProgress), "Invalid voting state");
        (, uint256 accept, uint256 reject) = getBallotVotingInfo(ballotIdx);
        require(accept >= getThreshold() || reject >= getThreshold() || (accept + reject) == 10000, "Not yet finalized");
    }

    function addMember(uint256 ballotIdx) private returns (bool) {
        fromValidBallot(ballotIdx, uint256(BallotTypes.MemberAdd));

        (
            ,
            address newStaker,
            address newVoter,
            address newReward,
            bytes memory name,
            bytes memory enode,
            bytes memory ip,
            uint port,
            uint256 lockAmount
        ) = getBallotMember(ballotIdx);
        if (isMember(newStaker)) {
            emit NotApplicable(ballotIdx, "Already a member");
            return false;
        }
        if (isReward(newReward)) {
            emit NotApplicable(ballotIdx, "Already a rewarder");
            return false;
        }

        // Lock
        if (lockAmount < getMinStaking() || getMaxStaking() < lockAmount) {
            emit NotApplicable(ballotIdx, "Invalid lock amount");
            return false;
        }

        if (availableBalanceOf(newStaker) < lockAmount) {
            emit NotApplicable(ballotIdx, "Insufficient balance that can be locked");
            return false;
        }

        if (newStaker != newVoter && newStaker != newReward) {
            emit NotApplicable(ballotIdx, "Invalid member address");
            return false;
        }

        // LEGACY (pre W1G-01): NO execution-time node-uniqueness re-validation here.

        lock(newStaker, lockAmount);

        // Add voting and reward member
        uint256 nMemIdx = memberLength + 1;
        voters[nMemIdx] = newVoter;
        voterIdx[newVoter] = nMemIdx;
        rewards[nMemIdx] = newReward;
        rewardIdx[newReward] = nMemIdx;
        stakers[nMemIdx] = newStaker;
        stakerIdx[newStaker] = nMemIdx;

        // Add node
        uint256 nNodeIdx = nodeLength + 1;
        Node storage node = nodes[nNodeIdx];

        node.name = name;
        node.enode = enode;
        node.ip = ip;
        node.port = port;
        checkNodeName[name] = true;
        checkNodeEnode[enode] = true;
        checkNodeIpPort[keccak256(abi.encodePacked(ip, port))] = true;

        nodeToMember[nNodeIdx] = newStaker;
        nodeIdxFromMember[newStaker] = nNodeIdx;
        memberLength = nMemIdx;
        nodeLength = nNodeIdx;
        modifiedBlock = block.number;
        emit MemberAdded(newStaker, newVoter);
        return true;
    }

    function removeMember(uint256 ballotIdx) private {
        fromValidBallot(ballotIdx, uint256(BallotTypes.MemberRemoval));

        (address oldStaker, , , , , , , , ) = getBallotMember(ballotIdx);
        if (!isMember(oldStaker)) {
            emit NotApplicable(ballotIdx, "Not already a member");
            return; // Non-member. it is abnormal case, but passed
        }
        // LEGACY (pre W1G-04 / W1G-02): no voter-only guard, no >1 re-check here.

        // Remove voting and reward member
        uint256 removeStakerIdx = stakerIdx[oldStaker];
        address oldVoter = voters[removeStakerIdx];
        address oldReward = rewards[removeStakerIdx];

        if (removeStakerIdx != memberLength) {
            address endStaker = stakers[memberLength];
            stakers[removeStakerIdx] = endStaker;
            stakerIdx[endStaker] = removeStakerIdx;

            uint256 removeRewardIdx = rewardIdx[oldReward];
            require(removeRewardIdx != 0, "Invalid reward index");
            address endReward = rewards[memberLength];
            rewards[removeRewardIdx] = endReward;
            rewardIdx[endReward] = removeRewardIdx;

            uint256 removeVoterIdx = voterIdx[oldVoter];
            require(removeVoterIdx != 0, "Invalid voter index");
            address endVoter = voters[memberLength];
            voters[removeVoterIdx] = endVoter;
            voterIdx[endVoter] = removeVoterIdx;
        }
        stakers[memberLength] = ZERO;
        stakerIdx[oldStaker] = 0;

        rewards[memberLength] = ZERO;
        rewardIdx[oldReward] = 0;

        voters[memberLength] = ZERO;
        voterIdx[oldVoter] = 0;

        memberLength = memberLength - 1;

        // Remove node
        uint256 removeNodeIdx = nodeIdxFromMember[oldStaker];
        require(removeNodeIdx != 0, "Invalid node index");
        Node storage node = nodes[removeNodeIdx];
        checkNodeEnode[node.enode] = false;
        checkNodeName[node.name] = false;
        checkNodeIpPort[keccak256(abi.encodePacked(node.ip, node.port))] = false;

        if (removeNodeIdx != nodeLength) {
            address endMember = nodeToMember[nodeLength];

            node.name = nodes[nodeLength].name;
            node.enode = nodes[nodeLength].enode;
            node.ip = nodes[nodeLength].ip;
            node.port = nodes[nodeLength].port;

            nodeToMember[removeNodeIdx] = endMember;
            nodeIdxFromMember[endMember] = removeNodeIdx;
        }
        nodeToMember[nodeLength] = ZERO;
        nodeIdxFromMember[oldStaker] = 0;
        delete nodes[nodeLength];
        nodeLength = nodeLength - 1;
        modifiedBlock = block.number;
        transferLockedAndUnlock(ballotIdx, oldStaker);

        emit MemberRemoved(oldStaker, oldVoter);
    }

    function checkChangeMember(
        uint256 ballotIdx,
        bool self,
        address oldStaker,
        address newStaker,
        address newVoter,
        address newReward,
        bytes memory name,
        bytes memory enode,
        bytes memory ip,
        uint256 port,
        uint256 lockAmount
    ) private returns (bool) {
        if (!self) {
            fromValidBallot(ballotIdx, uint256(BallotTypes.MemberChange));
        }

        if (!isMember(oldStaker)) {
            emit NotApplicable(ballotIdx, "Old address is not a member");
            return false;
        }
        // LEGACY (pre W1G-04): no voter-only (stakerIdx==0) guard here.

        //old staker
        uint256 memberIdx = stakerIdx[oldStaker];
        if (oldStaker != newStaker) {
            if (isMember(newStaker)) {
                emit NotApplicable(ballotIdx, "new address is already a member");
                return false;
            }
            if (newStaker != newVoter && newStaker != newReward) {
                emit NotApplicable(ballotIdx, "Invalid voter address");
                return false;
            }
            // Lock
            if (lockAmount < getMinStaking() || getMaxStaking() < lockAmount) {
                emit NotApplicable(ballotIdx, "Invalid lock amount");
                return false;
            }
            if (availableBalanceOf(newStaker) < lockAmount) {
                emit NotApplicable(ballotIdx, "Insufficient balance that can be locked");
                return false;
            }
        }
        // Change node
        uint256 nodeIdx = nodeIdxFromMember[oldStaker];
        {
            Node memory node = nodes[nodeIdx];

            if (!checkNodeInfoChange(name, enode, ip, port, node)) {
                emit NotApplicable(ballotIdx, "Duplicated node info");
                return false;
            }
        }

        {
            address oldReward = rewards[memberIdx];
            if ((oldStaker != newReward) && (oldReward != newReward) && (isMember(newReward) || isReward(newReward))) {
                emit NotApplicable(ballotIdx, "Invalid reward address");
                return false;
            }
        }
        {
            address oldVoter = voters[memberIdx];
            if ((oldStaker != newVoter) && (oldVoter != newVoter) && (isMember(newVoter) || isReward(newVoter))) {
                emit NotApplicable(ballotIdx, "Invalid voters address");
                return false;
            }
        }
        return true;
    }

    function changeMember(uint256 ballotIdx, bool self) private returns (bool) {
        if (!self) {
            fromValidBallot(ballotIdx, uint256(BallotTypes.MemberChange));
        }

        (
            address oldStaker,
            address newStaker,
            address newVoter,
            address newReward,
            bytes memory name,
            bytes memory enode,
            bytes memory ip,
            uint port,
            uint256 lockAmount
        ) = getBallotMember(ballotIdx);
        if (!isMember(oldStaker)) {
            emit NotApplicable(ballotIdx, "Old address is not a member");
            return false;
        }

        if (!checkChangeMember(ballotIdx, self, oldStaker, newStaker, newVoter, newReward, name, enode, ip, port, lockAmount)) return false;

        //old staker
        uint256 memberIdx = stakerIdx[oldStaker];
        if (oldStaker != newStaker) {
            stakers[memberIdx] = newStaker;
            stakerIdx[newStaker] = memberIdx;
            stakerIdx[oldStaker] = 0;

            lock(newStaker, lockAmount);
        }
        // Change node
        uint256 nodeIdx = nodeIdxFromMember[oldStaker];
        {
            Node storage node = nodes[nodeIdx];

            checkNodeName[node.name] = false;
            checkNodeEnode[node.enode] = false;
            checkNodeIpPort[keccak256(abi.encodePacked(node.ip, node.port))] = false;

            node.name = name;
            node.enode = enode;
            node.ip = ip;
            node.port = port;
            modifiedBlock = block.number;
            checkNodeName[name] = true;
            checkNodeEnode[enode] = true;
            checkNodeIpPort[keccak256(abi.encodePacked(ip, port))] = true;
        }

        {
            address oldReward = rewards[memberIdx];
            if (oldReward != newReward) {
                rewards[memberIdx] = newReward;
                rewardIdx[newReward] = memberIdx;
                rewardIdx[oldReward] = 0;
            }
        }
        {
            address oldVoter = voters[memberIdx];
            if (oldVoter != newVoter) {
                voters[memberIdx] = newVoter;
                voterIdx[newVoter] = memberIdx;
                voterIdx[oldVoter] = 0;
            }
        }

        if (oldStaker != newStaker) {
            nodeToMember[nodeIdx] = newStaker;
            nodeIdxFromMember[newStaker] = nodeIdx;
            nodeIdxFromMember[oldStaker] = 0;

            transferLockedAndUnlock(ballotIdx, oldStaker);

            emit MemberChanged(oldStaker, newStaker, newVoter);
        } else {
            emit MemberUpdated(oldStaker, newStaker);
        }
        return true;
    }

    function changeGov(uint256 ballotIdx) private {
        fromValidBallot(ballotIdx, uint256(BallotTypes.GovernanceChange));

        address newImp = IBallotStorage(getBallotStorageAddress()).getBallotAddress(ballotIdx);
        if (newImp != ZERO) {
            _authorizeUpgrade(newImp);
            _upgradeToAndCallUUPS(newImp, new bytes(0), false);
            modifiedBlock = block.number;
        }
    }

    function applyEnv(uint256 ballotIdx) private {
        fromValidBallot(ballotIdx, uint256(BallotTypes.EnvValChange));

        (bytes32 envKey, uint256 envType, bytes memory envVal) = IBallotStorage(getBallotStorageAddress()).getBallotVariable(ballotIdx);

        IEnvStorage envStorage = IEnvStorage(getEnvStorageAddress());
        envStorage.setVariable(envKey, envVal);
        modifiedBlock = block.number;

        emit EnvChanged(envKey, envType, envVal);
    }

    function createBallotForMember(uint256 id, uint256 bType, address creator, address oAddr, MemberInfo memory info) private {
        IBallotStorage(getBallotStorageAddress()).createBallotForMember(
            id, // ballot id
            bType, // ballot type
            info.duration,
            creator, // creator
            oAddr, // old member address
            info.staker, // new member address
            info.voter, // old staker address
            info.reward, // new staker address
            info.name, // new name
            info.enode, // new enode
            info.ip, // new ip
            info.port // new port
        );
    }

    function updateBallotLock(uint256 id, uint256 amount) private {
        IBallotStorage(getBallotStorageAddress()).updateBallotMemberLockAmount(id, amount);
    }

    function updateBallotMemo(uint256 id, bytes memory memo) private {
        IBallotStorage(getBallotStorageAddress()).updateBallotMemo(id, memo);
    }

    function createBallotForExit(uint256 id, uint256 unlockAmount, uint256 slashing) private {
        IBallotStorage(getBallotStorageAddress()).createBallotForExit(id, unlockAmount, slashing);
    }

    function startBallot(uint256 id, uint256 s, uint256 e) private {
        IBallotStorage(getBallotStorageAddress()).startBallot(id, s, e);
    }

    function finalizeBallot(uint256 id, uint256 state) private {
        IBallotStorage(getBallotStorageAddress()).finalizeBallot(id, state);
    }

    function getBallotState(uint256 id) private view returns (uint256, uint256, bool) {
        return IBallotStorage(getBallotStorageAddress()).getBallotState(id);
    }

    function getBallotPeriod(uint256 id) private view returns (uint256, uint256, uint256) {
        return IBallotStorage(getBallotStorageAddress()).getBallotPeriod(id);
    }

    function getBallotVotingInfo(uint256 id) private view returns (uint256, uint256, uint256) {
        return IBallotStorage(getBallotStorageAddress()).getBallotVotingInfo(id);
    }

    function getBallotMember(
        uint256 id
    ) private view returns (address, address, address, address, bytes memory, bytes memory, bytes memory, uint256, uint256) {
        return IBallotStorage(getBallotStorageAddress()).getBallotMember(id);
    }
    function getBallotForExit(uint256 id) private view returns (uint256, uint256) {
        return IBallotStorage(getBallotStorageAddress()).getBallotForExit(id);
    }

    function lock(address addr, uint256 amount) private {
        IStaking(getStakingAddress()).lock(addr, amount);
    }

    function unlock(address addr, uint256 amount) private {
        IStaking(getStakingAddress()).unlock(addr, amount);
    }

    function transferLockedAndUnlock(uint256 ballotIdx, address addr) private {
        (uint256 unlockAmount, uint256 slashing) = getBallotForExit(ballotIdx);

        require(unlockAmount + slashing <= getMinStaking(), "minStaking value must be greater than or equal to the sum of unlockAmount, slashing");

        IStaking staking = IStaking(getStakingAddress());
        uint256 locked = staking.lockedBalanceOf(addr);
        uint256 ext = locked - getMinStaking();

        if (locked > unlockAmount) {
            unlock(addr, unlockAmount);
            staking.transferLocked(addr, slashing, ext);
        } else {
            unlock(addr, locked);
        }
    }

    function lockedBalanceOf(address addr) private view returns (uint256) {
        return IStaking(getStakingAddress()).lockedBalanceOf(addr);
    }

    function availableBalanceOf(address addr) private view returns (uint256) {
        return IStaking(getStakingAddress()).availableBalanceOf(addr);
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyGovMem {}

    function checkVariableCondition(bytes32 envKey, bytes memory envVal) internal view returns (bool) {
        return IEnvStorage(getEnvStorageAddress()).checkVariableCondition(envKey, envVal);
    }

    function getStakerAddr(address _addr) public view returns (address staker) {
        if (isStaker(_addr)) staker = _addr;
        else if (isVoter(_addr)) staker = stakers[voterIdx[_addr]];
    }

    function setProposalTimePeriod(uint256 newPeriod) external onlyOwner {
        require(newPeriod < 1 hours, "newPeriod is too long");
        proposal_time_period = newPeriod;
        emit SetProposalTimePeriod(newPeriod);
    }

    function checkNodeInfoAdd(bytes memory name, bytes memory enode, bytes memory ip, uint port) internal view returns (bool check) {
        check = true;
        if (checkNodeEnode[enode]) check = false;
        if (checkNodeName[name]) check = false;

        bytes32 hvalue = keccak256(abi.encodePacked(ip, port));
        if (checkNodeIpPort[hvalue]) check = false;
    }

    function checkNodeInfoChange(
        bytes memory name,
        bytes memory enode,
        bytes memory ip,
        uint port,
        Node memory nodeInfo
    ) internal view returns (bool check) {
        check = true;
        if ((keccak256(nodeInfo.enode) != keccak256(enode) && checkNodeEnode[enode])) check = false;
        if ((keccak256(nodeInfo.name) != keccak256(name) && checkNodeName[name])) check = false;

        bytes32 hvalue = keccak256(abi.encodePacked(ip, port));
        if ((keccak256(abi.encodePacked(nodeInfo.ip, nodeInfo.port)) != hvalue && checkNodeIpPort[hvalue])) check = false;
    }

    uint256 public proposal_time_period;
    mapping(address => uint256) public lastAddProposalTime;

    uint256[46] private __gap;

    function reInit() external reinitializer(2) onlyOwner {
        unchecked {
            // LEGACY (pre W1G-03): 0-indexed, exclusive upper bound. Marks the
            // empty nodes[0] sentinel and skips the last real node nodes[N].
            for (uint256 i = 0; i < getMemberLength(); i++) {
                Node memory node = nodes[i];
                checkNodeName[node.name] = true;
                checkNodeEnode[node.enode] = true;
                checkNodeIpPort[keccak256(abi.encodePacked(node.ip, node.port))] = true;
            }
        }
    }

    function initMigration(address registry, uint256 oldModifiedBlock, address oldOwner) external override initializer {
        __ReentrancyGuard_init();
        __Ownable_init();
        setRegistry(registry);

        modifiedBlock = oldModifiedBlock;
        transferOwnership(oldOwner);
        emit GovDataMigrated(msg.sender);
    }

    function migrateFromLegacy(address oldGov) external initializer returns (int256) {
        __ReentrancyGuard_init();
        __Ownable_init();

        GovImpLegacy ogov = GovImpLegacy(oldGov);
        setRegistry(address(ogov.reg()));
        modifiedBlock = block.number;
        transferOwnership(ogov.owner());

        unchecked {
            for (uint256 i = 1; i <= ogov.getMemberLength(); i++) {
                stakers[i] = ogov.getMember(i);
                stakerIdx[stakers[i]] = i;
                voters[i] = ogov.getVoter(i);
                voterIdx[voters[i]] = i;
                rewards[i] = ogov.getReward(i);
                rewardIdx[rewards[i]] = i;
                memberLength = i;

                Node memory node;
                (node.name, node.enode, node.ip, node.port) = ogov.getNode(i);
                // LEGACY (pre W1G-03): dead check — checkNodeInfoChange(.., node)
                // compares node to itself, so every field reads as "unchanged"
                // and the function always returns true. Duplicates are NOT caught.
                require(checkNodeInfoChange(node.name, node.enode, node.ip, node.port, node), "node info is duplicated");
                checkNodeName[node.name] = true;
                checkNodeEnode[node.enode] = true;
                checkNodeIpPort[keccak256(abi.encodePacked(node.ip, node.port))] = true;
                nodes[i] = node;
                nodeIdxFromMember[stakers[i]] = i;
                nodeToMember[i] = stakers[i];
                nodeLength = i;
                lastAddProposalTime[stakers[i]] = ogov.lastAddProposalTime(stakers[i]);
            }
        }

        proposal_time_period = ogov.proposal_time_period();

        ballotLength = ogov.ballotLength();
        voteLength = ogov.voteLength();
        ballotInVoting = ogov.getBallotInVoting();

        return 0;
    }

    function upgradeTo(address) external override {
        revert("Invalid access");
    }

    function upgradeToAndCall(address, bytes memory) external payable override {
        revert("Invalid access");
    }
}
