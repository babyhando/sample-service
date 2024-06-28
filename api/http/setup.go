package http

import (
	"fmt"
	"log"
	"service/api/http/handlers"
	"service/api/http/middlerwares"
	"service/config"
	adapter "service/pkg/adapters"
	"service/pkg/jwt"
	"service/service"

	"github.com/gofiber/fiber/v2"
)

func Run(cfg config.Server, app *service.AppContainer) {
	fiberApp := fiber.New()
	api := fiberApp.Group("/api/v1", middlerwares.SetUserContext())

	// register global routes
	registerGlobalRoutes(api, app)

	secret := []byte(cfg.TokenSecret)

	// registering users APIs
	registerUsersAPI(api, app.UserService(), secret)

	// registering orders APIs
	registerOrderRoutes(api, app, secret)

	// run server
	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.Host, cfg.HttpPort)))
}

func registerUsersAPI(router fiber.Router, _ *service.UserService, secret []byte) {
	userGroup := router.Group("/users", middlerwares.Auth(secret), middlerwares.RoleChecker("user", "admin"))

	userGroup.Get("/folan", func(c *fiber.Ctx) error {
		claims := c.Locals(jwt.UserClaimKey).(*jwt.UserClaims)

		return c.JSON(map[string]any{
			"user_id": claims.UserID,
			"role":    claims.Role,
		})
	})
}

func registerGlobalRoutes(router fiber.Router, app *service.AppContainer) {
	router.Post("/login", handlers.LoginUser(app.AuthService()))
	router.Get("/refresh", handlers.RefreshCreds(app.AuthService()))
}

func registerOrderRoutes(router fiber.Router, app *service.AppContainer, secret []byte) {
	router = router.Group("/orders")

	router.Get("", middlerwares.Auth(secret), userRoleChecker(), handlers.UserOrders(app.OrderService()))

	router.Post("",
		middlerwares.SetTransaction(adapter.NewGormCommiter(app.RawRBConnection())),
		middlerwares.Auth(secret),
		userRoleChecker(),
		handlers.CreateUserOrder(app.OrderServiceFromCtx))
}

func userRoleChecker() fiber.Handler {
	return middlerwares.RoleChecker("user")
}
