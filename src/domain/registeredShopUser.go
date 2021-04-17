package domain

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type RegisteredShopUser struct {
	gorm.Model
	Email	string	`json:"email" gorm:"unique"`
	SecurityQuestion	string	`json:"question"`
	SecurityAnswer	string	`json:"answer"`
	Name	string	`json:"name"`
	Surname	string	`json:"surname"`
	ShopAccount   ShopAccount
	ShopAccountID uint

}

type RegisteredShopUserUsecase interface {
	Fetch(ctx echo.Context) ([]*RegisteredShopUser, error)
	GetByID(ctx echo.Context, id uint) (*RegisteredShopUser, error)
	Update(ctx echo.Context, reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Create(ctx echo.Context, reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Delete(ctx echo.Context, id uint) error
	ExistByUsernameOrEmail(ctx echo.Context, username string, email string) (*RegisteredShopUser, error)
}

type RegisteredShopUserRepository interface {
	Fetch() ([]*RegisteredShopUser, error)
	GetByID(id uint) (*RegisteredShopUser, error)
	Update(reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Create(reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Delete(id uint) error
	GetUserDetailsByAccount(account *ShopAccount) (*RegisteredShopUser, error)
	ExistByUsernameOrEmail(username string, email string) (*RegisteredShopUser, error)

}
