package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"web-shop/domain"
	"web-shop/infrastructure/security/auth"
	"web-shop/infrastructure/security/password-verification"
)

type AuthenticateHandler interface {
	Login(c echo.Context) error
	Logout(c echo.Context) error
	Refresh(c echo.Context) error
}


type Authenticate struct {
	us domain.RegisteredShopUserRepository
	rd auth.AuthInterface
	tk auth.TokenInterface
}

func NewAuthenticate(uApp domain.RegisteredShopUserRepository, rd auth.AuthInterface, tk auth.TokenInterface) AuthenticateHandler {
	return &Authenticate{
		us: uApp,
		rd: rd,
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


	metadata, err := au.tk.ExtractTokenMetadata(c.Request())
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")

	}

	deleteErr := au.rd.DeleteTokens(metadata)
	if deleteErr != nil {
		return c.JSON(http.StatusUnauthorized, deleteErr.Error())

	}
	return c.JSON(http.StatusOK, "Successfully logged out")
}

func (au *Authenticate) Refresh(c echo.Context) error {
	mapToken := map[string]string{}
	if err := c.Bind(&mapToken); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())

	}
	refreshToken := mapToken["refresh_token"]

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//any error may be due to token expiration
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())

	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return c.JSON(http.StatusUnauthorized, err)

	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			return c.JSON(http.StatusUnprocessableEntity, "Cannot get uuid")

		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, "Error occurred")

		}

		delErr := au.rd.DeleteRefresh(refreshUuid)
		if delErr != nil { //if any goes wrong
			return c.JSON(http.StatusUnauthorized, "unauthorized")

		}

		ts, createErr := au.tk.CreateToken(userId)
		if createErr != nil {
			return c.JSON(http.StatusForbidden, createErr.Error())

		}

		saveErr := au.rd.CreateAuth(userId, ts)
		if saveErr != nil {
			return c.JSON(http.StatusForbidden, saveErr.Error())

		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		return c.JSON(http.StatusCreated, tokens)
	} else {
		return c.JSON(http.StatusUnauthorized, "refresh token expired")
	}
}

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
