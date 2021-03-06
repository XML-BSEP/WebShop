package mapper

import (
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"web-shop/domain"
	"web-shop/infrastructure/dto"
)

func NewProductDtoToProduct (dto dto.ProductDTO) domain.Product {

	policy := bluemonday.UGCPolicy();
	dto.Name =  strings.TrimSpace(policy.Sanitize(dto.Name))

	return domain.Product{Model: gorm.Model{ID: dto.ID}, Name:dto.Name, Price : dto.Price}
}

func NewProductToProductDto (p domain.Product) dto.ProductDTO {

	return dto.ProductDTO{ID:p.ID, Name : p.Name, Price: p.Price}
}

func NewProductToProductViewDTO (p domain.Product) dto.ProductViewDTO {
	var imt_path string
 	if os.Getenv("DOCKER_ENV") == "" {
		imt_path = "https://localhost:443/static/"
	}  else {
		imt_path = "http://localhost:8099/static/"
	}
	files, _ := ioutil.ReadDir("./src/assets/"+strconv.FormatUint(uint64(p.Model.ID), 10))
	var images []string
	for _, file := range files{
		im := imt_path+strconv.FormatUint(uint64(p.Model.ID), 10) +"/"+file.Name()
		images=append(images,im)
	}

	return dto.ProductViewDTO{
		UserId: p.ShopAccountID,
		ProductId: p.ID,
		Category: p.Category.Name,
		Name: p.Name,
		Available: p.Available,
		Description: p.Description,
		Image: images,
		Price: p.Price,
		SerialNumber:  p.SerialNumber,
	}
}

func mapCurrency(currency domain.Currency) string {
	if currency == 0 {
		return "USD"
	}
	if currency == 1 {
		return "EUR"
	}
	if currency == 2 {
		return "RSD"
	}

	return ""
}

