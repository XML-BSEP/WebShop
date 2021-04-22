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


	//seeder.MigrateData()

	conn := database.NewDBConnection()
	i := interactor.NewInteractor(conn)
	handler := i.NewAppHandler()

	e := echo.New()

	middleware.NewMiddleware(e)
	router.NewRouter(e, handler)



	e.Logger.Fatal(e.StartTLS("localhost:443", "certificate/DukeStrategicTechnologies-SN-9946396461889217640.crt", "certificate/DukeStrategicTechnologies9946396461889217640-key.key"))
	//e.Logger.Fatal(e.Start(":8080"))


	fmt.Println("Successfully connected!")


}
