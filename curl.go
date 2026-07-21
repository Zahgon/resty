package resty

import (
	"net/http"
	"regexp"
)

const unexecutedRequestURL = "http://unexecuted-request"

func buildCurlCmd(req *Request) string { _ = "STUB: not implemented"; return "" }

func dumpCurlCookies(cookies []*http.Cookie) string { _ = "STUB: not implemented"; return "" }

func dumpCurlHeaders(req *http.Request) *[][2]string { _ = "STUB: not implemented"; return nil }

var regexCmdQuote = regexp.MustCompile(`[^\w@%+=:,./-]`)

func cmdQuote(s string) string { _ = "STUB: not implemented"; return "" }
