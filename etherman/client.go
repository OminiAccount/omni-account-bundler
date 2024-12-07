package etherman

import (
	"context"
	"errors"
	"fmt"
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/etherman/contracts/SimpleAccountFactory"
	"github.com/OAB/etherman/contracts/SyncRouter"
	"github.com/OAB/utils/chains"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

var (
	accountCreatedSignatureHash      = crypto.Keccak256Hash([]byte("AccountCreated(address,address)"))
	depositTicketAddedSignatureHash  = crypto.Keccak256Hash([]byte("DepositTicketAdded(bytes32,address,uint256,uint256)"))
	withdrawTicketAddedSignatureHash = crypto.Keccak256Hash([]byte("WithdrawTicketAdded(address,uint256,uint256)"))

	ErrNotFound = errors.New("Not found")
)

type ethereumClient struct {
	cfg            Network
	chainID        chains.ChainId
	ethClient      ethClientInterface
	entryPointAddr common.Address
	entryPoint     *EntryPoint.EntryPoint
	aaFactoryAddr  common.Address
	aaFactory      *SimpleAccountFactory.SimpleAccountFactory
	syncRouterAddr common.Address
	syncRouter     *SyncRouter.SyncRouter

	sender common.Address
	auth   bind.TransactOpts
}

func NewClient(cfg Network) (*ethereumClient, error) {
	ethClient, err := ethclient.Dial(cfg.Rpc)
	if err != nil {
		log.Errorf("chainID: %d, error connecting to %s: %+v", cfg.ChainId, cfg.Rpc, err)
		return nil, err
	}

	entryPoint, err := EntryPoint.NewEntryPoint(cfg.EntryPoint, ethClient)
	if err != nil {
		return nil, fmt.Errorf("chainID: %d, failed to init entryPoint: %v", cfg.ChainId, err)
	}

	accountFactory, err := SimpleAccountFactory.NewSimpleAccountFactory(cfg.AccountFactory, ethClient)
	if err != nil {
		return nil, fmt.Errorf("chainID: %d, failed to init simple account factory: %v", cfg.ChainId, err)
	}

	syncRouter, err := SyncRouter.NewSyncRouter(cfg.SyncRouter, ethClient)
	if err != nil {
		return nil, fmt.Errorf("chainID: %d, failed to init sync router: %v", cfg.ChainId, err)
	}

	return &ethereumClient{
		cfg:            cfg,
		chainID:        chains.ChainId(cfg.ChainId),
		entryPointAddr: cfg.EntryPoint,
		ethClient:      ethClient,
		entryPoint:     entryPoint,
		aaFactoryAddr:  cfg.AccountFactory,
		aaFactory:      accountFactory,
		syncRouterAddr: cfg.SyncRouter,
		syncRouter:     syncRouter,
	}, nil
}

func (c *ethereumClient) GetEventByBlockRange(ctx context.Context, fromBlock uint64, toBlock *uint64) ([]Block, error) {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		Addresses: []common.Address{c.entryPointAddr, c.aaFactoryAddr},
	}
	if toBlock != nil {
		query.ToBlock = new(big.Int).SetUint64(*toBlock)
	}
	blocks, err := c.readEvents(ctx, query)
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

func (c *ethereumClient) readEvents(ctx context.Context, query ethereum.FilterQuery) ([]Block, error) {
	logs, err := c.ethClient.FilterLogs(ctx, query)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var blocks []Block
	for _, vLog := range logs {
		err := c.processEvent(ctx, vLog, &blocks)
		if err != nil {
			log.Errorf("error processing event. Retrying... Error: %s. vLog: %+v", err.Error(), vLog)
			return nil, err
		}
	}
	return blocks, nil
}

func (c *ethereumClient) processEvent(ctx context.Context, vLog types.Log, blocks *[]Block) error {
	switch vLog.Topics[0] {
	case accountCreatedSignatureHash:
		return c.accountCreateEvent(vLog, blocks)
	case depositTicketAddedSignatureHash:
		return c.depositEvent(vLog, blocks)
	case withdrawTicketAddedSignatureHash:
		return c.withdrawEvent(vLog, blocks)
	}
	log.Infof("Event not registered: %+v", vLog)
	return nil
}

func (c *ethereumClient) accountCreateEvent(vLog types.Log, blocks *[]Block) error {
	log.Info("AccountCreate event detected")
	evt, err := c.aaFactory.ParseAccountCreated(vLog)
	if err != nil {
		return err
	}

	ac := AccountCreateData{
		Owner:   evt.Owner,
		Account: evt.Account,
	}
	if len(*blocks) == 0 || ((*blocks)[len(*blocks)-1].BlockHash != vLog.BlockHash || (*blocks)[len(*blocks)-1].BlockNumber != vLog.BlockNumber) {
		fullBlock, err := c.ethClient.BlockByHash(context.Background(), vLog.BlockHash)
		if err != nil {
			return fmt.Errorf("error getting hashParent. BlockNumber: %d. Error: %w", vLog.BlockNumber, err)
		}
		t := time.Unix(int64(fullBlock.Time()), 0)
		block := prepareBlock(vLog, t, fullBlock)
		block.ac = append(block.ac, ac)
		*blocks = append(*blocks, block)
	} else if (*blocks)[len(*blocks)-1].BlockHash == vLog.BlockHash && (*blocks)[len(*blocks)-1].BlockNumber == vLog.BlockNumber {
		(*blocks)[len(*blocks)-1].ac = append((*blocks)[len(*blocks)-1].ac, ac)
	} else {
		log.Error("Error processing accountCreate event. BlockHash:", vLog.BlockHash, ". BlockNumber: ", vLog.BlockNumber)
		return fmt.Errorf("error processing accountCreate event")
	}
	or := Order{
		Name: AccountCreatedOrder,
		Pos:  len((*blocks)[len(*blocks)-1].ac) - 1,
	}
	(*blocks)[len(*blocks)-1].or = append((*blocks)[len(*blocks)-1].or, or)
	return nil
}

func (c *ethereumClient) depositEvent(vLog types.Log, blocks *[]Block) error {
	log.Info("Deposit event detected")
	evt, err := c.entryPoint.ParseDepositTicketAdded(vLog)
	if err != nil {
		return err
	}

	dp := DepositData{
		Did:     string(evt.Did[:]),
		Account: evt.Account,
		Amount:  evt.Amount,
	}
	if len(*blocks) == 0 || ((*blocks)[len(*blocks)-1].BlockHash != vLog.BlockHash || (*blocks)[len(*blocks)-1].BlockNumber != vLog.BlockNumber) {
		fullBlock, err := c.ethClient.BlockByHash(context.Background(), vLog.BlockHash)
		if err != nil {
			return fmt.Errorf("error getting hashParent. BlockNumber: %d. Error: %w", vLog.BlockNumber, err)
		}
		t := time.Unix(int64(fullBlock.Time()), 0)
		block := prepareBlock(vLog, t, fullBlock)
		block.dp = append(block.dp, dp)
		*blocks = append(*blocks, block)
	} else if (*blocks)[len(*blocks)-1].BlockHash == vLog.BlockHash && (*blocks)[len(*blocks)-1].BlockNumber == vLog.BlockNumber {
		(*blocks)[len(*blocks)-1].dp = append((*blocks)[len(*blocks)-1].dp, dp)
	} else {
		log.Error("Error processing deposit event. BlockHash:", vLog.BlockHash, ". BlockNumber: ", vLog.BlockNumber)
		return fmt.Errorf("error processing deposit event")
	}
	or := Order{
		Name: DepositTicketAddedOrder,
		Pos:  len((*blocks)[len(*blocks)-1].dp) - 1,
	}
	(*blocks)[len(*blocks)-1].or = append((*blocks)[len(*blocks)-1].or, or)
	return nil
}

func (c *ethereumClient) withdrawEvent(vLog types.Log, blocks *[]Block) error {
	log.Info("Withdraw event detected")
	evt, err := c.entryPoint.ParseWithdrawTicketAdded(vLog)
	if err != nil {
		return err
	}

	wt := WithdrawData{
		Account: evt.Account,
		Amount:  evt.Amount,
	}
	if len(*blocks) == 0 || ((*blocks)[len(*blocks)-1].BlockHash != vLog.BlockHash || (*blocks)[len(*blocks)-1].BlockNumber != vLog.BlockNumber) {
		fullBlock, err := c.ethClient.BlockByHash(context.Background(), vLog.BlockHash)
		if err != nil {
			return fmt.Errorf("error getting hashParent. BlockNumber: %d. Error: %w", vLog.BlockNumber, err)
		}
		t := time.Unix(int64(fullBlock.Time()), 0)
		block := prepareBlock(vLog, t, fullBlock)
		block.wt = append(block.wt, wt)
		*blocks = append(*blocks, block)
	} else if (*blocks)[len(*blocks)-1].BlockHash == vLog.BlockHash && (*blocks)[len(*blocks)-1].BlockNumber == vLog.BlockNumber {
		(*blocks)[len(*blocks)-1].wt = append((*blocks)[len(*blocks)-1].wt, wt)
	} else {
		log.Error("Error processing Withdraw event. BlockHash:", vLog.BlockHash, ". BlockNumber: ", vLog.BlockNumber)
		return fmt.Errorf("error processing Withdraw event")
	}
	or := Order{
		Name: WithdrawTicketAddedOrder,
		Pos:  len((*blocks)[len(*blocks)-1].wt) - 1,
	}
	(*blocks)[len(*blocks)-1].or = append((*blocks)[len(*blocks)-1].or, or)
	return nil
}

func prepareBlock(vLog types.Log, t time.Time, fullBlock *types.Block) Block {
	var block Block
	block.BlockNumber = vLog.BlockNumber
	block.BlockHash = vLog.BlockHash
	block.ParentHash = fullBlock.ParentHash()
	block.ReceivedAt = t
	return block
}

func (c *ethereumClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return c.ethClient.HeaderByNumber(ctx, number)
}

func (c *ethereumClient) EthBlockByNumber(ctx context.Context, blockNumber uint64) (*types.Block, error) {
	block, err := c.ethClient.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		if errors.Is(err, ethereum.NotFound) || err.Error() == "block does not exist in blockchain" {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return block, nil
}
