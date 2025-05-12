package service

import (
	"context"
	"statistic-service/internal/repository"
	"time"

	pb "github.com/suyundykovv/protos/gen/go/events/v1"
	statisticspb "github.com/suyundykovv/protos/gen/go/statistics/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StatisticService interface {
	ProcessOrderCreated(event *pb.OrderEvent) error
	ProcessOrderUpdated(event *pb.OrderEvent) error
	ProcessOrderDeleted(event *pb.OrderEvent) error
	ProcessInventoryEvent(event *pb.InventoryEvent) error
	GetUserOrdersStatistics(ctx context.Context, userID string, timePeriod string) (*statisticspb.GetUserOrdersStatisticsResponse, error)
	GetUserStatistics(ctx context.Context, timePeriod string) (*statisticspb.GetUserStatisticsResponse, error)
}

type statisticService struct {
	repo repository.StatisticRepository
}

func NewStatisticService(repo repository.StatisticRepository) StatisticService {
	return &statisticService{repo: repo}
}

func (s *statisticService) ProcessOrderCreated(event *pb.OrderEvent) error {
	itemsCount := 0
	total := 0.0
	for _, item := range event.Items {
		itemsCount += int(item.Quantity)
		total += item.Price * float64(item.Quantity)
	}

	timestamp, err := time.Parse(time.RFC3339, event.Timestamp)
	if err != nil {
		return err
	}

	orderEvent := &repository.OrderEvent{
		EventType:  "created",
		OrderID:    event.OrderId,
		UserID:     event.UserId,
		Timestamp:  timestamp,
		ItemsCount: itemsCount,
		Total:      total,
	}

	return s.repo.SaveOrderEvent(context.Background(), orderEvent)
}

func (s *statisticService) ProcessOrderUpdated(event *pb.OrderEvent) error {
	// Similar to ProcessOrderCreated but for updates
	return nil
}

func (s *statisticService) ProcessOrderDeleted(event *pb.OrderEvent) error {
	// Handle order deletion statistics
	return nil
}

func (s *statisticService) ProcessInventoryEvent(event *pb.InventoryEvent) error {
	timestamp, err := time.Parse(time.RFC3339, event.Timestamp)
	if err != nil {
		return err
	}

	inventoryEvent := &repository.InventoryEvent{
		EventType:   event.EventType,
		ProductID:   event.ProductId,
		CategoryID:  event.CategoryId,
		Timestamp:   timestamp,
		StockChange: int(event.StockChange),
	}

	return s.repo.SaveInventoryEvent(context.Background(), inventoryEvent)
}

func (s *statisticService) GetUserOrdersStatistics(ctx context.Context, userID string, timePeriod string) (*statisticspb.GetUserOrdersStatisticsResponse, error) {
	stats, err := s.repo.GetUserOrderStatistics(ctx, userID, timePeriod)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user order statistics: %v", err)
	}

	return &statisticspb.GetUserOrdersStatisticsResponse{
		TotalOrders:       int32(stats.TotalOrders),
		AverageOrderValue: stats.AverageOrderValue,
		OrdersByHour:      convertMapToProto(stats.OrdersByHour),
		OrdersByDay:       convertMapToProto(stats.OrdersByDay),
		FavoriteCategory:  stats.FavoriteCategory,
	}, nil
}

func (s *statisticService) GetUserStatistics(ctx context.Context, timePeriod string) (*statisticspb.GetUserStatisticsResponse, error) {
	stats, err := s.repo.GetGeneralUserStatistics(ctx, timePeriod)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user statistics: %v", err)
	}

	return &statisticspb.GetUserStatisticsResponse{
		TotalUsers:           int32(stats.TotalUsers),
		ActiveUsers:          int32(stats.ActiveUsers),
		AverageOrdersPerUser: stats.AverageOrdersPerUser,
		MostActiveTime:       stats.MostActiveTime,
	}, nil
}

func convertMapToProto(m map[int]int) map[int32]int32 {
	result := make(map[int32]int32)
	for k, v := range m {
		result[int32(k)] = int32(v)
	}
	return result
}
