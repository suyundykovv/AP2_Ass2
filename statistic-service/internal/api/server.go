package api

import (
	"context"
	"log"
	"time"

	"statistic-service/internal/service"

	pb "github.com/suyundykovv/protos/gen/go/statistics/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StatisticsServer struct {
	pb.UnimplementedStatisticsServiceServer
	service service.StatisticService
}

func NewStatisticsServer(svc service.StatisticService) *StatisticsServer {
	return &StatisticsServer{service: svc}
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	resp, err = handler(ctx, req)
	log.Printf("Method: %s, Duration: %v, Error: %v", info.FullMethod, time.Since(start), err)
	return resp, err
}

func (s *StatisticsServer) GetUserOrdersStatistics(ctx context.Context, req *pb.GetUserOrdersStatisticsRequest) (*pb.GetUserOrdersStatisticsResponse, error) {
	stats, err := s.service.GetUserOrdersStatistics(ctx, req.UserId, req.TimePeriod)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user orders statistics: %v", err)
	}
	return stats, nil
}

func (s *StatisticsServer) GetUserStatistics(ctx context.Context, req *pb.GetUserStatisticsRequest) (*pb.GetUserStatisticsResponse, error) {
	stats, err := s.service.GetUserStatistics(ctx, req.TimePeriod)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user statistics: %v", err)
	}
	return stats, nil
}
