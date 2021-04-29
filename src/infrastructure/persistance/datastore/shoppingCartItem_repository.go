package datastore

import (
	"gorm.io/gorm"
	"web-shop/domain"
)

type shoppingCartItemRepository struct {
	Conn *gorm.DB
}

func (s *shoppingCartItemRepository) Fetch() ([]*domain.ShoppingCartItem, error) {
	var (
		cartItems []*domain.ShoppingCartItem
		err error
	)

	err = s.Conn.Order("id desc").Find(&cartItems).Error
	return cartItems, err
}

func (s *shoppingCartItemRepository) Update(cartItem *domain.ShoppingCartItem) (*domain.ShoppingCartItem, error) {
	err := s.Conn.Save(cartItem).Error
	return cartItem, err
}

func (s *shoppingCartItemRepository) Create(cartItem *domain.ShoppingCartItem) (*domain.ShoppingCartItem, error) {
	err := s.Conn.Create(cartItem).Error
	return  cartItem, err
}

func (s *shoppingCartItemRepository) Delete(id uint) error {

	cartItem := &domain.ShoppingCartItem{Model: gorm.Model{ID: id}}
	err := s.Conn.Delete(cartItem).Error
	return err
}

func (s *shoppingCartItemRepository) GetByID(id uint) (*domain.ShoppingCartItem, error) {
	cartItem := &domain.ShoppingCartItem{Model: gorm.Model{ID: id}}
	err := s.Conn.First(cartItem).Error
	return cartItem, err
}

func NewShoppingCartItemRepository(Conn *gorm.DB) domain.ShoppingCartItemRepository {
	return &shoppingCartItemRepository{Conn}
}

