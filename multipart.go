package resty

import (
	"io"
	"net/textproto"
	"strings"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string { _ = "STUB: not implemented"; return "" }

type MultipartField struct {
	Name string

	FileName string

	ContentType string

	Reader io.Reader

	FilePath string

	FileSize int64

	ProgressCallback MultipartFieldCallbackFunc

	Values []string

	tempBuf []byte
}

func (mf *MultipartField) Clone() *MultipartField { _ = "STUB: not implemented"; return nil }

func (mf *MultipartField) resetReader() error { _ = "STUB: not implemented"; return nil }

func (mf *MultipartField) isValues() bool { _ = "STUB: not implemented"; return false }

func (mf *MultipartField) close() { _ = "STUB: not implemented"; return }

func (mf *MultipartField) createHeader() textproto.MIMEHeader {
	_ = "STUB: not implemented"
	return *new(textproto.MIMEHeader)
}

func (mf *MultipartField) openFile() error { _ = "STUB: not implemented"; return nil }

func (mf *MultipartField) detectContentType() error { _ = "STUB: not implemented"; return nil }

func (mf *MultipartField) wrapProgressCallbackIfPresent(pw io.Writer) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

type MultipartFieldCallbackFunc func(MultipartFieldProgress)

type MultipartFieldProgress struct {
	Name     string
	FileName string
	FileSize int64
	Written  int64
}

func (mfp MultipartFieldProgress) String() string { _ = "STUB: not implemented"; return "" }

type multipartProgressWriter struct {
	w  io.Writer
	pb int64
	f  func(int64)
}

func (mpw *multipartProgressWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
