package etherman

import (
	"context"
	"github.com/OAB/utils/log"
	"github.com/jackc/pgx/v4/pgxpool"
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
	db             *PostgresStorage
	genBlockNumber uint64
	blockCheckNum  uint64
	networkID      uint64
	synced         bool

	acFunc  AccountCreateFunc
	vdpFunc DepositFunc
	dpFunc  DepositFunc
	wtFunc  WithdrawFunc
	eoFunc  ExecOpFunc
}

func NewSynchronizer(ctx context.Context, db *pgxpool.Pool, etherCli *EthereumClient,
	acFunc AccountCreateFunc, vdpFunc DepositFunc, dpFunc DepositFunc,
	wtFunc WithdrawFunc, eoFunc ExecOpFunc) (Synchronizer, error) {
	cliSync := &ClientSynchronizer{
		db:             NewPostgresStorage(db),
		etherCli:       etherCli,
		ctx:            ctx,
		genBlockNumber: etherCli.cfg.GenBlockNumber,
		blockCheckNum:  etherCli.GetBlockCheckNum(),
		networkID:      etherCli.chainID,
		acFunc:         acFunc,
		vdpFunc:        vdpFunc,
		dpFunc:         dpFunc,
		wtFunc:         wtFunc,
		eoFunc:         eoFunc,
	}
	cliSync.db.AddChainHeight(ctx, cliSync.networkID, nil)
	lastBlockSync := cliSync.db.GetChainHeight(ctx, cliSync.networkID, nil)
	if lastBlockSync > cliSync.genBlockNumber {
		cliSync.genBlockNumber = lastBlockSync
	}
	if cliSync.genBlockNumber < 1 {
		header, err := cliSync.etherCli.Cli().HeaderByNumber(context.Background(), nil)
		if err != nil || header == nil {
			log.Fatal("networkID: %d, error getting latest block from. Error: %+v", cliSync.networkID, err)
		}
		cliSync.genBlockNumber = header.Number.Uint64() - cliSync.blockCheckNum
	}
	_ = cliSync.db.SetChain(ctx, cliSync.networkID, cliSync.genBlockNumber, "block_id", nil)
	return cliSync, nil
}

var waitDuration = time.Second

func (s *ClientSynchronizer) Sync() {
	log.Infof("NetworkID: %d, Synchronization started", s.networkID)
	lastBlockSynced := s.db.GetChainHeight(s.ctx, s.networkID, nil)
	if lastBlockSynced < 1 {
		log.Errorf("networkID: %d, unexpected error getting the latest block.", s.networkID)
		return
	}
	log.Infof("NetworkID: %d, initial lastBlockSynced: %+v", s.networkID, lastBlockSynced)
	for {
		select {
		case <-s.ctx.Done():
			log.Infof("NetworkID: %d, synchronizer ctx done", s.networkID)
			_ = s.db.SetChain(s.ctx, s.networkID, lastBlockSynced, "block_id", nil)
			return
		case <-time.After(waitDuration):
			log.Infof("NetworkID: %d, syncing...", s.networkID)
			var err error
			lastBlockSynced, err = s.syncBlocks(lastBlockSynced)
			if err != nil {
				log.Errorf("networkID: %d, error syncing blocks: %+v", s.networkID, err)
			}
			_ = s.db.SetChain(s.ctx, s.networkID, lastBlockSynced, "block_id", nil)
			if !s.synced {
				header, err := s.etherCli.Cli().HeaderByNumber(s.ctx, nil)
				if err != nil {
					log.Errorf("networkID: %d, error getting latest block from. Error: %s", s.networkID, err.Error())
					continue
				}
				lastKnownBlock := header.Number.Uint64() - s.blockCheckNum
				if lastBlockSynced >= lastKnownBlock && !s.synced {
					log.Infof("NetworkID %d Synced!", s.networkID)
					waitDuration = time.Second * 20
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
	lastKnownBlock := header.Number.Uint64() - s.blockCheckNum

	var fromBlock uint64
	if lastBlockSynced > 0 {
		fromBlock = lastBlockSynced + 1
	}
	if fromBlock > lastKnownBlock {
		log.Infof("NetworkID: %d, block synced, from block: %d, last new block: %d", s.networkID, fromBlock, lastKnownBlock)
		return lastBlockSynced, nil
	}
	for {
		toBlock := fromBlock + SyncChunkSize
		if toBlock > lastKnownBlock {
			toBlock = lastKnownBlock
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
		if lastKnownBlock <= toBlock {
			if !s.synced {
				log.Infof("NetworkID %d Synced!", s.networkID)
				waitDuration = time.Second * 20
				s.synced = true
			}
			break
		}
	}
	return lastBlockSynced, nil
}

func (s *ClientSynchronizer) processBlockRange(blocks []Block) error {
	for i := range blocks {
		for _, ele := range blocks[i].or {
			switch ele.Name {
			case AccountCreatedOrder:
				s.acFunc(blocks[i].ac[ele.Pos])
			case VizingDepositTicketOrder:
				s.vdpFunc(blocks[i].dp[ele.Pos])
			case DepositTicketAddedOrder:
				s.dpFunc(blocks[i].dp[ele.Pos])
			case WithdrawTicketAddedOrder:
				s.wtFunc(blocks[i].wt[ele.Pos])
			case ExecOpOrder:
				s.eoFunc(blocks[i].eo[ele.Pos])
			}
		}
	}
	return nil
}
