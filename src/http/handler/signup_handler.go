package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
	"web-shop/usecase"
)

type SignUpHandler interface {
	UserRegister(ctx echo.Context) error
}


type SignUp struct {
	us domain.RegisteredShopUserRepository
}

func NewSignUpHandler() SignUpHandler {
	return &SignUp{nil}
}


func (SignUp *SignUp) UserRegister(ctx echo.Context) (err error){

	/*
	//var newUser *dto.NewUser
	var json map[string]interface{} = map[string]interface{}{}

	if err := ctx.Bind(&json); err != nil {
		return err
	}*/

	/*
	u := new(dto.NewUser)

	if err = ctx.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

	return ctx.String(http.StatusOK, u.Name)
*/

	decoder := json.NewDecoder(ctx.Request().Body)

	var t dto.NewUser
	err = decoder.Decode(&t)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	usecase.SendMail(t.Email, t.Name)



	return ctx.JSON(http.StatusOK, t)


	//AJAO IDE GAS AAAAAAAAAAAAAAA
	// SJUJUJUUUUU
	//return  ctx.JSON(http.StatusCreated, "")
}
