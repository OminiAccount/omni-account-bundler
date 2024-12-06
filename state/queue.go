package state

import "github.com/OAB/utils/log"

func (s *State) GetBatchProof() ([]*Batch, error) {
	if s.proofQueue.IsEmpty() {
		return nil, nil
	}
	var res []*Batch
	for {
		if ok := s.proofQueue.IsEmpty(); ok {
			break
		}
		batch, _ := s.proofQueue.Dequeue()
		res = append(res, &batch)
	}
	return res, nil
}

func (s *State) SetBatchProofResult(result *ProofResult) error {
	log.Infof("SetBatchProofResult: %+v", result)
	s.provenQueue.Enqueue(*result)
	return nil
}
