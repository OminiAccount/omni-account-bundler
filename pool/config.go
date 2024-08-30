package pool

import (
	"fmt"
	"time"
)

type Config struct {
	maxOps        uint64        `toml:"max-ops"`
	flushInterval time.Duration `toml:"flush-interval"`
}

func (c *Config) UnmarshalTOML(data interface{}) error {
	m, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected type: %T", data)
	}

	if maxOps, ok := m["max-ops"].(int64); ok {
		c.maxOps = uint64(maxOps)
	} else {
		return fmt.Errorf("invalid maxOps")
	}

	if flushInterval, ok := m["flush-interval"].(int64); ok {
		c.flushInterval = time.Duration(flushInterval)
	} else {
		return fmt.Errorf("invalid contract address")
	}

	return nil
}
