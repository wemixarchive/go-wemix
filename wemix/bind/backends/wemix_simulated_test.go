package backends_test

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	gov "github.com/ethereum/go-ethereum/wemix/bind"
	"github.com/ethereum/go-ethereum/wemix/bind/backends"
	wemix_miner "github.com/ethereum/go-ethereum/wemix/miner"
	"github.com/stretchr/testify/require"
)

func TestWemixBackends(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	opts, err := bind.NewKeyedTransactorWithChainID(key, backends.ChainID)
	require.NoError(t, err)
	b, err := backends.NewWemixSimulatedBackend(
		key,
		t.TempDir(),
		core.GenesisAlloc{
			opts.From: {Balance: new(big.Int).Mul(big.NewInt(256_000_000_000_000), big.NewInt(params.Ether))},
		},
		backends.SetLogLevel(2),
	)
	if err != nil {
		return
	}
	require.NotNil(t, b)

	defaultEnv := gov.DefaultInitEnvStorage

	require.True(t, defaultEnv.MAX_PRIORITY_FEE_PER_GAS.Cmp(wemix_miner.SuggestGasPrice()) == 0)
	coinBase, err := wemix_miner.GetCoinbase(common.Big0)
	require.NoError(t, err)
	require.Equal(t, opts.From, coinBase)

	callOpts := new(bind.CallOpts)
	contracts, err := gov.GetGovContractsByOwner(callOpts, b, opts.From)
	require.NoError(t, err)

	got, err := contracts.EnvStorageImp.GetBallotDurationMax(callOpts)
	require.NoError(t, err)
	require.Truef(t, defaultEnv.BALLOT_DURATION_MAX.Cmp(got) == 0, "mismatch BALLOT_DURATION_MAX, got: %v, want: %v", got, defaultEnv.BALLOT_DURATION_MAX)

	got, err = contracts.EnvStorageImp.GetBallotDurationMin(callOpts)
	require.NoError(t, err)
	require.Truef(t, defaultEnv.BALLOT_DURATION_MIN.Cmp(got) == 0, "mismatch BALLOT_DURATION_MIN, got: %v, want: %v", got, defaultEnv.BALLOT_DURATION_MIN)

	got, err = contracts.EnvStorageImp.GetBlockCreationTime(callOpts)
	require.NoError(t, err)
	require.Truef(t, defaultEnv.BLOCK_CREATION_TIME.Cmp(got) == 0, "mismatch BLOCK_CREATION_TIME, got: %v, want: %v", got, defaultEnv.BLOCK_CREATION_TIME)

	got, err = contracts.EnvStorageImp.GetBlockRewardAmount(callOpts)
	require.NoError(t, err)
	require.Truef(t, defaultEnv.BLOCK_REWARD_AMOUNT.Cmp(got) == 0, "mismatch BLOCK_REWARD_AMOUNT, got: %v, want: %v", got, defaultEnv.BLOCK_REWARD_AMOUNT)

	blockRewardDistributionBlockProducer,
		blockRewardDistributionStakingReward, blockRewardDistributionEcosystem,
		blockRewardDistributionMaintenance,
		err := contracts.EnvStorageImp.GetBlockRewardDistributionMethod(callOpts)
	require.NoError(t, err)
	require.Truef(t, defaultEnv.BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER.Cmp(blockRewardDistributionBlockProducer) == 0, "mismatch BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER, got: %v, want: %v", blockRewardDistributionBlockProducer, defaultEnv.BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER)
	require.Truef(t, defaultEnv.BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD.Cmp(blockRewardDistributionStakingReward) == 0, "mismatch BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD, got: %v, want: %v", blockRewardDistributionStakingReward, defaultEnv.BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD)
	require.Truef(t, defaultEnv.BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM.Cmp(blockRewardDistributionEcosystem) == 0, "mismatch BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM, got: %v, want: %v", blockRewardDistributionEcosystem, defaultEnv.BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM)
	require.Truef(t, defaultEnv.BLOCK_REWARD_DISTRIBUTION_MAINTENANCE.Cmp(blockRewardDistributionMaintenance) == 0, "mismatch BLOCK_REWARD_DISTRIBUTION_MAINTENANCE, got: %v, want: %v", blockRewardDistributionMaintenance, defaultEnv.BLOCK_REWARD_DISTRIBUTION_MAINTENANCE)

	got, err = contracts.EnvStorageImp.GetBlocksPer(callOpts)
	require.NoError(t, err)
	require.Truef(t, defaultEnv.BLOCKS_PER.Cmp(got) == 0, "mismatch BLOCKS_PER, got: %v, want: %v", got, defaultEnv.BLOCKS_PER)

	got, err = contracts.EnvStorageImp.GetMaxBaseFee(callOpts)
	require.NoError(t, err)
	require.Truef(t, defaultEnv.MAX_BASE_FEE.Cmp(got) == 0, "mismatch MAX_BASE_FEE, got: %v, want: %v", got, defaultEnv.MAX_BASE_FEE)

	got, err = contracts.EnvStorageImp.GetMaxIdleBlockInterval(callOpts)
	require.NoError(t, err)
	require.Truef(t, defaultEnv.MAX_IDLE_BLOCK_INTERVAL.Cmp(got) == 0, "mismatch MAX_IDLE_BLOCK_INTERVAL, got: %v, want: %v", got, defaultEnv.MAX_IDLE_BLOCK_INTERVAL)

	got, err = contracts.EnvStorageImp.GetMaxPriorityFeePerGas(callOpts)
	require.NoError(t, err)
	require.Truef(t, defaultEnv.MAX_PRIORITY_FEE_PER_GAS.Cmp(got) == 0, "mismatch MAX_PRIORITY_FEE_PER_GAS, got: %v, want: %v", got, defaultEnv.MAX_PRIORITY_FEE_PER_GAS)

	got, err = contracts.EnvStorageImp.GetStakingMax(callOpts)
	require.NoError(t, err)
	require.Truef(t, defaultEnv.STAKING_MAX.Cmp(got) == 0, "mismatch STAKING_MAX, got: %v, want: %v", got, defaultEnv.STAKING_MAX)

	got, err = contracts.EnvStorageImp.GetStakingMin(callOpts)
	require.NoError(t, err)
	require.Truef(t, defaultEnv.STAKING_MIN.Cmp(got) == 0, "mismatch STAKING_MIN, got: %v, want: %v", got, defaultEnv.STAKING_MIN)

	getNode, err := contracts.GovImp.GetNode(callOpts, common.Big1)
	require.NoError(t, err)
	t.Log(
		"\n Name:", string(getNode.Name),
		"\n Enode:", hex.EncodeToString(getNode.Enode),
		"\n Ip:", string(getNode.Ip),
		"\n Port:", getNode.Port,
	)
	newVoterKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	newVoter := crypto.PubkeyToAddress(newVoterKey.PublicKey)

	LOCK_AMOUNT, err := contracts.EnvStorageImp.GetStakingMax(callOpts)
	require.NoError(t, err)

	tx, err := contracts.GovImp.AddProposalToChangeMember(opts, gov.GovImpMemberInfo{
		Staker:     opts.From,
		Voter:      newVoter,
		Reward:     opts.From,
		Name:       getNode.Name,
		Enode:      getNode.Enode,
		Ip:         getNode.Ip,
		Port:       getNode.Port,
		LockAmount: LOCK_AMOUNT,
		Memo:       []byte("memo1"),
		Duration:   big.NewInt(86400),
	}, opts.From, common.Big0, common.Big0)
	require.NoError(t, err)

	receipt, err := bind.WaitMined(context.Background(), b, tx)
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	inVoting, err := contracts.GovImp.GetBallotInVoting(callOpts)
	require.NoError(t, err)
	require.True(t, inVoting.Sign() == 0)

	member, err := contracts.GovImp.GetMember(callOpts, common.Big1)
	require.NoError(t, err)
	require.Equal(t, opts.From, member)

	voter, err := contracts.GovImp.GetVoter(callOpts, common.Big1)
	require.NoError(t, err)
	require.Equal(t, newVoter, voter)

	newReward, err := contracts.GovImp.GetReward(callOpts, common.Big1)
	require.NoError(t, err)
	require.Equal(t, opts.From, newReward)

	newHead := make(chan *types.Header)
	sub, err := b.SubscribeNewHead(context.Background(), newHead)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	ticker := time.NewTicker(10e9)
	var headNumber *big.Int = nil
loop:
	for {
		b.Commit()
		select {
		case head := <-newHead:
			if headNumber == nil {
				headNumber = head.Number
			} else {
				require.Equal(t, new(big.Int).Add(headNumber, common.Big1), head.Number)
				headNumber = head.Number
			}
		case err := <-sub.Err():
			require.NoError(t, err)
		case <-ticker.C:
			break loop
		}
	}
	t.Log("end")
}
