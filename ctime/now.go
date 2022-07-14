package ctime

import "time"

// Now returns the current time, and set location information to China for display purposes.
func Now() time.Time {
	return time.Now().In(china)
}

// TodayString 当前时间转换为 20060102（年月日）的格式
func TodayString() string {
	return FormatNow("20060102")
}

// TimeString 当前时间转换为 150405（时分秒）的格式
func TimeString() string {
	return FormatNow("150405")
}

// NowString 当前时间转换为 200601021504（年月日时分）的格式
func NowString() string {
	return FormatNow("200601021504")
}
