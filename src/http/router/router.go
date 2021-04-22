package router

import (
	"github.com/labstack/echo"
	"web-shop/http/handler"
	"web-shop/http/middleware"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {


	e.POST("/login", h.Login)
	g := e.Group("/member")
	g.Use(middleware.Auth2())
	g.GET("/addresses", h.GetAddresses)

	/*e.GET("/addresses", h.GetAddresses)
	e.POST("/login", h.Login)
	e.POST("/logout", h.Logout)
	e.POST("/register", h.UserRegister)
	e.POST("/redisPost", h.AddKeyValueSet)
	e.POST("/redisVal", h.GetValueByKey)
	e.POST("/confirmAccount", h.ConfirmAccount)*/



}
