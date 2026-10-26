package test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gov "github.com/ethereum/go-ethereum/wemix/bind"
	compile "github.com/ethereum/go-ethereum/wemix/governance-contract"
)

const (
	contractsRoot = "../contracts"
	regenHint     = "regenerate with: cd wemix/governance-contract/contracts && go run ."
)

// solc 0.8.x appends a CBOR metadata trailer to the runtime code of every contract:
//
//	a2 64 "ipfs" 58 22 <34-byte multihash> 64 "solc" 43 <3-byte version> 00 33
//
// The multihash covers the metadata JSON, which records source paths as the compiler
// saw them, so it changes with the build directory even when the code is identical.
// The trailer is not executed, so it is removed from both sides before comparing.
var (
	metadataHead = []byte{0xa2, 0x64, 'i', 'p', 'f', 's', 0x58, 0x22}
	metadataMid  = []byte{0x64, 's', 'o', 'l', 'c', 0x43}
	metadataTail = []byte{0x00, 0x33}
)

const (
	multihashLen   = 34
	solcVersionLen = 3
	metadataLen    = 8 + multihashLen + 6 + solcVersionLen + 2 // 53 bytes
)

// boundContracts maps each governance contract to the metadata compiled into gwemix.
var boundContracts = map[string]*bind.MetaData{
	gov.REGISTRY:          gov.RegistryMetaData,
	gov.GOV:               gov.GovMetaData,
	gov.GOV_IMP:           gov.GovImpMetaData,
	gov.NCPEXIT:           gov.NCPExitMetaData,
	gov.NCPEXIT_IMP:       gov.NCPExitImpMetaData,
	gov.STAKING:           gov.StakingMetaData,
	gov.STAKING_IMP:       gov.StakingImpMetaData,
	gov.BALLOTSTORAGE:     gov.BallotStorageMetaData,
	gov.BALLOTSTORAGE_IMP: gov.BallotStorageImpMetaData,
	gov.ENVSTORAGE:        gov.EnvStorageMetaData,
	gov.ENVSTORAGE_IMP:    gov.EnvStorageImpMetaData,
}

// stripMetadata removes every solc metadata trailer from code and returns how many it removed.
// The trailer is located by its byte pattern rather than by trimming the tail, because in
// proxy contracts constructor-only data follows the runtime code inside the creation code.
func stripMetadata(code []byte) ([]byte, int) {
	var (
		out   = make([]byte, 0, len(code))
		found int
	)
	for {
		i := bytes.Index(code, metadataHead)
		if i < 0 || len(code)-i < metadataLen {
			return append(out, code...), found
		}
		trailer := code[i : i+metadataLen]
		mid := len(metadataHead) + multihashLen
		if !bytes.Equal(trailer[mid:mid+len(metadataMid)], metadataMid) ||
			!bytes.Equal(trailer[metadataLen-len(metadataTail):], metadataTail) {
			out = append(out, code[:i+1]...)
			code = code[i+1:]
			continue
		}
		out = append(out, code[:i]...)
		code = code[i+metadataLen:]
		found++
	}
}

func decodeBin(s string) ([]byte, error) {
	return hex.DecodeString(strings.TrimPrefix(s, "0x"))
}

// TestBindingsMatchSources compiles the governance Solidity sources and checks that the
// bytecode of every contract equals the Bin in wemix/bind/gen_*_abi.go, ignoring only the
// solc metadata trailer. It fails when a .sol change is committed without regenerating the
// bindings, or when a binding's bytecode is edited by hand.
// GovImp is compared with the deployed GovImpLegacy build, as CompileBindings produces it.
func TestBindingsMatchSources(t *testing.T) {
	compiled, err := compile.CompileBindings(contractsRoot)
	if err != nil {
		t.Fatalf("compile governance contracts: %v", err)
	}

	for _, b := range compile.Bindings {
		for _, name := range b.Contracts {
			name, file := name, b.File
			t.Run(name, func(t *testing.T) {
				meta, ok := boundContracts[name]
				if !ok {
					t.Fatalf("no bound metadata registered for %s; add it to boundContracts", name)
				}
				contract, ok := compiled[name]
				if !ok {
					t.Fatalf("%s not found in compiler output", name)
				}
				want, err := decodeBin(contract.Code)
				if err != nil {
					t.Fatalf("decode compiled bytecode: %v", err)
				}
				got, err := decodeBin(meta.Bin)
				if err != nil {
					t.Fatalf("decode %s Bin in %s: %v", name, file, err)
				}

				wantCode, nWant := stripMetadata(want)
				gotCode, nGot := stripMetadata(got)
				if nWant != 1 || nGot != 1 {
					t.Fatalf("expected one solc metadata trailer per contract, found compiled=%d committed=%d", nWant, nGot)
				}
				if !bytes.Equal(wantCode, gotCode) {
					t.Errorf("%s Bin in %s does not match the compiled %s.sol (%s)\n%s",
						name, file, name, regenHint, describeDiff(gotCode, wantCode))
				}
			})
		}
	}
}

// describeDiff reports the code sizes and the first differing byte offset.
func describeDiff(committed, compiled []byte) string {
	const window = 32
	i := 0
	for i < len(committed) && i < len(compiled) && committed[i] == compiled[i] {
		i++
	}
	clip := func(b []byte) string {
		end := i + window
		if end > len(b) {
			end = len(b)
		}
		return hex.EncodeToString(b[i:end])
	}
	return fmt.Sprintf("size without metadata: committed=%d compiled=%d, first difference at byte %d\n  committed: %s\n  compiled:  %s",
		len(committed), len(compiled), i, clip(committed), clip(compiled))
}
