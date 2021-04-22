package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo"
	"log"
	"net/http"
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

func Auth2() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return c.JSON(http.StatusUnauthorized, "user hasn't logged in yet")
			}
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

/*func Authorize(adapter persist.Adapter, next echo.HandlerFunc) echo.MiddlewareFunc {
	return func(c echo.Context){


		err := auth.TokenValid(c.Request())
		if err != nil {
			c.JSON(http.StatusUnauthorized, "user hasn't logged in yet")
			c.Error(err)
			return err
		}


		metadata, err := auth.NewToken2().ExtractTokenMetadata(c.Request())


		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return err
		}

		ok, err := enforce(metadata.UserId, c.Request().URL.Path, c.Request().Method, adapter)

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
}*/

func enforce(id uint64, obj string, act string) (bool, error) {
	enforcer, _ := casbin.NewEnforcer("src/security/rbac-model/rbac_model.conf", "src/security/rbac-model/policy.csv")
	err := enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok, _ := enforcer.Enforce(id, obj, act)
	return ok, nil
}
