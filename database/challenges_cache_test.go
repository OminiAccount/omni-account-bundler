package database

import (
	"fmt"
	"github.com/OAAC/database/leveldb"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestCacheLastUpdateTimeStamp(t *testing.T) {
	dbs, _ := leveldb.NewLevelDB("../../spv_level_db")
	hashList := []common.Hash{
		common.HexToHash("0x7e028521c49d6a139ef03a58c60b49050f76d3cd489d9a026bb552a707f80674"),
	}
	for _, hash := range hashList {
		exist := HasChallenge(dbs, hash.Bytes())
		fmt.Println("exist", exist)
		err := dbs.Put(hash.Bytes(), []byte("1"))
		if err != nil {
			fmt.Println("err", err)
		}
	}

}

func TestDelChallenge(t *testing.T) {
	dbs, _ := leveldb.NewLevelDB("../../spv_level_db")
	hashList := []common.Hash{
		common.HexToHash("0xbb95f79e7044364252cbab60e504bbeb7f9265b0603684fec6b85609c900b267"),
	}
	for _, hash := range hashList {
		exist := HasChallenge(dbs, hash.Bytes())
		fmt.Println("exist", exist)
		err := DelChallenge(dbs, hash.Bytes())
		if err != nil {
			fmt.Println("err", err)
		}
	}
}

func TestUpdate(t *testing.T) {
	dbs, _ := leveldb.NewLevelDB("../../spv_level_db")
	err := WriteLastCreateChallengeTimestamp(dbs, 1703565509)
	if err != nil {
		fmt.Println("err", err)
	}
}
