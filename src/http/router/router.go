package router

import (
	"github.com/labstack/echo"
	"web-shop/http/handler"
	"web-shop/http/middleware"
)

func NewRouter(e *echo.Echo, h handler.AppHandler, authMiddleware middleware.AuthMiddleware) {

	e.GET("/categories", h.GetAllCategories)
	e.POST("/login", h.Login, authMiddleware.Authenticated())
	e.POST("/register", h.UserRegister, authMiddleware.Authenticated())
	e.POST("/confirmAccount", h.ConfirmAccount, authMiddleware.Authenticated())
	e.GET("/products", h.FetchProducts)
	e.POST("/confirmAccount", h.ConfirmAccount, authMiddleware.Authenticated())
	e.POST("/resetPasswordMail", h.SendResetMail, authMiddleware.Authenticated())
	e.POST("/resetPassword", h.ResetPassword, authMiddleware.Authenticated())

	g := e.Group("/")
	g.Use(authMiddleware.Auth())
	g.GET("addresses", h.GetAddresses)
	g.POST("logout", h.Logout)
	g.POST("addProduct", h.AddProduct)

}
