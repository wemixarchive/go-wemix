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
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "gen_registry_abi.go"), gov.CNAME_Registry); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "gen_gov_abi.go"), gov.CNAME_Gov, gov.CNAME_GovImp); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "gen_ncpExit_abi.go"), gov.CNAME_NCPExit, gov.CNAME_NCPExitImp); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "gen_staking_abi.go"), gov.CNAME_Staking, gov.CNAME_StakingImp); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "gen_ballotStorage_abi.go"), gov.CNAME_BallotStorage, gov.CNAME_BallotStorageImp); err != nil {
		panic(err)
	} else if err := contracts.BindContracts(pkg, filepath.Join(outDir, "gen_envStorage_abi.go"), gov.CNAME_EnvStorage, gov.CNAME_EnvStorageImp); err != nil {
		panic(err)
	} else {
		fmt.Println("success!")
	}
}
