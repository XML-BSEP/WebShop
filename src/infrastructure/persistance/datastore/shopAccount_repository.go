package datastore

import (
	"context"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"web-shop/domain"
)

type shopAccountRepository struct {
	Conn *gorm.DB
}


func (s shopAccountRepository) FetchShops(ctx echo.Context) ([]*domain.ShopAccount, error) {
	var (
		shops []*domain.ShopAccount
		err error
	)

	err = s.Conn.Joins("JOIN registered_shop_users on registered_shop_users.shop_account_id=shop_accounts.id").Where("role_id=1").Find(&shops).Error
	return shops, err
}

func (s shopAccountRepository) GetUserDetailsByUsername(ctx context.Context, username string) (*domain.ShopAccount, error) {
	user := &domain.ShopAccount{}
	err := s.Conn.Where("username = ?", username).Take(&user).Error
	return user, err
}

func (s shopAccountRepository) GetUserDetailsByEmail(ctx context.Context, email string) (*domain.ShopAccount, error) {
	user := &domain.ShopAccount{}
	err := s.Conn.Joins("JOIN registered_shop_users on registered_shop_users.shop_account_id=shop_accounts.id").Where("registered_shop_users.email=?", email).Take(&user).Error
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