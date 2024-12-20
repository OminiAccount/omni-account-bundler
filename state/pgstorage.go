package state

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/0xPolygonHermez/zkevm-node/db"
	"github.com/OAB/database/pgstorage"
	"github.com/OAB/pool"
	"github.com/OAB/utils/hex"
	"github.com/OAB/utils/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"math/big"
	"strconv"
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

func (p *PostgresStorage) GetUserInfoLock(ctx context.Context, uid, nid uint64, dbTx pgx.Tx) (uint64, *big.Int, error) {
	e := p.getExecQuerier(dbTx)
	getSQL := `SELECT nonce, gas FROM omni.user_info WHERE user_id = $1 AND network_id = $2 FOR UPDATE;`
	var tNonce uint64
	var tGas string
	err := e.QueryRow(ctx, getSQL, uid, nid).Scan(&tNonce, &tGas)
	if err != nil {
		return 0, nil, err
	}
	tGasInt, ok := big.NewInt(0).SetString(tGas, 10)
	if !ok {
		return 0, nil, fmt.Errorf("convert gas error")
	}
	return tNonce, tGasInt, nil
}

func (p *PostgresStorage) ModUserInfo(ctx context.Context, uid, nid, nonce uint64, gas string, dbTx pgx.Tx) error {
	e := p.getExecQuerier(dbTx)
	updateSQL := "UPDATE omni.user_info "
	setSQL := ""
	where := []interface{}{}
	if nonce > 0 {
		setSQL += "nonce=$" + strconv.Itoa(len(where)+1) + ","
		where = append(where, nonce)
	}
	if gas != "" {
		setSQL += "gas=$" + strconv.Itoa(len(where)+1) + ","
		where = append(where, gas)
	}
	if len(where) > 0 {
		setSQL = "set " + setSQL
		setSQL = strings.TrimRight(setSQL, ",")
	} else {
		return nil
	}
	updateSQL += " where user_id=$" + strconv.Itoa(len(where)+1) + " and network_id=$" + strconv.Itoa(len(where)+2) + ";"
	where = append(where, uid, nid)
	_, err := e.Exec(ctx, updateSQL, where...)
	return err
}

func (p *PostgresStorage) AddUser(ctx context.Context, am AccountMapping, dbTx pgx.Tx) error {
	addSQL := "INSERT INTO omni.user (owner, account, salt, create_at) VALUES ($1, $2, $3, $4);"
	e := p.getExecQuerier(dbTx)
	_, err := e.Exec(ctx, addSQL, am.User.Hex(), am.Account.Hex(), am.Salt, time.Now())
	return err
}

func (p *PostgresStorage) AddUserInfo(ctx context.Context, uid, nid uint64, dbTx pgx.Tx) error {
	addSQL := "INSERT INTO omni.user_info (user_id, network_id, nonce, gas, create_at, update_at) VALUES ($1, $2, $3, $4, $5, $6);"
	e := p.getExecQuerier(dbTx)
	_, err := e.Exec(ctx, addSQL, uid, nid, 0, "0", time.Now(), time.Now())
	return err
}

func (p *PostgresStorage) GetUserInfoByAA(ctx context.Context, owner, account string, dbTx pgx.Tx) (*AccountInfo, error) {
	const getSQL = `
		SELECT
			u.id, u.salt, ui.network_id, ui.nonce, ui.gas
		FROM
			omni.user u
		LEFT JOIN
			omni.user_info ui
		ON
			u.id = ui.user_id
		WHERE
			u.owner = $1 AND u.account = $2;
	`
	e := p.getExecQuerier(dbTx)
	rows, err := e.Query(ctx, getSQL, owner, account)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ai := &AccountInfo{
		User:    common.HexToAddress(owner),
		Account: common.HexToAddress(account),
		Chain:   map[uint64]ChainInfo{},
	}
	var (
		tid   uint64
		salt  uint64
		nid   sql.NullInt32
		nonce sql.NullInt64
		gas   sql.NullString
	)
	for rows.Next() {
		err := rows.Scan(&tid, &salt, &nid, &nonce, &gas)
		if err != nil {
			return nil, err
		}
		if nid.Valid {
			ci := ChainInfo{
				NetworkID: uint64(nid.Int32),
			}
			ai.Chain[uint64(nid.Int32)] = ci
			if nonce.Valid {
				ci.Nonce = uint64(nonce.Int64)
			}
			if gas.Valid {
				ci.Gas = hex.DecodeBig(gas.String)
			}
		}
	}
	ai.Uid = tid
	ai.Salt = salt
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ai, nil
}

func (p *PostgresStorage) GetAccountAdr(ctx context.Context, owner string, dbTx pgx.Tx) (*common.Address, error) {
	const getSQL = `SELECT account FROM	omni.user WHERE u.owner = $1 ORDER BY id ASC LIMIT 1;`
	e := p.getExecQuerier(dbTx)
	var adr string
	err := e.QueryRow(ctx, getSQL, owner).Scan(&adr)
	if err != nil {
		return nil, err
	}
	if adr == "" {
		return nil, ErrStorageNotFound
	}
	acc := common.HexToAddress(adr)
	return &acc, nil
}

type Operation struct {
	ID                    uint64    `json:"id"`
	UserID                uint64    `json:"user_id"`
	Status                uint64    `json:"status"`
	Signature             string    `json:"signature"`
	Did                   string    `json:"did"`
	OperationType         uint      `json:"operation_type"`
	OperationValue        string    `json:"operation_value"`
	Phase                 uint      `json:"phase"`
	ExecNonce             uint64    `json:"exec_nonce"`
	ExecNetworkID         uint64    `json:"exec_network_id"`
	ExecCalldata          []byte    `json:"exec_calldata"`
	ExecMainGasLimit      uint64    `json:"exec_main_gas_limit"`
	ExecMainGasPrice      string    `json:"exec_main_gas_price"`
	ExecZkpGasLimit       uint64    `json:"exec_zkp_gas_limit"`
	ExecDestGasLimit      uint64    `json:"exec_dest_gas_limit"`
	ExecDestGasPrice      string    `json:"exec_dest_gas_price"`
	InnerExecNonce        uint64    `json:"inner_exec_nonce"`
	InnerExecNetworkID    uint64    `json:"inner_exec_network_id"`
	InnerExecCalldata     []byte    `json:"inner_exec_calldata"`
	InnerExecMainGasLimit uint64    `json:"inner_exec_main_gas_limit"`
	InnerExecMainGasPrice string    `json:"inner_exec_main_gas_price"`
	InnerExecZkpGasLimit  uint64    `json:"inner_exec_zkp_gas_limit"`
	InnerExecDestGasLimit uint64    `json:"inner_exec_dest_gas_limit"`
	InnerExecDestGasPrice string    `json:"inner_exec_dest_gas_price"`
	CreatedAt             time.Time `json:"create_at"`
	UpdatedAt             time.Time `json:"update_at"`
}

func (p *PostgresStorage) GetUserOps(ctx context.Context, uid uint64, dbTx pgx.Tx) ([]*pool.UserOperation, error) {
	const getSQL = `SELECT * FROM omni.operation o WHERE o.user_id = $1 ORDER BY id ASC;`
	e := p.getExecQuerier(dbTx)
	rows, err := e.Query(ctx, getSQL, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var uos []*pool.UserOperation
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

func scanOperation(row pgx.Row) (*pool.UserOperation, error) {
	op := &pool.UserOperation{}
	var (
		signature string
	)
	err := row.Scan(
		&op.OpId, &op.Uid, &op.Status, &signature, &op.Did, &op.OperationType, &op.OperationValue,
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

func (p *PostgresStorage) GetSigleUserOp(ctx context.Context, owner, account,
	did string, status int, dbTx pgx.Tx) (*pool.UserOperation, error) {
	const getSQL = `
		SELECT * FROM omni.user u 
		LEFT JOIN omni.operation o ON o.user_id = u.id 
		WHERE u.owner = $1 AND u.account = $2 AND o.did = $3 AND status = $4 LIMIT 1;
	`
	e := p.getExecQuerier(dbTx)
	row := e.QueryRow(ctx, getSQL, owner, account, did, status)
	uo, err := scanOperation(row)
	log.Infof("%+v", uo)
	return uo, err
}

func (p *PostgresStorage) AddOperation(ctx context.Context, uid uint64, uo *pool.SignedUserOperation, dbTx pgx.Tx) error {
	e := p.getExecQuerier(dbTx)
	addSQL := `
		INSERT INTO omni.operation (
			user_id, status, signature, did, operation_type, operation_value, phase,
			exec_nonce, exec_network_id, exec_calldata, exec_main_gas_limit, exec_main_gas_price,
			exec_zkp_gas_limit, exec_dest_gas_limit, exec_dest_gas_price,
			inner_exec_nonce, inner_exec_network_id, inner_exec_calldata, inner_exec_main_gas_limit,
			inner_exec_main_gas_price, inner_exec_zkp_gas_limit, inner_exec_dest_gas_limit,
			inner_exec_dest_gas_price, create_at, update_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15,
			$16, $17, $18, $19, $20, $21, $22, $23, $24, $25
		);
	`
	execMainGasPrice := "0"
	if uo.Exec.MainChainGasPrice != nil {
		execMainGasPrice = uo.Exec.MainChainGasPrice.String()
	}
	execDestGasPrice := "0"
	if uo.Exec.DestChainGasPrice != nil {
		execDestGasPrice = uo.Exec.DestChainGasPrice.String()
	}
	innerExecMainGasPrice := "0"
	if uo.InnerExec.MainChainGasPrice != nil {
		innerExecMainGasPrice = uo.InnerExec.MainChainGasPrice.String()
	}
	innerExecDestGasPrice := "0"
	if uo.InnerExec.MainChainGasPrice != nil {
		innerExecDestGasPrice = uo.InnerExec.DestChainGasPrice.String()
	}
	_, err := e.Exec(ctx, addSQL,
		uid, uo.Status, uo.Signature, uo.Did, uo.OperationType, uo.OperationValue.String(),
		uo.Phase, uo.Exec.Nonce, uo.Exec.ChainId, uo.Exec.CallData, uo.Exec.MainChainGasLimit,
		execMainGasPrice, uo.Exec.ZkVerificationGasLimit, uo.Exec.DestChainGasLimit,
		execDestGasPrice, uo.InnerExec.Nonce, uo.InnerExec.ChainId, uo.InnerExec.CallData,
		uo.InnerExec.MainChainGasLimit, innerExecMainGasPrice, uo.InnerExec.ZkVerificationGasLimit,
		uo.InnerExec.DestChainGasLimit, innerExecDestGasPrice, time.Now(), time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresStorage) UpdateUserOp(ctx context.Context, opid uint64, col string, val int, dbTx pgx.Tx) error {
	e := p.getExecQuerier(dbTx)
	updateSQL := "UPDATE omni.operation set $1=$2 where id=$3;"
	_, err := e.Exec(ctx, updateSQL, col, val, opid)
	return err
}
