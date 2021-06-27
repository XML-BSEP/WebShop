package dto

type FilterProductDTO struct {
	UserId 		 uint 	  `json:"userId"`

	Category string `json:"category"`
	PriceStart uint `json:"priceRangeStart"`
	PriceEnd uint `json:"priceRangeEnd"`

}
