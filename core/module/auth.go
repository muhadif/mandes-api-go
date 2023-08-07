package module

import (
	"context"
	"github.com/muhadif/mandes/core/entity"
	coreErr "github.com/muhadif/mandes/core/error"
	"github.com/muhadif/mandes/core/repository"
	"github.com/muhadif/mandes/pkg/fault"
	pkgString "github.com/muhadif/mandes/pkg/string"
	"golang.org/x/crypto/bcrypt"
)

type AuthModule interface {
	Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error)
	Register(ctx context.Context, req *entity.RegisterRequest) error
	RegisterFromAdmin(ctx context.Context, req *entity.RegisterFromAdminRequest) error
}

type authModule struct {
	userRepository repository.UserRepository
}

func NewAuthModule(userRepository repository.UserRepository) AuthModule {
	return &authModule{
		userRepository: userRepository,
	}
}


func (a *authModule) Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error) {
	panic("implement me")
}

func (a *authModule) Register(ctx context.Context, req *entity.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	checkUser, err := a.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if checkUser != nil {
		return fault.ErrorDictionary(fault.HTTPBadRequestError, coreErr.ErrEmailTaken)
	}

	user := &entity.User{
		Serial:       pkgString.GenerateSerial(entity.UserSerialPrefix, entity.UserSerialLength),
		Email:        req.Email,
		Role:         entity.UserRoleUser,
		Password:     string(hashedPassword),
		AccessStatus: entity.UserStatusDisabled,
		Status:            entity.UserStatusDisabled,
	}
	err = a.userRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (a *authModule) RegisterFromAdmin(ctx context.Context, req *entity.RegisterFromAdminRequest) error {
	panic("implement me")
}