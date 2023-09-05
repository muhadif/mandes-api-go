package route

import (
	"github.com/gin-gonic/gin"
	"github.com/muhadif/mandes/app"
	"github.com/muhadif/mandes/handler"
	"github.com/muhadif/mandes/pkg/auth"
)

func NewRouter(app *app.App) *gin.Engine {
	authHandler := handler.NewAuthHandler(app.AuthModule)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(auth.AuthMiddleware(true))

	router.POST("/auth/login", authHandler.Login)
	router.Handle("POST", "register/login", authHandler.Register)

	admin := router.Group("/admin")
	adminVillager := admin.Group("/villager")
	{
		adminVillager.GET("/")
	}
	return router
}
