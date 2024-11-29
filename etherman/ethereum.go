package etherman

import (
	"fmt"
	"github.com/OAB/utils/chains"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethereumType "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
)

type EtherMan struct {
	ChainsClient map[chains.ChainId]*ethereumClient

	db     ethdb.Database
	logger log.Logger
}

// NewEthereum create a EthereumClient that support multiple chains
func NewEthereum(cfg Config, ethDb ethdb.Database) (*EtherMan, error) {
	em := &EtherMan{
		ChainsClient: make(map[chains.ChainId]*ethereumClient, chains.MaxChainInfoLength),
		logger:       log.New("service", "ethereum"),
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

func (ether *EtherMan) UpdateEntryPointRoot(proof hexutil.Bytes, pubicValues hexutil.Bytes) (*ethereumType.Transaction, error) {
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

	tx, err := etherCli.entryPoint.VerifyBatch(&opts, proof, pubicValues, opts.From)
	if err != nil {
		return nil, err
	}
	fmt.Println("tx", tx.Hash())

	ether.logger.Info("update entryPoint root", "hash", tx.Hash().String())
	return tx, nil
}
