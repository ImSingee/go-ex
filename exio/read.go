package exio

import (
	"bytes"
	"io"
)

// ReadAll is like io.ReadAll but more efficient
func ReadAll(r io.Reader) ([]byte, error) {
	b := bytes.Buffer{}
	_, err := io.Copy(&b, r)
	return b.Bytes(), err
}
