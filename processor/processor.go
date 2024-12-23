package processor

import (
	"context"
	"github.com/OAB/config"
	"github.com/OAB/database/leveldb"
	"github.com/OAB/database/pgstorage"
	"github.com/OAB/etherman"
	"github.com/OAB/jsonrpc"
	"github.com/OAB/pool"
	"github.com/OAB/state"
	"github.com/OAB/synchronizer"
	"github.com/OAB/utils"
	"github.com/OAB/utils/db"
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/merkletree"
	"path/filepath"
)

type Processor struct {
	ctx context.Context
	cfg *config.Config

	pool         *pool.Pool
	server       *jsonrpc.Server
	ethereum     *etherman.EtherMan
	synchronizer *synchronizer.Synchronizer
	state        *state.State
}

var processor *Processor

// NewProcessor Initialize the processor
func NewProcessor(cfg *config.Config) (*Processor, error) {
	ctx := context.Background()

	storage, err := pgstorage.NewPostgresStorage(cfg.Db)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("DB successfully initialized")

	// Connect to levelDB
	levelDBDir, err := filepath.Abs("./db")
	if _, err = utils.PathExists(levelDBDir); err != nil {
		return nil, err
	}
	levelDB, err := leveldb.NewLevelDB(levelDBDir)
	if err != nil {
		return nil, err
	}

	// tree
	tree := merkletree.NewSMT(db.NewMemDb(levelDB), false)

	// Ethereum
	ethereum, err := etherman.NewEthereum(ctx, cfg.Ethereum, storage)
	if err != nil {
		return nil, err
	}
	log.Info("Ethereum successfully initialized")

	st, err := state.NewState(ctx, cfg.State, tree, ethereum, storage)
	if err != nil {
		return nil, err
	}

	// pool
	poolIns := pool.NewMemoryPool(cfg.Pool, ethereum, st, storage)
	log.Info("Pool successfully initialized")

	// Synchronizer
	sync, err := synchronizer.NewSynchronizer(ctx, ethereum, poolIns, st, levelDB, storage)
	if err != nil {
		return nil, err
	}
	log.Info("Synchronizer successfully initialized")

	// jsonrpc
	server := createJSONRPCServer(cfg.JsonRpc, poolIns, st)
	log.Info("JSONRPCServer successfully initialized")

	processor = &Processor{
		ctx:          ctx,
		cfg:          cfg,
		server:       server,
		ethereum:     ethereum,
		synchronizer: sync,
		state:        st,
		pool:         poolIns,
	}

	return processor, nil
}

// Start a processor
func (p *Processor) Start() error {
	// start rpc server
	p.server.Start()
	p.state.Start()
	p.pool.StartAutoFlush()
	p.ethereum.Start()
	p.synchronizer.Start()
	return nil
}

// Stop a processor
func (p *Processor) Stop() {
	log.Warn("stopping processor")
	//p.service.Stop()
	p.synchronizer.Stop()
	p.ethereum.Stop()
	p.pool.StopAutoFlush()
	p.state.Stop()
	log.Warn("processor stopped")
}
