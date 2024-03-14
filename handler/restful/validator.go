package restful

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (server *Server) setupValidator() {
	validator := validator.New()
	validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	server.validator = validator
}
