package compile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/wemix/governance-contract/solcdownloader"
	"github.com/pkg/errors"
)

// solcVersion is the solc release every governance contract is compiled with.
const solcVersion = "0.8.6"

// OpenZeppelin names the directories, relative to the compile root, that the
// @openzeppelin/contracts and @openzeppelin/contracts-upgradeable import
// prefixes resolve to. Each directory is an OpenZeppelin repository checkout.
type OpenZeppelin struct {
	Contracts   string
	Upgradeable string
}

// DefaultOpenZeppelin is the OpenZeppelin checkout pair Compile uses
// (contracts v4.6.0, contracts-upgradeable v4.7.0).
func DefaultOpenZeppelin() OpenZeppelin {
	return OpenZeppelin{
		Contracts:   "openzeppelin/contracts",
		Upgradeable: "openzeppelin/contracts-upgradeable",
	}
}

// Compile compiles sourceFiles against DefaultOpenZeppelin.
func Compile(root string, sourceFiles ...string) (compiledTy, error) {
	return CompileWith(root, DefaultOpenZeppelin(), sourceFiles...)
}

// CompileWith compiles sourceFiles with the @openzeppelin imports resolved to
// the checkouts named by oz. An empty root means "../contracts".
func CompileWith(root string, oz OpenZeppelin, sourceFiles ...string) (compiledTy, error) {
	if oz.Contracts == "" || oz.Upgradeable == "" {
		return nil, fmt.Errorf("openzeppelin directories must be set: %+v", oz)
	}
	solcPath, err := solcdownloader.GetSolcBin(solcVersion)
	if err != nil {
		return nil, err
	}
	if root == "" {
		root = "../contracts"
	}
	args := []string{
		"--combined-json", "bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes",
		"--optimize",                // code optimizer switched on
		"--allow-paths", ".,./,../", //default to support relative path： ./  ../  .
		fmt.Sprintf("@openzeppelin/contracts/=%s/%s/contracts/", root, oz.Contracts),
		fmt.Sprintf("@openzeppelin/contracts-upgradeable/=%s/%s/contracts/", root, oz.Upgradeable),
		"--",
	}
	// ~/.gsolc-select/artifacts/solc-0.8.14/0.8.14
	cmd := exec.Command(solcPath, append(args, sourceFiles...)...)

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

type compiledTy map[string]*compiler.Contract

// Pick returns a compiledTy holding only the named contracts. It fails if any
// name is missing, so a typo cannot silently drop a contract from a binding.
func (compiled compiledTy) Pick(names ...string) (compiledTy, error) {
	out := make(compiledTy, len(names))
	for _, name := range names {
		c, ok := compiled[name]
		if !ok {
			return nil, fmt.Errorf("not found contract : %v", name)
		}
		out[name] = c
	}
	return out, nil
}

// Merge combines sets into one compiledTy. A contract name present in more
// than one set is an error, because it would be ambiguous which build to bind.
func Merge(sets ...compiledTy) (compiledTy, error) {
	out := make(compiledTy)
	for _, set := range sets {
		for name, c := range set {
			if _, ok := out[name]; ok {
				return nil, fmt.Errorf("duplicated contract name: %s", name)
			}
			out[name] = c
		}
	}
	return out, nil
}

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
