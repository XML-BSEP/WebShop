package domain

import (
	"github.com/labstack/echo"
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
	Fetch(ctx echo.Context) ([]*Order, error)
	GetByID(ctx echo.Context, id uint) (*Order, error)
	Update(ctx echo.Context, o *Order) (*Order, error)
	Create(ctx echo.Context, o *Order) (*Order, error)
	Delete(ctx echo.Context, id uint) error
}

type OrderRepository interface {
	Fetch() ([]*Order, error)
	GetByID(id uint) (*Order, error)
	Update(o *Order) (*Order, error)
	Create(o *Order) (*Order, error)
	Delete(id uint) error
}
