package domain

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Currency int

const (
	USD Currency = iota
	EUR
	RSD
)

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price uint64 `json:"price"`
	Currency Currency `json:"currency"`
	Category string `json:"category"`
	Available uint `json:"available"`
	Description string `json:"description"`
	Images []Image
}

type ProductUsecase interface {
	Fetch(ctx echo.Context) ([]*Product, error)
	GetByID(ctx echo.Context, id uint) (*Product, error)
	Update(ctx echo.Context, pic *Product) (*Product, error)
	Create(ctx echo.Context, pic *Product) (*Product, error)
	Delete(ctx echo.Context, id uint) error
}

type ProductRepository interface {
	Fetch() ([]*Product, error)
	GetByID(id uint) (*Product, error)
	Update(pic *Product) (*Product, error)
	Create(pic *Product) (*Product, error)
	Delete(id uint) error
}
