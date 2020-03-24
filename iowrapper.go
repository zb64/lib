package lib

import (
	"io"
)

type PlainWriter struct {
	w io.Writer
}

func NewPlainWriter(w io.Writer) *PlainWriter {
	return &PlainWriter{w: w}
}

func (w *PlainWriter) Write(p []byte) (n int, err error) {
	return w.Write(p)
}

func (*PlainWriter) Close() error {
	return nil
}

type PlainReader struct {
	r io.Reader
}

func NewPlainReader(r io.Reader) *PlainReader {
	return &PlainReader{r: r}
}

func (r *PlainReader) Read(p []byte) (n int, err error) {
	return r.Read(p)
}

func (*PlainReader) Close() error {
	return nil
}
