package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

type ShoppingCartItemHandler interface {
	AddToCart(ctx echo.Context) error
	GetUsersShoppingCartItems(ctx echo.Context) error
}

type shoppingCartItemHandler struct {
	ShoppingCartItemUsecase domain.ShoppingCartItemUsecase

}

func (s shoppingCartItemHandler) GetUsersShoppingCartItems(ctx echo.Context) error {
	decoder := json.NewDecoder(ctx.Request().Body)
	var user dto.UserIdDTO

	if err := decoder.Decode(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid parameters")
	}
	items, err := s.ShoppingCartItemUsecase.GetAllUsersShoppingCartItems(ctx.Request().Context(), user.UserId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Error getting items")
	}
	var itemsInCart []dto.ShoppingCartItemFrontDTO
	for _,it := range items{
		image := "https://localhost:443/static/"
		itemsInCart=append(itemsInCart, dto.ShoppingCartItemFrontDTO{Id: it.ProductID, Description: it.Product.Description, Price : it.Product.Price, Name:it.Product.Name, Picture: image+it.Product.Images[0].Path})
	}
	return ctx.JSON(http.StatusOK, itemsInCart)
}

func (s shoppingCartItemHandler) AddToCart(ctx echo.Context) error {
	decoder := json.NewDecoder(ctx.Request().Body)
	var itemToCart dto.ItemToCart

	if err := decoder.Decode(&itemToCart); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid parameters")
	}

	err := s.ShoppingCartItemUsecase.AddToCart(ctx.Request().Context(), itemToCart.ProductId, itemToCart.UserId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Error while adding product to cart")
	}

	return ctx.JSON(http.StatusOK, "OK")
}

func NewShoppingCartItemHandler(us domain.ShoppingCartItemUsecase) ShoppingCartItemHandler {
	return &shoppingCartItemHandler{us}
}
