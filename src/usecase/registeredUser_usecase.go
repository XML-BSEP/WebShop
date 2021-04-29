package usecase

import (
	"fmt"
	"github.com/labstack/echo"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
	"web-shop/security/password-verification"
)
const (
	userNotFound = "user not found"
	emailNotSent = "email not sent"
	invalidCode = "invalidCode"
	invalidPass = "password can't be the same as last one"
	hashError = "error while hashing pass"
	updateError = "error while updating user"
	redisError = "error while deleting redis key"
	passwordsError = "enter same passwords"

)

type registeredUserUsecase struct {
	RedisUsecase RedisUsecase
	RegisteredUserRepository domain.RegisteredShopUserRepository
}

func (r *registeredUserUsecase) SaveCodeToRedis(code string, email string) error {
	expiration  := 1000000000 * 3600 * 2 //2h
	return r.RedisUsecase.AddKeyValueSet(email, code, expiration)
}

func (r *registeredUserUsecase) ResendResetCode(email string, code string) error {
	if !r.RedisUsecase.ExistsByKey(email) {
		return fmt.Errorf("invalid email")
	}

	r.SaveCodeToRedis(code, email)

	return nil



}
func (r *registeredUserUsecase) ResetPassword(dto dto.ResetPassDTO) string {

	if passwordCompare := dto.Password == dto.ConfirmedPassword; !passwordCompare {
		return passwordsError
	}

	user, err := r.RegisteredUserRepository.ExistByUsernameOrEmail("", dto.Email)

	if err != nil {
		return userNotFound
	}

	account, err := r.RegisteredUserRepository.GetAccountDetailsFromUser(user)

	if err != nil {
		return userNotFound
	}

	codeValue, err := r.RedisUsecase.GetValueByKey(dto.Email)

	if err != nil {
		return emailNotSent
	}

	err = password_verification.VerifyPassword(dto.VerificationCode, codeValue)
	if err != nil {
		return invalidCode
	}

	err = password_verification.VerifyPassword(dto.Password, account.Password)
	if err == nil {
		return invalidPass
	}

	newPass, err := password_verification.Hash(dto.Password)

	if err != nil {
		return hashError
	}

	account.Password = string(newPass)

	err = r.RegisteredUserRepository.SaveNewPassword(account)

	if err != nil {
		return updateError
	}

	err = r.RedisUsecase.DeleteValueByKey(dto.Email)

	if err != nil {
		return redisError
	}

	return ""

}

func NewRegisteredShopUserUsecase(r domain.RegisteredShopUserRepository, usecase RedisUsecase) domain.RegisteredShopUserUsecase {
	return &registeredUserUsecase{RegisteredUserRepository: r, RedisUsecase: usecase}
}

func (r *registeredUserUsecase) ExistByUsernameOrEmail(ctx echo.Context, username string, email string) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.ExistByUsernameOrEmail(username, email)
}

func (r *registeredUserUsecase) Fetch(ctx echo.Context) ([]*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.Fetch()
}

func (r *registeredUserUsecase) GetByID(ctx echo.Context, id uint) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.GetByID(id)
}

func (r *registeredUserUsecase) Update(ctx echo.Context, reg *domain.RegisteredShopUser) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.Update(reg)
}

func (r *registeredUserUsecase) Create(ctx echo.Context, reg *domain.RegisteredShopUser) (*domain.RegisteredShopUser, error) {
	return r.RegisteredUserRepository.Create(reg)
}

func (r *registeredUserUsecase) Delete(ctx echo.Context, id uint) error {
	return r.RegisteredUserRepository.Delete(id)
}




