package exstrings

import (
	"io"
	"strings"
)

func ToReader(s string) io.Reader {
	return strings.NewReader(s)
}
