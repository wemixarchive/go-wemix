package compile

import (
	"path/filepath"

	gov "github.com/ethereum/go-ethereum/wemix/bind"
)

// BindPackage is the Go package name of the generated governance bindings.
const BindPackage = "gov"

// Binding describes one generated file under wemix/bind and the contracts it holds.
type Binding struct {
	File      string
	Contracts []string
}

// buildGroup is a set of contracts compiled against one OpenZeppelin release.
type buildGroup struct {
	oz        OpenZeppelin
	sources   []string // relative to root
	contracts []string
}

// buildGroups records which OpenZeppelin release each contract is built with.
// Deployed contracts use the release their on-chain runtime code was built
// with, so the bindings match the deployed code apart from the metadata hash.
// GovImp is the next, not yet deployed, implementation and stays on the
// default release.
func buildGroups() []buildGroup {
	return []buildGroup{
		{
			// contracts v4.6.0, contracts-upgradeable v4.7.0
			oz: DefaultOpenZeppelin(),
			sources: []string{
				"Gov.sol",
				"GovImp.sol",
				"Staking.sol",
				filepath.Join("storage", "EnvStorage.sol"),
				filepath.Join("storage", "EnvStorageImp.sol"),
			},
			contracts: []string{gov.GOV, gov.GOV_IMP, gov.STAKING, gov.ENVSTORAGE, gov.ENVSTORAGE_IMP},
		},
		{
			oz: OpenZeppelin{
				Contracts:   "openzeppelin/v4.9.6/contracts",
				Upgradeable: "openzeppelin/v4.9.6/contracts-upgradeable",
			},
			sources: []string{
				"Registry.sol",
				"NCPExit.sol",
				"NCPExitImp.sol",
				"StakingImp.sol",
				filepath.Join("storage", "BallotStorage.sol"),
				filepath.Join("storage", "BallotStorageImp.sol"),
			},
			contracts: []string{gov.REGISTRY, gov.NCPEXIT, gov.NCPEXIT_IMP, gov.STAKING_IMP, gov.BALLOTSTORAGE, gov.BALLOTSTORAGE_IMP},
		},
	}
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

// CompileBindings compiles every build group under root with its OpenZeppelin
// release and merges the picked contracts, producing exactly the contracts
// that go into wemix/bind.
func CompileBindings(root string) (compiledTy, error) {
	contracts := make(compiledTy)
	for _, g := range buildGroups() {
		sources := make([]string, len(g.sources))
		for i, s := range g.sources {
			sources[i] = filepath.Join(root, s)
		}
		compiled, err := CompileWith(root, g.oz, sources...)
		if err != nil {
			return nil, err
		}
		picked, err := compiled.Pick(g.contracts...)
		if err != nil {
			return nil, err
		}
		if contracts, err = Merge(contracts, picked); err != nil {
			return nil, err
		}
	}
	return contracts, nil
}
