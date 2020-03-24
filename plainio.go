package lib

import (
	"io"
)

type PlainWriter struct {
	writer io.Writer
}

func NewPlainWriter(w io.Writer) *PlainWriter {
	return &PlainWriter{writer: w}
}

func (w *PlainWriter) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func (*PlainWriter) Close() error {
	return nil
}

type PlainReader struct {
	reader io.Reader
}

func NewPlainReader(r io.Reader) *PlainReader {
	return &PlainReader{reader: r}
}

func (r *PlainReader) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}

func (*PlainReader) Close() error {
	return nil
}
