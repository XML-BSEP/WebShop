package datastore

import (
	"gorm.io/gorm"
	"web-shop/domain"
)

type addressRepository struct {
	Conn *gorm.DB
}

func (a *addressRepository) GetByID(id uint) (*domain.Address, error) {
	adr := &domain.Address{ID: id}
	err := a.Conn.First(adr).Error
	return adr, err
}

func (a *addressRepository) Update(adr *domain.Address) (*domain.Address, error) {
	err := a.Conn.Save(adr).Error
	return adr, err
}

func (a *addressRepository) Create(adr *domain.Address) (*domain.Address, error) {
	err := a.Conn.Create(adr).Error
	return adr, err}

func (a *addressRepository) Delete(id uint) error {
	adr := &domain.Address{ID: id}
	err := a.Conn.Delete(adr).Error
	return err
}

func (a *addressRepository) Fetch() ([]*domain.Address, error) {
	var (
		addresses []*domain.Address
		err   error
	)
	err = a.Conn.Order("id desc").Find(&addresses).Error
	return addresses, err
}


func NewAddressRepository(Conn *gorm.DB) domain.AddressRepository {
	return &addressRepository{Conn}
}