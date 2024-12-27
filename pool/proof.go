package pool

import (
	"fmt"
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/etherman/contracts/SyncRouter"
	"github.com/OAB/state"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/jackc/pgx/v4"
	"math/big"
)

func (p *Pool) processBatch() {
	log.Infof("processBatch...")
	oldBatch, err := p.db.GetBatchNoOp(p.ctx, state.PendingStatus, nil)
	if err != nil {
		log.Errorf("get latest batch err: %+v", err)
		return
	}

	batch, err := NewBatch(oldBatch.NewNumBatch + 1)
	if err != nil {
		log.Errorf("new batch err: %+v", err)
		return
	}
	// Set old smt root
	batch.SetOldStateRoot(oldBatch.NewStateRoot)
	batch.SetOldAccInputHash(oldBatch.NewAccInputHash)
	dbTx, err := p.db.BeginDBTransaction(p.ctx)
	if err != nil {
		log.Errorf("use db transaction err: %+v", err)
		return
	}
	list, err := p.db.GetPendingOpList(p.ctx, p.cfg.MaxOps, dbTx)
	if err != nil {
		_ = dbTx.Rollback(p.ctx)
		log.Errorf("get pending op list err: %+v", err)
		return
	}
	for i, uo := range list {
		isR, err := p.calculateOp(batch.NewNumBatch, uo, dbTx)
		if err != nil {
			_ = dbTx.Rollback(p.ctx)
			log.Errorf("check op err: %+v", err)
			return
		}
		if isR {
			_ = i
			//list = append(list[:i], list[i+1:]...)
		}
	}
	// Set BatchL2Data and BatchHashData
	batch.SetBatchL2Data(list)
	// Set NewAccInputHash
	batch.SetNewAccInputHash(list)
	// Set new state root
	batch.SetNewStateRoot(common.BigToHash(p.state.GetTree().LastRoot()))
	err = p.db.AddBatch(p.ctx, batch, nil)
	if err != nil {
		_ = dbTx.Rollback(p.ctx)
		log.Errorf("add batch err: %+v", err)
		return
	}
	_ = dbTx.Commit(p.ctx)
	log.Infof("Successfully sealing a batch, number: %d, stateRoot: %s", batch.NewNumBatch, batch.NewStateRoot.Hex())
}

func (p *Pool) calculateOp(batchNum uint64, uo *state.SignedUserOperation, dbTx pgx.Tx) (bool, error) {
	log.Infof("operation: %+v ===============", uo)
	if uo.Phase != state.PhaseFirst {
		ai, _ := p.state.GetAccountInfoByChain(uo.Owner, uo.InnerExec.ChainId.Uint64())
		if ai == nil {
			return true, nil
		}
		return false, nil
	}
	/*uoOwner := uo.RecoverAddress()
	if uoOwner.Hex() != uo.Owner.Hex() {
		log.Errorf("recover address(%s) != owner(%s)", uoOwner.Hex(), uo.Owner.Hex())
		_ = p.db.UpdateOp(p.ctx, uo.OpId, "status", state.FailedStatus, nil)
		return true, nil
	}*/
	ai, _ := p.state.GetAccountInfoByChain(uo.Owner, uo.Exec.ChainId.Uint64())
	if ai == nil {
		return true, nil
	}
	// balance
	var balance big.Int
	var balanceU256 *uint256.Int
	balanceU256, err := p.state.GetTree().GetAccountBalance(uo.Sender)
	if err != nil {
		return false, err
	}
	balance.SetBytes(balanceU256.Bytes())
	log.Infof("oldBalance: %d", &balance)

	if uo.OperationType == 1 {
		log.Infof("balance %d + operationValue %d", &balance, uo.OperationValue.Uint64())
		balance.Add(&balance, uo.OperationValue.ToInt())
	} else if uo.OperationType == 2 {
		if balance.Cmp(uo.OperationValue.ToInt()) < 0 {
			log.Errorf("current account balance %d is insufficient to withdraw %d",
				&balance, uo.OperationValue.Uint64())
			_ = p.db.UpdateOp(p.ctx, uo.OpId, "status", state.FailedStatus, nil)
			return true, nil
		}
		log.Infof("balance %d - operationValue %d", &balance, uo.OperationValue.Uint64())
		balance.Sub(&balance, uo.OperationValue.ToInt())
	}

	log.Infof("balance %d - CalculateGasUsed %s", &balance, uo.CalculateGasUsed())
	balance.Sub(&balance, uo.CalculateGasUsed())

	if _, err = p.state.GetTree().SetAccountBalance(uo.Sender.String(), &balance); err != nil {
		return false, err
	}
	log.Infof("newBalance: %d", &balance)

	// nonce
	var nonce big.Int
	nonceU256, err := p.state.GetTree().GetAccountNonce(uo.Sender, uo.Exec.ChainId.String())
	if err != nil {
		return false, err
	}
	nonce.SetBytes(nonceU256.Bytes())
	log.Infof("phase first oldNonce: %d", &nonce)
	nonce.Add(&nonce, big.NewInt(1))
	log.Infof("phase first newNonce: %d", &nonce)
	_, err = p.state.GetTree().SetAccountNonce(uo.Sender.String(), &nonce, uo.Exec.ChainId.String())
	if err != nil {
		return false, err
	}
	if uo.InnerExec.ChainId.Uint64() > 0 {
		var nonce2 big.Int
		nonceU256, err = p.state.GetTree().GetAccountNonce(uo.Sender, uo.InnerExec.ChainId.String())
		if err != nil {
			return false, err
		}
		nonce2.SetBytes(nonceU256.Bytes())
		log.Infof("phase second oldNonce: %d", &nonce2)
		nonce2.Add(&nonce2, big.NewInt(1))
		log.Infof("phase second newNonce: %d", &nonce2)
		_, err = p.state.GetTree().SetAccountNonce(uo.Sender.String(), &nonce2, uo.InnerExec.ChainId.String())
		if err != nil {
			return false, err
		}
	}
	err = p.db.UpdateOp(p.ctx, uo.OpId, "batch_num", batchNum, dbTx)
	if err != nil {
		return false, err
	}
	err = p.db.UpdateOp(p.ctx, uo.OpId, "status", state.BatchStatus, dbTx)
	if err != nil {
		return false, err
	}
	return false, nil
}

func (p *Pool) verifyProof() error {
	log.Info("verify proof...")
	verifyBatch, err := p.db.GetBatchNoOp(p.ctx, state.SuccessStatus, nil)
	if err != nil {
		return err
	}
	res, err := p.db.GetEarliestProof(p.ctx, nil)
	if err != nil {
		return err
	}
	allOp := []*state.SignedUserOperation{}
	batches := []EntryPoint.BaseStructBatchData{}
	extraInfo := EntryPoint.BaseStructChainsExecuteInfo{}
	chainExtra := make(map[uint64]EntryPoint.BaseStructChainExecuteExtra)
	for i := res.Number; i <= res.FinalNumber; i++ {
		if i <= verifyBatch.NewNumBatch {
			return fmt.Errorf("batch number too low")
		}
		batch, err := p.db.GetBatch(p.ctx, state.BatchStatus, i, i, nil)
		if err != nil {
			return err
		}
		if batch == nil || len(batch) < 1 {
			return fmt.Errorf("batch(%d) no found data", i)
		}
		if i == res.FinalNumber {
			extraInfo.NewStateRoot = batch[0].NewStateRoot
		}
		suo := batch[0].BatchL2SrcData
		var puo []EntryPoint.BaseStructPackedUserOperation
		for _, op := range suo {
			puo = append(puo, op.ToEntryPointOp())
			chainID := op.Exec.ChainId
			if op.Phase == state.PhaseSecond {
				chainID = op.InnerExec.ChainId
			}
			fee, err := p.ethereum.EstimateGas(chainID.Uint64(), op.CalculateGasUsed(),
				[]SyncRouter.BaseStructPackedUserOperation{op.ToSyncRouterOp()})
			if err != nil {
				return err
			}
			log.Infof("EstimateGas fee: %d", fee.Uint64())
			if cee, ok := chainExtra[chainID.Uint64()]; ok {
				cee.ChainFee += fee.Uint64()
				cee.ChainUserOperationsNumber++
				chainExtra[chainID.Uint64()] = cee
			} else {
				chainExtra[chainID.Uint64()] = EntryPoint.BaseStructChainExecuteExtra{
					ChainId:                   chainID.Uint64(),
					ChainFee:                  fee.Uint64(),
					ChainUserOperationsNumber: 1,
				}
			}
		}
		batches = append(batches, EntryPoint.BaseStructBatchData{
			UserOperations: puo,
			AccInputHash:   batch[0].NewAccInputHash,
		})
		allOp = append(allOp, suo...)
	}
	if len(batches) < 1 {
		log.Warnf("batch(%d - %d) no data", res.Number, res.FinalNumber)
		err = p.db.DelProof(p.ctx, res.Number, res.FinalNumber, nil)
		return err
	}
	for _, v := range chainExtra {
		extraInfo.ChainExtra = append(extraInfo.ChainExtra, v)
	}
	tx, err := p.ethereum.UpdateEntryPointRoot(res.Proof, batches, extraInfo)
	if err != nil {
		return err
	}
	log.Infof("verify tx: %s", tx.Hex())
	dbTx, err := p.db.BeginDBTransaction(p.ctx)
	if err != nil {
		log.Errorf("use db transaction err: %+v", err)
		return err
	}
	log.Infof("delete proof(%d - %d)", res.Number, res.FinalNumber)
	err = p.db.DelProof(p.ctx, res.Number, res.FinalNumber, dbTx)
	if err != nil {
		_ = dbTx.Rollback(p.ctx)
		return err
	}
	log.Infof("update batch status")
	for i := res.Number; i <= res.FinalNumber; i++ {
		err = p.db.UpdateBatch(p.ctx, i, "status", state.SuccessStatus, dbTx)
		if err != nil {
			_ = dbTx.Rollback(p.ctx)
			return err
		}
	}
	log.Infof("update user op and history")
	for _, op := range allOp {
		err = p.db.UpdateOp(p.ctx, op.OpId, "status", state.SuccessStatus, dbTx)
		if err != nil {
			_ = dbTx.Rollback(p.ctx)
			return err
		}
		if op.OperationType == state.DepositAction {
			continue
		}
		if op.Phase == state.PhaseSecond {
			err = p.state.GetHisIns().UpdateAccountHis(op.Uid, op.Did,
				map[string]interface{}{
					"target_chain": op.InnerExec.ChainId.Uint64(),
					"target_hash":  tx.Hex(),
					"status":       state.CheckToTxStatus,
				},
			)
			if err != nil {
				_ = dbTx.Rollback(p.ctx)
				return err
			}
			continue
		}
		err = p.state.GetHisIns().UpdateAccountHis(op.Uid, op.Did,
			map[string]interface{}{
				"source_hash": tx.Hex(),
				"target_hash": tx.Hex(),
				"status":      state.CheckFromTxStatus,
			},
		)
		if err != nil {
			_ = dbTx.Rollback(p.ctx)
			log.Errorf("UpdateAccountHis err: %+v", err)
			return err
		}
		/*uoHis := state.ToUserOperationHis(tx.Hex(), op.UserOperation)
		err = p.state.GetHisIns().SaveAccountHis(&uoHis)
		if err != nil {
			log.Errorf("add account(%s, %s) history error: %v", op.Owner, op.Sender, err)
			continue
		}*/
	}
	err = dbTx.Commit(p.ctx)
	return err
}
