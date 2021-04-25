package dto


type SearchName struct {
	SearchName	string	`json:"search_name"`
	Offset   int		`json:"offset"`
	Limit    int		`json:"limit"`
	Order 	 int		`json:"order"`
}
