package interactor

import (
	"gorm.io/gorm"
	"web-shop/domain"
	"web-shop/http/handler"
	"web-shop/infrastructure/persistance/datastore"
	"web-shop/infrastructure/redisdb"
	auth2 "web-shop/security/auth"
	"web-shop/usecase"
)

type Interactor interface {
	NewAddressRepository() domain.AddressRepository
	NewAddressUsecase() domain.AddressUsecase
	NewAddressHandler() handler.AddressHandler
	NewPersonRepository() domain.PersonRepository
	NewShopAccountRepository() domain.ShopAccountRepository
	NewRegisteredUserRepository(repository domain.ShopAccountRepository) domain.RegisteredShopUserRepository
	NewPersonUsecase() domain.PersonUsecase
	NewTokenService() *auth2.Token
	NewPersonHandler() handler.PersonHandler
	NewAuthenticateHandler() handler.AuthenticateHandler
	NewRedisUsecase() usecase.RedisUsecase
	NewRedisHandler() handler.RedisHandlerSample
	NewAppHandler() handler.AppHandler
	NewSigUpUsecase() usecase.SignUpUseCase
	NewRegisterUserUsecase() usecase.RegisterUserUsecase
	NewRandomStringGeneratorUsecase() usecase.RandomStringGeneratorUsecase
}

type interactor struct {
	Conn *gorm.DB

}



type appHandler struct {
	handler.AddressHandler
	handler.PersonHandler
	handler.AuthenticateHandler
	handler.SignUpHandler
	handler.RedisHandlerSample
}

func NewInteractor(conn *gorm.DB) Interactor {
	return &interactor{conn}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.AddressHandler = i.NewAddressHandler()
	appHandler.PersonHandler = i.NewPersonHandler()
	appHandler.AuthenticateHandler = i.NewAuthenticateHandler()
	appHandler.SignUpHandler = i.NewSignUpHandler()
	appHandler.RedisHandlerSample = i.NewRedisHandler()
	return appHandler
}


func (i *interactor) NewAuthenticateHandler() handler.AuthenticateHandler {
	shopAccountRepo := i.NewShopAccountRepository()
	userRepo := i.NewRegisteredUserRepository(shopAccountRepo)
	tk := i.NewTokenService()

	return handler.NewAuthenticate(userRepo, tk)
}

func (i *interactor) NewTokenService() *auth2.Token {
	tk := auth2.NewToken()
	return tk
}

func (i *interactor) NewRegisteredUserRepository(shopAccountRepo domain.ShopAccountRepository) domain.RegisteredShopUserRepository {
	userRepo := datastore.NewRegisteredUserRepository(i.Conn, shopAccountRepo)
	return userRepo
}

func (i *interactor) NewShopAccountRepository() domain.ShopAccountRepository {
	shopAccountRepo := datastore.NewShopAccountRepository(i.Conn)
	return shopAccountRepo
}

func (i *interactor) NewPersonRepository() domain.PersonRepository {
	return datastore.NewPersonRepository(i.Conn)
}

func (i *interactor) NewPersonUsecase() domain.PersonUsecase {
	return usecase.NewPersonUsecase(i.NewPersonRepository())
}

func (i *interactor) NewPersonHandler() handler.PersonHandler {
	return handler.NewPersonHandler(i.NewPersonUsecase())
}

func (i *interactor) NewAddressUsecase() domain.AddressUsecase {
	return usecase.NewAddresUsecase(i.NewAddressRepository())
}

func (i *interactor) NewAddressHandler() handler.AddressHandler {
	return handler.NewAddressHandler(i.NewAddressUsecase())
}

func (i *interactor) NewAddressRepository() domain.AddressRepository {
	return datastore.NewAddressRepository(i.Conn)
}

func (i *interactor) NewSignUpHandler() handler.SignUpHandler {
	return handler.NewSignUpHandler(i.NewRegisteredUserRepository(i.NewShopAccountRepository()), i.NewSigUpUsecase())
}

func (i *interactor) NewRedisUsecase() usecase.RedisUsecase {
	redis := redisdb.NewReddisConn()
	return usecase.NewRedisUsecase(redis)
}

func (i *interactor) NewRedisHandler() handler.RedisHandlerSample {
	return handler.NewRedisHandlerSample(i.NewRedisUsecase())
}

func (i *interactor) NewRegisterUserUsecase() usecase.RegisterUserUsecase {
	return usecase.NewRegisteredUserUsecase(i.NewRegisteredUserRepository(i.NewShopAccountRepository()))
}


func (i *interactor) NewRandomStringGeneratorUsecase() usecase.RandomStringGeneratorUsecase {
	return usecase.NewRandomStringGenrator()
}


func (i *interactor) NewSigUpUsecase() usecase.SignUpUseCase {
	return usecase.NewSignUpUsecase(i.NewRedisUsecase(), i.NewRegisterUserUsecase(), i.NewRandomStringGeneratorUsecase())
}



