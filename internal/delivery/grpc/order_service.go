package grpc

import (
	"context"

	"github.com/sorroche-m/desafio-clean-architecture/internal/domain"
	pb "github.com/sorroche-m/desafio-clean-architecture/pkg/proto"
)

type OrderGRPCService struct {
	pb.UnimplementedOrderServiceServer
	orderUseCase domain.OrderUseCase
}

func NewOrderGRPCService(orderUseCase domain.OrderUseCase) *OrderGRPCService {
	return &OrderGRPCService{
		orderUseCase: orderUseCase,
	}
}

func (s *OrderGRPCService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	order, err := s.orderUseCase.CreateOrder(req.CustomerId, req.Amount)
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{
		Id:         order.ID,
		CustomerId: order.CustomerID,
		Amount:     order.Amount,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:  order.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

func (s *OrderGRPCService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.orderUseCase.ListOrders()
	if err != nil {
		return nil, err
	}

	var pbOrders []*pb.OrderResponse
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.OrderResponse{
			Id:         order.ID,
			CustomerId: order.CustomerID,
			Amount:     order.Amount,
			Status:     order.Status,
			CreatedAt:  order.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:  order.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}

	return &pb.ListOrdersResponse{
		Orders: pbOrders,
	}, nil
}

func (s *OrderGRPCService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	order, err := s.orderUseCase.GetOrder(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{
		Id:         order.ID,
		CustomerId: order.CustomerID,
		Amount:     order.Amount,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:  order.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}
