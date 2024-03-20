package port

import "github.com/gofiber/fiber/v2"

type Server interface {
	Start() error
}

type Middleware interface {
	AuthMiddleware(autoDenied bool) fiber.Handler
}
