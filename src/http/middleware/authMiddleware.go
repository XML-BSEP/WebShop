package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"web-shop/domain"
	"web-shop/security/auth"
	"web-shop/usecase"
)

type AuthMiddleware struct {
	ur domain.RegisteredShopUserRepository
	redisUseCase usecase.RedisUsecase
}

func NewAuthMiddleware(userRepo domain.RegisteredShopUserRepository, redisUseCase usecase.RedisUsecase ) *AuthMiddleware {
	return &AuthMiddleware{
		ur: userRepo,
		redisUseCase: redisUseCase,
	}
}

func (au *AuthMiddleware) Authenticated() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			err := auth.TokenValid(c.Request())
			if err == nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Already logged in")

			}
			return next(c)
		}
	}
}

func (au *AuthMiddleware) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role := ""
			err := auth.TokenValid(c.Request())
			if err != nil {
				role = "anonymous"
				return c.JSON(http.StatusUnauthorized, "user hasn't logged in yet")

			}
			metadata, err := auth.NewToken2().ExtractTokenMetadata(c.Request())

			if err != nil {
				return c.JSON(http.StatusUnauthorized, "unauthorized")
			} else {

				if !au.redisUseCase.ExistsByKey(metadata.TokenUuid) {
					return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
				}

				role, err = au.ur.GetRoleById(uint(metadata.UserId))
				if err != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
				}
			}

			ok, err := enforce(role, c.Request().URL.Path, c.Request().Method)

			if err != nil {
				log.Println(err)
				echo.NewHTTPError(http.StatusInternalServerError, "error occurred when authorizing user")
				return err
			}

			if !ok {
				echo.NewHTTPError(http.StatusForbidden, "forbidden")
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
