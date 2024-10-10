// governancedeploy.js

package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	gov "github.com/ethereum/go-ethereum/wemix/bind"
	"github.com/ethereum/go-ethereum/wemix/metclient"
	"gopkg.in/urfave/cli.v1"
)

func getInitialGovernanceInitDatas(configJsFile string) (
	domains map[string]common.Address,
	env gov.InitEnvStorage,
	members gov.InitMembers,
	err error,
) {
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

	domains = make(map[string]common.Address)
	zeroAddress := common.Address{}
	if cfg.StakingReward != zeroAddress {
		domains[gov.DOMAIN_StakingReward] = cfg.StakingReward
	}
	if cfg.Ecosystem != zeroAddress {
		domains[gov.DOMAIN_Ecosystem] = cfg.Ecosystem
	}
	if cfg.Maintenance != zeroAddress {
		domains[gov.DOMAIN_Maintenance] = cfg.Maintenance
	}
	if cfg.FeeCollector != zeroAddress {
		domains[gov.DOMAIN_FeeCollector] = cfg.FeeCollector
	}

	env = cfg.Env.ToInitData()

	members = make(gov.InitMembers, len(cfg.Members))
	for i, member := range cfg.Members {
		members[i] = member.ToInitData()
	}

	return
}

// governance-contract.js config.js
func deployGovernanceContracts(ctx *cli.Context) error {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stdout, log.TerminalFormat(true))))
	// get command line arguments
	args := ctx.Args()

	var (
		errInvalidArguments     error = errors.New("invalid Arguments")
		configFile, accountFile string
		lockAmount              *big.Int
	)
	switch len(args) {
	case 3:
		configFile, accountFile = args[0], args[1]
		lockAmount, ok := new(big.Int).SetString(args[2], 10)
		if !ok || lockAmount.Sign() <= 0 {
			return errInvalidArguments
		}
	case 2:
		configFile, accountFile = args[0], args[1]
		lockAmount = gov.DefaultInitEnvStorage.STAKING_MIN
	default:
		return errInvalidArguments
	}

	client, err := ethclient.Dial(ctx.String(urlFlag.Name))
	if err != nil {
		return err
	}
	chainID, err := client.ChainID(context.TODO())
	if err != nil {
		return err
	}

	var opts *bind.TransactOpts
	if from, err := metclient.LoadAccount(utils.GetPassPhraseWithList("", false, 0, utils.MakePasswordList(ctx)), accountFile); err != nil {
		return err
	} else if opts, err = bind.NewKeyedTransactorWithChainID(from.PrivateKey, chainID); err != nil {
		return err
	} else {
		if ctx.IsSet(gasFlag.Name) {
			opts.GasLimit = ctx.Uint64(gasFlag.Name)
		}
		if ctx.IsSet(gasPriceFlag.Name) {
			opts.GasPrice = big.NewInt(ctx.Int64(gasPriceFlag.Name))
		}
	}

	return deployGovernance(client, opts, lockAmount, configFile)
}

type Client interface {
	bind.ContractBackend
	bind.DeployBackend
}

func deployGovernance(cli Client, ownerOpts *bind.TransactOpts, lockAmount *big.Int, configJsFile string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ownerOpts.Context = ctx

	// initial members and nodes data
	domains, env, members, err := getInitialGovernanceInitDatas(configJsFile)
	if err != nil {
		return err
	}
	if lockAmount.Cmp(env.STAKING_MIN) < 0 || lockAmount.Cmp(env.STAKING_MAX) > 0 {
		return fmt.Errorf("invalid lock amount, input:%v, min:%v, max:%v", lockAmount, env.STAKING_MIN, env.STAKING_MAX)
	}

	// deploy & set domains
	contracts, err := gov.DeployGovContracts(ownerOpts, cli, domains)
	if err != nil {
		return err
	}
	// init contracts
	err = gov.ExecuteInitialize(contracts, ownerOpts, cli, lockAmount, env, members)
	if err != nil {
		return err
	}

	// print the addresses
	address := contracts.Address()
	fmt.Printf(`{
  "REGISTRY_ADDRESS": "%s",
  "STAKING_ADDRESS": "%s",
  "ENV_STORAGE_ADDRESS": "%s",
  "BALLOT_STORAGE_ADDRESS": "%s",
  "GOV_ADDRESS": "%s"
}
`,
		address.Registry, address.Staking, address.EnvStorage, address.BallotStorage, address.Gov)

	return nil
}

// EOF
