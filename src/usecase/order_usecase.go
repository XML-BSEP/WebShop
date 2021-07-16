package usecase

import (
	"context"
	"time"
	"web-shop/domain"
)

type orderUseCase struct {
	OrderRepository domain.OrderRepository
	ShoppingCartItemUsecase domain.ShoppingCartItemUsecase
}

func (o *orderUseCase) PlaceOrder(ctx context.Context, order *domain.Order) error {
	items, err := o.ShoppingCartItemUsecase.GetAllUsersShoppingCartItems(ctx, order.UserId)
	if err !=nil{
		return err
	}
	var itemIds []uint
	var totalPrice float64
	for _, it:=range items{
		itemIds = append(itemIds, it.ID)
		totalPrice += it.Product.Price
		err1 := o.ShoppingCartItemUsecase.Delete(ctx, it.ID)
		if err1 != nil {
			return err1
		}
	}
	newOrder := domain.Order{UserId: order.UserId, Zip: order.Zip, City: order.City, State: order.State, Address: order.Address, Timestamp: time.Now(),ShoppingCartItems: itemIds, TotalPrice: totalPrice}
	err1 := o.OrderRepository.PlaceOrder(ctx, &newOrder)
	if err1!=nil{
		return err1
	}
	return nil
}

func NewOrderUsecase(r domain.OrderRepository, 	s domain.ShoppingCartItemUsecase) domain.OrderUsecase {
	return &orderUseCase{r,s}
}