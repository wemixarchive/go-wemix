package core

import (
	"crypto/ecdsa"
	"math/big"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

// feeDelegatedTransfer builds a 0x16 transaction: the sender signs the inner
// dynamic-fee body and the fee payer signs over it.
func feeDelegatedTransfer(t *testing.T, pool *TxPool, nonce uint64, senderKey, payerKey *ecdsa.PrivateKey) *types.Transaction {
	t.Helper()
	inner := types.MustSignNewTx(senderKey, pool.signer, &types.DynamicFeeTx{
		ChainID:   params.TestChainConfig.ChainID,
		Nonce:     nonce,
		GasTipCap: big.NewInt(1),
		GasFeeCap: big.NewInt(1),
		Gas:       params.TxGas,
		To:        &common.Address{},
		Value:     big.NewInt(0),
	})
	v, r, s := inner.RawSignatureValues()
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)
	fd := &types.FeeDelegateDynamicFeeTx{FeePayer: &payer}
	fd.SetSenderTx(types.DynamicFeeTx{
		ChainID:   inner.ChainId(),
		Nonce:     nonce,
		GasTipCap: inner.GasTipCap(),
		GasFeeCap: inner.GasFeeCap(),
		Gas:       inner.Gas(),
		To:        inner.To(),
		Value:     inner.Value(),
		V:         v,
		R:         r,
		S:         s,
	})
	tx, err := types.SignTx(types.NewTx(fd), types.NewFeeDelegateSigner(params.TestChainConfig.ChainID), payerKey)
	if err != nil {
		t.Fatalf("sign fee delegated tx: %v", err)
	}
	return tx
}

// Tests that when demoteUnexecutables drops a fee-delegated pending transaction
// because its fee payer can no longer pay, the sender's higher-nonce
// transactions are moved back to the queue instead of being left in the lookup
// set with no list holding them. Followers that are themselves unpayable are
// dropped, not queued: a queue made only of them is what the eviction loop
// trips on.
func TestFeeDelegationDemoteRequeuesFollowingTxs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		// delegated reports whether the transaction at each nonce is fee
		// delegated by the payer that runs dry. Nonce 0 always is.
		delegated []bool
	}{
		{name: "ordinary follower", delegated: []bool{true, false}},
		{name: "ordinary then delegated followers", delegated: []bool{true, false, true}},
		{name: "delegated follower", delegated: []bool{true, true}},
		{name: "delegated then ordinary followers", delegated: []bool{true, true, false}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pool, senderKey := setupTxPool()
			defer pool.Stop()

			pool.mu.Lock()
			pool.feedelegation = true
			pool.mu.Unlock()

			payerKey, _ := crypto.GenerateKey()
			sender := crypto.PubkeyToAddress(senderKey.PublicKey)
			payer := crypto.PubkeyToAddress(payerKey.PublicKey)
			testAddBalance(pool, sender, big.NewInt(1_000_000))
			testAddBalance(pool, payer, big.NewInt(1_000_000))

			txs := make([]*types.Transaction, len(tt.delegated))
			for nonce, delegated := range tt.delegated {
				if delegated {
					txs[nonce] = feeDelegatedTransfer(t, pool, uint64(nonce), senderKey, payerKey)
				} else {
					txs[nonce] = dynamicFeeTx(uint64(nonce), params.TxGas, big.NewInt(1), big.NewInt(1), senderKey)
				}
			}
			for i, err := range pool.AddRemotesSync(txs) {
				if err != nil {
					t.Fatalf("add tx %d: %v", i, err)
				}
			}
			if pending, queued := pool.stats(); pending != len(txs) || queued != 0 {
				t.Fatalf("before demote: pending %d queued %d, want %d and 0", pending, queued, len(txs))
			}

			pool.mu.Lock()
			pool.currentState.SetBalance(payer, common.Big0)
			pool.demoteUnexecutables()
			pool.mu.Unlock()

			if err := validateTxPoolInternals(pool); err != nil {
				t.Fatalf("pool internal state corrupted: %v", err)
			}
			// Ordinary transactions are re-queued; delegated ones share the payer
			// that ran dry, so they leave the pool entirely.
			queue := pool.queue[sender]
			for nonce, delegated := range tt.delegated {
				queued := queue != nil && queue.txs.Get(uint64(nonce)) != nil
				known := pool.all.Get(txs[nonce].Hash()) != nil
				switch {
				case delegated && (queued || known):
					t.Errorf("unpayable fee delegated tx at nonce %d kept: queued=%v known=%v", nonce, queued, known)
				case !delegated && !queued:
					t.Errorf("tx at nonce %d not re-queued", nonce)
				}
			}
		})
	}
}

// evictChildEnv makes TestFeeDelegationEvictEmptiedQueue run its scenario in
// place. The parent test runs it in a child process because, without the fix,
// the panic happens on the pool loop goroutine and takes the binary down.
const evictChildEnv = "WEMIX_TXPOOL_EVICT_CHILD"

// Tests that the eviction tick survives when dropping an unpayable
// fee-delegated transaction empties an account's queue. removeTx then deletes
// queue[addr] and beats[addr], and the lifetime check that follows must not
// read them.
func TestFeeDelegationEvictEmptiedQueue(t *testing.T) {
	if os.Getenv(evictChildEnv) != "" {
		evictEmptiedQueueChild(t)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestFeeDelegationEvictEmptiedQueue$", "-test.v")
	cmd.Env = append(os.Environ(), evictChildEnv+"=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("eviction tick crashed the pool: %v\n%s", err, out)
	}
}

func evictEmptiedQueueChild(t *testing.T) {
	evictionInterval = 50 * time.Millisecond

	pool, senderKey := setupTxPool()
	defer pool.Stop()

	pool.mu.Lock()
	pool.feedelegation = true
	pool.mu.Unlock()

	payerKey, _ := crypto.GenerateKey()
	sender := crypto.PubkeyToAddress(senderKey.PublicKey)
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)
	testAddBalance(pool, sender, big.NewInt(1_000_000))
	testAddBalance(pool, payer, big.NewInt(1_000_000))

	// The nonce gap keeps the transaction in the queue as its only entry.
	tx := feeDelegatedTransfer(t, pool, 1, senderKey, payerKey)
	if err := pool.AddRemotesSync([]*types.Transaction{tx})[0]; err != nil {
		t.Fatalf("add tx: %v", err)
	}
	if pending, queued := pool.stats(); pending != 0 || queued != 1 {
		t.Fatalf("before evict: pending %d queued %d, want 0 and 1", pending, queued)
	}
	// Drain the fee payer without a new head, as when no block arrives before
	// the next eviction tick.
	pool.mu.Lock()
	pool.currentState.SetBalance(payer, common.Big0)
	pool.mu.Unlock()

	// The sweep and the lifetime check run under one hold of pool.mu, so once
	// the transaction is gone under the lock the whole iteration has finished.
	deadline := time.Now().Add(5 * time.Second)
	for {
		pool.mu.RLock()
		gone := pool.all.Get(tx.Hash()) == nil
		pool.mu.RUnlock()
		if gone {
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("unpayable fee delegated tx was never evicted")
		}
		time.Sleep(evictionInterval / 5)
	}
	if err := validateTxPoolInternals(pool); err != nil {
		t.Fatalf("pool internal state corrupted: %v", err)
	}
}
