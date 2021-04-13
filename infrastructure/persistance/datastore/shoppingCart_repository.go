package datastore

import (
	"context"
	"gorm.io/gorm"
	"web-shop/domain"
)

type shoppingCartReporsitory struct {
	Conn *gorm.DB
}

func (s *shoppingCartReporsitory) Fetch(ctx context.Context) ([]*domain.ShoppingCart, error) {
	var (
		shoppingCarts []*domain.ShoppingCart
		err error
	)

	err = s.Conn.Order("id desc").Find(&shoppingCarts).Error
	return shoppingCarts, err
}

func (s *shoppingCartReporsitory) Update(ctx context.Context, shoppingCart *domain.ShoppingCart) (*domain.ShoppingCart, error) {
	err := s.Conn.Save(shoppingCart).Error
	return shoppingCart, err
}

func (s *shoppingCartReporsitory) Create(ctx context.Context, shoppingCart *domain.ShoppingCart) (*domain.ShoppingCart, error) {
	err := s.Conn.Create(shoppingCart).Error
	return  shoppingCart, err
}

func (s *shoppingCartReporsitory) Delete(ctx context.Context, id uint) error {

	shoppingCart := &domain.ShoppingCart{Model: gorm.Model{ID: id}}
	err := s.Conn.Delete(shoppingCart).Error
	return err
}

func (s *shoppingCartReporsitory) GetByID(ctx context.Context, id uint) (*domain.ShoppingCart, error) {
	shoppingCart := &domain.ShoppingCart{Model: gorm.Model{ID: id}}
	err := s.Conn.First(shoppingCart).Error
	return shoppingCart, err
}

func NewShoppingCartRepository(Conn *gorm.DB) domain.ShoppingCartRepository {
	return &shoppingCartReporsitory{Conn}
}

