package repository

import (
	"context"
	"github.com/muhadif/mandes/core/entity"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
}
