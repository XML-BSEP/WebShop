package domain

import (
	"context"
	"gorm.io/gorm"
)

type Currency int

const (
	USD Currency = iota
	EUR
	RSD
)

type Product struct {
	gorm.Model
	Name  string
	Price uint64
	Image string
	Currency Currency
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
