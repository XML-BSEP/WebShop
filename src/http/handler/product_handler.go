package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

type ProductHandler interface {
	GetProductsWithPriceRange(c echo.Context) error
}

type productHandler struct {
	ProductUseCase domain.ProductUsecase
}

func NewProductHandler(u domain.ProductUsecase) ProductHandler {
	return &productHandler{u}
}

func (p productHandler) GetProductsWithPriceRange(ctx echo.Context) error {

	decoder := json.NewDecoder(ctx.Request().Body)

	var t dto.PriceRange
	err := decoder.Decode(&t)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	products, err := p.ProductUseCase.GetWithPriceRange(t.Low, t.High)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Addreses does not exists")
	}

	return ctx.JSON(http.StatusOK, products)
}

