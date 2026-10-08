package miner

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	wemixminer "github.com/ethereum/go-ethereum/wemix/miner"
)

var errInjectedTokenFailure = errors.New("injected failure")

// miningTokenModel mirrors the observable contract of the etcd mining token in
// wemix/sync.go: acquire sets the held token, release or abandon clears it,
// HasMiningToken reports it.
type miningTokenModel struct {
	held      bool
	acquired  int
	released  int
	abandoned int
}

// tokenHookOptions selects the failure injected into a single commitWork run.
type tokenHookOptions struct {
	coinbaseErr error
	signErr     error
	// onAcquire runs inside AcquireMiningToken, after the token is taken.
	onAcquire func()
}

// installMiningTokenHooks wires the wemix miner hooks to the model and restores
// the previous hooks when the test ends.
func installMiningTokenHooks(t *testing.T, m *miningTokenModel, opts tokenHookOptions) {
	t.Helper()
	prevMethod := params.ConsensusMethod
	prevAcquire := wemixminer.AcquireMiningTokenFunc
	prevRelease := wemixminer.ReleaseMiningTokenFunc
	prevHas := wemixminer.HasMiningTokenFunc
	prevCoinbase := wemixminer.GetCoinbaseFunc
	prevSign := wemixminer.SignBlockFunc
	prevAbandon := wemixminer.AbandonMiningTokenFunc
	t.Cleanup(func() {
		params.ConsensusMethod = prevMethod
		wemixminer.AcquireMiningTokenFunc = prevAcquire
		wemixminer.ReleaseMiningTokenFunc = prevRelease
		wemixminer.HasMiningTokenFunc = prevHas
		wemixminer.GetCoinbaseFunc = prevCoinbase
		wemixminer.SignBlockFunc = prevSign
		wemixminer.AbandonMiningTokenFunc = prevAbandon
	})
	params.ConsensusMethod = params.ConsensusPoA
	wemixminer.AcquireMiningTokenFunc = func(*big.Int, common.Hash) (bool, error) {
		m.held = true
		m.acquired++
		if opts.onAcquire != nil {
			opts.onAcquire()
		}
		return true, nil
	}
	wemixminer.ReleaseMiningTokenFunc = func(*big.Int, common.Hash, common.Hash) error {
		m.held = false
		m.released++
		return nil
	}
	wemixminer.HasMiningTokenFunc = func() bool { return m.held }
	wemixminer.AbandonMiningTokenFunc = func() error {
		m.held = false
		m.abandoned++
		return nil
	}
	wemixminer.GetCoinbaseFunc = func(*big.Int) (common.Address, error) {
		return testBankAddress, opts.coinbaseErr
	}
	wemixminer.SignBlockFunc = func(*big.Int, common.Hash) (common.Address, []byte, error) {
		return testBankAddress, nil, opts.signErr
	}
}

// newSyncTokenWorker builds a worker without background loops so commitWork
// runs synchronously on the test goroutine.
func newSyncTokenWorker(t *testing.T) *worker {
	t.Helper()
	engine := ethash.NewFaker()
	backend := newTestWorkerBackend(t, ethashChainConfig, engine, rawdb.NewMemoryDatabase(), 0)
	t.Cleanup(func() {
		backend.txPool.Stop()
		backend.chain.Stop()
	})
	return &worker{
		config:       testConfig,
		chainConfig:  ethashChainConfig,
		engine:       engine,
		eth:          backend,
		mux:          new(event.TypeMux),
		chain:        backend.chain,
		localUncles:  make(map[common.Hash]*types.Block),
		remoteUncles: make(map[common.Hash]*types.Block),
		unconfirmed:  newUnconfirmedBlocks(backend.chain, sealingLogAtDepth),
		pendingTasks: make(map[common.Hash]*task),
		taskCh:       make(chan *task, 1),
		exitCh:       make(chan struct{}),
		coinbase:     testBankAddress,
		running:      1,
	}
}

// TestCommitWorkGivesBackMiningTokenWithoutBlock checks that every commitWork
// exit after AcquireMiningToken that writes no block leaves the token free.
// A held token blocks the whole cluster at that height until its Till.
func TestCommitWorkGivesBackMiningTokenWithoutBlock(t *testing.T) {
	cases := []struct {
		name string
		opts func(w *worker) tokenHookOptions
	}{
		{
			// downloader.StartEvent calls worker.stop() while the token is being taken.
			name: "worker stopped after token acquired",
			opts: func(w *worker) tokenHookOptions { return tokenHookOptions{onAcquire: w.stop} },
		},
		{
			name: "no etherbase",
			opts: func(w *worker) tokenHookOptions {
				w.coinbase = common.Address{}
				return tokenHookOptions{}
			},
		},
		{
			name: "GetCoinbase error",
			opts: func(*worker) tokenHookOptions { return tokenHookOptions{coinbaseErr: errInjectedTokenFailure} },
		},
		{
			name: "SignBlock error",
			opts: func(*worker) tokenHookOptions { return tokenHookOptions{signErr: errInjectedTokenFailure} },
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := newSyncTokenWorker(t)
			m := &miningTokenModel{}
			installMiningTokenHooks(t, m, tc.opts(w))

			w.commitWork(nil, false, 0)

			if m.acquired != 1 {
				t.Fatalf("token not acquired: acquired=%d", m.acquired)
			}
			if head := w.chain.CurrentBlock().NumberU64(); head != 0 {
				t.Fatalf("block written (head=%d), failure was not injected", head)
			}
			if m.held {
				t.Fatalf("mining token still held after commitWork wrote no block: released=%d abandoned=%d",
					m.released, m.abandoned)
			}
		})
	}
}
