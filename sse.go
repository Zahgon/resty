package resty

import (
	"bufio"
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	defaultSseMaxBufSize = 1 << 15
	defaultEventName     = "message"
	defaultHTTPMethod    = MethodGet

	headerID    = []byte("id:")
	headerData  = []byte("data:")
	headerEvent = []byte("event:")
	headerRetry = []byte("retry:")

	hdrCacheControlKey = http.CanonicalHeaderKey("Cache-Control")
	hdrConnectionKey   = http.CanonicalHeaderKey("Connection")
	hdrLastEvevntID    = http.CanonicalHeaderKey("Last-Event-ID")
)

type (
	SSEOpenFunc func(url string, respHdr http.Header)

	SSEMessageFunc func(any)

	SSEErrorFunc func(error)

	SSERequestFailureFunc func(err error, res *http.Response)

	SSE struct {
		ID   string
		Name string
		Data string
	}

	SSESource struct {
		lock             *sync.RWMutex
		url              string
		method           string
		header           http.Header
		bodyBytes        []byte
		lastEventID      string
		retryCount       int
		retryWaitTime    time.Duration
		retryMaxWaitTime time.Duration
		retryConditions  []RetryConditionFunc
		serverSentRetry  time.Duration
		maxBufSize       int
		onOpen           SSEOpenFunc
		onError          SSEErrorFunc
		onRequestFailure SSERequestFailureFunc
		onEvent          map[string]*callback
		log              Logger
		ctx              context.Context
		closed           bool
		httpClient       *http.Client
	}

	callback struct {
		Func   SSEMessageFunc
		Result any
	}
)

func NewSSESource() *SSESource { _ = "STUB: not implemented"; return nil }

func (sse *SSESource) SetURL(url string) *SSESource { _ = "STUB: not implemented"; return nil }

func (sse *SSESource) SetMethod(method string) *SSESource { _ = "STUB: not implemented"; return nil }

func (sse *SSESource) SetHeader(header, value string) *SSESource {
	_ = "STUB: not implemented"
	return nil
}

func (sse *SSESource) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (sse *SSESource) SetContext(ctx context.Context) *SSESource {
	_ = "STUB: not implemented"
	return nil
}

func (sse *SSESource) SetBody(body io.Reader) *SSESource { _ = "STUB: not implemented"; return nil }

func (sse *SSESource) TLSClientConfig() *tls.Config { _ = "STUB: not implemented"; return nil }

func (sse *SSESource) SetTLSClientConfig(tlsConfig *tls.Config) *SSESource {
	_ = "STUB: not implemented"
	return nil
}

func (sse *SSESource) tlsConfig() (*tls.Config, error) { _ = "STUB: not implemented"; return nil, nil }

func (sse *SSESource) SetTransport(transport http.RoundTripper) *SSESource {
	_ = "STUB: not implemented"
	return nil
}

func (sse *SSESource) AddHeader(header, value string) *SSESource {
	_ = "STUB: not implemented"
	return nil
}

func (sse *SSESource) SetRetryCount(count int) *SSESource { _ = "STUB: not implemented"; return nil }

func (sse *SSESource) SetRetryWaitTime(waitTime time.Duration) *SSESource {
	_ = "STUB: not implemented"
	return nil
}

func (sse *SSESource) SetRetryMaxWaitTime(maxWaitTime time.Duration) *SSESource {
	_ = "STUB: not implemented"
	return nil
}

func (sse *SSESource) SetSizeMaxBuffer(bufSize int) *SSESource {
	_ = "STUB: not implemented"
	return nil
}

func (sse *SSESource) Logger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (sse *SSESource) SetLogger(l Logger) *SSESource { _ = "STUB: not implemented"; return nil }

func (sse *SSESource) outputLogTo(w io.Writer) *SSESource { _ = "STUB: not implemented"; return nil }

func (sse *SSESource) OnOpen(ef SSEOpenFunc) *SSESource { _ = "STUB: not implemented"; return nil }

func (sse *SSESource) OnError(ef SSEErrorFunc) *SSESource { _ = "STUB: not implemented"; return nil }

func (sse *SSESource) OnRequestFailure(ef SSERequestFailureFunc) *SSESource {
	_ = "STUB: not implemented"
	return nil
}

func (sse *SSESource) OnMessage(ef SSEMessageFunc, result any) *SSESource {
	_ = "STUB: not implemented"
	return nil
}

func (sse *SSESource) AddEventListener(eventName string, ef SSEMessageFunc, result any) *SSESource {
	_ = "STUB: not implemented"
	return nil
}

func (sse *SSESource) Get() error { _ = "STUB: not implemented"; return nil }

func (sse *SSESource) Close() { _ = "STUB: not implemented"; return }

func (sse *SSESource) enableConnect() { _ = "STUB: not implemented"; return }

func (sse *SSESource) isClosed() bool { _ = "STUB: not implemented"; return false }

func (sse *SSESource) triggerOnOpen(hdr http.Header) { _ = "STUB: not implemented"; return }

func (sse *SSESource) triggerOnError(err error) { _ = "STUB: not implemented"; return }

func (sse *SSESource) triggerOnRequestFailure(err error, res *http.Response) {
	_ = "STUB: not implemented"
	return
}

func (sse *SSESource) createRequest() (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sse *SSESource) connect() (*http.Response, error) { _ = "STUB: not implemented"; return nil, nil }

func (sse *SSESource) listenStream(res *http.Response) error { _ = "STUB: not implemented"; return nil }

func (sse *SSESource) processEvent(scanner *bufio.Scanner) error {
	_ = "STUB: not implemented"
	return nil
}

func (sse *SSESource) handleCallback(e *SSE) { _ = "STUB: not implemented"; return }

var readEvent = readEventFunc

func readEventFunc(scanner *bufio.Scanner) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func wrapResponse(res *http.Response, req *http.Request) *Response {
	_ = "STUB: not implemented"
	return nil
}

type rawSSE struct {
	ID    []byte
	Data  []byte
	Event []byte
	Retry []byte
}

var parseEvent = parseEventFunc

func parseEventFunc(msg []byte) (*rawSSE, error) { _ = "STUB: not implemented"; return nil, nil }

func trimHeader(size int, data []byte) []byte { _ = "STUB: not implemented"; return nil }

var rawEventPool = &sync.Pool{New: func() any { return new(rawSSE) }}

func newRawEvent() *rawSSE { _ = "STUB: not implemented"; return nil }

func putRawEvent(e *rawSSE) { _ = "STUB: not implemented"; return }
