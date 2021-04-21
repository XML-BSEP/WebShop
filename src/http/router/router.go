package router

import (
	"github.com/labstack/echo"
	"web-shop/http/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.GET("/addresses", h.GetAddresses)
	e.POST("/login", h.Login)
	e.POST("/logout", h.Logout)
	e.POST("/register", h.UserRegister)
	e.POST("/redisPost", h.AddKeyValueSet)
	e.POST("/redisVal", h.GetValueByKey)
	e.POST("/confirmAccount", h.ConfirmAccount)
	e.POST("/getProductsInPriceRange", h.GetProductsWithPriceRange)

}
