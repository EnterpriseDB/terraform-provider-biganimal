package utils

import "encoding/json"

// json marshals to bytes then unmarshals into struct to fill a struct. Out has to be a pointer
func CopyObjectJson(in any, out any) error {
	bytes, err := json.Marshal(in)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, out); err != nil {
		return err
	}

	return nil
}

// ToPointer returns a pointer of given data which kind of T
func ToPointer[T any](data T) *T {
	return &data
}
