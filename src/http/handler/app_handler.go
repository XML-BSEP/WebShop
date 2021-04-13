package handler


type AppHandler interface {
	AuthenticateHandler
	AddressHandler
	PersonHandler
}
