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

func TestAuthMiddleware(t *testing.T) {
	cases := []struct {
		name          string
		isValidToken  bool
		token         string
		hasAuthHeader bool
		respCode      int
	}{
		{
			name:          "200 ok",
			isValidToken:  true,
			token:         "",
			respCode:      http.StatusOK,
			hasAuthHeader: true,
		},
		{
			name:          "401 no authorization header",
			isValidToken:  false,
			token:         "",
			respCode:      http.StatusUnauthorized,
			hasAuthHeader: false,
		},
		{
			name:          "401 invalid authorization header format",
			isValidToken:  false,
			token:         "foobar",
			respCode:      http.StatusUnauthorized,
			hasAuthHeader: true,
		},
		{
			name:          "401 invalid authorization header type",
			isValidToken:  false,
			token:         "ApiKey foobar",
			respCode:      http.StatusUnauthorized,
			hasAuthHeader: true,
		},
		{
			name:          "401 invalid token",
			isValidToken:  false,
			token:         "Bearer foobar",
			respCode:      http.StatusUnauthorized,
			hasAuthHeader: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			// TODO mock tokenMaker
			secretKey := util.RandomString(32)
			maker, err := token.NewJwtMaker(secretKey)
			require.NoError(t, err)

			mw := NewMiddleware(maker)

			e.Get("/test", mw.AuthMiddleware(true), func(c *fiber.Ctx) error {
				return c.SendString("ok")
			})

			if tc.isValidToken {
				token, _, err := maker.CreateToken(token.CreateTokenParams{
					User:     domain.User{Model: gorm.Model{ID: 1}, Email: "test@email.com"},
					Duration: time.Minute,
				})
				require.NoError(t, err)
				tc.token = "Bearer " + token
			}

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			if tc.hasAuthHeader {
				req.Header.Set("Authorization", tc.token)
			}

			resp, err := e.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tc.respCode, resp.StatusCode)
		})
	}
}
