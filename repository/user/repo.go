package user

import (
	"context"
	"github.com/muhadif/mandes/core/entity"
	"github.com/muhadif/mandes/core/repository"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *gorm.DB
}

func (r repo) CreateUser(ctx context.Context, user *entity.User) error {
	err := r.db.Table("user").Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *entity.User
	err := r.db.Table("user").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

