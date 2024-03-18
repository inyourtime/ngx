package restful

import "github.com/gofiber/fiber/v2"

type appAPIHandler struct {
	router fiber.Router
}

func NewAppAPIHandler(router fiber.Router) *appAPIHandler {
	return &appAPIHandler{
		router: router,
	}
}

func (h *appAPIHandler) Init() {
	router := h.router
	router.Get("/", h.AppIndex)
}

// Root godoc
//	@Security		ApiKeyAuth
//	@Summary		Root
//	@Description	Root route
//	@Tags			root
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/api [get]
func (h *appAPIHandler) AppIndex(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hello World from Root ✨"})
}
