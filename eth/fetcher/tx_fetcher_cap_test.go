// Copyright 2026 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package fetcher

import (
	"encoding/binary"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// uniqueAnnounceHashes returns n distinct hashes tagged with prefix, so that
// the counts below are not reduced by duplicates collapsing in the maps.
func uniqueAnnounceHashes(prefix byte, n int) []common.Hash {
	hashes := make([]common.Hash, n)
	for i := range hashes {
		hashes[i][0] = prefix
		binary.BigEndian.PutUint32(hashes[i][1:], uint32(i))
	}
	return hashes
}

// trackedAnnounces counts the hashes a peer holds in the waitlist and the
// scheduler, the two maps the per-peer announcement cap bounds.
func trackedAnnounces(f *TxFetcher, peer string) int {
	return len(f.waitslots[peer]) + len(f.announces[peer])
}

// Tests that an announcement crossing the per-peer cap keeps exactly the room
// left under the cap: no more and no less, so the per-peer count stops at the
// cap for both a nearly-full peer and a single bulk announcement.
func TestTransactionFetcherAnnounceCapTruncation(t *testing.T) {
	tests := []struct {
		name  string
		used  int // hashes the peer already holds
		extra int // hashes in the announcement that crosses the cap
	}{
		// Already near the cap, then a small announcement crosses it.
		{name: "overshoot", used: maxTxAnnounces - 10, extra: 100},
		// Empty peer, then one announcement larger than the cap.
		{name: "undershoot", used: 0, extra: maxTxAnnounces + 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			first := uniqueAnnounceHashes(0x01, tt.used)
			second := uniqueAnnounceHashes(0x02, tt.extra)

			var fetcher *TxFetcher
			var steps []interface{}
			if len(first) > 0 {
				steps = append(steps, doTxNotify{peer: "A", hashes: first})
			}
			steps = append(steps,
				doTxNotify{peer: "A", hashes: second},
				doFunc(func() {
					if got := trackedAnnounces(fetcher, "A"); got != maxTxAnnounces {
						t.Errorf("tracked announcements = %d, want %d (cap)", got, maxTxAnnounces)
					}
				}))
			testTransactionFetcherParallel(t, txFetcherTest{
				init: func() *TxFetcher {
					fetcher = NewTxFetcher(
						func(common.Hash) bool { return false },
						nil,
						func(string, []common.Hash) error { return nil },
					)
					return fetcher
				},
				steps: steps,
			})
		})
	}
}
