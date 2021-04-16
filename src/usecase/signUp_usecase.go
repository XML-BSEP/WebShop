package usecase

import (
	"encoding/json"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

type SignUpUseCase interface {
	RegisterNewUser(ctx echo.Context, newUser domain.UserRegistrationRequest) (string, error)
	CheckIfExistUser(ctx echo.Context, newUser dto.NewUser) (*domain.RegisteredShopUser, error)
	Hash(password string) ([]byte, error)
}

type signUp struct {
	RedisUsecase                 RedisUsecase
	RegisteredUserUsecase        RegisterUserUsecase
	RandomStringGeneratorUSecase RandomStringGeneratorUsecase
}

func (s *signUp) Hash(password string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (s *signUp) CheckIfExistUser(ctx echo.Context, newUser dto.NewUser) (*domain.RegisteredShopUser, error) {

	acc, err  := s.RegisteredUserUsecase.GetByUsernameOrEmail(ctx, newUser.Username, newUser.Email)
	if err != nil {
		return nil, err
	}
	return acc, err

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
	errR := s.RedisUsecase.AddKeyValueSet(newUser.Username, string(newAcc), expiration)
	if errR != nil {
		return "", errR
	}

	return code, errR

}

func NewSignUpUsecase (redisUsecase RedisUsecase, userUsecase RegisterUserUsecase, generatorUsecase RandomStringGeneratorUsecase) SignUpUseCase {
	return &signUp{redisUsecase, userUsecase, generatorUsecase}
}




