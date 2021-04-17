package domain

import (
	"context"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type RegisteredShopUser struct {
	gorm.Model
	Email	string	`json:"email" gorm:"unique"`
	Name	string	`json:"name"`
	Surname	string	`json:"surname"`
	ShopAccount   ShopAccount
	ShopAccountID uint


	Person Person
	PersonID uint64

}

type RegisteredShopUserUsecase interface {
	Fetch(ctx context.Context) ([]*RegisteredShopUser, error)
	GetByID(ctx context.Context, id uint) (*RegisteredShopUser, error)
	Update(ctx context.Context, reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Create(ctx context.Context, reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Delete(ctx context.Context, id uint) error
	GetByUsernameOrEmail(ctx context.Context, username string, email string) (*RegisteredShopUser, error)
}

type RegisteredShopUserRepository interface {
	Fetch(ctx context.Context) ([]*RegisteredShopUser, error)
	GetByID(ctx context.Context, id uint) (*RegisteredShopUser, error)
	Update(ctx context.Context, reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Create(ctx context.Context, reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Delete(ctx context.Context, id uint) error
	GetUserDetailsByAccount(account *ShopAccount) (*RegisteredShopUser, error)
	GetByUsernameOrEmail(ctx echo.Context, username string, email string) (*RegisteredShopUser, error)

}
