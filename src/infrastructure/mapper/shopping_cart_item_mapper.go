package mapper

import (
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

func NewShoppingCartItemDtoToNewShoppingCartItem (dto dto.ShoppingCartItemDTO) domain.ShoppingCartItem {

	return domain.ShoppingCartItem{Product: NewProductDtoToProduct(dto.Product), Amount: dto.Amount}
}

func NewShoppingCartItemToNewShoppingCartItemDTO (s domain.ShoppingCartItem) dto.ShoppingCartItemDTO {

	return dto.ShoppingCartItemDTO{Product: NewProductToProductDto(s.Product), Amount: s.Amount}
}

func NewDtosFromShoppingCartItems(s []domain.ShoppingCartItem) []dto.ShoppingCartItemDTO{
	var dtos []dto.ShoppingCartItemDTO
	for _, sci := range s {
		dtos = append(dtos, NewShoppingCartItemToNewShoppingCartItemDTO(sci))
	}

	return dtos
}

func NewItemsFromDtos(s []dto.ShoppingCartItemDTO) []domain.ShoppingCartItem{
	var items []domain.ShoppingCartItem
	for _, sci := range s {
		items = append(items, NewShoppingCartItemDtoToNewShoppingCartItem(sci))
	}

	return items
}