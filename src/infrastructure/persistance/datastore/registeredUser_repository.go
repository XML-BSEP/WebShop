package datastore

import (
	"gorm.io/gorm"
	"web-shop/domain"
)

type registeredUserRepository struct {
	Conn *gorm.DB
	ShopAccountRepository domain.ShopAccountRepository
}

func (r registeredUserRepository) GetByUsernameOrEmail(username string, email string) (*domain.RegisteredShopUser, error) {

	var newUser *domain.RegisteredShopUser

	err := r.Conn.Joins("ShopAccount").Find(&newUser, "username = ? or email = ?", username, email).Error

	return newUser, err

}

func (r registeredUserRepository) GetUserDetailsByAccount(account *domain.ShopAccount) (*domain.RegisteredShopUser, error) {
	realAccount, accErr := r.ShopAccountRepository.GetUserDetailsByUsername(account)

	if accErr != nil {
		return nil, accErr
	}

	user := &domain.RegisteredShopUser{}
	err := r.Conn.Where("shop_account_id = ?", realAccount.ID).Take(&user).Error
	return user, err
}

func (r registeredUserRepository) Fetch() ([]*domain.RegisteredShopUser, error) {
	var(
		users []*domain.RegisteredShopUser
		err error
	)

	err = r.Conn.Order("id desc").Find(&users).Error

	return users, err
}

func (r registeredUserRepository) GetByID(id uint) (*domain.RegisteredShopUser, error) {
	user:=&domain.RegisteredShopUser{Model: gorm.Model{ID: id}}
	err := r.Conn.First(user).Error
	return user, err
}

func (r registeredUserRepository) Update(reg *domain.RegisteredShopUser) (*domain.RegisteredShopUser, error) {
	err := r.Conn.Save(reg).Error
	return reg, err
}

func (r registeredUserRepository) Create(reg *domain.RegisteredShopUser) (*domain.RegisteredShopUser, error) {
	err := r.Conn.Create(reg).Error
	return reg, err
}

func (r registeredUserRepository) Delete(id uint) error {
	reg:=&domain.RegisteredShopUser{Model: gorm.Model{ID: id}}
	err := r.Conn.Delete(reg).Error
	return err
}

func NewRegisteredUserRepository(Conn *gorm.DB, shopAccountRepository domain.ShopAccountRepository) domain.RegisteredShopUserRepository{
	return &registeredUserRepository{Conn, shopAccountRepository}
}