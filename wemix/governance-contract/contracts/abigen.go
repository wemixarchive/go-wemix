package main

import (
	"flag"
	"fmt"
	"path/filepath"

	gov "github.com/ethereum/go-ethereum/wemix/bind"
	compile "github.com/ethereum/go-ethereum/wemix/governance-contract"
)

const pkg string = "gov"

var (
	rootFlag = flag.String("root", "../contracts", "")
)

// buildGroup is a set of contracts compiled against one OpenZeppelin release.
type buildGroup struct {
	oz        compile.OpenZeppelin
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
			oz: compile.DefaultOpenZeppelin(),
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
			oz: compile.OpenZeppelin{
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

func main() {
	flag.Parse()
	root := *rootFlag
	outDir := filepath.Join(root, "../../bind")

	contracts, err := compile.Merge()
	if err != nil {
		panic(err)
	}
	for _, g := range buildGroups() {
		sources := make([]string, len(g.sources))
		for i, s := range g.sources {
			sources[i] = filepath.Join(root, s)
		}
		compiled, err := compile.CompileWith(root, g.oz, sources...)
		if err != nil {
			panic(err)
		}
		picked, err := compiled.Pick(g.contracts...)
		if err != nil {
			panic(err)
		}
		if contracts, err = compile.Merge(contracts, picked); err != nil {
			panic(err)
		}
	}

	if err := contracts.BindContracts(pkg, filepath.Join(outDir, "gen_registry_abi.go"), gov.REGISTRY); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "gen_gov_abi.go"), gov.GOV, gov.GOV_IMP); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "gen_ncpExit_abi.go"), gov.NCPEXIT, gov.NCPEXIT_IMP); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "gen_staking_abi.go"), gov.STAKING, gov.STAKING_IMP); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "gen_ballotStorage_abi.go"), gov.BALLOTSTORAGE, gov.BALLOTSTORAGE_IMP); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "gen_envStorage_abi.go"), gov.ENVSTORAGE, gov.ENVSTORAGE_IMP); err != nil {
		panic(err)
	} else {
		fmt.Println("success!")
	}
}
