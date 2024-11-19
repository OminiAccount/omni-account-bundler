package db

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"

	"github.com/OAB/utils/merkletreeutils"
)

type MemDb struct {
	Db          map[string][]string
	DbAccVal    map[string][]string
	DbKeySource map[string][]byte
	DbHashKey   map[string][]byte
	DbCode      map[string][]byte
	LastRoot    *big.Int
	Depth       uint8

	lock sync.RWMutex
}

func NewMemDb() *MemDb {
	return &MemDb{
		Db:          make(map[string][]string),
		DbAccVal:    make(map[string][]string),
		DbKeySource: make(map[string][]byte),
		DbHashKey:   make(map[string][]byte),
		DbCode:      make(map[string][]byte),
		LastRoot:    big.NewInt(0),
		Depth:       0,
	}
}

func (m *MemDb) OpenBatch(quitCh <-chan struct{}) {
}

func (m *MemDb) CommitBatch() error {
	return nil
}

func (m *MemDb) RollbackBatch() {
}

func (m *MemDb) GetLastRoot() (*big.Int, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.LastRoot, nil
}

func (m *MemDb) SetLastRoot(value *big.Int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.LastRoot = value
	return nil
}

func (m *MemDb) GetDepth() (uint8, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.Depth, nil
}

func (m *MemDb) SetDepth(depth uint8) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.Depth = depth
	return nil
}

func (m *MemDb) Get(key merkletreeutils.NodeKey) (merkletreeutils.NodeValue12, error) {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])

	k := merkletreeutils.ConvertBigIntToHex(keyConc)

	values := merkletreeutils.NodeValue12{}
	for i, v := range m.Db[k] {
		values[i] = merkletreeutils.ConvertHexToBigInt(v)
	}

	return values, nil
}

func (m *MemDb) Insert(key merkletreeutils.NodeKey, value merkletreeutils.NodeValue12) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)

	values := make([]string, 12)
	for i, v := range value {
		values[i] = merkletreeutils.ConvertBigIntToHex(v)
	}

	m.Db[k] = values
	return nil
}

func (m *MemDb) Delete(key string) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	delete(m.Db, key)
	return nil
}

func (m *MemDb) DeleteByNodeKey(key merkletreeutils.NodeKey) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)
	delete(m.Db, k)
	return nil
}

func (m *MemDb) GetAccountValue(key merkletreeutils.NodeKey) (merkletreeutils.NodeValue8, error) {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])

	k := merkletreeutils.ConvertBigIntToHex(keyConc)

	values := merkletreeutils.NodeValue8{}
	for i, v := range m.DbAccVal[k] {
		values[i] = merkletreeutils.ConvertHexToBigInt(v)
	}

	return values, nil
}

func (m *MemDb) InsertAccountValue(key merkletreeutils.NodeKey, value merkletreeutils.NodeValue8) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)

	values := make([]string, 8)
	for i, v := range value {
		values[i] = merkletreeutils.ConvertBigIntToHex(v)
	}

	m.DbAccVal[k] = values
	return nil
}

func (m *MemDb) InsertKeySource(key merkletreeutils.NodeKey, value []byte) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])

	m.DbKeySource[keyConc.String()] = value
	return nil
}

func (m *MemDb) DeleteKeySource(key merkletreeutils.NodeKey) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])

	delete(m.DbKeySource, keyConc.String())
	return nil
}

func (m *MemDb) GetKeySource(key merkletreeutils.NodeKey) ([]byte, error) {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])

	s, ok := m.DbKeySource[keyConc.String()]

	if !ok {
		return nil, fmt.Errorf("key not found")
	}

	return s, nil
}

func (m *MemDb) InsertHashKey(key merkletreeutils.NodeKey, value merkletreeutils.NodeKey) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)

	valConc := merkletreeutils.ArrayToScalar(value[:])

	m.DbHashKey[k] = valConc.Bytes()
	return nil
}

func (m *MemDb) DeleteHashKey(key merkletreeutils.NodeKey) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)

	delete(m.DbHashKey, k)
	return nil
}

func (m *MemDb) GetHashKey(key merkletreeutils.NodeKey) (merkletreeutils.NodeKey, error) {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done

	keyConc := merkletreeutils.ArrayToScalar(key[:])
	k := merkletreeutils.ConvertBigIntToHex(keyConc)

	s, ok := m.DbHashKey[k]

	if !ok {
		return merkletreeutils.NodeKey{}, fmt.Errorf("key not found")
	}

	nv := big.NewInt(0).SetBytes(s)

	na := merkletreeutils.ScalarToArray(nv)

	return merkletreeutils.NodeKey{na[0], na[1], na[2], na[3]}, nil
}

func (m *MemDb) GetCode(codeHash []byte) ([]byte, error) {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done

	codeHash = merkletreeutils.ResizeHashTo32BytesByPrefixingWithZeroes(codeHash)

	s, ok := m.DbCode["0x"+hex.EncodeToString(codeHash)]

	if !ok {
		return nil, fmt.Errorf("key not found")
	}

	return s, nil
}

func (m *MemDb) AddCode(code []byte) error {
	m.lock.Lock()         // Lock for writing
	defer m.lock.Unlock() // Make sure to unlock when done

	codeHash := merkletreeutils.HashContractBytecode(hex.EncodeToString(code))
	m.DbCode[codeHash] = code
	return nil
}

func (m *MemDb) IsEmpty() bool {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done

	return len(m.Db) == 0
}

func (m *MemDb) PrintDb() {
	m.lock.RLock()         // Lock for reading
	defer m.lock.RUnlock() // Make sure to unlock when done

	for k, v := range m.Db {
		println(k, v)
	}
}

func (m *MemDb) GetDb() map[string][]string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.Db
}
