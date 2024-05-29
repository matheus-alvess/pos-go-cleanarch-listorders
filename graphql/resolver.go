package graphql

import (
	"context"
	"pos-go-cleanarch-listorders/internal/model"
	"pos-go-cleanarch-listorders/internal/repository"
)

type Resolver struct {
	OrderRepo *repository.OrderRepository
}

func NewResolver(orderRepo *repository.OrderRepository) *Resolver {
	return &Resolver{OrderRepo: orderRepo}
}

func (r *Resolver) ListOrders() ([]*model.Order, error) {
	orders, err := r.OrderRepo.ListOrders()
	if err != nil {
		return nil, err
	}

	var result []*model.Order
	for _, order := range orders {
		result = append(result, &order)
	}
	return result, nil
}

func (r *Resolver) CreateOrder(ctx context.Context, args struct {
	Price float64
	Tax   float64
}) (*model.Order, error) {
	id, err := r.OrderRepo.CreateOrder(args.Price, args.Tax)
	if err != nil {
		return nil, err
	}
	return &model.Order{
		IDField:    int32(id),
		PriceField: args.Price,
		TaxField:   args.Tax,
	}, nil
}
