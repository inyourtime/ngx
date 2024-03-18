package restful

import "github.com/gofiber/fiber/v2"

type authAPIHandler struct {
	router    fiber.Router
	validator *Validator
}

func NewAuthAPIHandler(router fiber.Router, v *Validator) *authAPIHandler {
	return &authAPIHandler{
		router:    router,
		validator: v,
	}
}

func (h *authAPIHandler) Init() {
	router := h.router.Group("/auth")
	router.Post("/signup", h.AuthSignUp)
	router.Post("/login", h.AuthLogin)
	router.Post("/refresh", h.AuthRefresh)
}

// Registration API
// @Summary      Registration of user
// @Description  Signup route
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param body body model.AuthSignUpRequest true "body"
// @Success      200
// @Router       /api/auth/signup [POST]
func (h *authAPIHandler) AuthSignUp(c *fiber.Ctx) error {
	return c.SendString("signup")
}

// Login API
// @Summary      User login
// @Description  Login route
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param body body model.AuthLoginRequest true "body"
// @Success      200
// @Router       /api/auth/login [POST]
func (h *authAPIHandler) AuthLogin(c *fiber.Ctx) error {
	return c.SendString("login")
}

func (h *authAPIHandler) AuthRefresh(c *fiber.Ctx) error {
	return c.SendString("refresh")
}
