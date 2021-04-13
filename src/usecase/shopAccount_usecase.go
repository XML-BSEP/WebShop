package usecase

import (
	"context"
	"web-shop/domain"
)

type ShopAccountUsecase struct {
	ShopAccountRepository domain.ShopAccountRepository
}

func NewShopAccoutUsecase(r domain.ShopAccountRepository) domain.ShopAccountUsecase {
	return &ShopAccountUsecase{r}
}

func (s *ShopAccountUsecase) Fetch(ctx context.Context) ([]*domain.ShopAccount, error) {
	return s.ShopAccountRepository.Fetch(ctx)
}

func (s *ShopAccountUsecase) GetByID(ctx context.Context, id uint) (*domain.ShopAccount, error) {
	return s.ShopAccountRepository.GetByID(ctx, id)
}

func (s *ShopAccountUsecase) Update(ctx context.Context, account *domain.ShopAccount) (*domain.ShopAccount, error) {
	return s.ShopAccountRepository.Update(ctx, account)
}

func (s *ShopAccountUsecase) Create(ctx context.Context, account *domain.ShopAccount) (*domain.ShopAccount, error) {
	return s.ShopAccountRepository.Create(ctx, account)
}

func (s *ShopAccountUsecase) Delete(ctx context.Context, id uint) error {
	return s.ShopAccountRepository.Delete(ctx, id)
}

