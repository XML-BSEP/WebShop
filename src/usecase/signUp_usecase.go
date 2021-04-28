package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
	password_verification "web-shop/security/password-verification"
)

type SignUpUseCase interface {
	RegisterNewUser(ctx echo.Context, newUser domain.UserRegistrationRequest) (string, error)
	CheckIfExistUser(ctx echo.Context, newUser dto.NewUser) (*domain.RegisteredShopUser, error)
	IsCodeValid(ctx echo.Context, email, code string) (*domain.UserRegistrationRequest, error)
	ConfirmAccount(ctx echo.Context, user *domain.UserRegistrationRequest) (*domain.RegisteredShopUser, error)
	Hash(password string) ([]byte, error)
	ValidatePassword(password, confirmPassword string) bool
	ResendCode(email string) (error, string, string)
}


type signUp struct {
	RedisUsecase                 RedisUsecase
	RegisteredUserUsecase        domain.RegisteredShopUserUsecase
	RandomStringGeneratorUSecase RandomStringGeneratorUsecase
}

func (s *signUp) ResendCode(email string) (error, string, string) {
	if !s.RedisUsecase.ExistsByKey(email) {
		return fmt.Errorf("invalid email"), "", ""
	}
	userJson, err := s.RedisUsecase.GetValueByKey(email)

	if err != nil {
		return err, "", ""
	}

	var userObj domain.UserRegistrationRequest

	err = json.Unmarshal([]byte(userJson), &userObj)

	code := s.RandomStringGeneratorUSecase.RandomStringGenerator(8)

	expiration  := 1000000000 * 3600 * 2 //2h
	hash, _ := password_verification.Hash(code)

	userObj.VerificationCode = string(hash)

	newAcc, err := json.Marshal(userObj)
	if err != nil {
		return err, "", ""
	}

	s.RedisUsecase.AddKeyValueSet(email, string(newAcc), expiration)

	return nil, userObj.Username, code
}

func (s *signUp) IsCodeValid(ctx echo.Context, email, code string) (*domain.UserRegistrationRequest, error) {

	userJson, err := s.RedisUsecase.GetValueByKey(email)

	if err != nil {
		return nil, err
	}

	var userObj domain.UserRegistrationRequest

	err = json.Unmarshal([]byte(userJson), &userObj)

	hashedCode := userObj.VerificationCode

	err = VerifyCode(hashedCode, code)

	return &userObj, err
}

func (s *signUp) Hash(password string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (s *signUp) CheckIfExistUser(ctx echo.Context, newUser dto.NewUser) (*domain.RegisteredShopUser, error) {

	acc, err  := s.RegisteredUserUsecase.ExistByUsernameOrEmail(ctx, newUser.Username, newUser.Email)

	if s.RedisUsecase.CheckUsername(newUser.Username) {
		return nil, fmt.Errorf("already exists")
	}

	if s.RedisUsecase.ExistsByKey(newUser.Email) {
		return nil, fmt.Errorf("already exists")
	}

	if err != nil {
		return nil, nil
	}

	return acc, err

}

func (s *signUp) ConfirmAccount(ctx echo.Context, user *domain.UserRegistrationRequest) (*domain.RegisteredShopUser, error) {

	newUser := domain.RegisteredShopUser{
		Name: user.Name,
		Surname: user.Surname,
		Email: user.Email,
		ShopAccount: domain.ShopAccount{Username: user.Username, Password: user.Password},
		RoleId: 2,
	}

	s.RedisUsecase.DeleteValueByKey(newUser.Email)
	return s.RegisteredUserUsecase.Create(ctx, &newUser)



}

func (s *signUp) RegisterNewUser(ctx echo.Context, newUser domain.UserRegistrationRequest) (string, error) {

	newUser.VerificationCode = s.RandomStringGeneratorUSecase.RandomStringGenerator(8)

	code := newUser.VerificationCode

	hashPassword, _ := s.Hash(newUser.Password)
	newUser.Password = string(hashPassword)

	hashCode, _ := s.Hash(newUser.VerificationCode)
	newUser.VerificationCode = string(hashCode)

	newAcc, err := json.Marshal(newUser)
	if err != nil {
		return "", err
	}

	expiration  := 1000000000 * 3600 * 2 //2h
	errR := s.RedisUsecase.AddKeyValueSet(newUser.Email, string(newAcc), expiration)
	if errR != nil {
		return "", errR
	}

	return code, errR

}

func VerifyCode(hashedCode, code string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedCode), []byte(code))

}

func NewSignUpUsecase (redisUsecase RedisUsecase, userUsecase domain.RegisteredShopUserUsecase, generatorUsecase RandomStringGeneratorUsecase) SignUpUseCase {
	return &signUp{redisUsecase, userUsecase, generatorUsecase}
}

func (s *signUp) ValidatePassword(password, confirmPassword string) bool {
	return password == confirmPassword
}



