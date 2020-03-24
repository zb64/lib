package lib

import (
	"compress/flate"
	"compress/lzw"
	"io"
)

const (
	lzwLitWidth = 8
)

// Deflate is a wrapper of DEFLATE compressed data format, provided by compress/flate.
var Deflate = newDataFormat(
	func(buf io.Writer) io.WriteCloser {
		w, _ := flate.NewWriter(buf, flate.BestCompression)
		return w
	},
	func(buf io.Reader) (io.ReadCloser, error) {
		return flate.NewReader(buf), nil
	},
)

// LzwMSB is a wrapper of Lempel-Ziv-Welch compressed data format using Least Significant Bits, provided by compress/lzw.
var LzwLSB = newDataFormat(
	func(buf io.Writer) io.WriteCloser {
		return lzw.NewWriter(buf, lzw.LSB, lzwLitWidth)
	},
	func(buf io.Reader) (io.ReadCloser, error) {
		return lzw.NewReader(buf, lzw.LSB, lzwLitWidth), nil
	},
)

// LzwMSB is a wrapper of Lempel-Ziv-Welch compressed data format using Most Significant Bits, provided by compress/lzw.
var LzwMSB = newDataFormat(
	func(buf io.Writer) io.WriteCloser {
		return lzw.NewWriter(buf, lzw.MSB, lzwLitWidth)
	},
	func(buf io.Reader) (io.ReadCloser, error) {
		return lzw.NewReader(buf, lzw.MSB, lzwLitWidth), nil
	},
)

// Noop is a wrapper of plain text format, no compression.
var Noop = newDataFormat(
	func(buf io.Writer) io.WriteCloser {
		return NewPlainWriter(buf)
	},
	func(buf io.Reader) (io.ReadCloser, error) {
		return NewPlainReader(buf), nil
	},
)
