// Copyright 2019 The go-ethereum Authors
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
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/dbtest"
)

type EphemeralRDB struct {
	rdb  *RDBDatabase
	file string
}

func newEphemeralRDB(file string, cache int, handles int, namespace string, readonly bool) (*EphemeralRDB, error) {
	rdb, err := New(file, cache, handles, namespace, readonly)
	if err != nil {
		return nil, err
	}
	return &EphemeralRDB{
		rdb:  rdb,
		file: file,
	}, nil
}

func (db *EphemeralRDB) Path() string {
	return db.rdb.Path()
}

func (db *EphemeralRDB) Put(key []byte, value []byte) error {
	return db.rdb.Put(key, value)
}

func (db *EphemeralRDB) Has(key []byte) (bool, error) {
	return db.rdb.Has(key)
}

func (db *EphemeralRDB) Get(key []byte) ([]byte, error) {
	return db.rdb.Get(key)
}

func (db *EphemeralRDB) Delete(key []byte) error {
	return db.rdb.Delete(key)
}

func (db *EphemeralRDB) NewIterator(prefix, start []byte) ethdb.Iterator {
	return db.rdb.NewIterator(prefix, start)
}

func (db *EphemeralRDB) NewIteratorWithStart(start []byte) ethdb.Iterator {
	return db.rdb.NewIteratorWithStart(start)
}

func (db *EphemeralRDB) NewIteratorWithPrefix(prefix []byte) ethdb.Iterator {
	return db.rdb.NewIteratorWithPrefix(prefix)
}

func (db *EphemeralRDB) Close() error {
	err := db.rdb.Close()
	if len(db.file) > 0 {
		os.RemoveAll(db.file)
	}
	return err
}

func (db *EphemeralRDB) Stat(property string) (string, error) {
	return db.rdb.Stat(property)
}

func (db *EphemeralRDB) Compact(start []byte, limit []byte) error {
	return db.rdb.Compact(start, limit)
}

func (db *EphemeralRDB) Meter(prefix string) {
	return
}

func (db *EphemeralRDB) NewBatch() ethdb.Batch {
	return db.rdb.NewBatch()
}

func (db *EphemeralRDB) NewBatchWithSize(size int) ethdb.Batch {
	return db.rdb.NewBatchWithSize(size)
}

func (db *EphemeralRDB) NewSnapshot() (ethdb.Snapshot, error) {
	return db.rdb.NewSnapshot()
}

func TestRocksDB(t *testing.T) {
	t.Run("DatabaseSuite", func(t *testing.T) {
		dbtest.TestDatabaseSuite(t, func() ethdb.KeyValueStore {
			db, err := newEphemeralRDB("test", 1024, 1024, "test", false)
			if err != nil {
				t.Fatal(err)
			}
			return db
		})
	})
}
