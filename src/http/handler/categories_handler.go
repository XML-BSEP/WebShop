package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
	"web-shop/infrastructure/mapper"
)

type CategoryHandler interface {
	GetAllCategories(c echo.Context) error
}

type categoryHandler struct {
	CategoryUsecase domain.CategoryUsecase
}

func (c2 *categoryHandler) GetAllCategories(c echo.Context) error {
	categories, err := c2.CategoryUsecase.Fetch(c)
	var dtos []dto.ProductCategory
	for i, _ := range categories {
		dtos = append(dtos, mapper.NewCategoryToNewCategoryDTO(categories[i]))
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Categories does not exists")
	}

	return c.JSON(http.StatusOK, dtos)
}

func NewCategoryHandler(u domain.CategoryUsecase) CategoryHandler {
	return &categoryHandler{u}
}
