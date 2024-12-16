package processor

import (
	"context"
	"github.com/OAB/config"
	"github.com/OAB/database/leveldb"
	"github.com/OAB/etherman"
	"github.com/OAB/jsonrpc"
	"github.com/OAB/pool"
	"github.com/OAB/state"
	"github.com/OAB/synchronizer"
	"github.com/OAB/utils"
	"github.com/OAB/utils/log"
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

	// Connect to levelDB
	levelDBDir, err := filepath.Abs("./db")
	if _, err = utils.PathExists(levelDBDir); err != nil {
		return nil, err
	}
	levelDB, err := leveldb.NewLevelDB(levelDBDir)
	if err != nil {
		return nil, err
	}

	// pool
	poolInstance := pool.NewMemoryPool(cfg.Pool, levelDB)
	log.Info("Pool successfully initialized")

	// Ethereum
	ethereum, err := etherman.NewEthereum(ctx, cfg.Ethereum, levelDB)
	if err != nil {
		return nil, err
	}
	log.Info("Ethereum successfully initialized")

	// Init State
	stateConfig, err := state.NewConfig(ctx, levelDB)
	if err != nil {
		return nil, err
	}
	st, err := state.NewState(stateConfig, poolInstance, ethereum)
	if err != nil {
		return nil, err
	}

	// Synchronizer
	sync, err := synchronizer.NewSynchronizer(poolInstance, ethereum, st, levelDB)
	if err != nil {
		return nil, err
	}
	log.Info("Synchronizer successfully initialized")

	// jsonrpc
	server := createJSONRPCServer(cfg.JsonRpc, st)
	log.Info("JSONRPCServer successfully initialized")

	processor = &Processor{
		ctx:          ctx,
		cfg:          cfg,
		server:       server,
		ethereum:     ethereum,
		synchronizer: sync,
		state:        st,
		pool:         poolInstance,
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
