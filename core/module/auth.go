package module

import (
	"context"
	"github.com/muhadif/mandes/config"
	"github.com/muhadif/mandes/core/entity"
	coreErr "github.com/muhadif/mandes/core/error"
	"github.com/muhadif/mandes/core/repository"
	"github.com/muhadif/mandes/pkg/auth"
	"github.com/muhadif/mandes/pkg/fault"
	pkgString "github.com/muhadif/mandes/pkg/string"
	"golang.org/x/crypto/bcrypt"
)

type AuthModule interface {
	Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error)
	RefreshToken(ctx context.Context, req *entity.RefreshTokenRequest) (*entity.LoginResponse, error)
	Register(ctx context.Context, req *entity.RegisterRequest) error
	RegisterFromAdmin(ctx context.Context, req *entity.RegisterFromAdminRequest) error
	Logout(ctx context.Context, req *entity.LogoutRequest) error
}

type authModule struct {
	userRepository repository.UserRepository
	cfg            config.Config
}

func NewAuthModule(userRepository repository.UserRepository, cfg config.Config) AuthModule {
	return &authModule{
		userRepository: userRepository,
		cfg:            cfg,
	}
}

func (a *authModule) Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error) {
	checkUser, err := a.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	token, err := auth.CreateJWTToken(checkUser, a.cfg)
	if err != nil {
		return nil, err
	}

	err = a.userRepository.UpsertRefreshToken(ctx, checkUser.Serial, token.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{AccessToken: token}, nil
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
		Status:       entity.UserStatusDisabled,
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

func (a *authModule) RefreshToken(ctx context.Context, req *entity.RefreshTokenRequest) (*entity.LoginResponse, error) {
	if isValid := auth.ValidateRefreshToken(req.RefreshToken); !isValid {
		return nil, fault.ErrorDictionary(fault.HTTPUnauthorizedError, coreErr.ErrTokenNotValid)
	}

	userSerial, err := a.userRepository.GetUserSerialByRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	user, err := a.userRepository.GetUserByUserSerial(ctx, userSerial)
	if err != nil {
		return nil, err
	}

	token, err := auth.CreateJWTToken(user, a.cfg)
	if err != nil {
		return nil, err
	}

	err = a.userRepository.UpsertRefreshToken(ctx, user.Serial, token.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{AccessToken: token}, nil
}

func (a *authModule) Logout(ctx context.Context, req *entity.LogoutRequest) error {
	if err := a.userRepository.RemoveUserToken(ctx, req.UserSerial); err != nil {
		return err
	}
	return nil
}
