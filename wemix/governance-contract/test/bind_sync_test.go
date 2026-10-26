package test

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	compile "github.com/ethereum/go-ethereum/wemix/governance-contract"
)

const (
	contractsRoot = "../contracts"
	bindDir       = "../../bind"
	regenHint     = "regenerate with: cd wemix/governance-contract/contracts && go run ."
)

// solcMetadata matches the CBOR trailer solc 0.8.x appends to every contract:
// {"ipfs": <34-byte multihash>, "solc": <3-byte version>} followed by its 2-byte length (0x0033).
// The multihash covers the metadata JSON, which embeds source paths as the compiler saw them,
// so it differs between machines and working directories even when the code is identical.
var solcMetadata = regexp.MustCompile(`a264697066735822[0-9a-f]{68}(64736f6c6343[0-9a-f]{6}0033)`)

// maskedMetadataHash is what every multihash is replaced with before comparison.
// The solc version bytes are kept on purpose so a compiler change is still detected.
var maskedMetadataHash = "a264697066735822" + strings.Repeat("0", 68) + "${1}"

func maskMetadataHash(src string) string {
	return solcMetadata.ReplaceAllString(src, maskedMetadataHash)
}

// TestBindingsMatchSources fails when a governance .sol change was committed without
// regenerating wemix/bind/gen_*_abi.go, or when the committed bindings were edited by hand.
// Everything in the generated files must match except the metadata hash inside each Bin.
func TestBindingsMatchSources(t *testing.T) {
	contracts, err := compile.CompileBindings(contractsRoot)
	if err != nil {
		t.Fatalf("compile governance contracts: %v", err)
	}

	for _, b := range compile.Bindings {
		b := b
		t.Run(b.File, func(t *testing.T) {
			want, err := contracts.Bind(compile.BindPackage, b.Contracts...)
			if err != nil {
				t.Fatalf("render binding: %v", err)
			}
			got, err := os.ReadFile(filepath.Join(bindDir, b.File))
			if err != nil {
				t.Fatalf("read committed binding: %v", err)
			}

			wantMasked, gotMasked := maskMetadataHash(want), maskMetadataHash(string(got))
			if n := len(solcMetadata.FindAllStringIndex(wantMasked, -1)); n != len(b.Contracts) {
				t.Fatalf("expected %d solc metadata trailers in rendered binding, found %d", len(b.Contracts), n)
			}
			if wantMasked == gotMasked {
				return
			}
			t.Errorf("%s is out of date with the Solidity sources (%s)\n%s", b.File, regenHint, describeDiff(gotMasked, wantMasked))
		})
	}
}

// describeDiff lists the first lines that differ, truncating long hex lines.
func describeDiff(committed, compiled string) string {
	const (
		maxLine  = 160
		maxDiffs = 10
	)
	// window shows maxLine characters around the first differing column.
	window := func(l string, col int) string {
		start := col - maxLine/4
		if start < 0 {
			start = 0
		}
		if start > len(l) {
			start = len(l)
		}
		end := start + maxLine
		if end > len(l) {
			end = len(l)
		}
		prefix := ""
		if start > 0 {
			prefix = "..."
		}
		return prefix + l[start:end]
	}
	a, b := strings.Split(committed, "\n"), strings.Split(compiled, "\n")
	var out strings.Builder
	if len(a) != len(b) {
		fmt.Fprintf(&out, "line count: committed=%d compiled=%d\n", len(a), len(b))
	}
	diffs := 0
	for i := 0; i < len(a) && i < len(b) && diffs < maxDiffs; i++ {
		if a[i] == b[i] {
			continue
		}
		diffs++
		col := 0
		for col < len(a[i]) && col < len(b[i]) && a[i][col] == b[i][col] {
			col++
		}
		fmt.Fprintf(&out, "line %d, first difference at column %d (len committed=%d compiled=%d):\n  committed: %s\n  compiled:  %s\n",
			i+1, col+1, len(a[i]), len(b[i]), window(a[i], col), window(b[i], col))
	}
	return out.String()
}
