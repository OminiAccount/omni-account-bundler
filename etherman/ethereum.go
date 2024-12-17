package etherman

import (
	"context"
	"fmt"
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/etherman/contracts/SyncRouter"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/utils/chains"
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/msgpack"
	"github.com/OAB/utils/packutils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"math/big"
	"time"
)

const SaltKey = "%d-salt"
const PendingCreateAAKey = "pending-create-aa-%d"

type EtherMan struct {
	ctx          context.Context
	cancel       func()
	cfg          Config
	chainsClient map[chains.ChainId]*EthereumClient
	db           ethdb.Database
}

// NewEthereum create a EthereumClient that support multiple chains
func NewEthereum(ctx context.Context, cfg Config, ethDb ethdb.Database) (*EtherMan, error) {
	etherCtx, cancel := context.WithCancel(ctx)
	em := &EtherMan{
		ctx:          etherCtx,
		cancel:       cancel,
		cfg:          cfg,
		chainsClient: make(map[chains.ChainId]*EthereumClient, chains.MaxChainInfoLength),
		db:           ethDb,
	}
	for _, n := range cfg.Networks {
		chainId := chains.ChainId(n.ChainId)
		cli, err := NewClient(n)
		if err != nil {
			return nil, err
		}
		opts, _, err := em.LoadAuthFromKeyStore(n.PrivateKeys.Path, n.PrivateKeys.Password, n.ChainId)
		if err != nil {
			return nil, err
		}
		cli.sender = opts.From
		cli.opAuth = *opts
		em.chainsClient[chainId] = cli
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
	etherCli := ether.chainsClient[ether.cfg.VizingChainID]
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
	etherCli := ether.chainsClient[ether.cfg.VizingChainID]
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

func (ether *EtherMan) CreateAccount(user common.Address) (*common.Address, uint64) {
	ccli := ether.chainsClient[ether.cfg.VizingChainID]
	ccli.lock.Lock()
	defer ccli.lock.Unlock()
	var salt uint64
	chainSaltKey := fmt.Sprintf(SaltKey, uint64(ccli.chainID))
	has, err := ether.db.Has([]byte(chainSaltKey))
	if err != nil {
		log.Errorf("[CreateAccount] read db err: %+v", err)
		return nil, salt
	}
	if has {
		data, err := ether.db.Get([]byte(chainSaltKey))
		if err != nil {
			log.Errorf("[CreateAccount] from db get data err: %+v", err)
			return nil, salt
		}
		salt = packutils.BytesToUint64(data)
	}
	log.Infof("[CreateAccount] chainID: %d, salt: %d", ccli.chainID, salt)
	opts := new(bind.CallOpts)
	opts.From = ccli.sender
	adr, err := ccli.aaFactory.GetAccountAddress(opts, user, big.NewInt(0).SetUint64(salt))
	if err != nil {
		log.Errorf("get create account err: %+v", err)
		return nil, salt
	}
	err = ether.db.Put([]byte(chainSaltKey), packutils.Uint64ToBytes(salt))
	if err != nil {
		log.Errorf("put salt key to db err: %+v", err)
		return nil, salt
	}
	for _, cli := range ether.chainsClient {
		go ether.pendingCreateAA(PendingData{cli.chainID, user, salt})
	}
	return &adr, salt
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
	pendingKey := fmt.Sprintf(PendingCreateAAKey, pd.ChainID)
	pdm := ether.loadPendingData(pd.ChainID)
	if pdm == nil {
		return
	}
	if isDel {
		delete(pdm, pd.User.Hex())
	} else {
		pdm[pd.User.Hex()] = pd
	}
	by, err := msgpack.MarshalStruct(pdm)
	if err != nil {
		log.Errorf("chainID: %d, MarshalStruct err: %+v", pd.ChainID, err)
		return
	}
	err = ether.db.Put([]byte(pendingKey), by)
	if err != nil {
		log.Errorf("chainID: %d, put pending key to db err: %+v", pd.ChainID, err)
		return
	}
}

func (ether *EtherMan) loadPendingData(cid chains.ChainId) map[string]PendingData {
	pendingKey := fmt.Sprintf(PendingCreateAAKey, cid)
	has, err := ether.db.Has([]byte(pendingKey))
	if err != nil {
		log.Errorf("[saveFailedCreateAA] chainID: %d, read db err: %+v", cid, err)
		return nil
	}
	var pdm map[string]PendingData
	if has {
		data, err := ether.db.Get([]byte(pendingKey))
		if err != nil {
			log.Errorf("[saveFailedCreateAA] chainID: %d, from db get data err: %+v", cid, err)
			return nil
		}
		decodeData, err := msgpack.UnmarshalStruct[map[string]PendingData](data)
		if err != nil {
			log.Errorf("[saveFailedCreateAA] chainID: %d, UnmarshalStruct err: %+v", cid, err)
			return nil
		}
		pdm = decodeData
	} else {
		pdm = make(map[string]PendingData)
	}
	return pdm
}

func (ether *EtherMan) Start() {
	go func() {
		ticker := time.NewTicker(time.Minute)
		for {
			select {
			case <-ticker.C:
				for cid, _ := range ether.chainsClient {
					pdm := ether.loadPendingData(cid)
					if pdm == nil {
						return
					}
					for _, pd := range pdm {
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
