package handler

import (
	"context"
	"github.com/labstack/echo"
	"net/http"
	"web-shop/domain"
)

type AddressHandler interface {
	GetAddresses(c echo.Context) error
}

type addressHandler struct {
	AddressUsecase domain.AddressUsecase
}

func NewAddressHandler(u domain.AddressUsecase) AddressHandler {
	return &addressHandler{u}
}

func (a *addressHandler) GetAddresses(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	addresses, err := a.AddressUsecase.Fetch(ctx)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Addreses does not exists")
	}

	return c.JSON(http.StatusOK, addresses)
}


