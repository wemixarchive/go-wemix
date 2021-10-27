// offline_wallet.go

package jsre

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/url"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts/usbwallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pborman/uuid"
	"github.com/robertkrimen/otto"
)

var (
	offlineWalletLock = &sync.Mutex{}
	offlineWallets    = map[string]offlineWallet{}

	// this comes from console.Stdin.PromptPassword
	PromptPassword func(string) (string, error)
)

// from internal/ethapi/api.go
type SendTxArgs struct {
	From     common.Address  `json:"from"`
	To       *common.Address `json:"to"`
	Gas      *hexutil.Uint64 `json:"gas"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Value    *hexutil.Big    `json:"value"`
	Nonce    *hexutil.Uint64 `json:"nonce"`
	// We accept "data" and "input" for backwards-compatibility reasons. "input" is the
	// newer name and should be preferred by clients.
	Data  *hexutil.Bytes `json:"data"`
	Input *hexutil.Bytes `json:"input"`
}

// from accounts/usbwallet/wallet.go
type offlineWallet interface {
	Status() (string, error)
	Open(device io.ReadWriter, passphrase string) error
	Close() error
	Heartbeat() error
	Derive(path accounts.DerivationPath) (common.Address, error)
	SignTx(path accounts.DerivationPath, tx *types.Transaction, chainID *big.Int) (common.Address, *types.Transaction, error)
}

// from geth's keystore style account
type gethAccount struct {
	id  string
	key *keystore.Key
}

// copied from 'console' package.
// throwJSException panics on an otto.Value. The Otto VM will recover from the
// Go panic and throw msg as a JavaScript error.
func throwJSException(err error) otto.Value {
	val, err2 := otto.ToValue(err.Error())
	if err2 != nil {
		log.Error("Failed to serialize JavaScript exception", "exception", err, "err", err2)
	}
	panic(val)
}

// id is sha3 of random uuid
func offlineWalletNewId() string {
	return hex.EncodeToString(
		crypto.Keccak256(
			[]byte(uuid.NewRandom().String())))
}

// from metadium/metclient/util.go
// if password
// - || "": read password from stdin
// @<file-name>: <file-name> file has password
func loadGethAccount(password, fileName string) (string, *common.Address, error) {
	var err error
	id := offlineWalletNewId()

	if password == "" || password == "-" {
		if PromptPassword == nil {
			return "", nil, errors.New("Not Intiailized")
		}
		password, err = PromptPassword("Passphrase: ")
		if err != nil {
			return "", nil, err
		}
	} else if password[0] == '@' {
		pw, err := ioutil.ReadFile(password[1:])
		if err != nil {
			return "", nil, err
		}
		password = strings.TrimSpace(string(pw))
	}

	keyJson, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", nil, err
	}
	key, err := keystore.DecryptKey(keyJson, password)
	if err != nil {
		return "", nil, err
	}

	offlineWallets[id] = &gethAccount{id: id, key: key}
	return id, &key.Address, nil
}

// driver interface function implementations below
func (acct *gethAccount) Status() (string, error) {
	return "good", nil
}

func (acct *gethAccount) Open(device io.ReadWriter, passphrase string) error {
	return nil
}

func (acct *gethAccount) Close() error {
	// TODO: need to wipe the key
	acct.id = ""
	acct.key = nil
	return nil
}

func (acct *gethAccount) Heartbeat() error {
	return nil
}

func (acct *gethAccount) Derive(path accounts.DerivationPath) (common.Address, error) {
	return acct.key.Address, nil
}

func (acct *gethAccount) SignTx(path accounts.DerivationPath, tx *types.Transaction, chainID *big.Int) (common.Address, *types.Transaction, error) {
	signer := types.NewEIP155Signer(chainID)
	stx, err := types.SignTx(tx, signer, acct.key.PrivateKey)
	return acct.key.Address, stx, err
}

// open hardware wallet, either ledger or trezor
func openUsbWallet(scheme, path string) (string, *common.Address, error) {
	id := offlineWalletNewId()

	driver, err := usbwallet.OpenOffline(scheme, path)
	if err != nil {
		return "", nil, err
	} else if _, ok := driver.(offlineWallet); !ok {
		return "", nil, errors.New("Invalid Driver")
	}
	drv := driver.(offlineWallet)
	offlineWallets[id] = drv
	addr, err := drv.Derive(accounts.DefaultBaseDerivationPath)
	if err != nil {
		throwJSException(err)
	}
	return id, &addr, nil
}

// { "id": string, "address": address } offlneWalletOpen(string url, string password)
func (re *JSRE) offlineWalletOpen(call otto.FunctionCall) otto.Value {
	rawurl, err := call.Argument(0).ToString()
	if err != nil {
		throwJSException(err)
	}
	password := ""
	if call.Argument(1).IsString() {
		password, _ = call.Argument(1).ToString()
	}

	offlineWalletLock.Lock()
	defer offlineWalletLock.Unlock()

	u, err := url.Parse(rawurl)
	if err != nil {
		throwJSException(err)
	}

	path := u.Path
	if len(path) == 0 {
		// handle relative path
		path = u.Opaque
	}

	switch u.Scheme {
	case "", "geth", "gmet":
		id, addr, err := loadGethAccount(password, path)
		if err != nil {
			throwJSException(err)
		} else {
			r := map[string]string{
				"id":      id,
				"address": addr.Hex(),
			}
			env := otto.New()
			v, err := env.ToValue(r)
			if err != nil {
				throwJSException(err)
			}
			return v
		}
	case "ledger", "trezor":
		id, addr, err := openUsbWallet(u.Scheme, path)
		if err != nil {
			throwJSException(err)
		} else {
			r := map[string]string{
				"id":      id,
				"address": addr.Hex(),
			}
			env := otto.New()
			v, err := env.ToValue(r)
			if err != nil {
				throwJSException(err)
			}
			return v
		}
	default:
		// not supported
		throwJSException(errors.New("Not Supported"))
	}
	return otto.UndefinedValue()
}

// address offlneWalletAddress(string id)
func (re *JSRE) offlineWalletAddress(call otto.FunctionCall) otto.Value {
	id, err := call.Argument(0).ToString()
	if err != nil {
		throwJSException(err)
	}

	offlineWalletLock.Lock()
	w, ok := offlineWallets[id]
	offlineWalletLock.Unlock()

	if !ok {
		throwJSException(ethereum.NotFound)
		return otto.UndefinedValue()
	} else {
		addr, err := w.Derive(accounts.DefaultBaseDerivationPath)
		if err != nil {
			throwJSException(err)
		}
		v, err := otto.ToValue(addr.Hex())
		if err != nil {
			throwJSException(err)
		}
		return v
	}
}

// address offlneWalletClose(string id)
func (re *JSRE) offlineWalletClose(call otto.FunctionCall) otto.Value {
	id, err := call.Argument(0).ToString()
	if err != nil {
		throwJSException(err)
	}

	offlineWalletLock.Lock()
	defer offlineWalletLock.Unlock()

	if w, ok := offlineWallets[id]; !ok {
		throwJSException(ethereum.NotFound)
		return otto.FalseValue()
	} else {
		delete(offlineWallets, id)
		w.Close()
		return otto.TrueValue()
	}
}

// convert tx json string to SendTxArgs
func (re *JSRE) getTxArgs(jtx string) (*SendTxArgs, error) {
	var (
		args map[string]interface{}
		bi   *big.Int
		err  error
	)

	dec := json.NewDecoder(strings.NewReader(jtx))
	dec.UseNumber()
	if err = dec.Decode(&args); err != nil {
		return nil, err
	}

	fixNum := func(v interface{}) (*big.Int, error) {
		switch w := v.(type) {
		case json.Number:
			ui, err := w.Int64()
			if err != nil {
				return nil, err
			} else {
				return big.NewInt(ui), nil
			}
		case string:
			bi, _ := new(big.Int).SetString(w, 0)
			return bi, nil
		default:
			fmt.Printf("%T %v\n", v, v)
			return nil, errors.New("Unknown type")
		}
	}

	tx := &SendTxArgs{}
	for n, v := range args {
		switch n {
		case "from":
			if s, ok := v.(string); !ok || !common.IsHexAddress(s) {
				return nil, errors.New("From is not address")
			} else {
				tx.From = common.HexToAddress(s)
			}
		case "to":
			if s, ok := v.(string); !ok || !common.IsHexAddress(s) {
				return nil, errors.New("To is not address")
			} else {
				to := common.HexToAddress(s)
				tx.To = &to
			}
		case "gas":
			if bi, err = fixNum(v); err != nil {
				return nil, err
			} else {
				tx.Gas = new(hexutil.Uint64)
				*tx.Gas = hexutil.Uint64(uint64(bi.Int64()))
			}
		case "gasPrice":
			if bi, err = fixNum(v); err != nil {
				return nil, err
			} else {
				tx.GasPrice = (*hexutil.Big)(bi)
			}
		case "value":
			if bi, err = fixNum(v); err != nil {
				return nil, err
			} else {
				tx.Value = (*hexutil.Big)(bi)
			}
		case "nonce":
			if bi, err = fixNum(v); err != nil {
				return nil, err
			} else {
				tx.Nonce = new(hexutil.Uint64)
				*tx.Nonce = hexutil.Uint64(uint64(bi.Int64()))
			}
		case "data":
			if s, ok := v.(string); ok {
				data, err := hexutil.Decode(s)
				if err != nil {
					return nil, err
				}
				tx.Data = (*hexutil.Bytes)(&data)
			} else {
				return nil, errors.New("Invalid Data")
			}
		case "input":
			if s, ok := v.(string); ok {
				input, err := hexutil.Decode(s)
				if err != nil {
					return nil, err
				}
				tx.Input = (*hexutil.Bytes)(&input)
			} else {
				return nil, errors.New("Invalid Input")
			}
		}
	}
	return tx, nil
}

// []byte offlineWalletSignTx(string id, transaction tx, chainID int)
func (re *JSRE) offlineWalletSignTx(call otto.FunctionCall) otto.Value {
	id, err := call.Argument(0).ToString()
	if err != nil {
		throwJSException(err)
	}
	chainID, err := call.Argument(2).ToInteger()
	if err != nil {
		throwJSException(err)
	}

	var (
		tx    *types.Transaction
		input []byte
	)

	// javascript json object -> string -> jsTx -> types.Transaction
	JSON, _ := call.Otto.Object("JSON")
	jtx, err := JSON.Call("stringify", call.Argument(1))
	if err != nil {
		throwJSException(err)
	}

	txargs, err := re.getTxArgs(jtx.String())
	if err != nil {
		throwJSException(err)
	}

	if txargs.Data != nil {
		input = *txargs.Data
	} else if txargs.Input != nil {
		input = *txargs.Input
	}
	if txargs.To == nil {
		tx = types.NewContractCreation(uint64(*txargs.Nonce),
			(*big.Int)(txargs.Value), uint64(*txargs.Gas),
			(*big.Int)(txargs.GasPrice), input)
	} else {
		tx = types.NewTransaction(uint64(*txargs.Nonce), *txargs.To,
			(*big.Int)(txargs.Value), uint64(*txargs.Gas),
			(*big.Int)(txargs.GasPrice), input)
	}

	offlineWalletLock.Lock()
	w, ok := offlineWallets[id]
	offlineWalletLock.Unlock()

	if !ok {
		throwJSException(ethereum.NotFound)
		return otto.UndefinedValue()
	} else {
		_, stx, err := w.SignTx(accounts.DefaultBaseDerivationPath, tx,
			big.NewInt(chainID))
		if err != nil {
			throwJSException(err)
		}
		data, err := rlp.EncodeToBytes(stx)
		if err != nil {
			throwJSException(err)
		}
		v, err := otto.ToValue(hexutil.Encode(data))
		if err != nil {
			throwJSException(err)
		}
		return v
	}
}

// offlineWalletList returns the array of ledger or trezor device paths
// for mostly informational use
func (re *JSRE) offlineWalletList(call otto.FunctionCall) otto.Value {
	scheme := ""
	if call.Argument(0).IsString() {
		scheme, _ = call.Argument(0).ToString()
	}

	paths := usbwallet.ListDevices(scheme)
	env := otto.New()
	v, err := env.ToValue(paths)
	if err != nil {
		throwJSException(err)
	}
	return v
}

// EOF
