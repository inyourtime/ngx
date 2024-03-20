package port

import "ngx/util/token"

type Usecase interface {
	Auth() AuthUsecase
	User() UserUsecase
	TokenMaker() token.Maker
}
