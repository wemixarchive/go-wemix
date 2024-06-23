// util.go

package metclient

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/console/prompt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractData struct {
	Name     string
	Abi      abi.ABI
	Bytecode []byte
}

type RemoteContract struct {
	Cli  *ethclient.Client
	From *keystore.Key
	To   *common.Address
	Abi  abi.ABI
	Gas  int
}

// if password
// - || "": read password from stdin
// @<file-name>: <file-name> file has password
func LoadAccount(password, fileName string) (*keystore.Key, error) {
	var err error
	if password == "" || password == "-" {
		password, err = prompt.PromptPassword("Passphrase: ")
		if err != nil {
			return nil, err
		}
	} else if password[0] == '@' {
		var pw []byte
		pw, err = os.ReadFile(password[1:])
		if err != nil {
			return nil, err
		}
		password = strings.TrimSpace(string(pw))
	}

	var keyJson []byte
	keyJson, err = os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return keystore.DecryptKey(keyJson, password)
}

type contractJson struct {
	ContractName string                   `json:"contractName"`
	Abi          []map[string]interface{} `json:"abi"`
	Bytecode     string                   `json:"bytecode"`
}

// load a contract from .json file ala truffle
func LoadJsonContract(r io.Reader) (*ContractData, error) {
	var (
		name                        string
		fileData, bytecode, abiData []byte
		abii                        abi.ABI
		err                         error
	)

	fileData, err = io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var data contractJson
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return nil, err
	}

	if data.ContractName == "" || len(data.Abi) == 0 {
		return nil, fmt.Errorf("invalid contract json file")
	}

	name = data.ContractName
	if len(data.Bytecode) > 0 {
		bytecode, err = hexutil.Decode(data.Bytecode)
		if err != nil {
			return nil, err
		}
	}

	abiData, err = json.Marshal(data.Abi)
	if err != nil {
		return nil, err
	}

	abii, err = abi.JSON(bytes.NewReader(abiData))
	if err != nil {
		return nil, err
	}

	return &ContractData{
		Name:     name,
		Abi:      abii,
		Bytecode: bytecode,
	}, nil
}

// load contract data from .js file ala remix
// var <name>_contract = web3.eth.contract([{<abi>}]);
// var <name>_data = "0x...";
func LoadJsContract(r io.Reader) (map[string]*ContractData, error) {
	var (
		err       error
		contracts = map[string]*ContractData{}
	)

	re := regexp.MustCompile(`^var ([^_]+)_(contract)[^(]+\(([^)]+)\);$|^var ([^_]+)_(data) = "(0x[0-9a-zA-Z]+)";$`)
	b := bufio.NewReaderSize(r, 1024*1024)
	for {
		var (
			name, kind string
			line, data []byte
		)

		line, _, err = b.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		submatches := re.FindSubmatch(line)
		if submatches == nil || len(submatches) != 7 {
			continue
		} else if len(submatches[1]) > 0 {
			name, kind, data = string(submatches[1]), string(submatches[2]), submatches[3]
		} else if len(submatches[4]) > 0 {
			name, kind, data = string(submatches[4]), string(submatches[5]), submatches[6]
		} else {
			continue
		}

		c, ok := contracts[name]
		if !ok {
			c = &ContractData{}
			c.Name = name
			contracts[name] = c
		}

		if kind == "contract" {
			c.Abi, err = abi.JSON(bytes.NewReader(data))
			if err != nil {
				continue
			}
		} else if kind == "data" {
			c.Bytecode, err = hexutil.Decode(string(data))
			if err != nil {
				continue
			}
		}
	}

	return contracts, nil
}

func LoadContract(fn string, name string) (*ContractData, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	contract, err := LoadJsonContract(f)
	if err == nil && contract.Name == name {
		return contract, nil
	}

	f.Seek(0, 0)
	contracts, err := LoadJsContract(f)
	if err != nil {
		return nil, err
	} else if contract, ok := contracts[name]; ok {
		return contract, nil
	} else {
		return nil, ethereum.NotFound
	}
}

func getReceipt(ctx context.Context, cli *ethclient.Client, hash common.Hash, isContract bool, msinterval, count int) (receipt *types.Receipt, err error) {
	d := time.Millisecond * time.Duration(msinterval)
	nilAddr := common.Address{}
	for i := 0; i < count; i++ {
		receipt, err = cli.TransactionReceipt(ctx, hash)
		if err != nil && err != ethereum.NotFound {
			return
		} else if receipt != nil &&
			(!isContract || receipt.ContractAddress != nilAddr) {
			return
		}
		time.Sleep(d)
	}
	err = fmt.Errorf("timed out")
	return
}

func GetReceipt(ctx context.Context, cli *ethclient.Client, hash common.Hash, msinterval, count int) (receipt *types.Receipt, err error) {
	return getReceipt(ctx, cli, hash, false, msinterval, count)
}

func GetContractReceipt(ctx context.Context, cli *ethclient.Client, hash common.Hash, msinterval, count int) (receipt *types.Receipt, err error) {
	return getReceipt(ctx, cli, hash, true, msinterval, count)
}

func Deploy(ctx context.Context, cli *ethclient.Client, from *keystore.Key,
	contractData *ContractData, args []interface{}, gas, _gasPrice int) (
	hash common.Hash, err error) {
	// pull transaction parameters from wemix node
	chainId, gasPrice, nonce, err := GetOpportunisticTxParams(
		ctx, cli, from.Address, false, true)
	if err != nil {
		return
	}

	if _gasPrice > 0 {
		gasPrice = big.NewInt(int64(_gasPrice))
	}

	var data []byte
	if len(args) == 0 {
		data = contractData.Bytecode
	} else {
		data, err = contractData.Abi.Pack("", args...)
		if err != nil {
			return
		}
		data = append(contractData.Bytecode, data...)
	}

	var tx, stx *types.Transaction
	tx = types.NewContractCreation(nonce.Uint64(), nil, uint64(gas), gasPrice,
		data)

	signer := types.NewEIP155Signer(chainId)
	stx, err = types.SignTx(tx, signer, from.PrivateKey)
	if err != nil {
		return
	}

	err = cli.SendTransaction(ctx, stx)
	if err != nil {
		return
	}

	hash = stx.Hash()
	return
}

func IsArray(x interface{}) bool {
	if x == nil {
		return false
	}
	y := reflect.TypeOf(x)
	switch y.Kind() {
	case reflect.Slice:
		return true
	case reflect.Array:
		return true
	default:
		return false
	}
}

func CallContract(ctx context.Context, contract *RemoteContract,
	method string, input, output interface{}, block *big.Int) error {
	var in, out []byte
	var err error

	if !IsArray(input) {
		if input == nil {
			in, err = contract.Abi.Pack(method)
		} else {
			in, err = contract.Abi.Pack(method, input)
		}
	} else {
		if len(input.([]interface{})) == 0 {
			in, err = contract.Abi.Pack(method)
		} else {
			in, err = contract.Abi.Pack(method, input.([]interface{})...)
		}
	}

	if err != nil {
		return err
	}

	msg := ethereum.CallMsg{
		To:   contract.To,
		Data: in,
	}
	out, err = contract.Cli.CallContract(ctx, msg, block)
	if err != nil {
		return err
	}

	if !IsArray(output) {
		if output == nil {
			return fmt.Errorf("output is nil")
		} else {
			err = contract.Abi.UnpackIntoInterface(output, method, out)
		}
	} else {
		if len(output.([]interface{})) == 0 {
			return fmt.Errorf("output is empty array")
		} else {
			err = contract.Abi.UnpackIntoInterface(output.([]interface{}), method, out)
		}
	}
	return err
}

func SendContract(ctx context.Context, contract *RemoteContract, method string,
	args interface{}) (
	hash common.Hash, err error) {
	var data []byte

	if !IsArray(args) {
		if args == nil {
			data, err = contract.Abi.Pack(method)
		} else {
			data, err = contract.Abi.Pack(method, args)
		}
	} else {
		if len(args.([]interface{})) == 0 {
			data, err = contract.Abi.Pack(method)
		} else {
			data, err = contract.Abi.Pack(method, args.([]interface{})...)
		}
	}
	if err != nil {
		return
	}

	chainId, gasPrice, nonce, err := GetOpportunisticTxParams(
		ctx, contract.Cli, contract.From.Address, false, true)
	if err != nil {
		return
	}

	var tx, stx *types.Transaction
	tx = types.NewTransaction(nonce.Uint64(), *contract.To, nil,
		uint64(contract.Gas), gasPrice, data)

	signer := types.NewEIP155Signer(chainId)
	stx, err = types.SignTx(tx, signer, contract.From.PrivateKey)
	if err != nil {
		return
	}

	err = contract.Cli.SendTransaction(ctx, stx)
	if err != nil {
		return
	}

	hash = stx.Hash()
	return
}

func SendValue(ctx context.Context, cli *ethclient.Client, from *keystore.Key, to common.Address, amount, gas, _gasPrice int) (hash common.Hash, err error) {
	chainId, gasPrice, nonce, err := GetOpportunisticTxParams(
		ctx, cli, from.Address, false, true)
	if err != nil {
		return
	}
	if _gasPrice > 0 {
		gasPrice = big.NewInt(int64(_gasPrice))
	}

	var tx, stx *types.Transaction
	tx = types.NewTransaction(nonce.Uint64(), to, big.NewInt(int64(amount)),
		uint64(gas), gasPrice, nil)

	signer := types.NewEIP155Signer(chainId)
	stx, err = types.SignTx(tx, signer, from.PrivateKey)
	if err != nil {
		return
	}

	err = cli.SendTransaction(ctx, stx)
	if err != nil {
		return
	}

	hash = stx.Hash()
	return
}

func ToBytes32(b string) [32]byte {
	var b32 [32]byte
	if len(b) > len(b32) {
		b = b[len(b)-len(b32):]
	}
	copy(b32[:], []byte(b))
	return b32
}

// packNum packs the given number (using the reflect value) and will cast it to appropriate number representation
func PackNum(value reflect.Value) []byte {
	switch value.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return math.U256Bytes(new(big.Int).SetUint64(value.Uint()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return math.U256Bytes(big.NewInt(value.Int()))
	case reflect.Ptr:
		return math.U256Bytes(new(big.Int).Set(value.Interface().(*big.Int)))
	default:
		panic("abi: fatal error")
	}
}

// EOF
