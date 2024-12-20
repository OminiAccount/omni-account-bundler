package state

import (
	"fmt"
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/etherman/contracts/SyncRouter"
	"github.com/OAB/pool"
	"github.com/OAB/state/types"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (s *State) processBatch(batchContext *pool.BatchContext) error {
	batch, err := NewBatch(s.storage.BatchNumber + 1)
	if err != nil {
		return err
	}

	// Set old smt root
	batch.SetOldStateRoot(common.BigToHash(s.tree.LastRoot()))

	for _, uo := range batchContext.SignedUserOperations() {
		log.Infof("operationValue: %+v ===============", uo)

		if uo.Phase != pool.PhaseFirst {
			continue
		}
		// balance
		var balance big.Int
		balanceU256, err := s.tree.GetAccountBalance(uo.Sender)
		if err != nil {
			return err
		}
		balance.SetBytes(balanceU256.Bytes())
		log.Infof("oldBalance: %d", &balance)

		if uo.OperationType == 1 {
			log.Infof("balance %d + operationValue %d", &balance, uo.OperationValue.Uint64())
			balance.Add(&balance, uo.OperationValue.ToInt())
		} else if uo.OperationType == 2 {
			if balance.Cmp(uo.OperationValue.ToInt()) < 0 {
				log.Warnf("current account balance %d is insufficient to withdraw %d", &balance, uo.OperationValue.Uint64())
			}
			log.Infof("balance %d - operationValue %d", &balance, uo.OperationValue.Uint64())
			balance.Sub(&balance, uo.OperationValue.ToInt())
		}

		log.Infof("balance %d - CalculateGasUsed %s", &balance, uo.CalculateGasUsed())
		balance.Sub(&balance, uo.CalculateGasUsed())

		if _, err = s.tree.SetAccountBalance(uo.Sender.String(), &balance); err != nil {
			return err
		}
		log.Infof("newBalance: %d", &balance)

		// nonce
		var nonce big.Int
		nonceU256, err := s.tree.GetAccountNonce(uo.Sender, uo.Exec.ChainId.String())
		if err != nil {
			return err
		}
		nonce.SetBytes(nonceU256.Bytes())
		log.Infof("phase first oldNonce: %d", &nonce)
		nonce.Add(&nonce, big.NewInt(1))
		log.Infof("phase first newNonce: %d", &nonce)
		_, err = s.tree.SetAccountNonce(uo.Sender.String(), &nonce, uo.Exec.ChainId.String())
		if err != nil {
			return err
		}
		if uo.InnerExec.ChainId.Uint64() > 0 {
			var nonce2 big.Int
			nonceU256, err = s.tree.GetAccountNonce(uo.Sender, uo.InnerExec.ChainId.String())
			if err != nil {
				return err
			}
			nonce2.SetBytes(nonceU256.Bytes())
			log.Infof("phase second oldNonce: %d", &nonce2)
			nonce2.Add(&nonce2, big.NewInt(1))
			log.Infof("phase second newNonce: %d", &nonce2)
			_, err = s.tree.SetAccountNonce(uo.Sender.String(), &nonce2, uo.InnerExec.ChainId.String())
			if err != nil {
				return err
			}
		}
	}

	// Set BatchL2Data and BatchHashData
	batch.SetBatchL2Data(batchContext.SignedUserOperations())
	// Set NewAccInputHash
	batch.SetNewAccInputHash(batch.BatchHashData)
	// Set new state root
	batch.SetNewStateRoot(common.BigToHash(s.tree.LastRoot()))
	err = s.storage.cacheBatch(batch)
	if err != nil {
		return err
	}
	s.proofQueue.Enqueue(*batch)
	log.Infof("Successfully sealing a batch, number: %d, stateRoot: %s", batch.NewNumBatch, batch.NewStateRoot.Hex())

	return nil
}

func (s *State) executeBatch() (*ProofResult, error) {
	if s.provenQueue.IsEmpty() {
		return nil, nil
	}
	log.Info("executing batch...")
	res, ok := s.provenQueue.Dequeue()
	if !ok {
		return nil, nil
	}
	batches := EntryPoint.BaseStructBatchData{}
	extraInfo := EntryPoint.BaseStructChainsExecuteInfo{}
	var suo []*pool.SignedUserOperation
	for i := res.Number; i <= res.FinalNumber; i++ {
		if i <= s.storage.VerifyBatchNumber {
			return nil, fmt.Errorf("batch number too low")
		}
		batch, err := s.storage.loadBatch(i)
		if err != nil {
			return nil, err
		}
		if batch == nil {
			return nil, fmt.Errorf("batch(%d) no cache data", i)
		}
		if i == res.FinalNumber {
			extraInfo.NewStateRoot = batch.NewStateRoot
		}
		suo = append(suo, batch.BatchL2SrcData...)
	}
	if len(suo) < 1 {
		log.Warnf("batch(%d - %d) no data", res.Number, res.FinalNumber)
		s.storage.VerifyBatchNumber = res.FinalNumber
		return nil, nil
	}
	var puo []EntryPoint.BaseStructPackedUserOperation
	chainExtra := make(map[uint64]EntryPoint.BaseStructChainExecuteExtra)
	for _, op := range suo {
		puo = append(puo, op.ToEntryPointOp())
		chainID := op.Exec.ChainId
		if op.Phase == pool.PhaseSecond {
			chainID = op.InnerExec.ChainId
		}
		fee, err := s.ethereum.EstimateGas(chainID.Uint64(), op.CalculateGasUsed(), []SyncRouter.BaseStructPackedUserOperation{op.ToSyncRouterOp()})
		if err != nil {
			return &res, err
		}
		log.Infof("EstimateGas fee: %d", fee.Uint64())
		if cee, ok := chainExtra[chainID.Uint64()]; ok {
			cee.ChainFee += fee.Uint64()
			cee.ChainUserOperationsNumber++
			chainExtra[chainID.Uint64()] = cee
		} else {
			chainExtra[chainID.Uint64()] = EntryPoint.BaseStructChainExecuteExtra{
				ChainFee:                  fee.Uint64(),
				ChainUserOperationsNumber: 1,
			}
		}
	}
	for _, v := range chainExtra {
		extraInfo.ChainExtra = append(extraInfo.ChainExtra, v)
	}
	batches.AccInputHash = batchHashData(suo)
	batches.UserOperations = puo
	tx, err := s.ethereum.UpdateEntryPointRoot(res.Proof, []EntryPoint.BaseStructBatchData{batches}, extraInfo)
	if err != nil {
		return &res, err
	}
	for i := res.Number; i <= res.FinalNumber; i++ {
		err = s.storage.clearBatch(i)
		if err != nil {
			return nil, err
		}
	}
	/*for _, op := range suo {
		if op.OperationType == pool.DepositAction {
			continue
		}
		if op.Phase == pool.PhaseSecond {
			s.his.UpdateAccountHis(tx.Hex(), op.UserOperation)
			continue
		}
		uoHis := types.ToUserOperationHis(tx.Hex(), op.UserOperation)
		err = s.his.SaveAccountHis(op.Owner, op.Sender, &uoHis)
		if err != nil {
			log.Errorf("cache account(%s, %s) history error: %v", op.Owner, op.Sender, err)
			continue
		}
	}*/
	return nil, nil
}
