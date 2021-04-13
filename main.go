package main

import (
	"fmt"
	"github.com/labstack/echo"
	"web-shop/http/middleware"
	"web-shop/http/router"
	"web-shop/infrastructure/database"
	"web-shop/interactor"
)

func main() {

	//conn := database.NewDBConnection()

	//seeder.MigrateData()

	conn := database.NewDBConnection()
	i := interactor.NewInteractor(conn)
	handler := i.NewAppHandler()

	e := echo.New()

	middleware.NewMiddleware(e)
	router.NewRouter(e, handler)

	e.Logger.Fatal(e.Start("localhost:8080"))


	fmt.Println("Successfully connected!")

}
