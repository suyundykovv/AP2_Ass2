package service

import (
	"context"
	"fmt"
	"order-service/internal/domain"
	"order-service/internal/nats"
	"order-service/internal/repository"
	"time"

	eventspb "github.com/suyundykovv/protos/gen/go/events/v1"
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
	repo      repository.OrderRepository
	publisher *nats.Publisher // Add NATS publisher
}

func NewOrderService(repo repository.OrderRepository, publisher *nats.Publisher) OrderService {
	return &orderService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	order := protoToDomainOrder(req)
	if err := s.repo.Create(ctx, order); err != nil {
		return nil, err
	}

	// Publish order created event
	event := &eventspb.OrderEvent{
		EventType: "created",
		Id:        order.ID,
		UserId:    order.UserID,
		Items:     order.Items,
		Total:     order.Total,
		Status:    string(order.Status),
		CreatedAt: time.Now().Unix(),
	}

	if err := s.publisher.PublishOrderCreated(event); err != nil {
		// Log error but don't fail the operation
		fmt.Printf("Failed to publish order created event: %v\n", err)
	}

	return domainToProtoOrder(order), nil
}

func (s *orderService) UpdateOrderStatus(ctx context.Context, id string, newStatus string) (*pb.Order, error) {
	// Get current order first
	currentOrder, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, statusError(codes.Internal, "failed to fetch order")
	}
	if currentOrder == nil {
		return nil, statusError(codes.NotFound, fmt.Sprintf("order with ID %s not found", id))
	}

	// Update status
	err = s.repo.UpdateStatus(ctx, id, newStatus)
	if err != nil {
		return nil, statusError(codes.Internal, "failed to update order status")
	}

	// Publish order updated event
	event := &eventspb.OrderEvent{
		EventType: "updated",
		Id:        id,
		UserId:    currentOrder.UserID,
		Items:     currentOrder.Items,
		Total:     currentOrder.Total,
		Status:    newStatus,
		CreatedAt: time.Now().Unix(),
		// Include any additional fields that changed
	}

	if err := s.publisher.PublishOrderUpdated(event); err != nil {
		fmt.Printf("Failed to publish order updated event: %v\n", err)
	}

	// Return updated order
	updatedOrder, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, statusError(codes.Internal, "failed to fetch updated order")
	}
	return domainToProtoOrder(updatedOrder), nil
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
