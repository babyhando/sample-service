package storage

import (
	"context"
	"errors"
	"service/internal/order"
	"service/pkg/adapters/storage/entities"
	"service/pkg/adapters/storage/mappers"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) order.Repo {
	return &orderRepo{db}
}

func (r *orderRepo) GetUserOrders(ctx context.Context, userID uint, limit, offset uint) ([]order.Order, uint, error) {
	query := r.db.WithContext(ctx).Model(&entities.Order{}).Where("user_id = ?", userID)

	var total int64

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if offset > 0 {
		query = query.Offset(int(offset))
	}

	if limit > 0 {
		query = query.Limit(int(limit))
	}

	var orders []entities.Order

	if err := query.Find(&orders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	return mappers.OrderEntitiesToDomain(orders), uint(total), nil
}
