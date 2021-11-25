// tx_sender_resolver

package core

import (
	"sync"
	"sync/atomic"

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
	for {
		j, ok := <-s.jobs
		if !ok || j == nil {
			break
		}
		go func() {
			s.busy <- struct{}{}
			defer func() {
				<-s.busy
			}()
			j.f(j.param)
		}()
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
	s := pool.senderResolver
	var by_ecrecover, failed int64

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
}

// ResolveSender resolves sender address from a transaction
func (pool *TxPool) ResolveSender(signer types.Signer, tx *types.Transaction) {
	var txs []*types.Transaction
	txs = append(txs, tx)
	pool.ResolveSenders(signer, txs)
}

// EOF
