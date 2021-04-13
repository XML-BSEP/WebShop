package datastore

import (
	"context"
	"gorm.io/gorm"
	"web-shop/domain"
)

type storageRepository struct {
	Conn *gorm.DB
}

func (s *storageRepository) Fetch(ctx context.Context) ([]*domain.Storage, error) {
	var (
		storages []*domain.Storage
		err error
	)

	err = s.Conn.Order("id desc").Find(&storages).Error
	return storages, err
}

func (s *storageRepository) Update(ctx context.Context, storage *domain.Storage) (*domain.Storage, error) {
	err := s.Conn.Save(storage).Error
	return storage, err
}

func (s *storageRepository) Create(ctx context.Context, storage *domain.Storage) (*domain.Storage, error) {
	err := s.Conn.Create(storage).Error
	return  storage, err
}

func (s *storageRepository) Delete(ctx context.Context, id uint) error {

	storage := &domain.Storage{Model: gorm.Model{ID: id}}
	err := s.Conn.Delete(storage).Error
	return err
}

func (s *storageRepository) GetByID(ctx context.Context, id uint) (*domain.Storage, error) {
	storage := &domain.Storage{Model: gorm.Model{ID: id}}
	err := s.Conn.First(storage).Error
	return storage, err
}

func NewStorageRepository(Conn *gorm.DB) domain.StorageRepository {
	return &storageRepository{Conn}
}


