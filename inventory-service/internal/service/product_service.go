package service

import (
	"context"
	"fmt"
	"inventory-service/internal/handler"
	"inventory-service/internal/nats"
	"inventory-service/internal/repository"

	eventspb "github.com/suyundykovv/protos/gen/go/events/v1"
	pb "github.com/suyundykovv/protos/gen/go/inventory/v1"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error)
	GetProduct(ctx context.Context, id string) (*pb.Product, error)
	UpdateProduct(ctx context.Context, id string, req *pb.UpdateProductRequest) (*pb.Product, error)
}

type productService struct {
	repo      repository.ProductRepository
	publisher *nats.Publisher
}

func NewProductService(repo repository.ProductRepository, publisher *nats.Publisher) ProductService {
	return &productService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	product := handler.ProtoToDomainProduct(req)
	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	// Publish product created event
	event := &eventspb.InventoryEvent{
		EventType:  "created",
		ProductId:  product.ID,
		Name:       product.Name,
		CategoryId: product.CategoryID,
		Price:      product.Price,
	}

	if err := s.publisher.PublishInventoryCreated(event); err != nil {
		fmt.Printf("Failed to publish inventory created event: %v\n", err)
	}

	return handler.DomainToProtoProduct(product), nil
}

func (s *productService) UpdateProduct(ctx context.Context, id string, req *pb.UpdateProductRequest) (*pb.Product, error) {
	// Get current product first
	currentProduct, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	}
	if currentProduct == nil {
		return nil, fmt.Errorf("product not found")
	}

	// Update product
	updatedProduct := handler.ProtoToDomainProduct(req)
	if err := s.repo.Update(ctx, id, updatedProduct); err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	// Publish product updated event
	event := &eventspb.InventoryEvent{
		EventType:  "updated",
		ProductId:  id,
		Name:       updatedProduct.Name,
		CategoryId: updatedProduct.CategoryID,
		Price:      updatedProduct.Price,
	}

	if err := s.publisher.PublishInventoryUpdated(event); err != nil {
		fmt.Printf("Failed to publish inventory updated event: %v\n", err)
	}

	// Return updated product
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated product: %w", err)
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
