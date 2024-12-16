package state

import (
	"fmt"
	"github.com/OAB/pool"
	"github.com/OAB/state/types"
	"github.com/OAB/utils/chains"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func (s *State) InitAccountNonce(chainID chains.ChainId, am types.AccountMapping) {
	s.storage.Account.InitNonce(chainID, am)
}

func (s *State) AddNewMapping(mapping types.AccountMapping) error {
	return s.storage.Account.AddNewMapping(mapping)
}

func (s *State) AddSignedUserOperation(suop *pool.SignedUserOperation) error {
	if suop.OperationType == pool.DepositAction {
		if suop.Did == "" {
			var encodeBytes []byte
			encodeBytes = append(encodeBytes, suop.Sender.Bytes()...)
			encodeBytes = append(encodeBytes, common.LeftPadBytes(big.NewInt(int64(suop.Exec.ChainId)).Bytes(), 32)...)
			encodeBytes = append(encodeBytes, common.LeftPadBytes(big.NewInt(int64(suop.Exec.Nonce)).Bytes(), 32)...)
			encodeBytes = append(encodeBytes, common.LeftPadBytes(suop.OperationValue.ToInt().Bytes(), 32)...)
			suop.Did = crypto.Keccak256Hash(encodeBytes).Hex()
			log.Infof("deposit did: %s", suop.Did)
		}
		s.pool.AddTicket(&pool.TicketFull{
			SignedUserOp: suop,
			Ticket: pool.Ticket{
				Did:    suop.Did,
				User:   suop.Sender,
				Amount: suop.OperationValue.ToInt(),
			},
			Type: pool.Deposit,
		})
		return nil
	}
	log.Debug("Check SignedUserOperation")
	if suop.Did == "" {
		hashBytes := crypto.Keccak256Hash(suop.Encode())
		suop.Did = hashBytes.Hex()
		log.Infof("signedUserOperation did: %s", suop.Did)
	}
	if err := s.storage.Account.AddSignedUserOperation(suop); err != nil {
		return fmt.Errorf("check SignedUserOperation failed, error: %s", err.Error())
	}

	s.pool.AddSignedUserOperation(suop)

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

func (s *State) GetAccountAdrs(user common.Address) *[]common.Address {
	accountsInfo := s.storage.Account.GetAccountsForUser(user)
	var accounts []common.Address
	for account := range accountsInfo {
		accounts = append(accounts, account)
	}
	return &accounts
}

func (s *State) GetAccountInfo(user common.Address) (*common.Address, *types.AccountInfo) {
	accountsInfo := s.storage.Account.GetAccountsForUser(user)
	for adr, info := range accountsInfo {
		return &adr, &info
	}
	return nil, nil
}

func (s *State) GetAccountInfoByAA(user, account common.Address) (*types.AccountInfo, error) {
	accInfo, err := s.storage.Account.GetAccount(user, account)
	if err != nil {
		return nil, err
	}
	return accInfo, nil
}

func (s *State) CreateAccount(user common.Address) *common.Address {
	adr, salt := s.ethereum.CreateAccount(user)
	if adr == nil {
		return nil
	}
	log.Infof("user: %s, account: %s, salt: %d", user, adr, salt)
	err := s.AddNewMapping(types.AccountMapping{
		User:    user,
		Account: *adr,
		Salt:    salt,
	})
	if err != nil {
		log.Error("Add a new account mapping error", "error", err)
	}
	log.Infof("[CreateAccount] add new account success %s", adr)
	return adr
}
