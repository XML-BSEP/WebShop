package usecase

import (
	"context"
	"github.com/labstack/echo"
	"web-shop/domain"
)

type registeredUserUsecase struct {
	RegisteredUserRepository domain.RegisteredShopUserRepository
}


type RegisterUserUsecase interface {
	GetByUsernameOrEmail(ctx echo.Context, username string, email string) (*domain.RegisteredShopUser, error)
	Fetch(ctx context.Context) ([]*domain.RegisteredShopUser, error)
	GetByID(ctx context.Context, id uint) (*domain.RegisteredShopUser, error)
	Update(ctx context.Context, reg *domain.RegisteredShopUser) (*domain.RegisteredShopUser, error)
	Create(ctx echo.Context, reg *domain.RegisteredShopUser) (*domain.RegisteredShopUser, error)
	Delete(ctx context.Context, id uint) error
}


func NewRegisteredUserUsecase(r domain.RegisteredShopUserRepository) RegisterUserUsecase {
	return &registeredUserUsecase{r}
}

func (r *registeredUserUsecase) GetByUsernameOrEmail(ctx echo.Context, username string, email string) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.GetByUsernameOrEmail(username, email)
}

func (r *registeredUserUsecase) Fetch(ctx context.Context) ([]*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.Fetch()
}

func (r *registeredUserUsecase) GetByID(ctx context.Context, id uint) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.GetByID(id)
}

func (r *registeredUserUsecase) Update(ctx context.Context, reg *domain.RegisteredShopUser) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.Update(reg)
}

func (r *registeredUserUsecase) Create(ctx echo.Context, reg *domain.RegisteredShopUser) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.Create(reg)
}

func (r *registeredUserUsecase) Delete(ctx context.Context, id uint) error {
	return r.RegisteredUserRepository.Delete(id)
}




