package state

import (
	"bytes"
	"encoding/gob"
	"github.com/OAAC/state/types"
	"github.com/ethereum/go-ethereum/ethdb"
)

var AccountPersistenceKey = []byte("AccountPersistenceKey")

type Storage struct {
	Account *types.UserAccount
	db      ethdb.Database
}

func NewStorage(db ethdb.Database) *Storage {
	return &Storage{Account: types.NewUserAccount(), db: db}
}

func (s *Storage) persistence() error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(s.Account); err != nil {
		return err
	}
	if err := s.db.Put(AccountPersistenceKey, buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (s *Storage) loadDisk() error {
	has, err := s.db.Has(AccountPersistenceKey)
	if err != nil {
		return err
	}
	if has {
		accountData, err := s.db.Get(AccountPersistenceKey)
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer(accountData)
		decoder := gob.NewDecoder(buf)

		var account types.UserAccount
		if err = decoder.Decode(&account); err != nil {
			return err
		}
		s.Account = &account
	}

	return nil
}
