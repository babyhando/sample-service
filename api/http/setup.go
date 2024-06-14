package http

import (
	"fmt"
	"log"
	"service/api/http/handlers"
	"service/config"
	"service/service"

	"github.com/gofiber/fiber/v2"
)

func Run(cfg config.Server, app *service.AppContainer) {
	fiberApp := fiber.New()
	api := fiberApp.Group("/api/v1")

	// register global routes
	registerGlobalRoutes(api, app)

	// registering users APIs
	registerUsersAPI(fiberApp, app.UserService())

	// run server
	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.Host, cfg.HttpPort)))
}

func registerUsersAPI(router fiber.Router, userService *service.UserService) {
	// userGroup := fiberApp.Group("/users")
}

func registerGlobalRoutes(router fiber.Router, app *service.AppContainer) {
	router.Post("/login", handlers.LoginUser(app.AuthService()))
}
