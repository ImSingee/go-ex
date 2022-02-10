package exos

import (
	"os"
)

func CurrentHasTTY() bool {
	return HasTTY(os.Getpid())
}

func GetCurrentTTY() (string, error) {
	return GetTTY(os.Getpid())
}

func GetTTY(pid int) (string, error) {
	return GetPSResult(pid, "tty")
}

func HasTTY(pid int) bool {
	tty, err := GetTTY(pid)
	if err != nil {
		return false
	}

	return tty != "" && tty != "?"
}
