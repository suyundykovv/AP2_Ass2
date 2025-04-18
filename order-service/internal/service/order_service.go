package service

import (
	"context"
	"fmt"
	"order-service/internal/domain"
	"order-service/internal/repository"
	"time"

	pb "github.com/suyundykovv/protos/gen/go/order/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error)
	GetOrder(ctx context.Context, id string) (*pb.Order, error)
	UpdateOrderStatus(ctx context.Context, id string, status string) (*pb.Order, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	order := protoToDomainOrder(req)
	if err := s.repo.Create(ctx, order); err != nil {
		return nil, err
	}
	return domainToProtoOrder(order), nil
}

func (s *orderService) GetOrder(ctx context.Context, id string) (*pb.Order, error) {
	order, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, fmt.Errorf("order with ID %s not found", id)
	}
	return domainToProtoOrder(order), nil
}
func (s *orderService) UpdateOrderStatus(ctx context.Context, id string, status string) (*pb.Order, error) {
	err := s.repo.UpdateStatus(ctx, id, status)
	if err != nil {
		if err.Error() == fmt.Sprintf("order with ID %s not found", id) {
			return nil, statusError(codes.NotFound, err.Error())
		}
		return nil, statusError(codes.Internal, "failed to update order status")
	}

	order, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, statusError(codes.Internal, "failed to fetch updated order")
	}
	if order == nil {
		return nil, statusError(codes.NotFound, fmt.Sprintf("order with ID %s not found", id))
	}

	return domainToProtoOrder(order), nil
}

func statusError(code codes.Code, msg string) error {
	return status.Errorf(code, msg)
}

func protoToDomainOrder(req *pb.CreateOrderRequest) *domain.Order {
	return &domain.Order{
		UserID:    req.GetUserId(),
		Items:     req.GetItems(), // В нашем случае это просто список строк (ID товаров)
		Status:    domain.OrderStatusPending,
		Total:     req.GetTotal(),
		CreatedAt: time.Now().Unix(),
	}
}

// Обратное преобразование из domain в protobuf
func domainToProtoOrder(order *domain.Order) *pb.Order {
	return &pb.Order{
		Id:        order.ID,
		UserId:    order.UserID,
		Items:     order.Items,
		Total:     order.Total,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
	}
}
