package token

import (
	"ngx/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CreateTokenParams struct {
	User     domain.User
	Duration time.Duration
}

type Maker interface {
	CreateToken(CreateTokenParams) (string, *Payload, error)
	VerifyToken(string) (*Payload, error)
}

type Payload struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func NewPayload(user domain.User, duration time.Duration) *Payload {
	return &Payload{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
}
