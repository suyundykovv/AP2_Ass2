package handler

import (
	"context"
	"errors"
	"order-service/internal/service"
	"order-service/pkg/models"
	pb "order-service/proto/order"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	service service.OrderService
}

func NewOrderServer(svc service.OrderService) *OrderServer {
	return &OrderServer{service: svc}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// Convert items from repeated string to []string
	items := make([]string, len(req.Items))
	copy(items, req.Items)

	order := &models.Order{
		UserID: int(req.UserId),
		Items:  items,
		Status: req.Status,
		Total:  req.Total,
	}

	createdOrder, err := s.service.CreateOrder(ctx, order)
	if err != nil {
		return convertErrorToStatus4(err)
	}

	return &pb.CreateOrderResponse{
		OrderId: createdOrder.ID,
		UserId:  int32(createdOrder.UserID),
		Items:   createdOrder.Items,
		Status:  createdOrder.Status,
		Total:   createdOrder.Total,
	}, nil
}

func (s *OrderServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	if req.OrderId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid order ID")
	}

	order, err := s.service.GetOrder(ctx, req.OrderId)
	if err != nil {
		return convertErrorToStatus(err)
	}

	return &pb.GetOrderResponse{
		OrderId: order.ID,
		UserId:  int32(order.UserID),
		Items:   order.Items,
		Status:  order.Status,
		Total:   order.Total,
	}, nil
}

func (s *OrderServer) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	order := &models.Order{
		ID:     req.OrderId,
		UserID: int(req.UserId),
		Items:  req.Items,
		Status: req.Status,
		Total:  req.Total,
	}

	updatedOrder, err := s.service.UpdateOrder(ctx, order)
	if err != nil {
		return convertErrorToStatus3(err)
	}

	return &pb.UpdateOrderResponse{
		OrderId: updatedOrder.ID,
		UserId:  int32(updatedOrder.UserID),
		Items:   updatedOrder.Items,
		Status:  updatedOrder.Status,
		Total:   updatedOrder.Total,
	}, nil
}

func (s *OrderServer) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	if req.OrderId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid order ID")
	}

	err := s.service.DeleteOrder(ctx, req.OrderId)
	if err != nil {
		return convertErrorToStatus2(err)
	}

	return &pb.DeleteOrderResponse{
		Status: "successfully deleted",
	}, nil
}

func convertErrorToStatus(err error) (*pb.GetOrderResponse, error) {
	switch {
	case errors.Is(err, errors.New("invalid user ID")),
		errors.Is(err, errors.New("order must contain items")),
		errors.Is(err, errors.New("total must be positive")),
		errors.Is(err, errors.New("order total exceeds maximum limit")),
		errors.Is(err, errors.New("order cannot contain more than 10 items")):
		return nil, status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, errors.New("order not found")),
		errors.Is(err, errors.New("cannot retrieve cancelled orders")),
		errors.Is(err, errors.New("cannot delete orders in processing status")):
		return nil, status.Error(codes.NotFound, err.Error())
	case errors.Is(err, errors.New("cannot modify shipped orders")):
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
}
func convertErrorToStatus2(err error) (*pb.DeleteOrderResponse, error) {
	switch {
	case errors.Is(err, errors.New("invalid user ID")),
		errors.Is(err, errors.New("order must contain items")),
		errors.Is(err, errors.New("total must be positive")),
		errors.Is(err, errors.New("order total exceeds maximum limit")),
		errors.Is(err, errors.New("order cannot contain more than 10 items")):
		return nil, status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, errors.New("order not found")),
		errors.Is(err, errors.New("cannot retrieve cancelled orders")),
		errors.Is(err, errors.New("cannot delete orders in processing status")):
		return nil, status.Error(codes.NotFound, err.Error())
	case errors.Is(err, errors.New("cannot modify shipped orders")):
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
}
func convertErrorToStatus3(err error) (*pb.UpdateOrderResponse, error) {
	switch {
	case errors.Is(err, errors.New("invalid user ID")),
		errors.Is(err, errors.New("order must contain items")),
		errors.Is(err, errors.New("total must be positive")),
		errors.Is(err, errors.New("order total exceeds maximum limit")),
		errors.Is(err, errors.New("order cannot contain more than 10 items")):
		return nil, status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, errors.New("order not found")),
		errors.Is(err, errors.New("cannot retrieve cancelled orders")),
		errors.Is(err, errors.New("cannot delete orders in processing status")):
		return nil, status.Error(codes.NotFound, err.Error())
	case errors.Is(err, errors.New("cannot modify shipped orders")):
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
}
func convertErrorToStatus4(err error) (*pb.CreateOrderResponse, error) {
	switch {
	case errors.Is(err, errors.New("invalid user ID")),
		errors.Is(err, errors.New("order must contain items")),
		errors.Is(err, errors.New("total must be positive")),
		errors.Is(err, errors.New("order total exceeds maximum limit")),
		errors.Is(err, errors.New("order cannot contain more than 10 items")):
		return nil, status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, errors.New("order not found")),
		errors.Is(err, errors.New("cannot retrieve cancelled orders")),
		errors.Is(err, errors.New("cannot delete orders in processing status")):
		return nil, status.Error(codes.NotFound, err.Error())
	case errors.Is(err, errors.New("cannot modify shipped orders")):
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
}
