package datastore

import (
	"gorm.io/gorm"
	"web-shop/domain"
)

type orderRepository struct {
	Conn *gorm.DB
}

func (o *orderRepository) Fetch() ([]*domain.Order, error) {
	var (
		orders []*domain.Order
		err   error
	)
	err = o.Conn.Order("id desc").Find(&orders).Error
	return orders, err
}

func (o *orderRepository) GetByID(id uint) (*domain.Order, error) {
	ord := &domain.Order{Model: gorm.Model{ID: id}}
	err := o.Conn.First(ord).Error
	return ord, err
}

func (o *orderRepository) Update(order *domain.Order) (*domain.Order, error) {
	err := o.Conn.Save(order).Error
	return order, err
}

func (o *orderRepository) Create(order *domain.Order) (*domain.Order, error) {
	err := o.Conn.Create(order).Error
	return order, err
}

func (o *orderRepository) Delete(id uint) error {
	ord := &domain.Order{Model: gorm.Model{ID: id}}
	err := o.Conn.Delete(ord).Error
	return err
}

func NewOrderRepository(Conn *gorm.DB) domain.OrderRepository {
	return &orderRepository{Conn}
}