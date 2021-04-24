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
	Image string 	`json:"image"`
	Currency Currency `json:"currency"`
	Available uint `json:"available"`
	Description string `json:"description"`
	Category Category
	CategoryId uint
}

type ProductUsecase interface {
	Fetch(ctx echo.Context) ([]*Product, error)
	GetByID(ctx echo.Context, id uint) (*Product, error)
	Update(ctx echo.Context, pic *Product) (*Product, error)
	Create(ctx echo.Context, pic *Product) (*Product, error)
	Delete(ctx echo.Context, id uint) error
	GetWithPriceRange(low uint, high uint)([]*Product, error)
	GetProductsWithCategory(category string)([]*Product, error)
	GetProductsWithCondition(low uint, high uint, category string, limit int, offset int)([]*Product, error)
	GetByName(name string, limit int, offset int)([]*Product, error)
	GetProductsWithConditionOrderedByPrice(low uint, high uint, category string, limit int, offset int, order int)([]*Product, error)
	GetProductsWithConditionOrderedByName(low uint, high uint, category string, limit int, offset int, order int)([]*Product, error)
	GetByNameOrderByPrice(name string, limit int, offset int, order int)([]*Product, error)
	GetByNameOrderByName(name string, limit int, offset int, order int)([]*Product, error)
	FilterByCategory(category string, priceRangeStart uint, priceRangeEnd uint, limit int, offset int, order string) ([]*Product, error)
}

type ProductRepository interface {
	Fetch() ([]*Product, error)
	GetByID(id uint) (*Product, error)
	Update(pic *Product) (*Product, error)
	Create(pic *Product) (*Product, error)
	Delete(id uint) error
	GetWithPriceRange(low uint, high uint)([]*Product, error)
	GetProductsWithCategory(category string)([]*Product, error)
	GetProductsWithCondition(low uint, high uint, category string, limit int, offset int)([]*Product, error)
	GetByName(name string, limit int, offset int)([]*Product, error)
	GetProductsWithConditionOrderedByPrice(low uint, high uint, category string, limit int, offset int, order int)([]*Product, error)
	GetProductsWithConditionOrderedByName(low uint, high uint, category string, limit int, offset int, order int)([]*Product, error)
	GetByNameOrderByPrice(name string, limit int, offset int, order int)([]*Product, error)
	GetByNameOrderByName(name string, limit int, offset int, order int)([]*Product, error)
	FilterByCategory(category string, priceRangeStart uint, priceRangeEnd uint, limit int, offset int, order string) ([]*Product, error)
}
