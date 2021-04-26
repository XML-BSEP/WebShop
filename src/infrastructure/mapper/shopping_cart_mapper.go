package mapper
import (
	"github.com/microcosm-cc/bluemonday"
	"strings"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

func NewShoppingCartDtoToNewShoppingCart (dto dto.ShoppingCartDTO) domain.ShoppingCart {

	policy := bluemonday.UGCPolicy();
	dto.Buyer =  strings.TrimSpace(policy.Sanitize(dto.Buyer))

	return domain.ShoppingCart{ShoppingCartItems: NewItemsFromDtos(dto.ShoppingCartItems), Buyer: dto.Buyer}
}

func NewShoppingCartToNewShoppingCartDTO (s domain.ShoppingCart) dto.ShoppingCartDTO {
	return dto.ShoppingCartDTO{ShoppingCartItems: NewDtosFromShoppingCartItems(s.ShoppingCartItems), Buyer: s.Buyer}
}

