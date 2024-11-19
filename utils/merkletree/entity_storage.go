package merkletree

import (
	"github.com/OAB/utils/merkletreeutils"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// SetAccountState sets the balance and nonce of an account
func (s *SMT) SetAccountState(ethAddr string, balance, nonce *big.Int, chainId string) (*big.Int, error) {
	_, err := s.SetAccountBalance(ethAddr, balance)
	if err != nil {
		return nil, err
	}

	auxOut, err := s.SetAccountNonce(ethAddr, nonce, chainId)
	if err != nil {
		return nil, err
	}

	return auxOut, nil
}

// SetAccountBalance sets the balance of an account
func (s *SMT) SetAccountBalance(ethAddr string, balance *big.Int) (*big.Int, error) {
	keyBalance := merkletreeutils.KeyEthAddrBalance(ethAddr)

	response, err := s.InsertKA(keyBalance, balance)
	if err != nil {
		return nil, err
	}

	ks := merkletreeutils.EncodeKeySource(merkletreeutils.KeyBalance, merkletreeutils.ConvertHexToAddress(ethAddr), common.Hash{})
	err = s.Db.InsertKeySource(keyBalance, ks)
	if err != nil {
		return nil, err
	}

	return response.NewRootScalar.ToBigInt(), err
}

// SetAccountNonce sets the nonce of an account
func (s *SMT) SetAccountNonce(ethAddr string, nonce *big.Int, nonceChainId string) (*big.Int, error) {
	keyNonce := merkletreeutils.KeyEthAddrNonceChainId(ethAddr, nonceChainId)

	response, err := s.InsertKA(keyNonce, nonce)
	if err != nil {
		return nil, err
	}

	ks := merkletreeutils.EncodeKeySource(merkletreeutils.KeyNonce, merkletreeutils.ConvertHexToAddress(ethAddr), common.Hash{})
	err = s.Db.InsertKeySource(keyNonce, ks)
	if err != nil {
		return nil, err
	}

	return response.NewRootScalar.ToBigInt(), nil
}
