package domain

import (
	"github.com/labstack/echo"

	"gorm.io/gorm"
)

type ShoppingCart struct {
	gorm.Model
	ShoppingCartItems 	[]ShoppingCartItem
	OrderID 		  	uint
	RegisteredShopUser	RegisteredShopUser
	RegisteredShopUserID uint

}

type ShoppingCartUsecase interface {
	Fetch(ctx echo.Context) ([]*ShoppingCart, error)
	GetByID(ctx echo.Context, id uint) (*ShoppingCart, error)
	Update(ctx echo.Context, s *ShoppingCart) (*ShoppingCart, error)
	Create(ctx echo.Context, s *ShoppingCart) (*ShoppingCart, error)
	Delete(ctx echo.Context, id uint) error
}

type ShoppingCartRepository interface {
	Fetch() ([]*ShoppingCart, error)
	GetByID(d uint) (*ShoppingCart, error)
	Update(s *ShoppingCart) (*ShoppingCart, error)
	Create(s *ShoppingCart) (*ShoppingCart, error)
	Delete(id uint) error
}
