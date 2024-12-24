package etherman

import (
	"context"
	"errors"
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/etherman/contracts/SyncRouter"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/utils/chains"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4/pgxpool"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"
)

const SaltKey = "%d-salt"
const PendingCreateAAKey = "pending-create-aa-%d"

var (
	ErrExistAccount = errors.New("account already exists")
)

type lockUser struct {
	adr    common.Address
	salt   uint64
	timeAt time.Time
}

var tempUser sync.Map

type EtherMan struct {
	ctx          context.Context
	cancel       func()
	cfg          Config
	chainsClient map[chains.ChainId]*EthereumClient
	db           *PostgresStorage
}

// NewEthereum create a EthereumClient that support multiple chains
func NewEthereum(ctx context.Context, cfg Config, pg *pgxpool.Pool) (*EtherMan, error) {
	etherCtx, cancel := context.WithCancel(ctx)
	em := &EtherMan{
		ctx:          etherCtx,
		cancel:       cancel,
		cfg:          cfg,
		chainsClient: make(map[chains.ChainId]*EthereumClient, chains.MaxChainInfoLength),
		db:           NewPostgresStorage(pg),
	}
	var nw = []Network{
		cfg.VizingNetwork,
		cfg.ArbitNetwork,
		cfg.BaseNetwork,
		cfg.OpNetwork,
	}
	for _, n := range nw {
		if n.ChainId < 1 {
			continue
		}
		cli, err := NewClient(n)
		if err != nil {
			return nil, err
		}
		opts, _, err := em.LoadAuthFromKeyStore(n.PrivateKeys.Path, n.PrivateKeys.Password, uint64(n.ChainId))
		if err != nil {
			return nil, err
		}
		cli.sender = opts.From
		cli.opAuth = *opts
		em.chainsClient[n.ChainId] = cli
	}
	return em, nil
}

func (ether *EtherMan) GetChains() map[chains.ChainId]*EthereumClient {
	return ether.chainsClient
}

func (ether *EtherMan) GetChainCli(c chains.ChainId) *EthereumClient {
	return ether.chainsClient[c]
}

func (ether *EtherMan) EstimateGas(chainID uint64, useFee *big.Int, userOperations []SyncRouter.BaseStructPackedUserOperation) (*big.Int, error) {
	etherCli := ether.chainsClient[ether.cfg.VizingNetwork.ChainId]
	opts := new(bind.CallOpts)
	opts.From = etherCli.sender
	log.Infof("EstimateGas destChainId: %d, destContract: %s, userOperations: %+v",
		chainID, ether.chainsClient[chains.ChainId(chainID)].entryPointAddr, userOperations)
	return etherCli.syncRouter.FetchOmniMessageFee(
		opts,
		chainID,
		ether.chainsClient[chains.ChainId(chainID)].entryPointAddr,
		useFee,
		userOperations,
	)
}

func (ether *EtherMan) UpdateEntryPointRoot(proof hexutil.Bytes, batches []EntryPoint.BaseStructBatchData, extraInfo EntryPoint.BaseStructChainsExecuteInfo) (common.Hash, error) {
	etherCli := ether.chainsClient[ether.cfg.VizingNetwork.ChainId]
	opts := etherCli.opAuth
	/*nonce, err := etherCli.ethClient.NonceAt(context.Background(), opts.From, nil)
	if err != nil {
		log.Errorf("get nonce, error: %+v", err)
		return common.Hash{}, err
	}
	opts.Nonce = big.NewInt(0).SetUint64(nonce)
	gasPrice, err := etherCli.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("get SuggestGasPrice, err: %+v", err)
		return common.Hash{}, err
	}*/
	opts.Nonce = nil
	opts.GasPrice = nil
	opts.GasLimit = 0
	opts.Value = big.NewInt(0)
	extraInfo.Beneficiary = etherCli.sender
	log.Infof("VerifyBatches proof: %s, batches: %+v, extraInfo: %+v", proof.String(), batches, extraInfo)
	tx, err := etherCli.entryPoint.VerifyBatches(&opts, proof, batches, extraInfo)
	if err != nil {
		return common.Hash{}, err
	}
	log.Infof("verify batch tx: %s", tx.Hash())
	return tx.Hash(), nil
}

func (ether *EtherMan) CreateAccount(user common.Address) (*common.Address, uint64, error) {
	ccli := ether.chainsClient[ether.cfg.VizingNetwork.ChainId]
	ccli.lock.Lock()
	defer ccli.lock.Unlock()
	if u, ok := tempUser.Load(user.Hex()); ok {
		lu := u.(lockUser)
		return &lu.adr, lu.salt, ErrExistAccount
	}
	dbTx, err := ether.db.BeginDBTransaction(ether.ctx)
	if err != nil {
		log.Errorf("use db transaction err: %+v", err)
		return nil, 0, err
	}
	salt := ether.db.GetChainSalt(ether.ctx, uint64(ccli.chainID), dbTx)
	log.Infof("[CreateAccount] chainID: %d, salt: %d", ccli.chainID, salt)
	opts := new(bind.CallOpts)
	opts.From = ccli.sender
	adr, err := ccli.aaFactory.GetAccountAddress(opts, user, big.NewInt(0).SetUint64(salt))
	if err != nil {
		_ = dbTx.Rollback(ether.ctx)
		log.Errorf("get create account err: %+v", err)
		return nil, salt, err
	}
	ccli.opAuth.Nonce = nil
	ccli.opAuth.GasPrice = nil
	ccli.opAuth.GasLimit = 0
	ccli.opAuth.Value = big.NewInt(0)
	_, err = ccli.aaFactory.CreateAccount(&ccli.opAuth, user, big.NewInt(0).SetUint64(salt))
	if err != nil {
		_ = dbTx.Rollback(ether.ctx)
		log.Errorf("chain: %d, user: %s, create account err: %+v",
			ether.cfg.VizingNetwork.ChainId, user, err)
		return nil, salt, err
	}
	err = ether.db.SetChain(ether.ctx, uint64(ccli.chainID), salt+1, "salt", dbTx)
	if err != nil {
		_ = dbTx.Rollback(ether.ctx)
		log.Errorf("set chain salt err: %+v", err)
		return nil, salt, err
	}
	err = dbTx.Commit(ether.ctx)
	if err != nil {
		log.Errorf("db commit err: %+v", err)
	}
	log.Infof("==============================")
	salt2 := ether.db.GetChainSalt(ether.ctx, uint64(ccli.chainID), nil)
	log.Infof("check add salt: %d", ccli.chainID, salt2)
	log.Infof("==============================")
	for _, cli := range ether.chainsClient {
		if cli.chainID == ether.cfg.VizingNetwork.ChainId {
			continue
		}
		go ether.pendingCreateAA(PendingData{cli.chainID, user, salt})
	}
	tempUser.Store(user.Hex(), lockUser{adr, salt, time.Now()})
	return &adr, salt, nil
}

type PendingData struct {
	ChainID chains.ChainId
	User    common.Address
	Salt    uint64
}

func (ether *EtherMan) pendingCreateAA(pd PendingData) {
	c := ether.chainsClient[pd.ChainID]
	c.lock.Lock()
	defer c.lock.Unlock()
	c.opAuth.Nonce = nil
	c.opAuth.GasPrice = nil
	c.opAuth.GasLimit = 0
	c.opAuth.Value = big.NewInt(0)
	tx, err := c.aaFactory.CreateAccount(&c.opAuth, pd.User, big.NewInt(0).SetUint64(pd.Salt))
	if err != nil {
		log.Errorf("chain: %d, user: %s, create account err: %+v", c.chainID, pd.User, err)
		ether.saveFailedCreateAA(pd, false)
		return
	}
	ether.saveFailedCreateAA(pd, true)
	log.Infof("create account success, chain: %d, user: %s, tx: %s", c.chainID, pd.User, tx.Hash())
}

func (ether *EtherMan) saveFailedCreateAA(pd PendingData, isDel bool) {
	err := ether.db.UpdateUserFailedSalt(ether.ctx, pd.User.String(), uint64(pd.ChainID), isDel, nil)
	if err != nil {
		log.Errorf("chain: %d, user: %s, saveFailedCreateAA err: %+v", pd.ChainID, pd.User, err)
	}
}

func (ether *EtherMan) Start() {
	go func() {
		ticker := time.NewTicker(time.Minute)
		for {
			select {
			case <-ticker.C:
				tempUser.Range(func(key, value any) bool {
					lu := value.(lockUser)
					if time.Since(lu.timeAt) > time.Minute*10 {
						tempUser.Delete(key)
					}
					return true
				})
				list := ether.db.GetFailedSalts(ether.ctx, nil)
				if list == nil || len(list) < 1 {
					continue
				}
				for _, ufs := range list {
					cids := strings.Split(ufs.FailedSalt, ",")
					if len(cids) < 1 {
						continue
					}
					for _, cid := range cids {
						nid, err := strconv.ParseUint(cid, 10, 64)
						if err != nil {
							continue
						}
						pd := PendingData{
							User:    common.HexToAddress(ufs.User),
							ChainID: chains.ChainId(nid),
							Salt:    ufs.Salt,
						}
						ok, err := ether.db.IsExistChainByOwner(ether.ctx, ufs.User, uint64(pd.ChainID), nil)
						if err != nil {
							log.Errorf("IsExistChainByOwner err: %+v", err)
							continue
						}
						if ok {
							ether.saveFailedCreateAA(pd, true)
							continue
						}
						ether.pendingCreateAA(pd)
					}
				}
			case <-ether.ctx.Done():
				log.Info("EtherMan exit...")
				return
			}
		}
	}()
}

func (ether *EtherMan) Stop() {
	log.Info("EtherMan stop")
	ether.cancel()
}
