package dto

type NewUser struct {
	Email	string	`json:"email" validate:"required,email"`
	SecurityQuestion	string	`json:"question"`
	SecurityAnswer	string	`json:"answer"`
	Name	string	`json:"name" validate:"required,name"`
	Surname	string	`json:"surname" validate:"required"`
	Username     string	`json:"username" validate:"required"`
	Password   string	`json:"password" validate:"required"`
	ConfirmedPassword	string `json:"confirmedpassword" validate:"required"`
}