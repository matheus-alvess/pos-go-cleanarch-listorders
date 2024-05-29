// graphql/resolver.go

package graphql

import (
	"pos-go-cleanarch-listorders/internal/model"
	"pos-go-cleanarch-listorders/internal/repository"
)

type Resolver struct {
	OrderRepo *repository.OrderRepository
}

func NewResolver(orderRepo *repository.OrderRepository) *Resolver {
	return &Resolver{OrderRepo: orderRepo}
}

// ListOrders retorna a lista de pedidos.
func (r *Resolver) ListOrders() ([]*model.Order, error) {
	// Lógica para recuperar a lista de pedidos do repositório
	orders, err := r.OrderRepo.ListOrders()
	if err != nil {
		return nil, err
	}

	// Convertendo []model.Order para []*model.Order
	var result []*model.Order
	for _, order := range orders {
		result = append(result, &order)
	}
	return result, nil
}
