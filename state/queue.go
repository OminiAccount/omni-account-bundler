package state

func (s *State) GetBatchProof() (*Batch, error) {
	if s.proofQueue.IsEmpty() {
		return nil, nil
	}
	batch, _ := s.proofQueue.Dequeue()
	return &batch, nil
}

func (s *State) SetBatchProofResult(result *ProofResult) error {
	s.provenQueue.Enqueue(*result)
	return nil
}
