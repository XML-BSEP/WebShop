package mapper

import (
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

func NewUserDtoToNewUser (userDto dto.NewUser) domain.RegisteredShopUser {

	return domain.RegisteredShopUser{Email: userDto.Email, Name: userDto.Name, Surname: userDto.Surname,
		SecurityAnswer: userDto.SecurityAnswer, SecurityQuestion: userDto.SecurityQuestion,
		ShopAccount: domain.ShopAccount{Username: userDto.Username, Password: userDto.Password}}
}



