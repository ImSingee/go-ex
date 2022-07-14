package extime

import (
	"fmt"
	"time"
)

func SecondsToString(seconds int) string {
	if seconds == 0 {
		return ""
	} else if seconds < 60 { // [0, 1分钟)
		return fmt.Sprintf("%2ds ", seconds)
	} else if seconds < 60*60 { // [1分钟, 1小时)
		minutes := seconds / 60
		remainSeconds := seconds - minutes*60

		return fmt.Sprintf("%2dm %s", minutes, SecondsToString(remainSeconds))
	} else if seconds < 24*60*60 { // [1小时, 1天)
		hours := seconds / (60 * 60)
		remainSeconds := seconds - hours*60*60

		return fmt.Sprintf("%2dh %s", hours, SecondsToString(remainSeconds))
	} else { // [1天, +∞)
		days := seconds / (24 * 60 * 60)
		remainSeconds := seconds - days*60*60*24

		return fmt.Sprintf("%2dd %s", days, SecondsToString(remainSeconds))
	}
}

func TimeToRelative(d time.Time) string {
	seconds := int(time.Since(d) / time.Second)

	s := SecondsToString(seconds)
	if s == "" {
		s = "0s "
	}

	if seconds >= 0 {
		return s + "ago"
	} else {
		return s + "later"
	}
}
