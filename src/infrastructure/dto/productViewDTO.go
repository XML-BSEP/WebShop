package dto

type ProductViewDTO struct {
	UserId 		 uint 	  `json:"userId"`
	ProductId   uint 	`json:"productId"`
	Name        string `json:"name"`
	Price       float64 `json:"price"`
	Description string `json:"description"`
	Image       []string `json:"image"`
	Category    string `json:"category"`
	Available   uint   `json:"available"`
	Currency    string `json:"currency"`
	Count int64 `json:"count"`
	SerialNumber uint64 `json:"serial"`
}
