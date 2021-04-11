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
	Name     string    `json:"name"`
	Surname   string    `json:"surname"`
	Phone 	string	`json:"phone"`
	Gender Gender `json:"gender"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Address Address `json:"address" gorm:"embedded"`
}


type PersonUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Person, string, error)
	GetByID(ctx context.Context, id int64) (Person, error)
	Update(ctx context.Context, adr *Person) error
	GetByTitle(ctx context.Context, title string) (Person, error)
	Store(ctx context.Context, adr *Person) error
	Delete(ctx context.Context, id int64) error
}

type PersonRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Person, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (Person, error)
	GetByTitle(ctx context.Context, title string) (Person, error)
	Update(ctx context.Context, adr *Person) error
	Store(ctx context.Context, adr *Person) error
	Delete(ctx context.Context, id int64) error
}

