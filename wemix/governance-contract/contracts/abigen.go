package main

import (
	"flag"
	"fmt"
	"path/filepath"

	compile "github.com/ethereum/go-ethereum/wemix/governance-contract"
)

var (
	rootFlag = flag.String("root", "../contracts", "")
)

func main() {
	flag.Parse()
	root := *rootFlag
	outDir := filepath.Join(root, "../../bind")
	contracts, err := compile.CompileBindings(root)
	if err != nil {
		panic(err)
	}
	for _, b := range compile.Bindings {
		if err := contracts.BindContracts(compile.BindPackage, filepath.Join(outDir, b.File), b.Contracts...); err != nil {
			panic(err)
		}
	}
	fmt.Println("success!")
}
