package main

import (
	"fmt"
	"github.com/labstack/echo"
	"web-shop/infrastructure/seeder"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "webshop"
)

func main() {


	seeder.MigrateData()

	e := echo.New()

	e.Logger.Fatal(e.Start("localhost:8080"))

	fmt.Println("Successfully connected!")


}
