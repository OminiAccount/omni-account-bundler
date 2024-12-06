package state

import (
	"fmt"
	"github.com/OAB/state/types"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/common"
)

func (s *State) GetAccountHis(user, account common.Address, page uint64) ([]types.AccountHistory, error) {
	return s.storage.loadUserHistory(account, page)
}

func (s *State) SaveAccountHis(owner, account common.Address, data types.AccountHistory) error {
	user, err := s.storage.Account.GetAccount(owner, account)
	if err != nil || user == nil {
		log.Errorf("get account(%s, %s) no data, error: %v", owner, account, err)
		return fmt.Errorf("get account(%s, %s) no data", owner, account)
	}
	his, err := s.storage.loadUserHistory(account, user.HistoryPage)
	if (err != nil && err != NotHis) || (user.HistoryPage > 0 && his == nil) {
		log.Errorf("get account(%s, %s) history(%d) no data, err: %v", owner, account, user.HistoryPage, err)
		return fmt.Errorf("get account(%s, %s) no data", owner, account)
	}
	if len(his) >= types.HistoryPageSize {
		user.HistoryPage++
		his = []types.AccountHistory{}
	}
	if his == nil {
		his = []types.AccountHistory{}
	}
	his = append(his, data)
	err = s.storage.cacheUserHistory(account, user.HistoryPage, his)
	if err != nil {
		log.Errorf("cache account(%s, %s) history error: %v", err)
	}
	return err
}

