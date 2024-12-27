package etherman

import (
	"errors"
	"fmt"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync"
	"time"
)

const ZeroHash = "0x0000000000000000000000000000000000000000"

var (
	ErrExistAccount    = errors.New("account already exists")
	ErrAlreadyCreateAA = errors.New("aa already created")
)

type lockUser struct {
	adr    common.Address
	salt   uint64
	timeAt time.Time
}

var tempUser sync.Map

func (ether *EtherMan) CreateVizingAccount(user common.Address) (*common.Address, uint64, error) {
	ccli := ether.chainsClient[ether.cfg.VizingNetwork.ChainId]
	ccli.lock.Lock()
	defer ccli.lock.Unlock()
	var salt uint64
	lockKey := fmt.Sprintf("%s-%d", user.Hex(), ether.cfg.VizingNetwork.ChainId)
	if u, ok := tempUser.Load(lockKey); ok {
		lu := u.(lockUser)
		return &lu.adr, lu.salt, ErrExistAccount
	}
	ua, err := ccli.aaFactory.Accounts(new(bind.CallOpts), user)
	if err != nil {
		log.Errorf("chain: %d, user: %s, get account err: %+v",
			ether.cfg.VizingNetwork.ChainId, user, err)
		return nil, salt, err
	}
	if ua.Hex() != ZeroHash {
		return &ua, salt, ErrAlreadyCreateAA
	}
	//salt := ether.db.GetChainSalt(ether.ctx, ccli.chainID, nil)
	//log.Infof("[CreateAccount] chainID: %d, salt: %d", ccli.chainID, salt)
	opts := new(bind.CallOpts)
	opts.From = ccli.sender
	adr, err := ccli.aaFactory.GetAccountAddress(opts, user, big.NewInt(0).SetUint64(salt))
	if err != nil {
		log.Errorf("get create account err: %+v", err)
		return nil, salt, err
	}
	ccli.opAuth.Nonce = nil
	ccli.opAuth.GasPrice = nil
	ccli.opAuth.GasLimit = 0
	ccli.opAuth.Value = big.NewInt(0)
	tx, err := ccli.aaFactory.CreateAccount(&ccli.opAuth, user, big.NewInt(0).SetUint64(salt))
	if err != nil {
		log.Errorf("chain: %d, user: %s, create account err: %+v",
			ether.cfg.VizingNetwork.ChainId, user, err)
		return nil, salt, err
	}
	tempUser.Store(lockKey, lockUser{adr, salt, time.Now()})
	log.Infof("create account success, chain: vizing, user: %s, tx: %s", user, tx.Hash())
	return &adr, salt, nil
}

func (ether *EtherMan) CreateOtherAccount(user common.Address, salt, nid uint64) (*common.Address, error) {
	return ether.pendingCreateAA(PendingData{nid, user, salt})
}

type PendingData struct {
	ChainID uint64
	User    common.Address
	Salt    uint64
}

func (ether *EtherMan) pendingCreateAA(pd PendingData) (*common.Address, error) {
	c := ether.chainsClient[pd.ChainID]
	c.lock.Lock()
	defer c.lock.Unlock()
	ua, err := c.aaFactory.Accounts(new(bind.CallOpts), pd.User)
	if err != nil {
		log.Errorf("chain: %d, user: %s, get account err: %+v", c.chainID, pd.User, err)
		ether.saveFailedCreateAA(pd, false)
		return nil, err
	}
	if ua.Hex() != ZeroHash {
		return &ua, ErrAlreadyCreateAA
	}
	c.opAuth.Nonce = nil
	c.opAuth.GasPrice = nil
	c.opAuth.GasLimit = 0
	c.opAuth.Value = big.NewInt(0)
	tx, err := c.aaFactory.CreateAccount(&c.opAuth, pd.User, big.NewInt(0).SetUint64(pd.Salt))
	if err != nil {
		log.Errorf("chain: %d, user: %s, create account err: %+v", c.chainID, pd.User, err)
		ether.saveFailedCreateAA(pd, false)
		return nil, err
	}
	ether.saveFailedCreateAA(pd, true)
	log.Infof("create account success, chain: %d, user: %s, tx: %s", c.chainID, pd.User, tx.Hash())
	return nil, nil
}

func (ether *EtherMan) saveFailedCreateAA(pd PendingData, isDel bool) {
	err := ether.db.UpdateUserFailedSalt(ether.ctx, pd.User.String(), uint64(pd.ChainID), isDel, nil)
	if err != nil {
		log.Errorf("chain: %d, user: %s, saveFailedCreateAA err: %+v", pd.ChainID, pd.User, err)
	}
}
