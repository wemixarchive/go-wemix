package compile

import (
	"path/filepath"

	gov "github.com/ethereum/go-ethereum/wemix/bind"
)

// BindPackage is the Go package name of the generated governance bindings.
const BindPackage = "gov"

// The GovImp implementation deployed on mainnet is the pre-W1G source,
// preserved as mock/GovImpLegacy.sol, compiled with solc 0.8.2. The GovImp
// binding records that deployed bytecode; a follow-up PR updates GovImp to the
// current GovImp.sol (W1G-fixed).
const (
	govImpSolcVersion = "0.8.2"
	govImpSource      = "GovImpLegacy"
)

var govImpSourceFile = filepath.Join("mock", "GovImpLegacy.sol")

// Binding describes one generated file under wemix/bind and the contracts it holds.
type Binding struct {
	File      string
	Contracts []string
}

// govSources lists the production governance sources relative to the contracts root.
var govSources = []string{
	"Registry.sol",
	"Gov.sol",
	"GovImp.sol",
	"NCPExit.sol",
	"NCPExitImp.sol",
	"Staking.sol",
	"StakingImp.sol",
	filepath.Join("storage", "BallotStorage.sol"),
	filepath.Join("storage", "BallotStorageImp.sol"),
	filepath.Join("storage", "EnvStorage.sol"),
	filepath.Join("storage", "EnvStorageImp.sol"),
}

// Bindings lists every generated binding file and the contracts bound into it.
var Bindings = []Binding{
	{"gen_registry_abi.go", []string{gov.REGISTRY}},
	{"gen_gov_abi.go", []string{gov.GOV, gov.GOV_IMP}},
	{"gen_ncpExit_abi.go", []string{gov.NCPEXIT, gov.NCPEXIT_IMP}},
	{"gen_staking_abi.go", []string{gov.STAKING, gov.STAKING_IMP}},
	{"gen_ballotStorage_abi.go", []string{gov.BALLOTSTORAGE, gov.BALLOTSTORAGE_IMP}},
	{"gen_envStorage_abi.go", []string{gov.ENVSTORAGE, gov.ENVSTORAGE_IMP}},
}

// SourceFiles returns the production governance sources joined onto root.
func SourceFiles(root string) []string {
	files := make([]string, len(govSources))
	for i, src := range govSources {
		files[i] = filepath.Join(root, src)
	}
	return files
}

// CompileBindings compiles the governance sources under root and replaces the
// GovImp bytecode with the deployed GovImpLegacy build, producing exactly the
// contracts that go into wemix/bind.
func CompileBindings(root string) (compiledTy, error) {
	govImp, err := CompileWithVersion(govImpSolcVersion, root, filepath.Join(root, govImpSourceFile))
	if err != nil {
		return nil, err
	}
	contracts, err := Compile(root, SourceFiles(root)...)
	if err != nil {
		return nil, err
	}
	if err := contracts.ReplaceCode(govImp, govImpSource, gov.GOV_IMP); err != nil {
		return nil, err
	}
	return contracts, nil
}
