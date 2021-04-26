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

	NewShopAccountRepository() domain.ShopAccountRepository
	NewRegisteredUserRepository(repository domain.ShopAccountRepository) domain.RegisteredShopUserRepository

	NewTokenService() *auth2.Token
	NewAuthenticateHandler() handler.AuthenticateHandler
	NewRedisUsecase() usecase.RedisUsecase
	NewRedisHandler() handler.RedisHandlerSample
	NewAppHandler() handler.AppHandler
	NewSigUpUsecase() usecase.SignUpUseCase
	NewRandomStringGeneratorUsecase() usecase.RandomStringGeneratorUsecase
	NewRegisteredShopUserUsecase() domain.RegisteredShopUserUsecase
	NewProductRepository() domain.ProductRepository
	NewProductUsecase() domain.ProductUsecase
	NewProductHandler() handler.ProductHandler
	NewOrderHandler() handler.OrderHandler
	NewOrderUsecase() domain.OrderUsecase
	NewOrderRepository() domain.OrderRepository
	NewResetPasswordHandler() handler.ResetPasswordHandler

	NewCategoryHandler() handler.CategoryHandler
	NewCategoryUsecase() domain.CategoryUsecase
	NewCategoryRepository() domain.CategoryRepository
}

type interactor struct {
	Conn *gorm.DB

}

type appHandler struct {
	handler.AddressHandler
	handler.AuthenticateHandler
	handler.SignUpHandler
	handler.RedisHandlerSample
	handler.ProductHandler
	handler.OrderHandler
	handler.ResetPasswordHandler
	handler.CategoryHandler
}


func NewInteractor(conn *gorm.DB) Interactor {
	return &interactor{conn}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.AddressHandler = i.NewAddressHandler()
	appHandler.AuthenticateHandler = i.NewAuthenticateHandler()
	appHandler.SignUpHandler = i.NewSignUpHandler()
	appHandler.RedisHandlerSample = i.NewRedisHandler()
	appHandler.ProductHandler = i.NewProductHandler()
	appHandler.OrderHandler = i.NewOrderHandler()
	appHandler.ResetPasswordHandler = i.NewResetPasswordHandler()
	appHandler.CategoryHandler = i.NewCategoryHandler()
	return appHandler
}
func (i *interactor) NewCategoryHandler() handler.CategoryHandler{
	return handler.NewCategoryHandler(i.NewCategoryUsecase())
}

func (i *interactor) NewCategoryUsecase() domain.CategoryUsecase {
	return usecase.NewCategoryUsecase(i.NewCategoryRepository())
}

func (i *interactor) NewCategoryRepository() domain.CategoryRepository {
	return datastore.NewCategoryRepository(i.Conn)
}


func (i *interactor) NewOrderHandler() handler.OrderHandler{
	return handler.NewOrderHandler(i.NewOrderUsecase())
}
func (i *interactor) NewOrderUsecase() domain.OrderUsecase {
	return usecase.NewOrderUsecase(i.NewOrderRepository())
}

func (i *interactor) NewOrderRepository() domain.OrderRepository {
	return datastore.NewOrderRepository(i.Conn)
}

func (i *interactor) NewProductUsecase() domain.ProductUsecase {
	return usecase.NewProductUseCase(i.NewProductRepository(), i.NewCategoryRepository(), i.NewImageRepository())
}
func (i *interactor) NewProductRepository() domain.ProductRepository {
	return datastore.NewProductRepository(i.Conn)
}

func (i *interactor) NewProductHandler() handler.ProductHandler{
	return handler.NewProductHandler(i.NewProductUsecase())
}

func (i *interactor) NewImageRepository() domain.ImageRepository{
	return datastore.NewImageRepository(i.Conn)
}


func (i *interactor) NewRegisteredShopUserUsecase() domain.RegisteredShopUserUsecase {
	return usecase.NewRegisteredShopUserUsecase(i.NewRegisteredUserRepository(i.NewShopAccountRepository()), i.NewRedisUsecase())
}

func (i *interactor) NewAuthenticateHandler() handler.AuthenticateHandler {
	shopAccountRepo := i.NewShopAccountRepository()
	userRepo := i.NewRegisteredUserRepository(shopAccountRepo)
	tk := i.NewTokenService()
	au := i.NewAuthService()

	return handler.NewAuthenticate(userRepo, tk, au)
}

func (i *interactor) NewTokenService() *auth2.Token {
	tk := auth2.NewToken(i.NewRedisUsecase())
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

func (i *interactor) NewRandomStringGeneratorUsecase() usecase.RandomStringGeneratorUsecase {
	return usecase.NewRandomStringGenrator()
}


func (i *interactor) NewSigUpUsecase() usecase.SignUpUseCase {
	return usecase.NewSignUpUsecase(i.NewRedisUsecase(), i.NewRegisteredShopUserUsecase(), i.NewRandomStringGeneratorUsecase())
}

func (i *interactor) NewAuthService() auth2.AuthInterface {
	return auth2.NewAuth(i.NewRedisUsecase())
}


func (i *interactor) NewResetPasswordHandler() handler.ResetPasswordHandler {
	return  handler.NewResetPasswordHandler(i.NewRegisteredShopUserUsecase(), i.NewRandomStringGeneratorUsecase())
}




