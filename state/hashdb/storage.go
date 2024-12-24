package hashdb

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/OAB/database/pgstorage"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
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

type HashDB struct {
	Hash string `json:"hash"`
	Data string `json:"data"`
}

func (p *PostgresStorage) insertHashDB(ctx context.Context, hash string, data string, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := `INSERT INTO omni.hash_db (hash, data) VALUES ($1, $2) ON CONFLICT (hash) DO UPDATE SET data = EXCLUDED.data;`
	_, err := db.Exec(ctx, query, hash, data)
	return err
}

func (p *PostgresStorage) updateHashDB(ctx context.Context, hash string, data string, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := `UPDATE omni.hash_db SET data = $1 WHERE hash = $2;`
	_, err := db.Exec(ctx, query, data, hash)
	return err
}

func (p *PostgresStorage) deleteHashDB(ctx context.Context, hash string, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := `DELETE FROM omni.hash_db WHERE hash = $1;`
	_, err := db.Exec(ctx, query, hash)
	return err
}

func (p *PostgresStorage) getHashDB(ctx context.Context, hash string, dbTx pgx.Tx) ([]string, error) {
	db := p.getExecQuerier(dbTx)
	query := `SELECT data FROM omni.hash_db WHERE hash = $1;`
	var d string
	err := db.QueryRow(ctx, query, hash).Scan(&d)
	var res []string
	err = json.Unmarshal([]byte(d), &res)
	return res, err
}

func (p *PostgresStorage) getHashDBCount(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	db := p.getExecQuerier(dbTx)
	query := `SELECT count(1) as c FROM omni.hash_db;`
	var d uint64
	err := db.QueryRow(ctx, query).Scan(&d)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, nil
	}
	return d, err
}

func (p *PostgresStorage) getHashDBAllData(ctx context.Context, dbTx pgx.Tx) ([]HashDB, error) {
	db := p.getExecQuerier(dbTx)
	query := `SELECT hash, data FROM omni.hash_db;`
	rows, err := db.Query(ctx, query)
	var d []HashDB
	if errors.Is(err, pgx.ErrNoRows) {
		return d, nil
	}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var hd HashDB
		err := rows.Scan(&hd.Hash, &hd.Data)
		if err != nil {
			return nil, err
		}
		d = append(d, hd)
	}
	return d, err
}

type HashAccVal struct {
	Hash string `json:"hash"`
	Data string `json:"data"`
}

func (p *PostgresStorage) insertHashAccVal(ctx context.Context, hash, data string, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := "INSERT INTO omni.hash_accval (hash, data) VALUES ($1, $2) ON CONFLICT (hash) DO UPDATE SET data = EXCLUDED.data;"
	_, err := db.Exec(ctx, query, hash, data)
	return err
}

func (p *PostgresStorage) getHashAccVal(ctx context.Context, hash string, dbTx pgx.Tx) ([]string, error) {
	db := p.getExecQuerier(dbTx)
	query := "SELECT data FROM omni.hash_accval WHERE hash = $1;"
	var d string
	err := db.QueryRow(ctx, query, hash).Scan(&d)
	var res []string
	err = json.Unmarshal([]byte(d), &res)
	return res, err
}

func (p *PostgresStorage) updateHashAccVal(ctx context.Context, record HashAccVal, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := "UPDATE omni.hash_accval SET data = $1 WHERE hash = $2;"
	_, err := db.Exec(ctx, query, record.Data, record.Hash)
	return err
}

func (p *PostgresStorage) deleteHashAccVal(ctx context.Context, hash string, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := "DELETE FROM omni.hash_accval WHERE hash = $1;"
	_, err := db.Exec(ctx, query, hash)
	return err
}

type HashKeySource struct {
	Hash string
	Data []byte
}

func (p *PostgresStorage) insertHashKeySource(ctx context.Context, k string, d []byte, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := `INSERT INTO omni.hash_keysource (hash, data) VALUES ($1, $2) ON CONFLICT (hash) DO UPDATE SET data = EXCLUDED.data;`
	_, err := db.Exec(ctx, query, k, d)
	return err
}

func (p *PostgresStorage) getHashKeySource(ctx context.Context, hash string, dbTx pgx.Tx) ([]byte, error) {
	db := p.getExecQuerier(dbTx)
	query := `SELECT data FROM omni.hash_keysource WHERE hash = $1;`
	var d []byte
	err := db.QueryRow(ctx, query, hash).Scan(&d)
	return d, err
}

func (p *PostgresStorage) updateHashKeySource(ctx context.Context, record HashKeySource, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := `UPDATE omni.hash_keysource SET data = $2 WHERE hash = $1;`
	_, err := db.Exec(ctx, query, record.Hash, record.Data)
	return err
}

func (p *PostgresStorage) deleteHashKeySource(ctx context.Context, hash string, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := `DELETE FROM omni.hash_keysource WHERE hash = $1;`
	_, err := db.Exec(ctx, query, hash)
	return err
}

// HashHashKey represents a record in the omni.hash_hashkey table
type HashHashKey struct {
	Hash string
	Data []byte
}

// InsertHashHashKey inserts a new record into the omni.hash_hashkey table
func (p *PostgresStorage) insertHashHashKey(ctx context.Context, k string, d []byte, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := "INSERT INTO omni.hash_hashkey (hash, data) VALUES ($1, $2) ON CONFLICT (hash) DO UPDATE SET data = EXCLUDED.data;"
	_, err := db.Exec(ctx, query, k, d)
	return err
}

// GetHashHashKey retrieves a record from the omni.hash_hashkey table by hash
func (p *PostgresStorage) getHashHashKey(ctx context.Context, hash string, dbTx pgx.Tx) ([]byte, error) {
	db := p.getExecQuerier(dbTx)
	query := "SELECT data FROM omni.hash_hashkey WHERE hash = $1;"
	var d []byte
	err := db.QueryRow(ctx, query, hash).Scan(&d)
	return d, err
}

// updateHashHashKey updates the data of an existing record in the omni.hash_hashkey table
func (p *PostgresStorage) updateHashHashKey(ctx context.Context, record HashHashKey, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := "UPDATE omni.hash_hashkey SET data = $1 WHERE hash = $2;"
	_, err := db.Exec(ctx, query, record.Data, record.Hash)
	return err
}

// DeleteHashHashKey deletes a record from the omni.hash_hashkey table by hash
func (p *PostgresStorage) deleteHashHashKey(ctx context.Context, hash string, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := "DELETE FROM omni.hash_hashkey WHERE hash = $1;"
	_, err := db.Exec(ctx, query, hash)
	return err
}

type HashCode struct {
	Hash string
	Data []byte
}

// InsertHashCode inserts a new record into the omni.hash_code table
func (p *PostgresStorage) insertHashCode(ctx context.Context, k string, d []byte, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := `INSERT INTO omni.hash_code (hash, data) VALUES ($1, $2) ON CONFLICT (hash) DO UPDATE SET data = EXCLUDED.data;`
	_, err := db.Exec(ctx, query, k, d)
	return err
}

// GetHashCode retrieves a record from the omni.hash_code table by hash
func (p *PostgresStorage) getHashCode(ctx context.Context, hash string, dbTx pgx.Tx) ([]byte, error) {
	db := p.getExecQuerier(dbTx)
	query := `SELECT data FROM omni.hash_code WHERE hash = $1;`
	var d []byte
	err := db.QueryRow(ctx, query, hash).Scan(&d)
	return d, err
}

// UpdateHashCode updates an existing record in the omni.hash_code table
func (p *PostgresStorage) updateHashCode(ctx context.Context, hashCode HashCode, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := `UPDATE omni.hash_code SET data = $1 WHERE hash = $2;`
	_, err := db.Exec(ctx, query, hashCode.Data, hashCode.Hash)
	return err
}

// DeleteHashCode deletes a record from the omni.hash_code table by hash
func (p *PostgresStorage) deleteHashCode(ctx context.Context, hash string, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := `DELETE FROM omni.hash_code WHERE hash = $1;`
	_, err := db.Exec(ctx, query, hash)
	return err
}

// HashInfo represents a row in the omni.hash_info table
type HashInfo struct {
	LastRoot string
	Depth    uint8
}

// GetHashInfo retrieves a row from the omni.hash_info table
func (p *PostgresStorage) getHashInfo(ctx context.Context, dbTx pgx.Tx) (*HashInfo, error) {
	db := p.getExecQuerier(dbTx)
	query := `SELECT lastroot, depth FROM omni.hash_info;`
	var hi HashInfo
	err := db.QueryRow(ctx, query).Scan(&hi.LastRoot, &hi.Depth)
	return &hi, err
}

func (p *PostgresStorage) updateHashInfoRoot(ctx context.Context, r string, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := `UPDATE omni.hash_info SET lastroot = $1;`
	_, err := db.Exec(ctx, query, r)
	return err
}

func (p *PostgresStorage) updateHashInfoDepth(ctx context.Context, d uint8, dbTx pgx.Tx) error {
	db := p.getExecQuerier(dbTx)
	query := `UPDATE omni.hash_info SET depth = $1;`
	_, err := db.Exec(ctx, query, d)
	return err
}
