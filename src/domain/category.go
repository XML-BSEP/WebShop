package domain

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string
}


type CategoryUsecase interface {
	Fetch(ctx echo.Context) ([]*Category, error)
	GetByID(ctx echo.Context, id uint) (*Category, error)
	Update(ctx echo.Context, pic *Category) (*Category, error)
	Create(ctx echo.Context, pic *Category) (*Category, error)
	Delete(ctx echo.Context, id uint) error
}

type CategoryRepository interface {
	Fetch() ([]*Category, error)
	GetByID(id uint) (*Category, error)
	Update(pic *Category) (*Category, error)
	Create(pic *Category) (*Category, error)
	Delete(id uint) error
}