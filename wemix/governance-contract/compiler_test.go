package compile

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/compiler"
)

func TestPick(t *testing.T) {
	set := compiledTy{"A": {Code: "0xa"}, "B": {Code: "0xb"}}
	tests := []struct {
		name    string
		names   []string
		want    []string
		wantErr bool
	}{
		{name: "subset", names: []string{"A"}, want: []string{"A"}},
		{name: "all", names: []string{"A", "B"}, want: []string{"A", "B"}},
		{name: "missing", names: []string{"A", "C"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := set.Pick(tt.names...)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("Pick(%v) succeeded, want error", tt.names)
				}
				return
			}
			if err != nil {
				t.Fatalf("Pick(%v): %v", tt.names, err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("Pick(%v) returned %d contracts, want %d", tt.names, len(got), len(tt.want))
			}
			for _, n := range tt.want {
				if got[n] != set[n] {
					t.Errorf("Pick(%v)[%s] is not the source contract", tt.names, n)
				}
			}
		})
	}
}

func TestMerge(t *testing.T) {
	a := &compiler.Contract{Code: "0xa"}
	b := &compiler.Contract{Code: "0xb"}
	tests := []struct {
		name    string
		sets    []compiledTy
		want    compiledTy
		wantErr bool
	}{
		{name: "none", sets: nil, want: compiledTy{}},
		{name: "disjoint", sets: []compiledTy{{"A": a}, {"B": b}}, want: compiledTy{"A": a, "B": b}},
		{name: "duplicate", sets: []compiledTy{{"A": a}, {"A": b}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Merge(tt.sets...)
			if tt.wantErr {
				if err == nil {
					t.Fatal("Merge succeeded, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("Merge: %v", err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("Merge returned %d contracts, want %d", len(got), len(tt.want))
			}
			for n, c := range tt.want {
				if got[n] != c {
					t.Errorf("Merge()[%s] = %v, want %v", n, got[n], c)
				}
			}
		})
	}
}

func TestCompileWithRejectsEmptyOpenZeppelin(t *testing.T) {
	tests := []OpenZeppelin{
		{},
		{Contracts: "openzeppelin/contracts"},
		{Upgradeable: "openzeppelin/contracts-upgradeable"},
	}
	for _, oz := range tests {
		if _, err := CompileWith("", oz); err == nil {
			t.Errorf("CompileWith(%+v) succeeded, want error", oz)
		}
	}
}
