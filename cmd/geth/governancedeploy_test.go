// governancedeploy.js

package main

import (
	"context"
	"encoding/json"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	gov "github.com/ethereum/go-ethereum/wemix/bind"
	"github.com/stretchr/testify/require"
)

func TestDeployGoverananceContracts(t *testing.T) {
	owner, err := crypto.GenerateKey()
	require.NoError(t, err)

	opts, err := bind.NewKeyedTransactorWithChainID(owner, params.AllEthashProtocolChanges.ChainID)
	require.NoError(t, err)

	backend := backends.NewSimulatedBackend(
		core.GenesisAlloc{
			opts.From: {Balance: new(big.Int).Mul(big.NewInt(math.MaxInt64), big.NewInt(params.Ether))},
		},
		params.MaxGasLimit,
	)

	fin := make(chan struct{})
	defer close(fin)

	go func() {
		ticker := time.NewTicker(0.1e9)
		defer ticker.Stop()

		nonce, err := backend.NonceAt(context.Background(), opts.From, nil)
		require.NoError(t, err)

		for {
			select {
			case <-ticker.C:
				pending, err := backend.PendingNonceAt(context.Background(), opts.From)
				require.NoError(t, err)
				if nonce != pending {
					nonce = pending
					backend.Commit()
				}
			case <-fin:
				return
			}
		}
	}()

	configJSFile := filepath.Join(t.TempDir(), "config.json")
	require.NoError(t, os.WriteFile(configJSFile, []byte(configStr), 0775))

	domains, env, members, err := getInitialGovernanceInitDatas(configJSFile)
	require.NoError(t, err)
	for name, addr := range domains {
		t.Log(name, addr)
	}
	bytes, _ := json.MarshalIndent(env, "", "  ")
	t.Log(string(bytes))
	for i, member := range members {
		t.Log(i, member)
	}

	require.NoError(t, deployGovernance(backend, opts, gov.DefaultInitEnvStorage.STAKING_MIN, configJSFile))
}

// config.json.example
const configStr string = `
{
	"extraData": "The beginning of Wemix3.0 testnet on July 1st, 2022",
	"staker": "0xf00d9928ed1dada205aec56ab85e0e2ab5670ad5",
	"ecosystem": "0x1be19928ed1dada205aec56ab85e0e2ab5670ad5",
	"maintenance": "0x900d9928ed1dada205aec56ab85e0e2ab5670ad5",
	"feecollector": "0x900d9928ed1dada205aec56ab85e0e2ab5670ad5",
	"env": {
	  "ballotDurationMin": 86400,
	  "ballotDurationMax": 604800,
	  "stakingMin": 1500000000000000000000000,
	  "stakingMax": 1500000000000000000000000,
	  "MaxIdleBlockInterval": 5,
	  "blockCreationTime": 1000,
	  "blockRewardAmount": 1000000000000000000,
	  "maxPriorityFeePerGas": 100000000000,
	  "rewardDistributionMethod": [ 4000, 1000, 2500, 2500 ],
	  "maxBaseFee": 50000000000000,
	  "blockGasLimit": 105000000,
	  "baseFeeMaxChangeRate": 55,
	  "gasTargetPercentage": 30
	},
	"members": [
	  {
		"addr": "0x1be19928ed1dada205aec56ab85e0e2ab5670ad5",
		"stake": 2000000000000000000000000000,
		"name": "fak1",
		"id": "0xfa5f92fc954e4e45ac5773d5472bf8ab0b888979a5e65d49bac65e1b4345e82a745e255d1018d7a6de7ae8fd3a04b0e8eca4359f0fd35c2d0c45d29f7ffa0290",
		"ip": "172.18.0.1",
		"port": 8589,
		"bootnode": true
	  },
	  {
		"addr": "0xb4388353fd0f3b3a017e09f2b857052ff219e663",
		"stake": 2000000000000000000000000000,
		"name": "fak2",
		"id": "0xea5f92fc954e4e45ac5773d5472bf8ab0b888979a5e65d49bac65e1b4345e82a745e255d1018d7a6de7ae8fd3a04b0e8eca4359f0fd35c2d0c45d29f7ffa0290",
		"ip": "172.18.0.2",
		"port": 8589
	  }
	],
	"accounts": [
	  {
		"addr": "0x1be19928ed1dada205aec56ab85e0e2ab5670ad5",
		"balance": 200000000000000000000000
	  },
	  {
		"addr": "0xb4388353fd0f3b3a017e09f2b857052ff219e663",
		"balance": 200000000000000000000000
	  }
	]
}
`
