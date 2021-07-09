package domain

import (
	"context"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"web-shop/infrastructure/dto"
)

type Currency int



type Product struct {
	gorm.Model
	Name         string `json:"name"`
	Price        float64 `json:"price"`
	Available    uint `json:"available"`
	Description  string `json:"description"`
	Images       []Image
	Category     Category
	CategoryId   uint
	SerialNumber uint64 `json:"serial"`
	ShopAccount   ShopAccount
	ShopAccountID uint `json:"shopAccountId"`
}

type ProductUsecase interface {
	GetAllAvailableProductsInUsersShop(ctx echo.Context, userId uint) ([]*Product, error)
	Fetch(ctx echo.Context) ([]*Product, error)
	GetByID(ctx echo.Context, id uint) (*Product, error)
	Update(ctx echo.Context, pic *dto.EditProduct) (*Product, error)
	Create(ctx echo.Context, pic *dto.NewProduct) (*Product, error)
	Delete(ctx echo.Context, deletedProduct dto.DeleteProduct) error
	GetWithPriceRange(low uint, high uint)([]*Product, error)
	GetProductsWithCategory(category string)([]*Product, error)
	GetProductsWithCondition(low uint, high uint, category string, limit int, offset int)([]*Product, error)
	GetByName(name string, limit int, offset int)([]*Product, error)
	GetProductsWithConditionOrderedByPrice(low uint, high uint, category string, limit int, offset int, order int)([]*Product, error)
	GetProductsWithConditionOrderedByName(low uint, high uint, category string, limit int, offset int, order int)([]*Product, error)
	GetByNameOrderByPrice(name string, limit int, offset int, order int)([]*Product, error)
	GetByNameOrderByName(name string, limit int, offset int, order int)([]*Product, error)
	FilterByCategory(userId uint, name string, category string, priceRangeStart uint, priceRangeEnd uint, limit int, offset int, order string) ([]*Product, error)
	Count() (int64, error)
	GetBySerial(serial uint64) (*Product, error)
	GetAllProductsInUsersShop(ctx echo.Context, userId uint) ([]*Product, error)
	GetProductDetails(ctx context.Context, productId uint)(*Product, error)

}

type ProductRepository interface {
	GetAllAvailableProductsInUsersShop(ctx echo.Context, userId uint) ([]*Product, error)
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
	FilterByCategory(userId uint,name string, category string, priceRangeStart uint, priceRangeEnd uint, limit int, offset int, order string) ([]*Product, error)
	Count() (int64, error)
	MinMaxPrice() (int64)
	GetBySerial(serial uint64) (*Product, error)
	GetBySerialAndUserId(serial uint64, id uint) (*Product, error)
	GetAllProductsInUsersShop(ctx echo.Context, userId uint) ([]*Product, error)
	GetProductDetails(ctx context.Context, productId uint)(*Product, error)
}
