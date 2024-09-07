package smt

import (
	"encoding/json"
	"strconv"
)

func (d DeltaMerkleProof) MarshalJSON() ([]byte, error) {
	type Proof DeltaMerkleProof
	return json.Marshal(&struct {
		Index string `json:"index"`
		*Proof
	}{
		Index: strconv.FormatInt(int64(d.Index), 16),
		Proof: (*Proof)(&d),
	})
}
