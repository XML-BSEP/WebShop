package seeder

import (
	"context"
	"time"
	"web-shop/domain"
	"web-shop/infrastructure/database"
	"web-shop/infrastructure/persistance/datastore"

	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(db *gorm.DB) error
}

func MigrateData() {
	conn := database.NewDBConnection()

	conn.AutoMigrate(&domain.Address{})
	conn.AutoMigrate(&domain.Person{})
	conn.AutoMigrate(&domain.ShopAccount{})
	conn.AutoMigrate(&domain.RegisteredShopUser{})
	conn.AutoMigrate(&domain.Product{})
	conn.AutoMigrate(&domain.Storage{})

	seedAddresses(conn)
	seedPersons(conn)
	seedShopAccounts(conn)
	seedRegisteredUsers(conn)
	seedProducts(conn)
	seedStorages(conn)
}

func seedAddresses(conn *gorm.DB) {

	addrRepo := datastore.NewAddressRepository(conn)

	a := domain.Address{City: "Novi Sad", State: "Serbia", Street: "Mise Dimitrijevica 2c", Zip: 21000}
	a1 := domain.Address{City: "Novi Sad", State: "Serbia", Street: "Mise Dimitrijevica 3c", Zip: 21000}
	a2 := domain.Address{City: "Novi Sad", State: "Serbia", Street: "Mise Dimitrijevica 4c", Zip: 21000}
	a3 := domain.Address{City: "Novi Sad", State: "Serbia", Street: "Mise Dimitrijevica 5c", Zip: 21000}

	addrRepo.Create(context.TODO(), &a)
	addrRepo.Create(context.TODO(), &a1)
	addrRepo.Create(context.TODO(), &a2)
	addrRepo.Create(context.TODO(), &a3)

}

func seedPersons(conn *gorm.DB) {
	perRepo := datastore.NewPersonRepository(conn)
	a1, _ := datastore.NewAddressRepository(conn).GetByID(context.TODO(), 1)
	a2, _ := datastore.NewAddressRepository(conn).GetByID(context.TODO(), 2)

	p1 := domain.Person{Address: *a1, Name: "Pera", Surname: "Peric", Phone: "1223124", DateOfBirth: time.Now(), Gender: 1}
	p2 := domain.Person{Address: *a2, Name: "Jovica", Surname: "Jovic", Phone: "918246", DateOfBirth: time.Now(), Gender: 1}

	perRepo.Create(context.TODO(), &p1)
	perRepo.Create(context.TODO(), &p2)

}

func seedShopAccounts(conn *gorm.DB) {
	accRepo := datastore.NewShopAccountRepository(conn)
	acc := domain.ShopAccount{Username: "password", Password: "$2y$12$1duXzw4C3iYpZpU14rh0A.cjbF2kWdqKlUfsMWJOpRGmcFFHfok36 "}
	accRepo.Create(context.TODO(), &acc)
	acc = domain.ShopAccount{Username: "password1", Password: "$2y$12$1duXzw4C3iYpZpU14rh0A.cjbF2kWdqKlUfsMWJOpRGmcFFHfok36 "}
	accRepo.Create(context.TODO(), &acc)

}

func seedRegisteredUsers(conn *gorm.DB) {
	shopAccountRepo := datastore.NewShopAccountRepository(conn)
	regRepo := datastore.NewRegisteredUserRepository(conn, shopAccountRepo)

	p1, _ := datastore.NewPersonRepository(conn).GetByID(context.TODO(), 1)
	p2, _ := datastore.NewPersonRepository(conn).GetByID(context.TODO(), 2)

	acc1, _ := datastore.NewShopAccountRepository(conn).GetByID(context.TODO(), 1)
	acc2, _ := datastore.NewShopAccountRepository(conn).GetByID(context.TODO(), 2)

	regUser1 := domain.RegisteredShopUser{Email: "a@a.com", Person: *p1, ShopAccount: *acc1}
	regRepo.Create(context.TODO(), &regUser1)

	regUser2 := domain.RegisteredShopUser{Email: "a2@a.com", Person: *p2, ShopAccount: *acc2}
	regRepo.Create(context.TODO(), &regUser2)

}

func seedProducts(conn *gorm.DB) {
	prodRepo := datastore.NewProductRepository(conn)

	product1 := domain.Product{Name: "Product1", Price: 6969, Image: "assets/randompic1.jpg", Currency: 1}
	product2 := domain.Product{Name: "Product2", Price: 69420, Image: "assets/randompic2.jpg", Currency: 1}
	product3 := domain.Product{Name: "Product3", Price: 1512, Image: "assets/randompic4.jpg", Currency: 1}

	prodRepo.Create(context.TODO(), &product1)
	prodRepo.Create(context.TODO(), &product2)
	prodRepo.Create(context.TODO(), &product3)

}

func seedStorages(conn *gorm.DB) {

	p1, _ := datastore.NewProductRepository(conn).GetByID(context.TODO(), 1)
	p2, _ := datastore.NewProductRepository(conn).GetByID(context.TODO(), 2)
	p3, _ := datastore.NewProductRepository(conn).GetByID(context.TODO(), 3)

	storageRepo := datastore.NewStorageRepository(conn)

	s1 := domain.Storage{Product: *p1, Available: 666}
	s2 := domain.Storage{Product: *p2, Available: 12}
	s3 := domain.Storage{Product: *p3, Available: 420}

	storageRepo.Create(context.TODO(), &s1)
	storageRepo.Create(context.TODO(), &s2)
	storageRepo.Create(context.TODO(), &s3)
}

func seedShoppingCarts(conn *gorm.DB) {
	// prodRepo := datastore.NewProductRepository(conn)
	// p1, _ := datastore.NewProductRepository(conn).GetByID(context.TODO(), 1)
	// p2, _ := datastore.NewProductRepository(conn).GetByID(context.TODO(), 2)
	// p3, _ := datastore.NewProductRepository(conn).GetByID(context.TODO(), 3)

}
