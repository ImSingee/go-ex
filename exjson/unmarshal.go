package exjson

import (
	"encoding/json"
	"github.com/ImSingee/go-ex/exstrings"
)

func unmarshal[T any](data []byte) (v T, err error) {
	err = json.Unmarshal(data, &v)
	return
}

func Unmarshal[T any](data []byte) (T, error) {
	return unmarshal[T](data)
}

func UnmarshalString[T any](data string) (T, error) {
	return unmarshal[T](exstrings.UnsafeToBytes(data))
}
