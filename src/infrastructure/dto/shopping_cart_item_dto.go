package dto
type ShoppingCartItemDTO struct {
	Product        ProductDTO `json:"product"`
	Amount         uint	`json:"quantity"`
}
