package order

import (
	"context"
	"github.com/NuthChanReaksa/ap-with-golang-001/types"
	"github.com/go-kivik/kivik"
)

type Store struct {
	db *kivik.DB
}

func NewStore(db *kivik.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (string, error) {
	doc := map[string]interface{}{
		"userId":  order.UserID,
		"total":   order.Total,
		"status":  order.Status,
		"address": order.Address,
		"items":   order.Items,
	}

	// Create a context
	ctx := context.Background()

	// Use Put with context
	id, err := s.db.Put(ctx, order.ID, doc)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	doc := map[string]interface{}{
		"orderId":   orderItem.OrderID,
		"productId": orderItem.ProductID,
		"quantity":  orderItem.Quantity,
		"price":     orderItem.Price,
	}

	ctx := context.Background()

	_, err := s.db.Put(ctx, orderItem.ID, doc)
	return err
}
