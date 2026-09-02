//go:build !(js && wasm)
// +build !js !wasm

package resty

import (
	"context"
	"net"
)

func transportDialContext(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return nil
}
