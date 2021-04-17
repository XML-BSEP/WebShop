package domain

import (
	"context"

	"gorm.io/gorm"
)

type ShoppingCart struct {
	gorm.Model
	ShoppingCartItems []ShoppingCartItem
	Buyer             RegisteredShopUser
}

type ShoppingCartUsecase interface {
	Fetch(ctx context.Context) ([]*ShoppingCart, error)
	GetByID(ctx context.Context, id uint) (*ShoppingCart, error)
	Update(ctx context.Context, s *ShoppingCart) (*ShoppingCart, error)
	Create(ctx context.Context, s *ShoppingCart) (*ShoppingCart, error)
	Delete(ctx context.Context, id uint) error
}

type ShoppingCartRepository interface {
	Fetch() ([]*ShoppingCart, error)
	GetByID(d uint) (*ShoppingCart, error)
	Update(s *ShoppingCart) (*ShoppingCart, error)
	Create(s *ShoppingCart) (*ShoppingCart, error)
	Delete(id uint) error
}
