package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewMiddleware(e *echo.Echo) {
	e.Static("/static", "src/assets")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.HEAD},

	}))


}
