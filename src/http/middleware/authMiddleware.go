package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"web-shop/domain"
	"web-shop/security/auth"
)

type AuthMiddleware struct {
	ur domain.RegisteredShopUserRepository
	redisUseCase auth.RedisUseCase
}

func NewAuthMiddleware(userRepo domain.RegisteredShopUserRepository, redisUseCase auth.RedisUseCase ) *AuthMiddleware {
	return &AuthMiddleware{
		ur: userRepo,
		redisUseCase: redisUseCase,
	}
}

func (au *AuthMiddleware) Auth2() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role := ""
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				role = "anonymous"
				return c.JSON(http.StatusUnauthorized, "user hasn't logged in yet")
			}
			err := auth.TokenValid(c.Request())
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "user hasn't logged in yet")

			}

			metadata, err := auth.NewToken2().ExtractTokenMetadata(c.Request())

			if err != nil {
				return c.JSON(http.StatusUnauthorized, "unauthorized")
			} else {
				role, err = au.ur.GetRoleById(metadata.UserId)
				if err != nil {
					return c.JSON(http.StatusUnauthorized, "unauthorized")
				}
			}

			ok, err := enforce(role, c.Request().URL.Path, c.Request().Method)

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

func enforce(role string, obj string, act string) (bool, error) {
	enforcer, _ := casbin.NewEnforcer("src/security/rbac-model/rbac_model.conf", "src/security/rbac-model/policy.csv")
	err := enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok, _ := enforcer.Enforce(role, obj, act)
	return ok, nil
}
