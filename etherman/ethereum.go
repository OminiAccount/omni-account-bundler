package etherman

import (
	"context"
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
	"sync"
)

const SaltKey = "%d-salt"

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
	etherCli := ether.chainsClient[0]
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
	etherCli := ether.chainsClient[0]
	opts := etherCli.opAuth
	nonce, err := etherCli.ethClient.NonceAt(context.Background(), opts.From, nil)
	if err != nil {
		log.Errorf("get nonce, error: %+v", err)
		return common.Hash{}, err
	}
	opts.Nonce = big.NewInt(0).SetUint64(nonce)
	//gasPrice, err := etherCli.ethClient.SuggestGasPrice(context.Background())
	//if err != nil {
	//	ether.logger.Error("get SuggestGasPrice", "error", err.Error())
	//	return nil, err
	//}
	//opts.GasPrice = gasPrice
	//opts.GasLimit = uint64(47860)
	//opts.Value = big.NewInt(0)

	extraInfo.Beneficiary = etherCli.sender
	log.Infof("VerifyBatches proof: %s, batches: %+v, extraInfo: %+v", proof.String(), batches, extraInfo)
	tx, err := etherCli.entryPoint.VerifyBatches(&opts, proof, batches, extraInfo)
	if err != nil {
		return common.Hash{}, err
	}
	log.Infof("verify batch tx: %s", tx.Hash())
	return tx.Hash(), nil
}

func (ether *EtherMan) CreateAccount(user common.Address) *common.Address {
	etherCli := ether.chainsClient[0]
	etherCli.lock.Lock()
	var salt uint64
	chainSaltKey := fmt.Sprintf(SaltKey, etherCli.chainID)
	has, err := ether.db.Has([]byte(chainSaltKey))
	if err != nil {
		log.Errorf("[CreateAccount] read db err: %+v", err)
		etherCli.lock.Unlock()
		return nil
	}
	if has {
		data, err := ether.db.Get([]byte(chainSaltKey))
		if err != nil {
			log.Errorf("[CreateAccount] from db get data err: %+v", err)
			etherCli.lock.Unlock()
			return nil
		}
		salt = packutils.BytesToUint64(data)
	}
	opts := new(bind.CallOpts)
	opts.From = etherCli.sender
	adr, err := etherCli.aaFactory.GetAccountAddress(opts, user, big.NewInt(0).SetUint64(salt))
	if err != nil {
		log.Errorf("get create account err: %+v", err)
		etherCli.lock.Unlock()
		return nil
	}
	err = ether.db.Put([]byte(chainSaltKey), packutils.Uint64ToBytes(salt))
	if err != nil {
		log.Errorf("put key to db err: %+v", err)
		etherCli.lock.Unlock()
		return nil
	}
	etherCli.lock.Unlock()
	for _, cli := range ether.chainsClient {
		go func(c *EthereumClient) {
			c.lock.Lock()
			defer c.lock.Unlock()
			nonce, err := c.ethClient.NonceAt(context.Background(), opts.From, nil)
			if err != nil {
				log.Errorf("chain: %d, user: %s, get nonce, error: %+v", c.chainID, user, err)
				return
			}
			c.opAuth.Nonce = big.NewInt(0).SetUint64(nonce)
			tx, err := c.aaFactory.CreateAccount(&c.opAuth, user, salt)
			if err != nil {
				log.Errorf("chain: %d, user: %s, create account err: %+v", c.chainID, user, err)
				return
			}
			log.Infof("create account success, chain: %d, user: %s, tx: %s", c.chainID, user, tx.Hash())
		}(cli)
	}
	// set account
	return &adr
}