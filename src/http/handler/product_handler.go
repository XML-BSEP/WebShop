package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
	"web-shop/infrastructure/mapper"
)

type ProductHandler interface {
	GetProductsWithPriceRange(c echo.Context) error
	GetProductsWithCategory(ctx echo.Context) error
	GetProductsWithCondition(ctx echo.Context) error
	GetByName(ctx echo.Context) error
	GetProductsWithConditionOrderedByPrice(ctx echo.Context) error
	GetProductsWithConditionOrderedByName(ctx echo.Context) error
	GetByNameOrderByPrice(ctx echo.Context) error
	GetByNameOrderByName(ctx echo.Context) error
	FilterByCategory(ctx echo.Context) error
	FetchProducts(ctx echo.Context) error
	FilterSearch(ctx echo.Context) error
}

type productHandler struct {
	ProductUseCase domain.ProductUsecase
}

func (p *productHandler) FilterSearch(ctx echo.Context) error {

	decoder := json.NewDecoder(ctx.Request().Body)
	var product dto.FilterDTO

	if err := decoder.Decode(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid parameters")
	}

	products, err := p.ProductUseCase.FilterByCategory(product.Name, product.Category, product.PriceRangeStart, product.PriceRangeEnd, product.Limit, product.Offset, product.Order)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Errpr filtering")
	}

	count, _ := p.ProductUseCase.Count()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error counting")
	}

	var productsRet = make([]dto.ProductViewDTO, len(products))

	for i, p := range products {
		productsRet[i] = mapper.NewProductToProductViewDTO(*p)
		productsRet[i].Count = count
	}
	return ctx.JSON(http.StatusOK, productsRet)
}

func (p *productHandler) FetchProducts(ctx echo.Context) error {

	products, err := p.ProductUseCase.Fetch(ctx)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No results")
	}



	var productsRet = make([]dto.ProductViewDTO, len(products))

	for i, p := range products {
		productsRet[i] = mapper.NewProductToProductViewDTO(*p)
	}

	return ctx.JSON(http.StatusOK, productsRet)
}

func (p productHandler) FilterByCategory(ctx echo.Context) error {
	panic("implement me")
}

func (p productHandler) GetProductsWithConditionOrderedByPrice(ctx echo.Context) error{
	decoder := json.NewDecoder(ctx.Request().Body)

	var t dto.PriceRangeCategory
	err := decoder.Decode(&t)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	products, err := p.ProductUseCase.GetProductsWithConditionOrderedByPrice(t.Low,t.High,t.Category,t.Limit,t.Offset, t.Order)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Products do not exist")
	}

	return ctx.JSON(http.StatusOK, products)
}

func (p productHandler) GetProductsWithConditionOrderedByName(ctx echo.Context) error{
	decoder := json.NewDecoder(ctx.Request().Body)

	var t dto.PriceRangeCategory
	err := decoder.Decode(&t)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	products, err := p.ProductUseCase.GetProductsWithConditionOrderedByName(t.Low,t.High,t.Category,t.Limit,t.Offset, t.Order)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Products do not exist")
	}

	return ctx.JSON(http.StatusOK, products)
}

func (p productHandler) GetByNameOrderByPrice(ctx echo.Context) error{
	decoder := json.NewDecoder(ctx.Request().Body)
	var t dto.SearchName
	err := decoder.Decode(&t)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	name := "%"+t.SearchName+"%"
	products, err := p.ProductUseCase.GetByNameOrderByPrice(name,t.Limit,t.Offset, t.Order)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Products do not exist")
	}

	return ctx.JSON(http.StatusOK, products)

}

func (p productHandler) GetByNameOrderByName(ctx echo.Context) error{
	decoder := json.NewDecoder(ctx.Request().Body)
	var t dto.SearchName
	err := decoder.Decode(&t)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	name := "%"+t.SearchName+"%"
	products, err := p.ProductUseCase.GetByNameOrderByName(name,t.Limit,t.Offset, t.Order)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Products do not exist")
	}

	return ctx.JSON(http.StatusOK, products)}

func NewProductHandler(u domain.ProductUsecase) ProductHandler {
	return &productHandler{u}
}

func (p productHandler) GetByName(ctx echo.Context) error {
	decoder := json.NewDecoder(ctx.Request().Body)
	var t dto.SearchName
	err := decoder.Decode(&t)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	name := "%"+t.SearchName+"%"
	products, err := p.ProductUseCase.GetByName(name,t.Limit,t.Offset)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Products do not exist")
	}

	return ctx.JSON(http.StatusOK, products)
}

func (p productHandler) GetProductsWithCondition(ctx echo.Context) error {
	decoder := json.NewDecoder(ctx.Request().Body)

	var t dto.PriceRangeCategory
	err := decoder.Decode(&t)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	products, err := p.ProductUseCase.GetProductsWithCondition(t.Low,t.High,t.Category,t.Limit,t.Offset)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Products do not exist")
	}

	return ctx.JSON(http.StatusOK, products)
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
		return echo.NewHTTPError(http.StatusNotFound, "Products do not exist")
	}

	return ctx.JSON(http.StatusOK, products)
}

func (p productHandler) GetProductsWithCategory(ctx echo.Context) error{
	decoder := json.NewDecoder(ctx.Request().Body)

	var t dto.ProductCategory
	err := decoder.Decode(&t)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	products, err := p.ProductUseCase.GetProductsWithCategory(t.Category)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Products do not exist")
	}

	return ctx.JSON(http.StatusOK, products)
}
