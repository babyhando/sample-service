package handlers

import (
	"errors"
	"service/api/http/handlers/presenter"
	"service/internal/user"
	"service/pkg/jwt"
	"service/service"

	"github.com/gofiber/fiber/v2"
)

func UserOrders(orderService *service.OrderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userClaims, ok := c.Locals(UserClaimKey).(*jwt.UserClaims)
		if !ok {
			return SendError(c, errWrongClaimType, fiber.StatusBadRequest)
		}

		page, pageSize := PageAndPageSize(c)

		orders, total, err := orderService.GetUserOrders(c.UserContext(), userClaims.UserID, uint(page), uint(pageSize))
		if err != nil {
			status := fiber.StatusInternalServerError
			if errors.Is(err, user.ErrUserNotFound) {
				status = fiber.StatusBadRequest
			}
			return SendError(c, err, status)
		}

		return c.JSON(presenter.NewPagination(
			presenter.OrdersToUserOrders(orders),
			uint(page),
			uint(pageSize),
			total,
		))
	}
}
