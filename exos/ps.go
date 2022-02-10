package exos

import (
	"bytes"
	"os/exec"
	"strconv"

	"github.com/ImSingee/go-ex/exbytes"
)

func GetPSResult(pid int, option string) (string, error) {
	result, err := exec.Command("ps", "-p", strconv.Itoa(pid), "-o", option+"=").CombinedOutput()
	if err != nil {
		return "", err
	}

	return exbytes.ToString(bytes.TrimSpace(result)), nil
}
