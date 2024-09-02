package processor

import (
	"context"
	"github.com/OAAC/config"
	"github.com/OAAC/database/leveldb"
	"github.com/OAAC/ethereum"
	"github.com/OAAC/jsonrpc"
	"github.com/OAAC/pool"
	"github.com/OAAC/state"
	"github.com/OAAC/synchronizer"
	"github.com/OAAC/utils"
	"github.com/ethereum/go-ethereum/log"
	"path/filepath"
)

type Processor struct {
	ctx context.Context
	cfg config.Config

	pool         *pool.Pool
	server       *jsonrpc.Server
	ethereum     *ethereum.Ethereum
	synchronizer *synchronizer.Synchronizer
	service      *state.Service
}

var processor *Processor

// NewProcessor Initialize the processor
func NewProcessor(cfg config.Config) (*Processor, error) {
	ctx := context.Background()

	// Connect to levelDB
	levelDBDir, err := filepath.Abs("../../spv_level_db")
	if _, err = utils.PathExists(levelDBDir); err != nil {
		return nil, err
	}
	levelDB, err := leveldb.NewLevelDB(levelDBDir)
	if err != nil {
		return nil, err
	}

	// pool
	poolInstance := createPool(cfg.Pool)
	log.Info("Pool successfully initialized")

	// jsonrpc
	server := createJSONRPCServer(cfg.JsonRpc, poolInstance)
	log.Info("JSONRPCServer successfully initialized")

	// Init Service
	serviceConfig, err := state.NewServiceConfig(ctx, levelDB, cfg)
	if err != nil {
		return nil, err
	}
	service, err := state.NewService(serviceConfig)
	if err != nil {
		return nil, err
	}

	// Ethereum
	ethereum, err := ethereum.NewEthereum(cfg.Ethereum)
	if err != nil {
		return nil, err
	}
	log.Info("Ethereum successfully initialized")

	// Synchronizer
	synchronizer, err := synchronizer.NewSynchronizer(poolInstance, ethereum, synchronizer.Config{EthereumCfg: cfg.Ethereum})
	if err != nil {
		return nil, err
	}
	log.Info("Synchronizer successfully initialized")

	processor = &Processor{
		ctx:          ctx,
		cfg:          cfg,
		server:       server,
		ethereum:     ethereum,
		synchronizer: synchronizer,
		service:      service,
	}

	return processor, nil
}

// Start a processor
func (p *Processor) Start() error {
	// start rpc server
	p.server.Start()
	go p.synchronizer.Sync()
	return nil
}

// Stop a processor
func (p *Processor) Stop() {
	log.Warn("Stopping processor")
	//p.service.Stop()
	p.synchronizer.Stop()
	log.Warn("Processor stopped")

}
