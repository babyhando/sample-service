package http

import (
	"fmt"
	"log"
	"service/config"

	"github.com/gofiber/fiber/v2"
)

func Run(cfg config.Server) {
	fiberApp := fiber.New()

	// registering global middlewares
	registerMiddlewares(fiberApp)

	// registering users APIs
	registerUsersAPI(fiberApp)

	// run server
	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.Host, cfg.HttpPort)))
}

func registerMiddlewares(fiberApp *fiber.App) {

}

func registerUsersAPI(fiberApp *fiber.App) {
	userGroup := fiberApp.Group("/users")

	userGroup.Get("/:id", nil) // todo
}
