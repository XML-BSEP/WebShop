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

	e.Logger.Fatal(e.StartTLS("localhost:443", "certificate/DukeStrategicTechnologies-SN-17502617923117970082.pem", "certificate/DukeStrategicTechnologies17502617923117970082-key.pem"))


	fmt.Println("Successfully connected!")

}
