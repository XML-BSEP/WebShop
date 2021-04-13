package domain

import (
	"context"

	"gorm.io/gorm"
)

type Storage struct {
	gorm.Model
	Product   Product
	ProductID uint
	Available uint
}

type StorageUsecase interface {
	Fetch(ctx context.Context) ([]*Storage, error)
	GetByID(ctx context.Context, id uint) (*Storage, error)
	Update(ctx context.Context, s *Storage) (*Storage, error)
	Create(ctx context.Context, s *Storage) (*Storage, error)
	Delete(ctx context.Context, id uint) error
}

type StorageRepository interface {
	Fetch(ctx context.Context) ([]*Storage, error)
	GetByID(ctx context.Context, id uint) (*Storage, error)
	Update(ctx context.Context, s *Storage) (*Storage, error)
	Create(ctx context.Context, s *Storage) (*Storage, error)
	Delete(ctx context.Context, id uint) error
}
