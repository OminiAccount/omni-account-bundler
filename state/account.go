package state

import (
	"fmt"
	"github.com/OAB/pool"
	"github.com/OAB/state/types"
	"github.com/OAB/utils/chains"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jackc/pgx/v4"
	"math/big"
)

type AccountMapping struct {
	User    common.Address
	Account common.Address
	Salt    uint64
}

type ChainInfo struct {
	NetworkID uint64
	Nonce     uint64
	Gas       *big.Int
}

type AccountInfo struct {
	Uid     uint64
	User    common.Address
	Account common.Address
	Salt    uint64
	Chain   map[uint64]ChainInfo
}

func (s *State) CreateAccountInfo(chainID chains.ChainId, am types.AccountMapping) {
	ui, err := s.db.GetUserInfoByAA(s.ctx, am.User.Hex(), am.Account.Hex(), nil)
	if err != nil {
		log.Errorf("[CreateAccountInfo] chainID: %d, user: %+v, err: %+v", chainID, am, err)
		return
	}
	err = s.db.AddUserInfo(s.ctx, ui.Uid, uint64(chainID), nil)
	if err != nil {
		log.Errorf("[CreateAccountInfo] chainID: %d, user: %+v, err: %+v", chainID, am, err)
		return
	}
}

func (s *State) AddSignedUserOp(ai *AccountInfo, suop *pool.SignedUserOperation) error {
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
	} else if suop.Did == "" {
		hashBytes := crypto.Keccak256Hash(suop.Encode())
		suop.Did = hashBytes.Hex()
		log.Infof("signedUserOperation did: %s", suop.Did)
		suop.Status = pool.PendingStatus
	}
	dbTx, err := s.db.BeginDBTransaction(s.ctx)
	if err != nil {
		log.Errorf("%+v, use db transaction err: %+v", suop.UserOperation, err)
		return err
	}
	tNonce, tGas, err := s.db.GetUserInfoLock(s.ctx, ai.Uid, suop.Exec.ChainId.Uint64(), dbTx)
	if err != nil {
		log.Errorf("%+v, chain: %d, get userinfo err: %+v", suop.UserOperation, suop.Exec.ChainId, err)
		_ = dbTx.Rollback(s.ctx)
		return err
	}
	if suop.OperationType == pool.WithdrawAction {
		if tGas.Cmp(suop.OperationValue.ToInt()) < 0 {
			_ = dbTx.Rollback(s.ctx)
			return fmt.Errorf("gas insufficient")
		} else {
			tGas = tGas.Sub(tGas, suop.OperationValue.ToInt())
		}
	}
	if suop.OperationType != pool.DepositAction {
		usedGas := suop.UserOperation.CalculateGasUsed()
		if tGas.Cmp(usedGas) < 0 {
			_ = dbTx.Rollback(s.ctx)
			return fmt.Errorf("gas insufficient")
		} else {
			tGas = tGas.Sub(tGas, usedGas)
		}
	}
	tNonce++
	if tNonce != suop.Exec.Nonce.Uint64() {
		_ = dbTx.Rollback(s.ctx)
		err = fmt.Errorf("account:%s nonce mismatch, want:%d, get:%d", ai.User, tNonce, suop.Exec.Nonce.Uint64())
		return err
	}
	err = s.db.ModUserInfo(s.ctx, ai.Uid, suop.Exec.ChainId.Uint64(), tNonce, tGas.String(), dbTx)
	if err != nil {
		_ = dbTx.Rollback(s.ctx)
		log.Errorf("%+v, chain: %d, mod userinfo err: %+v", suop.UserOperation, suop.Exec.ChainId, err)
		return err
	}
	if suop.InnerExec.ChainId.Uint64() > 0 {
		tNonce, _, err = s.db.GetUserInfoLock(s.ctx, ai.Uid, suop.InnerExec.ChainId.Uint64(), dbTx)
		if err != nil {
			_ = dbTx.Rollback(s.ctx)
			log.Errorf("%+v, chain: %d, get userinfo err: %+v", suop.UserOperation, suop.InnerExec.ChainId, err)
			return err
		}
		tNonce++
		if tNonce != suop.InnerExec.Nonce.Uint64() {
			_ = dbTx.Rollback(s.ctx)
			err = fmt.Errorf("account:%s nonce mismatch, want:%d, get:%d", ai.User, tNonce, suop.InnerExec.Nonce.Uint64())
			return err
		}
		err = s.db.ModUserInfo(s.ctx, ai.Uid, suop.Exec.ChainId.Uint64(), tNonce, "", dbTx)
		if err != nil {
			_ = dbTx.Rollback(s.ctx)
			log.Errorf("%+v, chain: %d, mod userinfo err: %+v", suop.UserOperation, suop.InnerExec.ChainId, err)
			return err
		}
	}
	err = s.db.AddOperation(s.ctx, ai.Uid, suop, dbTx)
	if err != nil {
		_ = dbTx.Rollback(s.ctx)
		log.Errorf("%+v, add op err: %+v", suop.UserOperation, err)
		return err
	}
	err = dbTx.Commit(s.ctx)
	return err
}

func (s *State) AddAccountGas(suop *pool.SignedUserOperation, dbTx pgx.Tx) error {
	var (
		err  error
		flag bool
	)
	if dbTx == nil {
		flag = true
		dbTx, err = s.db.BeginDBTransaction(s.ctx)
		if err != nil {
			log.Errorf("[AddAccountGas] %+v, use db transaction err: %+v", suop.UserOperation, err)
			return err
		}
	}
	_, tGas, err1 := s.db.GetUserInfoLock(s.ctx, suop.Uid, suop.Exec.ChainId.Uint64(), dbTx)
	if err1 != nil {
		log.Errorf("[AddAccountGas] %+v, chain: %d, get userinfo err: %+v", suop.UserOperation, suop.Exec.ChainId, err1)
		if flag {
			_ = dbTx.Rollback(s.ctx)
		}
		return err1
	}
	tGas = tGas.Add(tGas, suop.OperationValue.ToInt())
	err = s.db.ModUserInfo(s.ctx, suop.Uid, suop.Exec.ChainId.Uint64(), 0, tGas.String(), dbTx)
	if err != nil {
		if flag {
			_ = dbTx.Rollback(s.ctx)
		}
		log.Errorf("[AddAccountGas] %+v, chain: %d, mod userinfo err: %+v", suop.UserOperation, suop.Exec.ChainId, err)
		return err
	}
	if flag {
		err = dbTx.Commit(s.ctx)
	}
	return err
}

func (s *State) UpdateUserOpStatus(opid uint64, status int, dbTx pgx.Tx) error {
	err := s.db.UpdateUserOp(s.ctx, opid, "status", status, dbTx)
	return err
}

func (s *State) UpdateUserOpPhase(opid uint64, phase int, dbTx pgx.Tx) error {
	err := s.db.UpdateUserOp(s.ctx, opid, "phase", phase, dbTx)
	return err
}

func (s *State) GetSignedUserOp(user, account common.Address, did string, status int) (*pool.SignedUserOperation, error) {
	op, err := s.db.GetSigleUserOp(s.ctx, user.Hex(), account.Hex(), did, status, nil)
	if err != nil {
		return nil, err
	}
	return &pool.SignedUserOperation{UserOperation: op}, nil
}

func (s *State) GetAccountAdr(user common.Address) *common.Address {
	adr, err := s.db.GetAccountAdr(s.ctx, user.Hex(), nil)
	if err != nil {
		log.Errorf("GetAccountAdr, user: %s, get user account adr err: %+v", user, err)
		return nil
	}
	return adr
}

func (s *State) GetAccountInfo(user, account common.Address) (*AccountInfo, error) {
	ai, err := s.db.GetUserInfoByAA(s.ctx, user.Hex(), account.Hex(), nil)
	if err != nil {
		log.Errorf("GetAccountInfoByAA, user: %s, acc: %s, get user info err: %+v", user, account, err)
		return nil, err
	}
	return ai, nil
}

func (s *State) GetAccountOps(uid uint64) ([]*pool.UserOperation, error) {
	uos, err := s.db.GetUserOps(s.ctx, uid, nil)
	if err != nil {
		log.Errorf("GetAccountOps, user id: %d, get user op err: %+v", uid, err)
		return nil, err
	}
	return uos, nil
}

func (s *State) CreateAccount(user common.Address) *common.Address {
	adr, salt := s.ethereum.CreateAccount(user)
	if adr == nil {
		return nil
	}
	log.Infof("user: %s, account: %s, salt: %d", user, adr, salt)
	err := s.db.AddUser(s.ctx, AccountMapping{
		User:    user,
		Account: *adr,
		Salt:    salt,
	}, nil)
	if err != nil {
		log.Error("Add a new account mapping error", "error", err)
	}
	log.Infof("[CreateAccount] add new account success %s", adr)
	return adr
}
