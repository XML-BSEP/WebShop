package mapper

import (
	"gorm.io/gorm"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

func NewProductDtoToProduct (dto dto.ProductDTO) domain.Product {

	return domain.Product{Model: gorm.Model{ID: dto.ID}, Name:dto.Name, Price : dto.Price}
}

func NewProductToProductDto (p domain.Product) dto.ProductDTO {

	return dto.ProductDTO{ID:p.ID, Name : p.Name, Price: p.Price}
}

