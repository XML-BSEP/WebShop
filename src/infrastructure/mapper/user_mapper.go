package mapper

import (
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

func NewUserDtoToNewUser (userDto dto.NewUser) domain.RegisteredShopUser {

	return domain.RegisteredShopUser{Email: userDto.Email, Name: userDto.Name, Surname: userDto.Surname,
		ShopAccount: domain.ShopAccount{Username: userDto.Username, Password: userDto.Password}}
}

func NewUserDtoToRequestUser (userDto dto.NewUser) domain.UserRegistrationRequest {

	return domain.UserRegistrationRequest{Email: userDto.Email, Name: userDto.Name, Surname: userDto.Surname,
		Username: userDto.Username, Password: userDto.Password}
}


