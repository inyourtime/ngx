package port

import "ngx/util/token"

type AuthParams struct {
	Token   string
	Payload *token.Payload
}

type UserUsecase interface {
}
