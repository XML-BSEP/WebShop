package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"unicode"
	"web-shop/domain"
	"web-shop/usecase"
)

type resetPassword struct {
	RegistoredShopUserUsecase domain.RegisteredShopUserUsecase
	RandomStringGeneratorUsecase usecase.RandomStringGeneratorUsecase
}



type ResetPasswordHandler interface {
	SendResetMail(ctx echo.Context) error
}


func NewResetPasswordHandler(r domain.RegisteredShopUserUsecase, generatorUsecase usecase.RandomStringGeneratorUsecase) ResetPasswordHandler {
	return &resetPassword{r,generatorUsecase }
}

func (r *resetPassword) SendResetMail(ctx echo.Context) (err error) {
	decoder := json.NewDecoder(ctx.Request().Body)

	type Email struct {
		Email	string	`json:"email"`
	}

	var req Email
	err = decoder.Decode(&req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, errE :=  r.RegistoredShopUserUsecase.ExistByUsernameOrEmail(ctx, "", req.Email)
	if errE != nil {
		return ctx.JSON(http.StatusBadRequest, "User not found!")
	}

	var code string
	code = r.RandomStringGeneratorUsecase.RandomStringGenerator(8)

	go usecase.SendRestartPasswordMail(user.Email , code)

	return ctx.JSON(http.StatusOK, "Successfully mail sent!")


}

func verifyResetPassword(s string) (eightOrMore, number, upper, special bool)  {
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
