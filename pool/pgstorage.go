package pool

import (
	"context"
	"errors"
	"github.com/OAB/database/pgstorage"
	"github.com/OAB/utils/log"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
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

func (p *PostgresStorage) GetUserOps(ctx context.Context, uid uint64, dbTx pgx.Tx) ([]*UserOperation, error) {
	const getSQL = `SELECT * FROM omni.operation WHERE user_id = $1 ORDER BY id ASC;`
	e := p.getExecQuerier(dbTx)
	rows, err := e.Query(ctx, getSQL, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var uos []*UserOperation
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

func scanOperation(row pgx.Row) (*UserOperation, error) {
	op := &UserOperation{}
	var (
		signature string
	)
	err := row.Scan(
		&op.OpId, &op.Uid, &op.BatchNum, &op.Status, &signature, &op.Did, &op.OperationType, &op.OperationValue,
		&op.Phase, &op.Exec.Nonce, &op.Exec.ChainId, &op.Exec.CallData, &op.Exec.MainChainGasLimit,
		&op.Exec.MainChainGasPrice, &op.Exec.ZkVerificationGasLimit, &op.Exec.DestChainGasLimit,
		&op.Exec.DestChainGasPrice, &op.InnerExec.Nonce, &op.InnerExec.ChainId, &op.InnerExec.CallData,
		&op.InnerExec.MainChainGasLimit, &op.InnerExec.MainChainGasPrice, &op.InnerExec.ZkVerificationGasLimit,
		&op.InnerExec.DestChainGasLimit, &op.InnerExec.DestChainGasPrice,
	)
	if err != nil {
		return nil, err
	}
	log.Infof("%+v", op)
	return op, nil
}
