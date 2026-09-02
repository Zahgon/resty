package resty

import (
	"io"
	"net/http"
	"time"
)

type Response struct {
	Request     *Request
	Body        io.ReadCloser
	RawResponse *http.Response
	IsRead      bool

	CascadeError error

	bodyBytes  []byte
	size       int64
	receivedAt time.Time
}

func (r *Response) Status() string { _ = "STUB: not implemented"; return "" }

func (r *Response) StatusCode() int { _ = "STUB: not implemented"; return 0 }

func (r *Response) Proto() string { _ = "STUB: not implemented"; return "" }

func (r *Response) Result() any { _ = "STUB: not implemented"; return *new(any) }

func (r *Response) ResultError() any { _ = "STUB: not implemented"; return *new(any) }

func (r *Response) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (r *Response) Cookies() []*http.Cookie { _ = "STUB: not implemented"; return nil }

func (r *Response) String() string { _ = "STUB: not implemented"; return "" }

func (r *Response) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (r *Response) Duration() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (r *Response) ReceivedAt() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (r *Response) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Response) IsStatusSuccess() bool { _ = "STUB: not implemented"; return false }

func (r *Response) IsStatusFailure() bool { _ = "STUB: not implemented"; return false }

func (r *Response) RedirectHistory() []*RedirectInfo { _ = "STUB: not implemented"; return nil }

func (r *Response) setReceivedAt() { _ = "STUB: not implemented"; return }

func (r *Response) fmtBodyString(sl int) string { _ = "STUB: not implemented"; return "" }

func (r *Response) readIfRequired() { _ = "STUB: not implemented"; return }

var ioReadAll = io.ReadAll

func (r *Response) readAll() (err error) { _ = "STUB: not implemented"; return nil }

func (r *Response) wrapLimitReadCloser() { _ = "STUB: not implemented"; return }

func (r *Response) wrapCopyReadCloser() { _ = "STUB: not implemented"; return }

func (r *Response) wrapContentDecompresser() error { _ = "STUB: not implemented"; return nil }

func (r *Response) wrapError(err error, preserve bool) error { _ = "STUB: not implemented"; return nil }
