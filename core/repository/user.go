package repository

import (
	"context"
	"github.com/muhadif/mandes/core/entity"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserSerialByRefreshToken(ctx context.Context, refreshToken string) (string, error)
	UpsertRefreshToken(ctx context.Context, userSerial string, token string) error
	GetUserByUserSerial(ctx context.Context, userSerial string) (*entity.User, error)
	RemoveUserToken(ctx context.Context, userSerial string) error
}
