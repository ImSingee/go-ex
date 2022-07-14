package extime

import "time"

// TodayString 当前时间转换为 20060102（年月日）的格式
func TodayString() string {
	return time.Now().Format("20060102")
}

// TimeString 当前时间转换为 150405（时分秒）的格式
func TimeString() string {
	return time.Now().Format("150405")
}

// NowString 当前时间转换为 200601021504（年月日时分）的格式
func NowString() string {
	return time.Now().Format("200601021504")
}
