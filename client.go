package resty

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"sync"
	"time"
)

const (
	MethodGet = "GET"

	MethodPost = "POST"

	MethodPut = "PUT"

	MethodDelete = "DELETE"

	MethodPatch = "PATCH"

	MethodHead = "HEAD"

	MethodOptions = "OPTIONS"

	MethodTrace = "TRACE"
)

const (
	defaultWatcherPoolingInterval = 24 * time.Hour
)

var (
	ErrNotHttpTransportType = errors.New("resty: not a http.Transport type")

	ErrUnsupportedRequestBodyKind = errors.New("resty: unsupported request body kind")

	ErrReaderNotSeekable = errors.New("resty: reader is not seekable on request retry")

	hdrUserAgentKey       = http.CanonicalHeaderKey("User-Agent")
	hdrAcceptKey          = http.CanonicalHeaderKey("Accept")
	hdrAcceptEncodingKey  = http.CanonicalHeaderKey("Accept-Encoding")
	hdrContentTypeKey     = http.CanonicalHeaderKey("Content-Type")
	hdrContentLengthKey   = http.CanonicalHeaderKey("Content-Length")
	hdrContentEncodingKey = http.CanonicalHeaderKey("Content-Encoding")
	hdrContentDisposition = http.CanonicalHeaderKey("Content-Disposition")
	hdrAuthorizationKey   = http.CanonicalHeaderKey("Authorization")
	hdrWwwAuthenticateKey = http.CanonicalHeaderKey("WWW-Authenticate")
	hdrRetryAfterKey      = http.CanonicalHeaderKey("Retry-After")
	hdrCookieKey          = http.CanonicalHeaderKey("Cookie")

	plainTextType   = "text/plain; charset=utf-8"
	jsonContentType = "application/json"
	formContentType = "application/x-www-form-urlencoded"

	jsonKey = "json"
	xmlKey  = "xml"

	defaultAuthScheme = "Bearer"

	hdrUserAgentValue = "go-resty/" + Version + " (https://resty.dev)"
	bufPool           = &sync.Pool{New: func() any { return &bytes.Buffer{} }}
)

type (
	RequestMiddleware func(*Client, *Request) error

	ResponseMiddleware func(*Client, *Response) error

	ErrorHook func(*Request, error)

	SuccessHook func(*Client, *Response)

	CloseHook func()

	RequestFunc func(*Request) *Request

	TLSClientConfiger interface {
		TLSClientConfig() *tls.Config
		SetTLSClientConfig(*tls.Config) error
	}
)

type TransportSettings struct {
	DialerTimeout time.Duration

	DialerKeepAlive time.Duration

	IdleConnTimeout time.Duration

	TLSHandshakeTimeout time.Duration

	ExpectContinueTimeout time.Duration

	ResponseHeaderTimeout time.Duration

	MaxIdleConns int

	MaxIdleConnsPerHost int

	MaxConnsPerHost int

	DisableKeepAlives bool

	MaxResponseHeaderBytes int64

	WriteBufferSize int

	ReadBufferSize int
}

type Client struct {
	lock                       *sync.RWMutex
	baseURL                    string
	queryParams                url.Values
	formData                   url.Values
	pathParams                 map[string]string
	header                     http.Header
	credentials                *credentials
	authToken                  string
	authScheme                 string
	cookies                    []*http.Cookie
	errorType                  reflect.Type
	debug                      bool
	disableWarn                bool
	isMethodGetAllowPayload    bool
	isMethodDeleteAllowPayload bool
	timeout                    time.Duration
	retryCount                 int
	retryWaitTime              time.Duration
	retryMaxWaitTime           time.Duration
	retryConditions            []RetryConditionFunc
	retryHooks                 []RetryHookFunc
	retryDelayStrategy         RetryDelayStrategyFunc
	isRetryDefaultConditions   bool
	isRetryAllowNonIdempotent  bool
	headerAuthorizationKey     string
	responseBodyLimit          int64
	resBodyUnlimitedReads      bool
	jsonEscapeHTML             bool
	closeConnection            bool
	isResponseDoNotParse       bool
	isTrace                    bool
	debugBodyLimit             int
	responseSaveDirectory      string
	isResponseSaveToFile       bool
	scheme                     string
	log                        Logger
	ctx                        context.Context
	httpClient                 *http.Client
	proxyURL                   *url.URL
	debugLogFormatter          DebugLogFormatterFunc
	debugLogCallback           DebugLogCallbackFunc
	isCurlCmdGenerate          bool
	isCurlCmdDebugLog          bool
	unescapeQueryParams        bool
	loadBalancer               LoadBalancer
	beforeRequest              []RequestMiddleware
	afterResponse              []ResponseMiddleware
	errorHooks                 []ErrorHook
	invalidHooks               []ErrorHook
	panicHooks                 []ErrorHook
	successHooks               []SuccessHook
	closeHooks                 []CloseHook
	contentTypeEncoders        map[string]ContentTypeEncoder
	contentTypeDecoders        map[string]ContentTypeDecoder
	contentDecompresserKeys    []string
	contentDecompressers       map[string]ContentDecompresser
	certWatcherStopChan        chan bool
	isClosed                   bool
	circuitBreaker             CircuitBreaker
	hedging                    Hedger
	rateLimiter                RateLimiter
}

type CertWatcherOptions struct {
	PoolInterval time.Duration
}

func (c *Client) BaseURL() string { _ = "STUB: not implemented"; return "" }

func (c *Client) SetBaseURL(url string) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) LoadBalancer() LoadBalancer { _ = "STUB: not implemented"; return *new(LoadBalancer) }

func (c *Client) SetLoadBalancer(b LoadBalancer) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (c *Client) SetHeader(header, value string) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SetHeaderAny(header string, value any) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetHeaders(headers map[string]string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetHeaderVerbatim(header, value string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetHeaderVerbatimAny(header string, value any) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Context() context.Context { _ = "STUB: not implemented"; return *new(context.Context) }

func (c *Client) SetContext(ctx context.Context) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) CookieJar() http.CookieJar { _ = "STUB: not implemented"; return *new(http.CookieJar) }

func (c *Client) SetCookieJar(jar http.CookieJar) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) Cookies() []*http.Cookie { _ = "STUB: not implemented"; return nil }

func (c *Client) SetCookie(hc *http.Cookie) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SetCookies(cs []*http.Cookie) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) QueryParams() url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

func (c *Client) SetQueryParam(param, value string) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SetQueryParamAny(param string, value any) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetQueryParams(params map[string]string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) FormData() url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

func (c *Client) SetFormData(data map[string]string) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SetBasicAuth(username, password string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) AuthToken() string { _ = "STUB: not implemented"; return "" }

func (c *Client) HeaderAuthorizationKey() string { _ = "STUB: not implemented"; return "" }

func (c *Client) SetHeaderAuthorizationKey(k string) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SetAuthToken(token string) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) AuthScheme() string { _ = "STUB: not implemented"; return "" }

func (c *Client) SetAuthScheme(scheme string) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SetDigestAuth(username, password string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) R() *Request { _ = "STUB: not implemented"; return nil }

func (c *Client) NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (c *Client) AddRequestMiddleware(m RequestMiddleware) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetRequestMiddlewares(middlewares ...RequestMiddleware) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) requestMiddlewares() []RequestMiddleware { _ = "STUB: not implemented"; return nil }

func (c *Client) AddResponseMiddleware(m ResponseMiddleware) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetResponseMiddlewares(middlewares ...ResponseMiddleware) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) responseMiddlewares() []ResponseMiddleware { _ = "STUB: not implemented"; return nil }

func (c *Client) OnError(hooks ...ErrorHook) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) OnSuccess(hooks ...SuccessHook) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) OnInvalid(hooks ...ErrorHook) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) OnPanic(hooks ...ErrorHook) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) OnClose(hooks ...CloseHook) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) ContentTypeEncoders() map[string]ContentTypeEncoder {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) AddContentTypeEncoder(ct string, e ContentTypeEncoder) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) inferContentTypeEncoder(ct ...string) (ContentTypeEncoder, bool) {
	_ = "STUB: not implemented"
	return *new(ContentTypeEncoder), false
}

func (c *Client) ContentTypeDecoders() map[string]ContentTypeDecoder {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) AddContentTypeDecoder(ct string, d ContentTypeDecoder) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) inferContentTypeDecoder(ct ...string) (ContentTypeDecoder, bool) {
	_ = "STUB: not implemented"
	return *new(ContentTypeDecoder), false
}

func (c *Client) ContentDecompressers() map[string]ContentDecompresser {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) AddContentDecompresser(k string, d ContentDecompresser) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) ContentDecompresserKeys() string { _ = "STUB: not implemented"; return "" }

func (c *Client) SetContentDecompresserKeys(keys []string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetCircuitBreaker(cb CircuitBreaker) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) RateLimiter() RateLimiter { _ = "STUB: not implemented"; return *new(RateLimiter) }

func (c *Client) SetRateLimiter(l RateLimiter) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) IsDebug() bool { _ = "STUB: not implemented"; return false }

func (c *Client) SetDebug(d bool) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) DebugBodyLimit() int { _ = "STUB: not implemented"; return 0 }

func (c *Client) SetDebugBodyLimit(sl int) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) debugLogCallbackFunc() DebugLogCallbackFunc {
	_ = "STUB: not implemented"
	return *new(DebugLogCallbackFunc)
}

func (c *Client) OnDebugLog(dlc DebugLogCallbackFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) debugLogFormatterFunc() DebugLogFormatterFunc {
	_ = "STUB: not implemented"
	return *new(DebugLogFormatterFunc)
}

func (c *Client) SetDebugLogFormatter(df DebugLogFormatterFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) IsDisableWarn() bool { _ = "STUB: not implemented"; return false }

func (c *Client) SetLoggerWarnLevel(d bool) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) IsMethodGetAllowPayload() bool { _ = "STUB: not implemented"; return false }

func (c *Client) SetMethodGetAllowPayload(allow bool) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) IsMethodDeleteAllowPayload() bool { _ = "STUB: not implemented"; return false }

func (c *Client) SetMethodDeleteAllowPayload(allow bool) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Logger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (c *Client) SetLogger(l Logger) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) Timeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (c *Client) SetTimeout(timeout time.Duration) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) ResultError() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (c *Client) SetResultError(v any) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) newErrorInterface() any { _ = "STUB: not implemented"; return *new(any) }

func (c *Client) SetRedirectPolicy(policies ...RedirectPolicy) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) RetryCount() int { _ = "STUB: not implemented"; return 0 }

func (c *Client) SetRetryCount(count int) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) RetryWaitTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Client) SetRetryWaitTime(waitTime time.Duration) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) RetryMaxWaitTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Client) SetRetryMaxWaitTime(maxWaitTime time.Duration) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) RetryDelayStrategy() RetryDelayStrategyFunc {
	_ = "STUB: not implemented"
	return *new(RetryDelayStrategyFunc)
}

func (c *Client) SetRetryDelayStrategy(rs RetryDelayStrategyFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) IsRetryDefaultConditions() bool { _ = "STUB: not implemented"; return false }

func (c *Client) SetRetryDefaultConditions(b bool) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) IsRetryAllowNonIdempotent() bool { _ = "STUB: not implemented"; return false }

func (c *Client) SetRetryAllowNonIdempotent(b bool) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) RetryConditions() []RetryConditionFunc { _ = "STUB: not implemented"; return nil }

func (c *Client) AddRetryConditions(conditions ...RetryConditionFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) RetryHooks() []RetryHookFunc { _ = "STUB: not implemented"; return nil }

func (c *Client) AddRetryHooks(hooks ...RetryHookFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) isHedgingEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *Client) Hedging() Hedger { _ = "STUB: not implemented"; return *new(Hedger) }

func (c *Client) SetHedging(h Hedger) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) TLSClientConfig() *tls.Config { _ = "STUB: not implemented"; return nil }

func (c *Client) SetTLSClientConfig(tlsConfig *tls.Config) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) ProxyURL() *url.URL { _ = "STUB: not implemented"; return nil }

func (c *Client) SetProxy(proxyURL string) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) RemoveProxy() *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SetCertificateFromFile(certFilePath, certKeyFilePath string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetCertificateFromString(certStr, certKeyStr string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetCertificates(certs ...tls.Certificate) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetRootCertificates(pemFilePaths ...string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetRootCertificatesWatcher(options *CertWatcherOptions, pemFilePaths ...string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetRootCertificateFromString(pemCerts string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetClientRootCertificates(pemFilePaths ...string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetClientRootCertificatesWatcher(options *CertWatcherOptions, pemFilePaths ...string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetClientRootCertificateFromString(pemCerts string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) handleCAs(scope string, permCerts []byte) { _ = "STUB: not implemented"; return }

func (c *Client) initCertWatcher(pemFilePath, scope string, options *CertWatcherOptions) {
	_ = "STUB: not implemented"
	return
}

func (c *Client) ResponseSaveDirectory() string { _ = "STUB: not implemented"; return "" }

func (c *Client) SetResponseSaveDirectory(dirPath string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) IsResponseSaveToFile() bool { _ = "STUB: not implemented"; return false }

func (c *Client) SetResponseSaveToFile(save bool) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) HTTPTransport() (*http.Transport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Transport() http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

func (c *Client) SetTransport(transport http.RoundTripper) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Scheme() string { _ = "STUB: not implemented"; return "" }

func (c *Client) SetScheme(scheme string) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SetCloseConnection(close bool) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SetResponseDoNotParse(notParse bool) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) PathParams() map[string]string { _ = "STUB: not implemented"; return nil }

func (c *Client) SetPathParam(param, value string) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SetPathParamAny(param string, value any) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetPathParams(params map[string]string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetPathRawParam(param, value string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetPathRawParamAny(param string, value any) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetPathRawParams(params map[string]string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SetJSONEscapeHTML(b bool) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) ResponseBodyLimit() int64 { _ = "STUB: not implemented"; return 0 }

func (c *Client) SetResponseBodyLimit(v int64) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) IsTrace() bool { _ = "STUB: not implemented"; return false }

func (c *Client) SetTrace(t bool) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SetCurlCmdGenerate(b bool) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SetCurlCmdDebugLog(b bool) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SetQueryParamsUnescape(unescape bool) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) ResponseBodyUnlimitedReads() bool { _ = "STUB: not implemented"; return false }

func (c *Client) SetResponseBodyUnlimitedReads(b bool) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) IsProxySet() bool { _ = "STUB: not implemented"; return false }

func (c *Client) Client() *http.Client { _ = "STUB: not implemented"; return nil }

func (c *Client) Clone(ctx context.Context) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Client) executeRequestMiddlewares(req *Request) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) cbRequestError() { _ = "STUB: not implemented"; return }

func (c *Client) execute(req *Request) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) tlsConfig() (*tls.Config, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Client) outputLogTo(w io.Writer) *Client { _ = "STUB: not implemented"; return nil }

type ResponseError struct {
	Response *Response
	Err      error
}

func (e *ResponseError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ResponseError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (c *Client) onErrorHooks(req *Request, res *Response, err error) {
	_ = "STUB: not implemented"
	return
}

func (c *Client) onPanicHooks(req *Request, err error) { _ = "STUB: not implemented"; return }

func (c *Client) onInvalidHooks(req *Request, err error) { _ = "STUB: not implemented"; return }

func (c *Client) onCloseHooks() { _ = "STUB: not implemented"; return }

func (c *Client) debugf(format string, v ...any) { _ = "STUB: not implemented"; return }
