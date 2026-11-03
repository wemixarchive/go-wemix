// Copyright 2026 The go-wemix Authors

package wemix

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	wemixminer "github.com/ethereum/go-ethereum/wemix/miner"
)

// returns42Code is runtime code that returns 42 as a 32-byte word.
var returns42Code = common.FromHex("602a60005260206000f3")

var (
	testContract = common.HexToAddress("0x0000000000000000000000000000000000c0de42")
	testBootAcct = common.HexToAddress("0x00000000000000000000000000000000000b0075")
)

// newTestChain builds a BlockChain holding only a genesis block whose extra
// data carries a 64-byte boot node id, as wemix genesis files do.
func newTestChain(t *testing.T) *core.BlockChain {
	t.Helper()
	db := rawdb.NewMemoryDatabase()
	genesis := &core.Genesis{
		Config:   params.TestChainConfig,
		GasLimit: 10_000_000,
		BaseFee:  big.NewInt(params.InitialBaseFee),
		Coinbase: testBootAcct,
		ExtraData: common.FromHex("0x" +
			"11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"),
		Alloc: core.GenesisAlloc{
			testContract: {Code: returns42Code, Balance: new(big.Int)},
		},
	}
	genesis.MustCommit(db)
	chain, err := core.NewBlockChain(db, nil, genesis.Config, ethash.NewFaker(), vm.Config{}, nil, nil)
	if err != nil {
		t.Fatalf("new blockchain: %v", err)
	}
	t.Cleanup(chain.Stop)
	return chain
}

// withAdmin replaces the package admin for the duration of the test.
func withAdmin(t *testing.T, ma *wemixAdmin) {
	t.Helper()
	old := admin
	admin = ma
	t.Cleanup(func() { admin = old })
}

func TestVerifyBlockSigWithoutAdmin(t *testing.T) {
	withAdmin(t, nil)
	if verifyBlockSig(big.NewInt(1), common.Address{}, nil, common.Hash{}, nil, true) {
		t.Fatal("verifyBlockSig accepted a block with no admin to verify against")
	}
}

func TestCalculateRewardsWithoutAdmin(t *testing.T) {
	withAdmin(t, nil)
	_, err := calculateRewards(params.TestChainConfig, big.NewInt(1), new(big.Int), nil)
	if !errors.Is(err, wemixminer.ErrNotInitialized) {
		t.Fatalf("err = %v, want %v", err, wemixminer.ErrNotInitialized)
	}
}

// startTestAdmin sets up an offline admin over chain for the duration of
// the test.
func startTestAdmin(t *testing.T, chain *core.BlockChain) {
	t.Helper()
	withAdmin(t, nil)
	stop, err := StartOfflineAdmin(chain)
	if err != nil {
		t.Fatalf("StartOfflineAdmin: %v", err)
	}
	t.Cleanup(stop)
}

func TestOfflineChainAPI(t *testing.T) {
	startTestAdmin(t, newTestChain(t))
	c := admin.cli
	ctx := context.Background()

	out, err := c.CallContract(ctx, ethereum.CallMsg{To: &testContract}, big0)
	if err != nil {
		t.Fatalf("CallContract: %v", err)
	}
	if got := new(big.Int).SetBytes(out); got.Int64() != 42 {
		t.Fatalf("CallContract = %v, want 42", got)
	}

	code, err := c.CodeAt(ctx, testContract, nil)
	if err != nil {
		t.Fatalf("CodeAt: %v", err)
	}
	if string(code) != string(returns42Code) {
		t.Fatalf("CodeAt = %x, want %x", code, returns42Code)
	}

	genesis, err := c.HeaderByNumber(ctx, big0)
	if err != nil {
		t.Fatalf("HeaderByNumber(0): %v", err)
	}
	if byHash, err := c.HeaderByHash(ctx, genesis.Hash()); err != nil || byHash.Hash() != genesis.Hash() {
		t.Fatalf("HeaderByHash = %v, %v, want %x", byHash, err, genesis.Hash())
	}

	// A block the chain does not have yet must look the same as it does on
	// a running node, so governance lookups treat it as not found.
	if _, err := c.HeaderByNumber(ctx, big.NewInt(1)); !errors.Is(err, ethereum.NotFound) {
		t.Fatalf("HeaderByNumber(1) err = %v, want %v", err, ethereum.NotFound)
	}
	if _, err := c.CallContract(ctx, ethereum.CallMsg{To: &testContract}, big.NewInt(1)); err == nil {
		t.Fatal("CallContract on a missing block succeeded")
	}
	// Methods outside what governance reads need are not served.
	if _, err := c.PendingNonceAt(ctx, testContract); err == nil {
		t.Fatal("PendingNonceAt succeeded on the offline server")
	}
}

func TestStartOfflineAdmin(t *testing.T) {
	withAdmin(t, nil)
	chain := newTestChain(t)

	stop, err := StartOfflineAdmin(chain)
	if err != nil {
		t.Fatalf("StartOfflineAdmin: %v", err)
	}
	if admin == nil || admin.bootAccount != testBootAcct {
		t.Fatalf("admin not set up from genesis: %+v", admin)
	}
	if _, err := StartOfflineAdmin(chain); !errors.Is(err, ErrAlreadyRunning) {
		t.Fatalf("second StartOfflineAdmin err = %v, want %v", err, ErrAlreadyRunning)
	}
	// Before governance is deployed the registry lookup must report not
	// found, as it does on a running node, so blocks from that period pass.
	if _, err := admin.getRegGovEnvContracts(nil, big0); !errors.Is(err, ethereum.NotFound) {
		t.Fatalf("getRegGovEnvContracts err = %v, want %v", err, ethereum.NotFound)
	}

	stop()
	if admin != nil {
		t.Fatal("admin still set after stop")
	}
}
