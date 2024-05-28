package grpc

import (
	"context"
	"pos-go-cleanarch-listorders/internal/repository"
)

type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer
	orderRepo *repository.OrderRepository
}

func NewOrderServiceServer(orderRepo *repository.OrderRepository) *OrderServiceServer {
	return &OrderServiceServer{orderRepo: orderRepo}
}

func (s *OrderServiceServer) ListOrders(ctx context.Context, in *pb.Empty) (*pb.OrderListResponse, error) {
	orders, err := s.orderRepo.ListOrders()
	if err != nil {
		return nil, err
	}

	var orderList []*pb.Order
	for _, order := range orders {
		orderList = append(orderList, &pb.Order{
			Id:    order.ID,
			Price: order.Price,
			Tax:   order.Tax,
		})
	}

	return &pb.OrderListResponse{Orders: orderList}, nil
}
