package dto

type FilterProductDTO struct {
	Category string `json:"category"`
	PriceStart uint `json:"priceRangeStart"`
	PriceEnd uint `json:"priceRangeEnd"`

}
