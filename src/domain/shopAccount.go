package domain

import (
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
}

type ShopAccountRepository interface {
	Fetch() ([]*ShopAccount, error)
	GetByID(id uint) (*ShopAccount, error)
	Update(account *ShopAccount) (*ShopAccount, error)
	Create(account *ShopAccount) (*ShopAccount, error)
	Delete(id uint) error
	GetUserDetailsByUsername(account *ShopAccount) (*ShopAccount, error)
}
