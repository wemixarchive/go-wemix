package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/wemix/governance-contract/common/compile"
)

var (
	rootFlag = flag.String("root", "../contracts", "")
	outDir   string
)

const pkg string = "gov"

func main() {
	flag.Parse()
	root := *rootFlag
	outDir = filepath.Join(root, "../../bind")
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
	} else if err := bindContracts(contracts, "registry", "Registry"); err != nil {
		panic(err)
	} else if err := bindContracts(contracts, "gov", "Gov", "GovImp"); err != nil {
		panic(err)
	} else if err := bindContracts(contracts, "ncpExit", "NCPExit", "NCPExitImp"); err != nil {
		panic(err)
	} else if err := bindContracts(contracts, "staking", "Staking", "StakingImp"); err != nil {
		panic(err)
	} else if err := bindContracts(contracts, "ballotStorage", "BallotStorage", "BallotStorageImp"); err != nil {
		panic(err)
	} else if err := bindContracts(contracts, "envStorage", "EnvStorage", "EnvStorageImp"); err != nil {
		panic(err)
	} else {
		fmt.Println("success!")
	}
}

func bindContracts(contracts map[string]*compiler.Contract, fname string, cnames ...string) error {
	length := len(cnames)
	types := make([]string, length)
	abis := make([]string, length)
	bytecodes := make([]string, length)

	for i, name := range cnames {
		contract, ok := contracts[name]
		if !ok {
			return ethereum.NotFound
		}
		types[i] = name
		abis[i] = abiToString(contract.Info.AbiDefinition)
		bytecodes[i] = contract.Code
	}

	str, err := bind.Bind(types, abis, bytecodes, nil, pkg, bind.LangGo, nil, nil)
	if err != nil {
		return err
	}

	if _, err := os.Stat(outDir); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err = os.MkdirAll(outDir, 0755); err != nil {
			return err
		}
	}
	return os.WriteFile(filepath.Join(outDir, fmt.Sprintf("%s.go", fname)), []byte(str), 0644)
}

func abiToString(definition interface{}) string {
	if str, ok := definition.(string); ok {
		return str
	}
	bytes, _ := json.Marshal(definition)
	return string(bytes)
}
