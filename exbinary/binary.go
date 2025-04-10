//go:build go1.23

package exbinary

import (
	"encoding/binary"
)

func EncodeBigEndian[T any](v T) []byte {
	return Encode(v, binary.BigEndian)
}

func DecodeBigEndian[T any](buf []byte) (v T, err error) {
	return Decode[T](buf, binary.BigEndian)
}

func EncodeLittleEndian[T any](v T) []byte {
	return Encode(v, binary.LittleEndian)
}

func DecodeLittleEndian[T any](buf []byte) (v T, err error) {
	return Decode[T](buf, binary.LittleEndian)
}

func Encode[T any](v T, order binary.ByteOrder) []byte {
	buf := make([]byte, binary.Size(v))
	_, _ = binary.Encode(buf, order, v)
	return buf
}

func Decode[T any](buf []byte, order binary.ByteOrder) (v T, err error) {
	_, err = binary.Decode(buf, order, &v)
	return
}
