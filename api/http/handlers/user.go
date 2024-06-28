package handlers

import (
	"errors"
	"fmt"
	"service/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoginUser(authService *service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		c.Cookie(&fiber.Cookie{
			Name:        "X-Session-ID",
			Value:       fmt.Sprint(time.Now().UnixNano()),
			HTTPOnly:    true,
			SessionOnly: true,
		})

		if err := c.BodyParser(&input); err != nil {
			return SendError(c, err, fiber.StatusBadRequest)
		}

		authToken, err := authService.Login(c.Context(), input.Email, input.Password)
		if err != nil {
			return SendError(c, err, fiber.StatusBadRequest)
		}

		return SendUserToken(c, authToken)
	}
}

func RefreshCreds(authService *service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		refToken := c.GetReqHeaders()["Authorization"]
		if len(refToken[0]) == 0 {
			return SendError(c, errors.New("token should be provided"), fiber.StatusBadRequest)
		}

		authToken, err := authService.RefreshAuth(c.UserContext(), refToken[0])
		if err != nil {
			return SendError(c, err, fiber.StatusUnauthorized)
		}

		return SendUserToken(c, authToken)
	}
}
