package sim

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ctx             = context.Background()
	defaultGasLimit = uint64(100000000)
)

// Contract struct holds data before compilation and information after compilation.
type Contract struct {
	Client   *Client
	Compiled *compiler.Contract
	Address  common.Address
	Logic    *Contract
	Alias    string
	Name     string
	Abi      *abi.ABI
	Code     []byte
}

// NewContract can specify the version of solidity
func NewContract(name, alias string, contracts map[string]*compiler.Contract) (*Contract, error) {
	r := &Contract{
		Alias: alias,
		Name:  name,
	}
	return r, r.init(contracts)
}

func (p *Contract) init(contracts map[string]*compiler.Contract) error {
	contract, ok := contracts[p.Name]
	if !ok {
		return fmt.Errorf("not found `%v` after compiling", p.Name)
	}
	//Get the contract to test from the compiled contracts.
	var err error
	p.Compiled = contract
	p.Abi, err = parseABI(contract.Info.AbiDefinition)
	if err != nil {
		return err
	}
	p.Code = common.FromHex(contract.Code)
	collectEvents(p.Abi.Events)
	collectErrors(p.Abi.Errors)

	fmt.Printf("compiled > %v [%v] bytes: %v\n", p.Alias, p.Name, len(p.Code))
	return nil
}

// Deploy makes creation contract tx and receives the result by receipt.
func (p *Contract) Deploy(t *testing.T, client *Client, args ...interface{}) (receipt *types.Receipt) {
	p.Client = client

	input, err := p.Abi.Pack("", args...) //constructor's inputs
	require.NoError(t, err, len(args))

	_, receipt = p.Client.SendTransaction(t, common.Address{}, common.Address{}, common.Big0, append(p.Code, input...))
	require.Truef(t, receipt.Status == types.ReceiptStatusSuccessful, "deploy error. name: %s, alias: %s", p.Name, p.Alias)
	p.Address = receipt.ContractAddress

	collectContract(p.Address, p.Alias)
	p.Logic = client.EnrollContract(t, p)

	return receipt
}

func (p *Contract) Deployed(t *testing.T, client *Client, ca common.Address) *Contract {
	p.Client = client
	p.Address = ca

	collectContract(p.Address, p.Alias)
	return p
}

func (p *Contract) PackInput(t *testing.T, method string, args ...interface{}) (data []byte) {
	var err error
	defer func() {
		require.NoError(t, err)
	}()

	data, err = p.Abi.Pack(method, args...)
	if err == nil {
		return
	}
	if p.Logic != nil {
		data, err = p.Logic.Abi.Pack(method, args...)
	}
	return
}

func (p *Contract) UnpackOutput(method string, output []byte) (ret []interface{}, err error) {
	if m, ok := p.Abi.Methods[method]; ok {
		ret, err = m.Outputs.UnpackValues(output)
	} else if m, ok := p.Logic.Abi.Methods[method]; ok {
		ret, err = m.Outputs.UnpackValues(output)
	} else {
		err = fmt.Errorf("not found method %v", method)
	}
	return
}

// LowCall returns method's output in a different way than Call.
func (p *Contract) UnpackCall(t *testing.T, ret interface{}, method string, args ...interface{}) {
	input := p.PackInput(t, method, args...)

	msg := ethereum.CallMsg{From: common.Address{}, To: &p.Address, Data: input}
	output, err := p.Client.Backend.CallContract(ctx, msg, nil)

	require.NoErrorf(t, err, "call contract: %v", method)

	if _, ok := p.Abi.Methods[method]; ok {
		err = p.Abi.UnpackIntoInterface(ret, method, output)
	} else if _, ok := p.Logic.Abi.Methods[method]; ok {
		err = p.Logic.Abi.UnpackIntoInterface(ret, method, output)
	} else {
		err = fmt.Errorf("not found method %v", method)
	}
	require.NoError(t, err)

}

// LowCall returns method's output in a different way than Call.
func (p *Contract) Call(t *testing.T, method string, args ...interface{}) (ret []interface{}) {
	input := p.PackInput(t, method, args...)

	msg := ethereum.CallMsg{From: common.Address{}, To: &p.Address, Data: input}
	output, err := p.Client.Backend.CallContract(ctx, msg, nil)

	require.NoErrorf(t, err, "call contract: %v", method)

	ret, err = p.UnpackOutput(method, output)
	assert.NoError(t, err, "unpack values:%v", method)

	return
}

// LowCall1 returns method's output only not single returns
func (p *Contract) Call1(t *testing.T, method string, args ...interface{}) interface{} {
	ret := p.Call(t, method, args...)
	require.Truef(t, len(ret) > 0, "mismatch return size:%v", len(ret))
	return ret[0]
}

// FormatCall return output according to the abi specification
// struct 타입으로 받을경우 사용하면 편리
// 호춣할때 out 변수를 위에 선언하지않고 한줄로 사용할수있도록 output의 type을 인자로 받게 했음.
func (p *Contract) FormatCall(t *testing.T, outType reflect.Type, method string, args ...interface{}) (output interface{}) {
	input := p.PackInput(t, method, args...)

	msg := ethereum.CallMsg{From: common.Address{}, To: &p.Address, Data: input}
	data, err := p.Client.Backend.CallContract(ctx, msg, nil)

	require.NoErrorf(t, err, "call contract: %v", method)

	output = reflect.New(outType).Interface()
	if err = p.Abi.UnpackIntoInterface(&output, method, data); err != nil {
		err = p.Abi.UnpackIntoInterface(output, method, data)
	}
	if err != nil && p.Logic != nil {
		if err = p.Logic.Abi.UnpackIntoInterface(&output, method, data); err != nil {
			err = p.Logic.Abi.UnpackIntoInterface(output, method, data)
		}
	}
	assert.NoError(t, err, "unpack values:%v", method)

	return
}

// ExecuteOk is that must be successful.
func (p *Contract) ExecuteOk(t *testing.T, sender common.Address, method string, args ...interface{}) (receipt *types.Receipt) {
	return p.ExecuteWithETHOk(t, sender, big.NewInt(0), method, args...)
}

// ExecuteFail is that must be failed.
func (p *Contract) ExecuteFail(t *testing.T, sender common.Address, method string, args ...interface{}) (err error) {
	return p.ExecuteWithETHFail(t, sender, big.NewInt(0), method, args...)
}

func (p *Contract) ExecuteWithETHOk(t *testing.T, sender common.Address, value *big.Int, method string, args ...interface{}) (receipt *types.Receipt) {
	var err error
	receipt, err = p.execute(t, sender, value, method, args...)
	require.Truef(t, receipt.Status == types.ReceiptStatusSuccessful, "execute error(%v): contract:%v, method:%v, gas:%v\n%v", receipt.Status, p.Alias, method, receipt.GasUsed, err)
	return
}

func (p *Contract) ExecuteWithETHFail(t *testing.T, sender common.Address, value *big.Int, method string, args ...interface{}) (err error) {
	var receipt *types.Receipt
	receipt, err = p.execute(t, sender, value, method, args...)
	require.False(t, receipt.Status == types.ReceiptStatusSuccessful, fmt.Sprintf("execute not error(%v): contract:%v, method:%v, gas:%v", receipt.Status, p.Alias, method, receipt.GasUsed))
	return
}

// ExecuteWithKlay includes value in the execution transaction.
func (p *Contract) execute(t *testing.T, sender common.Address, value *big.Int, method string, args ...interface{}) (receipt *types.Receipt, err error) {
	var tx *types.Transaction
	if sender == (common.Address{}) {
		sender = p.Client.Owner
	}

	tx, receipt = p.Client.SendTransaction(t, sender, p.Address, value, p.PackInput(t, method, args...))
	if receipt.Status != types.ReceiptStatusSuccessful {
		err = CallRevertReason(t, p.Client.Backend, sender, receipt.BlockNumber, tx)
	} else {
		// proxy의 logic 업데이트경우 처리.
		if p.Logic != nil {
			if e, err := p.Abi.EventByID(crypto.Keccak256Hash([]byte("Upgraded(address)"))); err == nil {
			search:
				for _, g := range receipt.Logs {
					if g.Topics[0] == e.ID {
						newLogic := p.Client.FindContractByAddress(common.BytesToAddress(g.Topics[1][:]))
						require.NotNil(t, newLogic)
						p.Logic = newLogic
						break search
					}
				}
			}
		}
	}
	return
}

// CallExpectedReverted is expects an 'evm: execution reverted' error to occur.
func (p *Contract) CallExpectedReverted(t *testing.T, method string, args ...interface{}) error {
	input := p.PackInput(t, method, args...)

	msg := ethereum.CallMsg{From: common.Address{}, To: &p.Address, Data: input}
	_, err := p.Client.Backend.CallContract(ctx, msg, nil)
	require.Error(t, err)
	return UnpackError(err)
}

// RevertReason returns the reason when 'evm: execution reverted' occurs in the contract method call.
func CallRevertReason(t *testing.T, client iCallContract, from common.Address, blockNumber *big.Int, tx *types.Transaction) error {
	msg := ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Data:     tx.Data(),
		Gas:      defaultGasLimit,
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
	}
	_, err := client.CallContract(ctx, msg, blockNumber)
	require.Error(t, err)
	return UnpackError(err)
}

type iCallContract interface {
	CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
}
