package processor

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/omni-account/client/config"
	"github.com/omni-account/client/services"
)

type Processor struct {
	ctx context.Context
	cfg config.Config
}

var processor *Processor

// NewProcessor Initialize the processor
func NewProcessor(cfg config.Config) (*Processor, error) {
	ctx := context.Background()

	processor = &Processor{
		ctx: ctx,
		cfg: cfg,
	}

	return processor, nil
}

// Start a processor
func (p *Processor) Start() error {
	// start rpc server
	services.NewService(&p.cfg.API)
	return nil
}

// Stop a processor
func (p *Processor) Stop() {
	log.Warn("Stopping processor")
	//p.service.Stop()
	log.Warn("Processor stopped")

}
