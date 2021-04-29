package datastore

import (
	"gorm.io/gorm"
	"web-shop/domain"
)

type categoryRepository struct {
	Conn *gorm.DB
}

func (c *categoryRepository) Fetch() ([]*domain.Category, error) {
	var (
		categories []*domain.Category
		err   error
	)
	err = c.Conn.Order("id desc").Find(&categories).Error
	return categories, err
}

func (c *categoryRepository) GetByID(id uint) (*domain.Category, error) {
	cat := &domain.Category{Model: gorm.Model{ID: id}}
	err := c.Conn.First(&cat).Error
	return cat, err
}

func (c *categoryRepository) Update(cat *domain.Category) (*domain.Category, error) {
	err := c.Conn.Save(cat).Error
	return cat, err
}

func (c *categoryRepository) Create(cat *domain.Category) (*domain.Category, error) {
	err := c.Conn.Create(cat).Error
	return cat, err
}

func (c *categoryRepository) Delete(id uint) error {
	ord := &domain.Category{Model: gorm.Model{ID: id}}
	err := c.Conn.Delete(ord).Error
	return err
}

func (c *categoryRepository) GetByName(name string)  (*domain.Category, error) {

	var(
		category *domain.Category
		err error
		)

	err = c.Conn.Where("name = ? ", name).Find(&category).Error

	return category, err
}

func NewCategoryRepository(conn *gorm.DB) domain.CategoryRepository {
	return &categoryRepository{conn}
}