package resty

import (
	"net/http"
)

type (
	RedirectPolicy interface {
		Apply(*http.Request, []*http.Request) error
	}

	RedirectPolicyFunc func(*http.Request, []*http.Request) error

	RedirectInfo struct {
		URL string

		StatusCode int
	}
)

func (f RedirectPolicyFunc) Apply(req *http.Request, via []*http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func RedirectNoPolicy() RedirectPolicy { _ = "STUB: not implemented"; return *new(RedirectPolicy) }

func RedirectFlexiblePolicy(noOfRedirect int) RedirectPolicy {
	_ = "STUB: not implemented"
	return *new(RedirectPolicy)
}

func RedirectDomainCheckPolicy(hostnames ...string) RedirectPolicy {
	_ = "STUB: not implemented"
	return *new(RedirectPolicy)
}

func RedirectHeaderStripSensitivePolicy(applyDefault bool, headers ...string) RedirectPolicy {
	_ = "STUB: not implemented"
	return *new(RedirectPolicy)
}

func checkHostAndAddHeaders(cur *http.Request, pre *http.Request) {
	_ = "STUB: not implemented"
	return
}
