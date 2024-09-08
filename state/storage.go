package state

import (
	"github.com/OAB/state/types"
	"github.com/OAB/utils/msgpack"
	"github.com/OAB/utils/smt"
	"github.com/ethereum/go-ethereum/ethdb"
)

var (
	AccountPersistenceKey = []byte("AccountKey")
	SmtPersistenceKey     = []byte("SmtKey")
)

type Storage struct {
	Account *types.UserAccount
	db      ethdb.Database
}

func NewStorage(db ethdb.Database) *Storage {
	return &Storage{Account: types.NewUserAccount(), db: db}
}

func (s *Storage) cache() error {
	data, err := msgpack.MarshalStruct(s.Account)
	if err != nil {
		return err
	}

	if err := s.db.Put(AccountPersistenceKey, data); err != nil {
		return err
	}
	return nil
}

func (s *Storage) loadCache() error {
	has, err := s.db.Has(AccountPersistenceKey)
	if err != nil {
		return err
	}
	if has {
		accountData, err := s.db.Get(AccountPersistenceKey)
		if err != nil {
			return err
		}

		decodeData, err := msgpack.UnmarshalStruct[types.UserAccount](accountData)
		if err != nil {
			return err
		}
		s.Account = &decodeData
	}

	return nil
}

func (s *Storage) cacheSmt(tree *smt.ZeroMerkleTree) error {
	data, err := msgpack.MarshalStruct(tree)
	if err != nil {
		return err
	}

	if err := s.db.Put(SmtPersistenceKey, data); err != nil {
		return err
	}
	return nil
}

func (s *Storage) loadCacheSmt() (*smt.ZeroMerkleTree, error) {
	has, err := s.db.Has(SmtPersistenceKey)
	if err != nil {
		return nil, err
	}
	if has {
		treeData, err := s.db.Get(SmtPersistenceKey)
		if err != nil {
			return nil, err
		}

		decodeData, err := msgpack.UnmarshalStruct[*smt.ZeroMerkleTree](treeData)
		if err != nil {
			return nil, err
		}
		return decodeData, nil
	}

	return nil, nil
}
