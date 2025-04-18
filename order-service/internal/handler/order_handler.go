package handler

import (
	"order-service/internal/domain"
	"time"

	pb "github.com/suyundykovv/protos/gen/go/order/v1"
)

// Преобразование из protobuf в domain
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
