package usecase

import (
	"context"
	"github.com/labstack/echo"
	"web-shop/domain"
)

type ShopAccountUsecase struct {
	ShopAccountRepository domain.ShopAccountRepository
}

func (s *ShopAccountUsecase) GetUserDetailsByEmail(ctx context.Context, email string) (*domain.ShopAccount, error) {
	return s.ShopAccountRepository.GetUserDetailsByEmail(ctx,email)
}

func (s *ShopAccountUsecase) GetShopAccountByUsername(ctx context.Context, username string) (*domain.ShopAccount, error) {
	return s.ShopAccountRepository.GetUserDetailsByUsername(ctx,username)

}

func (s *ShopAccountUsecase) FetchShops(ctx echo.Context) ([]*domain.ShopAccount, error) {
	return s.ShopAccountRepository.FetchShops(ctx)
}

func NewShopAccoutUsecase(r domain.ShopAccountRepository) domain.ShopAccountUsecase {
	return &ShopAccountUsecase{r}
}

func (s *ShopAccountUsecase) Fetch(ctx echo.Context) ([]*domain.ShopAccount, error) {
	return s.ShopAccountRepository.Fetch()
}

func (s *ShopAccountUsecase) GetByID(ctx echo.Context, id uint) (*domain.ShopAccount, error) {
	return s.ShopAccountRepository.GetByID(id)
}

func (s *ShopAccountUsecase) Update(ctx echo.Context, account *domain.ShopAccount) (*domain.ShopAccount, error) {
	return s.ShopAccountRepository.Update(account)
}

func (s *ShopAccountUsecase) Create(ctx echo.Context, account *domain.ShopAccount) (*domain.ShopAccount, error) {
	return s.ShopAccountRepository.Create(account)
}

func (s *ShopAccountUsecase) Delete(ctx echo.Context, id uint) error {
	return s.ShopAccountRepository.Delete(id)
}

