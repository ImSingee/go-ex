package pp

import "os"

var Stderr = NewPrinter(os.Stderr, false)

func EPrint(a ...interface{}) {
	_ = Stderr.Print(nil, a...)
}

func EPrintf(format string, a ...interface{}) {
	_ = Stderr.printf(nil, format, a...)
}

func EPrintln(a ...interface{}) {
	_ = Stderr.Println(nil, a...)
}
