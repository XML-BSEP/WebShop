package handler

import (
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"web-shop/domain"
	"web-shop/infrastructure/security/auth"
	"web-shop/infrastructure/security/password-verification"
)

type AuthenticateHandler interface {
	Login(c echo.Context) error
	Logout(c echo.Context) error
}



type Authenticate struct {
	us domain.RegisteredShopUserRepository
	tk auth.TokenInterface
}

func NewAuthenticate(uApp domain.RegisteredShopUserRepository, tk auth.TokenInterface) AuthenticateHandler {
	return &Authenticate{
		us: uApp,
		tk: tk,
	}
}

func (au *Authenticate) Login(c echo.Context) error {
	var account *domain.ShopAccount
	var tokenErr = map[string]string{}

	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")

	}

	validateUser := ValidateLoginInput(account)

	if len(validateUser) > 0 {
		return c.JSON(http.StatusUnprocessableEntity, validateUser)

	}

	u, userErr := au.us.GetUserDetailsByAccount(account)

	err := password_verification.VerifyPassword(account.Password, u.ShopAccount.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return c.JSON(http.StatusForbidden, "Invalid password")

	}

	if userErr != nil {
		return c.JSON(http.StatusInternalServerError, userErr)

	}

	ts, tErr := au.tk.CreateToken(u.PersonID)
	if tErr != nil {
		tokenErr["token_error"] = tErr.Error()
		return c.JSON(http.StatusUnprocessableEntity, tErr.Error())

	}

	//saveErr := au.rd.CreateAuth(u.PersonID, ts)
	//if saveErr != nil {
	//	return c.JSON(http.StatusInternalServerError, saveErr.Error())

	//}

	userData := make(map[string]interface{})
	userData["access_token"] = ts.AccessToken
	userData["refresh_token"] = ts.RefreshToken
	userData["id"] = u.PersonID
	userData["first_name"] = u.Person.Name
	userData["last_name"] = u.Person.Surname

	return c.JSON(http.StatusOK, userData)
}

func (au *Authenticate) Logout(c echo.Context) error {


	_, err := au.tk.ExtractTokenMetadata(c.Request())
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")

	}

	// deleteErr := au.rd.DeleteTokens(metadata)
	// if deleteErr != nil {
	// 	return c.JSON(http.StatusUnauthorized, deleteErr.Error())

	// }
	return c.JSON(http.StatusOK, "Successfully logged out")
}

// ValidateLoginInput /*
func ValidateLoginInput(account *domain.ShopAccount) map[string]string {
	var errorMessages = make(map[string]string)
	if account.Password == "" {
		errorMessages["password_required"] = "password is required"
	}
	if account.Username == "" {
		errorMessages["email_required"] = "email is required"
	}
	return errorMessages
}
