// governancedeploy.js

package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	gov "github.com/ethereum/go-ethereum/wemix/bind"
	"github.com/stretchr/testify/require"
)

func TestDeployGoverananceContracts(t *testing.T) {
	var Bootnodes []string = params.MainnetBootnodes

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

	address := make([]common.Address, 0)
	address = append(address, opts.From)
	for i := 1; i < len(Bootnodes); i++ {
		address = append(address, common.Address{byte(i)})
	}

	// make temp config.json
	configJSFile := filepath.Join(t.TempDir(), "config.json")
	memberStr := func(addr common.Address, enode string) string {
		enode = strings.TrimLeft(enode, "enode://")
		sp := strings.Split(enode, "@")
		t.Log(enode, sp)
		id := sp[0]
		if len(id) != 128 {
			return ""
		}
		sp = strings.Split(sp[1], ":")
		ip := sp[0]
		port := sp[1]
		return fmt.Sprintf(`{"addr":"%s","stake":4980000000000000000000000,"name":"%s","id":"%s","ip":"%s","port":%s,"bootnode":true},`, addr.Hex(), id[:8], id, ip, port)
	}
	var configStr string = `{"members":[`
	for i, enode := range Bootnodes {
		configStr += memberStr(address[i], enode)
	}
	configStr = configStr[:len(configStr)-1] + "],"

	// to avoid error 'at least one account and node are required'
	configStr += `"accounts":[{"addr":"0x0000000000000000000000000000000000000000","balance":0}]}`
	require.NoError(t, os.WriteFile(configJSFile, []byte(configStr), 0775))

	go func() {
		for {
			time.Sleep(0.5e9)
			backend.Commit()
		}
	}()

	require.NoError(t, deployGovernance(backend, opts, gov.DefaultInitEnvStorage.STAKING_MIN, configJSFile))
}
