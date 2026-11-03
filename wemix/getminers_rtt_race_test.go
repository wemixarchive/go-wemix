// Reproduction for the getMiners RttMs data race.
//
// minerStatusFeed.Send delivers the same *WemixMinerStatus pointer to every
// subscriber. getMiners stamps RttMs on the pointer it receives, while the
// background handleMinerStatusUpdate goroutine clones the same pointer. The two
// run concurrently with no synchronization, so the write in getMiners races the
// read inside Clone.
//
// This test drives the two real operations on one pointer delivered by the real
// feed: rttFromFeed (what getMiners calls) and Clone (what
// handleMinerStatusUpdate calls). Run with -race. It fails (RED) while
// rttFromFeed mutates the shared value in place and passes (GREEN) once it
// copies first.
package wemix

import (
	"math/big"
	"sync"
	"testing"
	"time"

	wemixapi "github.com/ethereum/go-ethereum/wemix/api"
)

func TestGetMinersRttDoesNotRaceFeed(t *testing.T) {
	ch1 := make(chan *wemixapi.WemixMinerStatus, 1) // getMiners' subscription
	ch2 := make(chan *wemixapi.WemixMinerStatus, 1) // handleMinerStatusUpdate's subscription
	sub1 := wemixapi.SubscribeToMinerStatus(ch1)
	sub2 := wemixapi.SubscribeToMinerStatus(ch2)
	defer sub1.Unsubscribe()
	defer sub2.Unsubscribe()

	status := &wemixapi.WemixMinerStatus{
		NodeName:          "node1",
		RttMs:             big.NewInt(0),
		LatestBlockHeight: big.NewInt(1),
		LatestBlockTd:     big.NewInt(1),
	}
	wemixapi.GotStatusEx(status) // one pointer to both subscribers

	start := time.Now().UnixNano()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { // getMiners stamps RttMs on the delivered status
		defer wg.Done()
		s := <-ch1
		_ = rttFromFeed(s, start)
	}()
	go func() { // handleMinerStatusUpdate clones the same delivered status
		defer wg.Done()
		s := <-ch2
		_ = s.Clone()
	}()
	wg.Wait()
}

func TestRttFromFeedNilStatus(t *testing.T) {
	if got := rttFromFeed(nil, time.Now().UnixNano()); got != nil {
		t.Fatalf("rttFromFeed(nil) = %+v, want nil", got)
	}
}
