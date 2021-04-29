package domain

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleName string
}


type RoleUsecase interface {
	Fetch(ctx echo.Context) ([]*Role, error)
	GetByID(ctx echo.Context, id uint) (*Role, error)
	Update(ctx echo.Context, o *Role) (*Role, error)
	Create(ctx echo.Context, o *Role) (*Role, error)
	Delete(ctx echo.Context, id uint) error
}

type RoleRepository interface {
	Fetch() ([]*Role, error)
	GetByID(id uint) (*Role, error)
	Update(o *Role) (*Role, error)
	Create(o *Role) (*Role, error)
	Delete(id uint) error
}