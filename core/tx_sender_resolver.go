// tx_sender_resolver

package core

import (
	_ "fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SenderResolver resolves sender accounts from transactions concurrently
// with worker threads
type SenderResolver struct {
	jobs chan func()
	busy chan interface{}
}

// NewSenderResolver creates a new sender resolver worker pool
func NewSenderResolver(count int) *SenderResolver {
	return &SenderResolver{
		jobs: make(chan func(), count),
		busy: make(chan interface{}, count),
	}
}

// sender resolver main loop
func (s *SenderResolver) Run() {
	eor := false
	for {
		select {
		case f := <-s.jobs:
			if f == nil {
				eor = true
			} else {
				go func() {
					s.busy <- struct{}{}
					defer func() {
						<-s.busy
					}()
					f()
				}()
			}
		}
		if eor {
			break
		}
	}
}

// Stop stops sender resolver
func (s *SenderResolver) Stop() {
	s.jobs <- nil
}

// Post a new sender resolver task
func (s *SenderResolver) Post(f func()) {
	s.jobs <- f
}

// ResolveSenders resolves sender accounts from given transactions
// concurrently using SenderResolver worker pool. If 'checkPool' is true,
// it checks the pending pool first
func (pool *TxPool) ResolveSenders(signer types.Signer, txs []*types.Transaction, checkPool bool) []*common.Address {
	var total, by_ecrecover, failed int64
	total = int64(len(txs))
	ot := time.Now()

	var addrs []*common.Address
	if checkPool {
		addrs = pool.all.ResolveSenders(signer, txs)
	} else {
		addrs = make([]*common.Address, len(txs))
	}

	var wg sync.WaitGroup
	for i, addr := range addrs {
		if addr != nil {
			types.SetSender(signer, txs[i], *addr)
			continue
		}
		atomic.AddInt64(&by_ecrecover, 1)

		wg.Add(1)

		f := func(ix int) {
			if from, err := types.Sender(signer, txs[ix]); err == nil {
				addrs[ix] = &from
			} else {
				atomic.AddInt64(&failed, 1)
			}
			wg.Done()
		}
		pool.senderResolver.Post(func() { f(i) })
	}

	wg.Wait()

	if total > 1 {
		dt := float64(time.Now().Sub(ot) / time.Millisecond)
		if dt <= 0 {
			dt = 1
		}
		ps := float64(total) * 1000.0 / dt
		_ = ps
		//fmt.Printf("=== %d/%d/%d : took %.3f ms %.3f/sec\n", total, total-by_ecrecover, failed, dt, ps)
	}

	return addrs
}

// EOF
