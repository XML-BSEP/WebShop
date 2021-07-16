package datastore

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"web-shop/domain"
)

type orderRepository struct {
	collection *mongo.Collection
	db *mongo.Client
}


func NewOrderRepository(db *mongo.Client) domain.OrderRepository {
	return &orderRepository{
		db: db,
		collection : db.Database("order_db").Collection("orders"),
	}
}

func (o *orderRepository) PlaceOrder(ctx context.Context, order *domain.Order) error {
	order.ID = uuid.NewString()
	_, err := o.collection.InsertOne(ctx, *order)
	if err != nil {
		return err
	}
	return nil
}

