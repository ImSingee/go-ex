package exbytes

import (
	"bytes"
	"io"
)

func ToWriterFunc(p []byte) func(writer io.Writer) error {
	return func(writer io.Writer) error {
		_, err := io.Copy(writer, bytes.NewReader(p))
		return err
	}
}
