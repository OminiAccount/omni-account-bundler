package state

import (
	"fmt"
	"github.com/OAB/pool"
	"github.com/OAB/state/types"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (s *State) AddNewMapping(mapping types.AccountMapping) error {
	return s.storage.Account.AddNewMapping(mapping)
}

func (s *State) AddSignedUserOperation(signedUserOp *pool.SignedUserOperation) error {
	if signedUserOp.OperationType == pool.DepositAction {
		if signedUserOp.Did == "" {
			return fmt.Errorf("invalid signedUserOperation, did mismatch")
		}
		s.pool.AddTicket(&pool.TicketFull{
			SignedUserOp: signedUserOp,
			Ticket: pool.Ticket{
				Did:    signedUserOp.Did,
				User:   signedUserOp.Sender,
				Amount: signedUserOp.OperationValue.ToInt(),
			},
			Type: pool.Deposit,
		})
		return nil
	}
	log.Debug("Check SignedUserOperation")
	if err := s.storage.Account.AddSignedUserOperation(signedUserOp); err != nil {
		return fmt.Errorf("check SignedUserOperation failed, error: %s", err.Error())
	}

	s.pool.AddSignedUserOperation(signedUserOp)

	return nil
}

func (s *State) AddAccountGas(signedUserOp *pool.SignedUserOperation) error {
	err := s.storage.Account.AddSignedUserOperation(signedUserOp)
	if err != nil {
		return err
	}
	return nil
}

func (s *State) GetSignedUserOp(user, account common.Address, id string) (*pool.SignedUserOperation, error) {
	accInfo, err := s.storage.Account.GetAccount(user, account)
	if err != nil {
		return nil, err
	}
	for _, uop := range accInfo.UserOperations {
		if uop.Did == id {
			return &pool.SignedUserOperation{UserOperation: uop}, nil
		}
	}
	return nil, nil
}

func (s *State) GetAccountsForUser(user common.Address) *[]common.Address {
	return s.storage.Account.GetAccountsForUser(user)
}

func (s *State) GetAccountInfo(user, account common.Address, chainId uint64) (*big.Int, uint64, uint64, []*pool.UserOperation, error) {
	accInfo, err := s.storage.Account.GetAccount(user, account)
	if err != nil {
		return nil, 0, 0, nil, err
	}
	return accInfo.Gas, accInfo.Nonce[chainId], accInfo.HistoryPage, accInfo.UserOperations, nil
}

func (s *State) CreateAccount(user common.Address) *common.Address {
	return s.ethereum.CreateAccount(user)
}