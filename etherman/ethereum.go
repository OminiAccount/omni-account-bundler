package etherman

import (
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/utils/chains"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"math/big"
)

type EtherMan struct {
	ChainsClient map[chains.ChainId]*ethereumClient
	db           ethdb.Database
}

// NewEthereum create a EthereumClient that support multiple chains
func NewEthereum(cfg Config, ethDb ethdb.Database) (*EtherMan, error) {
	em := &EtherMan{
		ChainsClient: make(map[chains.ChainId]*ethereumClient, chains.MaxChainInfoLength),
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
		em.ChainsClient[chainId] = cli
	}
	return em, nil
}

func (ether *EtherMan) EstimateGas(chainID uint64, useFee *big.Int, msg []byte) (*big.Int, error) {
	etherCli := ether.ChainsClient[0]
	opts := new(bind.CallOpts)
	opts.From = etherCli.sender
	return etherCli.syncRouter.FetchOmniMessageFee(
		opts,
		chainID,
		ether.ChainsClient[chains.ChainId(chainID)].entryPointAddr,
		useFee,
		msg,
	)
}

func (ether *EtherMan) UpdateEntryPointRoot(proof hexutil.Bytes, batches []EntryPoint.IEntryPointBatchData, extraInfo EntryPoint.IEntryPointChainsExecuteInfo) (common.Hash, error) {
	etherCli := ether.ChainsClient[0]
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
	tx, err := etherCli.entryPoint.VerifyBatch(&opts, proof, batches, extraInfo)
	if err != nil {
		return common.Hash{}, err
	}
	log.Infof("verify batch tx: %s", tx.Hash())
	return tx.Hash(), nil
}
