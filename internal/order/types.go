package order

import (
	"context"
	"errors"
	"time"
)

var (
	ErrQuantityGreater = errors.New("quantity should not be greater than price")
	ErrWrongOrderTime  = errors.New("wrong order time")
)

type Repo interface {
	GetUserOrders(ctx context.Context, userID uint, limit, offset uint) ([]Order, uint, error)
	Insert(ctx context.Context, order *Order) error
}

type Order struct {
	ID            uint
	CreatedAt     time.Time
	TotalPrice    uint
	TotalQuantity uint
	Description   string
	UserID        uint
}
