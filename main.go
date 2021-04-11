package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "webshop"
)


type Seed struct {
	Name string
	Run  func(*database.DBConn) error
}

func MigrateData(db *database.DBConn) {
	db.DB.AutoMigrate(&admin.Admin{})
	db.DB.AutoMigrate(&cert.ArchivedCert{})
	db.DB.AutoMigrate(&userOrService.UserOrService{})
}

func SeedData(db *database.DBConn) {
	for _, seed := range allSeeds() {
		if err := seed.Run(db); err != nil {
			log.Printf("Seed: '%s' failed with error: '%s'", seed.Name, err)
		}
	}
}

func allSeeds() []Seed {
	return []Seed{
		Seed{
			Name: "CreateAdmin1",
			Run: func(db *database.DBConn) error {
				a := admin.Admin{ID: 1, Username: "admin1", Email: "admin@email.com", Password: "admin1"}
				return admin.AddAdmin(&a, db)
			},
		},
		//Seed{
		//	Name: "User1",
		//	Run: func(db *database.DBConn) error {
		//		uos := userOrService.UserOrService{ID: 1, Username: "user1", Password: "user1"}
		//		return userOrService.AddUserOrServiceToDB(&uos, db)
		//	},
		//},
		//Seed{
		//	Name: "Service2",
		//	Run: func(db *database.DBConn) error {
		//		uos := userOrService.UserOrService{ID: 2, Username: "service2", Password: "service2"}
		//		return userOrService.AddUserOrServiceToDB(&uos, db)
		//	},
		//},
	}
}



func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")


}