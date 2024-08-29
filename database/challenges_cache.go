package database

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/ethdb"
)

var lastCreateChallengeTimestamp = []byte("lastCreateChallengeTimestamp")

func ReadLastCreateChallengeTimestamp(leveldb ethdb.KeyValueReader) (uint64, error) {
	var lastUpdateTimeStamp uint64
	isExist, err := hasExist(leveldb, lastCreateChallengeTimestamp)
	if err != nil {
		return 0, err
	}
	if isExist {
		lastUpdateTimeStampBytes, err := leveldb.Get(lastCreateChallengeTimestamp)
		if err != nil {
			return 0, err
		}
		lastUpdateTimeStamp = binary.LittleEndian.Uint64(lastUpdateTimeStampBytes)

	}

	return lastUpdateTimeStamp, nil

}

func WriteLastCreateChallengeTimestamp(leveldb ethdb.KeyValueWriter, value uint64) error {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, value)
	return leveldb.Put(lastCreateChallengeTimestamp, b)
}

func HasChallenge(leveldb ethdb.KeyValueReader, key []byte) bool {
	has, _ := hasExist(leveldb, key)
	return has
}

func WriteChallenges(leveldb ethdb.Database, keys [][]byte) error {
	batch := leveldb.NewBatch()
	for _, key := range keys {
		err := batch.Put(key, []byte(""))
		if err != nil {
			return err
		}
	}
	err := batch.Write()
	if err != nil {
		return err
	}
	return nil
}

func UpdateChallenge(leveldb ethdb.Database, key, value []byte) error {
	return leveldb.Put(key, value)
}

func GetChallenge(leveldb ethdb.Database, key []byte) ([]byte, error) {
	return leveldb.Get(key)
}

func DelChallenge(leveldb ethdb.Database, key []byte) error {
	return leveldb.Delete(key)
}

func hasExist(leveldb ethdb.KeyValueReader, key []byte) (bool, error) {
	return leveldb.Has(key)
}
