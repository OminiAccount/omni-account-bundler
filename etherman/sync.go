package etherman

import (
	"context"
	"fmt"
	"github.com/OAB/utils/chains"
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/packutils"
	"github.com/ethereum/go-ethereum/ethdb"
	"time"
)

const SyncChunkSize = 100
const LevalDBNotFound = "leveldb: not found"

type AccountCreateFunc func(acc AccountCreateData)
type DepositFunc func(acc DepositData)
type WithdrawFunc func(acc WithdrawData)
type ExecOpFunc func(wt ExecOpData)

type Synchronizer interface {
	Sync()
	Stop()
}

type ClientSynchronizer struct {
	ctx            context.Context
	etherCli       *EthereumClient
	storage        ethdb.Database
	genBlockNumber uint64
	networkID      chains.ChainId
	synced         bool

	acFunc AccountCreateFunc
	dpFunc DepositFunc
	wtFunc WithdrawFunc
	eoFunc ExecOpFunc
}

func NewSynchronizer(ctx context.Context, db ethdb.Database, etherCli *EthereumClient,
	acFunc AccountCreateFunc, dpFunc DepositFunc, wtFunc WithdrawFunc, eoFunc ExecOpFunc) (Synchronizer, error) {
	cliSync := &ClientSynchronizer{
		storage:        db,
		etherCli:       etherCli,
		ctx:            ctx,
		genBlockNumber: etherCli.cfg.GenBlockNumber,
		networkID:      etherCli.chainID,
		acFunc:         acFunc,
		dpFunc:         dpFunc,
		wtFunc:         wtFunc,
		eoFunc: 		eoFunc,
	}
	lastBlockSync, err := cliSync.storage.Get(cliSync.ChainKey())
	if err != nil {
		if err.Error() == LevalDBNotFound {
			lastBlockSync = []byte{0}
		} else {
			log.Fatal("get chain genblock number error: " + err.Error())
		}
	}
	lastBlockSynced := packutils.BytesToUint64(lastBlockSync)
	if lastBlockSynced > cliSync.genBlockNumber {
		cliSync.genBlockNumber = lastBlockSynced
	}
	if cliSync.genBlockNumber < 1 {
		header, err := cliSync.etherCli.Cli().HeaderByNumber(context.Background(), nil)
		if err != nil || header == nil {
			log.Fatal("networkID: %d, error getting latest block from. Error: %+v", cliSync.networkID, err)
		}
		cliSync.genBlockNumber = header.Number.Uint64()
	}
	_ = cliSync.storage.Put(cliSync.ChainKey(), packutils.Uint64ToBytes(cliSync.genBlockNumber))
	return cliSync, nil
}

var waitDuration = time.Second

func (s *ClientSynchronizer) ChainKey() []byte {
	return []byte(fmt.Sprintf("chain-sync-%d", s.networkID))
}

func (s *ClientSynchronizer) Sync() {
	log.Infof("NetworkID: %d, Synchronization started", s.networkID)
	lastBlockSync, err := s.storage.Get(s.ChainKey())
	if err != nil {
		log.Errorf("networkID: %d, unexpected error getting the latest block. Error: %s", s.networkID, err.Error())
		return
	}
	lastBlockSynced := packutils.BytesToUint64(lastBlockSync)
	log.Infof("NetworkID: %d, initial lastBlockSynced: %+v", s.networkID, lastBlockSynced)
	for {
		select {
		case <-s.ctx.Done():
			log.Infof("NetworkID: %d, synchronizer ctx done", s.networkID)
			err = s.storage.Put(s.ChainKey(), packutils.Uint64ToBytes(lastBlockSynced))
			if err != nil {
				log.Errorf("networkID: %d, put last sync block: %d, error: %+v", s.networkID, lastBlockSynced, err)
				return
			}
			return
		case <-time.After(waitDuration):
			log.Infof("NetworkID: %d, syncing...", s.networkID)
			lastBlockSynced, err = s.syncBlocks(lastBlockSynced)
			if err != nil {
				log.Errorf("networkID: %d, error syncing blocks: %+v", s.networkID, err)
			}
			err = s.storage.Put(s.ChainKey(), packutils.Uint64ToBytes(lastBlockSynced))
			if err != nil {
				log.Errorf("networkID: %d, put last sync block: %d, error: %+v", s.networkID, lastBlockSynced, err)
				return
			}
			if !s.synced {
				header, err := s.etherCli.Cli().HeaderByNumber(s.ctx, nil)
				if err != nil {
					log.Errorf("networkID: %d, error getting latest block from. Error: %s", s.networkID, err.Error())
					continue
				}
				lastKnownBlock := header.Number.Uint64()
				if lastBlockSynced >= lastKnownBlock && !s.synced {
					log.Infof("NetworkID %d Synced!", s.networkID)
					waitDuration = time.Second * 60
					s.synced = true
				}
			}
		}
	}
}

func (s *ClientSynchronizer) Stop() {

}

func (s *ClientSynchronizer) syncBlocks(lastBlockSynced uint64) (uint64, error) {
	header, err := s.etherCli.Cli().HeaderByNumber(s.ctx, nil)
	if err != nil {
		log.Error(err)
		return lastBlockSynced, err
	}
	lastKnownBlock := header.Number

	var fromBlock uint64
	if lastBlockSynced > 0 {
		fromBlock = lastBlockSynced + 1
	}
	if fromBlock > lastKnownBlock.Uint64() {
		log.Infof("NetworkID: %d, block synced, from block: %d, last new block: %d", s.networkID, fromBlock, lastKnownBlock)
		return lastBlockSynced, nil
	}
	for {
		toBlock := fromBlock + SyncChunkSize
		if toBlock > lastKnownBlock.Uint64() {
			toBlock = lastKnownBlock.Uint64()
		}
		if fromBlock > toBlock {
			break
		}
		log.Infof("NetworkID: %d, Getting event info from block %d to block %d", s.networkID, fromBlock, toBlock)
		blocks, err := s.etherCli.GetEventByBlockRange(s.ctx, fromBlock, &toBlock)
		if err != nil {
			log.Error(err)
			return lastBlockSynced, err
		}
		fromBlock = toBlock + 1
		for i := range blocks {
			log.Info("NetworkID: ", s.networkID, ", Position: ", i, ". BlockNumber: ", blocks[i].BlockNumber, ". BlockHash: ", blocks[i].BlockHash)
		}
		_ = s.processBlockRange(blocks)
		lastBlockSynced = toBlock
		if lastKnownBlock.Uint64() <= toBlock {
			if !s.synced {
				log.Infof("NetworkID %d Synced!", s.networkID)
				waitDuration = time.Second * 60
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
				s.dpFunc(blocks[i].dp[element.Pos])
			case WithdrawTicketAddedOrder:

			}
		}
	}
	return nil
}
