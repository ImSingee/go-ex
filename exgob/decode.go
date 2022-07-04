package exgob

import (
	"bytes"
	"encoding/gob"
	"github.com/ImSingee/go-ex/exstrings"
)

func Decode[T any](p []byte) (*T, error) {
	v, err := DecodeToValue[T](p)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func DecodeString[T any](s string) (*T, error) {
	return Decode[T](exstrings.UnsafeToBytes(s))
}

func DecodeToValue[T any](p []byte) (T, error) {
	var v T
	err := DecodeTo(p, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func DecodeStringToValue[T any](s string) (T, error) {
	return DecodeToValue[T](exstrings.UnsafeToBytes(s))
}

func DecodeTo[T any](p []byte, target *T) error {
	return gob.NewDecoder(bytes.NewReader(p)).Decode(target)
}

func DecodeStringTo[T any](s string, target *T) error {
	return DecodeTo[T](exstrings.UnsafeToBytes(s), target)
}
