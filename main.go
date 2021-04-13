package main

import (
	"fmt"

	//"web-shop/infrastructure/security/auth"
	"web-shop/infrastructure/seeder"
	//"web-shop/infrastructure"
	"github.com/labstack/echo"
	//"web-shop/infrastructure/persistance/datastore"
)

func main() {

	//conn := database.NewDBConnection()

	seeder.MigrateData()

	e := echo.New()


	e.Logger.Fatal(e.Start("localhost:8080"))

	//userRepo := datastore.NewRegisteredUserRepository(conn)
	//redisService := auth.RedisService{}
	//tk := auth.NewToken()
	//authenticate := infrastructure.NewAuthenticate(userRepo, redisService.Auth, tk)

	//e.POST("/login", authenticate.Login)
	fmt.Println("Successfully connected!")

}
