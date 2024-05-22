package wemix_backends_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	wemix_backends "github.com/ethereum/go-ethereum/accounts/abi/bind/backends/wemix"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	gov "github.com/ethereum/go-ethereum/wemix/governance-contract/bind"
	wemix_miner "github.com/ethereum/go-ethereum/wemix/miner"
	"github.com/stretchr/testify/require"
)

func TestWemixBackends(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	opts, err := bind.NewKeyedTransactorWithChainID(key, wemix_backends.ChainID)
	require.NoError(t, err)
	b, err := wemix_backends.NewWemixSimulatedBackend(
		key,
		t.TempDir(),
		core.GenesisAlloc{
			opts.From: {Balance: new(big.Int).Mul(big.NewInt(256_000_000_000_000), big.NewInt(params.Ether))},
		},
	)
	require.NoError(t, err)
	require.NotNil(t, b)

	t.Log(wemix_miner.HasMiningTokenFunc())
	t.Log(wemix_miner.SuggestGasPriceFunc())
	t.Log(wemix_miner.GetCoinbaseFunc(common.Big0))

	// t.Log(admin.Info())
	callOpts := new(bind.CallOpts)
	contracts, err := gov.GetGovContractsByOwner(callOpts, b, opts.From)
	require.NoError(t, err)
	t.Log(contracts.Registry.Address())
	t.Log(contracts.EnvStorageImp.Funcs.GetBallotDurationMax(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GetBallotDurationMin(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GetBallotDurationMinMax(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GetBlockCreationTime(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GetBlockRewardAmount(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GetBlockRewardDistributionMethod(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GetBlocksPer(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GetGasLimitAndBaseFee(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GetMaxBaseFee(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GetMaxIdleBlockInterval(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GetMaxPriorityFeePerGas(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GetStakingMax(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GetStakingMin(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GetStakingMinMax(callOpts))
	t.Log(contracts.EnvStorageImp.Funcs.GASTARGETPERCENTAGENAME(callOpts))
}
