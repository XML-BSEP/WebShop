package usecase

import (
	"context"
	"github.com/labstack/echo"
	"web-shop/domain"
)

type ShoppingCartItemUsecase struct {
	ShoppingCartItemRepository domain.ShoppingCartItemRepository
	ShoppingCartReporsitory domain.ShoppingCartRepository
	ProductUseCase domain.ProductUsecase
}

func (s1 ShoppingCartItemUsecase) GetAllUsersShoppingCartItems(ctx context.Context, userId uint) ([]*domain.ShoppingCartItem, error) {
	items, err := s1.ShoppingCartItemRepository.GetAllUsersShoppingCartItems(ctx, userId)
	if err!=nil{
		return nil, err
	}

	for i,it :=range items{
		p,err1 := s1.ProductUseCase.GetProductDetails(ctx, it.ProductID)
		if err1!=nil{
			return nil, err1
		}
		items[i].Product = *p

	}
	return items, nil
}

func (s1 ShoppingCartItemUsecase) Fetch(ctx echo.Context) ([]*domain.ShoppingCartItem, error) {
	return s1.ShoppingCartItemRepository.Fetch()
}

func (s1 ShoppingCartItemUsecase) GetByID(ctx echo.Context, id uint) (*domain.ShoppingCartItem, error) {
	return s1.ShoppingCartItemRepository.GetByID(id)
}

func (s1 ShoppingCartItemUsecase) Update(ctx echo.Context, s *domain.ShoppingCartItem) (*domain.ShoppingCartItem, error) {
	return s1.ShoppingCartItemRepository.Update(s)
}

func (s1 ShoppingCartItemUsecase) Create(ctx echo.Context, s *domain.ShoppingCartItem) (*domain.ShoppingCartItem, error) {
	return s1.ShoppingCartItemRepository.Create(s)
}

func (s1 ShoppingCartItemUsecase) Delete(ctx context.Context, id uint) error {
	return s1.ShoppingCartItemRepository.Delete(id)
}

func (s1 ShoppingCartItemUsecase) AddToCart(ctx context.Context, productId uint, userId uint) error {
	//return s1.ShoppingCartItemRepository.AddToCart(ctx, productId, userId)

	newShoppingCartItem := domain.ShoppingCartItem{ProductID: productId, RegisteredShopUserID: userId}

	_,err1 := s1.ShoppingCartItemRepository.Create(&newShoppingCartItem)
	if err1!=nil{
		return err1
	}
	return nil
}

func NewShoppingCartItemUsecase(r domain.ShoppingCartItemRepository, r2 domain.ShoppingCartRepository, r3 domain.ProductUsecase) domain.ShoppingCartItemUsecase {
	return &ShoppingCartItemUsecase{r, r2, r3}
}

