package usecase

import (
	"context"
	"web-shop/domain"
)

type PersonUsecase struct {
	PersonRepository domain.PersonRepository
}

func NewPersonUsecase(r domain.PersonRepository) domain.PersonUsecase {
	return &PersonUsecase{r}
}

func (p *PersonUsecase) Fetch(ctx context.Context) ([]*domain.Person, error) {
	return p.PersonRepository.Fetch()
}

func (p *PersonUsecase) GetByID(ctx context.Context, id uint) (*domain.Person, error) {
	return p.PersonRepository.GetByID(id)
}

func (p *PersonUsecase) Update(ctx context.Context, person *domain.Person) (*domain.Person, error) {
	return p.PersonRepository.Update(person)
}

func (p *PersonUsecase) Create(ctx context.Context, person *domain.Person) (*domain.Person, error) {
	return p.PersonRepository.Create(person)
}

func (p *PersonUsecase) Delete(ctx context.Context, id uint) error {
	return p.PersonRepository.Delete(id)
}


