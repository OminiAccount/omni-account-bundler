package state

import (
	"context"
	"errors"
	"fmt"
	"github.com/OAB/utils"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/common"
	ether_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"strconv"
	"strings"
	"time"
)

const (
	CheckFromTxStatus = iota
	CheckFromCrossTxStatus
	CheckToTxStatus
	CheckToCrossTxStatus
	CheckSuccessStatus
	CheckFailedStatus
)

type HistoryManager struct {
	ctx    context.Context
	cancel context.CancelFunc
	cfg    Config
	state  *State
}

func NewHistoryManager(ctx context.Context, cfg Config, s *State) *HistoryManager {
	hisCtx, cancel := context.WithCancel(ctx)
	return &HistoryManager{
		ctx:    hisCtx,
		cancel: cancel,
		cfg:    cfg,
		state:  s,
	}
}

func (h *HistoryManager) Start() {
	log.Infof("history manager start")
	ticker := time.NewTicker(h.state.cfg.HisInterval.Duration)
	for {
		select {
		case <-h.ctx.Done():
			log.Warn("history manager exit")
			return
		case <-ticker.C:
			err := h.scanHis()
			if err != nil {
				log.Errorf("scan history failed: %v", err)
			}
		}
	}
}

func (h *HistoryManager) scanHis() error {
	pHis, err := h.state.db.GetHistoryByStatus(h.ctx,
		[]int{CheckFromTxStatus, CheckFromCrossTxStatus, CheckToTxStatus, CheckToCrossTxStatus}, 100, nil)
	if err != nil {
		return err
	}
	if pHis == nil {
		return nil
	}
	log.Infof("found %v pending history to process", len(pHis))
	for _, v := range pHis {
		mTx := v
		diff := time.Now().Unix() - mTx.TimeAt.Unix()
		if diff > (24*60*60) && mTx.SourceHash != "" && mTx.TargetHash != "" { // 24h
			_ = h.state.db.UpdateHistory(h.ctx, mTx.ID, map[string]interface{}{"status": CheckFailedStatus}, nil)
			continue
		}
		switch mTx.Status {
		case CheckFromTxStatus:
			h.checkHisFromTx(mTx)
		case CheckFromCrossTxStatus:
			h.checkHisFromCrossTx(mTx)
		case CheckToTxStatus:
			h.checkHisToTx(mTx)
		case CheckToCrossTxStatus:
			h.checkHisToCrossTx(mTx)
		}
	}
	return nil
}

func (h *HistoryManager) checkHisFromTx(his AccountHistory) {
	if his.SourceHash == "" {
		return
	}
	l2Node := h.state.ethereum.GetChainCli(his.SourceChain)
	tx, isPending, err := l2Node.Cli().TransactionByHash(h.ctx, common.HexToHash(his.SourceHash))
	if err != nil {
		log.Errorf("[HistoryManager] error getting txByHash %s. Error: %v", his.SourceHash, err)
		return
	}
	if tx == nil || isPending {
		log.Infof("[HistoryManager] tx: %s not mined yet", his.SourceHash)
		return
	}
	signer := ether_types.NewEIP155Signer(tx.ChainId())
	from, err := signer.Sender(tx)
	if err != nil {
		log.Infof("[HistoryManager] from signer get sender err: %v", err)
		return
	}
	if his.From.Address != strings.ToLower(from.Hex()) {
		log.Warnf("[HistoryManager] tx from address(%s / %s) mismatch", his.From.Address, from.Hex())
		//return true, fmt.Errorf("tx from address mismatch")
	}
	if his.From.Value != tx.Value().String() {
		log.Warnf("[HistoryManager] tx value(%s / %s) mismatch", his.From.Value, tx.Value().String())
		//return true, fmt.Errorf("tx value mismatch")
	}
	_ = h.state.db.UpdateHistory(h.ctx, his.ID, map[string]interface{}{"status": CheckFromCrossTxStatus}, nil)
}

func (h *HistoryManager) checkHisFromCrossTx(his AccountHistory) {
	if his.OrderType == DepositAction ||
		his.OrderType == WithdrawAction {
		_ = h.state.db.UpdateHistory(h.ctx, his.ID, map[string]interface{}{"status": CheckToTxStatus}, nil)
		return
	}
	/*crossTx := his.SourceHash
	targetTx, err := h.getCrossTx(crossTx, his.SourceChain)
	if err != nil {
		log.Errorf("get cross tx err: %+v", err)
		return
	}
	if targetTx == "" {
		return
	}
	dbTx, err := h.state.db.BeginDBTransaction(h.ctx)
	if err != nil {
		log.Errorf("use db transaction err: %+v", err)
		return
	}
	err = h.state.db.UpdateHistory(h.ctx, his.ID, map[string]interface{}{
		"source_hash": targetTx,
		"status":      CheckToTxStatus,
	}, dbTx)
	if err != nil {
		_ = dbTx.Rollback(h.ctx)
		return
	}
	_ = dbTx.Commit(h.ctx)*/
}

func (h *HistoryManager) checkHisToTx(his AccountHistory) {
	if his.TargetHash == "" {
		return
	}
	l2Node := h.state.ethereum.GetChainCli(his.TargetChain)
	tx, isPending, err := l2Node.Cli().TransactionByHash(h.ctx, common.HexToHash(his.TargetHash))
	if err != nil {
		log.Errorf("[HistoryManager] error getting txByHash %s. Error: %v", his.SourceHash, err)
		return
	}
	if tx == nil || isPending {
		log.Infof("[HistoryManager] tx: %s not mined yet", his.SourceHash)
		return
	}
	signer := ether_types.NewEIP155Signer(tx.ChainId())
	from, err := signer.Sender(tx)
	if err != nil {
		log.Infof("[HistoryManager] from signer get sender err: %v", err)
		return
	}
	if his.From.Address != strings.ToLower(from.Hex()) {
		log.Warnf("[HistoryManager] tx from address(%s / %s) mismatch", his.From.Address, from.Hex())
		//return true, fmt.Errorf("tx from address mismatch")
	}
	if his.From.Value != tx.Value().String() {
		log.Warnf("[HistoryManager] tx value(%s / %s) mismatch", his.From.Value, tx.Value().String())
		//return true, fmt.Errorf("tx value mismatch")
	}
	_ = h.state.db.UpdateHistory(h.ctx, his.ID, map[string]interface{}{"status": CheckToCrossTxStatus}, nil)
}

func (h *HistoryManager) checkHisToCrossTx(his AccountHistory) {
	if his.OrderType == DepositAction ||
		his.OrderType == WithdrawAction {
		_ = h.state.db.UpdateHistory(h.ctx, his.ID, map[string]interface{}{"status": CheckSuccessStatus}, nil)
		return
	}
	/*crossTx := his.TargetHash
	targetTx, err := h.getCrossTx(crossTx, his.TargetChain)
	if err != nil {
		log.Errorf("get cross tx err: %+v", err)
		return
	}
	if targetTx == "" {
		return
	}
	dbTx, err := h.state.db.BeginDBTransaction(h.ctx)
	if err != nil {
		log.Errorf("use db transaction err: %+v", err)
		return
	}
	err = h.state.db.UpdateHistory(h.ctx, his.ID, map[string]interface{}{
		"target_hash": targetTx,
		"status":      CheckSuccessStatus,
	}, dbTx)
	if err != nil {
		_ = dbTx.Rollback(h.ctx)
		return
	}
	_ = dbTx.Commit(h.ctx)*/
}

func (h *HistoryManager) UpdateAccountHis(uid uint64, did string, m map[string]interface{}) error {
	his, err := h.state.db.GetHistoryByDid(h.ctx, uid, did, nil)
	if err != nil {
		return err
	}
	err = h.state.db.UpdateHistory(h.ctx, his.ID, m, nil)
	return err
}

func (h *HistoryManager) GetAccountHis(uid, ts, size uint64) ([]AccountHistory, error) {
	list, err := h.state.db.GetHistoryByTs(h.ctx, uid, size, time.Unix(int64(ts), 0), nil)
	if err != nil {
		log.Errorf("get account history err: %+v", err)
	}
	return list, err
}

func (h *HistoryManager) SaveAccountHis(data *AccountHistory) error {
	_, err := h.state.db.GetHistoryByDid(h.ctx, data.Uid, data.Did, nil)
	if !errors.Is(err, pgx.ErrNoRows) {
		log.Errorf("history already exist: %+v", data)
		return errors.New("history already exist")
	}
	err = h.state.db.AddUserHis(h.ctx, data, nil)
	return err
}

const Cross_FAILED = 97
const Cross_SUCCESS = 99

type TxStatus struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Result  []ResultData `json:"result"`
}

type ResultData struct {
	SourceChain  string    `json:"sourceChain"`
	SourceId     string    `json:"sourceId"`
	SourceAmount string    `json:"sourceAmount"`
	SourceSymbol string    `json:"sourceSymbol"`
	SourceTime   time.Time `json:"sourceTime"`
	TargetChain  string    `json:"targetChain"`
	TargetId     string    `json:"targetId"`
	TargetAmount string    `json:"targetAmount"`
	TargetSymbol string    `json:"targetSymbol"`
	TargetTime   time.Time `json:"targetTime"`
	Status       int       `json:"status"`
}

func (h *HistoryManager) getCrossTx(tx string, tId uint64) (string, error) {
	r := utils.NewHTTPCli()
	ret, err := r.Get(h.cfg.CrossChainAPI + tx)
	if err != nil {
		return "", err
	}
	retRes := &TxStatus{}
	err = ret.Parse(retRes)
	if err != nil {
		return "", err
	}
	log.Debugf("cross tx result: %+v", retRes)
	if retRes.Status != "success" {
		return "", fmt.Errorf("cross error")
	}
	for _, rd := range retRes.Result {
		if strconv.Itoa(int(tId)) == rd.TargetChain {
			targetHash := ""
			hashs := strings.Split(rd.TargetId, "-")
			if len(hashs) > 0 {
				targetHash = hashs[0]
			}
			if targetHash == "" || targetHash[:2] != "0x" || len(targetHash) < 30 {
				continue
			}
			return targetHash, nil
		}
	}
	return "", nil
}
