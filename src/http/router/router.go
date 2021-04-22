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
	e.POST("/getProductsWithCategory", h.GetProductsWithCategory)
	e.POST("/getProductsWithCategoryInPriceRange", h.GetProductsWithCondition)
	e.POST("/getByName", h.GetByName)
	e.POST("/getProductsWithCategoryInPriceRangeOrderByPrice", h.GetProductsWithConditionOrderedByPrice)
	e.POST("/getProductsWithCategoryInPriceRangeOrderByName", h.GetProductsWithConditionOrderedByName)
	e.POST("/getByNameOrderedByName", h.GetByNameOrderByName)
	e.POST("/getByNameOrderedByPrice", h.GetByNameOrderByPrice)
	e.POST("/placeOrder", h.PlaceOrder)

}
