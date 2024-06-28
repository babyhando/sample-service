package service

import (
	"context"
	"log"
	"service/config"
	"service/internal/order"
	"service/internal/user"
	"service/pkg/adapters/storage"
	"service/pkg/valuecontext"

	"gorm.io/gorm"
)

type AppContainer struct {
	cfg          config.Config
	dbConn       *gorm.DB
	userService  *UserService
	authService  *AuthService
	orderService *OrderService
}

func NewAppContainer(cfg config.Config) (*AppContainer, error) {
	app := &AppContainer{
		cfg: cfg,
	}

	app.mustInitDB()
	storage.Migrate(app.dbConn)

	app.setUserService()
	app.setAuthService()
	app.setOrderService()

	return app, nil
}

func (a *AppContainer) RawRBConnection() *gorm.DB {
	return a.dbConn
}

func (a *AppContainer) UserService() *UserService {
	return a.userService
}

func (a *AppContainer) OrderService() *OrderService {
	return a.orderService
}

func (a *AppContainer) OrderServiceFromCtx(ctx context.Context) *OrderService {
	tx, ok := valuecontext.TryGetTxFromContext(ctx)
	if !ok {
		return a.orderService
	}

	gc, ok := tx.Tx().(*gorm.DB)
	if !ok {
		return a.orderService
	}

	return NewOrderService(
		order.NewOps(storage.NewOrderRepo(gc)),
		user.NewOps(storage.NewUserRepo(gc)),
	)
}

func (a *AppContainer) AuthService() *AuthService {
	return a.authService
}

func (a *AppContainer) setUserService() {
	if a.userService != nil {
		return
	}
	a.userService = NewUserService(user.NewOps(storage.NewUserRepo(a.dbConn)))
}

func (a *AppContainer) setOrderService() {
	if a.orderService != nil {
		return
	}
	a.orderService = NewOrderService(order.NewOps(storage.NewOrderRepo(a.dbConn)), user.NewOps(storage.NewUserRepo(a.dbConn)))
}

func (a *AppContainer) mustInitDB() {
	if a.dbConn != nil {
		return
	}

	db, err := storage.NewMysqlGormConnection(a.cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	a.dbConn = db
}

func (a *AppContainer) setAuthService() {
	if a.authService != nil {
		return
	}

	a.authService = NewAuthService(user.NewOps(storage.NewUserRepo(a.dbConn)), []byte(a.cfg.Server.TokenSecret),
		a.cfg.Server.TokenExpMinutes,
		a.cfg.Server.RefreshTokenExpMinutes)
}
