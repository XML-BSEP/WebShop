package domain

import (
	"context"
)

type Address struct {
	ID uint
	Street     string    `json:"street"`
	City   string    `json:"city"`
	State    string    `json:"state"`
	Zip uint `json:"zip"`
}

type AddressUsecase interface {
	Fetch(ctx context.Context) ([]*Address, error)
	GetByID(ctx context.Context, id uint) (*Address, error)
	Update(ctx context.Context, adr *Address) (*Address, error)
	Create(ctx context.Context, adr *Address) (*Address, error)
	Delete(ctx context.Context, id uint) error
}

type AddressRepository interface {
	Fetch() ([]*Address, error)
	GetByID(id uint) (*Address, error)
	Update(adr *Address)(*Address, error)
	Create(adr *Address) (*Address, error)
	Delete(id uint) error
}

