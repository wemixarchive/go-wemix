// dbbench

package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethdb"
)

var (
	big1 = big.NewInt(1)
)

func read(db ethdb.Database, prefix string, start, end, numThreads int, verbose bool) error {

	doRead := func(idx int) error {
		ks := fmt.Sprintf("%s-%d", prefix, idx)
		k := sha3.Sum256([]byte(ks))
		v, err := db.Get(k[:])
		if err != nil {
			fmt.Printf("Failed to read %s: %v\n", ks, err)
			return err
		}
		if verbose {
			fmt.Printf("%s(<-%s): %d %s\n", hex.EncodeToString(k[:]), string(ks), len(v), hex.EncodeToString(v))
		}
		return nil
	}

	ix := int64(start)
	var wg sync.WaitGroup
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(gid int) {
			defer wg.Done()

			for {
				jx := atomic.AddInt64(&ix, 1) - 1
				if jx > int64(end) {
					break
				}
				doRead(int(jx))
			}
		}(i)
	}

	wg.Wait()
	return nil
}

func sha_256(b []byte) []byte {
	x := sha256.Sum256(b)
	return x[:]
}

func genVal(key []byte, sz int) []byte {
	sz = (sz + 31) / 32
	v := sha_256(key)
	for i := 1; i < sz; i++ {
		v = append(v, sha_256(v)...)
	}
	return v
}

func write(db ethdb.Database, prefix string, start, end, numThreads, batchCount, valueSize int) error {

	flush := func(dbb ethdb.Batch) error {
		var err error
		if dbb.ValueSize() > 0 {
			err = dbb.Write()
		}
		dbb.Reset()
		return err
	}

	var dbbs []ethdb.Batch
	for i := 0; i < numThreads; i++ {
		dbbs = append(dbbs, db.NewBatch())
	}

	defer func() {
		for i := 0; i < numThreads; i++ {
			flush(dbbs[i])
		}
	}()

	doWrite := func(gid, idx int) error {
		dbb := dbbs[gid]

		ks := fmt.Sprintf("%s-%d", prefix, idx)
		k := sha3.Sum256([]byte(ks))
		dbb.Put(k[:], genVal(k[:], valueSize))

		if dbb.ValueSize() >= batchCount {
			flush(dbb)
		}
		return nil
	}

	ix := int64(start)
	var wg sync.WaitGroup
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(gid int) {
			defer wg.Done()

			for {
				jx := atomic.AddInt64(&ix, 1) - 1
				if jx > int64(end) {
					break
				}
				doWrite(gid, int(jx))
			}
		}(i)
	}

	wg.Wait()
	return nil
}

func stats(device string) []uint64 {
	disk_r_count, disk_r_bytes, disk_w_couhnt, disk_w_bytes, r_count, r_bytes, w_count, w_bytes, l_count, d_count := ethdb.Stats(device)
	return []uint64{disk_r_count, disk_r_bytes, disk_w_couhnt, disk_w_bytes, r_count, r_bytes, w_count, w_bytes, l_count, d_count}
}

func diskUsage(dbPath string) (int, error) {
	cmd := exec.Command("du", "-s", dbPath)
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	ls := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(string(out)), -1)
	du, err := strconv.Atoi(ls[0])
	if err != nil {
		return 0, err
	}
	return du, nil
}

func header() {
	fmt.Printf("@,OP,Prefix,Start,End,Time,Elap,DB(KB),R(#),R(KB),R(KB/s),W(#),W(KB),W(KB/s),DbR(#),DbR(KB),DbR(KB/s),DbW(#),DbW(KB),DbW(KB/s),Has(#),Del(#)\n")
}

func pre(device string) (time.Time, []uint64) {
	return time.Now(), stats(device)
}

func post(dbPath, device, header string, ot time.Time, ss []uint64) {
	dur := uint64(time.Since(ot) / time.Millisecond)
	se := stats(device)
	for i := 0; i < len(se); i++ {
		se[i] = se[i] - ss[i]
	}

	du, err := diskUsage(dbPath)
	if err != nil {
		fmt.Printf("Failed to get disk usage: %v\n", err)
	}

	fmt.Printf("%s,%d,%d,%d", header, ot.Unix(), dur/1000, du)
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
	fmt.Printf(`Usage: dbbench [<options>...] <db-name>
	[<read>|<write> <prefix> <start> <end> [<batch> <value-size>]]

options:
-H:	no header
-t rocksdb|leveldb:	choose between rocksdb and leveldb (leveldb).
-d <device-name>:	where to collect disk stats from ("")
-r <num-threads>:	number of read threads (1)
-w <num-threads>:	number of write threads (1)
-v:	verbose

It's going to use about 1035 file descriptors, so don't forget to set open file descriptor limit to 2048, .e.g "ulimit -n 2048".
`)
}

func main() {
	var (
		which        string = "leveldb"
		device       string
		dbPath       string
		db           ethdb.Database
		err          error
		noHeader     bool = false
		readThreads  int = 1
		writeThreads int = 1
		verbose      bool = false
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
		case "-r":
			fallthrough
		case "-w":
			if i >= len(os.Args)-1 {
				usage()
				return
			}
			if os.Args[i] == "-r" {
				readThreads, err = strconv.Atoi(os.Args[i+1])
			} else {
				writeThreads, err = strconv.Atoi(os.Args[i+1])
			}
			if err != nil {
				usage()
				return
			}
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

	dbPath = nargs[0]

	ethdb.EnableStats(true)
	switch which {
	case "leveldb":
		db, err = ethdb.NewLDBDatabase(dbPath, 1024, 1024)
		if err != nil {
			fmt.Printf("Cannot open DB %s: %v\n", dbPath, err)
			return
		}
	case "rocksdb":
		db, err = ethdb.NewRDBDatabase(dbPath, 1024, 1024)
		if err != nil {
			fmt.Printf("Cannot open DB %s: %v\n", dbPath, err)
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
			read(db, prefix, six, eix, readThreads, verbose)
			post(dbPath, device, fmt.Sprintf("@,read,%s,%d,%d", prefix, six, eix), ot, stats)
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
			write(db, prefix, six, eix, writeThreads, batch, valueSize)
			post(dbPath, device, fmt.Sprintf("@,write,%s,%d,%d", prefix, six, eix), ot, stats)
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
				read(db, prefix, six, eix, readThreads, verbose)
				post(dbPath, device, fmt.Sprintf("@,read,%s,%d,%d", prefix, six, eix), ot, stats)
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
				write(db, prefix, six, eix, writeThreads, batch, valueSize)
				post(dbPath, device, fmt.Sprintf("@,write,%s,%d,%d", prefix, six, eix), ot, stats)
			}
		}
	}
}

// EOF
