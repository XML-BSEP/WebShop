package datastore

import (
	"gorm.io/gorm"
	"web-shop/domain"
)

type productRepository struct {
	Conn *gorm.DB
}

func (p *productRepository) Fetch() ([]*domain.Product, error) {
	var (
		products []*domain.Product
		err error
	)

	err = p.Conn.Order("id desc").Find(&products).Error
	return products, err
}

func (p *productRepository) GetWithPriceRange(low uint, high uint) ([]*domain.Product, error){
	var(
		products []*domain.Product
		err error
	)
	limit := 2
	offset := 0
	err = p.Conn.Limit(limit).Offset(offset).Where("price > ? and price < ? ", low, high).Find(&products).Error

	return products, err
}

func (p *productRepository) Update(product *domain.Product) (*domain.Product, error) {
	err := p.Conn.Save(product).Error
	return product, err
}

func (p *productRepository) Create(product *domain.Product) (*domain.Product, error) {
	err := p.Conn.Create(product).Error
	return  product, err
}

func (p *productRepository) Delete(id uint) error {

	product := &domain.Product{Model: gorm.Model{ID: id}}
	err := p.Conn.Delete(product).Error
	return err
}

func (p *productRepository) GetByID(id uint) (*domain.Product, error) {
	product := &domain.Product{Model: gorm.Model{ID: id}}
	err := p.Conn.First(product).Error
	return product, err
}

func NewProductRepository(Conn *gorm.DB) domain.ProductRepository {
	return &productRepository{Conn}
}

