package resty

import (
	"net/http"
	"time"
)

type (
	DebugLogCallbackFunc func(*DebugLog)

	DebugLogFormatterFunc func(*DebugLog) string

	DebugLog struct {
		Request   *DebugLogRequest  `json:"request"`
		Response  *DebugLogResponse `json:"response"`
		TraceInfo *TraceInfo        `json:"trace_info"`
	}

	DebugLogRequest struct {
		CorrelationID string `json:"correlation_id"`

		Host string `json:"host"`

		URI string `json:"uri"`

		Method string `json:"method"`

		Proto string `json:"proto"`

		Header http.Header `json:"header"`

		CurlCmd string `json:"curl_cmd"`

		Attempt int `json:"attempt"`

		Body string `json:"body"`
	}

	DebugLogResponse struct {
		StatusCode int `json:"status_code"`

		Status string `json:"status"`

		Proto string `json:"proto"`

		ReceivedAt time.Time `json:"received_at"`

		Duration time.Duration `json:"duration"`

		Size int64 `json:"size"`

		Header http.Header `json:"header"`

		Body string `json:"body"`
	}
)

func DebugLogFormatter(dl *DebugLog) string { _ = "STUB: not implemented"; return "" }

func DebugLogJSONFormatter(dl *DebugLog) string { _ = "STUB: not implemented"; return "" }

func debugLogger(c *Client, res *Response) { _ = "STUB: not implemented"; return }

const debugRequestLogKey = "__restyDebugRequestLog"

func prepareRequestDebugInfo(c *Client, r *Request) { _ = "STUB: not implemented"; return }
