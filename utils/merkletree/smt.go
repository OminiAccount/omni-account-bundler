package merkletree

import (
	"github.com/OAB/utils/log"
	"github.com/OAB/utils/merkletreeutils"
	"math/big"

	"context"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"sync"
)

type MemDB interface {
	Insert(key merkletreeutils.NodeKey, value merkletreeutils.NodeValue12) error
	InsertAccountValue(key merkletreeutils.NodeKey, value merkletreeutils.NodeValue8) error
	InsertKeySource(key merkletreeutils.NodeKey, value []byte) error
	DeleteKeySource(key merkletreeutils.NodeKey) error
	InsertHashKey(key merkletreeutils.NodeKey, value merkletreeutils.NodeKey) error
	AddCode(code []byte) error
	DeleteHashKey(key merkletreeutils.NodeKey) error
	Delete(string) error
	DeleteByNodeKey(key merkletreeutils.NodeKey) error
	SetLastRoot(lr *big.Int) error
	SetDepth(uint8) error
	CommitBatch() error
	OpenBatch(quitCh <-chan struct{})
	RollbackBatch()
	RoDB
}

type RoDB interface {
	GetDepth() (uint8, error)
	GetLastRoot() (*big.Int, error)
	GetCode(codeHash []byte) ([]byte, error)
	GetHashKey(key merkletreeutils.NodeKey) (merkletreeutils.NodeKey, error)
	GetKeySource(key merkletreeutils.NodeKey) ([]byte, error)
	Get(key merkletreeutils.NodeKey) (merkletreeutils.NodeValue12, error)
	GetAccountValue(key merkletreeutils.NodeKey) (merkletreeutils.NodeValue8, error)
}

type DebuggableDB interface {
	MemDB
	PrintDb()
	GetDb() map[string][]string
}

type SMT struct {
	noSaveOnInsert bool
	Db             MemDB
	*RoSMT
}

type RoSMT struct {
	DbRo         RoDB
	clearUpMutex sync.Mutex
}

type SMTResponse struct {
	NewRootScalar *merkletreeutils.NodeKey
	Mode          string
}

func NewSMT(database MemDB, noSaveOnInsert bool) *SMT {
	return &SMT{
		noSaveOnInsert: noSaveOnInsert,
		Db:             database,
		RoSMT:          NewRoSMT(database),
	}
}

func NewRoSMT(database RoDB) *RoSMT {
	return &RoSMT{
		DbRo: database,
	}
}

func (s *RoSMT) LastRoot() *big.Int {
	s.clearUpMutex.Lock()
	defer s.clearUpMutex.Unlock()
	lr, err := s.DbRo.GetLastRoot()
	if err != nil {
		panic(err)
	}
	cop := new(big.Int).Set(lr)
	return cop
}

func (s *SMT) SetLastRoot(lr *big.Int) {
	s.clearUpMutex.Lock()
	defer s.clearUpMutex.Unlock()
	err := s.Db.SetLastRoot(lr)
	if err != nil {
		panic(err)
	}
}

func (s *SMT) InsertBI(key *big.Int, value *big.Int) (*SMTResponse, error) {
	k := merkletreeutils.ScalarToNodeKey(key)
	v := merkletreeutils.ScalarToNodeValue8(value)
	return s.insertSingle(k, v, [4]uint64{})
}

func (s *SMT) InsertKA(key merkletreeutils.NodeKey, value *big.Int) (*SMTResponse, error) {
	x := merkletreeutils.ScalarToArrayBig(value)
	v, err := merkletreeutils.NodeValue8FromBigIntArray(x)
	if err != nil {
		return nil, err
	}

	return s.insertSingle(key, *v, [4]uint64{})
}

func (s *SMT) Insert(key merkletreeutils.NodeKey, value merkletreeutils.NodeValue8) (*SMTResponse, error) {
	return s.insertSingle(key, value, [4]uint64{})
}

func (s *SMT) InsertStorage(ethAddr string, storage *map[string]string, chm *map[string]*merkletreeutils.NodeValue8, vhm *map[string][4]uint64, progressChan chan uint64) (*SMTResponse, error) {
	s.clearUpMutex.Lock()
	defer s.clearUpMutex.Unlock()

	a := merkletreeutils.ConvertHexToBigInt(ethAddr)
	add := merkletreeutils.ScalarToArrayBig(a)

	or, err := s.getLastRoot()
	if err != nil {
		return nil, err
	}

	smtr := &SMTResponse{
		NewRootScalar: &or,
	}
	for k := range *storage {
		keyStoragePosition := merkletreeutils.KeyContractStorage(add, k)
		smtr, err = s.insert(keyStoragePosition, *(*chm)[k], (*vhm)[k], *smtr.NewRootScalar)
		if err != nil {
			return nil, err
		}

		sp, _ := merkletreeutils.StrValToBigInt(k)

		ks := merkletreeutils.EncodeKeySource(merkletreeutils.ScStorage, merkletreeutils.ConvertHexToAddress(ethAddr), common.BigToHash(sp))
		err = s.Db.InsertKeySource(keyStoragePosition, ks)

		if err != nil {
			return nil, err
		}

		if progressChan != nil {
			progressChan <- 1
		}
	}

	if err = s.setLastRoot(*smtr.NewRootScalar); err != nil {
		return nil, err
	}

	return smtr, nil
}

func (s *SMT) insertSingle(k merkletreeutils.NodeKey, v merkletreeutils.NodeValue8, newValH [4]uint64) (*SMTResponse, error) {
	s.clearUpMutex.Lock()
	defer s.clearUpMutex.Unlock()

	or, err := s.getLastRoot()
	if err != nil {
		return nil, err
	}

	smtr, err := s.insert(k, v, newValH, or)
	if err != nil {
		return nil, err
	}

	if err = s.setLastRoot(*smtr.NewRootScalar); err != nil {
		return nil, err
	}

	return smtr, nil
}

func (s *SMT) insert(k merkletreeutils.NodeKey, v merkletreeutils.NodeValue8, newValH [4]uint64, oldRoot merkletreeutils.NodeKey) (*SMTResponse, error) {
	newRoot := oldRoot

	smtResponse := &SMTResponse{
		Mode: "not run",
	}

	// split the key
	keys := k.GetPath()

	var usedKey []int
	var level int
	var foundKey *merkletreeutils.NodeKey
	var foundVal merkletreeutils.NodeValue8
	var foundRKey merkletreeutils.NodeKey
	var proofHashCounter int
	var foundOldValHash merkletreeutils.NodeKey

	siblings := map[int]*merkletreeutils.NodeValue12{}

	var err error
	// JS WHILE
	for !oldRoot.IsZero() && foundKey == nil {
		sl, err := s.Db.Get(oldRoot)
		if err != nil {
			return nil, err
		}
		siblings[level] = &sl
		if siblings[level].IsFinalNode() {
			foundOldValHash = merkletreeutils.NodeKeyFromBigIntArray(siblings[level][4:8])
			fva, err := s.Db.Get(foundOldValHash)
			if err != nil {
				return nil, err
			}
			foundValA := merkletreeutils.Value8FromBigIntArray(fva[0:8])
			foundRKey = merkletreeutils.NodeKeyFromBigIntArray(siblings[level][0:4])
			foundVal = foundValA

			foundKey = merkletreeutils.JoinKey(usedKey, foundRKey)
		} else {
			oldRoot = merkletreeutils.NodeKeyFromBigIntArray(siblings[level][keys[level]*4 : keys[level]*4+4])
			usedKey = append(usedKey, keys[level])
			level++
		}
	}

	level--
	if len(usedKey) != 0 {
		usedKey = usedKey[:len(usedKey)-1]
	}

	proofHashCounter = 0
	if !oldRoot.IsZero() {
		//RemoveOver(siblings, level+1)
		proofHashCounter += len(siblings)
		if foundVal.IsZero() {
			proofHashCounter += 2
		}
	}

	if !v.IsZero() { // we have a value - so we're updating or inserting
		if foundKey != nil {
			if foundKey.IsEqualTo(k) {
				// UPDATE MODE
				smtResponse.Mode = "update"

				if newValH == [4]uint64{} {
					newValH, err = s.hashcalcAndSave(v.ToUintArray(), merkletreeutils.BranchCapacity)
				} else {
					err = s.hashSave(v.ToUintArray(), merkletreeutils.BranchCapacity, newValH)
				}
				if err != nil {
					return nil, err
				}

				newLeafHash, err := s.hashcalcAndSave(merkletreeutils.ConcatArrays4(foundRKey, newValH), merkletreeutils.LeafCapacity)
				if err != nil {
					return nil, err
				}
				if err := s.Db.InsertHashKey(newLeafHash, k); err != nil {
					return nil, err
				}
				if level >= 0 {
					for j := 0; j < 4; j++ {
						siblings[level][keys[level]*4+j] = new(big.Int).SetUint64(newLeafHash[j])
					}
				} else {
					newRoot = newLeafHash
				}
			} else {
				smtResponse.Mode = "insertFound"
				// INSERT WITH FOUND KEY
				level2 := level + 1
				foundKeys := foundKey.GetPath()

				for {
					if level2 >= len(keys) || level2 >= len(foundKeys) {
						break
					}

					if keys[level2] != foundKeys[level2] {
						break
					}

					level2++
				}

				oldKey := merkletreeutils.RemoveKeyBits(*foundKey, level2+1)
				oldLeafHash, err := s.hashcalcAndSave(merkletreeutils.ConcatArrays4(oldKey, foundOldValHash), merkletreeutils.LeafCapacity)
				s.Db.InsertHashKey(oldLeafHash, *foundKey)
				if err != nil {
					return nil, err
				}

				newKey := merkletreeutils.RemoveKeyBits(k, level2+1)

				if newValH == [4]uint64{} {
					newValH, err = s.hashcalcAndSave(v.ToUintArray(), merkletreeutils.BranchCapacity)
				} else {
					err = s.hashSave(v.ToUintArray(), merkletreeutils.BranchCapacity, newValH)
				}

				if err != nil {
					return nil, err
				}

				newLeafHash, err := s.hashcalcAndSave(merkletreeutils.ConcatArrays4(newKey, newValH), merkletreeutils.LeafCapacity)
				if err != nil {
					return nil, err
				}

				s.Db.InsertHashKey(newLeafHash, k)

				var node [8]uint64
				for i := 0; i < 8; i++ {
					node[i] = 0
				}

				for j := 0; j < 4; j++ {
					node[keys[level2]*4+j] = newLeafHash[j]
					node[foundKeys[level2]*4+j] = oldLeafHash[j]
				}

				r2, err := s.hashcalcAndSave(node, merkletreeutils.BranchCapacity)
				if err != nil {
					return nil, err
				}
				proofHashCounter += 4
				level2 -= 1

				for level2 != level {
					for i := 0; i < 8; i++ {
						node[i] = 0
					}

					for j := 0; j < 4; j++ {
						node[keys[level2]*4+j] = r2[j]
					}

					r2, err = s.hashcalcAndSave(node, merkletreeutils.BranchCapacity)
					if err != nil {
						return nil, err
					}
					proofHashCounter += 1
					level2 -= 1
				}

				if level >= 0 {
					for j := 0; j < 4; j++ {
						siblings[level][keys[level]*4+j] = new(big.Int).SetUint64(r2[j])
					}
				} else {
					newRoot = r2
				}
			}

		} else {
			// INSERT NOT FOUND
			smtResponse.Mode = "insertNotFound"
			newKey := merkletreeutils.RemoveKeyBits(k, level+1)

			if newValH == [4]uint64{} {
				newValH, err = s.hashcalcAndSave(v.ToUintArray(), merkletreeutils.BranchCapacity)
			} else {
				err = s.hashSave(v.ToUintArray(), merkletreeutils.BranchCapacity, newValH)
			}
			if err != nil {
				return nil, err
			}

			nk := merkletreeutils.ConcatArrays4(newKey, newValH)

			newLeafHash, err := s.hashcalcAndSave(nk, merkletreeutils.LeafCapacity)
			if err != nil {
				return nil, err
			}

			s.Db.InsertHashKey(newLeafHash, k)

			proofHashCounter += 2

			if level >= 0 {
				for j := 0; j < 4; j++ {
					nlh := big.Int{}
					nlh.SetUint64(newLeafHash[j])
					siblings[level][keys[level]*4+j] = &nlh
				}
			} else {
				newRoot = newLeafHash
			}
		}
	} else if foundKey != nil && foundKey.IsEqualTo(k) { // we don't have a value so we're deleting
		if level >= 0 {
			for j := 0; j < 4; j++ {
				siblings[level][keys[level]*4+j] = big.NewInt(0)
			}

			uKey, err := siblings[level].IsUniqueSibling()
			if err != nil {
				return nil, err
			}

			if uKey >= 0 {
				// DELETE FOUND
				smtResponse.Mode = "deleteFound"
				dk := merkletreeutils.NodeKeyFromBigIntArray(siblings[level][uKey*4 : uKey*4+4])
				sl, err := s.Db.Get(dk)
				if err != nil {
					return nil, err
				}
				siblings[level+1] = &sl

				if siblings[level+1].IsFinalNode() {
					valH := siblings[level+1].Get4to8()

					rKey := siblings[level+1].Get0to4()
					proofHashCounter += 2

					insKey := merkletreeutils.JoinKey(append(usedKey, uKey), *rKey)

					for uKey >= 0 && level >= 0 {
						level -= 1
						if level >= 0 {
							uKey, err = siblings[level].IsUniqueSibling()
							if err != nil {
								return nil, err
							}
						}
					}

					oldKey := merkletreeutils.RemoveKeyBits(*insKey, level+1)
					oldLeafHash, err := s.hashcalcAndSave(merkletreeutils.ConcatArrays4(oldKey, *valH), merkletreeutils.LeafCapacity)
					s.Db.InsertHashKey(oldLeafHash, *insKey)
					if err != nil {
						return nil, err
					}
					proofHashCounter += 1

					if level >= 0 {
						for j := 0; j < 4; j++ {
							siblings[level][keys[level]*4+j] = new(big.Int).SetUint64(oldLeafHash[j])
						}
					} else {
						newRoot = oldLeafHash
					}
				} else {
					// DELETE NOT FOUND
					smtResponse.Mode = "deleteNotFound"
				}
			} else {
				// DELETE NOT FOUND
				smtResponse.Mode = "deleteNotFound"
			}
		} else {
			// DELETE LAST
			smtResponse.Mode = "deleteLast"
			newRoot = merkletreeutils.NodeKey{0, 0, 0, 0}
		}
	} else { // we're going zero to zero - do nothing
		smtResponse.Mode = "zeroToZero"
	}

	merkletreeutils.RemoveOver(siblings, level+1)

	s.updateDepth(len(siblings))

	for level >= 0 {
		hashValueIn, err := merkletreeutils.NodeValue8FromBigIntArray(siblings[level][0:8])
		if err != nil {
			return nil, err
		}
		hashCapIn := merkletreeutils.NodeKeyFromBigIntArray(siblings[level][8:12])
		newRoot, err = s.hashcalcAndSave(hashValueIn.ToUintArray(), hashCapIn)
		if err != nil {
			return nil, err
		}
		proofHashCounter += 1
		level -= 1
		if level >= 0 {
			for j := 0; j < 4; j++ {
				nrj := big.Int{}
				nrj.SetUint64(newRoot[j])
				siblings[level][keys[level]*4+j] = &nrj
			}
		}
	}

	_ = oldRoot

	smtResponse.NewRootScalar = &newRoot

	return smtResponse, nil
}

// used only by old smt (not by smt batch/create)
func (s *SMT) hashSave(in [8]uint64, capacity, h [4]uint64) error {
	if s.noSaveOnInsert {
		return nil
	}
	v := merkletreeutils.ConcatArrays8AndCapacityByPointers(&in, &capacity)

	return s.Db.Insert(h, *v)
}

func (s *SMT) hashSaveByPointers(in *[8]uint64, capacity, h *[4]uint64) error {
	if s.noSaveOnInsert {
		return nil
	}
	v := merkletreeutils.ConcatArrays8AndCapacityByPointers(in, capacity)

	return s.Db.Insert(*h, *v)
}

// used only by old smt (not by smt batch/create)
func (s *SMT) hashcalcAndSave(in [8]uint64, capacity [4]uint64) ([4]uint64, error) {
	h := merkletreeutils.Hash(in, capacity)
	return h, s.hashSave(in, capacity, h)
}

func (s *RoSMT) getLastRoot() (merkletreeutils.NodeKey, error) {
	or, err := s.DbRo.GetLastRoot()
	if err != nil {
		return merkletreeutils.NodeKey{}, err
	}
	return merkletreeutils.ScalarToRoot(or), nil
}

func (s *SMT) setLastRoot(newRoot merkletreeutils.NodeKey) error {
	return s.Db.SetLastRoot(newRoot.ToBigInt())
}

// Utility functions for debugging

func (s *RoSMT) PrintDb() {
	if debugDB, ok := s.DbRo.(DebuggableDB); ok {
		debugDB.PrintDb()
	}
}

func (s *RoSMT) PrintTree() {
	if debugDB, ok := s.DbRo.(DebuggableDB); ok {
		data := debugDB.GetDb()
		str, err := json.Marshal(data)
		if err != nil {
			log.Error(err)
		}
		log.Info(string(str))
	}
}

type VisitedNodesMap map[string]bool

func (s *SMT) updateDepth(newDepth int) {
	oldDepth, err := s.Db.GetDepth()
	if err != nil {
		oldDepth = 0
	}

	// if new depth is 255 we should be adding an addition +1 because the tree itself cannot have more than 255 levels [0, 255] (counting starts at 0)
	if newDepth > 255 {
		newDepth = 255
	}

	newDepthAsByte := byte(newDepth & 0xFF)
	if oldDepth < newDepthAsByte {
		_ = s.Db.SetDepth(newDepthAsByte)
	}
}

/*
depths are 0 based
0 means either only root leaf or empty tree
*/
func (s *RoSMT) GetDepth() int {
	depth, err := s.DbRo.GetDepth()
	if err != nil {
		return 0
	}
	return int(depth)
}

type TraverseAction func(prefix []byte, k merkletreeutils.NodeKey, v merkletreeutils.NodeValue12) (bool, error)

func (s *RoSMT) Traverse(ctx context.Context, node *big.Int, action TraverseAction) error {
	return s.traverse(ctx, node, action, []byte{})
}

func (s *RoSMT) traverse(ctx context.Context, node *big.Int, action TraverseAction, prefix []byte) error {
	if node == nil || node.Cmp(big.NewInt(0)) == 0 {
		return nil
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	ky := merkletreeutils.ScalarToRoot(node)

	nodeValue, err := s.DbRo.Get(ky)

	if err != nil {
		return err
	}

	shouldContinue, err := action(prefix, ky, nodeValue)

	if err != nil {
		return err
	}

	if nodeValue.IsFinalNode() || !shouldContinue {
		return nil
	}

	for i := 0; i < 2; i++ {
		if len(nodeValue) < i*4+4 {
			return errors.New("nodeValue has insufficient length")
		}
		child := merkletreeutils.NodeKeyFromBigIntArray(nodeValue[i*4 : i*4+4])
		childPrefix := make([]byte, len(prefix)+1)
		copy(childPrefix, prefix)
		childPrefix[len(prefix)] = byte(i)
		err = s.traverse(ctx, child.ToBigInt(), action, childPrefix)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

func (s *RoSMT) traverseAndMark(ctx context.Context, node *big.Int, visited VisitedNodesMap) error {
	return s.Traverse(ctx, node, func(prefix []byte, k merkletreeutils.NodeKey, v merkletreeutils.NodeValue12) (bool, error) {
		if visited[merkletreeutils.ConvertBigIntToHex(k.ToBigInt())] {
			return false, nil
		}

		visited[merkletreeutils.ConvertBigIntToHex(k.ToBigInt())] = true
		return true, nil
	})
}

// InsertHashNode inserts a hash node into the SMT. The SMT should not contain any other leaf nodes with the same path prefix. Otherwise, the new root hash will be incorrect.
// TODO: Support insertion of hash nodes even if there are leaf nodes with the same path prefix in SMT.
func (s *SMT) InsertHashNode(path []int, hash *big.Int) (*big.Int, error) {
	s.clearUpMutex.Lock()
	defer s.clearUpMutex.Unlock()

	or, err := s.getLastRoot()
	if err != nil {
		return nil, err
	}

	h := merkletreeutils.ScalarToArray(hash)

	var nodeHash [4]uint64
	copy(nodeHash[:], h[:4])

	lastRoot, err := s.insertHashNode(path, nodeHash, or)
	if err != nil {
		return nil, err
	}

	if err = s.setLastRoot(lastRoot); err != nil {
		return nil, err
	}

	return lastRoot.ToBigInt(), nil
}

func (s *SMT) insertHashNode(path []int, hash [4]uint64, root merkletreeutils.NodeKey) (merkletreeutils.NodeKey, error) {
	if len(path) == 0 {
		newValHBig := merkletreeutils.ArrayToScalar(hash[:])
		v := merkletreeutils.ScalarToNodeValue8(newValHBig)

		err := s.hashSave(v.ToUintArray(), merkletreeutils.LeafCapacity, hash)
		if err != nil {
			return merkletreeutils.NodeKey{}, err
		}

		return hash, nil
	}

	rootVal := merkletreeutils.NodeValue12{}

	if !root.IsZero() {
		v, err := s.Db.Get(root)
		if err != nil {
			return merkletreeutils.NodeKey{}, err
		}

		rootVal = v
	}

	childIndex := path[0]

	childOldRoot := rootVal[childIndex*4 : childIndex*4+4]

	childNewRoot, err := s.insertHashNode(path[1:], hash, merkletreeutils.NodeKeyFromBigIntArray(childOldRoot))

	if err != nil {
		return merkletreeutils.NodeKey{}, err
	}

	var newIn [8]uint64

	emptyRootVal := merkletreeutils.NodeValue12{}

	if childIndex == 0 {
		var sibling [4]uint64
		if rootVal == emptyRootVal {
			sibling = [4]uint64{0, 0, 0, 0}
		} else {
			sibling = *rootVal.Get4to8()
		}
		newIn = merkletreeutils.ConcatArrays4(childNewRoot, sibling)
	} else {
		var sibling [4]uint64
		if rootVal == emptyRootVal {
			sibling = [4]uint64{0, 0, 0, 0}
		} else {
			sibling = *rootVal.Get0to4()
		}
		newIn = merkletreeutils.ConcatArrays4(sibling, childNewRoot)
	}

	return s.hashcalcAndSave(newIn, merkletreeutils.BranchCapacity)
}
