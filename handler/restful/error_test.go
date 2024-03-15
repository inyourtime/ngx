package restful

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestErrorHandler(t *testing.T) {
	app := fiber.New()

	t.Run("fiber.Error case", func(t *testing.T) {
		err := fiber.ErrBadRequest
		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)

		_ = errorHandler(c, err)

		assert.Equal(t, http.StatusBadRequest, c.Response().Header.StatusCode())
	})

	t.Run("validator.ValidationErrors case", func(t *testing.T) {
		type MockRequest struct {
			Email string `validate:"required,email"`
		}

		mock := MockRequest{}
		validator := NewValidator()

		err := validator.Validate(&mock)

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)

		_ = errorHandler(c, err)

		assert.Equal(t, http.StatusBadRequest, c.Response().Header.StatusCode())

		var responseBody map[string]interface{}
		err = json.Unmarshal(c.Response().Body(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, fiber.ErrBadRequest.Message, responseBody["message"])
		assert.NotNil(t, responseBody["description"])
	})
}
