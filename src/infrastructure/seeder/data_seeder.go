package seeder

import (
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

	//conn.AutoMigrate(&domain.Address{})
	//conn.AutoMigrate(&domain.Person{})
	//conn.AutoMigrate(&domain.ShopAccount{})
	//conn.AutoMigrate(&domain.RegisteredShopUser{})
	//conn.AutoMigrate(&domain.Product{})
	//conn.AutoMigrate(&domain.Storage{})
	seedRoles(conn)
	seedAddresses(conn)
	//seedPersons(conn)
	//seedShopAccounts(conn)
	//seedRegisteredUsers(conn)
	//seedProducts(conn)
	//seedStorages(conn)
}

func seedAddresses(conn *gorm.DB) {

	addrRepo := datastore.NewAddressRepository(conn)

	a := domain.Address{City: "Novi Sad", State: "Serbia", Street: "Mise Dimitrijevica 2c", Zip: 21000}
	a1 := domain.Address{City: "Novi Sad", State: "Serbia", Street: "Mise Dimitrijevica 3c", Zip: 21000}
	a2 := domain.Address{City: "Novi Sad", State: "Serbia", Street: "Mise Dimitrijevica 4c", Zip: 21000}
	a3 := domain.Address{City: "Novi Sad", State: "Serbia", Street: "Mise Dimitrijevica 5c", Zip: 21000}

	addrRepo.Create(&a)
	addrRepo.Create(&a1)
	addrRepo.Create(&a2)
	addrRepo.Create(&a3)

}

func seedShopAccounts(conn *gorm.DB) {
	accRepo := datastore.NewShopAccountRepository(conn)

	acc := domain.ShopAccount{Username: "password", Password: "$2y$12$1duXzw4C3iYpZpU14rh0A.cjbF2kWdqKlUfsMWJOpRGmcFFHfok36"}
	accRepo.Create(&acc)
	acc = domain.ShopAccount{Username: "password1", Password: "$2y$12$1duXzw4C3iYpZpU14rh0A.cjbF2kWdqKlUfsMWJOpRGmcFFHfok36"}

	accRepo.Create(&acc)

}

func seedRegisteredUsers(conn *gorm.DB) {
	shopAccountRepo := datastore.NewShopAccountRepository(conn)
	regRepo := datastore.NewRegisteredUserRepository(conn, shopAccountRepo)

	acc1, _ := datastore.NewShopAccountRepository(conn).GetByID(1)
	acc2, _ := datastore.NewShopAccountRepository(conn).GetByID(2)

	role1, _ := datastore.NewRoleRepository(conn).GetByID(1)
	role2, _ := datastore.NewRoleRepository(conn).GetByID(2)

	regUser1 := domain.RegisteredShopUser{Email: "a@a.com", ShopAccount: *acc1, Role: *role1}
	regRepo.Create(&regUser1)

	regUser2 := domain.RegisteredShopUser{Email: "a2@a.com", ShopAccount: *acc2, Role: *role2}
	regRepo.Create(&regUser2)

}

func seedProducts(conn *gorm.DB) {
	prodRepo := datastore.NewProductRepository(conn)

	product1 := domain.Product{Name: "Product1", Price: 6969, Image: "assets/randompic1.jpg", Currency: 1, Category: "Tech", Available: 123}
	product2 := domain.Product{Name: "Product2", Price: 69420, Image: "assets/randompic2.jpg", Currency: 1, Category: "Tech", Available: 0}
	product3 := domain.Product{Name: "Product3", Price: 1512, Image: "assets/randompic4.jpg", Currency: 1, Category: "Clothes", Available: 69}

	prodRepo.Create(&product1)
	prodRepo.Create(&product2)
	prodRepo.Create(&product3)

}

func seedStorages(conn *gorm.DB) {

	p1, _ := datastore.NewProductRepository(conn).GetByID(1)
	p2, _ := datastore.NewProductRepository(conn).GetByID(2)
	p3, _ := datastore.NewProductRepository(conn).GetByID(3)

	storageRepo := datastore.NewStorageRepository(conn)

	s1 := domain.Storage{Product: *p1, Available: 666}
	s2 := domain.Storage{Product: *p2, Available: 12}
	s3 := domain.Storage{Product: *p3, Available: 420}

	storageRepo.Create(&s1)
	storageRepo.Create(&s2)
	storageRepo.Create(&s3)
}

func seedRoles(conn *gorm.DB) {
	roleRepo := datastore.NewRoleRepository(conn)

	r1 := &domain.Role{RoleName: "admin"}
	r2 := &domain.Role{RoleName: "user"}

	roleRepo.Create(r1)
	roleRepo.Create(r2)
}
func seedShoppingCarts(conn *gorm.DB) {
	// prodRepo := datastore.NewProductRepository(conn)
	// p1, _ := datastore.NewProductRepository(conn).GetByID(context.TODO(), 1)
	// p2, _ := datastore.NewProductRepository(conn).GetByID(context.TODO(), 2)
	// p3, _ := datastore.NewProductRepository(conn).GetByID(context.TODO(), 3)

}
