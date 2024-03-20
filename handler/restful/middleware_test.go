package restful

import (
	"net/http"
	"net/http/httptest"
	"ngx/domain"
	"ngx/util"
	"ngx/util/token"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestAuthMiddleware_200_OK(t *testing.T) {
	setup()
	// TODO mock tokenMaker
	secretKey := util.RandomString(32)
	maker, err := token.NewJwtMaker(secretKey)
	require.NoError(t, err)

	token, _, err := maker.CreateToken(token.CreateTokenParams{
		User:     domain.User{Model: gorm.Model{ID: 1}, Email: "test@email.com"},
		Duration: time.Minute,
	})
	require.NoError(t, err)

	mw := NewMiddleware(maker)

	e.Get("/test", mw.AuthMiddleware(true), func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := e.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAuthMiddleware_401_Unauthorized(t *testing.T) {
	setup()
	// TODO mock tokenMaker
	secretKey := util.RandomString(32)
	maker, err := token.NewJwtMaker(secretKey)
	require.NoError(t, err)

	mw := NewMiddleware(maker)

	e.Get("/test", mw.AuthMiddleware(true), func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)

	resp, err := e.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestAuthMiddleware_401_InvalidAuthorizationFormat(t *testing.T) {
	setup()
	// TODO mock tokenMaker
	secretKey := util.RandomString(32)
	maker, err := token.NewJwtMaker(secretKey)
	require.NoError(t, err)

	mw := NewMiddleware(maker)

	e.Get("/test", mw.AuthMiddleware(true), func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "foobar")

	resp, err := e.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestAuthMiddleware_401_InvalidAuthorizationType(t *testing.T) {
	setup()
	// TODO mock tokenMaker
	secretKey := util.RandomString(32)
	maker, err := token.NewJwtMaker(secretKey)
	require.NoError(t, err)

	mw := NewMiddleware(maker)

	e.Get("/test", mw.AuthMiddleware(true), func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "ApiKey foobar")

	resp, err := e.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestAuthMiddleware_401_InvalidToken(t *testing.T) {
	setup()
	// TODO mock tokenMaker
	secretKey := util.RandomString(32)
	maker, err := token.NewJwtMaker(secretKey)
	require.NoError(t, err)

	mw := NewMiddleware(maker)

	e.Get("/test", mw.AuthMiddleware(true), func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer foobar")

	resp, err := e.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
