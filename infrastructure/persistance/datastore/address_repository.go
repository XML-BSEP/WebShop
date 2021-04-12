package datastore

import(
	"context"
	"gorm.io/gorm"
	"web-shop/domain"
)

type addressRepository struct {
	Conn *gorm.DB
}

func newAddressRepository(Conn *gorm.DB) domain.AddressRepository {
	return &addressRepository{Conn}
}
func (a *addressRepository) GetByID(ctx context.Context, id uint) (*domain.Address, error) {
	adr := &domain.Address{ID: id}
	err := a.Conn.First(adr).Error
	return adr, err
}

func (a *addressRepository) Update(ctx context.Context, adr *domain.Address) (*domain.Address, error) {
	err := a.Conn.Save(adr).Error
	return adr, err
}

func (a *addressRepository) Create(ctx context.Context, adr *domain.Address) (*domain.Address, error) {
	err := a.Conn.Create(adr).Error
	return adr, err}

func (a *addressRepository) Delete(ctx context.Context, id uint) error {
	adr := &domain.Address{ID: id}
	err := a.Conn.Delete(adr).Error
	return err
}

func (a *addressRepository) Fetch(ctx context.Context) ([]*domain.Address, error) {
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