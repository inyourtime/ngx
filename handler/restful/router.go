package restful

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (server *Server) setupApplication() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New(cors.ConfigDefault))

	_ = customValidator()

	api := app.Group("/api")

	NewSwaggerAPIHandler(api).Init()
	NewAppAPIHandler(api).Init()

	// Not found route handler
	app.Use(NotFoundHandler)

	server.app = app
}

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": 404, "message": "route not found"})
}
