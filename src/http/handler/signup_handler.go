package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"strings"
	"unicode"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
	"web-shop/infrastructure/mapper"
	"web-shop/usecase"
	validator2 "web-shop/validator"
)

type SignUpHandler interface {
	UserRegister(ctx echo.Context) error
	ConfirmAccount(ctx echo.Context) error
	ResendCode(ctx echo.Context) error
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
	if errE != nil {

		return echo.NewHTTPError(http.StatusBadRequest, "User already exist!")
	}

	fmt.Print(user)
	newUser := mapper.NewUserDtoToRequestUser(t)

	if strings.Contains(newUser.Username, " ") {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid username")
	}

	customValidator := validator2.NewCustomValidator()
	translator, _ := customValidator.RegisterEnTranslation()
	errValidation := customValidator.Validator.Struct(newUser)
	errs := customValidator.TranslateError(errValidation, translator)
	errorsString := customValidator.GetErrorsString(errs)

	if errValidation != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorsString[0])
	}

	if pasval1, pasval2, pasval3, pasval4 := verifyPassword(newUser.Password); pasval1 == false || pasval2 == false || pasval3 == false || pasval4 == false {
		return echo.NewHTTPError(http.StatusBadRequest, "Password must have minimum 1 uppercase letter, 1 lowercase letter, 1 digit and 1 special character and needs to be minimum 8 characters long")
	}

	passwordCompare := signUp.SignUpUsecase.ValidatePassword(newUser.Password, t.ConfirmedPassword)

	if !passwordCompare {
		return echo.NewHTTPError(http.StatusBadRequest, "Enter same passwords")
	}
	code, errR := signUp.SignUpUsecase.RegisterNewUser(ctx, newUser)
	if errR != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Redis failed!")
	}

	go usecase.SendMail(newUser.Email, newUser.Username, code)

	return ctx.JSON(http.StatusOK, "Successfull registration, please check your mail!")
}

func (signUp *signUp) ResendCode(ctx echo.Context) error {
	decoder := json.NewDecoder(ctx.Request().Body)

	type Email struct {
		Email	string	`json:"email"`
	}

	var req Email
	err := decoder.Decode(&req)


	policy := bluemonday.UGCPolicy()
	email := strings.TrimSpace(policy.Sanitize(req.Email))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error while unmarshalling request")
	}
	err, username, code := signUp.SignUpUsecase.ResendCode(email)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Invalid email")
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid email")
	}


	usecase.SendMail(email, username, code)
	return ctx.JSON(http.StatusOK, "Resend request successful, please check your email")
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

	policy := bluemonday.UGCPolicy()
	credentials.Email = strings.TrimSpace(policy.Sanitize(credentials.Email))
	credentials.VerificationCode = strings.TrimSpace(policy.Sanitize(credentials.VerificationCode))


	if validateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorsString[0])
	}

	code := credentials.VerificationCode
	email := credentials.Email

	user, err := signUp.SignUpUsecase.IsCodeValid(ctx, email, code)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Entered code is not valid")
	}

	_, err2 := signUp.SignUpUsecase.ConfirmAccount(ctx, user)

	if err2 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Znaci...")
	}

	return ctx.JSON(http.StatusOK, "Ok")
}


func verifyPassword(s string) (eightOrMore, number, upper, special bool)  {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
			letters++
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
			letters++
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			return false, false, false, false
		}
	}
	eightOrMore = letters >= 8
	return
}


