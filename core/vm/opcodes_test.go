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

package vm

import "testing"

// TestIsPush checks that IsPush covers PUSH0 through PUSH32 and nothing
// around them. Tracers read op.isPush() from it, so PUSH0 must count.
func TestIsPush(t *testing.T) {
	for op := 0; op < 256; op++ {
		want := op >= int(PUSH0) && op <= int(PUSH32)
		if got := OpCode(op).IsPush(); got != want {
			t.Errorf("%v (0x%02x): IsPush() = %v, want %v", OpCode(op), op, got, want)
		}
	}
}
