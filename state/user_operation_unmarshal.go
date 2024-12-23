package state

import (
	"encoding/json"
	"fmt"
	"github.com/OAB/lib/common/hexutil"
	"math/big"
)

func (s *ExecData) UnmarshalJSON(input []byte) error {
	type dec ExecData
	aux := &struct {
		MainChainGasPrice string `json:"mainChainGasPrice"`
		DestChainGasPrice string `json:"destChainGasPrice"`
		*dec
	}{
		dec: (*dec)(s),
	}
	if err := json.Unmarshal(input, &aux); err != nil {
		return err
	}

	if aux.MainChainGasPrice != "" {
		mainChainGasPrice, ok := new(big.Int).SetString(aux.MainChainGasPrice, 0)
		if !ok {
			return fmt.Errorf("invalid MainChainGasPrice value: %s", aux.MainChainGasPrice)
		}
		s.MainChainGasPrice = (*hexutil.Big)(mainChainGasPrice)
	}
	if s.MainChainGasPrice == nil {
		s.MainChainGasPrice = (*hexutil.Big)(big.NewInt(0))
	}

	if aux.DestChainGasPrice != "" {
		destChainGasPrice, ok := new(big.Int).SetString(aux.DestChainGasPrice, 0)
		if !ok {
			return fmt.Errorf("invalid DestChainGasPrice value: %s", aux.DestChainGasPrice)
		}
		s.DestChainGasPrice = (*hexutil.Big)(destChainGasPrice)
	}
	if s.DestChainGasPrice == nil {
		s.DestChainGasPrice = (*hexutil.Big)(big.NewInt(0))
	}

	return nil
}

func (s *SignedUserOperation) UnmarshalJSON(input []byte) error {
	type dec SignedUserOperation
	aux := &struct {
		*dec
	}{
		dec: (*dec)(s),
	}

	if err := json.Unmarshal(input, &aux); err != nil {
		return err
	}

	if aux.Exec.MainChainGasPrice == nil ||
		aux.Exec.DestChainGasPrice == nil ||
		aux.Exec.Nonce < 1 ||
		aux.Exec.ChainId.Uint64() < 1 {
		return fmt.Errorf("invalid exec data")
	}

	if aux.InnerExec.Nonce > 0 && (aux.InnerExec.MainChainGasPrice == nil ||
		aux.InnerExec.DestChainGasPrice == nil ||
		aux.InnerExec.ChainId.Uint64() < 1) {
		return fmt.Errorf("invalid innerexec data")
	}

	s.RecoverAddress()

	return nil
}
