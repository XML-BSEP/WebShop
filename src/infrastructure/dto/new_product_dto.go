package dto
type NewProduct struct{
	Name  string `json:"name"`
	Price uint `json:"price"`
	Category string `json:"category"`
	Description string `json:"description"`
	Images []string `json:"images"`
	Currency uint `json:"currency"`
}