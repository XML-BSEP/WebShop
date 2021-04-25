package dto

import "time"

type OrderDTO struct {
	ID 				uint
	DateOfPlacement time.Time
	ShoppingCart    ShoppingCartDTO
	TotalPrice      uint
}