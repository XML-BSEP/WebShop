package datastore

import (
	"context"
	"gorm.io/gorm"
	"web-shop/domain"
)

type shoppingCartItemRepository struct {
	Conn *gorm.DB
}

func (s *shoppingCartItemRepository) Fetch(ctx context.Context) ([]*domain.ShoppingCartItem, error) {
	var (
		cartItems []*domain.ShoppingCartItem
		err error
	)

	err = s.Conn.Order("id desc").Find(&cartItems).Error
	return cartItems, err
}

func (s *shoppingCartItemRepository) Update(ctx context.Context, cartItem *domain.ShoppingCartItem) (*domain.ShoppingCartItem, error) {
	err := s.Conn.Save(cartItem).Error
	return cartItem, err
}

func (s *shoppingCartItemRepository) Create(ctx context.Context, cartItem *domain.ShoppingCartItem) (*domain.ShoppingCartItem, error) {
	err := s.Conn.Create(cartItem).Error
	return  cartItem, err
}

func (s *shoppingCartItemRepository) Delete(ctx context.Context, id uint) error {

	cartItem := &domain.ShoppingCartItem{Model: gorm.Model{ID: id}}
	err := s.Conn.Delete(cartItem).Error
	return err
}

func (s *shoppingCartItemRepository) GetByID(ctx context.Context, id uint) (*domain.ShoppingCartItem, error) {
	cartItem := &domain.ShoppingCartItem{Model: gorm.Model{ID: id}}
	err := s.Conn.First(cartItem).Error
	return cartItem, err
}

func NewShoppingCartItemRepository(Conn *gorm.DB) domain.ShoppingCartItemRepository {
	return &shoppingCartItemRepository{Conn}
}

