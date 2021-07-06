package router

import (
	"github.com/labstack/echo"
	"web-shop/http/handler"
	"web-shop/http/middleware"
)

func NewRouter(e *echo.Echo, h handler.AppHandler, authMiddleware middleware.AuthMiddleware) {

	e.GET("/categories", h.GetAllCategories)
	e.GET("/shopProducts", h.FetchShopProducts)
	e.GET("/allShops", h.GetAllShopAdminAccounts)
	e.GET("/products", h.FetchProducts)

	e.POST("/login", h.Login, authMiddleware.Authenticated())
	e.POST("/register", h.UserRegister, authMiddleware.Authenticated())
	e.POST("/confirmAccount", h.ConfirmAccount, authMiddleware.Authenticated())
	e.POST("/confirmAccount", h.ConfirmAccount, authMiddleware.Authenticated())
	e.POST("/resetPasswordMail", h.SendResetMail, authMiddleware.Authenticated())
	e.POST("/resetPassword", h.ResetPassword, authMiddleware.Authenticated())
	e.POST("/filterSearch", h.FilterSearch)
	e.POST("/resendRegistrationCode", h.ResendCode, authMiddleware.Authenticated())
	e.POST("/resendPassCode", h.ResendResetCode, authMiddleware.Authenticated())
	e.POST("/refresh", h.Refresh, authMiddleware.Authenticated())
	e.POST("/addToCart", h.AddToCart)
	e.POST("/getCart", h.GetUsersShoppingCartItems)
	e.POST("/removeFromCart", h.RemoveFromCart)
	e.POST("/placeOrder", h.PlaceOrder)

	g := e.Group("/")
	g.Use(authMiddleware.Auth())
	g.GET("addresses", h.GetAddresses)
	g.POST("logout", h.Logout)
	g.POST("addProduct", h.AddProduct)
	g.POST("editProduct", h.EditProduct)
	g.POST("deleteProduct", h.RemoveProduct)
	g.POST("createAd", h.CreateAd)
	g.POST("saveToken", h.SaveToken)
	g.GET("getAllAdsPerAgent", h.GetAllAdsPerAgent)
	g.POST("createDisposableCampaign", h.CreateDisposableCampaign)
	g.POST("createMultipleCampaign", h.CreateMultipleCampaign)
	g.GET("getAllMultipleCampaigns", h.GetAllMultipleCampaigns)
	g.GET("getAllDisposableCampaigns", h.GetAllDisposableCampaigns)
	g.POST("deleteDisposableCampaign", h.DeleteDisposableCampaign)
	g.POST("deleteMultipleCampaign", h.DeleteMultipleCampaign)
	g.POST("updateMultipleCampaign", h.UpdateMultipleCampaign)
}
