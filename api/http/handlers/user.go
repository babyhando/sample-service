package handlers

import (
	"service/service"

	"github.com/gofiber/fiber/v2"
)

func LoginUser(authService *service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&input); err != nil {
			return sendError(c, err, fiber.StatusBadRequest)
		}

		authToken, err := authService.Login(c.Context(), input.Email, input.Password)
		if err != nil {
			return sendError(c, err, fiber.StatusBadRequest)
		}

		return c.JSON(map[string]any{
			"auth":    authToken.AuthorizationToken,
			"refresh": authToken.RefreshToken,
			"exp":     authToken.ExpiresAt,
		})
	}
}
