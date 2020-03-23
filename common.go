package lib

import (
	"bytes"
	"encoding/base64"
	"io"
)

type (
	newFormatWriterFunc func(io.Writer) io.WriteCloser
	newFormatReaderFunc func(io.Reader) (io.ReadCloser, error)
)

type dataFormat struct {
	newWriter newFormatWriterFunc
	newReader newFormatReaderFunc
}

func newDataFormat(writer newFormatWriterFunc, reader newFormatReaderFunc) *dataFormat {
	return &dataFormat{
		newWriter: writer,
		newReader: reader,
	}
}

// Encode compresses data using the given compressed data writer, and encodes as base64-encoded result string.
func (z *dataFormat) Encode(data []byte) (result string, err error) {
	var buf bytes.Buffer
	w := z.newWriter(&buf)
	if _, err = w.Write(data); err != nil {
		return
	}
	if err = w.Close(); err != nil {
		return
	}
	result = base64.StdEncoding.EncodeToString(buf.Bytes())
	return
}

// Decode decodes base64-encoded string and decompresses data using the given compressed data reader.
func (z *dataFormat) Decode(raw string) (result []byte, err error) {
	var (
		data   []byte
		outBuf bytes.Buffer
		r      io.ReadCloser
	)
	if data, err = base64.StdEncoding.DecodeString(raw); err != nil {
		return
	}

	if r, err = z.newReader(bytes.NewReader(data)); err != nil {
		return
	}
	if _, err = io.Copy(&outBuf, r); err != nil {
		return
	}
	if err = r.Close(); err != nil {
		return
	}
	result = outBuf.Bytes()
	return
}
