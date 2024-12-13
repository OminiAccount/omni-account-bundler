package leveldb

import (
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
)

func NewLevelDB(dir string) (ethdb.Database, error) {
	db, err := leveldb.New(dir, 128, 1024, "spv", false)
	if err != nil {
		return nil, err
	}
	log.Info("Using LevelDB as the backing database")
	return rawdb.NewDatabase(db), nil
	//return rawdb.NewLevelDBDatabase(dir, 128, 1024, "spv", false)
}
