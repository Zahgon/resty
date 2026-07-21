package resty

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrRateLimitExceeded = errors.New("resty: rate limit exceeded")

type RateLimiter interface {
	Allow(ctx context.Context) error
}

func NewRateLimitTokenBucket(requestsPerSecond float64, burst int) *RateLimitTokenBucket {
	_ = "STUB: not implemented"
	return nil
}

var _ RateLimiter = (*RateLimitTokenBucket)(nil)

type RateLimitTokenBucket struct {
	mu         sync.Mutex
	rate       float64
	burst      int
	tokens     float64
	lastRefill time.Time
}

func (l *RateLimitTokenBucket) Rate() float64 { _ = "STUB: not implemented"; return 0 }

func (l *RateLimitTokenBucket) Burst() int { _ = "STUB: not implemented"; return 0 }

func (l *RateLimitTokenBucket) Allow(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *RateLimitTokenBucket) refill() { _ = "STUB: not implemented"; return }

func NewRateLimitSlidingWindow(limit int, windowSize time.Duration) *RateLimitSlidingWindow {
	_ = "STUB: not implemented"
	return nil
}

var _ RateLimiter = (*RateLimitSlidingWindow)(nil)

type RateLimitSlidingWindow struct {
	mu         sync.Mutex
	limit      int
	windowSize time.Duration
	timestamps []time.Time
}

func (l *RateLimitSlidingWindow) Limit() int { _ = "STUB: not implemented"; return 0 }

func (l *RateLimitSlidingWindow) WindowSize() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (l *RateLimitSlidingWindow) Allow(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
