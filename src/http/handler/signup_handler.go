package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
	"web-shop/infrastructure/mapper"
	"web-shop/usecase"
)

type SignUpHandler interface {
	UserRegister(ctx echo.Context) error
}


type signUp struct {
	us domain.RegisteredShopUserRepository
	SignUpUsecase usecase.SignUpUseCase
}

func NewSignUpHandler(us domain.RegisteredShopUserRepository,signUpUsecase usecase.SignUpUseCase) SignUpHandler {
	return &signUp{us, signUpUsecase}
}


func (signUp *signUp) UserRegister(ctx echo.Context) (err error){

	decoder := json.NewDecoder(ctx.Request().Body)

	var t dto.NewUser
	err = decoder.Decode(&t)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, errE :=  signUp.SignUpUsecase.CheckIfExistUser(ctx, t)
	if errE != nil {
		ctx.JSON(http.StatusBadRequest, "User does not exists!")
	}


	code, errR := signUp.SignUpUsecase.RegisterNewUser(ctx, mapper.NewUserDtoToRequestUser(t))
	if errR != nil {
		ctx.JSON(http.StatusBadRequest, "Redis failed!")
	}

	go usecase.SendMail(t.Email, t.Username, code)

	return ctx.JSON(http.StatusOK, "Successfull registration!")



}
