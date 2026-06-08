package test

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

func TestGov(t *testing.T) {
	// for mute chain log
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(0), log.StreamHandler(os.Stdout, log.TerminalFormat(true))))
	callOpts := new(bind.CallOpts)

	node1 := nodeInfo{
		name:  []byte("name1"),
		enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
		ip:    []byte("127.0.0.2"),
		port:  big.NewInt(8542),
	}

	node2 := nodeInfo{
		name:  []byte("name2"),
		enode: hexutil.MustDecode("0x888777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a1"),
		ip:    []byte("127.0.0.3"),
		port:  big.NewInt(8542),
	}

	t.Run("Staker is voter", func(t *testing.T) {
		t.Run("has enode and locked staking", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			var (
				locked          *big.Int
				idx             *big.Int
				name, enode, ip []uint8
				port            *big.Int
			)

			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&locked}, "lockedBalanceOf", gov.owner.From))
			require.Equal(t, LOCK_AMOUNT, locked)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&idx}, "getNodeIdxFromMember", gov.owner.From))
			require.True(t, idx.Sign() != 0)
			result := []interface{}{}
			require.NoError(t, gov.GovImp.Call(callOpts, &result, "getNode", idx))
			name, enode, ip, port = result[0].([]byte), result[1].([]byte), result[2].([]byte), result[3].(*big.Int)
			nodeinfo := gov.nodeInfos[0]
			require.Equal(t, nodeinfo.name, name)
			require.Equal(t, nodeinfo.enode, enode)
			require.Equal(t, nodeinfo.ip, ip)
			require.Equal(t, nodeinfo.port, port)
		})
		t.Run("cannot init twice", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			node := gov.nodeInfos[0]
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "init", gov.registry, LOCK_AMOUNT, node.name, node.enode, node.ip, node.port)),
				"Initializable: contract is already initialized",
			)
		})
		t.Run("cannot addProposal to add member self", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			node := gov.nodeInfos[0]
			info := MemberInfo{
				Staker:     gov.owner.From,
				Voter:      gov.owner.From,
				Reward:     gov.owner.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", info)),
				"Already member",
			)
		})
		t.Run("cannot addProposal to add member with different voter", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			member := gov.owner.From
			node := gov.nodeInfos[0]
			info := MemberInfo{
				Staker:     member,
				Voter:      common.Address{1},
				Reward:     member,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", info)),
				"Already member",
			)
		})
		t.Run("cannot addProposal to add member with different reward", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			member := gov.owner.From
			node := gov.nodeInfos[0]
			info := MemberInfo{
				Staker:     member,
				Voter:      member,
				Reward:     common.Address{1},
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", info)),
				"Already member",
			)
		})
		t.Run("can addProposal to add member", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			govMem1 := getTxOpt(t, "govMem1")
			// staking first
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(getTxOpt(t, "govMem1"), "deposit"))
			govMem1.Value = nil
			// proposal
			gov.nodeInfos = append(gov.nodeInfos, node1)
			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", info))
			// check proposal
			var (
				length           *big.Int
				creator          common.Address
				memo             []byte
				newStakerAddress common.Address
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big1, length)

			ballot, ballotDetail := []interface{}{}, []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &ballot, "getBallotBasic", length))
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &ballotDetail, "getBallotMember", length))
			creator, memo = ballot[3].(common.Address), ballot[4].([]byte)
			newStakerAddress = ballotDetail[1].(common.Address)
			require.Equal(t, gov.owner.From, creator)
			require.Equal(t, []byte("memo1"), memo)
			require.Equal(t, govMem1.From, newStakerAddress)

			// govMem1 is not member yet
			info.Memo = []byte("memo2")
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", info))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)
		})
		t.Run("cannot addProposal to remove non-member", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			govMem1 := getTxOpt(t, "govMem1").From
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToRemoveMember", govMem1, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int))),
				"Non-member",
			)
		})
		t.Run("cannot addProposal to remove a sole member", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToRemoveMember", gov.owner.From, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int))),
				"Cannot remove a sole member",
			)
		})
		t.Run("can addProposal to change member's other addresses self without voting", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			var (
				oldMember, oldVoter, oldReward common.Address
				length                         *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldReward}, "getReward", common.Big1))

			require.Equal(t, gov.owner.From, oldReward)
			node := gov.nodeInfos[0]
			voter1 := getTxOpt(t, "voter1")
			user1 := getTxOpt(t, "user1")

			info := MemberInfo{
				Staker:     gov.owner.From,
				Voter:      voter1.From,
				Reward:     user1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big1, length)
		})
		t.Run("can addProposal to change member's other addresses self without voting twice about node name", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			var (
				oldMember, oldVoter, oldReward common.Address
				length                         *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldReward}, "getReward", common.Big1))

			require.Equal(t, gov.owner.From, oldMember)
			require.Equal(t, gov.owner.From, oldVoter)
			require.Equal(t, gov.owner.From, oldReward)
			node := gov.nodeInfos[0]

			info := MemberInfo{
				Staker:     gov.owner.From,
				Voter:      gov.owner.From,
				Reward:     gov.owner.From,
				Name:       []byte("name1"), // change
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big1, length)

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				newMember, newVoter, newReward common.Address
			)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", common.Big1))
			require.Equal(t, gov.owner.From, newMember)
			require.Equal(t, gov.owner.From, newVoter)
			require.Equal(t, gov.owner.From, newReward)

			info.Name = node.name
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", common.Big1))
			require.Equal(t, gov.owner.From, newMember)
			require.Equal(t, gov.owner.From, newVoter)
			require.Equal(t, gov.owner.From, newReward)
		})
		t.Run("can addProposal to change member's other addresses self without voting twice about enode", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			var (
				oldMember, oldVoter, oldReward common.Address
				length                         *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldReward}, "getReward", common.Big1))

			require.Equal(t, gov.owner.From, oldMember)
			require.Equal(t, gov.owner.From, oldVoter)
			require.Equal(t, gov.owner.From, oldReward)
			node := gov.nodeInfos[0]

			info := MemberInfo{
				Staker:     gov.owner.From,
				Voter:      gov.owner.From,
				Reward:     gov.owner.From,
				Name:       []byte("name1"),
				Enode:      node1.enode, // change
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big1, length)

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				newMember, newVoter, newReward common.Address
			)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", common.Big1))
			require.Equal(t, gov.owner.From, newMember)
			require.Equal(t, gov.owner.From, newVoter)
			require.Equal(t, gov.owner.From, newReward)

			info.Enode = node.enode
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", common.Big1))
			require.Equal(t, gov.owner.From, newMember)
			require.Equal(t, gov.owner.From, newVoter)
			require.Equal(t, gov.owner.From, newReward)
		})
		t.Run("can addProposal to change member's other addresses self without voting twice about ipport", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			var (
				oldMember, oldVoter, oldReward common.Address
				length                         *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldReward}, "getReward", common.Big1))

			require.Equal(t, gov.owner.From, oldMember)
			require.Equal(t, gov.owner.From, oldVoter)
			require.Equal(t, gov.owner.From, oldReward)
			node := gov.nodeInfos[0]

			info := MemberInfo{
				Staker:     gov.owner.From,
				Voter:      gov.owner.From,
				Reward:     gov.owner.From,
				Name:       []byte("name1"),
				Enode:      node.enode,
				Ip:         node1.ip, // change
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big1, length)

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				newMember, newVoter, newReward common.Address
			)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", common.Big1))
			require.Equal(t, gov.owner.From, newMember)
			require.Equal(t, gov.owner.From, newVoter)
			require.Equal(t, gov.owner.From, newReward)

			info.Ip = node.ip
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", common.Big1))
			require.Equal(t, gov.owner.From, newMember)
			require.Equal(t, gov.owner.From, newVoter)
			require.Equal(t, gov.owner.From, newReward)
		})
		t.Run("can addProposal to change member's other addresses self without voting twice about import2", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			var (
				oldMember, oldVoter, oldReward common.Address
				length                         *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldReward}, "getReward", common.Big1))

			require.Equal(t, gov.owner.From, oldMember)
			require.Equal(t, gov.owner.From, oldVoter)
			require.Equal(t, gov.owner.From, oldReward)
			node := gov.nodeInfos[0]

			info := MemberInfo{
				Staker:     gov.owner.From,
				Voter:      gov.owner.From,
				Reward:     gov.owner.From,
				Name:       []byte("name1"),
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node1.port, // change
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big1, length)

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				newMember, newVoter, newReward common.Address
			)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", common.Big1))
			require.Equal(t, gov.owner.From, newMember)
			require.Equal(t, gov.owner.From, newVoter)
			require.Equal(t, gov.owner.From, newReward)

			info.Port = node.port
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", common.Big1))
			require.Equal(t, gov.owner.From, newMember)
			require.Equal(t, gov.owner.From, newVoter)
			require.Equal(t, gov.owner.From, newReward)
		})
		t.Run("cannot addProposal to change non-member", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			govMem1, govMem2 := getTxOpt(t, "govMem1"), getTxOpt(t, "govMem2")
			node := gov.nodeInfos[0]

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, govMem2.From, common.Big0, common.Big0)),
				"Non-member",
			)
		})
		t.Run("can addProposal to change governance", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			newGovImp, _, err := gov.Deploy(compiled.GovImp.Deploy(gov.backend, gov.owner))
			require.NoError(t, err)
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeGov", newGovImp, []byte("memo"), big.NewInt(86400)))

			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big1, length)
		})
		t.Run("can not addProposal to change governance using EOA", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToChangeGov", getTxOpt(t, "govMem1").From, []byte("memo"), big.NewInt(86400))),
				"",
			)
		})
		t.Run("can not addProposal to change governance using non-UUPS", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToChangeGov", gov.registry, []byte("memo"), big.NewInt(86400))),
				"ERC1967Upgrade: new implementation is not UUPS",
			)
		})
		t.Run("cannot addProposal to change governance with same address", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			var govImp common.Address
			require.NoError(t, gov.Gov.Call(callOpts, &[]interface{}{&govImp}, "implementation"))
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToChangeGov", govImp, []byte("memo"), big.NewInt(86400))),
				"Same contract address",
			)
		})
		t.Run("cannot addProposal to change governance with zero address", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToChangeGov", common.Address{}, []byte("memo"), big.NewInt(86400))),
				"Implementation cannot be zero",
			)
		})
		t.Run("can addProposal to change environment", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeEnv", ToBytes32("key"), EnvTypes.Bytes32, []byte("value"), []byte("memo"), big.NewInt(86400)))

			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big1, length)
		})
		t.Run("cannot addProposal to change environment with wrong type", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToChangeEnv", ToBytes32("key"), EnvTypes.Invalid, []byte("value"), []byte("memo"), big.NewInt(86400))),
				"Invalid type",
			)
		})
		t.Run("can vote approval to add member", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem1, "deposit"))
			govMem1.Value = nil
			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", info))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big1, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				memberLen, nodeLen, lock       *big.Int
				newMember, newVoter, newReward common.Address
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeLen}, "getNodeLength"))
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&lock}, "lockedBalanceOf", govMem1.From))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", memberLen))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", memberLen))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", memberLen))
			require.Equal(t, common.Big2, memberLen)
			require.Equal(t, memberLen, nodeLen)
			require.Equal(t, LOCK_AMOUNT, lock)
			require.Equal(t, govMem1.From, newMember)
			require.Equal(t, govMem1.From, newVoter)
			require.Equal(t, govMem1.From, newReward)
		})
		t.Run("cannot vote approval to add member with insufficient staking", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem1.From))
			govMem1.Value = new(big.Int).Div(LOCK_AMOUNT, common.Big2)
			gov.ExpectedOk(gov.StakingImp.Transact(govMem1, "deposit"))
			govMem1.Value = nil

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", info))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big1, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
				memberLen   *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Rejected, state)
			require.True(t, isFinalized)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.Equal(t, common.Big1, memberLen)
		})
		t.Run("can vote disapproval to deny adding member", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem1, "deposit"))
			govMem1.Value = nil

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", info))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big1, false))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
				memberLen   *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.Equal(t, common.Big1, length)
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Rejected, state)
			require.True(t, isFinalized)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.Equal(t, common.Big1, memberLen)
		})
		t.Run("can vote approval to change member totally", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem1, "deposit"))
			govMem1.Value = nil

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", info))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big1, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.Equal(t, common.Big1, length)
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			govMem2 := getTxOpt(t, "govMem2")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem2.From))
			govMem2.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem2, "deposit"))
			govMem2.Value = nil

			var (
				preDeployerAvail, preGovmem1Avail *big.Int
			)
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&preDeployerAvail}, "availableBalanceOf", govMem1.From))
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&preGovmem1Avail}, "availableBalanceOf", govMem2.From))

			info = MemberInfo{
				Staker:     govMem2.From,
				Voter:      govMem2.From,
				Reward:     govMem2.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo2"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, govMem1.From, LOCK_AMOUNT, common.Big0))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big2, true))
			gov.ExpectedOk(gov.GovImp.Transact(govMem1, "vote", common.Big2, true))

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, big.NewInt(3), length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				memberLen       *big.Int
				memberAddr      common.Address
				name, enode, ip []byte
				port            *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberAddr}, "getMember", common.Big2))
			getNode := []interface{}{}
			require.NoError(t, gov.GovImp.Call(callOpts, &getNode, "getNode", common.Big2))
			require.Equal(t, common.Big2, memberLen)
			require.Equal(t, govMem2.From, memberAddr)
			name, enode, ip, port = getNode[0].([]byte), getNode[1].([]byte), getNode[2].([]byte), getNode[3].(*big.Int)
			require.Equal(t, node1.name, name)
			require.Equal(t, node1.enode, enode)
			require.Equal(t, node1.ip, ip)
			require.Equal(t, node1.port, port)

			var (
				nodeIdxFromDeployer, nodeIdxFromGovMem1 *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeIdxFromDeployer}, "getNodeIdxFromMember", govMem1.From))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeIdxFromGovMem1}, "getNodeIdxFromMember", govMem2.From))
			require.True(t, nodeIdxFromDeployer.Sign() == 0)
			require.Equal(t, common.Big2, nodeIdxFromGovMem1)

			var (
				postDeployerAvail, postGovmem1Avail *big.Int
			)
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&postDeployerAvail}, "availableBalanceOf", govMem1.From))
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&postGovmem1Avail}, "availableBalanceOf", govMem2.From))
			require.Equal(t, LOCK_AMOUNT, new(big.Int).Sub(postDeployerAvail, preDeployerAvail))
			require.Equal(t, LOCK_AMOUNT, new(big.Int).Sub(preGovmem1Avail, postGovmem1Avail))
		})
		t.Run("can vote approval to change enode only without voting", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)

			info := MemberInfo{
				Staker:     gov.owner.From,
				Voter:      gov.owner.From,
				Reward:     gov.owner.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.True(t, length.Sign() == 0)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				memberLen       *big.Int
				memberAddr      common.Address
				name, enode, ip []byte
				port            *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberAddr}, "getMember", common.Big1))
			getNode := []interface{}{}
			require.NoError(t, gov.GovImp.Call(callOpts, &getNode, "getNode", common.Big1))
			require.Equal(t, common.Big1, memberLen)
			require.Equal(t, gov.owner.From, memberAddr)
			name, enode, ip, port = getNode[0].([]byte), getNode[1].([]byte), getNode[2].([]byte), getNode[3].(*big.Int)
			require.Equal(t, node1.name, name)
			require.Equal(t, node1.enode, enode)
			require.Equal(t, node1.ip, ip)
			require.Equal(t, node1.port, port)
		})
		t.Run("cannot vote approval to change member with insufficient staking", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = new(big.Int).Sub(LOCK_AMOUNT, big.NewInt(1000000000))
			gov.ExpectedOk(gov.StakingImp.Transact(getTxOpt(t, "govMem1"), "deposit"))
			govMem1.Value = nil

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big1, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Rejected, state)
			require.True(t, isFinalized)
		})
		t.Run("can vote approval to change governance", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)

			var (
				getGasLimitAndBaseFee []interface{}
				MBF                   *big.Int
			)
			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &getGasLimitAndBaseFee, "getGasLimitAndBaseFee"))
			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &[]interface{}{&MBF}, "getMaxBaseFee"))

			newGovImp, _, err := gov.Deploy(compiled.GovImp.Deploy(gov.backend, gov.owner))
			require.NoError(t, err)

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeGov", newGovImp, []byte("memo"), big.NewInt(86400)))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big1, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var imp common.Address
			require.NoError(t, gov.Gov.Call(callOpts, &[]interface{}{&imp}, "implementation"))
			require.Equal(t, newGovImp, imp)

			var (
				newGetGasLimitAndBaseFee []interface{}
				NEW_MBF                  *big.Int
			)
			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &newGetGasLimitAndBaseFee, "getGasLimitAndBaseFee"))
			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &[]interface{}{&NEW_MBF}, "getMaxBaseFee"))
			require.Equal(t, getGasLimitAndBaseFee[0], newGetGasLimitAndBaseFee[0])
			require.Equal(t, getGasLimitAndBaseFee[1], newGetGasLimitAndBaseFee[1])
			require.Equal(t, getGasLimitAndBaseFee[2], newGetGasLimitAndBaseFee[2])
			require.Equal(t, MBF, NEW_MBF)
		})
		t.Run("can vote approval to change environment", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			var (
				blocksPer *big.Int
			)
			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &[]interface{}{&blocksPer}, "getBlocksPer"))
			require.NotEqual(t, big.NewInt(100), blocksPer)

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				big.NewInt(86400),
			))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big1, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &[]interface{}{&blocksPer}, "getBlocksPer"))
			require.Equal(t, big.NewInt(100), blocksPer)
		})
		t.Run("cannot vote for a ballot already done", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem1.From))
			govMem1.Value = new(big.Int).Sub(LOCK_AMOUNT, big.NewInt(1000000000))
			gov.ExpectedOk(gov.StakingImp.Transact(govMem1, "deposit"))
			govMem1.Value = nil

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big1, true))
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "vote", common.Big1, true)),
				"Expired",
			)
		})
		t.Run("cannot add proposal during period time", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				big.NewInt(86400),
			))

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				big.NewInt(86400),
			))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "setProposalTimePeriod", big.NewInt(60)))

			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToChangeEnv",
					crypto.Keccak256Hash([]byte("blocksPer")),
					EnvTypes.Uint,
					hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
					[]byte("memo"),
					big.NewInt(86400),
				)), "Cannot add proposal too early",
			)
		})
		t.Run("cannot addProposal to add member which is already reward", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(getTxOpt(t, "govMem1"), "deposit"))
			govMem1.Value = nil

			gov.nodeInfos = append(gov.nodeInfos, node1)
			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", info))
			var (
				ballotLength *big.Int
				memberLength *big.Int
			)
			gov.GovImp.Call(callOpts, &[]interface{}{&ballotLength}, "ballotLength")
			require.Equal(t, common.Big1, ballotLength)

			gov.GovImp.Call(callOpts, &[]interface{}{&memberLength}, "getMemberLength")
			node := gov.nodeInfos[0]
			info = MemberInfo{
				Staker:     gov.owner.From,
				Voter:      gov.owner.From,
				Reward:     govMem1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))

			gov.GovImp.Call(callOpts, &[]interface{}{&ballotLength}, "ballotLength")
			require.Equal(t, common.Big2, ballotLength)

			var (
				newMember, newVoter, newReward common.Address
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", memberLength))
			require.Equal(t, gov.owner.From, newMember)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", memberLength))
			require.Equal(t, gov.owner.From, newVoter)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", memberLength))
			require.Equal(t, govMem1.From, newReward)
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big1, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Rejected, state)
			require.True(t, isFinalized)

			var (
				memberLen, nodeLen *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeLen}, "getNodeLength"))
			require.Equal(t, common.Big1, memberLen)
			require.Equal(t, memberLen, nodeLen)

			var (
				lock, bal *big.Int
			)
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&lock}, "lockedBalanceOf", govMem1.From))
			require.True(t, lock.Sign() == 0)
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&bal}, "balanceOf", govMem1.From))
			require.Equal(t, LOCK_AMOUNT, bal)
		})
		t.Run("cannot addProposal to change member which is already reward", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(getTxOpt(t, "govMem1"), "deposit"))
			govMem1.Value = nil

			node := gov.nodeInfos[0]
			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))

			var memberLen *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			info = MemberInfo{
				Staker:     gov.owner.From,
				Voter:      gov.owner.From,
				Reward:     govMem1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))

			var (
				newMember, newVoter, newReward common.Address
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", memberLen))
			require.Equal(t, gov.owner.From, newMember)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", memberLen))
			require.Equal(t, gov.owner.From, newVoter)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", memberLen))
			require.Equal(t, govMem1.From, newReward)

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big1, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Rejected, state)
			require.True(t, isFinalized)

			var (
				/* memberLen */ nodeLen, lock, bal *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.Equal(t, common.Big1, memberLen)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeLen}, "getNodeLength"))
			require.Equal(t, memberLen, nodeLen)
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&lock}, "lockedBalanceOf", govMem1.From))
			require.True(t, lock.Sign() == 0)
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&bal}, "balanceOf", govMem1.From))
			require.Equal(t, LOCK_AMOUNT, bal)
		})
	})
	t.Run("Staker is not a voter", func(t *testing.T) {
		deployGovernance := func(t *testing.T) (gov *Governance, voter *bind.TransactOpts) {
			gov = NewGovernance(t).DeployContracts(t)
			voter = getTxOpt(t, "voter")
			user1 := getTxOpt(t, "user1")
			balance, err := gov.backend.BalanceAt(context.TODO(), gov.owner.From, nil)
			require.NoError(t, err)

			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Div(balance, common.Big2), &voter.From))

			node := gov.nodeInfos[0]
			info := MemberInfo{
				Staker:     gov.owner.From,
				Voter:      voter.From,
				Reward:     user1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			return
		}
		t.Run("cannot addProposal to add member self", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			node := gov.nodeInfos[0]
			info := MemberInfo{
				Staker:     gov.owner.From,
				Voter:      gov.owner.From,
				Reward:     gov.owner.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "addProposalToAddMember", info)),
				"Already member",
			)
		})
		t.Run("cannot addProposal to add member with different voter", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			node := gov.nodeInfos[0]
			info := MemberInfo{
				Staker:     gov.owner.From,
				Voter:      voter.From,
				Reward:     gov.owner.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "addProposalToAddMember", info)),
				"Already member",
			)
		})
		t.Run("cannot addProposal to add member with same voter", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			node := gov.nodeInfos[0]
			info := MemberInfo{
				Staker:     voter.From,
				Voter:      voter.From,
				Reward:     voter.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "addProposalToAddMember", info)),
				"Already member",
			)
		})
		t.Run("cannot addProposal to add member with same reward", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			user1 := getTxOpt(t, "user1")
			node := gov.nodeInfos[0]
			info := MemberInfo{
				Staker:     user1.From,
				Voter:      user1.From,
				Reward:     user1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "addProposalToAddMember", info)),
				"Already member",
			)
		})
		t.Run("can addProposal to add member", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			govMem1 := getTxOpt(t, "govMem1")
			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToAddMember", info))

			var (
				length           *big.Int
				creator          common.Address
				memo             []byte
				newStakerAddress common.Address
			)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)
			ballot, ballotDetail := []interface{}{}, []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &ballot, "getBallotBasic", length))
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &ballotDetail, "getBallotMember", length))
			creator, memo, newStakerAddress = ballot[3].(common.Address), ballot[4].([]byte), ballotDetail[1].(common.Address)
			require.Equal(t, voter.From, creator)
			require.Equal(t, []byte("memo"), memo)
			require.Equal(t, govMem1.From, newStakerAddress)

			info = MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToAddMember", info))

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, big.NewInt(3), length)
		})
		t.Run("cannot addProposal to remove non-member", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			govMem1 := getTxOpt(t, "govMem1")
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "addProposalToRemoveMember", govMem1.From, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int))),
				"Non-member",
			)
		})
		t.Run("cannot addProposal to remove a sole member", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "addProposalToRemoveMember", gov.owner.From, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int))),
				"Cannot remove a sole member",
			)
		})
		t.Run("can addProposal to change member's other addresses self without voting", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			voter1 := getTxOpt(t, "voter1")
			user1 := getTxOpt(t, "user1")
			var (
				oldMember, oldVoter, oldReward common.Address
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&oldReward}, "getReward", common.Big1))
			require.Equal(t, user1.From, oldReward)

			node := gov.nodeInfos[0]
			info := MemberInfo{
				Staker:     gov.owner.From,
				Voter:      voter1.From,
				Reward:     user1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)
			gov.ExpectedOk(gov.GovImp.Transact(voter, "vote", common.Big2, true))

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				newMember, newVoter, newReward common.Address
			)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", common.Big1))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", common.Big1))
			require.Equal(t, gov.owner.From, newMember)
			require.Equal(t, voter1.From, newVoter)
			require.Equal(t, user1.From, newReward)
		})
		t.Run("cannot addProposal to change member's other addresses which is already member", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem1, "deposit"))
			govMem1.Value = nil

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToAddMember", info))
			gov.ExpectedOk(gov.GovImp.Transact(voter, "vote", common.Big2, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				memberLen, nodeLen, lock       *big.Int
				newMember, newVoter, newReward common.Address
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeLen}, "getNodeLength"))
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&lock}, "lockedBalanceOf", govMem1.From))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", memberLen))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", memberLen))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", memberLen))
			require.Equal(t, govMem1.From, newMember)
			require.Equal(t, govMem1.From, newVoter)
			require.Equal(t, govMem1.From, newReward)

			node := gov.nodeInfos[0]
			info = MemberInfo{
				Staker:     gov.owner.From,
				Voter:      govMem1.From,
				Reward:     getTxOpt(t, "user1").From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0)),
				"Already a member",
			)
		})
		t.Run("cannot addProposal to add member which is already voter", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(getTxOpt(t, "govMem1"), "deposit"))
			govMem1.Value = nil
			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToAddMember", info))
			var memberLen *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))

			node := gov.nodeInfos[0]
			user1 := getTxOpt(t, "user1")
			info = MemberInfo{
				Staker:     gov.owner.From,
				Voter:      govMem1.From,
				Reward:     user1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))

			var (
				newMember, newVoter, newReward common.Address
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", memberLen))
			require.Equal(t, gov.owner.From, newMember)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", memberLen))
			require.Equal(t, govMem1.From, newVoter)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", memberLen))
			require.Equal(t, user1.From, newReward)

			gov.ExpectedOk(gov.GovImp.Transact(govMem1, "vote", common.Big2, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big1))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				/*memberLen*/ nodeLen, lock, bal *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.Equal(t, common.Big1, memberLen)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeLen}, "getNodeLength"))
			require.Equal(t, memberLen, nodeLen)
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&lock}, "lockedBalanceOf", govMem1.From))
			require.True(t, lock.Sign() == 0)
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&bal}, "balanceOf", govMem1.From))
			require.Equal(t, LOCK_AMOUNT, bal)
		})
		t.Run("cannot addProposal to change member which is already voter", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(getTxOpt(t, "govMem1"), "deposit"))
			govMem1.Value = nil

			node := gov.nodeInfos[0]
			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))

			var memberLen *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))

			user1 := getTxOpt(t, "user1")
			info = MemberInfo{
				Staker:     gov.owner.From,
				Voter:      govMem1.From,
				Reward:     user1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))

			var (
				newMember, newVoter, newReward common.Address
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", memberLen))
			require.Equal(t, gov.owner.From, newMember)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", memberLen))
			require.Equal(t, govMem1.From, newVoter)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", memberLen))
			require.Equal(t, user1.From, newReward)

			gov.ExpectedOk(gov.GovImp.Transact(govMem1, "vote", common.Big2, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Rejected, state)
			require.True(t, isFinalized)

			var (
				/*memberLen*/ nodeLen, lock, bal *big.Int
			)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.Equal(t, common.Big1, memberLen)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeLen}, "getNodeLength"))
			require.Equal(t, memberLen, nodeLen)
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&lock}, "lockedBalanceOf", govMem1.From))
			require.True(t, lock.Sign() == 0)
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&bal}, "balanceOf", govMem1.From))
			require.Equal(t, LOCK_AMOUNT, bal)
		})
		t.Run("cannot addProposal to change non-member", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			govMem1, govMem2 := getTxOpt(t, "govMem1"), getTxOpt(t, "govMem2")

			node := gov.nodeInfos[0]
			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "addProposalToChangeMember", info, govMem2.From, common.Big0, common.Big0)),
				"Non-member",
			)
		})
		t.Run("can addProposal to change governance", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			newGovImp, _, err := gov.Deploy(compiled.GovImp.Deploy(gov.backend, gov.owner))
			require.NoError(t, err)
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeGov", newGovImp, []byte("memo"), big.NewInt(86400)))

			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)
		})
		t.Run("can not addProposal to change governance using EOA", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "addProposalToChangeGov", getTxOpt(t, "govMem1").From, []byte("memo"), big.NewInt(86400))),
				"",
			)
		})
		t.Run("can not addProposal to change governance using non-UUPS", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "addProposalToChangeGov", gov.registry, []byte("memo"), big.NewInt(86400))),
				"ERC1967Upgrade: new implementation is not UUPS",
			)
		})
		t.Run("cannot addProposal to change governance with same address", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			var govImp common.Address
			require.NoError(t, gov.Gov.Call(callOpts, &[]interface{}{&govImp}, "implementation"))
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "addProposalToChangeGov", govImp, []byte("memo"), big.NewInt(86400))),
				"Same contract address",
			)
		})
		t.Run("cannot addProposal to change governance with zero address", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "addProposalToChangeGov", common.Address{}, []byte("memo"), big.NewInt(86400))),
				"Implementation cannot be zero",
			)
		})
		t.Run("can addProposal to change environment", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeEnv", ToBytes32("key"), EnvTypes.Bytes32, []byte("value"), []byte("memo"), big.NewInt(86400)))

			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)
		})
		t.Run("cannot addProposal to change environment with wrong type", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "addProposalToChangeEnv", ToBytes32("key"), EnvTypes.Invalid, []byte("value"), []byte("memo"), big.NewInt(86400))),
				"Invalid type",
			)
		})
		t.Run("can vote approval to add member", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem1, "deposit"))
			govMem1.Value = nil
			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToAddMember", info))
			gov.ExpectedOk(gov.GovImp.Transact(voter, "vote", common.Big2, true))
			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				memberLen, nodeLen, lock       *big.Int
				newMember, newVoter, newReward common.Address
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeLen}, "getNodeLength"))
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&lock}, "lockedBalanceOf", govMem1.From))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newMember}, "getMember", memberLen))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newVoter}, "getVoter", memberLen))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&newReward}, "getReward", memberLen))
			require.Equal(t, common.Big2, memberLen)
			require.Equal(t, memberLen, nodeLen)
			require.Equal(t, LOCK_AMOUNT, lock)
			require.Equal(t, govMem1.From, newMember)
			require.Equal(t, govMem1.From, newVoter)
			require.Equal(t, govMem1.From, newReward)
		})
		t.Run("cannot vote approval to add member with insufficient staking", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = new(big.Int).Div(LOCK_AMOUNT, common.Big2)
			gov.ExpectedOk(gov.StakingImp.Transact(getTxOpt(t, "govMem1"), "deposit"))

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToAddMember", info))
			gov.ExpectedOk(gov.GovImp.Transact(voter, "vote", common.Big2, true))
			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool

				memberLen *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Rejected, state)
			require.True(t, isFinalized)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.Equal(t, common.Big1, memberLen)
		})
		t.Run("can vote disapproval to deny adding member", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem1, "deposit"))
			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToAddMember", info))
			gov.ExpectedOk(gov.GovImp.Transact(voter, "vote", common.Big2, false))
			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Rejected, state)
			require.True(t, isFinalized)
		})
		t.Run("can vote approval to change member totally", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem1, "deposit"))
			govMem1.Value = nil

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToAddMember", info))
			gov.ExpectedOk(gov.GovImp.Transact(voter, "vote", common.Big2, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.Equal(t, common.Big1, length)
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			govMem2 := getTxOpt(t, "govMem2")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem2.From))
			govMem2.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem2, "deposit"))
			govMem2.Value = nil

			var (
				preDeployerAvail, preGovmem1Avail *big.Int
			)
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&preDeployerAvail}, "availableBalanceOf", govMem1.From))
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&preGovmem1Avail}, "availableBalanceOf", govMem2.From))

			info = MemberInfo{
				Staker:     govMem2.From,
				Voter:      govMem2.From,
				Reward:     govMem2.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo2"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeMember", info, govMem1.From, LOCK_AMOUNT, common.Big0))
			gov.ExpectedOk(gov.GovImp.Transact(voter, "vote", common.Big3, true))
			gov.ExpectedOk(gov.GovImp.Transact(govMem1, "vote", common.Big3, true))

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, big.NewInt(3), length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big3))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				memberLen       *big.Int
				memberAddr      common.Address
				name, enode, ip []byte
				port            *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberAddr}, "getMember", common.Big2))
			getNode := []interface{}{}
			require.NoError(t, gov.GovImp.Call(callOpts, &getNode, "getNode", common.Big2))
			require.Equal(t, common.Big2, memberLen)
			require.Equal(t, govMem2.From, memberAddr)
			name, enode, ip, port = getNode[0].([]byte), getNode[1].([]byte), getNode[2].([]byte), getNode[3].(*big.Int)
			require.Equal(t, node1.name, name)
			require.Equal(t, node1.enode, enode)
			require.Equal(t, node1.ip, ip)
			require.Equal(t, node1.port, port)

			var (
				nodeIdxFromDeployer, nodeIdxFromGovMem1 *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeIdxFromDeployer}, "getNodeIdxFromMember", govMem1.From))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeIdxFromGovMem1}, "getNodeIdxFromMember", govMem2.From))
			require.True(t, nodeIdxFromDeployer.Sign() == 0)
			require.Equal(t, common.Big2, nodeIdxFromGovMem1)

			var (
				postDeployerAvail, postGovmem1Avail *big.Int
			)
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&postDeployerAvail}, "availableBalanceOf", govMem1.From))
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&postGovmem1Avail}, "availableBalanceOf", govMem2.From))
			require.Equal(t, LOCK_AMOUNT, new(big.Int).Sub(postDeployerAvail, preDeployerAvail))
			require.Equal(t, LOCK_AMOUNT, new(big.Int).Sub(preGovmem1Avail, postGovmem1Avail))
		})
		t.Run("can vote approval to change enode only without voting", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			info := MemberInfo{
				Staker:     gov.owner.From,
				Voter:      gov.owner.From,
				Reward:     gov.owner.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big2, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				memberLen       *big.Int
				memberAddr      common.Address
				name, enode, ip []byte
				port            *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberAddr}, "getMember", common.Big1))
			getNode := []interface{}{}
			require.NoError(t, gov.GovImp.Call(callOpts, &getNode, "getNode", common.Big1))
			require.Equal(t, common.Big1, memberLen)
			require.Equal(t, gov.owner.From, memberAddr)
			name, enode, ip, port = getNode[0].([]byte), getNode[1].([]byte), getNode[2].([]byte), getNode[3].(*big.Int)
			require.Equal(t, node1.name, name)
			require.Equal(t, node1.enode, enode)
			require.Equal(t, node1.ip, ip)
			require.Equal(t, node1.port, port)
		})
		t.Run("cannot vote approval to change member with insufficient staking", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = new(big.Int).Sub(LOCK_AMOUNT, big.NewInt(1000000000))
			gov.ExpectedOk(gov.StakingImp.Transact(getTxOpt(t, "govMem1"), "deposit"))
			govMem1.Value = nil

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			gov.ExpectedOk(gov.GovImp.Transact(voter, "vote", common.Big2, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Rejected, state)
			require.True(t, isFinalized)
		})
		t.Run("can vote approval to change governance", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			var (
				getGasLimitAndBaseFee []interface{}
				MBF                   *big.Int
			)
			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &getGasLimitAndBaseFee, "getGasLimitAndBaseFee"))
			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &[]interface{}{&MBF}, "getMaxBaseFee"))

			newGovImp, _, err := gov.Deploy(compiled.GovImp.Deploy(gov.backend, gov.owner))
			require.NoError(t, err)

			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeGov", newGovImp, []byte("memo"), big.NewInt(86400)))
			gov.ExpectedOk(gov.GovImp.Transact(voter, "vote", common.Big2, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var imp common.Address
			require.NoError(t, gov.Gov.Call(callOpts, &[]interface{}{&imp}, "implementation"))
			require.Equal(t, newGovImp, imp)

			var (
				newGetGasLimitAndBaseFee []interface{}
				NEW_MBF                  *big.Int
			)
			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &newGetGasLimitAndBaseFee, "getGasLimitAndBaseFee"))
			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &[]interface{}{&NEW_MBF}, "getMaxBaseFee"))
			require.Equal(t, getGasLimitAndBaseFee[0], newGetGasLimitAndBaseFee[0])
			require.Equal(t, getGasLimitAndBaseFee[1], newGetGasLimitAndBaseFee[1])
			require.Equal(t, getGasLimitAndBaseFee[2], newGetGasLimitAndBaseFee[2])
			require.Equal(t, MBF, NEW_MBF)
		})
		t.Run("can vote approval to change environment", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			var (
				blocksPer *big.Int
			)
			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &[]interface{}{&blocksPer}, "getBlocksPer"))
			require.NotEqual(t, big.NewInt(100), blocksPer)

			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				big.NewInt(86400),
			))
			gov.ExpectedOk(gov.GovImp.Transact(voter, "vote", common.Big2, true))

			var (
				length      *big.Int
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &[]interface{}{&blocksPer}, "getBlocksPer"))
			require.Equal(t, big.NewInt(100), blocksPer)
		})
		t.Run("cannot vote for a ballot already done", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem1.From))
			govMem1.Value = new(big.Int).Sub(LOCK_AMOUNT, big.NewInt(1000000000))
			gov.ExpectedOk(gov.StakingImp.Transact(govMem1, "deposit"))
			govMem1.Value = nil

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			gov.ExpectedOk(gov.GovImp.Transact(voter, "vote", common.Big2, true))
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "vote", common.Big2, true)),
				"Expired",
			)
		})
		t.Run("cannot vote for a ballot already staker done", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = new(big.Int).Sub(LOCK_AMOUNT, big.NewInt(1000000000))
			gov.ExpectedOk(gov.StakingImp.Transact(getTxOpt(t, "govMem1"), "deposit"))

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big2, true))
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "vote", common.Big2, true)),
				"Expired",
			)
		})
		t.Run("cannot vote for a ballot already voter done", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = new(big.Int).Sub(LOCK_AMOUNT, big.NewInt(1000000000))
			gov.ExpectedOk(gov.StakingImp.Transact(getTxOpt(t, "govMem1"), "deposit"))

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeMember", info, gov.owner.From, common.Big0, common.Big0))
			gov.ExpectedOk(gov.GovImp.Transact(voter, "vote", common.Big2, true))
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "vote", common.Big2, true)),
				"Expired",
			)
		})
		t.Run("cannot add proposal during period time", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				big.NewInt(86400),
			))

			gov.ExpectedOk(gov.GovImp.Transact(voter, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				big.NewInt(86400),
			))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "setProposalTimePeriod", big.NewInt(60)))

			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter, "addProposalToChangeEnv",
					crypto.Keccak256Hash([]byte("blocksPer")),
					EnvTypes.Uint,
					hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
					[]byte("memo"),
					big.NewInt(86400),
				)), "Cannot add proposal too early",
			)
		})
	})
	t.Run("Two Member", func(t *testing.T) {
		deployGovernance := func(t *testing.T) (gov *Governance, govMem1 *bind.TransactOpts) {
			gov = NewGovernance(t).DeployContracts(t)
			govMem1 = getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem1, "deposit"))
			govMem1.Value = nil

			info := MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", info))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big1, true))

			var memberLen *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.Equal(t, common.Big2, memberLen)

			gov.nodeInfos = append(gov.nodeInfos, node1)
			return
		}
		t.Run("cannot vote with changed voter address", func(t *testing.T) {
			gov, _ := deployGovernance(t)

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeEnv", ToBytes32("key"), EnvTypes.Bytes32, []byte("value"), []byte("memo"), big.NewInt(86400)))
			var (
				ballotIdx *big.Int
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&ballotIdx}, "ballotLength"))

			node := gov.nodeInfos[0]
			voter, voter1, user1 := getTxOpt(t, "voter"), getTxOpt(t, "voter1"), getTxOpt(t, "user1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, towei(1), &voter.From))
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, towei(1), &voter1.From))

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.owner.From,
				Voter:      voter.From,
				Reward:     user1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}, gov.owner.From, common.Big0, common.Big0))
			gov.ExpectedOk(gov.GovImp.Transact(voter, "vote", ballotIdx, true))

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.owner.From,
				Voter:      voter1.From,
				Reward:     user1.From,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}, gov.owner.From, common.Big0, common.Big0))

			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(voter1, "vote", ballotIdx, true)),
				"already voted",
			)
		})
		t.Run("cannot addProposal to add member self", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)
			node := gov.nodeInfos[0]
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(govMem1, "addProposalToAddMember",
					MemberInfo{
						Staker:     govMem1.From,
						Voter:      govMem1.From,
						Reward:     govMem1.From,
						Name:       node.name,
						Enode:      node.enode,
						Ip:         node.ip,
						Port:       node.port,
						LockAmount: LOCK_AMOUNT,
						Memo:       []byte("memo1"),
						Duration:   big.NewInt(86400),
					},
				)), "Already member",
			)
		})
		t.Run("can addProposal to remove member", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToRemoveMember", govMem1.From, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int)))

			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)
		})
		t.Run("can addProposal to add member where info is the removed member's with same govMem", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToRemoveMember", govMem1.From, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int)))

			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", length, true))

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.Equal(t, length, inVoting)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.InProgress, state)
			require.False(t, isFinalized)

			gov.ExpectedOk(gov.GovImp.Transact(govMem1, "vote", length, true))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0, inVoting)

			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)
		})
		t.Run("can addProposal to add member where info is the removed member's", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToRemoveMember", govMem1.From, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int)))

			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", length, true))

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.Equal(t, length, inVoting)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.InProgress, state)
			require.False(t, isFinalized)

			gov.ExpectedOk(gov.GovImp.Transact(govMem1, "vote", length, true))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0, inVoting)

			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			t.Log("member removed")
			govMem2 := getTxOpt(t, "govMem2")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem2.From))
			govMem2.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem2, "deposit"))
			govMem2.Value = nil

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem2.From,
				Voter:      govMem2.From,
				Reward:     govMem2.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}))

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			gov.ExpectedOk(gov.GovImp.Transact(govMem1, "vote", length, true))
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)
		})
		t.Run("can vote to add member", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)

			govMem2 := getTxOpt(t, "govMem2")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem2.From))
			govMem2.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem2, "deposit"))
			govMem2.Value = nil

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem2.From,
				Voter:      govMem2.From,
				Reward:     govMem2.From,
				Name:       node2.name,
				Enode:      node2.enode,
				Ip:         node2.ip,
				Port:       node2.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}))

			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", length, true))

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.Equal(t, length, inVoting)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.InProgress, state)
			require.False(t, isFinalized)

			gov.ExpectedOk(gov.GovImp.Transact(govMem1, "vote", length, true))

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0, inVoting)
			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)
		})
		t.Run("can vote to deny adding member", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)

			govMem2 := getTxOpt(t, "govMem2")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem2.From))
			govMem2.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem2, "deposit"))
			govMem2.Value = nil

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem2.From,
				Voter:      govMem2.From,
				Reward:     govMem2.From,
				Name:       node2.name,
				Enode:      node2.enode,
				Ip:         node2.ip,
				Port:       node2.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}))

			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", length, false))

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.Equal(t, length, inVoting)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.InProgress, state)
			require.False(t, isFinalized)

			gov.ExpectedOk(gov.GovImp.Transact(govMem1, "vote", length, false))

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0, inVoting)
			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Rejected, state)
			require.True(t, isFinalized)
		})
		t.Run("can vote to remove first member", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)

			var preAvail, postAvail *big.Int
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&preAvail}, "availableBalanceOf", gov.owner.From))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToRemoveMember", gov.owner.From, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int)))

			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", length, true))

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.Equal(t, length, inVoting)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.InProgress, state)
			require.False(t, isFinalized)

			gov.ExpectedOk(gov.GovImp.Transact(govMem1, "vote", length, true))

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0, inVoting)
			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				memberLen, nodeLen, nodeIdx, nodeIdx2 *big.Int
				isMem                                 bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.Equal(t, common.Big1, memberLen)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&isMem}, "isMember", gov.owner.From))
			require.False(t, isMem)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeLen}, "getNodeLength"))
			require.Equal(t, common.Big1, nodeLen)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeIdx}, "getNodeIdxFromMember", gov.owner.From))
			require.True(t, nodeIdx.Sign() == 0)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeIdx2}, "getNodeIdxFromMember", govMem1.From))
			require.Equal(t, common.Big1, nodeIdx2)

			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&postAvail}, "availableBalanceOf", gov.owner.From))
			require.Equal(t, LOCK_AMOUNT, new(big.Int).Sub(postAvail, preAvail))
		})
		t.Run("can vote to remove last member", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)

			var preAvail, postAvail *big.Int
			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&preAvail}, "availableBalanceOf", gov.owner.From))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToRemoveMember", govMem1.From, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int)))

			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", length, true))

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.Equal(t, length, inVoting)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.InProgress, state)
			require.False(t, isFinalized)

			gov.ExpectedOk(gov.GovImp.Transact(govMem1, "vote", length, true))

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0, inVoting)
			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Accepted, state)
			require.True(t, isFinalized)

			var (
				memberLen, nodeLen, nodeIdx, nodeIdx2 *big.Int
				isMem                                 bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.Equal(t, common.Big1, memberLen)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&isMem}, "isMember", govMem1.From))
			require.False(t, isMem)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeLen}, "getNodeLength"))
			require.Equal(t, common.Big1, nodeLen)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeIdx}, "getNodeIdxFromMember", govMem1.From))
			require.True(t, nodeIdx.Sign() == 0)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&nodeIdx2}, "getNodeIdxFromMember", gov.owner.From))
			require.Equal(t, common.Big1, nodeIdx2)

			require.NoError(t, gov.StakingImp.Call(callOpts, &[]interface{}{&postAvail}, "availableBalanceOf", govMem1.From))
			require.Equal(t, LOCK_AMOUNT, new(big.Int).Sub(postAvail, preAvail))
		})
		t.Run("cannot vote simultaneously", func(t *testing.T) {
			gov, _ := deployGovernance(t)

			govMem2 := getTxOpt(t, "govMem2")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem2.From))
			govMem2.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem2, "deposit"))
			govMem2.Value = nil

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem2.From,
				Voter:      govMem2.From,
				Reward:     govMem2.From,
				Name:       node2.name,
				Enode:      node2.enode,
				Ip:         node2.ip,
				Port:       node2.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}))

			govMem3 := getTxOpt(t, "govMem3")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem3.From))
			govMem3.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem3, "deposit"))
			govMem3.Value = nil

			node3 := nodeInfo{
				name:  []byte("name3"),
				enode: hexutil.MustDecode("0x999777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a1"),
				ip:    []byte("127.0.0.4"),
				port:  big.NewInt(8542),
			}

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem3.From,
				Voter:      govMem3.From,
				Reward:     govMem3.From,
				Name:       node3.name,
				Enode:      node3.enode,
				Ip:         node3.ip,
				Port:       node3.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}))

			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", new(big.Int).Sub(length, common.Big1), true))
			var inVoting *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))

			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "vote", length, true)),
				"Now in voting with different ballot",
			)
		})
		t.Run("vote is ended when the sum of voting power is max", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)

			govMem2 := getTxOpt(t, "govMem2")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), &govMem2.From))
			govMem2.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(govMem2, "deposit"))
			govMem2.Value = nil

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem2.From,
				Voter:      govMem2.From,
				Reward:     govMem2.From,
				Name:       node2.name,
				Enode:      node2.enode,
				Ip:         node2.ip,
				Port:       node2.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}))

			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "ballotLength"))
			require.Equal(t, common.Big2, length)

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", length, true))

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.Equal(t, length, inVoting)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.InProgress, state)
			require.False(t, isFinalized)

			gov.ExpectedOk(gov.GovImp.Transact(govMem1, "vote", length, false))

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0, inVoting)
			getBallotState = []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.Rejected, state)
			require.True(t, isFinalized)
		})
		t.Run("cannot vote approval when the voting is ended", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)

			delay_time := big.NewInt(86400)
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				delay_time,
			))
			var ballotLen *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&ballotLen}, "ballotLength"))

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", ballotLen, true))
			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big2, length)

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.Equal(t, length, inVoting)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.InProgress, state)
			require.False(t, isFinalized)

			gov.backend.AdjustTime(time.Second * time.Duration(new(big.Int).Mul(delay_time, big.NewInt(2000)).Int64()))
			gov.backend.Commit()

			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(govMem1, "vote", ballotLen, false)),
				"Expired",
			)
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "finalizeEndedVote"))

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)
		})
		t.Run("Non member cannot end voting", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)

			delay_time := big.NewInt(86400)
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				delay_time,
			))
			var ballotLen *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&ballotLen}, "ballotLength"))

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", ballotLen, true))
			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big2, length)

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.Equal(t, length, inVoting)
			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", length))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.Equal(t, BallotStates.InProgress, state)
			require.False(t, isFinalized)

			gov.backend.AdjustTime(time.Second * time.Duration(new(big.Int).Mul(delay_time, big.NewInt(2000)).Int64()))
			gov.backend.Commit()
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(govMem1, "vote", ballotLen, false)),
				"Expired",
			)
			govMem3 := getTxOpt(t, "getMem3")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, towei(1), &govMem3.From))
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(govMem3, "finalizeEndedVote")),
				"No Permission",
			)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.Equal(t, ballotLen, inVoting)
		})
		t.Run("reject proposal without voting about changing voter address if voter is already registered", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)
			node := gov.nodeInfos[0]
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "addProposalToChangeMember", MemberInfo{
					Staker:     gov.owner.From,
					Voter:      govMem1.From,
					Reward:     gov.owner.From,
					Name:       node.name,
					Enode:      node.enode,
					Ip:         node.ip,
					Port:       node.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				}, gov.owner.From, common.Big0, common.Big0)),
				"Already a member",
			)
			var length *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&length}, "voteLength"))
			require.Equal(t, common.Big1, length)

			var (
				inVoting    *big.Int
				state       *big.Int
				isFinalized bool
			)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&inVoting}, "getBallotInVoting"))
			require.True(t, inVoting.Sign() == 0)

			getBallotState := []interface{}{}
			require.NoError(t, gov.BallotStorageImp.Call(callOpts, &getBallotState, "getBallotState", common.Big2))
			state, isFinalized = getBallotState[1].(*big.Int), getBallotState[2].(bool)
			require.True(t, state.Sign() == 0)
			require.False(t, isFinalized)

			var (
				memberLen                         *big.Int
				memberAddr, voterAddr, rewardAddr common.Address
			)

			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.Equal(t, common.Big2, memberLen)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberAddr}, "getMember", common.Big1))
			require.Equal(t, gov.owner.From, memberAddr)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&voterAddr}, "getVoter", common.Big1))
			require.Equal(t, gov.owner.From, voterAddr)
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&rewardAddr}, "getReward", common.Big1))
			require.Equal(t, gov.owner.From, rewardAddr)

			var (
				name, enode, ip []byte
				port            *big.Int
			)
			getNode := []interface{}{}
			require.NoError(t, gov.GovImp.Call(callOpts, &getNode, "getNode", common.Big1))
			name, enode, ip, port = getNode[0].([]byte), getNode[1].([]byte), getNode[2].([]byte), getNode[3].(*big.Int)
			require.Equal(t, node.name, name)
			require.Equal(t, node.enode, enode)
			require.Equal(t, node.ip, ip)
			require.Equal(t, node.port, port)
		})
	})
	t.Run("Others", func(t *testing.T) {
		t.Run("cannot init", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			govMem1 := getTxOpt(t, "govMem1")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(getTxOpt(t, "govMem1"), "deposit"))
			govMem1.Value = nil

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem1.From,
				Voter:      govMem1.From,
				Reward:     govMem1.From,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}))
			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", common.Big1, true))

			var memberLen *big.Int
			require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&memberLen}, "getMemberLength"))
			require.Equal(t, common.Big2, memberLen)

			node := gov.nodeInfos[0]
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(govMem1, "init", gov.registry, LOCK_AMOUNT, node.name, node.enode, node.ip, node.port)),
				"Initializable: contract is already initialized",
			)
		})
		t.Run("cannot addProposal", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			govMem1, govMem2 := getTxOpt(t, "govMem1"), getTxOpt(t, "govMem2")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, towei(1), &govMem1.From))
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, towei(1), &govMem2.From))

			node := gov.nodeInfos[0]
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(govMem1, "addProposalToAddMember", MemberInfo{
					Staker:     govMem1.From,
					Voter:      govMem1.From,
					Reward:     govMem1.From,
					Name:       node.name,
					Enode:      node.enode,
					Ip:         node.ip,
					Port:       node.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo"),
					Duration:   big.NewInt(86400),
				})),
				"No Permission",
			)
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(govMem1, "addProposalToRemoveMember", govMem1.From, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int))),
				"No Permission",
			)

			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem1.From))
			govMem1.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(getTxOpt(t, "govMem1"), "deposit"))
			govMem1.Value = nil

			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(govMem1, "addProposalToChangeMember", MemberInfo{
					Staker:     govMem2.From,
					Voter:      govMem2.From,
					Reward:     govMem2.From,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo"),
					Duration:   big.NewInt(86400),
				}, govMem1.From, common.Big0, common.Big0)),
				"No Permission",
			)
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(govMem1, "addProposalToChangeGov", govMem1.From, []byte("memo"), big.NewInt(86400))),
				"No Permission",
			)
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(govMem1, "addProposalToChangeEnv", ToBytes32("key"), EnvTypes.Bytes32, []byte("value"), []byte("memo"), big.NewInt(86400))),
				"No Permission",
			)
		})
		t.Run("cannot vote", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)
			govMem1, govMem2 := getTxOpt(t, "govMem1"), getTxOpt(t, "govMem2")
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, towei(1), &govMem1.From))
			gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, new(big.Int).Add(LOCK_AMOUNT, towei(1)), &govMem2.From))
			govMem2.Value = LOCK_AMOUNT
			gov.ExpectedOk(gov.StakingImp.Transact(getTxOpt(t, "govMem2"), "deposit"))
			govMem2.Value = nil

			gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem2.From,
				Voter:      govMem2.From,
				Reward:     govMem2.From,
				Name:       node2.name,
				Enode:      node2.enode,
				Ip:         node2.ip,
				Port:       node2.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}))
			ExpectedRevert(t,
				gov.ExpectedFail(gov.GovImp.Transact(govMem1, "vote", common.Big1, true)),
				"No Permission",
			)
		})
		t.Run("upgradeTo", func(t *testing.T) {
			gov := NewGovernance(t).DeployContracts(t)

			var (
				callOpts = new(bind.CallOpts)

				getGasLimitAndBaseFee []interface{}
				MBF                   *big.Int
			)

			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &getGasLimitAndBaseFee, "getGasLimitAndBaseFee"))
			require.NoError(t, gov.EnvStorageImp.Call(callOpts, &[]interface{}{&MBF}, "getMaxBaseFee"))

			newGovImp, _, err := gov.Deploy(compiled.GovImp.Deploy(gov.backend, gov.owner))
			require.NoError(t, err)

			ExpectedRevert(t, gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "upgradeTo", newGovImp)), "Invalid access")
			ExpectedRevert(t, gov.ExpectedFail(gov.GovImp.Transact(gov.owner, "upgradeToAndCall", newGovImp, []byte{})), "Invalid access")
		})
	})
}

func TestGov_IndexCorruptionAfterRemoveMember(t *testing.T) {
	var (
		gov *Governance

		// member B
		memberB        = getTxOpt(t, "memberB")
		memberB_Reward = getTxOpt(t, "memberB_Reward")
		memberB_Voter  = getTxOpt(t, "memberB_Voter")
		memberB_Node   = nodeInfo{
			name:  []byte("name1"),
			enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
			ip:    []byte("127.0.0.2"),
			port:  big.NewInt(8542),
		}

		// member C
		memberC      = getTxOpt(t, "memberC")
		memberC_Node = nodeInfo{
			name:  []byte("name2"),
			enode: hexutil.MustDecode("0x888777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a1"),
			ip:    []byte("127.0.0.3"),
			port:  big.NewInt(8542),
		}

		callOpts = new(bind.CallOpts)
	)

	// setup for test
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(0), log.StreamHandler(os.Stdout, log.TerminalFormat(true))))
	gov = NewGovernance(t).DeployContracts(t)
	require.NoError(t, gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, towei(2_000_000), &memberB.From)))
	require.NoError(t, gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, towei(2_000_000), &memberC.From)))

	// staking for member B
	memberB.Value = LOCK_AMOUNT
	require.NoError(t, gov.ExpectedOk(gov.StakingImp.Transact(memberB, "deposit")))
	memberB.Value = nil

	// proposal to add member B
	gov.nodeInfos = append(gov.nodeInfos, memberB_Node)
	memeberB_info := MemberInfo{
		Staker:     memberB.From,
		Voter:      memberB.From,
		Reward:     memberB.From,
		Name:       memberB_Node.name,
		Enode:      memberB_Node.enode,
		Ip:         memberB_Node.ip,
		Port:       memberB_Node.port,
		LockAmount: LOCK_AMOUNT,
		Memo:       []byte("memo1"),
		Duration:   big.NewInt(86400),
	}
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", memeberB_info)))

	// vote for member B
	ballotIdx := getBallotIdx(t, gov)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", ballotIdx, true)))

	memberBIdx := getMemberIdx(t, gov, memberB.From)
	checkGovMember(t, gov, memberBIdx, memeberB_info)

	// staking for member C
	memberC.Value = LOCK_AMOUNT
	require.NoError(t, gov.ExpectedOk(gov.StakingImp.Transact(memberC, "deposit")))
	memberC.Value = nil

	// proposal to add member C
	gov.nodeInfos = append(gov.nodeInfos, memberC_Node)
	memberC_info := MemberInfo{
		Staker:     memberC.From,
		Voter:      memberC.From,
		Reward:     memberC.From,
		Name:       memberC_Node.name,
		Enode:      memberC_Node.enode,
		Ip:         memberC_Node.ip,
		Port:       memberC_Node.port,
		LockAmount: LOCK_AMOUNT,
		Memo:       []byte("memo2"),
		Duration:   big.NewInt(86400),
	}
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToAddMember", memberC_info)))

	// vote for member C (member B is already a member)
	ballotIdx = getBallotIdx(t, gov)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", ballotIdx, true)))
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(memberB, "vote", ballotIdx, true)))

	memberCIdx := getMemberIdx(t, gov, memberC.From)
	checkGovMember(t, gov, memberCIdx, memberC_info)

	// change member B's reward address
	memberB_info := MemberInfo{
		Staker:     memberB.From,
		Voter:      memberB_Voter.From,
		Reward:     memberB_Reward.From,
		Name:       memberB_Node.name,
		Enode:      memberB_Node.enode,
		Ip:         memberB_Node.ip,
		Port:       memberB_Node.port,
		LockAmount: LOCK_AMOUNT,
		Memo:       []byte("change memberB reward addr"),
		Duration:   big.NewInt(86400),
	}
	// Self-change by old staker is finalized immediately without a normal vote
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(memberB, "addProposalToChangeMember", memberB_info, memberB.From, big.NewInt(0), big.NewInt(0))))

	checkGovMember(t, gov, memberBIdx, memberB_info)
	{
		var isVoter, isReward bool
		gov.GovImp.Call(callOpts, &[]interface{}{&isVoter}, "isVoter", memberB.From)
		require.False(t, isVoter)
		gov.GovImp.Call(callOpts, &[]interface{}{&isReward}, "isReward", memberB.From)
		require.False(t, isReward)
	}

	// remove member B
	targetIdx := new(big.Int).Set(memberBIdx)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "addProposalToRemoveMember", memberB.From, LOCK_AMOUNT, []byte("remove"), big.NewInt(86400), LOCK_AMOUNT, big.NewInt(0))))
	ballotIdx = getBallotIdx(t, gov)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "vote", ballotIdx, true)))
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(memberC, "vote", ballotIdx, true)))

	checkGovMember(t, gov, big.NewInt(0), memberB_info)
	checkGovMember(t, gov, targetIdx, memberC_info)
}

// ==========================================================================
// Shared test helpers (ledger assertions, ballot voting, legacy-upgrade driver)
// ==========================================================================

func checkGovMember(t *testing.T, gov *Governance, idx *big.Int, info MemberInfo) {
	var callOpts = new(bind.CallOpts)

	// 1. Mapping table verification (Address -> Index)
	var stakerIdx, voterIdx, rewardIdx, nodeIdx *big.Int
	gov.GovImp.Call(callOpts, &[]interface{}{&stakerIdx}, "stakerIdx", info.Staker)
	gov.GovImp.Call(callOpts, &[]interface{}{&voterIdx}, "voterIdx", info.Voter)
	gov.GovImp.Call(callOpts, &[]interface{}{&rewardIdx}, "rewardIdx", info.Reward)
	gov.GovImp.Call(callOpts, &[]interface{}{&nodeIdx}, "getNodeIdxFromMember", info.Staker)

	// The provided idx must match all mapping indexes
	require.True(t, idx.Cmp(stakerIdx) == 0)
	require.True(t, idx.Cmp(voterIdx) == 0)
	require.True(t, idx.Cmp(rewardIdx) == 0)
	require.True(t, idx.Cmp(nodeIdx) == 0)

	if idx.Sign() == 0 {
		return
	}

	var (
		staker, voter, reward common.Address
		name, enode, ip       []byte
		port                  *big.Int
	)

	gov.GovImp.Call(callOpts, &[]interface{}{&staker}, "getMember", idx)
	require.Equal(t, info.Staker, staker)

	gov.GovImp.Call(callOpts, &[]interface{}{&voter}, "getVoter", idx)
	require.Equal(t, info.Voter, voter)

	gov.GovImp.Call(callOpts, &[]interface{}{&reward}, "getReward", idx)
	require.Equal(t, info.Reward, reward)

	getNode := []interface{}{}
	gov.GovImp.Call(callOpts, &getNode, "getNode", idx)
	require.Len(t, getNode, 4)
	name, enode, ip, port = getNode[0].([]byte), getNode[1].([]byte), getNode[2].([]byte), getNode[3].(*big.Int)
	require.Equal(t, info.Name, name)
	require.Equal(t, info.Enode, enode)
	require.Equal(t, info.Ip, ip)
	require.True(t, info.Port.Cmp(port) == 0)
}

func getMemberIdx(t *testing.T, gov *Governance, staker common.Address) *big.Int {
	var idx *big.Int
	require.NoError(t, gov.GovImp.Call(new(bind.CallOpts), &[]interface{}{&idx}, "stakerIdx", staker))
	return idx
}

func getBallotIdx(t *testing.T, gov *Governance) *big.Int {
	var length *big.Int
	require.NoError(t, gov.GovImp.Call(new(bind.CallOpts), &[]interface{}{&length}, "ballotLength"))
	return length
}

// enode helper: 64-byte enodes that only differ in their first byte so each
// scenario node is unique by name/enode/ip while staying a valid 64-byte value.
func w1gEnode(prefix string) []byte {
	const rest = "4311c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"
	return hexutil.MustDecode("0x" + prefix + rest)
}

func w1gFundAndStake(t *testing.T, gov *Governance, opt *bind.TransactOpts) {
	require.NoError(t, gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, towei(2_000_000), &opt.From)))
	opt.Value = LOCK_AMOUNT
	require.NoError(t, gov.ExpectedOk(gov.StakingImp.Transact(opt, "deposit")))
	opt.Value = nil
}

func w1gInfo(staker, voter, reward common.Address, node nodeInfo, memo string) MemberInfo {
	return MemberInfo{
		Staker: staker, Voter: voter, Reward: reward,
		Name: node.name, Enode: node.enode, Ip: node.ip, Port: node.port,
		LockAmount: LOCK_AMOUNT, Memo: []byte(memo), Duration: big.NewInt(86400),
	}
}

func w1gBool(t *testing.T, gov *Governance, method string, arg common.Address) bool {
	var b bool
	require.NoError(t, gov.GovImp.Call(new(bind.CallOpts), &[]interface{}{&b}, method, arg))
	return b
}

func w1gMemberLength(t *testing.T, gov *Governance) int64 {
	var n *big.Int
	require.NoError(t, gov.GovImp.Call(new(bind.CallOpts), &[]interface{}{&n}, "getMemberLength"))
	return n.Int64()
}

func w1gStakerAddr(t *testing.T, gov *Governance, a common.Address) common.Address {
	var out common.Address
	require.NoError(t, gov.GovImp.Call(new(bind.CallOpts), &[]interface{}{&out}, "getStakerAddr", a))
	return out
}

// like checkGovMember but against an arbitrary bound GovImp (used for the
// migrated proxy in the W1G-03 migration test).
func w1gCheckMemberOn(t *testing.T, bc *bind.BoundContract, idx *big.Int, info MemberInfo) {
	callOpts := new(bind.CallOpts)
	var stakerIdx, voterIdx, rewardIdx, nodeIdx *big.Int
	require.NoError(t, bc.Call(callOpts, &[]interface{}{&stakerIdx}, "stakerIdx", info.Staker))
	require.NoError(t, bc.Call(callOpts, &[]interface{}{&voterIdx}, "voterIdx", info.Voter))
	require.NoError(t, bc.Call(callOpts, &[]interface{}{&rewardIdx}, "rewardIdx", info.Reward))
	require.NoError(t, bc.Call(callOpts, &[]interface{}{&nodeIdx}, "getNodeIdxFromMember", info.Staker))
	require.Equal(t, 0, idx.Cmp(stakerIdx))
	require.Equal(t, 0, idx.Cmp(voterIdx))
	require.Equal(t, 0, idx.Cmp(rewardIdx))
	require.Equal(t, 0, idx.Cmp(nodeIdx))

	var staker, voter, reward common.Address
	require.NoError(t, bc.Call(callOpts, &[]interface{}{&staker}, "getMember", idx))
	require.NoError(t, bc.Call(callOpts, &[]interface{}{&voter}, "getVoter", idx))
	require.NoError(t, bc.Call(callOpts, &[]interface{}{&reward}, "getReward", idx))
	require.Equal(t, info.Staker, staker)
	require.Equal(t, info.Voter, voter)
	require.Equal(t, info.Reward, reward)

	getNode := []interface{}{}
	require.NoError(t, bc.Call(callOpts, &getNode, "getNode", idx))
	require.Len(t, getNode, 4)
	require.Equal(t, info.Name, getNode[0].([]byte))
	require.Equal(t, info.Enode, getNode[1].([]byte))
	require.Equal(t, info.Ip, getNode[2].([]byte))
	require.Equal(t, 0, info.Port.Cmp(getNode[3].(*big.Int)))
}

// ballotFinalized reads BallotStorage.getBallotState(ballotIdx) and returns its
// isFinalized flag (index 2 of the tuple).
func ballotFinalized(t *testing.T, gov *Governance, ballotIdx *big.Int) bool {
	state := []interface{}{}
	require.NoError(t, gov.BallotStorageImp.Call(new(bind.CallOpts), &state, "getBallotState", ballotIdx))
	require.Len(t, state, 3)
	return state[2].(bool)
}

// voteAccept casts an approval vote from each given voter in turn and stops as
// soon as the ballot finalizes (threshold reached). It asserts the ballot is
// finalized at the end. It does NOT assert acceptance — callers verify the
// concrete effect (member added / removed / impl changed) themselves, because
// a finalized ballot may legitimately end Rejected (e.g. NotApplicable).
func voteAccept(t *testing.T, gov *Governance, ballotIdx *big.Int, voters ...*bind.TransactOpts) {
	for _, v := range voters {
		if ballotFinalized(t, gov, ballotIdx) {
			break
		}
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(v, "vote", ballotIdx, true)))
	}
	require.True(t, ballotFinalized(t, gov, ballotIdx), "ballot %s did not finalize", ballotIdx)
}

// govImplementation returns the implementation address the Gov proxy delegates
// to right now.
func govImplementation(t *testing.T, gov *Governance) common.Address {
	var imp common.Address
	require.NoError(t, gov.Gov.Call(new(bind.CallOpts), &[]interface{}{&imp}, "implementation"))
	return imp
}

// upgradeGovToFixed deploys the production GovImp, drives a GovernanceChange
// ballot through `proposer` + `voters` until the proxy points at the fixed
// impl, then runs reInit (owner-only, reinitializer(2)). This is the real
// mainnet upgrade path. Returns the new (fixed) implementation address.
func upgradeGovToFixed(t *testing.T, gov *Governance, proposer *bind.TransactOpts, voters ...*bind.TransactOpts) common.Address {
	fixedImp, _, err := gov.Deploy(compiled.GovImp.Deploy(gov.backend, gov.owner))
	require.NoError(t, err)
	require.NotEqual(t, govImplementation(t, gov), fixedImp)

	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(
		proposer, "addProposalToChangeGov", fixedImp, []byte("upgrade to fixed GovImp"), big.NewInt(86400))))
	ballot := getBallotIdx(t, gov)
	for _, v := range voters {
		if govImplementation(t, gov) == fixedImp {
			break
		}
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(v, "vote", ballot, true)))
	}
	require.Equal(t, fixedImp, govImplementation(t, gov), "proxy did not upgrade to fixed impl")

	// reInit backfills node-uniqueness markers (1..N) on the now-fixed impl.
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(gov.owner, "reInit")))
	return fixedImp
}

// getNodeEnode returns the enode bytes stored for member index idx.
func getNodeEnode(t *testing.T, gov *Governance, idx *big.Int) []byte {
	out := []interface{}{}
	require.NoError(t, gov.GovImp.Call(new(bind.CallOpts), &out, "getNode", idx))
	require.Len(t, out, 4)
	return out[1].([]byte)
}

func getVoterAt(t *testing.T, gov *Governance, idx *big.Int) common.Address {
	var v common.Address
	require.NoError(t, gov.GovImp.Call(new(bind.CallOpts), &[]interface{}{&v}, "getVoter", idx))
	return v
}

// addPlainMember registers a staker==voter==reward member through the normal
// propose+vote path and returns its member index. `voters` must reach quorum.
func addPlainMember(t *testing.T, gov *Governance, proposer, member *bind.TransactOpts, node nodeInfo, memo string, voters ...*bind.TransactOpts) *big.Int {
	w1gFundAndStake(t, gov, member)
	info := w1gInfo(member.From, member.From, member.From, node, memo)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(proposer, "addProposalToAddMember", info)))
	voteAccept(t, gov, getBallotIdx(t, gov), voters...)
	idx := getMemberIdx(t, gov, member.From)
	require.True(t, w1gBool(t, gov, "isStaker", member.From))
	return idx
}

// ==========================================================================
// W1G-01 — duplicate node uniqueness at execution
// ==========================================================================

// W1G-01: two add-member proposals carrying the SAME node metadata are both
// queued (creation check passes because markers are only set at execution).
// After the first is applied, the second must be rejected at execution so no
// duplicate validator identity is written.
func TestW1G01_DuplicateNodeRejectedAtExecution(t *testing.T) {
	gov := NewGovernance(t).DeployContracts(t)
	owner := gov.owner
	C := getTxOpt(t, "w1g01_C")
	D := getTxOpt(t, "w1g01_D")
	dupNode := nodeInfo{[]byte("w1g01-dup"), w1gEnode("aaaaaaaa"), []byte("10.0.0.50"), big.NewInt(8542)}

	w1gFundAndStake(t, gov, C)
	w1gFundAndStake(t, gov, D)

	infoC := w1gInfo(C.From, C.From, C.From, dupNode, "C dup")
	infoD := w1gInfo(D.From, D.From, D.From, dupNode, "D dup") // identical node metadata

	// queue BOTH proposals before either executes -> both pass creation-time check
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", infoC)))
	ballotC := getBallotIdx(t, gov)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", infoD)))
	ballotD := getBallotIdx(t, gov)

	// apply C (owner is sole member, weight 10000) -> C added, node markers set
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "vote", ballotC, true)))
	cIdx := getMemberIdx(t, gov, C.From)
	checkGovMember(t, gov, cIdx, infoC)
	require.Equal(t, int64(2), w1gMemberLength(t, gov))

	// apply D with the SAME node: members={owner,C} (2) -> both must vote
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "vote", ballotD, true)))
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(C, "vote", ballotD, true)))

	// W1G-01: addMember re-validates node uniqueness -> D rejected, NOT added
	require.False(t, w1gBool(t, gov, "isStaker", D.From))
	require.Equal(t, int64(2), w1gMemberLength(t, gov))
}

// ============================================================================
// W1G-01 — duplicate node at execution.  RED on legacy, GREEN after upgrade.
//
// Two add-member proposals carry the SAME node metadata. Both pass the
// creation-time check (markers are only set at execution). On LEGACY the second
// addMember writes the duplicate node anyway. After upgrade the queued second
// proposal is re-validated at execution and rejected (NotApplicable).
// ============================================================================
func TestW1G01_Legacy_DuplicateNode_RedThenGreen(t *testing.T) {
	dupNode := nodeInfo{[]byte("w01-dup"), w1gEnode("01010101"), []byte("10.1.9.9"), big.NewInt(8542)}

	t.Run("RED_legacy", func(t *testing.T) {
		gov := NewGovernance(t).DeployContractsLegacy(t)
		owner := gov.owner
		C := getTxOpt(t, "w01_C")
		D := getTxOpt(t, "w01_D")
		w1gFundAndStake(t, gov, C)
		w1gFundAndStake(t, gov, D)

		// queue BOTH proposals before either executes.
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", w1gInfo(C.From, C.From, C.From, dupNode, "C dup"))))
		ballotC := getBallotIdx(t, gov)
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", w1gInfo(D.From, D.From, D.From, dupNode, "D dup"))))
		ballotD := getBallotIdx(t, gov)

		voteAccept(t, gov, ballotC, owner) // owner sole -> C added, markers set
		cIdx := getMemberIdx(t, gov, C.From)
		require.Equal(t, int64(2), w1gMemberLength(t, gov))

		voteAccept(t, gov, ballotD, owner, C) // legacy addMember does NOT re-check
		dIdx := getMemberIdx(t, gov, D.From)

		// RED: D was added with the SAME node identity -> duplicate validators.
		require.True(t, w1gBool(t, gov, "isStaker", D.From), "RED: D should be wrongly admitted on legacy")
		require.Equal(t, int64(3), w1gMemberLength(t, gov))
		require.Equal(t, getNodeEnode(t, gov, cIdx), getNodeEnode(t, gov, dIdx), "RED: two members share one enode")
	})

	t.Run("GREEN_after_upgrade", func(t *testing.T) {
		gov := NewGovernance(t).DeployContractsLegacy(t)
		owner := gov.owner
		C := getTxOpt(t, "w01_C")
		D := getTxOpt(t, "w01_D")
		w1gFundAndStake(t, gov, C)
		w1gFundAndStake(t, gov, D)

		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", w1gInfo(C.From, C.From, C.From, dupNode, "C dup"))))
		ballotC := getBallotIdx(t, gov)
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", w1gInfo(D.From, D.From, D.From, dupNode, "D dup"))))
		ballotD := getBallotIdx(t, gov)

		voteAccept(t, gov, ballotC, owner) // C added on legacy
		cIdx := getMemberIdx(t, gov, C.From)
		require.Equal(t, int64(2), w1gMemberLength(t, gov))

		// upgrade with the D proposal STILL QUEUED, then execute it on the fixed impl.
		upgradeGovToFixed(t, gov, owner, owner, C)
		voteAccept(t, gov, ballotD, owner, C)

		// GREEN: fixed addMember re-validates uniqueness -> D rejected.
		require.False(t, w1gBool(t, gov, "isStaker", D.From), "GREEN: D must be rejected at execution")
		require.Equal(t, int64(2), w1gMemberLength(t, gov))
		require.Equal(t, dupNode.enode, getNodeEnode(t, gov, cIdx))
	})
}

// ==========================================================================
// W1G-02 — stale removal must not empty governance
// ==========================================================================

// W1G-02: a removal proposal created while >1 members exist must not be able
// to drop the member count to zero when executed later as a stale ballot.
func TestW1G02_StaleRemovalCannotEmptyGovernance(t *testing.T) {
	gov := NewGovernance(t).DeployContracts(t)
	owner := gov.owner
	B := getTxOpt(t, "w1g02_B")
	nodeB := nodeInfo{[]byte("w1g02-B"), w1gEnode("ffffffff"), []byte("127.0.0.21"), big.NewInt(8542)}

	// register B -> members {owner(A), B}, memberLength=2
	w1gFundAndStake(t, gov, B)
	infoB := w1gInfo(B.From, B.From, B.From, nodeB, "add B")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", infoB)))
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "vote", getBallotIdx(t, gov), true)))
	require.Equal(t, int64(2), w1gMemberLength(t, gov))

	// STALE proposal to remove A, created while memberLength==2 (passes creation > 1)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToRemoveMember", owner.From, LOCK_AMOUNT, []byte("rm A stale"), big.NewInt(86400), LOCK_AMOUNT, big.NewInt(0))))
	staleBallotA := getBallotIdx(t, gov)

	// remove B (owner + B vote) -> memberLength 2 -> 1 (A sole remaining)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToRemoveMember", B.From, LOCK_AMOUNT, []byte("rm B"), big.NewInt(86400), LOCK_AMOUNT, big.NewInt(0))))
	rmB := getBallotIdx(t, gov)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "vote", rmB, true)))
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(B, "vote", rmB, true)))
	require.Equal(t, int64(1), w1gMemberLength(t, gov))
	require.True(t, w1gBool(t, gov, "isStaker", owner.From))

	// execute the stale removal against the SOLE member A (A votes, weight 10000)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "vote", staleBallotA, true)))

	// W1G-02: removeMember re-checks getMemberLength() > 1 -> A NOT removed
	require.Equal(t, int64(1), w1gMemberLength(t, gov))
	require.True(t, w1gBool(t, gov, "isStaker", owner.From))
}

// ============================================================================
// W1G-02 — stale removal empties governance.  RED on legacy, GREEN after upgrade.
//
// A removal proposal created while >1 members exist is executed later, after
// the member count has dropped to 1. On LEGACY removeMember has no >1 re-check,
// so the last member is removed and memberLength hits 0 (network bricked). The
// fixed contract re-checks and rejects (NotApplicable).
// ============================================================================
func TestW1G02_Legacy_EmptyGovernance_RedThenGreen(t *testing.T) {
	nodeB := nodeInfo{[]byte("w02-B"), w1gEnode("02020202"), []byte("10.2.0.2"), big.NewInt(8542)}

	// builds: members {owner, B}, a STALE removal-of-owner ballot, then removes B
	// so only owner remains. returns the stale ballot id (Ready, not yet executed).
	build := func(t *testing.T, gov *Governance, B *bind.TransactOpts) *big.Int {
		owner := gov.owner
		addPlainMember(t, gov, owner, B, nodeB, "add B", owner)
		require.Equal(t, int64(2), w1gMemberLength(t, gov))

		// stale removal of owner, created while memberLength==2 (passes creation >1).
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(
			owner, "addProposalToRemoveMember", owner.From, LOCK_AMOUNT, []byte("rm owner stale"), big.NewInt(86400), LOCK_AMOUNT, big.NewInt(0))))
		stale := getBallotIdx(t, gov)

		// remove B -> only owner remains (memberLength 2 -> 1).
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(
			owner, "addProposalToRemoveMember", B.From, LOCK_AMOUNT, []byte("rm B"), big.NewInt(86400), LOCK_AMOUNT, big.NewInt(0))))
		voteAccept(t, gov, getBallotIdx(t, gov), owner, B)
		require.Equal(t, int64(1), w1gMemberLength(t, gov))
		return stale
	}

	t.Run("RED_legacy", func(t *testing.T) {
		gov := NewGovernance(t).DeployContractsLegacy(t)
		owner := gov.owner
		B := getTxOpt(t, "w02_B")
		stale := build(t, gov, B)

		// execute the stale removal against the SOLE member (owner self-votes, weight 10000).
		voteAccept(t, gov, stale, owner)

		// RED: legacy removeMember has no >1 guard -> governance emptied.
		require.Equal(t, int64(0), w1gMemberLength(t, gov), "RED: memberLength should reach 0 on legacy")
		require.False(t, w1gBool(t, gov, "isStaker", owner.From))
	})

	t.Run("GREEN_after_upgrade", func(t *testing.T) {
		gov := NewGovernance(t).DeployContractsLegacy(t)
		owner := gov.owner
		B := getTxOpt(t, "w02_B")
		stale := build(t, gov, B)

		upgradeGovToFixed(t, gov, owner, owner) // owner is sole member, weight 10000

		// execute the stale removal on the fixed impl.
		voteAccept(t, gov, stale, owner)

		// GREEN: fixed removeMember re-checks >1 -> sole member preserved.
		require.Equal(t, int64(1), w1gMemberLength(t, gov), "GREEN: sole member must survive")
		require.True(t, w1gBool(t, gov, "isStaker", owner.From))
	})
}

// ==========================================================================
// W1G-03 — migration / reInit marker integrity
// ==========================================================================

// W1G-03 (migrateFromLegacy): migrating from a legacy gov must copy the member
// ledger 1..N consistently and re-establish node-uniqueness markers. This also
// exercises the fixed duplicate check (checkNodeInfoChange(.., node) -> the real
// checkNodeInfoAdd): unique legacy nodes must pass migration without reverting.
//
// Teeth limitation: the duplicate-rejection branch only fires for a legacy gov
// that already holds duplicate node metadata, which the uniqueness-enforcing
// contracts cannot produce; that case needs a mock legacy and is out of scope here.
func TestW1G03_MigrateFromLegacyIntegrity(t *testing.T) {
	gov := NewGovernance(t).DeployContractsLegacy(t) // legacy gov, owner = member 1
	owner := gov.owner
	B := getTxOpt(t, "w1g03_mig_B")
	nodeB := nodeInfo{[]byte("w1g03m-B"), w1gEnode("12121212"), []byte("127.0.0.31"), big.NewInt(8542)}

	// legacy gains a 2nd member
	w1gFundAndStake(t, gov, B)
	infoB := w1gInfo(B.From, B.From, B.From, nodeB, "add B")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", infoB)))
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "vote", getBallotIdx(t, gov), true)))
	require.Equal(t, int64(2), w1gMemberLength(t, gov))

	// legacy gov proxy address (resolved from the shared registry)
	var legacyAddr common.Address
	require.NoError(t, gov.Registry.Call(new(bind.CallOpts), &[]interface{}{&legacyAddr}, "getContractAddress", ToBytes32("GovernanceContract")))

	// deploy a fresh implementation + proxy, then migrate the legacy state into it
	newImpAddr, _, err := gov.Deploy(compiled.GovImp.Deploy(gov.backend, owner))
	require.NoError(t, err)
	newProxyAddr, _, err := gov.Deploy(compiled.Gov.Deploy(gov.backend, owner, newImpAddr))
	require.NoError(t, err)
	newGov := compiled.GovImp.New(gov.backend, newProxyAddr)

	require.NoError(t, gov.ExpectedOk(newGov.Transact(owner, "migrateFromLegacy", legacyAddr)))

	// the migrated ledger must be 1..N consistent (no off-by-one / index corruption)
	var ml *big.Int
	require.NoError(t, newGov.Call(new(bind.CallOpts), &[]interface{}{&ml}, "getMemberLength"))
	require.Equal(t, int64(2), ml.Int64())
	w1gCheckMemberOn(t, newGov, big.NewInt(1), w1gInfo(owner.From, owner.From, owner.From, gov.nodeInfos[0], ""))
	w1gCheckMemberOn(t, newGov, big.NewInt(2), infoB)
}

// W1G-03 (reInit): reInit backfills node-uniqueness markers across 1..N.
// After reInit, every member node (including the LAST one) must remain
// uniqueness-protected.
//
// Teeth limitation: the 1..N (vs buggy 0..N-1) loop only changes observable
// behaviour when markers start empty (a pre-marker in-place upgrade). In a fresh
// deploy addMember already sets markers, so even a buggy 0-indexed reInit that
// skips the last node leaves that node's marker set — this test alone cannot
// distinguish the buggy reInit from the fixed one. It is therefore a smoke test
// (reInit runs and does not corrupt uniqueness on a fresh deploy); the
// off-by-one's actual teeth live in TestW1G03_PreMarker_ReInitOffByOne_RedThenGreen,
// which uses the GovImpPreMarker fixture (empty markers) to make the skip observable.
func TestW1G03_ReInitPreservesNodeUniqueness(t *testing.T) {
	gov := NewGovernance(t).DeployContracts(t)
	owner := gov.owner
	B := getTxOpt(t, "w1g03_ri_B")
	nodeB := nodeInfo{[]byte("w1g03r-B"), w1gEnode("34343434"), []byte("127.0.0.41"), big.NewInt(8542)}

	w1gFundAndStake(t, gov, B)
	infoB := w1gInfo(B.From, B.From, B.From, nodeB, "add B")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", infoB)))
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "vote", getBallotIdx(t, gov), true)))
	require.Equal(t, int64(2), w1gMemberLength(t, gov)) // B is the last node here (index 2)

	// reInit (reinitializer(2), owner-only) must not revert
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "reInit")))

	// node uniqueness still enforced after reInit: reusing the LAST member's enode is rejected
	C := getTxOpt(t, "w1g03_ri_C")
	w1gFundAndStake(t, gov, C)
	dupOfB := w1gInfo(C.From, C.From, C.From, nodeInfo{[]byte("w1g03r-C"), nodeB.enode, []byte("127.0.0.42"), big.NewInt(8542)}, "dup enode of B")
	ExpectedRevert(t,
		gov.ExpectedFail(gov.GovImp.Transact(owner, "addProposalToAddMember", dupOfB)),
		"Duplicated node info",
	)
}

// ============================================================================
// W1G-03 (migrate) — dead duplicate-check.  RED on legacy, GREEN on fixed.
//
// A legacy gov is driven (via the legacy W1G-01 bug) into holding two members
// with the SAME node metadata. Migrating that ledger with the LEGACY
// migrateFromLegacy (checkNodeInfoChange(.., node) -> always true) silently
// accepts the duplicate; the FIXED migrateFromLegacy (checkNodeInfoAdd) rejects
// it with "node info is duplicated".
// ============================================================================
func TestW1G03_Legacy_MigrateDeadCheck_RedThenGreen(t *testing.T) {
	dupNode := nodeInfo{[]byte("w03m-dup"), w1gEnode("03030303"), []byte("10.3.0.9"), big.NewInt(8542)}

	// build a legacy source gov that already holds a duplicate node, return its proxy addr.
	buildDuplicateLegacy := func(t *testing.T, gov *Governance) common.Address {
		owner := gov.owner
		C := getTxOpt(t, "w03m_C")
		D := getTxOpt(t, "w03m_D")
		w1gFundAndStake(t, gov, C)
		w1gFundAndStake(t, gov, D)
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", w1gInfo(C.From, C.From, C.From, dupNode, "C"))))
		ballotC := getBallotIdx(t, gov)
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", w1gInfo(D.From, D.From, D.From, dupNode, "D"))))
		ballotD := getBallotIdx(t, gov)
		voteAccept(t, gov, ballotC, owner)
		voteAccept(t, gov, ballotD, owner, C)
		require.Equal(t, int64(3), w1gMemberLength(t, gov)) // owner, C, D (C & D duplicate node)

		var srcAddr common.Address
		require.NoError(t, gov.Registry.Call(new(bind.CallOpts), &[]interface{}{&srcAddr}, "getContractAddress", ToBytes32("GovernanceContract")))
		return srcAddr
	}

	t.Run("RED_legacy_migrate_accepts_duplicate", func(t *testing.T) {
		gov := NewGovernance(t).DeployContractsLegacy(t)
		owner := gov.owner
		srcAddr := buildDuplicateLegacy(t, gov)

		// fresh LEGACY impl + proxy, migrate the duplicate-bearing ledger into it.
		impl, _, err := gov.Deploy(compiled.GovImpLegacy.Deploy(gov.backend, owner))
		require.NoError(t, err)
		proxy, _, err := gov.Deploy(compiled.Gov.Deploy(gov.backend, owner, impl))
		require.NoError(t, err)
		newLegacy := compiled.GovImpLegacy.New(gov.backend, proxy)

		// RED: dead check -> migration succeeds despite the duplicate.
		require.NoError(t, gov.ExpectedOk(newLegacy.Transact(owner, "migrateFromLegacy", srcAddr)))
		var ml *big.Int
		require.NoError(t, newLegacy.Call(new(bind.CallOpts), &[]interface{}{&ml}, "getMemberLength"))
		require.Equal(t, int64(3), ml.Int64(), "RED: duplicate ledger migrated wholesale on legacy")
	})

	t.Run("GREEN_fixed_migrate_rejects_duplicate", func(t *testing.T) {
		gov := NewGovernance(t).DeployContractsLegacy(t)
		owner := gov.owner
		srcAddr := buildDuplicateLegacy(t, gov)

		// fresh FIXED impl + proxy, migrate the same ledger -> must revert.
		impl, _, err := gov.Deploy(compiled.GovImp.Deploy(gov.backend, owner))
		require.NoError(t, err)
		proxy, _, err := gov.Deploy(compiled.Gov.Deploy(gov.backend, owner, impl))
		require.NoError(t, err)
		newFixed := compiled.GovImp.New(gov.backend, proxy)

		// GREEN: fixed migrate re-checks uniqueness -> reverts at the duplicate.
		// (Asserted via Contains because the shared ExpectedRevert helper mangles
		// revert strings whose leading chars collide with "execution reverted:".)
		migErr := gov.ExpectedFail(newFixed.Transact(owner, "migrateFromLegacy", srcAddr))
		require.Error(t, migErr)
		require.Contains(t, migErr.Error(), "node info is duplicated")
	})
}

// ============================================================================
// W1G-03 (migrate) — ledger preservation for a SEPARATED member.
//
// migrateFromLegacy copies getMember(i)/getVoter(i)/getReward(i) independently,
// so a member whose staker / voter / reward keys were self-separated on the
// legacy gov must survive migration with all three index maps still pointing at
// the same slot. This exercises the migrate path (distinct from the changeGov
// upgrade path) on a separated ledger — the "ledger preservation" concern of
// Check 2 — using the fixed migrateFromLegacy on a clean (non-duplicate) source.
// ============================================================================
func TestW1G03_Legacy_MigratePreservesSeparation(t *testing.T) {
	gov := NewGovernance(t).DeployContractsLegacy(t)
	owner := gov.owner
	B := getTxOpt(t, "w03s_B")
	Bv := getTxOpt(t, "w03s_Bv")
	Br := getTxOpt(t, "w03s_Br")
	C := getTxOpt(t, "w03s_C")
	nodeB := nodeInfo{[]byte("w03s-B"), w1gEnode("06060606"), []byte("10.3.2.2"), big.NewInt(8542)}
	nodeC := nodeInfo{[]byte("w03s-C"), w1gEnode("07070707"), []byte("10.3.2.3"), big.NewInt(8542)}

	// legacy ledger: owner(1), B(2) self-separated to Bv/Br, C(3) plain.
	bIdx := addPlainMember(t, gov, owner, B, nodeB, "add B", owner)
	infoBsep := w1gInfo(B.From, Bv.From, Br.From, nodeB, "separate B")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(B, "addProposalToChangeMember", infoBsep, B.From, big.NewInt(0), big.NewInt(0))))
	checkGovMember(t, gov, bIdx, infoBsep)
	cIdx := addPlainMember(t, gov, owner, C, nodeC, "add C", owner, B)
	require.Equal(t, int64(3), w1gMemberLength(t, gov))

	// legacy source proxy address.
	var srcAddr common.Address
	require.NoError(t, gov.Registry.Call(new(bind.CallOpts), &[]interface{}{&srcAddr}, "getContractAddress", ToBytes32("GovernanceContract")))

	// fresh FIXED impl + proxy, migrate the separated ledger into it.
	impl, _, err := gov.Deploy(compiled.GovImp.Deploy(gov.backend, owner))
	require.NoError(t, err)
	proxy, _, err := gov.Deploy(compiled.Gov.Deploy(gov.backend, owner, impl))
	require.NoError(t, err)
	newGov := compiled.GovImp.New(gov.backend, proxy)
	require.NoError(t, gov.ExpectedOk(newGov.Transact(owner, "migrateFromLegacy", srcAddr)))

	// migrated ledger: every slot 1..N consistent, and the SEPARATED member B
	// keeps stakerIdx[B]==voterIdx[Bv]==rewardIdx[Br]==nodeIdx[B]==bIdx.
	var ml *big.Int
	require.NoError(t, newGov.Call(new(bind.CallOpts), &[]interface{}{&ml}, "getMemberLength"))
	require.Equal(t, int64(3), ml.Int64())
	w1gCheckMemberOn(t, newGov, big.NewInt(1), w1gInfo(owner.From, owner.From, owner.From, gov.nodeInfos[0], ""))
	w1gCheckMemberOn(t, newGov, bIdx, infoBsep) // separation preserved through migration
	w1gCheckMemberOn(t, newGov, cIdx, w1gInfo(C.From, C.From, C.From, nodeC, ""))

	// sanity: the separated voter resolves to its staker on the migrated gov.
	var stakerOfBv common.Address
	require.NoError(t, newGov.Call(new(bind.CallOpts), &[]interface{}{&stakerOfBv}, "getStakerAddr", Bv.From))
	require.Equal(t, B.From, stakerOfBv)
}

// ============================================================================
// W1G-03 (reInit) — marker-backfill off-by-one.  RED on pre-marker, GREEN fixed.
//
// On an implementation that never maintained node-uniqueness markers (the state
// reInit's backfill exists to repair), the LEGACY reInit loops 0..N-1: it marks
// the empty nodes[0] sentinel and SKIPS the last real node nodes[N]. That last
// member's identifier is therefore left reusable. The fixed reInit (1..N) marks
// every real node, so the last member's identifier is protected.
// ============================================================================
func TestW1G03_PreMarker_ReInitOffByOne_RedThenGreen(t *testing.T) {
	nodeB := nodeInfo{[]byte("w03r-B"), w1gEnode("04040404"), []byte("10.3.1.2"), big.NewInt(8542)}
	nodeC := nodeInfo{[]byte("w03r-C"), w1gEnode("05050505"), []byte("10.3.1.3"), big.NewInt(8542)} // the LAST member's node

	build := func(t *testing.T, gov *Governance) (B, C *bind.TransactOpts) {
		owner := gov.owner
		B = getTxOpt(t, "w03r_B")
		C = getTxOpt(t, "w03r_C")
		addPlainMember(t, gov, owner, B, nodeB, "add B", owner)
		addPlainMember(t, gov, owner, C, nodeC, "add C", owner, B) // C is the LAST member (index 3)
		require.Equal(t, int64(3), w1gMemberLength(t, gov))
		return
	}

	t.Run("RED_premarker_buggy_reInit_skips_last_node", func(t *testing.T) {
		gov := NewGovernance(t).DeployContractsPreMarker(t)
		owner := gov.owner
		B, C := build(t, gov)
		cIdx := getMemberIdx(t, gov, C.From) // the LAST member (index 3)

		// buggy reInit (0..N-1): marks nodes[0..2] (sentinel + owner + B), skips nodes[3]=C.
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "reInit")))

		// reusing B's enode (marked) is correctly blocked...
		E := getTxOpt(t, "w03r_E")
		w1gFundAndStake(t, gov, E)
		reuseB := w1gInfo(E.From, E.From, E.From, nodeInfo{[]byte("w03r-E1"), nodeB.enode, []byte("10.3.1.7"), big.NewInt(8542)}, "reuse B enode")
		ExpectedRevert(t, gov.ExpectedFail(gov.GovImp.Transact(owner, "addProposalToAddMember", reuseB)), "Duplicated node info")

		// ...but reusing the LAST member C's enode SLIPS THROUGH (off-by-one left it
		// unmarked), and PreMarker's addMember has no execution-time re-check either,
		// so the duplicate is fully ADMITTED (proposal -> vote -> executed add).
		reuseC := w1gInfo(E.From, E.From, E.From, nodeInfo{[]byte("w03r-E2"), nodeC.enode, []byte("10.3.1.8"), big.NewInt(8542)}, "reuse C enode")
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", reuseC)),
			"RED: last node's enode should be wrongly reusable after buggy reInit")
		voteAccept(t, gov, getBallotIdx(t, gov), owner, B) // owner + B reach quorum (E is the target)

		eIdx := getMemberIdx(t, gov, E.From)
		require.True(t, w1gBool(t, gov, "isStaker", E.From), "RED: duplicate-enode member wrongly admitted")
		require.Equal(t, int64(4), w1gMemberLength(t, gov))
		require.Equal(t, getNodeEnode(t, gov, cIdx), getNodeEnode(t, gov, eIdx), "RED: E and the LAST member C share one enode")
	})

	t.Run("GREEN_fixed_reInit_marks_last_node", func(t *testing.T) {
		gov := NewGovernance(t).DeployContractsPreMarker(t)
		owner := gov.owner
		B, C := build(t, gov)
		_ = C

		// upgrade runs the FIXED reInit (1..N) -> every real node marked, incl. the last.
		upgradeGovToFixed(t, gov, owner, owner, B)

		// reusing the LAST member C's enode is now correctly blocked.
		E := getTxOpt(t, "w03r_E")
		w1gFundAndStake(t, gov, E)
		reuseC := w1gInfo(E.From, E.From, E.From, nodeInfo{[]byte("w03r-E2"), nodeC.enode, []byte("10.3.1.8"), big.NewInt(8542)}, "reuse C enode")
		ExpectedRevert(t,
			gov.ExpectedFail(gov.GovImp.Transact(owner, "addProposalToAddMember", reuseC)),
			"Duplicated node info",
		)
	})
}

// ==========================================================================
// W1G-04 — staker/voter identification & separation
// ==========================================================================

// W1G-04: after a staker self-separates its voter/reward, verify the ledger,
// that a voter-only address cannot drive a self-finalized change (no slot-0
// corruption), and that both the separated voter and the staker can still
// operate normally afterwards.
func TestW1G04_StakerVoterSeparationLifecycle(t *testing.T) {
	callOpts := new(bind.CallOpts)
	gov := NewGovernance(t).DeployContracts(t)
	owner := gov.owner

	B := getTxOpt(t, "w1g04_B")
	Bv := getTxOpt(t, "w1g04_B_voter")
	Br := getTxOpt(t, "w1g04_B_reward")
	C := getTxOpt(t, "w1g04_C")

	nodeB := nodeInfo{[]byte("w1g04-B"), w1gEnode("bbbbbbbb"), []byte("127.0.0.11"), big.NewInt(8542)}
	nodeX := nodeInfo{[]byte("w1g04-X"), w1gEnode("cccccccc"), []byte("127.0.0.99"), big.NewInt(9999)}
	nodeY := nodeInfo{[]byte("w1g04-Y"), w1gEnode("dddddddd"), []byte("127.0.0.12"), big.NewInt(8543)}
	nodeC := nodeInfo{[]byte("w1g04-C"), w1gEnode("eeeeeeee"), []byte("127.0.0.13"), big.NewInt(8544)}

	// register member B (staker=voter=reward=B): memberLength 1 -> 2
	w1gFundAndStake(t, gov, B)
	infoB := w1gInfo(B.From, B.From, B.From, nodeB, "add B")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", infoB)))
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "vote", getBallotIdx(t, gov), true)))
	bIdx := getMemberIdx(t, gov, B.From)
	checkGovMember(t, gov, bIdx, infoB)

	// S1: staker B self-updates to SEPARATE voter & reward (self-finalized, no vote)
	infoBsep := w1gInfo(B.From, Bv.From, Br.From, nodeB, "separate B")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(B, "addProposalToChangeMember", infoBsep, B.From, big.NewInt(0), big.NewInt(0))))
	checkGovMember(t, gov, bIdx, infoBsep) // stakerIdx[B]=voterIdx[Bv]=rewardIdx[Br]=nodeIdx[B]=bIdx
	require.True(t, w1gBool(t, gov, "isStaker", B.From))
	require.False(t, w1gBool(t, gov, "isVoter", B.From))
	require.True(t, w1gBool(t, gov, "isVoter", Bv.From))
	require.False(t, w1gBool(t, gov, "isStaker", Bv.From))
	require.Equal(t, B.From, w1gStakerAddr(t, gov, Bv.From))

	// the separated voter key needs gas to send its own transactions (it stakes nothing)
	require.NoError(t, gov.ExpectedOk(TransferCoin(gov.backend, gov.owner, towei(2_000_000), &Bv.From)))

	// S2: voter-only address (Bv, stakerIdx==0) attempts a self-finalized change.
	// MUST be rejected ("Non-member") with no state mutation / slot-0 poisoning.
	infoExploit := w1gInfo(Bv.From, Bv.From, Bv.From, nodeX, "exploit")
	ExpectedRevert(t,
		gov.ExpectedFail(gov.GovImp.Transact(Bv, "addProposalToChangeMember", infoExploit, Bv.From, big.NewInt(0), big.NewInt(0))),
		"Non-member",
	)
	checkGovMember(t, gov, bIdx, infoBsep) // B entry unchanged
	require.True(t, w1gBool(t, gov, "isVoter", Bv.From))
	require.False(t, w1gBool(t, gov, "isStaker", Bv.From))
	{ // slot 0 stays empty (no node/role written into the sentinel index)
		var m0, v0 common.Address
		require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&m0}, "getMember", big.NewInt(0)))
		require.NoError(t, gov.GovImp.Call(callOpts, &[]interface{}{&v0}, "getVoter", big.NewInt(0)))
		require.Equal(t, common.Address{}, m0)
		require.Equal(t, common.Address{}, v0)
	}

	// S4: the real staker B can still self-update its own node (self-finalized)
	infoBnode := w1gInfo(B.From, Bv.From, Br.From, nodeY, "update B node")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(B, "addProposalToChangeMember", infoBnode, B.From, big.NewInt(0), big.NewInt(0))))
	checkGovMember(t, gov, bIdx, infoBnode)

	// S3: the separated VOTER (Bv) can both PROPOSE and VOTE (operational path).
	// Bv proposes adding C; owner + Bv (B's voter) vote -> 2/2 stakers accept -> C added.
	w1gFundAndStake(t, gov, C)
	infoC := w1gInfo(C.From, C.From, C.From, nodeC, "add C")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(Bv, "addProposalToAddMember", infoC)))
	cBallot := getBallotIdx(t, gov)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "vote", cBallot, true)))
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(Bv, "vote", cBallot, true)))
	cIdx := getMemberIdx(t, gov, C.From)
	checkGovMember(t, gov, cIdx, infoC)
	require.Equal(t, int64(3), w1gMemberLength(t, gov))
}

// D1: a VOTED full member change (proposed by another member, finalized by vote)
// of a previously-separated member must clear the SEPARATED voter/reward keys
// (read from the member index, not from oldStaker) and leave the ledger consistent.
func TestW1G04_VotedChangeOfSeparatedMember(t *testing.T) {
	gov := NewGovernance(t).DeployContracts(t)
	owner := gov.owner
	B := getTxOpt(t, "d1_B")
	Bv := getTxOpt(t, "d1_Bv")
	Br := getTxOpt(t, "d1_Br")
	T := getTxOpt(t, "d1_T")
	nodeB := nodeInfo{[]byte("d1-B"), w1gEnode("a1a1a1a1"), []byte("127.0.1.1"), big.NewInt(8542)}
	nodeT := nodeInfo{[]byte("d1-T"), w1gEnode("a2a2a2a2"), []byte("127.0.1.2"), big.NewInt(8542)}

	// add B, then B self-separates voter & reward
	w1gFundAndStake(t, gov, B)
	infoB := w1gInfo(B.From, B.From, B.From, nodeB, "add B")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", infoB)))
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "vote", getBallotIdx(t, gov), true)))
	bIdx := getMemberIdx(t, gov, B.From)
	require.NoError(t, gov.ExpectedOk(TransferCoin(gov.backend, owner, towei(2_000_000), &Bv.From))) // voter gas
	infoBsep := w1gInfo(B.From, Bv.From, Br.From, nodeB, "separate B")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(B, "addProposalToChangeMember", infoBsep, B.From, big.NewInt(0), big.NewInt(0))))
	checkGovMember(t, gov, bIdx, infoBsep)

	// VOTED full change B -> T (msg.sender=owner != B, so it goes through voting)
	w1gFundAndStake(t, gov, T)
	infoT := w1gInfo(T.From, T.From, T.From, nodeT, "change B to T")
	// unlockAmount = LOCK_AMOUNT (== minStaking) so the exit takes the plain-unlock
	// branch in transferLockedAndUnlock (matches the existing change-totally test).
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToChangeMember", infoT, B.From, LOCK_AMOUNT, big.NewInt(0))))
	chBallot := getBallotIdx(t, gov)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "vote", chBallot, true)))
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(Bv, "vote", chBallot, true))) // B votes via its separated voter

	// slot now holds T; every old (separated) key is cleared, no stale mappings
	checkGovMember(t, gov, bIdx, infoT)
	require.False(t, w1gBool(t, gov, "isStaker", B.From))
	require.False(t, w1gBool(t, gov, "isVoter", Bv.From))
	require.False(t, w1gBool(t, gov, "isReward", Br.From))
	require.Equal(t, int64(2), w1gMemberLength(t, gov))
}

// D2: the VOTED change path must also reject a voter-only target at creation
// ("Non-member"), complementing the self-finalized case (S2 in the lifecycle test).
func TestW1G04_VotedPathRejectsVoterOnlyTarget(t *testing.T) {
	gov := NewGovernance(t).DeployContracts(t)
	owner := gov.owner
	B := getTxOpt(t, "d2_B")
	Bv := getTxOpt(t, "d2_Bv")
	nodeB := nodeInfo{[]byte("d2-B"), w1gEnode("b1b1b1b1"), []byte("127.0.2.1"), big.NewInt(8542)}
	nodeX := nodeInfo{[]byte("d2-X"), w1gEnode("b2b2b2b2"), []byte("127.0.2.2"), big.NewInt(8542)}

	w1gFundAndStake(t, gov, B)
	infoB := w1gInfo(B.From, B.From, B.From, nodeB, "add B")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", infoB)))
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "vote", getBallotIdx(t, gov), true)))

	// B separates its voter -> Bv (voter-only, stakerIdx==0)
	infoBsep := w1gInfo(B.From, Bv.From, B.From, nodeB, "separate voter")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(B, "addProposalToChangeMember", infoBsep, B.From, big.NewInt(0), big.NewInt(0))))
	bIdx := getMemberIdx(t, gov, B.From)

	// owner proposes a voted change targeting the voter-only address Bv -> rejected at creation
	infoX := w1gInfo(Bv.From, Bv.From, Bv.From, nodeX, "voted exploit")
	ExpectedRevert(t,
		gov.ExpectedFail(gov.GovImp.Transact(owner, "addProposalToChangeMember", infoX, Bv.From, big.NewInt(0), big.NewInt(0))),
		"Non-member",
	)
	checkGovMember(t, gov, bIdx, infoBsep) // unchanged
}

// D3: reward-only separation. The staker keeps staker==voter but moves the reward
// to a distinct address; the reward index must move while staker/voter stay, and a
// later self-update must keep the separated reward consistent.
func TestW1G04_RewardSeparationLedger(t *testing.T) {
	gov := NewGovernance(t).DeployContracts(t)
	owner := gov.owner
	B := getTxOpt(t, "d3_B")
	Br := getTxOpt(t, "d3_Br")
	nodeB := nodeInfo{[]byte("d3-B"), w1gEnode("c1c1c1c1"), []byte("127.0.3.1"), big.NewInt(8542)}
	nodeY := nodeInfo{[]byte("d3-Y"), w1gEnode("c2c2c2c2"), []byte("127.0.3.2"), big.NewInt(8542)}

	w1gFundAndStake(t, gov, B)
	infoB := w1gInfo(B.From, B.From, B.From, nodeB, "add B")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", infoB)))
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "vote", getBallotIdx(t, gov), true)))
	bIdx := getMemberIdx(t, gov, B.From)

	// separate ONLY the reward (staker==voter==B, reward=Br)
	infoBsep := w1gInfo(B.From, B.From, Br.From, nodeB, "separate reward")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(B, "addProposalToChangeMember", infoBsep, B.From, big.NewInt(0), big.NewInt(0))))
	checkGovMember(t, gov, bIdx, infoBsep)
	require.True(t, w1gBool(t, gov, "isStaker", B.From))
	require.True(t, w1gBool(t, gov, "isVoter", B.From))
	require.True(t, w1gBool(t, gov, "isReward", Br.From))
	require.False(t, w1gBool(t, gov, "isReward", B.From))

	// a subsequent self-update (node change) keeps the separated reward consistent
	infoBnode := w1gInfo(B.From, B.From, Br.From, nodeY, "update node keep sep reward")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(B, "addProposalToChangeMember", infoBnode, B.From, big.NewInt(0), big.NewInt(0))))
	checkGovMember(t, gov, bIdx, infoBnode)
	require.True(t, w1gBool(t, gov, "isReward", Br.From))
	require.False(t, w1gBool(t, gov, "isReward", B.From))
}

// ============================================================================
// W1G-04 — voter-only address self-change.  RED on legacy, GREEN after upgrade.
//
// A staker B self-separates its voter key Bv (legitimate). Bv then has
// voterIdx[Bv]!=0 but stakerIdx[Bv]==0. On the LEGACY contract Bv can drive a
// self-finalized "change" of itself: changeMember indexes by stakerIdx[Bv]==0
// and writes the sentinel slot 0 — poisoning voters[0], destroying Bv's own
// voter mapping, and squatting attacker-chosen node markers (a DoS on future
// legitimate registration). The fixed contract rejects it at proposal time.
// ============================================================================
func TestW1G04_Legacy_Slot0Poison_RedThenGreen(t *testing.T) {
	nodeB := nodeInfo{[]byte("w04-B"), w1gEnode("0b0b0b0b"), []byte("10.4.0.2"), big.NewInt(8542)}
	// attacker-chosen identifier the exploit squats; a legitimate node later wants it.
	nodeX := nodeInfo{[]byte("w04-X"), w1gEnode("0a0a0a0a"), []byte("10.4.0.9"), big.NewInt(9999)}

	setupSeparated := func(t *testing.T, gov *Governance) (B, Bv, Br *bind.TransactOpts, bIdx *big.Int) {
		owner := gov.owner
		B = getTxOpt(t, "w04_B")
		Bv = getTxOpt(t, "w04_Bv")
		Br = getTxOpt(t, "w04_Br")
		bIdx = addPlainMember(t, gov, owner, B, nodeB, "add B", owner)
		infoBsep := w1gInfo(B.From, Bv.From, Br.From, nodeB, "separate B")
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(B, "addProposalToChangeMember", infoBsep, B.From, big.NewInt(0), big.NewInt(0))))
		require.True(t, w1gBool(t, gov, "isVoter", Bv.From))
		require.False(t, w1gBool(t, gov, "isStaker", Bv.From))
		require.NoError(t, gov.ExpectedOk(TransferCoin(gov.backend, owner, towei(2_000_000), &Bv.From)))
		return
	}

	t.Run("RED_legacy", func(t *testing.T) {
		gov := NewGovernance(t).DeployContractsLegacy(t)
		owner := gov.owner
		B, Bv, _, _ := setupSeparated(t, gov)

		// EXPLOIT: voter-only Bv drives a self-finalized change of itself.
		infoExploit := w1gInfo(Bv.From, Bv.From, Bv.From, nodeX, "exploit via voter-only key")
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(Bv, "addProposalToChangeMember", infoExploit, Bv.From, big.NewInt(0), big.NewInt(0))))

		// (1) Bv's own voter mapping was destroyed (voterIdx[Bv] overwritten to 0).
		require.False(t, w1gBool(t, gov, "isVoter", Bv.From), "RED: Bv voter mapping should be corrupted on legacy")
		// (2) sentinel slot 0 poisoned with Bv as a voter.
		require.Equal(t, Bv.From, getVoterAt(t, gov, big.NewInt(0)), "RED: voters[0] poisoned on legacy")
		// (3) attacker squatted nodeX's markers -> a legitimate member can no longer use nodeX (DoS).
		C := getTxOpt(t, "w04_C")
		w1gFundAndStake(t, gov, C)
		infoLegit := w1gInfo(C.From, C.From, C.From, nodeX, "legit C wants nodeX")
		ExpectedRevert(t,
			gov.ExpectedFail(gov.GovImp.Transact(owner, "addProposalToAddMember", infoLegit)),
			"Duplicated node info",
		)
		// B's real entry is otherwise intact (bIdx still maps to B as staker).
		require.True(t, w1gBool(t, gov, "isStaker", B.From))
	})

	t.Run("GREEN_after_upgrade", func(t *testing.T) {
		gov := NewGovernance(t).DeployContractsLegacy(t)
		owner := gov.owner
		B, Bv, _, bIdx := setupSeparated(t, gov)

		upgradeGovToFixed(t, gov, owner, owner, B) // 2 members: owner + B

		// exploit now rejected at proposal time ("Non-member": Bv is not a staker).
		infoExploit := w1gInfo(Bv.From, Bv.From, Bv.From, nodeX, "exploit via voter-only key")
		ExpectedRevert(t,
			gov.ExpectedFail(gov.GovImp.Transact(Bv, "addProposalToChangeMember", infoExploit, Bv.From, big.NewInt(0), big.NewInt(0))),
			"Non-member",
		)
		require.True(t, w1gBool(t, gov, "isVoter", Bv.From), "GREEN: Bv voter mapping intact")
		require.Equal(t, common.Address{}, getVoterAt(t, gov, big.NewInt(0)), "GREEN: slot 0 clean")
		_ = bIdx

		// and nodeX is still available to a legitimate member (no squatting happened).
		C := getTxOpt(t, "w04_C")
		w1gFundAndStake(t, gov, C)
		infoLegit := w1gInfo(C.From, C.From, C.From, nodeX, "legit C uses nodeX")
		require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToAddMember", infoLegit)))
		voteAccept(t, gov, getBallotIdx(t, gov), owner, B)
		require.True(t, w1gBool(t, gov, "isStaker", C.From))
	})
}

// ==========================================================================
// End-to-end — legacy ledger -> proxy upgrade -> exercise every core flow
// ==========================================================================

// ============================================================================
// Checks 2 / 3 / 5 — Legacy ledger -> upgrade -> exercise EVERY core flow.
//
// Builds a 3-member legacy ledger where member B has a self-separated
// staker/voter/reward, upgrades the proxy to the fixed GovImp, then drives
// propose / vote / execute for add, change (self & voted), env, and remove —
// initiated and voted via BOTH staker and the separated voter address — and
// asserts the W1G-04 staker/voter identification rule (passing a voter-only
// address as oldStaker reverts "Non-member", passing the staker works).
// ============================================================================
func TestW1G_LegacyUpgrade_FullLifecycle(t *testing.T) {
	gov := NewGovernance(t).DeployContractsLegacy(t)
	owner := gov.owner // member 1 (staker==voter==reward)

	B := getTxOpt(t, "lc_B")
	Bv := getTxOpt(t, "lc_Bv")
	Br := getTxOpt(t, "lc_Br")
	C := getTxOpt(t, "lc_C")
	D := getTxOpt(t, "lc_D")
	T := getTxOpt(t, "lc_T")

	nodeB := nodeInfo{[]byte("lc-B"), w1gEnode("b0b0b0b0"), []byte("10.1.0.2"), big.NewInt(8542)}
	nodeBy := nodeInfo{[]byte("lc-B2"), w1gEnode("b1b1b1b1"), []byte("10.1.0.3"), big.NewInt(8542)}
	nodeC := nodeInfo{[]byte("lc-C"), w1gEnode("c0c0c0c0"), []byte("10.1.0.4"), big.NewInt(8542)}
	nodeD := nodeInfo{[]byte("lc-D"), w1gEnode("d0d0d0d0"), []byte("10.1.0.5"), big.NewInt(8542)}
	nodeT := nodeInfo{[]byte("lc-T"), w1gEnode("70707070"), []byte("10.1.0.6"), big.NewInt(8542)}

	// --- build the LEGACY ledger (pre-fix contract) ------------------------
	require.Equal(t, int64(1), w1gMemberLength(t, gov))

	// member B (plain), then C (plain): owner is sole member -> owner's vote finalizes.
	bIdx := addPlainMember(t, gov, owner, B, nodeB, "add B (legacy)", owner)
	require.Equal(t, int64(2), w1gMemberLength(t, gov))
	cIdx := addPlainMember(t, gov, owner, C, nodeC, "add C (legacy)", owner, B)
	require.Equal(t, int64(3), w1gMemberLength(t, gov))

	// B self-separates voter+reward on the LEGACY contract (self-finalized).
	infoBsep := w1gInfo(B.From, Bv.From, Br.From, nodeB, "separate B (legacy)")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(B, "addProposalToChangeMember", infoBsep, B.From, big.NewInt(0), big.NewInt(0))))
	checkGovMember(t, gov, bIdx, infoBsep)
	require.True(t, w1gBool(t, gov, "isVoter", Bv.From))
	require.False(t, w1gBool(t, gov, "isStaker", Bv.From))
	require.Equal(t, B.From, w1gStakerAddr(t, gov, Bv.From))
	require.NoError(t, gov.ExpectedOk(TransferCoin(gov.backend, owner, towei(2_000_000), &Bv.From))) // gas for the voter key

	// --- UPGRADE proxy to the fixed GovImp (real governance path) + reInit --
	// quorum among {owner, B, C}: owner proposes, owner + B(via staker) vote.
	upgradeGovToFixed(t, gov, owner, owner, B, C)

	// ledger survived the upgrade unchanged (state preserved through proxy).
	checkGovMember(t, gov, bIdx, infoBsep)
	checkGovMember(t, gov, cIdx, w1gInfo(C.From, C.From, C.From, nodeC, ""))
	require.Equal(t, int64(3), w1gMemberLength(t, gov))

	// ====================================================================
	// FLOW 1 — propose ADD via the separated VOTER (Bv), execute via vote.
	//          vote with a mix of staker (owner, C) and voter (Bv) keys.
	// ====================================================================
	w1gFundAndStake(t, gov, D)
	infoD := w1gInfo(D.From, D.From, D.From, nodeD, "add D via voter Bv")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(Bv, "addProposalToAddMember", infoD))) // voter proposes
	dBallot := getBallotIdx(t, gov)
	voteAccept(t, gov, dBallot, owner, Bv, C) // owner(staker)+Bv(voter for B)+C(staker)
	dIdx := getMemberIdx(t, gov, D.From)
	checkGovMember(t, gov, dIdx, infoD)
	require.Equal(t, int64(4), w1gMemberLength(t, gov))

	// ====================================================================
	// FLOW 2 — Check 5: oldStaker identification on the FIXED contract.
	//   (a) passing the separated VOTER address Bv as oldStaker -> "Non-member"
	//   (b) passing the real STAKER address B as oldStaker -> OK (self-update)
	// ====================================================================
	infoViaVoter := w1gInfo(B.From, Bv.From, Br.From, nodeBy, "change via voter key (must fail)")
	ExpectedRevert(t,
		gov.ExpectedFail(gov.GovImp.Transact(B, "addProposalToChangeMember", infoViaVoter, Bv.From, big.NewInt(0), big.NewInt(0))),
		"Non-member",
	)
	checkGovMember(t, gov, bIdx, infoBsep) // unchanged by the rejected attempt

	// (b) same change but oldStaker = the real staker B -> self-finalized OK.
	infoBupd := w1gInfo(B.From, Bv.From, Br.From, nodeBy, "B self-updates node")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(B, "addProposalToChangeMember", infoBupd, B.From, big.NewInt(0), big.NewInt(0))))
	checkGovMember(t, gov, bIdx, infoBupd)

	// ====================================================================
	// FLOW 3 — VOTED full member change C -> T (proposed by the voter Bv),
	//          voted by staker(owner) + voter(Bv). Exercises the voted path.
	// ====================================================================
	w1gFundAndStake(t, gov, T)
	infoT := w1gInfo(T.From, T.From, T.From, nodeT, "change C to T")
	// unlockAmount == minStaking so transferLockedAndUnlock takes the plain-unlock branch.
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(Bv, "addProposalToChangeMember", infoT, C.From, LOCK_AMOUNT, big.NewInt(0))))
	chBallot := getBallotIdx(t, gov)
	voteAccept(t, gov, chBallot, owner, Bv, D) // staker + voter + staker
	checkGovMember(t, gov, cIdx, infoT)        // slot now holds T
	require.False(t, w1gBool(t, gov, "isStaker", C.From))
	require.True(t, w1gBool(t, gov, "isStaker", T.From))
	require.Equal(t, int64(4), w1gMemberLength(t, gov))

	// ====================================================================
	// FLOW 4 — propose+vote+execute an ENV change (proposed by staker owner).
	//          Sets blocksPer to 100 and verifies the value actually applied.
	// ====================================================================
	var blocksPer *big.Int
	require.NoError(t, gov.EnvStorageImp.Call(new(bind.CallOpts), &[]interface{}{&blocksPer}, "getBlocksPer"))
	require.NotEqual(t, big.NewInt(100), blocksPer)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(owner, "addProposalToChangeEnv",
		crypto.Keccak256Hash([]byte("blocksPer")),
		EnvTypes.Uint,
		hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 100
		[]byte("env memo"), big.NewInt(86400))))
	envBallot := getBallotIdx(t, gov)
	voteAccept(t, gov, envBallot, owner, Bv, D)
	require.NoError(t, gov.EnvStorageImp.Call(new(bind.CallOpts), &[]interface{}{&blocksPer}, "getBlocksPer"))
	require.Equal(t, big.NewInt(100), blocksPer) // env change applied through the fixed contract

	// ====================================================================
	// FLOW 5 — propose+vote+execute REMOVE of member D (proposed by voter Bv).
	// ====================================================================
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(
		Bv, "addProposalToRemoveMember", D.From, LOCK_AMOUNT, []byte("remove D"), big.NewInt(86400), LOCK_AMOUNT, big.NewInt(0))))
	rmBallot := getBallotIdx(t, gov)
	voteAccept(t, gov, rmBallot, owner, Bv, T) // owner + B(via Bv) + T
	require.False(t, w1gBool(t, gov, "isStaker", D.From))
	require.Equal(t, int64(3), w1gMemberLength(t, gov))

	// final ledger consistency for the surviving members.
	checkGovMember(t, gov, bIdx, infoBupd)
}

// ============================================================================
// Check 2 (supplement) — voter-driven ENV and GOV proposals.
//
// Complements the FullLifecycle test (where the separated voter drives add /
// change / remove). Here the separated voter key Bv ALSO proposes — and helps
// vote in — an env change and a governance (implementation) change, the two
// proposal kinds that carry no member-identification argument. Confirms a
// voter-only key drives every proposal kind end to end after the upgrade.
// ============================================================================
func TestW1G_VoterDrivenEnvAndGovProposals(t *testing.T) {
	gov := NewGovernance(t).DeployContractsLegacy(t)
	owner := gov.owner
	B := getTxOpt(t, "vd_B")
	Bv := getTxOpt(t, "vd_Bv")
	Br := getTxOpt(t, "vd_Br")
	nodeB := nodeInfo{[]byte("vd-B"), w1gEnode("08080808"), []byte("10.5.0.2"), big.NewInt(8542)}

	// legacy ledger: owner(1), B(2) self-separated to Bv/Br.
	bIdx := addPlainMember(t, gov, owner, B, nodeB, "add B", owner)
	infoBsep := w1gInfo(B.From, Bv.From, Br.From, nodeB, "separate B")
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(B, "addProposalToChangeMember", infoBsep, B.From, big.NewInt(0), big.NewInt(0))))
	checkGovMember(t, gov, bIdx, infoBsep)
	require.NoError(t, gov.ExpectedOk(TransferCoin(gov.backend, owner, towei(2_000_000), &Bv.From))) // gas for the voter key

	// upgrade to the fixed impl (owner + B vote) + reInit. members: {owner, B}.
	upgradeGovToFixed(t, gov, owner, owner, B)

	// ---- ENV change PROPOSED BY THE VOTER Bv, voted by owner + Bv ----------
	var blocksPer *big.Int
	require.NoError(t, gov.EnvStorageImp.Call(new(bind.CallOpts), &[]interface{}{&blocksPer}, "getBlocksPer"))
	require.NotEqual(t, big.NewInt(100), blocksPer)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(Bv, "addProposalToChangeEnv",
		crypto.Keccak256Hash([]byte("blocksPer")),
		EnvTypes.Uint,
		hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 100
		[]byte("env via voter"), big.NewInt(86400))))
	voteAccept(t, gov, getBallotIdx(t, gov), owner, Bv)
	require.NoError(t, gov.EnvStorageImp.Call(new(bind.CallOpts), &[]interface{}{&blocksPer}, "getBlocksPer"))
	require.Equal(t, big.NewInt(100), blocksPer) // env change applied via a voter-proposed ballot

	// ---- GOV (implementation) change PROPOSED BY THE VOTER Bv -------------
	nextImp, _, err := gov.Deploy(compiled.GovImp.Deploy(gov.backend, owner))
	require.NoError(t, err)
	require.NotEqual(t, govImplementation(t, gov), nextImp)
	require.NoError(t, gov.ExpectedOk(gov.GovImp.Transact(Bv, "addProposalToChangeGov", nextImp, []byte("gov via voter"), big.NewInt(86400))))
	voteAccept(t, gov, getBallotIdx(t, gov), owner, Bv)
	require.Equal(t, nextImp, govImplementation(t, gov), "voter-proposed governance change must apply")

	// ledger still intact after both voter-driven proposals.
	checkGovMember(t, gov, bIdx, infoBsep)
	require.Equal(t, int64(2), w1gMemberLength(t, gov))
}
