package domain

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type ShoppingCartItem struct {
	gorm.Model
	Product        Product
	Amount         uint
	ShoppingCartID uint
}

type ShoppingCartItemUsecase interface {
	Fetch(ctx echo.Context) ([]*ShoppingCartItem, error)
	GetByID(ctx echo.Context, id uint) (*ShoppingCartItem, error)
	Update(ctx echo.Context, s *ShoppingCartItem) (*ShoppingCartItem, error)
	Create(ctx echo.Context, s *ShoppingCartItem) (*ShoppingCartItem, error)
	Delete(ctx echo.Context, id uint) error
}

type ShoppingCartItemRepository interface {
	Fetch() ([]*ShoppingCartItem, error)
	GetByID(id uint) (*ShoppingCartItem, error)
	Update(s *ShoppingCartItem) (*ShoppingCartItem, error)
	Create(s *ShoppingCartItem) (*ShoppingCartItem, error)
	Delete(id uint) error
}
