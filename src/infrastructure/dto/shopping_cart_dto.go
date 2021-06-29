package dto

type ShoppingCartDTO struct {
	Address  string	`json:"address"`
	Zip	uint	`json:"zip"`
	City string `json:"city"`
	State string `json:"state"`
	TotalPrice	uint	`json:"totalPrice"`
	UserId 	uint	`json:"userId"`
}
