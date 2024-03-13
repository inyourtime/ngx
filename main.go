package main

import (
	"fmt"
	"log"
	"ngx/adapter/logger"
	"ngx/core/util"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "ngx/docs"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	config, err := util.LoadConfig(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load env config: %s", err)
		os.Exit(1)
	}

	logger := logger.NewLogger(config)
	logger.Info().Msg("This is a ngx project ✨")

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	app.Get("/", xxx)

	log.Fatal(app.Listen(":3000"))
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {string}  "ok"
// @Router       / [get]
func xxx(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
