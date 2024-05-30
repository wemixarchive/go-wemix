package compile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/ethereum/go-ethereum/common/compiler"
)

func Compile(root string, sourceFiles ...string) (map[string]*compiler.Contract, error) {
	if root == "" {
		root = "../contracts"
	}
	args := []string{
		"--combined-json", "bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes",
		"--optimize",                // code optimizer switched on
		"--allow-paths", ".,./,../", //default to support relative path： ./  ../  .
		fmt.Sprintf("@openzeppelin/=%s/openzeppelin/", root),
		"--",
	}
	cmd := exec.Command("solc", append(args, sourceFiles...)...)

	var stderr, stdout bytes.Buffer
	cmd.Stderr, cmd.Stdout = &stderr, &stdout
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("solc: %v\n%s", err, stderr.Bytes())
	}

	// compiler version
	var compilerVersion = struct {
		Version string
	}{}

	if err := json.Unmarshal(stdout.Bytes(), &compilerVersion); err != nil {
		return nil, err
	} else if contracts, err := compiler.ParseCombinedJSON(stdout.Bytes(), "", "", compilerVersion.Version, strings.Join(args, " ")); err != nil {
		return nil, err
	} else {
		out := make(map[string]*compiler.Contract)
		for name, v := range contracts {
			n := strings.Split(name, ":")[1]
			if _, ok := out[n]; ok {
				return nil, fmt.Errorf("duplicated contract name: %s", name)
			} else {
				out[n] = v
			}
		}
		return out, nil
	}
}
