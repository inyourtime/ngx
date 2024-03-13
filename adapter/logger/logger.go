package logger

import (
	"ngx/core/port"
	"ngx/core/util"
)

func NewLogger(config util.Config) port.Logger {
	return NewZeroLogLogger(config)
}
