package domain

import (
	"context"
	"gorm.io/gorm"
)

type RegisteredShopUser struct {
	gorm.Model
	Person     Person
	PersonID uint64
	ShopAccount   ShopAccount
	ShopAccountID uint

}

type RegisteredShopUserUsecase interface {
	Fetch(ctx context.Context) ([]*RegisteredShopUser, error)
	GetByID(ctx context.Context, id uint) (*RegisteredShopUser, error)
	Update(ctx context.Context, reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Create(ctx context.Context, reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Delete(ctx context.Context, id uint) error
}

type RegisteredShopUserRepository interface {
	Fetch(ctx context.Context) ([]*RegisteredShopUser, error)
	GetByID(ctx context.Context, id uint) (*RegisteredShopUser, error)
	Update(ctx context.Context, reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Create(ctx context.Context, reg *RegisteredShopUser) (*RegisteredShopUser, error)
	Delete(ctx context.Context, id uint) error
	GetUserDetailsByUsername(account *ShopAccount) (*RegisteredShopUser, error)
}
