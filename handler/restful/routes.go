package restful

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func (server *Server) RegisterRoutes() {
	// Validator
	validator := NewValidator()

	app := server.app
	uc := server.usecase

	// Middleware
	mw := NewMiddleware(uc.TokenMaker())

	// API group
	api := app.Group("/api")

	// Session store
	store := session.New(session.ConfigDefault)
	store.Expiration = 24 * time.Hour

	// Routes
	{
		NewSwaggerAPIHandler(api).Init()
		NewAppAPIHandler(api).Init()
		NewAuthAPIHandler(api, validator, uc.Auth(), server.config, store).Init()
		NewUserAPIHandler(api, validator, mw).Init()
	}

	// Not found route handler
	app.Use(NotFoundHandler)
}

// Not found route handler
func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": 404, "message": "route not found"})
}
