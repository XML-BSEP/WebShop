package mapper

import (
	"github.com/microcosm-cc/bluemonday"
	"strings"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

func NewUserDtoToNewUser (userDto dto.NewUser) domain.RegisteredShopUser {

	return domain.RegisteredShopUser{Email: userDto.Email, Name: userDto.Name, Surname: userDto.Surname,
		ShopAccount: domain.ShopAccount{Username: userDto.Username, Password: userDto.Password}}
}

func NewUserDtoToRequestUser (userDto dto.NewUser) domain.UserRegistrationRequest {

	policy := bluemonday.UGCPolicy();

	email := strings.TrimSpace(policy.Sanitize(userDto.Email))
	name := strings.TrimSpace(policy.Sanitize(userDto.Name))
	surname := strings.TrimSpace(policy.Sanitize(userDto.Surname))
	username := strings.TrimSpace(policy.Sanitize(userDto.Username))
	password := strings.TrimSpace(policy.Sanitize(userDto.Password))


	return domain.UserRegistrationRequest{Email: email, Name: name, Surname: surname,
		Username: username, Password: password}
}



