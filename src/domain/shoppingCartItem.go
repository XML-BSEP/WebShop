package domain

import (
	"context"
	"gorm.io/gorm"
)

type ShoppingCartItem struct {
	gorm.Model
	Product        Product
	Amount         uint
	ShoppingCartID uint
}

type ShoppingCartItemUsecase interface {
	Fetch(ctx context.Context) ([]*ShoppingCartItem, error)
	GetByID(ctx context.Context, id uint) (*ShoppingCartItem, error)
	Update(ctx context.Context, s *ShoppingCartItem) (*ShoppingCartItem, error)
	Create(ctx context.Context, s *ShoppingCartItem) (*ShoppingCartItem, error)
	Delete(ctx context.Context, id uint) error
}

type ShoppingCartItemRepository interface {
	Fetch(ctx context.Context) ([]*ShoppingCartItem, error)
	GetByID(ctx context.Context, id uint) (*ShoppingCartItem, error)
	Update(ctx context.Context, s *ShoppingCartItem) (*ShoppingCartItem, error)
	Create(ctx context.Context, s *ShoppingCartItem) (*ShoppingCartItem, error)
	Delete(ctx context.Context, id uint) error
}
