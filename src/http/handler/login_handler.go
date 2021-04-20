package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
	auth2 "web-shop/security/auth"
	password_verification2 "web-shop/security/password-verification"
)

type AuthenticateHandler interface {
	Login(c echo.Context) error
	Logout(c echo.Context) error
}

const (
	usernameIsRequired = "Username is required!"
	passwordIsRequired = "Password is required!"
	invalidJson        = "Invalid JSON!"
	invalidPassword    = "Invalid password"
	successfulLogout   = "Successfully logged out!"
	invalidEmail = "Invalid email!"
)


type Authenticate struct {
	us domain.RegisteredShopUserRepository
	tk auth2.TokenInterface
	au auth2.AuthInterface
}

func NewAuthenticate(uApp domain.RegisteredShopUserRepository, tk auth2.TokenInterface, au auth2.AuthInterface) AuthenticateHandler {
	return &Authenticate{
		us: uApp,
		tk: tk,
		au: au,
	}
}

func (au *Authenticate) Login(c echo.Context) error {
	var account *dto.AuthenticationDto
	var tokenErr = map[string]string{}


	decoder := json.NewDecoder(c.Request().Body)

	err := decoder.Decode(&account)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, invalidJson)
	}

	validateUser := ValidateLoginInput(account)

	if len(validateUser) > 0 {
		return c.JSON(http.StatusUnprocessableEntity, validateUser)
	}

	u, userErr := au.us.GetUserDetailsFromEmail(account.Email)

	if userErr != nil {
		return c.JSON(http.StatusInternalServerError, invalidEmail)

	}

	accDetails, accErr := au.us.GetAccountDetailsFromUser(u)

	if accErr != nil {
		return accErr
	}

	err = password_verification2.VerifyPassword(account.Password, accDetails.Password)
	if err != nil {
		return c.JSON(http.StatusForbidden, invalidPassword)

	}
	
	ts, tErr := au.tk.CreateToken(uint64(u.Model.ID))
	if tErr != nil {
		tokenErr["token_error"] = tErr.Error()
		return c.JSON(http.StatusUnprocessableEntity, tErr.Error())

	}

	saveErr := au.au.CreateAuth(u.ID, ts)
	if saveErr != nil {
		return c.JSON(http.StatusInternalServerError, saveErr.Error())

	}

	userData := make(map[string]interface{})
	userData["access_token"] = ts.AccessToken
	userData["refresh_token"] = ts.RefreshToken
	userData["id"] = u.Model.ID
	userData["first_name"] = u.Name
	userData["last_name"] = u.Surname

	return c.JSON(http.StatusOK, userData)
}


func (au *Authenticate) Logout(c echo.Context) error {

	metadata, err := au.tk.ExtractTokenMetadata(c.Request())
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")

	}

	deleteErr := au.au.DeleteTokens(metadata)

	if deleteErr != nil {
		c.JSON(http.StatusUnauthorized, deleteErr.Error())
		return deleteErr

	}
	return c.JSON(http.StatusOK, successfulLogout)
}

func ValidateLoginInput(account *dto.AuthenticationDto) map[string]string {
	var errorMessages = make(map[string]string)
	if account.Password == "" {
		errorMessages["password_required"] = passwordIsRequired
	}
	if account.Email == "" {
		errorMessages["username_required"] = usernameIsRequired
	}
	return errorMessages
}
