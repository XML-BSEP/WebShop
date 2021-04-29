package datastore

import (
	"gorm.io/gorm"
	"web-shop/domain"
)

type shoppingCartReporsitory struct {
	Conn *gorm.DB
}

func (s *shoppingCartReporsitory) Fetch() ([]*domain.ShoppingCart, error) {
	var (
		shoppingCarts []*domain.ShoppingCart
		err error
	)

	err = s.Conn.Order("id desc").Find(&shoppingCarts).Error
	return shoppingCarts, err
}

func (s *shoppingCartReporsitory) Update(shoppingCart *domain.ShoppingCart) (*domain.ShoppingCart, error) {
	err := s.Conn.Save(shoppingCart).Error
	return shoppingCart, err
}

func (s *shoppingCartReporsitory) Create(shoppingCart *domain.ShoppingCart) (*domain.ShoppingCart, error) {
	err := s.Conn.Create(shoppingCart).Error
	return  shoppingCart, err
}

func (s *shoppingCartReporsitory) Delete(id uint) error {

	shoppingCart := &domain.ShoppingCart{Model: gorm.Model{ID: id}}
	err := s.Conn.Delete(shoppingCart).Error
	return err
}

func (s *shoppingCartReporsitory) GetByID(id uint) (*domain.ShoppingCart, error) {
	shoppingCart := &domain.ShoppingCart{Model: gorm.Model{ID: id}}
	err := s.Conn.First(shoppingCart).Error
	return shoppingCart, err
}

func NewShoppingCartRepository(Conn *gorm.DB) domain.ShoppingCartRepository {
	return &shoppingCartReporsitory{Conn}
}

