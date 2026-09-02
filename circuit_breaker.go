package resty

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var ErrCircuitBreakerOpen = errors.New("resty: circuit breaker open")

const (
	CircuitBreakerStateClosed CircuitBreakerState = iota

	CircuitBreakerStateOpen

	CircuitBreakerStateHalfOpen
)

type (
	CircuitBreaker interface {
		Allow() error

		ApplyPolicies(*Response)
	}

	CircuitBreakerObserver interface {
		OnTrigger(...CircuitBreakerTriggerHook) CircuitBreakerObserver

		RunOnTriggerHooks(*Request, error)

		OnStateChange(...CircuitBreakerStateChangeHook) CircuitBreakerObserver

		RunOnStateChangeHooks(oldState, newState CircuitBreakerState)
	}

	CircuitBreakerTriggerHook func(*Request, error)

	CircuitBreakerStateChangeHook func(oldState, newState CircuitBreakerState)

	CircuitBreakerState uint32

	circuitBreakerMode interface {
		shouldOpenOnClosed(totalAndFailures) bool
		halfOpenSuccessThreshold() uint64
	}
)

var _ CircuitBreakerObserver = (*CircuitBreakerCount)(nil)
var _ CircuitBreaker = (*CircuitBreakerCount)(nil)

type CircuitBreakerCount struct {
	*circuitBreakerBase
	failureThreshold uint64
	successThreshold uint64
}

func (cb *CircuitBreakerCount) shouldOpenOnClosed(tf totalAndFailures) bool {
	_ = "STUB: not implemented"
	return false
}

func (cb *CircuitBreakerCount) halfOpenSuccessThreshold() uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (cb *CircuitBreakerCount) ApplyPolicies(resp *Response) { _ = "STUB: not implemented"; return }

var _ CircuitBreakerObserver = (*CircuitBreakerRatio)(nil)
var _ CircuitBreaker = (*CircuitBreakerRatio)(nil)

type CircuitBreakerRatio struct {
	*circuitBreakerBase
	failureRatio float64
	minRequests  uint64
}

func (cb *CircuitBreakerRatio) shouldOpenOnClosed(tf totalAndFailures) bool {
	_ = "STUB: not implemented"
	return false
}

func (cb *CircuitBreakerRatio) halfOpenSuccessThreshold() uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (cb *CircuitBreakerRatio) ApplyPolicies(resp *Response) { _ = "STUB: not implemented"; return }

type group[T any] interface {
	op(T) T
	empty() T
	inverse() T
}

type totalAndFailures struct {
	total    int
	failures int
}

func (tf totalAndFailures) op(g totalAndFailures) totalAndFailures {
	_ = "STUB: not implemented"
	return *new(totalAndFailures)
}

func (tf totalAndFailures) empty() totalAndFailures {
	_ = "STUB: not implemented"
	return *new(totalAndFailures)
}

func (tf totalAndFailures) inverse() totalAndFailures {
	_ = "STUB: not implemented"
	return *new(totalAndFailures)
}

type slidingWindow[G group[G]] struct {
	mutex     sync.RWMutex
	total     G
	values    []G
	idx       int
	lastStart time.Time
	interval  time.Duration
}

func newSlidingWindow[G group[G]](interval time.Duration, buckets int) *slidingWindow[G] {
	_ = "STUB: not implemented"
	return nil
}

func (sw *slidingWindow[G]) Add(val G) { _ = "STUB: not implemented"; return }

func (sw *slidingWindow[G]) AddAndGet(val G) G { _ = "STUB: not implemented"; return *new(G) }

func (sw *slidingWindow[G]) Get() G { _ = "STUB: not implemented"; return *new(G) }

func (sw *slidingWindow[G]) SetInterval(interval time.Duration) { _ = "STUB: not implemented"; return }

type cbRequestErrorObserver interface {
	onRequestError()
}

var _ cbRequestErrorObserver = (*circuitBreakerBase)(nil)

type circuitBreakerBase struct {
	lock                  sync.RWMutex
	policies              []CircuitBreakerPolicy
	resetTimeout          time.Duration
	resetTimerMu          sync.Mutex
	resetTimer            *time.Timer
	resetDeadlineUnixNano atomic.Int64
	state                 atomic.Value
	halfOpenProbeInFlight atomic.Uint32
	sw                    atomic.Pointer[slidingWindow[totalAndFailures]]

	triggerHooks     []CircuitBreakerTriggerHook
	stateChangeHooks []CircuitBreakerStateChangeHook
}

func NewCircuitBreakerCount(failureThreshold uint64, successThreshold uint64,
	resetTimeout time.Duration, policies ...CircuitBreakerPolicy) *CircuitBreakerCount {
	_ = "STUB: not implemented"
	return nil
}

func NewCircuitBreakerRatio(failureRatio float64, minRequests uint64,
	resetTimeout time.Duration, policies ...CircuitBreakerPolicy) *CircuitBreakerRatio {
	_ = "STUB: not implemented"
	return nil
}

func newCircuitBreakerBase(resetTimeout time.Duration, policies ...CircuitBreakerPolicy) *circuitBreakerBase {
	_ = "STUB: not implemented"
	return nil
}

func (cb *circuitBreakerBase) Allow() error { _ = "STUB: not implemented"; return nil }

func (cb *circuitBreakerBase) applyPolicies(resp *Response, mode circuitBreakerMode) {
	_ = "STUB: not implemented"
	return
}

func (cb *circuitBreakerBase) OnTrigger(hooks ...CircuitBreakerTriggerHook) CircuitBreakerObserver {
	_ = "STUB: not implemented"
	return *new(CircuitBreakerObserver)
}

func (cb *circuitBreakerBase) RunOnTriggerHooks(req *Request, err error) {
	_ = "STUB: not implemented"
	return
}

func (cb *circuitBreakerBase) OnStateChange(hooks ...CircuitBreakerStateChangeHook) CircuitBreakerObserver {
	_ = "STUB: not implemented"
	return *new(CircuitBreakerObserver)
}

func (cb *circuitBreakerBase) RunOnStateChangeHooks(oldState, newState CircuitBreakerState) {
	_ = "STUB: not implemented"
	return
}

type CircuitBreakerPolicy func(resp *Response) bool

func CircuitBreaker5xxPolicy(resp *Response) bool { _ = "STUB: not implemented"; return false }

func (cb *circuitBreakerBase) getState() CircuitBreakerState {
	_ = "STUB: not implemented"
	return *new(CircuitBreakerState)
}

func (cb *circuitBreakerBase) open() { _ = "STUB: not implemented"; return }

func (cb *circuitBreakerBase) onResetTimeout() { _ = "STUB: not implemented"; return }

func (cb *circuitBreakerBase) changeState(state CircuitBreakerState) {
	_ = "STUB: not implemented"
	return
}

func (cb *circuitBreakerBase) onRequestError() { _ = "STUB: not implemented"; return }
