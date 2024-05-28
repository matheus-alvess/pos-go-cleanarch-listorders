package service

import (
	"encoding/json"
	"net/http"
	"pos-go-cleanarch-listorders/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := s.repo.ListOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}
