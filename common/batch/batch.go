// batch.go

package batch

import (
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

type Batch struct {
	toInterval time.Duration
	timeout    time.Duration
	batchCount int
	ch         chan interface{}
	count      int32
	f          func(interface{}, int) error
}

func NewBatch(toInterval, timeout time.Duration, batchCount int, f func(interface{}, int) error) *Batch {
	return &Batch{
		toInterval: toInterval,
		timeout:    timeout,
		batchCount: batchCount,
		ch:         make(chan interface{}, batchCount*10),
		count:      0,
		f:          f,
	}
}

func (b *Batch) Run() {
	var (
		data []interface{}
		lt   time.Time = time.Now() // last time
		ln   int       = 0          // last count
	)

	timer := time.NewTimer(0)
	<-timer.C // drain the initial timeout

	eod := false
	for {
		itstimer := false
		fire := false

		select {
		case d := <-b.ch:
			atomic.AddInt32(&b.count, -1)
			if d == nil {
				eod = true
			} else {
				data = append(data, d)
			}
		case <-timer.C:
			itstimer = true
		}
		last := b.count == 0

		if eod {
			break
		}

		// when to fire
		// 1. timer fired
		//   1.1 no count change
		//   1.2 more than 50 ms passed from the initial
		// 2. count >= 100

		if !itstimer {
			if ln == 0 {
				timer.Stop()
				timer.Reset(b.toInterval)
			} else if len(data) >= b.batchCount {
				fire = true
			}
		} else if last {
			et := time.Since(lt)
			if (len(data) == ln && et > b.toInterval) || et > b.timeout {
				fire = true
			}
		}

		if fire {
			if len(data) < b.batchCount {
				// do it
				e := b.f(data, len(data))
				if e != nil {
					log.Error("Server", "Failed", e)
				} else {
					log.Debug("Server", "Count", len(data))
				}
				data = nil
			} else {
				for {
					if len(data) < b.batchCount {
						break
					}

					// do it
					e := b.f(data, b.batchCount)
					if e != nil {
						log.Error("Server", "Failed", e)
					} else {
						log.Debug("Server", "Count", b.batchCount)
					}
					data = data[b.batchCount:]
				}
			}
		}

		lt = time.Now()
		ln = len(data)

		if itstimer && ln != 0 {
			timer.Reset(b.toInterval)
		} else if !itstimer && ln == 0 {
			timer.Stop()
		}
	}

	// got eod, flush the remaining data
	for len(data) > 0 {
		l := len(data)
		if l > b.batchCount {
			l = b.batchCount
		}
		e := b.f(data, l)
		if e != nil {
			log.Error("Server", "Failed", e)
		} else {
			log.Debug("Server", "Count", l)
		}
		data = data[l:]
	}
}

func (b *Batch) Stop() {
	b.ch <- nil
}

func (b *Batch) Put(data interface{}) {
	b.ch <- data
	atomic.AddInt32(&b.count, 1)
}

// EOF
