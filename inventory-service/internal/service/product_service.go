package service

import (
	"context"
	"fmt"
	"inventory-service/internal/handler"
	"inventory-service/internal/repository"

	pb "github.com/suyundykovv/protos/gen/go/inventory/v1"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error)
	GetProduct(ctx context.Context, id string) (*pb.Product, error)
	UpdateProduct(ctx context.Context, id string) (*pb.Product, error)
}
type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *productService {
	return &productService{repo: repo}
}
func (s *productService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	product := handler.ProtoToDomainProduct(req)
	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}
	return handler.DomainToProtoProduct(product), nil
}
func (s *productService) GetProduct(ctx context.Context, id string) (*pb.Product, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, fmt.Errorf("product not found")
	}
	return handler.DomainToProtoProduct(product), nil
}
func (s *productService) UpdateProduct(ctx context.Context, id string) (*pb.Product, error) {
	err := s.repo.Update(ctx, id)
	if err != nil {
		if err.Error() == fmt.Sprintf("user with ID %s not found", id) {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to took data")
	}
	if product == nil {
		return nil, fmt.Errorf("product not found")
	}
	return handler.DomainToProtoProduct(product), nil
}
