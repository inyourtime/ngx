package logger

import (
	"ngx/port"
	"ngx/util"
)

func NewLogger(config util.Config) port.Logger {
	return NewZeroLogLogger(config)
}
