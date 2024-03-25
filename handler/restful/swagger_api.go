package restful

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type swaggerAPIHandler struct {
	router fiber.Router
}

func NewSwaggerAPIHandler(router fiber.Router) *swaggerAPIHandler {
	return &swaggerAPIHandler{
		router: router,
	}
}

func (h *swaggerAPIHandler) Init() {
	router := h.router

	router.Get("/docs/*", swagger.HandlerDefault) // default

	// router.Get("/swagger/*", swagger.New(swagger.Config{ // custom
	// 	URL:         "http://localhost:5000/api/docs/doc.json",
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
