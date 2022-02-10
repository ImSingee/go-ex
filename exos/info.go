package exos

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/ImSingee/go-ex/exbytes"
)

func GetPPID(pid int) (int, error) {
	ppidS, err := GetPSResult(pid, "ppid")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(ppidS)
}

func GetComm(pid int) (string, error) {
	result, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/comm", pid))
	if err != nil {
		return "", err
	}
	return exbytes.ToString(bytes.TrimSpace(result)), nil
}

func GetExe(pid int) (string, error) {
	return os.Readlink(fmt.Sprintf("/proc/%d/exe", pid))
}

func GetCmdline(pid int) ([]string, error) {
	result, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
	if err != nil {
		return nil, err
	}

	cmdBytes := bytes.Split(result, []byte{0})

	cmds := make([]string, 0, len(cmdBytes))
	for _, c := range cmdBytes {
		c = bytes.TrimSpace(c)
		if len(c) == 0 {
			continue
		}

		cmds = append(cmds, exbytes.ToString(c))
	}

	return cmds, nil
}
