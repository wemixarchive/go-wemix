// tx_prefetch.go

package miner

import (
	"context"
	"math"
	"time"

	"github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
)

var (
	doneTxs = lru.NewLruCache(5000, false)
)

func tx_prefetch(w *worker, to *TxOrderer, numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go func() {
			ctx := context.Background()
			vmCfg := vm.Config{}
			header := w.chain.GetHeaderByNumber(w.current.header.Number.Uint64() - 1)
			state, err := w.chain.StateAt(header.Root)
			if err != nil {
				log.Error("Prefetch", "failed to get state db", err)
				return
			}

			for {
				tx := to.NextForPrefetch()
				if tx == nil {
					break
				} else if doneTxs.Exists(tx.Hash()) {
					continue
				} else {
					doneTxs.Put(tx.Hash(), true)
				}

				from, err := types.Sender(
					types.MakeSigner(w.chainConfig, header.Number), tx)
				if err != nil {
					log.Error("Prefetch", "tx -> sender", err)
					continue
				}

				state.Reset(header.Root)
				msg := types.NewMessage(from, tx.To(), 0, tx.Value(), tx.Gas(),
					tx.GasPrice(), tx.Data(), false)
				ictx, cancel := context.WithTimeout(ctx, time.Second)
				vmCtx := core.NewEVMContext(msg, header, w.chain, nil)
				evm := vm.NewEVM(vmCtx, state, w.chainConfig, vmCfg)

				go func() {
					<-ictx.Done()
					evm.Cancel()
				}()

				gp := new(core.GasPool).AddGas(math.MaxUint64)
				_, _, _, _, err = core.ApplyMessage(evm, msg, gp)
				if err != nil {
					log.Error("Prefetch", "ApplyMessage failed", err)
				}
				cancel()
			}
		}()
	}
}

// EOF
