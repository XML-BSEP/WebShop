package dto
type NewProduct struct{
	UserId 		 uint 	  `json:"userId"`
	Name  string `json:"name" validate:"required"`
	Price string `json:"price" validate:"required,price"`
	Category string `json:"category" validate:"required"`
	Description string `json:"description" validate:"required"`
	Images []string `json:"images" validate:"required,images"`
	Currency string `json:"currency"`
	Available string `json:"available" validate:"required,available"`

}