// metadiumcmd.go

package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/metadium/metclient"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"gopkg.in/urfave/cli.v1"
)

// gmet metadium new-account
var (
	metadiumCommand = cli.Command{
		Name:      "metadium",
		Usage:     "Metadium helper commands",
		ArgsUsage: "",
		Category:  "METADIUM COMMANDS",
		Description: `

Metadium helper commands, create a new account, a new node id, a new genesis file, or a new admin contract file.`,
		Subcommands: []cli.Command{
			{
				Name:   "new-account",
				Usage:  "Create a new account",
				Action: utils.MigrateFlags(newAccount),
				Flags: []cli.Flag{
					utils.PasswordFileFlag,
					outFlag,
				},
				Description: `
    geth metadium new-account --out <file>

Creates a new account and saves it in the given file name.
`,
			},
			{
				Name:   "new-nodekey",
				Usage:  "Create a new node key",
				Action: utils.MigrateFlags(newNodeKey),
				Flags: []cli.Flag{
					outFlag,
				},
				Description: `
    geth metadium new-nodekey --out <file>

Creates a new node key and saves it in the given file name.
`,
			},
			{
				Name:   "nodeid",
				Usage:  "Print node id from node key",
				Action: utils.MigrateFlags(nodeKey2Id),
				Description: `
    geth metadium new-nodekey <file>

Print node id from node key.
`,
			},
			{
				Name:      "genesis",
				Usage:     "Create a new genesis file",
				Action:    utils.MigrateFlags(genGenesis),
				ArgsUsage: "<file-name>",
				Flags: []cli.Flag{
					dataFileFlag,
					genesisTemplateFlag,
					outFlag,
				},
				Description: `
    geth metadium genesis [--data <file> --genesis <file> --out <file>]

Generate a new genesis file from a template.

Stdin is used when --data is missing, and stdout is used for --out.

Data consists of "<account> <tokens>" or "<node id>".`,
			},
			{
				Name:   "admin-contract",
				Usage:  "Create an admin contract",
				Action: utils.MigrateFlags(genAdminContract),
				Flags: []cli.Flag{
					dataFileFlag,
					adminTemplateFlag,
					outFlag,
				},
				Description: `
    geth metadium admin-contract [--data <file> --admin <file> --out <file>]

Generate a new admin contract file from a template.

Stdin is used when --data is missing, and stdout is used for --out.

Data consists of "<account> <balance> <tokens>" or "<node id>".
The first account becomes the coinbase of the genesis block, and the creator of the admin contract.
The first node becomes the boot miner who's allowed to generate blocks before admin contract gets created.`,
			},
			{
				Name:   "deploy-contract",
				Usage:  "Deploy a contract",
				Action: utils.MigrateFlags(deployContract),
				Flags: []cli.Flag{
					utils.PasswordFileFlag,
					urlFlag,
					gasFlag,
					gasPriceFlag,
				},
				Description: `
    geth metadium deploy-contract [--password data <file> --url <url> --gas <gas> --gasprice <gas-price>] <account-file> <contract-name> <contract-file.[js|json]>

Deploy a contract from a contract file in .js or .json format.`,
			},
		},
	}

	dataFileFlag = cli.StringFlag{
		Name:  "data",
		Usage: "data file",
	}
	genesisTemplateFlag = cli.StringFlag{
		Name:  "genesis",
		Usage: "genesis template file",
	}
	adminTemplateFlag = cli.StringFlag{
		Name:  "admin",
		Usage: "admin contract template file",
	}
	outFlag = cli.StringFlag{
		Name:  "out",
		Usage: "out file",
	}
	gasFlag = cli.IntFlag{
		Name:  "gas",
		Usage: "gas amount",
	}
	gasPriceFlag = cli.IntFlag{
		Name:  "gasprice",
		Usage: "gas price",
	}
	urlFlag = cli.StringFlag{
		Name:  "url",
		Usage: "url of gmet node",
	}
)

func newAccount(ctx *cli.Context) error {
	var err error

	w := os.Stdout
	if fn := ctx.String(outFlag.Name); fn != "" {
		w, err = os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			utils.Fatalf("%v", err)
		}
	}

	password := getPassPhrase("Please give a password. Do not forget this password.", true, 0, utils.MakePasswordList(ctx))

	key, err := keystore.NewKey(rand.Reader)
	if err != nil {
		return err
	}

	defer func() {
		b := key.PrivateKey.D.Bits()
		for i := range b {
			b[i] = 0
		}
	}()

	keyjson, err := keystore.EncryptKey(key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(keyjson))
	return err
}

func newNodeKey(ctx *cli.Context) error {
	nodeKey, err := crypto.GenerateKey()
	if err != nil {
		return err
	}
	if err = crypto.SaveECDSA(ctx.String(outFlag.Name), nodeKey); err != nil {
		return err
	}
	return nil
}

func nodeKey2Id(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Nodekey file name is not given.")
	}
	nodeKey, err := crypto.LoadECDSA(ctx.Args()[0])
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", discover.PubkeyID(&nodeKey.PublicKey))
	return nil
}

type genesisConfig struct {
	Members []*struct {
		Addr    string `json:"addr"`
		Balance int64  `json:"balance"`
		Stake   int64  `json:"stake"`
	} `json:"members"`
	Nodes []*struct {
		Name string `json:"name"`
		Ip   string `json:"ip"`
		Port int    `json:"port"`
		Id   string `json:"id"`
	} `json:"nodes"`
}

func loadGenesisConfig(r io.Reader) (*genesisConfig, error) {
	var config genesisConfig
	if data, err := ioutil.ReadAll(r); err != nil {
		return nil, err
	} else if err = json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	if len(config.Members) == 0 || len(config.Nodes) == 0 {
		return nil, fmt.Errorf("At least one account and node are required.")
	}

	for _, m := range config.Members {
		// to conforming form to avoid checksum error
		m.Addr = common.HexToAddress(m.Addr).Hex()
	}
	for _, n := range config.Nodes {
		if !(len(n.Id) == 128 || len(n.Id) == 130) {
			return nil, fmt.Errorf("Not a node id: %s\n", n.Id)
		}
		if len(n.Id) == 128 {
			n.Id = "0x" + n.Id
		}
	}

	return &config, nil
}

func genGenesis(ctx *cli.Context) error {
	var err error

	var genesis map[string]interface{}
	if fn := ctx.String(genesisTemplateFlag.Name); fn == "" {
		utils.Fatalf("Genesis template is not specified.")
	} else if data, err := ioutil.ReadFile(fn); err != nil {
		return err
	} else if err = json.Unmarshal(data, &genesis); err != nil {
		return err
	}

	r := os.Stdin
	if fn := ctx.String(dataFileFlag.Name); fn != "" {
		r, err = os.Open(fn)
		if err != nil {
			utils.Fatalf("%v", err)
		}
	}

	config, err := loadGenesisConfig(r)
	if err != nil {
		return err
	}

	w := os.Stdout
	if fn := ctx.String(outFlag.Name); fn != "" {
		w, err = os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			utils.Fatalf("%v", err)
		}
	}

	if len(config.Members) <= 0 || len(config.Nodes) <= 0 {
		utils.Fatalf("At least one member and node are required.")
	}

	genesis["coinbase"] = config.Members[0].Addr
	genesis["extraData"] = config.Nodes[0].Id
	alloc := map[string]map[string]string{}
	for _, m := range config.Members {
		alloc[m.Addr] = map[string]string{
			"balance": strings.ToLower(hexutil.EncodeUint64(uint64(m.Balance))),
		}
	}
	genesis["alloc"] = alloc

	x, err := json.MarshalIndent(genesis, "", "  ")
	if err != nil {
		return err
	}
	w.Write(x)
	return nil
}

func genAdminContract(ctx *cli.Context) error {
	var err error

	var f io.Reader
	if fn := ctx.String(adminTemplateFlag.Name); fn == "" {
		utils.Fatalf("Admin contract template is not specified.")
	} else {
		f, err = os.Open(fn)
		if err != nil {
			utils.Fatalf("%v", err)
		}
	}

	r := os.Stdin
	if fn := ctx.String(dataFileFlag.Name); fn != "" {
		r, err = os.Open(fn)
		if err != nil {
			utils.Fatalf("%v", err)
		}
	}

	config, err := loadGenesisConfig(r)
	if err != nil {
		return err
	}

	w := os.Stdout
	if fn := ctx.String(outFlag.Name); fn != "" {
		w, err = os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			utils.Fatalf("%v", err)
		}
	}

	tokens := int64(0)
	for _, m := range config.Members {
		tokens += m.Stake
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		if strings.Index(l, "// To Be Substituted") < 0 {
			_, err = fmt.Fprintln(w, l)
			if err != nil {
				return err
			}
			continue
		}

		ll := strings.TrimSpace(l)
		if strings.Index(ll, "tokens") == 0 {
			_, err = fmt.Fprintf(w, "        tokens = %d;\n", tokens)
		} else if strings.Index(ll, "address[") == 0 {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("        address[%d] memory _members = [ ", len(config.Members)))
			first := true
			for _, m := range config.Members {
				if first {
					first = false
				} else {
					b.WriteString(", ")
				}
				b.WriteString(fmt.Sprintf("address(%s)", m.Addr))
			}
			b.Write([]byte(" ];\n"))
			_, err = b.WriteTo(w)
		} else if strings.Index(ll, "int[") == 0 {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("        int[%d] memory _stakes = [ ", len(config.Members)))
			first := true
			for _, m := range config.Members {
				if first {
					first = false
				} else {
					b.WriteString(", ")
				}
				b.WriteString(fmt.Sprintf("int(%d)", m.Stake))
			}
			b.Write([]byte(" ];\n"))
			_, err = b.WriteTo(w)
		} else if strings.Index(ll, "Node[") == 0 {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("        Node[%d] memory _nodes = [ ", len(config.Nodes)))
			first := true
			for _, n := range config.Nodes {
				if first {
					first = false
				} else {
					b.WriteString(", ")
				}
				b.WriteString(fmt.Sprintf(`Node(true, "%s", "%s", "%s", %d, 0, 0, "", "")`, n.Name, n.Id[2:], n.Ip, n.Port))
			}
			b.Write([]byte(" ];\n"))
			_, err = b.WriteTo(w)
		} else {
			_, err = fmt.Fprintln(w, l)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func deployContract(ctx *cli.Context) error {
	var err error

	passwd := ctx.String(utils.PasswordFileFlag.Name)
	url := ctx.String(urlFlag.Name)
	gas := ctx.Int(gasFlag.Name)
	gasPrice := ctx.Int(gasPriceFlag.Name)

	if len(url) == 0 || len(ctx.Args()) != 3 {
		return fmt.Errorf("Invalid Arguments")
	}

	accountFile, contractName, contractFile := ctx.Args()[0], ctx.Args()[1], ctx.Args()[2]

	var acct *keystore.Key
	acct, err = metclient.LoadAccount(passwd, accountFile)
	if err != nil {
		return err
	}

	var contractData *metclient.ContractData
	contractData, err = metclient.LoadContract(contractFile, contractName)
	if err != nil {
		return err
	}

	ctxx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var cli *ethclient.Client
	cli, err = ethclient.Dial(url)
	if err != nil {
		return err
	}

	var hash common.Hash
	hash, err = metclient.Deploy(ctxx, cli, acct, contractData, nil, gas,
		gasPrice)
	if err != nil {
		return err
	}

	var receipt *types.Receipt
	receipt, err = metclient.GetContractReceipt(ctxx, cli, hash, 500, 60)
	if err != nil {
		return err
	} else {
		if receipt.Status == 1 {
			fmt.Printf("Contract mined! ")
		} else {
			fmt.Printf("Contract failed with %d! ", receipt.Status)
		}
		fmt.Printf("address: %s transactionHash: %s\n",
			receipt.ContractAddress.String(), hash.String())
	}

	return nil
}

// EOF
