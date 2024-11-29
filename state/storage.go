package state

import (
	"github.com/OAB/state/types"
	"github.com/OAB/utils/msgpack"
	"github.com/ethereum/go-ethereum/ethdb"
)

var (
	storageKey = []byte("StateContextKey")
)

type Storage struct {
	Account     *types.UserAccount
	BatchNumber uint64
	db          ethdb.Database
}

func NewStorage(db ethdb.Database) *Storage {
	return &Storage{Account: types.NewUserAccount(), BatchNumber: 0, db: db}
}

func (s *Storage) cache() error {
	types.Lock.RLock()
	defer types.Lock.RUnlock()

	data, err := msgpack.MarshalStruct(s)
	if err != nil {
		return err
	}

	if err = s.db.Put(storageKey, data); err != nil {
		return err
	}

	return nil
}

func (s *Storage) loadCache() error {
	has, err := s.db.Has(storageKey)
	if err != nil {
		return err
	}
	if has {
		data, err := s.db.Get(storageKey)
		if err != nil {
			return err
		}

		decodeData, err := msgpack.UnmarshalStruct[Storage](data)
		if err != nil {
			return err
		}
		s.Account = decodeData.Account
		s.BatchNumber = decodeData.BatchNumber
	}

	return nil
}
