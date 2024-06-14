package handlers

import "github.com/gofiber/fiber/v2"

func SendError(c *fiber.Ctx, err error, status int) error {
	if status == 0 {
		status = fiber.StatusInternalServerError
	}
	return c.Status(status).JSON(map[string]any{
		"error_msg": err.Error(),
	})
}
