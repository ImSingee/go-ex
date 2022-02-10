package exints

import (
	"strconv"
	"strings"
)

// ListToInterface 将 []int 转换为 []interface
func ListToInterface(ss []int) []interface{} {
	result := make([]interface{}, len(ss))

	for i, s := range ss {
		result[i] = s
	}

	return result
}

func Join(ss []int, sep string) string {
	b := strings.Builder{}

	for i, x := range ss {
		if i != 0 {
			b.WriteString(sep)
		}
		b.WriteString(strconv.Itoa(x))
	}

	return b.String()
}
