package dto

type ProductViewDTO struct {
	Name        string `json:"name"`
	Price       float64 `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Category    string `json:"category"`
	Available   uint   `json:"available"`
	Currency    string `json:"currency"`
}
