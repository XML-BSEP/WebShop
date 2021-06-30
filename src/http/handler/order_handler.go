package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

type OrderHandler interface {
	PlaceOrder(ctx echo.Context) error
}


type orderHandler struct {
	OrderUseCase domain.OrderUsecase
}

func (o orderHandler) PlaceOrder(ctx echo.Context) error {
	decoder := json.NewDecoder(ctx.Request().Body)

	var t dto.OrderFrontendDto
	err := decoder.Decode(&t)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to decode")
	}
	order := domain.Order{Address: t.Address, Zip: t.Zip, City: t.City, UserId: t.UserId, State: t.State}

	err1 := o.OrderUseCase.PlaceOrder(ctx.Request().Context(), &order)

	if err1 != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Failed to place order")
	}

	return ctx.JSON(http.StatusOK, "Success")
}

func NewOrderHandler(u domain.OrderUsecase) OrderHandler {
	return &orderHandler{u}
}




