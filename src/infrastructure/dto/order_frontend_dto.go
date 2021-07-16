package dto

type OrderFrontendDto struct {
	UserId	uint	`json:"userId"`
	Address	string `json:"address"`
	City	string `json:"city"`
	Zip	uint `json:"zip"`
	State	string `json:"state"`
}
