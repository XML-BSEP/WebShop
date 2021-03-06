package datastore

import (
	"context"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"web-shop/domain"
)

type productRepository struct {
	Conn *gorm.DB
}

func (p *productRepository) GetProductDetails(ctx context.Context, productId uint) (*domain.Product, error) {
	var (
		product *domain.Product
		err error
	)

	err = p.Conn.Preload("Images").
		Joins("JOIN images on images.product_id = products.id").
		Where("products.id=?", productId).
		Take(&product).Error

	return product, err
}

func (p *productRepository) GetAllAvailableProductsInUsersShop(ctx echo.Context, userId uint) ([]*domain.Product, error) {
	var (
		products []*domain.Product
		err error
	)

	err = p.Conn.Preload("Images").Joins("Category").Order("id desc").Where("shop_account_id = ? and available>0", userId).Find(&products).Error
	return products, err
}
func (p *productRepository) GetAllProductsInUsersShop(ctx echo.Context, userId uint) ([]*domain.Product, error) {
	var (
		products []*domain.Product
		err error
	)

	err = p.Conn.Preload("Images").Joins("Category").Order("id desc").Where("shop_account_id = ?", userId).Find(&products).Error
	return products, err
}

func (p *productRepository) GetBySerialAndUserId(serial uint64, id uint) (*domain.Product, error) {
	var product *domain.Product

	err := p.Conn.Where("serial_number = ? and shop_account_id=?", serial, id).Take(&product).Error

	return product, err}

func (p *productRepository) MinMaxPrice() int64 {
	panic("implement me")
}

func (p *productRepository) Count() (int64, error) {
	var count int64
	err := p.Conn.Model(&domain.Product{}).Count(&count).Error
	return count, err
}

func (p *productRepository) GetBySerial(serial uint64) (*domain.Product, error) {

	var product *domain.Product

	err := p.Conn.Where("serial_number = ?", serial).Take(&product).Error

	return product, err
}

func (p *productRepository) FilterByCategory(userid uint,name string, category string, priceRangeStart uint, priceRangeEnd uint, limit int, offset int, order string) ([]*domain.Product, error) {
	var (
		products []*domain.Product
		err error
	)

	err = p.Conn.Preload("Images").
		Joins("JOIN categories on products.category_id = categories.id and lower(categories.name) like lower(?)", category).
		Where("lower(products.name) LIKE lower(?) and shop_account_id=?", name, userid).
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&products).Error

	return products, err
}

func (p *productRepository) GetProductsWithConditionOrderedByPrice(low uint, high uint, category string, limit int, offset int, order int) ([]*domain.Product, error) {
	var(
		products []*domain.Product
		err error
	)
	if len(category)<0 {
		if order == 1{
			err = p.Conn.Order("price asc").Limit(limit).Offset(offset).Where("price > ? and price < ? order by price asc", low, high).Find(&products).Error
		}else if order == 0{
			err = p.Conn.Order("price desc").Limit(limit).Offset(offset).Where("price > ? and price < ? order by price asc", low, high).Find(&products).Error
		}else{
			return nil, nil
		}
		return products, err
	}else{
		if order == 1{
			err = p.Conn.Order("price asc").Limit(limit).Offset(offset).Where("price > ? and price < ? and category = ?", low, high, category).Find(&products).Error
		}else if order == 0{
			err = p.Conn.Order("price desc").Limit(limit).Offset(offset).Where("price > ? and price < ? and category = ?", low, high, category).Find(&products).Error
		}else{
			return nil, nil
		}
		return products, err
	}
}

func (p *productRepository) GetProductsWithConditionOrderedByName(low uint, high uint, category string, limit int, offset int, order int) ([]*domain.Product, error) {
	var(
		products []*domain.Product
		err error
	)
	if len(category)<0 {
		if order == 1{
			err = p.Conn.Order("name asc").Limit(limit).Offset(offset).Where("price > ? and price < ? order by price asc", low, high).Find(&products).Error
		}else if order == 0{
			err = p.Conn.Order("name desc").Limit(limit).Offset(offset).Where("price > ? and price < ? order by price asc", low, high).Find(&products).Error
		}else{
			return nil, nil
		}
		return products, err
	}else{
		if order == 1{
			err = p.Conn.Order("name asc").Limit(limit).Offset(offset).Where("price > ? and price < ? and category = ?", low, high, category).Find(&products).Error
		}else if order == 0{
			err = p.Conn.Order("name desc").Limit(limit).Offset(offset).Where("price > ? and price < ? and category = ?", low, high, category).Find(&products).Error
		}else{
			return nil, nil
		}
		return products, err
	}}

func (p *productRepository) GetByNameOrderByPrice(name string, limit int, offset int, order int) ([]*domain.Product, error) {
	var(
		products []*domain.Product
		err error
	)
	if order == 1{
		err = p.Conn.Order("price asc").Limit(limit).Offset(offset).Where("LOWER(name) LIKE ?", name).Find(&products).Error
	}else if order == 0{
		err = p.Conn.Order("price desc").Limit(limit).Offset(offset).Where("LOWER(name) LIKE ?", name).Find(&products).Error
	}else{
		return nil, nil
	}
	return products, err
}

func (p *productRepository) GetByNameOrderByName(name string, limit int, offset int, order int) ([]*domain.Product, error) {
	var(
		products []*domain.Product
		err error
	)
	if order == 1{
		err = p.Conn.Order("name asc").Limit(limit).Offset(offset).Where("LOWER(name) LIKE ?", name).Find(&products).Error
	}else if order == 0{
		err = p.Conn.Order("name desc").Limit(limit).Offset(offset).Where("LOWER(name) LIKE ?", name).Find(&products).Error
	}else{
		return nil, nil
	}
	return products, err}

func (p *productRepository) GetByName(name string, limit int, offset int) ([]*domain.Product, error) {
	var(
		products []*domain.Product
		err error
	)

	err = p.Conn.Limit(limit).Offset(offset).Where("LOWER(name) LIKE ?", name).Find(&products).Error
	return products, err
}

func (p *productRepository) GetProductsWithCondition(low uint, high uint, category string, limit int, offset int) ([]*domain.Product, error) {
	var(
		products []*domain.Product
		err error
	)
	if len(category)<0 {
		err = p.Conn.Limit(int(limit)).Offset(int(offset)).Where("price > ? and price < ? ", low, high).Find(&products).Error
		return products, err
	}else{
		err = p.Conn.Limit(int(limit)).Offset(int(offset)).Where("price > ? and price < ? and category = ?", low, high, category).Find(&products).Error
		return products, err
	}

}

func (p *productRepository) GetProductsWithCategory(category string) ([]*domain.Product, error) {
	var(
		products []*domain.Product
		err error
	)

	err = p.Conn.Where("category = ? ", category).Find(&products).Error

	return products, err
}

func (p *productRepository) Fetch() ([]*domain.Product, error) {
	var (
		products []*domain.Product
		err error
	)

	err = p.Conn.Preload("Images").Joins("Category").Order("id desc").Find(&products).Error
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

