package handlers

import (
	"errors"
	"service/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

const UserClaimKey = jwt.UserClaimKey

var (
	errWrongClaimType = errors.New("wrong claim type")
)

func SendError(c *fiber.Ctx, err error, status int) error {
	if status == 0 {
		status = fiber.StatusInternalServerError
	}
	return c.Status(status).JSON(map[string]any{
		"error_msg": err.Error(),
	})
}

func PageAndPageSize(c *fiber.Ctx) (int, int) {
	page, pageSize := c.QueryInt("page"), c.QueryInt("page_size")
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 20
	}

	return page, pageSize
}
