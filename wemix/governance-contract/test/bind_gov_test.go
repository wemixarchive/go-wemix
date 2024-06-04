package test

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	gov "github.com/ethereum/go-ethereum/wemix/bind"
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
		deployed, err = gov.DeployGovContracts(opts, client.Backend, nil)
		require.NoError(t, err)
		zeroAddress := common.Address{}
		require.NotEqual(t, zeroAddress, deployed.Address().Registry)
		require.NotEqual(t, zeroAddress, deployed.Address().Gov)
		require.NotEqual(t, zeroAddress, deployed.Address().Staking)
		require.NotEqual(t, zeroAddress, deployed.Address().BallotStorage)
		require.NotEqual(t, zeroAddress, deployed.Address().EnvStorage)
	})
	var copyed = new(gov.GovContracts)
	t.Run("copy", func(t *testing.T) {
		require.NoError(t, deployed.Copy(copyed, client.Backend))
		require.Equal(t, deployed.Address().Registry, copyed.Address().Registry)
		require.Equal(t, deployed.Address().Gov, copyed.Address().Gov)
		require.Equal(t, deployed.Address().Staking, copyed.Address().Staking)
		require.Equal(t, deployed.Address().BallotStorage, copyed.Address().BallotStorage)
		require.Equal(t, deployed.Address().EnvStorage, copyed.Address().EnvStorage)
		require.True(t, deployed.Equal(copyed))
	})
	var byOwenr = new(gov.GovContracts)
	t.Run("GetGovContractsByOwner", func(t *testing.T) {
		byOwenr, err = gov.GetGovContractsByOwner(callOpts, client.Backend, client.Owner)
		require.NoError(t, err)
		require.Equal(t, copyed.Address().Registry, byOwenr.Address().Registry)
		require.Equal(t, copyed.Address().Gov, byOwenr.Address().Gov)
		require.Equal(t, copyed.Address().Staking, byOwenr.Address().Staking)
		require.Equal(t, copyed.Address().BallotStorage, byOwenr.Address().BallotStorage)
		require.Equal(t, copyed.Address().EnvStorage, byOwenr.Address().EnvStorage)
		require.True(t, deployed.Equal(byOwenr))
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
	contracts, err := gov.DeployGovContracts(opts, client.Backend, nil)
	require.NoError(t, err)
	lockAmount := gov.DefaultInitEnvStorage.STAKING_MIN
	gov.ExecuteInitialize(contracts, opts, client.Backend, lockAmount, gov.DefaultInitEnvStorage, gov.InitMembers{
		{
			Staker:  opts.From,
			Voter:   opts.From,
			Reward:  opts.From,
			Name:    "name",
			Enode:   "0x6f8a80d14311c39f35f516fa664deaaaa13e85b2f7493f37f6144d86991ec012937307647bd3b9a82abe2974e1407241d54947bbb39763a4cac9f77166ad92a0",
			Ip:      "127.0.0.1",
			Port:    8542,
			Deposit: lockAmount,
		},
	})

	checkMainnetEnvStorageValues(t, new(bind.CallOpts), contracts.EnvStorageImp)
}

func TestCheckMainnetEnvStorageValues(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := ethclient.DialContext(ctx, "https://api.wemix.com/")
	require.NoError(t, err)

	block, err := client.BlockByNumber(ctx, common.Big0)
	require.NoError(t, err)

	callOpts := &bind.CallOpts{Context: ctx}
	contracts, err := gov.GetGovContractsByOwner(callOpts, client, block.Coinbase())
	require.NoError(t, err)

	checkMainnetEnvStorageValues(t, callOpts, contracts.EnvStorageImp)
}

func checkMainnetEnvStorageValues(t *testing.T, callOpts *bind.CallOpts, envStorage *gov.EnvStorageImp) {
	BLOCKS_PER, _ := envStorage.GetBlocksPer(callOpts)
	t.Log("BLOCKS_PER:", BLOCKS_PER)

	BALLOT_DURATION_MIN, _ := envStorage.GetBallotDurationMin(callOpts)
	t.Log("BALLOT_DURATION_MIN:", BALLOT_DURATION_MIN)

	BALLOT_DURATION_MAX, _ := envStorage.GetBallotDurationMax(callOpts)
	t.Log("BALLOT_DURATION_MAX:", BALLOT_DURATION_MAX)

	STAKING_MIN, _ := envStorage.GetStakingMin(callOpts)
	t.Log("STAKING_MIN:", STAKING_MIN)

	STAKING_MAX, _ := envStorage.GetStakingMax(callOpts)
	t.Log("STAKING_MAX:", STAKING_MAX)

	MAX_IDLE_BLOCK_INTERVAL, _ := envStorage.GetMaxIdleBlockInterval(callOpts)
	t.Log("MAX_IDLE_BLOCK_INTERVAL:", MAX_IDLE_BLOCK_INTERVAL)

	BLOCK_CREATION_TIME, _ := envStorage.GetBlockCreationTime(callOpts)
	t.Log("BLOCK_CREATION_TIME:", BLOCK_CREATION_TIME)

	BLOCK_REWARD_AMOUNT, _ := envStorage.GetBlockRewardAmount(callOpts)
	t.Log("BLOCK_REWARD_AMOUNT:", BLOCK_REWARD_AMOUNT)

	MAX_PRIORITY_FEE_PER_GAS, _ := envStorage.GetMaxPriorityFeePerGas(callOpts)
	t.Log("MAX_PRIORITY_FEE_PER_GAS:", MAX_PRIORITY_FEE_PER_GAS)

	BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER, _ := envStorage.GetUint(callOpts, crypto.Keccak256Hash([]byte("blockRewardDistributionBlockProducer")))
	t.Log("BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER:", BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER)

	BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD, _ := envStorage.GetUint(callOpts, crypto.Keccak256Hash([]byte("blockRewardDistributionStakingReward")))
	t.Log("BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD:", BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD)

	BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM, _ := envStorage.GetUint(callOpts, crypto.Keccak256Hash([]byte("blockRewardDistributionEcosystem")))
	t.Log("BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM:", BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM)

	BLOCK_REWARD_DISTRIBUTION_MAINTENANCE, _ := envStorage.GetUint(callOpts, crypto.Keccak256Hash([]byte("blockRewardDistributionMaintenance")))
	t.Log("BLOCK_REWARD_DISTRIBUTION_MAINTENANCE:", BLOCK_REWARD_DISTRIBUTION_MAINTENANCE)

	MAX_BASE_FEE, _ := envStorage.GetMaxBaseFee(callOpts)
	t.Log("MAX_BASE_FEE:", MAX_BASE_FEE)

	BLOCK_GASLIMIT, _ := envStorage.GetUint(callOpts, crypto.Keccak256Hash([]byte("blockGasLimit")))
	t.Log("BLOCK_GASLIMIT:", BLOCK_GASLIMIT)

	BASE_FEE_MAX_CHANGE_RATE, _ := envStorage.GetUint(callOpts, crypto.Keccak256Hash([]byte("baseFeeMaxChangeRate")))
	t.Log("BASE_FEE_MAX_CHANGE_RATE:", BASE_FEE_MAX_CHANGE_RATE)

	GAS_TARGET_PERCENTAGE, _ := envStorage.GetUint(callOpts, crypto.Keccak256Hash([]byte("gasTargetPercentage")))
	t.Log("GAS_TARGET_PERCENTAGE:", GAS_TARGET_PERCENTAGE)
}
