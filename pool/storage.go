package pool

import (
	"fmt"
	"github.com/OAB/utils/msgpack"
	"github.com/ethereum/go-ethereum/ethdb"
)

var (
	UserOpsPersistenceKey = []byte("UserOpsKey")
)

type Storage struct {
	userOps []*SignedUserOperation

	db ethdb.Database
}

func NewStorage(db ethdb.Database) *Storage {
	return &Storage{
		userOps: []*SignedUserOperation{},
		db:      db,
	}
}

func (s *Storage) addUserOp(userOp *SignedUserOperation) {
	s.userOps = append(s.userOps, userOp)
}

func (s *Storage) getUserOps() []*SignedUserOperation {
	return s.userOps
}

func (s *Storage) empty() {
	s.userOps = []*SignedUserOperation{}
}

func (s *Storage) cacheUserOps() error {
	data, err := msgpack.MarshalStruct(s.userOps)
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

		decodeSigUserOps, err := msgpack.UnmarshalStruct[[]*SignedUserOperation](userOpsData)
		if err != nil {
			return err
		}
		s.userOps = decodeSigUserOps
		fmt.Println("load cache userOps length ", len(s.userOps))
		//s.userOps = []SignedUserOperation{}
	}
	return nil
}
