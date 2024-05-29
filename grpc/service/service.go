package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb2 "pos-go-cleanarch-listorders/grpc/pb"
	"pos-go-cleanarch-listorders/internal/model"
	"pos-go-cleanarch-listorders/internal/repository"
)

type CustomOrderServiceServer struct {
	pb2.UnimplementedOrderServiceServer
	orderRepo *repository.OrderRepository
}

func NewOrderServiceServer(orderRepo *repository.OrderRepository) *CustomOrderServiceServer {
	return &CustomOrderServiceServer{orderRepo: orderRepo}
}

func (s *CustomOrderServiceServer) ListOrders(ctx context.Context, in *pb2.Empty) (*pb2.OrderListResponse, error) {
	orders, err := s.orderRepo.ListOrders()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list orders: %v", err)
	}

	var orderList []*pb2.Order
	for _, order := range orders {
		orderList = append(orderList, &pb2.Order{
			Id:    order.ID(),
			Price: float32(order.Price()),
			Tax:   float32(order.Tax()),
		})
	}

	return &pb2.OrderListResponse{Orders: orderList}, nil
}

func (s *CustomOrderServiceServer) CreateOrder(ctx context.Context, in *pb2.Order) (*pb2.OrderResponse, error) {
	price := in.GetPrice()
	tax := in.GetTax()

	order := model.Order{
		PriceField: float64(price),
		TaxField:   float64(tax),
	}

	id, err := s.orderRepo.CreateOrder(order.Price(), order.Tax())
	if err != nil {
		return nil, err
	}

	return &pb2.OrderResponse{
		Id:      int32(id),
		Message: "Order created successfully",
	}, nil
}
