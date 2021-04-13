package usecase

import (
	"context"
	"web-shop/domain"
)

type RegisteredUserUsecase struct {
	RegisteredUserRepository domain.RegisteredShopUserRepository
}

func NewRegisteredUserUsecase(r domain.RegisteredShopUserRepository) domain.RegisteredShopUserUsecase {
	return &RegisteredUserUsecase{r}
}

func (r *RegisteredUserUsecase) Fetch(ctx context.Context) ([]*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.Fetch(ctx)
}

func (r *RegisteredUserUsecase) GetByID(ctx context.Context, id uint) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.GetByID(ctx, id)
}

func (r *RegisteredUserUsecase) Update(ctx context.Context, reg *domain.RegisteredShopUser) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.Update(ctx, reg)
}

func (r *RegisteredUserUsecase) Create(ctx context.Context, reg *domain.RegisteredShopUser) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.Create(ctx, reg)
}

func (r *RegisteredUserUsecase) Delete(ctx context.Context, id uint) error {
	return r.RegisteredUserRepository.Delete(ctx, id)
}

