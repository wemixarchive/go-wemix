// spinlock.go

package wemix

import (
	"runtime"
	"sync/atomic"
)

type SpinLock struct {
	v uint32
}

func (l *SpinLock) Lock() {
	for !l.TryLock() {
		runtime.Gosched()
	}
}

func (l *SpinLock) Unlock() {
	atomic.StoreUint32(&l.v, 0)
}

func (l *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapUint32(&l.v, 0, 1)
}

// EOF
