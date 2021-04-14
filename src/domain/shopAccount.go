package domain

import (
	"context"
	"gorm.io/gorm"
)

type ShopAccount struct {
	gorm.Model
	Username     string	`json: "username"`
	Password   string	`json: "password"`
}

type ShopAccountUsecase interface {
	Fetch(ctx context.Context) ([]*ShopAccount, error)
	GetByID(ctx context.Context, id uint) (*ShopAccount, error)
	Update(ctx context.Context, account *ShopAccount) (*ShopAccount, error)
	Create(ctx context.Context, account *ShopAccount) (*ShopAccount, error)
	Delete(ctx context.Context, id uint) error
}

type ShopAccountRepository interface {
	Fetch(ctx context.Context) ([]*ShopAccount, error)
	GetByID(ctx context.Context, id uint) (*ShopAccount, error)
	Update(ctx context.Context, account *ShopAccount) (*ShopAccount, error)
	Create(ctx context.Context, account *ShopAccount) (*ShopAccount, error)
	Delete(ctx context.Context, id uint) error
	GetUserDetailsByUsername(account *ShopAccount) (*ShopAccount, error)
}
