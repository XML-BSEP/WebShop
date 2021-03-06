package datastore

import (
	"gorm.io/gorm"
	"web-shop/domain"
)

type registeredUserRepository struct {
	Conn *gorm.DB
	ShopAccountRepository domain.ShopAccountRepository
}

func (r registeredUserRepository) SaveNewPassword(account *domain.ShopAccount) error {
	_, err := r.ShopAccountRepository.Update(account)
	return err
}

func (r registeredUserRepository) GetRoleById(id uint) (string, error) {
	user, err := r.GetByID(id)
	if err != nil {
		return "", err
	}
	return user.Role.RoleName, nil
}

func (r registeredUserRepository) GetAccountDetailsFromUser(u *domain.RegisteredShopUser) (*domain.ShopAccount, error) {
	user, err := r.ExistByUsernameOrEmail("", u.Email)
	account, err := r.ShopAccountRepository.GetByID(user.ShopAccountID)
	if err != nil {
		return nil, err
	}
	return account, nil


}

func (r registeredUserRepository) ExistByUsernameOrEmail(username string, email string) (*domain.RegisteredShopUser, error) {

	var newUser *domain.RegisteredShopUser

	err := r.Conn.Joins("ShopAccount").Take(&newUser, "username = ? or email = ?", username, email).Error

	return newUser, err

}

func (r registeredUserRepository) ExistByUsername(username string) (*domain.RegisteredShopUser, error) {

	var newUser *domain.RegisteredShopUser

	err := r.Conn.Joins("ShopAccount").Where("username = ?", username).Take(&newUser).Error

	return newUser, err

}

func (r registeredUserRepository) ExistByEmail(email string) (*domain.RegisteredShopUser, error) {

	var newUser *domain.RegisteredShopUser

	err := r.Conn.Where("email = ?", email).Take(&newUser).Error

	return newUser, err

}

func (r registeredUserRepository) GetUserDetailsFromEmail(email string) (*domain.RegisteredShopUser, error) {

	user := &domain.RegisteredShopUser{}
	err := r.Conn.Where("email = ?", email).Take(&user).Error
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
	err := r.Conn.Joins("Role").First(&user).Error
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