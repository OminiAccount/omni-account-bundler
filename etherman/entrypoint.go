package etherman

import (
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/etherman/contracts/SyncRouter"
	"github.com/OAB/etherman/contracts/ZKVizingAADataHelp"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

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
	for i, b := range batches {
		log.Infof("batchNum: %d, Accinput: %s", i, common.Bytes2Hex(b.AccInputHash[:]))
	}
	log.Infof("NewStateRoot: %s", common.Bytes2Hex(extraInfo.NewStateRoot[:]))
	tx, err := etherCli.entryPoint.VerifyBatches(&opts, proof, batches, extraInfo)
	if err != nil {
		return common.Hash{}, err
	}
	log.Infof("verify batch tx: %s", tx.Hash())
	return tx.Hash(), nil
}

func (ether *EtherMan) EstimateGas(chainID uint64, useFee *big.Int, userOperations []SyncRouter.BaseStructPackedUserOperation) (*big.Int, error) {
	etherCli := ether.chainsClient[ether.cfg.VizingNetwork.ChainId]
	opts := new(bind.CallOpts)
	opts.From = etherCli.sender
	log.Infof("EstimateGas destChainId: %d, destContract: %s, userOperations: %+v",
		chainID, ether.chainsClient[chainID].entryPointAddr, userOperations)
	return etherCli.syncRouter.FetchOmniMessageFee(
		opts,
		chainID,
		ether.chainsClient[chainID].entryPointAddr,
		useFee,
		userOperations,
	)
}

func (ether *EtherMan) DecodeSwapCalldata(chainID uint64, calldata hexutil.Bytes) (ZKVizingAADataHelp.BaseStructV3SwapParams, error) {
	etherCli := ether.chainsClient[ether.cfg.VizingNetwork.ChainId]
	opts := new(bind.CallOpts)
	opts.From = etherCli.sender
	log.Infof("DecodeCalldata destChainId: %d", chainID)
	return etherCli.dataHelp.DecodeUniswapV3Data(
		opts,
		calldata,
	)
}
