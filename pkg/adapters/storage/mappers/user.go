package mappers

import (
	"service/internal/user"
	"service/pkg/adapters/storage/entities"
)

func UserEntityToDomain(entity *entities.User) *user.User {
	return &user.User{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Email:     entity.Email,
		Password:  entity.Password,
		Role:      user.UserRole(entity.Role),
	}
}
