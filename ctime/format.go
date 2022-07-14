package ctime

import "time"

func Format(t time.Time, layout string) string {
	return ToChina(t).Format(layout)
}

func FormatNow(layout string) string {
	return Now().Format(layout)
}
