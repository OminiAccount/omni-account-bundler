package pool

import (
	"context"
	"errors"
	"fmt"
	"github.com/OAB/database/pgstorage"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/state"
	"github.com/OAB/utils/hex"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
	"time"
)

var (
	// ErrNilDBTransaction indicates the db transaction has not been properly initialized
	ErrNilDBTransaction = errors.New("database transaction not properly initialized")
)

// PostgresStorage implements the Storage interface.
type PostgresStorage struct {
	*pgxpool.Pool
}

// getExecQuerier determines which execQuerier to use, dbTx or the main pgxpool
func (p *PostgresStorage) getExecQuerier(dbTx pgx.Tx) pgstorage.ExecQuerier {
	if dbTx != nil {
		return pgstorage.GetExecQuerierReconnect(dbTx)
	}
	return pgstorage.GetExecQuerierReconnect(p)
}

// NewPostgresStorage creates a new Storage DB
func NewPostgresStorage(db *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{db}
}

// Rollback rollbacks a db transaction.
func (p *PostgresStorage) Rollback(ctx context.Context, dbTx pgx.Tx) error {
	if dbTx != nil {
		return dbTx.Rollback(ctx)
	}

	return ErrNilDBTransaction
}

// Commit commits a db transaction.
func (p *PostgresStorage) Commit(ctx context.Context, dbTx pgx.Tx) error {
	if dbTx != nil {
		return dbTx.Commit(ctx)
	}
	return ErrNilDBTransaction
}

// BeginDBTransaction starts a transaction block.
func (p *PostgresStorage) BeginDBTransaction(ctx context.Context) (pgx.Tx, error) {
	return p.Begin(ctx)
}

func (p *PostgresStorage) DelProof(ctx context.Context, num, finalNum uint64, dbTx pgx.Tx) error {
	delSQL := "DELETE FROM omni.proof WHERE batch_num >= $1 AND final_batch_num <= $2"
	e := p.getExecQuerier(dbTx)
	_, err := e.Exec(ctx, delSQL, num, finalNum)
	return err
}

func (p *PostgresStorage) GetEarliestProof(ctx context.Context, dbTx pgx.Tx) (*ProofResult, error) {
	getSQL := `
		SELECT batch_num, final_batch_num, proof, public_values
		FROM omni.proof 
		ORDER BY batch_num ASC LIMIT 1;
	`
	e := p.getExecQuerier(dbTx)
	var pr ProofResult
	var proofStr, valuesStr string
	err := e.QueryRow(ctx, getSQL).Scan(&pr.Number, &pr.FinalNumber, &proofStr, &valuesStr)
	if err != nil {
		return nil, err
	}
	pr.Proof = common.Hex2Bytes(proofStr)
	pr.PublicValues = common.Hex2Bytes(valuesStr)
	return &pr, nil
}

func (p *PostgresStorage) AddProof(ctx context.Context, r *ProofResult, dbTx pgx.Tx) error {
	addSQL := `
		INSERT INTO omni.proof (
			batch_num, final_batch_num, proof, public_values, create_at, update_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		);
	`
	e := p.getExecQuerier(dbTx)
	_, err := e.Exec(ctx, addSQL, r.Number, r.FinalNumber, r.Proof.String(), r.PublicValues.String(), time.Now(), time.Now())
	return err
}

func (p *PostgresStorage) AddBatch(ctx context.Context, b *Batch, dbTx pgx.Tx) error {
	addSQL := `
		INSERT INTO omni.batch (
			batch_num, batch_hash, old_state_root, state_root, acc_input_hash, encoded, status, create_at, update_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		);
	`
	e := p.getExecQuerier(dbTx)
	_, err := e.Exec(ctx, addSQL, b.NewNumBatch, b.BatchHashData.Hex(), b.OldStateRoot.Hex(), b.NewStateRoot.Hex(),
		b.NewAccInputHash.Hex(), b.BatchL2Data.String(), b.Status, time.Now(), time.Now(),
	)
	return err
}

func (p *PostgresStorage) GetLatestBatch(ctx context.Context, status uint64, dbTx pgx.Tx) (Batch, error) {
	getSQL := `
		SELECT 
			batch_num, batch_hash, old_state_root, state_root, acc_input_hash, encoded, status, create_at
		FROM omni.batch 
		WHERE status = $1 ORDER BY batch_num DESC LIMIT 1;
	`
	e := p.getExecQuerier(dbTx)
	var b Batch
	var hashStr, oldStateStr, stateStr, aihStr, encodeStr string
	err := e.QueryRow(ctx, getSQL, status).Scan(
		&b.NewNumBatch, &hashStr, &oldStateStr, &stateStr, &aihStr, &encodeStr, &b.Status, &b.Timestamp,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		b.OldStateRoot = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
		b.NewStateRoot = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
		return b, nil
	} else if err != nil {
		return b, err
	}
	b.BatchHashData = common.HexToHash(hashStr)
	b.OldStateRoot = common.HexToHash(oldStateStr)
	b.NewStateRoot = common.HexToHash(stateStr)
	b.NewAccInputHash = common.HexToHash(aihStr)
	b.BatchL2Data = common.Hex2Bytes(strings.TrimPrefix(encodeStr, "0x"))
	return b, err
}

func (p *PostgresStorage) GetBatch(ctx context.Context, status, bs, be uint64, dbTx pgx.Tx) ([]*Batch, error) {
	getSQL := `
		SELECT 
		b.batch_num, b.batch_hash, b.old_state_root, b.state_root, b.acc_input_hash, b.encoded, b.status, b.create_at, o.*
		FROM omni.batch b
		LEFT JOIN omni.operation o ON o.batch_num = b.batch_num 
		WHERE b.status=$1 AND o.status=$2 AND b.batch_num>=$3 AND b.batch_num<=$4 ORDER BY b.batch_num,o.id ASC;
	`
	e := p.getExecQuerier(dbTx)
	var list []*Batch
	rows, err := e.Query(ctx, getSQL, status, state.BatchStatus, bs, be)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	listM := make(map[uint64]*Batch)
	for rows.Next() {
		b := &Batch{}
		op := &state.SignedUserOperation{}
		op.UserOperation = &state.UserOperation{}
		var opValue, hashStr, oldStateStr, stateStr, aihStr, encodeStr string
		var execMainGasPrice, execDestGasPrice, iexecMainGasPrice, iexecDestGasPrice string
		var updateAt time.Time
		err := rows.Scan(
			&b.NewNumBatch, &hashStr, &oldStateStr, &stateStr, &aihStr, &encodeStr, &b.Status, &b.Timestamp,
			&op.OpId, &op.Uid, &op.BatchNum, &op.Status, &op.Signature, &op.Did, &op.OperationType, &opValue,
			&op.Phase, &op.Exec.Nonce, &op.Exec.ChainId, &op.Exec.CallData, &op.Exec.MainChainGasLimit,
			&execMainGasPrice, &op.Exec.ZkVerificationGasLimit, &op.Exec.DestChainGasLimit,
			&execDestGasPrice, &op.InnerExec.Nonce, &op.InnerExec.ChainId, &op.InnerExec.CallData,
			&op.InnerExec.MainChainGasLimit, &iexecMainGasPrice, &op.InnerExec.ZkVerificationGasLimit,
			&op.InnerExec.DestChainGasLimit, &iexecDestGasPrice, &op.TimeAt, &updateAt,
		)
		if err != nil {
			return nil, err
		}
		if opValue == "" {
			opValue = "0x0"
		}
		op.OperationValue = (*hexutil.Big)(hex.DecodeBig(opValue))
		if execMainGasPrice == "" {
			execMainGasPrice = "0x0"
		}
		op.Exec.MainChainGasPrice = (*hexutil.Big)(hex.DecodeBig(execMainGasPrice))
		if execDestGasPrice == "" {
			execDestGasPrice = "0x0"
		}
		op.Exec.DestChainGasPrice = (*hexutil.Big)(hex.DecodeBig(execDestGasPrice))
		if iexecMainGasPrice == "" {
			iexecMainGasPrice = "0x0"
		}
		op.InnerExec.MainChainGasPrice = (*hexutil.Big)(hex.DecodeBig(iexecMainGasPrice))
		if iexecDestGasPrice == "" {
			iexecDestGasPrice = "0x0"
		}
		op.InnerExec.DestChainGasPrice = (*hexutil.Big)(hex.DecodeBig(iexecDestGasPrice))
		b.BatchHashData = common.HexToHash(hashStr)
		b.OldStateRoot = common.HexToHash(oldStateStr)
		b.NewStateRoot = common.HexToHash(stateStr)
		b.NewAccInputHash = common.HexToHash(aihStr)
		b.BatchL2Data = common.Hex2Bytes(strings.TrimPrefix(encodeStr, "0x"))
		if lb, ok := listM[b.NewNumBatch]; !ok {
			b.BatchL2SrcData = append(b.BatchL2SrcData, op)
			listM[b.NewNumBatch] = b
			list = append(list, b)
		} else {
			//log.Infof("%+v", op.UserOperation)
			lb.BatchL2SrcData = append(lb.BatchL2SrcData, op)
			listM[b.NewNumBatch] = lb
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return list, err
}

func (p *PostgresStorage) UpdateBatch(ctx context.Context, batchNum uint64, col string, val uint64, dbTx pgx.Tx) error {
	e := p.getExecQuerier(dbTx)
	updateSQL := "UPDATE omni.batch set %s=$1,update_at=$2 where batch_num=$3;"
	updateSQL = fmt.Sprintf(updateSQL, col)
	_, err := e.Exec(ctx, updateSQL, val, time.Now(), batchNum)
	return err
}

func (p *PostgresStorage) UpdateOp(ctx context.Context, opid uint64, col string, val uint64, dbTx pgx.Tx) error {
	e := p.getExecQuerier(dbTx)
	updateSQL := "UPDATE omni.operation set %s=$1,update_at=$2 where id=$3;"
	updateSQL = fmt.Sprintf(updateSQL, col)
	_, err := e.Exec(ctx, updateSQL, val, time.Now(), opid)
	return err
}

func (p *PostgresStorage) GetPendingOpList(ctx context.Context, l uint64, dbTx pgx.Tx) ([]*state.SignedUserOperation, error) {
	getSQL := `SELECT * FROM omni.operation WHERE status = $1 ORDER BY id ASC LIMIT $2;`
	e := p.getExecQuerier(dbTx)
	rows, err := e.Query(ctx, getSQL, state.PendingStatus, l)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var uos []*state.SignedUserOperation
	for rows.Next() {
		uo, err := scanOperation(rows)
		if err != nil {
			return nil, err
		}
		uos = append(uos, uo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return uos, nil
}

func (p *PostgresStorage) GetPendingOpCount(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	getSQL := `SELECT count(1) as c FROM omni.operation WHERE status = $1;`
	e := p.getExecQuerier(dbTx)
	var n uint64
	err := e.QueryRow(ctx, getSQL, state.PendingStatus).Scan(&n)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (p *PostgresStorage) GetUserOps(ctx context.Context, uid uint64, dbTx pgx.Tx) ([]*state.UserOperation, error) {
	getSQL := `SELECT * FROM omni.operation WHERE user_id = $1 ORDER BY id ASC;`
	e := p.getExecQuerier(dbTx)
	rows, err := e.Query(ctx, getSQL, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var uos []*state.UserOperation
	for rows.Next() {
		uo, err := scanOperation(rows)
		if err != nil {
			return nil, err
		}
		uos = append(uos, uo.UserOperation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return uos, nil
}

func scanOperation(row pgx.Row) (*state.SignedUserOperation, error) {
	op := &state.SignedUserOperation{}
	op.UserOperation = &state.UserOperation{}
	var updateAt time.Time
	var opValue, execMainGasPrice, execDestGasPrice, iexecMainGasPrice, iexecDestGasPrice string
	err := row.Scan(
		&op.OpId, &op.Uid, &op.BatchNum, &op.Status, &op.Signature, &op.Did, &op.OperationType, &opValue,
		&op.Phase, &op.Exec.Nonce, &op.Exec.ChainId, &op.Exec.CallData, &op.Exec.MainChainGasLimit,
		&execMainGasPrice, &op.Exec.ZkVerificationGasLimit, &op.Exec.DestChainGasLimit,
		&execDestGasPrice, &op.InnerExec.Nonce, &op.InnerExec.ChainId, &op.InnerExec.CallData,
		&op.InnerExec.MainChainGasLimit, &iexecMainGasPrice, &op.InnerExec.ZkVerificationGasLimit,
		&op.InnerExec.DestChainGasLimit, &iexecDestGasPrice, &op.TimeAt, &updateAt,
	)
	if err != nil {
		return nil, err
	}
	if opValue == "" {
		opValue = "0x0"
	}
	op.OperationValue = (*hexutil.Big)(hex.DecodeBig(opValue))
	if execMainGasPrice == "" {
		execMainGasPrice = "0x0"
	}
	op.Exec.MainChainGasPrice = (*hexutil.Big)(hex.DecodeBig(execMainGasPrice))
	if execDestGasPrice == "" {
		execDestGasPrice = "0x0"
	}
	op.Exec.DestChainGasPrice = (*hexutil.Big)(hex.DecodeBig(execDestGasPrice))
	if iexecMainGasPrice == "" {
		iexecMainGasPrice = "0x0"
	}
	op.InnerExec.MainChainGasPrice = (*hexutil.Big)(hex.DecodeBig(iexecMainGasPrice))
	if iexecDestGasPrice == "" {
		iexecDestGasPrice = "0x0"
	}
	op.InnerExec.DestChainGasPrice = (*hexutil.Big)(hex.DecodeBig(iexecDestGasPrice))
	log.Infof("%+v %+v", op.Signature, op.UserOperation)
	return op, nil
}
