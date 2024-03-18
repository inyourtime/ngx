package restful

import (
	"ngx/domain"
	"ngx/models"
	"ngx/port"

	"github.com/gofiber/fiber/v2"
)

type authAPIHandler struct {
	router    fiber.Router
	validator *Validator
	authUc    port.AuthUsecase
}

func NewAuthAPIHandler(router fiber.Router, v *Validator, au port.AuthUsecase) *authAPIHandler {
	return &authAPIHandler{
		router:    router,
		validator: v,
		authUc:    au,
	}
}

func (h *authAPIHandler) Init() {
	router := h.router.Group("/auth")
	router.Post("/signup", h.AuthSignUp)
	router.Post("/login", h.AuthLogin)
	router.Post("/refresh", h.AuthRefresh)
}

// Registration API
//	@Summary		Registration of user
//	@Description	Signup route
//	@Tags			auth
//	@Produce		json
//	@Param			body	body	models.AuthSignUpRequest	true	"body"
//	@Success		200
//	@Router			/api/auth/signup [POST]
func (h *authAPIHandler) AuthSignUp(c *fiber.Ctx) error {
	req := models.AuthSignUpRequest{}

	if err := h.validator.Bind(c, &req); err != nil {
		return errorHandler(c, err)
	}

	user, err := h.authUc.SignUp(c.UserContext(), port.SignUpParams{
		User: domain.User{
			Email:    req.Email,
			Name:     req.Name,
			Password: req.Password,
		},
	})
	if err != nil {
		return errorHandler(c, err)
	}
	return c.JSON(user)
}

// Login API
//	@Summary		User login
//	@Description	Login route
//	@Tags			auth
//	@Produce		json
//	@Param			body	body	models.AuthLoginRequest	true	"body"
//	@Success		200
//	@Router			/api/auth/login [POST]
func (h *authAPIHandler) AuthLogin(c *fiber.Ctx) error {
	return c.SendString("login")
}

func (h *authAPIHandler) AuthRefresh(c *fiber.Ctx) error {
	return c.SendString("refresh")
}
