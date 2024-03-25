package domain

import (
	"ngx/util"
	"time"

	"gorm.io/gorm"
)

type roleType uint

const (
	AdminRole roleType = 101
	UserRole  roleType = 201
)

type User struct {
	gorm.Model
	Email          string  `gorm:"uniqueIndex;not null"`
	Password       *string `gorm:"default:null"`
	FirstName      string
	LastName       string
	Role           roleType   `gorm:"default:201"`
	SocialID       *string    `gorm:"default:null"`
	SocialProvider *string    `gorm:"default:null"`
	Verified       *time.Time `gorm:"default:null"`
}

// Type of Google user information
type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}

// Type of Github user information
type GithubUserInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
	Primary  bool   `json:"primary"`
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
		Email:          arg.Email,
		FirstName:      arg.FirstName,
		LastName:       arg.LastName,
		Role:           arg.Role,
		SocialID:       arg.SocialID,
		SocialProvider: arg.SocialProvider,
		Verified:       arg.Verified,
	}

	if err := newUser.SetPassword(*arg.Password); err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (u *User) SetPassword(pwd string) error {
	hashedPwd, err := util.HashPassword(pwd)
	if err != nil {
		return err
	}
	u.Password = &hashedPwd
	return nil
}
