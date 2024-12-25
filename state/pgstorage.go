package state

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OAB/database/pgstorage"
	"github.com/OAB/lib/common/hexutil"
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

func (p *PostgresStorage) GetUserInfoLock(ctx context.Context, uid, nid uint64, dbTx pgx.Tx) (uint64, *big.Int, error) {
	e := p.getExecQuerier(dbTx)
	getSQL := `SELECT nonce, gas FROM omni.user_info WHERE user_id = $1 AND network_id = $2 FOR UPDATE;`
	var tNonce uint64
	var tGas string
	err := e.QueryRow(ctx, getSQL, uid, nid).Scan(&tNonce, &tGas)
	if err != nil {
		return 0, nil, err
	}
	tGasInt := hex.DecodeBig(tGas)
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
		setSQL = "set " + setSQL + "update_at=$" + strconv.Itoa(len(where)+1)
		where = append(where, time.Now())
	} else {
		return nil
	}
	updateSQL += setSQL + " where user_id=$" + strconv.Itoa(len(where)+1) + " and network_id=$" + strconv.Itoa(len(where)+2) + ";"
	where = append(where, uid, nid)
	_, err := e.Exec(ctx, updateSQL, where...)
	return err
}

func (p *PostgresStorage) AddUser(ctx context.Context, am AccountMapping, dbTx pgx.Tx) error {
	addSQL := "INSERT INTO omni.user (owner, account, salt, failed_salt, create_at, update_at) " +
		"VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (owner) DO NOTHING;"
	e := p.getExecQuerier(dbTx)
	_, err := e.Exec(ctx, addSQL, strings.ToLower(am.User.Hex()),
		strings.ToLower(am.Account.Hex()), am.Salt, "", time.Now(), time.Now())
	return err
}

func (p *PostgresStorage) AddUserInfo(ctx context.Context, uid, nid uint64, dbTx pgx.Tx) error {
	addSQL := "INSERT INTO omni.user_info (user_id, network_id, nonce, gas, create_at, update_at) VALUES ($1, $2, $3, $4, $5, $6);"
	e := p.getExecQuerier(dbTx)
	_, err := e.Exec(ctx, addSQL, uid, nid, 0, "0x0", time.Now(), time.Now())
	return err
}

func (p *PostgresStorage) GetUserInfoByAA(ctx context.Context, owner, account string, dbTx pgx.Tx) (*AccountInfo, error) {
	getSQL := `
		SELECT u.id,u.owner,u.account,u.salt,ui.network_id,ui.nonce,ui.gas
		FROM omni.user u LEFT JOIN omni.user_info ui ON u.id = ui.user_id
		WHERE u.owner = $1 AND u.account = $2;
	`
	e := p.getExecQuerier(dbTx)
	rows, err := e.Query(ctx, getSQL, strings.ToLower(owner), strings.ToLower(account))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		ai, err := scanUser(rows)
		if err != nil {
			return nil, err
		}
		return ai, nil
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	//log.Infof("%+v", ai)
	return nil, pgx.ErrNoRows
}

func scanUser(row pgx.Row) (*AccountInfo, error) {
	ai := &AccountInfo{Chain: map[uint64]ChainInfo{}}
	var tid, salt uint64
	var ownerStr, accStr string
	var (
		nid   sql.NullInt32
		nonce sql.NullInt64
		gas   sql.NullString
	)
	err := row.Scan(&tid, &ownerStr, &accStr, &salt, &nid, &nonce, &gas)
	if err != nil {
		return nil, err
	}
	if nid.Valid {
		ci := ChainInfo{
			NetworkID: uint64(nid.Int32),
		}
		if nonce.Valid {
			ci.Nonce = uint64(nonce.Int64)
		}
		if gas.Valid {
			ci.Gas = hex.DecodeBig(gas.String)
		}
		ai.Chain[uint64(nid.Int32)] = ci
	}
	ai.User = common.HexToAddress(ownerStr)
	ai.Account = common.HexToAddress(accStr)
	ai.Uid = tid
	ai.Salt = salt
	return ai, nil
}

func (p *PostgresStorage) GetUserInfoByChain(ctx context.Context, owner string, nid uint64, dbTx pgx.Tx) (*AccountInfo, error) {
	getSQL := `
		SELECT u.id,u.owner,u.account,u.salt,ui.network_id,ui.nonce,ui.gas
		FROM omni.user u LEFT JOIN omni.user_info ui ON u.id = ui.user_id
		WHERE u.owner = $1 AND ui.network_id = $2;
	`
	e := p.getExecQuerier(dbTx)
	rows, err := e.Query(ctx, getSQL, strings.ToLower(owner), nid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		ai, err := scanUser(rows)
		if err != nil {
			return nil, err
		}
		return ai, nil
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	//log.Infof("%+v", ai)
	return nil, nil
}

func (p *PostgresStorage) GetUser(ctx context.Context, owner string, dbTx pgx.Tx) (*AccountInfo, error) {
	const getSQL = `SELECT id,owner,account,salt FROM omni.user WHERE owner = $1 ORDER BY id ASC LIMIT 1;`
	e := p.getExecQuerier(dbTx)
	var ai AccountInfo
	var ownerStr, accStr string
	err := e.QueryRow(ctx, getSQL, strings.ToLower(owner)).Scan(&ai.Uid, &ownerStr, &accStr, &ai.Salt)
	if err != nil {
		return nil, err
	}
	if ai.Uid < 1 {
		return nil, ErrStorageNotFound
	}
	ai.User = common.HexToAddress(ownerStr)
	ai.Account = common.HexToAddress(accStr)
	return &ai, nil
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
	var opValue, signature string
	var execMainGasPrice, execDestGasPrice, iexecMainGasPrice, iexecDestGasPrice string
	var updateAt time.Time
	err := row.Scan(
		&op.OpId, &op.Uid, &op.BatchNum, &op.Status, &signature, &op.Did, &op.OperationType, &opValue,
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
	log.Infof("%+v", op)
	return op, nil
}

func (p *PostgresStorage) GetSigleUserOp(ctx context.Context, owner, account,
	did string, status int, dbTx pgx.Tx) (*UserOperation, error) {
	const getSQL = `
		SELECT o.* FROM omni.user u 
		LEFT JOIN omni.operation o ON o.user_id = u.id 
		WHERE u.owner = $1 AND u.account = $2 AND o.did = $3 AND o.status = $4 LIMIT 1;
	`
	e := p.getExecQuerier(dbTx)
	row := e.QueryRow(ctx, getSQL, strings.ToLower(owner), strings.ToLower(account), did, status)
	uo, err := scanOperation(row)
	log.Infof("%+v", uo)
	return uo, err
}

func (p *PostgresStorage) AddOperation(ctx context.Context, uid uint64, uo *SignedUserOperation, dbTx pgx.Tx) error {
	e := p.getExecQuerier(dbTx)
	addSQL := `
		INSERT INTO omni.operation (
			user_id, batch_num, status, signature, did, operation_type, operation_value, phase,
			exec_nonce, exec_network_id, exec_calldata, exec_main_gas_limit, exec_main_gas_price,
			exec_zkp_gas_limit, exec_dest_gas_limit, exec_dest_gas_price,
			inner_exec_nonce, inner_exec_network_id, inner_exec_calldata, inner_exec_main_gas_limit,
			inner_exec_main_gas_price, inner_exec_zkp_gas_limit, inner_exec_dest_gas_limit,
			inner_exec_dest_gas_price, create_at, update_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15,
			$16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26
		);
	`
	execMainGasPrice := "0x0"
	if uo.Exec.MainChainGasPrice != nil {
		execMainGasPrice = uo.Exec.MainChainGasPrice.String()
	}
	execDestGasPrice := "0x0"
	if uo.Exec.DestChainGasPrice != nil {
		execDestGasPrice = uo.Exec.DestChainGasPrice.String()
	}
	innerExecMainGasPrice := "0x0"
	if uo.InnerExec.MainChainGasPrice != nil {
		innerExecMainGasPrice = uo.InnerExec.MainChainGasPrice.String()
	}
	innerExecDestGasPrice := "0x0"
	if uo.InnerExec.MainChainGasPrice != nil {
		innerExecDestGasPrice = uo.InnerExec.DestChainGasPrice.String()
	}
	_, err := e.Exec(ctx, addSQL,
		uid, 0, uo.Status, uo.Signature, uo.Did, uo.OperationType, uo.OperationValue.String(),
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
	updateSQL := "UPDATE omni.operation set %s=$1,update_at=$2 where id=$3;"
	updateSQL = fmt.Sprintf(updateSQL, col)
	_, err := e.Exec(ctx, updateSQL, val, time.Now(), opid)
	return err
}

func (p *PostgresStorage) AddUserFailedSalt(ctx context.Context, uid, nid uint64, dbTx pgx.Tx) error {
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
	getSQL := `SELECT failed_salt FROM omni.user where id = $1 FOR UPDATE;`
	e := p.getExecQuerier(dbTx)
	err = e.QueryRow(ctx, getSQL, uid).Scan(&failedSalt)
	if err != nil {
		if flag {
			_ = dbTx.Rollback(ctx)
		}
		return err
	}
	if len(failedSalt) > 0 {
		if strings.Contains(failedSalt, fmt.Sprintf("%d", nid)) {
			if flag {
				_ = dbTx.Rollback(ctx)
			}
			return nil
		}
		failedSalt += "," + fmt.Sprintf("%d", nid)
	} else {
		failedSalt = fmt.Sprintf("%d", nid)
	}
	updateSQL := `UPDATE omni.user SET failed_salt=$1,update_at=$2 where id = $3;`
	_, err = e.Exec(ctx, updateSQL, failedSalt, time.Now(), uid)
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

func (p *PostgresStorage) GetHistoryByDid(ctx context.Context, uid uint64, did string, dbTx pgx.Tx) (AccountHistory, error) {
	const getSQL = `
		SELECT 
		id,order_type,from_info,to_info,source_chain,target_chain,source_hash,target_hash,swap_hash,status,create_at
		FROM omni.history 
		WHERE user_id=$1 AND did=$2 LIMIT 1;
	`
	e := p.getExecQuerier(dbTx)
	his := AccountHistory{
		Did: did,
		Uid: uid,
	}
	var fromStr, toStr string
	err := e.QueryRow(ctx, getSQL, uid, did).Scan(
		&his.ID, &his.OrderType, &fromStr, &toStr, &his.SourceChain, &his.TargetChain, &his.SourceHash, &his.TargetHash,
		&his.SwapHash, &his.Status, &his.TimeAt,
	)
	if err != nil {
		return his, err
	}
	err = json.Unmarshal([]byte(fromStr), &his.From)
	if err != nil {
		return his, err
	}
	err = json.Unmarshal([]byte(toStr), &his.To)
	if err != nil {
		return his, err
	}
	return his, nil
}

func (p *PostgresStorage) AddUserHis(ctx context.Context, d *AccountHistory, dbTx pgx.Tx) error {
	addSQL := `
		INSERT INTO omni.history 
		(did,user_id,order_type,from_info,to_info,source_chain,target_chain,source_hash,target_hash,swap_hash,status,create_at,update_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);
	`
	e := p.getExecQuerier(dbTx)
	fromB, err := json.Marshal(d.From)
	if err != nil {
		return err
	}
	toB, err := json.Marshal(d.To)
	if err != nil {
		return err
	}
	_, err = e.Exec(ctx, addSQL, d.Did, d.Uid, d.OrderType, string(fromB), string(toB), d.SourceChain,
		d.TargetChain, d.SourceHash, d.TargetHash, d.SwapHash, d.Status, d.TimeAt, d.TimeAt)
	return err
}

func (p *PostgresStorage) GetHistoryByStatus(ctx context.Context, status []int, l uint64, dbTx pgx.Tx) ([]AccountHistory, error) {
	const getSQL = `
		SELECT 
		id,did,user_id,order_type,from_info,to_info,source_chain,target_chain,source_hash,target_hash,swap_hash,status,create_at
		FROM omni.history 
		WHERE status = ANY ($1) ORDER BY id ASC LIMIT $2;
	`
	e := p.getExecQuerier(dbTx)
	var list []AccountHistory
	rows, err := e.Query(ctx, getSQL, status, l)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		his, err := p.scanHisItem(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, his)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func (p *PostgresStorage) GetHistoryByTs(ctx context.Context, uid, l uint64, ts time.Time, dbTx pgx.Tx) ([]AccountHistory, error) {
	const getSQL = `
		SELECT 
		id,did,user_id,order_type,from_info,to_info,source_chain,target_chain,source_hash,target_hash,swap_hash,status,create_at
		FROM omni.history 
		WHERE user_id = $1 AND create_at < $2 ORDER BY id DESC LIMIT $3;
	`
	e := p.getExecQuerier(dbTx)
	var list []AccountHistory
	rows, err := e.Query(ctx, getSQL, uid, ts, l)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		his, err := p.scanHisItem(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, his)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func (p *PostgresStorage) scanHisItem(row pgx.Rows) (AccountHistory, error) {
	var fromStr, toStr string
	var his AccountHistory
	err := row.Scan(
		&his.ID, &his.Did, &his.Uid, &his.OrderType, &fromStr, &toStr, &his.SourceChain, &his.TargetChain,
		&his.SourceHash, &his.TargetHash, &his.SwapHash, &his.Status, &his.TimeAt,
	)
	if err != nil {
		return his, err
	}
	err = json.Unmarshal([]byte(fromStr), &his.From)
	if err != nil {
		return his, err
	}
	err = json.Unmarshal([]byte(toStr), &his.To)
	if err != nil {
		return his, err
	}
	return his, nil
}

func (p *PostgresStorage) UpdateHistory(ctx context.Context, id uint64, m map[string]interface{}, dbTx pgx.Tx) error {
	e := p.getExecQuerier(dbTx)
	updateSQL := "UPDATE omni.history set "
	var where []interface{}
	for col, val := range m {
		updateSQL += col + "=$" + strconv.Itoa(len(where)+1) + ","
		where = append(where, val)
	}
	updateSQL += "update_at=$" + strconv.Itoa(len(where)+1) + " where id=$" + strconv.Itoa(len(where)+2) + ";"
	where = append(where, time.Now(), id)
	_, err := e.Exec(ctx, updateSQL, where...)
	return err
}
