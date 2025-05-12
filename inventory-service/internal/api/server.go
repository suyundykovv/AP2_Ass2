package api

import (
	"context"
	"inventory-service/internal/service"
	"log"
	"time"

	pb "github.com/suyundykovv/protos/gen/go/inventory/v1"
	"google.golang.org/grpc"
)

type InventoryServer struct {
	pb.UnimplementedInventoryServiceServer
	service service.ProductService
}

func NewInventoryServer(svc service.ProductService) *InventoryServer {
	return &InventoryServer{service: svc}
}
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	resp, err = handler(ctx, req)
	log.Printf("Method: %s, Duration: %v, Error: %v", info.FullMethod, time.Since(start), err)
	return resp, err
}

func (s *InventoryServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	order, err := s.service.CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *InventoryServer) GetProductByID(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	order, err := s.service.GetProduct(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (s *InventoryServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
	return s.service.UpdateProduct(ctx, req.Id)
}
