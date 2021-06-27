package dto

type DeleteProduct struct{
	UserId 		 uint 	  `json:"userId"`
	SerialNumber string   `json:"serial" validate:"required"`
}
