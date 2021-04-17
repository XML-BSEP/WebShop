package domain

import (
	"github.com/labstack/echo"

	"gorm.io/gorm"
)

type Storage struct {
	gorm.Model
	Product   Product
	ProductID uint
	Available uint
}

type StorageUsecase interface {
	Fetch(ctx echo.Context) ([]*Storage, error)
	GetByID(ctx echo.Context, id uint) (*Storage, error)
	Update(ctx echo.Context, s *Storage) (*Storage, error)
	Create(ctx echo.Context, s *Storage) (*Storage, error)
	Delete(ctx echo.Context, id uint) error
}

type StorageRepository interface {
	Fetch() ([]*Storage, error)
	GetByID(id uint) (*Storage, error)
	Update(s *Storage) (*Storage, error)
	Create(s *Storage) (*Storage, error)
	Delete(id uint) error
}
