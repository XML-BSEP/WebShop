package usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo"
	"os"
	"strconv"
	"strings"
	"time"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

type productUseCase struct {
	ShopAccountRepository domain.ShopAccountRepository
	ProductRepository domain.ProductRepository
	CategoryRepository domain.CategoryRepository
	ImageRepository domain.ImageRepository
}

func (p *productUseCase) GetProductDetails(ctx context.Context, productId uint) (*domain.Product, error) {
	product, err := p.ProductRepository.GetProductDetails(ctx,productId)
	if err!=nil{
		return nil,err
	}
	images, err1 :=p.ImageRepository.GetByProduct(ctx, productId)
	if err1!=nil{
		return nil, err
	}

	product.Images = images
	return product, nil
}

func (p *productUseCase) GetAllProductsInUsersShop(ctx echo.Context, userId uint) ([]*domain.Product, error) {
	return p.ProductRepository.GetAllProductsInUsersShop(ctx, userId)
}

func (p *productUseCase) GetAllAvailableProductsInUsersShop(ctx echo.Context, userId uint) ([]*domain.Product, error) {
	return p.ProductRepository.GetAllAvailableProductsInUsersShop(ctx,userId)
}

func (p *productUseCase) GetBySerial(serial uint64) (*domain.Product, error) {
	return p.ProductRepository.GetBySerial(serial)
}

func (p *productUseCase) Count() (int64, error) {
	return p.ProductRepository.Count()
}

func (p *productUseCase) FilterByCategory(userId uint ,name string, category string, priceRangeStart uint, priceRangeEnd uint, limit int, offset int, order string) ([]*domain.Product, error) {
	if name == "" {
		name = "%"
	}
	if category == "" {
		category = "%"
	}
	return p.ProductRepository.FilterByCategory(userId,name, category, priceRangeStart, priceRangeEnd, limit, offset, order)
}

func (p *productUseCase) GetProductsWithConditionOrderedByPrice(low uint, high uint, category string, limit int, offset int, order int) ([]*domain.Product, error) {
	return p.ProductRepository.GetProductsWithConditionOrderedByPrice(low,high,category,limit,offset,order)
}

func (p *productUseCase) GetProductsWithConditionOrderedByName(low uint, high uint, category string, limit int, offset int, order int) ([]*domain.Product, error) {
	return p.ProductRepository.GetProductsWithConditionOrderedByName(low,high,category,limit,offset,order)
}

func (p *productUseCase) GetByNameOrderByPrice(name string, limit int, offset int, order int) ([]*domain.Product, error) {
	return p.ProductRepository.GetByNameOrderByPrice(name, limit, offset, order)
}

func (p *productUseCase) GetByNameOrderByName(name string, limit int, offset int, order int) ([]*domain.Product, error) {
	return p.ProductRepository.GetByNameOrderByName(name, limit, offset, order)
}

func (p *productUseCase) GetByName(name string, limit int, offset int) ([]*domain.Product, error) {
	return p.ProductRepository.GetByName(name, limit, offset)
}

func (p *productUseCase) GetProductsWithCondition(low uint, high uint, category string, limit int, offset int) ([]*domain.Product, error) {
	return p.ProductRepository.GetProductsWithCondition(low, high, category, limit, offset)
}

func (p *productUseCase) GetProductsWithCategory(category string) ([]*domain.Product, error) {
	return p.ProductRepository.GetProductsWithCategory(category)
}

func (p *productUseCase) GetWithPriceRange(low uint, high uint) ([]*domain.Product, error){
	return p.ProductRepository.GetWithPriceRange(low, high)
}
func (p *productUseCase) Fetch(ctx echo.Context) ([]*domain.Product, error) {
	return p.ProductRepository.Fetch()
}

func (p *productUseCase) GetByID(ctx echo.Context, id uint) (*domain.Product, error) {
	return p.ProductRepository.GetByID(id)
}
func CheckExistance(newImgs []string, oldImage domain.Image) bool{

	for _, newImg := range newImgs{
		if strings.Contains(newImg, oldImage.Path){
			return true
		}
	}
	return false
}
func (p *productUseCase) Update(ctx echo.Context, prod *dto.EditProduct) (*domain.Product, error) {
	cat, err := p.CategoryRepository.GetByName(prod.Category)
	if err != nil{
		return nil, err
	}
	serialNumber, _ := strconv.ParseUint(prod.SerialNumber,10, 64)

	oldProduct, _  := p.ProductRepository.GetBySerialAndUserId(serialNumber, prod.UserId)
	folderName := strconv.FormatUint(uint64(oldProduct.Model.ID), 10)
	path1 := "./src/assets/" + folderName

	oldDir, _ := os.Getwd()

	os.Chdir(path1)

	//os.Mkdir(folderName, 0755)
	//
	//os.Chdir(folderName)

	var images []domain.Image
	var oldImages []*domain.Image
	oldImagesLikePath := folderName + "%"
	oldImages,_ = p.ImageRepository.GetyByPath(oldImagesLikePath)
	fmt.Println(oldImages)
	size := cap(prod.Images)
	if size>0{
		for _,image := range oldImages{
			if !CheckExistance(prod.Images, *image){
				fmt.Println(image.Path)
				os.Remove(strings.Split(image.Path, "/")[1])
				p.ImageRepository.Delete(image.Model.ID)
			}
		}
	}

	localhost := "localhost"
	if size>0{
		for i,_ := range prod.Images {
			if !strings.Contains(prod.Images[i], localhost){
				s := strings.Split(prod.Images[i], ",")
				a := strings.Split(s[0], "/")
				format := strings.Split(a[1], ";")

				dec, err := base64.StdEncoding.DecodeString(s[1])

				if err != nil {
					panic(err)
				}
				f, err := os.Create(strconv.Itoa(i) + "." + format[0])

				if err != nil {
					panic(err)
				}

				defer f.Close()

				if _, err := f.Write(dec); err != nil {
					panic(err)
				}
				if err := f.Sync(); err != nil {
					panic(err)
				}

				images = append(images, domain.Image{Path: folderName+"/"+strconv.Itoa(i) + "." + format[0], Timestamp: time.Now(), ProductId: oldProduct.Model.ID})
			}

		}
	}
	os.Chdir(oldDir)

	price, _ := strconv.ParseFloat(prod.Price, 64)
	av,_ :=strconv.Atoi(prod.Available)

	oldProduct.Price = price
	oldProduct.Available = uint(av)
	oldProduct.Images = images
	oldProduct.Category = *cat
	oldProduct.Description = prod.Description
	oldProduct.Name = prod.Name


	return p.ProductRepository.Update(oldProduct)
}

func (p *productUseCase) Create(ctx echo.Context, newProd *dto.NewProduct) (*domain.Product, error) {

	user, err := p.ShopAccountRepository.GetByID(newProd.UserId)
	if err != nil{
		return nil, err
	}
	cat, err := p.CategoryRepository.GetByName(newProd.Category)
	if err != nil{
		return nil, err
	}
	var list []*domain.Product

	list, _ = p.ProductRepository.Fetch()

	path2, _ := os.Getwd()
	fmt.Println(path2)
	len := len(list)
	t := strconv.Itoa(len+1)
	path1 := "./src/assets"
	os.Chdir(path1)

	os.Mkdir(t, 0755)

	os.Chdir(t)

	var images []domain.Image

	size := cap(newProd.Images)

	if size>0{
		for i,_ := range newProd.Images {

			s := strings.Split(newProd.Images[i], ",")
			a := strings.Split(s[0], "/")
			format := strings.Split(a[1], ";")

			dec, err := base64.StdEncoding.DecodeString(s[1])

			if err != nil {
				panic(err)
			}
			f, err := os.Create(strconv.Itoa(i) + "." + format[0])

			if err != nil {
				panic(err)
			}

			defer f.Close()

			if _, err := f.Write(dec); err != nil {
				panic(err)
			}
			if err := f.Sync(); err != nil {
				panic(err)
			}

			images = append(images, domain.Image{Path: strconv.Itoa(len+1)+"/"+strconv.Itoa(i) + "." + format[0], Timestamp: time.Now(), ProductId: uint(len + 1)})
		}
	}


	os.Chdir(path2)

	price, _ := strconv.ParseFloat(newProd.Price, 64)
	av,_ :=strconv.Atoi(newProd.Available)
	prod:=domain.Product{Available: uint(av), Price: price, Name: newProd.Name, Category: *cat, Description: newProd.Description, Images: images, SerialNumber: makeTimestamp(),ShopAccount:  *user}

	return p.ProductRepository.Create(&prod)
}

func (p *productUseCase) Delete(ctx echo.Context, deletedProduct dto.DeleteProduct) error {
	serial, err := strconv.ParseUint(deletedProduct.SerialNumber, 10, 64)


	if err != nil{
		return err
	}
	deleted, err := p.ProductRepository.GetBySerialAndUserId(serial, deletedProduct.UserId)
	if err!=nil{
		return err
	}
	folderName := strconv.FormatUint(uint64(deleted.Model.ID), 10)
	path1 := "./src/assets/" + folderName
	os.RemoveAll(path1)
	oldDir, _ := os.Getwd()

	os.Chdir(path1)

	oldImagesLikePath := folderName + "%"
	images,_ := p.ImageRepository.GetyByPath(oldImagesLikePath)

	for _,image := range images{
		//os.Remove(strings.Split(image.Path, "/")[1])
		p.ImageRepository.Delete(image.Model.ID)

	}

	os.Chdir(oldDir)

	return p.ProductRepository.Delete(deleted.Model.ID)
}

func makeTimestamp() uint64 {
	return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}
func NewProductUseCase(p domain.ProductRepository, c domain.CategoryRepository, img domain.ImageRepository, s domain.ShopAccountRepository) domain.ProductUsecase {
	return &productUseCase{s,p,c, img}
}
