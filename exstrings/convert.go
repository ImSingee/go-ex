package exstrings

import "unsafe"

// UnsafeToBytes 将 string 转换为 []byte，避免拷贝导致的内存开销
// 需要注意的是，如果试图对返回的 []byte 进行修改可能会导致程序直接崩溃
func UnsafeToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// Bytes 将 string 转换为 []byte
func Bytes(s string) []byte {
	buf := make([]byte, len(s))
	copy(buf, s)
	return buf
}
