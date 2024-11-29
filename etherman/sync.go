package etherman

import (
	"context"
	"fmt"
	"github.com/OAB/utils/chains"
	"github.com/OAB/utils/packutils"
	"github.com/ethereum/go-ethereum/ethdb"
	"log"
	"math/big"
	"time"
)

const SyncChunkSize = 100

type AccountCreateFunc func(acc AccountCreateData)
type DepositFunc func(acc DepositData)
type WithdrawFunc func(acc WithdrawData)

type Synchronizer interface {
	Sync()
	Stop()
}

type ClientSynchronizer struct {
	ctx            context.Context
	etherCli       *ethereumClient
	storage        ethdb.Database
	cancelCtx      context.CancelFunc
	genBlockNumber uint64
	networkID      chains.ChainId
	synced         bool

	acFunc AccountCreateFunc
	dpFunc DepositFunc
	wtFunc WithdrawFunc
}

func NewSynchronizer(db ethdb.Database, etherCli *ethereumClient, genBlockNumber uint64,
	acFunc AccountCreateFunc, dpFunc DepositFunc, wtFunc WithdrawFunc) (Synchronizer, error) {
	ctx, cancel := context.WithCancel(context.Background())
	cliSync := &ClientSynchronizer{
		storage:        db,
		etherCli:       etherCli,
		ctx:            ctx,
		cancelCtx:      cancel,
		genBlockNumber: genBlockNumber,
		networkID:      etherCli.chainID,
		acFunc:         acFunc,
		dpFunc:         dpFunc,
		wtFunc:         wtFunc,
	}
	lastBlockSync, err := cliSync.storage.Get(cliSync.ChainKey())
	if err != nil {
		panic("get chain genblock number error: " + err.Error())
	}
	lastBlockSynced := packutils.BytesToUint64(lastBlockSync)
	if lastBlockSynced > cliSync.genBlockNumber {
		cliSync.genBlockNumber = lastBlockSynced
	}
	if cliSync.genBlockNumber < 1 {
		header, err := cliSync.etherCli.HeaderByNumber(context.Background(), nil)
		if err != nil {
			fmt.Printf("networkID: %d, error getting latest block from. Error: %s\n", cliSync.networkID, err.Error())
			panic("NewSynchronizer error")
		}
		cliSync.genBlockNumber = header.Number.Uint64()
	}
	return cliSync, nil
}

var waitDuration = time.Duration(0)

func (s *ClientSynchronizer) ChainKey() []byte {
	return []byte(fmt.Sprintf("chain-sync-%d", s.networkID))
}

func (s *ClientSynchronizer) Sync() {
	fmt.Printf("NetworkID: %d, Synchronization started\n", s.networkID)
	lastBlockSync, err := s.storage.Get(s.ChainKey())
	if err != nil {
		fmt.Println(fmt.Sprintf("networkID: %d, unexpected error getting the latest block. Error: %s", s.networkID, err.Error()))
		return
	}
	lastBlockSynced := packutils.BytesToUint64(lastBlockSync)
	fmt.Println(fmt.Sprintf("NetworkID: %d, initial lastBlockSynced: %+v", s.networkID, lastBlockSynced))
	for {
		select {
		case <-s.ctx.Done():
			fmt.Println(fmt.Sprintf("NetworkID: %d, synchronizer ctx done", s.networkID))
			return
		case <-time.After(waitDuration):
			fmt.Println(fmt.Sprintf("NetworkID: %d, syncing...", s.networkID))
			if lastBlockSynced, err = s.syncBlocks(lastBlockSynced); err != nil {
				fmt.Println(fmt.Sprintf("networkID: %d, error syncing blocks: %+v", s.networkID, err))
				err = s.storage.Put(s.ChainKey(), packutils.Uint64ToBytes(lastBlockSynced))
				if err != nil {
					fmt.Println(fmt.Sprintf("networkID: %d, put last sync block error: %+v", s.networkID, err))
					return
				}
				if s.ctx.Err() != nil {
					continue
				}
			}
			if !s.synced {
				header, err := s.etherCli.HeaderByNumber(s.ctx, nil)
				if err != nil {
					fmt.Printf("networkID: %d, error getting latest block from. Error: %s\n", s.networkID, err.Error())
					continue
				}
				lastKnownBlock := header.Number.Uint64()
				if lastBlockSynced == lastKnownBlock && !s.synced {
					fmt.Printf("NetworkID %d Synced!\n", s.networkID)
					waitDuration = time.Second * 10
					s.synced = true
				}
			}
		}
	}
}

func (s *ClientSynchronizer) Stop() {
	s.cancelCtx()
}

func (s *ClientSynchronizer) syncBlocks(lastBlockSynced uint64) (uint64, error) {
	log.Printf("NetworkID: %d, after checkReorg: no reorg detected", s.networkID)
	header, err := s.etherCli.HeaderByNumber(s.ctx, nil)
	if err != nil {
		return lastBlockSynced, err
	}
	lastKnownBlock := header.Number

	var fromBlock uint64
	if lastBlockSynced > 0 {
		fromBlock = lastBlockSynced + 1
	}

	for {
		toBlock := fromBlock + SyncChunkSize

		log.Printf("NetworkID: %d, Getting bridge info from block %d to block %d\n", s.networkID, fromBlock, toBlock)
		blocks, err := s.etherCli.GetEventByBlockRange(s.ctx, fromBlock, &toBlock)
		if err != nil {
			return lastBlockSynced, err
		}
		fromBlock = toBlock + 1
		err = s.processBlockRange(blocks)
		if err != nil {
			return lastBlockSynced, err
		}
		if len(blocks) > 0 {
			lastBlockSynced = blocks[len(blocks)-1].BlockNumber
			for i := range blocks {
				fmt.Println("NetworkID: ", s.networkID, ", Position: ", i, ". BlockNumber: ", blocks[i].BlockNumber, ". BlockHash: ", blocks[i].BlockHash)
			}
		}
		if len(blocks) < 1 {
			fb, err := s.etherCli.EthBlockByNumber(s.ctx, toBlock)
			if err != nil || fb == nil || fb.NumberU64() != toBlock {
				return lastBlockSynced, err
			}
			lastBlockSynced = toBlock
		}
		if lastKnownBlock.Cmp(new(big.Int).SetUint64(toBlock)) < 1 {
			if !s.synced {
				fmt.Printf("NetworkID %d Synced!\n", s.networkID)
				waitDuration = time.Second * 10
				s.synced = true
			}
			break
		}
	}
	return lastBlockSynced, nil
}

func (s *ClientSynchronizer) processBlockRange(blocks []Block) error {
	for i := range blocks {
		for _, element := range blocks[i].or {
			switch element.Name {
			case AccountCreatedOrder:
				s.acFunc(blocks[i].ac[element.Pos])
			case DepositTicketAddedOrder:

			case WithdrawTicketAddedOrder:

			}
		}
	}
	return nil
}
