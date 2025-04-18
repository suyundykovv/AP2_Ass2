package grpc

import (
	"context"
	"log"
	"time"

	"order-service/internal/service"

	pb "github.com/suyundykovv/protos/gen/go/order/v1"
	"google.golang.org/grpc"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	service service.OrderService
}

func NewOrderServer(svc service.OrderService) *OrderServer {
	return &OrderServer{service: svc}
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	resp, err = handler(ctx, req)
	log.Printf("Method: %s, Duration: %v, Error: %v", info.FullMethod, time.Since(start), err)
	return resp, err
}

// CreateOrder просто вызывает сервис, и все
func (s *OrderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	order, err := s.service.CreateOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// То же самое для других методов
func (s *OrderServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.Order, error) {
	order, err := s.service.GetOrder(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (s *OrderServer) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.Order, error) {
	return s.service.UpdateOrderStatus(ctx, req.GetId(), req.GetStatus())
}
