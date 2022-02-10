package exstrings

import "strings"

// TrimEmptyLine 将字符串首尾的空行删除
func TrimEmptyLine(s string) string {
	return strings.Trim(s, "\n")
}
