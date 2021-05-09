package domain


import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"time"
)

type Image struct {
	gorm.Model
	Timestamp time.Time
	Path    string
	ProductId uint
}

type ImageUseCase interface {
	Fetch(ctx echo.Context) ([]*Image, error)
	GetByID(ctx echo.Context, id uint) (*Image, error)
	Update(ctx echo.Context, o *Image) (*Image, error)
	Create(ctx echo.Context, o *Image) (*Image, error)
	Delete(ctx echo.Context, id uint) error
	GetyByPath(path string) ([]*Image, error)
}

type ImageRepository interface {
	Fetch() ([]*Image, error)
	GetByID(id uint) (*Image, error)
	Update(o *Image) (*Image, error)
	Create(o *Image) (*Image, error)
	Delete(id uint) error
	GetyByPath(path string) ([]*Image, error)
}