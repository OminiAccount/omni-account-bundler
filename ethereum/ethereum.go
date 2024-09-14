package ethereum

import (
	"context"
	"fmt"
	"github.com/OAB/ethereum/contracts/EntryPoint"
	"github.com/OAB/ethereum/contracts/SimpleAccountFactory"
	"github.com/OAB/pool"
	"github.com/OAB/state/types"
	"github.com/OAB/utils/chains"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethereumType "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"strings"
	"time"
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

func (ether *Ethereum) WatchEntryPointEvent(ctx context.Context, chainId chains.ChainId, fromBlock uint64, ticketChannel chan<- pool.TicketFull) error {

	entryPoint := ether.ChainsClient[chainId].EntryPoint
	opts := &bind.WatchOpts{
		Start:   &fromBlock,
		Context: ctx,
	}
	depositTicketAddedChannel := make(chan *EntryPoint.EntryPointDepositTicketAdded)
	withdrawTicketAddedChannel := make(chan *EntryPoint.EntryPointWithdrawTicketAdded)

	// Encapsulating function to start a subscription
	subscribe := func() (ethereum.Subscription, ethereum.Subscription, error) {
		depositTicketSubscription, err := entryPoint.WatchDepositTicketAdded(opts, depositTicketAddedChannel, nil, nil)
		if err != nil {
			return nil, nil, err
		}

		withdrawTicketSubscription, err := entryPoint.WatchWithdrawTicketAdded(opts, withdrawTicketAddedChannel, nil, nil)
		if err != nil {
			depositTicketSubscription.Unsubscribe()
			return nil, nil, err
		}

		return depositTicketSubscription, withdrawTicketSubscription, nil
	}

	// Initial Subscription
	depositTicketSubscription, withdrawTicketSubscription, err := subscribe()
	if err != nil {
		return err
	}
	defer func() {
		depositTicketSubscription.Unsubscribe()
		withdrawTicketSubscription.Unsubscribe()
	}()

	for {
		select {
		case err = <-depositTicketSubscription.Err():
			ether.logger.Error("Failed to subscribe to depositTicket", "error", err)

			// Re-subscribe
			depositTicketSubscription.Unsubscribe()
			withdrawTicketSubscription.Unsubscribe()
			for {
				depositTicketSubscription, withdrawTicketSubscription, err = subscribe()
				if err == nil {
					break
				}
				ether.logger.Error("Retrying subscription after failure", "error", err)
				time.Sleep(time.Second * 1)
			}
		case err = <-withdrawTicketSubscription.Err():
			ether.logger.Error("Failed to subscribe to withdrawTicket", "error", err)

			depositTicketSubscription.Unsubscribe()
			withdrawTicketSubscription.Unsubscribe()
			for {
				depositTicketSubscription, withdrawTicketSubscription, err = subscribe()
				if err == nil {
					break
				}
				ether.logger.Error("Retrying subscription after failure", "error", err)
				time.Sleep(time.Second * 1)
			}
		case event1 := <-depositTicketAddedChannel:
			ether.logger.Debug("Received deposit ticket event", "event", event1)
			ticket := pool.TicketFull{
				Ticket: pool.Ticket{
					User:      event1.User,
					Amount:    (*hexutil.Big)(event1.Amount),
					TimeStamp: (*hexutil.Big)(event1.Timestamp),
				},
				Type: pool.Deposit,
			}
			ticketChannel <- ticket
		case event2 := <-withdrawTicketAddedChannel:
			ether.logger.Debug("Received withdraw ticket event", "event", event2)
			ticket := pool.TicketFull{
				Ticket: pool.Ticket{
					User:      event2.User,
					Amount:    (*hexutil.Big)(event2.Amount),
					TimeStamp: (*hexutil.Big)(event2.Timestamp),
				},
				Type: pool.Withdraw,
			}
			ticketChannel <- ticket
		case <-ctx.Done():
			return nil
		}
	}
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
