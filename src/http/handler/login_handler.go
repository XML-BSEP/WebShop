package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"os"
	"strconv"
	"strings"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
	auth2 "web-shop/security/auth"
	password_verification2 "web-shop/security/password-verification"
)

type AuthenticateHandler interface {
	Login(c echo.Context) error
	Logout(c echo.Context) error
	Refresh(c echo.Context) error
}

const (
	usernameIsRequired = "Username is required!"
	passwordIsRequired = "Password is required!"
	invalidJson        = "Invalid JSON!"
	invalidPassword    = "Invalid password"
	successfulLogout   = "Successfully logged out!"
	invalidEmail = "Invalid email!"
	cannotFindUiid = "Cannot get uuid"
	parseError = "Error while parsing occured"
	unauthorized = "Unauthorized"
	refreshTokenExpired = "Refresh token expired"
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

	policy := bluemonday.UGCPolicy();
	account.Email = strings.TrimSpace(policy.Sanitize(account.Email))
	account.Password = strings.TrimSpace(policy.Sanitize(account.Password))

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, invalidJson)
	}

	validateUser := ValidateLoginInput(account)

	if len(validateUser) > 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, validateUser)
	}

	u, userErr := au.us.GetUserDetailsFromEmail(account.Email)

	if userErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, invalidEmail)

	}

	accDetails, accErr := au.us.GetAccountDetailsFromUser(u)

	if accErr != nil {
		return accErr
	}

	err = password_verification2.VerifyPassword(account.Password, accDetails.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, invalidPassword)

	}

	ts, tErr := au.tk.CreateToken(uint64(u.Model.ID))
	if tErr != nil {
		tokenErr["token_error"] = tErr.Error()
		return echo.NewHTTPError(http.StatusUnprocessableEntity, tErr.Error())

	}

	saveErr := au.au.CreateAuth(u.ID, ts)
	if saveErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, saveErr.Error())

	}

	userData := make(map[string]interface{})
	userData["access_token"] = ts.AccessToken
	userData["refresh_token"] = ts.RefreshToken
	userData["id"] = u.Model.ID
	role, err := au.us.GetRoleById(u.Model.ID)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, invalidEmail)
	}
	userData["role"] = role

	return c.JSON(http.StatusOK, userData)
}


func (au *Authenticate) Logout(c echo.Context) error {

	metadata, err := au.tk.ExtractTokenMetadata(c.Request())
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")

	}

	deleteErr := au.au.DeleteTokens(metadata)

	if deleteErr != nil {
		echo.NewHTTPError(http.StatusUnauthorized, deleteErr.Error())
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

func (au *Authenticate) Refresh(c echo.Context) error {
	mapToken := map[string]string{}

	decoder := json.NewDecoder(c.Request().Body)

	err := decoder.Decode(&mapToken)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, invalidJson)
	}
	refreshToken := mapToken["refresh_token"]

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return c.JSON(http.StatusUnauthorized, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, cannotFindUiid)
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, parseError)
		}

		delErr := au.au.DeleteRefresh(refreshUuid)
		if delErr != nil { //if any goes wrong
			return echo.NewHTTPError(http.StatusUnauthorized, unauthorized)
		}

		ts, createErr := au.tk.CreateToken(userId)
		if createErr != nil {
			return echo.NewHTTPError(http.StatusForbidden, createErr.Error())

		}

		saveErr := au.au.CreateAuth(uint(userId), ts)
		if saveErr != nil {
			return echo.NewHTTPError(http.StatusForbidden, saveErr.Error())

		}


		role, err := au.us.GetRoleById(uint(userId))
		if err != nil {
			echo.NewHTTPError(http.StatusInternalServerError, invalidEmail)
		}


		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
			"role" : role,
			"id" : string(userId),
		}
		return c.JSON(http.StatusCreated, tokens)
	} else {
		return echo.NewHTTPError(http.StatusUnauthorized, refreshTokenExpired)
	}
}