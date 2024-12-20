package pool

import (
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/msgpack"
	"github.com/ethereum/go-ethereum/ethdb"
)

var (
	UserOpsPersistenceKey = []byte("UserOpsKey")
)

type Storage struct {
	UserOps []*SignedUserOperation

	db ethdb.Database
}

func NewStorage(db ethdb.Database) *Storage {
	return &Storage{
		UserOps: []*SignedUserOperation{},
		db:      db,
	}
}

func (s *Storage) addUserOp(userOp *SignedUserOperation) {
	s.UserOps = append(s.UserOps, userOp)
}

func (s *Storage) getUserOps() []*SignedUserOperation {
	return s.UserOps
}

func (s *Storage) empty() {
	s.UserOps = []*SignedUserOperation{}
}

func (s *Storage) cacheUserOps() error {
	data, err := msgpack.MarshalStruct(s)
	if err != nil {
		return err
	}

	if err := s.db.Put(UserOpsPersistenceKey, data); err != nil {
		return err
	}

	return nil
}

func (s *Storage) loadUserOps() error {
	has, err := s.db.Has(UserOpsPersistenceKey)
	if err != nil {
		return err
	}
	if has {
		userOpsData, err := s.db.Get(UserOpsPersistenceKey)
		if err != nil {
			return err
		}

		decodeSigUserOps, err := msgpack.UnmarshalStruct[Storage](userOpsData)
		if err != nil {
			return err
		}
		s.UserOps = decodeSigUserOps.UserOps

		log.Info("load cache userOps length ", len(s.UserOps))
	}
	return nil
}
