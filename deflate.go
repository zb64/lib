package lib

import (
	"compress/flate"
	"io"
)

var (
	Deflate = newDataFormat(
		func(buf io.Writer) io.WriteCloser {
			w, _ := flate.NewWriter(buf, flate.BestCompression)
			return w
		}, func(buf io.Reader) (io.ReadCloser, error) {
			return flate.NewReader(buf), nil
		})
)
