package msgpack

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
)

// MarshalStruct Serialize an arbitrary structure
func MarshalStruct[T any](input T) ([]byte, error) {
	data, err := msgpack.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("error marshaling: %w", err)
	}
	return data, nil
}

// UnmarshalStruct Deserialize a byte array into an arbitrary structure
func UnmarshalStruct[T any](data []byte) (T, error) {
	var output T
	err := msgpack.Unmarshal(data, &output)
	if err != nil {
		return output, fmt.Errorf("error unmarshaling: %w", err)
	}
	return output, nil
}
