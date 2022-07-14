package pp

import (
	"os"
	"strings"
)

var Stdout = NewPrinter(os.Stdout, false)

func Print(a ...interface{}) {
	_ = Stdout.Print(nil, a...)
}

func Printf(format string, a ...interface{}) {
	_ = Stdout.printf(nil, format, a...)
}

func Println(a ...interface{}) {
	_ = Stdout.Println(nil, a...)
}

func PrintLine(s string) {
	Println(strings.TrimSpace(s))
}

func PrintLineIfNotEmpty(s string) {
	s = strings.TrimSpace(s)

	if s != "" {
		Println(s)
	}
}
