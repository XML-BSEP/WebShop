package dto
type NewProduct struct{
	Name  string `json:"name"`
	Price string `json:"price"`
	Category string `json:"category"`
	Description string `json:"description"`
	Images []string `json:"images"`
	Currency string `json:"currency"`
	Available string `json:"available"`

}