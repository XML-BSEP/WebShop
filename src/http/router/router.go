package router

import (
	"github.com/labstack/echo"
	"web-shop/http/handler"
	"web-shop/http/middleware"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {

	e.POST("/login", h.Login, middleware.Authenticated())
	e.POST("/register", h.UserRegister, middleware.Authenticated())
	e.POST("/confirmAccount", h.ConfirmAccount, middleware.Authenticated())
	e.POST("/resetPasswordMail", h.SendResetMail, middleware.Authenticated())
	g := e.Group("/member")
	g.Use(middleware.Auth())
	g.GET("/addresses", h.GetAddresses)
	g.POST("/logout", h.Logout)

}
