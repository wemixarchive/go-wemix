package compile

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/fabelx/go-solc-select/pkg/config"
	"github.com/fabelx/go-solc-select/pkg/installer"
	"github.com/fabelx/go-solc-select/pkg/versions"
)

var (
	solcVersion string = "0.8.14"
	solcCmd     string = fmt.Sprintf("solc-%s", solcVersion)
)

func Compile(root string, sourceFiles ...string) (map[string]*compiler.Contract, error) {
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
		fmt.Sprintf("@openzeppelin/=%s/openzeppelin/", root),
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
