package main

import (
	"flag"
	"fmt"
	"path/filepath"

	compile "github.com/ethereum/go-ethereum/wemix/governance-contract"
)

const pkg string = "gov"

var (
	rootFlag = flag.String("root", "../contracts", "")
)

func main() {
	flag.Parse()
	root := *rootFlag
	outDir := filepath.Join(root, "../../bind")
	if contracts, err := compile.Compile(root,
		filepath.Join(root, "Registry.sol"),
		filepath.Join(root, "Gov.sol"),
		filepath.Join(root, "GovImp.sol"),
		filepath.Join(root, "NCPExit.sol"),
		filepath.Join(root, "NCPExitImp.sol"),
		filepath.Join(root, "Staking.sol"),
		filepath.Join(root, "StakingImp.sol"),
		filepath.Join(root, "storage", "BallotStorage.sol"),
		filepath.Join(root, "storage", "BallotStorageImp.sol"),
		filepath.Join(root, "storage", "EnvStorage.sol"),
		filepath.Join(root, "storage", "EnvStorageImp.sol"),
	); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "registry.go"), "Registry"); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "gov.go"), "Gov", "GovImp"); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "ncpExit.go"), "NCPExit", "NCPExitImp"); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "staking.go"), "Staking", "StakingImp"); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "ballotStorage.go"), "BallotStorage", "BallotStorageImp"); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "envStorage.go"), "EnvStorage", "EnvStorageImp"); err != nil {
		panic(err)
	} else {
		fmt.Println("success!")
	}
}
