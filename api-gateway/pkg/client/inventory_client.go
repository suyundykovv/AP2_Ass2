package client

import (
	"context"
	"log"
	"time"

	pb "github.com/suyundykovv/protos/gen/go/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InventoryClient struct {
	client pb.InventoryServiceClient
}

func NewInventoryClient(address string) *InventoryClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	return &InventoryClient{
		client: pb.NewInventoryServiceClient(conn),
	}
}

func (oc *InventoryClient) CreateProduct(request *pb.CreateProductRequest) (*pb.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order, err := oc.client.CreateProduct(ctx, request)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (oc *InventoryClient) GetProduct(request *pb.GetProductRequest) (*pb.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order, err := oc.client.GetProduct(ctx, request)
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (oc *InventoryClient) UpdateProduct(request *pb.UpdateProductRequest) (*pb.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order, err := oc.client.UpdateProduct(ctx, request)
	if err != nil {
		return nil, err
	}
	return order, nil
}
