package handler

import (
	"context"
	"github.com/labstack/echo"
	"net/http"
	"web-shop/domain"
)

type PersonHandler interface {
	GetPersons(c echo.Context) error
}

type personHandler struct {
	PersonUsecase domain.PersonUsecase
}

func NewPersonHandler(p domain.PersonUsecase) PersonHandler {
	return &personHandler{p}
}

func (p *personHandler) GetPersons(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	persons, err := p.PersonUsecase.Fetch(ctx)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Persons not found")
	}

	return c.JSON(http.StatusOK, persons)
}



