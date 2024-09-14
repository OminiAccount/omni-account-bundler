package state

import (
	"github.com/OAB/state/types"
	"github.com/OAB/utils/msgpack"
	"github.com/OAB/utils/smt"
	"github.com/ethereum/go-ethereum/ethdb"
)

var (
	storageKey        = []byte("StateContextKey")
	SmtPersistenceKey = []byte("SmtKey")
)

type Storage struct {
	Account     *types.UserAccount
	BatchNumber uint64
	StateProof  map[uint64]*smt.DeltaMerkleProof
	db          ethdb.Database
}

func NewStorage(db ethdb.Database) *Storage {
	return &Storage{Account: types.NewUserAccount(), BatchNumber: 0, StateProof: make(map[uint64]*smt.DeltaMerkleProof), db: db}
}

func (s *Storage) updateStateDiff(diff *smt.DeltaMerkleProof) {
	s.StateProof[s.BatchNumber] = diff
	s.BatchNumber++
}

func (s *Storage) cache() error {
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
		s.StateProof = decodeData.StateProof
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
