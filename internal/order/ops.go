package order

import (
	"context"
	"time"
)

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo}
}

func (o *Ops) UserOrders(ctx context.Context, userID uint, page, pageSize uint) ([]Order, uint, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	return o.repo.GetUserOrders(ctx, userID, limit, offset)
}

func (o *Ops) Create(ctx context.Context, order *Order) error {
	if order.TotalPrice < order.TotalQuantity {
		return ErrQuantityGreater
	}

	if order.CreatedAt.After(time.Now()) {
		return ErrWrongOrderTime
	}

	return o.repo.Insert(ctx, order)
}
