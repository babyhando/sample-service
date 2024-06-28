package presenter

import (
	"service/internal/order"
	"service/pkg/fp"
	"time"
)

type UserOrder struct {
	ID            uint      `json:"order_id"`
	CreatedAt     Timestamp `json:"order_time"`
	TotalPrice    uint      `json:"price"`
	TotalQuantity uint      `json:"quantity"`
	Description   string    `json:"description"`
}

func OrderToUserOrder(o order.Order) UserOrder {
	return UserOrder{
		ID:            o.ID,
		CreatedAt:     Timestamp(o.CreatedAt),
		TotalPrice:    o.TotalPrice,
		TotalQuantity: o.TotalQuantity,
		Description:   o.Description,
	}
}

func OrdersToUserOrders(orders []order.Order) []UserOrder {
	return fp.Map(orders, OrderToUserOrder)
}

func UserOrderToOrder(userOrder *UserOrder, userID uint) *order.Order {
	return &order.Order{
		CreatedAt:     time.Time(userOrder.CreatedAt),
		TotalPrice:    userOrder.TotalPrice,
		TotalQuantity: userOrder.TotalQuantity,
		Description:   userOrder.Description,
		UserID:        userID,
	}
}
