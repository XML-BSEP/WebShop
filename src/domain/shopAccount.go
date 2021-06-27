package domain

import (
	"context"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type ShopAccount struct {
	gorm.Model
	Username     string	`json:"username" gorm:"unique"`
	Password   string	`json:"password"`
}

type ShopAccountUsecase interface {
	Fetch(ctx echo.Context) ([]*ShopAccount, error)
	GetByID(ctx echo.Context, id uint) (*ShopAccount, error)
	Update(ctx echo.Context, account *ShopAccount) (*ShopAccount, error)
	Create(ctx echo.Context, account *ShopAccount) (*ShopAccount, error)
	Delete(ctx echo.Context, id uint) error
	FetchShops(ctx echo.Context) ([]*ShopAccount, error)
	GetShopAccountByUsername(ctx context.Context, username string) (*ShopAccount, error)
	GetUserDetailsByEmail(ctx context.Context, email string) (*ShopAccount, error)
}

type ShopAccountRepository interface {
	Fetch() ([]*ShopAccount, error)
	GetByID(id uint) (*ShopAccount, error)
	Update(account *ShopAccount) (*ShopAccount, error)
	Create(account *ShopAccount) (*ShopAccount, error)
	Delete(id uint) error
	GetUserDetailsByUsername(ctx context.Context, username string) (*ShopAccount, error)
	FetchShops(ctx echo.Context) ([]*ShopAccount, error)
	GetUserDetailsByEmail(ctx context.Context, email string) (*ShopAccount, error)
}
