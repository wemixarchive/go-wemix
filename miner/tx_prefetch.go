// tx_prefetch.go

package miner

import (
	"context"
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

func tx_prefetch(w *worker, env *environment, to *TxOrderer, numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go func() {
			var (
				ctx          = context.Background()
				header       = w.chain.GetHeaderByNumber(env.header.Number.Uint64() - 1)
				gaspool      = new(core.GasPool).AddGas(header.GasLimit)
				blockContext = core.NewEVMBlockContext(header, w.chain, nil)
				statedb, _   = w.chain.StateAt(header.Root)
				evm          = vm.NewEVM(blockContext, vm.TxContext{}, statedb, w.chainConfig, vm.Config{})
				signer       = types.MakeSigner(w.chainConfig, header.Number)
				tix          = 0
			)

			for {
				tx := to.NextForPrefetch()
				if tx == nil {
					break
				} else if doneTxs.Exists(tx.Hash()) {
					continue
				} else {
					doneTxs.Put(tx.Hash(), true)
				}

				_, err := types.Sender(signer, tx)
				if err != nil {
					log.Error("Prefetch", "tx -> sender", err)
					continue
				}

				msg, err := tx.AsMessage(signer, header.BaseFee)
				if err != nil {
					log.Error("Prefetch", "failed to turn tx to msg", err)
					return
				}
				statedb.Prepare(tx.Hash(), tix)
				tix++

				ictx, cancel := context.WithTimeout(ctx, time.Second)
				evm.Reset(core.NewEVMTxContext(msg), statedb)
				go func() {
					<-ictx.Done()
					evm.Cancel()
				}()
				_, err = core.ApplyMessage(evm, msg, gaspool)
				if err != nil {
					log.Error("Prefetch", "ApplyMessage failed", err)
				}
				cancel()
			}
		}()
	}
}

// EOF
