package exgob

import (
	"bytes"
	"encoding/gob"

	"github.com/ImSingee/go-ex/exbytes"
)

func Encode(v any) ([]byte, error) {
	var w bytes.Buffer
	err := gob.NewEncoder(&w).Encode(v)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func EncodeToString(v any) (string, error) {
	p, err := Encode(v)
	if err != nil {
		return "", err
	}
	return exbytes.ToString(p), nil
}
