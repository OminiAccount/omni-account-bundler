package hashdb

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/OAB/utils/hex"
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/merkletree"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"math/big"
	"sync"

	"github.com/OAB/utils/merkletreeutils"
)

type HashDb struct {
	ctx         context.Context
	Db          map[string][]string
	DbAccVal    map[string][]string
	DbKeySource map[string][]byte
	DbHashKey   map[string][]byte
	DbCode      map[string][]byte
	LastRoot    *big.Int
	Depth       uint8

	lock sync.RWMutex
	db   *PostgresStorage
}

func NewHashDb(c context.Context, p *pgxpool.Pool) merkletree.MemDB {
	return &HashDb{
		Db:          make(map[string][]string),
		DbAccVal:    make(map[string][]string),
		DbKeySource: make(map[string][]byte),
		DbHashKey:   make(map[string][]byte),
		DbCode:      make(map[string][]byte),
		db:          NewPostgresStorage(p),
		ctx:         c,
		//LastRoot:    big.NewInt(0),
		//Depth: 0,
	}
}

func (m *HashDb) OpenBatch(quitCh <-chan struct{}) {
}

func (m *HashDb) CommitBatch() error {
	return nil
}

func (m *HashDb) RollbackBatch() {
}

func (m *HashDb) GetLastRoot() (*big.Int, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if m.LastRoot == nil {
		hi, err := m.db.getHashInfo(m.ctx, nil)
		if errors.Is(err, pgx.ErrNoRows) {
			return big.NewInt(0), nil
		}
		if err != nil {
			return nil, err
		}
		m.LastRoot = hex.DecodeBig(hi.LastRoot)
		m.Depth = hi.Depth
	}
	return m.LastRoot, nil
}

func (m *HashDb) SetLastRoot(value *big.Int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	err := m.db.updateHashInfoRoot(m.ctx, hex.EncodeBig(value), nil)
	if err != nil {
		return err
	}
	m.LastRoot = value
	return nil
}

func (m *HashDb) GetDepth() (uint8, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if m.Depth < 1 {
		hi, err := m.db.getHashInfo(m.ctx, nil)
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		if err != nil {
			return 0, err
		}
		m.LastRoot = hex.DecodeBig(hi.LastRoot)
		m.Depth = hi.Depth
	}
	return m.Depth, nil
}

func (m *HashDb) SetDepth(depth uint8) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	err := m.db.updateHashInfoDepth(m.ctx, depth, nil)
	if err != nil {
		return err
	}
	m.Depth = depth
	return nil
}

func (m *HashDb) Get(key merkletreeutils.NodeKey) (merkletreeutils.NodeValue12, error) {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)
	values := merkletreeutils.NodeValue12{}
	if _, ok := m.Db[k]; !ok {
		d, err := m.db.getHashDB(m.ctx, k, nil)
		if err != nil {
			return values, err
		}
		m.Db[k] = d
	}
	for i, v := range m.Db[k] {
		values[i] = merkletreeutils.ConvertHexToBigInt(v)
	}

	return values, nil
}

func (m *HashDb) Insert(key merkletreeutils.NodeKey, value merkletreeutils.NodeValue12) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)

	values := make([]string, 12)
	for i, v := range value {
		values[i] = merkletreeutils.ConvertBigIntToHex(v)
	}
	dbyte, err := json.Marshal(values)
	if err != nil {
		return err
	}
	err = m.db.insertHashDB(m.ctx, k, string(dbyte), nil)
	if err != nil {
		return err
	}
	m.Db[k] = values
	return nil
}

func (m *HashDb) Delete(key string) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done
	err := m.db.deleteHashDB(m.ctx, key, nil)
	if err != nil {
		return err
	}
	delete(m.Db, key)
	return nil
}

func (m *HashDb) DeleteByNodeKey(key merkletreeutils.NodeKey) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)
	err := m.db.deleteHashDB(m.ctx, k, nil)
	if err != nil {
		return err
	}
	delete(m.Db, k)
	return nil
}

func (m *HashDb) GetAccountValue(key merkletreeutils.NodeKey) (merkletreeutils.NodeValue8, error) {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)
	values := merkletreeutils.NodeValue8{}
	if _, ok := m.DbAccVal[k]; !ok {
		d, err := m.db.getHashAccVal(m.ctx, k, nil)
		if err != nil {
			return values, err
		}
		m.DbAccVal[k] = d
	}
	for i, v := range m.DbAccVal[k] {
		values[i] = merkletreeutils.ConvertHexToBigInt(v)
	}

	return values, nil
}

func (m *HashDb) InsertAccountValue(key merkletreeutils.NodeKey, value merkletreeutils.NodeValue8) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)

	values := make([]string, 8)
	for i, v := range value {
		values[i] = merkletreeutils.ConvertBigIntToHex(v)
	}
	dbyte, err := json.Marshal(values)
	if err != nil {
		return err
	}
	err = m.db.insertHashAccVal(m.ctx, k, string(dbyte), nil)
	if err != nil {
		return err
	}

	m.DbAccVal[k] = values
	return nil
}

func (m *HashDb) InsertKeySource(key merkletreeutils.NodeKey, value []byte) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := keyConc.String()
	err := m.db.insertHashKeySource(m.ctx, k, value, nil)
	if err != nil {
		return err
	}
	m.DbKeySource[k] = value
	return nil
}

func (m *HashDb) DeleteKeySource(key merkletreeutils.NodeKey) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := keyConc.String()
	err := m.db.deleteHashKeySource(m.ctx, k, nil)
	if err != nil {
		return err
	}
	delete(m.DbKeySource, k)
	return nil
}

func (m *HashDb) GetKeySource(key merkletreeutils.NodeKey) ([]byte, error) {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := keyConc.String()

	if _, ok := m.DbKeySource[k]; !ok {
		s, err := m.db.getHashKeySource(m.ctx, k, nil)
		if err != nil {
			return nil, err
		}
		m.DbKeySource[k] = s
	}
	return m.DbKeySource[k], nil
}

func (m *HashDb) InsertHashKey(key merkletreeutils.NodeKey, value merkletreeutils.NodeKey) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)

	valConc := merkletreeutils.ArrayToScalar(value[:])
	err := m.db.insertHashHashKey(m.ctx, k, valConc.Bytes(), nil)
	if err != nil {
		return err
	}
	m.DbHashKey[k] = valConc.Bytes()
	return nil
}

func (m *HashDb) DeleteHashKey(key merkletreeutils.NodeKey) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)
	err := m.db.deleteHashHashKey(m.ctx, k, nil)
	if err != nil {
		return err
	}
	delete(m.DbHashKey, k)
	return nil
}

func (m *HashDb) GetHashKey(key merkletreeutils.NodeKey) (merkletreeutils.NodeKey, error) {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)

	if _, ok := m.DbHashKey[k]; !ok {
		s, err := m.db.getHashHashKey(m.ctx, k, nil)
		if err != nil {
			return merkletreeutils.NodeKey{}, err
		}
		m.DbHashKey[k] = s
	}

	nv := big.NewInt(0).SetBytes(m.DbHashKey[k])
	na := merkletreeutils.ScalarToArray(nv)
	return merkletreeutils.NodeKey{na[0], na[1], na[2], na[3]}, nil
}

func (m *HashDb) GetCode(codeHash []byte) ([]byte, error) {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done

	codeHash = merkletreeutils.ResizeHashTo32BytesByPrefixingWithZeroes(codeHash)
	k := "0x" + hex.EncodeToString(codeHash)
	if _, ok := m.DbCode[k]; !ok {
		s, err := m.db.getHashCode(m.ctx, k, nil)
		if err != nil {
			return nil, err
		}
		m.DbCode[k] = s
	}
	return m.DbCode[k], nil
}

func (m *HashDb) AddCode(code []byte) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	codeHash := merkletreeutils.HashContractBytecode(hex.EncodeToString(code))
	err := m.db.insertHashCode(m.ctx, codeHash, code, nil)
	if err != nil {
		return err
	}
	m.DbCode[codeHash] = code
	return nil
}

func (m *HashDb) IsEmpty() bool {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done
	n, err := m.db.getHashDBCount(m.ctx, nil)
	if err != nil {
		panic(err)
	}
	return n == 0
}

func (m *HashDb) PrintDb() {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done

	d, err := m.db.getHashDBAllData(m.ctx, nil)
	if err != nil {
		log.Errorf("PrintDb -> getHashDBAllData err: %+v", err)
		return
	}
	for _, v := range d {
		println(v.Hash, v.Data)
	}
}

func (m *HashDb) GetDb() map[string][]string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.Db
}
