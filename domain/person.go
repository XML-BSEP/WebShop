package domain

import (
	"context"
	"gorm.io/gorm"
	"time"
)
type Gender int

const (
	Male Gender = iota
	Female
	Other
)

type Person struct {
	gorm.Model
	Name     string
	Surname   string
	Phone 	string
	Gender Gender
	DateOfBirth time.Time
	Address Address `gorm:"embedded"`

}


type PersonUsecase interface {
	Fetch(ctx context.Context) ([]*Person, error)
	GetByID(ctx context.Context, id uint) (*Person, error)
	Update(ctx context.Context, person *Person) (*Person, error)
	Create(ctx context.Context, person *Person) (*Person, error)
	Delete(ctx context.Context, id uint) error
}

type PersonRepository interface {
	Fetch(ctx context.Context) ([]*Person, error)
	GetByID(ctx context.Context, id uint) (*Person, error)
	Update(ctx context.Context, person *Person) (*Person, error)
	Create(ctx context.Context, person *Person) (*Person, error)
	Delete(ctx context.Context, person uint) error
}

