package synchronizer

import "github.com/OAB/utils/apitypes"

type Config struct {
	SyncInterval  apitypes.Duration `mapstructure:"sync-interval"`
	SyncChunkSize uint64            `mapstructure:"sync-chunk-size"`
}
