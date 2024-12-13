package etherman

import (
	"fmt"
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/etherman/contracts/SyncRouter"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/utils/chains"
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/packutils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"math/big"
)

const SaltKey = "%d-salt"
const VizingChain = 28516

type EtherMan struct {
	chainsClient map[chains.ChainId]*EthereumClient
	db           ethdb.Database
}

// NewEthereum create a EthereumClient that support multiple chains
func NewEthereum(cfg Config, ethDb ethdb.Database) (*EtherMan, error) {
	em := &EtherMan{
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
	etherCli := ether.chainsClient[VizingChain]
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
	etherCli := ether.chainsClient[VizingChain]
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
	ccli := ether.chainsClient[VizingChain]
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
		go func(c *EthereumClient) {
			c.lock.Lock()
			defer c.lock.Unlock()
			c.opAuth.Nonce = nil
			c.opAuth.GasPrice = nil
			c.opAuth.GasLimit = 0
			c.opAuth.Value = big.NewInt(0)
			tx, err := c.aaFactory.CreateAccount(&c.opAuth, user, big.NewInt(0).SetUint64(salt))
			if err != nil {
				// TODO handler error
				log.Errorf("chain: %d, user: %s, create account err: %+v", c.chainID, user, err)
				return
			}
			log.Infof("create account success, chain: %d, user: %s, tx: %s", c.chainID, user, tx.Hash())
		}(cli)
	}
	return &adr, salt
}