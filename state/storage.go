package state

import (
	"errors"
	"fmt"
	"github.com/OAB/state/types"
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/msgpack"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
)

var (
	storageKey = []byte("StateContextKey")
	batchKey   = []byte("BatchKey-%d")
	hisKey     = []byte("His-%s-%d")
	NotBatch   = errors.New("not batch")
	NotHis     = errors.New("not history")
)

type Storage struct {
	Account           *types.UserAccount
	BatchNumber       uint64
	VerifyBatchNumber uint64
	db                ethdb.Database
}

func NewStorage(db ethdb.Database) *Storage {
	return &Storage{Account: types.NewUserAccount(), BatchNumber: 0, db: db}
}

func (s *Storage) cacheUserHistory(u common.Address, page uint64, d []types.AccountHistory) error {
	data, err := msgpack.MarshalStruct(d)
	if err != nil {
		return err
	}
	key := fmt.Sprintf(string(hisKey), u.String(), page)
	if err = s.db.Put([]byte(key), data); err != nil {
		return err
	}
	return nil
}

func (s *Storage) loadUserHistory(u common.Address, page uint64) ([]types.AccountHistory, error) {
	key := fmt.Sprintf(string(hisKey), u.String(), page)
	has, err := s.db.Has([]byte(key))
	if err != nil {
		return nil, err
	}
	if has {
		data, err := s.db.Get([]byte(key))
		if err != nil {
			log.Errorf("%v", err)
			return nil, err
		}

		decodeData, err := msgpack.UnmarshalStruct[[]types.AccountHistory](data)
		if err != nil {
			return nil, err
		}
		return decodeData, nil
	}
	return nil, NotHis
}

func (s *Storage) cacheBatch(d *Batch) error {
	data, err := msgpack.MarshalStruct(d)
	if err != nil {
		return err
	}
	key := fmt.Sprintf(string(batchKey), d.NewNumBatch)
	if err = s.db.Put([]byte(key), data); err != nil {
		return err
	}
	s.BatchNumber = d.NewNumBatch
	return nil
}

func (s *Storage) loadBatch(num uint64) (*Batch, error) {
	key := fmt.Sprintf(string(batchKey), num)
	has, err := s.db.Has([]byte(key))
	if err != nil {
		return nil, err
	}
	if has {
		data, err := s.db.Get([]byte(key))
		if err != nil {
			return nil, err
		}

		decodeData, err := msgpack.UnmarshalStruct[*Batch](data)
		if err != nil {
			return nil, err
		}
		return decodeData, nil
	}
	return nil, NotBatch
}

func (s *Storage) clearBatch(num uint64) error {
	s.VerifyBatchNumber = num
	key := fmt.Sprintf(string(batchKey), num)
	if err := s.db.Delete([]byte(key)); err != nil {
		return err
	}
	return nil
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
		s.VerifyBatchNumber = decodeData.VerifyBatchNumber
	}

	return nil
}
