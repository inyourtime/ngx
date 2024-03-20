package restful

import (
	"github.com/gofiber/fiber/v2"
)

func (server *Server) RegisterRoutes() {
	validator := NewValidator()

	app := server.app
	uc := server.usecase
	mw := NewMiddleware(uc.TokenMaker())
	api := app.Group("/api")

	NewSwaggerAPIHandler(api).Init()
	NewAppAPIHandler(api).Init()
	NewAuthAPIHandler(api, validator, uc.Auth()).Init()
	NewUserAPIHandler(api, validator, mw).Init()

	// Not found route handler
	app.Use(NotFoundHandler)
}

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": 404, "message": "route not found"})
}
