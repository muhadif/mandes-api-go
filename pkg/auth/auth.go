package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhadif/mandes/api"
	"github.com/muhadif/mandes/config"
	"github.com/muhadif/mandes/core/entity"
	coreErr "github.com/muhadif/mandes/core/error"
	"github.com/muhadif/mandes/pkg/fault"
	"strconv"
	"time"
)

func CreateJWTToken(user *entity.User, cfg config.Config) (*entity.AccessToken, error) {
	expiredTime := time.Now().Add(time.Minute * time.Duration(cfg.JWTExpiredTime))
	claim := entity.UserClaimToken{
		UserSerial: user.Serial,
		Role:       user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expiredTime},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
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
		AtExpires:    expiredTime.Unix(),
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

func ValidateToken(tokenStr string) (*entity.UserClaimToken, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &entity.UserClaimToken{}, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(*entity.UserClaimToken); ok && token.Valid {
		return claims, nil
	}

	return nil, fault.ErrorDictionary(fault.HTTPUnauthorizedError, coreErr.ErrTokenNotValid)
}

func AuthMiddleware(useAuth bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		userClaim, err := ValidateToken(authHeader)
		if err == nil && userClaim != nil {
			c.Set("user-context", userClaim)
			c.Next()
		}
		api.ResponseFailed(c, fault.ErrorDictionary(fault.HTTPUnauthorizedError, err.Error()))
		return
	}
}

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaim := c.Request.Header.Get("Authorization")
	}
}
