package usecase

import (
	"github.com/labstack/echo"
	"web-shop/domain"
)

type addressUseCase struct {
	AddressRepository domain.AddressRepository
}

func NewAddresUsecase(r domain.AddressRepository) domain.AddressUsecase {
	return &addressUseCase{r}
}

func (a *addressUseCase) Fetch(ctx echo.Context) ([]*domain.Address, error) {
	return a.AddressRepository.Fetch()
}

func (a *addressUseCase) GetByID(ctx echo.Context, id uint) (*domain.Address, error) {
	return a.AddressRepository.GetByID(id)
}

func (a *addressUseCase) Update(ctx echo.Context, adr *domain.Address) (*domain.Address, error) {
	return a.AddressRepository.Update(adr)
}

func (a *addressUseCase) Create(ctx echo.Context, adr *domain.Address) (*domain.Address, error) {
	return a.AddressRepository.Create(adr)
}

func (a *addressUseCase) Delete(ctx echo.Context, id uint) error {
	return a.AddressRepository.Delete(id)
}

