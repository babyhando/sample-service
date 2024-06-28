package order

import "context"

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
