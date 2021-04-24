package usecase

import (
	"github.com/labstack/echo"
	"web-shop/domain"
)

type productUseCase struct {
	ProductRepository domain.ProductRepository
}

func (p *productUseCase) FilterByCategory(category string, priceRangeStart uint, priceRangeEnd uint, limit int, offset int, order string) ([]*domain.Product, error) {
	return p.ProductRepository.FilterByCategory(category, priceRangeStart, priceRangeEnd, limit, offset, order)
}

func (p *productUseCase) GetProductsWithConditionOrderedByPrice(low uint, high uint, category string, limit int, offset int, order int) ([]*domain.Product, error) {
	return p.ProductRepository.GetProductsWithConditionOrderedByPrice(low,high,category,limit,offset,order)
}

func (p *productUseCase) GetProductsWithConditionOrderedByName(low uint, high uint, category string, limit int, offset int, order int) ([]*domain.Product, error) {
	return p.ProductRepository.GetProductsWithConditionOrderedByName(low,high,category,limit,offset,order)
}

func (p *productUseCase) GetByNameOrderByPrice(name string, limit int, offset int, order int) ([]*domain.Product, error) {
	return p.ProductRepository.GetByNameOrderByPrice(name, limit, offset, order)
}

func (p *productUseCase) GetByNameOrderByName(name string, limit int, offset int, order int) ([]*domain.Product, error) {
	return p.ProductRepository.GetByNameOrderByName(name, limit, offset, order)
}

func (p *productUseCase) GetByName(name string, limit int, offset int) ([]*domain.Product, error) {
	return p.ProductRepository.GetByName(name, limit, offset)
}

func (p *productUseCase) GetProductsWithCondition(low uint, high uint, category string, limit int, offset int) ([]*domain.Product, error) {
	return p.ProductRepository.GetProductsWithCondition(low, high, category, limit, offset)
}

func (p *productUseCase) GetProductsWithCategory(category string) ([]*domain.Product, error) {
	return p.ProductRepository.GetProductsWithCategory(category)
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
