package resty

import (
	"net/http"
	"sync"
	"time"
)

type Hedger interface {
	http.RoundTripper

	SetTransport(http.RoundTripper)

	Transport() http.RoundTripper
}

var _ Hedger = (*Hedging)(nil)

func NewHedging() *Hedging { _ = "STUB: not implemented"; return nil }

type Hedging struct {
	lock                 *sync.RWMutex
	underlying           http.RoundTripper
	delay                time.Duration
	maxRequest           int
	maxRequestPerSecond  float64
	rateDelay            time.Duration
	isNonReadOnlyAllowed bool
}

func (h *Hedging) SetTransport(t http.RoundTripper) { _ = "STUB: not implemented"; return }

func (h *Hedging) Transport() http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

func (h *Hedging) Delay() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (h *Hedging) SetDelay(delay time.Duration) *Hedging { _ = "STUB: not implemented"; return nil }

func (h *Hedging) MaxRequest() int { _ = "STUB: not implemented"; return 0 }

func (h *Hedging) SetMaxRequest(count int) *Hedging { _ = "STUB: not implemented"; return nil }

func (h *Hedging) MaxRequestPerSecond() float64 { _ = "STUB: not implemented"; return 0 }

func (h *Hedging) SetMaxRequestPerSecond(count float64) *Hedging {
	_ = "STUB: not implemented"
	return nil
}

func (h *Hedging) IsNonReadOnlyAllowed() bool { _ = "STUB: not implemented"; return false }

func (h *Hedging) SetNonReadOnlyAllowed(allow bool) *Hedging { _ = "STUB: not implemented"; return nil }

func (h *Hedging) calculateRateDelay() { _ = "STUB: not implemented"; return }

func (ht *Hedging) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isReadOnlyMethod(method string) bool { _ = "STUB: not implemented"; return false }
