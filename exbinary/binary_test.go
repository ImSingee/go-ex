package exbinary

import (
	"encoding/binary"
	"testing"
)

func doTestEncodeDecode[T comparable](t *testing.T, value T, order binary.ByteOrder) {
	encoded := Encode(value, order)
	if len(encoded) == 0 {
		t.Error("Encoded data is empty")
	}

	decoded, err := Decode[T](encoded, order)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if decoded != value {
		t.Errorf("Decoded value = %v, want %v", decoded, value)
	}
}

func TestEncodeDecode(t *testing.T) {
	t.Run("int32 LE", func(t *testing.T) {
		doTestEncodeDecode[int32](t, 12345, binary.LittleEndian)
	})

	t.Run("int32 BE", func(t *testing.T) {
		doTestEncodeDecode[int32](t, 12345, binary.BigEndian)
	})

	t.Run("uint32 LE", func(t *testing.T) {
		doTestEncodeDecode[uint32](t, 12345, binary.LittleEndian)
	})

	t.Run("uint32 BE", func(t *testing.T) {
		doTestEncodeDecode[uint32](t, 12345, binary.BigEndian)
	})

	t.Run("int64 LE", func(t *testing.T) {
		doTestEncodeDecode[int64](t, 123456789, binary.LittleEndian)
	})

	t.Run("int64 BE", func(t *testing.T) {
		doTestEncodeDecode[int64](t, 123456789, binary.BigEndian)
	})

	t.Run("float32 LE", func(t *testing.T) {
		doTestEncodeDecode[float32](t, 123.45, binary.LittleEndian)
	})

	t.Run("float32 BE", func(t *testing.T) {
		doTestEncodeDecode[float32](t, 123.45, binary.BigEndian)
	})

	t.Run("float64 LE", func(t *testing.T) {
		doTestEncodeDecode[float64](t, 123.45, binary.LittleEndian)
	})

	t.Run("float64 BE", func(t *testing.T) {
		doTestEncodeDecode[float64](t, 123.45, binary.BigEndian)
	})
}
