package pool

import (
	"time"
)

type Config struct {
	MaxOps        uint64        `mapstructure:"max-ops"`
	FlushInterval time.Duration `mapstructure:"flush-interval"`
}
