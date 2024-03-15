package restful

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	_ = NewValidator()

	api := app.Group("/api")

	NewSwaggerAPIHandler(api).Init()
	NewAppAPIHandler(api).Init()

	// Not found route handler
	app.Use(NotFoundHandler)
}

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": 404, "message": "route not found"})
}
