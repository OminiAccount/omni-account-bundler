package state

import (
	"context"
	"errors"
	"fmt"
	"github.com/OAB/pool"
	"github.com/OAB/state/types"
	"github.com/OAB/utils/chains"
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/msgpack"
	"github.com/ethereum/go-ethereum/common"
	ether_types "github.com/ethereum/go-ethereum/core/types"
	"strconv"
	"sync"
	"time"
)

type HistoryManager struct {
	ctx     context.Context
	cancel  context.CancelFunc
	cfg     Config
	state   *State
	storage *Storage
	txMgr   *HisInfra
	chHis   chan *types.AccountHistory
}

func NewHistoryManager(ctx context.Context, cfg Config, s *State) *HistoryManager {
	hisCtx, cancel := context.WithCancel(ctx)
	return &HistoryManager{
		ctx:     hisCtx,
		cancel:  cancel,
		cfg:     cfg,
		state:   s,
		storage: s.storage,
		txMgr:   NewHisInfra(hisCtx, cfg, s),
		chHis:   make(chan *types.AccountHistory, 10000),
	}
}

func (h *HistoryManager) Start() {
	log.Infof("history manager start")
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-h.ctx.Done():
			if len(h.chHis) > 0 {
				continue
			}
			log.Warn("history manager exit")
			return
		case his := <-h.chHis:
			log.Infof("checkHis: %+v", his)
			isDel, err := h.checkHis(context.Background(), his)
			if err != nil {
				log.Errorf("checkHis failed, data: %v, error: %v", his, err)
			}
			if isDel {
				err = h.txMgr.DelExistHis(his)
				if err != nil {
					log.Errorf("del exist history error: %v", err)
				}
			}
		case <-ticker.C:
			err := h.scanHis(context.Background())
			if err != nil {
				log.Errorf("scan history failed: %v", err)
			}
		}
	}
}

func (h *HistoryManager) scanHis(ctx context.Context) error {
	pHis, err := h.txMgr.GetPendingHis()
	if err != nil {
		return err
	}
	if pHis == nil {
		return nil
	}
	log.Infof("found %v pending history to process", len(pHis))
	for _, v := range pHis {
		mTx := v
		diff := time.Now().Unix() - mTx.TimeAt
		if diff > (60 * 60) { // 1h
			err = h.txMgr.DelExistHis(mTx.AccountHistory)
			if err != nil {
				return err
			}
			continue
		}
		isDel, err := h.checkHis(ctx, mTx.AccountHistory)
		if isDel {
			err = h.txMgr.DelExistHis(mTx.AccountHistory)
			if err != nil {
				log.Errorf("del exist history error: %v", err)
			}
		}
	}
	return nil
}

func (h *HistoryManager) checkHis(ctx context.Context, his *types.AccountHistory) (bool, error) {
	l2Node := h.state.ethereum.GetChainCli(chains.ChainId(his.SourceChain))
	tx, isPending, err := l2Node.Cli().TransactionByHash(ctx, common.HexToHash(his.SourceHash))
	if err != nil {
		log.Errorf("[HistoryManager] error getting txByHash %s. Error: %v", his.SourceHash, err)
		return false, err
	}
	if tx == nil || isPending {
		log.Infof("[HistoryManager] tx: %s not mined yet", his.SourceHash)
		return false, nil
	}
	signer := ether_types.NewEIP155Signer(tx.ChainId())
	from, err := signer.Sender(tx)
	if err != nil {
		log.Infof("[HistoryManager] from signer get sender err: %v", err)
		return false, err
	}
	if his.From.Address != from.Hex() {
		log.Infof("[HistoryManager] tx from address(%s / %s) mismatch", his.From.Address, from.Hex())
		return true, fmt.Errorf("tx from address mismatch")
	}
	if his.From.Value != tx.Value().String() {
		log.Infof("[HistoryManager] tx value(%s / %s) mismatch", his.From.Value, tx.Value().String())
		return true, fmt.Errorf("tx value mismatch")
	}
	err = h.txMgr.AddHisPage(his)
	if err != nil {
		log.Errorf("add history error: %v", err)
	}
	return true, err
}

func (h *HistoryManager) GetAccountHis(user, account common.Address, page uint64) ([]types.AccountHistory, error) {
	return h.storage.loadUserHistory(account, page)
}

func (h *HistoryManager) UpdateAccountHis(tx string, op *pool.UserOperation) {
	page, err := h.txMgr.GetHisAccIndex(op)
	if err != nil || page < 0 {
		log.Errorf("get history index failed, data: %+v, err: %+v", op, err)
		return
	}
	hisPage, err := h.storage.loadUserHistory(op.Sender, uint64(page))
	if err != nil {
		log.Errorf("get history page failed, data: %+v, err: %+v", op, err)
		return
	}
	for i, his := range hisPage {
		if his.ID != op.Did {
			continue
		}
		hisPage[i].TargetChain = op.InnerExec.ChainId.Uint64()
		hisPage[i].TargetHash = tx
		err = h.storage.cacheUserHistory(common.HexToAddress(his.Account), uint64(page), hisPage)
		if err != nil {
			log.Errorf("cache account(%s, %s) history error: %v", his.Owner, his.Account, err)
		}
		log.Infof("updateUserHistory success: %+v", hisPage)
		break
	}
}

func (h *HistoryManager) SaveAccountHis(owner, account common.Address, data *types.AccountHistory) error {
	user, err := h.storage.Account.GetAccount(owner, account)
	if err != nil || user == nil {
		log.Errorf("get account(%s, %s) no data, error: %v", owner, account, err)
		return fmt.Errorf("get account(%s, %s) no data", owner, account)
	}
	h.txMgr.lock.Lock()
	isExist := h.txMgr.IsExistHis(data)
	if isExist {
		h.txMgr.lock.Unlock()
		log.Errorf("history already exist: %+v", data)
		return errors.New("history already exist")
	}
	err = h.txMgr.AddExistHis(data)
	if err != nil {
		h.txMgr.lock.Unlock()
		log.Errorf("AddExistHis err: %v", err)
		return err
	}
	h.txMgr.lock.Unlock()
	h.chHis <- data
	return nil
}

type HisInfra struct {
	ctx          context.Context
	cfg          Config
	state        *State
	lock         sync.Mutex
	existHis     map[string]int64
	HisIndexKey  string
	HisAccIdxKey string
}

type IndexAccountHistory struct {
	*types.AccountHistory
	Page   string `json:"page"`
	TimeAt int64  `json:"timeAt"`
}

func NewHisInfra(ctx context.Context, cfg Config, s *State) *HisInfra {
	return &HisInfra{
		ctx:          ctx,
		cfg:          cfg,
		state:        s,
		existHis:     make(map[string]int64),
		HisIndexKey:  "his-index",
		HisAccIdxKey: "his-idx-%s",
	}
}

func (tm *HisInfra) GetPendingHis() (map[string]*IndexAccountHistory, error) {
	tm.lock.Lock()
	defer tm.lock.Unlock()
	return tm.pendingHis()
}

func (tm *HisInfra) pendingHis() (map[string]*IndexAccountHistory, error) {
	has, err := tm.state.storage.db.Has([]byte(tm.HisIndexKey))
	if err != nil {
		log.Errorf("key: %s, read db err: %v", tm.HisIndexKey, err)
		return nil, err
	}
	if has {
		data, err := tm.state.storage.db.Get([]byte(tm.HisIndexKey))
		if err != nil {
			log.Errorf("key: %s, get data from db err: %v", tm.HisIndexKey, err)
			return nil, err
		}
		decodeData, err := msgpack.UnmarshalStruct[map[string]*IndexAccountHistory](data)
		if err != nil {
			log.Errorf("key: %s, UnmarshalStruct err: %v", tm.HisIndexKey, err)
			return nil, err
		}
		return decodeData, nil
	}
	return make(map[string]*IndexAccountHistory), nil
}

func (tm *HisInfra) putHis(iahm map[string]*IndexAccountHistory) error {
	data, err := msgpack.MarshalStruct(iahm)
	if err != nil {
		log.Errorf("key: %s, MarshalStruct err: %v", tm.HisIndexKey, err)
		return err
	}
	err = tm.state.storage.db.Put([]byte(tm.HisIndexKey), data)
	return err
}

func (tm *HisInfra) DelExistHis(his *types.AccountHistory) error {
	tm.lock.Lock()
	defer tm.lock.Unlock()
	delete(tm.existHis, fmt.Sprintf("%s", his.ID))
	iahm, err := tm.pendingHis()
	if err != nil {
		return err
	}
	delete(iahm, fmt.Sprintf("%s", his.ID))
	return tm.putHis(iahm)
}

func (tm *HisInfra) AddHisPage(his *types.AccountHistory) error {
	tm.lock.Lock()
	defer tm.lock.Unlock()
	user, err := tm.state.storage.Account.GetAccount(common.HexToAddress(his.Owner), common.HexToAddress(his.Account))
	if err != nil || user == nil {
		log.Errorf("get account(%s, %s) no data, error: %v", his.Owner, his.Account, err)
		return fmt.Errorf("get account(%s, %s) no data", his.Owner, his.Account)
	}
	hisPage, err := tm.state.storage.loadUserHistory(common.HexToAddress(his.Account), user.HistoryPage)
	if (err != nil && err != NotHis) || (user.HistoryPage > 0 && hisPage == nil) {
		log.Errorf("get account(%s, %s) history(%d) no data, err: %v", his.Owner, his.Account, user.HistoryPage, err)
		return fmt.Errorf("get account(%s, %s) no data", his.Owner, his.Account)
	}
	if len(hisPage) >= types.HistoryPageSize {
		user.HistoryPage++
		hisPage = []types.AccountHistory{}
	}
	if hisPage == nil {
		hisPage = []types.AccountHistory{}
	}
	hisPage = append([]types.AccountHistory{*his}, hisPage...)
	err = tm.state.storage.cacheUserHistory(common.HexToAddress(his.Account), user.HistoryPage, hisPage)
	if err != nil {
		log.Errorf("cache account(%s, %s) history error: %v", his.Owner, his.Account, err)
	}
	log.Infof("cacheUserHistory success: %+v", hisPage)
	err = tm.AddHisAccIndex(his, user.HistoryPage)
	if err != nil {
		log.Errorf("add accidx history: %+v, error: %v", his, err)
	}
	return err
}

func (tm *HisInfra) AddHisAccIndex(his *types.AccountHistory, page uint64) error {
	//tm.lock.Lock()
	//defer tm.lock.Unlock()
	haikm, err := tm.hisAccIndex(his.Account)
	if err != nil {
		return err
	}
	haikm[fmt.Sprintf("%s", his.ID)] = strconv.FormatUint(page, 10)
	err = tm.putHisAccIndex(his.Account, haikm)
	log.Infof("AddHisAccIndex success: %+v", haikm)
	return err
}

func (tm *HisInfra) AddExistHis(his *types.AccountHistory) error {
	log.Infof("AddExistHis: %+v", his)
	tm.existHis[fmt.Sprintf("%s", his.ID)] = time.Now().Unix()
	idxAccountHis, err := tm.pendingHis()
	if err != nil {
		return err
	}
	idxAccountHis[fmt.Sprintf("%s", his.ID)] = &IndexAccountHistory{
		AccountHistory: his,
		TimeAt:         time.Now().Unix(),
	}
	err = tm.putHis(idxAccountHis)
	log.Infof("AddExistHis put indexkey: %+v", idxAccountHis)
	return err
}

func (tm *HisInfra) IsExistHis(his *types.AccountHistory) bool {
	if _, ok := tm.existHis[fmt.Sprintf("%s", his.ID)]; ok {
		return true
	}
	iahm, err := tm.pendingHis()
	if err != nil {
		return false
	}
	if idxHis, ok := iahm[fmt.Sprintf("%s", his.ID)]; ok {
		if idxHis.Account == his.Account {
			return true
		}
	}
	haikm, err := tm.hisAccIndex(his.Account)
	if err != nil {
		return false
	}
	if idxHis, ok := haikm[fmt.Sprintf("%s", his.ID)]; ok && idxHis != "" {
		return true
	}
	return false
}

func (tm *HisInfra) GetHisAccIndex(op *pool.UserOperation) (int64, error) {
	tm.lock.Lock()
	defer tm.lock.Unlock()
	haikm, err := tm.hisAccIndex(op.Sender.Hex())
	if err != nil {
		return -1, err
	}
	if page, ok := haikm[op.Did]; ok {
		return strconv.ParseInt(page, 10, 64)
	}
	return -1, nil
}

func (tm *HisInfra) hisAccIndex(user string) (map[string]string, error) {
	key := fmt.Sprintf(tm.HisAccIdxKey, user)
	has, err := tm.state.storage.db.Has([]byte(key))
	if err != nil {
		log.Errorf("key: %s, read db err: %+v", key, err)
		return nil, err
	}
	var accIdx map[string]string
	if has {
		data, err := tm.state.storage.db.Get([]byte(key))
		if err != nil {
			log.Errorf("get data(%s) from db err: %+v", key, err)
			return nil, err
		}
		decodeData, err := msgpack.UnmarshalStruct[map[string]string](data)
		if err != nil {
			log.Errorf("key: %s, UnmarshalStruct err: %+v", key, err)
			return nil, err
		}
		accIdx = decodeData
	} else {
		accIdx = make(map[string]string)
	}
	return accIdx, nil
}

func (tm *HisInfra) putHisAccIndex(user string, d map[string]string) error {
	key := fmt.Sprintf(tm.HisAccIdxKey, user)
	data, err := msgpack.MarshalStruct(d)
	if err != nil {
		log.Errorf("key: %s, MarshalStruct err: %+v", key, err)
		return err
	}
	err = tm.state.storage.db.Put([]byte(key), data)
	return err
}
