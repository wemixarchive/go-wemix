// dbbench

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethdb"
)

var (
	big1 = big.NewInt(1)
)

func read(db ethdb.Database, prefix string, start, end int, verbose bool) error {
	for i := start; i <= end; i++ {
		ks := fmt.Sprintf("%s-%d", prefix, i)
		k := sha3.Sum256([]byte(ks))
		v, err := db.Get(k[:])
		if err != nil {
			fmt.Printf("Failed to read %s: %v\n", ks, err)
			return err
		}
		if verbose {
			fmt.Printf("%s(<-%s): %d %s\n", hex.EncodeToString(k[:]), string(ks), len(v), hex.EncodeToString(v))
		}
	}
	return nil
}

func genVal(key []byte, sz int) []byte {
	sz = (sz + 31) / 32
	v0 := sha3.Sum256(key)
	v := v0[:]
	bi := big.NewInt(0)
	bi.SetBytes(v)
	for i := 1; i < sz; i++ {
		bi.Add(bi, big1)
		bib := bi.Bytes()
		l := len(bib)
		if l < 32 {
			l = 32 - l
			bib = append(bib, v0[:l]...)
		}
		v = append(v, bib[:32]...)
	}
	return v
}

func write(db ethdb.Database, prefix string, start, end, batchCount, valueSize int) error {
	dbb := db.NewBatch()
	defer dbb.Reset()

	for i := start; i <= end; i++ {
		ks := fmt.Sprintf("%s-%d", prefix, i)
		k := sha3.Sum256([]byte(ks))
		dbb.Put(k[:], genVal(k[:], valueSize))

		if dbb.ValueSize() >= batchCount {
			err := dbb.Write()
			if err != nil {
				fmt.Printf("Failed to write: %v\n", err)
				return nil
			}
			dbb.Reset()
		}
	}

	if dbb.ValueSize() > 0 {
		err := dbb.Write()
		if err != nil {
			fmt.Printf("Failed to write: %v\n", err)
			return nil
		}
	}

	return nil
}

func stats(device string) []uint64 {
	disk_r_count, disk_r_bytes, disk_w_couhnt, disk_w_bytes, r_count, r_bytes, w_count, w_bytes, l_count, d_count := ethdb.Stats(device)
	return []uint64{disk_r_count, disk_r_bytes, disk_w_couhnt, disk_w_bytes, r_count, r_bytes, w_count, w_bytes, l_count, d_count}
}

func header() {
	fmt.Printf("@,OP,Prefix,Start,End,Elap(ms),R(#),R(KB),R(KB/s),W(#),W(KB),W(KB/s),DbR(#),DbR(KB),DbW(#),DbW(KB),Has(#),Del(#)\n")
}

func pre(device string) (time.Time, []uint64) {
	return time.Now(), stats(device)
}

func post(device, header string, ot time.Time, ss []uint64) {
	dur := uint64(time.Since(ot) / time.Millisecond)
	se := stats(device)
	for i := 0; i < len(se); i++ {
		se[i] = se[i] - ss[i]
	}

	fmt.Printf("%s,%d", header, dur)
	for i := 0; i < len(se); i++ {
		v := se[i]
		// 1: disk read bytes
		// 3: disk write bytes
		// 5: DB read bytes
		// 7: DB write bytes
		switch i {
		case 1:
			fallthrough
		case 3:
			fallthrough
		case 5:
			fallthrough
		case 7:
			v /= 1024
			r := v * 1000 / dur
			fmt.Printf(",%d,%d", v, r)
		default:
			fmt.Printf(",%d", v)
		}
	}
	fmt.Println()
}

func usage() {
	fmt.Printf(`Usage: dbbench <options>... <db-name>
	[<read>|<write> <prefix> <start> <end> [<batch> <value-size>]]

options:
-t rocksdb|leveldb:	choose between rocksdb and leveldb (leveldb).
-d <device-name>:	where to collect disk stats from
-s <value-size>:	value size in bytes (32 bytes).
-H:	no header
-v:	verbose
`)
}

func main() {
	var (
		which    string = "leveldb"
		device   string
		db       ethdb.Database
		err      error
		noHeader bool = false
		verbose  bool = false
	)

	var nargs []string
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-h":
			usage()
			os.Exit(1)
		case "-d":
			if i >= len(os.Args)-1 {
				usage()
				return
			}
			device = os.Args[i+1]
			i++
		case "-t":
			if i >= len(os.Args)-1 {
				usage()
				return
			}
			which = os.Args[i+1]
			i++
		case "-v":
			verbose = true
		case "-H":
			noHeader = true
		default:
			nargs = append(nargs, os.Args[i])
		}
	}

	if len(nargs) <= 0 {
		usage()
		return
	}

	ethdb.EnableStats(true)
	switch which {
	case "leveldb":
		db, err = ethdb.NewLDBDatabase(nargs[0], 1024, 1024)
		if err != nil {
			fmt.Printf("Cannot open DB %s: %v\n", nargs[0], err)
			return
		}
	case "rocksdb":
		db, err = ethdb.NewRDBDatabase(nargs[0], 1024, 1024)
		if err != nil {
			fmt.Printf("Cannot open DB %s: %v\n", nargs[0], err)
			return
		}
	default:
		usage()
		return
	}

	if len(nargs) > 1 {
		if nargs[1] == "read" && len(nargs) >= 5 {
			prefix := nargs[2]
			six, e1 := strconv.Atoi(nargs[3])
			eix, e2 := strconv.Atoi(nargs[4])
			if e1 != nil || e2 != nil {
				usage()
				return
			}
			if !noHeader {
				header()
			}
			ot, stats := pre(device)
			read(db, prefix, six, eix, verbose)
			post(device, fmt.Sprintf("@,read,%s,%d,%d", prefix, six, eix), ot, stats)
		} else if nargs[1] == "write" && len(nargs) >= 7 {
			prefix := nargs[2]
			six, e1 := strconv.Atoi(nargs[3])
			eix, e2 := strconv.Atoi(nargs[4])
			batch, e3 := strconv.Atoi(nargs[5])
			valueSize, e4 := strconv.Atoi(nargs[6])
			if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
				usage()
				return
			}
			if !noHeader {
				header()
			}
			ot, stats := pre(device)
			write(db, prefix, six, eix, batch, valueSize)
			post(device, fmt.Sprintf("@,write,%s,%d,%d", prefix, six, eix), ot, stats)
		} else {
			usage()
		}
	} else {
		// read commands from stdin

		if !noHeader {
			header()
		}

		scanner := bufio.NewScanner(os.Stdin)
		r := regexp.MustCompile(`\s+`)
		for scanner.Scan() {
			ls := r.Split(strings.TrimSpace(scanner.Text()), -1)

			if ls[0] == "read" && len(ls) >= 4 {
				prefix := ls[1]
				six, e1 := strconv.Atoi(ls[2])
				eix, e2 := strconv.Atoi(ls[3])
				if e1 != nil || e2 != nil {
					continue
				}

				ot, stats := pre(device)
				read(db, prefix, six, eix, verbose)
				post(device, fmt.Sprintf("@,read,%s,%d,%d", prefix, six, eix), ot, stats)
			} else if ls[0] == "write" && len(ls) >= 6 {
				prefix := ls[1]
				six, e1 := strconv.Atoi(ls[2])
				eix, e2 := strconv.Atoi(ls[3])
				batch, e3 := strconv.Atoi(ls[4])
				valueSize, e4 := strconv.Atoi(ls[5])
				if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
					continue
				}
				ot, stats := pre(device)
				write(db, prefix, six, eix, batch, valueSize)
				post(device, fmt.Sprintf("@,write,%s,%d,%d", prefix, six, eix), ot, stats)
			}
		}
	}
}

// EOF
