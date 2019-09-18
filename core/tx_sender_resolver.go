// tx_sender_resolver

package core

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/core/types"
)

// job structure is needed because param is necessary due to evaluation order
// uncertainty with closure and channel.
type job struct {
	f     func(interface{})
	param interface{}
}

// SenderResolver resolves sender accounts from transactions concurrently
// with worker threads
type SenderResolver struct {
	tx2addr *lru.LruCache
	jobs    chan *job
	busy    chan interface{}
}

// NewSenderResolver creates a new sender resolver worker pool
func NewSenderResolver(concurrency, cacheSize int) *SenderResolver {
	return &SenderResolver{
		tx2addr: lru.NewLruCache(cacheSize, true),
		jobs:    make(chan *job, concurrency),
		busy:    make(chan interface{}, concurrency),
	}
}

// sender resolver main loop
func (s *SenderResolver) Run() {
	eor := false
	for {
		select {
		case j := <-s.jobs:
			if j == nil {
				eor = true
			} else {
				go func() {
					s.busy <- struct{}{}
					defer func() {
						<-s.busy
					}()
					j.f(j.param)
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
func (s *SenderResolver) Post(f func(interface{}), p interface{}) {
	s.jobs <- &job{f: f, param: p}
}

// ResolveSenders resolves sender accounts from given transactions
// concurrently using SenderResolver worker pool.
func (pool *TxPool) ResolveSenders(signer types.Signer, txs []*types.Transaction) {
	ot := time.Now()
	s := pool.senderResolver
	var total, by_ecrecover, failed int64 = int64(len(txs)), 0, 0

	var wg sync.WaitGroup
	for _, tx := range txs {
		hash := tx.Hash()
		if addr := types.GetSender(signer, tx); addr != nil {
			s.tx2addr.Put(hash, *addr)
			continue
		}

		data := s.tx2addr.Get(hash)
		if data != nil {
			types.SetSender(signer, tx, data.(common.Address))
			continue
		}

		wg.Add(1)
		atomic.AddInt64(&by_ecrecover, 1)
		s.Post(func(param interface{}) {
			t := param.(*types.Transaction)
			if from, err := types.Sender(signer, t); err == nil {
				s.tx2addr.Put(t.Hash(), from)
			} else {
				atomic.AddInt64(&failed, 1)
			}
			wg.Done()
		}, tx)
	}

	wg.Wait()

	if false && total > 1 {
		dt := float64(time.Now().Sub(ot) / time.Millisecond)
		if dt <= 0 {
			dt = 1
		}
		ps := float64(total) * 1000.0 / dt
		fmt.Printf("=== %d/%d/%d : took %.3f ms %.3f/sec %d\n", total, total-by_ecrecover, failed, dt, ps, s.tx2addr.Count())
	}

	return
}

// ResolveSender resolves sender address from a transaction
func (pool *TxPool) ResolveSender(signer types.Signer, tx *types.Transaction) {
	var txs []*types.Transaction
	txs = append(txs, tx)
	pool.ResolveSenders(signer, txs)
}

// EOF
