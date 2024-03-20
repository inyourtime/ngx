package token

import (
	"errors"
	"ngx/util/exception"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JwtMaker struct {
	secretKey string
}

func NewJwtMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, exception.New(exception.TypeInvalidKey, "invalid key size", nil)
	}
	return &JwtMaker{secretKey}, nil
}

func (maker *JwtMaker) CreateToken(arg CreateTokenParams) (string, *Payload, error) {
	payload := NewPayload(arg.User, arg.Duration)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(maker.secretKey))
	return token, payload, err
}

func (maker *JwtMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, exception.New(exception.TypeTokenInvalid, "invalid token", nil)
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, exception.New(exception.TypeTokenExpired, "token expired", err)
		}
		return nil, exception.New(exception.TypeTokenInvalid, "invalid token", err)
	}

	return jwtToken.Claims.(*Payload), nil
}
