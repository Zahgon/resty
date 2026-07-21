package resty

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Request struct {
	CorrelationID string

	URL                          string
	Method                       string
	AuthToken                    string
	AuthScheme                   string
	QueryParams                  url.Values
	FormData                     url.Values
	PathParams                   map[string]string
	Header                       http.Header
	StartTime                    time.Time
	Body                         any
	Result                       any
	ResultError                  any
	RawRequest                   *http.Request
	Cookies                      []*http.Cookie
	IsDebug                      bool
	IsCloseConnection            bool
	IsResponseDoNotParse         bool
	ResponseSaveFileName         string
	ResponseExpectContentType    string
	ResponseForceContentType     string
	DebugBodyLimit               int
	ResponseBodyLimit            int64
	IsResponseBodyUnlimitedReads bool
	IsTrace                      bool
	IsMethodGetAllowPayload      bool
	IsMethodDeleteAllowPayload   bool
	IsDone                       bool
	IsResponseSaveToFile         bool
	Timeout                      time.Duration
	HeaderAuthorizationKey       string
	RetryCount                   int
	RetryWaitTime                time.Duration
	RetryMaxWaitTime             time.Duration
	RetryDelayStrategy           RetryDelayStrategyFunc
	IsRetryDefaultConditions     bool
	IsRetryAllowNonIdempotent    bool
	Label                        string

	Attempt int

	mu                   *sync.Mutex
	credentials          *credentials
	isMultiPart          bool
	isFormData           bool
	isContentLengthSet   bool
	contentLength        int64
	jsonEscapeHTML       bool
	ctx                  context.Context
	ctxCancelFunc        context.CancelFunc
	values               map[string]any
	client               *Client
	bodyBuf              *bytes.Buffer
	trace                *clientTrace
	log                  Logger
	baseURL              string
	multipartBoundary    string
	multipartFields      []*MultipartField
	retryConditions      []RetryConditionFunc
	isSetRetryConditions bool
	retryHooks           []RetryHookFunc
	isSetRetryHooks      bool
	curlCmdString        string
	isCurlCmdGenerate    bool
	isCurlCmdDebugLog    bool
	unescapeQueryParams  bool
	multipartErrChan     chan error
	multipartCancelFunc  context.CancelFunc
}

func (r *Request) SetCorrelationID(id string) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetMethod(m string) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetURL(url string) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (r *Request) SetContext(ctx context.Context) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) WithContext(ctx context.Context) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetContentType(ct string) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetHeader(header, value string) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetHeaderAny(header string, value any) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetHeaders(headers map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetHeaderMultiValues(headers map[string][]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetHeaderVerbatim(header, value string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetHeaderVerbatimAny(header string, value any) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetQueryParam(param, value string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetQueryParamAny(param string, value any) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetQueryParams(params map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetQueryParamsFromValues(params url.Values) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetQueryString(query string) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetFormData(data map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetFormDataFromValues(data url.Values) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetBody(body any) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetResult(v any) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetResultError(err any) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetFile(fieldName, filePath string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetFiles(files map[string]string) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetFileReader(fieldName, fileName string, reader io.Reader) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetMultipartFormData(data map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetMultipartOrderedFormData(name string, values []string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetMultipartField(fieldName, fileName, contentType string, reader io.Reader) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetMultipartFields(fields ...*MultipartField) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetMultipartBoundary(boundary string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetContentLength(v int64) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetBasicAuth(username, password string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetAuthToken(authToken string) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetAuthScheme(scheme string) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetHeaderAuthorizationKey(k string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetResponseSaveFileName(file string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetResponseSaveToFile(save bool) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetCloseConnection(close bool) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetResponseDoNotParse(notParse bool) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetResponseBodyLimit(v int64) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetResponseBodyUnlimitedReads(b bool) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetPathParam(param, value string) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetPathParamAny(param string, value any) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetPathParams(params map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetPathRawParam(param, value string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetPathRawParamAny(param string, value any) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetPathRawParams(params map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetResponseExpectContentType(contentType string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetResponseForceContentType(contentType string) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetJSONEscapeHTML(b bool) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetCookie(hc *http.Cookie) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetCookies(rs []*http.Cookie) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetTimeout(timeout time.Duration) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetLogger(l Logger) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetDebug(d bool) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) AddRetryConditions(conditions ...RetryConditionFunc) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetRetryConditions(conditions ...RetryConditionFunc) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) AddRetryHooks(hooks ...RetryHookFunc) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetRetryHooks(hooks ...RetryHookFunc) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetRetryCount(count int) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetRetryWaitTime(waitTime time.Duration) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetRetryMaxWaitTime(maxWaitTime time.Duration) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetRetryDelayStrategy(rs RetryDelayStrategyFunc) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetRetryDefaultConditions(b bool) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetRetryAllowNonIdempotent(b bool) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetTrace(t bool) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetCurlCmdGenerate(b bool) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) SetCurlCmdDebugLog(b bool) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) CurlCmd() string { _ = "STUB: not implemented"; return "" }

func (r *Request) generateCurlCommand() string { _ = "STUB: not implemented"; return "" }

func (r *Request) SetQueryParamsUnescape(unescape bool) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetMethodGetAllowPayload(allow bool) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetMethodDeleteAllowPayload(allow bool) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) SetLabel(label string) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) TraceInfo() TraceInfo { _ = "STUB: not implemented"; return *new(TraceInfo) }

func (r *Request) Get(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Request) Head(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Request) Post(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Request) Put(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Request) Patch(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Request) Delete(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Request) Options(url string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Request) Trace(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Request) Send() (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Request) Execute(method, url string) (res *Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Request) Clone(ctx context.Context) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) Funcs(funcs ...RequestFunc) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) fmtBodyString(sl int) (body string) { _ = "STUB: not implemented"; return "" }

func (r *Request) initValuesMap() { _ = "STUB: not implemented"; return }

func (r *Request) initTraceIfEnabled() { _ = "STUB: not implemented"; return }

func (r *Request) isHeaderExists(k string) bool { _ = "STUB: not implemented"; return false }

func (r *Request) isPayloadSupported() bool { _ = "STUB: not implemented"; return false }

func (r *Request) sendLoadBalancerFeedback(res *Response, err error) {
	_ = "STUB: not implemented"
	return
}

func (r *Request) resetFileReaders() error { _ = "STUB: not implemented"; return nil }

var idempotentMethods = map[string]struct{}{
	MethodDelete:  {},
	MethodGet:     {},
	MethodHead:    {},
	MethodOptions: {},
	MethodPut:     {},
	MethodTrace:   {},
}

func (r *Request) isIdempotent() bool { _ = "STUB: not implemented"; return false }

func (r *Request) withTimeout() *http.Request { _ = "STUB: not implemented"; return nil }

func jsonIndent(v []byte) []byte { _ = "STUB: not implemented"; return nil }
