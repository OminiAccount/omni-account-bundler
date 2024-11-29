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
	UserOps []*SignedUserOperation
	Tickets []*TicketFull

	db ethdb.Database
}

func NewStorage(db ethdb.Database) *Storage {
	return &Storage{
		UserOps: []*SignedUserOperation{},
		Tickets: []*TicketFull{},
		db:      db,
	}
}

func (s *Storage) addTicket(ticket *TicketFull) {
	s.Tickets = append(s.Tickets, ticket)
}

func (s *Storage) getTickets() []*TicketFull {
	return s.Tickets
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
		s.Tickets = decodeSigUserOps.Tickets

		fmt.Println("load cache userOps length ", len(s.UserOps))
		//s.userOps = []SignedUserOperation{}
	}
	return nil
}
