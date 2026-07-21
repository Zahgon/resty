//go:build (js && wasm) || wasip1
// +build js,wasm wasip1

package resty

import (
	"context"
	"net"
)

func transportDialContext(_ *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return nil
}
