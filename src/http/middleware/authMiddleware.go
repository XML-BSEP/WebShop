package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
	"web-shop/security/auth"
)

/*func TokenAuthMiddleware(c echo.Context)  echo.MiddlewareFunc {

	err := auth.TokenValid(c.Request())
	if err != nil {
		c.JSON(http.StatusUnauthorized, "user hasn't logged in yet")
		return echo.ErrForbidden
	}

	return nil
}*/

func Authenticated() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			err := auth.TokenValid(c.Request())
			if err == nil {
				return c.JSON(http.StatusBadRequest, "Already logged in")

			}

			return next(c)
		}
	}
}


func Auth() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			err := auth.TokenValid(c.Request())
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "user hasn't logged in yet")

			}

			metadata, err := auth.NewToken2().ExtractTokenMetadata(c.Request())


			if err != nil {
				return c.JSON(http.StatusUnauthorized, "unauthorized")
			}

			ok, err := enforce(metadata.UserId, c.Request().URL.Path, c.Request().Method)

			if err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, "error occurred when authorizing user")
				return err
			}

			if !ok {
				c.JSON(http.StatusForbidden, "forbidden")
				return err
			}

			return next(c)
		}
	}
}


func enforce(id uint64, obj string, act string) (bool, error) {
	enforcer, _ := casbin.NewEnforcer("src/security/rbac-model/rbac_model.conf", "src/security/rbac-model/policy.csv")
	err := enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok, _ := enforcer.Enforce(strconv.FormatUint(id, 10), obj, act)
	return ok, nil
}
