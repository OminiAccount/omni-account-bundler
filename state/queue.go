package state

import "github.com/OAB/utils/log"

func (s *State) GetBatchProof() ([]Batch, error) {
	if s.proofQueue.IsEmpty() {
		return nil, nil
	}
	res, _ := s.proofQueue.Lrange(s.cfg.MaxBatches)
	return res, nil
}

func (s *State) SetBatchProofResult(res *ProofResult) error {
	log.Infof("SetBatchProofResult: %+v", res)
	var batch Batch
	batch, _ = s.proofQueue.Dequeue()
	if batch.NewNumBatch != res.Number {
		s.proofQueue.Lpush(batch)
		return nil
	}
	for i := res.Number + 1; i <= res.FinalNumber; i++ {
		_, _ = s.proofQueue.Dequeue()
	}
	s.provenQueue.Enqueue(*res)
	return nil
}
