package seeder

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"web-shop/domain"
	"web-shop/infrastructure/persistance/datastore"

	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(db *gorm.DB) error
}

func MigrateData(conn *gorm.DB) {
	conn.Migrator().DropTable(&domain.Address{})
	conn.Migrator().DropTable(&domain.ShopAccount{})
	conn.Migrator().DropTable(&domain.RegisteredShopUser{})
	conn.Migrator().DropTable(&domain.Product{})
	conn.Migrator().DropTable(&domain.Storage{})
	conn.Migrator().DropTable(&domain.Category{})
	conn.Migrator().DropTable(&domain.Image{})
	conn.Migrator().DropTable(&domain.Role{})
	conn.Migrator().DropTable(&domain.ShoppingCartItem{})
	conn.Migrator().DropTable(&domain.AdminToken{})



	conn.AutoMigrate(&domain.Address{})
	//conn.AutoMigrate(&domain.{})
	conn.AutoMigrate(&domain.ShopAccount{})
	conn.AutoMigrate(&domain.RegisteredShopUser{})
	conn.AutoMigrate(&domain.Product{})
	conn.AutoMigrate(&domain.Storage{})
	conn.AutoMigrate(&domain.Category{})
	conn.AutoMigrate(&domain.Image{})
	conn.AutoMigrate(&domain.Role{})
	conn.AutoMigrate(&domain.ShoppingCartItem{})
	conn.AutoMigrate(&domain.AdminToken{})
	//conn.AutoMigrate(&domain.Storage{})
	seedRoles(conn)
	seedAddresses(conn)


	seedShopAccounts(conn)
	seedRegisteredUsers(conn)
	seedCategories(conn)
	seedProducts(conn)

}

func dropDatabase(db string, mongoCli *mongo.Client, ctx *context.Context){
	err := mongoCli.Database(db).Drop(*ctx)
	if err != nil {
		return
	}
}

func SeedMongoData(db string, mongoCli *mongo.Client, ctx *context.Context){

	dropDatabase(db,mongoCli, ctx)

	if cnt,_ := mongoCli.Database(db).Collection("orders").EstimatedDocumentCount(*ctx, nil); cnt == 0{
		_ = mongoCli.Database(db).Collection("orders")
	}



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

	acc := domain.ShopAccount{Username: "Agent1", Password: "$2y$12$4FQN3yE4HwYwDuI1grAar.KPqCc/kvki6Bps9m4t4qWRuLaILdnBK"}
	acc1 := domain.ShopAccount{Username: "Agent2", Password: "$2y$12$4FQN3yE4HwYwDuI1grAar.KPqCc/kvki6Bps9m4t4qWRuLaILdnBK"}

	accRepo.Create(&acc)
	accRepo.Create(&acc1)

	acc2 := domain.ShopAccount{Username: "User", Password: "$2y$12$4FQN3yE4HwYwDuI1grAar.KPqCc/kvki6Bps9m4t4qWRuLaILdnBK"}

	accRepo.Create(&acc2)


}

func seedRegisteredUsers(conn *gorm.DB) {
	shopAccountRepo := datastore.NewShopAccountRepository(conn)
	regRepo := datastore.NewRegisteredUserRepository(conn, shopAccountRepo)

	acc1, _ := datastore.NewShopAccountRepository(conn).GetByID(1)
	acc2, _ := datastore.NewShopAccountRepository(conn).GetByID(2)
	acc3, _ := datastore.NewShopAccountRepository(conn).GetByID(3)

	role1, _ := datastore.NewRoleRepository(conn).GetByID(1)
	role2, _ := datastore.NewRoleRepository(conn).GetByID(2)

	regUser1 := domain.RegisteredShopUser{Email: "vlajkovj31@gmail.com", ShopAccount: *acc1, Role: *role1}

	regUser2 := domain.RegisteredShopUser{Email: "alexignjat1998@gmail.com", ShopAccount: *acc2, Role: *role1}

	regUser3 := domain.RegisteredShopUser{Email: "a2@a.com", ShopAccount: *acc3, Role: *role2}


	regRepo.Create(&regUser1)

	regRepo.Create(&regUser2)
	regRepo.Create(&regUser3)

}

func seedProducts(conn *gorm.DB) {
	prodRepo := datastore.NewProductRepository(conn)
	catRepo := datastore.NewCategoryRepository(conn)
	shopAccRepo := datastore.NewShopAccountRepository(conn)

	cat1, _ := catRepo.GetByID(1)
	cat2, _ := catRepo.GetByID(2)

	u1, _ :=shopAccRepo.GetByID(1)
	u2, _ :=shopAccRepo.GetByID(2)

	images1 := make([]domain.Image, 2)
	images1[0] = domain.Image{Path: "1/randompic1.jpg", Timestamp: time.Now().Add(40)}
	images1[1] = domain.Image{Path: "1/randompic2.jpg", Timestamp: time.Now().Add(40)}

	images2 := make([]domain.Image, 2)
	images2[0] = domain.Image{Path: "2/randompic3.jpg", Timestamp: time.Now().Add(10)}
	images2[1] = domain.Image{Path: "2/randompic4.jpg", Timestamp: time.Now().Add(15)}

	images3 := make([]domain.Image, 2)
	images3[0] = domain.Image{Path: "3/randompic5.jpg", Timestamp: time.Now()}
	images3[1] = domain.Image{Path: "3/randompic6.jpg", Timestamp: time.Now()}

	product1 := domain.Product{Name: "Product1", Price: 699, Images: images1, Category: *cat1, SerialNumber: 123, Description: "Ide gas1", Available: 6, ShopAccount: *u1}
	product2 := domain.Product{Name: "Product2", Price: 420, Images: images2, Category: *cat2, SerialNumber: 1234, Description: "Ide gas2", Available: 7, ShopAccount: *u2}
	product3 := domain.Product{Name: "Product3", Price: 1512, Images: images3, Category: *cat1, SerialNumber: 12345, Description: "Ide gas3", Available: 69, ShopAccount: *u1}

	prodRepo.Create(&product1)
	prodRepo.Create(&product2)
	prodRepo.Create(&product3)

}

func seedCategories(conn *gorm.DB) {
	catRepo := datastore.NewCategoryRepository(conn)

	category1 := domain.Category{Name: "Tech"}
	category2 := domain.Category{Name: "Makeup"}

	catRepo.Create(&category1)
	catRepo.Create(&category2)
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
