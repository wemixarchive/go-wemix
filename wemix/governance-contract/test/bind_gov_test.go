package test

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	gov "github.com/ethereum/go-ethereum/wemix/governance-contract/bind"
	"github.com/ethereum/go-ethereum/wemix/governance-contract/common/bn"
	sim "github.com/ethereum/go-ethereum/wemix/governance-contract/common/simulated-backend"
	"github.com/stretchr/testify/require"
)

func TestBindFuncs(t *testing.T) {
	client := sim.NewClient(t)
	callOpts := new(bind.CallOpts)
	opts, err := bind.NewKeyedTransactorWithChainID(sim.EOAKey[sim.GetOrNewEOA("owner")], client.ChainID)
	require.NoError(t, err)
	go func() {
		for {
			time.Sleep(1e9)
			client.Commit()
		}
	}()
	var deployed = new(gov.GovContracts)
	t.Run("deploy", func(t *testing.T) {
		deployed, err = gov.DeployGovContracts(opts, client.Backend)
		require.NoError(t, err)
		nilAddress := common.Address{}
		require.NotEqual(t, nilAddress, deployed.Registry.Address())
		require.NotEqual(t, nilAddress, deployed.Gov.Address())
		require.NotEqual(t, nilAddress, deployed.GovImp.Address())
		require.NotEqual(t, nilAddress, deployed.Staking.Address())
		require.NotEqual(t, nilAddress, deployed.StakingImp.Address())
		require.NotEqual(t, nilAddress, deployed.BallotStorage.Address())
		require.NotEqual(t, nilAddress, deployed.BallotStorageImp.Address())
		require.NotEqual(t, nilAddress, deployed.EnvStorage.Address())
		require.NotEqual(t, nilAddress, deployed.EnvStorageImp.Address())
	})
	var copyed = new(gov.GovContracts)
	t.Run("copy", func(t *testing.T) {
		require.NoError(t, deployed.Copy(copyed, client.Backend))
		require.Equal(t, deployed.Registry.Address(), copyed.Registry.Address())
		require.Equal(t, deployed.Gov.Address(), copyed.Gov.Address())
		require.Equal(t, deployed.GovImp.Address(), copyed.GovImp.Address())
		require.Equal(t, deployed.Staking.Address(), copyed.Staking.Address())
		require.Equal(t, deployed.StakingImp.Address(), copyed.StakingImp.Address())
		require.Equal(t, deployed.BallotStorage.Address(), copyed.BallotStorage.Address())
		require.Equal(t, deployed.BallotStorageImp.Address(), copyed.BallotStorageImp.Address())
		require.Equal(t, deployed.EnvStorage.Address(), copyed.EnvStorage.Address())
		require.Equal(t, deployed.EnvStorageImp.Address(), copyed.EnvStorageImp.Address())
		require.True(t, deployed.Equal(copyed))
	})
	t.Run("init", func(t *testing.T) {
		require.NoError(t, copyed.Init(callOpts, client.Backend, deployed.Registry))
		require.Equal(t, deployed.Registry.Address(), copyed.Registry.Address())
		require.Equal(t, deployed.Gov.Address(), copyed.Gov.Address())
		require.NotEqual(t, deployed.GovImp.Address(), copyed.GovImp.Address())
		require.Equal(t, copyed.Gov.Address(), copyed.GovImp.Address())
		require.Equal(t, deployed.Staking.Address(), copyed.Staking.Address())
		require.NotEqual(t, deployed.StakingImp.Address(), copyed.StakingImp.Address())
		require.Equal(t, copyed.Staking.Address(), copyed.StakingImp.Address())
		require.Equal(t, deployed.BallotStorage.Address(), copyed.BallotStorage.Address())
		require.NotEqual(t, deployed.BallotStorageImp.Address(), copyed.BallotStorageImp.Address())
		require.Equal(t, copyed.BallotStorage.Address(), copyed.BallotStorageImp.Address())
		require.Equal(t, deployed.EnvStorage.Address(), copyed.EnvStorage.Address())
		require.NotEqual(t, deployed.EnvStorageImp.Address(), copyed.EnvStorageImp.Address())
		require.Equal(t, copyed.EnvStorage.Address(), copyed.EnvStorageImp.Address())
		require.False(t, deployed.Equal(copyed))
	})
	var byOwenr = new(gov.GovContracts)
	t.Run("GetGovContractsByOwner", func(t *testing.T) {
		byOwenr, err = gov.GetGovContractsByOwner(callOpts, client.Backend, client.Owner)
		require.NoError(t, err)
		require.Equal(t, copyed.Registry.Address(), byOwenr.Registry.Address())
		require.Equal(t, copyed.Gov.Address(), byOwenr.Gov.Address())
		require.Equal(t, copyed.GovImp.Address(), byOwenr.GovImp.Address())
		require.Equal(t, copyed.Staking.Address(), byOwenr.Staking.Address())
		require.Equal(t, copyed.StakingImp.Address(), byOwenr.StakingImp.Address())
		require.Equal(t, copyed.BallotStorage.Address(), byOwenr.BallotStorage.Address())
		require.Equal(t, copyed.BallotStorageImp.Address(), byOwenr.BallotStorageImp.Address())
		require.Equal(t, copyed.EnvStorage.Address(), byOwenr.EnvStorage.Address())
		require.Equal(t, copyed.EnvStorageImp.Address(), byOwenr.EnvStorageImp.Address())
		require.False(t, deployed.Equal(byOwenr))
		require.True(t, copyed.Equal(byOwenr))
	})
}

func TestDeploy(t *testing.T) {
	client := sim.NewClient(t)
	opts, err := bind.NewKeyedTransactorWithChainID(sim.EOAKey[sim.GetOrNewEOA("owner")], client.ChainID)
	require.NoError(t, err)
	go func() {
		for {
			time.Sleep(1e9)
			client.Commit()
		}
	}()
	contracts, err := gov.DeployGovContracts(opts, client.Backend)
	require.NoError(t, err)
	t.Run("not exec init", func(t *testing.T) {
		_, err := contracts.StakingImp.Funcs.Init(opts, contracts.Registry.Address(), []byte{})
		sim.ExpectedRevert(t, err, "Initializable: contract is already initialized")

		_, err = contracts.BallotStorageImp.Funcs.Initialize(opts, contracts.Registry.Address())
		sim.ExpectedRevert(t, err, "Initializable: contract is already initialized")

		envNames, envValues := makeEnvParams(
			EnvConstants.BLOCKS_PER,
			EnvConstants.BALLOT_DURATION_MIN,
			EnvConstants.BALLOT_DURATION_MAX,
			EnvConstants.STAKING_MIN,
			EnvConstants.STAKING_MAX,
			EnvConstants.MAX_IDLE_BLOCK_INTERVAL,
			EnvConstants.BLOCK_CREATION_TIME,
			EnvConstants.BLOCK_REWARD_AMOUNT,
			EnvConstants.MAX_PRIORITY_FEE_PER_GAS,
			EnvConstants.BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER,
			EnvConstants.BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD,
			EnvConstants.BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM,
			EnvConstants.BLOCK_REWARD_DISTRIBUTION_MAINTANANCE,
			EnvConstants.MAX_BASE_FEE,
			EnvConstants.BLOCK_GASLIMIT,
			EnvConstants.BASE_FEE_MAX_CHANGE_RATE,
			EnvConstants.GAS_TARGET_PERCENTAGE,
		)
		_, err = contracts.EnvStorageImp.Funcs.Initialize(opts, contracts.Registry.Address(), envNames, envValues)
		sim.ExpectedRevert(t, err, "Initializable: contract is already initialized")

		opts.Value = LOCK_AMOUNT
		_, err = contracts.StakingImp.Funcs.Deposit(opts)
		sim.ExpectedRevert(t, err, "") // init 이 되지 않았기 때문에 Registry 에 call 을 시도하면서 실패
		opts.Value = nil

		_, err = contracts.GovImp.Funcs.Init(opts, contracts.Registry.Address(), LOCK_AMOUNT,
			[]byte("name"),
			hexutil.MustDecode("0x6f8a80d14311c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
			[]byte("127.0.0.1"),
			bn.New(8542),
		)
		sim.ExpectedRevert(t, err, "Initializable: contract is already initialized")
	})
	t.Run("exec init", func(t *testing.T) {
		require.NoError(t, contracts.Init(&bind.CallOpts{}, client.Backend, contracts.Registry))
		txs := make([]*types.Transaction, 0)
		tx, err := contracts.StakingImp.Funcs.Init(opts, contracts.Registry.Address(), []byte{})
		require.NoError(t, err)
		txs = append(txs, tx)

		tx, err = contracts.BallotStorageImp.Funcs.Initialize(opts, contracts.Registry.Address())
		require.NoError(t, err)
		txs = append(txs, tx)

		envNames, envValues := makeEnvParams(
			EnvConstants.BLOCKS_PER,
			EnvConstants.BALLOT_DURATION_MIN,
			EnvConstants.BALLOT_DURATION_MAX,
			EnvConstants.STAKING_MIN,
			EnvConstants.STAKING_MAX,
			EnvConstants.MAX_IDLE_BLOCK_INTERVAL,
			EnvConstants.BLOCK_CREATION_TIME,
			EnvConstants.BLOCK_REWARD_AMOUNT,
			EnvConstants.MAX_PRIORITY_FEE_PER_GAS,
			EnvConstants.BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER,
			EnvConstants.BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD,
			EnvConstants.BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM,
			EnvConstants.BLOCK_REWARD_DISTRIBUTION_MAINTANANCE,
			EnvConstants.MAX_BASE_FEE,
			EnvConstants.BLOCK_GASLIMIT,
			EnvConstants.BASE_FEE_MAX_CHANGE_RATE,
			EnvConstants.GAS_TARGET_PERCENTAGE,
		)
		tx, err = contracts.EnvStorageImp.Funcs.Initialize(opts, contracts.Registry.Address(), envNames, envValues)
		require.NoError(t, err)
		txs = append(txs, tx)

		opts.Value = LOCK_AMOUNT
		tx, err = contracts.StakingImp.Funcs.Deposit(opts)
		require.NoError(t, err)
		txs = append(txs, tx)
		opts.Value = nil

		tx, err = contracts.GovImp.Funcs.Init(opts, contracts.Registry.Address(), LOCK_AMOUNT,
			[]byte("name"),
			hexutil.MustDecode("0x6f8a80d14311c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0"),
			[]byte("127.0.0.1"),
			bn.New(8542),
		)
		require.NoError(t, err)
		txs = append(txs, tx)
		for _, tx := range txs {
			receipt, err := bind.WaitMined(context.Background(), client.Backend, tx)
			require.NoError(t, err)
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		}
	})
}
