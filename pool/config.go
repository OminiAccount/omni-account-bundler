package pool

import (
	"time"
)

type Config struct {
	MaxOps        uint64        `mapstructure:"max-ops"`
	MaxBatches    uint64        `mapstructure:"max-batches"`
	FlushInterval time.Duration `mapstructure:"flush-interval"`
}
