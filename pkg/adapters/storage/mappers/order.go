package mappers

import (
	"service/internal/order"
	"service/pkg/adapters/storage/entities"
	"service/pkg/fp"

	"gorm.io/gorm"
)

func OrderEntityToDomain(orderEntity entities.Order) order.Order {
	return order.Order{
		ID:            orderEntity.ID,
		CreatedAt:     orderEntity.CreatedAt,
		TotalPrice:    orderEntity.TotalPrice,
		TotalQuantity: orderEntity.TotalQuantity,
		Description:   orderEntity.Description,
		UserID:        orderEntity.UserID,
	}
}

func OrderEntitiesToDomain(orderEntities []entities.Order) []order.Order {
	return fp.Map(orderEntities, OrderEntityToDomain)
}

func OrderDomainToEntity(o *order.Order) *entities.Order {
	return &entities.Order{
		Model: gorm.Model{
			CreatedAt: o.CreatedAt,
		},
		TotalPrice:    o.TotalPrice,
		TotalQuantity: o.TotalQuantity,
		Description:   o.Description,
		UserID:        o.UserID,
	}
}
