package test

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	sim "github.com/ethereum/go-ethereum/wemix/governance-contract/common/simulated-backend"
	"github.com/stretchr/testify/require"
)

func TestGov(t *testing.T) {
	// for mute chain log
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(0), log.StreamHandler(os.Stdout, log.TerminalFormat(true))))

	t.Run("Staker is voter", func(t *testing.T) {
		t.Run("has enode and locked staking", func(t *testing.T) {
			gov := DeployGovernance(t)
			locked := gov.Staking.Call1(t, "lockedBalanceOf", gov.client.Owner).(*big.Int)
			require.Equal(t, LOCK_AMOUNT, locked)
			idx := gov.Gov.Call1(t, "getNodeIdxFromMember", gov.client.Owner).(*big.Int)
			require.True(t, idx.Sign() != 0)
			getNode := gov.Gov.Call(t, "getNode", idx)
			name, enode, ip, port := getNode[0].([]byte), getNode[1].([]byte), getNode[2].([]byte), getNode[3].(*big.Int)
			nodeinfo := gov.nodeInfos[0]
			require.Equal(t, nodeinfo.name, name)
			require.Equal(t, nodeinfo.enode, enode)
			require.Equal(t, nodeinfo.ip, ip)
			require.Equal(t, nodeinfo.port, port)
		})
		t.Run("cannot init twice", func(t *testing.T) {
			gov := DeployGovernance(t)
			node := gov.nodeInfos[0]
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "init", gov.Registry.Address, LOCK_AMOUNT, node.name, node.enode, node.ip, node.port),
				"Initializable: contract is already initialized",
			)
		})
		t.Run("cannot addProposal to add member self", func(t *testing.T) {
			gov := DeployGovernance(t)
			member := gov.client.Owner
			node := gov.nodeInfos[0]
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "addProposalToAddMember",
					MemberInfo{
						Staker:     member,
						Voter:      member,
						Reward:     member,
						Name:       node.name,
						Enode:      node.enode,
						Ip:         node.ip,
						Port:       node.port,
						LockAmount: LOCK_AMOUNT,
						Memo:       []byte("memo1"),
						Duration:   big.NewInt(86400),
					},
				), "Already member",
			)
		})
		t.Run("cannot addProposal to add member with different voter", func(t *testing.T) {
			gov := DeployGovernance(t)
			member := gov.client.Owner
			voter := sim.GetOrNewEOA("voter")
			node := gov.nodeInfos[0]
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "addProposalToAddMember",
					MemberInfo{
						Staker:     member,
						Voter:      voter,
						Reward:     member,
						Name:       node.name,
						Enode:      node.enode,
						Ip:         node.ip,
						Port:       node.port,
						LockAmount: LOCK_AMOUNT,
						Memo:       []byte("memo1"),
						Duration:   big.NewInt(86400),
					},
				), "Already member",
			)
		})
		t.Run("cannot addProposal to add member with different reward", func(t *testing.T) {
			gov := DeployGovernance(t)
			member := gov.client.Owner
			rewarder := sim.GetOrNewEOA("rewarder")
			node := gov.nodeInfos[0]
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "addProposalToAddMember",
					MemberInfo{
						Staker:     member,
						Voter:      member,
						Reward:     rewarder,
						Name:       node.name,
						Enode:      node.enode,
						Ip:         node.ip,
						Port:       node.port,
						LockAmount: LOCK_AMOUNT,
						Memo:       []byte("memo1"),
						Duration:   big.NewInt(86400),
					},
				), "Already member",
			)
		})
		t.Run("can addProposal to add member", func(t *testing.T) {
			gov := DeployGovernance(t)
			govMem1 := sim.GetOrNewEOA("govMem1")
			// staking first
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Add(LOCK_AMOUNT, towei(1)), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")
			// proposal
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}
			gov.nodeInfos = append(gov.nodeInfos, node1)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				},
			)
			// check proposal
			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big1, length)

			ballot := gov.BallotStorage.Call(t, "getBallotBasic", length)
			ballotDetail := gov.BallotStorage.Call(t, "getBallotMember", length)
			require.Equal(t, gov.client.Owner, ballot[3])
			require.Equal(t, []byte("memo1"), ballot[4])
			require.Equal(t, ballotDetail[1], govMem1)

			// govMem1 is not member yet
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo2"),
					Duration:   big.NewInt(86400),
				},
			)
			require.Equal(t, common.Big2, gov.Gov.Call1(t, "ballotLength").(*big.Int))
		})
		t.Run("cannot addProposal to remove non-member", func(t *testing.T) {
			gov := DeployGovernance(t)
			govMem1 := sim.AliasToAddress["govMem1"]
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "addProposalToRemoveMember", govMem1, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int)),
				"Non-member",
			)
		})
		t.Run("cannot addProposal to remove a sole member", func(t *testing.T) {
			gov := DeployGovernance(t)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "addProposalToRemoveMember", gov.client.Owner, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int)),
				"Cannot remove a sole member",
			)
		})
		t.Run("can addProposal to change member's other addresses self without voting", func(t *testing.T) {
			gov := DeployGovernance(t)
			oldMember := gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldMember)
			oldVoter := gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldVoter)
			oldReward := gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldReward)
			node := gov.nodeInfos[0]
			voter1 := sim.GetOrNewEOA("voter1")
			user1 := sim.GetOrNewEOA("user1")

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      voter1,
				Reward:     user1,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big1, length)
		})
		t.Run("can addProposal to change member's other addresses self without voting twice about node name", func(t *testing.T) {
			gov := DeployGovernance(t)
			oldMember := gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldMember)
			oldVoter := gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldVoter)
			oldReward := gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldReward)
			node := gov.nodeInfos[0]

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      gov.client.Owner,
				Reward:     gov.client.Owner,
				Name:       []byte("name1"), // change
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big1, length)

			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			stat := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, stat[1].(*big.Int))
			require.True(t, stat[2].(bool))

			newMember := gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newMember)
			newVoter := gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newVoter)
			newReward := gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newReward)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      gov.client.Owner,
				Reward:     gov.client.Owner,
				Name:       node.name, // change
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			length = gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big2, length)

			inVoting = gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			stat = gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, stat[1].(*big.Int))
			require.True(t, stat[2].(bool))

			newMember = gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newMember)
			newVoter = gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newVoter)
			newReward = gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newReward)
		})
		t.Run("can addProposal to change member's other addresses self without voting twice about enode", func(t *testing.T) {
			gov := DeployGovernance(t)
			oldMember := gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldMember)
			oldVoter := gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldVoter)
			oldReward := gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldReward)
			node := gov.nodeInfos[0]
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      gov.client.Owner,
				Reward:     gov.client.Owner,
				Name:       []byte("name1"),
				Enode:      node1.enode, // change
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big1, length)

			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			stat := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, stat[1].(*big.Int))
			require.True(t, stat[2].(bool))

			newMember := gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newMember)
			newVoter := gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newVoter)
			newReward := gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newReward)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      gov.client.Owner,
				Reward:     gov.client.Owner,
				Name:       node.name,
				Enode:      node.enode, // change
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			length = gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big2, length)

			inVoting = gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			stat = gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, stat[1].(*big.Int))
			require.True(t, stat[2].(bool))

			newMember = gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newMember)
			newVoter = gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newVoter)
			newReward = gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newReward)
		})
		t.Run("can addProposal to change member's other addresses self without voting twice about ipport", func(t *testing.T) {
			gov := DeployGovernance(t)
			oldMember := gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldMember)
			oldVoter := gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldVoter)
			oldReward := gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldReward)
			node := gov.nodeInfos[0]
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      gov.client.Owner,
				Reward:     gov.client.Owner,
				Name:       []byte("name1"),
				Enode:      node.enode,
				Ip:         node1.ip, // change
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big1, length)

			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			stat := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, stat[1].(*big.Int))
			require.True(t, stat[2].(bool))

			newMember := gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newMember)
			newVoter := gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newVoter)
			newReward := gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newReward)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      gov.client.Owner,
				Reward:     gov.client.Owner,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip, // change
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			length = gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big2, length)

			inVoting = gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			stat = gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, stat[1].(*big.Int))
			require.True(t, stat[2].(bool))

			newMember = gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newMember)
			newVoter = gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newVoter)
			newReward = gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newReward)
		})
		t.Run("can addProposal to change member's other addresses self without voting twice about import2", func(t *testing.T) {
			gov := DeployGovernance(t)
			oldMember := gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldMember)
			oldVoter := gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldVoter)
			oldReward := gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldReward)
			node := gov.nodeInfos[0]
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      gov.client.Owner,
				Reward:     gov.client.Owner,
				Name:       []byte("name1"),
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node1.port, // change
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big1, length)

			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			stat := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, stat[1].(*big.Int))
			require.True(t, stat[2].(bool))

			newMember := gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newMember)
			newVoter := gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newVoter)
			newReward := gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newReward)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      gov.client.Owner,
				Reward:     gov.client.Owner,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port, // change
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			length = gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big2, length)

			inVoting = gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			stat = gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, stat[1].(*big.Int))
			require.True(t, stat[2].(bool))

			newMember = gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newMember)
			newVoter = gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newVoter)
			newReward = gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newReward)
		})
		t.Run("cannot addProposal to change non-member", func(t *testing.T) {
			gov := DeployGovernance(t)
			govMem1, govMem2 := sim.GetOrNewEOA("govMem1"), sim.GetOrNewEOA("govMem2")
			node := gov.nodeInfos[0]

			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node.name,
					Enode:      node.enode,
					Ip:         node.ip,
					Port:       node.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				}, govMem2, common.Big0, common.Big0),
				"Non-member",
			)
		})
		t.Run("can addProposal to change governance", func(t *testing.T) {
			gov := DeployGovernance(t)
			newGovImp := compiled.Copy(new(Governance)).GovImp
			newGovImp.Deploy(t, gov.client)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeGov", newGovImp.Address, []byte("memo"), big.NewInt(86400))
			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big1, length)
		})
		t.Run("can not addProposal to change governance using EOA", func(t *testing.T) {
			gov := DeployGovernance(t)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "addProposalToChangeGov", sim.GetOrNewEOA("govMem1"), []byte("memo"), big.NewInt(86400)),
				"", "function call to a non-contract account", // TODO check
			)
		})
		t.Run("can not addProposal to change governance using non-UUPS", func(t *testing.T) {
			gov := DeployGovernance(t)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "addProposalToChangeGov", gov.Registry.Address, []byte("memo"), big.NewInt(86400)),
				"ERC1967Upgrade: new implementation is not UUPS",
			)
		})
		t.Run("cannot addProposal to change governance with same address", func(t *testing.T) {
			gov := DeployGovernance(t)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "addProposalToChangeGov", gov.GovImp.Address, []byte("memo"), big.NewInt(86400)),
				"Same contract address",
			)
		})
		t.Run("cannot addProposal to change governance with zero address", func(t *testing.T) {
			gov := DeployGovernance(t)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "addProposalToChangeGov", common.Address{}, []byte("memo"), big.NewInt(86400)),
				"Implementation cannot be zero",
			)
		})
		t.Run("can addProposal to change environment", func(t *testing.T) {
			gov := DeployGovernance(t)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeEnv", sim.StringToBytes32("key"), EnvTypes.Bytes32, []byte("value"), []byte("memo"), big.NewInt(86400))
			length := gov.Gov.Call1(t, "ballotLength")
			require.Equal(t, common.Big1, length)
		})
		t.Run("cannot addProposal to change environment with wrong type", func(t *testing.T) {
			gov := DeployGovernance(t)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "addProposalToChangeEnv", sim.StringToBytes32("key"), EnvTypes.Invalid, []byte("value"), []byte("memo"), big.NewInt(86400)),
				"Invalid type",
			)
		})
		t.Run("can vote approval to add member", func(t *testing.T) {
			gov := DeployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				},
			)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big1, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big1)
			require.Equal(t, BallotStates.Accepted, state[1].(*big.Int))
			require.True(t, state[2].(bool))
			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big2, memberLen)
			nodeLen := gov.Gov.Call1(t, "getNodeLength").(*big.Int)
			require.Equal(t, memberLen, nodeLen)
			lock := gov.Staking.Call1(t, "lockedBalanceOf", govMem1)
			require.Equal(t, LOCK_AMOUNT, lock)

			newMember := gov.Gov.Call1(t, "getMember", memberLen).(common.Address)
			require.Equal(t, govMem1, newMember)
			newVoter := gov.Gov.Call1(t, "getVoter", memberLen).(common.Address)
			require.Equal(t, govMem1, newVoter)
			newReward := gov.Gov.Call1(t, "getReward", memberLen).(common.Address)
			require.Equal(t, govMem1, newReward)
		})
		t.Run("cannot vote approval to add member with insufficient staking", func(t *testing.T) {
			gov := DeployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, new(big.Int).Div(LOCK_AMOUNT, common.Big2), "deposit")
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				},
			)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big1, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big1)
			require.Equal(t, BallotStates.Rejected, state[1].(*big.Int))
			require.True(t, state[2].(bool))
			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big1, memberLen)
		})
		t.Run("can vote disapproval to deny adding member", func(t *testing.T) {
			gov := DeployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				},
			)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big1, false)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big1)
			require.Equal(t, BallotStates.Rejected, state[1].(*big.Int))
			require.True(t, state[2].(bool))
		})
		t.Run("can vote approval to change member totally", func(t *testing.T) {
			gov := DeployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				},
			)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big1, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big1)
			require.Equal(t, BallotStates.Accepted, state[1].(*big.Int))
			require.True(t, state[2].(bool))

			govMem2 := sim.GetOrNewEOA("govMem2")
			_, receipt = gov.client.SendTransaction(t, gov.client.Owner, govMem2, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem2, LOCK_AMOUNT, "deposit")

			preDeployerAvail := gov.Staking.Call1(t, "availableBalanceOf", govMem1).(*big.Int)
			preGovmem1Avail := gov.Staking.Call1(t, "availableBalanceOf", govMem2).(*big.Int)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember",
				MemberInfo{
					Staker:     govMem2,
					Voter:      govMem2,
					Reward:     govMem2,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo2"),
					Duration:   big.NewInt(86400),
				}, govMem1, LOCK_AMOUNT, common.Big0,
			)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big2, true)
			gov.Gov.ExecuteOk(t, govMem1, "vote", common.Big2, true)
			length = gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, big.NewInt(3), length)
			inVoting = gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state = gov.BallotStorage.Call(t, "getBallotState", common.Big2)
			require.Equal(t, BallotStates.Accepted, state[1].(*big.Int))
			require.True(t, state[2].(bool))

			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big2, memberLen)
			memberAddr := gov.Gov.Call1(t, "getMember", common.Big2).(common.Address)
			require.Equal(t, govMem2, memberAddr)
			getNode := gov.Gov.Call(t, "getNode", common.Big2)
			name, enode, ip, port := getNode[0].([]byte), getNode[1].([]byte), getNode[2].([]byte), getNode[3].(*big.Int)
			require.Equal(t, node1.name, name)
			require.Equal(t, node1.enode, enode)
			require.Equal(t, node1.ip, ip)
			require.Equal(t, node1.port, port)

			nodeIdxFromDeployer := gov.Gov.Call1(t, "getNodeIdxFromMember", govMem1).(*big.Int)
			require.True(t, nodeIdxFromDeployer.Sign() == 0)
			nodeIdxFromGovMem1 := gov.Gov.Call1(t, "getNodeIdxFromMember", govMem2).(*big.Int)
			require.Equal(t, common.Big2, nodeIdxFromGovMem1)

			postDeployerAvail := gov.Staking.Call1(t, "availableBalanceOf", govMem1).(*big.Int)
			postGovmem1Avail := gov.Staking.Call1(t, "availableBalanceOf", govMem2).(*big.Int)
			require.Equal(t, LOCK_AMOUNT, new(big.Int).Sub(postDeployerAvail, preDeployerAvail))
			require.Equal(t, LOCK_AMOUNT, new(big.Int).Sub(preGovmem1Avail, postGovmem1Avail))
		})
		t.Run("can vote approval to change enode only without voting", func(t *testing.T) {
			gov := DeployGovernance(t)

			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      gov.client.Owner,
				Reward:     gov.client.Owner,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.True(t, length.Sign() == 0)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big1)
			require.Equal(t, BallotStates.Accepted, state[1])
			require.True(t, state[2].(bool))

			memberLen := gov.Gov.Call1(t, ("getMemberLength")).(*big.Int)
			require.Equal(t, common.Big1, memberLen)
			memberAddr := gov.Gov.Call1(t, "getMember", common.Big1)
			require.Equal(t, gov.client.Owner, memberAddr)
			getNode := gov.Gov.Call(t, "getNode", common.Big1)
			name, enode, ip, port := getNode[0].([]byte), getNode[1].([]byte), getNode[2].([]byte), getNode[3].(*big.Int)
			require.Equal(t, node1.name, name)
			require.Equal(t, node1.enode, enode)
			require.Equal(t, node1.ip, ip)
			require.Equal(t, node1.port, port)
		})
		t.Run("cannot vote approval to change member with insufficient staking", func(t *testing.T) {
			gov := DeployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, new(big.Int).Sub(LOCK_AMOUNT, big.NewInt(1000000000)), "deposit")

			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     govMem1,
				Voter:      govMem1,
				Reward:     govMem1,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big1, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big1)
			require.Equal(t, BallotStates.Rejected, state[1])
			require.True(t, state[2].(bool))
		})
		t.Run("can vote approval to change governance", func(t *testing.T) {
			gov := DeployGovernance(t)

			result := gov.EnvStorage.Call(t, "getGasLimitAndBaseFee")
			MBF := gov.EnvStorage.Call1(t, "getMaxBaseFee").(*big.Int)

			newGovImp := compiled.Copy(new(Governance)).GovImp
			newGovImp.Deploy(t, gov.client)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeGov", newGovImp.Address, []byte("memo"), big.NewInt(86400))
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big1, true)

			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big1)
			require.Equal(t, BallotStates.Accepted, state[1])
			require.True(t, state[2].(bool))

			imp := gov.Gov.Call1(t, "implementation").(common.Address)
			require.Equal(t, newGovImp.Address, imp)

			newResult := gov.EnvStorage.Call(t, "getGasLimitAndBaseFee")
			require.Equal(t, result[0], newResult[0])
			require.Equal(t, result[1], newResult[1])
			require.Equal(t, result[2], newResult[2])
			newMBF := gov.EnvStorage.Call1(t, "getMaxBaseFee").(*big.Int)
			require.Equal(t, MBF, newMBF)
		})
		t.Run("can vote approval to change environment", func(t *testing.T) {
			gov := DeployGovernance(t)

			blocksPer := gov.EnvStorage.Call1(t, "getBlocksPer").(*big.Int)
			require.NotEqual(t, big.NewInt(100), blocksPer)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				big.NewInt(86400),
			)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big1, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big1)
			require.Equal(t, BallotStates.Accepted, state[1])
			require.True(t, state[2].(bool))

			blocksPer = gov.EnvStorage.Call1(t, "getBlocksPer").(*big.Int)
			require.Equal(t, big.NewInt(100), blocksPer)
		})
		t.Run("cannot vote for a ballot already done", func(t *testing.T) {
			gov := DeployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, new(big.Int).Sub(LOCK_AMOUNT, big.NewInt(1000000000)), "deposit")

			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     govMem1,
				Voter:      govMem1,
				Reward:     govMem1,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big1, true)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "vote", common.Big1, true),
				"Expired",
			)
		})
		t.Run("cannot add proposal durring period time", func(t *testing.T) {
			gov := DeployGovernance(t)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				big.NewInt(86400),
			)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				big.NewInt(86400),
			)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "setProposalTimePeriod", big.NewInt(60))
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "addProposalToChangeEnv",
					crypto.Keccak256Hash([]byte("blocksPer")),
					EnvTypes.Uint,
					hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
					[]byte("memo"),
					big.NewInt(86400),
				), "Cannot add proposal too early",
			)
		})
		t.Run("cannot addProposal to add member which is already reward", func(t *testing.T) {
			gov := DeployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Add(LOCK_AMOUNT, towei(1)), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}
			gov.nodeInfos = append(gov.nodeInfos, node1)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				},
			)
			ballotLength := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big1, ballotLength)

			memberLength := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			node := gov.nodeInfos[0]
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember",
				MemberInfo{
					Staker:     gov.client.Owner,
					Voter:      gov.client.Owner,
					Reward:     govMem1,
					Name:       node.name,
					Enode:      node.enode,
					Ip:         node.ip,
					Port:       node.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo"),
					Duration:   big.NewInt(86400),
				}, gov.client.Owner, common.Big0, common.Big0,
			)

			ballotLength = gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big2, ballotLength)

			newMember := gov.Gov.Call1(t, "getMember", memberLength).(common.Address)
			require.Equal(t, gov.client.Owner, newMember)
			newVoter := gov.Gov.Call1(t, "getVoter", memberLength).(common.Address)
			require.Equal(t, gov.client.Owner, newVoter)
			newReward := gov.Gov.Call1(t, "getReward", memberLength).(common.Address)
			require.Equal(t, govMem1, newReward)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big1, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big1)
			require.Equal(t, BallotStates.Rejected, state[1])
			require.True(t, state[2].(bool))

			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big1, memberLen)
			nodeLen := gov.Gov.Call1(t, "getNodeLength").(*big.Int)
			require.Equal(t, memberLen, nodeLen)

			lock := gov.Staking.Call1(t, "lockedBalanceOf", govMem1).(*big.Int)
			require.True(t, lock.Sign() == 0)
			bal := gov.Staking.Call1(t, "balanceOf", govMem1).(*big.Int)
			require.Equal(t, LOCK_AMOUNT, bal)
		})
		t.Run("cannot addProposal to change member which is already reward", func(t *testing.T) {
			gov := DeployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Add(LOCK_AMOUNT, towei(1)), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")
			node := gov.nodeInfos[0]
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     govMem1,
				Voter:      govMem1,
				Reward:     govMem1,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      gov.client.Owner,
				Reward:     govMem1,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)

			newMember := gov.Gov.Call1(t, "getMember", memberLen).(common.Address)
			require.Equal(t, gov.client.Owner, newMember)
			newVoter := gov.Gov.Call1(t, "getVoter", memberLen).(common.Address)
			require.Equal(t, gov.client.Owner, newVoter)
			newReward := gov.Gov.Call1(t, "getReward", memberLen).(common.Address)
			require.Equal(t, govMem1, newReward)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big1, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big1)
			require.Equal(t, BallotStates.Rejected, state[1])
			require.True(t, state[2].(bool))

			memberLen = gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big1, memberLen)
			nodeLen := gov.Gov.Call1(t, "getNodeLength").(*big.Int)
			require.Equal(t, memberLen, nodeLen)
			lock := gov.Staking.Call1(t, "lockedBalanceOf", govMem1).(*big.Int)
			require.True(t, lock.Sign() == 0)
			bal := gov.Staking.Call1(t, "balanceOf", govMem1).(*big.Int)
			require.Equal(t, LOCK_AMOUNT, bal)
		})
	})
	t.Run("Staker is not a voter", func(t *testing.T) {
		deployGovernance := func(t *testing.T) (gov *Governance, voter common.Address) {
			gov = DeployGovernance(t)
			voter = sim.GetOrNewEOA("voter")
			user1 := sim.GetOrNewEOA("user1")
			balance, err := gov.client.Backend.BalanceAt(context.TODO(), gov.client.Owner, nil)
			require.NoError(t, err)
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, voter, new(big.Int).Div(balance, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

			node := gov.nodeInfos[0]
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      voter,
				Reward:     user1,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			return
		}
		t.Run("cannot addProposal to add member self", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			node := gov.nodeInfos[0]

			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToAddMember",
					MemberInfo{
						Staker:     gov.client.Owner,
						Voter:      gov.client.Owner,
						Reward:     gov.client.Owner,
						Name:       node.name,
						Enode:      node.enode,
						Ip:         node.ip,
						Port:       node.port,
						LockAmount: LOCK_AMOUNT,
						Memo:       []byte("memo"),
						Duration:   big.NewInt(86400),
					},
				), "Already member",
			)
		})
		t.Run("cannot addProposal to add member with different voter", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			node := gov.nodeInfos[0]

			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToAddMember",
					MemberInfo{
						Staker:     gov.client.Owner,
						Voter:      voter,
						Reward:     gov.client.Owner,
						Name:       node.name,
						Enode:      node.enode,
						Ip:         node.ip,
						Port:       node.port,
						LockAmount: LOCK_AMOUNT,
						Memo:       []byte("memo"),
						Duration:   big.NewInt(86400),
					},
				), "Already member",
			)
		})
		t.Run("cannot addProposal to add member with same voter", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			node := gov.nodeInfos[0]

			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToAddMember",
					MemberInfo{
						Staker:     voter,
						Voter:      voter,
						Reward:     voter,
						Name:       node.name,
						Enode:      node.enode,
						Ip:         node.ip,
						Port:       node.port,
						LockAmount: LOCK_AMOUNT,
						Memo:       []byte("memo"),
						Duration:   big.NewInt(86400),
					},
				), "Already member",
			)
		})
		t.Run("cannot addProposal to add member with same reward", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			user1 := sim.GetOrNewEOA("user1")
			node := gov.nodeInfos[0]

			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToAddMember",
					MemberInfo{
						Staker:     user1,
						Voter:      user1,
						Reward:     user1,
						Name:       node.name,
						Enode:      node.enode,
						Ip:         node.ip,
						Port:       node.port,
						LockAmount: LOCK_AMOUNT,
						Memo:       []byte("memo"),
						Duration:   big.NewInt(86400),
					},
				), "Already member",
			)
		})
		t.Run("can addProposal to add member", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			govMem1 := sim.GetOrNewEOA("govMem1")

			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}
			gov.Gov.ExecuteOk(t, voter, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo"),
					Duration:   big.NewInt(86400),
				},
			)

			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big2, length)
			ballot := gov.BallotStorage.Call(t, "getBallotBasic", length)
			ballotDetail := gov.BallotStorage.Call(t, "getBallotMember", length)
			require.Equal(t, voter, ballot[3])
			require.Equal(t, []byte("memo"), ballot[4])
			require.Equal(t, ballotDetail[1], govMem1)

			gov.Gov.ExecuteOk(t, voter, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				},
			)

			length = gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, big.NewInt(3), length)
		})
		t.Run("cannot addProposal to remove non-member", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			govMem1 := sim.AliasToAddress["govMem1"]
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToRemoveMember", govMem1, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int)),
				"Non-member",
			)
		})
		t.Run("cannot addProposal to remove a sole member", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToRemoveMember", gov.client.Owner, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int)),
				"Cannot remove a sole member",
			)
		})
		t.Run("can addProposal to change member's other addresses self without voting", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			oldMember := gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, oldMember)
			oldVoter := gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, voter, oldVoter)
			oldReward := gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, sim.GetOrNewEOA("user1"), oldReward)
			node := gov.nodeInfos[0]
			voter1 := sim.GetOrNewEOA("voter1")
			user1 := sim.GetOrNewEOA("user1")

			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      voter1,
				Reward:     user1,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big2, length)
			gov.Gov.ExecuteOk(t, voter, "vote", common.Big2, true)

			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, state[1])
			require.True(t, state[2].(bool))

			newMember := gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, newMember)
			newVoter := gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, voter1, newVoter)
			newReward := gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, user1, newReward)
		})
		t.Run("cannot addProposal to change member's other addresses which is already member", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, voter, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				},
			)
			gov.Gov.ExecuteOk(t, voter, "vote", common.Big2, true)

			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big2)
			require.Equal(t, BallotStates.Accepted, state[1].(*big.Int))
			require.True(t, state[2].(bool))
			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big2, memberLen)
			nodeLen := gov.Gov.Call1(t, "getNodeLength").(*big.Int)
			require.Equal(t, memberLen, nodeLen)
			lock := gov.Staking.Call1(t, "lockedBalanceOf", govMem1)
			require.Equal(t, LOCK_AMOUNT, lock)

			newMember := gov.Gov.Call1(t, "getMember", memberLen).(common.Address)
			require.Equal(t, govMem1, newMember)
			newVoter := gov.Gov.Call1(t, "getVoter", memberLen).(common.Address)
			require.Equal(t, govMem1, newVoter)
			newReward := gov.Gov.Call1(t, "getReward", memberLen).(common.Address)
			require.Equal(t, govMem1, newReward)

			node := gov.nodeInfos[0]
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToChangeMember", MemberInfo{
					Staker:     gov.client.Owner,
					Voter:      govMem1,
					Reward:     sim.GetOrNewEOA("user1"),
					Name:       node.name,
					Enode:      node.enode,
					Ip:         node.ip,
					Port:       node.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				}, gov.client.Owner, common.Big0, common.Big0),
				"Already a member",
			)
		})
		t.Run("cannot addProposal to add member which is already voter", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, voter, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				},
			)
			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)

			node := gov.nodeInfos[0]
			user1 := sim.GetOrNewEOA("user1")
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      govMem1,
				Reward:     user1,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)

			newMember := gov.Gov.Call1(t, "getMember", memberLen).(common.Address)
			require.Equal(t, gov.client.Owner, newMember)
			newVoter := gov.Gov.Call1(t, "getVoter", memberLen).(common.Address)
			require.Equal(t, govMem1, newVoter)
			newReward := gov.Gov.Call1(t, "getReward", memberLen).(common.Address)
			require.Equal(t, user1, newReward)

			gov.Gov.ExecuteOk(t, govMem1, "vote", common.Big2, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big2)
			require.Equal(t, BallotStates.Rejected, state[1])
			require.True(t, state[2].(bool))

			memberLen = gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big1, memberLen)
			nodeLen := gov.Gov.Call1(t, "getNodeLength").(*big.Int)
			require.Equal(t, memberLen, nodeLen)
			lock := gov.Staking.Call1(t, "lockedBalanceOf", govMem1).(*big.Int)
			require.True(t, lock.Sign() == 0)
			bal := gov.Staking.Call1(t, "balanceOf", govMem1)
			require.Equal(t, LOCK_AMOUNT, bal)
		})
		t.Run("cannot addProposal to change member which is already voter", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")

			node := gov.nodeInfos[0]
			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node.name,
					Enode:      node.enode,
					Ip:         node.ip,
					Port:       node.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				}, gov.client.Owner, common.Big0, common.Big0,
			)
			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			user1 := sim.GetOrNewEOA("user1")
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember",
				MemberInfo{
					Staker:     gov.client.Owner,
					Voter:      govMem1,
					Reward:     user1,
					Name:       node.name,
					Enode:      node.enode,
					Ip:         node.ip,
					Port:       node.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				}, gov.client.Owner, common.Big0, common.Big0,
			)

			newMember := gov.Gov.Call1(t, "getMember", memberLen).(common.Address)
			require.Equal(t, gov.client.Owner, newMember)
			newVoter := gov.Gov.Call1(t, "getVoter", memberLen).(common.Address)
			require.Equal(t, govMem1, newVoter)
			newReward := gov.Gov.Call1(t, "getReward", memberLen).(common.Address)
			require.Equal(t, user1, newReward)

			gov.Gov.ExecuteOk(t, govMem1, "vote", common.Big2, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big2)
			require.Equal(t, BallotStates.Rejected, state[1])
			require.True(t, state[2].(bool))

			memberLen = gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big1, memberLen)
			nodeLen := gov.Gov.Call1(t, "getNodeLength").(*big.Int)
			require.Equal(t, memberLen, nodeLen)
			lock := gov.Staking.Call1(t, "lockedBalanceOf", govMem1).(*big.Int)
			require.True(t, lock.Sign() == 0)
			bal := gov.Staking.Call1(t, "balanceOf", govMem1)
			require.Equal(t, LOCK_AMOUNT, bal)
		})
		t.Run("cannot addProposal to change non-member", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			node := gov.nodeInfos[0]
			govMem1 := sim.GetOrNewEOA("govMem1")
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToChangeMember",
					MemberInfo{
						Staker:     govMem1,
						Voter:      govMem1,
						Reward:     govMem1,
						Name:       node.name,
						Enode:      node.enode,
						Ip:         node.ip,
						Port:       node.port,
						LockAmount: LOCK_AMOUNT,
						Memo:       []byte("memo1"),
						Duration:   big.NewInt(86400),
					}, sim.GetOrNewEOA("govMem2"), common.Big0, common.Big0,
				),
				"Non-member",
			)
		})
		t.Run("can addProposal to change governance", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			newGovImp := compiled.Copy(new(Governance)).GovImp
			newGovImp.Deploy(t, gov.client)
			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeGov", newGovImp.Address, []byte("memo"), big.NewInt(86400))
			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big2, length)
		})
		t.Run("can not addProposal to change governance using EOA", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToChangeGov", sim.GetOrNewEOA("govMem1"), []byte("memo"), big.NewInt(86400)),
				"", "function call to a non-contract account", // TODO check
			)
		})
		t.Run("can not addProposal to change governance using non-UUPS", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToChangeGov", gov.Registry.Address, []byte("memo"), big.NewInt(86400)),
				"ERC1967Upgrade: new implementation is not UUPS",
			)
		})
		t.Run("cannot addProposal to change governance with same address", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToChangeGov", gov.GovImp.Address, []byte("memo"), big.NewInt(86400)),
				"Same contract address",
			)
		})
		t.Run("cannot addProposal to change governance with zero address", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToChangeGov", common.Address{}, []byte("memo"), big.NewInt(86400)),
				"Implementation cannot be zero",
			)
		})
		t.Run("can addProposal to change environment", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeEnv", sim.StringToBytes32("key"), EnvTypes.Bytes32, []byte("value"), []byte("memo"), big.NewInt(86400))
			length := gov.Gov.Call1(t, "ballotLength")
			require.Equal(t, common.Big2, length)
		})
		t.Run("cannot addProposal to change environment with wrong type", func(t *testing.T) {
			gov, voter := deployGovernance(t)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToChangeEnv", sim.StringToBytes32("key"), EnvTypes.Invalid, []byte("value"), []byte("memo"), big.NewInt(86400)),
				"Invalid type",
			)
		})
		t.Run("can vote approval to add member", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, voter, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				},
			)
			gov.Gov.ExecuteOk(t, voter, "vote", common.Big2, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big2)
			require.Equal(t, BallotStates.Accepted, state[1].(*big.Int))
			require.True(t, state[2].(bool))
			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big2, memberLen)
			nodeLen := gov.Gov.Call1(t, "getNodeLength").(*big.Int)
			require.Equal(t, memberLen, nodeLen)
			lock := gov.Staking.Call1(t, "lockedBalanceOf", govMem1)
			require.Equal(t, LOCK_AMOUNT, lock)

			newMember := gov.Gov.Call1(t, "getMember", memberLen).(common.Address)
			require.Equal(t, govMem1, newMember)
			newVoter := gov.Gov.Call1(t, "getVoter", memberLen).(common.Address)
			require.Equal(t, govMem1, newVoter)
			newReward := gov.Gov.Call1(t, "getReward", memberLen).(common.Address)
			require.Equal(t, govMem1, newReward)
		})
		t.Run("cannot vote approval to add member with insufficient staking", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, new(big.Int).Div(LOCK_AMOUNT, common.Big2), "deposit")
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, voter, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				},
			)
			gov.Gov.ExecuteOk(t, voter, "vote", common.Big2, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big2)
			require.Equal(t, BallotStates.Rejected, state[1].(*big.Int))
			require.True(t, state[2].(bool))
			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big1, memberLen)
		})
		t.Run("can vote disapproval to deny adding member", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, voter, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				},
			)
			gov.Gov.ExecuteOk(t, voter, "vote", common.Big2, false)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big2)
			require.Equal(t, BallotStates.Rejected, state[1].(*big.Int))
			require.True(t, state[2].(bool))
		})
		t.Run("can vote approval to change member totally", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")
			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, voter, "addProposalToAddMember",
				MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				},
			)
			gov.Gov.ExecuteOk(t, voter, "vote", common.Big2, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big2)
			require.Equal(t, BallotStates.Accepted, state[1].(*big.Int))
			require.True(t, state[2].(bool))

			govMem2 := sim.GetOrNewEOA("govMem2")
			_, receipt = gov.client.SendTransaction(t, gov.client.Owner, govMem2, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem2, LOCK_AMOUNT, "deposit")

			preDeployerAvail := gov.Staking.Call1(t, "availableBalanceOf", govMem1).(*big.Int)
			preGovmem1Avail := gov.Staking.Call1(t, "availableBalanceOf", govMem2).(*big.Int)

			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeMember",
				MemberInfo{
					Staker:     govMem2,
					Voter:      govMem2,
					Reward:     govMem2,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo2"),
					Duration:   big.NewInt(86400),
				}, govMem1, LOCK_AMOUNT, common.Big0,
			)
			gov.Gov.ExecuteOk(t, voter, "vote", big.NewInt(3), true)
			gov.Gov.ExecuteOk(t, govMem1, "vote", big.NewInt(3), true)
			length = gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, big.NewInt(3), length)
			inVoting = gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state = gov.BallotStorage.Call(t, "getBallotState", big.NewInt(3))
			require.Equal(t, BallotStates.Accepted, state[1].(*big.Int))
			require.True(t, state[2].(bool))

			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big2, memberLen)
			memberAddr := gov.Gov.Call1(t, "getMember", common.Big2).(common.Address)
			require.Equal(t, govMem2, memberAddr)
			getNode := gov.Gov.Call(t, "getNode", common.Big2)
			name, enode, ip, port := getNode[0].([]byte), getNode[1].([]byte), getNode[2].([]byte), getNode[3].(*big.Int)
			require.Equal(t, node1.name, name)
			require.Equal(t, node1.enode, enode)
			require.Equal(t, node1.ip, ip)
			require.Equal(t, node1.port, port)

			nodeIdxFromDeployer := gov.Gov.Call1(t, "getNodeIdxFromMember", govMem1).(*big.Int)
			require.True(t, nodeIdxFromDeployer.Sign() == 0)
			nodeIdxFromGovMem1 := gov.Gov.Call1(t, "getNodeIdxFromMember", govMem2).(*big.Int)
			require.Equal(t, common.Big2, nodeIdxFromGovMem1)

			postDeployerAvail := gov.Staking.Call1(t, "availableBalanceOf", govMem1).(*big.Int)
			postGovmem1Avail := gov.Staking.Call1(t, "availableBalanceOf", govMem2).(*big.Int)
			require.Equal(t, LOCK_AMOUNT, new(big.Int).Sub(postDeployerAvail, preDeployerAvail))
			require.Equal(t, LOCK_AMOUNT, new(big.Int).Sub(preGovmem1Avail, postGovmem1Avail))
		})
		t.Run("can vote approval to change enode only without voting", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}
			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      gov.client.Owner,
				Reward:     gov.client.Owner,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			gov.Gov.ExecuteOk(t, voter, "vote", common.Big2, true)

			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big2)
			require.Equal(t, BallotStates.Accepted, state[1])
			require.True(t, state[2].(bool))

			memberLen := gov.Gov.Call1(t, ("getMemberLength")).(*big.Int)
			require.Equal(t, common.Big1, memberLen)
			memberAddr := gov.Gov.Call1(t, "getMember", common.Big1)
			require.Equal(t, gov.client.Owner, memberAddr)
			getNode := gov.Gov.Call(t, "getNode", common.Big1)
			name, enode, ip, port := getNode[0].([]byte), getNode[1].([]byte), getNode[2].([]byte), getNode[3].(*big.Int)
			require.Equal(t, node1.name, name)
			require.Equal(t, node1.enode, enode)
			require.Equal(t, node1.ip, ip)
			require.Equal(t, node1.port, port)
		})
		t.Run("cannot vote approval to change member with insufficient staking", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, new(big.Int).Sub(LOCK_AMOUNT, big.NewInt(1000000000)), "deposit")

			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}
			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeMember", MemberInfo{
				Staker:     govMem1,
				Voter:      govMem1,
				Reward:     govMem1,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			gov.Gov.ExecuteOk(t, voter, "vote", common.Big2, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big2)
			require.Equal(t, BallotStates.Rejected, state[1])
			require.True(t, state[2].(bool))
		})
		t.Run("can vote approval to change governance", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			newGovImp := compiled.Copy(new(Governance)).GovImp
			newGovImp.Deploy(t, gov.client)

			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeGov", newGovImp.Address, []byte("memo"), big.NewInt(86400))
			gov.Gov.ExecuteOk(t, voter, "vote", common.Big2, true)

			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big2)
			require.Equal(t, BallotStates.Accepted, state[1])
			require.True(t, state[2].(bool))

			imp := gov.Gov.Call1(t, "implementation").(common.Address)
			require.Equal(t, newGovImp.Address, imp)
		})
		t.Run("can vote approval to change environment", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			blocksPer := gov.EnvStorage.Call1(t, "getBlocksPer").(*big.Int)
			require.NotEqual(t, big.NewInt(100), blocksPer)

			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				big.NewInt(86400),
			)
			gov.Gov.ExecuteOk(t, voter, "vote", common.Big2, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big2)
			require.Equal(t, BallotStates.Accepted, state[1])
			require.True(t, state[2].(bool))

			blocksPer = gov.EnvStorage.Call1(t, "getBlocksPer").(*big.Int)
			require.Equal(t, big.NewInt(100), blocksPer)
		})
		t.Run("cannot vote for a ballot already done", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, new(big.Int).Sub(LOCK_AMOUNT, big.NewInt(1000000000)), "deposit")

			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}
			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeMember", MemberInfo{
				Staker:     govMem1,
				Voter:      govMem1,
				Reward:     govMem1,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			gov.Gov.ExecuteOk(t, voter, "vote", common.Big2, true)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "vote", common.Big2, true),
				"Expired",
			)
		})
		t.Run("cannot vote for a ballot already staker done", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, new(big.Int).Sub(LOCK_AMOUNT, big.NewInt(1000000000)), "deposit")

			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}
			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeMember", MemberInfo{
				Staker:     govMem1,
				Voter:      govMem1,
				Reward:     govMem1,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big2, true)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "vote", common.Big2, true),
				"Expired",
			)
		})
		t.Run("cannot vote for a ballot already voter done", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Mul(LOCK_AMOUNT, common.Big2), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, new(big.Int).Sub(LOCK_AMOUNT, big.NewInt(1000000000)), "deposit")

			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}
			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeMember", MemberInfo{
				Staker:     govMem1,
				Voter:      govMem1,
				Reward:     govMem1,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo1"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			gov.Gov.ExecuteOk(t, voter, "vote", common.Big2, true)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "vote", common.Big2, true),
				"Expired",
			)
		})
		t.Run("cannot add proposal durring period time", func(t *testing.T) {
			gov, voter := deployGovernance(t)

			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				big.NewInt(86400),
			)
			gov.Gov.ExecuteOk(t, voter, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				big.NewInt(86400),
			)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "setProposalTimePeriod", big.NewInt(60))
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter, "addProposalToChangeEnv",
					crypto.Keccak256Hash([]byte("blocksPer")),
					EnvTypes.Uint,
					hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
					[]byte("memo"),
					big.NewInt(86400),
				), "Cannot add proposal too early",
			)
		})
	})
	t.Run("Two Member", func(t *testing.T) {
		deployGovernance := func(t *testing.T) (gov *Governance, govMem1 common.Address) {
			gov = DeployGovernance(t)
			govMem1 = sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Add(LOCK_AMOUNT, towei(1)), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")

			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem1,
				Voter:      govMem1,
				Reward:     govMem1,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			})
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big1, true)
			require.Equal(t, common.Big2, gov.Gov.Call1(t, "getMemberLength").(*big.Int))

			gov.nodeInfos = append(gov.nodeInfos, node1)
			return
		}
		t.Run("cannot vote with changed voter address", func(t *testing.T) {
			gov, _ := deployGovernance(t)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeEnv", sim.StringToBytes32("key"), EnvTypes.Bytes32, []byte("value"), []byte("memo"), big.NewInt(86400))
			ballotIdx := gov.Gov.Call1(t, "ballotLength").(*big.Int)

			node := gov.nodeInfos[0]
			voter, voter1, user1 := sim.GetOrNewEOA("voter"), sim.GetOrNewEOA("voter1"), sim.GetOrNewEOA("user1")
			gov.client.SendTransaction(t, gov.client.Owner, voter, towei(1), []byte{})
			gov.client.SendTransaction(t, gov.client.Owner, voter1, towei(1), []byte{})

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      voter,
				Reward:     user1,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			gov.Gov.ExecuteOk(t, voter, "vote", ballotIdx, true)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
				Staker:     gov.client.Owner,
				Voter:      voter1,
				Reward:     user1,
				Name:       node.name,
				Enode:      node.enode,
				Ip:         node.ip,
				Port:       node.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			}, gov.client.Owner, common.Big0, common.Big0)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, voter1, "vote", ballotIdx, true),
				"already voted",
			)
		})
		t.Run("cannot addProposal to add member self", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)
			node := gov.nodeInfos[0]
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, govMem1, "addProposalToAddMember",
					MemberInfo{
						Staker:     govMem1,
						Voter:      govMem1,
						Reward:     govMem1,
						Name:       node.name,
						Enode:      node.enode,
						Ip:         node.ip,
						Port:       node.port,
						LockAmount: LOCK_AMOUNT,
						Memo:       []byte("memo1"),
						Duration:   big.NewInt(86400),
					},
				), "Already member",
			)
		})
		t.Run("can addProposal to remove member", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToRemoveMember", govMem1, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int))
			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big2, length)
		})
		t.Run("can addProposal to add member where info is the removed member's with same govMem", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToRemoveMember", govMem1, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int))
			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big2, length)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", length, true)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.Equal(t, length, inVoting)

			state := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.InProgress, state[1])
			require.False(t, state[2].(bool))

			gov.Gov.ExecuteOk(t, govMem1, "vote", length, true)
			inVoting = gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0, inVoting)

			state = gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, state[1])
			require.True(t, state[2].(bool))
		})
		t.Run("can addProposal to add member where info is the removed member's", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToRemoveMember", govMem1, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int))
			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			require.Equal(t, common.Big2, length)

			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", length, true)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.Equal(t, length, inVoting)

			state := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.InProgress, state[1])
			require.False(t, state[2].(bool))

			gov.Gov.ExecuteOk(t, govMem1, "vote", length, true)
			inVoting = gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0, inVoting)

			state = gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, state[1])
			require.True(t, state[2].(bool))

			t.Log("member removed")
			govMem2 := sim.GetOrNewEOA("govMem2")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem2, new(big.Int).Add(LOCK_AMOUNT, towei(1)), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem2, LOCK_AMOUNT, "deposit")

			node1 := gov.nodeInfos[1]

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem2,
				Voter:      govMem2,
				Reward:     govMem2,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			})

			length = gov.Gov.Call1(t, "ballotLength").(*big.Int)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", length, true)
			inVoting = gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state = gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, state[1])
			require.True(t, state[2].(bool))
		})
		t.Run("can vote to add member", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)

			govMem2 := sim.GetOrNewEOA("govMem2")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem2, new(big.Int).Add(LOCK_AMOUNT, towei(1)), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem2, LOCK_AMOUNT, "deposit")

			node2 := nodeInfo{
				name:  []byte("name2"),
				enode: hexutil.MustDecode("0x888777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a1"),
				ip:    []byte("127.0.0.3"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem2,
				Voter:      govMem2,
				Reward:     govMem2,
				Name:       node2.name,
				Enode:      node2.enode,
				Ip:         node2.ip,
				Port:       node2.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			})

			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", length, true)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.Equal(t, length, inVoting)
			state := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.InProgress, state[1])
			require.False(t, state[2].(bool))

			gov.Gov.ExecuteOk(t, govMem1, "vote", length, true)
			inVoting2 := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting2.Sign() == 0)
			state2 := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, state2[1])
			require.True(t, state2[2].(bool))
		})
		t.Run("can vote to deny adding member", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)

			govMem2 := sim.GetOrNewEOA("govMem2")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem2, new(big.Int).Add(LOCK_AMOUNT, towei(1)), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem2, LOCK_AMOUNT, "deposit")

			node2 := nodeInfo{
				name:  []byte("name2"),
				enode: hexutil.MustDecode("0x888777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a1"),
				ip:    []byte("127.0.0.3"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem2,
				Voter:      govMem2,
				Reward:     govMem2,
				Name:       node2.name,
				Enode:      node2.enode,
				Ip:         node2.ip,
				Port:       node2.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			})

			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", length, false)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.Equal(t, length, inVoting)
			state := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.InProgress, state[1])
			require.False(t, state[2].(bool))

			gov.Gov.ExecuteOk(t, govMem1, "vote", length, false)
			inVoting2 := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting2.Sign() == 0)
			state2 := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Rejected, state2[1])
			require.True(t, state2[2].(bool))
		})
		t.Run("can vote to remove first member", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)
			preAvail := gov.Staking.Call1(t, "availableBalanceOf", gov.client.Owner).(*big.Int)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToRemoveMember", gov.client.Owner, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int))

			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", length, true)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.Equal(t, length, inVoting)
			state := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.InProgress, state[1])
			require.False(t, state[2].(bool))

			gov.Gov.ExecuteOk(t, govMem1, "vote", length, true)
			inVoting2 := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting2.Sign() == 0)
			state2 := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, state2[1])
			require.True(t, state2[2].(bool))

			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big1, memberLen)
			isMem := gov.Gov.Call1(t, "isMember", gov.client.Owner).(bool)
			require.False(t, isMem)
			nodeLen := gov.Gov.Call1(t, "getNodeLength").(*big.Int)
			require.Equal(t, common.Big1, nodeLen)
			nodeIdx := gov.Gov.Call1(t, "getNodeIdxFromMember", gov.client.Owner).(*big.Int)
			require.True(t, nodeIdx.Sign() == 0)
			nodeIdx2 := gov.Gov.Call1(t, "getNodeIdxFromMember", govMem1).(*big.Int)
			require.Equal(t, common.Big1, nodeIdx2)

			postAvail := gov.Staking.Call1(t, "availableBalanceOf", gov.client.Owner).(*big.Int)
			require.Equal(t, LOCK_AMOUNT, new(big.Int).Sub(postAvail, preAvail))
		})
		t.Run("can vote to remove last member", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)
			preAvail := gov.Staking.Call1(t, "availableBalanceOf", govMem1).(*big.Int)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToRemoveMember", govMem1, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int))

			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", length, true)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.Equal(t, length, inVoting)
			state := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.InProgress, state[1])
			require.False(t, state[2].(bool))

			gov.Gov.ExecuteOk(t, govMem1, "vote", length, true)
			inVoting2 := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting2.Sign() == 0)
			state2 := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Accepted, state2[1])
			require.True(t, state2[2].(bool))

			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big1, memberLen)
			isMem := gov.Gov.Call1(t, "isMember", govMem1).(bool)
			require.False(t, isMem)
			nodeLen := gov.Gov.Call1(t, "getNodeLength").(*big.Int)
			require.Equal(t, common.Big1, nodeLen)
			nodeIdx := gov.Gov.Call1(t, "getNodeIdxFromMember", govMem1).(*big.Int)
			require.True(t, nodeIdx.Sign() == 0)
			nodeIdx2 := gov.Gov.Call1(t, "getNodeIdxFromMember", gov.client.Owner).(*big.Int)
			require.Equal(t, common.Big1, nodeIdx2)

			postAvail := gov.Staking.Call1(t, "availableBalanceOf", govMem1).(*big.Int)
			require.Equal(t, LOCK_AMOUNT, new(big.Int).Sub(postAvail, preAvail))
		})
		t.Run("cannot vote simultaneously", func(t *testing.T) {
			gov, _ := deployGovernance(t)

			govMem2 := sim.GetOrNewEOA("govMem2")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem2, new(big.Int).Add(LOCK_AMOUNT, towei(1)), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem2, LOCK_AMOUNT, "deposit")

			node2 := nodeInfo{
				name:  []byte("name2"),
				enode: hexutil.MustDecode("0x888777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a1"),
				ip:    []byte("127.0.0.3"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem2,
				Voter:      govMem2,
				Reward:     govMem2,
				Name:       node2.name,
				Enode:      node2.enode,
				Ip:         node2.ip,
				Port:       node2.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			})

			govMem3 := sim.GetOrNewEOA("govMem3")
			_, receipt = gov.client.SendTransaction(t, gov.client.Owner, govMem3, new(big.Int).Add(LOCK_AMOUNT, towei(1)), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem3, LOCK_AMOUNT, "deposit")

			node3 := nodeInfo{
				name:  []byte("name3"),
				enode: hexutil.MustDecode("0x999777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a1"),
				ip:    []byte("127.0.0.4"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem3,
				Voter:      govMem3,
				Reward:     govMem3,
				Name:       node3.name,
				Enode:      node3.enode,
				Ip:         node3.ip,
				Port:       node3.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			})

			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", new(big.Int).Sub(length, common.Big1), true)
			voting := gov.Gov.Call1(t, "getBallotInVoting")
			require.Equal(t, new(big.Int).Sub(length, common.Big1), voting)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "vote", length, true),
				"Now in voting with different ballot",
			)
		})
		t.Run("vote is ended when the sum of voting power is max", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)

			govMem2 := sim.GetOrNewEOA("govMem2")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem2, new(big.Int).Add(LOCK_AMOUNT, towei(1)), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem2, LOCK_AMOUNT, "deposit")

			node2 := nodeInfo{
				name:  []byte("name2"),
				enode: hexutil.MustDecode("0x888777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a1"),
				ip:    []byte("127.0.0.3"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem2,
				Voter:      govMem2,
				Reward:     govMem2,
				Name:       node2.name,
				Enode:      node2.enode,
				Ip:         node2.ip,
				Port:       node2.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			})

			length := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", length, true)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.Equal(t, length, inVoting)
			state := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.InProgress, state[1])
			require.False(t, state[2].(bool))

			gov.Gov.ExecuteOk(t, govMem1, "vote", length, false)
			inVoting2 := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting2.Sign() == 0)
			state2 := gov.BallotStorage.Call(t, "getBallotState", length)
			require.Equal(t, BallotStates.Rejected, state2[1])
			require.True(t, state2[2].(bool))
		})
		t.Run("cannot vote approval when the voting is ended", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)

			delay_time := big.NewInt(86400)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				delay_time,
			)
			ballotLen := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", ballotLen, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big2, length)

			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.Equal(t, ballotLen, inVoting)
			state := gov.BallotStorage.Call(t, "getBallotState", ballotLen)
			require.Equal(t, BallotStates.InProgress, state[1])
			require.False(t, state[2].(bool))

			gov.client.AdjustTime(t, time.Second*time.Duration(new(big.Int).Mul(delay_time, big.NewInt(2000)).Int64()))
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, govMem1, "vote", ballotLen, false),
				"Expired",
			)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "finalizeEndedVote")

			inVoting = gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
		})
		t.Run("Non member cannot end voting", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)

			delay_time := big.NewInt(86400)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToChangeEnv",
				crypto.Keccak256Hash([]byte("blocksPer")),
				EnvTypes.Uint,
				hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064"), // 32bytes, 100
				[]byte("memo"),
				delay_time,
			)
			ballotLen := gov.Gov.Call1(t, "ballotLength").(*big.Int)
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", ballotLen, true)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big2, length)

			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.Equal(t, ballotLen, inVoting)
			state := gov.BallotStorage.Call(t, "getBallotState", ballotLen)
			require.Equal(t, BallotStates.InProgress, state[1])
			require.False(t, state[2].(bool))

			gov.client.AdjustTime(t, time.Second*time.Duration(new(big.Int).Mul(delay_time, big.NewInt(2000)).Int64()))
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, govMem1, "vote", ballotLen, false),
				"Expired",
			)
			govMem3 := sim.GetOrNewEOA("getMem3")
			gov.client.SendTransaction(t, gov.client.Owner, govMem3, towei(1), []byte{})
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, govMem3, "finalizeEndedVote"),
				"No Permission",
			)

			inVoting = gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.Equal(t, ballotLen, inVoting)
		})
		t.Run("reject proposal without voting about changing voter address if voter is already registered", func(t *testing.T) {
			gov, govMem1 := deployGovernance(t)
			node := gov.nodeInfos[0]
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, gov.client.Owner, "addProposalToChangeMember", MemberInfo{
					Staker:     gov.client.Owner,
					Voter:      govMem1,
					Reward:     gov.client.Owner,
					Name:       node.name,
					Enode:      node.enode,
					Ip:         node.ip,
					Port:       node.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo1"),
					Duration:   big.NewInt(86400),
				}, gov.client.Owner, common.Big0, common.Big0),
				"Already a member",
			)
			length := gov.Gov.Call1(t, "voteLength").(*big.Int)
			require.Equal(t, common.Big1, length)
			inVoting := gov.Gov.Call1(t, "getBallotInVoting").(*big.Int)
			require.True(t, inVoting.Sign() == 0)
			state := gov.BallotStorage.Call(t, "getBallotState", common.Big2)
			require.True(t, state[1].(*big.Int).Sign() == 0)
			require.False(t, state[2].(bool))

			memberLen := gov.Gov.Call1(t, "getMemberLength").(*big.Int)
			require.Equal(t, common.Big2, memberLen)
			memberAddr := gov.Gov.Call1(t, "getMember", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, memberAddr)
			voterAddr := gov.Gov.Call1(t, "getVoter", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, voterAddr)
			rewardAddr := gov.Gov.Call1(t, "getReward", common.Big1).(common.Address)
			require.Equal(t, gov.client.Owner, rewardAddr)

			getNode := gov.Gov.Call(t, "getNode", common.Big1)
			name, enode, ip, port := getNode[0].([]byte), getNode[1].([]byte), getNode[2].([]byte), getNode[3].(*big.Int)
			require.Equal(t, node.name, name)
			require.Equal(t, node.enode, enode)
			require.Equal(t, node.ip, ip)
			require.Equal(t, node.port, port)
		})
	})
	t.Run("Others", func(t *testing.T) {
		t.Run("cannot init", func(t *testing.T) {
			gov := DeployGovernance(t)
			govMem1 := sim.GetOrNewEOA("govMem1")
			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem1, new(big.Int).Add(LOCK_AMOUNT, towei(1)), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem1, LOCK_AMOUNT, "deposit")

			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem1,
				Voter:      govMem1,
				Reward:     govMem1,
				Name:       node1.name,
				Enode:      node1.enode,
				Ip:         node1.ip,
				Port:       node1.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			})
			gov.Gov.ExecuteOk(t, gov.client.Owner, "vote", common.Big1, true)
			require.Equal(t, common.Big2, gov.Gov.Call1(t, "getMemberLength").(*big.Int))

			node := gov.nodeInfos[0]
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, govMem1, "init", gov.Registry.Address, LOCK_AMOUNT, node.name, node.enode, node.ip, node.port),
				"Initializable: contract is already initialized",
			)
		})
		t.Run("cannot addProposal", func(t *testing.T) {
			gov := DeployGovernance(t)
			govMem1, govMem2 := sim.GetOrNewEOA("govMem1"), sim.GetOrNewEOA("govMem2")
			gov.client.SendTransaction(t, gov.client.Owner, govMem1, towei(1), []byte{})
			gov.client.SendTransaction(t, gov.client.Owner, govMem2, towei(1), []byte{})

			node := gov.nodeInfos[0]
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, govMem1, "addProposalToAddMember", MemberInfo{
					Staker:     govMem1,
					Voter:      govMem1,
					Reward:     govMem1,
					Name:       node.name,
					Enode:      node.enode,
					Ip:         node.ip,
					Port:       node.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo"),
					Duration:   big.NewInt(86400),
				}),
				"No Permission",
			)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, govMem1, "addProposalToRemoveMember", govMem1, LOCK_AMOUNT, []byte("memo1"), big.NewInt(86400), LOCK_AMOUNT, new(big.Int)),
				"No Permission",
			)

			_, receipt := gov.client.SendTransaction(t, gov.client.Owner, govMem2, new(big.Int).Add(LOCK_AMOUNT, towei(1)), []byte{})
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			gov.Staking.ExecuteWithETHOk(t, govMem2, LOCK_AMOUNT, "deposit")

			node1 := nodeInfo{
				name:  []byte("name1"),
				enode: hexutil.MustDecode("0x777777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
				ip:    []byte("127.0.0.2"),
				port:  big.NewInt(8542),
			}
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, govMem1, "addProposalToChangeMember", MemberInfo{
					Staker:     govMem2,
					Voter:      govMem2,
					Reward:     govMem2,
					Name:       node1.name,
					Enode:      node1.enode,
					Ip:         node1.ip,
					Port:       node1.port,
					LockAmount: LOCK_AMOUNT,
					Memo:       []byte("memo"),
					Duration:   big.NewInt(86400),
				}, govMem1, common.Big0, common.Big0),
				"No Permission",
			)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, govMem1, "addProposalToChangeGov", govMem1, []byte("memo"), big.NewInt(86400)),
				"No Permission", // TODO check
			)
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, govMem1, "addProposalToChangeEnv", sim.StringToBytes32("key"), EnvTypes.Bytes32, []byte("value"), []byte("memo"), big.NewInt(86400)),
				"No Permission",
			)
		})
		t.Run("cannot vote", func(t *testing.T) {
			gov := DeployGovernance(t)
			govMem1, govMem2 := sim.GetOrNewEOA("govMem1"), sim.GetOrNewEOA("govMem2")
			gov.client.SendTransaction(t, gov.client.Owner, govMem1, towei(1), []byte{})
			gov.client.SendTransaction(t, gov.client.Owner, govMem2, new(big.Int).Add(LOCK_AMOUNT, towei(1)), []byte{})
			gov.Staking.ExecuteWithETHOk(t, govMem2, LOCK_AMOUNT, "deposit")

			node2 := nodeInfo{
				name:  []byte("name2"),
				enode: hexutil.MustDecode("0x888777777711c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a1"),
				ip:    []byte("127.0.0.3"),
				port:  big.NewInt(8542),
			}

			gov.Gov.ExecuteOk(t, gov.client.Owner, "addProposalToAddMember", MemberInfo{
				Staker:     govMem2,
				Voter:      govMem2,
				Reward:     govMem2,
				Name:       node2.name,
				Enode:      node2.enode,
				Ip:         node2.ip,
				Port:       node2.port,
				LockAmount: LOCK_AMOUNT,
				Memo:       []byte("memo"),
				Duration:   big.NewInt(86400),
			})
			sim.ExpectedRevert(t,
				gov.Gov.ExecuteFail(t, govMem1, "vote", common.Big1, true),
				"No Permission",
			)
		})
	})
}
