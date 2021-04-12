package domain

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Price string
	Image []string
}

type ProductUsecase interface {
	Fetch(ctx context.Context) ([]*Product, error)
	GetByID(ctx context.Context, id uint) (*Product, error)
	Update(ctx context.Context, pic *Product) (*Product, error)
	Create(ctx context.Context, pic *Product) (*Product, error)
	Delete(ctx context.Context, id uint) error
}

type ProductRepository interface {
	Fetch(ctx context.Context) ([]*Product, error)
	GetByID(ctx context.Context, id uint) (*Product, error)
	Update(ctx context.Context, pic *Product) (*Product, error)
	Create(ctx context.Context, pic *Product) (*Product, error)
	Delete(ctx context.Context, id uint) error
}
