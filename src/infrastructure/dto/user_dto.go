package dto

type NewUser struct {
	Email	string	`json: "email"`
	Name	string	`json: "name"`
	Surname	string	`json: "surname"`
	Username     string	`json: "username"`
	Password   string	`json: "password"`
	ConfirmedPassword	string `json: "confirmedpassword"`
}