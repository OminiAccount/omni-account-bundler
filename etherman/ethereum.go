package etherman

import (
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/etherman/contracts/SyncRouter"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/utils/chains"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"math/big"
)

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
		cli.auth = *opts
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
	opts := etherCli.auth
	//nonce, err := etherCli.ethClient.NonceAt(context.Background(), etherCli.sender, nil)
	//if err != nil {
	//	ether.logger.Error("get nonce", "error", err.Error())
	//	return nil, err
	//}
	//opts.Nonce = big.NewInt(0).SetUint64(nonce)
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
