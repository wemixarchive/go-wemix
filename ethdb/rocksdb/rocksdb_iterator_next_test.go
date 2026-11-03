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
//go:build rocksdb
// +build rocksdb

package rocksdb

import (
	"testing"
)

// Tests that Next on an exhausted iterator returns false instead of advancing.
//
// RDBIterator.Next called C.rocksdb_iter_next without first checking validity.
// rocksdb's iter_next is undefined on an exhausted iterator and crashes the
// process, so a caller that keeps calling Next past the end (as the snapshot
// generator does via removeStorageBefore) took the node down. leveldb and
// memorydb return false instead.
func TestRDBIteratorNextPastEnd(t *testing.T) {
	db, err := newEphemeralRDB(t.TempDir(), 0, 0, "", false)
	if err != nil {
		t.Fatalf("open rocksdb: %v", err)
	}
	defer db.Close()

	if err := db.Put([]byte("k1"), []byte("v1")); err != nil {
		t.Fatalf("put: %v", err)
	}
	if err := db.Put([]byte("k2"), []byte("v2")); err != nil {
		t.Fatalf("put: %v", err)
	}

	it := db.NewIterator(nil, nil)
	defer it.Release()

	seen := 0
	for it.Next() {
		seen++
	}
	if seen != 2 {
		t.Fatalf("iterated %d keys, want 2", seen)
	}
	// The iterator is exhausted. Calling Next again must not advance an invalid
	// iterator; it must keep returning false.
	for i := 0; i < 3; i++ {
		if it.Next() {
			t.Fatalf("Next returned true after exhaustion on call %d", i)
		}
	}
	if err := it.Error(); err != nil {
		t.Fatalf("iterator error: %v", err)
	}
}
