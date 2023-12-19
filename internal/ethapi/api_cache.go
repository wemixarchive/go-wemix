// Copyright 2023 The go-ethereum Authors
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

package ethapi

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
)

var marshaledBlockPrefix = []byte(".block.")
var marshaledReceiptsPrefix = []byte(".receipts.")

func apiCacheOpen(fn string) (ethdb.Database, error) {
	return rawdb.NewDB(fn, 0, 0, "", false)
}

func apiCacheClose(db ethdb.Database) {
	db.Close()
}

func apiCacheGetMarshaledBlock(db ethdb.Database, hash []byte) (map[string]interface{}, error) {
	key := append(marshaledBlockPrefix, hash...)
	data, err := db.Get(key)
	if err != nil {
		return nil, err
	}

	gzReader, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer gzReader.Close()

	decompressedData, err := io.ReadAll(gzReader)
	if err != nil {
		return nil, err
	}

	var fields map[string]interface{}
	if err = json.Unmarshal(decompressedData, &fields); err != nil {
		return nil, err
	}
	return fields, nil
}

func apiCacheHasMarshaledBlock(db ethdb.Database, hash []byte) (bool, error) {
	key := append(marshaledBlockPrefix, hash...)
	return db.Has(key)
}

func apiCachePutMarshaledBlock(db ethdb.Database, hash []byte, fields map[string]interface{}) error {
	data, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	gzWriter := gzip.NewWriter(&buf)
	if _, err = gzWriter.Write(data); err != nil {
		return err
	}
	if err = gzWriter.Close(); err != nil {
		return err
	}

	key := append(marshaledBlockPrefix, hash...)
	return db.Put(key, buf.Bytes())
}

func apiCacheDeleteMarshaledBlock(db ethdb.Database, hash []byte) error {
	key := append(marshaledBlockPrefix, hash...)
	return db.Delete(key)
}

// EoF
