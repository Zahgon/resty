package resty

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"
)

var ErrNoBaseURLs = errors.New("resty: no base URLs found")

type LoadBalancer interface {
	NextWithContext(ctx context.Context) (string, error)
	Feedback(*RequestFeedback)
	Close() error
}

type RequestFeedback struct {
	BaseURL string
	Success bool
	Attempt int
}

func NewRoundRobin(baseURLs ...string) (*RoundRobin, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ LoadBalancer = (*RoundRobin)(nil)

type RoundRobin struct {
	lock     *sync.Mutex
	baseURLs []string
	current  int
}

func (rr *RoundRobin) NextWithContext(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (rr *RoundRobin) Feedback(_ *RequestFeedback) { _ = "STUB: not implemented"; return }

func (rr *RoundRobin) Close() error { _ = "STUB: not implemented"; return nil }

func (rr *RoundRobin) Refresh(baseURLs ...string) error { _ = "STUB: not implemented"; return nil }

type Host struct {
	BaseURL string

	Weight int

	MaxFailures int

	state          HostState
	currentWeight  int
	failedRequests int
}

func (h *Host) addWeight() { _ = "STUB: not implemented"; return }

func (h *Host) resetWeight(totalWeight int) { _ = "STUB: not implemented"; return }

type HostState int

const (
	HostStateInActive HostState = iota
	HostStateActive
)

type HostStateChangeFunc func(baseURL string, from, to HostState)

var ErrNoActiveHost = errors.New("resty: no active host")

func NewWeightedRoundRobin(recovery time.Duration, hosts ...*Host) (*WeightedRoundRobin, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ LoadBalancer = (*WeightedRoundRobin)(nil)

type WeightedRoundRobin struct {
	lock          *sync.RWMutex
	hosts         []*Host
	totalWeight   int
	tick          *time.Ticker
	onStateChange HostStateChangeFunc

	recovery time.Duration
}

func (wrr *WeightedRoundRobin) NextWithContext(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (wrr *WeightedRoundRobin) Feedback(f *RequestFeedback) { _ = "STUB: not implemented"; return }

func (wrr *WeightedRoundRobin) Close() error { _ = "STUB: not implemented"; return nil }

func (wrr *WeightedRoundRobin) Refresh(hosts ...*Host) error { _ = "STUB: not implemented"; return nil }

func (wrr *WeightedRoundRobin) SetOnStateChange(fn HostStateChangeFunc) {
	_ = "STUB: not implemented"
	return
}

func (wrr *WeightedRoundRobin) SetRecoveryDuration(d time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (wrr *WeightedRoundRobin) ticker() { _ = "STUB: not implemented"; return }

func NewSRVWeightedRoundRobin(service, proto, domainName, httpScheme string) (*SRVWeightedRoundRobin, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ LoadBalancer = (*SRVWeightedRoundRobin)(nil)

type SRVWeightedRoundRobin struct {
	Service    string
	Proto      string
	DomainName string
	HttpScheme string

	wrr       *WeightedRoundRobin
	tick      *time.Ticker
	lock      *sync.Mutex
	lookupSRV func() ([]*net.SRV, error)
}

func (swrr *SRVWeightedRoundRobin) NextWithContext(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (swrr *SRVWeightedRoundRobin) Feedback(f *RequestFeedback) { _ = "STUB: not implemented"; return }

func (swrr *SRVWeightedRoundRobin) Close() error { _ = "STUB: not implemented"; return nil }

func (swrr *SRVWeightedRoundRobin) Refresh() error { _ = "STUB: not implemented"; return nil }

func (swrr *SRVWeightedRoundRobin) SetRefreshDuration(d time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (swrr *SRVWeightedRoundRobin) SetOnStateChange(fn HostStateChangeFunc) {
	_ = "STUB: not implemented"
	return
}

func (swrr *SRVWeightedRoundRobin) SetRecoveryDuration(d time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (swrr *SRVWeightedRoundRobin) ticker() { _ = "STUB: not implemented"; return }

func extractBaseURL(u string) (string, error) { _ = "STUB: not implemented"; return "", nil }
