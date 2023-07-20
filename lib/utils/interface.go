package utils

import "net/http"

type TinkoffService interface {
	Debug() bool
	Log() Logger
	Request(r *http.Request, data any) error
}

type Logger interface {
	Logf(format string, values ...interface{})
}
