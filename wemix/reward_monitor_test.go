package wemix

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"
)

type Payload struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type TxRes struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	To               string `json:"to"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	TransactionIndex string `json:"transactionIndex"`
	Type             string `json:"type"`
	Value            string `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

type BlockRes struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BaseFeePerGas    string   `json:"baseFeePerGas"`
		Difficulty       string   `json:"difficulty"`
		ExtraData        string   `json:"extraData"`
		GasLimit         string   `json:"gasLimit"`
		GasUsed          string   `json:"gasUsed"`
		Hash             string   `json:"hash"`
		LogsBloom        string   `json:"logsBloom"`
		Miner            string   `json:"miner"`
		MixHash          string   `json:"mixHash"`
		Nonce            string   `json:"nonce"`
		Number           string   `json:"number"`
		Rewards          string   `json:"rewards"`
		ParentHash       string   `json:"parentHash"`
		ReceiptsRoot     string   `json:"receiptsRoot"`
		Sha3Uncles       string   `json:"sha3Uncles"`
		Size             string   `json:"size"`
		StateRoot        string   `json:"stateRoot"`
		Timestamp        string   `json:"timestamp"`
		TotalDifficulty  string   `json:"totalDifficulty"`
		Transactions     []TxRes  `json:"transactions"`
		TransactionsRoot string   `json:"transactionsRoot"`
		Uncles           []string `json:"uncles"`
	} `json:"result"`
}

func GetBlock() (*BlockRes, error) {
	data := Payload{
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{"latest", true},
		ID:      1,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://172.207.48.146:8588", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	var result BlockRes
	_ = json.Unmarshal(bytes, &result)

	return &result, nil
}

func TestBlockStatistics(t *testing.T) {
	for {
		block, err := GetBlock()
		if err != nil {
			t.Errorf("error get block: %s", err)
			break
		}
		num := new(big.Int)
		num.SetString(strings.TrimPrefix(block.Result.Number, "0x"), 16)

		data, _ := hex.DecodeString(strings.TrimPrefix(block.Result.Rewards, "0x"))
		var reward []reward
		if err = json.Unmarshal(data, &reward); err != nil {
			t.Errorf(err.Error())
		}
		t.Logf("Number = %v, Rewards = %v", num, reward)

		time.Sleep(time.Second)
	}
}
