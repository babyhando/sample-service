package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	TotalPrice    uint
	TotalQuantity uint
	Description   string
	UserID        uint
	User          User
}
