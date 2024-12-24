package state

import (
	"errors"
	"fmt"
	"github.com/OAB/etherman"
	"github.com/OAB/utils/chains"
	"github.com/OAB/utils/hex"
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/packutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jackc/pgx/v4"
	"math/big"
	"time"
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

type Token struct {
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Logo     string `json:"logo"`
}

type HistoryDetail struct {
	Token   Token  `json:"token"`
	Address string `json:"address"`
	Value   string `json:"value"`
}

type AccountHistory struct {
	ID          uint64        `json:"id"`
	Uid         uint64        `json:"uid"`
	Did         string        `json:"did"`
	OrderType   uint8         `json:"orderType"`
	From        HistoryDetail `json:"from"`
	To          HistoryDetail `json:"to"`
	SourceChain uint64        `json:"sourceChain"`
	TargetChain uint64        `json:"targetChain"`
	SourceHash  string        `json:"sourceHash"`
	TargetHash  string        `json:"targetHash"`
	SwapHash    string        `json:"swapHash"`
	Status      uint64        `json:"status"`
	HisTime     int64         `json:"time"`
	TimeAt      time.Time     `json:"timeAt"`
}

func (a AccountHistory) UniqueID() string {
	hashBytes := crypto.Keccak256Hash(a.Encode())
	return hashBytes.Hex()
}

func (a AccountHistory) Encode() []byte {
	var b []byte
	b = append(b, packutils.Uint64ToBytes(a.Uid)...)
	b = append(b, a.OrderType)
	b = append(b, packutils.Uint64ToBytes(a.SourceChain)...)
	b = append(b, packutils.Uint64ToBytes(a.TargetChain)...)
	b = append(b, common.LeftPadBytes(common.HexToAddress(a.From.Address).Bytes(), 32)...)
	b = append(b, common.LeftPadBytes(common.Hex2Bytes(a.From.Value), 32)...)
	b = append(b, common.LeftPadBytes(common.HexToAddress(a.To.Address).Bytes(), 32)...)
	b = append(b, common.LeftPadBytes(common.Hex2Bytes(a.To.Value), 32)...)
	b = append(b, packutils.Uint64ToBytes(uint64(a.HisTime))...)
	return b
}

func ToUserOperationHis(tx string, op *UserOperation) AccountHistory {
	uoHis := AccountHistory{
		Uid:       op.Uid,
		Did:       op.Did,
		OrderType: op.OperationType,
		From: HistoryDetail{
			Address: op.Sender.Hex(),
			Value:   hex.EncodeBig(op.OperationValue.ToInt()),
		},
		SourceChain: op.Exec.ChainId.Uint64(),
		SourceHash:  tx,
		To: HistoryDetail{
			Address: op.Sender.Hex(),
			Value:   hex.EncodeBig(op.OperationValue.ToInt()),
		},
		TargetChain: op.Exec.ChainId.Uint64(),
		TargetHash:  tx,
		HisTime:     time.Now().Unix(),
		TimeAt:      time.Now(),
	}
	if op.Status == SuccessStatus {
		uoHis.Status = CheckSuccessStatus
	}
	return uoHis
}

func (s *State) CreateAccountInfo(chainID chains.ChainId, am AccountMapping) {
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

func getDepositID(suop *SignedUserOperation) string {
	var encodeBytes []byte
	encodeBytes = append(encodeBytes, suop.Sender.Bytes()...)
	encodeBytes = append(encodeBytes, common.LeftPadBytes(big.NewInt(int64(suop.Exec.ChainId)).Bytes(), 32)...)
	encodeBytes = append(encodeBytes, common.LeftPadBytes(big.NewInt(int64(suop.Exec.Nonce)).Bytes(), 32)...)
	encodeBytes = append(encodeBytes, common.LeftPadBytes(suop.OperationValue.ToInt().Bytes(), 32)...)
	return crypto.Keccak256Hash(encodeBytes).Hex()
}

func (s *State) AddSignedUserOp(ai *AccountInfo, suop *SignedUserOperation) error {
	if suop.OperationType == DepositAction {
		if suop.Did == "" {
			suop.Did = getDepositID(suop)
			log.Infof("deposit did: %s", suop.Did)
		}
	} else if suop.Did == "" {
		hashBytes := crypto.Keccak256Hash(suop.Encode())
		suop.Did = hashBytes.Hex()
		log.Infof("signedUserOperation did: %s", suop.Did)
		suop.Status = PendingStatus
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
	if suop.OperationType == WithdrawAction {
		if tGas.Cmp(suop.OperationValue.ToInt()) < 0 {
			_ = dbTx.Rollback(s.ctx)
			return fmt.Errorf("gas insufficient")
		} else {
			tGas = tGas.Sub(tGas, suop.OperationValue.ToInt())
		}
	}
	if suop.OperationType != DepositAction {
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
	err = s.db.ModUserInfo(s.ctx, ai.Uid, suop.Exec.ChainId.Uint64(), tNonce, hex.EncodeBig(tGas), dbTx)
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

func (s *State) AddAccountGas(suop *SignedUserOperation, dbTx pgx.Tx) error {
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
	err = s.db.ModUserInfo(s.ctx, suop.Uid, suop.Exec.ChainId.Uint64(), 0, hex.EncodeBig(tGas), dbTx)
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

func (s *State) GetSignedUserOp(user, account common.Address, did string, status int) (*SignedUserOperation, error) {
	op, err := s.db.GetSigleUserOp(s.ctx, user.Hex(), account.Hex(), did, status, nil)
	if err != nil {
		return nil, err
	}
	return &SignedUserOperation{UserOperation: op}, nil
}

func (s *State) GetAccountAdr(user common.Address) *common.Address {
	adr, err := s.db.GetAccountAdr(s.ctx, user.Hex(), nil)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Errorf("GetAccountAdr, user: %s, get user account adr err: %+v", user, err)
		}
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

func (s *State) GetAccountOps(uid uint64) ([]*UserOperation, error) {
	uos, err := s.db.GetUserOps(s.ctx, uid, nil)
	if err != nil {
		log.Errorf("GetAccountOps, user id: %d, get user op err: %+v", uid, err)
		return nil, err
	}
	return uos, nil
}

func (s *State) CreateAccount(user common.Address) *common.Address {
	adr, salt, err := s.ethereum.CreateAccount(user)
	if errors.Is(err, etherman.ErrExistAccount) {
		return adr
	}
	if adr == nil {
		return nil
	}
	log.Infof("user: %s, account: %s, salt: %d", user, adr, salt)
	err = s.db.AddUser(s.ctx, AccountMapping{
		User:    user,
		Account: *adr,
		Salt:    salt,
	}, nil)
	if err != nil {
		log.Errorf("Add a new account mapping err: %+v", err)
	}
	log.Infof("[CreateAccount] add new account success %s", adr)
	return adr
}

func (s *State) AddFailedCreateAA(uid, nid uint64) {
	err := s.db.AddUserFailedSalt(s.ctx, uid, nid, nil)
	if err != nil {
		log.Errorf("Add user failed salt, uid: %d, nid: %d err: %+v", uid, nid, err)
	}
}
