package state

import (
	"context"
	"github.com/ethereum/go-ethereum/ethdb"
)

type Config struct {
	Context   context.Context
	LevelDB   ethdb.Database
}

func NewConfig(ctx context.Context, db ethdb.Database) (Config, error) {
	c := Config{
		Context:   ctx,
		LevelDB:   db,
	}

	return c, nil
}
