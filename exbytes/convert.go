package exbytes

import "unsafe"

// ToString 将 []byte 转换为 string，避免拷贝导致的内存开销
func ToString(s []byte) string {
	if s == nil {
		return ""
	}

	return *(*string)(unsafe.Pointer(&s))
}
