package main

import (
	"fmt"
	"github.com/labstack/echo"
	mid2 "github.com/labstack/echo/middleware"
	middleware2 "github.com/labstack/echo/middleware"
	"web-shop/http/middleware"
	"web-shop/http/router"
	"web-shop/infrastructure/database"
	_ "web-shop/infrastructure/seeder"
	"web-shop/interactor"
)

func main() {
	//seeder.MigrateData()

	conn := database.NewDBConnection()
	i := interactor.NewInteractor(conn)

	handler := i.NewAppHandler()

	authMiddleware := middleware.NewAuthMiddleware(i.NewRegisteredUserRepository(i.NewShopAccountRepository()), i.NewRedisUsecase())
	e := echo.New()
	e.Use(mid2.Recover())
	e.Use(mid2.Logger())
	e.Pre(mid2.HTTPSRedirect())
	e.Use(mid2.CORSWithConfig(mid2.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.HEAD},

	}))
	e.Use(middleware2.Secure())

	middleware.NewMiddleware(e)
	router.NewRouter(e, handler, *authMiddleware)

	e.Logger.Fatal(e.StartTLS("localhost:443", "certificate/DukeStrategicTechnologies-SN-9946396461889217640.crt", "certificate/DukeStrategicTechnologies9946396461889217640-key.key"))

	fmt.Println("Successfully connected!")


}
