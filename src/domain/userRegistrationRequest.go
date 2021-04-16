package domain

//TODO stavi u malo
type UserRegistrationRequest struct {
	Email	string	`json:"email" validate:"required,email"`
	SecurityQuestion	string	`json:"question" validate:"required"`
	SecurityAnswer	string	`json:"answer" validate:"required"`
	Name	string	`json:"name" validate:"required"`
	Surname	string	`json:"surname" validate:"required"`
	Username     string	`json:"username" validate:"required,min=4"`
	Password   string	`json:"password" validate:"required"`
	VerificationCode	string `json:"code" validate:"required"`

}
