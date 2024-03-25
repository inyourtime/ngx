package restful

import (
	"ngx/domain"
	"ngx/models"
	"ngx/port"
	"ngx/util"
	"ngx/util/exception"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type authAPIHandler struct {
	router    fiber.Router
	validator *Validator
	authUc    port.AuthUsecase
	config    util.Config
	store     *session.Store
}

func NewAuthAPIHandler(router fiber.Router, v *Validator, au port.AuthUsecase, config util.Config, store *session.Store) *authAPIHandler {
	return &authAPIHandler{
		router:    router,
		validator: v,
		authUc:    au,
		config:    config,
		store:     store,
	}
}

func (h *authAPIHandler) Init() {
	router := h.router.Group("/auth")
	router.Post("/signup", h.AuthSignUp)
	router.Post("/login", h.AuthLogin)
	router.Post("/refresh", h.AuthRefresh)
	router.Get("/google", h.GoogleIndex)
	router.Get("/google/callback", h.GoogleCallback)
	router.Get("/github", h.GithubIndex)
	router.Get("/github/callback", h.GithubCallback)
}

// Registration API
//
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
			Email:     req.Email,
			Password:  req.Password,
			FirstName: req.FirstName,
			LastName:  req.LastName,
		},
	})
	if err != nil {
		return errorHandler(c, err)
	}
	return c.JSON(user)
}

// Login API
//
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

// Refresh API
//
//	@Summary		Refresh token
//	@Description	Refresh token route
//	@Tags			auth
//	@Produce		json
//	@Success		200
//	@Router			/api/auth/refresh [POST]
func (h *authAPIHandler) AuthRefresh(c *fiber.Ctx) error {
	return c.SendString("refresh")
}

// Google API
//
//	@Summary		Google OAuth
//	@Description	Google OAuth API
//	@Tags			auth
//	@Produce		json
//	@Success		200
//	@Router			/api/auth/google [GET]
func (h *authAPIHandler) GoogleIndex(c *fiber.Ctx) error {
	state, err := GenerateState()
	if err != nil {
		return errorHandler(c, err)
	}
	url := h.config.GetGoogleCfg().AuthCodeURL(state)

	sess, err := h.store.Get(c)
	if err != nil {
		return errorHandler(c, err)
	}
	sess.Set(state, true)

	if err := sess.Save(); err != nil {
		return errorHandler(c, err)
	}

	return c.Redirect(url)
}

func (h *authAPIHandler) GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")

	sess, err := h.store.Get(c)
	if err != nil {
		return errorHandler(c, err)
	}

	s := sess.Get(state)
	isValidState, ok := s.(bool)
	if !ok || !isValidState {
		return errorHandler(c, exception.New(exception.TypePermissionDenied, "invalid state", nil))
	}
	sess.Delete(state)

	if err := sess.Save(); err != nil {
		return errorHandler(c, err)
	}

	code := c.Query("code")

	token, err := h.config.GetGoogleCfg().Exchange(c.UserContext(), code)
	if err != nil {
		return errorHandler(c, exception.New(exception.TypePermissionDenied, "invalid code", err))
	}

	endpoint := h.config.GoogleUserInfoURL + "?access_token=" + token.AccessToken
	userInfo, err := GetGoogleUserInfo(endpoint)
	if err != nil {
		return errorHandler(c, err)
	}

	return c.JSON(userInfo)
}

func (h *authAPIHandler) GithubIndex(c *fiber.Ctx) error {
	state, err := GenerateState()
	if err != nil {
		return errorHandler(c, err)
	}
	url := h.config.GetGithubCfg().AuthCodeURL(state)

	sess, err := h.store.Get(c)
	if err != nil {
		return errorHandler(c, err)
	}
	sess.Set(state, true)

	if err := sess.Save(); err != nil {
		return errorHandler(c, err)
	}

	return c.Redirect(url)
}

func (h *authAPIHandler) GithubCallback(c *fiber.Ctx) error {
	state := c.Query("state")

	sess, err := h.store.Get(c)
	if err != nil {
		return errorHandler(c, err)
	}

	s := sess.Get(state)
	isValidState, ok := s.(bool)
	if !ok || !isValidState {
		return errorHandler(c, exception.New(exception.TypePermissionDenied, "invalid state", nil))
	}
	sess.Delete(state)

	if err := sess.Save(); err != nil {
		return errorHandler(c, err)
	}

	code := c.Query("code")

	token, err := h.config.GetGithubCfg().Exchange(c.UserContext(), code)
	if err != nil {
		return errorHandler(c, exception.New(exception.TypePermissionDenied, "invalid code", err))
	}

	userData, err := GetGithubUserInfo(h.config.GithubUserInfoURL, token.AccessToken)
	if err != nil {
		return errorHandler(c, err)
	}

	userEmail, err := GetGithubUserEmail("https://api.github.com/user/emails", token.AccessToken)
	if err != nil {
		return errorHandler(c, err)
	}

	githubInfo := &domain.GithubUserInfo{
		ID:       strconv.Itoa(userData.ID),
		Name:     userData.Name,
		Email:    userEmail.Email,
		Verified: userEmail.Verified,
		Primary:  userEmail.Primary,
	}
	return c.JSON(githubInfo)
}
