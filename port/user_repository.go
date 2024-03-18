package port

import (
	"context"
	"ngx/domain"
)

type UserRepository interface {
	Create(context.Context, domain.User) (domain.User, error)
}
