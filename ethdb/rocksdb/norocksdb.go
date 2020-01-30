// rocksdb_database.go
// +build !rocksdb

package rocksdb

import "github.com/ethereum/go-ethereum/ethdb/leveldb"

func New(file string, cache int, handles int, namespace string) (*leveldb.Database, error) {
	return leveldb.New(file, cache, handles, namespace)
}
