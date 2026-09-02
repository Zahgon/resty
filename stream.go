package resty

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"io"
	"sync"
)

var (
	ErrContentDecompresserNotFound = errors.New("resty: content decoder not found")

	maxDecodeObjects = 1000001
)

type (
	ContentTypeEncoder func(io.Writer, any) error

	ContentTypeDecoder func(io.Reader, any) error

	ContentDecompresser func(io.ReadCloser) (io.ReadCloser, error)
)

func encodeJSON(w io.Writer, v any) error { _ = "STUB: not implemented"; return nil }

func encodeJSONEscapeHTML(w io.Writer, v any, esc bool) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeJSONEscapeHTMLIndent(w io.Writer, v any, esc bool, indent string) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeJSON(r io.Reader, v any) error { _ = "STUB: not implemented"; return nil }

func doDecodeJSON(dec *json.Decoder, v any) error { _ = "STUB: not implemented"; return nil }

func encodeXML(w io.Writer, v any) error { _ = "STUB: not implemented"; return nil }

func decodeXML(r io.Reader, v any) error { _ = "STUB: not implemented"; return nil }

var gzipReaderPool = sync.Pool{
	New: func() any {

		return nil
	},
}

type gzipReaderWrapper struct {
	mu *sync.Mutex
	r  io.ReadCloser
	gr *gzip.Reader
}

func acquireGzipReader(r io.ReadCloser) (*gzipReaderWrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func releaseGzipReader(w *gzipReaderWrapper) { _ = "STUB: not implemented"; return }

func decompressGzip(r io.ReadCloser) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (w *gzipReaderWrapper) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *gzipReaderWrapper) Close() error { _ = "STUB: not implemented"; return nil }

var flateReaderPool = sync.Pool{
	New: func() any {

		return nil
	},
}

type deflateReaderWrapper struct {
	mu *sync.Mutex
	r  io.ReadCloser
	fr io.ReadCloser
}

func acquireDeflateReader(r io.ReadCloser) (*deflateReaderWrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func releaseDeflateReader(w *deflateReaderWrapper) { _ = "STUB: not implemented"; return }

func decompressDeflate(r io.ReadCloser) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (w *deflateReaderWrapper) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *deflateReaderWrapper) Close() error { _ = "STUB: not implemented"; return nil }

var ErrReadExceedsThresholdLimit = errors.New("resty: read exceeds the threshold limit")

var _ io.ReadCloser = (*limitReadCloser)(nil)
var _ resetter = (*limitReadCloser)(nil)

type resetter interface {
	Reset() error
}

const unlimitedRead = 0

type limitReadCloser struct {
	r io.Reader
	l int64
	t int64
	f func(s int64)
}

func (l *limitReadCloser) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (l *limitReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

func (l *limitReadCloser) Reset() error { _ = "STUB: not implemented"; return nil }

var _ io.ReadCloser = (*copyReadCloser)(nil)

type copyReadCloser struct {
	s io.Reader
	t *bytes.Buffer
	c bool
	f func(*bytes.Buffer)
}

func (r *copyReadCloser) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *copyReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

var _ io.ReadCloser = (*cancelReadCloser)(nil)

type cancelReadCloser struct {
	r      io.ReadCloser
	cancel context.CancelFunc
}

func (c *cancelReadCloser) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *cancelReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

var _ io.ReadCloser = (*nopReadCloser)(nil)

type nopReadCloser struct {
	r          io.Reader
	resetOnEOF bool
}

func (r *nopReadCloser) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *nopReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

func (r *nopReadCloser) Reset() { _ = "STUB: not implemented"; return }

var _ flate.Reader = (*nopReader)(nil)

type nopReader struct{}

func (nopReader) Read([]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
func (nopReader) ReadByte() (byte, error)  { _ = "STUB: not implemented"; return 0, nil }

type gracefulStopReader struct {
	ctx context.Context
	r   io.Reader
}

func (gsr *gracefulStopReader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
