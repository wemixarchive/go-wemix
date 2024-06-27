package compile

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fabelx/go-solc-select/pkg/config"
	"github.com/fabelx/go-solc-select/pkg/installer"
	"github.com/fabelx/go-solc-select/pkg/versions"
	"github.com/pkg/errors"
)

var (
	solcVersion string = "0.8.14"
	solcCmd     string = fmt.Sprintf("solc-%s", solcVersion)
)

func Compile(root string, sourceFiles ...string) (compiledTy, error) {
	if err := installSolc(); err != nil {
		return nil, err
	}
	if root == "" {
		root = "../contracts"
	}
	args := []string{
		"--combined-json", "bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes",
		"--optimize",                // code optimizer switched on
		"--allow-paths", ".,./,../", //default to support relative path： ./  ../  .
		fmt.Sprintf("@openzeppelin/contracts/=%s/openzeppelin/contracts/contracts/", root),
		fmt.Sprintf("@openzeppelin/contracts-upgradeable/=%s/openzeppelin/contracts-upgradeable/contracts/", root),
		"--",
	}
	// ~/.gsolc-select/artifacts/solc-0.8.14/0.8.14
	cmd := exec.Command(filepath.Join(config.SolcArtifacts, solcCmd, solcCmd), append(args, sourceFiles...)...)

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
		out := make(compiledTy)
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

func installSolc() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// check already installed
	installedVersions := versions.GetInstalled()
	if installedVersions[solcVersion] != "" {
		return nil
	}

	// install solc-0.8.14
	installed, _, err := installer.InstallSolcs(ctx, []string{solcVersion})
	if err != nil {
		return err
	}

	for _, v := range installed {
		if v == solcVersion {
			return nil
		}
	}
	return fmt.Errorf("failed to install version %s", solcVersion)
}

type compiledTy map[string]*compiler.Contract

func (compiled compiledTy) BindContracts(pkg, filename string, contracts ...string) error {
	var (
		length    = len(contracts)
		types     = make([]string, length)
		abis      = make([]string, length)
		bytecodes = make([]string, length)
		sigs      = make([]map[string]string, length)
	)

	var err error
	for i, name := range contracts {
		contract, ok := compiled[name]
		if !ok {
			return fmt.Errorf("not found contract : %v", name)
		}
		types[i] = name
		if abis[i], err = abiToString(contract); err != nil {
			return errors.Wrap(err, name)
		}
		bytecodes[i] = contract.Code
		sigs[i] = contract.Hashes
	}

	str, err := bind.Bind(types, abis, bytecodes, sigs, pkg, bind.LangGo, nil, nil)
	if err != nil {
		return err
	}

	filedata := []byte(str)
	if read, err := os.ReadFile(filename); err == nil {
		// if out file is already exists, compare the file contents
		if crypto.Keccak256Hash(read) == crypto.Keccak256Hash(filedata) {
			return nil
		}
	} else {
		// check dir is exist
		outDir := filepath.Dir(filename)
		if _, err := os.Stat(outDir); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
			if err = os.MkdirAll(outDir, 0755); err != nil {
				return err
			}
		}
	}

	return os.WriteFile(filename, filedata, 0644)
}

func abiToString(contract *compiler.Contract) (abi string, err error) {
	switch v := contract.Info.AbiDefinition.(type) {
	case string:
		abi = v
	default:
		var bytes []byte
		if bytes, err = json.Marshal(v); err == nil {
			abi = string(bytes)
		}
	}
	return
}
