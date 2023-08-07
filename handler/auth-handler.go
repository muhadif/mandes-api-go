package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muhadif/mandes/api"
	"github.com/muhadif/mandes/core/entity"
	"github.com/muhadif/mandes/core/module"
	"github.com/muhadif/mandes/pkg/fault"
	"net/http"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

func NewAuthHandler(authModule module.AuthModule) AuthHandler {
	return &authHandler{
		authModule: authModule,
	}
}

type authHandler struct {
	authModule		module.AuthModule
}

func (a *authHandler) Register(ctx *gin.Context) {
	var req *entity.RegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.ResponseFailed(ctx, fault.ErrorDictionary(fault.HTTPBadRequestError, err.Error()))
		return
	}

	err := a.authModule.Register(ctx, req)
	if err != nil {
		api.ResponseFailed(ctx, err)
	}

	api.ResponseSuccess(ctx, http.StatusOK, nil)
}

func (a *authHandler) Login(ctx *gin.Context) {
	var req *entity.LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.ResponseFailed(ctx, fault.ErrorDictionary(fault.HTTPBadRequestError, err.Error()))
		return
	}

	resp, err := a.authModule.Login(ctx, req)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, resp)
}


