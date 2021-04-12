package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"web-shop/infrastructure/database"
	"web-shop/infrastructure/persistance/datastore"
	"web-shop/infrastructure/seeder"
	"web-shop/usecase"
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
	//
	//conn.AutoMigrate(&domain.Address{}, &domain.Person{})
	//
	//a := domain.Address{City: "Novi Sad", State: "Serbia", Street: "Mise Dimitrijevica 2c", Zip: 21000}
	//a1 := domain.Address{City: "Novi Sad", State: "Serbia", Street: "Mise Dimitrijevica 3c", Zip: 21000}
	//a2 := domain.Address{City: "Novi Sad", State: "Serbia", Street: "Mise Dimitrijevica 4c", Zip: 21000}
	//a3 := domain.Address{City: "Novi Sad", State: "Serbia", Street: "Mise Dimitrijevica 5c", Zip: 21000}
	//
	//p1 := domain.Person{Address: a, Name: "Pera", Surname: "Peric", Phone: "1223124", DateOfBirth: time.Now(), Gender: 1}
	//
	//fmt.Print(a)
	//

	seeder.MigrateData()

	e := echo.New()

	//_, _ = addrRepo.Create(context.TODO(), &a)
	//_, _ = addrRepo.Create(context.TODO(), &a1)
	//_, _ = addrRepo.Create(context.TODO(), &a2)
	//_, _ = addrRepo.Create(context.TODO(), &a3)
	//_,_ = perRepo.Create(context.TODO(), &p1)

	//addrRepo := datastore.NewAddressRepository(conn)
	perRepo := datastore.NewPersonRepository(conn)
	regRepo := datastore.NewRegisteredUserRepository(conn)



	personUsecase := usecase.NewPersonUsecase(perRepo)
	pers, _ := personUsecase.Fetch(context.TODO())




	e.GET("/persons", func(c echo.Context) error {
		return c.JSON(200, pers)
	})

	regUsersUsecase := usecase.NewRegisteredUserUsecase(regRepo)
	regusers, _ := regUsersUsecase.Fetch(context.TODO())

	e.GET("/regusers", func(c echo.Context) error {
		return c.JSON(200, regusers)
	})

	e.Logger.Fatal(e.Start("localhost:8080"))

	fmt.Println("Successfully connected!")


}
