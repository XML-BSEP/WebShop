package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

type ShopAccountHandler interface {
	GetAllShopAdminAccounts(c echo.Context) error
}

type shopAccountHandler struct {
	ShopAccountUsecase domain.ShopAccountUsecase
}

func (acc *shopAccountHandler) GetAllShopAdminAccounts(c echo.Context) error {
	shops, err := acc.ShopAccountUsecase.FetchShops(c)
	var dtos []dto.ShopViewDTO
	for _, it := range shops {
		dtos = append(dtos, dto.ShopViewDTO{ID: it.ID, ShopName: it.Username+"'s web shop"})
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "There are no web shops")
	}

	return c.JSON(http.StatusOK, dtos)
}

func NewShopAccountHandler(u domain.ShopAccountUsecase) ShopAccountHandler {
	return &shopAccountHandler{u}
}
