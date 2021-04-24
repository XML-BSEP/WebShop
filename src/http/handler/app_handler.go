package handler


type AppHandler interface {
	AuthenticateHandler
	AddressHandler
	SignUpHandler
	RedisHandlerSample
	ResetPasswordHandler
}
