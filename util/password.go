package util

import (
	"ngx/util/exception"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a hashed password from the input password string.
//
// It takes a password string as a parameter and returns the hashed password string and an error.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", exception.New(exception.TypePermissionDenied, "invalid password", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword is a function to compare a plain-text password with its hashed version.
//
// It takes two parameters: password of type string and hashedPassword of type string.
// It returns an error.
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
