package main

import (
	"fmt"
	"web-shop/infrastructure/seeder"
	"web-shop/usecase"
)

	"github.com/labstack/echo"
)

func main() {

	seeder.MigrateData()

	e := echo.New()

	e.Logger.Fatal(e.Start("localhost:8080"))

	fmt.Println("Successfully connected!")

}
