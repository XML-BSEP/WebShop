package datastore

import (
	"context"
	"gorm.io/gorm"
	"web-shop/domain"
)

type productRepository struct {
	Conn *gorm.DB
}

func (p *productRepository) Fetch(ctx context.Context) ([]*domain.Product, error) {
	var (
		products []*domain.Product
		err error
	)

	err = p.Conn.Order("id desc").Find(&products).Error
	return products, err
}

func (p *productRepository) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	err := p.Conn.Save(product).Error
	return product, err
}

func (p *productRepository) Create(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	err := p.Conn.Create(product).Error
	return  product, err
}

func (p *productRepository) Delete(ctx context.Context, id uint) error {

	product := &domain.Product{Model: gorm.Model{ID: id}}
	err := p.Conn.Delete(product).Error
	return err
}

func (p *productRepository) GetByID(ctx context.Context, id uint) (*domain.Product, error) {
	product := &domain.Product{Model: gorm.Model{ID: id}}
	err := p.Conn.First(product).Error
	return product, err
}

func NewProductRepository(Conn *gorm.DB) domain.ProductRepository {
	return &productRepository{Conn}
}

