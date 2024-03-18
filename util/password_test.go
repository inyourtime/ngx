package util

import (
	"ngx/util/exception"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(8)
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err)

	err = CheckPassword(password, hashedPassword)
	assert.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPassword)
	assert.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	differentHashedPassword, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, differentHashedPassword)
	assert.NotEqual(t, hashedPassword, differentHashedPassword)
}

func TestPasswordTooLong(t *testing.T) {
	password := RandomString(73)
	_, err := HashPassword(password)

	assert.EqualError(t, err, bcrypt.ErrPasswordTooLong.Error())
	assert.ErrorIs(t, err, err.(*exception.Exception))
	assert.Equal(t, exception.TypePermissionDenied, err.(*exception.Exception).Type)
}
