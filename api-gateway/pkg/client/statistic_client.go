package client

import (
	"context"
	"log"
	"time"

	pb "github.com/suyundykovv/protos/gen/go/statistics/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type StatisticsClient struct {
	client pb.StatisticsServiceClient
}

func NewStatisticsClient(address string) *StatisticsClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to Statistics service: %v", err)
	}
	return &StatisticsClient{
		client: pb.NewStatisticsServiceClient(conn),
	}
}

func (oc *StatisticsClient) GetUserOrdersStatistics(request *pb.GetUserOrdersStatisticsRequest) (*pb.GetUserOrdersStatisticsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	statistics, err := oc.client.GetUserOrdersStatistics(ctx, request)
	if err != nil {
		return nil, err
	}
	return statistics, nil
}
func (oc *StatisticsClient) GetUserStatistics(request *pb.GetUserStatisticsRequest) (*pb.GetUserStatisticsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	statistics, err := oc.client.GetUserStatistics(ctx, request)
	if err != nil {
		return nil, err
	}
	return statistics, nil
}
