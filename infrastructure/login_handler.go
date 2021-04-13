package infrastructure

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

type Authenticate struct {
	us domain.RegisteredShopUserRepository
	rd auth.AuthInterface
	tk auth.TokenInterface
}

func NewAuthenticate(uApp domain.RegisteredShopUserRepository, rd auth.AuthInterface, tk auth.TokenInterface) *Authenticate {
	return &Authenticate{
		us: uApp,
		rd: rd,
		tk: tk,
	}
}

func (au *Authenticate) Login(c *echo.Context) {
	var account *domain.ShopAccount
	var tokenErr = map[string]string{}

	if err := (*c).Bind(&account); err != nil {
		(*c).JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	validateUser := ValidateLoginInput(account)

	if len(validateUser) > 0 {
		(*c).JSON(http.StatusUnprocessableEntity, validateUser)
		return
	}

	u, userErr := au.us.GetUserDetailsByUsername(account)

	err := password_verification.VerifyPassword(account.Password, u.ShopAccount.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		(*c).JSON(http.StatusForbidden, "Invalid password")
		return
	}

	if userErr != nil {
		(*c).JSON(http.StatusInternalServerError, userErr)
		return
	}

	ts, tErr := au.tk.CreateToken(u.PersonID)
	if tErr != nil {
		tokenErr["token_error"] = tErr.Error()
		(*c).JSON(http.StatusUnprocessableEntity, tErr.Error())
		return
	}

	saveErr := au.rd.CreateAuth(u.PersonID, ts)
	if saveErr != nil {
		(*c).JSON(http.StatusInternalServerError, saveErr.Error())
		return
	}

	userData := make(map[string]interface{})
	userData["access_token"] = ts.AccessToken
	userData["refresh_token"] = ts.RefreshToken
	userData["id"] = u.PersonID
	userData["first_name"] = u.Person.Name
	userData["last_name"] = u.Person.Surname

	(*c).JSON(http.StatusOK, userData)
}

func (au *Authenticate) Logout(c echo.Context) {


	metadata, err := au.tk.ExtractTokenMetadata(c.Request())
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	deleteErr := au.rd.DeleteTokens(metadata)
	if deleteErr != nil {
		c.JSON(http.StatusUnauthorized, deleteErr.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

func (au *Authenticate) Refresh(c echo.Context) {
	mapToken := map[string]string{}
	if err := c.Bind(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
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
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, "Cannot get uuid")
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Error occurred")
			return
		}

		delErr := au.rd.DeleteRefresh(refreshUuid)
		if delErr != nil { //if any goes wrong
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}

		ts, createErr := au.tk.CreateToken(userId)
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}

		saveErr := au.rd.CreateAuth(userId, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		c.JSON(http.StatusUnauthorized, "refresh token expired")
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
