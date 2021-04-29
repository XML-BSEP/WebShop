package mapper

import (
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)


func NewCategoryToNewCategoryDTO (c *domain.Category) dto.ProductCategory {
	return dto.ProductCategory{Category: c.Name}
}
