package ctime

import "time"

// TimeFromUnix 使用秒级时间戳获取时间对象
func TimeFromUnix(sec int64) time.Time {
	if sec == 0 {
		return time.Time{}
	}

	return ToChina(time.Unix(sec, 0))
}

// TimeFromUnixMicro 使用微秒级时间戳获取时间对象
func TimeFromUnixMicro(msec uint64) time.Time {
	if msec == 0 {
		return time.Time{}
	}

	sec := msec / 1e6

	return ToChina(time.Unix(int64(sec), int64(msec-sec*1e6)))
}

// TimeFromUnixNano 使用纳秒级时间戳获取时间对象
func TimeFromUnixNano(nsec uint64) time.Time {
	if nsec == 0 {
		return time.Time{}
	}

	sec := nsec / 1e9

	return ToChina(time.Unix(int64(sec), int64(nsec-sec*1e9)))
}
