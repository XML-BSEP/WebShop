package main

import (
	"fmt"
	"web-shop/http/middleware"
	"web-shop/http/router"
	"web-shop/infrastructure/database"
	"web-shop/infrastructure/seeder"
	"github.com/labstack/echo"
	"web-shop/interactor"
)

func main() {

	seeder.MigrateData()

	conn := database.NewDBConnection()
	i := interactor.NewInteractor(conn)
	handler := i.NewAppHandler()

	e := echo.New()

	middleware.NewMiddleware(e)
	router.NewRouter(e, handler)

	e.Logger.Fatal(e.Start("localhost:8080"))

	fmt.Println("Successfully connected!")

}
