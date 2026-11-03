// Reproduction for the unbounded trace-RPC concurrency.
//
// The trace API ran every request immediately. Each rebuilds historical state,
// and block and chain traces also fan out to runtime.NumCPU workers. A burst of
// requests therefore saturated CPU and disk with no ceiling. This test sends
// more concurrent TraceCall requests than the configured limit and asserts that
// no more than the limit reach the expensive state rebuild at once. It fails
// (RED) without the traceSem gate and passes (GREEN) with it.
package tracers

import (
	"context"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

// gateBackend wraps a testBackend and makes StateAtBlock announce its entry and
// block until released, so a test can observe how many trace requests reach the
// state rebuild at the same time.
type gateBackend struct {
	*testBackend
	entered chan struct{}
	release chan struct{}
}

func (b *gateBackend) StateAtBlock(ctx context.Context, block *types.Block, reexec uint64, base *state.StateDB, checkLive, preferDisk bool) (*state.StateDB, error) {
	b.entered <- struct{}{} // announce arrival
	<-b.release             // hold the slot until the test lets go
	return b.testBackend.StateAtBlock(ctx, block, reexec, base, checkLive, preferDisk)
}

func TestTraceConcurrencyIsBounded(t *testing.T) {
	const limit = 2
	const callers = 5

	accounts := newAccounts(2)
	genesis := &core.Genesis{Alloc: core.GenesisAlloc{
		accounts[0].addr: {Balance: big.NewInt(params.Ether)},
		accounts[1].addr: {Balance: big.NewInt(params.Ether)},
	}}
	signer := types.HomesteadSigner{}
	tb := newTestBackend(t, 2, genesis, func(i int, b *core.BlockGen) {
		tx, _ := types.SignTx(types.NewTransaction(uint64(i), accounts[1].addr, big.NewInt(1000), params.TxGas, b.BaseFee(), nil), signer, accounts[0].key)
		b.AddTx(tx)
	})
	gb := &gateBackend{testBackend: tb, entered: make(chan struct{}, callers), release: make(chan struct{})}

	api := NewAPI(gb)
	api.traceSem = make(chan struct{}, limit) // deterministic limit for the test

	blockNum := rpc.BlockNumber(1)
	var wg sync.WaitGroup
	for i := 0; i < callers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _ = api.TraceCall(context.Background(),
				ethapi.TransactionArgs{From: &accounts[0].addr, To: &accounts[1].addr, Value: (*hexutil.Big)(big.NewInt(1))},
				rpc.BlockNumberOrHash{BlockNumber: &blockNum},
				nil)
		}()
	}

	// Exactly `limit` requests should reach StateAtBlock; the rest wait on the
	// semaphore. Drain the allowed arrivals, then confirm no further arrival
	// happens while they hold their slots.
	for i := 0; i < limit; i++ {
		select {
		case <-gb.entered:
		case <-time.After(5 * time.Second):
			t.Fatalf("only %d trace requests started; expected %d", i, limit)
		}
	}
	select {
	case <-gb.entered:
		t.Fatalf("more than %d trace requests ran concurrently; concurrency is not bounded", limit)
	case <-time.After(300 * time.Millisecond):
		// good: the extra callers are gated
	}

	close(gb.release) // let everything finish
	wg.Wait()
}
