package resty

import (
	"math/rand/v2"
	"regexp"
	"time"
)

const (
	defaultWaitTime    = time.Duration(100) * time.Millisecond
	defaultMaxWaitTime = time.Duration(2000) * time.Millisecond
)

type (
	RetryConditionFunc func(*Response, error) bool

	RetryHookFunc func(*Response, error)

	RetryDelayStrategyFunc func(*Response, error) (time.Duration, error)
)

func RetryConstantDelayStrategy(delay time.Duration) RetryDelayStrategyFunc {
	_ = "STUB: not implemented"
	return *new(RetryDelayStrategyFunc)
}

var (
	regexErrTooManyRedirects = regexp.MustCompile(`stopped after \d+ redirects\z`)
	regexErrScheme           = regexp.MustCompile("unsupported protocol scheme")
	regexErrInvalidHeader    = regexp.MustCompile("invalid header")
)

func RetryConditionStatusTooManyRequests(res *Response, _ error) bool {
	_ = "STUB: not implemented"
	return false
}

func RetryConditionStatus5XX(res *Response, _ error) bool { _ = "STUB: not implemented"; return false }

func RetryConditionStatusZero(res *Response, _ error) bool { _ = "STUB: not implemented"; return false }

func isDoNotRetryError(err error) bool { _ = "STUB: not implemented"; return false }

func newBackoffWithJitter(min, max time.Duration) *backoffWithJitter {
	_ = "STUB: not implemented"
	return nil
}

type backoffWithJitter struct {
	rnd *rand.Rand
	min time.Duration
	max time.Duration
}

func (b *backoffWithJitter) NextWaitDuration(c *Client, res *Response, err error, attempt int) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (b *backoffWithJitter) defaultDelayStrategy(attempt int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (b *backoffWithJitter) randDuration(center time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (b *backoffWithJitter) balanceMinMax(delay time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

var timeNow = time.Now

func parseRetryAfterHeader(v string) (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}
