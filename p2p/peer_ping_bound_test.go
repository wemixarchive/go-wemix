// Regression test for inbound ping handling: serving inbound pings must not
// grow the goroutine count without bound.
//
// Earlier, handle() sent each pong from a freshly spawned goroutine, so the
// number in flight scaled with inbound ping volume. This test sends many
// inbound pings to a peer whose pong writes cannot drain and asserts the
// goroutine count stays bounded. It fails (RED) when each ping spawns a
// goroutine and passes (GREEN) once pongs are served by the single pingLoop
// through a bounded channel.
package p2p

import (
	"runtime"
	"testing"
	"time"
)

func TestInboundPingDoesNotSpawnUnboundedGoroutines(t *testing.T) {
	const pingCount = 3000
	closer, rw, _, _ := testPeer(nil)
	defer closer()

	// Measure the baseline first so the growth is attributable to the pings.
	base := runtime.NumGoroutine()

	// Do not read the pongs rw sends back, so a pong write blocks. Send the
	// pings from a sender goroutine (the writes themselves may block once the
	// peer stops reading).
	go func() {
		for i := 0; i < pingCount; i++ {
			if err := SendItems(rw, pingMsg); err != nil {
				return
			}
		}
	}()

	// Let the peer process the pings.
	deadline := time.After(3 * time.Second)
	peak := base
	for {
		select {
		case <-deadline:
			// A per-ping goroutine would leave roughly `pingCount` extra
			// goroutines alive; the bounded path leaves only a handful.
			if grew := peak - base; grew > 500 {
				t.Fatalf("goroutines grew by %d while answering pings; pong handling is not bounded", grew)
			}
			return
		case <-time.After(50 * time.Millisecond):
			if n := runtime.NumGoroutine(); n > peak {
				peak = n
			}
		}
	}
}
