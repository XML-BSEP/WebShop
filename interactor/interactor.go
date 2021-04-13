package interactor

import (
	"gorm.io/gorm"
	"web-shop/domain"
	"web-shop/http/handler"
	"web-shop/infrastructure/persistance/datastore"
	"web-shop/infrastructure/security/auth"
	"web-shop/usecase"
)

type Interactor interface {
	NewAddressRepository() domain.AddressRepository
	NewAddressUsecase() domain.AddressUsecase
	NewAddressHandler() handler.AddressHandler
	NewPersonRepository() domain.PersonRepository
	NewPersonUsecase() domain.PersonUsecase
	NewPersonHandler() handler.PersonHandler
	NewAuthenticateHandler() handler.AuthenticateHandler

	NewAppHandler() handler.AppHandler
}

type interactor struct {
	Conn *gorm.DB
}

func (i *interactor) NewAuthenticateHandler() handler.AuthenticateHandler {
	shopAccountRepo := datastore.NewShopAccountRepository(i.Conn)
	userRepo := datastore.NewRegisteredUserRepository(i.Conn, shopAccountRepo)

	tk := auth.NewToken()
	return handler.NewAuthenticate(userRepo, tk)
}

type appHandler struct {
	handler.AddressHandler
	handler.PersonHandler
	handler.AuthenticateHandler
}

func NewInteractor(conn *gorm.DB) Interactor {
	return &interactor{conn}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.AddressHandler = i.NewAddressHandler()
	appHandler.PersonHandler = i.NewPersonHandler()
	appHandler.AuthenticateHandler = i.NewAuthenticateHandler()
	return appHandler
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


