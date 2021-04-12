package datastore

import (
	"context"
	"web-shop/domain"

	"gorm.io/gorm"
)

type personRepository struct {
	Conn *gorm.DB
}

func (p *personRepository) Fetch(ctx context.Context) ([]*domain.Person, error) {
	var (
		persons []*domain.Person
		err     error
	)
	err = p.Conn.Order("id desc").Find(&persons).Error
	return persons, err
}

func (p *personRepository) Update(ctx context.Context, person *domain.Person) (*domain.Person, error) {
	err := p.Conn.Save(person).Error
	return person, err
}

func (p *personRepository) Create(ctx context.Context, person *domain.Person) (*domain.Person, error) {
	err := p.Conn.Create(person).Error
	return person, err
}

func (p *personRepository) Delete(ctx context.Context, id uint) error {

	person := &domain.Person{Model: gorm.Model{ID: id}}

	err := p.Conn.Delete(person).Error
	return err
}

func (p *personRepository) GetByID(ctx context.Context, id uint) (*domain.Person, error) {
	person := &domain.Person{Model: gorm.Model{ID: id}}
	err := p.Conn.First(person).Error
	return person, err
}

func NewPersonRepository(Conn *gorm.DB) domain.PersonRepository {
	return &personRepository{Conn}
}
