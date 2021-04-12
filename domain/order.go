package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	DateOfPlacement time.Time
	ShoppingCart    ShoppingCart
	TotalPrice      uint
}

type OrderUsecase interface {
	Fetch(ctx context.Context) ([]*Order, error)
	GetByID(ctx context.Context, id uint) (*Order, error)
	Update(ctx context.Context, o *Order) (*Order, error)
	Create(ctx context.Context, o *Order) (*Order, error)
	Delete(ctx context.Context, id uint) error
}

type OrderRepository interface {
	Fetch(ctx context.Context) ([]*Order, error)
	GetByID(ctx context.Context, id uint) (*Order, error)
	Update(ctx context.Context, o *Order) (*Order, error)
	Create(ctx context.Context, o *Order) (*Order, error)
	Delete(ctx context.Context, id uint) error
}
