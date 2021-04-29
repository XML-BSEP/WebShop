package datastore

import (
	"gorm.io/gorm"
	"web-shop/domain"
)

type storageRepository struct {
	Conn *gorm.DB
}

func (s *storageRepository) Fetch() ([]*domain.Storage, error) {
	var (
		storages []*domain.Storage
		err error
	)

	err = s.Conn.Order("id desc").Find(&storages).Error
	return storages, err
}

func (s *storageRepository) Update(storage *domain.Storage) (*domain.Storage, error) {
	err := s.Conn.Save(storage).Error
	return storage, err
}

func (s *storageRepository) Create(storage *domain.Storage) (*domain.Storage, error) {
	err := s.Conn.Create(storage).Error
	return  storage, err
}

func (s *storageRepository) Delete(id uint) error {

	storage := &domain.Storage{Model: gorm.Model{ID: id}}
	err := s.Conn.Delete(storage).Error
	return err
}

func (s *storageRepository) GetByID(id uint) (*domain.Storage, error) {
	storage := &domain.Storage{Model: gorm.Model{ID: id}}
	err := s.Conn.First(storage).Error
	return storage, err
}

func NewStorageRepository(Conn *gorm.DB) domain.StorageRepository {
	return &storageRepository{Conn}
}


