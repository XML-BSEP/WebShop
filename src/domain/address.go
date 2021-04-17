package domain

import (
	"github.com/labstack/echo"
)

type Address struct {
	ID uint
	Street     string    `json:"street"`
	City   string    `json:"city"`
	State    string    `json:"state"`
	Zip uint `json:"zip"`
}

type AddressUsecase interface {
	Fetch(ctx echo.Context) ([]*Address, error)
	GetByID(ctx echo.Context, id uint) (*Address, error)
	Update(ctx echo.Context, adr *Address) (*Address, error)
	Create(ctx echo.Context, adr *Address) (*Address, error)
	Delete(ctx echo.Context, id uint) error
}

type AddressRepository interface {
	Fetch() ([]*Address, error)
	GetByID(id uint) (*Address, error)
	Update(adr *Address)(*Address, error)
	Create(adr *Address) (*Address, error)
	Delete(id uint) error
}

