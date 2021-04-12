package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"web-shop/infrastructure/database"
	"web-shop/infrastructure/persistance/datastore"
)
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "webshop"
)




func main() {
	conn := database.NewDBConnection()

	addrRepo := datastore.NewAddressRepository(conn)


	e := echo.New()

	addr, _ := addrRepo.GetByID(context.TODO(), 1)

	e.GET("/addresses", func(c echo.Context) error {
		return c.JSON(200, addr)
	})

	e.Logger.Fatal(e.Start("localhost:8080"))


	fmt.Println("Successfully connected!")

		//r := gin.Default()
		//r.GET("/ping", func(c *gin.Context) {
		//	c.JSON(200, gin.H{
		//		"message": "pong",
		//	})
		//})
		//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}