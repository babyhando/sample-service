package service

import (
	"service/config"
	"service/internal/user"
)

type AppContainer struct {
	cfg         config.Config
	userService *UserService
}

func NewAppContainer(cfg config.Config) (*AppContainer, error) {
	// setup database connection based on config
	// create repos

	// create services
	app := new(AppContainer)

	app.userService = NewUserService(user.NewOps(nil))

	return app, nil
}

func (a *AppContainer) UserService() *UserService {
	return a.userService
}
