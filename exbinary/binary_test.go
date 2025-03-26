package exbinary

import (
	"testing"
)

func doTestEncodeDecodeBE[T comparable](t *testing.T, value T) {
	encoded := EncodeBigEndian(value)
	if len(encoded) == 0 {
		t.Error("Encoded data is empty")
	}

	decoded, err := DecodeBigEndian[T](encoded)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if decoded != value {
		t.Errorf("Decoded value = %v, want %v", decoded, value)
	}
}

func doTestEncodeDecodeLE[T comparable](t *testing.T, value T) {
	encoded := EncodeLittleEndian(value)
	if len(encoded) == 0 {
		t.Error("Encoded data is empty")
	}

	decoded, err := DecodeLittleEndian[T](encoded)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if decoded != value {
		t.Errorf("Decoded value = %v, want %v", decoded, value)
	}
}

func TestEncodeDecode(t *testing.T) {
	t.Run("int32 LE", func(t *testing.T) {
		doTestEncodeDecodeBE[int32](t, 12345)
	})

	t.Run("int32 BE", func(t *testing.T) {
		doTestEncodeDecodeBE[int32](t, 12345)
	})

	t.Run("uint32 LE", func(t *testing.T) {
		doTestEncodeDecodeLE[uint32](t, 12345)
	})

	t.Run("uint32 BE", func(t *testing.T) {
		doTestEncodeDecodeBE[uint32](t, 12345)
	})

	t.Run("int64 LE", func(t *testing.T) {
		doTestEncodeDecodeLE[int64](t, 123456789)
	})

	t.Run("int64 BE", func(t *testing.T) {
		doTestEncodeDecodeBE[int64](t, 123456789)
	})

	t.Run("float32 LE", func(t *testing.T) {
		doTestEncodeDecodeLE[float32](t, 123.45)
	})

	t.Run("float32 BE", func(t *testing.T) {
		doTestEncodeDecodeBE[float32](t, 123.45)
	})

	t.Run("float64 LE", func(t *testing.T) {
		doTestEncodeDecodeLE[float64](t, 123.45)
	})

	t.Run("float64 BE", func(t *testing.T) {
		doTestEncodeDecodeBE[float64](t, 123.45)
	})
}
