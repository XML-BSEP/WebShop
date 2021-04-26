package dto

type PriceRangeCategory struct{
	Category string 	`json:"category"`
	Low		 uint 		`json:"low"`
	High	 uint 		`json:"high"`
	Offset   int		`json:"offset"`
	Limit    int		`json:"limit"`
	Order 	 int		`json:"order"`
}
