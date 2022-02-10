package exbytes

import (
	"bytes"
	"io"
)

func ToReader(p []byte) io.Reader {
	return bytes.NewBuffer(p)
}

// ZeroReader return a reader can read unlimited of zeros
func ZeroReader() io.Reader {
	return zeroReader{}
}

type zeroReader struct{}

func (zeroReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = 0
	}

	return len(p), nil
}
