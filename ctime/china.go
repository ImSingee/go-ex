package ctime

import (
	"time"
)

var china = time.FixedZone("Asia/Shanghai", 8*60*60)

func China() *time.Location {
	return china
}

func ToChina(t time.Time) time.Time {
	return t.In(china)
}
