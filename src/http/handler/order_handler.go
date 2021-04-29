package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
	"web-shop/infrastructure/mapper"
)

type OrderHandler interface {
	PlaceOrder(ctx echo.Context) error
}


type orderHandler struct {
	OrderUseCase domain.OrderUsecase
}

func (o orderHandler) PlaceOrder(ctx echo.Context) error {
	decoder := json.NewDecoder(ctx.Request().Body)

	var t dto.OrderDTO
	err := decoder.Decode(&t)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	order:=mapper.NewOrderDtoToNewOrder(t)

	products, err := o.OrderUseCase.Create(ctx, &order)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Products do not exist")
	}

	return ctx.JSON(http.StatusOK, products)
}

func NewOrderHandler(u domain.OrderUsecase) OrderHandler {
	return &orderHandler{u}
}




