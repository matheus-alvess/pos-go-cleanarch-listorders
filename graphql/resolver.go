package graphql

import (
	"context"
	"pos-go-cleanarch-listorders/internal/repository"
)

type Resolver struct {
	OrderRepo *repository.OrderRepository
}

func NewResolver(orderRepo *repository.OrderRepository) *Resolver {
	return &Resolver{OrderRepo: orderRepo}
}

func (r *Resolver) ListOrders(ctx context.Context) ([]*Order, error) {
	orders, err := r.OrderRepo.ListOrders()
	if err != nil {
		return nil, err
	}

	var result []*Order
	for _, order := range orders {
		result = append(result, &Order{
			ID:    order.ID,
			Price: order.Price,
			Tax:   order.Tax,
		})
	}
	return result, nil
}

type Order struct {
	ID    int     `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}
