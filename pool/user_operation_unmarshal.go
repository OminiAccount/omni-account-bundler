package pool

import (
	"encoding/json"
)

func (s *SignedUserOperation) UnmarshalJSON(input []byte) error {
	type dec SignedUserOperation
	aux := &struct {
		ChainId              string `json:"chainId"`
		MainChainGasPrice    string `json:"mainChainGasPrice"`
		DestChainGasPrice    string `json:"destChainGasPrice"`
		MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
		*dec
	}{
		dec: (*dec)(s),
	}

	if err := json.Unmarshal(input, &aux); err != nil {
		return err
	}

	/*if aux.ChainId != "" {
		chainId, ok := new(big.Int).SetString(aux.ChainId, 0)
		if !ok {
			return fmt.Errorf("invalid ChainId value: %s", aux.ChainId)
		}
		s.ChainId = (*hexutil.Big)(chainId)
	}*/

	/*if aux.MainChainGasPrice != "" {
		mainChainGasPrice, ok := new(big.Int).SetString(aux.MainChainGasPrice, 0)
		if !ok {
			return fmt.Errorf("invalid MainChainGasPrice value: %s", aux.MainChainGasPrice)
		}
		s.MainChainGasPrice = (*hexutil.Big)(mainChainGasPrice)
	}

	if aux.DestChainGasPrice != "" {
		destChainGasPrice, ok := new(big.Int).SetString(aux.DestChainGasPrice, 0)
		if !ok {
			return fmt.Errorf("invalid DestChainGasPrice value: %s", aux.DestChainGasPrice)
		}
		s.DestChainGasPrice = (*hexutil.Big)(destChainGasPrice)
	}*/

	//if aux.MaxPriorityFeePerGas != "" {
	//	maxPriorityFeePerGas, ok := new(big.Int).SetString(aux.MaxPriorityFeePerGas, 0)
	//	if !ok {
	//		return fmt.Errorf("invalid MaxPriorityFeePerGas value: %s", aux.MaxPriorityFeePerGas)
	//	}
	//	s.MaxPriorityFeePerGas = (*hexutil.Big)(maxPriorityFeePerGas)
	//}

	s.RecoverAddress()

	return nil
}
