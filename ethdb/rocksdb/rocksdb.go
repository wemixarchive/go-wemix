// rocksdb.go
//go:build rocksdb
// +build rocksdb

package rocksdb

// #include <stdlib.h>
// #include <string.h>
// #include "rocksdb/c.h"
import "C"

import (
	"errors"
	"runtime"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/metrics"
)

// Wemix: db stats
// (reads, read bytes, writes, written bytes, lookups, deletes)
var (
	_stats_enabled                                             = false
	_r_count, _r_bytes, _w_count, _w_bytes, _l_count, _d_count uint64

	errRocksdbNotFound = errors.New("not found")
)

type RDBDatabase struct {
	fn    string
	db    *C.rocksdb_t
	opts  *C.rocksdb_options_t
	wopts *C.rocksdb_writeoptions_t
	ropts *C.rocksdb_readoptions_t
}

type RDBIterator struct {
	it         *C.rocksdb_iterator_t
	opts       *C.rocksdb_readoptions_t
	first      bool
	lowerBound *C.char
	upperBound *C.char
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

func New(file string, cache int, handles int, namespace string, readonly bool) (*RDBDatabase, error) {
	var cerr *C.char

	// null terminated c string
	file0 := file + string(rune(0))

	opts := C.rocksdb_options_create()
	C.rocksdb_options_set_create_if_missing(opts, 1)
	C.rocksdb_options_set_max_open_files(opts, C.int(handles))

	wopts := C.rocksdb_writeoptions_create()
	ropts := C.rocksdb_readoptions_create()

	var db *C.rocksdb_t
	if readonly {
		db = C.rocksdb_open_for_read_only(opts, b2c([]byte(file0)), 0, &cerr)
	} else {
		db = C.rocksdb_open(opts, b2c([]byte(file0)), &cerr)
	}
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
		atomic.AddUint64(&_w_bytes, uint64(len(key)+len(value)))
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
		atomic.AddUint64(&_r_bytes, uint64(len(key))+uint64(C.int(cvl)))
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

func memdup(b []byte) *C.char {
	x := (*C.char)(C.malloc(C.size_t(len(b))))
	C.memcpy(unsafe.Pointer(x), unsafe.Pointer(b2c(b)), C.size_t(len(b)))
	return x
}

func (db *RDBDatabase) NewIterator(prefix, start []byte) ethdb.Iterator {
	var begin, end []byte
	begin = prefix
	if len(start) > 0 {
		begin = append(begin, start...)
	}
	if len(prefix) > 0 {
		end = incrBytes(prefix)
	}
	var lowerBound, upperBound *C.char
	opts := C.rocksdb_readoptions_create()
	if len(begin) > 0 {
		lowerBound = memdup(begin)
		C.rocksdb_readoptions_set_iterate_lower_bound(opts, lowerBound, C.size_t(len(begin)))
	}
	if len(end) > 0 {
		upperBound = memdup(end)
		C.rocksdb_readoptions_set_iterate_upper_bound(opts, upperBound, C.size_t(len(end)))
	}
	it := C.rocksdb_create_iterator(db.db, opts)
	C.rocksdb_iter_seek_to_first(it)
	rit := &RDBIterator{
		it:    it,
		opts:  opts,
		first: true,
	}
	if len(begin) > 0 {
		rit.lowerBound = lowerBound
		rit.upperBound = upperBound
	}
	return rit
}

func (db *RDBDatabase) NewIteratorWithStart(start []byte) ethdb.Iterator {
	opts := C.rocksdb_readoptions_create()
	lowerBound := memdup(start)
	C.rocksdb_readoptions_set_iterate_lower_bound(opts, lowerBound, C.size_t(len(start)))
	it := C.rocksdb_create_iterator(db.db, opts)
	C.rocksdb_iter_seek_to_first(it)
	return &RDBIterator{
		it:         it,
		opts:       opts,
		first:      true,
		lowerBound: lowerBound,
	}
}

func (db *RDBDatabase) NewIteratorWithPrefix(prefix []byte) ethdb.Iterator {
	start := prefix
	end := incrBytes(start)

	opts := C.rocksdb_readoptions_create()
	lowerBound := memdup(start)
	C.rocksdb_readoptions_set_iterate_lower_bound(opts, lowerBound, C.size_t(len(start)))
	upperBound := memdup(end)
	C.rocksdb_readoptions_set_iterate_upper_bound(opts, upperBound, C.size_t(len(end)))
	it := C.rocksdb_create_iterator(db.db, opts)
	C.rocksdb_iter_seek_to_first(it)
	return &RDBIterator{
		it:         it,
		opts:       opts,
		first:      true,
		lowerBound: lowerBound,
		upperBound: upperBound,
	}
}

// NewSnapshot creates a database snapshot based on the current state.
// The created snapshot will not be affected by all following mutations
// happened on the database.
func (db *RDBDatabase) NewSnapshot() (ethdb.Snapshot, error) {
	return newSnapshot(db), nil
}

func incrBytes(bz []byte) []byte {
	if len(bz) == 0 {
		return nil
	}
	cz := make([]byte, len(bz))
	copy(cz, bz)
	for i := len(bz) - 1; i >= 0; i-- {
		if cz[i] == 0xFF {
			cz[i] = 0
		} else {
			cz[i]++
			return cz
		}
	}
	return nil
}

func (it *RDBIterator) Next() bool {
	if it.first {
		it.first = false
	} else {
		// Added conditions to prevent Rocksdb Iterator error.
		// Valid() call is a RocksDB requirement.
		if C.rocksdb_iter_valid(it.it) == 0 {
			return false
		}
		C.rocksdb_iter_next(it.it)
	}
	return C.rocksdb_iter_valid(it.it) != 0
}

func (it *RDBIterator) Error() error {
	var cerr *C.char
	if it.it == nil {
		return nil
	}
	C.rocksdb_iter_get_error(it.it, &cerr)
	if cerr != nil {
		err := errors.New(C.GoString(cerr))
		C.rocksdb_free(unsafe.Pointer(cerr))
		return err
	}
	return nil
}

func (it *RDBIterator) Key() []byte {
	var cvl C.size_t
	cv := C.rocksdb_iter_key(it.it, &cvl)
	if cv == nil {
		return nil
	}
	return C.GoBytes(unsafe.Pointer(cv), C.int(cvl))
}

func (it *RDBIterator) Value() []byte {
	var cvl C.size_t
	cv := C.rocksdb_iter_value(it.it, &cvl)
	if cv == nil {
		return nil
	}
	return C.GoBytes(unsafe.Pointer(cv), C.int(cvl))
}

func (it *RDBIterator) Release() {
	if it.it != nil {
		C.rocksdb_iter_destroy(it.it)
	}
	if it.opts != nil {
		C.rocksdb_readoptions_destroy(it.opts)
	}
	if it.lowerBound != nil {
		C.free(unsafe.Pointer(it.lowerBound))
	}
	if it.upperBound != nil {
		C.free(unsafe.Pointer(it.upperBound))
	}
	it.it, it.opts, it.lowerBound, it.upperBound = nil, nil, nil, nil
}

func (db *RDBDatabase) Stat(property string) (string, error) {
	return "", errors.New("Not implemented")
}

func (db *RDBDatabase) Compact(start []byte, limit []byte) error {
	cs, cl := b2c(start), b2c(limit)
	C.rocksdb_compact_range(db.db, cs, C.size_t(len(start)), cl, C.size_t(len(limit)))
	return nil
}

func (db *RDBDatabase) Close() error {
	C.rocksdb_options_destroy(db.opts)
	C.rocksdb_writeoptions_destroy(db.wopts)
	C.rocksdb_readoptions_destroy(db.ropts)
	C.rocksdb_close(db.db)
	return nil
}

func (db *RDBDatabase) Meter(prefix string) {
	return
}

func rdbBatchFinalizer(b *rdbBatch) {
	if b.b != nil {
		bb := b.b
		b.b = nil
		go func() {
			// a little bit of delay here seems to abate crash
			time.Sleep(5 * time.Second)
			C.rocksdb_writebatch_destroy(bb)
		}()
	}
}

func (db *RDBDatabase) NewBatch() ethdb.Batch {
	b := C.rocksdb_writebatch_create()
	bb := &rdbBatch{db: db.db, b: b, wopts: db.wopts, data: nil}
	runtime.SetFinalizer(bb, rdbBatchFinalizer)
	return bb
}

func (db *RDBDatabase) NewBatchWithSize(size int) ethdb.Batch {
	b := C.rocksdb_writebatch_create()
	bb := &rdbBatch{db: db.db, b: b, wopts: db.wopts, data: nil, size: size}
	runtime.SetFinalizer(bb, rdbBatchFinalizer)
	return bb
}

type rdbBatchOp struct {
	del   bool
	key   []byte
	value []byte
}

type rdbBatch struct {
	db    *C.rocksdb_t
	b     *C.rocksdb_writebatch_t
	wopts *C.rocksdb_writeoptions_t
	data  []*rdbBatchOp
	size  int
}

func (b *rdbBatch) Put(key, value []byte) error {
	if _stats_enabled {
		atomic.AddUint64(&_w_count, 1)
		atomic.AddUint64(&_w_bytes, uint64(len(key)+len(value)))
	}
	ck, cv := b2c(key), b2c(value)
	C.rocksdb_writebatch_put(b.b, ck, C.size_t(len(key)), cv, C.size_t(len(value)))
	b.data = append(b.data, &rdbBatchOp{del: false, key: key, value: value})
	b.size += len(value)
	return nil
}

func (b *rdbBatch) Delete(key []byte) error {
	if _stats_enabled {
		atomic.AddUint64(&_d_count, 1)
	}
	C.rocksdb_writebatch_delete(b.b, b2c(key), C.size_t(len(key)))
	b.data = append(b.data, &rdbBatchOp{del: true, key: key, value: nil})
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
	C.rocksdb_writebatch_destroy(b.b)
	b.b = C.rocksdb_writebatch_create()
	b.data = nil
	b.size = 0
}

// Replay replays the batch contents.
func (b *rdbBatch) Replay(w ethdb.KeyValueWriter) error {
	for _, i := range b.data {
		if i.del {
			w.Delete(i.key)
		} else {
			w.Put(i.key, i.value)
		}
	}
	return nil
}

func EnableStats(b bool) {
	_stats_enabled = b
}

func Stats(device string) (disk_r_count, disk_r_bytes, disk_w_couhnt, disk_w_bytes, r_count, r_bytes, w_count, w_bytes, l_count, d_count uint64) {
	var diskStats metrics.DiskStats
	metrics.ReadProcDiskStats(device, &diskStats)
	return uint64(diskStats.ReadCount), uint64(diskStats.ReadBytes), uint64(diskStats.WriteCount), uint64(diskStats.WriteBytes), _r_count, _r_bytes, _w_count, _w_bytes, _l_count, _d_count
}

// snapshot wraps a batch of key-value entries deep copied from the in-memory
// database for implementing the Snapshot interface.
type snapshot struct {
	db    *C.rocksdb_t
	snap  *C.rocksdb_snapshot_t
	ropts *C.rocksdb_readoptions_t
}

// newSnapshot initializes the snapshot with the given database instance.
func newSnapshot(db *RDBDatabase) *snapshot {

	var ropts *C.rocksdb_readoptions_t
	var snap *C.rocksdb_snapshot_t

	ropts = C.rocksdb_readoptions_create()
	snap = C.rocksdb_create_snapshot(db.db)

	C.rocksdb_readoptions_set_snapshot(ropts, snap)

	return &snapshot{db: db.db, snap: snap, ropts: ropts}
}

// Has retrieves if a key is present in the snapshot backing by a key-value
// data store.
func (snap *snapshot) Has(key []byte) (bool, error) {
	if _stats_enabled {
		atomic.AddUint64(&_l_count, 1)
	}
	var cerr *C.char
	var cvl C.size_t
	ck := b2c(key)
	cv := C.rocksdb_get(snap.db, snap.ropts, ck, C.size_t(len(key)), &cvl, &cerr)
	if cerr != nil {
		return false, cerror(cerr)
	}
	if cv == nil {
		return false, errRocksdbNotFound
	}
	defer C.free(unsafe.Pointer(cv))
	return true, nil
}

// Get retrieves the given key if it's present in the snapshot backing by
// key-value data store.
func (snap *snapshot) Get(key []byte) ([]byte, error) {
	if _stats_enabled {
		atomic.AddUint64(&_r_count, 1)
	}
	var cerr *C.char
	var cvl C.size_t
	ck := b2c(key)
	cv := C.rocksdb_get(snap.db, snap.ropts, ck, C.size_t(len(key)), &cvl, &cerr)
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
		return nil, errRocksdbNotFound
	}
	if _stats_enabled {
		atomic.AddUint64(&_r_bytes, uint64(len(key))+uint64(C.int(cvl)))
	}
	defer C.free(unsafe.Pointer(cv))
	return C.GoBytes(unsafe.Pointer(cv), C.int(cvl)), nil
}

// Release releases associated resources. Release should always succeed and can
// be called multiple times without causing error.
func (snap *snapshot) Release() {

	C.rocksdb_release_snapshot(snap.db, snap.snap)

	snap.db = nil
	snap.snap = nil
	snap.ropts = nil
}
