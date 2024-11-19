package merkletree

import (
	"bytes"
	"context"
	"github.com/OAB/utils/merkletreeutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/holiman/uint256"
)

// GetAccountBalance returns the balance of an account from the SMT
func (s *SMT) GetAccountBalance(address common.Address) (*uint256.Int, error) {
	if s.IsEmptyAccountBalance(address) {
		return uint256.NewInt(0), nil
	}
	valueInBytes, err := s.getValue(merkletreeutils.KeyBalance, address, nil)
	if err != nil {
		log.Error("failed to get balance", "error", err)
		return nil, err
	}

	balance := uint256.NewInt(0).SetBytes(valueInBytes)

	return balance, nil
}

func (s *SMT) IsEmptyAccountBalance(address common.Address) bool {
	keyBalance := merkletreeutils.KeyEthAddrBalance(address.String())
	source, _ := s.Db.GetKeySource(keyBalance)
	return len(source) == 0
}

// GetAccountNonce returns the nonce of an account from the SMT
func (s *SMT) GetAccountNonce(address common.Address, nonceChainId string) (*uint256.Int, error) {
	if s.IsEmptyAccountNonce(address, nonceChainId) {
		return uint256.NewInt(0), nil
	}
	chainId := common.HexToHash(nonceChainId)
	valueInBytes, err := s.getValue(10, address, &chainId)
	if err != nil {
		log.Error("failed to get nonce", "error", err)
		return nil, err
	}

	nonce := uint256.NewInt(0).SetBytes(valueInBytes)

	return nonce, nil
}

func (s *SMT) IsEmptyAccountNonce(address common.Address, nonceChainId string) bool {
	keyNonce := merkletreeutils.KeyEthAddrNonceChainId(address.String(), nonceChainId)
	source, _ := s.Db.GetKeySource(keyNonce)
	return len(source) == 0
}

// getValue returns the value of a key from SMT by traversing the SMT
func (s *SMT) getValue(key int, address common.Address, storageKey *common.Hash) ([]byte, error) {
	var kn merkletreeutils.NodeKey

	if storageKey == nil {
		kn = merkletreeutils.Key(address.String(), key)
	} else {
		//a := merkletreeutils.ConvertHexToBigInt(address.String())
		//add := merkletreeutils.ScalarToArrayBig(a)
		kn = merkletreeutils.KeyEthAddrNonceChainId(address.String(), storageKey.String())
	}

	return s.getValueInBytes(kn)
}

// getValueInBytes returns the value of a key from SMT in bytes by traversing the SMT
func (s *SMT) getValueInBytes(nodeKey merkletreeutils.NodeKey) ([]byte, error) {
	var value []byte

	keyPath := nodeKey.GetPath()

	keyPathBytes := make([]byte, len(keyPath))
	for i, k := range keyPath {
		keyPathBytes[i] = byte(k)
	}

	action := func(prefix []byte, _ merkletreeutils.NodeKey, v merkletreeutils.NodeValue12) (bool, error) {
		if !bytes.HasPrefix(keyPathBytes, prefix) {
			return false, nil
		}

		if v.IsFinalNode() {
			valHash := v.Get4to8()
			v, err := s.Db.Get(*valHash)
			if err != nil {
				return false, err
			}
			vInBytes := merkletreeutils.ArrayBigToScalar(merkletreeutils.BigIntArrayFromNodeValue8(v.GetNodeValue8())).Bytes()
			value = vInBytes
			return false, nil
		}

		return true, nil
	}

	root, err := s.Db.GetLastRoot()
	if err != nil {
		return nil, err
	}

	err = s.Traverse(context.Background(), root, action)
	if err != nil {
		return nil, err
	}

	return value, nil
}
