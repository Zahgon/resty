package resty

import (
	"context"
	"net/http/httptrace"
	"sync"
	"time"
)

type TraceInfo struct {
	DNSLookup time.Duration `json:"dns_lookup_time"`

	ConnTime time.Duration `json:"connection_time"`

	TCPConnTime time.Duration `json:"tcp_connection_time"`

	TLSHandshake time.Duration `json:"tls_handshake_time"`

	ServerTime time.Duration `json:"server_time"`

	ResponseTime time.Duration `json:"response_time"`

	TotalTime time.Duration `json:"total_time"`

	IsConnReused bool `json:"is_connection_reused"`

	IsConnWasIdle bool `json:"is_connection_was_idle"`

	ConnIdleTime time.Duration `json:"connection_idle_time"`

	RequestAttempt int `json:"request_attempt"`

	RemoteAddr string `json:"remote_address"`
}

func (ti TraceInfo) String() string { _ = "STUB: not implemented"; return "" }

func (ti TraceInfo) JSON() string { _ = "STUB: not implemented"; return "" }

func (ti TraceInfo) Clone() *TraceInfo { _ = "STUB: not implemented"; return nil }

type clientTrace struct {
	lock                 sync.RWMutex
	getConn              time.Time
	dnsStart             time.Time
	dnsDone              time.Time
	connectDone          time.Time
	tlsHandshakeStart    time.Time
	tlsHandshakeDone     time.Time
	gotConn              time.Time
	gotFirstResponseByte time.Time
	endTime              time.Time
	gotConnInfo          httptrace.GotConnInfo
}

func (t *clientTrace) createContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
