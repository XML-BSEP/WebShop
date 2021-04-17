package usecase

import (
	"github.com/labstack/echo"
	"web-shop/domain"
)

type registeredUserUsecase struct {
	RegisteredUserRepository domain.RegisteredShopUserRepository
}



func NewRegisteredShopUserUsecase(r domain.RegisteredShopUserRepository) domain.RegisteredShopUserUsecase {
	return &registeredUserUsecase{r}
}

func (r *registeredUserUsecase) ExistByUsernameOrEmail(ctx echo.Context, username string, email string) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.ExistByUsernameOrEmail(username, email)
}

func (r *registeredUserUsecase) Fetch(ctx echo.Context) ([]*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.Fetch()
}

func (r *registeredUserUsecase) GetByID(ctx echo.Context, id uint) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.GetByID(id)
}

func (r *registeredUserUsecase) Update(ctx echo.Context, reg *domain.RegisteredShopUser) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.Update(reg)
}

func (r *registeredUserUsecase) Create(ctx echo.Context, reg *domain.RegisteredShopUser) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.Create(reg)
}

func (r *registeredUserUsecase) Delete(ctx echo.Context, id uint) error {
	return r.RegisteredUserRepository.Delete(id)
}




