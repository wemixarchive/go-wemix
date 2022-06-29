// Copyright 2015 The go-ethereum Authors
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

// Contains the Linux implementation of process disk IO counter retrieval.

package metrics

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ReadDiskStats retrieves the disk IO stats belonging to the current process.
func ReadDiskStats(stats *DiskStats) error {
	// Open the process disk IO counter file
	inf, err := os.Open(fmt.Sprintf("/proc/%d/io", os.Getpid()))
	if err != nil {
		return err
	}
	defer inf.Close()
	in := bufio.NewReader(inf)

	// Iterate over the IO counter, and extract what we need
	for {
		// Read the next line and split to key and value
		line, err := in.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		if err != nil {
			return err
		}

		// Update the counter based on the key
		switch key {
		case "syscr":
			stats.ReadCount = value
		case "syscw":
			stats.WriteCount = value
		case "rchar":
			stats.ReadBytes = value
		case "wchar":
			stats.WriteBytes = value
		}
	}
}

// ReadProcDiskStats retrieves the disk IO stats from /proc/diskstats
func ReadProcDiskStats(device string, stats *DiskStats) error {
	stats.ReadCount, stats.ReadBytes, stats.WriteCount, stats.WriteBytes = 0, 0, 0, 0

	// Open the process disk IO counter file
	inf, err := os.Open("/proc/diskstats")
	if err != nil {
		return err
	}
	defer inf.Close()
	in := bufio.NewReader(inf)
	r := regexp.MustCompile(`\s+`)

	// Iterate over the IO counter, and extract what we need
	for {
		// Read the next line and split to key and value
		line, err := in.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		parts := r.Split(strings.TrimSpace(line), -1)
		if len(parts) < 10 {
			continue
		}
		if len(device) > 0 && device != parts[2] {
			continue
		} else if len(device) == 0 && parts[1] != "0" {
			continue
		}

		r_count, e1 := strconv.ParseInt(parts[3], 10, 64)
		r_sectors, e2 := strconv.ParseInt(parts[5], 10, 64)
		w_count, e3 := strconv.ParseInt(parts[7], 10, 64)
		w_sectors, e4 := strconv.ParseInt(parts[9], 10, 64)
		if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
			continue
		}

		stats.ReadCount += r_count
		stats.ReadBytes += r_sectors * 512
		stats.WriteCount += w_count
		stats.WriteBytes += w_sectors * 512
	}

	return nil
}
