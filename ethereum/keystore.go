package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
	"path/filepath"
)

// getAuthByAddress tries to get an authorization from the authorizations map
func (ether *Ethereum) getAuthByAddress(addr common.Address) (bind.TransactOpts, error) {
	auth, found := ether.auth[addr]
	if !found {
		return bind.TransactOpts{}, fmt.Errorf("not found %s", addr.Hex())
	}
	return auth, nil
}

// LoadAuthFromKeyStore loads an authorization from a key store file
func (ether *Ethereum) LoadAuthFromKeyStore(path, password string, chainID uint64) (*bind.TransactOpts, *ecdsa.PrivateKey, error) {
	auth, pk, err := newAuthFromKeystore(path, password, chainID)
	if err != nil {
		return nil, nil, err
	}

	ether.logger.Info("loaded authorization for address", "address", auth.From.String())
	ether.auth[auth.From] = auth
	return &auth, pk, nil
}

// newKeyFromKeystore creates an instance of a keystore key from a keystore file
func newKeyFromKeystore(path, password string) (*keystore.Key, error) {
	if path == "" && password == "" {
		return nil, nil
	}
	keystoreEncrypted, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(keystoreEncrypted, password)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// newAuthFromKeystore an authorization instance from a keystore file
func newAuthFromKeystore(path, password string, chainID uint64) (bind.TransactOpts, *ecdsa.PrivateKey, error) {
	key, err := newKeyFromKeystore(path, password)
	if err != nil {
		return bind.TransactOpts{}, nil, err
	}
	if key == nil {
		return bind.TransactOpts{}, nil, nil
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, new(big.Int).SetUint64(chainID))
	if err != nil {
		return bind.TransactOpts{}, nil, err
	}
	return *auth, key.PrivateKey, nil
}
