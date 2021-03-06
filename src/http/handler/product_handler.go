package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"strings"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
	"web-shop/infrastructure/mapper"
	validator2 "web-shop/validator"
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
	AddProduct(ctx echo.Context) error
	EditProduct(ctx echo.Context) error
	RemoveProduct(ctx echo.Context) error
	FetchShopProducts(ctx echo.Context) error
}

type productHandler struct {
	ProductUseCase domain.ProductUsecase
}
func (p *productHandler) RemoveProduct(ctx echo.Context) error{
	decoder := json.NewDecoder(ctx.Request().Body)
	var deletedProduct dto.DeleteProduct

	if err := decoder.Decode(&deletedProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid parameters")
	}

	policy := bluemonday.UGCPolicy();
	deletedProduct.SerialNumber =  strings.TrimSpace(policy.Sanitize(deletedProduct.SerialNumber))


	customValidator := validator2.NewCustomValidator()
	translator, _ := customValidator.RegisterEnTranslation()
	errValidation := customValidator.Validator.Struct(deletedProduct)
	errs := customValidator.TranslateError(errValidation, translator)
	errorsString := customValidator.GetErrorsString(errs)

	if errValidation != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorsString)
	}

	err := p.ProductUseCase.Delete(ctx, deletedProduct)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Error while deleting product")
	}

	return ctx.JSON(http.StatusOK, "OK")
}


func (p *productHandler) EditProduct(ctx echo.Context) error {
	decoder := json.NewDecoder(ctx.Request().Body)
	var editedProduct dto.EditProduct

	if err := decoder.Decode(&editedProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid parameters")
	}

	policy := bluemonday.UGCPolicy();
	editedProduct.Category =  strings.TrimSpace(policy.Sanitize(editedProduct.Category))
	editedProduct.Price = strings.TrimSpace(policy.Sanitize(editedProduct.Price))
	editedProduct.Available = strings.TrimSpace(policy.Sanitize(editedProduct.Available))
	editedProduct.Description = strings.TrimSpace(policy.Sanitize(editedProduct.Description))
	editedProduct.Name = strings.TrimSpace(policy.Sanitize(editedProduct.Name))
	editedProduct.Currency = strings.TrimSpace(policy.Sanitize(editedProduct.Currency))

	for i,_ := range editedProduct.Images {
		editedProduct.Images[i] = 	strings.TrimSpace(policy.Sanitize(editedProduct.Images[i]))
	}

	customValidator := validator2.NewCustomValidator()
	translator, _ := customValidator.RegisterEnTranslation()
	errValidation := customValidator.Validator.Struct(editedProduct)
	errs := customValidator.TranslateError(errValidation, translator)
	errorsString := customValidator.GetErrorsString(errs)

	if errValidation != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorsString)
	}

	product, err := p.ProductUseCase.Update(ctx, &editedProduct)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Error while editing product")
	}

	return ctx.JSON(http.StatusOK, product)
}

func (p *productHandler) FilterSearch(ctx echo.Context) error {

	decoder := json.NewDecoder(ctx.Request().Body)
	var product dto.FilterDTO

	if err := decoder.Decode(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid parameters")
	}
	//TODO: Fix filter search because joining on table categories doesnt work!
	products, err := p.ProductUseCase.FilterByCategory(product.UserId,product.Name, product.Category, product.PriceRangeStart, product.PriceRangeEnd, product.Limit, product.Offset, product.Order)
	//products, err := p.ProductUseCase.Fetch(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error filtering")
	}

	count, _ := p.ProductUseCase.Count()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error counting")
	}

	var productsRet = make([]dto.ProductViewDTO, len(products))

	for i, p := range products {
		productsRet[i] = mapper.NewProductToProductViewDTO(*p)
		productsRet[i].ProductId = p.ID
		productsRet[i].Count = count
	}
	return ctx.JSON(http.StatusOK, productsRet)
}

func (p *productHandler) AddProduct(ctx echo.Context) error {
	decoder := json.NewDecoder(ctx.Request().Body)

	var t dto.NewProduct
	err := decoder.Decode(&t)

	policy := bluemonday.UGCPolicy();
	t.Category =  strings.TrimSpace(policy.Sanitize(t.Category))
	t.Price = strings.TrimSpace(policy.Sanitize(t.Price))
	t.Available = strings.TrimSpace(policy.Sanitize(t.Available))
	t.Description = strings.TrimSpace(policy.Sanitize(t.Description))
	t.Name = strings.TrimSpace(policy.Sanitize(t.Name))
	t.Currency = strings.TrimSpace(policy.Sanitize(t.Currency))

	for i,_ := range t.Images {
		t.Images[i] = 	strings.TrimSpace(policy.Sanitize(t.Images[i]))
	}

	customValidator := validator2.NewCustomValidator()
	translator, _ := customValidator.RegisterEnTranslation()
	errValidation := customValidator.Validator.Struct(t)
	errs := customValidator.TranslateError(errValidation, translator)
	errorsString := customValidator.GetErrorsString(errs)

	if errValidation != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorsString)
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}


	product, err := p.ProductUseCase.Create(ctx, &t)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Error while creating product")
	}

	return ctx.JSON(http.StatusOK, product)

}

func (p *productHandler) FetchShopProducts(ctx echo.Context) error {
	decoder := json.NewDecoder(ctx.Request().Body)
	var t dto.ShopIdDTO
	decodeErr := decoder.Decode(&t)

	if decodeErr !=nil{
		return echo.NewHTTPError(http.StatusBadRequest, "Error decoding dto")
	}

	products, err := p.ProductUseCase.GetAllAvailableProductsInUsersShop(ctx,t.UserId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No results")
	}

	var productsRet = make([]dto.ProductViewDTO, len(products))

	for i, p := range products {
		productsRet[i] = mapper.NewProductToProductViewDTO(*p)
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
