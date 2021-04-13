package domain

import (
	"gorm.io/gorm"
)

type ShoppingCartItem struct {
	gorm.Model
	Product        Product
	Amount         uint
	ShoppingCartID uint
}
