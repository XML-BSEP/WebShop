package usecase

import (
	"github.com/labstack/echo"
	"web-shop/domain"
)

type orderUseCase struct {
	OrderRepository domain.OrderRepository
}

func (o *orderUseCase) Fetch(ctx echo.Context) ([]*domain.Order, error) {
	return o.OrderRepository.Fetch()
}

func (o *orderUseCase) GetByID(ctx echo.Context, id uint) (*domain.Order, error) {
	return o.OrderRepository.GetByID(id)
}

func (o *orderUseCase) Update(ctx echo.Context, order *domain.Order) (*domain.Order, error) {
	return o.OrderRepository.Update(order)
}

func (o *orderUseCase) Create(ctx echo.Context, order *domain.Order) (*domain.Order, error) {
	return o.OrderRepository.Create(order)
}

func (o *orderUseCase) Delete(ctx echo.Context, id uint) error {
	return o.OrderRepository.Delete(id)
}

func NewOrderUsecase(r domain.OrderRepository) domain.OrderUsecase {
	return &orderUseCase{r}
}