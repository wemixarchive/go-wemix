// rocksdb_database.go
// +build rocksdb

package ethdb

// #include <stdlib.h>
// #include "rocksdb/c.h"
import "C"

import (
	"errors"
	"runtime"
	"sync/atomic"
	"unsafe"
)

type RDBDatabase struct {
	fn    string
	db    *C.rocksdb_t
	opts  *C.rocksdb_options_t
	wopts *C.rocksdb_writeoptions_t
	ropts *C.rocksdb_readoptions_t
}

func cerror(cerr *C.char) error {
	if cerr == nil {
		return nil
	}
	err := errors.New(C.GoString(cerr))
	C.free(unsafe.Pointer(cerr))
	return err
}

func b2c(b []byte) *C.char {
	if len(b) == 0 {
		return nil
	} else {
		return (*C.char)(unsafe.Pointer(&b[0]))
	}
}

func NewRDBDatabase(file string, cache int, handles int) (*RDBDatabase, error) {
	var cerr *C.char

	opts := C.rocksdb_options_create()
	C.rocksdb_options_set_create_if_missing(opts, 1)

	wopts := C.rocksdb_writeoptions_create()
	ropts := C.rocksdb_readoptions_create()

	db := C.rocksdb_open(opts, b2c([]byte(file)), &cerr)
	if cerr != nil {
		C.rocksdb_options_destroy(opts)
		return nil, cerror(cerr)
	}
	return &RDBDatabase{
		fn:    file,
		db:    db,
		opts:  opts,
		wopts: wopts,
		ropts: ropts,
	}, nil
}

func (db *RDBDatabase) Path() string {
	return db.fn
}

func (db *RDBDatabase) Put(key []byte, value []byte) error {
	if _stats_enabled {
		atomic.AddUint64(&_w_count, 1)
		atomic.AddUint64(&_w_bytes, uint64(len(key) + len(value)))
	}
	var cerr *C.char
	ck, cv := b2c(key), b2c(value)
	C.rocksdb_put(db.db, db.wopts, ck, C.size_t(len(key)), cv, C.size_t(len(value)),
		&cerr)
	if cerr != nil {
		return cerror(cerr)
	}
	return nil
}

func (db *RDBDatabase) Has(key []byte) (bool, error) {
	if _stats_enabled {
		atomic.AddUint64(&_l_count, 1)
	}
	var cerr *C.char
	var cvl C.size_t
	ck := b2c(key)
	cv := C.rocksdb_get(db.db, db.ropts, ck, C.size_t(len(key)), &cvl, &cerr)
	if cerr != nil {
		return false, cerror(cerr)
	}
	if cv == nil {
		return false, nil
	}
	defer C.free(unsafe.Pointer(cv))
	return true, nil
}

func (db *RDBDatabase) Get(key []byte) ([]byte, error) {
	if _stats_enabled {
		atomic.AddUint64(&_r_count, 1)
	}
	var cerr *C.char
	var cvl C.size_t
	ck := b2c(key)
	cv := C.rocksdb_get(db.db, db.ropts, ck, C.size_t(len(key)), &cvl, &cerr)
	if cerr != nil {
		if _stats_enabled {
			atomic.AddUint64(&_r_bytes, uint64(len(key)))
		}
		return nil, cerror(cerr)
	}
	if cv == nil {
		if _stats_enabled {
			atomic.AddUint64(&_r_bytes, uint64(len(key)))
		}
		return nil, nil
	}
	if _stats_enabled {
		atomic.AddUint64(&_r_bytes, uint64(len(key)) + uint64(C.int(cvl)))
	}
	defer C.free(unsafe.Pointer(cv))
	return C.GoBytes(unsafe.Pointer(cv), C.int(cvl)), nil
}

func (db *RDBDatabase) Delete(key []byte) error {
	if _stats_enabled {
		atomic.AddUint64(&_d_count, 1)
	}
	var cerr *C.char
	ck := b2c(key)
	C.rocksdb_delete(db.db, db.wopts, ck, C.size_t(len(key)), &cerr)
	if cerr != nil {
		return cerror(cerr)
	}
	return nil
}

/*
func (db *RDBDatabase) NewIterator() iterator.Iterator {
	return nil
}

func (db *RDBDatabase) NewIteratorWithPrefix(prefix []byte) iterator.Iterator {
	return nil
}
*/

func (db *RDBDatabase) Close() {
	C.rocksdb_options_destroy(db.opts)
	C.rocksdb_writeoptions_destroy(db.wopts)
	C.rocksdb_readoptions_destroy(db.ropts)
	C.rocksdb_close(db.db)
}

func (db *RDBDatabase) Meter(prefix string) {
	return
}

func (db *RDBDatabase) NewBatch() Batch {
	b := C.rocksdb_writebatch_create()
	bb := &rdbBatch{db: db.db, b: b, wopts: db.wopts}
	runtime.SetFinalizer(bb, func(bb *rdbBatch) {
		if bb.b != nil {
			C.rocksdb_writebatch_destroy(bb.b)
			bb.b = nil
		}
	})
	return bb
}

type rdbBatch struct {
	db    *C.rocksdb_t
	b     *C.rocksdb_writebatch_t
	wopts *C.rocksdb_writeoptions_t
	size  int
}

func (b *rdbBatch) Put(key, value []byte) error {
	if _stats_enabled {
		atomic.AddUint64(&_w_count, 1)
		atomic.AddUint64(&_w_bytes, uint64(len(key) + len(value)))
	}
	ck, cv := b2c(key), b2c(value)
	C.rocksdb_writebatch_put(b.b, ck, C.size_t(len(key)), cv, C.size_t(len(value)))
	b.size += len(value)
	return nil
}

func (b *rdbBatch) Delete(key []byte) error {
	if _stats_enabled {
		atomic.AddUint64(&_d_count, 1)
	}
	C.rocksdb_writebatch_delete(b.b, b2c(key), C.size_t(len(key)))
	b.size += 1
	return nil
}

func (b *rdbBatch) Write() error {
	var cerr *C.char
	C.rocksdb_write(b.db, b.wopts, b.b, &cerr)
	return cerror(cerr)
}

func (b *rdbBatch) ValueSize() int {
	return b.size
}

func (b *rdbBatch) Reset() {
	C.rocksdb_writebatch_clear(b.b)
	b.size = 0
}
