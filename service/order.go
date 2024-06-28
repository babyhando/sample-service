package service

import (
	"context"
	"service/internal/order"
	u "service/internal/user"
)

type OrderService struct {
	orderOps *order.Ops
	userOps  *u.Ops
}

func NewOrderService(orderOps *order.Ops, userOps *u.Ops) *OrderService {
	return &OrderService{
		orderOps: orderOps,
		userOps:  userOps,
	}
}

func (s *OrderService) GetUserOrders(ctx context.Context, userID uint, page, pageSize uint) ([]order.Order, uint, error) {
	user, err := s.userOps.GetUserByID(ctx, userID)
	if err != nil {
		return nil, 0, err
	}

	if user == nil {
		return nil, 0, u.ErrUserNotFound
	}

	return s.orderOps.UserOrders(ctx, userID, page, pageSize)
}

func (s *OrderService) CreateOrder(ctx context.Context, o *order.Order) error {
	user, err := s.userOps.GetUserByID(ctx, o.UserID)
	if err != nil {
		return err
	}

	if user == nil {
		return u.ErrUserNotFound
	}

	if err := s.orderOps.Create(ctx, o); err != nil {
		return err
	}

	return nil
	// return errors.New("test error") -> for testing transaction
}
