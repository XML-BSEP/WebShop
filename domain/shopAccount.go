package domain

import (
	"context"
	"gorm.io/gorm"
)

type ShopAccount struct {
	gorm.Model
	Username     string    `json:"username"`
	Password   string    `json:"password"`

}

type ShopAccountUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]ShopAccount, string, error)
	GetByID(ctx context.Context, id int64) (ShopAccount, error)
	Update(ctx context.Context, adr *ShopAccount) error
	GetByTitle(ctx context.Context, title string) (ShopAccount, error)
	Store(ctx context.Context, adr *ShopAccount) error
	Delete(ctx context.Context, id int64) error
}

type ShopAccountRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []ShopAccount, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (ShopAccount, error)
	GetByTitle(ctx context.Context, title string) (ShopAccount, error)
	Update(ctx context.Context, adr *ShopAccount) error
	Store(ctx context.Context, adr *ShopAccount) error
	Delete(ctx context.Context, id int64) error
}

