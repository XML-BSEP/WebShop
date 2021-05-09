package dto

type DeleteProduct struct{
	SerialNumber string   `json:"serial" validate:"required"`
}
