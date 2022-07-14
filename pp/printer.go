package pp

import (
	"io"

	"github.com/ImSingee/go-ex/exstrings"
)

type Printer struct {
	w       io.Writer
	NoColor bool
}

func NewPrinter(w io.Writer, noColor bool) *Printer {
	return &Printer{w, noColor}
}

func (p *Printer) ChangeWriter(newWriter io.Writer) {
	p.w = newWriter
}

func (p *Printer) CurrentWriter() io.Writer {
	return p.w
}

func (p *Printer) Write(b []byte) (n int, err error) {
	if p.w == nil { // no output
		return len(b), nil
	} else {
		return p.w.Write(b)
	}
}

func (p *Printer) Print(color *Color, a ...interface{}) error {
	_, err := p.Write(exstrings.UnsafeToBytes(color.Sprint(p.NoColor, a...)))
	return err
}

func (p *Printer) Printf(color *Color, format string, a ...interface{}) error {
	_, err := p.Write(exstrings.UnsafeToBytes(color.Sprintf(p.NoColor, format, a...)))
	return err
}

func (p *Printer) Println(color *Color, a ...interface{}) error {
	_, err := p.Write(exstrings.UnsafeToBytes(color.Sprintln(p.NoColor, a...)))
	return err
}

func (p *Printer) printf(color *Color, format string, a ...interface{}) error { //nolint:forbidigo
	if len(a) == 0 {
		return p.Print(color, format)
	} else {
		return p.Printf(color, format, a...)
	}
}
