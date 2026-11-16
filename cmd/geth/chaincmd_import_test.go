// Copyright 2026 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/params"
)

// importGenesisFormat is a minimal genesis whose extra data is filled in per
// test. The wemix admin reads the boot node id from it.
const importGenesisFormat = `{
	"alloc"      : {},
	"coinbase"   : "0x0000000000000000000000000000000000000000",
	"difficulty" : "0x20000",
	"extraData"  : "%s",
	"gasLimit"   : "0x2fefd8",
	"nonce"      : "0x0000000000001338",
	"mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
	"parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
	"timestamp"  : "0x00",
	"config"     : {}
}`

// bootNodeIDBytes is the length of the boot node id the wemix admin expects
// in the genesis extra data.
const bootNodeIDBytes = 64

// Tests that import runs under PoA only. Any other consensus method is
// refused with a message naming the required flag value, and PoA still needs
// a genesis the wemix admin can be set up from.
func TestImportRequiresPoA(t *testing.T) {
	var (
		noBootNode   = ""
		withBootNode = "0x" + strings.Repeat("11", bootNodeIDBytes)
	)
	tests := []struct {
		name      string
		consensus int
		extra     string
		wantErr   string // empty means import must succeed
	}{
		{
			name:      "pow is refused",
			consensus: params.ConsensusPoW,
			extra:     noBootNode,
			wantErr:   "import requires PoA (--consensusmethod 2), got --consensusmethod 1",
		},
		{
			name:      "poa without boot node id fails admin setup",
			consensus: params.ConsensusPoA,
			extra:     noBootNode,
			wantErr:   "Failed to set up wemix admin",
		},
		{
			name:      "poa imports",
			consensus: params.ConsensusPoA,
			extra:     withBootNode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datadir := t.TempDir()
			consensus := strconv.Itoa(tt.consensus)

			genesis := filepath.Join(datadir, "genesis.json")
			if err := os.WriteFile(genesis, []byte(fmt.Sprintf(importGenesisFormat, tt.extra)), 0600); err != nil {
				t.Fatalf("failed to write genesis file: %v", err)
			}
			init := runGeth(t, "--datadir", datadir, "--consensusmethod", consensus, "init", genesis)
			init.WaitExit()
			if status := init.ExitStatus(); status != 0 {
				t.Fatalf("init exit status = %d, want 0\n%s", status, init.StderrText())
			}

			// No blocks are needed: both checks run before the first one is read.
			blocks := filepath.Join(datadir, "blocks.rlp")
			if err := os.WriteFile(blocks, nil, 0600); err != nil {
				t.Fatalf("failed to write block file: %v", err)
			}
			imp := runGeth(t, "--datadir", datadir, "--consensusmethod", consensus, "import", blocks)
			imp.WaitExit()

			status, stderr := imp.ExitStatus(), imp.StderrText()
			if tt.wantErr == "" {
				if status != 0 {
					t.Fatalf("import exit status = %d, want 0\n%s", status, stderr)
				}
				return
			}
			if status == 0 {
				t.Fatalf("import succeeded, want failure with %q", tt.wantErr)
			}
			if !strings.Contains(stderr, tt.wantErr) {
				t.Fatalf("import stderr does not contain %q\n%s", tt.wantErr, stderr)
			}
		})
	}
}
