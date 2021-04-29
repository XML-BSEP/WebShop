package datastore

import (
	"gorm.io/gorm"
	"web-shop/domain"
)

type shopAccountRepository struct {
	Conn *gorm.DB
}

func (s shopAccountRepository) GetUserDetailsByUsername(account *domain.ShopAccount) (*domain.ShopAccount, error) {
	user := &domain.ShopAccount{}
	err := s.Conn.Where("username = ?", account.Username).Take(&user).Error
	return user, err
}

func (s shopAccountRepository) Fetch() ([]*domain.ShopAccount, error) {
	var(
		 accounts []*domain.ShopAccount
		 err error
	)

	err = s.Conn.Order("id desc").Find(&accounts).Error

	return accounts, err
}

func (s shopAccountRepository) GetByID(id uint) (*domain.ShopAccount, error) {
	account:=&domain.ShopAccount{Model: gorm.Model{ID: id}}

	err := s.Conn.First(account).Error
	return account, err
}

func (s shopAccountRepository) Update(account *domain.ShopAccount) (*domain.ShopAccount, error) {
	err := s.Conn.Save(account).Error
	return account, err
}

func (s shopAccountRepository) Create(account *domain.ShopAccount) (*domain.ShopAccount, error) {
	err := s.Conn.Create(account).Error
	return account, err
}

func (s shopAccountRepository) Delete(id uint) error {
	acc:=&domain.ShopAccount{Model: gorm.Model{ID: id}}
	err := s.Conn.Delete(acc).Error
	return err
}


func NewShopAccountRepository(Conn *gorm.DB) domain.ShopAccountRepository{
	return &shopAccountRepository{Conn}
}