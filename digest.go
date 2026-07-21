package resty

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"net/http"
)

var (
	ErrDigestBadChallenge = errors.New("resty: digest: challenge is bad")

	ErrDigestInvalidCharset = errors.New("resty: digest: invalid charset")

	ErrDigestAlgNotSupported = errors.New("resty: digest: algorithm is not supported")

	ErrDigestQopNotSupported = errors.New("resty: digest: qop is not supported")
)

var digestHashFuncs = map[string]func() hash.Hash{
	"":                 md5.New,
	"MD5":              md5.New,
	"MD5-sess":         md5.New,
	"SHA-256":          sha256.New,
	"SHA-256-sess":     sha256.New,
	"SHA-512":          sha512.New,
	"SHA-512-sess":     sha512.New,
	"SHA-512-256":      sha512.New512_256,
	"SHA-512-256-sess": sha512.New512_256,
}

const (
	qopAuth    = "auth"
	qopAuthInt = "auth-int"
)

type digestTransport struct {
	*credentials
	transport http.RoundTripper
}

func (dt *digestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dt *digestTransport) cloneReq(r *http.Request, first bool) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

func (dt *digestTransport) parseChallenge(input string) (*digestChallenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dt *digestTransport) createCredentials(cha *digestChallenge, req *http.Request) (*digestCredentials, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dt *digestTransport) prepareBody(req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

type digestChallenge struct {
	realm     string
	domain    string
	nonce     string
	opaque    string
	stale     string
	algorithm string
	qop       []string
	nc        int
	userHash  string
}

func (dc *digestChallenge) isQopSupported(qop string) bool { _ = "STUB: not implemented"; return false }

func (dc *digestChallenge) setValue(k, v string) error { _ = "STUB: not implemented"; return nil }

type digestCredentials struct {
	username      string
	password      string
	userHash      string
	method        string
	uri           string
	realm         string
	nonce         string
	algorithm     string
	sessAlgorithm bool
	cnonce        string
	opaque        string
	qop           string
	nc            int
	response      string
	bodyHash      string
}

func (dc *digestCredentials) parseQop(cha *digestChallenge) error {
	_ = "STUB: not implemented"
	return nil
}

func (dc *digestCredentials) h(data string) string { _ = "STUB: not implemented"; return "" }

func (dc *digestCredentials) digest(cha *digestChallenge) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (dc *digestCredentials) ha1() string { _ = "STUB: not implemented"; return "" }

func (dc *digestCredentials) ha2() string { _ = "STUB: not implemented"; return "" }

func (dc *digestCredentials) String() string { _ = "STUB: not implemented"; return "" }

func newHashFunc(algorithm string) hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }
