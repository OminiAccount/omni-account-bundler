package ethereum

import (
	"context"
	"fmt"
	"github.com/OAB/ethereum/contracts/EntryPoint"
	"github.com/OAB/ethereum/contracts/SimpleAccountFactory"
	"github.com/OAB/state/types"
	"github.com/OAB/utils/chains"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethereumType "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"strings"
)

type Client struct {
	EthClient  ethereumClient
	EntryPoint *EntryPoint.EntryPoint
}

type Ethereum struct {
	ChainsClient   map[chains.ChainId]*Client
	AccountFactory *SimpleAccountFactory.SimpleAccountFactory // We only need to listen to the aa contract of the main chain.
	auth           map[common.Address]bind.TransactOpts

	sender  common.Address
	chainId uint64

	logger log.Logger
}

// NewEthereum create a EthereumClient that support multiple chains
func NewEthereum(cfg Config) (*Ethereum, error) {
	ethereum := &Ethereum{
		ChainsClient: make(map[chains.ChainId]*Client, chains.MaxChainInfoLength),
		auth:         map[common.Address]bind.TransactOpts{},
		logger:       log.New("service", "ethereum"),
	}
	ethereum.chainId = cfg.Networks[0].ChainId

	for _, network := range cfg.Networks {
		chainId := chains.ChainId(network.ChainId)
		client, err := Dial(network.Rpc)
		if err != nil {
			return nil, err
		}

		wsClient, err := Dial(strings.Replace(network.Rpc, "https://", "wss://", 1))
		if err != nil {
			return nil, fmt.Errorf("failed to init wsClient: %w", err)
		}

		entryPoint, err := EntryPoint.NewEntryPoint(network.EntryPoint, wsClient)
		if err != nil {
			return nil, err
		}

		if ethereum.chainId == network.ChainId {
			if ethereum.AccountFactory, err = SimpleAccountFactory.NewSimpleAccountFactory(cfg.AccountFactory, wsClient); err != nil {
				return nil, err
			}
		}

		ethereum.ChainsClient[chainId] = &Client{
			EthClient:  client,
			EntryPoint: entryPoint,
		}
	}

	for _, key := range cfg.PrivateKeys {
		opts, _, err := ethereum.LoadAuthFromKeyStore(key.Path, key.Password, ethereum.chainId)
		if err != nil {
			return nil, err
		}
		ethereum.sender = opts.From
	}

	return ethereum, nil
}

func (ether *Ethereum) UpdateEntryPointRoot(proof hexutil.Bytes, pubicValues hexutil.Bytes) (*ethereumType.Transaction, error) {
	opts, err := ether.getAuthByAddress(ether.sender)
	if err != nil {
		return nil, err
	}

	// force nonce, gas limit and gas price to avoid querying it from the chain
	//opts.Nonce = big.NewInt(1)
	//opts.GasLimit = uint64(47860)
	//opts.GasPrice = big.NewInt(1)
	opts.Value = big.NewInt(9000000000000000)

	entryPoint := ether.ChainsClient[chains.ChainId(ether.chainId)].EntryPoint

	tx, err := entryPoint.VerifyBatch(&opts, proof, pubicValues, opts.From)

	if err != nil {
		return nil, err
	}
	fmt.Println("tx", tx.Hash())

	ether.logger.Info("update entryPoint root", "hash", tx.Hash().String())
	return tx, nil
}

func (ether *Ethereum) WatchAAFactoryEvent(ctx context.Context, fromBlock uint64, accountMappingChannel chan<- types.AccountMapping) error {
	opts := &bind.WatchOpts{
		Start:   &fromBlock,
		Context: ctx,
	}

	accountCreatedChannel := make(chan *SimpleAccountFactory.SimpleAccountFactoryAccountCreated)
	accountCreatedSubscription, err := ether.AccountFactory.WatchAccountCreated(opts, accountCreatedChannel, nil)
	if err != nil {
		return err
	}
	defer accountCreatedSubscription.Unsubscribe()

	for {
		select {
		case err = <-accountCreatedSubscription.Err():
			ether.logger.Crit("Failed to subscribe to accountCreated", "error", err)
		case event := <-accountCreatedChannel:
			ether.logger.Debug("Received accountCreated event", "event", event)
			mapping := types.AccountMapping{
				User:    event.Owner,
				Account: event.Account,
			}
			accountMappingChannel <- mapping
		case <-ctx.Done():
			return nil
		}
	}
}
