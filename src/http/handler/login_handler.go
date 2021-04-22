package handler

import (
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"web-shop/domain"
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
)


type Authenticate struct {
	us domain.RegisteredShopUserRepository
	tk auth2.TokenInterface
}

func NewAuthenticate(uApp domain.RegisteredShopUserRepository, tk auth2.TokenInterface) AuthenticateHandler {
	return &Authenticate{
		us: uApp,
		tk: tk,
	}
}

func (au *Authenticate) Login(c echo.Context) error {
	var account *domain.ShopAccount
	var tokenErr = map[string]string{}

	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, invalidJson)

	}

	validateUser := ValidateLoginInput(account)

	if len(validateUser) > 0 {
		return c.JSON(http.StatusUnprocessableEntity, validateUser)

	}

	u, userErr := au.us.GetUserDetailsByAccount(account)

	err := password_verification2.VerifyPassword(account.Password, u.ShopAccount.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return c.JSON(http.StatusForbidden, invalidPassword)

	}

	if userErr != nil {
		return c.JSON(http.StatusInternalServerError, userErr)
	}

	ts, tErr := au.tk.CreateToken(uint64(u.Model.ID))
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
	userData["id"] = u.Model.ID
	userData["first_name"] = u.Name
	userData["last_name"] = u.Surname

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
	return c.JSON(http.StatusOK, successfulLogout)
}

// ValidateLoginInput /*
func ValidateLoginInput(account *domain.ShopAccount) map[string]string {
	var errorMessages = make(map[string]string)
	if account.Password == "" {
		errorMessages["password_required"] = passwordIsRequired
	}
	if account.Username == "" {
		errorMessages["username_required"] = usernameIsRequired
	}
	return errorMessages
}
