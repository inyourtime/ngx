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

func NewUser(arg User) (User, error) {
	newUser := User{
		Email:    arg.Email,
		Name:     arg.Name,
		Password: arg.Password,
	}

	return newUser, nil
}
