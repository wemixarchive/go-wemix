// Copyright 2026 The go-ethereum Authors
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

package ethconfig

import "testing"

// wemixBlockGasLimit is the gasLimit of the Wemix mainnet and testnet genesis.
// Update it, along with defaultRPCGasCap, when governance changes blockGasLimit.
const wemixBlockGasLimit = 105000000

// TestDefaultRPCGasCap checks that the default RPC gas cap is bounded and does
// not fall below the block gas limit. An unbounded cap leaves the resources of a
// single call unbounded, and a cap below the block gas limit makes
// eth_estimateGas fail for transactions that would still fit in a block (see
// internal/ethapi.DoEstimateGas).
func TestDefaultRPCGasCap(t *testing.T) {
	if Defaults.RPCGasCap == 0 {
		t.Fatal("RPCGasCap default is 0 (unlimited), want a bounded cap")
	}
	if Defaults.RPCGasCap < wemixBlockGasLimit {
		t.Fatalf("RPCGasCap default = %d, below the block gas limit %d: eth_estimateGas fails for transactions that fit in a block",
			Defaults.RPCGasCap, wemixBlockGasLimit)
	}
}
