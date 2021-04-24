package mapper
import (
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

func NewShoppingCartDtoToNewShoppingCart (dto dto.ShoppingCartDTO) domain.ShoppingCart {
	return domain.ShoppingCart{ShoppingCartItems: NewItemsFromDtos(dto.ShoppingCartItems), Buyer: dto.Buyer}
}

func NewShoppingCartToNewShoppingCartDTO (s domain.ShoppingCart) dto.ShoppingCartDTO {
	return dto.ShoppingCartDTO{ShoppingCartItems: NewDtosFromShoppingCartItems(s.ShoppingCartItems), Buyer: s.Buyer}
}

