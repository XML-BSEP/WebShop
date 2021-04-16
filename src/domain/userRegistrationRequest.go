package domain

//TODO stavi u malo
type UserRegistrationRequest struct {
	Email	string	`json:"email"`
	SecurityQuestion	string	`json:"question"`
	SecurityAnswer	string	`json:"answer"`
	Name	string	`json:"name"`
	Surname	string	`json:"surname"`
	Username     string	`json:"username"`
	Password   string	`json:"password"`
	VerificationCode	string `json:"code"`

}
