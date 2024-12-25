package etherman

import (
	"context"
	"errors"
	"fmt"
	"github.com/OAB/database/pgstorage"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
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
	addSQL := `INSERT INTO omni.chain_height (network_id, block_id, salt, create_at, update_at) VALUES ($1, $2, $3, $4, $5);`
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
	getSQL := "UPDATE omni.chain_height SET %s = $1, update_at = $2 WHERE network_id = $3;"
	getSQL = fmt.Sprintf(getSQL, col)
	e := p.getExecQuerier(dbTx)
	_, err := e.Exec(ctx, getSQL, val, time.Now(), nid)
	return err
}

func (p *PostgresStorage) GetChainSalt(ctx context.Context, nid uint64, dbTx pgx.Tx) uint64 {
	getSQL := `SELECT salt FROM omni.chain_height where network_id = $1 FOR UPDATE;`
	e := p.getExecQuerier(dbTx)
	var salt uint64
	_ = e.QueryRow(ctx, getSQL, nid).Scan(&salt)
	return salt + 2
}

func (p *PostgresStorage) UpdateUserFailedSalt(ctx context.Context, user string, nid uint64, idDel bool, dbTx pgx.Tx) error {
	var err error
	var flag bool
	if dbTx == nil {
		flag = true
		dbTx, err = p.BeginDBTransaction(ctx)
		if err != nil {
			return err
		}
	}
	var failedSalt string
	getSQL := `SELECT failed_salt FROM omni.user where owner = $1 FOR UPDATE;`
	e := p.getExecQuerier(dbTx)
	err = e.QueryRow(ctx, getSQL, strings.ToLower(user)).Scan(&failedSalt)
	if err != nil {
		if flag {
			_ = dbTx.Rollback(ctx)
		}
		return err
	}
	nidStr := fmt.Sprintf("%d", nid)
	if len(failedSalt) > 0 {
		if idDel {
			cids := strings.Split(failedSalt, ",")
			for i, cid := range cids {
				if cid == nidStr {
					cids = append(cids[:i], cids[i+1:]...)
					break
				}
			}
			failedSalt = strings.Join(cids, ",")
		} else if !strings.Contains(failedSalt, nidStr) {
			failedSalt += "," + nidStr
		}
	} else if !idDel {
		failedSalt = nidStr
	}
	updateSQL := `UPDATE omni.user SET failed_salt = $1 where owner = $2;`
	_, err = e.Exec(ctx, updateSQL, failedSalt, strings.ToLower(user))
	if err != nil {
		if flag {
			_ = dbTx.Rollback(ctx)
		}
		return err
	}
	if flag {
		return dbTx.Commit(ctx)
	}
	return nil
}

type UserFailedSalt struct {
	User       string
	Salt       uint64
	FailedSalt string
}

func (p *PostgresStorage) GetFailedSalts(ctx context.Context, dbTx pgx.Tx) []UserFailedSalt {
	getSQL := `SELECT owner, salt, failed_salt FROM omni.user where failed_salt != '';`
	e := p.getExecQuerier(dbTx)
	rows, err := e.Query(ctx, getSQL)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var list []UserFailedSalt
	for rows.Next() {
		var ufs UserFailedSalt
		err := rows.Scan(&ufs.User, &ufs.Salt, &ufs.FailedSalt)
		if err != nil {
			return nil
		}
		list = append(list, ufs)
	}
	return list
}

func (p *PostgresStorage) IsExistChainByOwner(ctx context.Context, owner string, nid uint64, dbTx pgx.Tx) (bool, error) {
	getSQL := `
		SELECT u.id FROM omni.user u INNER JOIN omni.user_info ui ON u.id=ui.user_id WHERE u.owner=$1 AND ui.network_id=$2 LIMIT 1;
	`
	e := p.getExecQuerier(dbTx)
	var uid uint64
	err := e.QueryRow(ctx, getSQL, strings.ToLower(owner), nid).Scan(&uid)
	if errors.Is(err, pgx.ErrNoRows) || uid < 1 {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
