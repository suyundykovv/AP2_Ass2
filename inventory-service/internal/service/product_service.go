package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"inventory-service/internal/domain"
	"inventory-service/internal/handler"
	"inventory-service/internal/nats"
	"inventory-service/internal/repository"

	"github.com/patrickmn/go-cache"

	eventspb "github.com/suyundykovv/protos/gen/go/events/v1"
	pb "github.com/suyundykovv/protos/gen/go/inventory/v1"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error)
	GetProduct(ctx context.Context, id string) (*pb.Product, error)
	GetAllProducts(ctx context.Context) ([]*pb.Product, error)
	UpdateProduct(ctx context.Context, id string) (*pb.Product, error)
	WarmCache(ctx context.Context) error
	StartCacheRefresh(ctx context.Context)
}

type productService struct {
	repo      repository.ProductRepository
	publisher *nats.Publisher
	cache     *cache.Cache
}

func NewProductService(repo repository.ProductRepository, publisher *nats.Publisher) ProductService {
	c := cache.New(12*time.Hour, 1*time.Hour)
	service := &productService{
		repo:      repo,
		publisher: publisher,
		cache:     c,
	}
	return service
}

func (s *productService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	product := handler.ProtoToDomainProduct(req)
	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	// Invalidate relevant cache entries
	s.cache.Delete("all_products")

	// Publish product created event
	event := &eventspb.InventoryEvent{
		EventType:  "created",
		ProductId:  product.ID,
		Name:       product.Name,
		CategoryId: product.CategoryID,
		Price:      product.Price,
	}

	if err := s.publisher.PublishInventoryCreated(event); err != nil {
		log.Printf("Failed to publish inventory created event: %v\n", err)
	}

	return handler.DomainToProtoProduct(product), nil
}

func (s *productService) GetProduct(ctx context.Context, id string) (*pb.Product, error) {
	cacheKey := "product_" + id
	if item, found := s.cache.Get(cacheKey); found {
		return item.(*pb.Product), nil
	}

	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, fmt.Errorf("product not found")
	}

	protoProduct := handler.DomainToProtoProduct(product)
	s.cache.Set(cacheKey, protoProduct, cache.DefaultExpiration)
	return protoProduct, nil
}

func (s *productService) UpdateProduct(ctx context.Context, id string) (*pb.Product, error) {
	err := s.repo.Update(ctx, id)
	if err != nil {
		if err.Error() == fmt.Sprintf("user with ID %s not found", id) {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}

	// Invalidate cache
	s.cache.Delete("product_" + id)
	s.cache.Delete("all_products")

	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated product: %v", err)
	}
	if product == nil {
		return nil, fmt.Errorf("product not found after update")
	}

	protoProduct := handler.DomainToProtoProduct(product)
	s.cache.Set("product_"+id, protoProduct, cache.DefaultExpiration)
	return protoProduct, nil
}
func (s *productService) GetAllProducts(ctx context.Context) ([]*pb.Product, error) {
	// Try to get from cache first
	if items, found := s.cache.Get("all_products"); found {
		if products, ok := items.([]*domain.Product); ok {
			// Convert domain products to protobuf products
			var protoProducts []*pb.Product
			for _, product := range products {
				protoProducts = append(protoProducts, handler.DomainToProtoProduct(product))
			}
			return protoProducts, nil
		}
		// If type assertion fails, continue to fetch from database
	}

	// Cache miss - get from database
	domainProducts, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get products from repository: %v", err)
	}

	// Convert to protobuf and cache individual products
	var protoProducts []*pb.Product
	for _, product := range domainProducts {
		protoProduct := handler.DomainToProtoProduct(product)
		protoProducts = append(protoProducts, protoProduct)

		// Cache individual product
		cacheKey := "product_" + product.ID
		s.cache.Set(cacheKey, protoProduct, cache.DefaultExpiration)
	}

	// Cache the full list (store domain products for consistency with WarmCache)
	s.cache.Set("all_products", domainProducts, cache.DefaultExpiration)

	return protoProducts, nil
}
func (s *productService) WarmCache(ctx context.Context) error {
	products, err := s.repo.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("failed to get products for cache warming: %v", err)
	}

	for _, product := range products {
		cacheKey := "product_" + product.ID
		s.cache.Set(cacheKey, handler.DomainToProtoProduct(product), cache.DefaultExpiration)
	}

	s.cache.Set("all_products", products, cache.DefaultExpiration)
	return nil
}

func (s *productService) StartCacheRefresh(ctx context.Context) {
	ticker := time.NewTicker(12 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Refreshing cache...")
			if err := s.WarmCache(ctx); err != nil {
				log.Printf("Cache refresh failed: %v", err)
			}
		case <-ctx.Done():
			return
		}
	}
}
