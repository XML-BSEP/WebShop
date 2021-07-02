package domain

import (
	"context"
	"time"
)

type Order struct {
	ID string `bson:"_id,omitempty" json:"id"`
	Timestamp time.Time	`bson:"timestamp" json:"timestamp"`
	ShoppingCartItems    []uint	`bson:"items" json:"items"`
	TotalPrice      float64	`bson:"totalPrice" json:"totalPrice"`
	UserId	uint	`bson:"userId" json:"userId"`
	Address	string `bson:"address" json:"address"`
	City	string `bson:"city" json:"city"`
	Zip	uint `bson:"zip" json:"zip"`
	State	string `bson:"state" json:"state"`
}



type OrderUsecase interface {

	PlaceOrder(ctx context.Context, order *Order) error

}

type OrderRepository interface {
	PlaceOrder(ctx context.Context, order *Order) error
}
