package route

import (
	"github.com/gin-gonic/gin"
	"github.com/muhadif/mandes/app"
	"github.com/muhadif/mandes/handler"
)

func NewRouter(app *app.App) *gin.Engine {
	authHandler := handler.NewAuthHandler(app.AuthModule)

	router := gin.Default()
	router.Handle("POST", "/auth/login", authHandler.Login)
	router.Handle("POST", "register/login", authHandler.Register)

	return router
}
