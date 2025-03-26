package exbinary

import (
	"encoding/binary"
)

func Encode[T any](v T, order binary.ByteOrder) []byte {
	buf := make([]byte, binary.Size(v))
	_, _ = binary.Encode(buf, order, v)
	return buf
}

func Decode[T any](buf []byte, order binary.ByteOrder) (v T, err error) {
	_, err = binary.Decode(buf, order, &v)
	return
}
