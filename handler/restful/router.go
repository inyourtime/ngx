package restful

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func (server *Server) setupApplication() {
	app := fiber.New()

	app.Use(recover.New())

	// Root
	app.Get("/", RootHandler)

	// Swagger documentation
	docsRouter(app.Group("/swagger"))

	// Start setup entire route

	// End setup route

	// Not found route handler
	app.Use(NotFoundHandler)

	server.app = app
}

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": 404, "message": "route not found"})
}

// Root godoc
// @Summary      Root
// @Description  Root route
// @Tags         root
// @Accept       json
// @Produce      json
// @Success      200
// @Router       / [get]
func RootHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hello World from Root ✨"})
}

func docsRouter(router fiber.Router) {
	router.Get("/*", swagger.HandlerDefault) // default

	// router.Get("/swagger/*", swagger.New(swagger.Config{ // custom
	// 	URL:         "http://example.com/doc.json",
	// 	DeepLinking: false,
	// 	// Expand ("list") or Collapse ("none") tag groups by default
	// 	DocExpansion: "none",
	// 	// Prefill OAuth ClientId on Authorize popup
	// 	OAuth: &swagger.OAuthConfig{
	// 		AppName:  "OAuth Provider",
	// 		ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
	// 	},
	// 	// Ability to change OAuth2 redirect uri location
	// 	OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	// }))
}
