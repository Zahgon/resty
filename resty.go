package resty

import (
	"net"
	"net/http"
	"net/http/cookiejar"
)

const Version = "3.0.0-rc.3+devrc4"

func New() *Client { _ = "STUB: not implemented"; return nil }

func NewWithTransportSettings(transportSettings *TransportSettings) *Client {
	_ = "STUB: not implemented"
	return nil
}

func NewWithClient(hc *http.Client) *Client { _ = "STUB: not implemented"; return nil }

func NewWithDialer(dialer *net.Dialer) *Client { _ = "STUB: not implemented"; return nil }

func NewWithLocalAddr(localAddr net.Addr) *Client { _ = "STUB: not implemented"; return nil }

func NewWithDialerAndTransportSettings(dialer *net.Dialer, transportSettings *TransportSettings) *Client {
	_ = "STUB: not implemented"
	return nil
}

func createTransport(dialer *net.Dialer, transportSettings *TransportSettings) *http.Transport {
	_ = "STUB: not implemented"
	return nil
}

func createCookieJar() *cookiejar.Jar { _ = "STUB: not implemented"; return nil }

func createClient(hc *http.Client) *Client { _ = "STUB: not implemented"; return nil }
