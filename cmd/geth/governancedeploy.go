// governancedeploy.js

package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/wemix/metclient"
	cli "github.com/urfave/cli/v2"
)

func getInitialGovernanceMembersAndNodes(configJsFile string) (nodes []byte, stakes []byte, rewardPoolAccount, maintenanceAccount *common.Address, err error) {
	var fin *os.File
	if fin, err = os.Open(configJsFile); err != nil {
		return
	}
	defer fin.Close()

	var cfg *genesisConfig
	cfg, err = loadGenesisConfig(fin)
	if err != nil {
		return
	}

	l := len(cfg.Members)
	var b1, b2 bytes.Buffer
	for i := 0; i < l; i++ {
		m := cfg.Members[i]
		var (
			sid string
			id  []byte
		)
		if len(m.Id) == 128 {
			sid = m.Id
		} else if len(m.Id) == 130 {
			sid = m.Id[2:]
		} else {
			return nil, nil, nil, nil, fmt.Errorf("Invalid enode id %s", m.Id)
		}
		if id, err = hex.DecodeString(sid); err != nil {
			return nil, nil, nil, nil, err
		}

		addr := new(big.Int).SetBytes(m.Addr[:])
		b1.Write(metclient.PackNum(reflect.ValueOf(addr)))
		b1.Write(metclient.PackNum(reflect.ValueOf(len(m.Name))))
		b1.Write([]byte(m.Name))
		b1.Write(metclient.PackNum(reflect.ValueOf(len(id))))
		b1.Write(id)
		b1.Write(metclient.PackNum(reflect.ValueOf(len(m.Ip))))
		b1.Write([]byte(m.Ip))
		b1.Write(metclient.PackNum(reflect.ValueOf(m.Port)))

		b2.Write(metclient.PackNum(reflect.ValueOf(addr)))
		b2.Write(metclient.PackNum(reflect.ValueOf(m.Stake)))
	}
	nodes = b1.Bytes()
	stakes = b2.Bytes()
	nilAddr := common.Address{}
	if cfg.RewardPool != nilAddr {
		rewardPoolAccount = &cfg.RewardPool
	}
	if cfg.Maintenance != nilAddr {
		maintenanceAccount = &cfg.Maintenance
	}

	return
}

// governance-contract.js config.js
func deployGovernanceContracts(cliCtx *cli.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error

	// get command line arguments
	url := cliCtx.String(urlFlag.Name)
	gas := cliCtx.Int(gasFlag.Name)
	gasPrice := cliCtx.Int(gasPriceFlag.Name)

	if gas <= 0 {
		gas = 0xF000000
	}
	if gasPrice <= 0 {
		gasPrice = 80000000000
	}

	if len(url) == 0 || cliCtx.NArg() != 3 {
		return fmt.Errorf("Invalid Arguments")
	}

	passwd := utils.GetPassPhraseWithList("", false, 0, utils.MakePasswordList(cliCtx))
	if len(passwd) == 0 {
		return fmt.Errorf("Invalid Arguments")
	}

	args := cliCtx.Args().Slice()
	contractsFile, configJsFile, accountFile := args[0], args[1], args[2]

	// account
	var from *keystore.Key
	from, err = metclient.LoadAccount(passwd, accountFile)
	if err != nil {
		return err
	}

	// initial members and nodes data
	var (
		membersAndNodes, stakes               []byte
		rewardPoolAccount, maintenanceAccount *common.Address
	)
	membersAndNodes, stakes, rewardPoolAccount, maintenanceAccount, err = getInitialGovernanceMembersAndNodes(configJsFile)
	if err != nil {
		return nil
	}

	// cli connection
	var cli *ethclient.Client
	cli, err = ethclient.Dial(url)
	if err != nil {
		return err
	}

	// contract variables
	registry := &metclient.RemoteContract{Cli: cli, From: from, Gas: gas}
	envStorageImp := &metclient.RemoteContract{Cli: cli, From: from, Gas: gas}
	staking := &metclient.RemoteContract{Cli: cli, From: from, Gas: gas}
	ballotStorage := &metclient.RemoteContract{Cli: cli, From: from, Gas: gas}
	envStorage := &metclient.RemoteContract{Cli: cli, From: from, Gas: gas}
	govImp := &metclient.RemoteContract{Cli: cli, From: from, Gas: gas}
	gov := &metclient.RemoteContract{Cli: cli, From: from, Gas: gas}

	// load contract codes
	var contracts map[string]*metclient.ContractData
	if fin, e2 := os.Open(contractsFile); e2 != nil {
		return e2
	} else {
		defer fin.Close()
		if contracts, err = metclient.LoadJsContract(fin); err != nil {
			return nil
		}
	}

	// check if contracts exist
	contractNames := []string{"Registry", "EnvStorageImp", "Staking", "BallotStorage", "EnvStorage", "GovImp", "Gov"}
	for _, name := range contractNames {
		if _, ok := contracts[name]; !ok {
			return fmt.Errorf("Cannot find %s contract", name)
		}
	}
	registry.Abi = contracts["Registry"].Abi
	envStorageImp.Abi = contracts["EnvStorageImp"].Abi
	staking.Abi = contracts["Staking"].Abi
	ballotStorage.Abi = contracts["BallotStorage"].Abi
	envStorage.Abi = contracts["EnvStorage"].Abi
	govImp.Abi = contracts["GovImp"].Abi
	gov.Abi = contracts["Gov"].Abi

	txs := make([]common.Hash, 10)

	// 1. deploy Registry and EnvStorageImp contracts
	ixTxs := 0
	fmt.Println("Deploying Registry...")
	if txs[ixTxs], err = metclient.Deploy(ctx, cli, from, contracts["Registry"], nil, gas, gasPrice); err != nil {
		return err
	}
	ixTxs++
	fmt.Println("Deploying EnvStorageImp...")
	if txs[ixTxs], err = metclient.Deploy(ctx, cli, from, contracts["EnvStorageImp"], nil, gas, gasPrice); err != nil {
		return err
	}

	fmt.Print("Waiting for receipts...")
	for i := ixTxs; i >= 0; i-- {
		var receipt *types.Receipt
		receipt, err = metclient.GetContractReceipt(ctx, cli, txs[i], 200, 300)
		if err != nil {
			return err
		}
		switch i {
		case 0:
			registry.To = &receipt.ContractAddress
		case 1:
			envStorageImp.To = &receipt.ContractAddress
		}
	}
	fmt.Println("good.")

	// 2. deploy Staking, BallotStorage, EnvStorage, GovImp, Gov
	ixTxs = 0
	fmt.Println("Deploying Staking...")
	if txs[ixTxs], err = metclient.Deploy(ctx, cli, from, contracts["Staking"],
		[]interface{}{registry.To, stakes}, gas, gasPrice); err != nil {
		return err
	}
	ixTxs++
	fmt.Println("Deploying BalloStorage...")
	if txs[ixTxs], err = metclient.Deploy(ctx, cli, from, contracts["BallotStorage"],
		[]interface{}{registry.To}, gas, gasPrice); err != nil {
		return err
	}
	ixTxs++
	fmt.Println("Deploying EnvStorage...")
	if txs[ixTxs], err = metclient.Deploy(ctx, cli, from, contracts["EnvStorage"],
		[]interface{}{registry.To, envStorageImp.To}, gas, gasPrice); err != nil {
		return err
	}
	ixTxs++
	fmt.Println("Deploying GovImp...")
	if txs[ixTxs], err = metclient.Deploy(ctx, cli, from, contracts["GovImp"], nil, gas, gasPrice); err != nil {
		return err
	}
	ixTxs++
	fmt.Println("Deploying Gov...")
	if txs[ixTxs], err = metclient.Deploy(ctx, cli, from, contracts["Gov"],
		nil, gas, gasPrice); err != nil {
		return err
	}
	fmt.Println("Gov tx is", txs[ixTxs].Hex())

	fmt.Printf("Waiting for receipts...")
	for i := ixTxs; i >= 0; i-- {
		var receipt *types.Receipt
		receipt, err = metclient.GetContractReceipt(ctx, cli, txs[i], 200, 300)
		if err != nil {
			return err
		}
		switch i {
		case 0:
			staking.To = &receipt.ContractAddress
		case 1:
			ballotStorage.To = &receipt.ContractAddress
		case 2:
			envStorage.To = &receipt.ContractAddress
		case 3:
			govImp.To = &receipt.ContractAddress
		case 4:
			gov.To = &receipt.ContractAddress
		}
	}
	fmt.Printf("good. Governance address %v.\n", gov.To.Hex())

	// 3. setup registry
	fmt.Println("Setting registry...")
	ixTxs = 0
	if txs[ixTxs], err = metclient.SendContract(ctx, registry, "setContractDomain", []interface{}{metclient.ToBytes32("Staking"), staking.To}); err != nil {
		return err
	}
	ixTxs++
	if txs[ixTxs], err = metclient.SendContract(ctx, registry, "setContractDomain", []interface{}{metclient.ToBytes32("BallotStorage"), ballotStorage.To}); err != nil {
		return err
	}
	ixTxs++
	if txs[ixTxs], err = metclient.SendContract(ctx, registry, "setContractDomain", []interface{}{metclient.ToBytes32("EnvStorage"), envStorage.To}); err != nil {
		return err
	}
	ixTxs++
	if txs[ixTxs], err = metclient.SendContract(ctx, registry, "setContractDomain", []interface{}{metclient.ToBytes32("GovernanceContract"), gov.To}); err != nil {
		return err
	}
	if rewardPoolAccount != nil {
		ixTxs++
		if txs[ixTxs], err = metclient.SendContract(ctx, registry, "setContractDomain", []interface{}{metclient.ToBytes32("RewardPool"), rewardPoolAccount}); err != nil {
			return err
		}
	}
	if maintenanceAccount != nil {
		ixTxs++
		if txs[ixTxs], err = metclient.SendContract(ctx, registry, "setContractDomain", []interface{}{metclient.ToBytes32("Maintenance"), maintenanceAccount}); err != nil {
			return err
		}
	}

	// no need to wait for the receipts for the above

	// 4. deposit staking - not needed
	fmt.Println("Depositing stakes...")

	// 5. Gov.initOnce()
	fmt.Printf("Initializing governance members and nodes...")
	if txs[0], err = metclient.SendContract(ctx, gov, "initOnce", []interface{}{registry.To, govImp.To, membersAndNodes}); err != nil {
		return err
	}

	if receipt, err2 := metclient.GetReceipt(ctx, cli, txs[0], 200, 300); err2 != nil {
		return err2
	} else if receipt.Status != 1 {
		fmt.Printf("Transaction %v failed with status %d.\n",
			txs[0].Hex(), receipt.Status)
		return fmt.Errorf("Transaction failed with status %d.", receipt.Status)
	}
	fmt.Println("good.")

	// 6. initialize environment storage data:
	// blocksPer, ballotDurationMin, ballotDurationMax, stakingMin, stakingMax,
	// gasPrice
	defaultBlocksPer := big.NewInt(100)
	defaultBallotDurationMin := big.NewInt(86400)
	defaultBallotDurationMax := big.NewInt(604800)
	defaultStakingMin, _ := big.NewInt(0).SetString("4980000000000000000000000", 0)
	defaultStakingMax, _ := big.NewInt(0).SetString("39840000000000000000000000", 0)
	defaultGasPrice := big.NewInt(80000000000)
	defaultMaxIdleBlockInterval := big.NewInt(5)
	envDefaults := []interface{}{
		defaultBlocksPer,
		defaultBallotDurationMin,
		defaultBallotDurationMax,
		defaultStakingMin,
		defaultStakingMax,
		defaultGasPrice,
		defaultMaxIdleBlockInterval,
	}
	fmt.Printf("Initializing environment storage.\n")
	envStorageImp.To = envStorage.To
	if txs[0], err = metclient.SendContract(ctx, envStorageImp, "initialize", envDefaults); err != nil {
		return err
	}

	// 7. print the addresses
	fmt.Printf(`{
  "REGISTRY_ADDRESS": "%s",
  "STAKING_ADDRESS": "%s",
  "ENV_STORAGE_ADDRESS": "%s",
  "BALLOT_STORAGE_ADDRESS": "%s",
  "GOV_ADDRESS": "%s",
  "GOV_IMP_ADDRESS": "%s"
}
`,
		registry.To.Hex(), staking.To.Hex(), envStorage.To.Hex(),
		ballotStorage.To.Hex(), gov.To.Hex(), govImp.To.Hex())

	return nil
}

// EOF
