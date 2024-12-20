package etherman

import (
	"context"
	"database/sql"
	"errors"
	"github.com/0xPolygonHermez/zkevm-node/db"
	"github.com/OAB/database/pgstorage"
	"github.com/OAB/utils/log"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

var (
	// ErrStorageNotFound is used when the object is not found in the storage
	ErrStorageNotFound = errors.New("not found in the Storage")
	// ErrStorageNotRegister is used when the object is not found in the synchronizer
	ErrStorageNotRegister = errors.New("not registered storage")
	// ErrNilDBTransaction indicates the db transaction has not been properly initialized
	ErrNilDBTransaction = errors.New("database transaction not properly initialized")
	// ErrRestServerHealth indicates the health check of rest server failed
	ErrRestServerHealth = errors.New("not ready for the rest server")
	// ErrDepositNotSynced is used when the deposit is not synchronized in nodes
	ErrDepositNotSynced = errors.New("not synchronized deposit")
	// ErrNetworkNotRegister is used when the networkID is not registered in the bridge
	ErrNetworkNotRegister = errors.New("not registered network")
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

func (p *PostgresStorage) AddChainHeight(ctx context.Context, nid uint64, dbTx pgx.Tx) {
	addSQL := `INSERT INTO omni.chain_height (network_id, block_id, salt, create_at, update_at) VALUES ($1, $2, $3, $4);`
	e := p.getExecQuerier(dbTx)
	_, _ = e.Exec(ctx, addSQL, nid, 0, 0, time.Now(), time.Now())
}

func (p *PostgresStorage) GetChainHeight(ctx context.Context, nid uint64, dbTx pgx.Tx) uint64 {
	getSQL := `SELECT block_id FROM omni.chain_height where network_id=$1 LIMIT 1;`
	e := p.getExecQuerier(dbTx)
	var h uint64
	_ = e.QueryRow(ctx, getSQL, nid).Scan(&h)
	return h
}

func (p *PostgresStorage) SetChain(ctx context.Context, nid, val uint64, col string, dbTx pgx.Tx) error {
	getSQL := `UPDATE omni.chain_height SET $1 = $2, update_at = $3 WHERE network_id = $4;`
	e := p.getExecQuerier(dbTx)
	_, err := e.Exec(ctx, getSQL, col, val, time.Now(), nid)
	return err
}

func (p *PostgresStorage) GetChainSalt(ctx context.Context, nid uint64, dbTx pgx.Tx) uint64 {
	getSQL := `SELECT salt FROM omni.chain_height where network_id = $1 FOR UPDATE;`
	e := p.getExecQuerier(dbTx)
	var salt uint64
	_ = e.QueryRow(ctx, getSQL, nid).Scan(&salt)
	return salt
}
