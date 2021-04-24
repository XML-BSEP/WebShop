package dto

type Product struct {
	Name string `json:"name"`
	Price uint64 `json:"price"`
	Image string 	`json:"image"`
	Currency string `json:"currency"`
	Available uint `json:"available"`
	Description string `json:"description"`
	Category string `json:"category"'`
}
