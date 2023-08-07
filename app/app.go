package app

import (
	"github.com/muhadif/mandes/core/module"
	"github.com/muhadif/mandes/repository/user"
)

type App struct {
	AuthModule module.AuthModule
}

func NewApp(dep *Dependency) *App {
	authRepo := user.NewUserRepository(dep.Database)

	authModule := module.NewAuthModule(authRepo)
	return &App{
		AuthModule: authModule,
	}
}
