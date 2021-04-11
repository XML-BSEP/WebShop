package domain

import (
	"context"
	"gorm.io/gorm"
)

type RegisteredShopUser struct {
	gorm.Model
	Person     Person    `json:"Person"`
	ShopAccount   ShopAccount    `json:"shopAccount"`

}

type RegisteredShopUserUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]RegisteredShopUser, string, error)
	GetByID(ctx context.Context, id int64) (RegisteredShopUser, error)
	Update(ctx context.Context, adr *RegisteredShopUser) error
	GetByTitle(ctx context.Context, title string) (RegisteredShopUser, error)
	Store(ctx context.Context, adr *RegisteredShopUser) error
	Delete(ctx context.Context, id int64) error
}

type RegisteredShopUserRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []RegisteredShopUser, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (RegisteredShopUser, error)
	GetByTitle(ctx context.Context, title string) (RegisteredShopUser, error)
	Update(ctx context.Context, adr *RegisteredShopUser) error
	Store(ctx context.Context, adr *RegisteredShopUser) error
	Delete(ctx context.Context, id int64) error
}
