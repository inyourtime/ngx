package sql

import (
	"context"
	"ngx/domain"
	"ngx/port"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, u domain.User) (domain.User, error) {
	return domain.User{}, nil
}
