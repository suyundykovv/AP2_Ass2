package client

import (
	"context"
	"log"
	"time"

	pb "github.com/suyundykovv/protos/gen/go/order/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClient struct {
	client pb.OrderServiceClient
}

func NewOrderClient(address string) *OrderClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	return &OrderClient{
		client: pb.NewOrderServiceClient(conn),
	}
}

func (oc *OrderClient) CreateOrder(request *pb.CreateOrderRequest) (*pb.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order, err := oc.client.CreateOrder(ctx, request)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (oc *OrderClient) GetOrder(request *pb.GetOrderRequest) (*pb.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order, err := oc.client.GetOrder(ctx, request)
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (oc *OrderClient) UpdateOrderStatus(request *pb.UpdateOrderStatusRequest) (*pb.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order, err := oc.client.UpdateOrderStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	return order, nil
}
