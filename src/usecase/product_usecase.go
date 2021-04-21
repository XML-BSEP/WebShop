package usecase

import (
	"github.com/labstack/echo"
	"web-shop/domain"
)

type productUseCase struct {
	ProductRepository domain.ProductRepository
}
func (p *productUseCase) GetWithPriceRange(low uint, high uint) ([]*domain.Product, error){
	return p.ProductRepository.GetWithPriceRange(low, high)
}
func (p *productUseCase) Fetch(ctx echo.Context) ([]*domain.Product, error) {
	return p.ProductRepository.Fetch()
}

func (p *productUseCase) GetByID(ctx echo.Context, id uint) (*domain.Product, error) {
	return p.ProductRepository.GetByID(id)
}

func (p *productUseCase) Update(ctx echo.Context, pic *domain.Product) (*domain.Product, error) {
	return p.ProductRepository.Update(pic)
}

func (p *productUseCase) Create(ctx echo.Context, pic *domain.Product) (*domain.Product, error) {
	return p.ProductRepository.Create(pic)
}

func (p *productUseCase) Delete(ctx echo.Context, id uint) error {
	return p.ProductRepository.Delete(id)
}


func NewProductUseCase(p domain.ProductRepository) domain.ProductUsecase {
	return &productUseCase{p}
}
