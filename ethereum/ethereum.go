package ethereum

import (
	"context"
	"fmt"
	EntryPoint "github.com/OAAC/ethereum/contracts"
	"github.com/OAAC/pool"
	"github.com/OAAC/utils/chains"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"strings"
)

type Client struct {
	EthClient  ethereumClient
	EntryPoint *EntryPoint.EntryPoint
}

type Ethereum struct {
	ChainsClient map[chains.ChainId]*Client
	auth         map[common.Address]bind.TransactOpts

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

		ethereum.ChainsClient[chainId] = &Client{
			EthClient:  client,
			EntryPoint: entryPoint,
		}
	}

	ethereum.chainId = cfg.Networks[0].ChainId

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
	depositTicketSubscription, err := entryPoint.WatchDepositTicketAdded(opts, depositTicketAddedChannel, nil, nil)
	if err != nil {
		return err
	}
	defer depositTicketSubscription.Unsubscribe()

	withdrawTicketSubscription, err := entryPoint.WatchWithdrawTicketAdded(opts, withdrawTicketAddedChannel, nil, nil)
	if err != nil {
		return err
	}
	defer withdrawTicketSubscription.Unsubscribe()

	for {
		select {
		case err = <-depositTicketSubscription.Err():
			ether.logger.Crit("Failed to subscribe to depositTicket", "error", err)
		case err = <-withdrawTicketSubscription.Err():
			ether.logger.Crit("Failed to subscribe to withdrawTicket", "error", err)
		case event1 := <-depositTicketAddedChannel:
			ether.logger.Debug("Received deposit ticket event", "event", event1)
			ticket := pool.TicketFull{
				Ticket: pool.Ticket{
					User:      event1.User,
					Amount:    (*hexutil.Big)(event1.Amount),
					TimeStamp: event1.Timestamp.Uint64(),
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
					TimeStamp: event2.Timestamp.Uint64(),
				},
				Type: pool.Withdraw,
			}
			ticketChannel <- ticket
		case <-ctx.Done():
			return nil
		}
	}
}

func (ether *Ethereum) UpdateEntryPointRoot(proof hexutil.Bytes, pubicValues hexutil.Bytes) error {
	_, err := ether.getAuthByAddress(ether.sender)
	if err != nil {
		return err
	}

	entryPoint := ether.ChainsClient[chains.ChainId(ether.chainId)].EntryPoint

	//tx, err := entryPoint.ZkAAEntryPoint(&opts, proof, [32]byte(pubInput), opts.From)
	//if err != nil {
	//	return err
	//}

	ether.logger.Info("update entryPoint root", "entryPoint", entryPoint)
	return nil
}
