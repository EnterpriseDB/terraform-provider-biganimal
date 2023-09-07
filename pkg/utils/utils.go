package utils

import (
	"encoding/json"
	"reflect"
)

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

// ToValue returns an object to which p points, if p is nil, will return a zero value
func ToValue[T any](p *T) T {
	if p == nil {
		t := reflect.ValueOf(p).Type().Elem()
		v := reflect.New(t)

		return v.Elem().Interface().(T)
	}
	return *p
}
