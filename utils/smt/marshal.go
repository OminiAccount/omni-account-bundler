package smt

import (
	"encoding/json"
)

func (d DeltaMerkleProof) MarshalJSON() ([]byte, error) {
	type Proof DeltaMerkleProof
	return json.Marshal(&struct {
		Index string `json:"index"`
		*Proof
	}{
		Index: d.Index.Text(16),
		Proof: (*Proof)(&d),
	})
}
