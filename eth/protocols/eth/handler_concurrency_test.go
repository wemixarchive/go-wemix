// Reproduction for the unbounded SPoA handler goroutines.
//
// Under PoA every request handler ran as `go f()` with no bound, no
// backpressure, and the error discarded, so one peer sending a stream of
// requests piled up handler goroutines without limit. All handlers now go
// through peer.dispatchHandler, which holds a per-peer semaphore. These tests
// drive that method directly: the first asserts concurrency never exceeds the
// bound, the second asserts a terminating peer starts no new handler.
package eth

import (
	"sync"
	"testing"
	"time"
)

func TestDispatchHandlerBoundsConcurrency(t *testing.T) {
	const floods = maxHandlerConcurrency + 5
	p := &Peer{
		handlerSem: make(chan struct{}, maxHandlerConcurrency),
		term:       make(chan struct{}),
	}

	var mu sync.Mutex
	cur, peak := 0, 0
	entered := make(chan struct{}, floods)
	release := make(chan struct{})
	f := func() error {
		mu.Lock()
		cur++
		if cur > peak {
			peak = cur
		}
		mu.Unlock()
		entered <- struct{}{}
		<-release
		mu.Lock()
		cur--
		mu.Unlock()
		return nil
	}

	var wg sync.WaitGroup
	for i := 0; i < floods; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = p.dispatchHandler(f)
		}()
	}

	// Exactly maxHandlerConcurrency handlers should run; the rest block in
	// dispatchHandler waiting for a slot.
	for i := 0; i < maxHandlerConcurrency; i++ {
		select {
		case <-entered:
		case <-time.After(5 * time.Second):
			t.Fatalf("only %d handlers started; expected %d", i, maxHandlerConcurrency)
		}
	}
	select {
	case <-entered:
		t.Fatalf("more than %d handlers ran concurrently; not bounded", maxHandlerConcurrency)
	case <-time.After(300 * time.Millisecond):
	}

	mu.Lock()
	gotPeak := peak
	mu.Unlock()
	if gotPeak != maxHandlerConcurrency {
		t.Fatalf("peak concurrency = %d, want %d", gotPeak, maxHandlerConcurrency)
	}

	close(release)
	wg.Wait()
}

func TestDispatchHandlerStopsOnTerm(t *testing.T) {
	p := &Peer{
		handlerSem: make(chan struct{}, 1),
		term:       make(chan struct{}),
	}
	p.handlerSem <- struct{}{} // occupy the only slot

	close(p.term) // peer is terminating

	ran := make(chan struct{}, 1)
	if err := p.dispatchHandler(func() error { ran <- struct{}{}; return nil }); err != nil {
		t.Fatalf("dispatchHandler returned error: %v", err)
	}
	select {
	case <-ran:
		t.Fatal("handler ran after peer termination; it should have been dropped")
	case <-time.After(200 * time.Millisecond):
	}
}
