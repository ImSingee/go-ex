package exstrings

import "github.com/ImSingee/go-ex/exbytes"

// repeat 重复字符串到指定长度, []byte必须有充足的容量。
func repeat(b []byte, pad string, padLen int) {
	bp := copy(b[:padLen], pad)
	for bp < padLen {
		copy(b[bp:padLen], b[:bp])
		bp *= 2
	}
}

// RightPad 使用另一个字符串从右端填充字符串为指定长度。
func RightPad(s, pad string, c int) string {
	padLen := c - len(s)
	if padLen <= 0 {
		return s
	}

	b := make([]byte, c)
	l := copy(b, s)
	repeat(b[l:], pad, padLen)

	return exbytes.ToString(b)
}
func RightPadSpace(s string, minLength int) string {

	return RightPad(s, " ", minLength)
}
