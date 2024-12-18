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
	"sync"
	"time"
)

var (
	accountCreatedSignatureHash = crypto.Keccak256Hash([]byte("AccountCreated(address,address)"))
	vizingDepositTicketHash     = crypto.Keccak256Hash([]byte("DepositTicketAdded(bytes32,address,uint256,uint256)"))
	depositTicketHash           = crypto.Keccak256Hash([]byte("ValueDepositAmount(address,address,uint256)"))
	withdrawTicketHash          = crypto.Keccak256Hash([]byte("WithdrawTicketAdded(address,uint256,uint256)"))
	operationPhaseEventHash     = crypto.Keccak256Hash([]byte("UserOperationEvent(bytes32,address,address,uint8,bool,bool,uint256,uint256)"))

	ErrNotFound = errors.New("Not found")
)

type EthereumClient struct {
	cfg            Network
	chainID        chains.ChainId
	ethClient      ethClientInterface
	entryPointAddr common.Address
	entryPoint     *EntryPoint.EntryPoint
	aaFactoryAddr  common.Address
	aaFactory      *SimpleAccountFactory.SimpleAccountFactory
	syncRouterAddr common.Address
	syncRouter     *SyncRouter.SyncRouter
	sender         common.Address
	opAuth         bind.TransactOpts
	lock           sync.Mutex
}

func NewClient(cfg Network) (*EthereumClient, error) {
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

	return &EthereumClient{
		cfg:            cfg,
		chainID:        cfg.ChainId,
		entryPointAddr: cfg.EntryPoint,
		ethClient:      ethClient,
		entryPoint:     entryPoint,
		aaFactoryAddr:  cfg.AccountFactory,
		aaFactory:      accountFactory,
		syncRouterAddr: cfg.SyncRouter,
		syncRouter:     syncRouter,
	}, nil
}

func (c *EthereumClient) GetEventByBlockRange(ctx context.Context, fromBlock uint64, toBlock *uint64) ([]Block, error) {
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

func (c *EthereumClient) readEvents(ctx context.Context, query ethereum.FilterQuery) ([]Block, error) {
	logs, err := c.ethClient.FilterLogs(ctx, query)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var blocks []Block
	for _, vLog := range logs {
		err := c.processEvent(vLog, &blocks)
		if err != nil {
			log.Errorf("chainID: %d, processing event error: %s. vLog: %+v", c.chainID, err.Error(), vLog)
			return nil, err
			//continue
		}
	}
	return blocks, nil
}

func (c *EthereumClient) processEvent(vLog types.Log, blocks *[]Block) error {
	switch vLog.Topics[0] {
	case accountCreatedSignatureHash:
		return c.accountCreateEvent(vLog, blocks)
	case vizingDepositTicketHash:
		return c.vizingDepositEvent(vLog, blocks)
	case depositTicketHash:
		return c.depositEvent(vLog, blocks)
	case withdrawTicketHash:
		return c.withdrawEvent(vLog, blocks)
	case operationPhaseEventHash:
		return c.execOpEvent(vLog, blocks)
	}
	//log.Infof("chainID: %d, Event not registered: %+v", c.chainID, vLog)
	return nil
}

func (c *EthereumClient) accountCreateEvent(vLog types.Log, blocks *[]Block) error {
	log.Infof("chainID: %d, accountCreate event detected", c.chainID)
	evt, err := c.aaFactory.ParseAccountCreated(vLog)
	if err != nil {
		return err
	}

	ac := ToAccountCreateData(evt)
	ac.ChainID = c.chainID
	if len(*blocks) == 0 || ((*blocks)[len(*blocks)-1].BlockHash != vLog.BlockHash || (*blocks)[len(*blocks)-1].BlockNumber != vLog.BlockNumber) {
		fullBlock, err := c.ethClient.HeaderByHash(context.Background(), vLog.BlockHash)
		if err != nil {
			return fmt.Errorf("error getting hashParent. BlockNumber: %d. Error: %w", vLog.BlockNumber, err)
		}
		t := time.Unix(int64(fullBlock.Time), 0)
		block := prepareBlock(vLog, t, fullBlock)
		block.ac = append(block.ac, ac)
		*blocks = append(*blocks, block)
	} else if (*blocks)[len(*blocks)-1].BlockHash == vLog.BlockHash && (*blocks)[len(*blocks)-1].BlockNumber == vLog.BlockNumber {
		(*blocks)[len(*blocks)-1].ac = append((*blocks)[len(*blocks)-1].ac, ac)
	} else {
		log.Errorf("chainID: %d, Error processing accountCreate event. BlockHash: %s. BlockNumber: %d",
			c.chainID, vLog.BlockHash, vLog.BlockNumber)
		return fmt.Errorf("error processing accountCreate event")
	}
	or := Order{
		Name: AccountCreatedOrder,
		Pos:  len((*blocks)[len(*blocks)-1].ac) - 1,
	}
	(*blocks)[len(*blocks)-1].or = append((*blocks)[len(*blocks)-1].or, or)
	return nil
}

func (c *EthereumClient) depositEvent(vLog types.Log, blocks *[]Block) error {
	log.Infof("chainID: %d, deposit event detected", c.chainID)
	evt, err := c.entryPoint.ParseValueDepositAmount(vLog)
	if err != nil {
		return err
	}

	dp := ToDepositData(evt)
	dp.ChainID = c.chainID
	if len(*blocks) == 0 || ((*blocks)[len(*blocks)-1].BlockHash != vLog.BlockHash || (*blocks)[len(*blocks)-1].BlockNumber != vLog.BlockNumber) {
		fullBlock, err := c.ethClient.HeaderByHash(context.Background(), vLog.BlockHash)
		if err != nil {
			return fmt.Errorf("error getting hashParent. BlockNumber: %d. Error: %w", vLog.BlockNumber, err)
		}
		t := time.Unix(int64(fullBlock.Time), 0)
		block := prepareBlock(vLog, t, fullBlock)
		block.dp = append(block.dp, dp)
		*blocks = append(*blocks, block)
	} else if (*blocks)[len(*blocks)-1].BlockHash == vLog.BlockHash && (*blocks)[len(*blocks)-1].BlockNumber == vLog.BlockNumber {
		(*blocks)[len(*blocks)-1].dp = append((*blocks)[len(*blocks)-1].dp, dp)
	} else {
		log.Errorf("chainID: %d, Error processing deposit event. BlockHash: %s, . BlockNumber: %d",
			c.chainID, vLog.BlockHash, vLog.BlockNumber)
		return fmt.Errorf("error processing deposit event")
	}
	or := Order{
		Name: DepositTicketAddedOrder,
		Pos:  len((*blocks)[len(*blocks)-1].dp) - 1,
	}
	(*blocks)[len(*blocks)-1].or = append((*blocks)[len(*blocks)-1].or, or)
	return nil
}

func (c *EthereumClient) vizingDepositEvent(vLog types.Log, blocks *[]Block) error {
	log.Infof("chainID: %d, vizing deposit event detected", c.chainID)
	evt, err := c.entryPoint.ParseDepositTicketAdded(vLog)
	if err != nil {
		return err
	}

	dp := ToVizingDepositData(evt)
	dp.ChainID = c.chainID
	if len(*blocks) == 0 || ((*blocks)[len(*blocks)-1].BlockHash != vLog.BlockHash || (*blocks)[len(*blocks)-1].BlockNumber != vLog.BlockNumber) {
		fullBlock, err := c.ethClient.HeaderByHash(context.Background(), vLog.BlockHash)
		if err != nil {
			return fmt.Errorf("error getting hashParent. BlockNumber: %d. Error: %w", vLog.BlockNumber, err)
		}
		t := time.Unix(int64(fullBlock.Time), 0)
		block := prepareBlock(vLog, t, fullBlock)
		block.dp = append(block.dp, dp)
		*blocks = append(*blocks, block)
	} else if (*blocks)[len(*blocks)-1].BlockHash == vLog.BlockHash && (*blocks)[len(*blocks)-1].BlockNumber == vLog.BlockNumber {
		(*blocks)[len(*blocks)-1].dp = append((*blocks)[len(*blocks)-1].dp, dp)
	} else {
		log.Errorf("chainID: %d, Error processing deposit event. BlockHash: %s, . BlockNumber: %d",
			c.chainID, vLog.BlockHash, vLog.BlockNumber)
		return fmt.Errorf("error processing deposit event")
	}
	or := Order{
		Name: VizingDepositTicketOrder,
		Pos:  len((*blocks)[len(*blocks)-1].dp) - 1,
	}
	(*blocks)[len(*blocks)-1].or = append((*blocks)[len(*blocks)-1].or, or)
	return nil
}

func (c *EthereumClient) withdrawEvent(vLog types.Log, blocks *[]Block) error {
	log.Infof("chainID: %d, withdraw event detected", c.chainID)
	evt, err := c.entryPoint.ParseWithdrawTicketAdded(vLog)
	if err != nil {
		return err
	}

	wt := ToWithdrawData(evt)
	if len(*blocks) == 0 || ((*blocks)[len(*blocks)-1].BlockHash != vLog.BlockHash || (*blocks)[len(*blocks)-1].BlockNumber != vLog.BlockNumber) {
		fullBlock, err := c.ethClient.HeaderByHash(context.Background(), vLog.BlockHash)
		if err != nil {
			return fmt.Errorf("error getting hashParent. BlockNumber: %d. Error: %w", vLog.BlockNumber, err)
		}
		t := time.Unix(int64(fullBlock.Time), 0)
		block := prepareBlock(vLog, t, fullBlock)
		block.wt = append(block.wt, wt)
		*blocks = append(*blocks, block)
	} else if (*blocks)[len(*blocks)-1].BlockHash == vLog.BlockHash && (*blocks)[len(*blocks)-1].BlockNumber == vLog.BlockNumber {
		(*blocks)[len(*blocks)-1].wt = append((*blocks)[len(*blocks)-1].wt, wt)
	} else {
		log.Errorf("chainID: %d, Error processing Withdraw event. BlockHash: %s. BlockNumber: %d",
			c.chainID, vLog.BlockHash, vLog.BlockNumber)
		return fmt.Errorf("error processing Withdraw event")
	}
	or := Order{
		Name: WithdrawTicketAddedOrder,
		Pos:  len((*blocks)[len(*blocks)-1].wt) - 1,
	}
	(*blocks)[len(*blocks)-1].or = append((*blocks)[len(*blocks)-1].or, or)
	return nil
}

func (c *EthereumClient) execOpEvent(vLog types.Log, blocks *[]Block) error {
	log.Infof("chainID: %d, execop event detected", c.chainID)
	evt, err := c.entryPoint.ParseUserOperationEvent(vLog)
	if err != nil {
		return err
	}
	log.Infof("%+v", evt)
	eo := ToExecOpData(evt)
	if len(*blocks) == 0 || ((*blocks)[len(*blocks)-1].BlockHash != vLog.BlockHash || (*blocks)[len(*blocks)-1].BlockNumber != vLog.BlockNumber) {
		fullBlock, err := c.ethClient.HeaderByHash(context.Background(), vLog.BlockHash)
		if err != nil {
			return fmt.Errorf("error getting hashParent. BlockNumber: %d. Error: %w", vLog.BlockNumber, err)
		}
		t := time.Unix(int64(fullBlock.Time), 0)
		block := prepareBlock(vLog, t, fullBlock)
		block.eo = append(block.eo, eo)
		*blocks = append(*blocks, block)
	} else if (*blocks)[len(*blocks)-1].BlockHash == vLog.BlockHash && (*blocks)[len(*blocks)-1].BlockNumber == vLog.BlockNumber {
		(*blocks)[len(*blocks)-1].eo = append((*blocks)[len(*blocks)-1].eo, eo)
	} else {
		log.Errorf("chainID: %d, Error processing Withdraw event. BlockHash: %s. BlockNumber: %d",
			c.chainID, vLog.BlockHash, vLog.BlockNumber)
		return fmt.Errorf("error processing Withdraw event")
	}
	or := Order{
		Name: ExecOpOrder,
		Pos:  len((*blocks)[len(*blocks)-1].eo) - 1,
	}
	(*blocks)[len(*blocks)-1].or = append((*blocks)[len(*blocks)-1].or, or)
	return nil
}

func prepareBlock(vLog types.Log, t time.Time, fullBlock *types.Header) Block {
	var block Block
	block.BlockNumber = vLog.BlockNumber
	block.BlockHash = vLog.BlockHash
	block.ParentHash = fullBlock.ParentHash
	block.ReceivedAt = t
	return block
}

func (c *EthereumClient) Cli() ethClientInterface {
	return c.ethClient
}

func (c *EthereumClient) IsNeedSync() bool {
	return !(c.cfg.IsSync == 0)
}

func (c *EthereumClient) EthBlockByNumber(ctx context.Context, blockNumber uint64) (*types.Block, error) {
	block, err := c.ethClient.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		if errors.Is(err, ethereum.NotFound) || err.Error() == "block does not exist in blockchain" {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return block, nil
}
