package mapper

import (
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
	"strings"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

func NewProductDtoToProduct (dto dto.ProductDTO) domain.Product {

	policy := bluemonday.UGCPolicy();
	dto.Name =  strings.TrimSpace(policy.Sanitize(dto.Name))

	return domain.Product{Model: gorm.Model{ID: dto.ID}, Name:dto.Name, Price : dto.Price}
}

func NewProductToProductDto (p domain.Product) dto.ProductDTO {

	return dto.ProductDTO{ID:p.ID, Name : p.Name, Price: p.Price}
}

func NewProductToProductViewDTO (p domain.Product) dto.ProductViewDTO {

	imt_path := "https://localhost:443/static/"
	return dto.ProductViewDTO{
		Category: p.Category.Name,
		Name: p.Name,
		Available: p.Available,
		Currency: mapCurrency(p.Currency),
		Description: p.Description,
		Image: imt_path + p.Images[0].Path,
		Price: p.Price,
	}
}

func mapCurrency(currency domain.Currency) string {
	if currency == 0 {
		return "USD"
	}
	if currency == 1 {
		return "EUR"
	}
	if currency == 2 {
		return "RSD"
	}

	return ""
}

