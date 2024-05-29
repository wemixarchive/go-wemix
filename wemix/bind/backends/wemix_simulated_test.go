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
	require.NoError(t, err)
	require.NotNil(t, b)

	t.Log(wemix_miner.HasMiningTokenFunc())
	t.Log(wemix_miner.SuggestGasPriceFunc())
	t.Log(wemix_miner.GetCoinbaseFunc(common.Big0))

	callOpts := new(bind.CallOpts)
	contracts, err := gov.GetGovContractsByOwner(callOpts, b, opts.From)
	require.NoError(t, err)
	t.Log(contracts.Address().Registry)
	t.Log(contracts.EnvStorageImp.GetBallotDurationMax(callOpts))
	t.Log(contracts.EnvStorageImp.GetBallotDurationMin(callOpts))
	t.Log(contracts.EnvStorageImp.GetBallotDurationMinMax(callOpts))
	t.Log(contracts.EnvStorageImp.GetBlockCreationTime(callOpts))
	t.Log(contracts.EnvStorageImp.GetBlockRewardAmount(callOpts))
	t.Log(contracts.EnvStorageImp.GetBlockRewardDistributionMethod(callOpts))
	t.Log(contracts.EnvStorageImp.GetBlocksPer(callOpts))
	t.Log(contracts.EnvStorageImp.GetGasLimitAndBaseFee(callOpts))
	t.Log(contracts.EnvStorageImp.GetMaxBaseFee(callOpts))
	t.Log(contracts.EnvStorageImp.GetMaxIdleBlockInterval(callOpts))
	t.Log(contracts.EnvStorageImp.GetMaxPriorityFeePerGas(callOpts))
	t.Log(contracts.EnvStorageImp.GetStakingMax(callOpts))
	t.Log(contracts.EnvStorageImp.GetStakingMin(callOpts))
	t.Log(contracts.EnvStorageImp.GetStakingMinMax(callOpts))
	t.Log(contracts.EnvStorageImp.GASTARGETPERCENTAGENAME(callOpts))

	getNode, err := contracts.GovImp.GetNode(callOpts, common.Big1)
	require.NoError(t, err)
	t.Log(
		"\n Name:", string(getNode.Name),
		"\n Enode:", hex.EncodeToString(getNode.Enode),
		"\n Ip:", string(getNode.Ip),
		"\n Port:", getNode.Port,
	)
	// /*
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

	ticker := time.NewTicker(10e9)
loop:
	for {
		b.Commit()
		select {
		case head := <-newHead:
			t.Log("head.Number", head.Number)
			time.Sleep(0.2e9)
		case err := <-sub.Err():
			require.NoError(t, err)
		case <-ticker.C:
			break loop
		}
	}
	t.Log("end")
}
