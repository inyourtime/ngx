package restful

import (
	"fmt"
	"ngx/port"
	"ngx/util/exception"
	"ngx/util/token"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	authorizationHeaderKey = "Authorization"
	authorizationTypeToken = "Bearer"
	authorizationArgKey    = "authorization_arg"
)

type middleware struct {
	tokenMaker token.Maker
}

func NewMiddleware(tm token.Maker) port.Middleware {
	return &middleware{
		tokenMaker: tm,
	}
}

func (m *middleware) AuthMiddleware(autoDenied bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authArg, err := m.parseToken(c)
		if err != nil {
			if hasToken(c) {
				return errorHandler(c, err)
			}
			if autoDenied {
				return errorHandler(c, err)
			}
		}
		c.Locals(authorizationArgKey, authArg)
		return c.Next()
	}
}

func (m *middleware) parseToken(c *fiber.Ctx) (port.AuthParams, error) {
	authorizationHeader := c.Get(authorizationHeaderKey)
	if len(authorizationHeader) == 0 {
		msg := "authorization header not provided"
		err := exception.New(exception.TypePermissionDenied, msg, nil)
		return port.AuthParams{}, err
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		msg := "invalid authorization format"
		err := exception.New(exception.TypePermissionDenied, msg, nil)
		return port.AuthParams{}, err
	}

	authorizationType := fields[0]
	if authorizationType != authorizationTypeToken {
		msg := fmt.Sprintf("authorization type %s not supported", authorizationType)
		err := exception.New(exception.TypePermissionDenied, msg, nil)
		return port.AuthParams{}, err
	}

	token := fields[1]
	payload, err := m.tokenMaker.VerifyToken(token)
	if err != nil {
		return port.AuthParams{}, err
	}
	return port.AuthParams{Token: token, Payload: payload}, nil
}

func hasToken(c *fiber.Ctx) bool {
	authorizationHeader := c.Get(authorizationHeaderKey)
	return len(authorizationHeader) > 0
}
