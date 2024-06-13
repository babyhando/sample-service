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

	// registering global middlewares
	registerMiddlewares(fiberApp)

	// registering users APIs
	registerUsersAPI(fiberApp, app.UserService())

	// run server
	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.Host, cfg.HttpPort)))
}

func registerMiddlewares(fiberApp *fiber.App) {

}

func registerUsersAPI(fiberApp *fiber.App, userService *service.UserService) {
	userGroup := fiberApp.Group("/users")

	userGroup.Post("", handlers.CreateUserHandler(userService)) // todo
}
