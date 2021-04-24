package datastore

import (
	"gorm.io/gorm"
	"web-shop/domain"
)

type imageRepository struct {
	Conn *gorm.DB
}

func (p *imageRepository) Fetch() ([]*domain.Image, error) {
	var (
		images []*domain.Image
		err error
	)

	err = p.Conn.Order("id desc").Find(&images).Error
	return images, err
}

func (p *imageRepository) Update(image *domain.Image) (*domain.Image, error) {
	err := p.Conn.Save(image).Error
	return image, err
}

func (p *imageRepository) Create(image *domain.Image) (*domain.Image, error) {
	err := p.Conn.Create(image).Error
	return  image, err
}

func (p *imageRepository) Delete(id uint) error {

	image := &domain.Image{Model: gorm.Model{ID: id}}
	err := p.Conn.Delete(image).Error
	return err
}

func (p *imageRepository) GetByID(id uint) (*domain.Image, error) {
	image := &domain.Image{Model: gorm.Model{ID: id}}
	err := p.Conn.First(image).Error
	return image, err
}

func NewImageRepository(Conn *gorm.DB) domain.ImageRepository {
	return &imageRepository{Conn}
}


