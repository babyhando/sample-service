package handlers

import (
	"errors"
	"service/api/http/handlers/presenter"
	"service/internal/order"
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

func CreateUserOrder(orderService *service.OrderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.UserOrder

		if err := c.BodyParser(&req); err != nil {
			return SendError(c, err, fiber.StatusBadRequest)
		}

		userClaims, ok := c.Locals(UserClaimKey).(*jwt.UserClaims)
		if !ok {
			return SendError(c, errWrongClaimType, fiber.StatusBadRequest)
		}

		o := presenter.UserOrderToOrder(&req, userClaims.UserID)

		if err := orderService.CreateOrder(c.UserContext(), o); err != nil {
			status := fiber.StatusInternalServerError
			if errors.Is(err, order.ErrQuantityGreater) || errors.Is(err, order.ErrWrongOrderTime) {
				status = fiber.StatusBadRequest
			}

			return SendError(c, err, status)
		}

		return c.JSON(fiber.Map{
			"order_id": o.ID,
		})
	}
}
