// rocksdb_database.go
// +build !rocksdb

package ethdb

func NewRDBDatabase(file string, cache int, handles int) (Database, error) {
	return NewLDBDatabase(file, cache, handles)
}
