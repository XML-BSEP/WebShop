package usecase

import (
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
	ProductRepository domain.ProductRepository
	CategoryRepository domain.CategoryRepository
	ImageRepository domain.ImageRepository
}

func (p *productUseCase) FilterByCategory(category string, priceRangeStart uint, priceRangeEnd uint, limit int, offset int, order string) ([]*domain.Product, error) {
	return p.ProductRepository.FilterByCategory(category, priceRangeStart, priceRangeEnd, limit, offset, order)
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

func (p *productUseCase) Update(ctx echo.Context, pic *domain.Product) (*domain.Product, error) {
	return p.ProductRepository.Update(pic)
}

func (p *productUseCase) Create(ctx echo.Context, newProd *dto.NewProduct) (*domain.Product, error) {
	var cat *domain.Category
	cat, _ = p.CategoryRepository.GetByName(newProd.Category)

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

	//for i, _ := range images{
	//	p.ImageRepository.Create(images[i])
	//
	//}


	os.Chdir(path2)
	curr,_ :=strconv.Atoi(newProd.Currency)
	price,_ :=strconv.Atoi(newProd.Price)
	av,_ :=strconv.Atoi(newProd.Available)

	prod:=domain.Product{Currency: domain.Currency(curr), Available: uint(av), Price: uint64(uint(price)), Name: newProd.Name, Category: *cat, Description: newProd.Description, Images: images}

	//Name  string `json:"name"`
	//Price uint64 `json:"price"`
	//Currency Currency `json:"currency"`
	//Images []Image
	return p.ProductRepository.Create(&prod)
}

func (p *productUseCase) Delete(ctx echo.Context, id uint) error {
	return p.ProductRepository.Delete(id)
}


func NewProductUseCase(p domain.ProductRepository, c domain.CategoryRepository, img domain.ImageRepository) domain.ProductUsecase {
	return &productUseCase{p,c, img}
}
