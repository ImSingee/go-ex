package ctime

import (
	"time"
)

// Parse 解析时间字符串（基于 UTC）
func Parse(layout, value string) (time.Time, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return time.Time{}, err
	} else {
		return ToChina(t), nil
	}
}

// ParseInChina 解析时间字符串（基于 GMT+8）
func ParseInChina(layout, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, china)
}
