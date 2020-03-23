package lib

import (
	"compress/flate"
	"compress/lzw"
	"io"
)

const (
	lzwLitWidth = 8
)

var (
	Deflate = newDataFormat(
		func(buf io.Writer) io.WriteCloser {
			w, _ := flate.NewWriter(buf, flate.BestCompression)
			return w
		},
		func(buf io.Reader) (io.ReadCloser, error) {
			return flate.NewReader(buf), nil
		})
	LzwMSB = newDataFormat(
		func(buf io.Writer) io.WriteCloser {
			return lzw.NewWriter(buf, lzw.MSB, lzwLitWidth)
		}, func(buf io.Reader) (io.ReadCloser, error) {
			return lzw.NewReader(buf, lzw.MSB, lzwLitWidth), nil
		})
	LzwLSB = newDataFormat(
		func(buf io.Writer) io.WriteCloser {
			return lzw.NewWriter(buf, lzw.LSB, lzwLitWidth)
		}, func(buf io.Reader) (io.ReadCloser, error) {
			return lzw.NewReader(buf, lzw.LSB, lzwLitWidth), nil
		})
)
