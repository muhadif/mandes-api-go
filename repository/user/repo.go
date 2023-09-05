package user

import (
	"context"
	"github.com/muhadif/mandes/core/entity"
	"github.com/muhadif/mandes/core/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *gorm.DB
}

func (r repo) GetUserByUserSerial(ctx context.Context, userSerial string) (*entity.User, error) {
	var user *entity.User
	err := r.db.Table("user").Where("serial = ?", userSerial).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r repo) GetUserSerialByRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	var userToken *entity.UserToken

	err := r.db.Table("user_token").Where("refresh_token = ?", refreshToken).First(&userToken).Error
	if err != nil {
		return "", nil
	}

	return userToken.UserSerial, nil
}

func (r repo) UpsertRefreshToken(ctx context.Context, serial string, token string) error {
	userToken := &entity.UserToken{
		UserSerial:   serial,
		RefreshToken: token,
	}
	err := r.db.Table("user_token").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_serial"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"refresh_token": token}),
	}).Create(&userToken).Error

	if err != nil {
		return err
	}

	return nil
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

func (r repo) RemoveUserToken(ctx context.Context, userSerial string) error {
	var user *entity.User
	err := r.db.Table("user_token").Where("user_serial = ?", userSerial).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
