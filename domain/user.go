package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null"`
	Name     *string
	Password *string
}

// NewUser creates a new User based on the provided User argument.
//
// Parameter:
//
//	arg User - the User object containing Email, Name, and Password.
//
// Return:
//
//	User - the newly created User object.
//	error - an error if any.
func NewUser(arg User) (User, error) {
	newUser := User{
		Email:    arg.Email,
		Name:     arg.Name,
		Password: arg.Password,
	}

	return newUser, nil
}
