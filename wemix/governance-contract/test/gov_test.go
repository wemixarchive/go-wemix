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
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}
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

			node1 := gov.nodeInfos[1]

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
	})
}
