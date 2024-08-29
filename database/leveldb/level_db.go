package leveldb

import (
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
)

func NewLevelDB(dir string) (ethdb.Database, error) {
	return rawdb.NewLevelDBDatabase(dir, 128, 1024, "spv", false)
}
