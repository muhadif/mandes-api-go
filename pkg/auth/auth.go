package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhadif/mandes/config"
	"github.com/muhadif/mandes/core/entity"
	"strconv"
	"time"
)

func CreateJWTToken(user *entity.User, cfg config.Config) (*entity.AccessToken, error) {
	expiredTime := time.Now().Add(time.Minute * time.Duration(cfg.JWTExpiredTime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role,
		"exp":   expiredTime,
	})
	signed, err := token.SignedString(cfg.AppSecretKey)
	if err != nil {
		return nil, err
	}

	refreshTokenExpiredTime := time.Now().Add(time.Minute * time.Duration(cfg.JWTExpiredTime)).Unix()
	refreshToken := jwt.New(jwt.SigningMethodES256)
	signedRefreshToken, err := refreshToken.SignedString(cfg.AppSecretKey)
	if err != nil {
		return nil, err
	}

	return &entity.AccessToken{
		AccessToken:  signed,
		RefreshToken: signedRefreshToken,
		AtExpires:    expiredTime,
		RtExpires:    refreshTokenExpiredTime,
	}, nil
}

func ValidateRefreshToken(refreshToken string) bool {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}

	var expiredAt = time.Now()
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		i, err := strconv.ParseInt(fmt.Sprintf("%s", claims["exp"]), 10, 64)
		if err != nil {
			return false
		}
		expiredAt = time.Unix(i, 0)
	}

	if expiredAt.After(time.Now()) {
		return false
	}

	return true
}
