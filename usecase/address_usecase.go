package usecase

import (
	"context"
	"web-shop/domain"
)

type addressUseCase struct {
	AddressRepository domain.AddressRepository
}

func NewAddresUseCase(r domain.AddressRepository) domain.AddressUsecase {
	return &addressUseCase{r}
}

func (a *addressUseCase) Fetch(ctx context.Context) ([]*domain.Address, error) {
	return a.AddressRepository.Fetch(ctx)
}

func (a *addressUseCase) GetByID(ctx context.Context, id uint) (*domain.Address, error) {
	return a.AddressRepository.GetByID(ctx, id)
}

func (a *addressUseCase) Update(ctx context.Context, adr *domain.Address) (*domain.Address, error) {
	return a.AddressRepository.Update(ctx, adr)
}

func (a *addressUseCase) Create(ctx context.Context, adr *domain.Address) (*domain.Address, error) {
	return a.AddressRepository.Create(ctx, adr)
}

func (a *addressUseCase) Delete(ctx context.Context, id uint) error {
	return a.AddressRepository.Delete(ctx, id)
}

