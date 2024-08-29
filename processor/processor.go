package processor

import (
	"context"
	"github.com/OAAC/config"
	"github.com/OAAC/database/leveldb"
	"github.com/OAAC/server"
	"github.com/OAAC/services"
	"github.com/OAAC/utils"
	"github.com/ethereum/go-ethereum/log"
	"path/filepath"
)

type Processor struct {
	ctx context.Context
	cfg config.Config

	service *services.Service
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

	// Init Service
	serviceConfig, err := services.NewServiceConfig(ctx, levelDB, cfg)
	if err != nil {
		return nil, err
	}
	service, err := services.NewService(serviceConfig)
	if err != nil {
		return nil, err
	}

	processor = &Processor{
		ctx:     ctx,
		cfg:     cfg,
		service: service,
	}

	return processor, nil
}

// Start a processor
func (p *Processor) Start() error {
	// start rpc server
	server.NewService(&p.cfg.API)
	return nil
}

// Stop a processor
func (p *Processor) Stop() {
	log.Warn("Stopping processor")
	//p.service.Stop()
	log.Warn("Processor stopped")

}
