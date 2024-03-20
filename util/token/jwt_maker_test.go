package token

import (
	"ngx/domain"
	"ngx/util"
	"ngx/util/exception"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestInvalidKey(t *testing.T) {
	maker, err := NewJwtMaker("invalid")
	assert.Error(t, err)
	assert.Nil(t, maker)
}

func TestJwtMaker(t *testing.T) {
	maker, err := NewJwtMaker(util.RandomString(32))
	require.NoError(t, err, "error creating JWT maker")

	user := domain.User{
		Model: gorm.Model{
			ID: 1,
		},
		Email: "test@email.com",
	}
	duration := time.Minute
	expiredAt := time.Now().Add(duration)
	token, payload, err := maker.CreateToken(CreateTokenParams{
		User:     user,
		Duration: duration,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	assert.NoError(t, err)
	assert.NotEmpty(t, payload)

	assert.Equal(t, user.ID, payload.UserID)
	assert.WithinDuration(t, expiredAt, payload.ExpiresAt.Time, time.Second)
}

func TestExpiredJwtToken(t *testing.T) {
	maker, err := NewJwtMaker(util.RandomString(32))
	require.NoError(t, err, "error creating JWT maker")

	user := domain.User{
		Model: gorm.Model{
			ID: 1,
		},
		Email: "test@email.com",
	}
	token, payload, err := maker.CreateToken(CreateTokenParams{
		User:     user,
		Duration: -time.Minute,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	assert.Error(t, err)
	assert.Nil(t, payload)

	fail, ok := err.(*exception.Exception)
	assert.True(t, ok)
	assert.Equal(t, exception.TypeTokenExpired, fail.Type)
}

func TestInvalidJwtToken(t *testing.T) {
	payload := NewPayload(domain.User{Model: gorm.Model{ID: 1}, Email: "test@email.com"}, time.Minute)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	assert.NoError(t, err)

	maker, err := NewJwtMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	assert.Error(t, err)

	fail, ok := err.(*exception.Exception)
	assert.True(t, ok)
	assert.Equal(t, exception.TypeTokenInvalid, fail.Type)
	assert.Nil(t, payload)
}
