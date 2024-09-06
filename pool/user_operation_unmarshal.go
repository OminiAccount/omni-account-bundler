package pool

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func (s *SignedUserOperation) UnmarshalJSON(input []byte) error {
	type dec SignedUserOperation
	aux := &struct {
		ChainId              string `json:"chainId"`
		PreVerificationGas   string `json:"preVerificationGas"`
		MaxFeePerGas         string `json:"maxFeePerGas"`
		MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
		*dec
	}{
		dec: (*dec)(s),
	}

	if err := json.Unmarshal(input, &aux); err != nil {
		return err
	}

	if aux.ChainId != "" {
		chainId, ok := new(big.Int).SetString(aux.ChainId, 0)
		if !ok {
			return fmt.Errorf("invalid ChainId value: %s", aux.ChainId)
		}
		s.ChainId = (*hexutil.Big)(chainId)
	}

	if aux.PreVerificationGas != "" {
		preVerificationGas, ok := new(big.Int).SetString(aux.PreVerificationGas, 0)
		if !ok {
			return fmt.Errorf("invalid PreVerificationGas value: %s", aux.PreVerificationGas)
		}
		s.PreVerificationGas = (*hexutil.Big)(preVerificationGas)
	}

	if aux.MaxFeePerGas != "" {
		maxFeePerGas, ok := new(big.Int).SetString(aux.MaxFeePerGas, 0)
		if !ok {
			return fmt.Errorf("invalid MaxFeePerGas value: %s", aux.MaxFeePerGas)
		}
		s.MaxFeePerGas = (*hexutil.Big)(maxFeePerGas)
	}

	if aux.MaxPriorityFeePerGas != "" {
		maxPriorityFeePerGas, ok := new(big.Int).SetString(aux.MaxPriorityFeePerGas, 0)
		if !ok {
			return fmt.Errorf("invalid MaxPriorityFeePerGas value: %s", aux.MaxPriorityFeePerGas)
		}
		s.MaxPriorityFeePerGas = (*hexutil.Big)(maxPriorityFeePerGas)
	}

	return nil
}
