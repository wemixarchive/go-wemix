// tx_orderer.go

package miner

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TxOrdererList struct {
	addr       common.Address
	txs        types.Transactions
	ix         int
	sz         int
	prev, next *TxOrdererList
}

func (to *TxOrderer) listAppend(l *TxOrdererList) {
	if to.head == nil {
		to.head, to.tail, l.prev, l.next = l, l, nil, nil
	} else {
		l.prev, l.next = to.tail, nil
		to.tail.next, to.tail = l, l
	}
}

func (to *TxOrderer) listDelete(l *TxOrdererList) {
	if l.prev != nil {
		l.prev.next = l.next
	}
	if l.next != nil {
		l.next.prev = l.prev
	}
	if to.head == l {
		to.head = l.next
	}
	if to.tail == l {
		to.tail = l.prev
	}
}

type TxOrderer struct {
	lock             *sync.Mutex
	txs              []*types.Transaction
	head, tail, curr *TxOrdererList
	committedTxs     map[common.Hash]*types.Transaction
	done             bool // set when block generation is done, so that pre-fetching is not needed
	ix               int  // current index for real execution
	pix              int  // current index for pre-fetching
	sz               int
}

func NewTxOrderer(pending map[common.Address]types.Transactions, committedTxs map[common.Hash]*types.Transaction) *TxOrderer {
	to := &TxOrderer{
		lock:         &sync.Mutex{},
		committedTxs: committedTxs,
		done:         false,
		ix:           0,
		pix:          0,
	}
	for a, i := range pending {
		to.listAppend(&TxOrdererList{
			addr: a,
			txs:  i,
			ix:   0,
			sz:   len(i),
			prev: nil,
			next: nil,
		})
		to.sz += len(i)
	}
	to.curr = to.head
	return to
}

func (to *TxOrderer) Close() {
	to.done = true
}

// lock should be held by the caller
func (to *TxOrderer) pull(count int) {
	for count > 0 {
		if to.curr == nil {
			break
		}

		if to.curr.ix >= to.curr.sz {
			curr := to.curr.next
			to.listDelete(to.curr)
			to.curr = curr
			if to.curr == nil {
				to.curr = to.head
			}
			continue
		}

		to.txs = append(to.txs, to.curr.txs[to.curr.ix])
		to.curr.ix++

		to.curr = to.curr.next
		if to.curr == nil {
			to.curr = to.head
		}

		count--
	}
}

func (to *TxOrderer) Peek() *types.Transaction {
	to.lock.Lock()
	defer to.lock.Unlock()

	for {
		if to.ix >= len(to.txs) {
			to.pull(1000)
			if to.ix >= len(to.txs) {
				return nil
			}
		}

		tx := to.txs[to.ix]
		if _, ok := to.committedTxs[tx.Hash()]; !ok {
			return tx
		} else {
			to.ix++
			continue
		}
	}
}

func (to *TxOrderer) Shift() {
	to.lock.Lock()
	defer to.lock.Unlock()
	to.ix++
}

func (to *TxOrderer) Pop() {
	to.lock.Lock()
	defer to.lock.Unlock()
	to.ix++
}

func (to *TxOrderer) MarkCommitted(tx *types.Transaction) {
	to.lock.Lock()
	defer to.lock.Unlock()
	to.committedTxs[tx.Hash()] = tx
}

func (to *TxOrderer) CommittedLength() int {
	to.lock.Lock()
	defer to.lock.Unlock()
	return len(to.committedTxs)
}

func (to *TxOrderer) NextForPrefetch() *types.Transaction {
	to.lock.Lock()
	defer to.lock.Unlock()

	for {
		if to.done {
			return nil
		}
		if to.ix >= to.pix {
			to.pix = to.ix + 1
		} else if to.pix > to.ix+100 {
			to.lock.Unlock()
			time.Sleep(time.Millisecond * 20)
			to.lock.Lock()
			continue
		}
		if to.pix >= len(to.txs) {
			to.pull(1000)
			if to.pix >= len(to.txs) {
				return nil
			}
		}

		tx := to.txs[to.pix]
		to.pix++
		if _, ok := to.committedTxs[tx.Hash()]; !ok {
			return tx
		} else {
			continue
		}
	}
}

// EOF
