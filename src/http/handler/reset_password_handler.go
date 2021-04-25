package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"strings"
	"unicode"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
	password_verification "web-shop/security/password-verification"
	"web-shop/usecase"
)

type resetPassword struct {
	RegisteredShopUserUsecase    domain.RegisteredShopUserUsecase
	RandomStringGeneratorUsecase usecase.RandomStringGeneratorUsecase
}



type ResetPasswordHandler interface {
	SendResetMail(ctx echo.Context) error
	ResetPassword(ctx echo.Context) error
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

	policy := bluemonday.UGCPolicy();
	req.Email = strings.TrimSpace(policy.Sanitize(req.Email))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, errE :=  r.RegisteredShopUserUsecase.ExistByUsernameOrEmail(ctx, "", req.Email)
	if errE != nil {
		return ctx.JSON(http.StatusBadRequest, "User not found!")
	}

	var code string
	code = r.RandomStringGeneratorUsecase.RandomStringGenerator(8)

	go usecase.SendRestartPasswordMail(user.Email , code)
	hashedCode, err := password_verification.Hash(code)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error")
	}


	r.RegisteredShopUserUsecase.SaveCodeToRedis(string(hashedCode), req.Email)

	return ctx.JSON(http.StatusOK, "Successfully mail sent!")
}

func (r *resetPassword) ResetPassword(ctx echo.Context) error {
	decoder := json.NewDecoder(ctx.Request().Body)

	var resetDto dto.ResetPassDTO

	err := decoder.Decode(&resetDto)

	policy := bluemonday.UGCPolicy();
	resetDto.Email = strings.TrimSpace(policy.Sanitize(resetDto.Email))
	resetDto.Password = strings.TrimSpace(policy.Sanitize(resetDto.Password))
	resetDto.VerificationCode = strings.TrimSpace(policy.Sanitize(resetDto.VerificationCode))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if pasval1, pasval2, pasval3, pasval4 := verifyPassword(resetDto.Password); pasval1 == false || pasval2 == false || pasval3 == false || pasval4 == false {
		return ctx.JSON(http.StatusBadRequest, "Password must have minimum 1 uppercase letter, 1 lowercase letter, 1 digit and 1 special character and needs to be minimum 8 characters long")
	}

	errorMessage := r.RegisteredShopUserUsecase.ResetPassword(resetDto)

	if errorMessage != "" {
		return ctx.JSON(http.StatusInternalServerError, errorMessage)
	}

	return ctx.JSON(http.StatusOK, "password successfully changed")


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
