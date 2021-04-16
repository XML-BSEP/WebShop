package dto

type NewUser struct {
	Email	string	`json: "email"`
	SecurityQuestion	string	`json: "question"`
	SecurityAnswer	string	`json: "answer"`
	Name	string	`json: "name"`
	Surname	string	`json: "surname"`
	Username     string	`json: "username"`
	Password   string	`json: "password"`
	ConfirmedPassword	string `json: "confirmedpassword"`
}