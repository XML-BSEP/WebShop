package handler

import (
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


	addresses, err := a.AddressUsecase.Fetch(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Addreses does not exists")
	}

	return c.JSON(http.StatusOK, addresses)
}


