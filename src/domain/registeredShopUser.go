package domain

import (
	"context"
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
	Fetch() ([]*RegisteredShopUser, error)
	GetByID(id uint) (*RegisteredShopUser, error)
	Update(reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Create(reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Delete(id uint) error
	GetUserDetailsByAccount(account *ShopAccount) (*RegisteredShopUser, error)
	GetByUsernameOrEmail(username string, email string) (*RegisteredShopUser, error)

}
