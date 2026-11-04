// Reproduction for the unbounded SPoA handler goroutines.
//
// Under PoA every request handler ran as `go f()` with no bound, no
// backpressure, and the error discarded, so one peer sending a stream of
// requests piled up handler goroutines without limit. All handlers now go
// through peer.dispatchHandler, which holds a per-peer semaphore and a
// node-global semaphore. These tests drive that method directly: concurrency
// never exceeds the per-peer bound, nor the global bound across many peers,
// both slots are released when a handler finishes, and handler errors are
// classified (connection teardown vs genuine fault) rather than dropping the peer.
package eth

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

func TestDispatchHandlerBoundsConcurrency(t *testing.T) {
	const floods = maxHandlerConcurrency + 5
	p := &Peer{handlerSem: make(chan struct{}, maxHandlerConcurrency)}

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

// TestDispatchHandlerBoundsGlobalConcurrency oversubscribes the node-global
// bound across many peers (each still capped per-peer) and asserts concurrency
// never exceeds maxGlobalHandlerConcurrency. No t.Parallel(): globalHandlerSem
// is a package global shared with the other dispatch tests.
func TestDispatchHandlerBoundsGlobalConcurrency(t *testing.T) {
	const peers = (maxGlobalHandlerConcurrency / maxHandlerConcurrency) + 4
	total := peers * maxHandlerConcurrency // > maxGlobalHandlerConcurrency

	var mu sync.Mutex
	cur, peak := 0, 0
	entered := make(chan struct{}, total)
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
	for p := 0; p < peers; p++ {
		peer := &Peer{handlerSem: make(chan struct{}, maxHandlerConcurrency)}
		for i := 0; i < maxHandlerConcurrency; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_ = peer.dispatchHandler(f)
			}()
		}
	}

	// Exactly maxGlobalHandlerConcurrency handlers should run; the rest block
	// in dispatchHandler waiting for a global slot.
	for i := 0; i < maxGlobalHandlerConcurrency; i++ {
		select {
		case <-entered:
		case <-time.After(5 * time.Second):
			t.Fatalf("only %d handlers started; expected %d", i, maxGlobalHandlerConcurrency)
		}
	}
	select {
	case <-entered:
		t.Fatalf("more than %d handlers ran concurrently; global bound not enforced", maxGlobalHandlerConcurrency)
	case <-time.After(300 * time.Millisecond):
	}
	mu.Lock()
	gotPeak := peak
	mu.Unlock()
	if gotPeak != maxGlobalHandlerConcurrency {
		t.Fatalf("peak concurrency = %d, want %d", gotPeak, maxGlobalHandlerConcurrency)
	}
	close(release)
	wg.Wait()
}

// TestDispatchHandlerReleasesBothSlots verifies that finishing a handler frees
// both its per-peer and its global slot, so later requests are not starved.
func TestDispatchHandlerReleasesBothSlots(t *testing.T) {
	peer := &Peer{handlerSem: make(chan struct{}, maxHandlerConcurrency)}
	ran := make(chan struct{})
	_ = peer.dispatchHandler(func() error { ran <- struct{}{}; return nil })
	<-ran // handler body reached; its deferred releases run right after it returns

	deadline := time.After(2 * time.Second)
	for {
		if len(peer.handlerSem) == 0 && len(globalHandlerSem) == 0 {
			return
		}
		select {
		case <-deadline:
			t.Fatalf("slots not released: perPeer=%d global=%d", len(peer.handlerSem), len(globalHandlerSem))
		case <-time.After(5 * time.Millisecond):
		}
	}
}

// TestIsConnClosedErr pins the error classification that decides the log level:
// connection/shutdown errors are expected (quiet), genuine faults are not.
func TestIsConnClosedErr(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"shutting down", p2p.ErrShuttingDown, true},
		{"wrapped shutting down", fmt.Errorf("reply failed: %w", p2p.ErrShuttingDown), true},
		{"eof", io.EOF, true},
		{"closed pipe", io.ErrClosedPipe, true},
		{"net closed", net.ErrClosed, true},
		{"use of closed string", errors.New("write tcp: use of closed network connection"), true},
		{"broken pipe", errors.New("write: broken pipe"), true},
		{"genuine fault", errors.New("rlp: encoding failed"), false},
	}
	for _, c := range cases {
		if got := isConnClosedErr(c.err); got != c.want {
			t.Errorf("%s: isConnClosedErr(%v) = %v, want %v", c.name, c.err, got, c.want)
		}
	}
}

// TestDispatchHandlerNoSlotLeak stresses acquire/release across many peers,
// oversubscribing both bounds with a mix of nil- and error-returning handlers,
// and asserts both the per-peer and the global semaphore drain completely. A
// slot that is acquired but not released would leave a non-zero count here;
// over time that would exhaust the global pool and stall every peer's read loop.
func TestDispatchHandlerNoSlotLeak(t *testing.T) {
	const peers = 6
	const perPeer = maxHandlerConcurrency * 3 // oversubscribe the per-peer bound

	ps := make([]*Peer, peers)
	for i := range ps {
		// A real embedded *p2p.Peer so the error path's logHandlerError -> Log()
		// works, as it does in production (NewPeer always sets it).
		ps[i] = &Peer{
			Peer:       p2p.NewPeer(enode.ID{byte(i)}, "leaktest", nil),
			handlerSem: make(chan struct{}, maxHandlerConcurrency),
		}
	}

	var wg sync.WaitGroup
	for pi := range ps {
		peer := ps[pi]
		for i := 0; i < perPeer; i++ {
			wg.Add(1)
			i := i
			go func() {
				defer wg.Done()
				_ = peer.dispatchHandler(func() error {
					// Both the nil and the error path must release both slots.
					if i%3 == 0 {
						return errors.New("boom")
					}
					return nil
				})
			}()
		}
	}
	wg.Wait() // every dispatchHandler call returned (handler goroutine spawned)

	// Wait for all handler goroutines to run their deferred releases, then
	// assert every semaphore is empty — no slot leaked.
	deadline := time.After(5 * time.Second)
	for {
		leaked := len(globalHandlerSem)
		for _, p := range ps {
			leaked += len(p.handlerSem)
		}
		if leaked == 0 {
			return
		}
		select {
		case <-deadline:
			perPeerHeld := make([]int, len(ps))
			for i, p := range ps {
				perPeerHeld[i] = len(p.handlerSem)
			}
			t.Fatalf("slots leaked: global=%d per-peer=%v", len(globalHandlerSem), perPeerHeld)
		case <-time.After(5 * time.Millisecond):
		}
	}
}
