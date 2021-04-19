package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
	"web-shop/infrastructure/mapper"
	"web-shop/usecase"
	validator2 "web-shop/validator"
)

type SignUpHandler interface {
	UserRegister(ctx echo.Context) error
	ConfirmAccount(ctx echo.Context) error
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

	user, errE :=  signUp.SignUpUsecase.CheckIfExistUser(ctx, t)
	if errE == nil {
		return ctx.JSON(http.StatusBadRequest, "User already exist!")
	}

	fmt.Print(user)
	newUser := mapper.NewUserDtoToRequestUser(t)

	customValidator := validator2.NewCustomValidator()
	translator, _ := customValidator.RegisterEnTranslation()
	errValidation := customValidator.Validator.Struct(newUser)
	errs := customValidator.TranslateError(errValidation, translator)
	errorsString := customValidator.GetErrorsString(errs)

	if errValidation != nil {
		return ctx.JSON(http.StatusBadRequest, errorsString[0])
	}

	passwordCompare := signUp.SignUpUsecase.ValidatePassword(t.Password, t.ConfirmedPassword)

	if !passwordCompare {
		return ctx.JSON(http.StatusBadRequest, "Enter same passwords")
	}
	code, errR := signUp.SignUpUsecase.RegisterNewUser(ctx, newUser)
	if errR != nil {
		return ctx.JSON(http.StatusBadRequest, "Redis failed!")
	}

	go usecase.SendMail(t.Email, t.Username, code)

	return ctx.JSON(http.StatusOK, "Successfull registration, please check your mail!")
}


func (signUp *signUp) ConfirmAccount(ctx echo.Context) error {

	decoder := json.NewDecoder(ctx.Request().Body)

	var credentials dto.ConfirmRegistrationDTO
	_ = decoder.Decode(&credentials)

	customValidator := validator2.NewCustomValidator()
	translator, _ := customValidator.RegisterEnTranslation()
	validateErr := customValidator.Validator.Struct(credentials)
	errs := customValidator.TranslateError(validateErr, translator)
	errorsString := customValidator.GetErrorsString(errs)



	if validateErr != nil {
		return ctx.JSON(http.StatusBadRequest, errorsString[0])
	}

	code := credentials.VerificationCode
	email := credentials.Email

	user, err := signUp.SignUpUsecase.IsCodeValid(ctx, email, code)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Entered code is not valid")
	}

	_, err2 := signUp.SignUpUsecase.ConfirmAccount(ctx, user)

	if err2 != nil {
		return ctx.JSON(http.StatusBadRequest, "Znaci...")
	}

	return ctx.JSON(http.StatusOK, "Ok")
}

