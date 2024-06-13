package handlers

import (
	"service/service"

	"github.com/gofiber/fiber/v2"
)

func CreateUserHandler(service *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
