package restful

import (
	"ngx/port"

	"github.com/gofiber/fiber/v2"
)

type userAPIHandler struct {
	router    fiber.Router
	validator *Validator
	mw        port.Middleware
}

func NewUserAPIHandler(router fiber.Router, v *Validator, mw port.Middleware) *userAPIHandler {
	return &userAPIHandler{
		router:    router,
		validator: v,
		mw:        mw,
	}
}

func (h *userAPIHandler) Init() {
	router := h.router.Group("/users")
	// apply auth middleware
	router.Get("/me", h.Me)
}

// Me API
//
//	@Security		ApiKeyAuth
//	@Summary		User account (me)
//	@Description	User account route
//	@Tags			user
//	@Produce		json
//	@Success		200
//	@Router			/api/users/me [GET]
func (h *userAPIHandler) Me(c *fiber.Ctx) error {
	return c.SendString("me")
}
