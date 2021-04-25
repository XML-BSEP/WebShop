package dto

type FilterDTO struct {
	Name string `json:"name"`
	Category string `json:"category"`
	PriceRangeStart uint `json:"priceRangeStart"`
	PriceRangeEnd uint `json:"priceRangeEnd"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
	Order string `json:"order"`

}
