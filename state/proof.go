package state

import (
	"fmt"
	"github.com/OAB/etherman/contracts/EntryPoint"
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

	for _, userOperation := range batchContext.SignedUserOperations() {
		log.Infof("Address: %s ,Nonce: %d ,ChainId %s ,OperationType: %d ,operationValue: %d ===============", userOperation.Owner, userOperation.Nonce, userOperation.ChainId.String(), userOperation.OperationType, userOperation.OperationValue)

		// balance
		var balance big.Int
		balanceU256, err := s.tree.GetAccountBalance(userOperation.Sender)
		if err != nil {
			return err
		}
		balance.SetBytes(balanceU256.Bytes())
		log.Infof("oldBalance: %d", &balance)
		operationValue := new(big.Int).SetUint64(userOperation.OperationValue.Uint64())

		if userOperation.OperationType == 1 {
			log.Infof("balance %d + operationValue %d", &balance, operationValue)
			balance.Add(&balance, operationValue)
		} else if userOperation.OperationType == 2 {
			if balance.Cmp(operationValue) < 0 {
				log.Infof("current account balance %s is insufficient to withdraw %s", &balance, operationValue)
			}
			log.Infof("balance %d - operationValue %d", &balance, operationValue.Uint64())
			balance.Sub(&balance, operationValue)
		}

		log.Infof("balance %d - CalculateGasUsed %d", &balance, userOperation.CalculateGasUsed())
		balance.Sub(&balance, userOperation.CalculateGasUsed())

		if _, err = s.tree.SetAccountBalance(userOperation.Sender.String(), &balance); err != nil {
			return err
		}
		log.Infof("newBalance: %d", &balance)

		// nonce
		var nonce big.Int
		nonceU256, err := s.tree.GetAccountNonce(userOperation.Sender, userOperation.ChainId.String())
		if err != nil {
			return err
		}
		nonce.SetBytes(nonceU256.Bytes())
		log.Infof("oldNonce: %d", &nonce)
		nonce.Add(&nonce, big.NewInt(1))
		log.Infof("newNonce: %d", &nonce)
		_, err = s.tree.SetAccountNonce(userOperation.Sender.String(), &nonce, userOperation.ChainId.String())
		if err != nil {
			return err
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
	if !s.provenQueue.IsEmpty() {
		log.Info("Executing batch")
		res, ok := s.provenQueue.Dequeue()
		if !ok {
			return nil, nil
		}
		batches := EntryPoint.IEntryPointBatchData{}
		var suo []*pool.SignedUserOperation
		for i := res.Number; i <= res.FinalNumber; i++ {
			batch, err := s.storage.loadBatch(i)
			if err != nil {
				return nil, err
			}
			if batch == nil {
				return nil, fmt.Errorf("batch(%d) no cache data", i)
			}
			if i == res.Number {
				batches.OldStateRoot = batch.OldStateRoot
			} else if i == res.FinalNumber {
				batches.NewStateRoot = batch.NewStateRoot
			}
			suo = append(suo, batch.BatchL2SrcData...)
		}
		if len(suo) < 1 {
			log.Warnf("batch(%d - %d) no data", res.Number, res.FinalNumber)
			s.storage.VerifyBatchNumber = res.FinalNumber
			return nil, nil
		}
		var puo []EntryPoint.PackedUserOperation
		chainExtra := make(map[uint64]EntryPoint.IEntryPointChainExecuteExtra)
		for _, op := range suo {
			puo = append(puo, op.ToEntryPointOp())
			fee, err := s.ethereum.EstimateGas(op.ChainId.Uint64(), op.CalculateGasUsed(), op.Encode())
			if err != nil {
				return &res, err
			}
			if cee, ok := chainExtra[op.ChainId.Uint64()]; ok {
				cee.ChainFee.Add(cee.ChainFee, fee)
				cee.ChainUserOperationsNumber.Add(cee.ChainUserOperationsNumber, big.NewInt(1))
				chainExtra[op.ChainId.Uint64()] = cee
			} else {
				chainExtra[op.ChainId.Uint64()] = EntryPoint.IEntryPointChainExecuteExtra{
					ChainFee:                  fee,
					ChainUserOperationsNumber: big.NewInt(1),
				}
			}
		}
		extraInfo := EntryPoint.IEntryPointChainsExecuteInfo{}
		for _, v := range chainExtra {
			extraInfo.ChainExtra = append(extraInfo.ChainExtra, v)
		}
		batches.AccInputHash = batchHashData(suo)
		batches.UserOperations = puo
		tx, err := s.ethereum.UpdateEntryPointRoot(res.Proof, []EntryPoint.IEntryPointBatchData{batches}, extraInfo)
		if err != nil {
			return &res, err
		}
		for i := res.Number; i <= res.FinalNumber; i++ {
			err = s.storage.clearBatch(i)
			if err != nil {
				return nil, err
			}
		}
		for _, op := range suo {
			uoHis := types.ToUserOperationHis(tx.Hex(), *op.UserOperation)
			err = s.SaveAccountHis(op.Owner, op.Sender, uoHis)
			if err != nil {
				log.Errorf("cache account(%s, %s) history error: %v", op.Owner, op.Sender, err)
				continue
			}
		}
	}
	return nil, nil
}
