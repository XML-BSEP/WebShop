package dto

type ShoppingCartItemFrontDTO struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Price float64	`json:"price"`
	Picture string `json:"picture"`
	Description string	`json:"description"`
}
