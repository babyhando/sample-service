package order

import (
	"context"
	"time"
)

type Repo interface {
	GetUserOrders(ctx context.Context, userID uint, limit, offset uint) ([]Order, uint, error)
}

type Order struct {
	ID            uint
	CreatedAt     time.Time
	TotalPrice    uint
	TotalQuantity uint
	Description   string
	UserID        uint
}
